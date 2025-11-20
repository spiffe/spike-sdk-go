//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/net"
)

// Get retrieves a specific version of a secret at the given path.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to retrieve
//   - version: Version number of the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil if secret not found
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrNotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (nil, nil) if the secret is not found (ErrNotFound)
//
// Example:
//
//	secret, err := Get(x509Source, "secret/path", 1)
func Get(
	source *workloadapi.X509Source,
	path string, version int,
) (*data.Secret, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source
	}

	r := reqres.SecretReadRequest{Path: path, Version: version}

	mr, err := json.Marshal(r)
	if err != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(err)
		failErr.Msg = "problem generating the payload"
		return nil, failErr
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.SecretGet(), mr)
	if err != nil {
		if err.Is(sdkErrors.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.SecretReadResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(err)
		failErr.Msg = "problem parsing response body"
		return nil, failErr
	}
	if res.Err != "" {
		return nil, sdkErrors.FromCode(res.Err)
	}

	return &data.Secret{Data: res.Data}, nil
}
