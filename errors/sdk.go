//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

import (
	"errors"
	"fmt"
)

// SDKError represents a structured error in the SPIKE SDK. It provides error
// codes for programmatic handling, human-readable messages, and support for
// error wrapping to maintain error chains.
//
// Usage patterns (see ADR-0028):
//  1. All SDK and non-CLI SPIKE errors shall use SDKError
//  2. All comparisons shall be done with errors.Is()
//  3. Context information shall be included in the Msg field
//  4. Import SDK errors as `sdkErrors` consistently for easier code search
//  5. Use predefined errors and wrap them with context, don't create from codes
//
// Example:
//
//	// Use predefined errors
//	return sdkErrors.ErrEntityNotFound
//
//	// Or wrap with additional context
//	return sdkErrors.ErrEntityNotFound.Wrap(dbErr)
//
//	// Check error types
//	if errors.Is(err, sdkErrors.ErrEntityNotFound) {
//	    // Handle not found error
//	}
type SDKError struct {
	// Code is the error code for programmatic error handling
	Code ErrorCode

	// Msg is the human-readable error message
	Msg string

	// Wrapped is the underlying error, if any
	Wrapped error
}

// New creates a new SDKError with the specified error code, message, and
// optional wrapped error.
//
// Note: In most cases, you should use predefined errors
// (e.g., ErrEntityNotFound) and wrap them with .Wrap() instead of creating new
// errors from codes directly.
//
// Parameters:
//   - code: the error code identifying the error type
//   - msg: human-readable error message providing context
//   - wrapped: optional underlying error to wrap (can be nil)
//
// Returns:
//   - *SDKError: a new SDK error instance
//
// Example:
//
//	// Creating a custom error (rare, prefer using predefined errors)
//	err := sdkErrors.New(
//	    sdkErrors.ErrEntityNotFound.Code,
//	    "secret 'prod-api-key' not found in vault 'production'",
//	    dbErr,
//	)
func New(code ErrorCode, msg string, wrapped error) *SDKError {
	return &SDKError{
		Code:    code,
		Msg:     msg,
		Wrapped: wrapped,
	}
}

// Error implements the error interface, returning a formatted error message
// that includes the error code, message, and recursively includes wrapped
// error messages.
//
// Returns:
//   - string: formatted error message with error code and full error chain
func (e *SDKError) Error() string {
	if e.Wrapped != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Msg, e.Wrapped)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Msg)
}

// Unwrap returns the wrapped error, enabling error chain traversal with
// errors.Is() and errors.As() from the standard library.
//
// Returns:
//   - error: the wrapped error, or nil if no error was wrapped
func (e *SDKError) Unwrap() error {
	return e.Wrapped
}

// Wrap creates a new SDKError that wraps the provided error, preserving
// the current error's code and message while adding the new error to the
// error chain.
//
// Parameters:
//   - err: the error to wrap in the error chain
//
// Returns:
//   - *SDKError: a new SDK error with the same code and message but with
//     the provided error wrapped
//
// Example:
//
//	// Wrap a database error with entity not found error
//	return sdkErrors.ErrEntityNotFound.Wrap(dbErr)
func (e *SDKError) Wrap(err error) *SDKError {
	return &SDKError{
		Code:    e.Code,
		Msg:     e.Msg,
		Wrapped: err,
	}
}

// Is enables error comparison by error code using errors.Is() from the
// standard library. Two SDKErrors are considered equal if they have the
// same error code.
//
// Parameters:
//   - target: the error to compare against
//
// Returns:
//   - bool: true if target is an SDKError with the same error code
//
// Example:
//
//	if errors.Is(err, sdkErrors.ErrEntityNotFound) {
//	    // Handle not found error
//	}
func (e *SDKError) Is(target error) bool {
	var t *SDKError
	if errors.As(target, &t) {
		return e.Code == t.Code
	}
	return false
}

// Clone creates a shallow copy of the SDKError. This is useful when you need
// to modify the Msg field of a sentinel error without mutating the original.
//
// Returns:
//   - *SDKError: a new SDK error with the same code, message, and wrapped
//     error as the original
//
// Example:
//
//	// Copy a sentinel error to customize the message
//	failErr := sdkErrors.ErrEntityNotFound.Clone()
//	failErr.Msg = "secret 'prod-api-key' not found"
//	return failErr
func (e *SDKError) Clone() *SDKError {
	return &SDKError{
		Code:    e.Code,
		Msg:     e.Msg,
		Wrapped: e.Wrapped,
	}
}
