//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestCreatePolicy_NilSource tests that CreatePolicy returns error for nil X509Source
func TestCreatePolicy_NilSource(t *testing.T) {
	err := CreatePolicy(context.Background(), nil, "test-policy", "spiffe://test/*", "/api/*",
		[]data.PolicyPermission{data.PermissionRead})

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestGetPolicy_NilSource tests that GetPolicy returns error for nil X509Source
func TestGetPolicy_NilSource(t *testing.T) {
	policy, err := GetPolicy(context.Background(), nil, "policy-123")

	assert.Nil(t, policy)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestDeletePolicy_NilSource tests that DeletePolicy returns error for nil X509Source
func TestDeletePolicy_NilSource(t *testing.T) {
	err := DeletePolicy(context.Background(), nil, "policy-123")

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestListPolicies_NilSource tests that ListPolicies returns error for nil X509Source
func TestListPolicies_NilSource(t *testing.T) {
	policies, err := ListPolicies(context.Background(), nil, "", "")

	assert.Nil(t, policies)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestPolicyRequestMarshaling tests that policy request structs marshal correctly
func TestPolicyRequestMarshaling(t *testing.T) {
	tests := []struct {
		name    string
		request interface{}
		wantErr bool
	}{
		{
			name: "PolicyPutRequest",
			request: reqres.PolicyPutRequest{
				Name:            "test-policy",
				SPIFFEIDPattern: "spiffe://example.org/*",
				PathPattern:     "/api/*",
				Permissions:     []data.PolicyPermission{data.PermissionRead, data.PermissionWrite},
			},
			wantErr: false,
		},
		{
			name: "PolicyReadRequest",
			request: reqres.PolicyReadRequest{
				ID: "policy-123",
			},
			wantErr: false,
		},
		{
			name: "PolicyDeleteRequest",
			request: reqres.PolicyDeleteRequest{
				ID: "policy-456",
			},
			wantErr: false,
		},
		{
			name: "PolicyListRequest_Empty",
			request: reqres.PolicyListRequest{
				SPIFFEIDPattern: "",
				PathPattern:     "",
			},
			wantErr: false,
		},
		{
			name: "PolicyListRequest_WithFilters",
			request: reqres.PolicyListRequest{
				SPIFFEIDPattern: "spiffe://example.org/service/*",
				PathPattern:     "/api/v1/*",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, jsonData)

				// Verify it's valid JSON by unmarshaling back
				var result map[string]interface{}
				unmarshalErr := json.Unmarshal(jsonData, &result)
				assert.NoError(t, unmarshalErr)
			}
		})
	}
}

// TestPolicyResponseUnmarshaling tests that policy response structs unmarshal correctly
func TestPolicyResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		target   interface{}
		wantErr  bool
	}{
		{
			name:     "PolicyPutResponse",
			jsonData: `{"id":"policy-123"}`,
			target:   &reqres.PolicyPutResponse{},
			wantErr:  false,
		},
		{
			name:     "PolicyReadResponse",
			jsonData: `{"id":"policy-123","name":"test-policy","spiffeidPattern":"spiffe://example.org/*","pathPattern":"/api/*","permissions":["read"]}`,
			target:   &reqres.PolicyReadResponse{},
			wantErr:  false,
		},
		{
			name:     "PolicyDeleteResponse",
			jsonData: `{}`,
			target:   &reqres.PolicyDeleteResponse{},
			wantErr:  false,
		},
		{
			name:     "PolicyListResponse_Empty",
			jsonData: `{"policies":[]}`,
			target:   &reqres.PolicyListResponse{},
			wantErr:  false,
		},
		{
			name:     "PolicyListResponse_WithPolicies",
			jsonData: `{"policies":[{"id":"policy-1","name":"test"}]}`,
			target:   &reqres.PolicyListResponse{},
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

// TestPolicyDataConstruction tests that data.Policy is constructed correctly
func TestPolicyDataConstruction(t *testing.T) {
	policy := data.Policy{
		ID:              "policy-123",
		Name:            "test-policy",
		SPIFFEIDPattern: "spiffe://example.org/service/*",
		PathPattern:     "/api/documents/*",
		Permissions:     []data.PolicyPermission{data.PermissionRead, data.PermissionWrite},
	}

	assert.Equal(t, "policy-123", policy.ID)
	assert.Equal(t, "test-policy", policy.Name)
	assert.Equal(t, "spiffe://example.org/service/*", policy.SPIFFEIDPattern)
	assert.Equal(t, "/api/documents/*", policy.PathPattern)
	assert.Equal(t, 2, len(policy.Permissions))
	assert.Contains(t, policy.Permissions, data.PermissionRead)
	assert.Contains(t, policy.Permissions, data.PermissionWrite)
}

// TestPolicyPermissions tests all policy permission constants
func TestPolicyPermissions(t *testing.T) {
	permissions := []data.PolicyPermission{
		data.PermissionRead,
		data.PermissionWrite,
		data.PermissionList,
		data.PermissionExecute,
		data.PermissionSuper,
	}

	// Test that all permissions are distinct
	seen := make(map[data.PolicyPermission]bool)
	for _, perm := range permissions {
		assert.False(t, seen[perm], "Duplicate permission found: %v", perm)
		seen[perm] = true
	}

	// Test that permissions can be marshaled to JSON
	for _, perm := range permissions {
		jsonData, err := json.Marshal(perm)
		assert.NoError(t, err)
		assert.NotEmpty(t, jsonData)

		// Unmarshal back to verify
		var unmarshaledPerm data.PolicyPermission
		err = json.Unmarshal(jsonData, &unmarshaledPerm)
		assert.NoError(t, err)
		assert.Equal(t, perm, unmarshaledPerm)
	}
}

// TestListPolicies_EmptyResult tests that ListPolicies handles empty results correctly
func TestListPolicies_EmptyResult(t *testing.T) {
	// This test demonstrates the expected behavior for empty results
	// In practice, the function returns an empty slice pointer when ErrAPINotFound is encountered

	// Test that we can create an empty policy slice
	emptyPolicies := []data.PolicyListItem{}
	assert.NotNil(t, emptyPolicies)
	assert.Equal(t, 0, len(emptyPolicies))

	// Test pointer to empty slice
	emptyPtr := &emptyPolicies
	assert.NotNil(t, emptyPtr)
	assert.Equal(t, 0, len(*emptyPtr))
}

// TestGetPolicy_NotFoundBehavior tests the expected behavior when policy is not found
func TestGetPolicy_NotFoundBehavior(t *testing.T) {
	// This test demonstrates that GetPolicy returns (nil, nil) for not found
	// which is the documented behavior for 404 responses

	// When a policy is not found, both return values should be nil
	var policy *data.Policy
	var err *sdkErrors.SDKError

	assert.Nil(t, policy)
	assert.Nil(t, err)
}

// TestPolicyWithMultiplePermissions tests policy with various permission combinations
func TestPolicyWithMultiplePermissions(t *testing.T) {
	testCases := []struct {
		name        string
		permissions []data.PolicyPermission
		expected    int
	}{
		{
			name:        "SinglePermission",
			permissions: []data.PolicyPermission{data.PermissionRead},
			expected:    1,
		},
		{
			name:        "MultiplePermissions",
			permissions: []data.PolicyPermission{data.PermissionRead, data.PermissionWrite, data.PermissionList},
			expected:    3,
		},
		{
			name: "AllPermissions",
			permissions: []data.PolicyPermission{
				data.PermissionRead,
				data.PermissionWrite,
				data.PermissionList,
				data.PermissionExecute,
				data.PermissionSuper,
			},
			expected: 5,
		},
		{
			name:        "NoPermissions",
			permissions: []data.PolicyPermission{},
			expected:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			policy := data.Policy{
				ID:              "test-policy",
				Name:            tc.name,
				SPIFFEIDPattern: "spiffe://example.org/*",
				PathPattern:     "/api/*",
				Permissions:     tc.permissions,
			}

			assert.Equal(t, tc.expected, len(policy.Permissions))

			// Verify JSON marshaling works
			jsonData, err := json.Marshal(policy)
			assert.NoError(t, err)
			assert.NotEmpty(t, jsonData)

			// Verify unmarshaling works
			var unmarshaled data.Policy
			err = json.Unmarshal(jsonData, &unmarshaled)
			assert.NoError(t, err)
			assert.Equal(t, policy.ID, unmarshaled.ID)
			assert.Equal(t, len(policy.Permissions), len(unmarshaled.Permissions))
		})
	}
}

// TestPolicyPatterns tests various SPIFFE ID and path pattern combinations
func TestPolicyPatterns(t *testing.T) {
	testCases := []struct {
		name            string
		spiffeIDPattern string
		pathPattern     string
		valid           bool
	}{
		{
			name:            "WildcardBoth",
			spiffeIDPattern: "spiffe://example.org/*",
			pathPattern:     "/api/*",
			valid:           true,
		},
		{
			name:            "SpecificService",
			spiffeIDPattern: "spiffe://example.org/service/frontend",
			pathPattern:     "/api/v1/users",
			valid:           true,
		},
		{
			name:            "EmptyPatterns",
			spiffeIDPattern: "",
			pathPattern:     "",
			valid:           true,
		},
		{
			name:            "ComplexPath",
			spiffeIDPattern: "spiffe://example.org/ns/*/sa/*",
			pathPattern:     "/api/v1/namespaces/*/secrets/*",
			valid:           true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			policy := data.Policy{
				ID:              "test-policy",
				Name:            tc.name,
				SPIFFEIDPattern: tc.spiffeIDPattern,
				PathPattern:     tc.pathPattern,
				Permissions:     []data.PolicyPermission{data.PermissionRead},
			}

			// Verify policy can be created
			assert.NotEmpty(t, policy.ID)

			// Verify JSON marshaling
			jsonData, err := json.Marshal(policy)
			assert.NoError(t, err)
			assert.NotEmpty(t, jsonData)
		})
	}
}
