package ddex

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// Builder provides a fluent interface for creating DDEX ERN 4.3 messages
type Builder struct {
	Message *NewReleaseMessage
}

// NewDDEXBuilder creates a new builder for ERN 4.3 messages
func NewDDEXBuilder() *Builder {
	return &Builder{
		Message: &NewReleaseMessage{
			XmlnsErn:                XmlnsErn,
			XmlnsXsi:                XmlnsXsi,
			XsiSchemaLocation:       XsiSchemaLocation,
			ReleaseProfileVersionId: "Video",
			LanguageAndScriptCode:   "en",
			AvsVersionId:            "3",
			PartyList:               &PartyList{},
			ResourceList:            &ResourceList{},
			ReleaseList:             &ReleaseList{},
			DealList:                &DealList{},
		},
	}
}

// WithMessageHeader sets the message header
func (b *Builder) WithMessageHeader(messageId, threadId, senderDPID, senderName string) *Builder {
	sender := &MessageSender{
		PartyId: []PartyID{
			{Value: senderDPID},
		},
		PartyName: []Name{
			{FullName: senderName},
		},
	}

	b.Message.MessageHeader = &MessageHeader{
		MessageThreadId:        threadId,
		MessageId:              messageId,
		MessageSender:          sender,
		MessageCreatedDateTime: &DateTime{Time: time.Now()},
	}

	return b
}

// AddRecipient adds a message recipient (e.g., YouTube)
func (b *Builder) AddRecipient(dpid, name string) *Builder {
	if b.Message.MessageHeader == nil {
		b.Message.MessageHeader = &MessageHeader{}
	}

	recipient := &MessageRecipient{
		PartyId: []PartyID{
			{Value: dpid},
		},
		PartyName: []Name{
			{FullName: name},
		},
	}

	b.Message.MessageHeader.MessageRecipient = append(b.Message.MessageHeader.MessageRecipient, recipient)
	return b
}

// AddYouTubeRecipient adds YouTube as the message recipient
func (b *Builder) AddYouTubeRecipient() *Builder {
	return b.AddRecipient("PADPIDA2013020802I", "YouTube")
}

// AddYouTubeRecipient adds YouTube as the message recipient
func (b *Builder) AddYouTubeContentIDRecipient() *Builder {
	return b.AddRecipient("PADPIDA2015120100H", "YouTube_ContentID")
}

// AddParty adds a party (artist, label, etc.) to the party list
func (b *Builder) AddParty(reference, fullName, fullNameIndexed string) *Builder {
	var party Party
	if fullNameIndexed != "" {
		party = *NewPartyWithIndexedName(reference, fullName, fullNameIndexed)
	} else {
		party = *NewParty(reference, fullName)
	}

	b.Message.PartyList.Party = append(b.Message.PartyList.Party, party)
	return b
}

// AddVideo adds a video resource
func (b *Builder) AddVideo(resourceRef, videoType string) *VideoBuilder {
	video := &Video{
		ResourceReference: resourceRef,
		Type:              videoType,
	}

	b.Message.ResourceList.Video = append(b.Message.ResourceList.Video, *video)
	videoIndex := len(b.Message.ResourceList.Video) - 1

	return &VideoBuilder{
		builder: b,
		video:   &b.Message.ResourceList.Video[videoIndex],
	}
}

// AddImage adds an image resource
func (b *Builder) AddImage(resourceRef, imageType string) *ImageBuilder {
	image := &Image{
		ResourceReference: resourceRef,
		Type:              imageType,
	}

	b.Message.ResourceList.Image = append(b.Message.ResourceList.Image, *image)
	imageIndex := len(b.Message.ResourceList.Image) - 1

	return &ImageBuilder{
		builder: b,
		image:   &b.Message.ResourceList.Image[imageIndex],
	}
}

// AddRelease adds a release to the release list
func (b *Builder) AddRelease(releaseRef, releaseType string) *ReleaseBuilder {
	release := &Release{
		ReleaseReference: releaseRef,
		ReleaseType:      releaseType,
	}

	b.Message.ReleaseList.Release = append(b.Message.ReleaseList.Release, *release)
	releaseIndex := len(b.Message.ReleaseList.Release) - 1

	return &ReleaseBuilder{
		builder: b,
		release: &b.Message.ReleaseList.Release[releaseIndex],
	}
}

// AddDeal adds a deal to the deal list
// AddReleaseDeal adds a release deal to the deal list
func (b *Builder) AddReleaseDeal(releaseRef string) *ReleaseDealBuilder {
	releaseDeal := &ReleaseDeal{
		DealReleaseReference: releaseRef,
		Deal:                 []Deal{},
	}

	b.Message.DealList.ReleaseDeal = append(b.Message.DealList.ReleaseDeal, *releaseDeal)
	dealIndex := len(b.Message.DealList.ReleaseDeal) - 1

	return &ReleaseDealBuilder{
		builder:     b,
		releaseDeal: &b.Message.DealList.ReleaseDeal[dealIndex],
	}
}

// Build returns the completed NewReleaseMessage
func (b *Builder) Build() *NewReleaseMessage {
	return b.Message
}

// ToXML converts the message to XML bytes
func (b *Builder) ToXML() ([]byte, error) {
	return xml.MarshalIndent(b.Message, "", "    ")
}

// WriteToFile writes the message to an XML file
func (b *Builder) WriteToFile(filename string) error {
	xmlData, err := b.ToXML()
	if err != nil {
		return fmt.Errorf("failed to marshal XML: %w", err)
	}

	// Add XML declaration
	xmlWithDeclaration := []byte(xml.Header + string(xmlData))

	if err := os.WriteFile(filename, xmlWithDeclaration, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// VideoBuilder provides fluent interface for building video resources
type VideoBuilder struct {
	builder *Builder
	video   *Video
}

// WithTitle sets the video title
func (vb *VideoBuilder) WithTitle(title, subtitle string) *VideoBuilder {
	vb.video.DisplayTitleText = &DisplayTitleText{
		Value:                   title,
		ApplicableTerritoryCode: "Worldwide",
		LanguageAndScriptCode:   "en",
		IsDefault:               true,
	}

	displayTitle := &DisplayTitle{
		TitleText: []TitleText{{Value: title}},
	}
	if subtitle != "" {
		displayTitle.TitleText = append(displayTitle.TitleText, TitleText{Value: subtitle, TitleType: "SubTitle"})
	}
	vb.video.DisplayTitle = displayTitle

	return vb
}

// WithDisplayArtistName sets the display artist name for the video
func (vb *VideoBuilder) WithDisplayArtistName(artistName, territoryCode string) *VideoBuilder {
	vb.video.DisplayArtistName = append(vb.video.DisplayArtistName, DisplayArtistNameWithOriginalLanguage{
		Value:                   artistName,
		ApplicableTerritoryCode: territoryCode,
	})
	return vb
}

// WithArtist adds a display artist reference to the video
func (vb *VideoBuilder) WithArtist(partyRef, role string, sequence int) *VideoBuilder {
	if partyRef != "" {
		vb.video.DisplayArtist = append(vb.video.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    role,
			SequenceNumber:       sequence,
		})
	}

	return vb
}

// WithContributor adds a contributor to the video resource
// role can be multiple values like "Producer", "Director", "Cinematographer", etc.
func (vb *VideoBuilder) WithContributor(partyRef string, roles []string, sequence int) *VideoBuilder {
	if partyRef != "" && len(roles) > 0 {
		vb.video.Contributor = append(vb.video.Contributor, Contributor{
			PartyReference: partyRef,
			Role:           roles,
			SequenceNumber: sequence,
		})
	}

	return vb
}

// WithRightsController sets the rights controller
func (vb *VideoBuilder) WithRightsController(partyRef string, percentage float64, territories []string) *VideoBuilder {
	// If no territories provided, default to Worldwide
	if len(territories) == 0 {
		territories = []string{"Worldwide"}
	}

	vb.video.ResourceRightsController = append(vb.video.ResourceRightsController, ResourceRightsController{
		RightsControllerPartyReference: partyRef,
		RightsControlType:              "RightsController",
		RightSharePercentage:           fmt.Sprintf("%.2f", percentage),
		DelegatedUsageRights: []DelegatedUsageRights{
			{
				UseType:                     []string{"UserMakeAvailableUserProvided"},
				TerritoryOfRightsDelegation: territories,
			},
		},
	})

	return vb
}

// WithDuration sets the video duration (e.g., "PT3M10S")
func (vb *VideoBuilder) WithDuration(duration string) *VideoBuilder {
	vb.video.Duration = duration
	return vb
}

// WithCreationDate sets the creation date
func (vb *VideoBuilder) WithCreationDate(date string, isApproximate bool) *VideoBuilder {
	vb.video.CreationDate = &CreationDate{
		Value:         date,
		IsApproximate: isApproximate,
	}
	return vb
}

// WithParentalWarning sets the parental warning type
func (vb *VideoBuilder) WithParentalWarning(warningType string) *VideoBuilder {
	vb.video.ParentalWarningType = warningType
	return vb
}

// WithPLine sets the P-Line information
func (vb *VideoBuilder) WithPLine(year int, text string) *VideoBuilder {
	if len(vb.video.VideoEdition) == 0 {
		vb.video.VideoEdition = append(vb.video.VideoEdition, VideoEdition{})
	}

	vb.video.VideoEdition[0].PLine = []PLine{
		{Year: year, PLineText: text},
	}

	return vb
}

// WithTechnicalDetails adds technical details and file URI
func (vb *VideoBuilder) WithTechnicalDetails(techRef, fileURI string) *VideoBuilder {
	if len(vb.video.VideoEdition) == 0 {
		vb.video.VideoEdition = append(vb.video.VideoEdition, VideoEdition{})
	}

	vb.video.VideoEdition[0].TechnicalDetails = []TechnicalVideoDetails{
		{
			TechnicalResourceDetailsReference: techRef,
			DeliveryFile: []AudioVisualDeliveryFile{
				{
					Type: "AudioVisualFile",
					File: File{
						URI: fileURI,
					},
				},
			},
		},
	}

	return vb
}

// WithISRC sets the ISRC for the video
func (vb *VideoBuilder) WithISRC(isrc string) *VideoBuilder {
	if len(vb.video.VideoEdition) == 0 {
		vb.video.VideoEdition = append(vb.video.VideoEdition, VideoEdition{})
	}

	vb.video.VideoEdition[0].ResourceId = append(vb.video.VideoEdition[0].ResourceId, VideoId{
		ISRC: isrc,
	})

	return vb
}

// AddKeywords adds keywords for enhanced search and display
func (vb *VideoBuilder) AddKeywords(keywords ...string) *VideoBuilder {
	for _, keyword := range keywords {
		vb.video.Keywords = append(vb.video.Keywords, Keywords{
			Value:                   keyword,
			ApplicableTerritoryCode: "Worldwide",
		})
	}
	return vb
}

// AddKeywordsWithTerritory adds keywords with specific territory and language
func (vb *VideoBuilder) AddKeywordsWithTerritory(territoryCode, languageCode string, keywords ...string) *VideoBuilder {
	for _, keyword := range keywords {
		vb.video.Keywords = append(vb.video.Keywords, Keywords{
			Value:                   keyword,
			ApplicableTerritoryCode: territoryCode,
			LanguageAndScriptCode:   languageCode,
		})
	}
	return vb
}

// AddProprietaryId adds a proprietary ID (e.g., YouTube channel ID)
func (vb *VideoBuilder) AddProprietaryId(namespace, value string) *VideoBuilder {
	if len(vb.video.VideoEdition) == 0 {
		vb.video.VideoEdition = append(vb.video.VideoEdition, VideoEdition{})
	}

	vb.video.VideoEdition[0].ResourceId = append(vb.video.VideoEdition[0].ResourceId, VideoId{
		ProprietaryId: []ProprietaryId{
			{Namespace: namespace, Value: value},
		},
	})

	return vb
}

// Done returns to the main builder
func (vb *VideoBuilder) Done() *Builder {
	return vb.builder
}

// ImageBuilder provides fluent interface for building image resources
type ImageBuilder struct {
	builder *Builder
	image   *Image
}

// WithProprietaryId adds a proprietary ID to the image
func (ib *ImageBuilder) WithProprietaryId(namespace, value string) *ImageBuilder {
	ib.image.ResourceId = []ImageId{
		{
			ProprietaryId: []ProprietaryId{
				{Namespace: namespace, Value: value},
			},
		},
	}
	return ib
}

// WithParentalWarning sets the parental warning type
func (ib *ImageBuilder) WithParentalWarning(warningType string) *ImageBuilder {
	ib.image.ParentalWarningType = warningType
	return ib
}

// WithTechnicalDetails adds technical details and file URI
func (ib *ImageBuilder) WithTechnicalDetails(techRef, fileURI string) *ImageBuilder {
	ib.image.TechnicalDetails = []TechnicalImageDetails{
		{
			TechnicalResourceDetailsReference: techRef,
			File: File{
				URI: fileURI,
			},
		},
	}
	return ib
}

// Done returns to the main builder
func (ib *ImageBuilder) Done() *Builder {
	return ib.builder
}

// ReleaseBuilder provides fluent interface for building releases
type ReleaseBuilder struct {
	builder *Builder
	release *Release
}

// WithTitle sets the release title
func (rb *ReleaseBuilder) WithTitle(title, subtitle string) *ReleaseBuilder {
	rb.release.DisplayTitleText = []DisplayTitleText{
		{
			Value:                   title,
			ApplicableTerritoryCode: "Worldwide",
			LanguageAndScriptCode:   "en",
			IsDefault:               true,
		},
	}

	displayTitle := DisplayTitle{
		TitleText: []TitleText{{Value: title}},
	}
	if subtitle != "" {
		displayTitle.TitleText = append(displayTitle.TitleText, TitleText{Value: subtitle, TitleType: "SubTitle"})
	}
	rb.release.DisplayTitle = append(rb.release.DisplayTitle, displayTitle)

	return rb
}

// WithDisplayArtistName sets the display artist name for the release
func (rb *ReleaseBuilder) WithDisplayArtistName(artistName, territoryCode string) *ReleaseBuilder {
	rb.release.DisplayArtistName = append(rb.release.DisplayArtistName, DisplayArtistNameWithOriginalLanguage{
		Value:                   artistName,
		ApplicableTerritoryCode: territoryCode,
	})
	return rb
}

// WithArtist adds a display artist reference to the release
func (rb *ReleaseBuilder) WithArtist(partyRef, role string, sequence int) *ReleaseBuilder {
	if partyRef != "" {
		rb.release.DisplayArtist = append(rb.release.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    role,
			SequenceNumber:       sequence,
		})
	}

	return rb
}

// WithLabel sets the release label
func (rb *ReleaseBuilder) WithLabel(labelPartyRef, territoryCode string) *ReleaseBuilder {
	rb.release.ReleaseLabelReference = append(rb.release.ReleaseLabelReference, ReleaseLabelReference{
		Value:                   labelPartyRef,
		ApplicableTerritoryCode: territoryCode,
	})
	return rb
}

// WithPLine adds P-Line information
func (rb *ReleaseBuilder) WithPLine(year int, text string) *ReleaseBuilder {
	rb.release.PLine = append(rb.release.PLine, PLine{
		Year:      year,
		PLineText: text,
	})
	return rb
}

// WithCLine adds C-Line information
func (rb *ReleaseBuilder) WithCLine(year int, text string) *ReleaseBuilder {
	rb.release.CLine = append(rb.release.CLine, CLine{
		Year:      year,
		CLineText: text,
	})
	return rb
}

// WithDuration sets the release duration
func (rb *ReleaseBuilder) WithDuration(duration string) *ReleaseBuilder {
	rb.release.Duration = duration
	return rb
}

// WithReleaseDate sets both ReleaseDate and OriginalReleaseDate to the same date value
// Date should be in ISO 8601 format (YYYY, YYYY-MM, or YYYY-MM-DD)
func (rb *ReleaseBuilder) WithReleaseDate(date string) *ReleaseBuilder {
	releaseDateEntry := EventDateWithDefault{
		XMLName: xml.Name{Local: "ReleaseDate"},
		Value:   date,
	}
	originalReleaseDateEntry := EventDateWithDefault{
		XMLName: xml.Name{Local: "OriginalReleaseDate"},
		Value:   date,
	}

	rb.release.ReleaseDate = append(rb.release.ReleaseDate, releaseDateEntry)
	rb.release.OriginalReleaseDate = append(rb.release.OriginalReleaseDate, originalReleaseDateEntry)
	return rb
}

// WithGenre adds genre information
func (rb *ReleaseBuilder) WithGenre(genreText, territoryCode string) *ReleaseBuilder {
	rb.release.DisplayGenre = append(rb.release.DisplayGenre, DisplayGenre{
		GenreText:               genreText,
		ApplicableTerritoryCode: territoryCode,
	})
	return rb
}

// WithGenreAndSubGenre adds genre information with a subgenre
func (rb *ReleaseBuilder) WithGenreAndSubGenre(genreText, subGenre, territoryCode string) *ReleaseBuilder {
	rb.release.DisplayGenre = append(rb.release.DisplayGenre, DisplayGenre{
		GenreText:               genreText,
		SubGenre:                subGenre,
		ApplicableTerritoryCode: territoryCode,
	})
	return rb
}

// WithParentalWarning sets the parental warning type
func (rb *ReleaseBuilder) WithParentalWarning(warningType string) *ReleaseBuilder {
	rb.release.ParentalWarningType = warningType
	return rb
}

// WithAvRating adds an AvRating to the release
// Common use case for YouTube: WithAvRating("MadeForKids", "UserDefined", "YOUTUBE")
func (rb *ReleaseBuilder) WithAvRating(ratingText, agencyValue, agencyNamespace string) *ReleaseBuilder {
	avRating := AvRating{
		RatingText: ratingText,
		RatingAgency: RatingAgency{
			Value:     agencyValue,
			Namespace: agencyNamespace,
		},
	}

	rb.release.AvRating = append(rb.release.AvRating, avRating)
	return rb
}

// WithMadeForKids is a convenience method to set the YouTube MadeForKids rating
func (rb *ReleaseBuilder) WithMadeForKids() *ReleaseBuilder {
	return rb.WithAvRating("MadeForKids", "UserDefined", "YOUTUBE")
}

// WithMarketingComment adds a marketing comment to the release
func (rb *ReleaseBuilder) WithMarketingComment(comment, territoryCode, languageCode string) *ReleaseBuilder {
	commentEntry := MarketingComment{
		Value:                   comment,
		ApplicableTerritoryCode: territoryCode,
		LanguageAndScriptCode:   languageCode,
	}

	rb.release.MarketingComment = append(rb.release.MarketingComment, commentEntry)
	return rb
}

// WithMarketingCommentSimple adds a marketing comment to the release with worldwide territory
func (rb *ReleaseBuilder) WithMarketingCommentSimple(comment string) *ReleaseBuilder {
	return rb.WithMarketingComment(comment, "Worldwide", "")
}

// WithKeywords adds keywords to the release
func (rb *ReleaseBuilder) WithKeywords(keywords, territoryCode, languageCode string) *ReleaseBuilder {
	keywordsEntry := Keywords{
		Value:                   keywords,
		ApplicableTerritoryCode: territoryCode,
		LanguageAndScriptCode:   languageCode,
	}

	rb.release.Keywords = append(rb.release.Keywords, keywordsEntry)
	return rb
}

// WithKeywordsSimple adds keywords to the release with worldwide territory
func (rb *ReleaseBuilder) WithKeywordsSimple(keywords string) *ReleaseBuilder {
	return rb.WithKeywords(keywords, "Worldwide", "")
}

// WithContainsAI sets the AI contribution type for the release
// Valid values typically include: "AIGenerated", "AIAssisted", "AITraining", "NoAI"
func (rb *ReleaseBuilder) WithContainsAI(containsAI string) *ReleaseBuilder {
	rb.release.ContainsAI = containsAI
	return rb
}

// WithUPC sets the UPC identifier for the release
func (rb *ReleaseBuilder) WithUPC(upc string) *ReleaseBuilder {
	rb.release.ReleaseId = append(rb.release.ReleaseId, ReleaseId{
		ICPN: &ICPN{
			Value: upc,
			IsEan: false,
		},
	})
	return rb
}

// WithEAN sets the EAN identifier for the release
func (rb *ReleaseBuilder) WithEAN(ean string) *ReleaseBuilder {
	rb.release.ReleaseId = append(rb.release.ReleaseId, ReleaseId{
		ICPN: &ICPN{
			Value: ean,
			IsEan: true,
		},
	})
	return rb
}

// WithICPN sets a generic ICPN identifier (use WithUPC or WithEAN for better clarity)
func (rb *ReleaseBuilder) WithICPN(icpn string) *ReleaseBuilder {
	rb.release.ReleaseId = append(rb.release.ReleaseId, ReleaseId{
		ICPN: &ICPN{
			Value: icpn,
			IsEan: false, // Default to UPC format
		},
	})
	return rb
}

// WithGRid sets the GRid identifier for the release
func (rb *ReleaseBuilder) WithGRid(grid string) *ReleaseBuilder {
	rb.release.ReleaseId = append(rb.release.ReleaseId, ReleaseId{
		GRid: grid,
	})
	return rb
}

// AddProprietaryId adds a proprietary identifier to the release ID
func (rb *ReleaseBuilder) AddProprietaryId(namespace, value string) *ReleaseBuilder {
	// Find or create the first ReleaseId entry
	if len(rb.release.ReleaseId) == 0 {
		rb.release.ReleaseId = append(rb.release.ReleaseId, ReleaseId{})
	}

	// Add the ProprietaryId to the first ReleaseId
	rb.release.ReleaseId[0].ProprietaryId = append(rb.release.ReleaseId[0].ProprietaryId, ProprietaryId{
		Namespace: namespace,
		Value:     value,
	})
	return rb
}

// AddRelatedResource adds a related resource to the release
func (rb *ReleaseBuilder) AddRelatedResource(relationshipType, isrc string) *ReleaseBuilder {
	rb.release.RelatedResource = append(rb.release.RelatedResource, RelatedResource{
		ResourceRelationshipType: relationshipType,
		ResourceId: RelatedResourceId{
			ISRC: isrc,
		},
	})
	return rb
}

// AddResourceGroup adds a resource group to the release
func (rb *ReleaseBuilder) AddResourceGroup(titleText string, sequenceNumber int) *ResourceGroupBuilder {
	group := ResourceGroup{
		SequenceNumber: sequenceNumber,
	}

	if titleText != "" {
		group.AdditionalTitle = AdditionalTitle{
			TitleText: titleText,
		}
	}

	rb.release.ResourceGroup = append(rb.release.ResourceGroup, group)
	groupIndex := len(rb.release.ResourceGroup) - 1

	return &ResourceGroupBuilder{
		releaseBuilder: rb,
		group:          &rb.release.ResourceGroup[groupIndex],
	}
}

// Done returns to the main builder
func (rb *ReleaseBuilder) Done() *Builder {
	return rb.builder
}

// ResourceGroupBuilder provides fluent interface for building resource groups
type ResourceGroupBuilder struct {
	releaseBuilder *ReleaseBuilder
	group          *ResourceGroup
}

// AddContentItem adds a content item to the resource group
func (rgb *ResourceGroupBuilder) AddContentItem(sequenceNumber int, resourceRef string) *ResourceGroupBuilder {
	item := ResourceGroupContentItem{
		SequenceNumber:           sequenceNumber,
		ReleaseResourceReference: resourceRef,
	}

	rgb.group.ResourceGroupContentItem = append(rgb.group.ResourceGroupContentItem, item)
	return rgb
}

// AddLinkedResource adds a linked resource (e.g., cover art)
func (rgb *ResourceGroupBuilder) AddLinkedResource(linkDescription, resourceRef string) *ResourceGroupBuilder {
	if len(rgb.group.ResourceGroupContentItem) > 0 {
		lastIndex := len(rgb.group.ResourceGroupContentItem) - 1
		rgb.group.ResourceGroupContentItem[lastIndex].LinkedReleaseResourceReference = append(
			rgb.group.ResourceGroupContentItem[lastIndex].LinkedReleaseResourceReference,
			LinkedReleaseResourceReference{
				LinkDescription: linkDescription,
				Value:           resourceRef,
			},
		)
	}
	return rgb
}

// Done returns to the release builder
func (rgb *ResourceGroupBuilder) Done() *ReleaseBuilder {
	return rgb.releaseBuilder
}

// ReleaseDealBuilder provides fluent interface for building release deals
type ReleaseDealBuilder struct {
	builder     *Builder
	releaseDeal *ReleaseDeal
}

// AddDeal adds a new deal to the release deal
func (rdb *ReleaseDealBuilder) AddDeal() *DealBuilder {
	newDeal := Deal{}
	rdb.releaseDeal.Deal = append(rdb.releaseDeal.Deal, newDeal)
	dealIndex := len(rdb.releaseDeal.Deal) - 1

	return &DealBuilder{
		builder:            rdb.builder,
		releaseDealBuilder: rdb,
		deal:               &rdb.releaseDeal.Deal[dealIndex],
	}
}

// Done returns to the main builder
func (rdb *ReleaseDealBuilder) Done() *Builder {
	return rdb.builder
}

// DealBuilder provides fluent interface for building deals
type DealBuilder struct {
	builder            *Builder
	releaseDealBuilder *ReleaseDealBuilder
	deal               *Deal
}

// WithTerritory sets the deal territory
// WithTerritories sets the deal territories
func (db *DealBuilder) WithTerritories(territoryCodes []string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.TerritoryCode = territoryCodes
	return db
}

// WithValidityPeriod sets the deal validity period with a start date (YYYY-MM-DD)
func (db *DealBuilder) WithValidityPeriod(startDate string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	db.deal.DealTerms.ValidityPeriod = &ValidityPeriod{
		StartDate: startDate,
	}

	return db
}

// WithValidityPeriodDateTime sets the deal validity period with a start date-time (YYYY-MM-DDTHH:MM:SS)
func (db *DealBuilder) WithValidityPeriodDateTime(startDateTime string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	db.deal.DealTerms.ValidityPeriod = &ValidityPeriod{
		StartDateTime: startDateTime,
	}

	return db
}

// AddCommercialModel adds a commercial model type
func (db *DealBuilder) AddCommercialModel(modelType string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.CommercialModelType = append(db.deal.DealTerms.CommercialModelType, modelType)
	return db
}

// AddUseType adds a use type
func (db *DealBuilder) AddUseType(useType string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.UseType = append(db.deal.DealTerms.UseType, useType)
	return db
}

// WithRightsClaimPolicy sets the rights claim policy for the deal
func (db *DealBuilder) WithRightsClaimPolicy(policyType string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.RightsClaimPolicy = &RightsClaimPolicy{
		RightsClaimPolicyType: policyType,
	}
	return db
}

// Done returns to the release deal builder
func (db *DealBuilder) Done() *ReleaseDealBuilder {
	return db.releaseDealBuilder
}
