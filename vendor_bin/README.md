# Vendored Binaries

This directory contains the exact GoReleaser package used by Linux `amd64` release builds. Normal development builds use the pinned Go toolchain directly and do not require GoReleaser.

`script/build --release` verifies `goreleaser_2.11.2_amd64.deb` against its committed SHA-256 file, extracts it into the ignored `.tools/` directory without global installation, verifies the reported version, and invokes that binary. Release builds therefore do not trust an ambient GoReleaser installation.

Vendored release tools must have a committed checksum and proper upstream attestations before they are added here. For example:

```bash
$ gh attestation verify --repo goreleaser/goreleaser ./vendor_bin/goreleaser_2.11.2_amd64.deb
Loaded digest sha256:a9ddd4791bc0f862a665dbd7a8f077cf2861fc6d41c153c47252243cea3c1d67 for file://vendor_bin/goreleaser_2.11.2_amd64.deb
Loaded 1 attestation from GitHub API

The following policy criteria will be enforced:
- Predicate type must match:................ https://slsa.dev/provenance/v1
- Source Repository Owner URI must match:... https://github.com/goreleaser
- Source Repository URI must match:......... https://github.com/goreleaser/goreleaser
- Subject Alternative Name must match regex: (?i)^https://github.com/goreleaser/goreleaser/
- OIDC Issuer must match:................... https://token.actions.githubusercontent.com

✓ Verification succeeded!

The following 1 attestation matched the policy criteria

- Attestation #1
  - Build repo:..... goreleaser/goreleaser
  - Build workflow:. .github/workflows/release.yml@refs/tags/v2.11.2
  - Signer repo:.... goreleaser/goreleaser
  - Signer workflow: .github/workflows/release.yml@refs/tags/v2.11.2
```
