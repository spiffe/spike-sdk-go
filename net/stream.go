//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"errors"
	"io"
	"net/http"

	"github.com/spiffe/spike-sdk-go/log"
)

// StreamPostWithContentType performs an HTTP POST request with streaming data
// and a custom content type, returning the response body as a stream.
//
// This function is designed for streaming large amounts of data without loading
// the entire payload into memory. The caller is responsible for closing the
// returned io.ReadCloser.
//
// Parameters:
//   - client *http.Client: The HTTP client to use for the request
//   - path string: The URL path to POST to
//   - body io.Reader: The request body data stream
//   - contentType string: The MIME type of the request body
//     (e.g., "application/json", "text/plain")
//
// Returns:
//   - io.ReadCloser: The response body stream if successful
//     (must be closed by caller)
//   - error: nil on success, or one of the following well-known errors:
//   - ErrNotFound (404): Resource not found
//   - ErrUnauthorized (401): Authentication required
//   - ErrBadRequest (400): Invalid request
//   - ErrNotReady (503): Service unavailable
//   - Generic error for other non-200 status codes
//
// Example:
//
//		data := strings.NewReader("large data payload")
//		response, err := StreamPostWithContentType(client,
//	 	"/impl/upload", data, "text/plain")
//		if err != nil {
//		    return err
//		}
//		defer response.Close()
//		// Process streaming response...
func StreamPostWithContentType(
	client *http.Client, path string, body io.Reader, contentType string,
) (io.ReadCloser, error) {
	const fName = "StreamPostWithContentType"

	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		return nil, errors.Join(
			errors.New("streamPost: Failed to create request"),
			err,
		)
	}
	req.Header.Set("Content-Type", contentType)

	//nolint:bodyclose // Response body is properly closed in defer block
	r, err := client.Do(req)
	if err != nil {
		return nil, errors.Join(
			errors.New("streamPost: Problem connecting to peer"),
			err,
		)
	}
	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Log().Info(fName,
				"message", "Failed to close response body",
				"err", err.Error(),
			)
		}
	}(r.Body)

	if r.StatusCode != http.StatusOK {
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

// StreamPost is a convenience wrapper for StreamPostWithContentType that uses
// the default content type "application/octet-stream".
//
// This function is ideal for posting binary data or when the specific content
// type doesn't matter. The caller is responsible for closing the returned
// io.ReadCloser.
//
// Parameters:
//   - client *http.Client: The HTTP client to use for the request
//   - path string: The URL path to POST to
//   - body io.Reader: The request body data stream
//
// Returns:
//   - io.ReadCloser: The response body stream if successful
//     (must be closed by caller)
//   - error: nil on success, or a well-known error
//     (see StreamPostWithContentType)
//
// Example:
//
//	binaryData := bytes.NewReader(fileBytes)
//	response, err := StreamPost(client, "/impl/upload", binaryData)
//	if err != nil {
//	    return err
//	}
//	defer response.Close()
//	// Process response...
func StreamPost(
	client *http.Client, path string, body io.Reader,
) (io.ReadCloser, error) {
	return StreamPostWithContentType(
		client, path, body, "application/octet-stream",
	)
}
