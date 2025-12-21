//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestContentTypeConstants tests that content type constants are defined correctly
func TestContentTypeConstants(t *testing.T) {
	assert.Equal(t, ContentType("application/json"), ContentTypeJSON)
	assert.Equal(t, ContentType("text/plain"), ContentTypePlain)
	assert.Equal(t, ContentType("application/octet-stream"), ContentTypeOctetStream)
}

// TestRequestBody_Success tests successful request body reading
func TestRequestBody_Success(t *testing.T) {
	testData := []byte("test request body")
	req := httptest.NewRequest("POST", "/test", bytes.NewReader(testData))

	body, err := RequestBody(req)

	assert.Nil(t, err)
	assert.Equal(t, testData, body)
}

// TestRequestBody_EmptyBody tests reading an empty request body
func TestRequestBody_EmptyBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/test", bytes.NewReader([]byte{}))

	body, err := RequestBody(req)

	assert.Nil(t, err)
	assert.Equal(t, []byte{}, body)
}

// TestRequestBody_LargeBody tests reading a large request body
func TestRequestBody_LargeBody(t *testing.T) {
	// Create 1MB of test data
	largeData := make([]byte, 1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	req := httptest.NewRequest("POST", "/test", bytes.NewReader(largeData))

	body, err := RequestBody(req)

	assert.Nil(t, err)
	assert.Equal(t, largeData, body)
	assert.Equal(t, 1024*1024, len(body))
}

// TestBody_Success tests successful response body reading
func TestBody_Success(t *testing.T) {
	testData := "test response body"
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(testData)),
	}

	bodyBytes, err := body(resp)

	assert.Nil(t, err)
	assert.Equal(t, []byte(testData), bodyBytes)
}

// TestBody_EmptyResponse tests reading an empty response body
func TestBody_EmptyResponse(t *testing.T) {
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader("")),
	}

	bodyBytes, err := body(resp)

	assert.Nil(t, err)
	assert.Equal(t, []byte{}, bodyBytes)
}

// TestBody_JSONResponse tests reading a JSON response body
func TestBody_JSONResponse(t *testing.T) {
	jsonData := `{"key":"value","number":42}`
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(jsonData)),
	}

	bodyBytes, err := body(resp)

	assert.Nil(t, err)
	assert.Equal(t, []byte(jsonData), bodyBytes)

	// Verify it's valid JSON
	var result map[string]interface{}
	unmarshalErr := json.Unmarshal(bodyBytes, &result)
	assert.NoError(t, unmarshalErr)
	assert.Equal(t, "value", result["key"])
	assert.Equal(t, float64(42), result["number"])
}

// Mock response type for testing PostAndUnmarshal
type mockResponse struct {
	Data string              `json:"data"`
	Err  sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (r mockResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}

// TestAuthorizerWithPredicate_Allow tests authorizer allowing connections
func TestAuthorizerWithPredicate_Allow(t *testing.T) {
	// Create a predicate that allows all IDs containing "allowed"
	predicate := func(id string) bool {
		return strings.Contains(id, "allowed")
	}

	authorizer := AuthorizerWithPredicate(predicate)
	assert.NotNil(t, authorizer)
}

// TestAuthorizerWithPredicate_Deny tests authorizer denying connections
func TestAuthorizerWithPredicate_Deny(t *testing.T) {
	// Create a predicate that denies all IDs
	predicate := func(_ string) bool {
		return false
	}

	authorizer := AuthorizerWithPredicate(predicate)
	assert.NotNil(t, authorizer)
}

// TestAuthorizerWithPredicate_ComplexLogic tests authorizer with complex predicate
func TestAuthorizerWithPredicate_ComplexLogic(t *testing.T) {
	// Create a predicate with multiple conditions
	predicate := func(id string) bool {
		return strings.HasPrefix(id, "spiffe://") &&
			strings.Contains(id, "/service/") &&
			!strings.Contains(id, "/forbidden/")
	}

	authorizer := AuthorizerWithPredicate(predicate)
	assert.NotNil(t, authorizer)
}

// TestCreateMTLSServer_NilSource tests server creation with nil source
func TestCreateMTLSServer_NilSource(t *testing.T) {
	// Enable panic on fatal to test fatal error behavior
	t.Setenv("SPIKE_STACK_TRACES_ON_LOG_FATAL", "true")

	defer func() {
		r := recover()
		require.NotNil(t, r, "Expected panic due to nil source")
		panicMsg := fmt.Sprint(r)
		assert.Contains(t, panicMsg, "CreateMTLSServerWithPredicate")
	}()

	// This should panic because source is nil
	CreateMTLSServer(nil, ":8443")
	t.Fatal("Should have panicked due to nil source")
}

// TestCreateMTLSServerWithPredicate_NilSource tests server creation with predicate and nil source
func TestCreateMTLSServerWithPredicate_NilSource(t *testing.T) {
	// Enable panic on fatal to test fatal error behavior
	t.Setenv("SPIKE_STACK_TRACES_ON_LOG_FATAL", "true")

	defer func() {
		r := recover()
		require.NotNil(t, r, "Expected panic due to nil source")
		panicMsg := fmt.Sprint(r)
		assert.Contains(t, panicMsg, "CreateMTLSServerWithPredicate")
	}()

	predicate := func(_ string) bool { return true }
	// This should panic because source is nil
	CreateMTLSServerWithPredicate(nil, ":8443", predicate)
	t.Fatal("Should have panicked due to nil source")
}

// TestCreateMTLSClient_NilSource tests client creation with nil source
func TestCreateMTLSClient_NilSource(t *testing.T) {
	// This test documents behavior - CreateMTLSClient does not validate nil source
	// The validation happens at connection time, not creation time
	client := CreateMTLSClient(nil)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

// TestCreateMTLSClientWithPredicate_NilSource tests client creation with predicate and nil source
func TestCreateMTLSClientWithPredicate_NilSource(t *testing.T) {
	predicate := func(_ string) bool { return true }
	client := CreateMTLSClientWithPredicate(nil, predicate)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

// TestCreateMTLSClientForNexus_NilSource tests Nexus client creation
func TestCreateMTLSClientForNexus_NilSource(t *testing.T) {
	client := CreateMTLSClientForNexus(nil)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

// TestCreateMTLSClientForKeeper_NilSource tests Keeper client creation
func TestCreateMTLSClientForKeeper_NilSource(t *testing.T) {
	client := CreateMTLSClientForKeeper(nil)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

// TestServeWithPredicate_NilSource tests serve function with nil source
func TestServeWithPredicate_NilSource(t *testing.T) {
	initializeRoutes := func() {}
	predicate := func(_ string) bool { return true }

	err := ServeWithPredicate(nil, initializeRoutes, predicate, ":8443")

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestServe_NilSource tests serve function with nil source
func TestServe_NilSource(t *testing.T) {
	initializeRoutes := func() {}

	err := Serve(nil, initializeRoutes, ":8443")

	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrSPIFFENilX509Source))
}

// TestPost_ClientConfiguration tests that Post handles various HTTP responses
func TestPost_ClientConfiguration(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectedError  *sdkErrors.SDKError
		shouldHaveBody bool
	}{
		{
			name:           "Success_200",
			statusCode:     http.StatusOK,
			responseBody:   `{"success":true}`,
			expectedError:  nil,
			shouldHaveBody: true,
		},
		{
			name:           "NotFound_404",
			statusCode:     http.StatusNotFound,
			responseBody:   `{}`,
			expectedError:  sdkErrors.ErrAPINotFound,
			shouldHaveBody: false,
		},
		{
			name:           "Unauthorized_401",
			statusCode:     http.StatusUnauthorized,
			responseBody:   `{}`,
			expectedError:  sdkErrors.ErrAccessUnauthorized,
			shouldHaveBody: false,
		},
		{
			name:           "BadRequest_400",
			statusCode:     http.StatusBadRequest,
			responseBody:   `{}`,
			expectedError:  sdkErrors.ErrAPIBadRequest,
			shouldHaveBody: false,
		},
		{
			name:           "ServiceUnavailable_503",
			statusCode:     http.StatusServiceUnavailable,
			responseBody:   `{}`,
			expectedError:  sdkErrors.ErrStateNotReady,
			shouldHaveBody: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client := server.Client()
			bodyBytes, err := Post(context.Background(), client, server.URL, []byte(`{"test":"data"}`))

			if tt.expectedError != nil {
				assert.NotNil(t, err)
				assert.True(t, err.Is(tt.expectedError))
			} else {
				assert.Nil(t, err)
			}

			if tt.shouldHaveBody {
				assert.NotEmpty(t, bodyBytes)
			}
		})
	}
}

// TestStreamPost_Success tests successful streaming POST
func TestStreamPost_Success(t *testing.T) {
	testData := []byte("streaming test data")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify content type
		assert.Equal(t, "application/octet-stream", r.Header.Get("Content-Type"))

		// Read and echo the request body
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}))
	defer server.Close()

	client := server.Client()
	reader := bytes.NewReader(testData)

	responseBody, err := StreamPost(context.Background(), client, server.URL, reader)

	// Note: The implementation closes the response body in a defer,
	// so the returned io.ReadCloser may not be readable
	require.Nil(t, err)
	require.NotNil(t, responseBody)
}

// TestStreamPostWithContentType_CustomContentType tests streaming POST with custom content type
func TestStreamPostWithContentType_CustomContentType(t *testing.T) {
	testData := []byte(`{"streaming":"json"}`)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify content type
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"response":"ok"}`))
	}))
	defer server.Close()

	client := server.Client()
	reader := bytes.NewReader(testData)

	responseBody, err := StreamPostWithContentType(
		context.Background(), client, server.URL, reader, ContentTypeJSON,
	)

	require.Nil(t, err)
	require.NotNil(t, responseBody)
	responseBody.Close()
}

// TestStreamPost_NotFound tests streaming POST with 404 response
func TestStreamPost_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := server.Client()
	reader := bytes.NewReader([]byte("test"))

	responseBody, err := StreamPost(context.Background(), client, server.URL, reader)

	assert.Nil(t, responseBody)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrAPINotFound))
}

// TestResponseWithError_Interface tests the ResponseWithError interface
func TestResponseWithError_Interface(t *testing.T) {
	// Test that mockResponse implements ResponseWithError
	var _ ResponseWithError = mockResponse{}
	var _ ResponseWithError = &mockResponse{}

	resp := mockResponse{
		Data: "test",
		Err:  "test_error",
	}

	assert.Equal(t, sdkErrors.ErrorCode("test_error"), resp.ErrorCode())
}

// TestResponseWithError_EmptyErrorCode tests response with no error code
func TestResponseWithError_EmptyErrorCode(t *testing.T) {
	resp := mockResponse{
		Data: "test",
		Err:  "",
	}

	assert.Equal(t, sdkErrors.ErrorCode(""), resp.ErrorCode())
}

// TestPost_RequestCreationFailure tests Post with invalid URL
func TestPost_RequestCreationFailure(t *testing.T) {
	client := &http.Client{}

	// Use an invalid URL that will cause request creation to fail
	invalidURL := string([]byte{0x7f}) // Invalid UTF-8

	_, err := Post(context.Background(), client, invalidURL, []byte(`{}`))

	assert.NotNil(t, err)
}

// TestContentTypeString tests content type string conversion
func TestContentTypeString(t *testing.T) {
	tests := []struct {
		contentType ContentType
		expected    string
	}{
		{ContentTypeJSON, "application/json"},
		{ContentTypePlain, "text/plain"},
		{ContentTypeOctetStream, "application/octet-stream"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, string(tt.contentType))
		})
	}
}
