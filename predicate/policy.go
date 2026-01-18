//	  \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//	\\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package predicate

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/config/auth"
	"github.com/spiffe/spike-sdk-go/spiffeid"
)

// PolicyAccessChecker is a function type that determines whether a SPIFFE ID
// has the required permissions for a given path.
//
// This type is used as a dependency injection point for policy-based access
// control checks. Implementations should verify if the peer identified by
// peerSPIFFEID has any of the specified permissions for the given path.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - path: string - The resource path being accessed
//   - perms: []data.PolicyPermission - The permissions required for access
//
// Returns:
//   - bool: true if access is allowed, false otherwise
type PolicyAccessChecker func(
	peerSPIFFEID string,
	path string,
	perms []data.PolicyPermission,
) bool

// AllowAllPolicies is a permissive PolicyAccessChecker that grants access
// regardless of the SPIFFE ID, path, or requested permissions.
//
// This function is intended for use cases where policy-level authorization
// is not required, but the calling code still expects a PolicyAccessChecker.
// It is commonly used with AuthorizeAndRespondOnFailNoPolicy to perform
// SPIFFE ID validation without fine-grained policy checks.
//
// This is especially useful in SPIKE components where no operator engagement
// happens, such as SPIKE Keeper.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer (ignored)
//   - path: string - The resource path being accessed (ignored)
//   - perms: []data.PolicyPermission - The permissions required (ignored)
//
// Returns:
//   - bool: Always returns true
func AllowAllPolicies(_ string, _ string, _ []data.PolicyPermission) bool {
	return true
}

// WithPolicyAccessChecker is a function type that validates access by
// delegating to a PolicyAccessChecker.
//
// This type enables flexible access control validation by accepting a resource
// identifier and a PolicyAccessChecker function. Implementations use the
// provided checker to determine if the current peer has the necessary
// permissions for the specified resource.
//
// Parameters:
//   - string: The SPIFFEID of the peer requesting access.
//   - PolicyAccessChecker: The function to use for policy-based access
//     validation
//
// Returns:
//   - bool: true if access is granted, false otherwise
type WithPolicyAccessChecker func(
	peerSPIFFEID string, check PolicyAccessChecker,
) bool

// ForPathWithPolicyAccessChecker is a function type that validates access by
// delegating to a PolicyAccessChecker, with an explicit resource path
// parameter.
//
// This is a path-aware variant of WithPolicyAccessChecker that includes the
// resource path in the function signature. Use this type when the access
// decision depends on the specific resource path being accessed, allowing
// for path-based authorization logic before delegating to the policy checker.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - path: string - The resource path being accessed
//   - check: PolicyAccessChecker - The function to use for policy-based access
//     validation
//
// Returns:
//   - bool: true if access is granted, false otherwise
type ForPathWithPolicyAccessChecker func(
	peerSPIFFEID string, path string, check PolicyAccessChecker,
) bool

// AllowSPIFFEIDForPolicyDelete checks if a SPIFFE ID is authorized to delete
// policies.
//
// This function verifies that the peer has write permission on the system
// policy access path (auth.PathSystemPolicyAccess).
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to delete policies, false otherwise
func AllowSPIFFEIDForPolicyDelete(
	peerSPIFFEID string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to delete policies.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemPolicyAccess,
		[]data.PolicyPermission{data.PermissionWrite}, checkAccess,
	)
}

// AllowSPIFFEIDForPolicyRead checks if a SPIFFE ID is authorized to read
// policies.
//
// This function verifies that the peer has read permission on the system
// policy access path (auth.PathSystemPolicyAccess).
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to read policies, false otherwise
func AllowSPIFFEIDForPolicyRead(
	peerSPIFFEID string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to read policies.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemPolicyAccess,
		[]data.PolicyPermission{data.PermissionRead}, checkAccess,
	)
}

// AllowSPIFFEIDForPolicyList checks if a SPIFFE ID is authorized to list
// policies.
//
// This function verifies that the peer has list permission on the system
// policy access path (auth.PathSystemPolicyAccess).
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to list policies, false otherwise
func AllowSPIFFEIDForPolicyList(
	peerSPIFFEID string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to list policies.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemPolicyAccess,
		[]data.PolicyPermission{data.PermissionList}, checkAccess,
	)
}

// AllowSPIFFEIDForPolicyWrite checks if a SPIFFE ID is authorized to write
// policies.
//
// This function verifies that the peer has write permission on the system
// policy access path (auth.PathSystemPolicyAccess).
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to write policies, false otherwise
func AllowSPIFFEIDForPolicyWrite(
	peerSPIFFEID string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to write policies.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemPolicyAccess,
		[]data.PolicyPermission{data.PermissionWrite}, checkAccess,
	)
}
