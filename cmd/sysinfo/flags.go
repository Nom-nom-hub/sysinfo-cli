package main

import (
	"flag"
	"fmt"
	"os"
)

// Config holds CLI configuration from flags
type Config struct {
	Command     string
	Format      string
	OutputFile  string
	Pretty      bool
	Watch       bool
	SortBy      string
	Limit       int
	MountPoint  string
	Color       string
}

func parseFlags() Config {
	fs := flag.NewFlagSet("sysinfo", flag.ContinueOnError)

	format := fs.String("format", "table", "Output format: json, table, or csv")
	output := fs.String("output", "", "Output file (default: stdout)")
	pretty := fs.Bool("pretty", false, "Pretty-print JSON")
	watch := fs.Bool("watch", false, "Watch mode (continuous updates)")
	sortBy := fs.String("sort", "cpu", "Sort processes by: cpu or memory")
	limit := fs.Int("limit", 10, "Number of top processes to display")
	mount := fs.String("mount", "", "Filter disk by mount point")
	color := fs.String("color", "auto", "Color output: auto, on, or off")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: sysinfo <command> [flags]

Commands:
  os        Display operating system information
  cpu       Display CPU information and usage
  memory    Display memory/RAM information
  disk      Display disk/storage information
  network   Display network interface information
  process   Display top processes by CPU/memory

Flags:
`)
		fs.PrintDefaults()
	}

	if len(os.Args) < 2 {
		fs.Usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	if err := fs.Parse(os.Args[2:]); err != nil {
		fs.Usage()
		os.Exit(1)
	}

	return Config{
		Command:    cmd,
		Format:     *format,
		OutputFile: *output,
		Pretty:     *pretty,
		Watch:      *watch,
		SortBy:     *sortBy,
		Limit:      *limit,
		MountPoint: *mount,
		Color:      *color,
	}
}

func (c Config) Validate() error {
	validCommands := map[string]bool{
		"os": true, "cpu": true, "memory": true,
		"disk": true, "network": true, "process": true,
	}

	if !validCommands[c.Command] {
		return fmt.Errorf("unknown command: %s", c.Command)
	}

	validFormats := map[string]bool{
		"json": true, "table": true, "csv": true,
	}

	if !validFormats[c.Format] {
		return fmt.Errorf("invalid format: %s", c.Format)
	}

	if c.SortBy != "cpu" && c.SortBy != "memory" {
		return fmt.Errorf("invalid sort: %s (must be cpu or memory)", c.SortBy)
	}

	if c.Limit < 1 {
		return fmt.Errorf("limit must be >= 1")
	}

	validColors := map[string]bool{
		"auto": true, "on": true, "off": true,
	}

	if !validColors[c.Color] {
		return fmt.Errorf("invalid color: %s (must be auto, on, or off)", c.Color)
	}

	return nil
}
