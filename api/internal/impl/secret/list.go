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

// ListKeys retrieves all secret keys.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//
// Returns:
//   - *[]string: Array of secret keys if found, empty array if no secrets exist
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrNotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (*[]string{}, nil) if no secrets are found (ErrNotFound)
//
// Example:
//
//	keys, err := ListKeys(x509Source)
func ListKeys(
	source *workloadapi.X509Source,
) (*[]string, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source
	}

	r := reqres.SecretListRequest{}
	mr, err := json.Marshal(r)
	if err != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(err)
		failErr.Msg = "problem generating the payload"
		return nil, failErr
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.SecretList(), mr)
	if err != nil {
		if err.Is(sdkErrors.ErrNotFound) {
			return &[]string{}, nil
		}
		return nil, err
	}

	var res reqres.SecretListResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(err)
		failErr.Msg = "problem parsing response body"
		return nil, failErr
	}
	if res.Err != "" {
		return nil, sdkErrors.FromCode(res.Err)
	}

	return &res.Keys, nil
}
