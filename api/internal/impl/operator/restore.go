//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package operator

import (
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
// Parameters:
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
//	status, err := Restore(x509Source, shardIndex, shardValue)
func Restore(
	source *workloadapi.X509Source, shardIndex int, shardValue *[32]byte,
) (*data.RestorationStatus, *sdkErrors.SDKError) {
	const fName = "restore"

	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source
	}

	r := reqres.RestoreRequest{ID: shardIndex, Shard: shardValue}

	svid, err := source.GetX509SVID()
	if err != nil {
		failErr := sdkErrors.ErrSPIFFEFailedToExtractX509SVID.Wrap(err)
		failErr.Msg = "could not acquire SVID"
		log.FatalErr(fName, *failErr)
		return nil, failErr // To make linter happy.
	}
	if svid == nil {
		failErr := sdkErrors.ErrSPIFFEFailedToExtractX509SVID
		failErr.Msg = "no X509SVID in source"
		log.FatalErr(fName, *failErr)
		return nil, failErr // To make linter happy.
	}

	selfSPIFFEID := svid.ID.String()

	// Security: Recovery and Restoration can ONLY be done via SPIKE Pilot.
	if !spiffeid.IsPilot(selfSPIFFEID) {
		failErr := sdkErrors.ErrAccessUnauthorized
		failErr.Msg = "restoration can only be performed from SPIKE Pilot"
		log.FatalErr(fName, *failErr)
		return nil, failErr // To make linter happy.
	}

	mr, err := json.Marshal(r)
	// Security: Zero out r.Shard as soon as we're done with it
	for i := range r.Shard {
		r.Shard[i] = 0
	}

	if err != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(err)
		failErr.Msg = "failed to marshal restore request"
		return nil, failErr
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.Restore(), mr)
	// Security: Zero out mr after the POST request is complete
	for i := range mr {
		mr[i] = 0
	}

	if err != nil {
		return nil, err
	}

	var res reqres.RestoreResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(err)
		failErr.Msg = "problem parsing response body"
		return nil, failErr
	}
	if res.Err != "" {
		return nil, sdkErrors.FromCode(res.Err)
	}

	return &data.RestorationStatus{
		ShardsCollected: res.ShardsCollected,
		ShardsRemaining: res.ShardsRemaining,
		Restored:        res.Restored,
	}, nil
}
