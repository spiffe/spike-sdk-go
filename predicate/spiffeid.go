//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
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

// AllowSPIFFEIDForPathAndPermissions checks if a SPIFFE ID is authorized to
// access a specific path with the given permissions.
//
// This is a general-purpose authorization function that delegates to the
// provided PolicyAccessChecker. It serves as the foundation for more specific
// authorization functions like AllowSPIFFEIDForPolicyDelete and
// AllowSPIFFEIDForPolicyRead.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - path: string - The resource path being accessed
//   - permissions: []data.PolicyPermission - The permissions required for access
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer has the required permissions, false otherwise
func AllowSPIFFEIDForPathAndPermissions(
	peerSPIFFEID string,
	path string, permissions []data.PolicyPermission,
	checkAccess PolicyAccessChecker,
) bool {
	return checkAccess(peerSPIFFEID, path, permissions)
}

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
	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemPolicyAccess,
		[]data.PolicyPermission{data.PermissionList}, checkAccess,
	)
}

func AllowSPIFFEIDForPolicyWrite(
	peerSPIFFEID string, checkAccess PolicyAccessChecker,
) bool {
	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemPolicyAccess,
		[]data.PolicyPermission{data.PermissionWrite}, checkAccess,
	)
}

// AllowSPIFFEIDForCipherDecrypt checks if a SPIFFE ID is authorized to perform
// cipher decryption operations.
//
// This function first checks if the peer is a "lite" workload, which is always
// permitted to decrypt. If not, it verifies that the peer has "execute"
// permission on the system cipher decrypt path (auth.PathSystemCipherDecrypt).
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to decrypt, false otherwise
func AllowSPIFFEIDForCipherDecrypt(
	peerSPIFFEID string, checkAccess PolicyAccessChecker,
) bool {
	// Lite workloads are always allowed:
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return true
	}
	// If not, do a policy check to determine if the request is allowed:
	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemCipherDecrypt,
		[]data.PolicyPermission{data.PermissionExecute}, checkAccess,
	)
}
