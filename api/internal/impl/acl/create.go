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
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrMarshalFailure: if request serialization fails
//   - ErrPostFailed: if the HTTP request fails
//   - ErrUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error (e.g.,
//     ErrUnauthorized, ErrBadRequest, etc.)
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
) *sdkErrors.SDKError {
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source
	}

	r := reqres.PolicyCreateRequest{
		Name:            name,
		SPIFFEIDPattern: SPIFFEIDPattern,
		PathPattern:     pathPattern,
		Permissions:     permissions,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		failErr := sdkErrors.ErrMarshalFailure.Wrap(err)
		failErr.Msg = "problem generating the payload"
		return failErr
	}

	client := net.CreateMTLSClientForNexus(source)

	body, err := net.Post(client, url.PolicyCreate(), mr)
	if err != nil {
		return err
	}

	res := reqres.PolicyCreateResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		failErr := sdkErrors.ErrUnmarshalFailure.Wrap(err)
		failErr.Msg = "problem parsing response body"
		return failErr
	}
	if res.Err != "" {
		return sdkErrors.FromCode(res.Err)
	}

	return nil
}
