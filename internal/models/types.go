package models

// OSInfo represents operating system information
type OSInfo struct {
	Hostname       string `json:"hostname"`
	OS             string `json:"os"`
	Platform       string `json:"platform"`
	Release        string `json:"release"`
	Architecture   string `json:"arch"`
	UptimeSeconds  int64  `json:"uptime_seconds"`
}

// CPUInfo represents CPU information
type CPUInfo struct {
	Cores        int     `json:"cores"`
	Threads       int     `json:"threads"`
	Model         string  `json:"model"`
	FrequencyGHz  float64 `json:"frequency_ghz"`
	UsagePercent  float64 `json:"usage_percent"`
}

// MemoryInfo represents memory/RAM information
type MemoryInfo struct {
	TotalGB       float64 `json:"total_gb"`
	AvailableGB   float64 `json:"available_gb"`
	UsedGB        float64 `json:"used_gb"`
	UsagePercent  float64 `json:"usage_percent"`
	SwapTotalGB   float64 `json:"swap_total_gb"`
	SwapUsedGB    float64 `json:"swap_used_gb"`
}

// DiskInfo represents a single disk/partition
type DiskInfo struct {
	Filesystem    string  `json:"filesystem"`
	MountPoint    string  `json:"mount_point"`
	SizeGB        float64 `json:"size_gb"`
	UsedGB        float64 `json:"used_gb"`
	AvailableGB   float64 `json:"available_gb"`
	UsagePercent  float64 `json:"usage_percent"`
}

// NetworkInterface represents network interface information
type NetworkInterface struct {
	Name         string   `json:"name"`
	IPAddresses  []string `json:"ip_addresses"`
	MACAddress   string   `json:"mac_address"`
	MTU          int      `json:"mtu"`
	Status       string   `json:"status"`
}

// ProcessInfo represents a single process
type ProcessInfo struct {
	PID         int     `json:"pid"`
	Name        string  `json:"name"`
	CPUPercent  float64 `json:"cpu_percent"`
	MemoryMB    float64 `json:"memory_mb"`
}

// OutputFormats defines supported output types
type OutputFormat string

const (
	FormatJSON  OutputFormat = "json"
	FormatTable OutputFormat = "table"
	FormatCSV   OutputFormat = "csv"
)
