# go-template

A super simple starter Go template for building CLI applications with Cobra. This project provides a solid foundation with modern Go best practices, including 100% dependency vendoring, hermetic builds, and SLSA Level 3 security attestations.

## 🚀 Features

- **Simple CLI**: Built with [Cobra](https://github.com/spf13/cobra)
- **Hermetic Builds**: All dependencies are vendored and builds work offline
- **SLSA Level 3**: Supply chain security with build provenance attestations
- **Cross-Platform**: Builds for Linux, macOS, Windows, and FreeBSD
- **Modern Go**: Uses Go modules with 100% dependency vendoring
- **CI/CD Ready**: GitHub Actions workflows for testing, linting, and releasing

## 📦 Quick Start

```bash
# Clone the template
git clone https://github.com/grantbirki/go-template.git my-cli
cd my-cli

# Bootstrap the project (verify hermetic setup)
script/bootstrap

# Test the CLI
script/test

# Build the CLI
script/build

# Try it out
./go-template --help
./go-template --name "World"
```

## 🛠️ Development Scripts

This project uses script-based workflows for consistency and ease of use:

### `script/bootstrap`

Verifies the hermetic build setup by ensuring all dependencies resolve from the vendored modules only. This guarantees builds work without network access.

```bash
script/bootstrap
```

### `script/test`

Runs the complete test suite with coverage reporting.

```bash
script/test
```

### `script/lint`

Formats code and runs `golangci-lint` with auto-fixing enabled.

```bash
script/lint
```

### `script/build`

Builds binaries using GoReleaser with two modes:

```bash
# Development build (snapshot mode)
script/build [flags]

# Release build (for tagged releases and publishing to GitHub)
script/build --release [flags]
```

### `script/update`

Updates Go dependencies while maintaining the vendored setup:

```bash
# Update to latest compatible versions (minor/patch only)
script/update

# Update ALL dependencies including major versions
script/update --all
```

This script temporarily enables network access, updates dependencies, runs `go mod tidy`, re-vendors everything, and verifies the modules.

### `script/release`

Interactive script to create and tag new releases:

```bash
script/release
```

Prompts for a new version tag in `vX.X.X` format and creates the git tag that triggers the release workflow.

## 🔒 Security & Supply Chain

### Hermetic Builds

This project uses **100% dependency vendoring** for hermetic builds:

- All dependencies are committed in the `vendor/` directory
- `GOPROXY=off` and `GOSUMDB=off` ensure no network access during builds
- Builds are reproducible and work completely offline
- `SOURCE_DATE_EPOCH` is set for deterministic builds

### SLSA Level 3 Compliance

The release workflow implements SLSA Level 3 security:

1. **Build**: Creates binaries with full provenance tracking
2. **Sign**: Generates cryptographic attestations using GitHub's OIDC
3. **Verify**: Validates all artifacts against their attestations

The workflow uses:

- Build provenance attestations (`actions/attest-build-provenance`)
- Signed artifacts with verifiable supply chain metadata
- Multi-step verification process

### Environment Configuration

The `script/env` script sets up the hermetic environment:

```bash
export GOPROXY="off"      # No proxy access
export GOSUMDB="off"      # No checksum database
export GOFLAGS="-mod=vendor"  # Force vendor mode
```

## 🏗️ Project Structure

```text
├── cmd/go-template/     # Main CLI entry point
├── internal/cmd/        # Cobra command implementations
├── internal/version/    # Version information with build-time injection
├── script/              # Development and build scripts
├── vendor/              # All vendored dependencies (committed)
├── .github/workflows/   # CI/CD pipelines
└── .goreleaser.yaml     # GoReleaser configuration
```

## 🔧 Customizing the Template

1. **Update module name**: Change `go.mod` and import paths
2. **Rename binary**: Update `.goreleaser.yaml` and `cmd/` directory
3. **Add commands**: Create new files in `internal/cmd/`
4. **Update repository**: Change GitHub repository references in workflows

## 📋 CI/CD Workflows

- **Test**: Runs tests on every push and PR
- **Lint**: Code formatting and linting checks
- **Build**: Verifies builds work on multiple platforms
- **Release**: Triggered by git tags, creates signed releases with SLSA attestations
