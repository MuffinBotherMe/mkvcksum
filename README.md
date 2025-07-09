# ✨ Features

Matroska (based on EBML) allows embedding CRC-32 elements to verify data integrity.

This is a small command-line tool written in Go to handle CRC-32 checksums inside MKV (Matroska) files.

- Verify all existing CRC-32 checksums already present
- In write mode, add or update a CRC-32 element at the beginning of the Segment to protect the entire file content

# 📚 Example usage

```
Usage: mkvcksum input.mkv [output.mkv]
```

```
$ ./mkvcksum test.mkv
CRC  0 ✅ (9c85568f)
CRC  1 ✅ (5fb66baa)
CRC  2 ✅ (fad30b9d)
CRC  3 ✅ (0d5b9517)
CRC  4 ✅ (3642ff42)
CRC  5 ✅ (eb6c2c13)
CRC  6 ✅ (83d8875c)
CRC  7 ✅ (fad30b9d)
CRC  8 ✅ (f4bee3ad)
CRC  9 ✅ (89c62342)
CRC 10 ✅ (91d2a6c0)
CRC 11 ✅ (eb6c2c13)
CRC 12 ✅ (44a9b67c)
CRC 13 ✅ (fad30b9d)
CRC 14 ✅ (60216c67)
CRC 15 ✅ (dc63ec29)
CRC 16 ✅ (f98b4eba)
```

```
$ ./mkvcksum test-no-crc.mkv
No CRC present
```

```
$ ./mkvcksum test-no-crc.mkv output.mkv
CRC  0 ✅ (4e758c65)
```

```
$ mkvinfo output.mkv | head -n 10
+ EBML head
|+ Document type: matroska
|+ Document type version: 2
|+ Document type read version: 2
+ Segment: size 23339311
|+ EBML CRC-32: 0x4e758c65 <-----------------------------------------------
|+ Seek head (subentries will be skipped)
|+ Segment information
| + Duration: 00:01:27.336000000
| + Multiplexing application: libebml2 v0.10.0 + libmatroska2 v0.10.1
```


# 🚀 Building

```
go build
```

# 🖥 Platform support

Fully cross-platform (tested on Linux + Windows).

# 📌 TODO

 - Patch SeekHead+Cues offsets
 - Code cleanup

# 📄 License

MIT. See [LICENSE](./LICENSE)
