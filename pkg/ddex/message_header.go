package ddex

import (
	"encoding/xml"
	"time"
)

// MessageHeader contains information about the sender and recipient, including their unique
// DDEX Party IDs (DPIDs), and a timestamp indicating when the message was created.
type MessageHeader struct {
	XMLName                xml.Name           `xml:"MessageHeader"`
	MessageThreadId        string             `xml:"MessageThreadId"`
	MessageId              string             `xml:"MessageId"`
	MessageFileName        string             `xml:"MessageFileName,omitempty"`
	MessageSender          *MessageSender     `xml:"MessageSender"`
	SentOnBehalfOf         string             `xml:"SentOnBehalfOf,omitempty"`
	MessageRecipient       []*MessageRecipient  `xml:"MessageRecipient"`
	MessageCreatedDateTime *DateTime          `xml:"MessageCreatedDateTime"`
	MessageControlType     string             `xml:"MessageControlType,omitempty"`
	MessageAuditTrail      *MessageAuditTrail `xml:"MessageAuditTrail,omitempty"`
	Comment                string             `xml:"Comment,omitempty"`
}

// MessageSender represents the sender of the DDEX message
type MessageSender struct {
	XMLName     xml.Name  `xml:"MessageSender"`
	PartyId     []PartyID `xml:"PartyId"`
	PartyName   []Name    `xml:"PartyName,omitempty"`
	TradingName string    `xml:"TradingName,omitempty"`
}

// MessageRecipient represents the recipient of the DDEX message
type MessageRecipient struct {
	XMLName     xml.Name  `xml:"MessageRecipient"`
	PartyId     []PartyID `xml:"PartyId"`
	PartyName   []Name    `xml:"PartyName,omitempty"`
	TradingName string    `xml:"TradingName,omitempty"`
}

// MessageAuditTrail represents audit trail information for the message
type MessageAuditTrail struct {
	XMLName                xml.Name                 `xml:"MessageAuditTrail"`
	MessageAuditTrailEvent []MessageAuditTrailEvent `xml:"MessageAuditTrailEvent"`
}

// MessageAuditTrailEvent represents a single audit trail event
type MessageAuditTrailEvent struct {
	XMLName                        xml.Name  `xml:"MessageAuditTrailEvent"`
	MessagingPartyReference        string    `xml:"MessagingPartyReference"`
	MessageAuditTrailEventDateTime *DateTime `xml:"MessageAuditTrailEventDateTime"`
	MessageAuditTrailEventTypeCode string    `xml:"MessageAuditTrailEventTypeCode"`
}

// NewMessageHeader creates a new MessageHeader with required fields for YouTube DDEX
func NewMessageHeader(threadId, messageId string, sender *MessageSender) *MessageHeader {
	now := &DateTime{Time: time.Now()}

	return &MessageHeader{
		MessageThreadId:        threadId,
		MessageId:              messageId,
		MessageSender:          sender,
		MessageCreatedDateTime: now,
		MessageControlType:     "TestMessage", // Default to test, should be changed for production
	}
}

func (m *MessageHeader) AddMessageRecipient(recipient *MessageRecipient) {
    m.MessageRecipient = append(m.MessageRecipient, recipient)
}

// NewMessageSender creates a new MessageSender with DPID for YouTube
func NewMessageSender(dpid, name string) *MessageSender {
	return &MessageSender{
		PartyId: []PartyID{
			{
				Value:     dpid,
				Namespace: "DPID",
			},
		},
		PartyName: []Name{
			{
				FullName: name,
			},
		},
	}
}

// NewMessageRecipient creates a new MessageRecipient for YouTube
func NewMessageRecipient(dpid, name string) *MessageRecipient {
	return &MessageRecipient{
		PartyId: []PartyID{
			{
				Value:     dpid, // YouTube's DPID
				Namespace: "DPID",
			},
		},
		PartyName: []Name{
			{
				FullName: name,
			},
		},
	}
}
