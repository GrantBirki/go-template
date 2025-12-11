# Vendored tools

This directory is where binaries built from vendored Go sources live (no prebuilt archives). `script/bootstrap` builds `goreleaser` from the vendored module pinned in `tools/tools.go` and `go.mod`, keeping CI offline and reproducible.

To bump goreleaser:

1. Update `.goreleaser-version` and the version in `go.mod`.
2. `GOFLAGS='-tags=tools' go mod tidy && GOFLAGS='-tags=tools' go mod vendor`
3. `go build -mod=vendor -tags=tools -o vendor_bin/goreleaser github.com/goreleaser/goreleaser/v2`
