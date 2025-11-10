package main

import (
	"fmt"
	"log"

	"github.com/manosdetijera/ddex/pkg/ddex"
)

func main() {
	// Create a new DDEX ERN 4.3 message builder
	builder := ddex.NewDDEXBuilder()

	// Set up message header
	builder.WithMessageHeader(
		"VideoSingleTest", // Message ID
		"VideoSingleTest", // Thread ID
		"Your DPID",       // Sender DPID
		"Your party name", // Sender name
	).AddYouTubeRecipient()

	// Add parties
	builder.AddParty("PJohnDoe", "John Doe", "Doe, John").
		AddParty("PACME", "ACME music", "")

	// Add video resource
	builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
		WithISRC("QZ6GL1732999").
		WithTitle("Video display title", "Video subtitle").
		WithDisplayArtistName("John Doe").
		WithArtist("PJohnDoe", "MainArtist", 1).
		WithRightsController("PACME", 100.00, []string{"Worldwide"}).
		WithDuration("PT3M10S").
		WithCreationDate("2023-01-01", true).
		WithParentalWarning("NoAdviceAvailable").
		WithPLine(2023, "(P) 2023 Some Pline text").
		WithTechnicalDetails("T1", "vid.mpg").
		AddKeywords("music video", "pop", "john doe").
		AddProprietaryId("YOUTUBE:CHANNEL_ID", "UCQ0qe7vLz7uE_-sdtM9WB_w").
		Done()

	// Add image resource (cover art)
	builder.AddImage("A2", "VideoScreenCapture").
		WithProprietaryId("Your DPID", "VidCapPID").
		WithParentalWarning("NotExplicit").
		WithTechnicalDetails("T3", "vidCap.jpg").
		Done()

	// Add release
	builder.AddRelease("R0", "VideoSingle").
		WithICPN("2023121700021").
		WithTitle("Video display title", "Video").
		WithDisplayArtistName("John Doe").
		WithArtist("PJohnDoe", "MainArtist", 1).
		WithLabel("PACME", "Worldwide").
		WithPLine(2023, "(P) 2023 Some Pline text").
		WithCLine(2023, "(C) 2023 Some CLine text").
		WithDuration("PT6M36S").
		WithGenreAndSubGenre("Pop", "Synthpop", "Worldwide").
		WithParentalWarning("NoAdviceAvailable").
		AddRelatedResource("HasContentFrom", "US1111111111").
		AddResourceGroup("Component 1", 1).
		AddContentItem(1, "A1").
		AddLinkedResource("VideoScreenCapture", "A2").
		Done().
		Done()

	// Add deal
	builder.AddDeal("R0").
		WithTerritories([]string{"Worldwide"}).
		WithValidityPeriod("2023-12-01").
		AddCommercialModel("SubscriptionModel").
		AddCommercialModel("AdvertisementSupportedModel").
		AddUseType("NonInteractiveStream").
		AddUseType("OnDemandStream").
		AddUseType("Stream").
		Done()

	// Write to file
	if err := builder.WriteToFile("video_single_example.xml"); err != nil {
		log.Fatalf("Failed to write DDEX file: %v", err)
	}

	fmt.Println("Successfully created video_single_example.xml")
}
