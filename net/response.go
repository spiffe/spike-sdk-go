//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"io"
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// body reads and returns the entire response body from an HTTP response.
// The response body is read completely and returned as a byte slice.
//
// This is an internal helper function used by the net package to process
// HTTP responses.
//
// Parameters:
//   - r: The HTTP response to read from
//
// Returns:
//   - []byte: The complete response body, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrNetReadingResponseBody: if reading the response body fails
//
// Example:
//
//	resp, err := http.Get(url)
//	if err != nil {
//	    return nil, err
//	}
//	defer resp.Body.Close()
//
//	bodyBytes, sdkErr := body(resp)
//	if sdkErr != nil {
//	    log.Printf("Failed to read response body: %v", sdkErr)
//	    return nil, sdkErr
//	}
func body(r *http.Response) ([]byte, *sdkErrors.SDKError) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		failErr := sdkErrors.ErrNetReadingResponseBody.Wrap(err)
		failErr.Msg = "failed to read HTTP response body"
		return nil, failErr
	}

	return bodyBytes, nil
}
