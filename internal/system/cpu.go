package system

import (
	"os"
	"os/exec"
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

	// Get model name and frequency from platform-specific sources
	model := "Unknown"
	frequencyGHz := 0.0

	if runtime.GOOS == "linux" {
		model, frequencyGHz = getCPUInfoLinux()
	} else if runtime.GOOS == "darwin" {
		model, frequencyGHz = getCPUInfoDarwin()
	} else if runtime.GOOS == "windows" {
		model, frequencyGHz = getCPUInfoWindows()
	}

	// Frequency defaults if not detected
	if frequencyGHz == 0.0 {
		frequencyGHz = 2.4
	}

	// CPU usage percent (requires more complex monitoring, set to 0 for now)
	usagePercent := 0.0

	return &models.CPUInfo{
		Cores:        cores,
		Threads:      threads,
		Model:        model,
		FrequencyGHz: frequencyGHz,
		UsagePercent: usagePercent,
	}, nil
}

// getCPUInfoLinux extracts CPU model and frequency from /proc/cpuinfo
func getCPUInfoLinux() (string, float64) {
	model := "Unknown"
	frequency := 0.0

	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return model, frequency
	}

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

	// Try to get frequency from cpu_mhz
	for _, line := range lines {
		if strings.Contains(line, "cpu MHz") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				freqStr := strings.TrimSpace(parts[1])
				if mhz, err := strconv.ParseFloat(freqStr, 64); err == nil {
					frequency = mhz / 1000.0 // Convert MHz to GHz
					break
				}
			}
		}
	}

	return model, frequency
}

// getCPUInfoDarwin extracts CPU model and frequency on macOS
func getCPUInfoDarwin() (string, float64) {
	model := "Unknown"
	frequency := 0.0

	// Try system_profiler for detailed info
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Processor Name") {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					model = strings.TrimSpace(parts[1])
				}
			} else if strings.Contains(line, "Processor Speed") {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					speedStr := strings.TrimSpace(parts[1])
					// Parse "2.4 GHz" format
					speedStr = strings.TrimSuffix(speedStr, " GHz")
					if ghz, err := strconv.ParseFloat(speedStr, 64); err == nil {
						frequency = ghz
					}
				}
			}
		}
	}

	return model, frequency
}

// getCPUInfoWindows extracts CPU model and frequency on Windows
func getCPUInfoWindows() (string, float64) {
	model := "Unknown"
	frequency := 0.0

	// Use wmic to get CPU info
	cmd := exec.Command("wmic", "cpu", "get", "Name,MaxClockSpeed", "/value")
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Name=") {
				parts := strings.Split(line, "=")
				if len(parts) > 1 {
					model = strings.TrimSpace(parts[1])
				}
			} else if strings.Contains(line, "MaxClockSpeed=") {
				parts := strings.Split(line, "=")
				if len(parts) > 1 {
					speedStr := strings.TrimSpace(parts[1])
					// MaxClockSpeed is in MHz
					if mhz, err := strconv.ParseFloat(speedStr, 64); err == nil {
						frequency = mhz / 1000.0 // Convert MHz to GHz
					}
				}
			}
		}
	}

	return model, frequency
}
