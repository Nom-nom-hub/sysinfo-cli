package system

import (
	"strconv"
	"strings"
)

// parseCSVLine parses a CSV line accounting for quoted fields
// Used by Windows tasklist parser, available cross-platform for testing
func parseCSVLine(line string) []string {
	fields := make([]string, 0)
	current := ""
	inQuotes := false

	for i := 0; i < len(line); i++ {
		ch := line[i]

		if ch == '"' {
			inQuotes = !inQuotes
			current += string(ch)
		} else if ch == ',' && !inQuotes {
			fields = append(fields, current)
			current = ""
		} else {
			current += string(ch)
		}
	}

	if current != "" {
		fields = append(fields, current)
	}

	return fields
}

// parseMemoryString converts memory format "12,345 K" to MB
// Used by Windows tasklist parser, available cross-platform for testing
func parseMemoryString(memStr string) float64 {
	// Format: "12,345 K" or "1,234,567 K"
	memStr = strings.TrimSpace(memStr)

	// Remove " K" suffix (TrimSuffix handles missing suffix gracefully)
	memStr = strings.TrimSuffix(memStr, " K")

	// Remove commas
	memStr = strings.ReplaceAll(memStr, ",", "")

	// Parse as integer (KB)
	kb, err := strconv.ParseFloat(memStr, 64)
	if err != nil {
		return 0.0
	}

	// Convert KB to MB
	return kb / 1024.0
}
