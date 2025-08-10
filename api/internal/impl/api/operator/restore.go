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
	"github.com/spiffe/spike-sdk-go/net"
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
	r := reqres.RestoreRequest{
		ID:    shardIndex,
		Shard: shardValue,
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

	client, err := net.CreateMTLSClient(source)
	if err != nil {
		// Security: Zero out mr before returning error
		for i := range mr {
			mr[i] = 0
		}
		return nil, err
	}

	body, err := net.Post(client, url.Restore(), mr)
	// Security: Zero out mr after the post request is complete
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
