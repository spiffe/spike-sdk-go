//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Destroy permanently removes a secret path from the store, including all
// versions (both active and deleted). This is a hard delete operation that
// cannot be undone.
//
// Unlike Delete(), which soft-deletes versions by marking them with
// DeletedTime, Destroy() completely removes the path from the internal map,
// reclaiming the memory.
//
// This operation is useful for:
//   - Purging secrets that have all versions deleted
//   - Removing obsolete paths to prevent unbounded map growth
//   - Compliance requirements for data removal
//
// Parameters:
//   - path: The path to permanently remove from the store
//
// Returns:
//   - *sdkErrors.SDKError: ErrEntityNotFound if the path does not exist,
//     nil on success
//
// Example:
//
//	// Delete all versions first
//	kv.Delete("secret/path", []int{})
//
//	// Check if empty and destroy
//	secret, _ := kv.GetRawSecret("secret/path")
//	if secret.IsEmpty() {
//	    err := kv.Destroy("secret/path")
//	    if err != nil {
//	        log.Printf("Failed to destroy secret: %v", err)
//	    }
//	}
//
//	// Or destroy directly (removes regardless of deletion state)
//	err := kv.Destroy("secret/path")
func (kv *KV) Destroy(path string) *sdkErrors.SDKError {
	if _, exists := kv.data[path]; !exists {
		return sdkErrors.ErrEntityNotFound
	}

	delete(kv.data, path)
	return nil
}
