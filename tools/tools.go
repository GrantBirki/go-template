//go:build tools
// +build tools

package tools

import (
	// Tooling dependencies pinned in go.mod and vendored for offline builds.
	_ "github.com/goreleaser/goreleaser/v2"
)
