# DDEX ERN 4.3 Go Package

A Go package for creating and parsing DDEX ERN 4.3 (Electronic Release Notification) messages, specifically optimized for YouTube content delivery.

## Overview

This package provides a fluent builder API for easily creating DDEX ERN 4.3 compliant XML messages for video releases. It's designed to simplify the process of generating complex DDEX messages with proper structure and validation.

## Features

- ✅ Full DDEX ERN 4.3 specification support
- ✅ Fluent builder API for easy message construction
- ✅ YouTube-specific configurations
- ✅ Type-safe Go structs with proper XML marshaling
- ✅ Support for:
  - Video resources with technical details
  - Image resources (cover art, screenshots)
  - Multiple parties (artists, labels, rights holders)
  - Complex deal structures
  - Territory-specific metadata

## Installation

```bash
go get github.com/yourusername/ddex
```

## Quick Start

Here's a simple example creating a video single release for YouTube:

```go
package main

import (
    "log"
    "github.com/yourusername/ddex/pkg/ddex"
)

func main() {
    // Create builder
    builder := ddex.NewDDEXBuilder()

    // Set up message header
    builder.WithMessageHeader(
        "MSG001",           // Message ID
        "THREAD001",        // Thread ID
        "PADPIDA1234567",   // Your DPID
        "My Label",         // Your name
    ).AddYouTubeRecipient()

    // Add artist party
    builder.AddParty("P1", "Artist Name", "Name, Artist")

    // Add video resource
    builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
        WithISRC("USXX12300001").
        WithTitle("My Video Title", "Subtitle").
        WithDisplayArtistName("Artist Name").
        WithArtist("P1", "MainArtist", 1).
        WithDuration("PT3M30S").
        WithCreationDate("2024-01-01", false).
        WithParentalWarning("NoAdviceAvailable").
        WithTechnicalDetails("T1", "video.mp4").
        Done()

    // Add release
    builder.AddRelease("R1", "VideoSingle").
        WithUPC("1234567890123").
        WithTitle("My Video Title", "").
        WithDisplayArtistName("Artist Name").
        WithArtist("P1", 1).
        WithGenre("Pop", "Worldwide").
        AddResourceGroup("", 1).
        AddContentItem(1, "A1").
        Done().
        Done()

    // Add deal
    builder.AddDeal("R1").
        WithTerritory("Worldwide").
        WithValidityPeriod("2024-01-01").
        AddCommercialModel("AdvertisementSupportedModel").
        AddUseType("Stream").
        Done()

    // Write to file
    if err := builder.WriteToFile("release.xml"); err != nil {
        log.Fatal(err)
    }
}
```

## Complete Example: Video Single with Related Resource

Here's a complete example that creates a video single with a related resource (matching the structure of a typical YouTube DDEX delivery):

```go
package main

import (
    "log"
    "github.com/yourusername/ddex/pkg/ddex"
)

func main() {
    // Create builder
    builder := ddex.NewDDEXBuilder()

    // Set up message header
    builder.WithMessageHeader(
        "VideoSingleTest",        // Message ID
        "VideoSingleTest",        // Thread ID
        "Your DPID",              // Your DPID
        "Your party name",        // Your company name
    ).AddYouTubeRecipient()

    // Add parties (artist and label)
    builder.AddParty("PJohnDoe", "John Doe", "Doe, John").
        AddParty("PACME", "ACME music", "")

    // Add video resource
    builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
        WithISRC("QZ6GL1732999").
        WithTitle("Video display title", "Video subtitle").
        WithDisplayArtistName("John Doe").
        WithArtist("PJohnDoe", "MainArtist", 1).
        WithRightsController("PACME", 100.00).
        WithDuration("PT3M10S").
        WithCreationDate("2023-01-01", true).
        WithParentalWarning("NoAdviceAvailable").
        WithPLine(2023, "(P) 2023 Some Pline text").
        WithTechnicalDetails("T1", "vid.mpg").
        AddProprietaryId("YOUTUBE:CHANNEL_ID", "UCQ0qe7vLz7uE_-sdtM9WB_w").
        Done()

    // Add image resource (video screen capture)
    builder.AddImage("A2", "VideoScreenCapture").
        WithProprietaryId("Your DPID", "VidCapPID").
        WithParentalWarning("NotExplicit").
        WithTechnicalDetails("T3", "vidCap.jpg").
        Done()

    // Add release with related resource
    builder.AddRelease("R0", "VideoSingle").
        WithICPN("2023121700021").
        WithTitle("Video display title", "Video").
        WithDisplayArtistName("John Doe").
        WithArtist("PJohnDoe", 1).
        WithLabel("PACME", "Worldwide").
        WithPLine(2023, "(P) 2023 Some Pline text").
        WithCLine(2023, "(C) 2023 Some CLine text").
        WithDuration("PT6M36S").
        WithGenre("Pop", "Worldwide").
        WithParentalWarning("NoAdviceAvailable").
        AddRelatedResource("HasContentFrom", "US1111111111").  // Related resource
        AddResourceGroup("Component 1", 1).
            AddContentItem(1, "A1").
            AddLinkedResource("VideoScreenCapture", "A2").
            Done().
        Done()

    // Add deal
    builder.AddDeal("R0").
        WithTerritory("Worldwide").
        WithValidityPeriod("2023-12-01").
        AddCommercialModel("SubscriptionModel").
        AddCommercialModel("AdvertisementSupportedModel").
        AddUseType("NonInteractiveStream").
        AddUseType("OnDemandStream").
        AddUseType("Stream").
        Done()

    // Write to file
    if err := builder.WriteToFile("video_single_with_related_resource.xml"); err != nil {
        log.Fatal(err)
    }
    
    log.Println("Successfully created video_single_with_related_resource.xml")
}
```

This example demonstrates:
- Message header with YouTube as recipient
- Multiple parties (artist and label)
- Video resource with technical details and YouTube channel ID
- Image resource for video screen capture
- **Related resource** indicating content source (e.g., underlying music track)
- Resource group linking video and image
- Multi-territory deal with various commercial models

## Building a Complete Video Release

### 1. Initialize the Builder

```go
builder := ddex.NewDDEXBuilder()
```

### 2. Set Message Header

```go
builder.WithMessageHeader(
    "MESSAGE_ID",        // Unique message identifier
    "THREAD_ID",         // Message thread identifier
    "YOUR_DPID",         // Your DDEX Party ID
    "Your Company Name", // Your company name
).AddYouTubeRecipient() // Adds YouTube as recipient
```

### 3. Add Parties

Parties represent all entities involved (artists, labels, etc.):

```go
builder.AddParty("PART_REF", "Full Name", "Indexed Name")
```

Example:
```go
builder.AddParty("P_ARTIST", "John Doe", "Doe, John")
builder.AddParty("P_LABEL", "ACME Records", "")
```

### 4. Add Video Resource

```go
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    WithISRC("ISRC_CODE").
    WithTitle("Video Title", "Subtitle").
    WithDisplayArtistName("Artist Name").
    WithArtist("PARTY_REF", "MainArtist", 1).
    WithRightsController("LABEL_REF", 100.00).
    WithDuration("PT3M10S").                    // ISO 8601 duration
    WithCreationDate("2024-01-01", false).      // Date and isApproximate flag
    WithParentalWarning("NoAdviceAvailable").
    WithPLine(2024, "(P) 2024 Label Name").
    WithTechnicalDetails("T1", "video.mp4").
    AddKeywords("music video", "pop", "rock").          // Add keywords for search/display
    AddProprietaryId("YOUTUBE:CHANNEL_ID", "UCxxxxxxx").
    Done()
```

#### Setting Artist Information

Artist information consists of two parts - the display name and the artist reference:

```go
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    WithDisplayArtistName("Artist Name").           // How the artist name appears
    WithArtist("PARTY_REF", "MainArtist", 1).      // Reference to party definition
    Done()

// Multiple artists with separate references
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    WithDisplayArtistName("Artist One feat. Artist Two").
    WithArtist("PARTY_REF_1", "MainArtist", 1).
    WithArtist("PARTY_REF_2", "FeaturedArtist", 2).
    Done()
```

#### Adding Keywords

Keywords enhance user experiences by providing DSPs with information for content display and music searches:

```go
// Add simple keywords (worldwide territory)
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    AddKeywords("music video", "pop", "rock", "ballad").
    Done()

// Add keywords with specific territory and language
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    AddKeywordsWithTerritory("US", "en", "american", "billboard").
    AddKeywordsWithTerritory("GB", "en", "british", "uk charts").
    Done()
```

**Important**: Each keyword must be in a separate tag. Do not combine multiple keywords in a single tag.

#### Adding Marketing Comments

Marketing comments provide information about the promotion and marketing of releases:

```go
// Add simple marketing comment (worldwide territory)
release := builder.AddRelease("R1", "Single").
    WithTitle("My Song", "").
    WithMarketingCommentSimple("A powerful ballad about love and loss").
    Done()

// Add detailed marketing comment with territory and language
release := builder.AddRelease("R1", "Album").
    WithTitle("My Album", "").
    WithMarketingComment("This compelling album explores themes of human connection through ten original compositions, blending elements of pop, rock, and soul.", "Worldwide", "en").
    WithMarketingComment("Un album fascinant qui explore les thèmes de la connexion humaine.", "FR", "fr").
    Done()
```

**Benefits**: 
- Provides marketing context for releases
- Supports promotional content distribution
- Supports multiple languages and territories
- ERN 4.3 compliant implementation (ddexC:MarketingComment)

#### Adding AI Content Metadata

Specify AI contribution types for releases using the ContainsAI element:

```go
// AI generated content
release := builder.AddRelease("R1", "Single").
    WithTitle("AI Generated Song", "").
    WithContainsAI("AIGenerated").
    Done()

// AI assisted content
release := builder.AddRelease("R1", "Album").
    WithTitle("Hybrid Album", "").
    WithContainsAI("AIAssisted").
    Done()

// No AI content
release := builder.AddRelease("R1", "Album").
    WithTitle("Human Created Album", "").
    WithContainsAI("NoAI").
    Done()
```

**Common Values**:
- `AIGenerated`: Content created entirely by AI
- `AIAssisted`: Human-created content with AI assistance
- `AITraining`: Content used for AI training purposes
- `NoAI`: No AI involvement in creation

**Benefits**:
- Transparency about AI involvement in content creation
- Compliance with platform AI disclosure requirements
- Helps streaming services categorize content appropriately
- ERN 4.3 compliant implementation

### 5. Add Image Resources (Optional)

```go
builder.AddImage("A2", "VideoScreenCapture").
    WithProprietaryId("YOUR_DPID", "IMAGE_ID").
    WithParentalWarning("NotExplicit").
    WithTechnicalDetails("T2", "screenshot.jpg").
    Done()
```

### 6. Add Release

```go
builder.AddRelease("R1", "VideoSingle").
    WithUPC("123456789012").          // or WithEAN("1234567890123")
    WithTitle("Release Title", "Subtitle").
    WithDisplayArtistName("Artist Name").
    WithArtist("PARTY_REF", 1).
    WithLabel("LABEL_REF", "Worldwide").
    WithPLine(2024, "(P) 2024 Label").
    WithCLine(2024, "(C) 2024 Copyright Holder").
    WithDuration("PT3M10S").
    WithGenre("Pop", "Worldwide").
    WithParentalWarning("NoAdviceAvailable").
    AddRelatedResource("HasContentFrom", "US1111111111").  // Optional: adds a related resource
    AddResourceGroup("Component 1", 1).
        AddContentItem(1, "A1").
        AddLinkedResource("VideoScreenCapture", "A2").
        Done().
    Done()
```

#### Genre and SubGenre

Set genre information for releases with optional subgenres:

```go
// Simple genre
builder.AddRelease("R1", "VideoSingle").
    WithGenre("Rock", "Worldwide").
    Done()

// Genre with subgenre
builder.AddRelease("R1", "VideoSingle").
    WithGenreAndSubGenre("Blues", "Bluesrock", "Worldwide").
    Done()

// Multiple genres (no hierarchy implied)
builder.AddRelease("R1", "VideoSingle").
    WithGenreAndSubGenre("Pop", "Synthpop", "Worldwide").
    WithGenre("Electronic", "US").
    Done()
```

**Note**: Multiple genres do not imply hierarchy. Each represents an equally valid classification.

#### Release Identifiers

Set UPC, EAN, or other identifiers for the release:

```go
// Using UPC (12 digits)
builder.AddRelease("R1", "VideoSingle").
    WithUPC("123456789012").
    Done()

// Using EAN (13 digits)
builder.AddRelease("R1", "VideoSingle").
    WithEAN("1234567890123").
    Done()

// Using GRid
builder.AddRelease("R1", "VideoSingle").
    WithGRid("A12425GABC1234002M").
    Done()
```

#### Related Resources

To indicate that a release uses content from another resource (e.g., a music track), use `AddRelatedResource`:

```go
builder.AddRelease("R1", "VideoSingle").
    WithUPC("123456789012").
    // ... other details ...
    AddRelatedResource("HasContentFrom", "US1111111111").  // ISRC of the related resource
    AddResourceGroup("Component 1", 1).
        // ... content items ...
        Done().
    Done()
```

Common relationship types:
- `HasContentFrom` - The release contains content from another resource
- `IsRelatedTo` - General relationship

### 7. Add Deal

```go
builder.AddDeal("R1").
    WithTerritory("Worldwide").
    WithValidityPeriod("2024-01-01").
    AddCommercialModel("SubscriptionModel").
    AddCommercialModel("AdvertisementSupportedModel").
    AddUseType("Stream").
    AddUseType("OnDemandStream").
    Done()
```

### 8. Generate XML

```go
// Write to file
err := builder.WriteToFile("release.xml")

// Or get as bytes
xmlBytes, err := builder.ToXML()

// Or get the message struct
message := builder.Build()
```

## Common Values

### Video Types
- `ShortFormMusicalWorkVideo` - Music videos under 10 minutes
- `LongFormMusicalWorkVideo` - Music videos over 10 minutes

### Release Types
- `VideoSingle` - Single video release
- `VideoAlbum` - Multiple video compilation

### Parental Warning Types
- `NoAdviceAvailable`
- `NotExplicit`
- `Explicit`

### Commercial Model Types
- `AdvertisementSupportedModel`
- `SubscriptionModel`
- `PayAsYouGoModel`
- `FreeOfChargeModel`

### Use Types
- `Stream`
- `OnDemandStream`
- `NonInteractiveStream`
- `ConditionalDownload`
- `PermanentDownload`

## Duration Format

Durations use ISO 8601 format: `PT[hours]H[minutes]M[seconds]S`

Examples:
- `PT3M30S` = 3 minutes 30 seconds
- `PT1H5M20S` = 1 hour 5 minutes 20 seconds
- `PT45S` = 45 seconds

## Date Format

Dates use ISO 8601 format: `YYYY-MM-DD`

Example: `2024-01-15`

## Identifiers

### ISRC (International Standard Recording Code)
Format: `CC-XXX-YY-NNNNN`
- CC = Country code (2 letters)
- XXX = Registrant code (3 alphanumeric)
- YY = Year (2 digits)
- NNNNN = Designation code (5 digits)

Example: `USRC17607839`

### UPC/EAN (Release Identifier)
- UPC: 12 digits
- EAN: 13 digits

Example: `123456789012`

### DPID (DDEX Party Identifier)
Format: `PADPIDA` + alphanumeric
Example: `PADPIDA2013020802I` (YouTube's DPID)

## Advanced Usage

### Custom Recipients

```go
builder.AddRecipient("PADPIDA1234567", "Recipient Name")
```

### Multiple Resource Groups

```go
builder.AddRelease("R1", "VideoAlbum").
    WithUPC("123456789012").
    // ... other details ...
    AddResourceGroup("Disc 1", 1).
        AddContentItem(1, "A1").
        AddContentItem(2, "A2").
        Done().
    AddResourceGroup("Disc 2", 2).
        AddContentItem(1, "A3").
        Done().
    Done()
```

### Territory-Specific Deals

```go
builder.AddDeal("R1").
    WithTerritory("US").  // US only
    // ... other details ...
    Done()

builder.AddDeal("R1").
    WithTerritory("GB").  // UK only
    // ... other details ...
    Done()
```

## Error Handling

The builder returns errors when writing files:

```go
if err := builder.WriteToFile("release.xml"); err != nil {
    log.Fatalf("Failed to create DDEX file: %v", err)
}
```

## Validation

The package ensures:
- Required fields are present
- Proper XML structure
- Correct namespaces and schema locations
- Valid reference relationships between elements

## Utility Functions

The package includes utility functions for common tasks:

### ID Generation

```go
import "github.com/yourusername/ddex/pkg/ddex"

// Generate unique IDs
messageID := ddex.GenerateMessageID("MSG")        // MSG_20241108150000_a1b2c3d4
threadID := ddex.GenerateThreadID("THR")          // THR_20241108_a1b2c3d4e5f6
reference := ddex.GenerateReference("RES")        // RES_a1b2c3d4e5f6g7h8
```

### Validation

```go
// Validate identifiers
isValid := ddex.ValidateUPC("123456789012")       // UPC validation
isValid = ddex.ValidateEAN("1234567890123")       // EAN validation
isValid = ddex.ValidateISRC("USRC17607839")       // ISRC validation
isValid = ddex.ValidateDPID("PADPIDA2013020802I") // DPID validation
```

### Duration Formatting

```go
// Convert seconds to ISO 8601 duration
duration := ddex.FormatDuration(210)              // Returns "PT3M30S"

// Parse ISO 8601 duration to seconds
seconds, err := ddex.ParseDuration("PT3M30S")     // Returns 210
```

### Date Formatting

```go
import "time"

// Format dates for DDEX
date := ddex.FormatDate(time.Now())               // Returns "2024-11-08"
dateTime := ddex.FormatDateTime(time.Now())       // Returns "2024-11-08T15:30:45"
```

## Examples

See the `examples/` directory for complete working examples:

- `create_video_single.go` - Complete video single release
- `create_video_album.go` - Multi-video release
- `parse_example.go` - Parsing existing DDEX files

## DDEX Resources

- [DDEX Official Site](https://ddex.net/)
- [ERN 4.3 Specification](https://ddex.net/standards/ern/43/)
- [YouTube DDEX Guide](https://support.google.com/youtube/answer/7127884)

## License

MIT License - See LICENSE file for details

## Contributing

Contributions welcome! Please open an issue or submit a pull request.

## Support

For issues and questions:
- GitHub Issues: [your-repo-url]
- DDEX Knowledge Base: https://kb.ddex.net/
