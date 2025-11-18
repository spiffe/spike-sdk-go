//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/net"
)

// DeletePolicy removes an existing policy from the system using its ID.
// It establishes a mutual TLS connection to SPIKE Nexus using the X.509 source
// and sends a policy deletion request.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - id: The unique identifier of the policy to be deleted
//
// Returns:
//   - nil if successful
//   - error if marshaling, network request, or server-side deletion fails
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	err = DeletePolicy(source, "policy-123")
//	if err != nil {
//	    log.Printf("Failed to delete policy: %v", err)
//	}
func DeletePolicy(
	source *workloadapi.X509Source,
	id string,
) error {
	if source == nil {
		return sdkErrors.ErrNilX509Source
	}

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

	client := net.CreateMTLSClientForNexus(source)

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
