//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/net"
)

// GetMetadata retrieves a specific version of a secret metadata at the
// given path.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to retrieve
//   - version: Version number of the secret to retrieve
//
// Returns:
//   - *Secret: Secret metadata if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	metadata, err := GetMetadata(x509Source, "secret/path", 1)
func GetMetadata(
	source *workloadapi.X509Source, path string, version int,
) (*data.SecretMetadata, error) {
	if source == nil {
		return nil, sdkErrors.ErrNilX509Source
	}

	r := reqres.SecretMetadataRequest{
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

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.SecretMetadataGet(), mr)
	if err != nil {
		if errors.Is(err, sdkErrors.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.SecretMetadataResponse
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

	return &data.SecretMetadata{
		Versions: res.Versions,
		Metadata: res.Metadata,
	}, nil
}
