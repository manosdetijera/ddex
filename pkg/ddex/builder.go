package ddex

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// Builder provides a fluent interface for creating DDEX ERN 3.8 messages
type Builder struct {
	Message *NewReleaseMessage
}

// NewDDEXBuilder creates a new builder for ERN 3.8 messages
func NewDDEXBuilder() *Builder {
	return &Builder{
		Message: &NewReleaseMessage{
			XmlnsErn:               XmlnsErn,
			XmlnsXsi:               XmlnsXsi,
			XsiSchemaLocation:      XsiSchemaLocation,
			MessageSchemaVersionId: MessageSchemaVersionId,
			LanguageAndScriptCode:  "en",
			ResourceList:           &ResourceList{},
			ReleaseList:            &ReleaseList{},
			DealList:               &DealList{},
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

// WithUpdateIndicator sets the update indicator
// Valid values: "OriginalMessage" or "UpdateMessage"
// Note: This element is deprecated in ERN 3.8
func (b *Builder) WithUpdateIndicator(indicator string) *Builder {
	b.Message.UpdateIndicator = indicator
	return b
}

// AddVideo adds a video resource
func (b *Builder) AddVideo(resourceRef, videoType string) *VideoBuilder {
	video := &Video{
		ResourceReference: resourceRef,
		VideoType:         &VideoType{Value: videoType},
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
	}

	if imageType != "" {
		image.ImageType = &ImageType{Value: imageType}
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
	}

	if releaseType != "" {
		release.ReleaseType = []ReleaseType{{Value: releaseType}}
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
	builder                 *Builder
	video                   *Video
	currentTerritoryDetails *VideoDetailsByTerritory
	currentTerritoryIndex   int
}

// VideoDetailsByTerritoryBuilder provides fluent interface for building video territory details
type VideoDetailsByTerritoryBuilder struct {
	videoBuilder     *VideoBuilder
	territoryDetails *VideoDetailsByTerritory
}

// AddVideoDetailsByTerritory creates a new territory details section and returns a builder for it
func (vb *VideoBuilder) AddVideoDetailsByTerritory(territoryCodes []string) *VideoDetailsByTerritoryBuilder {
	// Validate that at least one territory code is provided
	if len(territoryCodes) == 0 {
		territoryCodes = []string{"Worldwide"}
	}

	// Create new territory details
	newDetails := VideoDetailsByTerritory{
		TerritoryCode: territoryCodes,
	}
	vb.video.VideoDetailsByTerritory = append(vb.video.VideoDetailsByTerritory, newDetails)
	vb.currentTerritoryIndex = len(vb.video.VideoDetailsByTerritory) - 1
	vb.currentTerritoryDetails = &vb.video.VideoDetailsByTerritory[vb.currentTerritoryIndex]

	return &VideoDetailsByTerritoryBuilder{
		videoBuilder:     vb,
		territoryDetails: vb.currentTerritoryDetails,
	}
}

// Done returns to the video builder
func (vtb *VideoDetailsByTerritoryBuilder) Done() *VideoBuilder {
	return vtb.videoBuilder
}

// WithTitle sets the video title (goes to territory details in ERN 3.8)
func (vtb *VideoDetailsByTerritoryBuilder) WithTitle(title, subtitle, titleType string) *VideoDetailsByTerritoryBuilder {
	// Add title to territory details
	titleStruct := Title{
		TitleText: title,
		TitleType: titleType,
	}
	if subtitle != "" {
		titleStruct.SubTitle = subtitle
	}
	vtb.territoryDetails.Title = append(vtb.territoryDetails.Title, titleStruct)

	return vtb
}

// WithDisplayArtistName sets the display artist name for the video (ERN 3.8 - territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithDisplayArtistName(artistName, languageCode string) *VideoDetailsByTerritoryBuilder {
	if languageCode == "" {
		languageCode = "en"
	}
	vtb.territoryDetails.DisplayArtistName = append(vtb.territoryDetails.DisplayArtistName, DisplayArtistName{
		Value:                 artistName,
		LanguageAndScriptCode: languageCode,
	})
	return vtb
}

// WithArtist adds a display artist reference to the video (territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithArtist(partyRef, role string, sequence int) *VideoDetailsByTerritoryBuilder {
	if partyRef != "" {
		vtb.territoryDetails.DisplayArtist = append(vtb.territoryDetails.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    role,
			SequenceNumber:       sequence,
		})
	}

	return vtb
}

// WithLabel adds a label name for the video (territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithLabel(labelName, labelNameType, languageCode string) *VideoDetailsByTerritoryBuilder {
	if languageCode == "" {
		languageCode = "en"
	}
	vtb.territoryDetails.LabelName = append(vtb.territoryDetails.LabelName, LabelName{
		Value:                 labelName,
		LabelNameType:         labelNameType,
		LanguageAndScriptCode: languageCode,
	})
	return vtb
}

// WithContributor adds a contributor to the video resource (territory specific)
// role can be multiple values like "Producer", "Director", "Cinematographer", etc.
func (vtb *VideoDetailsByTerritoryBuilder) WithContributor(partyRef string, roles []string, sequence int) *VideoDetailsByTerritoryBuilder {
	if partyRef != "" && len(roles) > 0 {
		vtb.territoryDetails.ResourceContributor = append(vtb.territoryDetails.ResourceContributor, ResourceContributor{
			PartyReference: partyRef,
			Role:           roles,
		})
	}

	return vtb
}

// WithRightsController sets the rights controller (territory specific)
// Parameters: partyName, partyId, and percentage
func (vtb *VideoDetailsByTerritoryBuilder) WithRightsController(partyName, partyId string, percentage float64) *VideoDetailsByTerritoryBuilder {
	rightsController := RightsController{
		PartyName: []Name{
			{FullName: partyName},
		},
		PartyId: []PartyID{
			{Value: partyId},
		},
		RightsControllerRole: []string{"RightsController"},
		RightSharePercentage: fmt.Sprintf("%.2f", percentage),
	}

	vtb.territoryDetails.RightsController = append(vtb.territoryDetails.RightsController, rightsController)

	return vtb
}

// WithDuration sets the video duration (e.g., "PT3M10S") - at video level, not territory
func (vb *VideoBuilder) WithDuration(duration string) *VideoBuilder {
	vb.video.Duration = duration
	return vb
}

// WithCreationDate sets the creation date - at video level, not territory
func (vb *VideoBuilder) WithCreationDate(date string, isApproximate bool) *VideoBuilder {
	vb.video.CreationDate = &EventDate{
		Value:         date,
		IsApproximate: isApproximate,
	}
	return vb
}

// WithReferenceTitle sets the reference title for the video - at video level, not territory
func (vb *VideoBuilder) WithReferenceTitle(titleText, subtitle string) *VideoBuilder {
	vb.video.ReferenceTitle = &ReferenceTitle{
		TitleText: titleText,
		SubTitle:  subtitle,
	}
	return vb
}

// WithParentalWarning sets the parental warning type (territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithParentalWarning(warningType string) *VideoDetailsByTerritoryBuilder {
	vtb.territoryDetails.ParentalWarningType = append(vtb.territoryDetails.ParentalWarningType, warningType)
	return vtb
}

// WithPLine sets the P-Line information for ERN 3.8 (territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithPLine(year int, text string) *VideoDetailsByTerritoryBuilder {
	vtb.territoryDetails.PLine = append(vtb.territoryDetails.PLine, PLine{
		Year:      year,
		PLineText: text,
	})
	return vtb
}

// WithCLine sets the C-Line information for ERN 3.8 (territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithCLine(year int, text string) *VideoDetailsByTerritoryBuilder {
	vtb.territoryDetails.CLine = append(vtb.territoryDetails.CLine, CLine{
		Year:      year,
		CLineText: text,
	})
	return vtb
}

// WithGenre adds genre information (territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithGenre(genreText string) *VideoDetailsByTerritoryBuilder {
	genre := Genre{
		GenreText: genreText,
	}
	vtb.territoryDetails.Genre = append(vtb.territoryDetails.Genre, genre)
	return vtb
}

// WithTechnicalDetails adds technical details and file URI for ERN 3.8 (territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) WithTechnicalDetails(techRef, fileURI string) *VideoDetailsByTerritoryBuilder {
	vtb.territoryDetails.TechnicalVideoDetails = append(vtb.territoryDetails.TechnicalVideoDetails, TechnicalVideoDetails{
		TechnicalResourceDetailsReference: techRef,
		File: &File{
			URI: fileURI,
		},
	})
	return vtb
}

// WithISRC sets the ISRC for the video in ERN 3.8 - at video level, not territory
func (vb *VideoBuilder) WithISRC(isrc string) *VideoBuilder {
	if vb.video.VideoId == nil {
		vb.video.VideoId = &VideoId{}
	}
	vb.video.VideoId.ISRC = isrc
	return vb
}

// AddKeywordsWithLanguage adds keywords with specific language (ERN 3.8 - territory specific)
func (vtb *VideoDetailsByTerritoryBuilder) AddKeywordsWithLanguage(keywords []string, languageCode string) *VideoDetailsByTerritoryBuilder {
	for _, keyword := range keywords {
		vtb.territoryDetails.Keywords = append(vtb.territoryDetails.Keywords, Keywords{
			Value:                 keyword,
			LanguageAndScriptCode: languageCode,
		})
	}
	return vtb
}

// AddProprietaryId adds a proprietary ID (e.g., YouTube channel ID) for ERN 3.8 - at video level
func (vb *VideoBuilder) AddProprietaryId(namespace, value string) *VideoBuilder {
	if vb.video.VideoId == nil {
		vb.video.VideoId = &VideoId{}
	}
	vb.video.VideoId.ProprietaryId = append(vb.video.VideoId.ProprietaryId, ProprietaryId{
		Namespace: namespace,
		Value:     value,
	})
	return vb
}

// Done returns to the main builder
func (vb *VideoBuilder) Done() *Builder {
	return vb.builder
}

// ImageBuilder provides fluent interface for building image resources
type ImageBuilder struct {
	builder                 *Builder
	image                   *Image
	currentTerritoryDetails *ImageDetailsByTerritory
	currentTerritoryIndex   int
}

// ImageDetailsByTerritoryBuilder provides fluent interface for building image territory details
type ImageDetailsByTerritoryBuilder struct {
	imageBuilder     *ImageBuilder
	territoryDetails *ImageDetailsByTerritory
}

// AddImageDetailsByTerritory creates a new territory details section and returns a builder for it
func (ib *ImageBuilder) AddImageDetailsByTerritory(territoryCodes []string) *ImageDetailsByTerritoryBuilder {
	// Validate that at least one territory code is provided
	if len(territoryCodes) == 0 {
		territoryCodes = []string{"Worldwide"}
	}

	// Create new territory details
	newDetails := ImageDetailsByTerritory{
		TerritoryCode: territoryCodes,
	}
	ib.image.ImageDetailsByTerritory = append(ib.image.ImageDetailsByTerritory, newDetails)
	ib.currentTerritoryIndex = len(ib.image.ImageDetailsByTerritory) - 1
	ib.currentTerritoryDetails = &ib.image.ImageDetailsByTerritory[ib.currentTerritoryIndex]

	return &ImageDetailsByTerritoryBuilder{
		imageBuilder:     ib,
		territoryDetails: ib.currentTerritoryDetails,
	}
}

// Done returns to the image builder
func (itb *ImageDetailsByTerritoryBuilder) Done() *ImageBuilder {
	return itb.imageBuilder
}

// WithProprietaryId adds a proprietary ID to the image (image level, not territory)
func (ib *ImageBuilder) WithProprietaryId(namespace, value string) *ImageBuilder {
	ib.image.ImageId = []ImageId{
		{
			ProprietaryId: []ProprietaryId{
				{Namespace: namespace, Value: value},
			},
		},
	}
	return ib
}

// WithCreationDate sets the creation date - at image level, not territory
func (ib *ImageBuilder) WithCreationDate(date string, isApproximate bool) *ImageBuilder {
	ib.image.CreationDate = &EventDate{
		Value:         date,
		IsApproximate: isApproximate,
	}
	return ib
}

// WithParentalWarning sets the parental warning type (territory specific)
func (itb *ImageDetailsByTerritoryBuilder) WithParentalWarning(warningType string) *ImageDetailsByTerritoryBuilder {
	itb.territoryDetails.ParentalWarningType = append(itb.territoryDetails.ParentalWarningType, warningType)
	return itb
}

// WithCLine sets the C-Line information (territory specific)
func (itb *ImageDetailsByTerritoryBuilder) WithCLine(year int, text string) *ImageDetailsByTerritoryBuilder {
	itb.territoryDetails.CLine = append(itb.territoryDetails.CLine, CLine{
		Year:      year,
		CLineText: text,
	})
	return itb
}

// Note: RightsController is not part of ImageDetailsByTerritory in ERN 3.8
// Rights information for images should be managed at the Image resource level, not territory level

// WithTechnicalDetails adds technical details and file URI for images (ERN 3.8 - territory specific)
func (itb *ImageDetailsByTerritoryBuilder) WithTechnicalDetails(techRef, fileURI string) *ImageDetailsByTerritoryBuilder {
	itb.territoryDetails.TechnicalImageDetails = append(itb.territoryDetails.TechnicalImageDetails, TechnicalImageDetails{
		TechnicalResourceDetailsReference: techRef,
		File: &File{
			URI: fileURI,
		},
	})
	return itb
}

// Done returns to the main builder
func (ib *ImageBuilder) Done() *Builder {
	return ib.builder
}

// ReleaseBuilder provides fluent interface for building releases
type ReleaseBuilder struct {
	builder                 *Builder
	release                 *Release
	currentTerritoryDetails *ReleaseDetailsByTerritory
	currentTerritoryIndex   int
}

// ReleaseDetailsByTerritoryBuilder provides fluent interface for building release territory details
type ReleaseDetailsByTerritoryBuilder struct {
	releaseBuilder   *ReleaseBuilder
	territoryDetails *ReleaseDetailsByTerritory
}

// WithTitle sets the reference title for the release (mandatory in ERN 3.8)
func (rb *ReleaseBuilder) WithTitle(title, subtitle string) *ReleaseBuilder {
	rb.release.ReferenceTitle = &ReferenceTitle{
		TitleText: title,
		SubTitle:  subtitle,
	}
	return rb
}

// AddReleaseDetailsByTerritory creates a new territory details section and returns a builder for it
// This is mandatory in ERN 3.8 - at least one territory must be specified
func (rb *ReleaseBuilder) AddReleaseDetailsByTerritory(territoryCodes []string) *ReleaseDetailsByTerritoryBuilder {
	// Validate that at least one territory code is provided
	if len(territoryCodes) == 0 {
		territoryCodes = []string{"Worldwide"}
	}

	// Create new territory details
	territoryDetails := ReleaseDetailsByTerritory{
		TerritoryCode: territoryCodes,
	}
	rb.release.ReleaseDetailsByTerritory = append(rb.release.ReleaseDetailsByTerritory, territoryDetails)
	rb.currentTerritoryIndex = len(rb.release.ReleaseDetailsByTerritory) - 1
	rb.currentTerritoryDetails = &rb.release.ReleaseDetailsByTerritory[rb.currentTerritoryIndex]

	return &ReleaseDetailsByTerritoryBuilder{
		releaseBuilder:   rb,
		territoryDetails: rb.currentTerritoryDetails,
	}
}

// Done returns to the release builder
func (rtb *ReleaseDetailsByTerritoryBuilder) Done() *ReleaseBuilder {
	return rtb.releaseBuilder
}

// WithDisplayArtistName sets the display artist name for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithDisplayArtistName(artistName, languageCode string) *ReleaseDetailsByTerritoryBuilder {
	if languageCode == "" {
		languageCode = "en"
	}
	rtb.territoryDetails.DisplayArtistName = append(rtb.territoryDetails.DisplayArtistName, Name{
		FullName:     artistName,
		LanguageCode: languageCode,
	})
	return rtb
}

// WithArtist adds a display artist reference for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithArtist(partyRef, role string, sequence int) *ReleaseDetailsByTerritoryBuilder {
	if partyRef != "" {
		rtb.territoryDetails.DisplayArtist = append(rtb.territoryDetails.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    role,
			SequenceNumber:       sequence,
		})
	}
	return rtb
}

// WithLabel adds a label name for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithLabel(labelName, languageCode string) *ReleaseDetailsByTerritoryBuilder {
	if languageCode == "" {
		languageCode = "en"
	}
	rtb.territoryDetails.LabelName = append(rtb.territoryDetails.LabelName, LabelName{
		Value:                 labelName,
		LanguageAndScriptCode: languageCode,
	})
	return rtb
}

// WithPLine adds P-Line information (can be global or territory-specific)
func (rb *ReleaseBuilder) WithPLine(year int, text string) *ReleaseBuilder {
	pline := PLine{
		Year:      year,
		PLineText: text,
	}
	// Add to global release
	rb.release.PLine = append(rb.release.PLine, pline)
	return rb
}

// WithTerritoryPLine adds P-Line information for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithTerritoryPLine(year int, text string) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.PLine = append(rtb.territoryDetails.PLine, PLine{
		Year:      year,
		PLineText: text,
	})
	return rtb
}

// WithCLine adds C-Line information (can be global or territory-specific)
func (rb *ReleaseBuilder) WithCLine(year int, text string) *ReleaseBuilder {
	cline := CLine{
		Year:      year,
		CLineText: text,
	}
	// Add to global release
	rb.release.CLine = append(rb.release.CLine, cline)
	return rb
}

// WithTerritoryCLine adds C-Line information for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithTerritoryCLine(year int, text string) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.CLine = append(rtb.territoryDetails.CLine, CLine{
		Year:      year,
		CLineText: text,
	})
	return rtb
}

// WithDuration sets the release duration
func (rb *ReleaseBuilder) WithDuration(duration string) *ReleaseBuilder {
	rb.release.Duration = duration
	return rb
}

// WithReleaseDate sets ReleaseDate for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithReleaseDate(date string) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.ReleaseDate = &EventDate{
		XMLName: xml.Name{Local: "ReleaseDate"},
		Value:   date,
	}
	return rtb
}

// WithOriginalReleaseDate sets OriginalReleaseDate for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithOriginalReleaseDate(date string) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.OriginalReleaseDate = &EventDate{
		XMLName: xml.Name{Local: "OriginalReleaseDate"},
		Value:   date,
	}
	return rtb
}

// WithGenre adds genre information for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithGenre(genreText string) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.Genre = append(rtb.territoryDetails.Genre, Genre{
		GenreText: genreText,
	})
	return rtb
}

// WithGenreAndSubGenre adds genre information with a subgenre for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithGenreAndSubGenre(genreText, subGenre string) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.Genre = append(rtb.territoryDetails.Genre, Genre{
		GenreText: genreText,
		SubGenre:  subGenre,
	})
	return rtb
}

// WithParentalWarning sets the parental warning type for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithParentalWarning(warningType string) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.ParentalWarningType = append(rtb.territoryDetails.ParentalWarningType, ParentalWarningType{
		Value: warningType,
	})
	return rtb
}

// WithAvRating adds an AvRating for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithAvRating(ratingText, agencyValue, agencyNamespace string) *ReleaseDetailsByTerritoryBuilder {
	avRating := AvRating{
		RatingText: ratingText,
		RatingAgency: &RatingAgency{
			Value:     agencyValue,
			Namespace: agencyNamespace,
		},
	}
	rtb.territoryDetails.AvRating = append(rtb.territoryDetails.AvRating, avRating)
	return rtb
}

// WithMarketingComment adds a marketing comment for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) WithMarketingComment(comment, languageCode string) *ReleaseDetailsByTerritoryBuilder {
	if languageCode == "" {
		languageCode = "en"
	}
	rtb.territoryDetails.MarketingComment = &Comment{
		Value:                 comment,
		LanguageAndScriptCode: languageCode,
	}
	return rtb
}

// AddKeywordsWithLanguage adds keywords with specific language for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) AddKeywordsWithLanguage(keywords []string, languageCode string) *ReleaseDetailsByTerritoryBuilder {
	if languageCode == "" {
		languageCode = "en"
	}
	for _, keyword := range keywords {
		keywordsEntry := Keywords{
			Value:                 keyword,
			LanguageAndScriptCode: languageCode,
		}
		rtb.territoryDetails.Keywords = append(rtb.territoryDetails.Keywords, keywordsEntry)
	}
	return rtb
}

// WithICPN sets the ICPN identifier for the release (ERN 3.8)
func (rb *ReleaseBuilder) WithICPN(icpn string) *ReleaseBuilder {
	rb.release.ReleaseId = append(rb.release.ReleaseId, ReleaseId{
		ICPN: icpn,
	})
	return rb
}

// WithISRC sets the ISRC identifier for the release
// Only applicable when the Release contains only one SoundRecording or one MusicalWorkVideo
func (rb *ReleaseBuilder) WithISRC(isrc string) *ReleaseBuilder {
	rb.release.ReleaseId = append(rb.release.ReleaseId, ReleaseId{
		ISRC: isrc,
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

// AddReleaseResourceReference adds a resource reference to the release
// In ERN 3.8, this is used at the Release level to reference resources
// releaseResourceType can be "PrimaryResource", "SecondaryResource", etc.
func (rb *ReleaseBuilder) AddReleaseResourceReference(resourceRef, releaseResourceType string) *ReleaseBuilder {
	if rb.release.ReleaseResourceReferenceList == nil {
		rb.release.ReleaseResourceReferenceList = &ReleaseResourceReferenceList{}
	}
	rb.release.ReleaseResourceReferenceList.ReleaseResourceReference = append(
		rb.release.ReleaseResourceReferenceList.ReleaseResourceReference,
		ReleaseResourceReference{
			ReleaseResourceType: releaseResourceType,
			Value:               resourceRef,
		},
	)
	return rb
}

// AddRelatedRelease adds a related release for the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) AddRelatedRelease(relationshipType string, releaseId ReleaseId) *ReleaseDetailsByTerritoryBuilder {
	rtb.territoryDetails.RelatedRelease = append(rtb.territoryDetails.RelatedRelease, RelatedRelease{
		ReleaseRelationshipType: relationshipType,
		ReleaseId:               releaseId,
	})
	return rtb
}

// AddResourceGroup adds a resource group to the current territory
func (rtb *ReleaseDetailsByTerritoryBuilder) AddResourceGroup(titleText string, sequenceNumber int) *ResourceGroupBuilder {
	group := ResourceGroup{
		SequenceNumber: sequenceNumber,
	}

	if titleText != "" {
		group.AdditionalTitle = AdditionalTitle{
			TitleText: titleText,
		}
	}

	rtb.territoryDetails.ResourceGroup = append(rtb.territoryDetails.ResourceGroup, group)
	groupIndex := len(rtb.territoryDetails.ResourceGroup) - 1

	return &ResourceGroupBuilder{
		releaseDetailsByTerritoryBuilder: rtb,
		group:                            &rtb.territoryDetails.ResourceGroup[groupIndex],
	}
}

// Done returns to the main builder
func (rb *ReleaseBuilder) Done() *Builder {
	return rb.builder
}

// ResourceGroupBuilder provides fluent interface for building resource groups
type ResourceGroupBuilder struct {
	releaseDetailsByTerritoryBuilder *ReleaseDetailsByTerritoryBuilder
	group                            *ResourceGroup
}

// AddContentItem adds a content item to the resource group
// resourceType can be "Video", "Image", "SoundRecording", etc.
// releaseResourceType can be "PrimaryResource", "SecondaryResource", etc.
func (rgb *ResourceGroupBuilder) AddContentItem(sequenceNumber int, resourceType, resourceRef, releaseResourceType string) *ResourceGroupBuilder {
	item := ResourceGroupContentItem{
		SequenceNumber: sequenceNumber,
		ResourceType:   resourceType,
		ReleaseResourceReference: ReleaseResourceReference{
			ReleaseResourceType: releaseResourceType,
			Value:               resourceRef,
		},
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

// Done returns to the release details by territory builder
func (rgb *ResourceGroupBuilder) Done() *ReleaseDetailsByTerritoryBuilder {
	return rgb.releaseDetailsByTerritoryBuilder
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

// WithTerritories sets the deal territories for ERN 3.8
func (db *DealBuilder) WithTerritories(territoryCodes []string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.TerritoryCode = append(db.deal.DealTerms.TerritoryCode, territoryCodes...)
	return db
}

// WithValidityPeriodStartDate sets the deal validity period start date (YYYY-MM-DD)
func (db *DealBuilder) WithValidityPeriodStartDate(startDate string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	// Ensure at least one ValidityPeriod exists
	if len(db.deal.DealTerms.ValidityPeriod) == 0 {
		db.deal.DealTerms.ValidityPeriod = append(db.deal.DealTerms.ValidityPeriod, ValidityPeriod{})
	}

	// Set the start date on the first ValidityPeriod
	db.deal.DealTerms.ValidityPeriod[0].StartDate = startDate

	return db
}

// WithValidityPeriodEndDate sets the deal validity period end date (YYYY-MM-DD)
func (db *DealBuilder) WithValidityPeriodEndDate(endDate string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	// Ensure at least one ValidityPeriod exists
	if len(db.deal.DealTerms.ValidityPeriod) == 0 {
		db.deal.DealTerms.ValidityPeriod = append(db.deal.DealTerms.ValidityPeriod, ValidityPeriod{})
	}

	// Set the end date on the first ValidityPeriod
	db.deal.DealTerms.ValidityPeriod[0].EndDate = endDate

	return db
}

// WithValidityPeriodDateTime sets the deal validity period with a start date-time (YYYY-MM-DDTHH:MM:SS)
func (db *DealBuilder) WithValidityPeriodDateTime(startDateTime string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	db.deal.DealTerms.ValidityPeriod = append(db.deal.DealTerms.ValidityPeriod, ValidityPeriod{
		StartDateTime: startDateTime,
	})

	return db
}

// WithCommercialModel adds a commercial model type for ERN 3.8 (can be called multiple times)
func (db *DealBuilder) WithCommercialModel(modelType string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.CommercialModelType = append(db.deal.DealTerms.CommercialModelType, modelType)
	return db
}

// WithUseType adds a use type for ERN 3.8 (can be called multiple times)
func (db *DealBuilder) WithUseType(useType string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	// Ensure Usage array exists
	if len(db.deal.DealTerms.Usage) == 0 {
		db.deal.DealTerms.Usage = append(db.deal.DealTerms.Usage, Usage{})
	}

	// Add to the first Usage element's UseType array
	db.deal.DealTerms.Usage[0].UseType = append(db.deal.DealTerms.Usage[0].UseType, useType)
	return db
}

// WithRightsClaimPolicy adds a rights claim policy for the deal (can be called multiple times)
func (db *DealBuilder) WithRightsClaimPolicy(policyType string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.RightsClaimPolicy = append(db.deal.DealTerms.RightsClaimPolicy, RightsClaimPolicy{
		RightsClaimPolicyType: policyType,
	})
	return db
}

// Done returns to the release deal builder
func (db *DealBuilder) Done() *ReleaseDealBuilder {
	return db.releaseDealBuilder
}
