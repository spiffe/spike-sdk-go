//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "os"

func TrustRoot() string {
	tr := os.Getenv("SPIKE_TRUST_ROOT")
	if tr == "" {
		return "spike.ist"
	}
	return tr
}
