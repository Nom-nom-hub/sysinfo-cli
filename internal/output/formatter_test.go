package output

import (
	"strings"
	"testing"

	"github.com/example/sysinfo-cli/internal/models"
)

func TestFormatterJSON(t *testing.T) {
	formatter := NewFormatter("json", false)

	info := &models.OSInfo{
		Hostname: "test",
		OS:       "linux",
	}

	output, err := formatter.Format(info, "os")
	if err != nil {
		t.Fatalf("Format failed: %v", err)
	}

	if !strings.Contains(output, "test") {
		t.Errorf("Expected hostname in output, got: %s", output)
	}
}

func TestFormatterTable(t *testing.T) {
	formatter := NewFormatter("table", false)

	info := &models.MemoryInfo{
		TotalGB:      16.0,
		AvailableGB:  8.0,
		UsedGB:       8.0,
		UsagePercent: 50.0,
	}

	output, err := formatter.Format(info, "memory")
	if err != nil {
		t.Fatalf("Format failed: %v", err)
	}

	if !strings.Contains(output, "Memory Information") {
		t.Errorf("Expected header in output")
	}
}

func TestFormatterCSV(t *testing.T) {
	formatter := NewFormatter("csv", false)

	info := &models.CPUInfo{
		Cores:       4,
		Threads:     8,
		Model:       "Intel",
		FrequencyGHz: 2.4,
	}

	output, err := formatter.Format(info, "cpu")
	if err != nil {
		t.Fatalf("Format failed: %v", err)
	}

	if !strings.Contains(output, "cores") {
		t.Errorf("Expected CSV header")
	}
}

func TestFormatterInvalidFormat(t *testing.T) {
	formatter := NewFormatter("invalid", false)

	info := &models.OSInfo{}
	_, err := formatter.Format(info, "os")

	if err == nil {
		t.Errorf("Expected error for invalid format")
	}
}

func TestPrettyJSON(t *testing.T) {
	formatter := NewFormatter("json", true)

	info := &models.CPUInfo{Cores: 4}
	output, err := formatter.Format(info, "cpu")
	if err != nil {
		t.Fatalf("Format failed: %v", err)
	}

	if !strings.Contains(output, "\n") {
		t.Errorf("Expected pretty-printed JSON with newlines")
	}
}
