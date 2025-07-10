package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"os"
	"strings"
	"time"
)

// ElementID represents the EBML ID of an element.
// The supported EBML IDs are documented in the Matroska specification:
// https://www.matroska.org/technical/specs/index.html
type ElementID uint64

type elementType int

const (
	_ = iota
	uintegerType
	integerType
	binaryType
	stringType
	utf8Type
	floatType
	dateType
	masterType
)

type Element struct {
	ID          ElementID
	Size        uint64
	Children    []*Element
	Parent      *Element
	Offset      int64
	WriteOffset int64
}

type BlockCRC struct {
	Element     *Element
	StartOffset int64
	EndOffset   int64
	ExpectedCRC uint32
	CRC32       hash.Hash32
	Dirty       bool
}

var blockCRCs []BlockCRC
var rootElement *Element

var progressBytesReadCurrent int64
var progressBytesReadTotal int64
var progressStart time.Time
var progressLastPrint time.Time

func parseElement(reader *bufio.Reader, writer *os.File, level int64, currentOffset int64, limit int64, parent *Element) error {
	var bytesRead int64 = 0

	if parent == nil {
		currentElement := &Element{
			Offset:   0,
			ID:       0,
			Size:     0,
			Children: make([]*Element, 0),
			Parent:   nil,
		}
		parent = currentElement
		rootElement = currentElement
	}

	for limit == 0 || currentOffset+bytesRead < limit {

		// Read ID
		id, idLen, err := readVint(reader, true)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error parsing ID (value):", err)
			break
		}

		idBytes := make([]byte, idLen)
		_, err = io.ReadFull(reader, idBytes)
		if err != nil {
			fmt.Println("Error parsing ID (length):", err)
			break
		}

		bytesRead += int64(idLen)

		// Read size
		size, sizeLen, err := readVint(reader, false)
		if err != nil {
			fmt.Println("Error parsing size (value):", err)
			break
		}

		sizeBytes := make([]byte, sizeLen)
		_, err = io.ReadFull(reader, sizeBytes)
		if err != nil {
			fmt.Println("Error parsing size (length):", err)
			break
		}

		bytesRead += int64(sizeLen)

		// ...

		elementType := getElementType(ElementID(id))
		elementName := getElementName(ElementID(id))

		currentElement := &Element{
			Offset:   currentOffset + bytesRead - int64(sizeLen) - int64(idLen),
			ID:       ElementID(id),
			Size:     size,
			Children: make([]*Element, 0),
			Parent:   parent,
		}

		// ...

		if writer != nil && parent.ID == SegmentElement && parent.Parent.ID == 0 && len(parent.Children) == 0 {
			if currentElement.ID != CRC32Element {
				tmp := &Element{
					ID:          CRC32Element,
					Size:        4,
					Children:    make([]*Element, 0),
					Parent:      parent,
					Offset:      0,
					WriteOffset: 0,
				}
				parent.Children = append(parent.Children, tmp)
				parent.Size = parent.Size + 6

				writer.Seek(parent.Offset+int64(len(ebmlIDToBytes(uint64(parent.ID)))), io.SeekStart)
				sizeParent, _ := encodeVint(parent.Size, 0)
				writer.Write(sizeParent)

				writer.Write(ebmlIDToBytes(uint64(CRC32Element)))
				sizeCRC, _ := encodeVint(4, 0)
				writer.Write(sizeCRC)
				writer.Write([]byte{0x0, 0x0, 0x0, 0x0})

				tmp.WriteOffset = parent.Offset + int64(len(ebmlIDToBytes(uint64(parent.ID)))) + int64(len(sizeParent))

				blockCRCs = append(blockCRCs, BlockCRC{
					Element:     tmp,
					StartOffset: currentElement.Offset,
					EndOffset:   limit,
					ExpectedCRC: 0,
					CRC32:       crc32.NewIEEE(),
					Dirty:       true,
				})
			}
		}

		parent.Children = append(parent.Children, currentElement)

		// ...

		if false {
			fmt.Printf("%s %s(0x%X) size=%d @=%d\n", strings.Repeat("  ", int(level)), elementName, id, size, currentElement.Offset)
		}

		// ...

		for i := range blockCRCs {
			//fmt.Printf("CRC %d cur=%d start=%d end=%d\n", i, currentElement.Offset, blockCRCs[i].StartOffset, blockCRCs[i].EndOffset)
			if currentElement.Offset >= blockCRCs[i].StartOffset && currentElement.Offset < blockCRCs[i].EndOffset {
				blockCRCs[i].CRC32.Write(idBytes)
				blockCRCs[i].CRC32.Write(sizeBytes)
			}
		}

		if writer != nil {
			writer.Write(idBytes)
			writer.Write(sizeBytes)
		}

		// ...

		if ElementID(id) == CRC32Element {
			data := make([]byte, size)
			_, err := reader.Read(data)
			if err != nil {
				return err
			}

			if writer != nil {
				writer.Write(data)
			}

			for i := range blockCRCs {
				if currentOffset+bytesRead >= blockCRCs[i].StartOffset && currentOffset+bytesRead <= blockCRCs[i].EndOffset {
					blockCRCs[i].CRC32.Write(data)
				}
			}

			blockCRCs = append(blockCRCs, BlockCRC{
				Element:     currentElement,
				StartOffset: currentElement.Offset + 6,
				EndOffset:   limit,
				ExpectedCRC: binary.LittleEndian.Uint32(data),
				CRC32:       crc32.NewIEEE(),
				Dirty:       false,
			})
		} else {
			if elementType == masterType /* && ElementID(id) != ClusterElement */ {
				if size > 0 {
					err = parseElement(reader, writer, level+1, currentOffset+bytesRead, currentOffset+bytesRead+int64(size), currentElement)
					if err != nil {
						return err
					}
				}
			} else {

				remaining := int(size)

				for remaining > 0 {
					toPeek := remaining
					if toPeek > reader.Buffered() {
						toPeek = reader.Buffered()
					}
					if toPeek == 0 {
						n := remaining
						if n > 4096 {
							n = 4096 // petite chunk pour éviter Peek trop gros
						}
						buf, err := reader.Peek(n)
						if err != nil {
							return err
						}
						toPeek = len(buf)
					}

					buf, err := reader.Peek(toPeek)
					if err != nil {
						return err
					}

					for i := range blockCRCs {
						if currentOffset+bytesRead >= blockCRCs[i].StartOffset && currentOffset+bytesRead <= blockCRCs[i].EndOffset {
							blockCRCs[i].CRC32.Write(buf)
						}
					}

					if writer != nil {
						writer.Write(buf)
					}

					if _, err := reader.Discard(len(buf)); err != nil {
						return err
					}

					remaining -= len(buf)
				}

				progressBytesReadCurrent = progressBytesReadCurrent + int64(size)
			}
		}

		bytesRead += int64(size)

		updateProgressBar(false)

	}

	return nil
}

func updateProgressBar(force bool) {
	colorFilled := "\033[32m"
	colorEmpty := "\033[30m"
	colorNone := "\033[0m"

	now := time.Now()
	if force || now.Sub(progressLastPrint) > 100*time.Millisecond {
		elapsed := time.Since(progressStart).Seconds()
		speed := float64(progressBytesReadCurrent) / (1024 * 1024) / elapsed // Mo/s
		eta := float64(progressBytesReadTotal-progressBytesReadCurrent) / (float64(progressBytesReadCurrent) / elapsed)
		percent := float64(progressBytesReadCurrent) / float64(progressBytesReadTotal) * 100

		var bar strings.Builder
		bar.Grow(50 + len(colorFilled) + len(colorEmpty))
		bar.WriteString(colorFilled)
		for i := 0; i < int(percent/2); i++ {
			bar.WriteString("─")
		}
		bar.WriteString(colorEmpty)
		for i := 0; i < 50-int(percent/2); i++ {
			bar.WriteString("─")
		}
		bar.WriteString(colorNone)

		line := fmt.Sprintf("Progress %s %6.2f%% • %6.2f MiB/s • ETA: %5.0fs", bar.String(), percent, speed, eta)
		fmt.Printf("\r%s\033[K", line)

		progressLastPrint = now
	}
}

func readVint(reader *bufio.Reader, keepLeadingBits bool) (value uint64, length uint64, err error) {
	firstByte, err := reader.Peek(1)
	if err != nil {
		return 0, 0, err
	}

	length = 0
	mask := byte(0x80) // 10000000
	for length = 1; length <= 8; length++ {
		if firstByte[0]&mask != 0 {
			break
		}
		mask >>= 1
	}

	if length > 8 {
		return 0, 0, fmt.Errorf("invalid EBML VINT: leading zero")
	}

	peeked, err := reader.Peek(int(length))
	if err != nil {
		return 0, 0, err
	}

	if keepLeadingBits {
		value = uint64(peeked[0])
	} else {
		value = uint64(peeked[0] & ^(0x80 >> (length - 1)))
	}

	for i := uint64(1); i < length; i++ {
		value = (value << 8) | uint64(peeked[i])
	}

	/*
		// check "unknown size" ??
		maxValue := uint64((1 << (7 * length)) - 1)
		if value == maxValue {
			fmt.Printf("Detected unknown size element (length=%d)\n", length)
			value = ^uint64(0)
		}
	*/

	return value, length, nil
}

func encodeVint(value uint64, length uint64) ([]byte, error) {
	if length == 0 {
		for l := uint64(1); l <= 8; l++ {
			maxValue := uint64((1 << (7 * l)) - 1)
			if value <= maxValue {
				length = l
				break
			}
		}
		if length == 0 {
			return nil, fmt.Errorf("value too large to encode as EBML VINT")
		}
	}

	if length < 1 || length > 8 {
		return nil, fmt.Errorf("invalid VINT length: %d", length)
	}

	maxValue := uint64((1 << (7 * length)) - 1)
	if value > maxValue {
		return nil, fmt.Errorf("value %d too large for VINT length %d", value, length)
	}

	data := make([]byte, length)

	for i := length - 1; i > 0; i-- {
		data[i] = byte(value & 0xFF)
		value >>= 8
	}

	data[0] = byte(value)
	data[0] |= 0x80 >> (length - 1)

	return data, nil
}

func getElementName(id ElementID) string {
	name, ok := elementNames[id]
	if !ok {
		return fmt.Sprintf("UNKNOWN:%x", id)
	}
	return name
}

func printElement(element *Element, level int) {
	if element.ID != 0 {
		elementName := getElementName(ElementID(element.ID))
		fmt.Printf("%s %s(0x%X) size=%d Offset=%d\n", strings.Repeat("  ", int(level-1)), elementName, element.ID, element.Size, element.Offset)
	}

	for _, child := range element.Children {
		printElement(child, level+1)
	}
}

func ebmlIDToBytes(id uint64) []byte {
	var buf [8]byte
	i := 8
	for id > 0 {
		i--
		buf[i] = byte(id & 0xFF)
		id >>= 8
	}
	return buf[i:]
}

/*
func writeDirtyElement(writer *os.File, element *Element) {
    elementName := getElementName(ElementID(element.ID))

    if element.Dirty {
        fmt.Printf("WRITE %s(0x%X) size=%d Offset=%d\n", elementName, element.ID, element.Size, element.Offset)

        a := ebmlIDToBytes(uint64(element.ID))
        fmt.Printf("FOO %x %x\n", element.ID, a)

        //writer.Seek(element.Offset + int64(len(a)), io.SeekStart)

        if element.Offset == 0 {
            return
        }

        data, _ := encodeVint(element.Size, 0)
        writer.Write(data)
    }

    for _, child := range element.Children {
        writeDirtyElement(writer, child)
    }
}
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./mkvcksum input.mkv [output.mkv]")
		return
	}

	input, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer input.Close()

	fileInfo, err := input.Stat()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	progressBytesReadTotal = fileInfo.Size()
	progressStart = time.Now()

	inputReader := bufio.NewReaderSize(input, 64*1024)

	var writer *os.File = nil

	if len(os.Args) == 3 {
		writer, _ = os.Create(os.Args[2])
		defer writer.Close()
	}

	parseElement(inputReader, writer, 0, 0, 0, nil)

	progressBytesReadCurrent = progressBytesReadTotal
	updateProgressBar(true)

	fmt.Printf("\r\033[K")

	//printElement(rootElement, 0)

	//if writer != nil {
	//    writeDirtyElement(writer, rootElement)
	//}

	if len(blockCRCs) == 0 {
		fmt.Printf("No CRC present\n")
	} else {
		for i := range blockCRCs {
			expectedCRC := blockCRCs[i].ExpectedCRC
			computedCRC := blockCRCs[i].CRC32.Sum32()
			if blockCRCs[i].Dirty {
				writer.Seek(blockCRCs[i].Element.WriteOffset+2, io.SeekStart)
				data := make([]byte, 4)
				binary.LittleEndian.PutUint32(data, uint32(computedCRC))
				writer.Write(data)
				expectedCRC = computedCRC
			}
			/*
			   path := ""
			   tmp := blockCRCs[i].Element.Parent
			   for tmp.Parent != nil {
			       path = getElementName(ElementID(tmp.ID)) + "/" + path
			       tmp = tmp.Parent
			   }
			*/
			if computedCRC == expectedCRC {
				fmt.Printf("CRC %2d ✅ (%08x)\n", i, computedCRC)
			} else {
				fmt.Printf("CRC %2d ❌ (%08x, expected %08x)\n", i, computedCRC, expectedCRC)
			}
		}
	}
}
