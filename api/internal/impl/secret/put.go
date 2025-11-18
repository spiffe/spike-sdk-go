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

// Put creates or updates a secret at the specified path with the given
// values.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - path: Path where the secret should be stored
//   - values: Map of key-value pairs representing the secret data
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	err := Put(x509Source, "secret/path",
//		map[string]string{"key": "value"})
func Put(
	source *workloadapi.X509Source,
	path string, values map[string]string,
) error {
	if source == nil {
		return sdkErrors.ErrNilX509Source
	}

	r := reqres.SecretPutRequest{
		Path:   path,
		Values: values,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New("putSecret: I am having problem generating the payload"),
			err,
		)
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.SecretPut(), mr)
	if err != nil {
		return err
	}

	res := reqres.SecretPutResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Join(
			errors.New("putSecret: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return errors.New(string(res.Err))
	}

	return nil
}
