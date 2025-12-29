//go:build windows
// +build windows

package system

import (
	"syscall"
	"unsafe"

	"github.com/example/sysinfo-cli/internal/models"
)

// GetMemoryInfo returns memory/RAM information on Windows
func GetMemoryInfo() (*models.MemoryInfo, error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	globalMemoryStatusEx := kernel32.NewProc("GlobalMemoryStatusEx")

	type MEMORYSTATUSEX struct {
		Length               uint32
		MemoryLoad           uint32
		TotalPhys            uint64
		AvailPhys            uint64
		TotalPageFile        uint64
		AvailPageFile        uint64
		TotalVirtual         uint64
		AvailVirtual         uint64
		AvailExtendedVirtual uint64
	}

	memStatus := MEMORYSTATUSEX{
		Length: uint32(unsafe.Sizeof(MEMORYSTATUSEX{})),
	}

	ret, _, _ := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memStatus)))
	if ret == 0 {
		// Fallback values
		return &models.MemoryInfo{
			TotalGB:      0,
			AvailableGB:  0,
			UsedGB:       0,
			UsagePercent: 0,
			SwapTotalGB:  0,
			SwapUsedGB:   0,
		}, nil
	}

	totalGB := bytesToGB(memStatus.TotalPhys)
	availGB := bytesToGB(memStatus.AvailPhys)
	usedGB := totalGB - availGB
	usagePercent := 0.0
	if totalGB > 0 {
		usagePercent = (usedGB / totalGB) * 100
	}

	swapTotalGB := bytesToGB(memStatus.TotalPageFile)
	swapUsedGB := bytesToGB(memStatus.TotalPageFile - memStatus.AvailPageFile)

	return &models.MemoryInfo{
		TotalGB:      totalGB,
		AvailableGB:  availGB,
		UsedGB:       usedGB,
		UsagePercent: usagePercent,
		SwapTotalGB:  swapTotalGB,
		SwapUsedGB:   swapUsedGB,
	}, nil
}
