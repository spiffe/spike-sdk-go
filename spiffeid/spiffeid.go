//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import (
	"path"

	"github.com/spiffe/spike-sdk-go/spiffeid/internal/env"
)

// SpikeKeeper constructs and returns the SPIKE Keeper's SPIFFE ID string.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/keeper"
func SpikeKeeper(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "keeper")
}

// SpikeNexus constructs and returns the SPIFFE ID for SPIKE Nexus.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/nexus"
func SpikeNexus(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "nexus")
}

// SpikePilot generates the SPIFFE ID for a SPIKE Pilot superuser role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/pilot/role/superuser"
func SpikePilot(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "superuser")
}

// SpikeLiteWorkload generates the SPIFFE ID for a SPIKE Lite workload role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/workload/role/lite"
func SpikeLiteWorkload(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "workload", "role", "lite")
}

// SpikePilotRecover generates the SPIFFE ID for a SPIKE Pilot recovery role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/pilot/role/recover"
func SpikePilotRecover(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "recover")
}

// SpikePilotRestore generates the SPIFFE ID for a SPIKE Pilot restore role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/pilot/role/restore"
func SpikePilotRestore(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "restore")
}
