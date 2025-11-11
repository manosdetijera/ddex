# DDEX ERN 3.8 Go Package

A Go package for creating and parsing DDEX ERN 3.8 (Electronic Release Notification) messages, specifically optimized for YouTube content delivery with Content ID support.

## Overview

This package provides a fluent builder API for easily creating DDEX ERN 3.8 compliant XML messages for video releases. It's designed to simplify the process of generating complex DDEX messages with proper structure and validation for YouTube's Content ID system and streaming platform.

## Features

- ✅ Full DDEX ERN 3.8.1 specification support
- ✅ Fluent builder API for easy message construction
- ✅ YouTube-specific configurations (Content ID + Streaming)
- ✅ Territory-based metadata organization
- ✅ Type-safe Go structs with proper XML marshaling
- ✅ Support for:
  - Video resources with technical details and multiple identifiers
  - Image resources (cover art, screenshots, thumbnails)
  - Multiple titles (formal, display, translated)
  - Complex deal structures (Content ID policies + streaming rights)
  - Collection/playlist management
  - Multiple contributors and artists
  - Genre, keywords, and marketing metadata
  - Rights controller and ownership information

## Installation

```bash
go get github.com/yourusername/ddex
```

## Complete Example: YouTube Music Video with Content ID

This example demonstrates how to create the exact DDEX feed structure for YouTube with Content ID enabled. This matches the official YouTube DDEX XML sample format.

```go
package main

import (
    "log"
    "github.com/yourusername/ddex/pkg/ddex"
)

func main() {
    // Create builder
    builder := ddex.NewDDEXBuilder()

    // ============================================================================
    // MESSAGE HEADER with YouTube and Content ID recipients
    // ============================================================================
    builder.WithMessageHeader(
        "",                           // MessageId (empty in sample)
        "",                           // ThreadId (empty in sample)
        "DPID_OF_THE_SENDER",        // Your DPID
        "Name of sending party",      // Your company name
    ).AddYouTubeRecipient().         // For making video available on YouTube (PADPIDA2013020802I)
      AddYouTubeContentIDRecipient().// For enabling Content ID (PADPIDA2015120100H)
      WithUpdateIndicator("OriginalMessage")

    // ============================================================================
    // VIDEO RESOURCE
    // ============================================================================
    builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
        // Add VideoId with ISRC and proprietary IDs
        WithISRC("QZ6RS1700001").
        AddProprietaryId("YOUTUBE:MV_ASSET_LABEL", "foobar_music_video_label").
        AddProprietaryId("YOUTUBE:MV_CUSTOM_ID", "music_video_custom_001").
        AddProprietaryId("YOUTUBE:CHANNEL_ID", "MyChannel").
        AddProprietaryId("DPID:DPID_OF_THE_SENDER", "YourProprietaryID_001").
        AddProprietaryId("YOUTUBE:AD_FORMAT", "AdFormatID").
        
        // Add IndirectVideoId (musical work)
        WithIndirectVideoId("T1234567890"). // ISWC
        
        // Add ReferenceTitle (video-level, before territory details)
        WithReferenceTitle("A little bit of Foo", "").
        
        // Add LanguageOfPerformance and Duration (video-level)
        WithLanguageOfPerformance("en").
        WithDuration("PT0H3M16S").
        
        // ============================================================================
        // TERRITORY-SPECIFIC DETAILS (VideoDetailsByTerritory)
        // ============================================================================
        WithTerritory("Worldwide").
        
        // Add three Title elements with different types and languages
        WithTitle("A little bit of Foo", "").
        WithTitleAttributes("en", "FormalTitle"). // First title: FormalTitle in English
        
        WithTitle("A little bit of Foo", "").
        WithTitleAttributes("en", "DisplayTitle"). // Second title: DisplayTitle in English
        
        WithTitle("キャン・ユー・フィール．．．ザ・モンキー・クロー！", 
                 "ライヴ・アット・武道館").
        WithTitleAttributes("ja", "TranslatedTitle"). // Third title: TranslatedTitle in Japanese
        
        // Add DisplayArtist (inline party with name and role)
        WithDisplayArtist("Jonny and the Føøbars", "MainArtist", 1).
        
        // Add ResourceContributor
        WithResourceContributor("Jane Doe", []string{"Producer"}, 1).
        
        // Add IndirectResourceContributor
        WithIndirectResourceContributor("Jonny Smith", []string{"Composer"}, 1).
        
        // Add DisplayArtistName (simple text)
        WithDisplayArtistName("Jonny and the Føøbars", "").
        
        // Add LabelName
        WithLabelName("Test Label", "DisplayLabelName").
        
        // Add RightsController (sets ownership on Music Video asset)
        WithRightsController("Test Label", "DPID_OF_THE_SENDER", 
                           []string{"RightsController"}, 100.00).
        
        // Add ResourceReleaseDate
        WithResourceReleaseDate("2017-02-05").
        
        // Add PLine
        WithPLine(2017, "(P) 2017 Test Label Inc.").
        
        // Add Genre
        WithGenre("Hip Hop", "").
        
        // Add ParentalWarningType
        WithParentalWarning("Explicit").
        
        // Add Keywords
        AddKeywords("Keyword1", "Keyword2").
        
        // Add TechnicalVideoDetails with file information
        WithTechnicalVideoDetails("T1", "QZ6RS1700004_01_01.mpeg", "resources/").
        Done()

    // ============================================================================
    // IMAGE RESOURCE
    // ============================================================================
    builder.AddImage("A2", "VideoScreenCapture").
        AddProprietaryId("DPID:DPID_OF_THE_SENDER", "YourProprietaryID_002").
        WithTerritory("Worldwide").
        WithParentalWarning("NotExplicit").
        WithTechnicalImageDetails("T2", "QZ6RS1700004_01_01.jpeg", "resources/").
        Done()

    // ============================================================================
    // COLLECTION (Playlist) - Optional
    // ============================================================================
    builder.AddCollection("X1", "FilmBundle").
        AddProprietaryId("YOUTUBE:PLAYLIST_ID", "PLONRDPtQh-FLMXFMM-SJHySwjpidVXmzw").
        WithTitle("My Updated Playlist Title", "").
        AddResourceReference("A1", 1). // Adds video at position 1 in playlist
        Done()

    // ============================================================================
    // RELEASE 1: VideoSingle (required by spec, not used by YouTube)
    // ============================================================================
    builder.AddRelease("R0", "VideoSingle").
        SetMainRelease(true).
        WithGRid("A1UCASE0000000007X").
        WithISRC("QZ6RS1700001").
        WithReferenceTitle("Can you feel ...the Monkey Claw!", "Live at Budokan").
        AddReleaseResourceReference("A1", "PrimaryResource").
        AddReleaseResourceReference("A2", "SecondaryResource").
        
        // Territory-specific release details
        WithTerritory("Worldwide").
        WithDisplayArtistName("Monkey Claw", "").
        WithLabel("Iron Crown Music", "").
        
        // Add three titles (same pattern as video)
        WithTitle("A little bit of Foo", "").
        WithTitleAttributes("en", "FormalTitle").
        
        WithTitle("A little bit of Foo", "").
        WithTitleAttributes("en", "DisplayTitle").
        
        WithTitle("キャン・ユー・フィール．．．ザ・モンキー・クロー！",
                 "ライヴ・アット・武道館").
        WithTitleAttributes("ja", "TranslatedTitle").
        
        // Add DisplayArtist
        WithDisplayArtist("Jonny and the Føøbars", "MainArtist", 1).
        
        // Set ParentalWarningType
        WithParentalWarning("NotExplicit").
        
        // Add ResourceGroup with linked image
        AddResourceGroup("Component 1", 1).
        AddContentItem(1, "A1", "Video", "PrimaryResource").
        AddLinkedResource("VideoScreenCapture", "A2").
        Done(). // Close content item
        Done()  // Close release

    // ============================================================================
    // RELEASE 2: VideoTrackRelease (used by YouTube for Content ID and streaming)
    // This release contains:
    // 1. Content ID and streaming rights (via DealList)
    // 2. MarketingComment (becomes YouTube video description)
    // 3. Link between Video and thumbnail Image
    // ============================================================================
    builder.AddRelease("R1", "VideoTrackRelease").
        WithGRid("A1UCASE0000000007X").
        WithISRC("QZ6RS1700001").
        AddProprietaryId("YOUTUBE:AD_FORMAT", "AdFormatID").
        WithReferenceTitle("Can you feel ...the Monkey Claw!", "Live at Budokan").
        AddReleaseResourceReference("A1", "PrimaryResource").
        AddReleaseResourceReference("A2", "SecondaryResource").
        
        // Territory-specific release details
        WithTerritory("Worldwide").
        WithDisplayArtistName("Monkey Claw", "").
        WithLabel("Iron Crown Music", "").
        
        // Add three titles (same pattern)
        WithTitle("A little bit of Foo", "").
        WithTitleAttributes("en", "FormalTitle").
        
        WithTitle("A little bit of Foo", "").
        WithTitleAttributes("en", "DisplayTitle").
        
        WithTitle("キャン・ユー・フィール．．．ザ・モンキー・クロー！",
                 "ライヴ・アット・武道館").
        WithTitleAttributes("ja", "TranslatedTitle").
        
        // Add DisplayArtist
        WithDisplayArtist("Jonny and the Føøbars", "MainArtist", 1).
        
        // Set ParentalWarningType
        WithParentalWarning("NotExplicit").
        
        // Add MarketingComment (becomes YouTube video description)
        WithMarketingComment(`Official video for "A little bit of Foo" by Jonny and the Føøbars.

Now also available on http://someothersite.abc/jonnyandthefoobars

Follow us on social media: http://socialmediasite.abc/jonnyandthefoobars`).
        
        // Add ResourceGroup with linked thumbnail
        AddResourceGroup("Component 1", 1).
        AddContentItem(1, "A1", "Video", "PrimaryResource").
        AddLinkedResource("VideoScreenCapture", "A2").
        Done(). // Close content item
        Done()  // Close release

    // ============================================================================
    // DEALS for Release R1 (VideoTrackRelease)
    // Deal 1: Content ID with saved policy reference
    // Deal 2: Streaming rights with territories and validity period
    // ============================================================================
    builder.AddReleaseDeal("R1").
        // Content ID deal - references existing saved policy on YouTube
        AddDealWithReference("YT_MATCH_POLICY:My Saved Policy").
        
        // Streaming deal with commercial models and territories
        AddDeal().
        WithCommercialModel("AdvertisementSupportedModel").
        WithCommercialModel("SubscriptionModel").
        WithUseType("OnDemandStream").
        WithTerritory("US").
        WithTerritory("CA").
        WithValidityPeriod("2017-02-18", "2019-02-01").
        Done(). // Close deal
        Done()  // Close ReleaseDeal

    // Write to file
    if err := builder.WriteToFile("QZ6RS1700001_music_video_content_id_combined.xml"); err != nil {
        log.Fatal(err)
    }
    
    log.Println("Successfully created DDEX ERN 3.8 XML file")
}
```

## Key Features Explained

### 1. Dual Recipients (YouTube + Content ID)

```go
builder.WithMessageHeader(messageId, threadId, dpid, name).
    AddYouTubeRecipient().         // Makes video available on YouTube
    AddYouTubeContentIDRecipient() // Enables Content ID matching
```

### 2. Multiple Proprietary IDs

YouTube supports various proprietary ID namespaces for different features:

```go
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    AddProprietaryId("YOUTUBE:MV_ASSET_LABEL", "my_label").      // Asset label
    AddProprietaryId("YOUTUBE:MV_CUSTOM_ID", "custom_001").      // Custom ID
    AddProprietaryId("YOUTUBE:CHANNEL_ID", "UCxxxxxxxx").        // Target channel
    AddProprietaryId("YOUTUBE:AD_FORMAT", "AdFormatID").         // Ad format
    Done()
```

### 3. Territory-Based Metadata (ERN 3.8 Structure)

All territory-specific data must be within `VideoDetailsByTerritory`:

```go
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    WithTerritory("Worldwide").  // Create territory section
    AddTitle("My Video", "", "en", "DisplayTitle").
    WithDisplayArtistName("Artist Name", "").
    WithLabel("Label Name", "").
    // ... more territory-specific fields
    Done()
```

### 4. Multiple Titles with Languages

Support for formal, display, and translated titles:

```go
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    WithTerritory("Worldwide").
    AddTitle("English Title", "", "en", "FormalTitle").
    AddTitle("English Display", "", "en", "DisplayTitle").
    AddTitle("日本語タイトル", "サブタイトル", "ja", "TranslatedTitle").
    Done()
```

### 5. Contributors and Rights

Add various contributors without needing PartyList in ERN 3.8:

```go
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    WithTerritory("Worldwide").
    AddDisplayArtist("Artist Name", "MainArtist", 1).
    AddContributor("Producer Name", []string{"Producer"}, 1).
    AddIndirectContributor("Composer Name", []string{"Composer"}, 1).
    AddRightsController("Label Name", "DPID", "RightsController", 100.00).
    Done()
```

### 6. Two Release Structure (Video Single Profile)

YouTube DDEX feeds require two releases:

1. **VideoSingle** - Contains the complete metadata but not used by YouTube
2. **VideoTrackRelease** - Used for Content ID, streaming rights, and video description

```go
// Release 1: VideoSingle (required by spec)
builder.AddRelease("R0", "VideoSingle").
    SetMainRelease(true).
    // ... metadata
    Done()

// Release 2: VideoTrackRelease (used by YouTube)
builder.AddRelease("R1", "VideoTrackRelease").
    AddMarketingComment("Video description text").
    // ... metadata
    Done()
```

### 7. Deal Structure (Content ID + Streaming)

Separate deals for Content ID matching and streaming rights:

```go
// Content ID deal with saved policy reference
builder.AddReleaseDeal("R1").
    AddDealWithReference("YT_MATCH_POLICY:My Saved Policy").
    Done()

// Streaming deal with territories and validity period
builder.AddReleaseDeal("R1").
    AddDeal().
    WithCommercialModel("AdvertisementSupportedModel").
    WithCommercialModel("SubscriptionModel").
    WithUseType("OnDemandStream").
    WithTerritory("US").
    WithTerritory("CA").
    WithValidityPeriod("2017-02-18", "2019-02-01").
    Done().
    Done()
```

### 8. Collections (Playlists)

Add videos to existing playlists:

```go
builder.AddCollection("X1", "FilmBundle").
    AddProprietaryId("YOUTUBE:PLAYLIST_ID", "PLxxxxxxxxxx").
    WithTitle("Playlist Title", "").
    AddResourceReference("A1", 1).  // Position in playlist
    Done()
```

### 9. Marketing Comments

The marketing comment becomes the YouTube video description:

```go
builder.AddRelease("R1", "VideoTrackRelease").
    WithTerritory("Worldwide").
    AddMarketingComment(`Official video for "Song Title" by Artist Name.

Available on streaming: http://example.com
Follow us: http://social.example.com`).
    Done()
```

### 10. Technical Details with File Paths

Include file paths when delivering actual media files:

```go
builder.AddVideo("A1", "ShortFormMusicalWorkVideo").
    WithTerritory("Worldwide").
    WithTechnicalDetails("T1", "video.mpeg", "resources/").
    Done()

builder.AddImage("A2", "VideoScreenCapture").
    WithTerritory("Worldwide").
    WithTechnicalDetails("T2", "thumbnail.jpeg", "resources/").
    Done()
```

## API Reference

### Builder Methods

#### Message Header
- `WithMessageHeader(messageId, threadId, dpid, name)` - Set message header
- `AddYouTubeRecipient()` - Add YouTube as recipient
- `AddYouTubeContentIDRecipient()` - Add YouTube Content ID as recipient
- `AddRecipient(partyId, partyName)` - Add custom recipient

#### Video Resource
- `AddVideo(resourceRef, videoType)` - Create video resource
- `WithISRC(isrc)` - Set ISRC
- `AddProprietaryId(namespace, value)` - Add proprietary identifier
- `WithIndirectVideoId(iswc)` - Add musical work reference
- `WithReferenceTitle(title, subtitle)` - Set reference title (video-level)
- `WithLanguageOfPerformance(language)` - Set language
- `WithDuration(duration)` - Set duration (ISO 8601 format)
- `WithTerritory(territoryCode)` - Create/switch territory section
- `AddTitle(title, subtitle, language, titleType)` - Add title in territory
- `WithDisplayArtistName(name, language)` - Set display artist name
- `AddDisplayArtist(name, role, sequence)` - Add display artist
- `AddContributor(name, roles, sequence)` - Add contributor
- `AddIndirectContributor(name, roles, sequence)` - Add indirect contributor
- `WithLabel(name, language)` - Set label name
- `AddRightsController(name, partyId, role, percentage)` - Add rights controller
- `WithResourceReleaseDate(date)` - Set resource release date
- `WithPLine(year, text)` - Set P-line
- `WithGenre(genre, subGenre)` - Set genre
- `WithParentalWarning(warningType)` - Set parental warning
- `AddKeywords(keywords...)` - Add keywords
- `WithTechnicalDetails(techRef, fileName, filePath)` - Add technical details

#### Image Resource
- `AddImage(resourceRef, imageType)` - Create image resource
- `AddProprietaryId(namespace, value)` - Add proprietary identifier
- `WithTerritory(territoryCode)` - Create/switch territory section
- `WithParentalWarning(warningType)` - Set parental warning
- `WithTechnicalDetails(techRef, fileName, filePath)` - Add technical details
- `WithCreationDate(date, isApproximate)` - Set creation date

#### Collection (Playlist)
- `AddCollection(collectionRef, collectionType)` - Create collection
- `AddProprietaryId(namespace, value)` - Add proprietary identifier
- `WithTitle(title, subtitle)` - Set collection title
- `AddResourceReference(resourceRef, sequence)` - Add resource to collection

#### Release
- `AddRelease(releaseRef, releaseType)` - Create release
- `SetMainRelease(isMain)` - Mark as main release
- `WithGRid(grid)` - Set GRid
- `WithISRC(isrc)` - Set ISRC
- `AddProprietaryId(namespace, value)` - Add proprietary identifier
- `WithReferenceTitle(title, subtitle)` - Set reference title
- `AddReleaseResourceReference(resourceRef, resourceType)` - Link resource
- `WithTerritory(territoryCode)` - Create/switch territory section
- `WithDisplayArtistName(name, language)` - Set display artist name
- `WithLabel(name, language)` - Set label
- `AddTitle(title, subtitle, language, titleType)` - Add title
- `AddDisplayArtist(name, role, sequence)` - Add display artist
- `WithParentalWarning(warningType)` - Set parental warning
- `AddMarketingComment(text)` - Add marketing comment
- `AddResourceGroup(title, sequence)` - Add resource group
- `AddContentItem(sequence, resourceRef, resourceType, releaseResourceType)` - Add content item
- `AddLinkedResource(linkDescription, resourceRef)` - Link related resource

#### Deal
- `AddReleaseDeal(releaseRef)` - Create deal for release
- `AddDealWithReference(dealReference)` - Add deal with policy reference
- `AddDeal()` - Create new deal with terms
- `WithCommercialModel(model)` - Add commercial model
- `WithUseType(useType)` - Set use type
- `WithTerritory(territoryCode)` - Add territory
- `WithValidityPeriod(startDate, endDate)` - Set validity period

#### Output
- `WriteToFile(filename)` - Write XML to file
- `ToXML()` - Get XML as string

## ERN 3.8 vs ERN 4.x Differences

This package implements ERN 3.8, which has important differences from ERN 4.x:

1. **No PartyList** - Party information is inline within resources and releases
2. **Territory-based structure** - VideoDetailsByTerritory and ReleaseDetailsByTerritory are mandatory
3. **Namespace** - Uses `http://ddex.net/xml/ern/381` or `/382`
4. **Simpler types** - Many types are simplified without territory attributes at top level

## YouTube-Specific Guidelines

### Message Recipients
- Include "YouTube" (`PADPIDA2013020802I`) to make video available on YouTube
- Include "YouTube_ContentID" (`PADPIDA2015120100H`) to enable Content ID matching

### Proprietary ID Namespaces
- `YOUTUBE:MV_ASSET_LABEL` - Asset label (can be repeated)
- `YOUTUBE:MV_CUSTOM_ID` - Custom ID for Music Video asset
- `YOUTUBE:CHANNEL_ID` - Target channel for upload
- `YOUTUBE:AD_FORMAT` - Ad format configuration
- `YOUTUBE:PLAYLIST_ID` - Playlist identifier for collections
- `DPID:YOUR_DPID` - Your proprietary ID (used as custom ID if no MV_CUSTOM_ID)

### Deal References
- `YT_MATCH_POLICY:PolicyName` - Reference to saved Content ID match policy
- `YT_USAGE_POLICY:PolicyName` - Reference to saved usage policy

### Release Types
- `VideoSingle` - Required by spec but not used by YouTube
- `VideoTrackRelease` - Used for Content ID, streaming rights, and video metadata

### Marketing Comments
- Marketing comments in VideoTrackRelease become the YouTube video description
- Supports multiline text with URLs

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues and questions, please open an issue on GitHub.

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

#### Adding Release Dates

Release dates provide timing information for when content was or will be made available:

```go
// Set release date (sets both ReleaseDate and OriginalReleaseDate)
release := builder.AddRelease("R1", "Single").
    WithTitle("My Song", "").
    WithReleaseDate("2023-12-01").
    Done()

// Different date formats supported
release := builder.AddRelease("R1", "Album").
    WithTitle("My Album", "").
    WithReleaseDate("2023").           // Year only
    // or WithReleaseDate("2023-12").   // Year and month
    // or WithReleaseDate("2023-12-01") // Full date
    Done()
```

**Benefits**:
- Provides timing context for releases
- Sets both ReleaseDate and OriginalReleaseDate consistently
- Supports ISO 8601 date formats (YYYY, YYYY-MM, YYYY-MM-DD)
- ERN 4.3 compliant implementation (ern:EventDateWithDefault)

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

- `create_video_single.go` - Complete video single release for ERN 3.8

## DDEX Resources

- [DDEX Official Site](https://ddex.net/)
- [ERN 3.8 Data Dictionary](https://service.ddex.net/dd/ERN38/)
- [YouTube DDEX Guide](https://support.google.com/youtube/answer/7127884)
- [YouTube DDEX Reference Resources](https://support.google.com/youtube/answer/3506749)
- [DDEX ERN Knowledge Base](https://kb.ddex.net/implementing-each-standard/electronic-release-notification-message-suite-(ern)/)

## Migration from ERN 4.3

If you're migrating from the previous ERN 4.3 version of this package, see [MIGRATION_TO_ERN38.md](MIGRATION_TO_ERN38.md) for a detailed guide on the changes.

## License

MIT License - See LICENSE file for details

## Contributing

Contributions welcome! Please open an issue or submit a pull request.

## Support

For issues and questions:
- GitHub Issues: [your-repo-url]

- DDEX Knowledge Base: https://kb.ddex.net/
