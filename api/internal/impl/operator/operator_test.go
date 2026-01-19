//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package operator

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestRecover_NilSource tests that Recover returns error for nil X509Source
func TestRecover_NilSource(t *testing.T) {
	shards, err := Recover(context.Background(), nil)

	assert.Nil(t, shards)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestRestore_NilSource tests that Restore returns error for nil X509Source
func TestRestore_NilSource(t *testing.T) {
	shardValue := &[32]byte{}
	status, err := Restore(context.Background(), nil, 1, shardValue)

	assert.Nil(t, status)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestRestore_NilShardValue tests that Restore handles nil shard value
func TestRestore_NilShardValue(t *testing.T) {
	status, err := Restore(context.Background(), nil, 1, nil)

	assert.Nil(t, status)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestRecoverRequestMarshaling tests that RecoverRequest marshals correctly
func TestRecoverRequestMarshaling(t *testing.T) {
	request := reqres.RecoverRequest{}

	jsonData, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Verify it's valid JSON by unmarshaling back
	var result map[string]interface{}
	unmarshalErr := json.Unmarshal(jsonData, &result)
	assert.NoError(t, unmarshalErr)
}

// TestRestoreRequestMarshaling tests that RestoreRequest marshals correctly
func TestRestoreRequestMarshaling(t *testing.T) {
	shardValue := &[32]byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
	}

	tests := []struct {
		name    string
		request reqres.RestoreRequest
		wantErr bool
	}{
		{
			name: "ValidRequest",
			request: reqres.RestoreRequest{
				ID:    1,
				Shard: shardValue,
			},
			wantErr: false,
		},
		{
			name: "ZeroID",
			request: reqres.RestoreRequest{
				ID:    0,
				Shard: shardValue,
			},
			wantErr: false,
		},
		{
			name: "NilShard",
			request: reqres.RestoreRequest{
				ID:    1,
				Shard: nil,
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
				var unmarshaled reqres.RestoreRequest
				unmarshalErr := json.Unmarshal(jsonData, &unmarshaled)
				assert.NoError(t, unmarshalErr)
				assert.Equal(t, tt.request.ID, unmarshaled.ID)
			}
		})
	}
}

// TestRecoverResponseUnmarshaling tests that RecoverResponse unmarshals correctly
func TestRecoverResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
	}{
		{
			name:     "EmptyShards",
			jsonData: `{"shards":{}}`,
			wantErr:  false,
		},
		{
			name:     "WithShards",
			jsonData: `{"shards":{"0":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32]}}`,
			wantErr:  false,
		},
		{
			name:     "NullShards",
			jsonData: `{"shards":null}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response reqres.RecoverResponse
			err := json.Unmarshal([]byte(tt.jsonData), &response)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestRestoreResponseUnmarshaling tests that RestoreResponse unmarshals correctly
func TestRestoreResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
	}{
		{
			name:     "InProgress",
			jsonData: `{"collected":2,"remaining":3,"restored":false}`,
			wantErr:  false,
		},
		{
			name:     "Completed",
			jsonData: `{"collected":5,"remaining":0,"restored":true}`,
			wantErr:  false,
		},
		{
			name:     "Initial",
			jsonData: `{"collected":0,"remaining":5,"restored":false}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response reqres.RestoreResponse
			err := json.Unmarshal([]byte(tt.jsonData), &response)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestRestorationStatusConstruction tests that RestorationStatus is constructed correctly
func TestRestorationStatusConstruction(t *testing.T) {
	tests := []struct {
		name     string
		status   data.RestorationStatus
		expected data.RestorationStatus
	}{
		{
			name: "InProgress",
			status: data.RestorationStatus{
				ShardsCollected: 2,
				ShardsRemaining: 3,
				Restored:        false,
			},
			expected: data.RestorationStatus{
				ShardsCollected: 2,
				ShardsRemaining: 3,
				Restored:        false,
			},
		},
		{
			name: "Completed",
			status: data.RestorationStatus{
				ShardsCollected: 5,
				ShardsRemaining: 0,
				Restored:        true,
			},
			expected: data.RestorationStatus{
				ShardsCollected: 5,
				ShardsRemaining: 0,
				Restored:        true,
			},
		},
		{
			name: "Initial",
			status: data.RestorationStatus{
				ShardsCollected: 0,
				ShardsRemaining: 5,
				Restored:        false,
			},
			expected: data.RestorationStatus{
				ShardsCollected: 0,
				ShardsRemaining: 5,
				Restored:        false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected.ShardsCollected, tt.status.ShardsCollected)
			assert.Equal(t, tt.expected.ShardsRemaining, tt.status.ShardsRemaining)
			assert.Equal(t, tt.expected.Restored, tt.status.Restored)

			// Verify JSON marshaling works
			jsonData, err := json.Marshal(tt.status)
			assert.NoError(t, err)
			assert.NotEmpty(t, jsonData)

			// Verify unmarshaling works
			var unmarshaled data.RestorationStatus
			err = json.Unmarshal(jsonData, &unmarshaled)
			assert.NoError(t, err)
			assert.Equal(t, tt.status, unmarshaled)
		})
	}
}

// TestRecoverResponseMethods tests the response builder methods
func TestRecoverResponseMethods(t *testing.T) {
	tests := []struct {
		name         string
		method       func(reqres.RecoverResponse) reqres.RecoverResponse
		expectedCode sdkErrors.ErrorCode
	}{
		{
			name:         "Success",
			method:       reqres.RecoverResponse.Success,
			expectedCode: "",
		},
		{
			name:         "BadRequest",
			method:       reqres.RecoverResponse.BadRequest,
			expectedCode: sdkErrors.ErrAPIBadRequest.Code,
		},
		{
			name:         "Unauthorized",
			method:       reqres.RecoverResponse.Unauthorized,
			expectedCode: sdkErrors.ErrAccessUnauthorized.Code,
		},
		{
			name:         "Internal",
			method:       reqres.RecoverResponse.Internal,
			expectedCode: sdkErrors.ErrAPIInternal.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := reqres.RecoverResponse{}
			result := tt.method(response)
			assert.Equal(t, tt.expectedCode, result.ErrorCode())
		})
	}
}

// TestRestoreResponseMethods tests the response builder methods
func TestRestoreResponseMethods(t *testing.T) {
	tests := []struct {
		name         string
		method       func(reqres.RestoreResponse) reqres.RestoreResponse
		expectedCode sdkErrors.ErrorCode
	}{
		{
			name:         "Success",
			method:       reqres.RestoreResponse.Success,
			expectedCode: "",
		},
		{
			name:         "BadRequest",
			method:       reqres.RestoreResponse.BadRequest,
			expectedCode: sdkErrors.ErrAPIBadRequest.Code,
		},
		{
			name:         "Unauthorized",
			method:       reqres.RestoreResponse.Unauthorized,
			expectedCode: sdkErrors.ErrAccessUnauthorized.Code,
		},
		{
			name:         "Internal",
			method:       reqres.RestoreResponse.Internal,
			expectedCode: sdkErrors.ErrAPIInternal.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := reqres.RestoreResponse{}
			result := tt.method(response)
			assert.Equal(t, tt.expectedCode, result.ErrorCode())
		})
	}
}

// TestShardByteArrayHandling tests handling of 32-byte shard arrays
func TestShardByteArrayHandling(t *testing.T) {
	// Create a shard with known values
	shard := &[32]byte{}
	for i := 0; i < 32; i++ {
		shard[i] = byte(i + 1)
	}

	// Test JSON marshaling of shard
	request := reqres.RestoreRequest{
		ID:    1,
		Shard: shard,
	}

	jsonData, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Test unmarshaling
	var unmarshaled reqres.RestoreRequest
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, 1, unmarshaled.ID)
	assert.NotNil(t, unmarshaled.Shard)

	// Verify shard values
	for i := 0; i < 32; i++ {
		assert.Equal(t, byte(i+1), unmarshaled.Shard[i])
	}
}

// TestRecoverShardMapHandling tests handling of shard maps in RecoverResponse
func TestRecoverShardMapHandling(t *testing.T) {
	// Create multiple shards
	shards := make(map[int]*[32]byte)
	for i := 0; i < 3; i++ {
		shard := &[32]byte{}
		for j := 0; j < 32; j++ {
			shard[j] = byte(i*32 + j)
		}
		shards[i] = shard
	}

	response := reqres.RecoverResponse{
		Shards: shards,
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Test unmarshaling
	var unmarshaled reqres.RecoverResponse
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(unmarshaled.Shards))

	// Verify each shard
	for i := 0; i < 3; i++ {
		assert.NotNil(t, unmarshaled.Shards[i])
		for j := 0; j < 32; j++ {
			assert.Equal(t, byte(i*32+j), unmarshaled.Shards[i][j])
		}
	}
}
