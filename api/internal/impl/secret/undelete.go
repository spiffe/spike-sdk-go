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

// Undelete restores previously deleted versions of a secret at the
// specified path.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to restore
//   - versions: Integer array of version numbers to restore. Empty array
//     attempts no restoration
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
//	err := Undelete(x509Source, "secret/path", []int{1, 2})
func Undelete(source *workloadapi.X509Source,
	path string, versions []int,
) *sdkErrors.SDKError {
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source
	}

	var vv []int
	if len(versions) == 0 {
		vv = []int{}
	}

	r := reqres.SecretUndeleteRequest{Path: path, Versions: vv}

	mr, err := json.Marshal(r)
	if err != nil {
		return sdkErrors.ErrMarshalFailure.Wrap(err)
	}

	client := net.CreateMTLSClientForNexus(source)
	body, err := net.Post(client, url.SecretUndelete(), mr)
	if err != nil {
		return err
	}

	res := reqres.SecretUndeleteResponse{}
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
