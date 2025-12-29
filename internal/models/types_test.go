package models

import (
	"encoding/json"
	"testing"
)

func TestOSInfoJSON(t *testing.T) {
	info := &OSInfo{
		Hostname:      "testhost",
		OS:            "linux",
		Platform:      "linux",
		Release:       "go1.21.0",
		Architecture:  "x86_64",
		UptimeSeconds: 3600,
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded OSInfo
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.Hostname != info.Hostname {
		t.Errorf("Expected hostname %s, got %s", info.Hostname, decoded.Hostname)
	}
}

func TestCPUInfoJSON(t *testing.T) {
	info := &CPUInfo{
		Cores:       4,
		Threads:     8,
		Model:       "Intel Core i7",
		FrequencyGHz: 2.4,
		UsagePercent: 25.5,
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded CPUInfo
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.Cores != info.Cores {
		t.Errorf("Expected cores %d, got %d", info.Cores, decoded.Cores)
	}
}

func TestMemoryInfoJSON(t *testing.T) {
	info := &MemoryInfo{
		TotalGB:      16.0,
		AvailableGB:  8.5,
		UsedGB:       7.5,
		UsagePercent: 46.875,
		SwapTotalGB:  4.0,
		SwapUsedGB:   0.5,
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded MemoryInfo
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.TotalGB != info.TotalGB {
		t.Errorf("Expected total %.2f, got %.2f", info.TotalGB, decoded.TotalGB)
	}
}

func TestDiskInfoJSON(t *testing.T) {
	info := DiskInfo{
		Filesystem:   "/dev/sda1",
		MountPoint:   "/",
		SizeGB:       100.0,
		UsedGB:       50.0,
		AvailableGB:  50.0,
		UsagePercent: 50.0,
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded DiskInfo
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.MountPoint != info.MountPoint {
		t.Errorf("Expected mount %s, got %s", info.MountPoint, decoded.MountPoint)
	}
}

func TestNetworkInterfaceJSON(t *testing.T) {
	info := NetworkInterface{
		Name:        "eth0",
		IPAddresses: []string{"192.168.1.1/24"},
		MACAddress:  "00:11:22:33:44:55",
		MTU:         1500,
		Status:      "up",
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded NetworkInterface
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(decoded.IPAddresses) != 1 {
		t.Errorf("Expected 1 IP address, got %d", len(decoded.IPAddresses))
	}
}

func TestProcessInfoJSON(t *testing.T) {
	info := ProcessInfo{
		PID:        1234,
		Name:       "sysinfo",
		CPUPercent: 5.2,
		MemoryMB:   25.5,
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded ProcessInfo
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.PID != info.PID {
		t.Errorf("Expected PID %d, got %d", info.PID, decoded.PID)
	}
}
