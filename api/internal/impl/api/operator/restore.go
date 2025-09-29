//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package operator

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/net"
	"github.com/spiffe/spike-sdk-go/predicate"
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
//     and restoration state if successful, nil if not found
//   - error: nil on success, error if:
//   - Failed to marshal restore request
//   - Failed to create mTLS client
//   - Request failed (except for not found case)
//   - Failed to parse response body
//   - Server returned error in response
//
// Example:
//
//	status, err := Restore(x509Source, shardIndex, shardValue)
func Restore(
	source *workloadapi.X509Source, shardIndex int, shardValue *[32]byte,
) (*data.RestorationStatus, error) {
	const fName = "restore"

	r := reqres.RestoreRequest{
		ID:    shardIndex,
		Shard: shardValue,
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.FatalLn(fName, "message", "Problem acquiring SVID", "err", err.Error())
	}
	if svid == nil {
		log.FatalLn("no X509SVID in source")
	}
	if svid != nil {
		selfSPIFFEID := svid.ID.String()
		// Security: Recovery and Restoration can ONLY be done via SPIKE Pilot.
		if !spiffeid.IsPilot(selfSPIFFEID) {
			log.FatalLn(fName,
				"message",
				"You can restore only from SPIKE Pilot: spiffeid is not SPIKE Pilot",
			)
		}
	}

	mr, err := json.Marshal(r)
	// Security: Zero out r.Shard as soon as we're done with it
	for i := range r.Shard {
		r.Shard[i] = 0
	}

	if err != nil {
		return nil, errors.Join(
			errors.New("restore: failed to marshal recover request"),
			err,
		)
	}

	client, err := net.CreateMTLSClientWithPredicate(
		source, predicate.AllowNexus,
	)
	if err != nil {
		// Security: Zero out mr before returning error
		for i := range mr {
			mr[i] = 0
		}
		return nil, err
	}

	body, err := net.Post(client, url.Restore(), mr)
	// Security: Zero out mr after the POST request is complete
	for i := range mr {
		mr[i] = 0
	}

	if err != nil {
		if errors.Is(err, net.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.RestoreResponse
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

	return &data.RestorationStatus{
		ShardsCollected: res.ShardsCollected,
		ShardsRemaining: res.ShardsRemaining,
		Restored:        res.Restored,
	}, nil
}
