//go:build windows
// +build windows

package system

import (
	"github.com/example/sysinfo-cli/internal/models"
)

func getProcessesLinux() []models.ProcessInfo {
	return []models.ProcessInfo{}
}

func getProcessesDarwin() []models.ProcessInfo {
	return []models.ProcessInfo{}
}

func getProcessesWindows() []models.ProcessInfo {
	// Windows process enumeration would require WMI or toolhelp32 API
	// For now, return empty to avoid errors on Windows
	// In production, would use: github.com/shirou/gopsutil or syscall toolhelp32
	return []models.ProcessInfo{}
}
