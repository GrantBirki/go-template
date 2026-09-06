# go-template

A super simple starter Go template for building CLI applications with Cobra. This project provides a solid foundation with modern Go best practices, including 100% dependency vendoring, hermetic builds, and build provenance attestations.

## 🚀 Features

- **Simple CLI**: Built with [Cobra](https://github.com/spf13/cobra)
- **Hermetic Builds**: All dependencies are vendored and builds work offline
- **Build Provenance**: Release artifacts include verifiable provenance attestations
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

# Exercise the compiled CLI as a consumer
script/acceptance

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

Runs the unit test suite once with the race detector and enforces 100% statement and function coverage for first-party packages under `internal/`.

```bash
script/test
```

### `script/acceptance`

Builds the real CLI from vendored modules and verifies its default, flag, help, and error behavior.

```bash
script/acceptance
```

### `script/lint`

Formats first-party Go code and fails in CI when formatting changes are required.

```bash
script/lint
```

### `script/build`

Builds the current-platform development binary directly with the pinned Go toolchain and vendored modules. Release builds use the checksum-verified GoReleaser package committed under `vendor_bin/`.

```bash
# Development build
script/build [go build flags]

# Linux amd64 release build (used by the release workflow)
script/build --release [goreleaser flags]
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
- Bootstrap, tests, acceptance tests, linting, and builds work completely offline after the pinned Go toolchain is installed
- Release builds verify and use the committed GoReleaser package instead of an ambient installation
- Release builds set `SOURCE_DATE_EPOCH` from the source commit

### Build Provenance

The release workflow creates and verifies build provenance attestations:

1. **Build**: Creates binaries with full provenance tracking
2. **Sign**: Generates cryptographic attestations using GitHub's OIDC
3. **Verify**: Validates all artifacts against their attestations

The workflow uses:

- Build provenance attestations (`actions/attest-build-provenance`)
- Signed artifacts with verifiable supply chain metadata
- Multi-step verification process

### Environment Configuration

`script/env` sets `GOTMPDIR` to the ignored `tmp/go` directory under the repository root and creates it before builds and tests. Use the repository scripts to inherit this setting; direct Go commands and editor test runners need the same environment. This keeps temporary Go build binaries with the checkout without changing the process-wide `TMPDIR`.

For a direct command from the repository root, invoke Bash explicitly so the example also works from an interactive zsh session:

```bash
bash <<'BASH'
source script/env
go test ./...
BASH
```

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

- **Test**: Runs unit tests and enforces coverage on every push and PR
- **Acceptance**: Exercises the compiled CLI as a consumer on every push and PR
- **Lint**: Code formatting and linting checks
- **Build**: Verifies the offline current-platform build
- **Release**: Triggered by git tags, creates releases with build provenance attestations
