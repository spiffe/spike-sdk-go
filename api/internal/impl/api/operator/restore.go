//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package operator

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// Restore submits a recovery shard to continue the restoration process.
//
// Parameters:
//   - source: X509Source used for mTLS client authentication
//   - shard: Recovery shard identifier to submit
//
// Returns:
//   - *RestorationStatus: Status containing shards collected, remaining, and
//     restoration state if successful, nil if not found
//   - error: nil on success, error if:
//   - Failed to marshal restore request
//   - Failed to create mTLS client
//   - Request failed (except for not found case)
//   - Failed to parse response body
//   - Server returned error in response
//
// Example:
//
//	status, err := Restore(x509Source, "randomshardentry")
func Restore(
	source *workloadapi.X509Source, shard string,
) (*data.RestorationStatus, error) {
	r := reqres.RestoreRequest{Shard: shard}

	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New("restore: failed to marshal recover request"),
			err,
		)
	}

	client, err := net.CreateMtlsClient(source)
	if err != nil {
		return nil, err
	}

	body, err := net.Post(client, url.Restore(), mr)
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
