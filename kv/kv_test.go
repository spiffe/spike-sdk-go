//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		verify func(*testing.T, *KV)
	}{
		{
			name: "basic construction with typical MaxSecretVersions",
			config: Config{
				MaxSecretVersions: 10,
			},
			verify: func(t *testing.T, kv *KV) {
				if kv == nil {
					t.Fatal("New() returned nil")
				}
				if kv.maxSecretVersions != 10 {
					t.Errorf("maxSecretVersions = %d, want 10", kv.maxSecretVersions)
				}
				if kv.data == nil {
					t.Error("data map not initialized")
				}
				if len(kv.data) != 0 {
					t.Errorf("data map should be empty, got %d entries", len(kv.data))
				}
			},
		},
		{
			name: "construction with MaxSecretVersions=1",
			config: Config{
				MaxSecretVersions: 1,
			},
			verify: func(t *testing.T, kv *KV) {
				if kv.maxSecretVersions != 1 {
					t.Errorf("maxSecretVersions = %d, want 1", kv.maxSecretVersions)
				}
				// Verify it works correctly by doing a Put
				kv.Put("test", map[string]string{"key": "v1"})
				kv.Put("test", map[string]string{"key": "v2"})
				secret := kv.data["test"]
				if len(secret.Versions) != 1 {
					t.Errorf("with MaxVersions=1, should only keep 1 version, got %d",
						len(secret.Versions))
				}
			},
		},
		{
			name: "construction with large MaxSecretVersions",
			config: Config{
				MaxSecretVersions: 1000,
			},
			verify: func(t *testing.T, kv *KV) {
				if kv.maxSecretVersions != 1000 {
					t.Errorf("maxSecretVersions = %d, want 1000", kv.maxSecretVersions)
				}
			},
		},
		{
			name: "construction with MaxSecretVersions=0",
			config: Config{
				MaxSecretVersions: 0,
			},
			verify: func(t *testing.T, kv *KV) {
				// Document current behavior: 0 is accepted but may cause issues
				if kv.maxSecretVersions != 0 {
					t.Errorf("maxSecretVersions = %d, want 0", kv.maxSecretVersions)
				}
				// Note: This may cause all versions to be pruned immediately
				// Consider adding validation in New() if this is problematic
			},
		},
		{
			name: "construction with negative MaxSecretVersions",
			config: Config{
				MaxSecretVersions: -5,
			},
			verify: func(t *testing.T, kv *KV) {
				// Document current behavior: negative values are accepted
				if kv.maxSecretVersions != -5 {
					t.Errorf("maxSecretVersions = %d, want -5", kv.maxSecretVersions)
				}
				// Note: This may cause unexpected pruning behavior
				// Consider adding validation in New() if this is problematic
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := New(tt.config)
			tt.verify(t, kv)
		})
	}
}
