//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package errors provides structured error handling for the SPIKE SDK.
//
// This package defines SDKError, a structured error type with error codes for
// programmatic handling, and provides predefined sentinel errors for common
// failure conditions such as authentication failures, validation errors,
// initialization issues, and communication problems.
//
// # Sentinel Errors and Cloning
//
// All predefined errors (e.g., ErrEntityNotFound, ErrAccessUnauthorized) are
// pointer types (*SDKError) pointing to shared global instances. This design
// enables efficient error comparison using errors.Is().
//
// IMPORTANT: Because sentinel errors are shared pointers, you must NEVER
// mutate them directly. If you need to customize the error message, always
// use Clone() first:
//
//	// WRONG - mutates the shared global sentinel:
//	failErr := sdkErrors.ErrEntityNotFound
//	failErr.Msg = "custom message"  // BUG: corrupts the sentinel!
//
//	// CORRECT - clone before mutating:
//	failErr := sdkErrors.ErrEntityNotFound.Clone()
//	failErr.Msg = "custom message"  // Safe: only affects the clone
//
// When returning sentinel errors without modification, use Clone() for
// defensive programming to prevent consumers from accidentally mutating them:
//
//	// Defensive return - prevents consumer mutation:
//	return sdkErrors.ErrEntityNotFound.Clone()
//
// The Wrap() method is safe because it creates a new instance:
//
//	// Safe - Wrap() returns a new SDKError:
//	failErr := sdkErrors.ErrEntityNotFound.Wrap(dbErr)
//	failErr.Msg = "custom message"  // Safe: modifies the new instance
//
// # Error Comparison
//
// Always use errors.Is() for error comparison. Two SDKErrors are considered
// equal if they have the same error code, regardless of message or wrapped
// error:
//
//	if errors.Is(err, sdkErrors.ErrEntityNotFound) {
//	    // Handle not found case
//	}
//
// # Usage Patterns (ADR-0028)
//
//  1. All SDK and non-CLI SPIKE errors shall use SDKError
//  2. All comparisons shall be done with errors.Is()
//  3. Context information shall be included in the Msg field (after cloning)
//  4. Import SDK errors as `sdkErrors` consistently for easier code search
//  5. Use predefined errors and wrap them with context, don't create from codes
//
// # Example
//
//	import sdkErrors "github.com/spiffe/spike-sdk-go/errors"
//
//	func GetSecret(path string) (*Secret, *sdkErrors.SDKError) {
//	    if path == "" {
//	        return nil, sdkErrors.ErrDataInvalidInput.Clone()
//	    }
//
//	    secret, err := store.Get(path)
//	    if err != nil {
//	        if errors.Is(err, store.ErrNotFound) {
//	            failErr := sdkErrors.ErrEntityNotFound.Clone()
//	            failErr.Msg = fmt.Sprintf("secret not found: %s", path)
//	            return nil, failErr
//	        }
//	        return nil, sdkErrors.ErrEntityLoadFailed.Wrap(err)
//	    }
//
//	    return secret, nil
//	}
package errors
