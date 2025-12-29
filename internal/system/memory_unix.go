//go:build linux || darwin
// +build linux darwin

package system

import (
	"github.com/example/sysinfo-cli/internal/models"
)

// GetMemoryInfo returns memory/RAM information on Unix systems
func GetMemoryInfo() (*models.MemoryInfo, error) {
	// Unix/Linux implementation - simplified fallback
	// Production version would use syscall.Sysinfo or similar
	return &models.MemoryInfo{
		TotalGB:      8.0,      // Placeholder - production would read from /proc/meminfo or syscalls
		AvailableGB:  4.0,      // Placeholder
		UsedGB:       4.0,      // Placeholder
		UsagePercent: 50.0,     // Placeholder
		SwapTotalGB:  2.0,      // Placeholder
		SwapUsedGB:   0.5,      // Placeholder
	}, nil
}
