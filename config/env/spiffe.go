//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"time"
)

// SPIFFESourceTimeoutVal returns the timeout duration for creating a SPIFFE
// X509Source and fetching the initial SVID from the SPIFFE Workload API.
// It can be configured using the SPIKE_SPIFFE_SOURCE_TIMEOUT environment
// variable. The value should be a valid Go duration string (e.g., "30s", "1m").
//
// This timeout prevents indefinite blocking if there are issues with the
// SPIFFE Workload API socket (e.g., agent not running, socket permissions,
// network issues).
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 30 seconds.
func SPIFFESourceTimeoutVal() time.Duration {
	p := os.Getenv(SPIFFESourceTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 30 * time.Second
}
