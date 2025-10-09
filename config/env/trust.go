//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strings"
)

const defaultTrustRoot = "spike.ist"

// TrustRootFromEnv retrieves the trust root from an environment variable.
// It takes the name of an environment variable and returns its value.
// The environment variable name must start with "SPIKE_TRUST_ROOT" for
// security. If the environment variable name doesn't follow this pattern,
// is not set, or is empty, it returns the default trust root "spike.ist".
//
// Parameters:
//   - trustRootEnvVar: The name of the environment variable to read
//     (must start with "SPIKE_TRUST_ROOT")
//
// Returns:
//   - The value of the environment variable, or "spike.ist" if not set or
//     invalid name
func TrustRootFromEnv(trustRootEnvVar string) string {
	// Validate that the environment variable follows the expected pattern.
	// If the pattern does not match, return the default trust root.
	if !strings.HasPrefix(trustRootEnvVar, "SPIKE_TRUST_ROOT") {
		return defaultTrustRoot
	}

	tr := os.Getenv(trustRootEnvVar)
	if tr == "" {
		return defaultTrustRoot
	}
	return tr
}

// TrustRootVal returns the default trust root from the SPIKE_TRUST_ROOT
// environment variable. This is a convenience function that calls
// TrustRootFromEnv with the default trust root environment variable name.
//
// Returns:
//   - The value of SPIKE_TRUST_ROOT environment variable, or "spike.ist"
//     if not set
func TrustRootVal() string {
	return TrustRootFromEnv(TrustRoot)
}

// TrustRootForKeeperVal returns the trust root for SPIKE Keeper from the
// SPIKE_TRUST_ROOT_KEEPER environment variable. This is a convenience function
// that calls TrustRootFromEnv with the Keeper-specific environment variable.
//
// Returns:
//   - The value of SPIKE_TRUST_ROOT_KEEPER environment variable, or "spike.ist"
//     if not set
func TrustRootForKeeperVal() string {
	return TrustRootFromEnv(TrustRootKeeper)
}

// TrustRootForPilotVal returns the trust root for SPIKE Pilot from the
// SPIKE_TRUST_ROOT_PILOT environment variable. This is a convenience function
// that calls TrustRootFromEnv with the Pilot-specific environment variable.
//
// Returns:
//   - The value of SPIKE_TRUST_ROOT_PILOT environment variable, or "spike.ist"
//     if not set
func TrustRootForPilotVal() string {
	return TrustRootFromEnv(TrustRootPilot)
}

// TrustRootForLiteWorkloadVal returns the trust root for SPIKE Lite Workloads
// from the SPIKE_TRUST_ROOT_LITE_WORKLOAD environment variable. This is a
// convenience function that calls TrustRootFromEnv with the Lite
// Workload-specific environment variable.
//
// Returns:
//   - The value of SPIKE_TRUST_ROOT_LITE_WORKLOAD environment variable, or
//     "spike.ist" if not set
func TrustRootForLiteWorkloadVal() string {
	return TrustRootFromEnv(TrustRootLiteWorkload)
}
