package ddex

import "encoding/xml"

// DealList lists all Deal composites
type DealList struct {
	XMLName     xml.Name      `xml:"DealList"`
	ReleaseDeal []ReleaseDeal `xml:"ReleaseDeal"`
}

// ReleaseDeal represents a deal for a specific release
type ReleaseDeal struct {
	XMLName              xml.Name `xml:"ReleaseDeal"`
	DealReleaseReference string   `xml:"DealReleaseReference"`
	Deal                 []Deal   `xml:"Deal"`
}

// Deal represents commercial terms for a release
type Deal struct {
	XMLName   xml.Name   `xml:"Deal"`
	DealTerms *DealTerms `xml:"DealTerms"`
}

// DealTerms represents the commercial terms of a deal
type DealTerms struct {
	XMLName             xml.Name        `xml:"DealTerms"`
	TerritoryCode       string          `xml:"TerritoryCode,omitempty"`
	ValidityPeriod      *ValidityPeriod `xml:"ValidityPeriod,omitempty"`
	CommercialModelType []string        `xml:"CommercialModelType,omitempty"`
	UseType             []string        `xml:"UseType,omitempty"`
}

// ValidityPeriod represents time period validity information
type ValidityPeriod struct {
	XMLName       xml.Name `xml:"ValidityPeriod"`
	StartDate     string   `xml:"StartDate,omitempty"`
	StartDateTime string   `xml:"StartDateTime,omitempty"`
	EndDate       string   `xml:"EndDate,omitempty"`
}
