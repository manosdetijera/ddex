package ddex

import "encoding/xml"

// ReleaseList lists all the Release composites
type ReleaseList struct {
	XMLName xml.Name  `xml:"ReleaseList"`
	Release []Release `xml:"Release"`
}

// Release represents a single release for ERN 3.8
// Following ERN 3.8 specification with mandatory ReferenceTitle and ReleaseDetailsByTerritory
type Release struct {
	XMLName                   xml.Name                    `xml:"Release"`
	LanguageAndScriptCode     string                      `xml:"LanguageAndScriptCode,attr,omitempty"`
	IsMainRelease             bool                        `xml:"IsMainRelease,attr,omitempty"`
	ReleaseReference          string                      `xml:"ReleaseReference,omitempty"`
	ReleaseId                 []ReleaseId                 `xml:"ReleaseId"`
	ExternalResourceLink      []ExternalResourceLink      `xml:"ExternalResourceLink,omitempty"`
	ReferenceTitle            *ReferenceTitle             `xml:"ReferenceTitle"`
	ReleaseResourceReference  []string                    `xml:"ReleaseResourceReference,omitempty"`
	ReleaseType               []ReleaseType               `xml:"ReleaseType,omitempty"`
	ReleaseDetailsByTerritory []ReleaseDetailsByTerritory `xml:"ReleaseDetailsByTerritory"`
	LanguageOfPerformance     []string                    `xml:"LanguageOfPerformance,omitempty"`
	LanguageOfDubbing         []string                    `xml:"LanguageOfDubbing,omitempty"`
	SubTitleLanguage          []string                    `xml:"SubTitleLanguage,omitempty"`
	Duration                  string                      `xml:"Duration,omitempty"`
	PLine                     []PLine                     `xml:"PLine,omitempty"`
	CLine                     []CLine                     `xml:"CLine,omitempty"`
	GlobalReleaseDate         *EventDate                  `xml:"GlobalReleaseDate,omitempty"`
	GlobalOriginalReleaseDate *EventDate                  `xml:"GlobalOriginalReleaseDate,omitempty"`
}

// ReferenceTitle represents the reference title of a release (mandatory in ERN 3.8)
type ReferenceTitle struct {
	XMLName   xml.Name `xml:"ReferenceTitle"`
	TitleText string   `xml:"TitleText"`
	SubTitle  string   `xml:"SubTitle,omitempty"`
}

// ReleaseType represents the form in which a release is offered
type ReleaseType struct {
	XMLName xml.Name `xml:"ReleaseType"`
	Value   string   `xml:",chardata"`
}

// ExternalResourceLink represents promotional or other material related to the release
type ExternalResourceLink struct {
	XMLName xml.Name `xml:"ExternalResourceLink"`
	URL     string   `xml:"URL"`
}

// ReleaseDetailsByTerritory contains territory-specific release details (mandatory in ERN 3.8)
type ReleaseDetailsByTerritory struct {
	XMLName                     xml.Name                      `xml:"ReleaseDetailsByTerritory"`
	LanguageAndScriptCode       string                        `xml:"LanguageAndScriptCode,attr,omitempty"`
	TerritoryCode               []string                      `xml:"TerritoryCode,omitempty"`
	ExcludedTerritoryCode       []string                      `xml:"ExcludedTerritoryCode,omitempty"`
	DisplayArtistName           []Name                        `xml:"DisplayArtistName,omitempty"`
	LabelName                   []LabelName                   `xml:"LabelName,omitempty"`
	Title                       []Title                       `xml:"Title,omitempty"`
	DisplayArtist               []DisplayArtist               `xml:"DisplayArtist,omitempty"`
	IsMultiArtistCompilation    bool                          `xml:"IsMultiArtistCompilation,omitempty"`
	AdministratingRecordCompany []AdministratingRecordCompany `xml:"AdministratingRecordCompany,omitempty"`
	ReleaseType                 []ReleaseType                 `xml:"ReleaseType,omitempty"`
	RelatedRelease              []RelatedRelease              `xml:"RelatedRelease,omitempty"`
	ParentalWarningType         []ParentalWarningType         `xml:"ParentalWarningType,omitempty"`
	AvRating                    []AvRating                    `xml:"AvRating,omitempty"`
	MarketingComment            *Comment                      `xml:"MarketingComment,omitempty"`
	ResourceGroup               []ResourceGroup               `xml:"ResourceGroup,omitempty"`
	Genre                       []Genre                       `xml:"Genre,omitempty"`
	PLine                       []PLine                       `xml:"PLine,omitempty"`
	CLine                       []CLine                       `xml:"CLine,omitempty"`
	ReleaseDate                 *EventDate                    `xml:"ReleaseDate,omitempty"`
	OriginalReleaseDate         *EventDate                    `xml:"OriginalReleaseDate,omitempty"`
	Keywords                    []Keywords                    `xml:"Keywords,omitempty"`
	Synopsis                    *Synopsis                     `xml:"Synopsis,omitempty"`
}

// LabelName represents the label name
type LabelName struct {
	XMLName               xml.Name `xml:"LabelName"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// Title represents a title (different from DisplayTitle)
type Title struct {
	XMLName   xml.Name `xml:"Title"`
	TitleText string   `xml:"TitleText"`
	SubTitle  string   `xml:"SubTitle,omitempty"`
}

// AdministratingRecordCompany represents the administrating record company
type AdministratingRecordCompany struct {
	XMLName     xml.Name  `xml:"AdministratingRecordCompany"`
	PartyId     []PartyId `xml:"PartyId,omitempty"`
	PartyName   []Name    `xml:"PartyName,omitempty"`
	TradingName string    `xml:"TradingName,omitempty"`
}

// ParentalWarningType represents parental warning classification
type ParentalWarningType struct {
	XMLName xml.Name `xml:"ParentalWarningType"`
	Value   string   `xml:",chardata"`
}

// Comment represents a comment (used for MarketingComment, etc.)
type Comment struct {
	XMLName               xml.Name `xml:",omitempty"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// RelatedRelease represents a related release
type RelatedRelease struct {
	XMLName                 xml.Name  `xml:"RelatedRelease"`
	ReleaseRelationshipType string    `xml:"ReleaseRelationshipType"`
	ReleaseId               ReleaseId `xml:"ReleaseId"`
}

// ReleaseId represents release identification (ICPN, GRid, ISRC, etc.) for ERN 3.8
type ReleaseId struct {
	XMLName       xml.Name        `xml:"ReleaseId"`
	GRid          string          `xml:"GRid,omitempty"`          // 0-1
	ISRC          string          `xml:"ISRC,omitempty"`          // 0-1
	ICPN          string          `xml:"ICPN,omitempty"`          // 0-1
	ISAN          string          `xml:"ISAN,omitempty"`          // 0-1
	CatalogNumber *CatalogNumber  `xml:"CatalogNumber,omitempty"` // 0-1
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"` // 0-n
}

// CatalogNumber represents a catalog number
type CatalogNumber struct {
	XMLName   xml.Name `xml:"CatalogNumber"`
	Value     string   `xml:",chardata"`
	Namespace string   `xml:"Namespace,attr,omitempty"`
}

// ReleaseLabelReference has been simplified to just string in ERN 3.8

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
