//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"encoding/json"
	"errors"
	code "github.com/spiffe/spike-sdk-go/api/errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity"
	"github.com/spiffe/spike-sdk-go/api/internal/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// GetSecret retrieves a specific version of a secret at the given path using
// mTLS authentication.
//
// Parameters:
//   - source: X509Source for mTLS client authentication
//   - path: Path to the secret to retrieve
//   - version: Version number of the secret to retrieve
//
// Returns:
//   - *Secret: Secret data if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	secret, err := GetSecret(x509Source, "secret/path", 1)
func GetSecret(source *workloadapi.X509Source,
	path string, version int) (*entity.Secret, error) {
	r := reqres.SecretReadRequest{
		Path:    path,
		Version: version,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New("getSecret: I am having problem generating the payload"),
			err,
		)
	}

	truer := func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)
	if err != nil {
		return nil, err
	}

	body, err := net.Post(client, url.SecretGet(), mr)
	if err != nil {
		if errors.Is(err, code.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.SecretReadResponse
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

	return &entity.Secret{Data: res.Data}, nil
}
