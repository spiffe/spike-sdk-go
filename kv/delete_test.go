//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	"errors"
	"testing"
	"time"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

func TestKV_Delete(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *KV
		path     string
		versions []int
		wantErr  error
	}{
		{
			name: "non_existent_path",
			setup: func() *KV {
				return &KV{
					data: make(map[string]*Value),
				}
			},
			path:     "non/existent/path",
			versions: nil,
			wantErr:  sdkErrors.ErrEntityNotFound,
		},
		{
			name: "delete_current_version_no_versions_specified",
			setup: func() *KV {
				kv := &KV{
					data: make(map[string]*Value),
				}
				kv.data["test/path"] = &Value{
					Metadata: Metadata{
						CurrentVersion: 1,
					},
					Versions: map[int]Version{
						1: {
							Data: map[string]string{
								"key": "test_value",
							},
						},
					},
				}
				return kv
			},
			path:     "test/path",
			versions: nil,
			wantErr:  nil,
		},
		{
			name: "delete_specific_versions",
			setup: func() *KV {
				kv := &KV{
					data: make(map[string]*Value),
				}
				kv.data["test/path"] = &Value{
					Metadata: Metadata{
						CurrentVersion: 2,
					},
					Versions: map[int]Version{
						1: {
							Data: map[string]string{
								"key": "value1",
							},
						},
						2: {
							Data: map[string]string{
								"key": "value2",
							},
						},
					},
				}
				return kv
			},
			path:     "test/path",
			versions: []int{1, 2},
			wantErr:  nil,
		},
		{
			name: "delete_version_0_when_current_version_does_not_exist",
			setup: func() *KV {
				kv := &KV{
					data: make(map[string]*Value),
				}
				// Create secret with CurrentVersion=2 but version 2 doesn't exist
				kv.data["test/path"] = &Value{
					Metadata: Metadata{
						CurrentVersion: 2,
						OldestVersion:  1,
					},
					Versions: map[int]Version{
						1: {
							Data: map[string]string{"key": "value1"},
						},
						// Version 2 missing!
					},
				}
				return kv
			},
			path:     "test/path",
			versions: []int{0}, // 0 means current version
			wantErr:  nil,
		},
		{
			name: "delete_already_deleted_version_idempotent",
			setup: func() *KV {
				kv := &KV{
					data: make(map[string]*Value),
				}
				deletedTime := time.Now()
				kv.data["test/path"] = &Value{
					Metadata: Metadata{
						CurrentVersion: 1,
						OldestVersion:  1,
					},
					Versions: map[int]Version{
						1: {
							Data:        map[string]string{"key": "value"},
							DeletedTime: &deletedTime, // Already deleted
						},
					},
				}
				return kv
			},
			path:     "test/path",
			versions: []int{1},
			wantErr:  nil,
		},
		{
			name: "delete_non_existent_version_silent_skip",
			setup: func() *KV {
				kv := &KV{
					data: make(map[string]*Value),
				}
				kv.data["test/path"] = &Value{
					Metadata: Metadata{
						CurrentVersion: 1,
						OldestVersion:  1,
					},
					Versions: map[int]Version{
						1: {
							Data: map[string]string{"key": "value"},
						},
					},
				}
				return kv
			},
			path:     "test/path",
			versions: []int{99}, // Version doesn't exist
			wantErr:  nil,
		},
		{
			name: "delete_empty_versions_when_current_version_missing",
			setup: func() *KV {
				kv := &KV{
					data: make(map[string]*Value),
				}
				// Create secret with CurrentVersion but that version doesn't exist
				kv.data["test/path"] = &Value{
					Metadata: Metadata{
						CurrentVersion: 5,
						OldestVersion:  1,
					},
					Versions: map[int]Version{
						// Version 5 is missing
					},
				}
				return kv
			},
			path:     "test/path",
			versions: []int{}, // Empty means delete current
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup()
			_, err := kv.Delete(tt.path, tt.versions)

			// Handle nil case explicitly to avoid typed nil vs untyped nil issues
			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("Delete() error = %v, wantErr nil", err)
					return
				}
			} else if !errors.Is(err, tt.wantErr) {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				secret, exists := kv.data[tt.path]
				if !exists {
					t.Errorf("Value should still exist after deletion")
					return
				}

				if len(tt.versions) == 0 {
					cv := secret.Metadata.CurrentVersion
					if v, exists := secret.Versions[cv]; exists {
						if v.DeletedTime == nil {
							t.Errorf("Current version should be marked as deleted")
						}
					}
				} else {
					for _, version := range tt.versions {
						if version == 0 {
							version = secret.Metadata.CurrentVersion
						}
						if v, exists := secret.Versions[version]; exists {
							if v.DeletedTime == nil {
								t.Errorf("Version %d should be marked as deleted", version)
							}
						}
					}
				}
			}
		})
	}
}
