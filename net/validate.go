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

// RespondErrOnBadPolicyID validates the given policy ID and writes an error
// response if the validation fails.
//
// This function checks if the policy ID is a valid UUID format. If validation
// fails, it sends the provided error response to the client with a 400 Bad
// Request status code.
//
// Type Parameters:
//   - T: The response type to send to the client in case of validation failure
//
// Parameters:
//   - policyID: string - The policy ID to validate (must be a valid UUID)
//   - w: http.ResponseWriter - The response writer for error handling
//   - errorResponse: T - The error response object to send if validation fails
//
// Returns:
//   - *sdkErrors.SDKError: ErrDataInvalidInput if the policy ID is invalid,
//     nil if validation succeeds
func RespondErrOnBadPolicyID[T any](
	policyID string, w http.ResponseWriter, errorResponse T,
) *sdkErrors.SDKError {
	validationErr := validation.ValidatePolicyID(policyID)
	if invalidPolicy := validationErr != nil; invalidPolicy {
		failErr := Fail(
			errorResponse, w, http.StatusBadRequest,
		)
		validationErr.Msg = "invalid policy ID: " + policyID
		if failErr != nil {
			return validationErr.Wrap(failErr)
		}
		return validationErr
	}
	return nil
}

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
