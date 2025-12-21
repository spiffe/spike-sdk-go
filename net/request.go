//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"context"
	"encoding/json"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

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
