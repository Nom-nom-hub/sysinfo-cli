package output

import (
	"encoding/json"
	"fmt"

	"github.com/example/sysinfo-cli/internal/models"
)

// Formatter handles output formatting
type Formatter struct {
	format string
	pretty bool
}

// NewFormatter creates a new formatter
func NewFormatter(format string, pretty bool) *Formatter {
	return &Formatter{
		format: format,
		pretty: pretty,
	}
}

// Format converts data to the specified format
func (f *Formatter) Format(data interface{}, command string) (string, error) {
	switch f.format {
	case "json":
		return f.formatJSON(data)
	case "table":
		return f.formatTable(data, command)
	case "csv":
		return f.formatCSV(data, command)
	default:
		return "", fmt.Errorf("unknown format: %s", f.format)
	}
}

func (f *Formatter) formatJSON(data interface{}) (string, error) {
	var (
		bytes []byte
		err   error
	)

	if f.pretty {
		bytes, err = json.MarshalIndent(data, "", "  ")
	} else {
		bytes, err = json.Marshal(data)
	}

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (f *Formatter) formatTable(data interface{}, command string) (string, error) {
	var result string

	switch command {
	case "os":
		result = formatOSTable(data.(*models.OSInfo))
	case "cpu":
		result = formatCPUTable(data.(*models.CPUInfo))
	case "memory":
		result = formatMemoryTable(data.(*models.MemoryInfo))
	case "disk":
		result = formatDiskTable(data.([]models.DiskInfo))
	case "network":
		result = formatNetworkTable(data.([]models.NetworkInterface))
	case "process":
		result = formatProcessTable(data.([]models.ProcessInfo))
	default:
		return "", fmt.Errorf("unknown command: %s", command)
	}

	return result, nil
}

func (f *Formatter) formatCSV(data interface{}, command string) (string, error) {
	var result string

	switch command {
	case "os":
		result = formatOSCSV(data.(*models.OSInfo))
	case "cpu":
		result = formatCPUCSV(data.(*models.CPUInfo))
	case "memory":
		result = formatMemoryCSV(data.(*models.MemoryInfo))
	case "disk":
		result = formatDiskCSV(data.([]models.DiskInfo))
	case "network":
		result = formatNetworkCSV(data.([]models.NetworkInterface))
	case "process":
		result = formatProcessCSV(data.([]models.ProcessInfo))
	default:
		return "", fmt.Errorf("unknown command: %s", command)
	}

	return result, nil
}

// Table formatters
func formatOSTable(info *models.OSInfo) string {
	return fmt.Sprintf(`OS Information:
  Hostname:      %s
  OS:            %s
  Platform:      %s
  Release:       %s
  Architecture:  %s
  Uptime:        %d seconds
`,
		info.Hostname, info.OS, info.Platform, info.Release, info.Architecture, info.UptimeSeconds)
}

func formatCPUTable(info *models.CPUInfo) string {
	return fmt.Sprintf(`CPU Information:
  Cores:         %d
  Threads:       %d
  Model:         %s
  Frequency:     %.2f GHz
  Usage:         %.2f%%
`,
		info.Cores, info.Threads, info.Model, info.FrequencyGHz, info.UsagePercent)
}

func formatMemoryTable(info *models.MemoryInfo) string {
	return fmt.Sprintf(`Memory Information:
  Total:         %.2f GB
  Available:     %.2f GB
  Used:          %.2f GB
  Usage:         %.2f%%
  Swap Total:    %.2f GB
  Swap Used:     %.2f GB
`,
		info.TotalGB, info.AvailableGB, info.UsedGB, info.UsagePercent, info.SwapTotalGB, info.SwapUsedGB)
}

func formatDiskTable(disks []models.DiskInfo) string {
	result := "Disk Information:\n"
	result += "  Filesystem          Mount Point     Size      Used      Available  Usage%\n"
	result += "  ------------------  -----------  ----------  ---------  ----------  -------\n"

	for _, d := range disks {
		result += fmt.Sprintf("  %-18s  %-11s  %9.2f  %8.2f  %9.2f  %6.2f%%\n",
			d.Filesystem, d.MountPoint, d.SizeGB, d.UsedGB, d.AvailableGB, d.UsagePercent)
	}

	return result
}

func formatNetworkTable(ifaces []models.NetworkInterface) string {
	result := "Network Interfaces:\n"
	result += "  Name            Status  IP Addresses\n"
	result += "  -----------  ---------  ------------------------------------------\n"

	for _, i := range ifaces {
		ips := ""
		for _, ip := range i.IPAddresses {
			if ips != "" {
				ips += ", "
			}
			ips += ip
		}
		result += fmt.Sprintf("  %-11s  %-7s  %s\n", i.Name, i.Status, ips)
	}

	return result
}

func formatProcessTable(processes []models.ProcessInfo) string {
	result := "Top Processes:\n"
	result += "  PID      Name            CPU%    Memory(MB)\n"
	result += "  ------  ---------------  -------  ----------\n"

	for _, p := range processes {
		result += fmt.Sprintf("  %-6d  %-15s  %6.2f  %9.2f\n", p.PID, p.Name, p.CPUPercent, p.MemoryMB)
	}

	return result
}

// CSV formatters
func formatOSCSV(info *models.OSInfo) string {
	return fmt.Sprintf("hostname,os,platform,release,architecture,uptime_seconds\n%s,%s,%s,%s,%s,%d\n",
		info.Hostname, info.OS, info.Platform, info.Release, info.Architecture, info.UptimeSeconds)
}

func formatCPUCSV(info *models.CPUInfo) string {
	return fmt.Sprintf("cores,threads,model,frequency_ghz,usage_percent\n%d,%d,%s,%.2f,%.2f\n",
		info.Cores, info.Threads, info.Model, info.FrequencyGHz, info.UsagePercent)
}

func formatMemoryCSV(info *models.MemoryInfo) string {
	return fmt.Sprintf("total_gb,available_gb,used_gb,usage_percent,swap_total_gb,swap_used_gb\n%.2f,%.2f,%.2f,%.2f,%.2f,%.2f\n",
		info.TotalGB, info.AvailableGB, info.UsedGB, info.UsagePercent, info.SwapTotalGB, info.SwapUsedGB)
}

func formatDiskCSV(disks []models.DiskInfo) string {
	result := "filesystem,mount_point,size_gb,used_gb,available_gb,usage_percent\n"
	for _, d := range disks {
		result += fmt.Sprintf("%s,%s,%.2f,%.2f,%.2f,%.2f\n",
			d.Filesystem, d.MountPoint, d.SizeGB, d.UsedGB, d.AvailableGB, d.UsagePercent)
	}
	return result
}

func formatNetworkCSV(ifaces []models.NetworkInterface) string {
	result := "name,ip_addresses,mac_address,mtu,status\n"
	for _, i := range ifaces {
		ips := ""
		for _, ip := range i.IPAddresses {
			if ips != "" {
				ips += ";"
			}
			ips += ip
		}
		result += fmt.Sprintf("%s,%s,%s,%d,%s\n", i.Name, ips, i.MACAddress, i.MTU, i.Status)
	}
	return result
}

func formatProcessCSV(processes []models.ProcessInfo) string {
	result := "pid,name,cpu_percent,memory_mb\n"
	for _, p := range processes {
		result += fmt.Sprintf("%d,%s,%.2f,%.2f\n", p.PID, p.Name, p.CPUPercent, p.MemoryMB)
	}
	return result
}
