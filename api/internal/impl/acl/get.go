//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/net"
)

// GetPolicy retrieves a policy from the system using its ID.
// It establishes a mutual TLS connection to SPIKE Nexus using the X.509 source
// and sends a policy retrieval request.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - id: The unique identifier of the policy to retrieve
//
// Returns:
//   - (*data.Policy, nil) if the policy is found
//   - (nil, nil) if the policy is not found
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPIPostFailed: if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error (e.g.,
//     ErrAccessUnauthorized, ErrAPIBadRequest, etc.)
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	policy, err := GetPolicy(source, "policy-123")
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
) (*data.Policy, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source
	}

	r := reqres.PolicyReadRequest{ID: id}

	mr, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return nil, failErr
	}

	res, postErr := net.PostAndUnmarshal[reqres.PolicyReadResponse](
		source, url.PolicyGet(), mr)
	if postErr != nil {
		if postErr.Is(sdkErrors.ErrAPINotFound) {
			return nil, nil
		}
		return nil, postErr
	}

	return &res.Policy, nil
}
