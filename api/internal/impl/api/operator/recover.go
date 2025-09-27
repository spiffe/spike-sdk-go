//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package operator

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/config/env"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/spiffeid"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/net"
	"github.com/spiffe/spike-sdk-go/predicate"
)

// Recover makes a request to initiate recovery of secrets, returning the
// recovery shards.
//
// Parameters:
//   - source: X509Source used for mTLS client authentication
//
// Returns:
//   - map[int]*[32]byte: Map of shard indices to shard byte arrays if
//     successful, nil if not found
//   - error: nil on success, error if:
//   - Failed to marshal recover request
//   - Failed to create mTLS client
//   - Request failed (except for the "not found" case)
//   - Failed to parse response body
//   - Server returned error in response
//
// Example:
//
//	shards, err := Recover(x509Source)
func Recover(source *workloadapi.X509Source) (map[int]*[32]byte, error) {
	svid, err := source.GetX509SVID()
	if err != nil {
		// TODO: FatalLn takes fName as first arg and it is formatted. fix it across API.
		log.FatalLn(err.Error())
	}
	if svid == nil {
		log.FatalLn("no X509SVID in source")
	}
	if svid != nil {
		selfSPIFFEID := svid.ID.String()
		// Security: Recovery and Restoration can ONLY be done via SPIKE Pilot.
		if !spiffeid.IsPilot(env.TrustRoot, selfSPIFFEID) {
			log.FatalLn("spiffeid is not pilot")
		}
	}

	r := reqres.RecoverRequest{}

	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New("recover: failed to marshal recover request"),
			err,
		)
	}

	client, err := net.CreateMTLSClientWithPredicate(
		source, predicate.AllowNexus(env.TrustRootNexus))
	if err != nil {
		return nil, err
	}

	body, err := net.Post(client, url.Recover(), mr)
	if err != nil {
		if errors.Is(err, net.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.RecoverResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Join(
			errors.New("recover: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return nil, errors.New(string(res.Err))
	}

	result := make(map[int]*[32]byte)

	for i, shard := range res.Shards {
		result[i] = shard
	}

	return result, nil
}
