//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"crypto/cipher"
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// CryptoHandler is a function type for HTTP handlers that require cryptographic
// operations. It wraps standard HTTP handling with cipher retrieval, allowing
// handlers to perform encrypted request/response processing.
//
// The getCipher function provides lazy cipher retrieval with built-in error
// handling - if the cipher is unavailable, it automatically sends an
// appropriate error response to the client.
//
// Parameters:
//   - w: The HTTP response writer
//   - r: The HTTP request
//   - getCipher: A function that retrieves the cipher and handles errors
//
// Returns:
//   - *sdkErrors.SDKError: An error if the handler fails, nil on success
type CryptoHandler func(
	w http.ResponseWriter, r *http.Request,
	getCipher func() (cipher.AEAD, *sdkErrors.SDKError),
) *sdkErrors.SDKError

// const spikeCipherVersion = byte('1')
const headerKeyContentType = "Content-Type"
const headerValueOctetStream = "application/octet-stream"

// getCipherOrFailStreaming retrieves the system cipher from the backend
// and handles errors for streaming mode requests.
//
// If the cipher is unavailable, sends a plain HTTP error response.
//
// Parameters:
//   - w: The HTTP response writer for sending error responses
//
// Returns:
//   - cipher.AEAD: The system cipher if available, nil otherwise
//   - *sdkErrors.SDKError: An error if the cipher is unavailable, nil otherwise
func getCipherOrFailStreaming(
	w http.ResponseWriter,
	getCipher func() cipher.AEAD,
) (cipher.AEAD, *sdkErrors.SDKError) {
	c := getCipher()

	if c == nil {
		http.Error(
			w, string(sdkErrors.ErrCryptoCipherNotAvailable.Code),
			http.StatusInternalServerError,
		)
		return nil, sdkErrors.ErrCryptoCipherNotAvailable.Clone()
	}

	return c, nil
}

// getCipherOrFailJSON retrieves the system cipher from the backend and
// handles errors for JSON mode requests.
//
// If the cipher is unavailable, sends a structured JSON error response.
//
// Parameters:
//   - w: The HTTP response writer for sending error responses
//   - errorResponse: The error response of type T to send as JSON
//
// Returns:
//   - cipher.AEAD: The system cipher if available, nil otherwise
//   - *sdkErrors.SDKError: An error if the cipher is unavailable, nil otherwise
func getCipherOrFailJSON[T any](
	w http.ResponseWriter, errorResponse T,
	getCipher func() cipher.AEAD,
) (cipher.AEAD, *sdkErrors.SDKError) {
	c := getCipher()
	if c == nil {
		failErr := Fail(errorResponse, w, http.StatusInternalServerError)
		if failErr != nil {
			return nil, sdkErrors.ErrCryptoCipherNotAvailable.Wrap(failErr)
		}
		return nil, sdkErrors.ErrCryptoCipherNotAvailable.Clone()
	}

	return c, nil
}

// DispatchByContentType routes requests to the appropriate handler based on
// the Content-Type header. Streaming mode (application/octet-stream) uses
// binary I/O, while all other content types use JSON encoding.
//
// Parameters:
//   - w: The HTTP response writer
//   - r: The HTTP request
//   - streamHandler: Handler for streaming (binary) mode
//   - jsonHandler: Handler for JSON mode
//   - errResponse: Error response to use for JSON mode failures
//
// Returns:
//   - *sdkErrors.SDKError: An error if the handler fails
func DispatchByContentType(
	w http.ResponseWriter, r *http.Request,
	streamHandler CryptoHandler,
	jsonHandler CryptoHandler,
	getCipher func() cipher.AEAD,
	errResponse any,
) *sdkErrors.SDKError {
	contentType := r.Header.Get(headerKeyContentType)
	streamModeActive := contentType == headerValueOctetStream

	if streamModeActive {
		getCipherStream := func() (cipher.AEAD, *sdkErrors.SDKError) {
			return getCipherOrFailStreaming(w, getCipher)
		}
		return streamHandler(w, r, getCipherStream)
	}

	getCipherJSON := func() (cipher.AEAD, *sdkErrors.SDKError) {
		return getCipherOrFailJSON(w, errResponse, getCipher)
	}
	return jsonHandler(w, r, getCipherJSON)
}
