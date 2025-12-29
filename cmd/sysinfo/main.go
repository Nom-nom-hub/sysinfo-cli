package main

import (
	"fmt"
	"os"
	"time"

	"github.com/example/sysinfo-cli/internal/output"
	"github.com/example/sysinfo-cli/internal/system"
)

func main() {
	config := parseFlags()

	if err := config.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Run once, or loop if watch mode
	ticker := time.NewTicker(time.Duration(config.WatchInterval) * time.Second)
	defer ticker.Stop()

	for {
		var (
			data interface{}
			err  error
		)

		switch config.Command {
		case "os":
			data, err = system.GetOSInfo()
		case "cpu":
			data, err = system.GetCPUInfo()
		case "memory":
			data, err = system.GetMemoryInfo()
		case "disk":
			data, err = system.GetDiskInfo(config.MountPoint)
		case "network":
			data, err = system.GetNetworkInfo()
		case "process":
			data, err = system.GetProcessInfo(config.SortBy, config.Limit)
		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", config.Command)
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		formatter := output.NewFormatter(config.Format, config.Pretty)
		formatted, err := formatter.Format(data, config.Command)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
			os.Exit(1)
		}

		writer := output.NewWriter(config.OutputFile)
		if err := writer.Write(formatted); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
			os.Exit(1)
		}

		if !config.Watch {
			break
		}

		<-ticker.C
	}
}
