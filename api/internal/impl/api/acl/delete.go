//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/spike-sdk-go/api/url"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/net"
)

// DeletePolicy removes an existing policy from the system using its ID.
// It requires a SPIFFE X.509 source for establishing a mutual TLS connection
// to make the deletion request.
//
// The function takes the following parameters:
//   - source: A pointer to a workloadapi.X509Source for establishing mTLS
//     connection
//   - id: The unique identifier of the policy to be deleted
//
// The function returns an error if any of the following operations fail:
//   - Marshaling the policy deletion request
//   - Creating the mTLS client
//   - Making the HTTP POST request
//   - Unmarshaling the response
//   - Server-side policy deletion (indicated in the response)
//
// Example usage:
//
//	source, err := workloadapi.NewX509Source(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	err = DeletePolicy(source, "policy-123")
//	if err != nil {
//	    log.Printf("Failed to delete policy: %v", err)
//	    return
//	}
func DeletePolicy(source *workloadapi.X509Source, id string) error {
	r := reqres.PolicyDeleteRequest{
		ID: id,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New(
				"deletePolicy: I am having problem generating the payload",
			),
			err,
		)
	}

	client, err := net.CreateMtlsClient(source)
	if err != nil {
		return err
	}

	body, err := net.Post(client, url.PolicyDelete(), mr)
	if err != nil {
		return err
	}

	res := reqres.PolicyDeleteResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Join(
			errors.New("deletePolicy: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return errors.New(string(res.Err))
	}

	return nil
}
