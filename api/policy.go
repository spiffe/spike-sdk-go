//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/acl"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// CreatePolicy creates a new policy in the system.
//
// It establishes a mutual TLS connection using the X.509 source and sends a
// policy creation request to SPIKE Nexus.
//
// Parameters:
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
//	permissions := []data.PolicyPermission{
//	    {Action: "read", Resource: "documents/*"},
//	}
//	err := api.CreatePolicy(
//	    "doc-reader",
//	    "spiffe://example.org/service/*",
//	    "/api/documents/*",
//	    permissions,
//	)
//	if err != nil {
//	    log.Printf("Failed to create policy: %v", err)
//	}
func (a *API) CreatePolicy(
	name string, SPIFFEIDPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) *sdkErrors.SDKError {
	return acl.CreatePolicy(a.source,
		name, SPIFFEIDPattern, pathPattern, permissions)
}

// DeletePolicy removes an existing policy from the system using its unique ID.
//
// Parameters:
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
//	err := api.DeletePolicy("policy-123")
//	if err != nil {
//	    log.Printf("Failed to delete policy: %v", err)
//	}
func (a *API) DeletePolicy(id string) *sdkErrors.SDKError {
	return acl.DeletePolicy(a.source, id)
}

// GetPolicy retrieves a policy from the system using its unique ID.
//
// Parameters:
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
//	policy, err := api.GetPolicy("policy-123")
//	if err != nil {
//	    if err.Is(sdkErrors.ErrAPINotFound) {
//	        log.Printf("Policy not found")
//	        return
//	    }
//	    log.Printf("Error retrieving policy: %v", err)
//	    return
//	}
//	log.Printf("Found policy: %+v", policy)
func (a *API) GetPolicy(id string) (*data.Policy, *sdkErrors.SDKError) {
	return acl.GetPolicy(a.source, id)
}

// ListPolicies retrieves policies from the system, optionally filtering by
// SPIFFE ID and path patterns.
//
// Parameters:
//   - SPIFFEIDPattern: The SPIFFE ID pattern to filter policies (empty string
//     matches all SPIFFE IDs)
//   - pathPattern: The path pattern to filter policies (empty string matches
//     all paths)
//
// Returns:
//   - *[]reqres.PolicyListItem: Array of policy list items (ID and Name),
//     empty array if none found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrAPINotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (&[]reqres.PolicyListItem{}, nil) if no policies are found (ErrAPINotFound)
//
// Example:
//
//	result, err := api.ListPolicies("", "")
//	if err != nil {
//	    log.Printf("Error listing policies: %v", err)
//	    return
//	}
//	policies := *result // slice of reqres.PolicyListItem
//	for _, policy := range policies {
//	    log.Printf("Found policy: %+v", policy)
//	}
func (a *API) ListPolicies(
	SPIFFEIDPattern, pathPattern string,
) (*[]reqres.PolicyListItem, *sdkErrors.SDKError) {
	return acl.ListPolicies(a.source, SPIFFEIDPattern, pathPattern)
}
