//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	"time"

	"github.com/spiffe/spike-sdk-go/errors"
)

// Delete marks secret versions as deleted for a given path. If no versions are
// specified, it marks only the current version as deleted. If specific versions
// are provided, it marks each existing version in the list as deleted. The
// deletion is performed by setting the DeletedTime to the current time.
//
// Parameters:
//   - path: Path to the secret to delete
//   - versions: Array of version numbers to delete (empty array deletes current
//     version only, 0 in the array represents current version)
//
// Returns:
//   - *errors.SDKError: nil on success, or one of the following errors:
//   - ErrEntityNotFound: if the path doesn't exist
//
// Example:
//
//	// Delete current version only
//	err := kv.Delete("secret/path", []int{})
//	if err != nil {
//	    log.Printf("Failed to delete secret: %v", err)
//	}
//
//	// Delete specific versions
//	err = kv.Delete("secret/path", []int{1, 2, 3})
//	if err != nil {
//	    log.Printf("Failed to delete versions: %v", err)
//	}
func (kv *KV) Delete(path string, versions []int) *errors.SDKError {
	secret, exists := kv.data[path]
	if !exists {
		return errors.ErrEntityNotFound
	}

	now := time.Now()
	cv := secret.Metadata.CurrentVersion

	// If no versions specified, mark the latest version as deleted
	if len(versions) == 0 {
		if v, exists := secret.Versions[cv]; exists {
			v.DeletedTime = &now // Mark as deleted.
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

			v.DeletedTime = &now // Mark as deleted.
			secret.Versions[cv] = v
			continue
		}

		if v, exists := secret.Versions[version]; exists {
			v.DeletedTime = &now // Mark as deleted.
			secret.Versions[version] = v
		}
	}

	return nil
}
