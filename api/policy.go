//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/acl"
)

// CreatePolicy creates a new policy in the system. It establishes a mutual
// TLS connection using the X.509 source and sends a policy creation request
// to the server.
//
// The function takes the following parameters:
//   - name string: The name of the policy to be created
//   - spiffeIdPattern string: The SPIFFE ID pattern that this policy will apply
//     to
//   - pathPattern string: The path pattern that this policy will match against
//   - permissions []data.PolicyPermission: A slice of PolicyPermission defining
//     the access rights for this policy
//
// The function returns an error if any of the following operations fail:
//   - Marshaling the policy creation request
//   - Creating the mTLS client
//   - Making the HTTP POST request
//   - Unmarshaling the response
//   - Server-side policy creation (indicated in the response)
//
// Example usage:
//
//	permissions := []data.PolicyPermission{
//	    {
//	        Action: "read",
//	        Resource: "documents/*",
//	    },
//	}
//
//	err = api.CreatePolicy(
//	    "doc-reader",
//	    "spiffe://example.org/service/*",
//	    "/api/documents/*",
//	    permissions,
//	)
//	if err != nil {
//	    log.Printf("Failed to create policy: %v", err)
//	    return
//	}
func (a *API) CreatePolicy(
	name string, SPIFFEIDPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) error {
	return acl.CreatePolicy(a.source,
		name, SPIFFEIDPattern, pathPattern, permissions)
}

// DeletePolicy removes an existing policy from the system using its name.
//
// The function takes the following parameters:
//   - name string: The name of the policy to be deleted
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
//	err = api.DeletePolicy("doc-reader")
//	if err != nil {
//	    log.Printf("Failed to delete policy: %v", err)
//	    return
//	}
func (a *API) DeletePolicy(name string) error {
	return acl.DeletePolicy(a.source, name)
}

// GetPolicy retrieves a policy from the system using its name.
//
// The function takes the following parameters:
//   - name string: The name of the policy to retrieve
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
//	policy, err := api.GetPolicy("doc-reader")
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
func (a *API) GetPolicy(name string) (*data.Policy, error) {
	return acl.GetPolicy(a.source, name)
}

// ListPolicies retrieves policies from the system, optionally filtering by
// SPIFFE ID and path patterns.
//
// The function takes the following parameters:
//   - spiffeIdPattern string: The SPIFFE ID pattern to filter policies.
//     An empty string matches all SPIFFE IDs.
//   - pathPattern string: The path pattern to filter policies.
//     An empty string matches all paths.
//
// The function returns:
//   - (*[]data.Policy, nil) containing all matching policies if successful
//   - (nil, nil) if no policies are found
//   - (nil, error) if an error occurs during the operation
//
// Note: The returned slice pointer should be dereferenced before use:
//
//	policies := *result
//
// Errors can occur during:
//   - Marshaling the policy list request
//   - Creating the mTLS client
//   - Making the HTTP POST request (except for not found cases)
//   - Unmarshaling the response
//   - Server-side policy listing (indicated in the response)
//
// Example usage:
//
//	// List all policies
//	result, err := api.ListPolicies("", "")
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
func (a *API) ListPolicies(
	SPIFFEIDPattern, pathPattern string,
) (*[]data.Policy, error) {
	return acl.ListPolicies(a.source, SPIFFEIDPattern, pathPattern)
}
