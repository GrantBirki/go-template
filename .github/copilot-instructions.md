# GitHub Copilot Guidelines

This is a public Go template for building small command-line applications. Follow the repository guidance in `AGENTS.md`.

## Code Standards

### Development Flow

- Test: `script/test`
- Acceptance: `script/acceptance`
- Lint: `script/lint`
- Build: `script/build`

## Repository Structure

- `cmd/*`: Main cli entry points and executables
- `internal/`: Logic related to the core functionality of the CLI
- `script/`: Scripts for building, testing, and releasing the CLI
- `.github/`: GitHub Actions workflows for CI/CD
- `vendor/`: Vendor directory for Go modules (committed to the repository for reproducibility)

## Key Guidelines

1. Prefer idiomatic Go and the smallest complete change.
2. Prefer the standard library and do not add dependencies without a concrete need.
3. Use Go's built-in `testing` package and table-driven tests where they improve clarity.
4. Exercise production code rather than recreating application behavior in test helpers.
5. Maintain 100% statement and function coverage for first-party packages under `internal/`, including errors and edge cases.
6. Preserve vendored, offline behavior for normal bootstrap, test, lint, and build paths.
7. Keep responses concise and changes focused on the requested problem.
