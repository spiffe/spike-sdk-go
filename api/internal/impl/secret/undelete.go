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
	"github.com/spiffe/spike-sdk-go/net"
)

// Undelete restores previously deleted versions of a secret at the
// specified path using mTLS authentication.
//
// Parameters:
//   - source: X509Source for mTLS client authentication
//   - path: Path to the secret to restore
//   - versions: Integer array of version numbers to restore. Empty array
//     attempts no restoration
//   - allow: A predicate.Predicate that determines which server certificates
//     to trust during the mTLS connection
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	err := Undelete(x509Source, "secret/path", []int{1, 2}, predicate.AllowAll)
func Undelete(source *workloadapi.X509Source,
	path string, versions []int,
) error {
	if source == nil {
		return errors.New("nil X509Source")
	}

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

	client := net.CreateMTLSClientForNexus(source)
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
