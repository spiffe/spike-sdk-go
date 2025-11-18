//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import (
	"errors"
	"testing"

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup()
			err := kv.Delete(tt.path, tt.versions)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == nil {
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
