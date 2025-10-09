//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strconv"
)

// MaxSecretVersionsVal returns the maximum number of versions to retain
// for each secret. It reads from the SPIKE_NEXUS_MAX_SECRET_VERSIONS
// environment variable which should contain a positive integer value.
// If the environment variable is not set, contains an invalid integer, or
// specifies a non-positive value, it returns the default of 10 versions.
func MaxSecretVersionsVal() int {
	p := os.Getenv(NexusMaxEntryVersions)
	if p != "" {
		mv, err := strconv.Atoi(p)
		if err == nil && mv > 0 {
			return mv
		}
	}

	return 10
}
