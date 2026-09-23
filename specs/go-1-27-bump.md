# Spec: Align the Go toolchain and tooling with SPIKE

## Problem Statement

SPIKE requires Go 1.27.1 (`go.mod`, all builder images) and runs
golangci-lint v2 (`specs/go-1-27-bump.md` in the SPIKE repository).
The SDK still requires Go 1.25.5 and runs golangci-lint through the v1
module path, which `@latest` resolves to v1.64.8, the final v1 release.
v1.64.8 cannot typecheck under Go 1.27 (`undefined: backoff` in
`retry/retry.go`), so the golangci-lint step of `make audit` fails on
the toolchain SPIKE uses, independently of the code.

CI reads the Go version from `go.mod` (`go-version-file`), so the
directive is the single switch for CI and local builds.

## Proposed Solution

1. `go.mod`: `go 1.25.5` -> `go 1.27.1`, no `toolchain` directive
   (same as SPIKE). `go mod tidy` changes nothing else.
2. `Makefile`: the audit's golangci-lint step runs
   `github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest`.
3. `.golangci.yml`: migrated to the v2 format with
   `golangci-lint migrate`; same linter selection:
   - `gosimple` and `stylecheck` are part of `staticcheck` in v2.
   - `goimports` moves to `formatters`.
   - The v1 default exclusions become the `comments`,
     `common-false-positives`, `legacy`, and `std-error-handling`
     presets.
   - `run.timeout: 1m` is dropped; v2 has no timeout by default.
4. Fix every finding the v2 linters report, at the site, with no
   `//nolint` and no new exclusions (SPIKE's `lint-hardening.md` rule).

## Findings (golangci-lint v2.13.2, uncapped: 20)

Baseline: golangci-lint v1.64.8 on Go 1.25.5 at `1801bd7` reports 0
issues with the same linters. All 20 findings come from newer analyzer
versions (gosec 2.28 G703 taint analysis, extended G115, staticcheck QF
checks, prealloc).

- `strings/char.go` G115, **a real defect**: `byte(len(chars))` wraps
  to 0 when the class covers all 256 byte values (e.g. `\x00-\xff`), so
  `secureRandomStringFromCharClass` panics with a division by zero. The
  class comes from a caller-supplied template (`StringFromTemplate`).
  The same line has modulo bias: whenever 256 is not a multiple of the
  set size, the first `256 % n` characters are more likely. Fix:
  rejection sampling in `int` arithmetic, so every character is equally
  likely and a full 256-byte set works.
- `strings/char.go`, found while testing the above, **a real defect**:
  `expandCharacterClass` loops `for c := start; c <= end; c++` over a
  `byte`, which never terminates when the range ends at `0xff`; a
  template range such as `\x00-\xff` hangs the caller. Fix: stop on
  `c == end` inside the loop.
- `config/fs/init.go` G703 x4:
  - `createTemp*Dir` builds `/tmp/.spike-$USER` from the raw `USER`
    variable; `USER=/../../etc` resolves outside `/tmp`. Fix: reduce
    `USER` to a single path element with `filepath.Base`, falling back to
    `spike` for `.`, `..`, and separator-only values.
  - `tryCustom*Dir` validates the absolute form of the environment path
    (`validateDataDirectory` calls `filepath.Abs`) but creates the raw
    one. Fix: the caller resolves the path once with `filepath.Abs`, and
    validates and creates that same absolute path. (Returning the path
    from `validateDataDirectory` was tried first; gosec's taint analysis
    does not follow sanitization across the return, so the resolution
    lives in the caller.)
- `strings/char.go`, `strings/char_test.go` QF1001: apply De Morgan.
- `strings/char_test.go` G115 x3: iterate the ASCII result as bytes, no
  rune -> byte conversion.
- `crypto/crypto_test.go` G115 x4: the fake readers keep their counters
  as `byte`, which wraps exactly like the old `byte(int)` conversion.
- `predicate/predicate_test.go` QF1011 x4: `var _ Predicate = X` is
  tautological for values declared as `Predicate(...)`, and would still
  compile if one became a plain `func`. Replace with
  `assert.IsType(t, Predicate(nil), X)`, which checks the declared type.
- `retry/retry.go` prealloc x2: allocate the option slice with capacity.

## Non-Goals

- No vulnerability remediation (govulncheck's findings are the next,
  separate change).
- No adoption of SPIKE's hardened lint gate (strict errcheck, all
  staticcheck checks); the SDK keeps its own linter selection.
- No change to the standalone staticcheck step.

## Verification (2026-09-23, go1.27.1 linux/arm64, GOTOOLCHAIN=local)

- `go mod tidy -diff` empty, `go mod verify`, `gofmt`, `go vet`,
  staticcheck: pass. `make test` (race): all packages ok.
- golangci-lint v2.13.2, uncapped: 0 issues.
- govulncheck: 1 called finding, GO-2026-5970 in `golang.org/x/text`
  v0.37.0 (fixed in v0.39.0), present on the base commit; it is the
  next change, so `make audit` still exits non-zero at that step.
- New tests: `TestSecureRandomStringFromCharClass_FullByteRange` (times
  out on the old code), `_ZeroLength`; `config/fs/init_test.go`
  (`USER` sanitization, custom directories created at the validated
  absolute path, restricted path refused), all failing on the old code.
