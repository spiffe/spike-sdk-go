//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"
	"errors"

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
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path to the secret to delete
//   - versions: Integer array of version numbers to delete
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or wrapped
//     error on request/parsing failure
//
// Example:
//
//	err := Delete(x509Source, "secret/path", []int{1, 2})
func Delete(
	source *workloadapi.X509Source,
	path string, versions []int,
) error {
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source
	}

	r := reqres.SecretDeleteRequest{
		Path:     path,
		Versions: versions,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New(
				"deleteSecret: I am having problem generating the payload",
			),
			err,
		)
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.SecretDelete(), mr)
	if err != nil {
		return err
	}

	res := reqres.SecretDeleteResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Join(
			errors.New("deleteSecret: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return sdkErrors.FromCode(res.Err)
	}

	return err
}
