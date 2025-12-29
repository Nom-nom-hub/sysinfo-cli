package system

import (
	"testing"
)

// Tests for Windows process parsing functions (works on all platforms)

func TestParseCSVLineQuoted(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		name     string
	}{
		{
			input:    `"sysinfo.exe","1234","Console","1","12,345 K"`,
			expected: 5,
			name:     "Standard tasklist CSV",
		},
		{
			input:    `"explorer.exe","2048","Console","1","98,765 K"`,
			expected: 5,
			name:     "Explorer process",
		},
		{
			input:    `"Name with spaces.exe","512","Console","0","5,432 K"`,
			expected: 5,
			name:     "Process with spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := parseCSVLine(tt.input)
			if len(fields) != tt.expected {
				t.Errorf("parseCSVLine got %d fields, want %d", len(fields), tt.expected)
			}
		})
	}
}

func TestParseMemoryStringKB(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		name     string
	}{
		{
			input:    "12,345 K",
			expected: 12.0458,
			name:     "Standard memory format",
		},
		{
			input:    "1,234,567 K",
			expected: 1205.8301,
			name:     "Large memory format",
		},
		{
			input:    "512 K",
			expected: 0.5,
			name:     "Small memory",
		},
		{
			input:    "0 K",
			expected: 0.0,
			name:     "Zero memory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseMemoryString(tt.input)
			// Allow floating point differences (within 1MB or 0.1% of expected)
			diff := result - tt.expected
			relDiff := diff / tt.expected
			if (diff < -1.0 || diff > 1.0) && (relDiff < -0.001 || relDiff > 0.001) {
				t.Errorf("parseMemoryString(%q) = %.4f, want %.4f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseMemoryStringInvalid(t *testing.T) {
	tests := []struct {
		input string
		name  string
	}{
		{"invalid", "Non-numeric string"},
		{"", "Empty string"},
		{"ABC K", "Non-numeric with suffix"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseMemoryString(tt.input)
			if result != 0.0 {
				t.Errorf("parseMemoryString(%q) should return 0.0 for invalid input, got %.4f", tt.input, result)
			}
		})
	}
}
