# Learnings

<!--
UPDATE WHEN:
- Discover a gotcha, bug, or unexpected behavior
- Debugging reveals non-obvious root cause
- External dependency has quirks worth documenting
- "I wish I knew this earlier" moments
- Production incidents reveal gaps

DO NOT UPDATE FOR:
- Well-documented behavior (link to docs instead)
- Temporary workarounds (use TASKS.md for follow-up)
- Opinions without evidence
-->

<!-- INDEX:START -->
<!-- INDEX:END -->

<!-- Add gotchas, tips, and lessons learned here -->
## [2026-09-20-131942] Local golangci-lint typecheck failure is pre-existing

**Context**: Ran golangci-lint@latest (as make audit does) on PR #195 with Go 1.27.1; it reported 'undefined: backoff' in retry/retry.go

**Lesson**: The same failure occurs on main; it is a toolchain mismatch between golangci-lint@latest and Go 1.27.1, not a code defect. go build and go test pass.

**Application**: Compare lint output against main before attributing failures to a PR branch.

---

## [2026-07-17-080343] log.FatalLn fatal paths are unit-testable via panic-recover

**Context**: The env-accessor-sentinel-errors brief argued config/env.KeepersVal 'cannot be unit-tested without process-exit workarounds', justifying a breaking signature change. Probed the claim on 2026-07-17 before accepting it.

**Lesson**: Every log.FatalLn path in this SDK is testable in-process: fatalExit (log/fatal.go) panics instead of os.Exit(1) when SPIKE_STACK_TRACES_ON_LOG_FATAL=true. A test sets it with t.Setenv, calls the fatal-calling function, and recover()s in a deferred func to assert on the panic message. Verified against KeepersVal's unset/malformed/duplicate/valid paths — all four passed without killing the test process. Precedent already lives at net/net_test.go:162 (CreateMTLSServer nil-source).

**Application**: Never accept 'a fatal-calling accessor is untestable' as a reason to change its signature. To test a fatal path: t.Setenv("SPIKE_STACK_TRACES_ON_LOG_FATAL","true") + deferred recover(). Reserve fatal-to-sentinel refactors for the real reason: giving callers the fail-or-recover decision.
