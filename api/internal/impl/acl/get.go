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
	"github.com/spiffe/spike-sdk-go/predicate"
)

// GetPolicy retrieves a policy from the system using its ID.
// It requires a SPIFFE X.509 source for establishing a mutual TLS connection
// to make the retrieval request.
//
// The function takes the following parameters:
//   - source: A pointer to a workloadapi.X509Source for establishing mTLS
//     connection
//   - id: The unique identifier of the policy to retrieve
//   - allow: A predicate.Predicate that determines which server certificates
//     to trust during the mTLS connection
//
// The function returns:
//   - (*data.Policy, nil) if the policy is found
//   - (nil, nil) if the policy is not found
//   - (nil, error) if an error occurs during the operation
//
// Errors can occur during:
//   - Marshaling the policy retrieval request
//   - Creating the mTLS client
//   - Making the HTTP POST request (except for not found cases)
//   - Unmarshaling the response
//   - Server-side policy retrieval (indicated in the response)
//
// Example usage:
//
//	source, err := workloadapi.NewX509Source(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	policy, err := GetPolicy(source, "policy-123", predicate.AllowAll)
//	if err != nil {
//	    log.Printf("Error retrieving policy: %v", err)
//	    return
//	}
//	if policy == nil {
//	    log.Printf("Policy not found")
//	    return
//	}
//
//	log.Printf("Found policy: %+v", policy)
func GetPolicy(
	source *workloadapi.X509Source, id string,
	allow predicate.Predicate,
) (*data.Policy, error) {
	if source == nil {
		return nil, errors.New("nil X509Source")
	}

	r := reqres.PolicyReadRequest{ID: id}

	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New("getPolicy: I am having problem generating the payload"),
			err,
		)
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.PolicyGet(), mr)
	if err != nil {
		if errors.Is(err, net.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.PolicyReadResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Join(
			errors.New("getPolicy: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return nil, errors.New(string(res.Err))
	}

	return &res.Policy, nil
}
