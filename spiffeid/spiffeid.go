//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import "github.com/spiffe/spike-sdk-go/spiffeid/internal/env"

// SpikeKeeper constructs and returns the SPIKE Keeper's
// SPIFFE ID string.
// The output is constructed based on the trust root from the environment.
func SpikeKeeper() string {
	return "spiffe://" + env.TrustRoot() + "/spike/keeper"
}

// SpikeNexus constructs and returns the SPIFFE ID for SPIKE Nexus.
// The output is constructed based on the trust root from the environment.
func SpikeNexus() string {
	return "spiffe://" + env.TrustRoot() + "/spike/nexus"
}

// SpikePilot generates the SPIFFE ID for a SPIKE Pilot superuser role.
// The output is constructed based on the trust root from the environment.
func SpikePilot() string {
	return "spiffe://" + env.TrustRoot() + "/spike/pilot/role/superuser"
}

// SpikePilotRecover generates the SPIFFE ID for a SPIKE Pilot recovery role.
// The output is constructed based on the trust root from the environment.
func SpikePilotRecover() string {
	return "spiffe://" + env.TrustRoot() + "/spike/pilot/role/recover"
}

// SpikePilotRestore generates the SPIFFE ID for a SPIKE Pilot restore role.
// The output is constructed based on the trust root from the environment.
func SpikePilotRestore() string {
	return "spiffe://" + env.TrustRoot() + "/spike/pilot/role/restore"
}
