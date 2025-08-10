//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import (
	"path"

	"github.com/spiffe/spike-sdk-go/spiffeid/internal/env"
)

// Keeper constructs and returns the SPIKE Keeper's SPIFFE ID string.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/keeper"
func Keeper(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "keeper")
}

// Nexus constructs and returns the SPIFFE ID for SPIKE Nexus.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/nexus"
func Nexus(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "nexus")
}

// Pilot generates the SPIFFE ID for a SPIKE Pilot superuser role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/pilot/role/superuser"
func Pilot(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "superuser")
}

// LiteWorkload generates the SPIFFE ID for a SPIKE Lite workload role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/workload/role/lite"
func LiteWorkload(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "workload", "role", "lite")
}

// PilotRecover generates the SPIFFE ID for a SPIKE Pilot recovery role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/pilot/role/recover"
func PilotRecover(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "recover")
}

// PilotRestore generates the SPIFFE ID for a SPIKE Pilot restore role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/pilot/role/restore"
func PilotRestore(trustRoot string) string {
	if trustRoot == "" {
		trustRoot = env.TrustRoot()
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "restore")
}
