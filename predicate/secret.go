//	  \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//	\\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package predicate

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/spiffeid"
)

// AllowSPIFFEIDForSecretList checks if a SPIFFE ID is authorized to list
// secrets at a specific path.
//
// Secrets are path-bound (similar to Vault KVv2), so access control is
// evaluated against the provided path. This function first checks if the peer
// is a SPIKE Pilot operator (system workload), which is always permitted. Lite
// workloads are explicitly denied. For other workloads, it verifies that the
// peer has "list" permission on the specified path.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - path: string - The secret path to check access for
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to list secrets, false otherwise
func AllowSPIFFEIDForSecretList(
	peerSPIFFEID string, path string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to list secrets.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, path,
		[]data.PolicyPermission{data.PermissionList}, checkAccess,
	)
}

// AllowSPIFFEIDForSecretDelete checks if a SPIFFE ID is authorized to delete
// secrets at a specific path.
//
// Secrets are path-bound (similar to Vault KVv2), so access control is
// evaluated against the provided path. This function first checks if the peer
// is a SPIKE Pilot operator (system workload), which is always permitted. Lite
// workloads are explicitly denied. For other workloads, it verifies that the
// peer has "write" permission on the specified path.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - path: string - The secret path to check access for
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to delete secrets, false otherwise
func AllowSPIFFEIDForSecretDelete(
	peerSPIFFEID string, path string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to delete secrets.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, path,
		[]data.PolicyPermission{data.PermissionWrite}, checkAccess,
	)
}

// AllowSPIFFEIDForSecretWrite checks if a SPIFFE ID is authorized to write
// secrets at a specific path.
//
// Secrets are path-bound (similar to Vault KVv2), so access control is
// evaluated against the provided path. This function first checks if the peer
// is a SPIKE Pilot operator (system workload), which is always permitted. Lite
// workloads are explicitly denied. For other workloads, it verifies that the
// peer has "write" permission on the specified path.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - path: string - The secret path to check access for
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to write secrets, false otherwise
func AllowSPIFFEIDForSecretWrite(
	peerSPIFFEID string, path string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to write secrets.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, path,
		[]data.PolicyPermission{data.PermissionWrite}, checkAccess,
	)
}

// AllowSPIFFEIDForSecretRead checks if a SPIFFE ID is authorized to read
// secrets at a specific path.
//
// Secrets are path-bound (similar to Vault KVv2), so access control is
// evaluated against the provided path. This function first checks if the peer
// is a SPIKE Pilot operator (system workload), which is always permitted. Lite
// workloads are explicitly denied. For other workloads, it verifies that the
// peer has "read" permission on the specified path.
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - path: string - The secret path to check access for
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to read secrets, false otherwise
func AllowSPIFFEIDForSecretRead(
	peerSPIFFEID string, path string, checkAccess PolicyAccessChecker,
) bool {
	// SPIKE Pilot is a system workload; no policy check needed.
	if spiffeid.IsPilotOperator(peerSPIFFEID) {
		return true
	}
	// Lite workloads are not allowed to read secrets.
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return false
	}

	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, path,
		[]data.PolicyPermission{data.PermissionRead}, checkAccess,
	)
}
