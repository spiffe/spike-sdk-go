//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/spiffe/spike-sdk-go/log"
)

// KeepersVal retrieves and parses the keeper peer configurations from the
// environment. It reads SPIKE_NEXUS_KEEPER_PEERS environment variable which
// should contain a comma-separated list of keeper URLs.
//
// The environment variable should be formatted as:
// 'https://localhost:8443,https://localhost:8543,https://localhost:8643'
//
// The SPIKE Keeper address mappings will be automatically assigned starting
// with the key "1" and incrementing by 1 for each subsequent SPIKE Keeper.
//
// Returns:
//   - map[string]string: Mapping of keeper IDs to their URLs
//
// Panics if:
//   - SPIKE_NEXUS_KEEPER_PEERS is not set
func KeepersVal() map[string]string {
	const fName = "KeepersVal"

	p := os.Getenv(NexusKeeperPeers)

	if p == "" {
		log.FatalLn(
			fName,
			"message",
			"SPIKE_NEXUS_KEEPER_PEERS must be configured in the environment",
		)
	}

	urls := strings.Split(p, ",")

	// Check for duplicate and empty URLs
	urlMap := make(map[string]bool)
	for i, u := range urls {
		trimmedURL := strings.TrimSpace(u)
		if trimmedURL == "" {
			log.FatalLn(
				fName,
				"message", "empty url found",
				"position", i+1,
			)
		}

		// Validate URL format and security
		if !validURL(trimmedURL) {
			log.FatalLn(
				fName,
				"message", "invalid url format",
				"position", i+1,
			)
		}

		if urlMap[trimmedURL] {
			log.FatalLn(
				fName,
				"message", "duplicate url found",
				"position", i+1,
			)
		}

		urlMap[trimmedURL] = true
	}

	// The key of the map is the Shamir Shard index (starting from 1), and
	// the value is the Keeper URL that corresponds to that shard index.
	peers := make(map[string]string)
	for i, u := range urls {
		peers[strconv.Itoa(i+1)] = strings.TrimSpace(u)
	}

	return peers
}
