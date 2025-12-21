//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"context"
	"encoding/json"

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
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
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
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	r := reqres.PolicyDeleteRequest{ID: id}

	mr, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return failErr
	}

	_, postErr := net.PostAndUnmarshal[reqres.PolicyDeleteResponse](
		ctx, source, url.PolicyDelete(), mr)
	return postErr
}
