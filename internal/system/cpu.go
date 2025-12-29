package system

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/example/sysinfo-cli/internal/models"
)

// GetCPUInfo returns CPU information
func GetCPUInfo() (*models.CPUInfo, error) {
	cores := runtime.NumCPU()

	// For simplicity, threads = cores * 2 (typical for modern processors)
	// In production, would parse /proc/cpuinfo or use syscalls
	threads := cores * 2

	// Get model name from /proc/cpuinfo on Linux
	model := "Unknown"
	if runtime.GOOS == "linux" {
		if data, err := os.ReadFile("/proc/cpuinfo"); err == nil {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				if strings.Contains(line, "model name") {
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						model = strings.TrimSpace(parts[1])
						break
					}
				}
			}
		}
	} else if runtime.GOOS == "darwin" {
		// macOS: system_profiler would be needed for accurate model
		model = "Apple Silicon/Intel (use system_profiler for details)"
	} else if runtime.GOOS == "windows" {
		model = "Intel/AMD (use wmic for details)"
	}

	// For demo: static frequency (production would use actual CPU frequency)
	frequencyGHz := 2.4
	// For demo: calculate approximate usage (would use real CPU stats in production)
	usagePercent := 0.0

	return &models.CPUInfo{
		Cores:        cores,
		Threads:      threads,
		Model:        model,
		FrequencyGHz: frequencyGHz,
		UsagePercent: usagePercent,
	}, nil
}
