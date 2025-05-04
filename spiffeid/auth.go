//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

// IsPilot checks if a given SPIFFE ID matches the SPIKE Pilot's SPIFFE ID.
//
// This function is used for identity verification to determine if the provided
// SPIFFE ID belongs to a SPIKE pilot instance. It compares the input against
// the expected pilot SPIFFE ID returned by SpikePilotSpiffeId().
//
// Parameters:
//   - trustRoot: The trust domain root (e.g., "example.org")
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the pilot's ID, false
//     otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/pilot"
//	if IsPilot("example.org", id) {
//	    // Handle pilot-specific logic
//	}
func IsPilot(trustRoot, id string) bool {
	return id == SpikePilot(trustRoot)
}

// IsPilotRecover checks if a given SPIFFE ID matches the SPIKE Pilot's
// recovery SPIFFE ID.
//
// This function verifies if the provided SPIFFE ID corresponds to a SPIKE Pilot
// instance with recovery capabilities by comparing it against the expected
// recovery SPIFFE ID returned by SpikePilotRecoverSpiffeId().
//
// Parameters:
//   - trustRoot: The trust domain root (e.g., "example.org")
//   - spiffeId: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the pilot's recovery ID,
//     false otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/pilot/recover"
//	if IsPilotRecover("example.org", id) {
//	    // Handle recovery-specific logic
//	}
func IsPilotRecover(trustRoot, id string) bool {
	return id == SpikePilotRecover(trustRoot)
}

// IsPilotRestore checks if a given SPIFFE ID matches the SPIKE Pilot's restore
// SPIFFE ID.
//
// This function verifies if the provided SPIFFE ID corresponds to a pilot
// instance with restore capabilities by comparing it against the expected
// restore SPIFFE ID returned by SpikePilotRestoreSpiffeId().
//
// Parameters:
//   - trustRoot: The trust domain root (e.g., "example.org")
//   - spiffeId: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the pilot's restore ID,
//     false otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/pilot/restore"
//	if IsPilotRestore("example.org", id) {
//	    // Handle restore-specific logic
//	}
func IsPilotRestore(trustRoot, spiffeId string) bool {
	return spiffeId == SpikePilotRestore(trustRoot)
}

// IsKeeper checks if a given SPIFFE ID matches the SPIKE Keeper's SPIFFE ID.
//
// This function is used for identity verification to determine if the provided
// SPIFFE ID belongs to a SPIKE Keeper instance. It compares the input against
// the expected pilot SPIFFE ID returned by SpikeKeeperSpiffeId().
//
// Parameters:
//   - trustRoot: The trust domain root (e.g., "example.org")
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches the SPIKE Keeper's ID, false
//     otherwise
//
// Example usage:
//
//	id := "spiffe://example.org/spike/keeper"
//	if IsKeeper("example.org", id) {
//	    // Handle pilot-specific logic
//	}
func IsKeeper(trustRoot, id string) bool {
	return id == SpikeKeeper(trustRoot)
}

// IsNexus checks if the provided SPIFFE ID matches the SPIKE Nexus SPIFFE ID.
//
// The function compares the input SPIFFE ID against the configured Spike Nexus
// SPIFFE ID from the environment. This is typically used for validating whether
// a given identity represents the Nexus service.
//
// Parameters:
//   - trustRoot: The trust domain root (e.g., "example.org")
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches the Nexus SPIFFE ID, false otherwise
func IsNexus(trustRoot, id string) bool {
	return id == SpikeNexus(trustRoot)
}

// PeerCanTalkToAnyone is used for debugging purposes
func PeerCanTalkToAnyone(_, _ string) bool {
	return true
}

// PeerCanTalkToKeeper checks if the provided SPIFFE ID matches the SPIKE Nexus
// SPIFFE ID.
//
// This is used as a validator in SPIKE Keeper, because currently only SPIKE
// Nexus can talk to SPIKE Keeper.
//
// Parameters:
//   - trustRoot: The trust domain root (e.g., "example.org")
//   - peerSpiffeId: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches SPIKE Nexus' SPIFFE ID,
//     false otherwise
func PeerCanTalkToKeeper(trustRoot, peerSpiffeId string) bool {
	return peerSpiffeId == SpikeNexus(trustRoot)
}
