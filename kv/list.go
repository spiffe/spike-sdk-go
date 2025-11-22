//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

// List returns a slice containing all paths stored in the key-value store.
// The order of paths in the returned slice is not guaranteed to be stable
// between calls.
//
// Note: List returns all paths regardless of whether their versions have been
// deleted. A path is only removed from the store when all of its data is
// explicitly removed, not when versions are soft-deleted.
//
// Returns:
//   - []string: A slice containing all paths present in the store
//
// Example:
//
//	kv := New(Config{MaxSecretVersions: 10})
//	kv.Put("app/config", map[string]string{"key": "value"})
//	kv.Put("app/database", map[string]string{"host": "localhost"})
//
//	paths := kv.List()
//	// Returns: ["app/config", "app/database"] (order not guaranteed)
func (kv *KV) List() []string {
	keys := make([]string, 0, len(kv.data))

	for k := range kv.data {
		keys = append(keys, k)
	}

	return keys
}
