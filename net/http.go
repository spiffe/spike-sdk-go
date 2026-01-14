//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// RespondWithHTTPError processes errors from state operations (database,
// storage) and sends appropriate HTTP responses. It uses generics to work
// with any response type that implements the ErrorResponder interface.
//
// Use this function in route handlers after state operations (Get, Put, Delete,
// List, etc.) that may return "not found" or internal errors. Do NOT use this
// for authentication/authorization or input validation errors in
// guard/intercept functions; those have different semantics (400 Bad Request,
// 401 Unauthorized) that don't map to the 404/500 distinction this function
// provides, so they should use net.Fail directly.
//
// The function distinguishes between two types of errors:
//   - sdkErrors.ErrEntityNotFound: Returns HTTP 404 Not Found when the
//     requested resource does not exist
//   - Other errors: Returns HTTP 500 Internal Server Error for backend or
//     server-side failures
//
// Parameters:
//   - err: The error that occurred during the state operation
//   - w: The HTTP response writer for sending error responses
//   - response: A zero-value response instance used to generate error responses
//
// Returns:
//   - *sdkErrors.SDKError: The error that was passed in (for chaining),
//     or nil if err was nil
//
// Example usage:
//
//	// In a route handler after a state operation:
//	if err != nil {
//	    return net.RespondWithHTTPError(err, w, reqres.SecretGetResponse{})
//	}
//
//	// In guard/intercept functions, use net.Fail directly instead:
//	if !authorized {
//	    net.Fail(response.Unauthorized(), w, http.StatusUnauthorized)
//	    return sdkErrors.ErrAccessUnauthorized
//	}
func RespondWithHTTPError[T ErrorResponder[T]](
	err *sdkErrors.SDKError, w http.ResponseWriter, response T,
) *sdkErrors.SDKError {
	if err == nil {
		return nil
	}
	if err.Is(sdkErrors.ErrEntityNotFound) {
		failErr := Fail(response.NotFound(), w, http.StatusNotFound)
		if failErr != nil {
			return err.Wrap(failErr)
		}
		return err
	}
	// Backend or other server-side failure
	failErr := Fail(response.Internal(), w, http.StatusInternalServerError)
	if failErr != nil {
		return err.Wrap(failErr)
	}
	return err
}

// RespondWithInternalError sends an HTTP 500 Internal Server Error response and
// returns the provided SDK error. Use this for operations where the only
// possible error is an internal/server error (no "not found" case), such as
// cryptographic operations, Shamir secret sharing validation, or system
// initialization checks.
//
// Like HandleError, this is intended for route handlers after state or system
// operations. Do NOT use this for authentication/authorization or input
// validation errors in guard/intercept functions; those have different
// semantics (400 Bad Request, 401 Unauthorized) that this function doesn't
// handle, so they should use net.Fail directly.
//
// Parameters:
//   - err: The SDK error that occurred
//   - w: The HTTP response writer for sending error responses
//   - response: A zero-value response instance used to generate the error
//
// Returns:
//   - *sdkErrors.SDKError: The error that was passed in
//
// Example usage:
//
//	if cipher == nil {
//	    return net.RspondWithInternalError(
//	        sdkErrors.ErrCryptoCipherNotAvailable, w,
//	        reqres.BootstrapVerifyResponse{},
//	    )
//	}
func RespondWithInternalError[T ErrorResponder[T]](
	err *sdkErrors.SDKError, w http.ResponseWriter, response T,
) *sdkErrors.SDKError {
	// ErrAPIInternal forces an `http.InternalServerError`.
	return RespondWithHTTPError(sdkErrors.ErrAPIInternal.Wrap(err), w, response)
}
