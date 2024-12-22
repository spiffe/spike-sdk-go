//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// CreatePolicy creates a new policy in the system using the provided SPIFFE
// X.509 source and policy details. It establishes a mutual TLS connection using
// the X.509 source and sends a policy creation request to the server.
//
// The function takes the following parameters:
//   - source: A pointer to a workloadapi.X509Source for establishing mTLS
//     connection
//   - name: The name of the policy to be created
//   - spiffeIdPattern: The SPIFFE ID pattern that this policy will apply to
//   - pathPattern: The path pattern that this policy will match against
//   - permissions: A slice of PolicyPermission defining the access rights for
//     this policy
//
// The function returns an error if any of the following operations fail:
//   - Marshaling the policy creation request
//   - Creating the mTLS client
//   - Making the HTTP POST request
//   - Unmarshaling the response
//   - Server-side policy creation (indicated in the response)
//
// Example usage:
//
//	source, err := workloadapi.NewX509Source(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	permissions := []data.PolicyPermission{
//	    {
//	        Action: "read",
//	        Resource: "documents/*",
//	    },
//	}
//
//	err = CreatePolicy(
//	    source,
//	    "doc-reader",
//	    "spiffe://example.org/service/*",
//	    "/api/documents/*",
//	    permissions,
//	)
//	if err != nil {
//	    log.Printf("Failed to create policy: %v", err)
//	    return
//	}
func CreatePolicy(source *workloadapi.X509Source,
	name string, spiffeIdPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) error {
	r := reqres.PolicyCreateRequest{
		Name:            name,
		SpiffeIdPattern: spiffeIdPattern,
		PathPattern:     pathPattern,
		Permissions:     permissions,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New(
				"createPolicy: I am having problem generating the payload",
			),
			err,
		)
	}

	client, err := net.CreateMtlsClient(source)
	if err != nil {
		return err
	}

	body, err := net.Post(client, url.PolicyCreate(), mr)
	if err != nil {
		return err
	}

	res := reqres.PolicyCreateResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Join(
			errors.New("createPolicy: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return errors.New(string(res.Err))
	}

	return nil
}
