//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import sdkErrors "github.com/spiffe/spike-sdk-go/errors"

// Undelete restores previously deleted versions of a secret at the specified
// path. It sets the DeletedTime to nil for each specified version that exists.
//
// The function supports flexible version restoration with the following behavior:
//   - If versions is empty, restores only the current version
//   - If versions contains specific numbers, restores those versions
//   - Version 0 in the array represents the current version
//   - Non-existent versions are silently skipped without error
//
// Parameters:
//   - path: The location of the secret in the store
//   - versions: Array of version numbers to restore (empty array restores
//     current version only, 0 in the array represents current version)
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrEntityNotFound: if the path doesn't exist
//
// Example:
//
//	// Restore current version only
//	err := kv.Undelete("secret/path", []int{})
//	if err != nil {
//	    log.Printf("Failed to undelete secret: %v", err)
//	}
//
//	// Restore specific versions
//	err = kv.Undelete("secret/path", []int{1, 2, 3})
//	if err != nil {
//	    log.Printf("Failed to undelete versions: %v", err)
//	}
func (kv *KV) Undelete(path string, versions []int) *sdkErrors.SDKError {
	secret, exists := kv.data[path]
	if !exists {
		return sdkErrors.ErrEntityNotFound
	}

	cv := secret.Metadata.CurrentVersion

	// If no versions specified, mark the latest version as undeleted
	if len(versions) == 0 {
		if v, exists := secret.Versions[cv]; exists {
			v.DeletedTime = nil // Mark as undeleted.
			secret.Versions[cv] = v
		}

		return nil
	}

	// Delete specific versions
	for _, version := range versions {
		if version == 0 {
			v, exists := secret.Versions[cv]
			if !exists {
				continue
			}

			v.DeletedTime = nil // Mark as undeleted.
			secret.Versions[cv] = v
			continue
		}

		if v, exists := secret.Versions[version]; exists {
			v.DeletedTime = nil // Mark as undeleted.
			secret.Versions[version] = v
		}
	}

	return nil
}
