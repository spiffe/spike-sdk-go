//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package operator

import (
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/net"
	"github.com/spiffe/spike-sdk-go/spiffeid"
)

// Recover makes a request to initiate recovery of secrets, returning the
// recovery shards.
//
// Parameters:
//   - source: X509Source used for mTLS client authentication
//
// Returns:
//   - map[int]*[32]byte: Map of shard indices to shard byte arrays if successful
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
//	shards, err := Recover(x509Source)
func Recover(source *workloadapi.X509Source) (map[int]*[32]byte, *sdkErrors.SDKError) {
	const fName = "recover"

	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source
	}

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
		failErr.Msg = "recovery can only be performed from SPIKE Pilot"
		log.FatalErr(fName, *failErr)
	}

	r := reqres.RecoverRequest{}

	mr, err := json.Marshal(r)
	if err != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(err)
		failErr.Msg = "failed to marshal recover request"
		return nil, failErr
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.Recover(), mr)
	if err != nil {
		return nil, err
	}

	var res reqres.RecoverResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(err)
		failErr.Msg = "problem parsing response body"
		return nil, failErr
	}
	if res.Err != "" {
		return nil, sdkErrors.FromCode(res.Err)
	}

	result := make(map[int]*[32]byte)
	for i, shard := range res.Shards {
		result[i] = shard
	}
	return result, nil
}
