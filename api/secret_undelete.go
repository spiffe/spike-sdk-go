//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/internal/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// UndeleteSecret restores previously deleted versions of a secret at the
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
//	err := UndeleteSecret(x509Source, "secret/path", []string{"1", "2"})
func UndeleteSecret(source *workloadapi.X509Source,
	path string, versions []string) error {
	var vv []int
	if len(versions) == 0 {
		vv = []int{}
	}

	for _, version := range versions {
		v, e := strconv.Atoi(version)
		if e != nil {
			continue
		}
		vv = append(vv, v)
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

	var truer = func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)
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
