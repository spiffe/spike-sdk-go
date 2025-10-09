//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "os"

// NexusTLSPortVal returns the TLS port for the Spike Nexus service.
// It reads from the SPIKE_NEXUS_TLS_PORT environment variable.
// If the environment variable is not set, it returns the default port ":8553".
func NexusTLSPortVal() string {
	p := os.Getenv(NexusTLSPort)
	if p != "" {
		return p
	}

	return ":8553"
}

// KeeperTLSPortVal returns the TLS port for the Spike Keeper service.
// It first checks for a port specified in the SPIKE_KEEPER_TLS_PORT
// environment variable.
// If no environment variable is set, it defaults to ":8443".
//
// The returned string is in the format ":port" suitable for use with
// net/http Listen functions.
func KeeperTLSPortVal() string {
	p := os.Getenv(KeeperTLSPort)

	if p != "" {
		return p
	}

	return ":8443"
}
