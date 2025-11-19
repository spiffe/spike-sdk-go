//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/net"
)

// Put creates or updates a secret at the specified path with the given
// values.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path where the secret should be stored
//   - values: Map of key-value pairs representing the secret data
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrMarshalFailure: if request serialization fails
//   - ErrPostFailed: if the HTTP request fails
//   - ErrUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error (e.g.,
//     ErrUnauthorized, ErrNotFound, ErrBadRequest, etc.)
//
// Example:
//
//	err := Put(x509Source, "secret/path",
//		map[string]string{"key": "value"})
func Put(
	source *workloadapi.X509Source,
	path string, values map[string]string,
) *sdkErrors.SDKError {
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source
	}

	r := reqres.SecretPutRequest{
		Path:   path,
		Values: values,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return sdkErrors.ErrMarshalFailure.Wrap(err)
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.SecretPut(), mr)
	if err != nil {
		return sdkErrors.ErrPostFailed.Wrap(err)
	}

	res := reqres.SecretPutResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		failErr := sdkErrors.ErrUnmarshalFailure.Wrap(err)
		failErr.Msg = "problem parsing response body"
		return failErr
	}
	if res.Err != "" {
		return sdkErrors.FromCode(res.Err)
	}

	return nil
}
