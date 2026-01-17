//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/log"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// ResponseWithError is an interface for response types that include an error
// code field.
// This allows generic error handling across different API response types.
type ResponseWithError interface {
	ErrorCode() sdkErrors.ErrorCode
}

// PostAndUnmarshal performs a complete request/response cycle for SPIKE Nexus
// API calls. It handles client creation, request posting, response
// unmarshaling, and error checking.
//
// Type parameter T must be a response type that implements ResponseWithError.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - urlPath: The URL path to send the POST request to
//   - requestBody: Marshaled JSON request body
//
// Returns:
//   - (*T, nil) containing the unmarshaled response if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - Errors from Post(): including ErrAPINotFound, ErrAccessUnauthorized, etc.
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the response contains an error code
//
// Note: Callers should check for specific errors and handle them as needed:
//
//	response, err := net.PostAndUnmarshal[MyResponse](ctx, source, url, body)
//	if err != nil {
//	    if err.Is(sdkErrors.ErrAPINotFound) {
//	        // Handle not found case (e.g., return empty slice for lists)
//	        return &[]MyType{}, nil
//	    }
//	    return nil, err
//	}
//
// Example:
//
//	type MyResponse struct {
//	    Data string              `json:"data"`
//	    Err  sdkErrors.ErrorCode `json:"err,omitempty"`
//	}
//
//	func (r *MyResponse) ErrorCode() sdkErrors.ErrorCode { return r.Err }
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	response, err := net.PostAndUnmarshal[MyResponse](
//	    ctx, source, "https://api.example.com/endpoint", requestBody)
func PostAndUnmarshal[T ResponseWithError](
	ctx context.Context,
	source *workloadapi.X509Source,
	urlPath string,
	requestBody []byte,
) (*T, *sdkErrors.SDKError) {
	client := CreateMTLSClientForNexus(source)

	postBody, err := Post(ctx, client, urlPath, requestBody)
	if err != nil {
		return nil, err
	}

	var response T
	if unmarshalErr := json.Unmarshal(postBody, &response); unmarshalErr != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(unmarshalErr)
		failErr.Msg = "problem parsing response body"
		return nil, failErr
	}

	if errCode := response.ErrorCode(); errCode != "" {
		return nil, sdkErrors.FromCode(errCode)
	}

	return &response, nil
}

// RequestBody reads and returns the entire request body as a byte slice.
// It reads all data from r.Body and ensures the body is properly closed
// after reading, even if an error occurs during the read operation.
//
// Close errors are logged but not returned to the caller, as the primary
// operation (reading the body data) has already completed. If reading fails,
// the error is returned immediately.
//
// Parameters:
//   - r: HTTP request containing the body to read
//
// Returns:
//   - bod: byte slice containing the full request body data on success, nil on
//     error
//   - err: *sdkErrors.SDKError with ErrNetReadingRequestBody if reading fails,
//     nil on success (close errors are only logged)
//
// Example:
//
//	body, err := RequestBody(req)
//	if err != nil {
//	    log.Printf("Failed to read request body: %v", err)
//	    return
//	}
//	// Process body data...
func RequestBody(r *http.Request) (bod []byte, err *sdkErrors.SDKError) {
	const fName = "RequestBody"

	readBody, e := io.ReadAll(r.Body)
	if e != nil {
		failErr := sdkErrors.ErrNetReadingRequestBody.Wrap(e)
		return nil, failErr
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		// This would almost never happen:
		if closeErr := b.Close(); closeErr != nil {
			failErr := sdkErrors.ErrFSStreamCloseFailed.Wrap(e)
			log.WarnErr(fName, *failErr)
		}
	}(r.Body)

	return readBody, err
}
