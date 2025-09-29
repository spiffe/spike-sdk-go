//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import (
	"path"
	"strings"

	"github.com/spiffe/spike-sdk-go/log"
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
	const fName = "Keeper"
	if trustRoot == "" {
		log.FatalLn(fName, "message", "Keeper: trustRoot cannot be an empty string")
	}
	if strings.Contains(trustRoot, ",") {
		log.FatalLn(fName, "message", "Keeper: provide a single trust root")
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
	const fName = "Nexus"
	if trustRoot == "" {
		log.FatalLn(fName, "message", "Nexus: trustRoot cannot be an empty string")
	}
	if strings.Contains(trustRoot, ",") {
		log.FatalLn(fName, "message", "Nexus: provide a single trust root")
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
	const fName = "Pilot"
	if trustRoot == "" {
		log.FatalLn(fName, "message", "Pilot: trustRoot cannot be an empty string")
	}
	if strings.Contains(trustRoot, ",") {
		log.FatalLn(fName, "message", "Pilot: provide a single trust root")
	}

	return "spiffe://" + path.Join(trustRoot,
		"spike", "pilot", "role", "superuser")
}

// Bootstrap generates the SPIFFE ID for a SPIKE Bootstrap role.
//
// Parameters:
//   - trustRoot: The trust domain for the SPIFFE ID. If empty, the value is
//     obtained from the environment.
//
// Returns:
//   - string: The complete SPIFFE ID in the format:
//     "spiffe://<trustRoot>/spike/bootstrap"
func Bootstrap(trustRoot string) string {
	const fName = "Bootstrap"
	if trustRoot == "" {
		log.FatalLn(fName,
			"message", "Bootstrap: trustRoot cannot be an empty string")
	}
	if strings.Contains(trustRoot, ",") {
		log.FatalLn(fName, "message", "Bootstrap: provide a single trust root")
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "bootstrap")
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
	const fName = "LiteWorkload"
	if trustRoot == "" {
		log.FatalLn(fName,
			"message", "LiteWorkload: trustRoot cannot be an empty string")
	}
	if strings.Contains(trustRoot, ",") {
		log.FatalLn(fName, "message", "LiteWorkload: provide a single trust root")
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
	const fName = "PilotRecover"
	if trustRoot == "" {
		log.FatalLn(fName, "message",
			"PilotRecover: trustRoot cannot be an empty string")
	}
	if strings.Contains(trustRoot, ",") {
		log.FatalLn(fName, "message",
			"PilotRecover: provide a single trust root")
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
	const fName = "PilotRestore"
	if trustRoot == "" {
		log.FatalLn(fName, "message",
			"PilotRestore: trustRoot cannot be an empty string")
	}
	if strings.Contains(trustRoot, ",") {
		log.FatalLn(fName, "message", "PilotRestore: provide a single trust root")
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "restore")
}
