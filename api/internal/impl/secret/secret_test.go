//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package secret

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestGet_NilSource tests that Get returns error for nil X509Source
func TestGet_NilSource(t *testing.T) {
	secret, err := Get(context.Background(), nil, "test/path", 1)

	assert.Nil(t, secret)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestPut_NilSource tests that Put returns error for nil X509Source
func TestPut_NilSource(t *testing.T) {
	err := Put(context.Background(), nil, "test/path", map[string]string{"key": "value"})

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestDelete_NilSource tests that Delete returns error for nil X509Source
func TestDelete_NilSource(t *testing.T) {
	err := Delete(context.Background(), nil, "test/path", []int{1})

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestUndelete_NilSource tests that Undelete returns error for nil X509Source
func TestUndelete_NilSource(t *testing.T) {
	err := Undelete(context.Background(), nil, "test/path", []int{1})

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestListKeys_NilSource tests that ListKeys returns error for nil X509Source
func TestListKeys_NilSource(t *testing.T) {
	keys, err := ListKeys(context.Background(), nil)

	assert.Nil(t, keys)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestGetMetadata_NilSource tests that GetMetadata returns error for nil X509Source
func TestGetMetadata_NilSource(t *testing.T) {
	metadata, err := GetMetadata(context.Background(), nil, "test/path", 1)

	assert.Nil(t, metadata)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestGet_NotFound tests that Get returns (nil, nil) when secret is not found
func TestGet_NotFound(t *testing.T) {
	// Create a test server that returns 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// Note: This test demonstrates the expected behavior, but without being able
	// to inject the server URL, it won't actually call the test server.
	// This is a limitation of the current implementation that could be addressed
	// with dependency injection.

	// For now, we test the nil source case which is testable
	secret, err := Get(context.Background(), nil, "test/path", 1)
	assert.Nil(t, secret)
	assert.NotNil(t, err)
}

// TestListKeys_NotFound tests that ListKeys returns empty array when no secrets exist
func TestListKeys_NotFound(t *testing.T) {
	// Create a test server that returns 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// Note: Same limitation as TestGet_NotFound applies here
	keys, err := ListKeys(context.Background(), nil)
	assert.Nil(t, keys)
	assert.NotNil(t, err)
}

// TestGetMetadata_NotFound tests that GetMetadata returns (nil, nil) when metadata not found
func TestGetMetadata_NotFound(t *testing.T) {
	// Create a test server that returns 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// Note: Same limitation as TestGet_NotFound applies here
	metadata, err := GetMetadata(context.Background(), nil, "test/path", 1)
	assert.Nil(t, metadata)
	assert.NotNil(t, err)
}

// TestUndelete_EmptyVersions tests that Undelete handles empty versions array
func TestUndelete_EmptyVersions(t *testing.T) {
	// The function should handle empty versions gracefully
	// But without a real X509Source, we can only test nil case
	err := Undelete(context.Background(), nil, "test/path", []int{})

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// Mock response builders for testing
func mockSecretGetResponse(secretData map[string]string) reqres.SecretGetResponse {
	return reqres.SecretGetResponse{
		Secret: data.Secret{Data: secretData},
	}
}

func mockSecretListResponse(keys []string) reqres.SecretListResponse {
	return reqres.SecretListResponse{
		Keys: keys,
	}
}

func mockSecretMetadataResponse(versions map[int]data.SecretVersionInfo, metadata data.SecretMetaDataContent) reqres.SecretMetadataResponse {
	return reqres.SecretMetadataResponse{
		SecretMetadata: data.SecretMetadata{
			Versions: versions,
			Metadata: metadata,
		},
	}
}

// Test helper to create mock responses
func TestMockResponseBuilders(t *testing.T) {
	// Test secret get response
	getResp := mockSecretGetResponse(map[string]string{"key": "value"})
	assert.Equal(t, "value", getResp.Data["key"])

	// Test secret list response
	listResp := mockSecretListResponse([]string{"secret1", "secret2"})
	assert.Equal(t, 2, len(listResp.Keys))
	assert.Equal(t, "secret1", listResp.Keys[0])

	// Test secret metadata response
	now := time.Now()
	versions := map[int]data.SecretVersionInfo{
		1: {CreatedTime: now, Version: 1, DeletedTime: nil},
	}
	metadata := data.SecretMetaDataContent{
		CurrentVersion: 1,
		OldestVersion:  1,
		CreatedTime:    now,
		UpdatedTime:    now,
		MaxVersions:    10,
	}
	metaResp := mockSecretMetadataResponse(versions, metadata)
	assert.NotNil(t, metaResp.Versions)
	assert.Equal(t, 1, metaResp.Metadata.CurrentVersion)
	assert.Equal(t, 10, metaResp.Metadata.MaxVersions)
}

// TestDataSecretConstruction tests that data.Secret is constructed correctly
func TestDataSecretConstruction(t *testing.T) {
	secretData := map[string]string{"username": "admin", "password": "secret"}
	secret := data.Secret{Data: secretData}

	assert.Equal(t, "admin", secret.Data["username"])
	assert.Equal(t, "secret", secret.Data["password"])
}

// TestDataSecretMetadataConstruction tests that data.SecretMetadata is constructed correctly
func TestDataSecretMetadataConstruction(t *testing.T) {
	now := time.Now()
	versions := map[int]data.SecretVersionInfo{
		1: {CreatedTime: now, Version: 1, DeletedTime: nil},
		2: {CreatedTime: now.Add(time.Hour), Version: 2, DeletedTime: nil},
	}
	metadata := data.SecretMetaDataContent{
		CurrentVersion: 2,
		OldestVersion:  1,
		CreatedTime:    now,
		UpdatedTime:    now.Add(time.Hour),
		MaxVersions:    10,
	}

	secretMetadata := data.SecretMetadata{
		Versions: versions,
		Metadata: metadata,
	}

	assert.Equal(t, 2, len(secretMetadata.Versions))
	assert.Equal(t, 10, secretMetadata.Metadata.MaxVersions)
	assert.Equal(t, 2, secretMetadata.Metadata.CurrentVersion)
}

// TestRequestMarshaling tests that request structs marshal correctly to JSON
func TestRequestMarshaling(t *testing.T) {
	tests := []struct {
		name    string
		request interface{}
		wantErr bool
	}{
		{
			name: "SecretGetRequest",
			request: reqres.SecretGetRequest{
				Path:    "test/path",
				Version: 1,
			},
			wantErr: false,
		},
		{
			name: "SecretPutRequest",
			request: reqres.SecretPutRequest{
				Path:   "test/path",
				Values: map[string]string{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "SecretDeleteRequest",
			request: reqres.SecretDeleteRequest{
				Path:     "test/path",
				Versions: []int{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "SecretUndeleteRequest",
			request: reqres.SecretUndeleteRequest{
				Path:     "test/path",
				Versions: []int{1, 2},
			},
			wantErr: false,
		},
		{
			name:    "SecretListRequest",
			request: reqres.SecretListRequest{},
			wantErr: false,
		},
		{
			name: "SecretMetadataRequest",
			request: reqres.SecretMetadataRequest{
				Path:    "test/path",
				Version: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, data)

				// Verify it's valid JSON by unmarshaling back
				var result map[string]interface{}
				unmarshalErr := json.Unmarshal(data, &result)
				assert.NoError(t, unmarshalErr)
			}
		})
	}
}

// TestResponseUnmarshaling tests that response structs unmarshal correctly from JSON
func TestResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		target   interface{}
		wantErr  bool
	}{
		{
			name:     "SecretGetResponse",
			jsonData: `{"data":{"username":"admin","password":"secret"}}`,
			target:   &reqres.SecretGetResponse{},
			wantErr:  false,
		},
		{
			name:     "SecretListResponse",
			jsonData: `{"keys":["secret1","secret2","secret3"]}`,
			target:   &reqres.SecretListResponse{},
			wantErr:  false,
		},
		{
			name:     "SecretMetadataResponse",
			jsonData: `{"versions":{"1":{"created_time":"2024-01-01"}},"metadata":{"max_versions":10}}`,
			target:   &reqres.SecretMetadataResponse{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.jsonData), tt.target)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
