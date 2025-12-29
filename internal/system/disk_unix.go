//go:build linux || darwin
// +build linux darwin

package system

import (
	"bufio"
	"os"
	"runtime"
	"strings"

	"github.com/example/sysinfo-cli/internal/models"
	"golang.org/x/sys/unix"
)

func getDiskInfoPlatform() []models.DiskInfo {
	if runtime.GOOS == "linux" {
		return getDiskInfoLinux()
	}
	return getDiskInfoDarwin()
}

func getDiskInfoLinux() []models.DiskInfo {
	var disks []models.DiskInfo

	file, err := os.Open("/etc/mtab")
	if err != nil {
		return disks
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		filesystem := parts[0]
		mountPoint := parts[1]

		// Skip non-device mounts
		if !strings.HasPrefix(filesystem, "/dev") {
			continue
		}

		var stat unix.Statfs_t
		if err := unix.Statfs(mountPoint, &stat); err != nil {
			continue
		}

		blockSize := uint64(stat.Bsize)
		totalBytes := stat.Blocks * blockSize
		availBytes := stat.Bavail * blockSize
		usedBytes := totalBytes - availBytes
		usagePercent := 0.0
		if totalBytes > 0 {
			usagePercent = (float64(usedBytes) / float64(totalBytes)) * 100
		}

		disks = append(disks, models.DiskInfo{
			Filesystem:   filesystem,
			MountPoint:   mountPoint,
			SizeGB:       bytesToGB(totalBytes),
			UsedGB:       bytesToGB(usedBytes),
			AvailableGB:  bytesToGB(availBytes),
			UsagePercent: usagePercent,
		})
	}

	return disks
}

func getDiskInfoDarwin() []models.DiskInfo {
	var disks []models.DiskInfo

	// On macOS, check common mount points
	for _, mount := range []string{"/", "/Volumes"} {
		var stat unix.Statfs_t
		if err := unix.Statfs(mount, &stat); err != nil {
			continue
		}

		blockSize := uint64(stat.Bsize)
		totalBytes := stat.Blocks * blockSize
		availBytes := stat.Bavail * blockSize
		usedBytes := totalBytes - availBytes
		usagePercent := 0.0
		if totalBytes > 0 {
			usagePercent = (float64(usedBytes) / float64(totalBytes)) * 100
		}

		disks = append(disks, models.DiskInfo{
			Filesystem:   "disk",
			MountPoint:   mount,
			SizeGB:       bytesToGB(totalBytes),
			UsedGB:       bytesToGB(usedBytes),
			AvailableGB:  bytesToGB(availBytes),
			UsagePercent: usagePercent,
		})
	}

	return disks
}
