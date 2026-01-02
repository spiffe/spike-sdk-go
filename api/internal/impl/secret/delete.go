//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"context"
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/net"
)

// Delete deletes specified versions of a secret at the given path.
//
// It converts string version numbers to integers, constructs a delete request,
// and sends it to the secrets API endpoint. If no versions are specified or
// the conversion fails, no versions will be deleted.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to delete
//   - versions: Integer array of version numbers to delete
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	err := Delete(ctx, x509Source, "secret/path", []int{1, 2})
func Delete(
	ctx context.Context,
	source *workloadapi.X509Source,
	path string, versions []int,
) *sdkErrors.SDKError {
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	r := reqres.SecretDeleteRequest{Path: path, Versions: versions}

	mr, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return failErr
	}

	_, postErr := net.PostAndUnmarshal[reqres.SecretDeleteResponse](
		ctx, source, url.SecretDelete(), mr)
	return postErr
}
