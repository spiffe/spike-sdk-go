//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "os"

// NexusAPIRoot returns the URL of the Nexus API.
func NexusAPIRoot() string {
	p := os.Getenv("SPIKE_NEXUS_API_URL")
	if p != "" {
		return p
	}
	return "https://localhost:8553"
}
