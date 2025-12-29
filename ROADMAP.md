# Roadmap

## v1.0.0 (Current) ✓

**Released Features:**
- [x] OS information (hostname, OS, release, architecture, uptime)
- [x] CPU information (cores, threads, model)
- [x] Memory information (RAM, swap, usage %)
- [x] Disk information (mount points, usage, capacity)
- [x] Network interfaces (IP addresses, MAC, status)
- [x] Cross-platform support (Linux, macOS, Windows)
- [x] Multiple output formats (JSON, table, CSV)
- [x] CI/CD automation (GitHub Actions)
- [x] Comprehensive testing (unit + integration)
- [x] 80%+ code coverage
- [x] Production-grade error handling

**Known Limitations:**
- Process enumeration not available on Windows/macOS
- CPU frequency simplified on macOS/Windows
- No real-time CPU usage tracking

---

## v1.1.0 (Planned)

**Target: Q1 2025**

### Windows Process Enumeration
- **Task:** Implement Windows process enumeration using WMI
- **Scope:** Top N processes by CPU/memory on Windows
- **Test:** Verify on Windows 10+ systems
- **Effort:** Medium (WMI syscall integration)

### macOS Process Enumeration
- **Task:** Implement macOS process enumeration using BSD syscalls
- **Scope:** Top N processes by CPU/memory on macOS
- **Test:** Verify on macOS 12+ systems
- **Effort:** Medium (BSD syscall integration)

### Enhanced CPU Frequency Detection
- **Task:** Real CPU frequency on all platforms
- **Linux:** Parse /proc/cpuinfo more accurately
- **macOS:** Use system_profiler or IOKit
- **Windows:** Use WMI Win32_Processor
- **Effort:** Low

### Watch Mode Refinement
- **Task:** Real-time monitoring with interval control
- **Flags:** `--watch --interval 1` (seconds)
- **Scope:** All commands support continuous updates
- **Effort:** Medium

### Color Output
- **Task:** Optional colored table output
- **Flags:** `--color` (auto/on/off)
- **Scope:** Tables only (JSON/CSV unaffected)
- **Effort:** Low

---

## v1.2.0 (Planned)

**Target: Q2 2025**

### Package Manager Distribution
- Homebrew formula for macOS/Linux
- Chocolatey package for Windows
- APT/Yum packages for Linux distributions
- Effort: Low (build system automation)

### Shell Completions
- Bash completion script
- Zsh completion script
- Fish completion script
- PowerShell completion (Windows)
- Effort: Low

### Human-Friendly Output Enhancements
- Uptime formatting: "2 days, 3 hours, 45 minutes" instead of seconds
- Memory sizes: Human-readable (GiB, TiB) with configurable precision
- Percentage formatting: Aligned columns
- Effort: Low

### Additional System Info Commands
- `sysinfo env` — Environment variables and system paths
- `sysinfo gpu` — GPU information (if available)
- `sysinfo battery` — Battery status (laptops)
- Effort: Medium per command

---

## v2.0.0 (Future)

**Target: 2025 H2+**

### Advanced Monitoring
- Export to Prometheus format
- InfluxDB integration
- Statsd metrics
- Effort: High

### System Performance Analysis
- Benchmark mode: Compare against baseline
- Anomaly detection: Alert on unusual metrics
- Trend analysis: Historical data collection
- Effort: High

### Interactive TUI Mode
- Terminal UI with real-time updates
- Keyboard navigation
- Resource filtering/sorting
- Effort: Very High

---

## Backlog (Future Consideration)

- [ ] Configuration file support (e.g., ~/.sysinfo/config.yaml)
- [ ] Plugin system for custom commands
- [ ] Remote system monitoring (SSH)
- [ ] Alert integration (Slack, email, PagerDuty)
- [ ] Web API server mode
- [ ] Grafana data source plugin
- [ ] Mobile app (data sync from CLI)

---

## Contributing to Roadmap Items

Want to help? Pick any roadmap item and open an issue/PR:

1. **Fork** the repository
2. **Create feature branch:** `git checkout -b feat/windows-process-enum`
3. **Implement** with tests (maintain 80%+ coverage)
4. **Test** on target platforms
5. **Submit PR** with description of changes

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

---

## Feedback

Help shape the roadmap:
- 💬 **Suggest features:** Open an issue with tag `enhancement`
- 🐛 **Report bugs:** Open an issue with tag `bug`
- 👍 **Vote on features:** React to issues with thumbs up
- 📝 **Contribute:** Submit PRs for any roadmap item

**Most requested features will be prioritized.**
