//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/spiffe/spike-sdk-go/log"
)

func body(r *http.Response) (bod []byte, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

// Post performs an HTTP POST request with a JSON payload and returns the
// response body. It handles the common cases of connection errors, non-200
// status codes, and proper response body handling.
//
// Parameters:
//   - client: An *http.Client used to make the request, typically
//     configured with TLS settings.
//   - path: The URL path to send the POST request to.
//   - mr: A byte slice containing the marshaled JSON request body.
//
// Returns:
//   - []byte: The response body if the request is successful.
//   - error: An error if any of the following occur:
//   - Connection failure during POST request
//   - Non-200 status code in response
//   - Failure to read response body
//   - Failure to close response body
//
// The function ensures proper cleanup by always attempting to close the
// response body, even if an error occurs during reading. Any error from closing
// the body is joined with any existing error using errors.Join.
//
// Example:
//
//	client := &http.Client{}
//	data := []byte(`{"key": "value"}`)
//	response, err := Post(client, "https://api.example.com/endpoint", data)
//	if err != nil {
//	    log.Fatalf("failed to post: %v", err)
//	}
func Post(client *http.Client, path string, mr []byte) ([]byte, error) {
	const fName = "Post"

	// Create the request while preserving the mTLS client
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(mr))
	if err != nil {
		return nil, errors.Join(
			errors.New("post: Failed to create request"),
			err,
		)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Use the existing mTLS client to make the request
	r, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Join(
			errors.New("post: Problem connecting to peer"),
			err,
		)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Log().Info(
				fName,
				"msg", "Failed to close response body",
				"err", err.Error(),
			)
		}
	}(r.Body)

	if r.StatusCode != http.StatusOK {
		if r.StatusCode == http.StatusNotFound {
			return []byte{}, ErrNotFound
		}

		if r.StatusCode == http.StatusUnauthorized {
			return []byte{}, ErrUnauthorized
		}

		if r.StatusCode == http.StatusBadRequest {
			return []byte{}, ErrBadRequest
		}

		// SPIKE Nexus is likely not initialized or in bad shape:
		if r.StatusCode == http.StatusServiceUnavailable {
			return []byte{}, ErrNotReady
		}

		return []byte{}, errors.New("post: Problem connecting to peer")
	}

	b, err := body(r)
	if err != nil {
		return []byte{}, errors.Join(
			errors.New("post: Problem reading response body"),
			err,
		)
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err = errors.Join(err, b.Close())
	}(r.Body)

	return b, nil
}
