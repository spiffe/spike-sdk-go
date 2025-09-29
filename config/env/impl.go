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
