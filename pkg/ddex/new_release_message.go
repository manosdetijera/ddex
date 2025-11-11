package ddex

import (
	"encoding/xml"
	"fmt"
)

// NewReleaseMessage represents the complete DDEX ERN 3.8 NewReleaseMessage structure
// specifically configured for YouTube delivery
type NewReleaseMessage struct {
	XMLName                 xml.Name        `xml:"ern:NewReleaseMessage"`
	XmlnsErn                string          `xml:"xmlns:ern,attr"`
	XmlnsXsi                string          `xml:"xmlns:xsi,attr,omitempty"`
	XsiSchemaLocation       string          `xml:"xsi:schemaLocation,attr,omitempty"`
	MessageSchemaVersionId  string          `xml:"MessageSchemaVersionId,attr"`
	ReleaseProfileVersionId string          `xml:"ReleaseProfileVersionId,attr,omitempty"`
	LanguageAndScriptCode   string          `xml:"LanguageAndScriptCode,attr,omitempty"`
	MessageHeader           *MessageHeader  `xml:"MessageHeader"`
	UpdateIndicator         string          `xml:"UpdateIndicator,omitempty"` // Deprecated: OriginalMessage or UpdateMessage
	ResourceList            *ResourceList   `xml:"ResourceList,omitempty"`
	CollectionList          *CollectionList `xml:"CollectionList,omitempty"`
	ReleaseList             *ReleaseList    `xml:"ReleaseList"`
	DealList                *DealList       `xml:"DealList"`
}

// CollectionList represents collections (playlists, compilations)
type CollectionList struct {
	XMLName    xml.Name     `xml:"CollectionList"`
	Collection []Collection `xml:"Collection"`
}

// Collection represents a collection of releases
type Collection struct {
	XMLName                      xml.Name                       `xml:"Collection"`
	CollectionReference          string                         `xml:"CollectionReference"`
	CollectionType               string                         `xml:"CollectionType,omitempty"`
	CollectionId                 []ReleaseId                    `xml:"CollectionId,omitempty"`
	DisplayTitleText             []TitleText                    `xml:"DisplayTitleText"`
	DisplayArtistName            []string                       `xml:"DisplayArtistName,omitempty"`
	DisplayArtist                []DisplayArtist                `xml:"DisplayArtist,omitempty"`
	CollectionDetailsByTerritory []CollectionDetailsByTerritory `xml:"CollectionDetailsByTerritory,omitempty"`
}

// CollectionDetailsByTerritory represents territory-specific collection details
type CollectionDetailsByTerritory struct {
	XMLName           xml.Name    `xml:"CollectionDetailsByTerritory"`
	TerritoryCode     string      `xml:"TerritoryCode"`
	DisplayTitleText  []TitleText `xml:"DisplayTitleText,omitempty"`
	DisplayArtistName []string    `xml:"DisplayArtistName,omitempty"`
	Genre             []Genre     `xml:"Genre,omitempty"`
}

// YouTube-specific constants for ERN 3.8
const (
	MessageSchemaVersionId = "ern/382"
	XmlnsErn               = "http://ddex.net/xml/ern/382"
	XmlnsXsi               = "http://www.w3.org/2001/XMLSchema-instance"
	XsiSchemaLocation      = "http://ddex.net/xml/ern/382 http://ddex.net/xml/ern/382/release-notification.xsd"
)

// NewReleaseMessageBuilder provides a fluent interface for building DDEX messages
type NewReleaseMessageBuilder struct {
	message *NewReleaseMessage
}

// NewNewReleaseMessage creates a new ERN 3.8 NewReleaseMessage for YouTube
func NewNewReleaseMessage(messageId, threadId, senderDPID, senderName, releaseProfileVersionId string) *NewReleaseMessage {
	// Create message header
	sender := NewMessageSender(senderDPID, senderName)
	header := NewMessageHeader(threadId, messageId, sender)

	return &NewReleaseMessage{
		MessageSchemaVersionId:  MessageSchemaVersionId,
		XmlnsErn:                XmlnsErn,
		XmlnsXsi:                XmlnsXsi,
		XsiSchemaLocation:       XsiSchemaLocation,
		ReleaseProfileVersionId: releaseProfileVersionId,
		LanguageAndScriptCode:   "en",
		MessageHeader:           header,
		ResourceList:            &ResourceList{},
		ReleaseList:             &ReleaseList{},
		DealList:                &DealList{},
	}
}

// NewBuilder creates a new builder for constructing NewReleaseMessage
func NewBuilder(messageId, threadId, senderDPID, senderName, releaseProfileVersionId string) *NewReleaseMessageBuilder {
	return &NewReleaseMessageBuilder{
		message: NewNewReleaseMessage(messageId, threadId, senderDPID, senderName, releaseProfileVersionId),
	}
}

// SetLanguage sets the language and script code for the message
func (b *NewReleaseMessage) SetLanguage(languageCode string) *NewReleaseMessage {
	b.LanguageAndScriptCode = languageCode
	return b
}

// SetReleaseProfile sets the release profile version
func (b *NewReleaseMessage) SetReleaseProfile(profileVersion string) *NewReleaseMessage {
	b.ReleaseProfileVersionId = profileVersion
	return b
}

// SetUpdateIndicator sets the update indicator
// Valid values: "OriginalMessage" or "UpdateMessage"
// Note: This element is deprecated in ERN 3.8
func (b *NewReleaseMessage) SetUpdateIndicator(indicator string) *NewReleaseMessage {
	b.UpdateIndicator = indicator
	return b
}

// AddSoundRecording adds a sound recording to the resource list
func (b *NewReleaseMessage) AddSoundRecording(recording *SoundRecording) *NewReleaseMessage {
	if b.ResourceList == nil {
		b.ResourceList = &ResourceList{}
	}
	b.ResourceList.SoundRecording = append(b.ResourceList.SoundRecording, *recording)
	return b
}

// AddVideo adds a video to the resource list
func (b *NewReleaseMessage) AddVideo(video *Video) *NewReleaseMessage {
	if b.ResourceList == nil {
		b.ResourceList = &ResourceList{}
	}
	b.ResourceList.Video = append(b.ResourceList.Video, *video)
	return b
}

// AddImage adds an image to the resource list
func (b *NewReleaseMessage) AddImage(image *Image) *NewReleaseMessage {
	if b.ResourceList == nil {
		b.ResourceList = &ResourceList{}
	}
	b.ResourceList.Image = append(b.ResourceList.Image, *image)
	return b
}

// AddRelease adds a release to the release list
func (b *NewReleaseMessage) AddRelease(release *Release) *NewReleaseMessage {
	if b.ReleaseList == nil {
		b.ReleaseList = &ReleaseList{}
	}
	b.ReleaseList.Release = append(b.ReleaseList.Release, *release)
	return b
}

// AddDeal adds a deal to the deal list (deprecated - use Builder.AddDeal instead)
func (b *NewReleaseMessage) AddDeal(deal *ReleaseDeal) *NewReleaseMessage {
	if b.DealList == nil {
		b.DealList = &DealList{}
	}
	b.DealList.ReleaseDeal = append(b.DealList.ReleaseDeal, *deal)
	return b
}

// Build returns the constructed NewReleaseMessage
func (b *NewReleaseMessageBuilder) Build() *NewReleaseMessage {
	return b.message
}

// ToXML converts the NewReleaseMessage to XML
func (nrm *NewReleaseMessage) ToXML() ([]byte, error) {
	return xml.MarshalIndent(nrm, "", "  ")
}

// ToXMLWithHeader converts the NewReleaseMessage to XML with XML declaration
func (nrm *NewReleaseMessage) ToXMLWithHeader() ([]byte, error) {
	xmlData, err := nrm.ToXML()
	if err != nil {
		return nil, err
	}

	header := `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	return append([]byte(header), xmlData...), nil
}

// FromXML parses XML data into a NewReleaseMessage
func FromXML(data []byte) (*NewReleaseMessage, error) {
	var nrm NewReleaseMessage
	err := xml.Unmarshal(data, &nrm)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}
	return &nrm, nil
}

// Validate performs basic validation on the NewReleaseMessage structure
func (nrm *NewReleaseMessage) Validate() error {
	if nrm.MessageHeader == nil {
		return fmt.Errorf("MessageHeader is required")
	}

	if nrm.MessageHeader.MessageId == "" {
		return fmt.Errorf("MessageHeader.MessageId is required")
	}

	if nrm.MessageHeader.MessageThreadId == "" {
		return fmt.Errorf("MessageHeader.MessageThreadId is required")
	}

	if nrm.MessageHeader.MessageSender == nil {
		return fmt.Errorf("MessageHeader.MessageSender is required")
	}

	if nrm.MessageHeader.MessageRecipient == nil {
		return fmt.Errorf("MessageHeader.MessageRecipient is required")
	}

	if nrm.ReleaseList == nil || len(nrm.ReleaseList.Release) == 0 {
		return fmt.Errorf("at least one Release is required")
	}

	if nrm.DealList == nil || len(nrm.DealList.ReleaseDeal) == 0 {
		return fmt.Errorf("at least one Deal is required")
	}

	// Validate that all releases have corresponding deals
	dealReleaseRefs := make(map[string]bool)
	for _, releaseDeal := range nrm.DealList.ReleaseDeal {
		dealReleaseRefs[releaseDeal.DealReleaseReference] = true
	}

	for _, release := range nrm.ReleaseList.Release {
		if !dealReleaseRefs[release.ReleaseReference] {
			return fmt.Errorf("no deal found for release reference: %s", release.ReleaseReference)
		}
	}

	return nil
}

// GetReleaseIDs returns all release IDs from the message (ERN 3.8)
func (nrm *NewReleaseMessage) GetReleaseIDs() []string {
	var ids []string
	if nrm.ReleaseList != nil {
		for _, release := range nrm.ReleaseList.Release {
			for _, releaseID := range release.ReleaseId {
				if releaseID.ICPN != "" {
					ids = append(ids, releaseID.ICPN)
				}
				if releaseID.GRid != "" {
					ids = append(ids, releaseID.GRid)
				}
				if releaseID.ISAN != "" {
					ids = append(ids, releaseID.ISAN)
				}
			}
		}
	}
	return ids
}

// GetMainRelease returns the main release from the release list (returns first release)
func (nrm *NewReleaseMessage) GetMainRelease() *Release {
	if nrm.ReleaseList != nil && len(nrm.ReleaseList.Release) > 0 {
		return &nrm.ReleaseList.Release[0]
	}
	return nil
}

// SetMessageControlType sets the message control type (TestMessage or LiveMessage)
func (nrm *NewReleaseMessage) SetMessageControlType(controlType string) {
	if nrm.MessageHeader != nil {
		nrm.MessageHeader.MessageControlType = controlType
	}
}

// AddComment adds a comment to the message header
func (nrm *NewReleaseMessage) AddComment(comment string) {
	if nrm.MessageHeader != nil {
		nrm.MessageHeader.Comment = comment
	}
}
