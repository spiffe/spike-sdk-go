//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/spike-sdk-go/api/url"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/net"
)

// ListPolicies retrieves policies from the system, optionally filtering by
// SPIFFE ID and path patterns. It requires a SPIFFE X.509 source for
// establishing a mutual TLS connection to make the list request.
//
// The function takes:
//   - source: A pointer to a workloadapi.X509Source for establishing mTLS
//     connection
//   - spiffeIdPattern: The SPIFFE ID pattern to filter policies. An empty
//     string matches all SPIFFE IDs.
//   - pathPattern: The path pattern to filter policies. An empty string
//     matches all paths.
//
// The function returns:
//   - (*[]data.Policy, nil) containing all matching policies if successful
//   - (nil, nil) if no policies are found
//   - (nil, error) if an error occurs during the operation
//
// Note: The returned slice pointer should be dereferenced before use:
//
//	policies := *result
//
// Errors can occur during:
//   - Marshaling the policy list request
//   - Creating the mTLS client
//   - Making the HTTP POST request (except for not found cases)
//   - Unmarshaling the response
//   - Server-side policy listing (indicated in the response)
//
// Example usage:
//
//	source, err := workloadapi.NewX509Source(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	// List all policies
//	result, err := ListPolicies(source, "", "")
//	if err != nil {
//	    log.Printf("Error listing policies: %v", err)
//	    return
//	}
//	if result == nil {
//	    log.Printf("No policies found")
//	    return
//	}
//
//	policies := *result
//	for _, policy := range policies {
//	    log.Printf("Found policy: %+v", policy)
//	}
func ListPolicies(
	source *workloadapi.X509Source,
	SPIFFEIDPattern string, pathPattern string,
) (*[]data.Policy, error) {
	r := reqres.PolicyListRequest{
		SPIFFEIDPattern: SPIFFEIDPattern,
		PathPattern:     pathPattern,
	}
	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New(
				"listPolicies: I am having problem generating the payload",
			),
			err,
		)
	}

	client, err := net.CreateMtlsClient(source)
	if err != nil {
		return nil, err
	}

	body, err := net.Post(client, url.PolicyList(), mr)
	if err != nil {
		if errors.Is(err, net.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.PolicyListResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Join(
			errors.New("listPolicies: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return nil, errors.New(string(res.Err))
	}

	return &res.Policies, nil
}
