package system

import (
	"runtime"
	"sort"

	"github.com/example/sysinfo-cli/internal/models"
)

// GetProcessInfo returns top processes by CPU or memory usage
func GetProcessInfo(sortBy string, limit int) ([]models.ProcessInfo, error) {
	var processes []models.ProcessInfo

	if runtime.GOOS == "linux" {
		processes = getProcessesLinux()
	} else if runtime.GOOS == "darwin" {
		// On macOS, would use BSD syscalls or ps command
		processes = getProcessesDarwin()
	} else if runtime.GOOS == "windows" {
		// On Windows, would use WMI or tasklist
		processes = getProcessesWindows()
	}

	// Sort by requested field
	if sortBy == "cpu" {
		sort.Slice(processes, func(i, j int) bool {
			return processes[i].CPUPercent > processes[j].CPUPercent
		})
	} else if sortBy == "memory" {
		sort.Slice(processes, func(i, j int) bool {
			return processes[i].MemoryMB > processes[j].MemoryMB
		})
	}

	// Limit results
	if len(processes) > limit {
		processes = processes[:limit]
	}

	return processes, nil
}
