//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"time"
)

// BootstrapConfigMapNameVal returns the name of the ConfigMap used to store
// SPIKE Bootstrap state information.
//
// It retrieves the ConfigMap name from the SPIKE_BOOTSTRAP_CONFIGMAP_NAME
// environment variable. If the environment variable is not set, it returns
// the default value "spike-bootstrap-state".
//
// Returns:
//   - A string containing the ConfigMap name for storing bootstrap state
func BootstrapConfigMapNameVal() string {
	cn := os.Getenv(BootstrapConfigMapName)
	if cn == "" {
		return "spike-bootstrap-state"
	}
	return cn
}

// BootstrapInitVerificationTimeoutVal returns the timeout duration for
// bootstrap initialization verification.
//
// It retrieves the timeout from the SPIKE_BOOTSTRAP_INIT_VERIFICATION_TIMEOUT
// environment variable. The value should be a valid Go duration string
// (e.g., "30m", "1h", "45m30s"). If the environment variable is not set
// or contains an invalid duration format, it returns a default of 30 minutes.
//
// Returns:
//   - A time.Duration representing the bootstrap initialization verification
//     timeout
func BootstrapInitVerificationTimeoutVal() time.Duration {
	p := os.Getenv(BootstrapInitVerificationTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}
	return 30 * time.Minute
}

// BootstrapTimeoutVal returns the maximum duration for the entire bootstrap
// process.
//
// It retrieves the timeout from the SPIKE_BOOTSTRAP_TIMEOUT environment
// variable. The value should be a valid Go duration string (e.g., "24h",
// "48h", "72h"). A value of 0 means no timeout (infinite), which is also
// the default behavior if the environment variable is not set or contains
// an invalid duration format.
//
// Returns:
//   - A time.Duration representing the maximum bootstrap process duration.
//     A value of 0 indicates no timeout.
func BootstrapTimeoutVal() time.Duration {
	p := os.Getenv(BootstrapTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}
	return 0
}
