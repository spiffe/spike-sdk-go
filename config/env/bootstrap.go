//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strconv"
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

// BootstrapKeeperTimeoutVal returns the timeout duration used by the bootstrap
// keeper when attempting to complete its operation.
//
// It retrieves the timeout value from the SPIKE_BOOTSTRAP_KEEPER_TIMEOUT
// environment variable. The value must be a valid Go duration string
// (e.g., "30s", "1m", "5m").
//
// If the environment variable is not set or contains an invalid duration
// format, it returns a default value of 30 seconds.
//
// Returns:
//   - A time.Duration representing the bootstrap keeper timeout.
func BootstrapKeeperTimeoutVal() time.Duration {
	p := os.Getenv(BootstrapKeeperTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}
	return 30 * time.Second
}

// BootstrapKeeperMaxRetriesVal returns the maximum number of retry attempts
// allowed for the bootstrap keeper operation.
//
// It retrieves the retry count from the SPIKE_BOOTSTRAP_KEEPER_MAX_RETRIES
// environment variable. The value must be a positive integer.
//
// If the environment variable is not set, contains an invalid value,
// or is less than or equal to zero, it returns a default value of 5.
//
// Returns:
//   - An integer representing the maximum number of bootstrap keeper retries.
func BootstrapKeeperMaxRetriesVal() int {
	p := os.Getenv(BootstrapKeeperMaxRetries)
	if p != "" {
		mi, err := strconv.Atoi(p)
		if err == nil && mi > 0 {
			return mi
		}
	}
	return 5
}

// BootstrapKeeperRetryInitialIntervalVal returns the initial interval between
// retry attempts when broadcasting to keepers during bootstrap.
//
// It retrieves the interval from the
// SPIKE_BOOTSTRAP_KEEPER_RETRY_INITIAL_INTERVAL environment variable.
// The value must be a valid Go duration string (e.g., "2s", "5s", "1m").
//
// If the environment variable is not set or contains an invalid duration
// format, it returns a default value of 2 seconds.
//
// Returns:
//   - A time.Duration representing the initial retry interval.
func BootstrapKeeperRetryInitialIntervalVal() time.Duration {
	p := os.Getenv(BootstrapKeeperRetryInitialInterval)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}
	return 2 * time.Second
}

// BootstrapKeeperRetryMaxIntervalVal returns the maximum interval between
// retry attempts when broadcasting to keepers during bootstrap.
//
// It retrieves the interval from the SPIKE_BOOTSTRAP_KEEPER_RETRY_MAX_INTERVAL
// environment variable. The value must be a valid Go duration string
// (e.g., "30s", "1m", "5m").
//
// If the environment variable is not set or contains an invalid duration
// format, it returns a default value of 30 seconds.
//
// Returns:
//   - A time.Duration representing the maximum retry interval.
func BootstrapKeeperRetryMaxIntervalVal() time.Duration {
	p := os.Getenv(BootstrapKeeperRetryMaxInterval)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}
	return 30 * time.Second
}
