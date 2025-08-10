//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/spike-sdk-go/api/url"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/net"
)

// Undelete restores previously deleted versions of a secret at the
// specified path using mTLS authentication.
//
// Parameters:
//   - source: X509Source for mTLS client authentication
//   - path: Path to the secret to restore
//   - versions: String array of version numbers to restore. Empty array
//     attempts no restoration
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	err := undeleteSecret(x509Source, "secret/path", []string{"1", "2"})
func Undelete(source *workloadapi.X509Source,
	path string, versions []int) error {
	var vv []int
	if len(versions) == 0 {
		vv = []int{}
	}

	r := reqres.SecretUndeleteRequest{
		Path:     path,
		Versions: vv,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New(
				"undeleteSecret: I am having problem generating the payload",
			),
			err,
		)
	}

	client, err := net.CreateMtlsClient(source)
	if err != nil {
		return err
	}

	body, err := net.Post(client, url.SecretUndelete(), mr)
	if err != nil {
		return nil
	}

	res := reqres.SecretUndeleteResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Join(
			errors.New("undeleteSecret: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return errors.New(string(res.Err))
	}

	return nil
}
