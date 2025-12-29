package output

import (
	"testing"
)

func TestShouldUseColor(t *testing.T) {
	tests := []struct {
		mode     ColorMode
		expected bool
	}{
		{ColorOn, true},
		{ColorOff, false},
		{ColorAuto, isTerminal()}, // Depends on environment
	}

	for _, tt := range tests {
		t.Run(string(tt.mode), func(t *testing.T) {
			result := ShouldUseColor(tt.mode)
			if tt.mode != ColorAuto && result != tt.expected {
				t.Errorf("ShouldUseColor(%s) = %v, want %v", tt.mode, result, tt.expected)
			}
		})
	}
}

func TestGetColors(t *testing.T) {
	tests := []struct {
		mode ColorMode
		name string
	}{
		{ColorOn, "ColorOn should return colors"},
		{ColorOff, "ColorOff should return no colors"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := GetColors(tt.mode)

			if tt.mode == ColorOn {
				if c.Reset == "" {
					t.Error("ColorOn: Expected color codes, got empty")
				}
				if c.Header == "" {
					t.Error("ColorOn: Expected header color, got empty")
				}
			} else if tt.mode == ColorOff {
				if c.Reset != "" {
					t.Error("ColorOff: Expected no colors, got color codes")
				}
			}
		})
	}
}

func TestColorizePercent(t *testing.T) {
	tests := []struct {
		value    float64
		mode     ColorMode
		hasColor bool
	}{
		{50.0, ColorOn, true},   // Low usage
		{85.0, ColorOn, true},   // High usage (>80%)
		{95.0, ColorOn, true},   // Critical usage
		{50.0, ColorOff, false}, // No color when disabled
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := ColorizePercent(tt.value, tt.mode)

			if tt.hasColor {
				if result == "" {
					t.Errorf("ColorizePercent(%.1f) should return color code, got empty", tt.value)
				}
			} else {
				if result != "" {
					t.Errorf("ColorizePercent(%.1f) with ColorOff should return empty, got %q", tt.value, result)
				}
			}
		})
	}
}

func TestFormatHeader(t *testing.T) {
	text := "Test Header"

	t.Run("with color", func(t *testing.T) {
		result := FormatHeader(text, ColorOn)
		if result == text {
			t.Error("FormatHeader with ColorOn should add color codes")
		}
	})

	t.Run("without color", func(t *testing.T) {
		result := FormatHeader(text, ColorOff)
		if result != text {
			t.Errorf("FormatHeader with ColorOff should return plain text, got %q", result)
		}
	})
}

func TestFormatValue(t *testing.T) {
	text := "Test Value"

	t.Run("with color", func(t *testing.T) {
		result := FormatValue(text, ColorOn)
		if result == text {
			t.Error("FormatValue with ColorOn should add color codes")
		}
	})

	t.Run("without color", func(t *testing.T) {
		result := FormatValue(text, ColorOff)
		if result != text {
			t.Errorf("FormatValue with ColorOff should return plain text, got %q", result)
		}
	})
}
