//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/operator"
)

// Recover returns recovery partitions for SPIKE Nexus to be used in a
// break-the-glass recovery operation if the SPIKE Nexus auto-recovery mechanism
// isn't successful.
//
// The returned shards are sensitive and should be securely stored out-of-band
// in encrypted form.
//
// Returns:
//   - *[][32]byte: Pointer to an array of recovery shards as 32-byte arrays
//   - error: nil on success, unauthorized error if not authorized, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	shards, err := api.Recover()
func (a *API) Recover() (map[int]*[32]byte, error) {
	return operator.Recover(a.source)
}

// Restore SPIKE Nexus backing using recovery shards when SPIKE Keepers cannot
// provide adequate shards and SPIKE Nexus cannot recall its root key either.
//
// This is a break-the-glass superuser-only operation that a well-architected
// SPIKE deployment should not need.
//
// Parameters:
//   - shard *[32]byte: Pointer to a 32-byte array containing the shard to seed
//
// Returns:
//   - *data.RestorationStatus: Status of the restoration process if successful
//   - error: nil on success, unauthorized error if not authorized, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	status, err := api.Restore(shardPtr)
func (a *API) Restore(
	index int, shard *[32]byte,
) (*data.RestorationStatus, error) {
	return operator.Restore(a.source, index, shard)
}
