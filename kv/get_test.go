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

func TestKV_Get(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *KV
		path    string
		version int
		want    map[string]string
		wantErr error
	}{
		{
			name: "non_existent_path",
			setup: func() *KV {
				return &KV{
					data: make(map[string]*Value),
				}
			},
			path:    "non/existent/path",
			version: 0,
			want:    nil,
			wantErr: sdkErrors.ErrEntityNotFound,
		},
		{
			name: "get_current_version",
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
								"key": "current_value",
							},
							Version: 1,
						},
					},
				}
				return kv
			},
			path:    "test/path",
			version: 0,
			want: map[string]string{
				"key": "current_value",
			},
			wantErr: nil,
		},
		{
			name: "get_specific_version",
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
								"key": "old_value",
							},
							Version: 1,
						},
						2: {
							Data: map[string]string{
								"key": "current_value",
							},
							Version: 2,
						},
					},
				}
				return kv
			},
			path:    "test/path",
			version: 1,
			want: map[string]string{
				"key": "old_value",
			},
			wantErr: nil,
		},
		{
			name: "get_deleted_version",
			setup: func() *KV {
				deletedTime := time.Now()
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
								"key": "deleted_value",
							},
							Version:     1,
							DeletedTime: &deletedTime,
						},
					},
				}
				return kv
			},
			path:    "test/path",
			version: 1,
			want:    nil,
			wantErr: sdkErrors.ErrStoreItemSoftDeleted,
		},
		{
			name: "non_existent_version",
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
								"key": "value",
							},
							Version: 1,
						},
					},
				}
				return kv
			},
			path:    "test/path",
			version: 999,
			want:    nil,
			wantErr: sdkErrors.ErrStoreItemSoftDeleted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup()
			got, err := kv.Get(tt.path, tt.version)

			// Handle nil case explicitly to avoid typed nil vs untyped nil issues
			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("Get() error = %v, wantErr nil", err)
					return
				}
			} else if !errors.Is(err, tt.wantErr) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if len(got) != len(tt.want) {
					t.Errorf("Get() got = %v, want %v", got, tt.want)
					return
				}
				for k, v := range got {
					if tt.want[k] != v {
						t.Errorf("Get() got[%s] = %v, want[%s] = %v", k, v, k, tt.want[k])
					}
				}
			}
		})
	}
}

func TestKV_GetRawSecret(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *KV
		path    string
		want    *Value
		wantErr error
	}{
		{
			name: "non_existent_path",
			setup: func() *KV {
				return &KV{
					data: make(map[string]*Value),
				}
			},
			path:    "non/existent/path",
			want:    nil,
			wantErr: sdkErrors.ErrEntityNotFound,
		},
		{
			name: "existing_secret",
			setup: func() *KV {
				secret := &Value{
					Metadata: Metadata{
						CurrentVersion: 1,
					},
					Versions: map[int]Version{
						1: {
							Data: map[string]string{
								"key": "value",
							},
							Version: 1,
						},
					},
				}
				kv := &KV{
					data: make(map[string]*Value),
				}
				kv.data["test/path"] = secret
				return kv
			},
			path: "test/path",
			want: &Value{
				Metadata: Metadata{
					CurrentVersion: 1,
				},
				Versions: map[int]Version{
					1: {
						Data: map[string]string{
							"key": "value",
						},
						Version: 1,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup()
			got, err := kv.GetRawSecret(tt.path)

			// Handle nil case explicitly to avoid typed nil vs untyped nil issues
			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("GetRawSecret() error = %v, wantErr nil", err)
					return
				}
			} else if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetRawSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if got.Metadata.CurrentVersion != tt.want.Metadata.CurrentVersion {
					t.Errorf("GetRawSecret() got CurrentVersion = %v, want %v",
						got.Metadata.CurrentVersion, tt.want.Metadata.CurrentVersion)
				}

				if len(got.Versions) != len(tt.want.Versions) {
					t.Errorf("GetRawSecret() got Versions length = %v, want %v",
						len(got.Versions), len(tt.want.Versions))
					return
				}

				for version, gotV := range got.Versions {
					wantV, exists := tt.want.Versions[version]
					if !exists {
						t.Errorf("GetRawSecret() unexpected version %v in result", version)
						continue
					}

					if gotV.Version != wantV.Version {
						t.Errorf("GetRawSecret() version %v: got Version = %v, want %v",
							version, gotV.Version, wantV.Version)
					}

					if len(gotV.Data) != len(wantV.Data) {
						t.Errorf("GetRawSecret() version %v: got Data length = %v, want %v",
							version, len(gotV.Data), len(wantV.Data))
						continue
					}

					for k, v := range gotV.Data {
						if wantV.Data[k] != v {
							t.Errorf("GetRawSecret() version %v: got Data[%s] = %v, want %v",
								version, k, v, wantV.Data[k])
						}
					}
				}
			}
		})
	}
}
