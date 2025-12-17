//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"io"
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// StreamPostWithContentType performs an HTTP POST request with streaming data
// and a custom content type, returning the response body as a stream.
//
// This function is designed for streaming large amounts of data without loading
// the entire payload into memory.
//
// Resource Management: On success, returns an open io.ReadCloser that the caller
// MUST close (typically with defer). On error, any response body is automatically
// closed by this function and nil is returned, following the canonical Go pattern
// of returning (zero-value, error) on failures.
//
// Parameters:
//   - client *http.Client: The HTTP client to use for the request
//   - path string: The URL path to POST to
//   - body io.Reader: The request body data stream
//   - contentType ContentType: The MIME type of the request body
//     (e.g., ContentTypeJSON, ContentTypeTextPlain, ContentTypeOctetStream)
//
// Returns:
//   - io.ReadCloser: The response body stream on success (must be closed by caller),
//     nil on error (already closed by this function)
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrAPINotFound (404): Resource not found
//   - ErrAccessUnauthorized (401): Authentication required
//   - ErrAPIBadRequest (400): Invalid request
//   - ErrStateNotReady (503): Service unavailable
//   - Generic error for other non-200 status codes
//
// Example:
//
//		data := strings.NewReader("large data payload")
//		response, err := StreamPostWithContentType(client,
//	 	"/api/upload", data, "text/plain")
//		if err != nil {
//		    return err
//		}
//		defer response.Close()
//		// Process streaming response...
func StreamPostWithContentType(
	client *http.Client, path string, body io.Reader,
	contentType ContentType,
) (io.ReadCloser, *sdkErrors.SDKError) {
	const fName = "StreamPostWithContentType"

	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		failErr := sdkErrors.ErrAPIPostFailed.Wrap(err)
		failErr.Msg = "failed to create request"
		return nil, failErr
	}
	req.Header.Set("Content-Type", string(contentType))

	r, err := client.Do(req)
	if err != nil {
		failErr := sdkErrors.ErrNetPeerConnection.Wrap(err)
		return nil, failErr
	}

	if r.StatusCode != http.StatusOK {
		// Close body on error paths before returning
		if r.Body != nil {
			closeErr := r.Body.Close()
			if closeErr != nil {
				failErr := sdkErrors.ErrFSStreamCloseFailed.Clone()
				failErr.Msg = "failed to close response body on error path"
				log.WarnErr(fName, *failErr)
			}
		}

		switch r.StatusCode {
		case http.StatusNotFound:
			return nil, sdkErrors.ErrAPINotFound
		case http.StatusUnauthorized:
			return nil, sdkErrors.ErrAccessUnauthorized
		case http.StatusBadRequest:
			return nil, sdkErrors.ErrAPIBadRequest
		case http.StatusServiceUnavailable:
			return nil, sdkErrors.ErrStateNotReady
		default:
			failErr := sdkErrors.ErrNetPeerConnection
			return nil, failErr
		}
	}

	// Success: return open body for caller to close
	return r.Body, nil
}

// StreamPost is a convenience wrapper for StreamPostWithContentType that uses
// the default content type ContentTypeOctetStream ("application/octet-stream").
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
//   - *sdkErrors.SDKError: nil on success, or a well-known error
//     (see StreamPostWithContentType)
//
// Example:
//
//	binaryData := bytes.NewReader(fileBytes)
//	response, err := StreamPost(client, "/api/upload", binaryData)
//	if err != nil {
//	    return err
//	}
//	defer response.Close()
//	// Process response...
func StreamPost(
	client *http.Client, path string, body io.Reader,
) (io.ReadCloser, *sdkErrors.SDKError) {
	return StreamPostWithContentType(
		client, path, body, ContentTypeOctetStream,
	)
}
