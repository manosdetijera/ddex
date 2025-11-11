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

// DealTerms represents the commercial terms of a deal for ERN 3.8
type DealTerms struct {
	XMLName               xml.Name `xml:"DealTerms"`
	LanguageAndScriptCode string   `xml:"LanguageAndScriptCode,attr,omitempty"`

	// Pre-order and deal flags
	IsPreOrderDeal *bool `xml:"IsPreOrderDeal,omitempty"` // 0-1

	// Commercial model
	CommercialModelType []string `xml:"CommercialModelType,omitempty"` // 0-n

	// Usage (choice with AllDealsCancelled or TakeDown, at least one required)
	Usage             []Usage `xml:"Usage,omitempty"`             // 1-n (if used)
	AllDealsCancelled *bool   `xml:"AllDealsCancelled,omitempty"` // 1 (if used, deprecated)
	TakeDown          *bool   `xml:"TakeDown,omitempty"`          // 1 (if used, deprecated)

	// Territory (choice: TerritoryCode OR ExcludedTerritoryCode, at least one required)
	TerritoryCode         []string `xml:"TerritoryCode,omitempty"`         // 1-n (if used)
	ExcludedTerritoryCode []string `xml:"ExcludedTerritoryCode,omitempty"` // 1-n (if used)

	// Distribution channels (choice: DistributionChannel OR ExcludedDistributionChannel)
	DistributionChannel         []DSP `xml:"DistributionChannel,omitempty"`         // 1-n (if used)
	ExcludedDistributionChannel []DSP `xml:"ExcludedDistributionChannel,omitempty"` // 1-n (if used)

	// Pricing
	PriceInformation []PriceInformation `xml:"PriceInformation,omitempty"` // 0-n

	// Promotional (choice: IsPromotional OR PromotionalCode)
	IsPromotional   *bool            `xml:"IsPromotional,omitempty"`   // 1 (if used)
	PromotionalCode *PromotionalCode `xml:"PromotionalCode,omitempty"` // 1 (if used)

	// Validity period
	ValidityPeriod []ValidityPeriod `xml:"ValidityPeriod,omitempty"` // 1-n

	// Rental and pre-order
	ConsumerRentalPeriod *ConsumerRentalPeriod `xml:"ConsumerRentalPeriod,omitempty"` // 0-1
	PreOrderReleaseDate  *EventDate            `xml:"PreOrderReleaseDate,omitempty"`  // 0-1

	// Display dates (choice: structured dates OR PreOrderPreviewDate)
	ReleaseDisplayStartDate      string     `xml:"ReleaseDisplayStartDate,omitempty"`      // 1 (if used)
	TrackListingPreviewStartDate string     `xml:"TrackListingPreviewStartDate,omitempty"` // 1 (if used)
	CoverArtPreviewStartDate     string     `xml:"CoverArtPreviewStartDate,omitempty"`     // 1 (if used)
	ClipPreviewStartDate         string     `xml:"ClipPreviewStartDate,omitempty"`         // 1 (if used)
	PreOrderPreviewDate          *EventDate `xml:"PreOrderPreviewDate,omitempty"`          // 1 (if used, deprecated)

	// Incentive resources
	PreOrderIncentiveResourceList    *DealResourceReferenceList `xml:"PreOrderIncentiveResourceList,omitempty"`    // 0-1
	InstantGratificationResourceList *DealResourceReferenceList `xml:"InstantGratificationResourceList,omitempty"` // 0-1

	// Exclusivity and offers
	IsExclusive            *bool                    `xml:"IsExclusive,omitempty"`            // 0-1
	RelatedReleaseOfferSet []RelatedReleaseOfferSet `xml:"RelatedReleaseOfferSet,omitempty"` // 0-n

	// Physical distribution
	PhysicalReturns           *PhysicalReturns `xml:"PhysicalReturns,omitempty"`           // 0-1
	NumberOfProductsPerCarton *int             `xml:"NumberOfProductsPerCarton,omitempty"` // 0-1

	// Policies
	RightsClaimPolicy []RightsClaimPolicy `xml:"RightsClaimPolicy,omitempty"` // 0-n
	WebPolicy         []WebPolicy         `xml:"WebPolicy,omitempty"`         // 0-n
}

// Usage represents usage types and restrictions
type Usage struct {
	XMLName xml.Name `xml:"Usage"`
	UseType []string `xml:"UseType"` // 1-n
}

// DSP represents a Digital Service Provider
type DSP struct {
	XMLName xml.Name `xml:",omitempty"`
	// DSP fields would be defined based on ddexC:DSP composite
}

// PromotionalCode represents a promotional code composite
type PromotionalCode struct {
	XMLName xml.Name `xml:"PromotionalCode"`
	// PromotionalCode fields would be defined based on ddexC:PromotionalCode composite
}

// ConsumerRentalPeriod represents the rental period for consumers
type ConsumerRentalPeriod struct {
	XMLName xml.Name `xml:"ConsumerRentalPeriod"`
	// ConsumerRentalPeriod fields would be defined based on ddexC:ConsumerRentalPeriod composite
}

// DealResourceReferenceList represents a list of resources in a deal
type DealResourceReferenceList struct {
	XMLName xml.Name `xml:",omitempty"`
	// DealResourceReferenceList fields would be defined based on ern:DealResourceReferenceList composite
}

// RelatedReleaseOfferSet represents related offers for a release
type RelatedReleaseOfferSet struct {
	XMLName xml.Name `xml:"RelatedReleaseOfferSet"`
	// RelatedReleaseOfferSet fields would be defined based on ern:RelatedReleaseOfferSet composite
}

// PhysicalReturns represents physical returns information
type PhysicalReturns struct {
	XMLName xml.Name `xml:"PhysicalReturns"`
	// PhysicalReturns fields would be defined based on ern:PhysicalReturns composite
}

// WebPolicy represents UserGeneratedContent permissions
type WebPolicy struct {
	XMLName xml.Name `xml:"WebPolicy"`
	// WebPolicy fields would be defined based on ern:WebPolicy composite
}

// PriceInformation represents pricing information for a deal
type PriceInformation struct {
	XMLName                        xml.Name `xml:"PriceInformation"`
	BulkOrderWholesalePricePerUnit float64  `xml:"BulkOrderWholesalePricePerUnit,omitempty"`
}

// ValidityPeriod represents time period validity information
type ValidityPeriod struct {
	XMLName       xml.Name `xml:"ValidityPeriod"`
	StartDate     string   `xml:"StartDate,omitempty"`
	StartDateTime string   `xml:"StartDateTime,omitempty"`
	EndDate       string   `xml:"EndDate,omitempty"`
}

// RightsClaimPolicy represents a policy for claiming rights
type RightsClaimPolicy struct {
	XMLName               xml.Name `xml:"RightsClaimPolicy"`
	RightsClaimPolicyType string   `xml:"RightsClaimPolicyType"`
}
