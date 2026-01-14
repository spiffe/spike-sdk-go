//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/predicate"
)

// RespondUnauthorizedOnPredicateFail extracts the peer SPIFFE ID from an HTTP
// request and validates it using the provided predicate function. If the
// SPIFFE ID extraction fails or the predicate returns false, it sends an
// HTTP 401 Unauthorized response with the provided failure response body.
//
// This function combines SPIFFE ID extraction with custom authorization logic,
// making it useful for route handlers that need to verify the caller's identity
// against specific criteria (e.g., checking if the caller is a known service,
// validating trust domain membership, or matching against an allowlist).
//
// Parameters:
//   - predicateFn: A function that takes a SPIFFE ID string and returns true
//     if the caller is authorized, false otherwise
//   - failureResponse: The response object to send if authorization fails
//   - w: The HTTP response writer for error responses
//   - r: The incoming HTTP request containing the peer's SPIFFE ID
//
// Returns:
//   - *sdkErrors.SDKError: nil if the SPIFFE ID was successfully extracted and
//     the predicate returned true; otherwise returns ErrAccessUnauthorized
//     (potentially wrapping additional errors from response writing)
//
// Example usage:
//
//	isAuthorizedService := func(spiffeID string) bool {
//	    return strings.HasPrefix(spiffeID, "spiffe://example.org/service/")
//	}
//
//	if err := net.RespondUnauthorizedOnPredicateFail(
//	    isAuthorizedService,
//	    reqres.SecretGetResponse{Err: data.ErrUnauthorized},
//	    w, r,
//	); err != nil {
//	    return err
//	}
func RespondUnauthorizedOnPredicateFail(
	predicateFn predicate.Predicate, failureResponse any,
	w http.ResponseWriter, r *http.Request,
) *sdkErrors.SDKError {
	peerSPIFFEID, err := ExtractPeerSPIFFEIDAndRespondOnFail(
		w, r, failureResponse,
	)
	if err != nil {
		return err
	}

	if !predicateFn(peerSPIFFEID.String()) {
		failErr := Fail(
			failureResponse, w,
			http.StatusUnauthorized,
		)
		if failErr != nil {
			return sdkErrors.ErrAccessUnauthorized.Wrap(failErr)
		}
		return sdkErrors.ErrAccessUnauthorized.Clone()
	}

	return nil
}

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
