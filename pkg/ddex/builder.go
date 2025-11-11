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
			XmlnsErn:                XmlnsErn,
			XmlnsXsi:                XmlnsXsi,
			XsiSchemaLocation:       XsiSchemaLocation,
			MessageSchemaVersionId:  MessageSchemaVersionId,
			ReleaseProfileVersionId: "Video",
			LanguageAndScriptCode:   "en",
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

// ensureTerritory ensures there's a territory to work with (defaults to Worldwide)
func (vb *VideoBuilder) ensureTerritory() {
	if vb.currentTerritoryDetails == nil {
		vb.WithTerritory([]string{"Worldwide"})
	}
}

// WithTerritory creates or switches to a territory section
func (vb *VideoBuilder) WithTerritory(territoryCodes []string) *VideoBuilder {
	// Check if territory already exists (compare first code for simplicity)
	if len(territoryCodes) > 0 {
		for i, details := range vb.video.VideoDetailsByTerritory {
			if len(details.TerritoryCode) > 0 && details.TerritoryCode[0] == territoryCodes[0] {
				vb.currentTerritoryDetails = &vb.video.VideoDetailsByTerritory[i]
				vb.currentTerritoryIndex = i
				return vb
			}
		}
	}

	// Create new territory details
	newDetails := VideoDetailsByTerritory{
		TerritoryCode: territoryCodes,
	}
	vb.video.VideoDetailsByTerritory = append(vb.video.VideoDetailsByTerritory, newDetails)
	vb.currentTerritoryIndex = len(vb.video.VideoDetailsByTerritory) - 1
	vb.currentTerritoryDetails = &vb.video.VideoDetailsByTerritory[vb.currentTerritoryIndex]

	return vb
}

// WithTitle sets the video title (goes to territory details in ERN 3.8)
func (vb *VideoBuilder) WithTitle(title, subtitle string) *VideoBuilder {
	vb.ensureTerritory()

	// Add title to territory details
	titleStruct := Title{
		TitleText: title,
	}
	if subtitle != "" {
		titleStruct.SubTitle = subtitle
	}
	vb.currentTerritoryDetails.Title = append(vb.currentTerritoryDetails.Title, titleStruct)

	return vb
}

// WithDisplayArtistName sets the display artist name for the video (ERN 3.8 - territory specific)
func (vb *VideoBuilder) WithDisplayArtistName(artistName, languageCode string) *VideoBuilder {
	vb.ensureTerritory()

	if languageCode == "" {
		languageCode = "en"
	}
	vb.currentTerritoryDetails.DisplayArtistName = append(vb.currentTerritoryDetails.DisplayArtistName, DisplayArtistName{
		Value:                 artistName,
		LanguageAndScriptCode: languageCode,
	})
	return vb
}

// WithArtist adds a display artist reference to the video (territory specific)
func (vb *VideoBuilder) WithArtist(partyRef, role string, sequence int) *VideoBuilder {
	vb.ensureTerritory()

	if partyRef != "" {
		vb.currentTerritoryDetails.DisplayArtist = append(vb.currentTerritoryDetails.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    role,
			SequenceNumber:       sequence,
		})
	}

	return vb
}

// WithContributor adds a contributor to the video resource (territory specific)
// role can be multiple values like "Producer", "Director", "Cinematographer", etc.
func (vb *VideoBuilder) WithContributor(partyRef string, roles []string, sequence int) *VideoBuilder {
	vb.ensureTerritory()

	if partyRef != "" && len(roles) > 0 {
		vb.currentTerritoryDetails.ResourceContributor = append(vb.currentTerritoryDetails.ResourceContributor, ResourceContributor{
			PartyReference: partyRef,
			Role:           roles,
		})
	}

	return vb
}

// WithRightsController sets the rights controller (territory specific)
func (vb *VideoBuilder) WithRightsController(partyRef string, percentage float64, territories []string) *VideoBuilder {
	vb.ensureTerritory()

	// If no territories provided, default to Worldwide
	if len(territories) == 0 {
		territories = []string{"Worldwide"}
	}

	vb.currentTerritoryDetails.RightsController = append(vb.currentTerritoryDetails.RightsController, RightsController{
		PartyReference:       partyRef,
		Role:                 []string{"RightsController"},
		RightSharePercentage: fmt.Sprintf("%.2f", percentage),
		DelegatedUsageRights: []DelegatedUsageRights{
			{
				UseType:                     []string{"UserMakeAvailableUserProvided"},
				TerritoryOfRightsDelegation: territories,
			},
		},
	})

	return vb
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

// WithParentalWarning sets the parental warning type (territory specific)
func (vb *VideoBuilder) WithParentalWarning(warningType string) *VideoBuilder {
	vb.ensureTerritory()

	vb.currentTerritoryDetails.ParentalWarningType = append(vb.currentTerritoryDetails.ParentalWarningType, warningType)
	return vb
}

// WithPLine sets the P-Line information for ERN 3.8 (territory specific)
func (vb *VideoBuilder) WithPLine(year int, text string) *VideoBuilder {
	vb.ensureTerritory()

	vb.currentTerritoryDetails.PLine = append(vb.currentTerritoryDetails.PLine, PLine{
		Year:      year,
		PLineText: text,
	})
	return vb
}

// WithCLine sets the C-Line information for ERN 3.8 (territory specific)
func (vb *VideoBuilder) WithCLine(year int, text string) *VideoBuilder {
	vb.ensureTerritory()

	vb.currentTerritoryDetails.CLine = append(vb.currentTerritoryDetails.CLine, CLine{
		Year:      year,
		CLineText: text,
	})
	return vb
}

// WithGenre adds genre information (territory specific)
func (vb *VideoBuilder) WithGenre(genreText, subGenre string) *VideoBuilder {
	vb.ensureTerritory()

	genre := Genre{
		GenreText: genreText,
	}
	if subGenre != "" {
		genre.SubGenre = subGenre
	}
	vb.currentTerritoryDetails.Genre = append(vb.currentTerritoryDetails.Genre, genre)
	return vb
}

// WithTechnicalDetails adds technical details and file URI for ERN 3.8 (territory specific)
func (vb *VideoBuilder) WithTechnicalDetails(techRef, fileURI string) *VideoBuilder {
	vb.ensureTerritory()

	vb.currentTerritoryDetails.TechnicalVideoDetails = append(vb.currentTerritoryDetails.TechnicalVideoDetails, TechnicalVideoDetails{
		TechnicalResourceDetailsReference: techRef,
		File: &File{
			URI: fileURI,
		},
	})
	return vb
}

// WithISRC sets the ISRC for the video in ERN 3.8 - at video level, not territory
func (vb *VideoBuilder) WithISRC(isrc string) *VideoBuilder {
	vb.video.VideoId = append(vb.video.VideoId, VideoId{
		ISRC: isrc,
	})
	return vb
}

// AddKeywords adds keywords for enhanced search and display (ERN 3.8 - territory specific)
func (vb *VideoBuilder) AddKeywords(keywords ...string) *VideoBuilder {
	vb.ensureTerritory()

	for _, keyword := range keywords {
		vb.currentTerritoryDetails.Keywords = append(vb.currentTerritoryDetails.Keywords, Keywords{
			Value: keyword,
		})
	}
	return vb
}

// AddKeywordsWithLanguage adds keywords with specific language (ERN 3.8 - territory specific)
func (vb *VideoBuilder) AddKeywordsWithLanguage(languageCode string, keywords ...string) *VideoBuilder {
	vb.ensureTerritory()

	for _, keyword := range keywords {
		vb.currentTerritoryDetails.Keywords = append(vb.currentTerritoryDetails.Keywords, Keywords{
			Value:                 keyword,
			LanguageAndScriptCode: languageCode,
		})
	}
	return vb
}

// AddProprietaryId adds a proprietary ID (e.g., YouTube channel ID) for ERN 3.8 - at video level
func (vb *VideoBuilder) AddProprietaryId(namespace, value string) *VideoBuilder {
	vb.video.VideoId = append(vb.video.VideoId, VideoId{
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
	builder                 *Builder
	image                   *Image
	currentTerritoryDetails *ImageDetailsByTerritory
	currentTerritoryIndex   int
}

// ensureTerritory ensures there's a territory to work with (defaults to Worldwide)
func (ib *ImageBuilder) ensureTerritory() {
	if ib.currentTerritoryDetails == nil {
		ib.WithTerritory([]string{"Worldwide"})
	}
}

// WithTerritory creates or switches to a territory section
func (ib *ImageBuilder) WithTerritory(territoryCodes []string) *ImageBuilder {
	// Check if territory already exists (compare first code for simplicity)
	if len(territoryCodes) > 0 {
		for i, details := range ib.image.ImageDetailsByTerritory {
			if len(details.TerritoryCode) > 0 && details.TerritoryCode[0] == territoryCodes[0] {
				ib.currentTerritoryDetails = &ib.image.ImageDetailsByTerritory[i]
				ib.currentTerritoryIndex = i
				return ib
			}
		}
	}

	// Create new territory details
	newDetails := ImageDetailsByTerritory{
		TerritoryCode: territoryCodes,
	}
	ib.image.ImageDetailsByTerritory = append(ib.image.ImageDetailsByTerritory, newDetails)
	ib.currentTerritoryIndex = len(ib.image.ImageDetailsByTerritory) - 1
	ib.currentTerritoryDetails = &ib.image.ImageDetailsByTerritory[ib.currentTerritoryIndex]

	return ib
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
func (ib *ImageBuilder) WithParentalWarning(warningType string) *ImageBuilder {
	ib.ensureTerritory()

	ib.currentTerritoryDetails.ParentalWarningType = append(ib.currentTerritoryDetails.ParentalWarningType, warningType)
	return ib
}

// WithCLine sets the C-Line information (territory specific)
func (ib *ImageBuilder) WithCLine(year int, text string) *ImageBuilder {
	ib.ensureTerritory()

	ib.currentTerritoryDetails.CLine = append(ib.currentTerritoryDetails.CLine, CLine{
		Year:      year,
		CLineText: text,
	})
	return ib
}

// Note: RightsController is not part of ImageDetailsByTerritory in ERN 3.8
// Rights information for images should be managed at the Image resource level, not territory level

// WithTechnicalDetails adds technical details and file URI for images (ERN 3.8 - territory specific)
func (ib *ImageBuilder) WithTechnicalDetails(techRef, fileURI string) *ImageBuilder {
	ib.ensureTerritory()

	ib.currentTerritoryDetails.TechnicalImageDetails = append(ib.currentTerritoryDetails.TechnicalImageDetails, TechnicalImageDetails{
		TechnicalResourceDetailsReference: techRef,
		File: &File{
			URI: fileURI,
		},
	})
	return ib
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

// WithTitle sets the reference title for the release (mandatory in ERN 3.8)
func (rb *ReleaseBuilder) WithTitle(title, subtitle string) *ReleaseBuilder {
	rb.release.ReferenceTitle = &ReferenceTitle{
		TitleText: title,
		SubTitle:  subtitle,
	}
	return rb
}

// WithTerritory creates or switches to a territory-specific details section
// This is mandatory in ERN 3.8 - at least one territory must be specified
// WithTerritory creates or switches to a territory-specific details section
// This is mandatory in ERN 3.8 - at least one territory must be specified
func (rb *ReleaseBuilder) WithTerritory(territoryCodes []string) *ReleaseBuilder {
	// Create new territory details
	territoryDetails := ReleaseDetailsByTerritory{
		TerritoryCode: territoryCodes,
	}
	rb.release.ReleaseDetailsByTerritory = append(rb.release.ReleaseDetailsByTerritory, territoryDetails)
	rb.currentTerritoryIndex = len(rb.release.ReleaseDetailsByTerritory) - 1
	rb.currentTerritoryDetails = &rb.release.ReleaseDetailsByTerritory[rb.currentTerritoryIndex]
	return rb
}

// ensureTerritory creates a default Worldwide territory if none exists
func (rb *ReleaseBuilder) ensureTerritory() {
	if rb.currentTerritoryDetails == nil {
		rb.WithTerritory([]string{"Worldwide"})
	}
}

// WithDisplayArtistName sets the display artist name for the current territory
func (rb *ReleaseBuilder) WithDisplayArtistName(artistName, languageCode string) *ReleaseBuilder {
	rb.ensureTerritory()
	if languageCode == "" {
		languageCode = "en"
	}
	rb.currentTerritoryDetails.DisplayArtistName = append(rb.currentTerritoryDetails.DisplayArtistName, Name{
		FullName:     artistName,
		LanguageCode: languageCode,
	})
	return rb
}

// WithArtist adds a display artist reference for the current territory
func (rb *ReleaseBuilder) WithArtist(partyRef, role string, sequence int) *ReleaseBuilder {
	rb.ensureTerritory()
	if partyRef != "" {
		rb.currentTerritoryDetails.DisplayArtist = append(rb.currentTerritoryDetails.DisplayArtist, DisplayArtist{
			ArtistPartyReference: partyRef,
			DisplayArtistRole:    role,
			SequenceNumber:       sequence,
		})
	}
	return rb
}

// WithLabel adds a label name for the current territory
func (rb *ReleaseBuilder) WithLabel(labelName, languageCode string) *ReleaseBuilder {
	rb.ensureTerritory()
	if languageCode == "" {
		languageCode = "en"
	}
	rb.currentTerritoryDetails.LabelName = append(rb.currentTerritoryDetails.LabelName, LabelName{
		Value:                 labelName,
		LanguageAndScriptCode: languageCode,
	})
	return rb
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
func (rb *ReleaseBuilder) WithTerritoryPLine(year int, text string) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.PLine = append(rb.currentTerritoryDetails.PLine, PLine{
		Year:      year,
		PLineText: text,
	})
	return rb
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
func (rb *ReleaseBuilder) WithTerritoryCLine(year int, text string) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.CLine = append(rb.currentTerritoryDetails.CLine, CLine{
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

// WithReleaseDate sets ReleaseDate for the current territory
func (rb *ReleaseBuilder) WithReleaseDate(date string) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.ReleaseDate = &EventDate{
		XMLName: xml.Name{Local: "ReleaseDate"},
		Value:   date,
	}
	return rb
}

// WithOriginalReleaseDate sets OriginalReleaseDate for the current territory
func (rb *ReleaseBuilder) WithOriginalReleaseDate(date string) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.OriginalReleaseDate = &EventDate{
		XMLName: xml.Name{Local: "OriginalReleaseDate"},
		Value:   date,
	}
	return rb
}

// WithGenre adds genre information for the current territory
func (rb *ReleaseBuilder) WithGenre(genreText string) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.Genre = append(rb.currentTerritoryDetails.Genre, Genre{
		GenreText: genreText,
	})
	return rb
}

// WithGenreAndSubGenre adds genre information with a subgenre for the current territory
func (rb *ReleaseBuilder) WithGenreAndSubGenre(genreText, subGenre string) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.Genre = append(rb.currentTerritoryDetails.Genre, Genre{
		GenreText: genreText,
		SubGenre:  subGenre,
	})
	return rb
}

// WithParentalWarning sets the parental warning type for the current territory
func (rb *ReleaseBuilder) WithParentalWarning(warningType string) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.ParentalWarningType = append(rb.currentTerritoryDetails.ParentalWarningType, ParentalWarningType{
		Value: warningType,
	})
	return rb
}

// WithAvRating adds an AvRating for the current territory
func (rb *ReleaseBuilder) WithAvRating(ratingText, agencyValue string) *ReleaseBuilder {
	rb.ensureTerritory()
	avRating := AvRating{
		RatingText:   ratingText,
		RatingAgency: agencyValue,
	}
	rb.currentTerritoryDetails.AvRating = append(rb.currentTerritoryDetails.AvRating, avRating)
	return rb
}

// WithMadeForKids is a convenience method to set the YouTube MadeForKids rating
func (rb *ReleaseBuilder) WithMadeForKids() *ReleaseBuilder {
	return rb.WithAvRating("MadeForKids", "UserDefined")
}

// WithMarketingComment adds a marketing comment for the current territory
func (rb *ReleaseBuilder) WithMarketingComment(comment, languageCode string) *ReleaseBuilder {
	rb.ensureTerritory()
	if languageCode == "" {
		languageCode = "en"
	}
	rb.currentTerritoryDetails.MarketingComment = &Comment{
		Value:                 comment,
		LanguageAndScriptCode: languageCode,
	}
	return rb
}

// WithKeywords adds keywords for the current territory
func (rb *ReleaseBuilder) WithKeywords(keywords, languageCode string) *ReleaseBuilder {
	rb.ensureTerritory()
	if languageCode == "" {
		languageCode = "en"
	}
	keywordsEntry := Keywords{
		Value:                 keywords,
		LanguageAndScriptCode: languageCode,
	}
	rb.currentTerritoryDetails.Keywords = append(rb.currentTerritoryDetails.Keywords, keywordsEntry)
	return rb
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
func (rb *ReleaseBuilder) AddReleaseResourceReference(resourceRef string) *ReleaseBuilder {
	rb.release.ReleaseResourceReference = append(rb.release.ReleaseResourceReference, resourceRef)
	return rb
}

// AddRelatedRelease adds a related release for the current territory
func (rb *ReleaseBuilder) AddRelatedRelease(relationshipType string, releaseId ReleaseId) *ReleaseBuilder {
	rb.ensureTerritory()
	rb.currentTerritoryDetails.RelatedRelease = append(rb.currentTerritoryDetails.RelatedRelease, RelatedRelease{
		ReleaseRelationshipType: relationshipType,
		ReleaseId:               releaseId,
	})
	return rb
}

// AddResourceGroup adds a resource group to the current territory
func (rb *ReleaseBuilder) AddResourceGroup(titleText string, sequenceNumber int) *ResourceGroupBuilder {
	rb.ensureTerritory()

	group := ResourceGroup{
		SequenceNumber: sequenceNumber,
	}

	if titleText != "" {
		group.AdditionalTitle = AdditionalTitle{
			TitleText: titleText,
		}
	}

	rb.currentTerritoryDetails.ResourceGroup = append(rb.currentTerritoryDetails.ResourceGroup, group)
	groupIndex := len(rb.currentTerritoryDetails.ResourceGroup) - 1

	return &ResourceGroupBuilder{
		releaseBuilder: rb,
		group:          &rb.currentTerritoryDetails.ResourceGroup[groupIndex],
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

// WithTerritories sets the deal territories for ERN 3.8 (can be called multiple times)
func (db *DealBuilder) WithTerritory(territoryCode string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}
	db.deal.DealTerms.TerritoryCode = append(db.deal.DealTerms.TerritoryCode, territoryCode)
	return db
}

// WithValidityPeriod sets the deal validity period with a start date (YYYY-MM-DD)
func (db *DealBuilder) WithValidityPeriod(startDate string, endDate string) *DealBuilder {
	if db.deal.DealTerms == nil {
		db.deal.DealTerms = &DealTerms{}
	}

	db.deal.DealTerms.ValidityPeriod = append(db.deal.DealTerms.ValidityPeriod, ValidityPeriod{
		StartDate: startDate,
		EndDate:   endDate,
	})

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
