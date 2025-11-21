//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	"time"
)

// Put stores a new version of key-value pairs at the specified path in the
// store. It implements automatic versioning as a bounded cache with a
// configurable maximum number of versions per path.
//
// When storing values:
//   - If the path doesn't exist, it creates new data with initial metadata
//   - Each put operation creates a new version with an incremented version
//     number
//   - Old versions are automatically pruned when they fall outside the version
//     window (CurrentVersion - MaxVersions)
//   - All versions exceeding MaxVersions are pruned in a single operation,
//     maintaining the most recent MaxVersions versions
//   - Timestamps are updated for both creation and modification times
//
// Version pruning behavior (bounded cache):
//   - Pruning occurs on each Put when versions exceed MaxVersions
//   - All versions older than (CurrentVersion - MaxVersions) are deleted
//   - Example: If CurrentVersion=15 and MaxVersions=10, versions 1-5 are
//     deleted, keeping versions 6-15
//   - This ensures O(n) pruning where n is the number of excess versions,
//     providing predictable performance
//
// Parameters:
//   - path: The location where the data will be stored
//   - values: A map of key-value pairs to store at this path
//
// Example:
//
//	kv := New(Config{MaxSecretVersions: 10})
//	kv.Put("app/config", map[string]string{
//	    "api_key": "secret123",
//	    "timeout": "30s",
//	})
//	// Creates version 1 at path "app/config"
//
//	kv.Put("app/config", map[string]string{
//	    "api_key": "newsecret456",
//	    "timeout": "60s",
//	})
//	// Creates version 2, version 1 is still available
//
// The function maintains metadata including:
//   - CreatedTime: When the data at this path was first created
//   - UpdatedTime: When the most recent version was added
//   - CurrentVersion: The latest version number
//   - OldestVersion: The oldest available version number after pruning
//   - MaxVersions: Maximum number of versions to keep (configurable at KV
//     creation)
func (kv *KV) Put(path string, values map[string]string) {
	rightNow := time.Now()

	secret, exists := kv.data[path]
	if !exists {
		secret = &Value{
			Versions: make(map[int]Version),
			Metadata: Metadata{
				CreatedTime: rightNow,
				UpdatedTime: rightNow,
				MaxVersions: kv.maxSecretVersions,
				// Versions start at 1, so that passing 0 as the version will
				// default to the current version.
				CurrentVersion: 1,
				OldestVersion:  1,
			},
		}
		kv.data[path] = secret
	} else {
		secret.Metadata.CurrentVersion++
	}

	newVersion := secret.Metadata.CurrentVersion

	// Add a new version:
	secret.Versions[newVersion] = Version{
		Data:        values,
		CreatedTime: rightNow,
		Version:     newVersion,
	}

	// Update metadata
	secret.Metadata.UpdatedTime = rightNow

	// Clean up the old versions if exceeding MaxVersions
	var deletedAny bool
	for version := range secret.Versions {
		if newVersion-version >= secret.Metadata.MaxVersions {
			delete(secret.Versions, version)
			deletedAny = true
		}
	}

	if deletedAny {
		oldestVersion := secret.Metadata.CurrentVersion
		for version := range secret.Versions {
			if version < oldestVersion {
				oldestVersion = version
			}
		}
		secret.Metadata.OldestVersion = oldestVersion
	}
}
