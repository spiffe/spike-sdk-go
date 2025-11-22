//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package kv provides an in-memory key-value store with automatic versioning
// and bounded cache semantics.
//
// # Concurrency Safety
//
// This package is NOT safe for concurrent use. All methods on KV must be
// externally synchronized. Callers are responsible for providing appropriate
// locking mechanisms (e.g., sync.RWMutex) to protect concurrent access.
//
// Concurrent operations without synchronization will cause data races and
// undefined behavior.
//
// Example of safe concurrent usage:
//
//	type SafeStore struct {
//	    kv *kv.KV
//	    mu sync.RWMutex
//	}
//
//	func (s *SafeStore) Put(path string, data map[string]string) {
//	    s.mu.Lock()
//	    defer s.mu.Unlock()
//	    s.kv.Put(path, data)
//	}
//
//	func (s *SafeStore) Get(path string, version int) (map[string]string, error) {
//	    s.mu.RLock()
//	    defer s.mu.RUnlock()
//	    return s.kv.Get(path, version)
//	}
//
// Use sync.RWMutex to allow concurrent reads while serializing writes for
// optimal performance.
package kv

// KV represents an in-memory key-value store with automatic versioning and
// bounded cache semantics. Each path maintains a configurable maximum number
// of versions, with older versions automatically pruned when the limit is
// exceeded.
//
// The store supports:
//   - Versioned storage with automatic version numbering
//   - Soft deletion with undelete capability
//   - Bounded cache with automatic pruning of old versions
//   - Version-specific retrieval and metadata tracking
type KV struct {
	maxSecretVersions int
	data              map[string]*Value
}

// Config represents the configuration for a KV instance.
type Config struct {
	// MaxSecretVersions is the maximum number of versions to retain per path.
	// When exceeded, older versions are automatically pruned.
	// Must be positive. A typical value is 10.
	MaxSecretVersions int
}

// New creates a new KV instance with the specified configuration.
//
// The store is initialized as an empty in-memory key-value store with
// versioning enabled. All paths stored in this instance will retain up to
// MaxSecretVersions versions.
//
// Parameters:
//   - config: Configuration specifying MaxSecretVersions
//
// Returns:
//   - *KV: A new KV instance ready for use
//
// Example:
//
//	kv := New(Config{MaxSecretVersions: 10})
//	kv.Put("app/config", map[string]string{"key": "value"})
func New(config Config) *KV {
	return &KV{
		maxSecretVersions: config.MaxSecretVersions,
		data:              make(map[string]*Value),
	}
}
