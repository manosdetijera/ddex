package ddex

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// Builder provides a fluent interface for creating DDEX ERN 4.3 messages
type Builder struct {
	message *NewReleaseMessage
}

// NewDDEXBuilder creates a new builder for ERN 4.3 messages
func NewDDEXBuilder() *Builder {
	return &Builder{
		message: &NewReleaseMessage{
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

	b.message.MessageHeader = &MessageHeader{
		MessageThreadId:        threadId,
		MessageId:              messageId,
		MessageSender:          sender,
		MessageCreatedDateTime: &DateTime{Time: time.Now()},
	}

	return b
}

// AddRecipient adds a message recipient (e.g., YouTube)
func (b *Builder) AddRecipient(dpid, name string) *Builder {
	if b.message.MessageHeader == nil {
		b.message.MessageHeader = &MessageHeader{}
	}

	recipient := &MessageRecipient{
		PartyId: []PartyID{
			{Value: dpid},
		},
		PartyName: []Name{
			{FullName: name},
		},
	}

	b.message.MessageHeader.MessageRecipient = append(b.message.MessageHeader.MessageRecipient, recipient)
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

	b.message.PartyList.Party = append(b.message.PartyList.Party, party)
	return b
}

// AddVideo adds a video resource
func (b *Builder) AddVideo(resourceRef, videoType, isrc string) *VideoBuilder {
	video := &Video{
		ResourceReference: resourceRef,
		Type:              videoType,
	}

	if isrc != "" {
		video.VideoEdition = append(video.VideoEdition, VideoEdition{
			ResourceId: []VideoId{
				{ISRC: isrc},
			},
		})
	}

	b.message.ResourceList.Video = append(b.message.ResourceList.Video, *video)
	videoIndex := len(b.message.ResourceList.Video) - 1

	return &VideoBuilder{
		builder: b,
		video:   &b.message.ResourceList.Video[videoIndex],
	}
}

// AddImage adds an image resource
func (b *Builder) AddImage(resourceRef, imageType string) *ImageBuilder {
	image := &Image{
		ResourceReference: resourceRef,
		Type:              imageType,
	}

	b.message.ResourceList.Image = append(b.message.ResourceList.Image, *image)
	imageIndex := len(b.message.ResourceList.Image) - 1

	return &ImageBuilder{
		builder: b,
		image:   &b.message.ResourceList.Image[imageIndex],
	}
}

// AddRelease adds a release to the release list
func (b *Builder) AddRelease(releaseRef, releaseType, icpn string) *ReleaseBuilder {
	release := &Release{
		ReleaseReference: releaseRef,
		ReleaseType:      releaseType,
	}

	if icpn != "" {
		release.ReleaseId = []ReleaseId{
			{ICPN: icpn},
		}
	}

	b.message.ReleaseList.Release = append(b.message.ReleaseList.Release, *release)
	releaseIndex := len(b.message.ReleaseList.Release) - 1

	return &ReleaseBuilder{
		builder: b,
		release: &b.message.ReleaseList.Release[releaseIndex],
	}
}

// AddDeal adds a deal to the deal list
func (b *Builder) AddDeal(releaseRef string) *DealBuilder {
	deal := &ReleaseDeal{
		DealReleaseReference: releaseRef,
		Deal:                 []Deal{{}},
	}

	b.message.DealList.ReleaseDeal = append(b.message.DealList.ReleaseDeal, *deal)
	dealIndex := len(b.message.DealList.ReleaseDeal) - 1

	return &DealBuilder{
		builder:     b,
		releaseDeal: &b.message.DealList.ReleaseDeal[dealIndex],
		deal:        &b.message.DealList.ReleaseDeal[dealIndex].Deal[0],
	}
}

// Build returns the completed NewReleaseMessage
func (b *Builder) Build() *NewReleaseMessage {
	return b.message
}

// ToXML converts the message to XML bytes
func (b *Builder) ToXML() ([]byte, error) {
	return xml.MarshalIndent(b.message, "", "    ")
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
	vb.video.DisplayTitleText = DisplayTitleText{
		Value:                   title,
		ApplicableTerritoryCode: "Worldwide",
		LanguageAndScriptCode:   "en",
		IsDefault:               true,
	}

	displayTitle := DisplayTitle{
		TitleText: []TitleText{{Value: title}},
	}
	if subtitle != "" {
		displayTitle.TitleText = append(displayTitle.TitleText, TitleText{Value: subtitle, TitleType: "SubTitle"})
	}
	vb.video.DisplayTitle = displayTitle

	return vb
}

// WithArtist sets the video artist
func (vb *VideoBuilder) WithArtist(artistName, partyRef string, sequence int) *VideoBuilder {
	vb.video.DisplayArtistName = []string{artistName}

	if partyRef != "" {
		vb.video.DisplayArtist = append(vb.video.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    "MainArtist",
			SequenceNumber:       sequence,
		})
	}

	return vb
}

// WithRightsController sets the rights controller
func (vb *VideoBuilder) WithRightsController(partyRef string, percentage float64) *VideoBuilder {
	vb.video.ResourceRightsController = append(vb.video.ResourceRightsController, ResourceRightsController{
		RightsControllerPartyReference: partyRef,
		RightsControlType:              "RightsController",
		RightSharePercentage:           fmt.Sprintf("%.2f", percentage),
		DelegatedUsageRights: []DelegatedUsageRights{
			{
				UseType:                     []string{"UserMakeAvailableUserProvided"},
				TerritoryOfRightsDelegation: []string{"Worldwide"},
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
	vb.video.CreationDate = CreationDate{
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

// WithArtist sets the release artist
func (rb *ReleaseBuilder) WithArtist(artistName, partyRef string, sequence int) *ReleaseBuilder {
	rb.release.DisplayArtistName = []string{artistName}

	if partyRef != "" {
		rb.release.DisplayArtist = append(rb.release.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    "MainArtist",
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

// WithGenre adds genre information
func (rb *ReleaseBuilder) WithGenre(genreText, territoryCode string) *ReleaseBuilder {
	rb.release.Genre = append(rb.release.Genre, Genre{
		GenreText:               genreText,
		ApplicableTerritoryCode: territoryCode,
	})
	return rb
}

// WithParentalWarning sets the parental warning type
func (rb *ReleaseBuilder) WithParentalWarning(warningType string) *ReleaseBuilder {
	rb.release.ParentalWarningType = warningType
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

// DealBuilder provides fluent interface for building deals
type DealBuilder struct {
	builder     *Builder
	releaseDeal *ReleaseDeal
	deal        *Deal
}

// WithTerritory sets the deal territory
func (db *DealBuilder) WithTerritory(territoryCode string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.TerritoryCode = territoryCode
	return db
}

// WithValidityPeriod sets the deal validity period
func (db *DealBuilder) WithValidityPeriod(startDate string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	db.deal.DealTerms.ValidityPeriod = &ValidityPeriod{
		StartDate: startDate,
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

// Done returns to the main builder
func (db *DealBuilder) Done() *Builder {
	return db.builder
}
