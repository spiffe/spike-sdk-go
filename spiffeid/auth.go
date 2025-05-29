//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import "strings"

// IsPilot checks if a given SPIFFE ID matches the SPIKE Pilot's SPIFFE ID pattern.
//
// This function is used for identity verification to determine if the provided
// SPIFFE ID belongs to a SPIKE pilot instance. It compares the input against
// the expected pilot SPIFFE ID pattern.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/pilot"
//   - Extended match with metadata: "spiffe://<trustRoot>/spike/pilot/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base pilot identity.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - id: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact pilot ID
//     or an extended ID with additional path segments for any of the trust roots,
//     false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/pilot"
//	extendedId := "spiffe://example.org/spike/pilot/instance-0"
//
//	// Both will return true
//	if IsPilot("example.org,other.org", baseId) {
//	    // Handle pilot-specific logic
//	}
//
//	if IsPilot("example.org,other.org", extendedId) {
//	    // Also recognized as a pilot, with instance metadata
//	}
func IsPilot(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		baseId := SpikePilot(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if id == baseId || strings.HasPrefix(id, baseId+"/") {
			return true
		}
	}
	return false
}

// IsPilotRecover checks if a given SPIFFE ID matches the SPIKE Pilot's
// recovery SPIFFE ID pattern.
//
// This function verifies if the provided SPIFFE ID corresponds to a SPIKE Pilot
// instance with recovery capabilities by comparing it against the expected
// recovery SPIFFE ID pattern.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/pilot/recover"
//   - Extended match with metadata: "spiffe://<trustRoot>/spike/pilot/recover/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base pilot recovery identity.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - id: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact pilot recovery ID
//     or an extended ID with additional path segments for any of the trust roots,
//     false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/pilot/recover"
//	extendedId := "spiffe://example.org/spike/pilot/recover/instance-0"
//
//	// Both will return true
//	if IsPilotRecover("example.org,other.org", baseId) {
//	    // Handle recovery-specific logic
//	}
//
//	if IsPilotRecover("example.org,other.org", extendedId) {
//	    // Also recognized as a pilot recovery, with instance metadata
//	}
func IsPilotRecover(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		baseId := SpikePilotRecover(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if id == baseId || strings.HasPrefix(id, baseId+"/") {
			return true
		}
	}
	return false
}

// IsPilotRestore checks if a given SPIFFE ID matches the SPIKE Pilot's restore
// SPIFFE ID pattern.
//
// This function verifies if the provided SPIFFE ID corresponds to a pilot
// instance with restore capabilities by comparing it against the expected
// restore SPIFFE ID pattern.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/pilot/restore"
//   - Extended match with metadata: "spiffe://<trustRoot>/spike/pilot/restore/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base pilot restore identity.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - id: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact pilot restore ID
//     or an extended ID with additional path segments for any of the trust roots,
//     false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/pilot/restore"
//	extendedId := "spiffe://example.org/spike/pilot/restore/instance-0"
//
//	// Both will return true
//	if IsPilotRestore("example.org,other.org", baseId) {
//	    // Handle restore-specific logic
//	}
//
//	if IsPilotRestore("example.org,other.org", extendedId) {
//	    // Also recognized as a pilot restore, with instance metadata
//	}
func IsPilotRestore(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		baseId := SpikePilotRestore(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if id == baseId || strings.HasPrefix(id, baseId+"/") {
			return true
		}
	}
	return false
}

// IsKeeper checks if a given SPIFFE ID matches the SPIKE Keeper's SPIFFE ID.
//
// This function is used for identity verification to determine if the provided
// SPIFFE ID belongs to a SPIKE Keeper instance. It compares the input against
// the expected keeper SPIFFE ID pattern.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/keeper"
//   - Extended match with metadata: "spiffe://<trustRoot>/spike/keeper/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base keeper identity.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - id: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact
//     SPIKE Keeper's ID or an extended ID with additional path segments for any
//     of the trust roots, false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/keeper"
//	extendedId := "spiffe://example.org/spike/keeper/instance-0"
//
//	// Both will return true
//	if IsKeeper("example.org", baseId) {
//	    // Handle keeper-specific logic
//	}
//
//	if IsKeeper("example.org", extendedId) {
//	    // Also recognized as a keeper, with instance metadata
//	}
func IsKeeper(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		baseId := SpikeKeeper(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if id == baseId || strings.HasPrefix(id, baseId+"/") {
			return true
		}
	}
	return false
}

// IsNexus checks if the provided SPIFFE ID matches the SPIKE Nexus SPIFFE ID.
//
// The function compares the input SPIFFE ID against the configured Spike Nexus
// SPIFFE ID pattern. This is typically used for validating whether a given
// identity represents the Nexus service.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/nexus"
//   - Extended match with metadata: "spiffe://<trustRoot>/spike/nexus/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base Nexus identity.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - id: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches either the exact Nexus SPIFFE ID
//     or an extended ID with additional path segments for any of the
//     trust roots, false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/nexus"
//	extendedId := "spiffe://example.org/spike/nexus/instance-0"
//
//	// Both will return true
//	if IsNexus("example.org", baseId) {
//	    // Handle Nexus-specific logic
//	}
//
//	if IsNexus("example.org", extendedId) {
//	    // Also recognized as a Nexus, with instance metadata
//	}
func IsNexus(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		baseId := SpikeNexus(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if id == baseId || strings.HasPrefix(id, baseId+"/") {
			return true
		}
	}
	return false
}

// PeerCanTalkToAnyone is used for debugging purposes
func PeerCanTalkToAnyone(_, _ string) bool {
	return true
}

// PeerCanTalkToKeeper checks if the provided SPIFFE ID matches the SPIKE Nexus
// SPIFFE ID.
//
// This is used as a validator in SPIKE Keeper because currently only SPIKE
// Nexus can talk to SPIKE Keeper.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - peerSpiffeId: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches SPIKE Nexus' SPIFFE ID for any of
//     the trust roots, false otherwise
func PeerCanTalkToKeeper(trustRoots, peerSpiffeId string) bool {
	return IsNexus(trustRoots, peerSpiffeId)
}
