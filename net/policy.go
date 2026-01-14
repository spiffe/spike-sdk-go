//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/validation"
)

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
