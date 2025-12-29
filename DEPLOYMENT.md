# Deployment Guide

## Phase 7: Deploy

### Prerequisites

- Git installed and configured
- GitHub repository created
- Go 1.21+

### Local Repository Setup

```bash
cd sysinfo-cli

# Initialize git
git init

# Configure git (if not already configured)
git config user.email "you@example.com"
git config user.name "Your Name"

# Add all files
git add .

# Initial commit
git commit -m "Initial commit: sysinfo-cli v1.0.0"
```

### Push to GitHub

```bash
# Add remote (replace with your repo URL)
git remote add origin https://github.com/yourusername/sysinfo-cli.git

# Rename branch to main if needed
git branch -M main

# Push to GitHub
git push -u origin main
```

### GitHub Actions Setup

GitHub Actions workflows are already configured:

**Files:**
- `.github/workflows/test.yml` — Runs on every push/PR
- `.github/workflows/release.yml` — Runs on git tags (v*)

**Automatic CI/CD:**
- Tests run on Ubuntu, macOS, Windows with Go 1.21, 1.22
- Coverage reports uploaded to Codecov
- Binaries auto-built on release tags

### Creating a Release

```bash
# Tag a release
git tag -a v1.0.0 -m "Release v1.0.0: Initial production release"

# Push the tag (triggers GitHub Actions release workflow)
git push origin v1.0.0
```

**GitHub Actions will:**
1. Build binaries for all 5 platforms
2. Calculate SHA256 checksums
3. Create GitHub Release with binaries attached
4. Make binaries available for download

### Release Checklist

Before tagging v1.0.0:

- [ ] All tests passing locally: `go test -v -cover ./...`
- [ ] Coverage >= 80%: `go test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`
- [ ] Code linted: `golangci-lint run ./...`
- [ ] README updated with latest features
- [ ] All commands tested on target platforms
- [ ] No hallucinated APIs or fake dependencies
- [ ] Security review completed

### Installation for End Users

After release, users download binaries from GitHub Releases:

```bash
# Linux
wget https://github.com/yourusername/sysinfo-cli/releases/download/v1.0.0/sysinfo-linux-x86_64
chmod +x sysinfo-linux-x86_64
./sysinfo-linux-x86_64 os

# macOS
wget https://github.com/yourusername/sysinfo-cli/releases/download/v1.0.0/sysinfo-darwin-arm64
chmod +x sysinfo-darwin-arm64
./sysinfo-darwin-arm64 os

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/yourusername/sysinfo-cli/releases/download/v1.0.0/sysinfo-windows-x86_64.exe" -OutFile "sysinfo.exe"
.\sysinfo.exe os
```

### Verify Release Integrity

```bash
# Check SHA256 (Linux/macOS)
sha256sum -c sysinfo-linux-x86_64.sha256

# Check SHA256 (Windows PowerShell)
certutil -hashfile sysinfo-windows-x86_64.exe SHA256
# Compare against .sha256 file from release
```

### Future Distribution

**Optional enhancements:**
- Homebrew formula: `brew install sysinfo-cli`
- Chocolatey package: `choco install sysinfo-cli`
- APT/Yum packages for Linux distributions
- GitHub Releases page branding and documentation

### Troubleshooting

**GitHub Actions not triggering:**
- Verify workflows are in `.github/workflows/`
- Check branch is `main`
- Ensure commits are pushed to GitHub

**Build fails on specific platform:**
- Check platform-specific build files (`*_unix.go`, `*_windows.go`)
- Verify dependencies available on that platform
- Review GitHub Actions logs

**Binary doesn't run:**
- Confirm SHA256 checksum matches release
- Ensure correct architecture/OS downloaded
- Check execute permissions: `chmod +x sysinfo-*` (Unix)

## Summary

1. **Git setup:** Initialize repo, configure user, add origin
2. **Push code:** `git push -u origin main`
3. **Create release:** Tag with `git tag -a v1.0.0` and push
4. **Automated build:** GitHub Actions builds all binaries
5. **Release published:** Users download from GitHub Releases

**No manual binary building required** — fully automated via GitHub Actions.
