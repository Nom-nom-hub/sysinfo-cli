//go:build linux || darwin
// +build linux darwin

package system

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/example/sysinfo-cli/internal/models"
)

func getProcessesLinux() []models.ProcessInfo {
	var processes []models.ProcessInfo

	procDir, err := os.Open("/proc")
	if err != nil {
		return processes
	}
	defer procDir.Close()

	entries, err := procDir.Readdirnames(-1)
	if err != nil {
		return processes
	}

	for _, entry := range entries {
		pid, err := strconv.Atoi(entry)
		if err != nil {
			continue
		}

		statusPath := filepath.Join("/proc", entry, "status")

		name := getProcessName(statusPath, pid)

		// Get memory from status
		memoryMB := getProcessMemory(statusPath)

		// Get CPU percent (simplified)
		cpuPercent := 0.0

		processes = append(processes, models.ProcessInfo{
			PID:        pid,
			Name:       name,
			CPUPercent: cpuPercent,
			MemoryMB:   memoryMB,
		})
	}

	return processes
}

// getProcessesDarwin returns process information using ps command
// Parses output to extract PID, Name, CPU%, and memory usage
func getProcessesDarwin() []models.ProcessInfo {
	processes := make([]models.ProcessInfo, 0)

	// ps aux format: USER PID %CPU %MEM VSZ RSS STAT START TIME COMMAND
	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		return processes
	}

	lines := strings.Split(string(output), "\n")
	for i, line := range lines {
		if i == 0 {
			// Skip header
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 11 {
			continue
		}

		// Parse PID (field 1)
		pid, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}

		// Parse CPU% (field 2)
		cpuStr := fields[2]
		cpuPercent, _ := strconv.ParseFloat(cpuStr, 64)

		// Parse MEM% (field 3)
		memStr := fields[3]
		memPercent, _ := strconv.ParseFloat(memStr, 64)

		// Parse RSS in KB (field 5)
		rssStr := fields[5]
		rssKB, err := strconv.ParseFloat(rssStr, 64)
		if err != nil {
			rssKB = 0
		}
		memMB := rssKB / 1024.0

		// Command name is the last field (field 10+)
		cmdName := strings.Join(fields[10:], " ")
		// Extract just the executable name
		if idx := strings.LastIndex(cmdName, "/"); idx >= 0 {
			cmdName = cmdName[idx+1:]
		}

		processes = append(processes, models.ProcessInfo{
			PID:        pid,
			Name:       cmdName,
			CPUPercent: cpuPercent,
			MemoryMB:   memMB,
		})
	}

	return processes
}

func getProcessName(statusPath string, pid int) string {
	data, err := os.ReadFile(statusPath)
	if err != nil {
		return fmt.Sprintf("pid-%d", pid)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Name:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return fmt.Sprintf("pid-%d", pid)
}

func getProcessMemory(statusPath string) float64 {
	data, err := os.ReadFile(statusPath)
	if err != nil {
		return 0
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "VmRSS:") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				if kb, err := strconv.ParseFloat(parts[1], 64); err == nil {
					return kb / 1024.0 // Convert KB to MB
				}
			}
		}
	}

	return 0
}
