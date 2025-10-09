//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strconv"
)

// ShamirSharesVal returns the total number of shares to be used in Shamir's
// Secret Sharing. It reads the value from the SPIKE_NEXUS_SHAMIR_SHARES
// environment variable.
//
// Returns:
//   - The number of shares specified in the environment variable if it's a
//     valid positive integer
//   - The default value of 3 if the environment variable is unset, empty,
//     or invalid
//
// This determines the total number of shares that will be created when
//
//	splitting a secret.
func ShamirSharesVal() int {
	p := os.Getenv(NexusShamirShares)
	if p != "" {
		mv, err := strconv.Atoi(p)
		if err == nil && mv > 0 {
			return mv
		}
	}

	return 3
}

// ShamirThresholdVal returns the minimum number of shares required to
// reconstruct the secret in Shamir's Secret Sharing scheme.
// It reads the value from the SPIKE_NEXUS_SHAMIR_THRESHOLD environment
// variable.
//
// Returns:
//   - The threshold specified in the environment variable if it's a valid
//     positive integer
//   - The default value of 2 if the environment variable is unset, empty,
//     or invalid
//
// This threshold value determines how many shares are needed to recover the
// original secret. It should be less than or equal to the total number of
// shares (ShamirShares()).
func ShamirThresholdVal() int {
	p := os.Getenv(NexusShamirThreshold)
	if p != "" {
		mv, err := strconv.Atoi(p)
		if err == nil && mv > 0 {
			return mv
		}
	}

	return 2
}
