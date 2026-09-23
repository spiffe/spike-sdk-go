# Decisions

<!-- INDEX:START -->
<!-- INDEX:END -->

<!-- DECISION FORMATS

## Quick Format (Y-Statement)

For lightweight decisions, a single statement suffices:

> "In the context of [situation], facing [constraint], we decided for [choice]
> and against [alternatives], to achieve [benefit], accepting that [trade-off]."

## Full Format

For significant decisions:

## [YYYY-MM-DD] Decision Title

**Status**: Accepted | Superseded | Deprecated

**Context**: What situation prompted this decision? What constraints exist?

**Alternatives Considered**:
- Option A: [Pros] / [Cons]
- Option B: [Pros] / [Cons]

**Decision**: What was decided?

**Rationale**: Why this choice over the alternatives?

**Consequence**: What are the implications? (Include both positive and negative)

**Related**: See also [other decision] | Supersedes [old decision]

## When to Record a Decision

✓ Trade-offs between alternatives
✓ Non-obvious design choices
✓ Choices that affect architecture
✓ "Why" that needs preservation

✗ Minor implementation details
✗ Routine maintenance
✗ Configuration changes
✗ No real alternatives existed

-->
## [2026-07-17-080408] Shelve config/env fatal-to-sentinel refactor; if revived, justify on caller control only

**Status**: Accepted

**Context**: SPIKE requested KeepersVal (config/env/keeper.go) stop calling log.FatalLn and instead return (map[string]string, *sdkErrors.SDKError). The env-accessor-sentinel-errors brief gave three justifications: (1) untestable without process-exit workarounds, (2) breaks the env->log coupling, (3) takes the fail-or-recover decision from callers. Another agent rolled back the brief on 2026-07-17; the spec has been removed from specs/. This records the analysis so it is not re-derived when SPIKE re-raises it.

**Decision**: Shelve config/env fatal-to-sentinel refactor; if revived, justify on caller control only

**Rationale**: Of the three justifications only (3) holds. (1) is false: fatal paths ARE unit-testable via SPIKE_STACK_TRACES_ON_LOG_FATAL panic-recover, verified against KeepersVal and already used at net/net_test.go:162 (see paired learning). (2) is marginal: the dependency already runs env->log, the safe direction; log locally re-declares its env-var consts to avoid importing env, so there is no cycle to break, only tidiness. (3) is legitimate: an accessor that calls os.Exit(1) forces process death on SPIKE Nexus/Bootstrap, which want to choose fail-or-recover; panic-recover suits tests but abuses panic as production control flow, so (map, *SDKError) is the idiomatic answer. Chose to shelve over both an additive KeepersValErr variant (leaves the fatal version around indefinitely) and doing the change now on the brief's stated grounds.

**Consequence**: No SDK change now; KeepersVal keeps fataling and stays testable via the panic-recover pattern. If revived, the refactor must be justified on caller control alone, not testability or coupling, and remains a breaking signature change SPIKE adapts (5 call sites) in the same release window. See also the paired learning on panic-recover testing.
