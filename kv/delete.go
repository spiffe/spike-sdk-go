//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	"time"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Delete marks secret versions as deleted for a given path. The deletion is
// performed by setting the DeletedTime to the current time.
//
// IMPORTANT: This is a soft delete. The path remains in the store even if all
// versions are deleted. To completely remove a path and reclaim memory, use
// Destroy() after deleting all versions.
//
// The function supports flexible version deletion with the following behavior:
//   - If versions is empty, deletes only the current version
//   - If versions contains specific numbers, deletes those versions
//   - Version 0 in the array represents the current version
//   - Non-existent versions are silently skipped without error
//
// This idempotent behavior is useful for batch operations where you want to
// ensure certain versions are deleted without failing if some don't exist.
//
// Parameters:
//   - path: Path to the secret to delete
//   - versions: Array of version numbers to delete (empty array deletes current
//     version only, 0 in the array represents current version)
//
// Returns:
//   - []int: Array of version numbers that were actually modified (had their
//     DeletedTime changed from nil to now). Already-deleted versions are not
//     included in this list.
//   - *errors.SDKError: nil on success, or one of the following sdkErrors:
//   - ErrEntityNotFound: if the path doesn't exist
//
// Example:
//
//	// Delete current version only
//	modified, err := kv.Delete("secret/path", []int{})
//	if err != nil {
//	    log.Printf("Failed to delete secret: %v", err)
//	}
//	log.Printf("Deleted %d version(s): %v", len(modified), modified)
//
//	// Delete specific versions
//	modified, err = kv.Delete("secret/path", []int{1, 2, 3})
//	if err != nil {
//	    log.Printf("Failed to delete versions: %v", err)
//	}
//	log.Printf("Actually deleted: %v", modified)
func (kv *KV) Delete(path string, versions []int) ([]int, *sdkErrors.SDKError) {
	secret, exists := kv.data[path]
	if !exists {
		return nil, sdkErrors.ErrEntityNotFound.Clone()
	}

	now := time.Now()
	cv := secret.Metadata.CurrentVersion
	var modified []int

	// If no versions specified, mark the latest version as deleted
	if len(versions) == 0 {
		if v, versionExists := secret.Versions[cv]; versionExists && v.DeletedTime == nil {
			v.DeletedTime = &now // Mark as deleted.
			secret.Versions[cv] = v
			modified = append(modified, cv)
		}

		return modified, nil
	}

	// Delete specific versions
	for _, version := range versions {
		if version == 0 {
			v, versionExists := secret.Versions[cv]
			if !versionExists || v.DeletedTime != nil {
				continue
			}

			v.DeletedTime = &now // Mark as deleted.
			secret.Versions[cv] = v
			modified = append(modified, cv)
			continue
		}

		if v, versionExists := secret.Versions[version]; versionExists && v.DeletedTime == nil {
			v.DeletedTime = &now // Mark as deleted.
			secret.Versions[version] = v
			modified = append(modified, version)
		}
	}

	return modified, nil
}
