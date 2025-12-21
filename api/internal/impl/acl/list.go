//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"context"
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/net"
)

// ListPolicies retrieves policies from the system, optionally filtering by
// SPIFFE ID and path patterns. It establishes a mutual TLS connection to
// SPIKE Nexus using the X.509 source and sends a policy list request.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - SPIFFEIDPattern: The SPIFFE ID pattern to filter policies. An empty
//     string matches all SPIFFE IDs.
//   - pathPattern: The path pattern to filter policies. An empty string
//     matches all paths.
//
// Returns:
//   - (*[]data.Policy, nil) containing all matching policies if successful
//   - (nil, nil) if no policies are found
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPIPostFailed: if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error (e.g.,
//     ErrAccessUnauthorized, ErrAPIBadRequest, etc.)
//
// Note: The returned slice pointer should be dereferenced before use:
//
//	policies := *result
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
//	// List all policies
//	result, err := ListPolicies(ctx, source, "", "")
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
	ctx context.Context,
	source *workloadapi.X509Source,
	SPIFFEIDPattern string, pathPattern string,
) (*[]data.PolicyListItem, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	r := reqres.PolicyListRequest{
		SPIFFEIDPattern: SPIFFEIDPattern,
		PathPattern:     pathPattern,
	}
	mr, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return nil, failErr
	}

	res, postErr := net.PostAndUnmarshal[reqres.PolicyListResponse](
		ctx, source, url.PolicyList(), mr)
	if postErr != nil {
		if postErr.Is(sdkErrors.ErrAPINotFound) {
			return &([]data.PolicyListItem{}), nil
		}
		return nil, postErr
	}

	return &res.Policies, nil
}
