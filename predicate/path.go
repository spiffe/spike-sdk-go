//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package predicate

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/spiffeid"
)

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
//   - permissions: []data.PolicyPermission - The permissions required for
//     access
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
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}

	return checkAccess(peerSPIFFEID, path, permissions)
}
