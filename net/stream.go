//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"errors"
	"io"
	"net/http"
)

// StreamPostWithContentType performs an HTTP POST request with the provided
// content-type and body and returns the response body as an io.ReadCloser on
// success. The caller is responsible for closing the returned body.
//
// It mirrors error handling of Post: maps common non-200 statuses to well-known
// errors and ensures response bodies are closed on error.
func StreamPostWithContentType(client *http.Client, path string, body io.Reader, contentType string) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		return nil, errors.Join(
			errors.New("streamPost: Failed to create request"),
			err,
		)
	}
	req.Header.Set("Content-Type", contentType)

	r, err := client.Do(req)
	if err != nil {
		return nil, errors.Join(
			errors.New("streamPost: Problem connecting to peer"),
			err,
		)
	}

	if r.StatusCode != http.StatusOK {
		defer r.Body.Close()

		if r.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		if r.StatusCode == http.StatusUnauthorized {
			return nil, ErrUnauthorized
		}
		if r.StatusCode == http.StatusBadRequest {
			return nil, ErrBadRequest
		}
		if r.StatusCode == http.StatusServiceUnavailable {
			return nil, ErrNotReady
		}
		return nil, errors.New("streamPost: Problem connecting to peer")
	}

	return r.Body, nil
}

// StreamPost is a convenience wrapper that posts with
// Content-Type: application/octet-stream.
func StreamPost(client *http.Client, path string, body io.Reader) (io.ReadCloser, error) {
	return StreamPostWithContentType(client, path, body, "application/octet-stream")
}
