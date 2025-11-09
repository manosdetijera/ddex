package ddex

import (
	"encoding/xml"
	"time"
)

// Common types used throughout DDEX ERN 4.3 messages

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
type DisplayTitleText struct {
	XMLName                 xml.Name `xml:"DisplayTitleText"`
	Value                   string   `xml:",chardata"`
	ApplicableTerritoryCode string   `xml:"ApplicableTerritoryCode,attr,omitempty"`
	LanguageAndScriptCode   string   `xml:"LanguageAndScriptCode,attr,omitempty"`
	IsDefault               bool     `xml:"IsDefault,attr,omitempty"`
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
type Keywords struct {
	XMLName                 xml.Name `xml:"Keywords"`
	Value                   string   `xml:",chardata"`
	ApplicableTerritoryCode string   `xml:"ApplicableTerritoryCode,attr,omitempty"`
	LanguageAndScriptCode   string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}
