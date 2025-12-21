//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"bytes"
	"context"
	"io"
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// Post performs an HTTP POST request with a JSON payload and returns the
// response body. It handles the common cases of connection errors, non-200
// status codes, and proper response body handling.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - client: An *http.Client used to make the request, typically
//     configured with TLS settings.
//   - path: The URL path to send the POST request to.
//   - mr: A byte slice containing the marshaled JSON request body.
//
// Returns:
//   - []byte: The response body if the request is successful
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrAPIBadRequest: if request creation fails or server returns 400
//   - ErrNetPeerConnection: if connection to peer fails or unexpected status
//     code
//   - ErrAPINotFound: if server returns 404
//   - ErrAccessUnauthorized: if server returns 401
//   - ErrStateNotReady: if server returns 503
//   - ErrNetReadingResponseBody: if reading response body fails
//
// The function ensures proper cleanup by always attempting to close the
// response body, even if an error occurs during reading.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	client := &http.Client{}
//	data := []byte(`{"key": "value"}`)
//	response, err := Post(ctx, client, "https://api.example.com/endpoint", data)
//	if err != nil {
//	    log.Fatalf("failed to post: %v", err)
//	}
func Post(
	ctx context.Context, client *http.Client, path string, mr []byte,
) ([]byte, *sdkErrors.SDKError) {
	const fName = "Post"

	if ctx == nil {
		ctx = context.Background()
	}

	// Create the request while preserving the mTLS client
	req, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewBuffer(mr))
	if err != nil {
		failErr := sdkErrors.ErrAPIBadRequest.Wrap(err)
		failErr.Msg = "failed to create request"
		return nil, failErr
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Use the existing mTLS client to make the request
	//nolint:bodyclose // Response body is properly closed in defer block
	r, err := client.Do(req)
	if err != nil {
		failErr := sdkErrors.ErrNetPeerConnection.Wrap(err)
		return []byte{}, failErr
	}
	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		closeErr := b.Close()
		if closeErr != nil {
			failErr := sdkErrors.ErrFSStreamCloseFailed.Wrap(err)
			failErr.Msg = "failed to close response body"
			log.WarnErr(fName, *failErr)
		}
	}(r.Body)

	if r.StatusCode != http.StatusOK {
		if r.StatusCode == http.StatusNotFound {
			return []byte{}, sdkErrors.ErrAPINotFound.Clone()
		}

		if r.StatusCode == http.StatusUnauthorized {
			return []byte{}, sdkErrors.ErrAccessUnauthorized.Clone()
		}

		if r.StatusCode == http.StatusBadRequest {
			return []byte{}, sdkErrors.ErrAPIBadRequest.Clone()
		}

		// SPIKE Nexus is likely not initialized or in bad shape:
		if r.StatusCode == http.StatusServiceUnavailable {
			return []byte{}, sdkErrors.ErrStateNotReady.Clone()
		}

		failErr := sdkErrors.ErrNetPeerConnection.Clone()
		failErr.Msg = "unexpected status code from peer"
		return []byte{}, failErr
	}

	b, sdkErr := body(r)
	if sdkErr != nil {
		return nil, sdkErr
	}

	return b, nil
}
