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

func TestKV_Undelete(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *KV
		path     string
		values   map[string]string
		versions []int
		wantErr  error
	}{
		{
			name: "undelete latest version if no versions specified",
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
							Data:        map[string]string{"key": "value"},
							Version:     1,
							DeletedTime: &time.Time{},
						},
					},
				}
				return kv
			},
			path:     "test/path",
			versions: []int{},
			wantErr:  nil,
		},
		{
			name: "undelete specific versions",
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
							Data:        map[string]string{"key": "value1"},
							Version:     1,
							DeletedTime: &time.Time{},
						},
						2: {
							Data:        map[string]string{"key": "value2"},
							Version:     2,
							DeletedTime: &time.Time{},
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
			name: "if secret does not exist",
			setup: func() *KV {
				return &KV{
					data:              make(map[string]*Value),
					maxSecretVersions: 10,
				}
			},
			path:     "path/undelete/notExist",
			versions: []int{1},
			values:   map[string]string{"key": "value"},
			wantErr:  sdkErrors.ErrEntityNotFound,
		},
		{
			name: "skip non-existent versions",
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
							Data:        map[string]string{"key": "value"},
							Version:     1,
							DeletedTime: &time.Time{},
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
			name: "skip non-existent versions",
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
							Data:        map[string]string{"key": "value"},
							Version:     1,
							DeletedTime: &time.Time{},
						},
					},
				}
				return kv
			},
			path:     "test/path",
			versions: []int{0},
			wantErr:  nil,
		},
		{
			name: "if version is 0 undelete current version",
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
							Data:        map[string]string{"key": "value"},
							Version:     1,
							DeletedTime: &time.Time{},
						},
					},
				}
				return kv
			},
			path:     "test/path",
			values:   map[string]string{"key": "value"},
			versions: []int{0},
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup()
			_, err := kv.Undelete(tt.path, tt.versions)

			// Handle nil case explicitly to avoid typed nil vs untyped nil issues
			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("Undelete() error = %v, wantErr nil", err)
					return
				}
			} else if !errors.Is(err, tt.wantErr) {
				t.Errorf("Undelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				secret, exist := kv.data[tt.path]
				if !exist {
					t.Errorf("Secret should exist at path %s", tt.path)
					return
				}

				for _, version := range tt.versions {
					if version == 0 {
						version = secret.Metadata.CurrentVersion
					}
					if v, exist := secret.Versions[version]; exist {
						if v.DeletedTime != nil {
							t.Errorf("Version %d should not be deleted", version)
						}
					}
				}
			}
		})
	}
}
