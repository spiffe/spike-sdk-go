//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import (
	"strings"

	"github.com/spiffe/spike-sdk-go/config/env"
)

// IsPilot checks if a given SPIFFE ID matches the SPIKE Pilot's SPIFFE ID
// pattern.
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
//   - SPIFFEID: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact pilot ID
//     or an extended ID with additional path segments for any of the trust
//     roots, false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/pilot"
//	extendedId := "spiffe://example.org/spike/pilot/instance-0"
//
//	// Both will return true
//	if IsPilot(baseId) {
//	    // Handle pilot-specific logic
//	}
//
//	if IsPilot(extendedId) {
//	    // Also recognized as a SPIKE Pilot, with instance metadata
//	}
func IsPilot(SPIFFEID string) bool {
	trustRoots := env.TrustRootFromEnv(env.TrustRootPilot)
	for _, root := range strings.Split(trustRoots, ",") {
		baseID := Pilot(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if SPIFFEID == baseID || strings.HasPrefix(SPIFFEID, baseID+"/") {
			return true
		}
	}
	return false
}

// IsLiteWorkload checks if a given SPIFFE ID matches the SPIKE Lite Workload's
// SPIFFE ID pattern.
//
// A SPIKE Lite workload can freely use SPIKE Nexus encryption and decryption
// RESTful APIs without needing any specific policies assigned to it. A SPIKE
// Lite workload cannot use any other SPIKE Nexus API unless a relevant policy
// is attached to it.
//
// This function is used for identity verification to determine if the provided
// SPIFFE ID belongs to a SPIKE lite workload instance. It compares the input
// against the expected lite workload SPIFFE ID pattern.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/workload/role/lite"
//   - Extended match with metadata:
//     "spiffe://<trustRoot>/spike/workload/role/lite/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base lite workload identity.
//
// Parameters:
//   - SPIFFEID: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact lite
//     workload ID or an extended ID with additional path segments for any of
//     the trust roots, false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/workload/role/lite"
//	extendedId := "spiffe://example.org/spike/workload/role/lite/instance-0"
//
//	// Both will return true
//	if IsLiteWorkload(baseId) {
//	    // Handle lite workload-specific logic
//	}
//
//	if IsLiteWorkload(extendedId) {
//	    // Also recognized as a SPIKE Lite Workload, with instance metadata
//	}
func IsLiteWorkload(SPIFFEID string) bool {
	trustRoots := env.TrustRootFromEnv(env.TrustRootLiteWorkload)
	for _, root := range strings.Split(trustRoots, ",") {
		baseID := LiteWorkload(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if SPIFFEID == baseID || strings.HasPrefix(SPIFFEID, baseID+"/") {
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
//   - Extended match with metadata:
//     "spiffe://<trustRoot>/spike/pilot/recover/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base pilot recovery identity.
//
// Parameters:
//   - SPIFFEID: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact pilot
//     recovery ID or an extended ID with additional path segments for any of
//     the trust roots, false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/pilot/recover"
//	extendedId := "spiffe://example.org/spike/pilot/recover/instance-0"
//
//	// Both will return true
//	if IsPilotRecover(baseId) {
//	    // Handle recovery-specific logic
//	}
//
//	if IsPilotRecover(extendedId) {
//	    // Also recognized as a SPIKE Pilot recovery, with instance metadata
//	}
func IsPilotRecover(SPIFFEID string) bool {
	trustRoots := env.TrustRootFromEnv(env.TrustRootPilot)
	for _, root := range strings.Split(trustRoots, ",") {
		baseID := PilotRecover(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if SPIFFEID == baseID || strings.HasPrefix(SPIFFEID, baseID+"/") {
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
//   - Extended match with metadata:
//     "spiffe://<trustRoot>/spike/pilot/restore/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base pilot restore identity.
//
// Parameters:
//   - SPIFFEID: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact pilot
//     restore ID or an extended ID with additional path segments for any of the
//     trust roots, false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/pilot/restore"
//	extendedId := "spiffe://example.org/spike/pilot/restore/instance-0"
//
//	// Both will return true
//	if IsPilotRestore(baseId) {
//			// Handle restore-specific logic
//	}
//
//	if IsPilotRestore(extendedId) {
//			// Also recognized as a SPIKE Pilot restore, with instance metadata
//	}
func IsPilotRestore(SPIFFEID string) bool {
	trustRoots := env.TrustRootFromEnv(env.TrustRootPilot)
	for _, root := range strings.Split(trustRoots, ",") {
		baseID := PilotRestore(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if SPIFFEID == baseID || strings.HasPrefix(SPIFFEID, baseID+"/") {
			return true
		}
	}
	return false
}

// IsBootstrap checks if a given SPIFFE ID matches the SPIKE Bootstrap's
// SPIFFE ID pattern.
//
// This function verifies if the provided SPIFFE ID corresponds to a bootstrap
// instance by comparing it against the expected bootstrap SPIFFE ID pattern.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/bootstrap"
//   - Extended match with metadata:
//     "spiffe://<trustRoot>/spike/bootstrap/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base bootstrap identity.
//
// Parameters:
//   - SPIFFEID: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the provided SPIFFE ID matches either the exact bootstrap
//     ID or an extended ID with additional path segments for any of the
//     trust roots, false otherwise
//
// Example usage:
//
//	baseId := "spiffe://example.org/spike/bootstrap"
//	extendedId := "spiffe://example.org/spike/bootstrap/instance-0"
//
//	// Both will return true
//	if IsBootstrap(baseId) {
//			// Handle bootstrap-specific logic
//	}
//
//	if IsBootstrap(extendedId) {
//			// Also recognized as a SPIKE Bootstrap, with instance metadata
//	}
func IsBootstrap(SPIFFEID string) bool {
	trustRoots := env.TrustRootFromEnv(env.TrustRootBootstrap)
	for _, root := range strings.Split(trustRoots, ",") {
		baseID := Bootstrap(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if SPIFFEID == baseID || strings.HasPrefix(SPIFFEID, baseID+"/") {
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
//   - Extended match with metadata:
//     "spiffe://<trustRoot>/spike/keeper/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base keeper identity.
//
// Parameters:
//   - SPIFFEID: The SPIFFE ID string to check
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
//	if IsKeeper(baseId) {
//	    // Handle keeper-specific logic
//	}
//
//	if IsKeeper(extendedId) {
//	    // Also recognized as a SPIKE Keeper, with instance metadata
//	}
func IsKeeper(SPIFFEID string) bool {
	trustRoots := env.TrustRootFromEnv(env.TrustRootKeeper)
	for _, root := range strings.Split(trustRoots, ",") {
		baseID := Keeper(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if SPIFFEID == baseID || strings.HasPrefix(SPIFFEID, baseID+"/") {
			return true
		}
	}
	return false
}

// IsNexus checks if the provided SPIFFE ID matches the SPIKE Nexus SPIFFE ID.
//
// The function compares the input SPIFFE ID against the configured SPIKE Nexus
// SPIFFE ID pattern. This is typically used for validating whether a given
// identity represents the Nexus service.
//
// The function supports two formats:
//   - Exact match: "spiffe://<trustRoot>/spike/nexus"
//   - Extended match with metadata:
//     "spiffe://<trustRoot>/spike/nexus/<metadata>"
//
// This allows for instance-specific identifiers while maintaining compatibility
// with the base Nexus identity.
//
// Parameters:
//   - SPIFFEID: The SPIFFE ID string to check
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
//	if IsNexus(baseId) {
//	    // Handle Nexus-specific logic
//	}
//
//	if IsNexus(extendedId) {
//	    // Also recognized as a SPIKE Nexus, with instance metadata
//	}
func IsNexus(SPIFFEID string) bool {
	trustRoots := env.TrustRootFromEnv(env.TrustRootNexus)
	for _, root := range strings.Split(trustRoots, ",") {
		baseID := Nexus(strings.TrimSpace(root))
		// Check if the ID is either exactly the base ID or starts with the base ID
		// followed by "/"
		if SPIFFEID == baseID || strings.HasPrefix(SPIFFEID, baseID+"/") {
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
//   - peerSPIFFEID: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches SPIKE Nexus' or SPIKE Bootstrap's
//     SPIFFE ID for any of the trust roots, false otherwise
func PeerCanTalkToKeeper(peerSPIFFEID string) bool {
	return IsNexus(peerSPIFFEID) || IsBootstrap(peerSPIFFEID)
}
