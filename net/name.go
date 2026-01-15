//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/validation"
)

// RespondErrOnBadName validates the given name and writes an error response
// if the validation fails.
//
// This function checks if the name meets length and format constraints using
// validation.ValidateName. The name must be between 1 and 250 characters and
// contain only alphanumeric characters, hyphens, underscores, and spaces. If
// validation fails, it sends the provided error response to the client with a
// 400 Bad Request status code.
//
// Type Parameters:
//   - T: The response type to send to the client in case of validation failure
//
// Parameters:
//   - name: string - The name to validate
//   - badInputResp: T - The error response object to send if validation fails
//   - w: http.ResponseWriter - The response writer for error handling
//
// Returns:
//   - *sdkErrors.SDKError: ErrDataInvalidInput if the name is invalid,
//     nil if validation succeeds
func RespondErrOnBadName[T any](
	name string, badInputResp T, w http.ResponseWriter,
) *sdkErrors.SDKError {
	if err := validation.ValidateName(name); err != nil {
		failErr := Fail(badInputResp, w, http.StatusBadRequest)
		if failErr != nil {
			return sdkErrors.ErrDataInvalidInput.Wrap(failErr)
		}
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	return nil
}
