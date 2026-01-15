//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package predicate

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/config/auth"
	"github.com/spiffe/spike-sdk-go/spiffeid"
)

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

// AllowSPIFFEIDForCipherEncrypt checks if a SPIFFE ID is authorized to perform
// cipher encryption operations.
//
// This function first checks if the peer is a "lite" workload, which is always
// permitted to encrypt. If not, it verifies that the peer has "execute"
// permission on the system cipher encrypt path (auth.PathSystemCipherEncrypt).
//
// Parameters:
//   - peerSPIFFEID: string - The SPIFFE ID of the peer requesting access
//   - checkAccess: PolicyAccessChecker - The function to perform the access
//     check
//
// Returns:
//   - bool: true if the peer is authorized to encrypt, false otherwise
func AllowSPIFFEIDForCipherEncrypt(
	peerSPIFFEID string, checkAccess PolicyAccessChecker,
) bool {
	// Lite workloads are always allowed:
	if spiffeid.IsLiteWorkload(peerSPIFFEID) {
		return true
	}
	// If not, do a policy check to determine if the request is allowed:
	return AllowSPIFFEIDForPathAndPermissions(
		peerSPIFFEID, auth.PathSystemCipherEncrypt,
		[]data.PolicyPermission{data.PermissionExecute}, checkAccess,
	)
}
