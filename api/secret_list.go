//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/internal/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// ListSecretKeys retrieves all secret keys using mTLS authentication.
//
// Parameters:
//   - source: X509Source for mTLS client authentication
//
// Returns:
//   - []string: Array of secret keys if found, empty array if none found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	keys, err := ListSecretKeys(x509Source)
func ListSecretKeys(source *workloadapi.X509Source) ([]string, error) {
	r := reqres.SecretListRequest{}
	mr, err := json.Marshal(r)
	if err != nil {
		return []string{}, errors.Join(
			errors.New(
				"listSecretKeys: I am having problem generating the payload",
			),
			err,
		)
	}

	var truer = func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)
	if err != nil {
		return []string{}, err
	}

	body, err := net.Post(client, url.SecretList(), mr)
	if err != nil {
		if errors.Is(err, net.ErrNotFound) {
			return []string{}, nil
		}
		return []string{}, err
	}

	var res reqres.SecretListResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return []string{}, errors.Join(
			errors.New("getSecret: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return []string{}, errors.New(string(res.Err))
	}

	return res.Keys, nil
}