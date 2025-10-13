//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	code "github.com/spiffe/spike-sdk-go/api/errors"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// ListKeys retrieves all secret keys.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//
// Returns:
//   - []string: Array of secret keys if found, empty array if none found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	keys, err := ListKeys(x509Source)
func ListKeys(
	source *workloadapi.X509Source,
) (*[]string, error) {
	if source == nil {
		return nil, code.ErrNilX509Source
	}

	r := reqres.SecretListRequest{}
	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New(
				"listSecretKeys: Having problem generating the payload",
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
