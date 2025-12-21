//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package bootstrap

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestVerify_NilSource tests that Verify returns error for nil X509Source
func TestVerify_NilSource(t *testing.T) {
	randomText := "test random text"
	nonce := []byte("test nonce")
	ciphertext := []byte("test ciphertext")

	err := Verify(context.Background(), nil, randomText, nonce, ciphertext)

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestVerify_EmptyRandomText tests Verify with empty random text
func TestVerify_EmptyRandomText(t *testing.T) {
	err := Verify(context.Background(), nil, "", []byte("nonce"), []byte("ciphertext"))

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestVerify_NilNonce tests Verify with nil nonce
func TestVerify_NilNonce(t *testing.T) {
	err := Verify(context.Background(), nil, "random text", nil, []byte("ciphertext"))

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestVerify_NilCiphertext tests Verify with nil ciphertext
func TestVerify_NilCiphertext(t *testing.T) {
	err := Verify(context.Background(), nil, "random text", []byte("nonce"), nil)

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestBootstrapVerifyRequestMarshaling tests that BootstrapVerifyRequest marshals correctly
func TestBootstrapVerifyRequestMarshaling(t *testing.T) {
	tests := []struct {
		name    string
		request reqres.BootstrapVerifyRequest
		wantErr bool
	}{
		{
			name: "ValidRequest",
			request: reqres.BootstrapVerifyRequest{
				Nonce:      []byte("test-nonce-12345"),
				Ciphertext: []byte("encrypted-data-here"),
			},
			wantErr: false,
		},
		{
			name: "EmptyNonce",
			request: reqres.BootstrapVerifyRequest{
				Nonce:      []byte{},
				Ciphertext: []byte("encrypted-data"),
			},
			wantErr: false,
		},
		{
			name: "EmptyCiphertext",
			request: reqres.BootstrapVerifyRequest{
				Nonce:      []byte("nonce"),
				Ciphertext: []byte{},
			},
			wantErr: false,
		},
		{
			name: "NilValues",
			request: reqres.BootstrapVerifyRequest{
				Nonce:      nil,
				Ciphertext: nil,
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
				var unmarshaled reqres.BootstrapVerifyRequest
				unmarshalErr := json.Unmarshal(jsonData, &unmarshaled)
				assert.NoError(t, unmarshalErr)
			}
		})
	}
}

// TestBootstrapVerifyResponseUnmarshaling tests that BootstrapVerifyResponse unmarshals correctly
func TestBootstrapVerifyResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
	}{
		{
			name:     "ValidHash",
			jsonData: `{"hash":"abc123def456","err":""}`,
			wantErr:  false,
		},
		{
			name:     "EmptyHash",
			jsonData: `{"hash":"","err":""}`,
			wantErr:  false,
		},
		{
			name:     "WithError",
			jsonData: `{"hash":"","err":"api_internal"}`,
			wantErr:  false,
		},
		{
			name:     "SHA256Hash",
			jsonData: `{"hash":"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response reqres.BootstrapVerifyResponse
			err := json.Unmarshal([]byte(tt.jsonData), &response)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestShardPutRequestMarshaling tests that ShardPutRequest marshals correctly
func TestShardPutRequestMarshaling(t *testing.T) {
	shard := &[32]byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
	}

	tests := []struct {
		name    string
		request reqres.ShardPutRequest
		wantErr bool
	}{
		{
			name: "ValidShard",
			request: reqres.ShardPutRequest{
				Shard: shard,
			},
			wantErr: false,
		},
		{
			name: "NilShard",
			request: reqres.ShardPutRequest{
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
				var unmarshaled reqres.ShardPutRequest
				unmarshalErr := json.Unmarshal(jsonData, &unmarshaled)
				assert.NoError(t, unmarshalErr)
			}
		})
	}
}

// TestShardPutResponseUnmarshaling tests that ShardPutResponse unmarshals correctly
func TestShardPutResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
	}{
		{
			name:     "Success",
			jsonData: `{"err":""}`,
			wantErr:  false,
		},
		{
			name:     "WithError",
			jsonData: `{"err":"api_bad_request"}`,
			wantErr:  false,
		},
		{
			name:     "EmptyObject",
			jsonData: `{}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response reqres.ShardPutResponse
			err := json.Unmarshal([]byte(tt.jsonData), &response)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestShardGetRequestMarshaling tests that ShardGetRequest marshals correctly
func TestShardGetRequestMarshaling(t *testing.T) {
	request := reqres.ShardGetRequest{}

	jsonData, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Verify it's valid JSON by unmarshaling back
	var result map[string]interface{}
	unmarshalErr := json.Unmarshal(jsonData, &result)
	assert.NoError(t, unmarshalErr)
}

// TestShardGetResponseUnmarshaling tests that ShardGetResponse unmarshals correctly
func TestShardGetResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
	}{
		{
			name:     "WithShard",
			jsonData: `{"shard":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32]}`,
			wantErr:  false,
		},
		{
			name:     "NullShard",
			jsonData: `{"shard":null}`,
			wantErr:  false,
		},
		{
			name:     "WithError",
			jsonData: `{"shard":null,"err":"api_not_found"}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response reqres.ShardGetResponse
			err := json.Unmarshal([]byte(tt.jsonData), &response)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestBootstrapVerifyResponseMethods tests the response builder methods
func TestBootstrapVerifyResponseMethods(t *testing.T) {
	tests := []struct {
		name         string
		method       func(reqres.BootstrapVerifyResponse) reqres.BootstrapVerifyResponse
		expectedCode sdkErrors.ErrorCode
	}{
		{
			name:         "Success",
			method:       reqres.BootstrapVerifyResponse.Success,
			expectedCode: "",
		},
		{
			name:         "BadRequest",
			method:       reqres.BootstrapVerifyResponse.BadRequest,
			expectedCode: sdkErrors.ErrAPIBadRequest.Code,
		},
		{
			name:         "Unauthorized",
			method:       reqres.BootstrapVerifyResponse.Unauthorized,
			expectedCode: sdkErrors.ErrAccessUnauthorized.Code,
		},
		{
			name:         "Internal",
			method:       reqres.BootstrapVerifyResponse.Internal,
			expectedCode: sdkErrors.ErrAPIInternal.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := reqres.BootstrapVerifyResponse{}
			result := tt.method(response)
			assert.Equal(t, tt.expectedCode, result.ErrorCode())
		})
	}
}

// TestShardPutResponseMethods tests the response builder methods
func TestShardPutResponseMethods(t *testing.T) {
	tests := []struct {
		name         string
		method       func(reqres.ShardPutResponse) reqres.ShardPutResponse
		expectedCode sdkErrors.ErrorCode
	}{
		{
			name:         "Success",
			method:       reqres.ShardPutResponse.Success,
			expectedCode: "",
		},
		{
			name:         "BadRequest",
			method:       reqres.ShardPutResponse.BadRequest,
			expectedCode: sdkErrors.ErrAPIBadRequest.Code,
		},
		{
			name:         "Unauthorized",
			method:       reqres.ShardPutResponse.Unauthorized,
			expectedCode: sdkErrors.ErrAccessUnauthorized.Code,
		},
		{
			name:         "Internal",
			method:       reqres.ShardPutResponse.Internal,
			expectedCode: sdkErrors.ErrAPIInternal.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := reqres.ShardPutResponse{}
			result := tt.method(response)
			assert.Equal(t, tt.expectedCode, result.ErrorCode())
		})
	}
}

// TestShardGetResponseMethods tests the response builder methods
func TestShardGetResponseMethods(t *testing.T) {
	tests := []struct {
		name         string
		method       func(reqres.ShardGetResponse) reqres.ShardGetResponse
		expectedCode sdkErrors.ErrorCode
	}{
		{
			name:         "Success",
			method:       reqres.ShardGetResponse.Success,
			expectedCode: "",
		},
		{
			name:         "NotFound",
			method:       reqres.ShardGetResponse.NotFound,
			expectedCode: sdkErrors.ErrAPINotFound.Code,
		},
		{
			name:         "BadRequest",
			method:       reqres.ShardGetResponse.BadRequest,
			expectedCode: sdkErrors.ErrAPIBadRequest.Code,
		},
		{
			name:         "Unauthorized",
			method:       reqres.ShardGetResponse.Unauthorized,
			expectedCode: sdkErrors.ErrAccessUnauthorized.Code,
		},
		{
			name:         "Internal",
			method:       reqres.ShardGetResponse.Internal,
			expectedCode: sdkErrors.ErrAPIInternal.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := reqres.ShardGetResponse{}
			result := tt.method(response)
			assert.Equal(t, tt.expectedCode, result.ErrorCode())
		})
	}
}

// TestShardByteArraySerialization tests 32-byte shard array serialization
func TestShardByteArraySerialization(t *testing.T) {
	// Create a shard with known values
	shard := &[32]byte{}
	for i := 0; i < 32; i++ {
		shard[i] = byte(i + 1)
	}

	// Test ShardPutRequest marshaling
	putRequest := reqres.ShardPutRequest{
		Shard: shard,
	}

	jsonData, err := json.Marshal(putRequest)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Test unmarshaling
	var unmarshaledPut reqres.ShardPutRequest
	err = json.Unmarshal(jsonData, &unmarshaledPut)
	assert.NoError(t, err)
	assert.NotNil(t, unmarshaledPut.Shard)

	// Verify shard values
	for i := 0; i < 32; i++ {
		assert.Equal(t, byte(i+1), unmarshaledPut.Shard[i])
	}

	// Test ShardGetResponse marshaling
	getResponse := reqres.ShardGetResponse{
		Shard: shard,
	}

	jsonData, err = json.Marshal(getResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Test unmarshaling
	var unmarshaledGet reqres.ShardGetResponse
	err = json.Unmarshal(jsonData, &unmarshaledGet)
	assert.NoError(t, err)
	assert.NotNil(t, unmarshaledGet.Shard)

	// Verify shard values
	for i := 0; i < 32; i++ {
		assert.Equal(t, byte(i+1), unmarshaledGet.Shard[i])
	}
}

// TestBootstrapVerifyRequestByteArrays tests byte array handling in BootstrapVerifyRequest
func TestBootstrapVerifyRequestByteArrays(t *testing.T) {
	nonce := make([]byte, 12)
	for i := 0; i < 12; i++ {
		nonce[i] = byte(i)
	}

	ciphertext := make([]byte, 48)
	for i := 0; i < 48; i++ {
		ciphertext[i] = byte(i + 100)
	}

	request := reqres.BootstrapVerifyRequest{
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}

	// Test marshaling
	jsonData, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Test unmarshaling
	var unmarshaled reqres.BootstrapVerifyRequest
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)

	// Verify nonce
	assert.Equal(t, len(nonce), len(unmarshaled.Nonce))
	for i := 0; i < len(nonce); i++ {
		assert.Equal(t, nonce[i], unmarshaled.Nonce[i])
	}

	// Verify ciphertext
	assert.Equal(t, len(ciphertext), len(unmarshaled.Ciphertext))
	for i := 0; i < len(ciphertext); i++ {
		assert.Equal(t, ciphertext[i], unmarshaled.Ciphertext[i])
	}
}
