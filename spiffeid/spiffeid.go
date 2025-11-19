//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import (
	"path"
	"strings"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
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
		failErr := sdkErrors.ErrSPIFFEEmptyTrustDomain
		log.FatalErr(fName, *failErr)
	}
	if strings.Contains(trustRoot, ",") {
		failErr := sdkErrors.ErrSPIFFEMultipleTrustDomains
		log.FatalErr(fName, *failErr)
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
		failErr := sdkErrors.ErrSPIFFEEmptyTrustDomain
		log.FatalErr(fName, *failErr)
	}
	if strings.Contains(trustRoot, ",") {
		failErr := sdkErrors.ErrSPIFFEMultipleTrustDomains
		log.FatalErr(fName, *failErr)
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
		failErr := sdkErrors.ErrSPIFFEEmptyTrustDomain
		log.FatalErr(fName, *failErr)
	}
	if strings.Contains(trustRoot, ",") {
		failErr := sdkErrors.ErrSPIFFEMultipleTrustDomains
		log.FatalErr(fName, *failErr)
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
		failErr := sdkErrors.ErrSPIFFEEmptyTrustDomain
		log.FatalErr(fName, *failErr)
	}
	if strings.Contains(trustRoot, ",") {
		failErr := sdkErrors.ErrSPIFFEMultipleTrustDomains
		log.FatalErr(fName, *failErr)
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
		failErr := sdkErrors.ErrSPIFFEEmptyTrustDomain
		log.FatalErr(fName, *failErr)
	}
	if strings.Contains(trustRoot, ",") {
		failErr := sdkErrors.ErrSPIFFEMultipleTrustDomains
		log.FatalErr(fName, *failErr)
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
		failErr := sdkErrors.ErrSPIFFEEmptyTrustDomain
		log.FatalErr(fName, *failErr)
	}
	if strings.Contains(trustRoot, ",") {
		failErr := sdkErrors.ErrSPIFFEMultipleTrustDomains
		log.FatalErr(fName, *failErr)
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
		failErr := sdkErrors.ErrSPIFFEEmptyTrustDomain
		log.FatalErr(fName, *failErr)
	}
	if strings.Contains(trustRoot, ",") {
		failErr := sdkErrors.ErrSPIFFEMultipleTrustDomains
		log.FatalErr(fName, *failErr)
	}

	return "spiffe://" + path.Join(trustRoot, "spike", "pilot", "role", "restore")
}
