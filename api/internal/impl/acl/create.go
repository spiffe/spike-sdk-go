//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// CreatePolicy creates a new policy in the system using the provided SPIFFE
// X.509 source and policy details. It establishes a mutual TLS connection to
// SPIKE Nexus using the X.509 source and sends a policy creation request.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - name: The name of the policy to be created
//   - SPIFFEIDPattern: The SPIFFE ID pattern that this policy will apply to
//   - pathPattern: The path pattern that this policy will match against
//   - permissions: A slice of PolicyPermission defining the access rights
//
// Returns:
//   - nil if successful
//   - error if marshaling, network request, or server-side creation fails
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
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
//	}
func CreatePolicy(source *workloadapi.X509Source,
	name string, SPIFFEIDPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) error {
	if source == nil {
		return errors.New("nil X509Source")
	}

	r := reqres.PolicyCreateRequest{
		Name:            name,
		SPIFFEIDPattern: SPIFFEIDPattern,
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

	client := net.CreateMTLSClientForNexus(source)

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
