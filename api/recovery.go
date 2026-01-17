//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/operator"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Recover returns recovery partitions for SPIKE Nexus to be used in a
// break-the-glass recovery operation.
//
// This should be used when the SPIKE Nexus auto-recovery mechanism isn't
// successful. The returned shards are sensitive and should be securely stored
// out-of-band in encrypted form.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//
// Returns:
//   - map[int]*[32]byte: Map of shard indices to shard byte arrays if
//     successful, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: The function will fatally crash (via log.FatalErr) if:
//   - SVID acquisition fails
//   - SVID is nil
//   - Caller is not SPIKE Pilot (security requirement)
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	shards, err := api.Recover(ctx)
//	if err != nil {
//	    log.Fatalf("Failed to recover shards: %v", err)
//	}
func (a *API) Recover(ctx context.Context) (map[int]*[32]byte, *sdkErrors.SDKError) {
	return operator.Recover(ctx, a.source)
}

// Restore submits a recovery shard to continue the SPIKE Nexus restoration
// process.
//
// This is used when SPIKE Keepers cannot provide adequate shards and SPIKE
// Nexus cannot recall its root key. This is a break-the-glass superuser-only
// operation that a well-architected SPIKE deployment should not need.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - index: Index of the recovery shard
//   - shard: Pointer to a 32-byte array containing the recovery shard
//
// Returns:
//   - *data.RestorationStatus: Status containing shards collected, remaining,
//     and restoration state if successful, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: The function will fatally crash (via log.FatalErr) if:
//   - SVID acquisition fails
//   - SVID is nil
//   - Caller is not SPIKE Pilot (security requirement)
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	status, err := api.Restore(ctx, shardIndex, shardPtr)
//	if err != nil {
//	    log.Fatalf("Failed to restore shard: %v", err)
//	}
//	log.Printf("Shards collected: %d, remaining: %d",
//	    status.ShardsCollected, status.ShardsRemaining)
func (a *API) Restore(
	ctx context.Context, index int, shard *[32]byte,
) (*data.RestorationStatus, *sdkErrors.SDKError) {
	return operator.Restore(ctx, a.source, index, shard)
}
