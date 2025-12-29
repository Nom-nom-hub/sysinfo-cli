//go:build windows
// +build windows

package system

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/example/sysinfo-cli/internal/models"
)

// getProcessesWindows returns process information using tasklist command
// Parses output to extract PID, Name, and memory usage
func getProcessesWindows() []models.ProcessInfo {
	processes := make([]models.ProcessInfo, 0)

	// Use tasklist with /v flag for verbose output including memory
	// Format: tasklist /v /fo csv /nh
	cmd := exec.Command("tasklist", "/v", "/fo", "csv", "/nh")
	output, err := cmd.Output()
	if err != nil {
		return processes
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Parse CSV format: "Name","PID","Session Name","Session#","Memory Usage"
		fields := parseCSVLine(line)
		if len(fields) < 5 {
			continue
		}

		name := strings.Trim(fields[0], "\"")
		pidStr := strings.Trim(fields[1], "\"")
		memStr := strings.Trim(fields[4], "\"")

		// Parse PID
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			continue
		}

		// Parse memory (format: "12,345 K" -> 12.345 MB)
		memMB := parseMemoryString(memStr)

		processes = append(processes, models.ProcessInfo{
			PID:        pid,
			Name:       name,
			MemoryMB:   memMB,
			CPUPercent: 0.0, // CPU usage not available via tasklist
		})
	}

	return processes
}

// parseCSVLine parses a CSV line accounting for quoted fields
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
func parseMemoryString(memStr string) float64 {
	// Format: "12,345 K" or "1,234,567 K"
	memStr = strings.TrimSpace(memStr)

	// Remove " K" suffix
	if strings.HasSuffix(memStr, " K") {
		memStr = strings.TrimSuffix(memStr, " K")
	}

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

// Stubs for non-Windows platforms (defined in process_unix.go)
func getProcessesLinux() []models.ProcessInfo {
	return []models.ProcessInfo{}
}

func getProcessesDarwin() []models.ProcessInfo {
	return []models.ProcessInfo{}
}

