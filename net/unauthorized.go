//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// RespondUnauthorizedAndWrapError sends an HTTP 401 Unauthorized response if no
// response has been sent yet, and wraps the provided failure error with
// ErrAccessUnauthorized.
//
// This function handles the common pattern of sending an unauthorized response
// while preserving error context for logging. It checks if a response has
// already been sent (indicated by err being non-nil) and only writes the
// response if needed.
//
// Parameters:
//   - err: An error from a previous operation (e.g., MarshalBodyAndRespondOnFail).
//     If nil, indicates no response has been sent yet, and this function will
//     send one.
//   - failErr: The underlying failure error to wrap with ErrAccessUnauthorized
//   - w: The HTTP response writer for sending the unauthorized response
//   - responseBody: The pre-marshaled JSON response body to send
//
// Returns:
//   - *sdkErrors.SDKError: ErrAccessUnauthorized wrapping failErr (and any
//     response writing errors if they occurred)
//
// Example usage:
//
//	responseBody, marshalErr := net.MarshalBodyAndRespondOnMarshalFail(
//	    errorResponse, w,
//	)
//	return net.RespondUnauthorizedAndWrapError(
//	    marshalErr, sdkErrors.ErrSPIFFEInvalidSPIFFEID, w, responseBody,
//	)
func RespondUnauthorizedAndWrapError(
	err *sdkErrors.SDKError,
	failErr *sdkErrors.SDKError,
	w http.ResponseWriter,
	responseBody []byte,
) *sdkErrors.SDKError {
	if notRespondedYet := err == nil; notRespondedYet {
		respondErr := Respond(http.StatusUnauthorized, responseBody, w)
		if respondErr != nil {
			notAuthorizedErr := sdkErrors.ErrAccessUnauthorized.Wrap(
				failErr.Wrap(respondErr),
			)
			return notAuthorizedErr
		}
	}
	notAuthorizedErr := sdkErrors.ErrAccessUnauthorized.Wrap(failErr)
	return notAuthorizedErr
}
