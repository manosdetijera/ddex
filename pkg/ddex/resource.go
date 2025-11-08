package ddex

import "encoding/xml"

// ResourceList lists all Resources composites in a release
type ResourceList struct {
	XMLName        xml.Name         `xml:"ResourceList"`
	SoundRecording []SoundRecording `xml:"SoundRecording,omitempty"`
	Video          []Video          `xml:"Video,omitempty"`
	Image          []Image          `xml:"Image,omitempty"`
	Text           []Text           `xml:"Text,omitempty"`
}

// Video represents a video resource
type Video struct {
	XMLName                  xml.Name                   `xml:"Video"`
	ResourceReference        string                     `xml:"ResourceReference"`
	Type                     string                     `xml:"Type,omitempty"`
	VideoEdition             []VideoEdition             `xml:"VideoEdition,omitempty"`
	DisplayTitleText         DisplayTitleText           `xml:"DisplayTitleText"`
	DisplayTitle             DisplayTitle               `xml:"DisplayTitle"`
	DisplayArtistName        []string                   `xml:"DisplayArtistName,omitempty"`
	DisplayArtist            []DisplayArtist            `xml:"DisplayArtist,omitempty"`
	ResourceRightsController []ResourceRightsController `xml:"ResourceRightsController,omitempty"`
	Duration                 string                     `xml:"Duration,omitempty"`
	CreationDate             CreationDate               `xml:"CreationDate,omitempty"`
	ParentalWarningType      string                     `xml:"ParentalWarningType,omitempty"`
}

// VideoEdition represents different editions of a video
type VideoEdition struct {
	XMLName          xml.Name                `xml:"VideoEdition,omitempty"`
	ResourceId       []VideoId               `xml:"ResourceId,omitempty"`
	PLine            []PLine                 `xml:"PLine,omitempty"`
	TechnicalDetails []TechnicalVideoDetails `xml:"TechnicalDetails,omitempty"`
}

// VideoId represents video identification
type VideoId struct {
	XMLName       xml.Name        `xml:"ResourceId"`
	ISRC          string          `xml:"ISRC,omitempty"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// Image represents an image resource
type Image struct {
	XMLName             xml.Name                `xml:"Image"`
	ResourceReference   string                  `xml:"ResourceReference"`
	Type                string                  `xml:"Type,omitempty"`
	ResourceId          []ImageId               `xml:"ResourceId,omitempty"`
	ParentalWarningType string                  `xml:"ParentalWarningType,omitempty"`
	TechnicalDetails    []TechnicalImageDetails `xml:"TechnicalDetails,omitempty"`
}

// ImageId represents image identification
type ImageId struct {
	XMLName       xml.Name        `xml:"ResourceId"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

// SoundRecording represents an audio resource
type SoundRecording struct {
	XMLName           xml.Name         `xml:"SoundRecording"`
	ResourceReference string           `xml:"ResourceReference"`
	Type              string           `xml:"Type,omitempty"`
	ResourceId        []ResourceID     `xml:"ResourceId,omitempty"`
	DisplayTitleText  DisplayTitleText `xml:"DisplayTitleText"`
	DisplayTitle      DisplayTitle     `xml:"DisplayTitle"`
}

// Text represents a text resource
type Text struct {
	XMLName           xml.Name         `xml:"Text"`
	ResourceReference string           `xml:"ResourceReference"`
	Type              string           `xml:"Type,omitempty"`
	ResourceId        []ResourceID     `xml:"ResourceId,omitempty"`
	DisplayTitleText  DisplayTitleText `xml:"DisplayTitleText"`
}

// ResourceRightsController represents rights controller for a resource
type ResourceRightsController struct {
	XMLName                        xml.Name               `xml:"ResourceRightsController"`
	RightsControllerPartyReference string                 `xml:"RightsControllerPartyReference"`
	RightsControlType              string                 `xml:"RightsControlType,omitempty"`
	RightSharePercentage           string                 `xml:"RightSharePercentage,omitempty"`
	DelegatedUsageRights           []DelegatedUsageRights `xml:"DelegatedUsageRights,omitempty"`
}

// DelegatedUsageRights represents delegated rights
type DelegatedUsageRights struct {
	XMLName                     xml.Name `xml:"DelegatedUsageRights"`
	UseType                     []string `xml:"UseType"`
	TerritoryOfRightsDelegation []string `xml:"TerritoryOfRightsDelegation,omitempty"`
}

// Technical details types
type TechnicalVideoDetails struct {
	XMLName                           xml.Name                  `xml:"TechnicalDetails"`
	TechnicalResourceDetailsReference string                    `xml:"TechnicalResourceDetailsReference"`
	DeliveryFile                      []AudioVisualDeliveryFile `xml:"DeliveryFile,omitempty"`
}

type TechnicalImageDetails struct {
	XMLName                           xml.Name `xml:"TechnicalDetails"`
	TechnicalResourceDetailsReference string   `xml:"TechnicalResourceDetailsReference"`
	File                              File     `xml:"File"`
}

type AudioVisualDeliveryFile struct {
	XMLName xml.Name `xml:"DeliveryFile"`
	Type    string   `xml:"Type"`
	File    File     `xml:"File"`
}

type File struct {
	XMLName xml.Name `xml:"File"`
	URI     string   `xml:"URI"`
}

// Supporting types
type CreationDate struct {
	XMLName       xml.Name `xml:"CreationDate"`
	Value         string   `xml:",chardata"`
	IsApproximate bool     `xml:"IsApproximate,attr,omitempty"`
}

type ProprietaryId struct {
	XMLName   xml.Name `xml:"ProprietaryId"`
	Namespace string   `xml:"Namespace,attr,omitempty"`
	Value     string   `xml:",chardata"`
}

type PLine struct {
	XMLName   xml.Name `xml:"PLine"`
	Year      int      `xml:"Year,omitempty"`
	PLineText string   `xml:"PLineText"`
}

type CLine struct {
	XMLName   xml.Name `xml:"CLine"`
	Year      int      `xml:"Year,omitempty"`
	CLineText string   `xml:"CLineText"`
}

type Genre struct {
	XMLName                 xml.Name `xml:"Genre"`
	GenreText               string   `xml:"GenreText"`
	ApplicableTerritoryCode string   `xml:"ApplicableTerritoryCode,attr,omitempty"`
}

type Contributor struct {
	XMLName        xml.Name `xml:"Contributor"`
	SequenceNumber int      `xml:"SequenceNumber,attr,omitempty"`
	PartyReference string   `xml:"PartyReference"`
	Role           []string `xml:"Role"`
}
