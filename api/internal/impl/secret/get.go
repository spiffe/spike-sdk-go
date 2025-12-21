//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"context"
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
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to retrieve
//   - version: Version number of the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPINotFound: if the secret is not found
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	secret, err := Get(ctx, x509Source, "secret/path", 1)
func Get(
	ctx context.Context,
	source *workloadapi.X509Source,
	path string, version int,
) (*data.Secret, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	r := reqres.SecretGetRequest{Path: path, Version: version}

	mr, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return nil, failErr
	}

	res, postErr := net.PostAndUnmarshal[reqres.SecretGetResponse](
		ctx, source, url.SecretGet(), mr)
	if postErr != nil {
		return nil, postErr
	}

	return &data.Secret{Data: res.Data}, nil
}
