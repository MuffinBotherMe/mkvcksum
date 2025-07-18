// This file is largely based on code from the go-mkvparse project,
// distributed under the MIT License: https://opensource.org/licenses/MIT

package main

// Supported ElementIDs
const (
	AlphaModeElement                   ElementID = 0x53C0
	AspectRatioTypeElement             ElementID = 0x54B3
	AttachedFileElement                ElementID = 0x61A7
	AttachmentLinkElement              ElementID = 0x7446
	AttachmentsElement                 ElementID = 0x1941A469
	AudioElement                       ElementID = 0xE1
	BitDepthElement                    ElementID = 0x6264
	BitsPerChannelElement              ElementID = 0x55B2
	BlockAddIDElement                  ElementID = 0xEE
	BlockAdditionIDElement             ElementID = 0xCB
	BlockAdditionalElement             ElementID = 0xA5
	BlockAdditionsElement              ElementID = 0x75A1
	BlockDurationElement               ElementID = 0x9B
	BlockElement                       ElementID = 0xA1
	BlockGroupElement                  ElementID = 0xA0
	BlockMoreElement                   ElementID = 0xA6
	BlockVirtualElement                ElementID = 0xA2
	CRC32Element                       ElementID = 0xBF
	CbSubsamplingHorzElement           ElementID = 0x55B5
	CbSubsamplingVertElement           ElementID = 0x55B6
	ChannelPositionsElement            ElementID = 0x7D7B
	ChannelsElement                    ElementID = 0x9F
	ChapCountryElement                 ElementID = 0x437E
	ChapLanguageElement                ElementID = 0x437C
	ChapLanguageIETFElement            ElementID = 0x437D
	ChapProcessCodecIDElement          ElementID = 0x6955
	ChapProcessCommandElement          ElementID = 0x6911
	ChapProcessDataElement             ElementID = 0x6933
	ChapProcessElement                 ElementID = 0x6944
	ChapProcessPrivateElement          ElementID = 0x450D
	ChapProcessTimeElement             ElementID = 0x6922
	ChapStringElement                  ElementID = 0x85
	ChapterAtomElement                 ElementID = 0xB6
	ChapterDisplayElement              ElementID = 0x80
	ChapterFlagEnabledElement          ElementID = 0x4598
	ChapterFlagHiddenElement           ElementID = 0x98
	ChapterPhysicalEquivElement        ElementID = 0x63C3
	ChapterSegmentEditionUIDElement    ElementID = 0x6EBC
	ChapterSegmentUIDElement           ElementID = 0x6E67
	ChapterStringUIDElement            ElementID = 0x5654
	ChapterTimeEndElement              ElementID = 0x92
	ChapterTimeStartElement            ElementID = 0x91
	ChapterTrackElement                ElementID = 0x8F
	ChapterTrackNumberElement          ElementID = 0x89
	ChapterTranslateCodecElement       ElementID = 0x69BF
	ChapterTranslateEditionUIDElement  ElementID = 0x69FC
	ChapterTranslateElement            ElementID = 0x6924
	ChapterTranslateIDElement          ElementID = 0x69A5
	ChapterUIDElement                  ElementID = 0x73C4
	ChaptersElement                    ElementID = 0x1043A770
	ChromaSitingHorzElement            ElementID = 0x55B7
	ChromaSitingVertElement            ElementID = 0x55B8
	ChromaSubsamplingHorzElement       ElementID = 0x55B3
	ChromaSubsamplingVertElement       ElementID = 0x55B4
	ClusterElement                     ElementID = 0x1F43B675
	CodecDecodeAllElement              ElementID = 0xAA
	CodecDelayElement                  ElementID = 0x56AA
	CodecDownloadURLElement            ElementID = 0x26B240
	CodecIDElement                     ElementID = 0x86
	CodecInfoURLElement                ElementID = 0x3B4040
	CodecNameElement                   ElementID = 0x258688
	CodecPrivateElement                ElementID = 0x63A2
	CodecSettingsElement               ElementID = 0x3A9697
	CodecStateElement                  ElementID = 0xA4
	ColourElement                      ElementID = 0x55B0
	ColourSpaceElement                 ElementID = 0x2EB524
	ContentCompAlgoElement             ElementID = 0x4254
	ContentCompSettingsElement         ElementID = 0x4255
	ContentCompressionElement          ElementID = 0x5034
	ContentEncAlgoElement              ElementID = 0x47E1
	ContentEncKeyIDElement             ElementID = 0x47E2
	ContentEncodingElement             ElementID = 0x6240
	ContentEncodingOrderElement        ElementID = 0x5031
	ContentEncodingScopeElement        ElementID = 0x5032
	ContentEncodingTypeElement         ElementID = 0x5033
	ContentEncodingsElement            ElementID = 0x6D80
	ContentEncryptionElement           ElementID = 0x5035
	ContentSigAlgoElement              ElementID = 0x47E5
	ContentSigHashAlgoElement          ElementID = 0x47E6
	ContentSigKeyIDElement             ElementID = 0x47E4
	ContentSignatureElement            ElementID = 0x47E3
	CueBlockNumberElement              ElementID = 0x5378
	CueClusterPositionElement          ElementID = 0xF1
	CueCodecStateElement               ElementID = 0xEA
	CueDurationElement                 ElementID = 0xB2
	CuePointElement                    ElementID = 0xBB
	CueRefClusterElement               ElementID = 0x97
	CueRefCodecStateElement            ElementID = 0xEB
	CueRefNumberElement                ElementID = 0x535F
	CueRefTimeElement                  ElementID = 0x96
	CueReferenceElement                ElementID = 0xDB
	CueRelativePositionElement         ElementID = 0xF0
	CueTimeElement                     ElementID = 0xB3
	CueTrackElement                    ElementID = 0xF7
	CueTrackPositionsElement           ElementID = 0xB7
	CuesElement                        ElementID = 0x1C53BB6B
	DateUTCElement                     ElementID = 0x4461
	DefaultDecodedFieldDurationElement ElementID = 0x234E7A
	DefaultDurationElement             ElementID = 0x23E383
	DelayElement                       ElementID = 0xCE
	DiscardPaddingElement              ElementID = 0x75A2
	DisplayHeightElement               ElementID = 0x54BA
	DisplayUnitElement                 ElementID = 0x54B2
	DisplayWidthElement                ElementID = 0x54B0
	DocTypeElement                     ElementID = 0x4282
	DocTypeExtensionElement            ElementID = 0x4281
	DocTypeExtensionNameElement        ElementID = 0x4283
	DocTypeExtensionVersionElement     ElementID = 0x4284
	DocTypeReadVersionElement          ElementID = 0x4285
	DocTypeVersionElement              ElementID = 0x4287
	DurationElement                    ElementID = 0x4489
	EBMLElement                        ElementID = 0x1A45DFA3
	EBMLMaxIDLengthElement             ElementID = 0x42F2
	EBMLMaxSizeLengthElement           ElementID = 0x42F3
	EBMLReadVersionElement             ElementID = 0x42F7
	EBMLVersionElement                 ElementID = 0x4286
	EditionEntryElement                ElementID = 0x45B9
	EditionFlagDefaultElement          ElementID = 0x45DB
	EditionFlagHiddenElement           ElementID = 0x45BD
	EditionFlagOrderedElement          ElementID = 0x45DD
	EditionUIDElement                  ElementID = 0x45BC
	EncryptedBlockElement              ElementID = 0xAF
	FieldOrderElement                  ElementID = 0x9D
	FileDataElement                    ElementID = 0x465C
	FileDescriptionElement             ElementID = 0x467E
	FileMimeTypeElement                ElementID = 0x4660
	FileNameElement                    ElementID = 0x466E
	FileReferralElement                ElementID = 0x4675
	FileUIDElement                     ElementID = 0x46AE
	FileUsedEndTimeElement             ElementID = 0x4662
	FileUsedStartTimeElement           ElementID = 0x4661
	FlagDefaultElement                 ElementID = 0x88
	FlagEnabledElement                 ElementID = 0xB9
	FlagForcedElement                  ElementID = 0x55AA
	FlagInterlacedElement              ElementID = 0x9A
	FlagLacingElement                  ElementID = 0x9C
	FrameNumberElement                 ElementID = 0xCD
	FrameRateElement                   ElementID = 0x2383E3
	GammaValueElement                  ElementID = 0x2FB523
	InfoElement                        ElementID = 0x1549A966
	LaceNumberElement                  ElementID = 0xCC
	LanguageElement                    ElementID = 0x22B59C
	LanguageIETFElement                ElementID = 0x22B59D
	LuminanceMaxElement                ElementID = 0x55D9
	LuminanceMinElement                ElementID = 0x55DA
	MasteringMetadataElement           ElementID = 0x55D0
	MatrixCoefficientsElement          ElementID = 0x55B1
	MaxBlockAdditionIDElement          ElementID = 0x55EE
	MaxCLLElement                      ElementID = 0x55BC
	MaxCacheElement                    ElementID = 0x6DF8
	MaxFALLElement                     ElementID = 0x55BD
	MinCacheElement                    ElementID = 0x6DE7
	MuxingAppElement                   ElementID = 0x4D80
	NameElement                        ElementID = 0x536E
	NextFilenameElement                ElementID = 0x3E83BB
	NextUIDElement                     ElementID = 0x3EB923
	OldStereoModeElement               ElementID = 0x53B9
	OutputSamplingFrequencyElement     ElementID = 0x78B5
	PixelCropBottomElement             ElementID = 0x54AA
	PixelCropLeftElement               ElementID = 0x54CC
	PixelCropRightElement              ElementID = 0x54DD
	PixelCropTopElement                ElementID = 0x54BB
	PixelHeightElement                 ElementID = 0xBA
	PixelWidthElement                  ElementID = 0xB0
	PositionElement                    ElementID = 0xA7
	PrevFilenameElement                ElementID = 0x3C83AB
	PrevSizeElement                    ElementID = 0xAB
	PrevUIDElement                     ElementID = 0x3CB923
	PrimariesElement                   ElementID = 0x55BB
	PrimaryBChromaticityXElement       ElementID = 0x55D5
	PrimaryBChromaticityYElement       ElementID = 0x55D6
	PrimaryGChromaticityXElement       ElementID = 0x55D3
	PrimaryGChromaticityYElement       ElementID = 0x55D4
	PrimaryRChromaticityXElement       ElementID = 0x55D1
	PrimaryRChromaticityYElement       ElementID = 0x55D2
	ProjectionElement                  ElementID = 0x7670
	ProjectionPosePitchElement         ElementID = 0x7674
	ProjectionPoseRollElement          ElementID = 0x7675
	ProjectionPoseYawElement           ElementID = 0x7673
	ProjectionPrivateElement           ElementID = 0x7672
	ProjectionTypeElement              ElementID = 0x7671
	RangeElement                       ElementID = 0x55B9
	ReferenceBlockElement              ElementID = 0xFB
	ReferenceFrameElement              ElementID = 0xC8
	ReferenceOffsetElement             ElementID = 0xC9
	ReferencePriorityElement           ElementID = 0xFA
	ReferenceTimeCodeElement           ElementID = 0xCA
	ReferenceVirtualElement            ElementID = 0xFD
	SamplingFrequencyElement           ElementID = 0xB5
	SeekElement                        ElementID = 0x4DBB
	SeekHeadElement                    ElementID = 0x114D9B74
	SeekIDElement                      ElementID = 0x53AB
	SeekPositionElement                ElementID = 0x53AC
	SeekPreRollElement                 ElementID = 0x56BB
	SegmentElement                     ElementID = 0x18538067
	SegmentFamilyElement               ElementID = 0x4444
	SegmentFilenameElement             ElementID = 0x7384
	SegmentUIDElement                  ElementID = 0x73A4
	SilentTrackNumberElement           ElementID = 0x58D7
	SilentTracksElement                ElementID = 0x5854
	SimpleBlockElement                 ElementID = 0xA3
	SimpleTagElement                   ElementID = 0x67C8
	SliceDurationElement               ElementID = 0xCF
	SlicesElement                      ElementID = 0x8E
	StereoModeElement                  ElementID = 0x53B8
	TagAttachmentUIDElement            ElementID = 0x63C6
	TagBinaryElement                   ElementID = 0x4485
	TagChapterUIDElement               ElementID = 0x63C4
	TagDefaultElement                  ElementID = 0x4484
	TagEditionUIDElement               ElementID = 0x63C9
	TagElement                         ElementID = 0x7373
	TagLanguageElement                 ElementID = 0x447A
	TagLanguageIETFElement             ElementID = 0x447B
	TagNameElement                     ElementID = 0x45A3
	TagStringElement                   ElementID = 0x4487
	TagTrackUIDElement                 ElementID = 0x63C5
	TagsElement                        ElementID = 0x1254C367
	TargetTypeElement                  ElementID = 0x63CA
	TargetTypeValueElement             ElementID = 0x68CA
	TargetsElement                     ElementID = 0x63C0
	TimeSliceElement                   ElementID = 0xE8
	TimecodeElement                    ElementID = 0xE7
	TimecodeScaleElement               ElementID = 0x2AD7B1
	TitleElement                       ElementID = 0x7BA9
	TrackCombinePlanesElement          ElementID = 0xE3
	TrackEntryElement                  ElementID = 0xAE
	TrackJoinBlocksElement             ElementID = 0xE9
	TrackJoinUIDElement                ElementID = 0xED
	TrackNumberElement                 ElementID = 0xD7
	TrackOffsetElement                 ElementID = 0x537F
	TrackOperationElement              ElementID = 0xE2
	TrackOverlayElement                ElementID = 0x6FAB
	TrackPlaneElement                  ElementID = 0xE4
	TrackPlaneTypeElement              ElementID = 0xE6
	TrackPlaneUIDElement               ElementID = 0xE5
	TrackTimecodeScaleElement          ElementID = 0x23314F
	TrackTranslateCodecElement         ElementID = 0x66BF
	TrackTranslateEditionUIDElement    ElementID = 0x66FC
	TrackTranslateElement              ElementID = 0x6624
	TrackTranslateTrackIDElement       ElementID = 0x66A5
	TrackTypeElement                   ElementID = 0x83
	TrackUIDElement                    ElementID = 0x73C5
	TracksElement                      ElementID = 0x1654AE6B
	TransferCharacteristicsElement     ElementID = 0x55BA
	TrickMasterTrackSegmentUIDElement  ElementID = 0xC4
	TrickMasterTrackUIDElement         ElementID = 0xC7
	TrickTrackFlagElement              ElementID = 0xC6
	TrickTrackSegmentUIDElement        ElementID = 0xC1
	TrickTrackUIDElement               ElementID = 0xC0
	VideoElement                       ElementID = 0xE0
	VoidElement                        ElementID = 0xEC
	WhitePointChromaticityXElement     ElementID = 0x55D7
	WhitePointChromaticityYElement     ElementID = 0x55D8
	WritingAppElement                  ElementID = 0x5741
	BlockAdditionMapping               ElementID = 0x41E4
	BlockAddIDValue                    ElementID = 0x41F0
	BlockAddIDName                     ElementID = 0x41A4
	BlockAddIDType                     ElementID = 0x41E7
	BlockAddIDExtraData                ElementID = 0x41ED
	FlagHearingImpaired                ElementID = 0x55AB
	FlagVisualImpaired                 ElementID = 0x55AC
	FlagTextDescriptions               ElementID = 0x55AD
	FlagOriginal                       ElementID = 0x55AE
	FlagCommentary                     ElementID = 0x55AF
)

func getElementType(el ElementID) elementType {
	switch el {
	case AlphaModeElement:
		return uintegerType
	case AspectRatioTypeElement:
		return uintegerType
	case AttachedFileElement:
		return masterType
	case AttachmentLinkElement:
		return uintegerType
	case AttachmentsElement:
		return masterType
	case AudioElement:
		return masterType
	case BitDepthElement:
		return uintegerType
	case BitsPerChannelElement:
		return uintegerType
	case BlockAddIDElement:
		return uintegerType
	case BlockAdditionIDElement:
		return uintegerType
	case BlockAdditionalElement:
		return binaryType
	case BlockAdditionsElement:
		return masterType
	case BlockDurationElement:
		return uintegerType
	case BlockElement:
		return binaryType
	case BlockGroupElement:
		return masterType
	case BlockMoreElement:
		return masterType
	case BlockVirtualElement:
		return binaryType
	case CRC32Element:
		return binaryType
	case CbSubsamplingHorzElement:
		return uintegerType
	case CbSubsamplingVertElement:
		return uintegerType
	case ChannelPositionsElement:
		return binaryType
	case ChannelsElement:
		return uintegerType
	case ChapCountryElement:
		return stringType
	case ChapLanguageElement:
		return stringType
	case ChapLanguageIETFElement:
		return stringType
	case ChapProcessCodecIDElement:
		return uintegerType
	case ChapProcessCommandElement:
		return masterType
	case ChapProcessDataElement:
		return binaryType
	case ChapProcessElement:
		return masterType
	case ChapProcessPrivateElement:
		return binaryType
	case ChapProcessTimeElement:
		return uintegerType
	case ChapStringElement:
		return utf8Type
	case ChapterAtomElement:
		return masterType
	case ChapterDisplayElement:
		return masterType
	case ChapterFlagEnabledElement:
		return uintegerType
	case ChapterFlagHiddenElement:
		return uintegerType
	case ChapterPhysicalEquivElement:
		return uintegerType
	case ChapterSegmentEditionUIDElement:
		return uintegerType
	case ChapterSegmentUIDElement:
		return binaryType
	case ChapterStringUIDElement:
		return utf8Type
	case ChapterTimeEndElement:
		return uintegerType
	case ChapterTimeStartElement:
		return uintegerType
	case ChapterTrackElement:
		return masterType
	case ChapterTrackNumberElement:
		return uintegerType
	case ChapterTranslateCodecElement:
		return uintegerType
	case ChapterTranslateEditionUIDElement:
		return uintegerType
	case ChapterTranslateElement:
		return masterType
	case ChapterTranslateIDElement:
		return binaryType
	case ChapterUIDElement:
		return uintegerType
	case ChaptersElement:
		return masterType
	case ChromaSitingHorzElement:
		return uintegerType
	case ChromaSitingVertElement:
		return uintegerType
	case ChromaSubsamplingHorzElement:
		return uintegerType
	case ChromaSubsamplingVertElement:
		return uintegerType
	case ClusterElement:
		return masterType
	case CodecDecodeAllElement:
		return uintegerType
	case CodecDelayElement:
		return uintegerType
	case CodecDownloadURLElement:
		return stringType
	case CodecIDElement:
		return stringType
	case CodecInfoURLElement:
		return stringType
	case CodecNameElement:
		return utf8Type
	case CodecPrivateElement:
		return binaryType
	case CodecSettingsElement:
		return utf8Type
	case CodecStateElement:
		return binaryType
	case ColourElement:
		return masterType
	case ColourSpaceElement:
		return binaryType
	case ContentCompAlgoElement:
		return uintegerType
	case ContentCompSettingsElement:
		return binaryType
	case ContentCompressionElement:
		return masterType
	case ContentEncAlgoElement:
		return uintegerType
	case ContentEncKeyIDElement:
		return binaryType
	case ContentEncodingElement:
		return masterType
	case ContentEncodingOrderElement:
		return uintegerType
	case ContentEncodingScopeElement:
		return uintegerType
	case ContentEncodingTypeElement:
		return uintegerType
	case ContentEncodingsElement:
		return masterType
	case ContentEncryptionElement:
		return masterType
	case ContentSigAlgoElement:
		return uintegerType
	case ContentSigHashAlgoElement:
		return uintegerType
	case ContentSigKeyIDElement:
		return binaryType
	case ContentSignatureElement:
		return binaryType
	case CueBlockNumberElement:
		return uintegerType
	case CueClusterPositionElement:
		return uintegerType
	case CueCodecStateElement:
		return uintegerType
	case CueDurationElement:
		return uintegerType
	case CuePointElement:
		return masterType
	case CueRefClusterElement:
		return uintegerType
	case CueRefCodecStateElement:
		return uintegerType
	case CueRefNumberElement:
		return uintegerType
	case CueRefTimeElement:
		return uintegerType
	case CueReferenceElement:
		return masterType
	case CueRelativePositionElement:
		return uintegerType
	case CueTimeElement:
		return uintegerType
	case CueTrackElement:
		return uintegerType
	case CueTrackPositionsElement:
		return masterType
	case CuesElement:
		return masterType
	case DateUTCElement:
		return dateType
	case DefaultDecodedFieldDurationElement:
		return uintegerType
	case DefaultDurationElement:
		return uintegerType
	case DelayElement:
		return uintegerType
	case DiscardPaddingElement:
		return integerType
	case DisplayHeightElement:
		return uintegerType
	case DisplayUnitElement:
		return uintegerType
	case DisplayWidthElement:
		return uintegerType
	case DocTypeElement:
		return stringType
	case DocTypeExtensionElement:
		return masterType
	case DocTypeExtensionNameElement:
		return stringType
	case DocTypeExtensionVersionElement:
		return uintegerType
	case DocTypeReadVersionElement:
		return uintegerType
	case DocTypeVersionElement:
		return uintegerType
	case DurationElement:
		return floatType
	case EBMLElement:
		return masterType
	case EBMLMaxIDLengthElement:
		return uintegerType
	case EBMLMaxSizeLengthElement:
		return uintegerType
	case EBMLReadVersionElement:
		return uintegerType
	case EBMLVersionElement:
		return uintegerType
	case EditionEntryElement:
		return masterType
	case EditionFlagDefaultElement:
		return uintegerType
	case EditionFlagHiddenElement:
		return uintegerType
	case EditionFlagOrderedElement:
		return uintegerType
	case EditionUIDElement:
		return uintegerType
	case EncryptedBlockElement:
		return binaryType
	case FieldOrderElement:
		return uintegerType
	case FileDataElement:
		return binaryType
	case FileDescriptionElement:
		return utf8Type
	case FileMimeTypeElement:
		return stringType
	case FileNameElement:
		return utf8Type
	case FileReferralElement:
		return binaryType
	case FileUIDElement:
		return uintegerType
	case FileUsedEndTimeElement:
		return uintegerType
	case FileUsedStartTimeElement:
		return uintegerType
	case FlagDefaultElement:
		return uintegerType
	case FlagEnabledElement:
		return uintegerType
	case FlagForcedElement:
		return uintegerType
	case FlagInterlacedElement:
		return uintegerType
	case FlagLacingElement:
		return uintegerType
	case FrameNumberElement:
		return uintegerType
	case FrameRateElement:
		return floatType
	case GammaValueElement:
		return floatType
	case InfoElement:
		return masterType
	case LaceNumberElement:
		return uintegerType
	case LanguageElement:
		return stringType
	case LanguageIETFElement:
		return stringType
	case LuminanceMaxElement:
		return floatType
	case LuminanceMinElement:
		return floatType
	case MasteringMetadataElement:
		return masterType
	case MatrixCoefficientsElement:
		return uintegerType
	case MaxBlockAdditionIDElement:
		return uintegerType
	case MaxCLLElement:
		return uintegerType
	case MaxCacheElement:
		return uintegerType
	case MaxFALLElement:
		return uintegerType
	case MinCacheElement:
		return uintegerType
	case MuxingAppElement:
		return utf8Type
	case NameElement:
		return utf8Type
	case NextFilenameElement:
		return utf8Type
	case NextUIDElement:
		return binaryType
	case OldStereoModeElement:
		return uintegerType
	case OutputSamplingFrequencyElement:
		return floatType
	case PixelCropBottomElement:
		return uintegerType
	case PixelCropLeftElement:
		return uintegerType
	case PixelCropRightElement:
		return uintegerType
	case PixelCropTopElement:
		return uintegerType
	case PixelHeightElement:
		return uintegerType
	case PixelWidthElement:
		return uintegerType
	case PositionElement:
		return uintegerType
	case PrevFilenameElement:
		return utf8Type
	case PrevSizeElement:
		return uintegerType
	case PrevUIDElement:
		return binaryType
	case PrimariesElement:
		return uintegerType
	case PrimaryBChromaticityXElement:
		return floatType
	case PrimaryBChromaticityYElement:
		return floatType
	case PrimaryGChromaticityXElement:
		return floatType
	case PrimaryGChromaticityYElement:
		return floatType
	case PrimaryRChromaticityXElement:
		return floatType
	case PrimaryRChromaticityYElement:
		return floatType
	case ProjectionElement:
		return masterType
	case ProjectionPosePitchElement:
		return floatType
	case ProjectionPoseRollElement:
		return floatType
	case ProjectionPoseYawElement:
		return floatType
	case ProjectionPrivateElement:
		return binaryType
	case ProjectionTypeElement:
		return uintegerType
	case RangeElement:
		return uintegerType
	case ReferenceBlockElement:
		return integerType
	case ReferenceFrameElement:
		return masterType
	case ReferenceOffsetElement:
		return uintegerType
	case ReferencePriorityElement:
		return uintegerType
	case ReferenceTimeCodeElement:
		return uintegerType
	case ReferenceVirtualElement:
		return integerType
	case SamplingFrequencyElement:
		return floatType
	case SeekElement:
		return masterType
	case SeekHeadElement:
		return masterType
	case SeekIDElement:
		return binaryType
	case SeekPositionElement:
		return uintegerType
	case SeekPreRollElement:
		return uintegerType
	case SegmentElement:
		return masterType
	case SegmentFamilyElement:
		return binaryType
	case SegmentFilenameElement:
		return utf8Type
	case SegmentUIDElement:
		return binaryType
	case SilentTrackNumberElement:
		return uintegerType
	case SilentTracksElement:
		return masterType
	case SimpleBlockElement:
		return binaryType
	case SimpleTagElement:
		return masterType
	case SliceDurationElement:
		return uintegerType
	case SlicesElement:
		return masterType
	case StereoModeElement:
		return uintegerType
	case TagAttachmentUIDElement:
		return uintegerType
	case TagBinaryElement:
		return binaryType
	case TagChapterUIDElement:
		return uintegerType
	case TagDefaultElement:
		return uintegerType
	case TagEditionUIDElement:
		return uintegerType
	case TagElement:
		return masterType
	case TagLanguageElement:
		return stringType
	case TagLanguageIETFElement:
		return stringType
	case TagNameElement:
		return utf8Type
	case TagStringElement:
		return utf8Type
	case TagTrackUIDElement:
		return uintegerType
	case TagsElement:
		return masterType
	case TargetTypeElement:
		return stringType
	case TargetTypeValueElement:
		return uintegerType
	case TargetsElement:
		return masterType
	case TimeSliceElement:
		return masterType
	case TimecodeElement:
		return uintegerType
	case TimecodeScaleElement:
		return uintegerType
	case TitleElement:
		return utf8Type
	case TrackCombinePlanesElement:
		return masterType
	case TrackEntryElement:
		return masterType
	case TrackJoinBlocksElement:
		return masterType
	case TrackJoinUIDElement:
		return uintegerType
	case TrackNumberElement:
		return uintegerType
	case TrackOffsetElement:
		return integerType
	case TrackOperationElement:
		return masterType
	case TrackOverlayElement:
		return uintegerType
	case TrackPlaneElement:
		return masterType
	case TrackPlaneTypeElement:
		return uintegerType
	case TrackPlaneUIDElement:
		return uintegerType
	case TrackTimecodeScaleElement:
		return floatType
	case TrackTranslateCodecElement:
		return uintegerType
	case TrackTranslateEditionUIDElement:
		return uintegerType
	case TrackTranslateElement:
		return masterType
	case TrackTranslateTrackIDElement:
		return binaryType
	case TrackTypeElement:
		return uintegerType
	case TrackUIDElement:
		return uintegerType
	case TracksElement:
		return masterType
	case TransferCharacteristicsElement:
		return uintegerType
	case TrickMasterTrackSegmentUIDElement:
		return binaryType
	case TrickMasterTrackUIDElement:
		return uintegerType
	case TrickTrackFlagElement:
		return uintegerType
	case TrickTrackSegmentUIDElement:
		return binaryType
	case TrickTrackUIDElement:
		return uintegerType
	case VideoElement:
		return masterType
	case VoidElement:
		return binaryType
	case WhitePointChromaticityXElement:
		return floatType
	case WhitePointChromaticityYElement:
		return floatType
	case WritingAppElement:
		return utf8Type
	case BlockAdditionMapping:
		return masterType
	case BlockAddIDValue:
		return uintegerType
	case BlockAddIDName:
		return stringType
	case BlockAddIDType:
		return uintegerType
	case BlockAddIDExtraData:
		return binaryType
	case FlagHearingImpaired:
		return uintegerType
	case FlagVisualImpaired:
		return uintegerType
	case FlagTextDescriptions:
		return uintegerType
	case FlagOriginal:
		return uintegerType
	case FlagCommentary:
		return uintegerType
	default:
		return elementType(0)
	}
}

var elementNames = map[ElementID]string{
	AlphaModeElement:                   "AlphaMode",
	AspectRatioTypeElement:             "AspectRatioType",
	AttachedFileElement:                "AttachedFile",
	AttachmentLinkElement:              "AttachmentLink",
	AttachmentsElement:                 "Attachments",
	AudioElement:                       "Audio",
	BitDepthElement:                    "BitDepth",
	BitsPerChannelElement:              "BitsPerChannel",
	BlockAddIDElement:                  "BlockAddID",
	BlockAdditionIDElement:             "BlockAdditionID",
	BlockAdditionalElement:             "BlockAdditional",
	BlockAdditionsElement:              "BlockAdditions",
	BlockDurationElement:               "BlockDuration",
	BlockElement:                       "Block",
	BlockGroupElement:                  "BlockGroup",
	BlockMoreElement:                   "BlockMore",
	BlockVirtualElement:                "BlockVirtual",
	CRC32Element:                       "CRC32",
	CbSubsamplingHorzElement:           "CbSubsamplingHorz",
	CbSubsamplingVertElement:           "CbSubsamplingVert",
	ChannelPositionsElement:            "ChannelPositions",
	ChannelsElement:                    "Channels",
	ChapCountryElement:                 "ChapCountry",
	ChapLanguageElement:                "ChapLanguage",
	ChapLanguageIETFElement:            "ChapLanguageIETF",
	ChapProcessCodecIDElement:          "ChapProcessCodecID",
	ChapProcessCommandElement:          "ChapProcessCommand",
	ChapProcessDataElement:             "ChapProcessData",
	ChapProcessElement:                 "ChapProcess",
	ChapProcessPrivateElement:          "ChapProcessPrivate",
	ChapProcessTimeElement:             "ChapProcessTime",
	ChapStringElement:                  "ChapString",
	ChapterAtomElement:                 "ChapterAtom",
	ChapterDisplayElement:              "ChapterDisplay",
	ChapterFlagEnabledElement:          "ChapterFlagEnabled",
	ChapterFlagHiddenElement:           "ChapterFlagHidden",
	ChapterPhysicalEquivElement:        "ChapterPhysicalEquiv",
	ChapterSegmentEditionUIDElement:    "ChapterSegmentEditionUID",
	ChapterSegmentUIDElement:           "ChapterSegmentUID",
	ChapterStringUIDElement:            "ChapterStringUID",
	ChapterTimeEndElement:              "ChapterTimeEnd",
	ChapterTimeStartElement:            "ChapterTimeStart",
	ChapterTrackElement:                "ChapterTrack",
	ChapterTrackNumberElement:          "ChapterTrackNumber",
	ChapterTranslateCodecElement:       "ChapterTranslateCodec",
	ChapterTranslateEditionUIDElement:  "ChapterTranslateEditionUID",
	ChapterTranslateElement:            "ChapterTranslate",
	ChapterTranslateIDElement:          "ChapterTranslateID",
	ChapterUIDElement:                  "ChapterUID",
	ChaptersElement:                    "Chapters",
	ChromaSitingHorzElement:            "ChromaSitingHorz",
	ChromaSitingVertElement:            "ChromaSitingVert",
	ChromaSubsamplingHorzElement:       "ChromaSubsamplingHorz",
	ChromaSubsamplingVertElement:       "ChromaSubsamplingVert",
	ClusterElement:                     "Cluster",
	CodecDecodeAllElement:              "CodecDecodeAll",
	CodecDelayElement:                  "CodecDelay",
	CodecDownloadURLElement:            "CodecDownloadURL",
	CodecIDElement:                     "CodecID",
	CodecInfoURLElement:                "CodecInfoURL",
	CodecNameElement:                   "CodecName",
	CodecPrivateElement:                "CodecPrivate",
	CodecSettingsElement:               "CodecSettings",
	CodecStateElement:                  "CodecState",
	ColourElement:                      "Colour",
	ColourSpaceElement:                 "ColourSpace",
	ContentCompAlgoElement:             "ContentCompAlgo",
	ContentCompSettingsElement:         "ContentCompSettings",
	ContentCompressionElement:          "ContentCompression",
	ContentEncAlgoElement:              "ContentEncAlgo",
	ContentEncKeyIDElement:             "ContentEncKeyID",
	ContentEncodingElement:             "ContentEncoding",
	ContentEncodingOrderElement:        "ContentEncodingOrder",
	ContentEncodingScopeElement:        "ContentEncodingScope",
	ContentEncodingTypeElement:         "ContentEncodingType",
	ContentEncodingsElement:            "ContentEncodings",
	ContentEncryptionElement:           "ContentEncryption",
	ContentSigAlgoElement:              "ContentSigAlgo",
	ContentSigHashAlgoElement:          "ContentSigHashAlgo",
	ContentSigKeyIDElement:             "ContentSigKeyID",
	ContentSignatureElement:            "ContentSignature",
	CueBlockNumberElement:              "CueBlockNumber",
	CueClusterPositionElement:          "CueClusterPosition",
	CueCodecStateElement:               "CueCodecState",
	CueDurationElement:                 "CueDuration",
	CuePointElement:                    "CuePoint",
	CueRefClusterElement:               "CueRefCluster",
	CueRefCodecStateElement:            "CueRefCodecState",
	CueRefNumberElement:                "CueRefNumber",
	CueRefTimeElement:                  "CueRefTime",
	CueReferenceElement:                "CueReference",
	CueRelativePositionElement:         "CueRelativePosition",
	CueTimeElement:                     "CueTime",
	CueTrackElement:                    "CueTrack",
	CueTrackPositionsElement:           "CueTrackPositions",
	CuesElement:                        "Cues",
	DateUTCElement:                     "DateUTC",
	DefaultDecodedFieldDurationElement: "DefaultDecodedFieldDuration",
	DefaultDurationElement:             "DefaultDuration",
	DelayElement:                       "Delay",
	DiscardPaddingElement:              "DiscardPadding",
	DisplayHeightElement:               "DisplayHeight",
	DisplayUnitElement:                 "DisplayUnit",
	DisplayWidthElement:                "DisplayWidth",
	DocTypeElement:                     "DocType",
	DocTypeExtensionElement:            "DocTypeExtension",
	DocTypeExtensionNameElement:        "DocTypeExtensionName",
	DocTypeExtensionVersionElement:     "DocTypeExtensionVersion",
	DocTypeReadVersionElement:          "DocTypeReadVersion",
	DocTypeVersionElement:              "DocTypeVersion",
	DurationElement:                    "Duration",
	EBMLElement:                        "EBML",
	EBMLMaxIDLengthElement:             "EBMLMaxIDLength",
	EBMLMaxSizeLengthElement:           "EBMLMaxSizeLength",
	EBMLReadVersionElement:             "EBMLReadVersion",
	EBMLVersionElement:                 "EBMLVersion",
	EditionEntryElement:                "EditionEntry",
	EditionFlagDefaultElement:          "EditionFlagDefault",
	EditionFlagHiddenElement:           "EditionFlagHidden",
	EditionFlagOrderedElement:          "EditionFlagOrdered",
	EditionUIDElement:                  "EditionUID",
	EncryptedBlockElement:              "EncryptedBlock",
	FieldOrderElement:                  "FieldOrder",
	FileDataElement:                    "FileData",
	FileDescriptionElement:             "FileDescription",
	FileMimeTypeElement:                "FileMimeType",
	FileNameElement:                    "FileName",
	FileReferralElement:                "FileReferral",
	FileUIDElement:                     "FileUID",
	FileUsedEndTimeElement:             "FileUsedEndTime",
	FileUsedStartTimeElement:           "FileUsedStartTime",
	FlagDefaultElement:                 "FlagDefault",
	FlagEnabledElement:                 "FlagEnabled",
	FlagForcedElement:                  "FlagForced",
	FlagInterlacedElement:              "FlagInterlaced",
	FlagLacingElement:                  "FlagLacing",
	FrameNumberElement:                 "FrameNumber",
	FrameRateElement:                   "FrameRate",
	GammaValueElement:                  "GammaValue",
	InfoElement:                        "Info",
	LaceNumberElement:                  "LaceNumber",
	LanguageElement:                    "Language",
	LanguageIETFElement:                "LanguageIETF",
	LuminanceMaxElement:                "LuminanceMax",
	LuminanceMinElement:                "LuminanceMin",
	MasteringMetadataElement:           "MasteringMetadata",
	MatrixCoefficientsElement:          "MatrixCoefficients",
	MaxBlockAdditionIDElement:          "MaxBlockAdditionID",
	MaxCLLElement:                      "MaxCLL",
	MaxCacheElement:                    "MaxCache",
	MaxFALLElement:                     "MaxFALL",
	MinCacheElement:                    "MinCache",
	MuxingAppElement:                   "MuxingApp",
	NameElement:                        "Name",
	NextFilenameElement:                "NextFilename",
	NextUIDElement:                     "NextUID",
	OldStereoModeElement:               "OldStereoMode",
	OutputSamplingFrequencyElement:     "OutputSamplingFrequency",
	PixelCropBottomElement:             "PixelCropBottom",
	PixelCropLeftElement:               "PixelCropLeft",
	PixelCropRightElement:              "PixelCropRight",
	PixelCropTopElement:                "PixelCropTop",
	PixelHeightElement:                 "PixelHeight",
	PixelWidthElement:                  "PixelWidth",
	PositionElement:                    "Position",
	PrevFilenameElement:                "PrevFilename",
	PrevSizeElement:                    "PrevSize",
	PrevUIDElement:                     "PrevUID",
	PrimariesElement:                   "Primaries",
	PrimaryBChromaticityXElement:       "PrimaryBChromaticityX",
	PrimaryBChromaticityYElement:       "PrimaryBChromaticityY",
	PrimaryGChromaticityXElement:       "PrimaryGChromaticityX",
	PrimaryGChromaticityYElement:       "PrimaryGChromaticityY",
	PrimaryRChromaticityXElement:       "PrimaryRChromaticityX",
	PrimaryRChromaticityYElement:       "PrimaryRChromaticityY",
	ProjectionElement:                  "Projection",
	ProjectionPosePitchElement:         "ProjectionPosePitch",
	ProjectionPoseRollElement:          "ProjectionPoseRoll",
	ProjectionPoseYawElement:           "ProjectionPoseYaw",
	ProjectionPrivateElement:           "ProjectionPrivate",
	ProjectionTypeElement:              "ProjectionType",
	RangeElement:                       "Range",
	ReferenceBlockElement:              "ReferenceBlock",
	ReferenceFrameElement:              "ReferenceFrame",
	ReferenceOffsetElement:             "ReferenceOffset",
	ReferencePriorityElement:           "ReferencePriority",
	ReferenceTimeCodeElement:           "ReferenceTimeCode",
	ReferenceVirtualElement:            "ReferenceVirtual",
	SamplingFrequencyElement:           "SamplingFrequency",
	SeekElement:                        "Seek",
	SeekHeadElement:                    "SeekHead",
	SeekIDElement:                      "SeekID",
	SeekPositionElement:                "SeekPosition",
	SeekPreRollElement:                 "SeekPreRoll",
	SegmentElement:                     "Segment",
	SegmentFamilyElement:               "SegmentFamily",
	SegmentFilenameElement:             "SegmentFilename",
	SegmentUIDElement:                  "SegmentUID",
	SilentTrackNumberElement:           "SilentTrackNumber",
	SilentTracksElement:                "SilentTracks",
	SimpleBlockElement:                 "SimpleBlock",
	SimpleTagElement:                   "SimpleTag",
	SliceDurationElement:               "SliceDuration",
	SlicesElement:                      "Slices",
	StereoModeElement:                  "StereoMode",
	TagAttachmentUIDElement:            "TagAttachmentUID",
	TagBinaryElement:                   "TagBinary",
	TagChapterUIDElement:               "TagChapterUID",
	TagDefaultElement:                  "TagDefault",
	TagEditionUIDElement:               "TagEditionUID",
	TagElement:                         "Tag",
	TagLanguageElement:                 "TagLanguage",
	TagLanguageIETFElement:             "TagLanguageIETF",
	TagNameElement:                     "TagName",
	TagStringElement:                   "TagString",
	TagTrackUIDElement:                 "TagTrackUID",
	TagsElement:                        "Tags",
	TargetTypeElement:                  "TargetType",
	TargetTypeValueElement:             "TargetTypeValue",
	TargetsElement:                     "Targets",
	TimeSliceElement:                   "TimeSlice",
	TimecodeElement:                    "Timecode",
	TimecodeScaleElement:               "TimecodeScale",
	TitleElement:                       "Title",
	TrackCombinePlanesElement:          "TrackCombinePlanes",
	TrackEntryElement:                  "TrackEntry",
	TrackJoinBlocksElement:             "TrackJoinBlocks",
	TrackJoinUIDElement:                "TrackJoinUID",
	TrackNumberElement:                 "TrackNumber",
	TrackOffsetElement:                 "TrackOffset",
	TrackOperationElement:              "TrackOperation",
	TrackOverlayElement:                "TrackOverlay",
	TrackPlaneElement:                  "TrackPlane",
	TrackPlaneTypeElement:              "TrackPlaneType",
	TrackPlaneUIDElement:               "TrackPlaneUID",
	TrackTimecodeScaleElement:          "TrackTimecodeScale",
	TrackTranslateCodecElement:         "TrackTranslateCodec",
	TrackTranslateEditionUIDElement:    "TrackTranslateEditionUID",
	TrackTranslateElement:              "TrackTranslate",
	TrackTranslateTrackIDElement:       "TrackTranslateTrackID",
	TrackTypeElement:                   "TrackType",
	TrackUIDElement:                    "TrackUID",
	TracksElement:                      "Tracks",
	TransferCharacteristicsElement:     "TransferCharacteristics",
	TrickMasterTrackSegmentUIDElement:  "TrickMasterTrackSegmentUID",
	TrickMasterTrackUIDElement:         "TrickMasterTrackUID",
	TrickTrackFlagElement:              "TrickTrackFlag",
	TrickTrackSegmentUIDElement:        "TrickTrackSegmentUID",
	TrickTrackUIDElement:               "TrickTrackUID",
	VideoElement:                       "Video",
	VoidElement:                        "Void",
	WhitePointChromaticityXElement:     "WhitePointChromaticityX",
	WhitePointChromaticityYElement:     "WhitePointChromaticityY",
	WritingAppElement:                  "WritingApp",
	BlockAdditionMapping:               "BlockAdditionMapping",
	BlockAddIDValue:                    "BlockAddIDValue",
	BlockAddIDName:                     "BlockAddIDName",
	BlockAddIDType:                     "BlockAddIDType",
	BlockAddIDExtraData:                "BlockAddIDExtraData",
	FlagHearingImpaired:                "FlagHearingImpaired",
	FlagVisualImpaired:                 "FlagVisualImpaired",
	FlagTextDescriptions:               "FlagTextDescriptions",
	FlagOriginal:                       "FlagOriginal",
	FlagCommentary:                     "FlagCommentary",
}
