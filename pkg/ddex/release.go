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
	XMLName                        xml.Name                        `xml:"Release"`
	LanguageAndScriptCode          string                          `xml:"LanguageAndScriptCode,attr,omitempty"`
	IsMainRelease                  bool                            `xml:"IsMainRelease,attr,omitempty"`
	ReleaseId                      []ReleaseId                     `xml:"ReleaseId"`                                // 1-n
	ReleaseReference               string                          `xml:"ReleaseReference,omitempty"`               // Mandatory (ID)
	DisplayTitleText               []DisplayTitleText              `xml:"DisplayTitleText,omitempty"`               // 0-n
	DisplayTitle                   []DisplayTitle                  `xml:"DisplayTitle,omitempty"`                   // 0-n
	AdditionalTitle                []AdditionalTitle               `xml:"AdditionalTitle,omitempty"`                // 0-n
	ExternalResourceLink           []ExternalResourceLink          `xml:"ExternalResourceLink,omitempty"`           // 0-n
	ReferenceTitle                 *ReferenceTitle                 `xml:"ReferenceTitle"`                           // Mandatory (1)
	ReleaseResourceReferenceList   *ReleaseResourceReferenceList   `xml:"ReleaseResourceReferenceList,omitempty"`   // 0-1
	ReleaseCollectionReferenceList *ReleaseCollectionReferenceList `xml:"ReleaseCollectionReferenceList,omitempty"` // 0-1
	IsCompilation                  *bool                           `xml:"IsCompilation,omitempty"`                  // 0-1
	ReleaseType                    []ReleaseType                   `xml:"ReleaseType,omitempty"`                    // 0-n
	ReleaseDetailsByTerritory      []ReleaseDetailsByTerritory     `xml:"ReleaseDetailsByTerritory"`                // 1-n (Mandatory)
	LanguageOfPerformance          []string                        `xml:"LanguageOfPerformance,omitempty"`          // 0-n
	LanguageOfDubbing              []string                        `xml:"LanguageOfDubbing,omitempty"`              // 0-n
	SubTitleLanguage               []string                        `xml:"SubTitleLanguage,omitempty"`               // 0-n
	Duration                       string                          `xml:"Duration,omitempty"`                       // 0-1
	PLine                          []PLine                         `xml:"PLine,omitempty"`                          // 0-n
	CLine                          []CLine                         `xml:"CLine,omitempty"`                          // 0-n
	GlobalReleaseDate              *EventDate                      `xml:"GlobalReleaseDate,omitempty"`              // 0-1
	GlobalOriginalReleaseDate      *EventDate                      `xml:"GlobalOriginalReleaseDate,omitempty"`      // 0-1
}

// ReleaseResourceReferenceList represents a list of resource references
type ReleaseResourceReferenceList struct {
	XMLName                  xml.Name                   `xml:"ReleaseResourceReferenceList"`
	ReleaseResourceReference []ReleaseResourceReference `xml:"ReleaseResourceReference"`
}

// ReleaseResourceReference represents a single resource reference with its type
type ReleaseResourceReference struct {
	ReleaseResourceType string `xml:"ReleaseResourceType,attr,omitempty"`
	Value               string `xml:",chardata"`
}

// ReleaseCollectionReferenceList represents a list of collection references
type ReleaseCollectionReferenceList struct {
	XMLName                    xml.Name `xml:"ReleaseCollectionReferenceList"`
	ReleaseCollectionReference []string `xml:"ReleaseCollectionReference"`
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
	DisplayArtistName           []DisplayArtistName           `xml:"DisplayArtistName,omitempty"`
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
	LabelNameType         string   `xml:"LabelNameType,attr,omitempty"`
	Value                 string   `xml:",chardata"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
}

// Title represents a title (different from DisplayTitle)
type Title struct {
	XMLName               xml.Name `xml:"Title"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`
	TitleType             string   `xml:"TitleType,attr,omitempty"`
	TitleText             string   `xml:"TitleText"`
	SubTitle              string   `xml:"SubTitle,omitempty"`
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
	ReleaseId               ReleaseId `xml:"ReleaseId"`
	ReleaseRelationshipType string    `xml:"ReleaseRelationshipType"`
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
	Title                    Title                      `xml:"Title,omitempty"`
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
	ResourceType                   string                           `xml:"ResourceType,omitempty"`
	ReleaseResourceReference       ReleaseResourceReference         `xml:"ReleaseResourceReference"`
	LinkedReleaseResourceReference []LinkedReleaseResourceReference `xml:"LinkedReleaseResourceReference,omitempty"`
}

// LinkedReleaseResourceReference represents a linked resource reference (e.g., cover art)
type LinkedReleaseResourceReference struct {
	XMLName         xml.Name `xml:"LinkedReleaseResourceReference"`
	LinkDescription string   `xml:"LinkDescription,attr,omitempty"`
	Value           string   `xml:",chardata"`
}
