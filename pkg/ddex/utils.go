package ddex

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Utils provides utility functions for DDEX message creation and validation

// GenerateMessageID generates a unique message ID following DDEX conventions
func GenerateMessageID(prefix string) string {
	timestamp := time.Now().Format("20060102150405")
	randomBytes := make([]byte, 4)
	rand.Read(randomBytes)
	randomHex := fmt.Sprintf("%x", randomBytes)

	if prefix == "" {
		prefix = "MSG"
	}

	return fmt.Sprintf("%s_%s_%s", prefix, timestamp, randomHex)
}

// GenerateThreadID generates a unique thread ID following DDEX conventions
func GenerateThreadID(prefix string) string {
	timestamp := time.Now().Format("20060102")
	randomBytes := make([]byte, 6)
	rand.Read(randomBytes)
	randomHex := fmt.Sprintf("%x", randomBytes)

	if prefix == "" {
		prefix = "THR"
	}

	return fmt.Sprintf("%s_%s_%s", prefix, timestamp, randomHex)
}

// GenerateReference generates a unique reference ID for resources, releases, deals, etc.
func GenerateReference(prefix string) string {
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	randomHex := fmt.Sprintf("%x", randomBytes)

	if prefix == "" {
		prefix = "REF"
	}

	return fmt.Sprintf("%s_%s", prefix, randomHex)
}

// ValidateUPC validates a UPC (Universal Product Code)
func ValidateUPC(upc string) bool {
	// UPC should be 12 digits
	if len(upc) != 12 {
		return false
	}

	// Check if all characters are digits
	matched, _ := regexp.MatchString(`^\d{12}$`, upc)
	if !matched {
		return false
	}

	// Validate check digit using UPC algorithm
	sum := 0
	for i, char := range upc[:11] {
		digit := int(char - '0')
		if i%2 == 0 {
			sum += digit * 3
		} else {
			sum += digit
		}
	}

	checkDigit := (10 - (sum % 10)) % 10
	expectedCheckDigit := int(upc[11] - '0')

	return checkDigit == expectedCheckDigit
}

// ValidateEAN validates an EAN (European Article Number)
func ValidateEAN(ean string) bool {
	// EAN should be 13 digits
	if len(ean) != 13 {
		return false
	}

	// Check if all characters are digits
	matched, _ := regexp.MatchString(`^\d{13}$`, ean)
	if !matched {
		return false
	}

	// Validate check digit using EAN algorithm
	sum := 0
	for i, char := range ean[:12] {
		digit := int(char - '0')
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	checkDigit := (10 - (sum % 10)) % 10
	expectedCheckDigit := int(ean[12] - '0')

	return checkDigit == expectedCheckDigit
}

// ValidateISRC validates an ISRC (International Standard Recording Code)
func ValidateISRC(isrc string) bool {
	// ISRC format: CC-XXX-YY-NNNNN (12 characters without hyphens, 15 with)
	isrcClean := strings.ReplaceAll(strings.ToUpper(isrc), "-", "")

	if len(isrcClean) != 12 {
		return false
	}

	// First 2 characters: country code (letters)
	// Next 3 characters: registrant code (alphanumeric)
	// Next 2 characters: year (digits)
	// Last 5 characters: designation code (digits)
	pattern := `^[A-Z]{2}[A-Z0-9]{3}\d{7}$`
	matched, _ := regexp.MatchString(pattern, isrcClean)

	return matched
}

// ValidateISWC validates an ISWC (International Standard Musical Work Code)
func ValidateISWC(iswc string) bool {
	// ISWC format: T-DDD.DDD.DDD-C (where D=digit, C=check digit)
	iswcClean := strings.ReplaceAll(iswc, ".", "")
	iswcClean = strings.ReplaceAll(iswcClean, "-", "")

	if len(iswcClean) != 11 || !strings.HasPrefix(iswcClean, "T") {
		return false
	}

	// Validate format: T followed by 9 digits and 1 check digit
	pattern := `^T\d{10}$`
	matched, _ := regexp.MatchString(pattern, iswcClean)

	return matched
}

// ValidateDPID validates a DDEX Party ID
func ValidateDPID(dpid string) bool {
	// DPID format varies but typically 18 characters
	if len(dpid) < 10 || len(dpid) > 20 {
		return false
	}

	// Should contain only alphanumeric characters
	pattern := `^[A-Z0-9]+$`
	matched, _ := regexp.MatchString(pattern, dpid)

	return matched
}

// FormatDuration formats a duration in seconds to ISO 8601 duration format (PT3M30S)
func FormatDuration(seconds int) string {
	if seconds <= 0 {
		return "PT0S"
	}

	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	duration := "PT"
	if hours > 0 {
		duration += fmt.Sprintf("%dH", hours)
	}
	if minutes > 0 {
		duration += fmt.Sprintf("%dM", minutes)
	}
	if secs > 0 || (hours == 0 && minutes == 0) {
		duration += fmt.Sprintf("%dS", secs)
	}

	return duration
}

// ParseDuration parses an ISO 8601 duration format to seconds
func ParseDuration(duration string) (int, error) {
	if !strings.HasPrefix(duration, "PT") {
		return 0, fmt.Errorf("invalid duration format: %s", duration)
	}

	// Remove PT prefix
	d := duration[2:]

	totalSeconds := 0

	// Parse hours
	if idx := strings.Index(d, "H"); idx != -1 {
		var hours int
		fmt.Sscanf(d[:idx], "%d", &hours)
		totalSeconds += hours * 3600
		d = d[idx+1:]
	}

	// Parse minutes
	if idx := strings.Index(d, "M"); idx != -1 {
		var minutes int
		fmt.Sscanf(d[:idx], "%d", &minutes)
		totalSeconds += minutes * 60
		d = d[idx+1:]
	}

	// Parse seconds
	if idx := strings.Index(d, "S"); idx != -1 {
		var seconds int
		fmt.Sscanf(d[:idx], "%d", &seconds)
		totalSeconds += seconds
	}

	return totalSeconds, nil
}

// FormatDate formats a time.Time to ISO 8601 date format (YYYY-MM-DD)
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime formats a time.Time to ISO 8601 datetime format for DDEX
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")
}
