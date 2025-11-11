package ddex

import "encoding/xml"

// ReleaseList lists all the Release composites
type ReleaseList struct {
	XMLName xml.Name  `xml:"ReleaseList"`
	Release []Release `xml:"Release"`
}

// Release represents a single release
type Release struct {
	XMLName               xml.Name                                `xml:"Release"`
	ReleaseReference      string                                  `xml:"ReleaseReference"`
	ReleaseType           string                                  `xml:"ReleaseType,omitempty"`
	ReleaseId             []ReleaseId                             `xml:"ReleaseId,omitempty"`
	DisplayTitleText      []DisplayTitleText                      `xml:"DisplayTitleText"`
	DisplayTitle          []DisplayTitle                          `xml:"DisplayTitle,omitempty"`
	DisplayArtistName     []DisplayArtistNameWithOriginalLanguage `xml:"DisplayArtistName,omitempty"`
	DisplayArtist         []DisplayArtist                         `xml:"DisplayArtist,omitempty"`
	ReleaseLabelReference []ReleaseLabelReference                 `xml:"ReleaseLabelReference,omitempty"`
	PLine                 []PLine                                 `xml:"PLine,omitempty"`
	CLine                 []CLine                                 `xml:"CLine,omitempty"`
	Duration              string                                  `xml:"Duration,omitempty"`
	ReleaseDate           []EventDateWithDefault                  `xml:"ReleaseDate,omitempty"`
	OriginalReleaseDate   []EventDateWithDefault                  `xml:"OriginalReleaseDate,omitempty"`
	DisplayGenre          []DisplayGenre                          `xml:"DisplayGenre,omitempty"`
	ParentalWarningType   string                                  `xml:"ParentalWarningType,omitempty"`
	AvRating              []AvRating                              `xml:"AvRating,omitempty"`
	RelatedResource       []RelatedResource                       `xml:"RelatedResource,omitempty"`
	ResourceGroup         []ResourceGroup                         `xml:"ResourceGroup,omitempty"`
	Keywords              []Keywords                              `xml:"Keywords,omitempty"`
	ContainsAI            string                                  `xml:"ContainsAI,omitempty"`
	MarketingComment      []MarketingComment                      `xml:"MarketingComment,omitempty"`
}

// RelatedResource represents a resource that is related to the release
type RelatedResource struct {
	XMLName                  xml.Name          `xml:"RelatedResource"`
	ResourceRelationshipType string            `xml:"ResourceRelationshipType"`
	ResourceId               RelatedResourceId `xml:"ResourceId"`
}

// RelatedResourceId represents the identifier for a related resource
type RelatedResourceId struct {
	XMLName xml.Name `xml:"ResourceId"`
	ISRC    string   `xml:"ISRC,omitempty"`
	ISNI    string   `xml:"ISNI,omitempty"`
}

// ReleaseId represents release identification (ICPN, GRid, UPC, EAN, etc.)
type ReleaseId struct {
	XMLName       xml.Name        `xml:"ReleaseId"`
	ICPN          *ICPN           `xml:"ICPN,omitempty"`
	GRid          string          `xml:"GRid,omitempty"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// ICPN represents UPC/EAN identifiers with proper IsEan attribute
type ICPN struct {
	XMLName xml.Name `xml:"ICPN"`
	Value   string   `xml:",chardata"`
	IsEan   bool     `xml:"IsEan,attr"`
}

// ReleaseLabelReference represents a reference to a label party
type ReleaseLabelReference struct {
	XMLName                 xml.Name `xml:"ReleaseLabelReference"`
	Value                   string   `xml:",chardata"`
	ApplicableTerritoryCode string   `xml:"ApplicableTerritoryCode,attr,omitempty"`
}

// ResourceGroup represents a grouping of resources within a release
type ResourceGroup struct {
	XMLName                  xml.Name                   `xml:"ResourceGroup"`
	AdditionalTitle          AdditionalTitle            `xml:"AdditionalTitle,omitempty"`
	SequenceNumber           int                        `xml:"SequenceNumber,omitempty"`
	ResourceGroupContentItem []ResourceGroupContentItem `xml:"ResourceGroupContentItem"`
}

// AdditionalTitle represents additional title information
type AdditionalTitle struct {
	XMLName   xml.Name `xml:"AdditionalTitle"`
	TitleText string   `xml:"TitleText"`
}

// ResourceGroupContentItem represents an item within a resource group
type ResourceGroupContentItem struct {
	XMLName                        xml.Name                         `xml:"ResourceGroupContentItem"`
	SequenceNumber                 int                              `xml:"SequenceNumber,omitempty"`
	ReleaseResourceReference       string                           `xml:"ReleaseResourceReference"`
	LinkedReleaseResourceReference []LinkedReleaseResourceReference `xml:"LinkedReleaseResourceReference,omitempty"`
}

// LinkedReleaseResourceReference represents a linked resource reference (e.g., cover art)
type LinkedReleaseResourceReference struct {
	XMLName         xml.Name `xml:"LinkedReleaseResourceReference"`
	LinkDescription string   `xml:"LinkDescription,attr,omitempty"`
	Value           string   `xml:",chardata"`
}
