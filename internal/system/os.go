package system

import (
	"os"
	"runtime"

	"github.com/example/sysinfo-cli/internal/models"
)

// GetOSInfo returns operating system information
func GetOSInfo() (*models.OSInfo, error) {
	hostname, _ := os.Hostname()
	uptime := getSystemUptime()

	return &models.OSInfo{
		Hostname:      hostname,
		OS:            runtime.GOOS,
		Platform:      runtime.GOOS,
		Release:       runtime.Version(),
		Architecture:  runtime.GOARCH,
		UptimeSeconds: uptime,
	}, nil
}

func getSystemUptime() int64 {
	// Platform-specific implementation in os_*.go files
	return 0
}
