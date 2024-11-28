package api

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// DeleteSecret deletes specified versions of a secret at the given path using
// mTLS authentication.
//
// It converts string version numbers to integers, constructs a delete request,
// and sends it to the secrets API endpoint. If no versions are specified or
// conversion fails, no versions will be deleted.
//
// Parameters:
//   - source: X509Source for mTLS client authentication
//   - path: Path to the secret to delete
//   - versions: String array of version numbers to delete
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or wrapped
//     error on request/parsing failure
//
// Example:
//
//	err := DeleteSecret(x509Source, "secret/path", []string{"1", "2"})
func DeleteSecret(source *workloadapi.X509Source,
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

	r := reqres.SecretDeleteRequest{
		Path:     path,
		Versions: vv,
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

	var truer = func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)
	if err != nil {
		return err
	}

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
		return errors.New(string(res.Err))
	}

	return err
}
