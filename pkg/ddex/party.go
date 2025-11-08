package ddex

import "encoding/xml"

// PartyList is a new composite in ERN 4 containing Party composites, consolidating data
// for all parties in the message—such as artists, writers, and labels.
type PartyList struct {
	XMLName xml.Name `xml:"PartyList"`
	Party   []Party  `xml:"Party"`
}

// Party represents a party (artist, writer, label, etc.) in the DDEX message
type Party struct {
	XMLName        xml.Name  `xml:"Party"`
	PartyReference string    `xml:"PartyReference"`
	PartyName      PartyName `xml:"PartyName"`
	PartyId        []PartyId `xml:"PartyId,omitempty"`
}

type PartyId struct {
	XMLName       xml.Name        `xml:"PartyId"`
	ISNI          string          `xml:"ISNI,omitempty"`
	DPID          string          `xml:"DPID,omitempty"`
	IpiNameNumber string          `xml:"IpiNameNumber,omitempty"`
	ProprietaryId []ProprietaryId `xml:"ProprietaryId,omitempty"`
}

type PartyName struct {
	XMLName         xml.Name `xml:"PartyName"`
	FullName        string   `xml:"FullName"`
	FullNameIndexed string   `xml:"FullNameIndexed,omitempty"`
}

// DisplayArtist represents how an artist should be displayed
type DisplayArtist struct {
	XMLName                 xml.Name                  `xml:"DisplayArtist"`
	SequenceNumber          int                       `xml:"SequenceNumber,attr,omitempty"`
	ArtistPartyReference    string                    `xml:"ArtistPartyReference"`
	DisplayArtistRole       string                    `xml:"DisplayArtistRole"`
	ArtisticRole            []string                  `xml:"ArtisticRole,omitempty"`
	TitleDisplayInformation []TitleDisplayInformation `xml:"TitleDisplayInformation,omitempty"`
}

// TitleDisplayInformation represents how artist info should be displayed in titles
type TitleDisplayInformation struct {
	XMLName            xml.Name `xml:"TitleDisplayInformation"`
	IsDisplayedInTitle bool     `xml:"IsDisplayedInTitle"`
	Prefix             string   `xml:"Prefix,omitempty"`
	Suffix             string   `xml:"Suffix,omitempty"`
}

// Location represents location information for a party
type Location struct {
	XMLName       xml.Name `xml:"Location"`
	CountryCode   string   `xml:"CountryCode,omitempty"`
	TerritoryCode string   `xml:"TerritoryCode,omitempty"`
	Address       *Address `xml:"Address,omitempty"`
}

// Address represents physical address information
type Address struct {
	XMLName     xml.Name `xml:"Address"`
	AddressLine []string `xml:"AddressLine,omitempty"`
	City        string   `xml:"City,omitempty"`
	PostalCode  string   `xml:"PostalCode,omitempty"`
	Country     string   `xml:"Country,omitempty"`
}

// ContactInformation represents contact details for a party
type ContactInformation struct {
	XMLName      xml.Name `xml:"ContactInformation"`
	EmailAddress []string `xml:"EmailAddress,omitempty"`
	PhoneNumber  []string `xml:"PhoneNumber,omitempty"`
	WebPage      []string `xml:"WebPage,omitempty"`
}

// NewParty creates a new Party with the specified reference and name
func NewParty(reference, name string) *Party {
	return &Party{
		PartyReference: reference,
		PartyName: PartyName{
			FullName: name,
		},
	}
}

// NewPartyWithIndexedName creates a new Party with full name and indexed name
func NewPartyWithIndexedName(reference, name, indexedName string) *Party {
	return &Party{
		PartyReference: reference,
		PartyName: PartyName{
			FullName:        name,
			FullNameIndexed: indexedName,
		},
	}
}
