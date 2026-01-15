//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/validation"
)

// RespondErrOnBadPermission validates the given permissions slice and writes an
// error response if the validation fails.
//
// This function checks that the permissions slice is not empty and that each
// permission is a valid PolicyPermission value. If validation fails, it sends
// the provided error response to the client with a 400 Bad Request status code.
//
// Type Parameters:
//   - T: The response type to send to the client in case of validation failure
//
// Parameters:
//   - permissions: []data.PolicyPermission - The permissions to validate
//   - badInputResp: T - The error response object to send if validation fails
//   - w: http.ResponseWriter - The response writer for error handling
//
// Returns:
//   - *sdkErrors.SDKError: ErrDataInvalidInput if permissions is empty or
//     contains an invalid permission, nil if validation succeeds
func RespondErrOnBadPermission[T any](
	permissions []data.PolicyPermission, badInputResp T, w http.ResponseWriter,
) *sdkErrors.SDKError {
	if len(permissions) == 0 {
		failErr := Fail(badInputResp, w, http.StatusBadRequest)
		if failErr != nil {
			return sdkErrors.ErrDataInvalidInput.Wrap(failErr)
		}
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	for _, perm := range permissions {
		if !validation.ValidPermission(string(perm)) {
			failErr := Fail(badInputResp, w, http.StatusBadRequest)
			if failErr != nil {
				return sdkErrors.ErrDataInvalidInput.Wrap(failErr)
			}
			return sdkErrors.ErrDataInvalidInput.Clone()
		}
	}
	return nil
}
