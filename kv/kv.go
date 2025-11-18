//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

// KV represents an in-memory key-value store with versioning
type KV struct {
	maxSecretVersions int
	data              map[string]*Value
}

// Config represents the configuration for a KV instance
type Config struct {
	MaxSecretVersions int
}

// New creates a new KV instance
//
// Parameters:
//   - config: KV Configuration
func New(config Config) *KV {
	return &KV{
		maxSecretVersions: config.MaxSecretVersions,
		data:              make(map[string]*Value),
	}
}
