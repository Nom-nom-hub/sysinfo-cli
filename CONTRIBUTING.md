# Contributing to sysinfo-cli

Thank you for your interest in contributing! This document provides guidelines for getting started.

## Code of Conduct

Be respectful, inclusive, and constructive in all interactions.

## How to Contribute

### Reporting Bugs

1. **Check existing issues** to avoid duplicates
2. **Create issue** with:
   - Title: Clear, concise description
   - Environment: OS, architecture, Go version
   - Steps to reproduce: Exact commands that fail
   - Expected vs actual output
   - Logs/error messages

**Example:**
```
Title: `sysinfo memory` crashes on Windows 11 x86_64

Environment:
- OS: Windows 11 Pro
- Architecture: x86_64
- Go version: 1.21.3

Steps:
1. Build: go build -o bin\sysinfo.exe .\cmd\sysinfo
2. Run: .\bin\sysinfo.exe memory

Expected: Memory info table
Actual: Panic: runtime error...

Error log:
[paste error output]
```

### Suggesting Features

1. **Check roadmap** in [ROADMAP.md](ROADMAP.md)
2. **Create issue** with tag `enhancement`
3. **Describe:**
   - What feature/improvement
   - Why it's useful
   - How it should work (examples)

**Example:**
```
Title: Add GPU information command

Description:
It would be useful to have a `sysinfo gpu` command that shows:
- GPU name/model
- Memory (VRAM)
- Usage percentage
- Temperature (if available)

This would complement CPU/memory info for gaming/ML workloads.
```

### Submitting Code

#### Setup

```bash
# Fork on GitHub, then:
git clone https://github.com/YOUR-USERNAME/sysinfo-cli.git
cd sysinfo-cli

# Create feature branch
git checkout -b feat/your-feature-name

# Install dependencies
go mod download

# Set up pre-commit checks (optional)
make lint
```

#### Development Workflow

```bash
# Write code + tests
# Keep 80%+ coverage target

# Format code
make fmt

# Run linter
make lint

# Run tests
make test

# Check coverage
make test-coverage
```

#### Code Style

- Follow Go conventions (effective Go)
- Use clear variable names
- Add comments for exported functions
- Keep functions focused and small
- Use error wrapping: `fmt.Errorf("context: %w", err)`

**Example:**
```go
// GetCPUInfo returns CPU information from the system
func GetCPUInfo() (*models.CPUInfo, error) {
	cores := runtime.NumCPU()
	
	model, err := detectCPUModel()
	if err != nil {
		return nil, fmt.Errorf("detecting CPU model: %w", err)
	}
	
	// ... implementation
}
```

#### Testing Requirements

- Add unit tests for new functions
- Add integration tests for commands
- Maintain or improve coverage (80%+ target)
- Test on all supported platforms (Linux, macOS, Windows)

**Example test:**
```go
func TestGetCPUInfo(t *testing.T) {
	info, err := GetCPUInfo()
	
	if err != nil {
		t.Fatalf("GetCPUInfo failed: %v", err)
	}
	
	if info.Cores < 1 {
		t.Errorf("Expected cores > 0, got %d", info.Cores)
	}
}
```

#### Platform-Specific Code

Use build tags for OS-specific implementations:

```go
// +build linux darwin
package system

func getProcessesLinux() []models.ProcessInfo {
	// Linux implementation
}

// +build windows
package system

func getProcessesWindows() []models.ProcessInfo {
	// Windows implementation
}
```

### Submitting a Pull Request

1. **Commit with clear messages:**
   ```bash
   git commit -m "feat: add GPU information command"
   git commit -m "fix: resolve memory leak in disk enumeration"
   git commit -m "test: add coverage for CPU frequency detection"
   ```

2. **Push to your fork:**
   ```bash
   git push origin feat/your-feature-name
   ```

3. **Create Pull Request with:**
   - Clear title and description
   - Reference related issue (e.g., "Fixes #123")
   - Checklist:
     - [ ] Code follows style guidelines
     - [ ] Tests added/updated
     - [ ] Coverage maintained (80%+)
     - [ ] Tested on target platforms
     - [ ] No breaking changes to CLI
     - [ ] Documentation updated (if needed)

**Example PR description:**
```markdown
## Description

Adds GPU information command to display VRAM and GPU usage.

## Fixes

Closes #456 (GPU info request)

## Testing

- Tested on Windows 11 RTX 3080
- Tested on macOS M1 with Metal
- Tested on Ubuntu 22.04 with NVIDIA

## Checklist

- [x] Tests pass locally (coverage 82%)
- [x] Code formatted and linted
- [x] No hallucinated APIs
- [x] Works on all 3 OS
```

### Documentation

- Update README.md for user-facing changes
- Add comments to exported types/functions
- Update ROADMAP.md if affecting roadmap
- Add CHANGELOG entry for releases

## Project Structure

```
sysinfo-cli/
├── cmd/sysinfo/          # CLI entry point
├── internal/system/      # OS-specific system info
├── internal/output/      # Output formatting
├── internal/models/      # Data structures
├── tests/               # Integration tests
├── .github/workflows/   # CI/CD
├── README.md           # User guide
├── ROADMAP.md          # Feature roadmap
└── CONTRIBUTING.md     # This file
```

## Design Principles

When adding features, follow these principles:

1. **No Hallucinations:** Only real APIs, no fake logic
2. **Correctness over Speed:** Slow and right > fast and wrong
3. **Explicit over Implicit:** Clear code, minimal magic
4. **Production-Grade:** Proper error handling, tested
5. **Cross-Platform:** Works on Linux, macOS, Windows
6. **Zero Dependencies:** Only stdlib + golang.org/x/sys

## Getting Help

- **Questions:** Open discussion on GitHub Discussions
- **Issues:** Ask in GitHub Issues
- **Security:** Email security@example.com (don't use issues)

## Recognition

Contributors are listed in:
- GitHub repository contributors page
- CHANGELOG.md releases
- Release notes for contributed features

Thank you for making sysinfo-cli better! 🎉
