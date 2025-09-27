//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package predicate provides SPIFFE ID validation predicates for SPIKE API
// access control.
//
// This package defines predicate functions that can be used to validate
// SPIFFE IDs in API calls, enabling fine-grained access control based on
// workload identity.
// Predicates are used by API methods to restrict access to specific types of
// workloads (e.g., only SPIKE Pilot instances).
package predicate

import "github.com/spiffe/spike-sdk-go/spiffeid"

// Predicate is a function type that validates a SPIFFE ID string.
// It returns true if the SPIFFE ID should be allowed access, false otherwise.
//
// Predicates are used throughout the SPIKE API to implement access control
// policies based on workload identity. They are typically passed to API methods
// to restrict which workloads can perform specific operations.
//
// Example usage:
//
//	// Create a predicate that only allows pilot workloads
//	pilotPredicate := AllowPilot("example.org")
//
//	// Use in an API call
//	policy, err := acl.GetPolicy(source, policyID, pilotPredicate)
type Predicate func(string) bool

// AllowAll is a predicate that accepts any SPIFFE ID.
// This effectively disables access control and should be used with caution.
// It's typically used when policy-based access control is handled at a higher level.
//
// Example usage:
//
//	// Allow any workload to access the API
//	secret, err := secret.Get(source, path, version, AllowAll)
var AllowAll = Predicate(func(_ string) bool { return true })

// DenyAll is a predicate that rejects all SPIFFE IDs.
// This can be used to temporarily disable access or as a default restrictive
// policy.
//
// Example usage:
//
//	// Deny all access during maintenance
//	policy, err := acl.GetPolicy(source, policyID, DenyAll)
var DenyAll = Predicate(func(_ string) bool { return false })

// AllowPilot creates a predicate that only allows SPIKE Pilot workloads.
// It returns a predicate function that validates whether a given SPIFFE ID
// matches the SPIKE Pilot identity pattern for the specified trust domains.
//
// This is used to restrict API access to only SPIKE Pilot instances, providing
// an additional layer of security for sensitive operations that should only
// be performed by the control plane.
//
// Parameters:
//   - trustRoots: Comma-delimited list of trust domain roots (e.g.,
//     "example.org,other.org")
//
// Returns:
//   - Predicate: A function that returns true only for SPIKE Pilot SPIFFE IDs
//
// Example usage:
//
//	// Create predicate for pilot-only access
//	pilotOnly := AllowPilot("example.org,dev.example.org")
//
//	// Use in API calls to restrict access
//	policy, err := acl.GetPolicy(source, policyID, pilotOnly)
//	secret, err := secret.Get(source, secretPath, version, pilotOnly)
//
// The returned predicate will accept SPIFFE IDs matching:
//   - "spiffe://example.org/spike/pilot"
//   - "spiffe://example.org/spike/pilot/instance-1"
//   - "spiffe://dev.example.org/spike/pilot"
//   - etc.
//
// if the trust root of the SPIFFE ID belongs to one of the specified domains
// in the trustRoots input parameter.
func AllowPilot(trustRoots string) Predicate {
	return func(spiffeID string) bool {
		return spiffeid.IsPilot(trustRoots, spiffeID)
	}
}
