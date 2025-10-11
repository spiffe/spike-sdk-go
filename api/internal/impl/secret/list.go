//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// ListKeys retrieves all secret keys using mTLS authentication.
//
// Parameters:
//   - source: X509Source for mTLS client authentication
//   - allow: A predicate.Predicate that determines which server certificates
//     to trust during the mTLS connection
//
// Returns:
//   - []string: Array of secret keys if found, empty array if none found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	keys, err := ListKeys(x509Source, predicate.AllowAll)
func ListKeys(
	source *workloadapi.X509Source,
) (*[]string, error) {
	if source == nil {
		return nil, errors.New("nil X509Source")
	}

	r := reqres.SecretListRequest{}
	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New(
				"listSecretKeys: I am having problem generating the payload",
			),
			err,
		)
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.SecretList(), mr)
	if err != nil {
		if errors.Is(err, net.ErrNotFound) {
			return &[]string{}, nil
		}
		return nil, err
	}

	var res reqres.SecretListResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Join(
			errors.New("getSecret: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return nil, errors.New(string(res.Err))
	}

	return &res.Keys, nil
}
