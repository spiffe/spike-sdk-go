//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import "github.com/spiffe/spike-sdk-go/spiffeid/internal/env"

// SpikeKeeperSpiffeId constructs and returns the SPIKE Keeper's
// SPIFFE ID string.
// The output is constructed based on the trust root from the environment.
func SpikeKeeperSpiffeId() string {
	return "spiffe://" + env.TrustRoot() + "/spike/keeper"
}

// SpikeNexusSpiffeId constructs and returns the SPIFFE ID for SPIKE Nexus.
// The output is constructed based on the trust root from the environment.
func SpikeNexusSpiffeId() string {
	return "spiffe://" + env.TrustRoot() + "/spike/nexus"
}

// SpikePilotSpiffeId generates the SPIFFE ID for a SPIKE Pilot superuser role.
// The output is constructed based on the trust root from the environment.
func SpikePilotSpiffeId() string {
	return "spiffe://" + env.TrustRoot() + "/spike/pilot/role/superuser"
}

func SpikePilotRecoverSpiffeId() string {
	return "spiffe://" + env.TrustRoot() + "/spike/pilot/role/recover"
}

func SpikePilotRestoreSpiffeId() string {
	return "spiffe://" + env.TrustRoot() + "/spike/pilot/role/restore"
}
