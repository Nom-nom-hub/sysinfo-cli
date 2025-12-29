package main

import (
	"testing"
)

func TestValidateValidCommand(t *testing.T) {
	config := Config{
		Command: "os",
		Format:  "json",
		SortBy:  "cpu",
		Limit:   10,
		Color:   "auto",
	}

	if err := config.Validate(); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestValidateInvalidCommand(t *testing.T) {
	config := Config{
		Command: "invalid",
		Format:  "json",
	}

	if err := config.Validate(); err == nil {
		t.Errorf("Expected error for invalid command")
	}
}

func TestValidateInvalidFormat(t *testing.T) {
	config := Config{
		Command: "os",
		Format:  "xml",
	}

	if err := config.Validate(); err == nil {
		t.Errorf("Expected error for invalid format")
	}
}

func TestValidateInvalidSort(t *testing.T) {
	config := Config{
		Command: "process",
		Format:  "json",
		SortBy:  "invalid",
	}

	if err := config.Validate(); err == nil {
		t.Errorf("Expected error for invalid sort")
	}
}

func TestValidateLimitTooLow(t *testing.T) {
	config := Config{
		Command: "process",
		Format:  "json",
		SortBy:  "cpu",
		Limit:   0,
	}

	if err := config.Validate(); err == nil {
		t.Errorf("Expected error for limit < 1")
	}
}

func TestValidAllCommands(t *testing.T) {
	commands := []string{"os", "cpu", "memory", "disk", "network", "process"}

	for _, cmd := range commands {
		config := Config{
			Command: cmd,
			Format:  "json",
			SortBy:  "cpu",
			Limit:   10,
			Color:   "auto",
		}

		if err := config.Validate(); err != nil {
			t.Errorf("Command %s should be valid: %v", cmd, err)
		}
	}
}

func TestValidAllFormats(t *testing.T) {
	formats := []string{"json", "table", "csv"}

	for _, fmt := range formats {
		config := Config{
			Command: "os",
			Format:  fmt,
			SortBy:  "cpu",
			Limit:   10,
			Color:   "auto",
		}

		if err := config.Validate(); err != nil {
			t.Errorf("Format %s should be valid: %v", fmt, err)
		}
	}
}
