//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	"testing"
	"time"
)

// TestKV_Delete_ModifiedReturnValue tests the modified versions return value
// to ensure Delete properly reports which versions were actually modified.
func TestKV_Delete_ModifiedReturnValue(t *testing.T) {
	tests := []struct {
		name         string
		setup        func() *KV
		path         string
		versions     []int
		wantModified []int
		wantErr      bool
	}{
		{
			name: "delete current version returns modified",
			setup: func() *KV {
				kv := New(Config{MaxSecretVersions: 10})
				kv.Put("test/path", map[string]string{"key": "value"})
				return kv
			},
			path:         "test/path",
			versions:     []int{},
			wantModified: []int{1}, // Current version is 1
			wantErr:      false,
		},
		{
			name: "delete already deleted returns empty",
			setup: func() *KV {
				kv := New(Config{MaxSecretVersions: 10})
				kv.Put("test/path", map[string]string{"key": "value"})
				_, _ = kv.Delete("test/path", []int{1}) // Delete it first
				return kv
			},
			path:         "test/path",
			versions:     []int{1},
			wantModified: []int{}, // Already deleted, no modifications
			wantErr:      false,
		},
		{
			name: "delete multiple versions returns all modified",
			setup: func() *KV {
				kv := New(Config{MaxSecretVersions: 10})
				kv.Put("test/path", map[string]string{"key": "v1"})
				kv.Put("test/path", map[string]string{"key": "v2"})
				kv.Put("test/path", map[string]string{"key": "v3"})
				return kv
			},
			path:         "test/path",
			versions:     []int{1, 2, 3},
			wantModified: []int{1, 2, 3},
			wantErr:      false,
		},
		{
			name: "delete mix of existing and non-existing returns only existing",
			setup: func() *KV {
				kv := New(Config{MaxSecretVersions: 10})
				kv.Put("test/path", map[string]string{"key": "v1"})
				kv.Put("test/path", map[string]string{"key": "v2"})
				return kv
			},
			path:         "test/path",
			versions:     []int{1, 2, 99, 100}, // 99, 100 don't exist
			wantModified: []int{1, 2},
			wantErr:      false,
		},
		{
			name: "delete version 0 returns current version number",
			setup: func() *KV {
				kv := New(Config{MaxSecretVersions: 10})
				kv.Put("test/path", map[string]string{"key": "v1"})
				kv.Put("test/path", map[string]string{"key": "v2"})
				kv.Put("test/path", map[string]string{"key": "v3"})
				return kv
			},
			path:         "test/path",
			versions:     []int{0}, // 0 means current
			wantModified: []int{3}, // Current is version 3
			wantErr:      false,
		},
		{
			name: "delete non-existent versions returns empty",
			setup: func() *KV {
				kv := New(Config{MaxSecretVersions: 10})
				kv.Put("test/path", map[string]string{"key": "v1"})
				return kv
			},
			path:         "test/path",
			versions:     []int{99, 100, 101},
			wantModified: []int{},
			wantErr:      false,
		},
		{
			name: "delete some already deleted returns only newly deleted",
			setup: func() *KV {
				kv := New(Config{MaxSecretVersions: 10})
				kv.Put("test/path", map[string]string{"key": "v1"})
				kv.Put("test/path", map[string]string{"key": "v2"})
				kv.Put("test/path", map[string]string{"key": "v3"})
				_, _ = kv.Delete("test/path", []int{1}) // Pre-delete version 1
				return kv
			},
			path:         "test/path",
			versions:     []int{1, 2, 3}, // 1 already deleted
			wantModified: []int{2, 3},    // Only 2 and 3 are newly deleted
			wantErr:      false,
		},
		{
			name: "delete empty versions on empty current returns empty",
			setup: func() *KV {
				kv := &KV{
					data: make(map[string]*Value),
				}
				// Manually create state where current version doesn't exist
				kv.data["test/path"] = &Value{
					Metadata: Metadata{
						CurrentVersion: 1,
					},
					Versions: map[int]Version{}, // No versions!
				}
				return kv
			},
			path:         "test/path",
			versions:     []int{},
			wantModified: []int{}, // Current doesn't exist, nothing modified
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup()
			modified, err := kv.Delete(tt.path, tt.versions)

			if tt.wantErr {
				if err == nil {
					t.Error("Delete() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Delete() unexpected error: %v", err)
				return
			}

			// Check modified length
			if len(modified) != len(tt.wantModified) {
				t.Errorf("Delete() modified length = %d, want %d (got: %v, want: %v)",
					len(modified), len(tt.wantModified), modified, tt.wantModified)
				return
			}

			// Check modified contains all expected versions
			modifiedMap := make(map[int]bool)
			for _, v := range modified {
				modifiedMap[v] = true
			}

			for _, wantVer := range tt.wantModified {
				if !modifiedMap[wantVer] {
					t.Errorf("Delete() modified missing version %d, got: %v, want: %v",
						wantVer, modified, tt.wantModified)
				}
			}
		})
	}
}

// TestKV_Delete_StateVerification verifies internal state consistency
// after delete operations.
func TestKV_Delete_StateVerification(t *testing.T) {
	t.Run("deleted version has DeletedTime set", func(t *testing.T) {
		kv := New(Config{MaxSecretVersions: 10})
		kv.Put("test/path", map[string]string{"key": "value"})

		beforeTime := time.Now()
		modified, err := kv.Delete("test/path", []int{1})
		afterTime := time.Now()

		if err != nil {
			t.Fatalf("Delete() error = %v", err)
		}

		if len(modified) != 1 {
			t.Fatalf("modified length = %d, want 1", len(modified))
		}

		secret := kv.data["test/path"]
		version := secret.Versions[1]

		if version.DeletedTime == nil {
			t.Error("DeletedTime should be set after delete")
		}

		// Verify DeletedTime is reasonable
		if version.DeletedTime.Before(beforeTime) || version.DeletedTime.After(afterTime) {
			t.Errorf("DeletedTime %v not between %v and %v",
				version.DeletedTime, beforeTime, afterTime)
		}
	})

	t.Run("second delete of same version returns empty modified", func(t *testing.T) {
		kv := New(Config{MaxSecretVersions: 10})
		kv.Put("test/path", map[string]string{"key": "value"})

		// First delete
		modified1, err := kv.Delete("test/path", []int{1})
		if err != nil {
			t.Fatalf("First Delete() error = %v", err)
		}
		if len(modified1) != 1 {
			t.Errorf("First delete modified = %d, want 1", len(modified1))
		}

		// Second delete (idempotent)
		modified2, err := kv.Delete("test/path", []int{1})
		if err != nil {
			t.Fatalf("Second Delete() error = %v", err)
		}
		if len(modified2) != 0 {
			t.Errorf("Second delete modified = %d, want 0 (idempotent)", len(modified2))
		}
	})

	t.Run("data still accessible via GetRawSecret after delete", func(t *testing.T) {
		kv := New(Config{MaxSecretVersions: 10})
		kv.Put("test/path", map[string]string{"key": "value"})

		_, err := kv.Delete("test/path", []int{1})
		if err != nil {
			t.Fatalf("Delete() error = %v", err)
		}

		// GetRawSecret should still return the data
		secret, err := kv.GetRawSecret("test/path")
		if err != nil {
			t.Errorf("GetRawSecret() error = %v", err)
		}
		if secret == nil {
			t.Error("GetRawSecret() should return secret even if deleted")
		}
	})
}
