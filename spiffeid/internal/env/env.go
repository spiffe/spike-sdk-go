//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "os"

// TrustRoot returns the trust root domain for SPIKE operations.
// It first checks the SPIKE_TRUST_ROOT environment variable.
// If the environment variable is not set or empty, it defaults to "spike.ist".
func TrustRoot() string {
	tr := os.Getenv("SPIKE_TRUST_ROOT")
	if tr == "" {
		return "spike.ist"
	}
	return tr
}
