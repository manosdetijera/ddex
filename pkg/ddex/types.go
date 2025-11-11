package ddex

import (
	"encoding/xml"
	"time"
)

// Common types used throughout DDEX ERN 3.8 messages

// DateTime represents a date and time in ISO 8601 format
type DateTime struct {
	time.Time
}

// MarshalXML marshals DateTime to XML
func (dt *DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if dt.Time.IsZero() {
		return nil
	}
	return e.EncodeElement(dt.Time.Format(time.RFC3339), start)
}

// UnmarshalXML unmarshals DateTime from XML
func (dt *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

// EventDate represents a date for ERN 3.8 events (ISO 8601 format)
// Following ERN 3.8 standard specification for ReleaseDate and OriginalReleaseDate
// In ERN 3.8, dates don't have territory attributes at this level
type EventDate struct {
	XMLName               xml.Name `xml:",omitempty"`
	Value                 string   `xml:",chardata"`
	IsApproximate         bool     `xml:"IsApproximate,attr,omitempty"`
	IsBefore              bool     `xml:"IsBefore,attr,omitempty"`
	IsAfter               bool     `xml:"IsAfter,attr,omitempty"`
	TerritoryCode         string   `xml:"TerritoryCode,attr,omitempty"`
	LocationDescription   string   `xml:"LocationDescription,attr,omitempty"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// PartyID represents various party identification types
type PartyID struct {
	XMLName   xml.Name `xml:"PartyId"`
	Value     string   `xml:",chardata"`
	Namespace string   `xml:"Namespace,attr,omitempty"`
}

// ResourceID represents unique resource identification
type ResourceID struct {
	XMLName   xml.Name `xml:"ResourceId"`
	Value     string   `xml:",chardata"`
	Namespace string   `xml:"Namespace,attr,omitempty"`
}

// DisplayTitle
type DisplayTitle struct {
	XMLName   xml.Name    `xml:"DisplayTitle"`
	TitleText []TitleText `xml:"TitleText"`
}

// TitleText represents localized title information
type TitleText struct {
	XMLName               xml.Name `xml:"TitleText"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
	TitleType             string   `xml:"TitleType,attr,omitempty"`
}

// DisplayTitleText represents title suggested to show consumer
// ERN 3.8 version - simpler structure without territory attributes
type DisplayTitleText struct {
	XMLName               xml.Name `xml:"DisplayTitleText"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// Name represents party names with localization
type Name struct {
	//XMLName       xml.Name `xml:"Name"`
	FullName      string `xml:"FullName"`
	FullNameAscii string `xml:"FullNameAscii,omitempty"`
	LanguageCode  string `xml:"LanguageAndScriptCode,attr,omitempty"`
	NameType      string `xml:"NameType,attr,omitempty"`
}

// Territory represents geographic territories
type Territory struct {
	XMLName               xml.Name `xml:"Territory"`
	TerritoryCode         string   `xml:"TerritoryCode"`
	ExcludedTerritoryCode []string `xml:"ExcludedTerritoryCode,omitempty"`
}

// Duration represents time duration in ISO 8601 format
type Duration struct {
	XMLName xml.Name `xml:"Duration"`
	Value   string   `xml:",chardata"` // ISO 8601 duration format (PT3M30S)
}

// Keywords represents keywords for enhanced search and display
// ERN 3.8 version
type Keywords struct {
	XMLName               xml.Name `xml:"Keywords"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// Synopsis represents a synopsis with language attributes
// Following ERN 3.8 standard specification
type Synopsis struct {
	XMLName               xml.Name `xml:"Synopsis"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// MarketingComment represents a comment about the promotion and marketing of the Release
// Following ERN 3.8 standard specification
type MarketingComment struct {
	XMLName               xml.Name `xml:"MarketingComment"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// AvRating represents an audio-visual rating for a Release
// Following ERN 3.8 standard specification
type AvRating struct {
	XMLName      xml.Name      `xml:"AvRating"`
	RatingText   string        `xml:"RatingText,omitempty"`
	RatingAgency *RatingAgency `xml:"RatingAgency,omitempty"`
}

// RatingAgency represents a rating agency with optional namespace
type RatingAgency struct {
	Value     string `xml:",chardata"`
	Namespace string `xml:"Namespace,attr,omitempty"`
}

// VideoType represents the type of a video.
type VideoType struct {
	XMLName xml.Name `xml:"VideoType"`
	Value   string   `xml:",chardata"`
}

// DisplayArtistName represents a display artist name with language attributes
// Following ERN 3.8 standard specification - simpler than ERN 4.3
type DisplayArtistName struct {
	XMLName               xml.Name `xml:"DisplayArtistName"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// ResourceContributor represents a contributor to a resource (ERN 3.8)
type ResourceContributor struct {
	XMLName                       xml.Name `xml:"ResourceContributor"`
	PartyReference                string   `xml:"ResourceContributorPartyReference"`
	Role                          []string `xml:"ResourceContributorRole,omitempty"`
	InstrumentType                []string `xml:"InstrumentType,omitempty"`
	HasMadeFeaturedContribution   *bool    `xml:"HasMadeFeaturedContribution,omitempty"`
	HasMadeContractedContribution *bool    `xml:"HasMadeContractedContribution,omitempty"`
}

// IndirectResourceContributor represents an indirect contributor (ERN 3.8)
type IndirectResourceContributor struct {
	XMLName        xml.Name `xml:"IndirectResourceContributor"`
	PartyReference string   `xml:"IndirectResourceContributorPartyReference"`
	Role           []string `xml:"IndirectResourceContributorRole,omitempty"`
}

// RightsController represents a rights controller (TypedRightsController in ERN 3.8)
type RightsController struct {
	XMLName                        xml.Name  `xml:"RightsController"`
	SequenceNumber                 *int      `xml:"SequenceNumber,omitempty"`
	PartyName                      []Name    `xml:"PartyName,omitempty"`
	PartyId                        []PartyID `xml:"PartyId,omitempty"`
	RightsControllerPartyReference string    `xml:"RightsControllerPartyReference,omitempty"`
	RightsControllerRole           []string  `xml:"RightsControllerRole,omitempty"`
	RightSharePercentage           string    `xml:"RightSharePercentage,omitempty"`
	RightShareUnknown              string    `xml:"RightShareUnknown,omitempty"`
}

// HostSoundCarrier represents the sound carrier on which a resource was originally released (ERN 3.8)
type HostSoundCarrier struct {
	XMLName             xml.Name            `xml:"HostSoundCarrier"`
	ReleaseId           []ReleaseId         `xml:"ReleaseId,omitempty"`
	CatalogNumber       *CatalogNumber      `xml:"CatalogNumber,omitempty"`
	Title               []Title             `xml:"Title,omitempty"`
	DisplayArtistName   []DisplayArtistName `xml:"DisplayArtistName,omitempty"`
	DisplayArtist       []DisplayArtist     `xml:"DisplayArtist,omitempty"`
	OriginalReleaseDate *EventDate          `xml:"OriginalReleaseDate,omitempty"`
}
