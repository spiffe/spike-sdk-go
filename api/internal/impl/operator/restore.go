//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package operator

import (
	"context"
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/net"
	"github.com/spiffe/spike-sdk-go/spiffeid"
)

// Restore submits a recovery shard to continue the restoration process.
//
// SVID Acquisition Error Handling:
//
// This function attempts to acquire an X.509 SVID from the SPIFFE Workload API
// via Unix domain socket. While UDS connections are generally more reliable than
// network sockets, SVID acquisition can fail in both fatal and transient ways:
//
// Fatal failures (indicate misconfiguration):
//   - Socket file doesn't exist (SPIRE agent never started)
//   - Permission denied (deployment/configuration error)
//   - Wrong socket path (configuration error)
//
// Transient failures (may succeed on retry):
//   - SPIRE agent restarting (brief unavailability, recovers in seconds)
//   - SVID not yet provisioned (startup race condition after attestation)
//   - File descriptor exhaustion (resource pressure may clear)
//   - SVID rotation failure (temporary SPIRE server issue)
//   - Workload API connection lost after source creation (agent crash/restart)
//
// Since restoration is often performed during emergency procedures when
// infrastructure may be unstable, this function returns errors rather than
// crashing to allow retry logic. Callers can implement exponential backoff
// or other retry strategies for transient failures.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source *workloadapi.X509Source: X509Source used for mTLS client
//     authentication
//   - shardIndex int: Index of the recovery shard
//   - shardValue *[32]byte: Pointer to a 32-byte array containing the recovery
//     shard
//
// Returns:
//   - *data.RestorationStatus: Status containing shards collected, remaining,
//     and restoration state if successful
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrSPIFFEFailedToExtractX509SVID: if SVID acquisition fails (may be
//     transient - see above for retry guidance)
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Security Note: The function will fatally crash (via log.FatalErr) if the
// caller is not SPIKE Pilot. This is a programming error, not a runtime
// condition, as restoration operations must only be performed by Pilot roles.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	status, err := Restore(ctx, x509Source, shardIndex, shardValue)
//	if err != nil {
//	    // SVID acquisition failures may be transient - consider retry logic
//	    return nil, err
//	}
func Restore(
	ctx context.Context,
	source *workloadapi.X509Source, shardIndex int, shardValue *[32]byte,
) (*data.RestorationStatus, *sdkErrors.SDKError) {
	const fName = "restore"

	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	r := reqres.RestoreRequest{ID: shardIndex, Shard: shardValue}

	svid, err := source.GetX509SVID()
	if err != nil {
		failErr := sdkErrors.ErrSPIFFEFailedToExtractX509SVID.Wrap(err)
		failErr.Msg = "could not acquire SVID"
		return nil, failErr
	}
	if svid == nil {
		failErr := sdkErrors.ErrSPIFFEFailedToExtractX509SVID.Clone()
		failErr.Msg = "no X509SVID in source"
		return nil, failErr
	}

	selfSPIFFEID := svid.ID.String()

	// Security: Recovery and Restoration can ONLY be done via SPIKE Pilot.
	if !spiffeid.IsPilotRestore(selfSPIFFEID) {
		failErr := sdkErrors.ErrAccessUnauthorized.Clone()
		failErr.Msg = "restoration can only be performed from SPIKE Pilot"
		log.FatalErr(fName, *failErr)
	}

	mr, marshalErr := json.Marshal(r)
	// Security: Zero out r.Shard as soon as we're done with it
	for i := range r.Shard {
		r.Shard[i] = 0
	}
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "failed to marshal restore request"
		return nil, failErr
	}

	res, postErr := net.PostAndUnmarshal[reqres.RestoreResponse](
		ctx, source, url.Restore(), mr)
	// Security: Zero out mr after the POST request is complete
	for i := range mr {
		mr[i] = 0
	}
	if postErr != nil {
		return nil, postErr
	}

	return &data.RestorationStatus{
		ShardsCollected: res.ShardsCollected,
		ShardsRemaining: res.ShardsRemaining,
		Restored:        res.Restored,
	}, nil
}
