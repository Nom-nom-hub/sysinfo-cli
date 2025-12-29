//go:build windows
// +build windows

package system

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/example/sysinfo-cli/internal/models"
)

func getDiskInfoPlatform() []models.DiskInfo {
	return getDiskInfoWindows()
}

func getDiskInfoWindows() []models.DiskInfo {
	var disks []models.DiskInfo

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getDiskFreeSpaceEx := kernel32.NewProc("GetDiskFreeSpaceExW")

	for drive := 'C'; drive <= 'Z'; drive++ {
		drivePath := fmt.Sprintf("%c:\\", drive)
		driveName := fmt.Sprintf("%c:", drive)

		var lpFreeBytesAvailable, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes int64

		ret, _, _ := getDiskFreeSpaceEx.Call(
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(drivePath))),
			uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)),
		)

		if ret == 0 {
			continue
		}

		totalBytes := uint64(lpTotalNumberOfBytes)
		availBytes := uint64(lpTotalNumberOfFreeBytes)
		usedBytes := totalBytes - availBytes
		usagePercent := 0.0
		if totalBytes > 0 {
			usagePercent = (float64(usedBytes) / float64(totalBytes)) * 100
		}

		disks = append(disks, models.DiskInfo{
			Filesystem:   driveName,
			MountPoint:   drivePath,
			SizeGB:       bytesToGB(totalBytes),
			UsedGB:       bytesToGB(usedBytes),
			AvailableGB:  bytesToGB(availBytes),
			UsagePercent: usagePercent,
		})
	}

	return disks
}
