package ddex

import "encoding/xml"

// ResourceList lists all Resources composites in a release
type ResourceList struct {
	XMLName        xml.Name         `xml:"ResourceList"`
	SoundRecording []SoundRecording `xml:"SoundRecording,omitempty"`
	Video          []Video          `xml:"Video,omitempty"`
	Image          []Image          `xml:"Image,omitempty"`
	Text           []Text           `xml:"Text,omitempty"`
}

// Video represents a video resource for ERN 3.8
type Video struct {
	XMLName                                xml.Name                                `xml:"Video"`
	IsUpdated                              *bool                                   `xml:"IsUpdated,attr,omitempty"` // Deprecated
	LanguageAndScriptCode                  string                                  `xml:"LanguageAndScriptCode,attr,omitempty"`
	ResourceReference                      string                                  `xml:"ResourceReference"` // Mandatory (ID)
	Type                                   string                                  `xml:"Type,omitempty"`    // VideoType
	IsArtistRelated                        *bool                                   `xml:"IsArtistRelated,omitempty"`
	VideoId                                []VideoId                               `xml:"VideoId,omitempty"`         // 0-n
	IndirectVideoId                        []MusicalWorkId                         `xml:"IndirectVideoId,omitempty"` // 0-n
	ResourceMusicalWorkReferenceList       *ResourceMusicalWorkReferenceList       `xml:"ResourceMusicalWorkReferenceList,omitempty"`
	ResourceContainedResourceReferenceList *ResourceContainedResourceReferenceList `xml:"ResourceContainedResourceReferenceList,omitempty"`

	// Choice: VideoCueSheetReference OR ReasonForCueSheetAbsence
	VideoCueSheetReference   []VideoCueSheetReference `xml:"VideoCueSheetReference,omitempty"`   // 1-n if present
	ReasonForCueSheetAbsence *Reason                  `xml:"ReasonForCueSheetAbsence,omitempty"` // 1 if present

	ReferenceTitle             *ReferenceTitle `xml:"ReferenceTitle,omitempty"`
	Title                      []Title         `xml:"Title,omitempty"` // 0-n
	InstrumentationDescription *Description    `xml:"InstrumentationDescription,omitempty"`

	// Boolean flags
	IsMedley                     *bool `xml:"IsMedley,omitempty"`
	IsPotpourri                  *bool `xml:"IsPotpourri,omitempty"`
	IsInstrumental               *bool `xml:"IsInstrumental,omitempty"`
	IsBackground                 *bool `xml:"IsBackground,omitempty"`
	IsHiddenResource             *bool `xml:"IsHiddenResource,omitempty"`
	IsBonusResource              *bool `xml:"IsBonusResource,omitempty"` // Deprecated
	HasPreOrderFulfillment       *bool `xml:"HasPreOrderFulfillment,omitempty"`
	IsRemastered                 *bool `xml:"IsRemastered,omitempty"`
	NoSilenceBefore              *bool `xml:"NoSilenceBefore,omitempty"`
	NoSilenceAfter               *bool `xml:"NoSilenceAfter,omitempty"`
	PerformerInformationRequired *bool `xml:"PerformerInformationRequired,omitempty"`

	// Language fields
	LanguageOfPerformance []string `xml:"LanguageOfPerformance,omitempty"` // ISO 639-2
	LanguageOfDubbing     []string `xml:"LanguageOfDubbing,omitempty"`     // ISO 639-2
	SubTitleLanguage      []string `xml:"SubTitleLanguage,omitempty"`      // ISO 639-2

	Duration                     string                                 `xml:"Duration"` // Mandatory
	RightsAgreementId            *RightsAgreementId                     `xml:"RightsAgreementId,omitempty"`
	VideoCollectionReferenceList *SoundRecordingCollectionReferenceList `xml:"VideoCollectionReferenceList,omitempty"`

	// Date fields
	CreationDate   *EventDate `xml:"CreationDate,omitempty"`
	MasteredDate   *EventDate `xml:"MasteredDate,omitempty"`
	RemasteredDate *EventDate `xml:"RemasteredDate,omitempty"`

	VideoDetailsByTerritory  []VideoDetailsByTerritory `xml:"VideoDetailsByTerritory"` // Mandatory 1-n
	TerritoryOfCommissioning string                    `xml:"TerritoryOfCommissioning,omitempty"`

	// Artist count fields
	NumberOfFeaturedArtists      *int `xml:"NumberOfFeaturedArtists,omitempty"`
	NumberOfNonFeaturedArtists   *int `xml:"NumberOfNonFeaturedArtists,omitempty"`
	NumberOfContractedArtists    *int `xml:"NumberOfContractedArtists,omitempty"`
	NumberOfNonContractedArtists *int `xml:"NumberOfNonContractedArtists,omitempty"`
}

// VideoDetailsByTerritory contains territory-specific video details for ERN 3.8
type VideoDetailsByTerritory struct {
	XMLName               xml.Name `xml:"VideoDetailsByTerritory"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`

	// Territory (choice: TerritoryCode OR ExcludedTerritoryCode, at least one required)
	TerritoryCode         []string `xml:"TerritoryCode,omitempty"`         // 1-n (if used)
	ExcludedTerritoryCode []string `xml:"ExcludedTerritoryCode,omitempty"` // 1-n (if used)

	// Title and display information
	Title            []Title         `xml:"Title,omitempty"`            // 0-n
	DisplayArtist    []DisplayArtist `xml:"DisplayArtist,omitempty"`    // 0-n
	DisplayConductor []DisplayArtist `xml:"DisplayConductor,omitempty"` // 0-n (uses Artist type)

	// Contributors
	ResourceContributor         []ResourceContributor         `xml:"ResourceContributor,omitempty"`         // 0-n
	IndirectResourceContributor []IndirectResourceContributor `xml:"IndirectResourceContributor,omitempty"` // 0-n

	// Rights and agreements
	RightsAgreementId *RightsAgreementId  `xml:"RightsAgreementId,omitempty"` // 0-1
	DisplayArtistName []DisplayArtistName `xml:"DisplayArtistName,omitempty"` // 0-n
	LabelName         []LabelName         `xml:"LabelName,omitempty"`         // 0-n
	RightsController  []RightsController  `xml:"RightsController,omitempty"`  // 0-n (TypedRightsController)

	// Dates
	RemasteredDate              *EventDate `xml:"RemasteredDate,omitempty"`              // 0-1
	ResourceReleaseDate         *EventDate `xml:"ResourceReleaseDate,omitempty"`         // 0-1
	OriginalResourceReleaseDate *EventDate `xml:"OriginalResourceReleaseDate,omitempty"` // 0-1

	// Copyright and credits
	PLine        []PLine       `xml:"PLine,omitempty"`        // 0-n
	CourtesyLine *CourtesyLine `xml:"CourtesyLine,omitempty"` // 0-1

	// Sequencing
	SequenceNumber *int `xml:"SequenceNumber,omitempty"` // 0-1

	// Descriptive metadata
	HostSoundCarrier    []HostSoundCarrier `xml:"HostSoundCarrier,omitempty"`    // 0-n
	MarketingComment    *Comment           `xml:"MarketingComment,omitempty"`    // 0-1
	Genre               []Genre            `xml:"Genre,omitempty"`               // 0-n
	ParentalWarningType []string           `xml:"ParentalWarningType,omitempty"` // 0-n (ParentalWarningType)
	AvRating            []AvRating         `xml:"AvRating,omitempty"`            // 0-n
	FulfillmentDate     *FulfillmentDate   `xml:"FulfillmentDate,omitempty"`     // 0-1
	Keywords            []Keywords         `xml:"Keywords,omitempty"`            // 0-n
	Synopsis            *Synopsis          `xml:"Synopsis,omitempty"`            // 0-1
	CLine               []CLine            `xml:"CLine,omitempty"`               // 0-n

	// Technical details
	TechnicalVideoDetails []TechnicalVideoDetails `xml:"TechnicalVideoDetails,omitempty"` // 0-n

	// Characters
	Character []Character `xml:"Character,omitempty"` // 0-n
}

// MusicalWorkId represents a musical work identifier
type MusicalWorkId struct {
	XMLName       xml.Name        `xml:"IndirectVideoId"`
	ISWC          string          `xml:"ISWC,omitempty"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// ResourceMusicalWorkReferenceList contains references to musical works
type ResourceMusicalWorkReferenceList struct {
	XMLName                      xml.Name                       `xml:"ResourceMusicalWorkReferenceList"`
	ResourceMusicalWorkReference []ResourceMusicalWorkReference `xml:"ResourceMusicalWorkReference,omitempty"`
}

// ResourceMusicalWorkReference references a musical work
type ResourceMusicalWorkReference struct {
	XMLName       xml.Name        `xml:"ResourceMusicalWorkReference"`
	MusicalWorkId []MusicalWorkId `xml:"MusicalWorkId,omitempty"`
	Duration      string          `xml:"Duration,omitempty"`
	StartPoint    string          `xml:"StartPoint,omitempty"`
}

// ResourceContainedResourceReferenceList contains references to contained resources
type ResourceContainedResourceReferenceList struct {
	XMLName                            xml.Name                             `xml:"ResourceContainedResourceReferenceList"`
	ResourceContainedResourceReference []ResourceContainedResourceReference `xml:"ResourceContainedResourceReference,omitempty"`
}

// ResourceContainedResourceReference references a contained resource
type ResourceContainedResourceReference struct {
	XMLName                            xml.Name `xml:"ResourceContainedResourceReference"`
	ResourceContainedResourceReference string   `xml:",chardata"`
	DurationUsed                       string   `xml:"DurationUsed,omitempty"`
	StartPoint                         string   `xml:"StartPoint,omitempty"`
}

// VideoCueSheetReference references a cue sheet
type VideoCueSheetReference struct {
	XMLName xml.Name `xml:"VideoCueSheetReference"`
	Value   string   `xml:",chardata"`
}

// Reason provides a textual reason
type Reason struct {
	XMLName               xml.Name `xml:"ReasonForCueSheetAbsence"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
	Value                 string   `xml:",chardata"`
}

// Description provides textual description
type Description struct {
	XMLName               xml.Name `xml:"InstrumentationDescription"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
	Value                 string   `xml:",chardata"`
}

// RightsAgreementId identifies rights agreements
type RightsAgreementId struct {
	XMLName       xml.Name        `xml:"RightsAgreementId"`
	MWLI          string          `xml:"MWLI,omitempty"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// SoundRecordingCollectionReferenceList contains collection references (used for VideoCollectionReferenceList)
type SoundRecordingCollectionReferenceList struct {
	XMLName                           xml.Name                            `xml:"VideoCollectionReferenceList"`
	SoundRecordingCollectionReference []SoundRecordingCollectionReference `xml:"SoundRecordingCollectionReference,omitempty"`
}

// SoundRecordingCollectionReference references a collection
type SoundRecordingCollectionReference struct {
	XMLName xml.Name `xml:"SoundRecordingCollectionReference"`
	Value   string   `xml:",chardata"`
}

// Character represents a character in the video
type Character struct {
	XMLName                 xml.Name `xml:"Character"`
	CharacterPartyReference string   `xml:"CharacterPartyReference,omitempty"`
	Name                    string   `xml:"Name,omitempty"`
}

// CourtesyLine represents a courtesy line
type CourtesyLine struct {
	XMLName          xml.Name `xml:"CourtesyLine"`
	Year             int      `xml:"Year,omitempty"`
	CourtesyLineText string   `xml:"CourtesyLineText"`
}

// FulfillmentDate represents a fulfillment date
type FulfillmentDate struct {
	XMLName                  xml.Name `xml:"FulfillmentDate"`
	FulfillmentDate          string   `xml:"FulfillmentDate"`
	ResourceReleaseReference string   `xml:"ResourceReleaseReference,omitempty"`
}

// VideoId represents video identification for ERN 3.8
type VideoId struct {
	XMLName       xml.Name        `xml:"VideoId"`
	ISRC          string          `xml:"ISRC,omitempty"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// Image represents an image resource for ERN 3.8
type Image struct {
	XMLName               xml.Name `xml:"Image"`
	IsUpdated             *bool    `xml:"IsUpdated,attr,omitempty"` // Deprecated (0-1)
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`

	// Type and classification
	ImageType       *ImageType `xml:"ImageType,omitempty"`       // 0-1
	IsArtistRelated *bool      `xml:"IsArtistRelated,omitempty"` // 0-1

	// Identifiers
	ImageId           []ImageId `xml:"ImageId"`           // Mandatory 1-n
	ResourceReference string    `xml:"ResourceReference"` // Mandatory (ID)

	// Descriptive information
	Title        []Title    `xml:"Title,omitempty"`        // 0-n
	CreationDate *EventDate `xml:"CreationDate,omitempty"` // 0-1

	// Territory-specific details
	ImageDetailsByTerritory []ImageDetailsByTerritory `xml:"ImageDetailsByTerritory"` // Mandatory 1-n
}

// ImageType represents the type of an image
type ImageType struct {
	XMLName xml.Name `xml:"ImageType"`
	Value   string   `xml:",chardata"`
}

// ImageDetailsByTerritory contains territory-specific image details for ERN 3.8
type ImageDetailsByTerritory struct {
	XMLName               xml.Name `xml:"ImageDetailsByTerritory"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`

	// Territory (choice: TerritoryCode OR ExcludedTerritoryCode, at least one required)
	TerritoryCode         []string `xml:"TerritoryCode,omitempty"`         // 1-n (if used)
	ExcludedTerritoryCode []string `xml:"ExcludedTerritoryCode,omitempty"` // 1-n (if used)

	// Title and contributors
	Title                       []Title                       `xml:"Title,omitempty"`                       // 0-n
	ResourceContributor         []ResourceContributor         `xml:"ResourceContributor,omitempty"`         // 0-n
	IndirectResourceContributor []IndirectResourceContributor `xml:"IndirectResourceContributor,omitempty"` // 0-n
	DisplayArtistName           []DisplayArtistName           `xml:"DisplayArtistName,omitempty"`           // 0-n

	// Copyright and credits
	CLine        []CLine       `xml:"CLine,omitempty"`        // 0-n
	Description  *Description  `xml:"Description,omitempty"`  // 0-1
	CourtesyLine *CourtesyLine `xml:"CourtesyLine,omitempty"` // 0-1

	// Dates
	ResourceReleaseDate         *EventDate       `xml:"ResourceReleaseDate,omitempty"`         // 0-1
	OriginalResourceReleaseDate *EventDate       `xml:"OriginalResourceReleaseDate,omitempty"` // 0-1
	FulfillmentDate             *FulfillmentDate `xml:"FulfillmentDate,omitempty"`             // 0-1

	// Descriptive metadata
	Keywords            []Keywords `xml:"Keywords,omitempty"`            // 0-n
	Synopsis            *Synopsis  `xml:"Synopsis,omitempty"`            // 0-1
	Genre               []Genre    `xml:"Genre,omitempty"`               // 0-n
	ParentalWarningType []string   `xml:"ParentalWarningType,omitempty"` // 0-n

	// Technical details
	TechnicalImageDetails []TechnicalImageDetails `xml:"TechnicalImageDetails,omitempty"` // 0-n
}

// ImageId represents image identification
type ImageId struct {
	XMLName       xml.Name        `xml:"ResourceId"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// IndirectResourceId represents an indirect resource identifier
type IndirectResourceId struct {
	XMLName       xml.Name        `xml:"IndirectResourceId"`
	ISRC          string          `xml:"ISRC,omitempty"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// SoundRecording represents an audio resource
type SoundRecording struct {
	XMLName           xml.Name          `xml:"SoundRecording"`
	ResourceReference string            `xml:"ResourceReference"`
	Type              string            `xml:"Type,omitempty"`
	ResourceId        []ResourceID      `xml:"ResourceId,omitempty"`
	DisplayTitleText  *DisplayTitleText `xml:"DisplayTitleText,omitempty"`
	DisplayTitle      *DisplayTitle     `xml:"DisplayTitle,omitempty"`
}

// Text represents a text resource
type Text struct {
	XMLName           xml.Name          `xml:"Text"`
	ResourceReference string            `xml:"ResourceReference"`
	Type              string            `xml:"Type,omitempty"`
	ResourceId        []ResourceID      `xml:"ResourceId,omitempty"`
	DisplayTitleText  *DisplayTitleText `xml:"DisplayTitleText,omitempty"`
}

// ResourceRightsController represents rights controller for a resource
type ResourceRightsController struct {
	XMLName                        xml.Name               `xml:"ResourceRightsController"`
	RightsControllerPartyReference string                 `xml:"RightsControllerPartyReference"`
	RightsControlType              string                 `xml:"RightsControlType,omitempty"`
	RightSharePercentage           string                 `xml:"RightSharePercentage,omitempty"`
	DelegatedUsageRights           []DelegatedUsageRights `xml:"DelegatedUsageRights,omitempty"`
}

// DelegatedUsageRights represents delegated rights
type DelegatedUsageRights struct {
	XMLName                     xml.Name `xml:"DelegatedUsageRights"`
	UseType                     []string `xml:"UseType"`
	TerritoryOfRightsDelegation []string `xml:"TerritoryOfRightsDelegation,omitempty"`
}

// WorkRightsController represents rights controller for musical works
type WorkRightsController struct {
	XMLName                        xml.Name               `xml:"WorkRightsController"`
	RightsControllerPartyReference string                 `xml:"RightsControllerPartyReference"`
	RightsControllerRole           string                 `xml:"RightsControllerRole,omitempty"`
	RightSharePercentage           string                 `xml:"RightSharePercentage,omitempty"`
	DelegatedUsageRights           []DelegatedUsageRights `xml:"DelegatedUsageRights,omitempty"`
}

// Technical details types for ERN 3.8
type TechnicalVideoDetails struct {
	XMLName                           xml.Name `xml:"TechnicalVideoDetails"`
	TechnicalResourceDetailsReference string   `xml:"TechnicalResourceDetailsReference"`
	VideoCodecType                    string   `xml:"VideoCodecType,omitempty"`
	VideoDefinitionType               string   `xml:"VideoDefinitionType,omitempty"`
	File                              *File    `xml:"File,omitempty"`
}

type TechnicalImageDetails struct {
	XMLName                           xml.Name `xml:"TechnicalImageDetails"`
	TechnicalResourceDetailsReference string   `xml:"TechnicalResourceDetailsReference"`
	ImageCodecType                    string   `xml:"ImageCodecType,omitempty"`
	ImageHeight                       int      `xml:"ImageHeight,omitempty"`
	ImageWidth                        int      `xml:"ImageWidth,omitempty"`
	File                              *File    `xml:"File,omitempty"`
}

type File struct {
	XMLName  xml.Name `xml:"File"`
	URI      string   `xml:"URI,omitempty"`
	HashSum  *HashSum `xml:"HashSum,omitempty"`
	FileSize int      `xml:"FileSize,omitempty"`
}

type HashSum struct {
	XMLName              xml.Name `xml:"HashSum"`
	HashSum              string   `xml:"HashSum"`
	HashSumAlgorithmType string   `xml:"HashSumAlgorithmType,omitempty"`
}

// Supporting types
type CreationDate struct {
	XMLName       xml.Name `xml:"CreationDate"`
	Value         string   `xml:",chardata"`
	IsApproximate bool     `xml:"IsApproximate,attr,omitempty"`
}

type ProprietaryId struct {
	XMLName   xml.Name `xml:"ProprietaryId"`
	Namespace string   `xml:"Namespace,attr,omitempty"`
	Value     string   `xml:",chardata"`
}

type PLine struct {
	XMLName   xml.Name `xml:"PLine"`
	Year      int      `xml:"Year,omitempty"`
	PLineText string   `xml:"PLineText"`
}

type CLine struct {
	XMLName   xml.Name `xml:"CLine"`
	Year      int      `xml:"Year,omitempty"`
	CLineText string   `xml:"CLineText"`
}

type Genre struct {
	XMLName                 xml.Name `xml:"Genre"`
	GenreText               string   `xml:"GenreText"`
	SubGenre                string   `xml:"SubGenre,omitempty"`
	ApplicableTerritoryCode string   `xml:"ApplicableTerritoryCode,attr,omitempty"`
}

// DisplayGenre represents genre information for display purposes (used in Release)
// Following ERN 4.3 standard specification
type DisplayGenre struct {
	XMLName                 xml.Name `xml:"DisplayGenre"`
	GenreText               string   `xml:"GenreText"`
	SubGenre                string   `xml:"SubGenre,omitempty"`
	ApplicableTerritoryCode string   `xml:"ApplicableTerritoryCode,attr,omitempty"`
}

type Contributor struct {
	XMLName                   xml.Name `xml:"Contributor"`
	SequenceNumber            int      `xml:"SequenceNumber,attr,omitempty"`
	ContributorPartyReference string   `xml:"ContributorPartyReference"`
	Role                      []string `xml:"Role"`
}
