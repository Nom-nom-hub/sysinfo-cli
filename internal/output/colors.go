package output

import (
	"os"
)

// ColorMode defines the color output behavior
type ColorMode string

const (
	ColorAuto ColorMode = "auto"
	ColorOn   ColorMode = "on"
	ColorOff  ColorMode = "off"
)

// ANSIColors defines ANSI color codes
type ANSIColors struct {
	Reset   string
	Bold    string
	Header  string // Cyan for headers
	Value   string // Green for values
	Percent string // Yellow for percentages
	Alert   string // Red for alerts (>80%)
	Dim     string // Gray for secondary info
}

var (
	// Default ANSI colors
	colors = ANSIColors{
		Reset:   "\033[0m",
		Bold:    "\033[1m",
		Header:  "\033[36m", // Cyan
		Value:   "\033[32m", // Green
		Percent: "\033[33m", // Yellow
		Alert:   "\033[31m", // Red
		Dim:     "\033[90m", // Gray
	}

	// NoColors for when color is disabled
	noColors = ANSIColors{
		Reset:   "",
		Bold:    "",
		Header:  "",
		Value:   "",
		Percent: "",
		Alert:   "",
		Dim:     "",
	}
)

// ShouldUseColor determines if colors should be used based on mode
func ShouldUseColor(mode ColorMode) bool {
	switch mode {
	case ColorOn:
		return true
	case ColorOff:
		return false
	case ColorAuto:
		// Auto: use colors if stdout is a terminal
		return isTerminal()
	default:
		return isTerminal()
	}
}

// GetColors returns the appropriate color set
func GetColors(mode ColorMode) ANSIColors {
	if ShouldUseColor(mode) {
		return colors
	}
	return noColors
}

// isTerminal checks if stdout is connected to a terminal
func isTerminal() bool {
	stat, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	// Check if mode has character device flag (portable way)
	// Mode&0o20000 checks for character device on Unix-like systems
	return (stat.Mode() & 0o20000) != 0
}

// ColorizePercent colors a percentage based on threshold
func ColorizePercent(value float64, mode ColorMode) string {
	c := GetColors(mode)
	if c.Reset == "" {
		return ""
	}

	if value > 80 {
		return c.Alert
	}
	return c.Percent
}

// FormatHeader formats a header with color
func FormatHeader(text string, mode ColorMode) string {
	c := GetColors(mode)
	return c.Bold + c.Header + text + c.Reset
}

// FormatValue formats a value with color
func FormatValue(text string, mode ColorMode) string {
	c := GetColors(mode)
	return c.Value + text + c.Reset
}
