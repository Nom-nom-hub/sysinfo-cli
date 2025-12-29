//go:build linux || darwin
// +build linux darwin

package system

import (
	"github.com/example/sysinfo-cli/internal/models"
	"golang.org/x/sys/unix"
)

// GetMemoryInfo returns memory/RAM information
func GetMemoryInfo() (*models.MemoryInfo, error) {
	var info unix.Sysinfo_t
	if err := unix.Sysinfo(&info); err != nil {
		return nil, err
	}

	unitSize := uint64(info.Unit)
	if unitSize == 0 {
		unitSize = 1
	}

	totalBytes := info.Totalram * unitSize
	availBytes := info.Freeram * unitSize
	usedBytes := totalBytes - availBytes
	swapTotal := info.Totalswap * unitSize
	swapUsed := (info.Totalswap - info.Freeswap) * unitSize

	totalGB := bytesToGB(totalBytes)
	availGB := bytesToGB(availBytes)
	usedGB := bytesToGB(usedBytes)
	usagePercent := 0.0
	if totalGB > 0 {
		usagePercent = (usedGB / totalGB) * 100
	}

	swapTotalGB := bytesToGB(swapTotal)
	swapUsedGB := bytesToGB(swapUsed)

	return &models.MemoryInfo{
		TotalGB:      totalGB,
		AvailableGB:  availGB,
		UsedGB:       usedGB,
		UsagePercent: usagePercent,
		SwapTotalGB:  swapTotalGB,
		SwapUsedGB:   swapUsedGB,
	}, nil
}
