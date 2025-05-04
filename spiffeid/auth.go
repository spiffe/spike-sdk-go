//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import "strings"

// IsPilot checks if a given SPIFFE ID matches the SPIKE Pilot's SPIFFE ID.
//
// This function is used for identity verification to determine if the provided
// SPIFFE ID belongs to a SPIKE pilot instance. It compares the input against
// the expected pilot SPIFFE ID returned by SpikePilotSpiffeId().
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the pilot's ID for any of
//     the trust roots, false otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/pilot"
//	if IsPilot("example.org,other.org", id) {
//	    // Handle pilot-specific logic
//	}
func IsPilot(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		if id == SpikePilot(strings.TrimSpace(root)) {
			return true
		}
	}
	return false
}

// IsPilotRecover checks if a given SPIFFE ID matches the SPIKE Pilot's
// recovery SPIFFE ID.
//
// This function verifies if the provided SPIFFE ID corresponds to a SPIKE Pilot
// instance with recovery capabilities by comparing it against the expected
// recovery SPIFFE ID returned by SpikePilotRecoverSpiffeId().
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - spiffeId: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the pilot's recovery ID for
//     any of the trust roots, false otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/pilot/recover"
//	if IsPilotRecover("example.org,other.org", id) {
//	    // Handle recovery-specific logic
//	}
func IsPilotRecover(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		if id == SpikePilotRecover(strings.TrimSpace(root)) {
			return true
		}
	}
	return false
}

// IsPilotRestore checks if a given SPIFFE ID matches the SPIKE Pilot's restore
// SPIFFE ID.
//
// This function verifies if the provided SPIFFE ID corresponds to a pilot
// instance with restore capabilities by comparing it against the expected
// restore SPIFFE ID returned by SpikePilotRestoreSpiffeId().
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - spiffeId: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the pilot's restore ID for
//     any of the trust roots, false otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/pilot/restore"
//	if IsPilotRestore("example.org,other.org", id) {
//	    // Handle restore-specific logic
//	}
func IsPilotRestore(trustRoots, spiffeId string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		if spiffeId == SpikePilotRestore(strings.TrimSpace(root)) {
			return true
		}
	}
	return false
}

// IsKeeper checks if a given SPIFFE ID matches the SPIKE Keeper's SPIFFE ID.
//
// This function is used for identity verification to determine if the provided
// SPIFFE ID belongs to a SPIKE Keeper instance. It compares the input against
// the expected pilot SPIFFE ID returned by SpikeKeeperSpiffeId().
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the SPIKE Keeper's ID for
//     and of the trust roots, false otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/keeper"
//	if IsKeeper("example.org,other.org", id) {
//	    // Handle pilot-specific logic
//	}
func IsKeeper(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		if id == SpikeKeeper(strings.TrimSpace(root)) {
			return true
		}
	}
	return false
}

// IsNexus checks if the provided SPIFFE ID matches the SPIKE Nexus SPIFFE ID.
//
// The function compares the input SPIFFE ID against the configured Spike Nexus
// SPIFFE ID from the environment. This is typically used for validating whether
// a given identity represents the Nexus service.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots
//     (e.g., "example.org,other.org")
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches the Nexus SPIFFE ID for any of the
//     trust roots, false otherwise
func IsNexus(trustRoots, id string) bool {
	for _, root := range strings.Split(trustRoots, ",") {
		if id == SpikeNexus(strings.TrimSpace(root)) {
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
	for _, root := range strings.Split(trustRoots, ",") {
		if peerSpiffeId == SpikeNexus(strings.TrimSpace(root)) {
			return true
		}
	}
	return false
}
