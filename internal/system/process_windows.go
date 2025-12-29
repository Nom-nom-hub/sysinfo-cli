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



// Stubs for non-Windows platforms (defined in process_unix.go)
func getProcessesLinux() []models.ProcessInfo {
	return []models.ProcessInfo{}
}

func getProcessesDarwin() []models.ProcessInfo {
	return []models.ProcessInfo{}
}

