//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"context"

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
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPIPostFailed: if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error (e.g.,
//     ErrAccessUnauthorized, ErrAPINotFound, ErrAPIBadRequest, etc.)
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	err = DeletePolicy(ctx, source, "policy-123")
//	if err != nil {
//	    log.Printf("Failed to delete policy: %v", err)
//	}
func DeletePolicy(
	ctx context.Context,
	source *workloadapi.X509Source,
	id string,
) *sdkErrors.SDKError {
	r := reqres.PolicyDeleteRequest{ID: id}

	_, err := net.DoPost[reqres.PolicyDeleteResponse](
		ctx, source, url.PolicyDelete(), r,
	)
	return err
}
