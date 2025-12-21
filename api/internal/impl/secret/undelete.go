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

// Undelete restores previously deleted versions of a secret at the
// specified path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to restore
//   - versions: Integer array of version numbers to restore. Empty array
//     attempts no restoration
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
//	err := Undelete(ctx, x509Source, "secret/path", []int{1, 2})
func Undelete(
	ctx context.Context,
	source *workloadapi.X509Source,
	path string, versions []int,
) *sdkErrors.SDKError {
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	var vv []int
	if len(versions) == 0 {
		vv = []int{}
	} else {
		vv = versions
	}

	r := reqres.SecretUndeleteRequest{Path: path, Versions: vv}

	mr, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return failErr
	}

	_, postErr := net.PostAndUnmarshal[reqres.SecretUndeleteResponse](
		ctx, source, url.SecretUndelete(), mr)
	return postErr
}
