//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/validation"
)

// RespondErrOnBadPath validates the given path and writes an error response
// if the validation fails.
//
// This function checks if the path conforms to the expected format using
// validation.ValidatePath. If validation fails, it sends the provided error
// response to the client with a 400 Bad Request status code.
//
// Type Parameters:
//   - T: The response type to send to the client in case of validation failure
//
// Parameters:
//   - path: string - The secret path to validate
//   - badInputResp: T - The error response object to send if validation fails
//   - w: http.ResponseWriter - The response writer for error handling
//
// Returns:
//   - *sdkErrors.SDKError: ErrDataInvalidInput if the path is invalid,
//     nil if validation succeeds
func RespondErrOnBadPath[T any](
	path string, badInputResp T, w http.ResponseWriter,
) *sdkErrors.SDKError {
	pathErr := validation.ValidatePath(path)
	if pathErr != nil {
		failErr := Fail(badInputResp, w, http.StatusBadRequest)
		pathErr.Msg = "invalid secret path: " + path
		if failErr != nil {
			return pathErr.Wrap(failErr)
		}
		return pathErr
	}
	return nil
}

// RespondErrOnBadPathPattern validates the given path pattern and writes an
// error response if the validation fails.
//
// This function checks if the path pattern conforms to the expected format
// using validation.ValidatePathPattern. The path pattern may include regex
// metacharacters (^, $, ?, +, *, |, [], {}, \).
//
// If validation fails, it sends the provided error response to the client with
// a 400 Bad Request status code.
//
// Type Parameters:
//   - T: The response type to send to the client in case of validation failure
//
// Parameters:
//   - pathPattern: string - The path pattern to validate
//   - badInputResp: T - The error response object to send if validation fails
//   - w: http.ResponseWriter - The response writer for error handling
//
// Returns:
//   - *sdkErrors.SDKError: ErrDataInvalidInput if the path pattern is invalid,
//     nil if validation succeeds
func RespondErrOnBadPathPattern[T any](
	pathPattern string, badInputResp T, w http.ResponseWriter,
) *sdkErrors.SDKError {
	if err := validation.ValidatePathPattern(pathPattern); err != nil {
		failErr := Fail(badInputResp, w, http.StatusBadRequest)
		if failErr != nil {
			return sdkErrors.ErrDataInvalidInput.Wrap(failErr)
		}
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	return nil
}
