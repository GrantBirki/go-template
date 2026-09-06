# AGENTS.md

This repository is a public Go template for small command-line applications. Keep changes portable, minimal, and useful to downstream projects without assuming a contributor's local machine.

## Core Principles

- Use the repository-owned `script/*` entrypoints for routine work instead of ad hoc command sequences.
- Keep `.go-version` and the `go` directive in `go.mod` aligned to the exact supported Go version.
- Prefer the Go standard library. Add a dependency only when it provides enough value to justify its code, trust, and maintenance cost.
- Preserve offline behavior for bootstrap, test, acceptance, lint, and build paths. Application dependencies belong in `vendor/` and routine commands must not resolve from the network.
- Keep changes simple and focused. Fix the root cause, add regression coverage, and avoid unrelated refactoring or speculative abstractions.
- Treat this as a public repository. Never commit secrets, private infrastructure details, host-specific paths, personal configuration, or generated machine state.

## Scripts

`script/env` sets `GOTMPDIR` to the ignored `tmp/go` directory under the repository root and creates it before builds and tests. Use the repository scripts to inherit this setting; direct Go commands and editor test runners need the same environment. This keeps temporary Go build binaries with the checkout without changing the process-wide `TMPDIR`.

- `script/bootstrap` verifies that the vendored dependency graph resolves offline.
- `script/test` runs the Go unit tests and enforces 100% statement and function coverage for first-party packages under `internal/`.
- `script/acceptance` builds the real CLI offline and exercises its consumer-facing behavior.
- `script/lint` formats first-party Go code and fails in CI when formatting changes are required.
- `script/build` creates the current-platform binary with the Go toolchain and vendored modules.
- `script/build --release` uses the checksum-verified GoReleaser package committed under `vendor_bin/`.
- `script/update` is the intentional networked path for updating and re-vendoring Go modules.

New shell entrypoints should use `#!/usr/bin/env bash`, `set -euo pipefail`, and paths derived from `script/env`. Keep temporary files and extracted tools in ignored repository-local directories.

## Dependencies And Builds

- Keep `go.mod`, `go.sum`, and `vendor/` consistent in the same change.
- Do not edit vendored source by hand.
- Do not add Git dependencies unless they are justified and pinned to an immutable full commit revision.
- Keep GitHub Actions pinned to full commit SHAs and checkout credentials disabled unless a job explicitly needs them.
- Pin container images by immutable manifest digest if a workflow or build introduces one.
- Build and test through vendored modules with the proxy and checksum database disabled. Only update workflows may enable dependency network access.

## Testing

Use Go's built-in `testing` package and table-driven tests where they improve clarity. Tests must exercise production logic rather than duplicate it in test helpers. Cover successful behavior, errors, edge cases, and every meaningful branch. Bug fixes require regression tests.

Keep unit tests fast, deterministic, and independent of live services. Consumer-visible CLI changes also require acceptance coverage through `script/acceptance`.

## Documentation

Update `README.md`, scripts, workflows, and this file together when their contracts change. Keep Markdown prose unwrapped and keep examples generic enough for a public template.
