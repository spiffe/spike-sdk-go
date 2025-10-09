//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "os"

// NexusAPIRootVal retrieves the SPIKE Nexus API root URL from the environment.
// It reads the value from the SPIKE_NEXUS_API_URL environment variable.
//
// Returns:
//   - The Nexus API root URL from the environment variable if set
//   - The default value of "https://localhost:8553" if the environment
//     variable is unset or empty
func NexusAPIRootVal() string {
	apiRoot := os.Getenv(NexusAPIURL)
	if apiRoot == "" {
		return "https://localhost:8553"
	}
	return apiRoot
}
