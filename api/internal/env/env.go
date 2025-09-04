//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"

	appEnv "github.com/spiffe/spike-sdk-go/config/env"
)

// NexusAPIRoot returns the URL of the Nexus API.
func NexusAPIRoot() string {
	p := os.Getenv(appEnv.NexusAPIURL)
	if p != "" {
		return p
	}
	return "https://localhost:8553"
}
