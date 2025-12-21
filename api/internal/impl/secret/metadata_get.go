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

// GetMetadata retrieves a specific version of a secret metadata at the
// given path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to retrieve
//   - version: Version number of the secret to retrieve
//
// Returns:
//   - *data.SecretMetadata: Secret metadata if found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPINotFound: if the secret metadata is not found
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	metadata, err := GetMetadata(ctx, x509Source, "secret/path", 1)
func GetMetadata(
	ctx context.Context,
	source *workloadapi.X509Source, path string, version int,
) (*data.SecretMetadata, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	r := reqres.SecretMetadataRequest{Path: path, Version: version}

	mr, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return nil, failErr
	}

	res, postErr := net.PostAndUnmarshal[reqres.SecretMetadataResponse](
		ctx, source, url.SecretMetadataGet(), mr)
	if postErr != nil {
		return nil, postErr
	}

	return &data.SecretMetadata{
		Versions: res.Versions,
		Metadata: res.Metadata,
	}, nil
}
