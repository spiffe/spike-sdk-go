//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"crypto/cipher"
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/journal"
)

// ContentType represents HTTP Content-Type header values.
type ContentType string

// Common content type constants for HTTP requests and responses.
var (
	ContentTypeJSON        ContentType = "application/json"
	ContentTypePlain       ContentType = "text/plain"
	ContentTypeOctetStream ContentType = "application/octet-stream"
)

// Handler is a function type that processes HTTP requests with audit
// logging support.
//
// Parameters:
//   - w: HTTP response writer for sending the response
//   - r: HTTP request containing the incoming request data
//   - audit: Audit entry for logging the request lifecycle
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, error on failure
type Handler func(
	w http.ResponseWriter, r *http.Request, audit *journal.AuditEntry,
) *sdkErrors.SDKError

// HandlerWithReturn is a generic function type for HTTP handlers that return
// a typed response. It processes HTTP requests and returns both the response
// data and any error that occurred.
//
// Type Parameters:
//   - T: The type of the response data to return
//
// Parameters:
//   - w: HTTP response writer for sending the response
//   - r: HTTP request containing the incoming request data
//
// Returns:
//   - *T: Pointer to the response data on success, nil on failure
//   - *sdkErrors.SDKError: nil on success, error on failure
type HandlerWithReturn[T any] func(
	w http.ResponseWriter, r *http.Request,
) (*T, *sdkErrors.SDKError)

// HandlerWithEntity is a generic function type for HTTP handlers that receive
// a pre-parsed request entity. It processes requests where the request body
// has already been deserialized into a typed struct.
//
// Type Parameters:
//   - T: The type of the request entity
//
// Parameters:
//   - req: The deserialized request entity
//   - w: HTTP response writer for sending the response
//   - r: HTTP request containing the incoming request data
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, error on failure
type HandlerWithEntity[T any] func(
	req T, w http.ResponseWriter, r *http.Request,
) *sdkErrors.SDKError

// Encryptor is a function type for encrypting plaintext data using an AEAD
// cipher. It handles the encryption process and writes any necessary error
// responses.
//
// Parameters:
//   - plaintext: The data to encrypt
//   - c: The AEAD cipher to use for encryption
//   - w: HTTP response writer for sending error responses
//
// Returns:
//   - []byte: The nonce used for encryption
//   - []byte: The encrypted ciphertext
//   - *sdkErrors.SDKError: nil on success, error on failure
type Encryptor func(
	plaintext []byte, c cipher.AEAD, w http.ResponseWriter,
) ([]byte, []byte, *sdkErrors.SDKError)

// ErrorResponder defines an interface for response types that can generate
// standard error responses. All SDK response types in the reqres package
// implement this interface through their NotFound() and Internal() methods.
type ErrorResponder[T any] interface {
	NotFound() T
	Internal() T
}

// InternalErrorResponder defines an interface for response types that can
// generate internal error responses. This is a subset of ErrorResponder for
// cases where only internal errors are possible (no "not found" scenario).
type InternalErrorResponder[T any] interface {
	Internal() T
}
