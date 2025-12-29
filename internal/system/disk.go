package system

import (
	"strings"

	"github.com/example/sysinfo-cli/internal/models"
)

// GetDiskInfo returns disk/partition information
// Platform-specific implementations in disk_unix.go and disk_windows.go
func GetDiskInfo(mountPointFilter string) ([]models.DiskInfo, error) {
	disks := getDiskInfoPlatform()

	if mountPointFilter != "" {
		var filtered []models.DiskInfo
		for _, disk := range disks {
			if strings.Contains(disk.MountPoint, mountPointFilter) {
				filtered = append(filtered, disk)
			}
		}
		return filtered, nil
	}

	return disks, nil
}
