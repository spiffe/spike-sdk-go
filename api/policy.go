//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/acl"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// CreatePolicy creates a new policy in the system.
//
// It establishes a mutual TLS connection using the X.509 source and sends a
// policy creation request to SPIKE Nexus.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - name: The name of the policy to be created
//   - SPIFFEIDPattern: The SPIFFE ID pattern that this policy will apply to
//   - pathPattern: The path pattern that this policy will match against
//   - permissions: A slice of PolicyPermission defining the access rights
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	permissions := []data.PolicyPermission{
//	    {Action: "read", Resource: "documents/*"},
//	}
//	err := api.CreatePolicy(
//	    ctx,
//	    "doc-reader",
//	    "spiffe://example.org/service/*",
//	    "/api/documents/*",
//	    permissions,
//	)
//	if err != nil {
//	    log.Printf("Failed to create policy: %v", err)
//	}
func (a *API) CreatePolicy(
	ctx context.Context,
	name string, SPIFFEIDPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) *sdkErrors.SDKError {
	return acl.CreatePolicy(ctx, a.source,
		name, SPIFFEIDPattern, pathPattern, permissions)
}

// DeletePolicy removes an existing policy from the system using its unique ID.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - id: The unique identifier of the policy to be deleted
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	err := api.DeletePolicy(ctx, "policy-123")
//	if err != nil {
//	    log.Printf("Failed to delete policy: %v", err)
//	}
func (a *API) DeletePolicy(ctx context.Context, id string) *sdkErrors.SDKError {
	return acl.DeletePolicy(ctx, a.source, id)
}

// GetPolicy retrieves a policy from the system using its unique ID.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - id: The unique identifier of the policy to retrieve
//
// Returns:
//   - *data.Policy: The policy if found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPINotFound: if the policy is not found
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	policy, err := api.GetPolicy(ctx, "policy-123")
//	if err != nil {
//	    if err.Is(sdkErrors.ErrAPINotFound) {
//	        log.Printf("Policy not found")
//	        return
//	    }
//	    log.Printf("Error retrieving policy: %v", err)
//	    return
//	}
//	log.Printf("Found policy: %+v", policy)
func (a *API) GetPolicy(
	ctx context.Context, id string,
) (*data.Policy, *sdkErrors.SDKError) {
	return acl.GetPolicy(ctx, a.source, id)
}

// ListPolicies retrieves policies from the system, optionally filtering by
// SPIFFE ID and path patterns.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - SPIFFEIDPattern: The SPIFFE ID pattern to filter policies (empty string
//     matches all SPIFFE IDs)
//   - pathPattern: The path pattern to filter policies (empty string matches
//     all paths)
//
// Returns:
//   - *[]data.PolicyListItem: Array of policy list items (ID and Name),
//     empty array if none found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrAPINotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (&[]data.PolicyListItem{}, nil) if no policies are found (ErrAPINotFound)
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	result, err := api.ListPolicies(ctx, "", "")
//	if err != nil {
//	    log.Printf("Error listing policies: %v", err)
//	    return
//	}
//	policies := *result // slice of data.PolicyListItem
//	for _, policy := range policies {
//	    log.Printf("Found policy: %+v", policy)
//	}
func (a *API) ListPolicies(
	ctx context.Context, SPIFFEIDPattern, pathPattern string,
) (*[]data.PolicyListItem, *sdkErrors.SDKError) {
	return acl.ListPolicies(ctx, a.source, SPIFFEIDPattern, pathPattern)
}
