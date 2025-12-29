//go:build linux || darwin
// +build linux darwin

package system

import (
	"fmt"
	"os"
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

func getProcessesDarwin() []models.ProcessInfo {
	var processes []models.ProcessInfo
	// Simplified for demo - production would use syscalls or ps command
	return processes
}

func getProcessesWindows() []models.ProcessInfo {
	var processes []models.ProcessInfo
	// Simplified for demo - production would use WMI
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
