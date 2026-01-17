//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	"github.com/spiffe/spike-sdk-go/config/env"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// minRequestID is the minimum valid request ID for Shamir secret sharing
// operations. Request IDs must be at least 1 to be valid.
const minRequestID = 1

// RespondErrOnBadRequestID validates that a request ID falls within the
// acceptable range for Shamir secret sharing operations and responds with
// a bad request error if validation fails.
//
// The request ID must be between minRequestID (1) and the configured maximum
// share count (env.ShamirMaxShareCountVal()). This ensures that shard indices
// are valid for the secret sharing scheme.
//
// Type Parameters:
//   - T: The type of the bad input response body to send on failure
//
// Parameters:
//   - requestID: The request ID to validate
//   - badInputResp: The response body to return if validation fails
//   - w: The HTTP response writer
//
// Returns:
//   - *sdkErrors.SDKError: nil if the request ID is valid, otherwise
//     ErrAPIBadRequest describing the validation failure
func RespondErrOnBadRequestID[T any](
	requestID int, badInputResp T, w http.ResponseWriter,
) *sdkErrors.SDKError {
	if requestID < minRequestID || requestID > env.ShamirMaxShareCountVal() {
		if failErr := Fail(badInputResp, w, http.StatusBadRequest); failErr != nil {
			return sdkErrors.ErrAPIBadRequest.Wrap(failErr)
		}
		return sdkErrors.ErrAPIBadRequest.Clone()
	}
	return nil
}

// RespondErrOnEmptyShard validates that a shard is non-nil and contains
// non-zero data, responding with a bad request error if validation fails.
//
// This function performs two checks on the shard:
//  1. The shard pointer must not be nil
//  2. The shard must contain at least one non-zero byte
//
// These validations ensure that the shard contains meaningful cryptographic
// data for Shamir secret sharing operations.
//
// Type Parameters:
//   - T: The type of the bad input response body to send on failure
//
// Parameters:
//   - shard: Pointer to a 32-byte array containing the shard data
//   - badInputResp: The response body to return if the shard is nil
//   - w: The HTTP response writer
//
// Returns:
//   - *sdkErrors.SDKError: nil if the shard is valid, otherwise
//     ErrAPIBadRequest with message "shard is nil or empty"
func RespondErrOnEmptyShard[T any](
	shard *[32]byte, badInputResp T, w http.ResponseWriter) *sdkErrors.SDKError {
	badReqErr := sdkErrors.ErrAPIBadRequest.Clone()
	badReqErr.Msg = "shard is nil or empty"

	if shard == nil {
		failErr := Fail(badInputResp, w, http.StatusBadRequest)
		if failErr != nil {
			return badReqErr.Wrap(failErr)
		}
		return badReqErr
	}

	allZero := true
	for _, b := range shard {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		failErr := Fail(badInputResp, w, http.StatusBadRequest)
		if failErr != nil {
			return badReqErr.Wrap(failErr)
		}
		return badReqErr
	}
	return nil
}
