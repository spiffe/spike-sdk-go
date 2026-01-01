//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/spiffe"
	"github.com/spiffe/spike-sdk-go/validation"
)

// Respond writes a JSON response with the specified status code and body.
//
// This function sets the Content-Type header to application/json, adds cache
// invalidation headers (Cache-Control, Pragma, Expires), writes the provided
// status code, and sends the response body.
//
// Parameters:
//   - statusCode: int - The HTTP status code to send
//   - body: []byte - The pre-marshaled JSON response body
//   - w: http.ResponseWriter - The response writer to use
//
// Returns:
//   - *sdkErrors.SDKError: sdkErrors.ErrAPIInternal if writing fails,
//     nil on success
func Respond(
	statusCode int, body []byte, w http.ResponseWriter,
) *sdkErrors.SDKError {
	w.Header().Set("Content-Type", "application/json")

	// Add cache invalidation headers
	w.Header().Set(
		"Cache-Control",
		"no-store, no-cache, must-revalidate, private",
	)
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	w.WriteHeader(statusCode)

	_, err := w.Write(body)
	if err != nil {
		// At this point, we cannot respond. So there is little to send
		// back to the client. We can only log the error.
		// This should rarely, if ever, happen.
		return sdkErrors.ErrAPIInternal.Wrap(err)
	}

	return nil
}

// MarshalBodyAndRespondOnMarshalFail serializes a response object to JSON and
// handles error cases.
//
// This function attempts to marshal the provided response object to JSON bytes.
// If marshaling fails, it sends a 500 Internal Server Error response to the
// client and returns nil. The function handles all error logging and response
// writing for the error case.
//
// Parameters:
//   - res: any - The response object to marshal to JSON
//   - w: http.ResponseWriter - The response writer for error handling
//
// Returns:
//   - []byte: The marshaled JSON bytes, or nil if marshaling failed
//   - *sdkErrors.SDKError: sdkErrors.ErrAPIInternal if marshaling failed,
//     nil otherwise
func MarshalBodyAndRespondOnMarshalFail(
	res any, w http.ResponseWriter,
) ([]byte, *sdkErrors.SDKError) {
	body, err := json.Marshal(res)

	// Since this function is typically called with sentinel error values,
	// this error should, "typically", never happen.
	// That's why, instead of sending a "marshal failure" sentinel error,
	// we return an internal sentinel error (sdkErrors.ErrAPIInternal)
	if err != nil {
		// Chain an error for detailed internal logging.
		failErr := *sdkErrors.ErrAPIInternal.Clone()
		failErr.Msg = "problem generating response"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		internalErrJSON, marshalErr := json.Marshal(failErr)

		// Add extra info "after" marshaling to avoid leaking internal error details
		wrappedErr := failErr.Wrap(err)

		if marshalErr != nil {
			wrappedErr = wrappedErr.Wrap(marshalErr)
			// Cannot marshal; try a generic message instead.
			internalErrJSON = []byte(`{"error":"internal server error"}`)
		}
		_, err = w.Write(internalErrJSON)
		if err != nil {
			wrappedErr = wrappedErr.Wrap(err)
			// At this point, we cannot respond. So there is little to send.
			// We cannot even send a generic error message.
			// We can only log the error.
		}

		return nil, wrappedErr
	}

	// body marshaled successfully
	return body, nil
}

// Fail sends an error response to the client.
//
// This function marshals the client response and sends it with the specified
// HTTP status code.
//
// Type Parameters:
//   - T: The response type to send to the client (e.g.,
//     reqres.ShardPutBadInput)
//
// Parameters:
//   - clientResponse: The response object to send to the client
//   - w: The HTTP response writer for error responses
//   - statusCode: The HTTP status code to send (e.g., http.StatusBadRequest)
//
// Returns:
//   - *sdkErrors.SDKError: An error if writing the response fails,
//     nil on success
//
// Example usage:
//
//	if request.Shard == nil {
//	    net.Fail(reqres.ShardPutBadInput, w, http.StatusBadRequest)
//	    return errors.ErrInvalidInput
//	}
func Fail[T any](
	clientResponse T,
	w http.ResponseWriter,
	statusCode int,
) *sdkErrors.SDKError {
	responseBody, marshalErr := MarshalBodyAndRespondOnMarshalFail(
		clientResponse, w,
	)
	if notRespondedYet := marshalErr == nil; notRespondedYet {
		return Respond(statusCode, responseBody, w)
	}
	return nil
}

// ErrorResponder defines an interface for response types that can generate
// standard error responses. All SDK response types in the reqres package
// implement this interface through their NotFound() and Internal() methods.
type ErrorResponder[T any] interface {
	NotFound() T
	Internal() T
}

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

// InternalErrorResponder defines an interface for response types that can
// generate internal error responses. This is a subset of ErrorResponder for
// cases where only internal errors are possible (no "not found" scenario).
type InternalErrorResponder[T any] interface {
	Internal() T
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
func RespondWithInternalError[T InternalErrorResponder[T]](
	err *sdkErrors.SDKError, w http.ResponseWriter, response T,
) *sdkErrors.SDKError {
	failErr := Fail(response.Internal(), w, http.StatusInternalServerError)
	if failErr != nil {
		if err != nil {
			return err.Wrap(failErr)
		}
		return failErr
	}
	return err
}

// Success sends a success response with HTTP 200 OK.
//
// This is a convenience wrapper around Fail that sends a 200 OK status.
// It maintains semantic clarity by using the name "Success" rather than
// calling Fail directly at call sites.
//
// Type Parameters:
//   - T: The response type to send to the client (e.g.,
//     reqres.ShardPutSuccess)
//
// Parameters:
//   - clientResponse: The response object to send to the client
//   - w: The HTTP response writer
//
// Returns:
//   - *sdkErrors.SDKError: An error if writing the response fails,
//     nil on success
//
// Example usage:
//
//	state.SetShard(request.Shard)
//	net.Success(reqres.ShardPutSuccess, w)
//	return nil
func Success[T any](
	clientResponse T, w http.ResponseWriter,
) *sdkErrors.SDKError {
	return Fail(clientResponse, w, http.StatusOK)
}

// SuccessWithResponseBody sends a success response with HTTP 200 OK and
// returns the response body for cleanup.
//
// This variant is used when the response body needs to be explicitly cleared
// from memory for security reasons, such as when returning sensitive
// cryptographic data. The caller is responsible for clearing the returned
// byte slice.
//
// Type Parameters:
//   - T: The response type to send to the client (e.g.,
//     reqres.ShardGetResponse)
//
// Parameters:
//   - clientResponse: The response object to send to the client
//   - w: The HTTP response writer
//
// Returns:
//   - []byte: The marshaled response body that should be cleared for security
//   - *sdkErrors.SDKError: An error if writing the response fails,
//     nil on success
//
// Example usage:
//
//	responseBody, err := net.SuccessWithResponseBody(
//	    reqres.ShardGetResponse{Shard: sh}.Success(), w,
//	)
//	if err != nil {
//	    return err
//	}
//	defer func() {
//	    mem.ClearBytes(responseBody)
//	}()
//	return nil
func SuccessWithResponseBody[T any](
	clientResponse T, w http.ResponseWriter,
) ([]byte, *sdkErrors.SDKError) {
	responseBody, marshalErr := MarshalBodyAndRespondOnMarshalFail(
		clientResponse, w,
	)

	if alreadyResponded := marshalErr != nil; alreadyResponded {
		// Headers already sent. Just return the response body.
		return responseBody, nil
	}

	respondErr := Respond(http.StatusOK, responseBody, w)
	if respondErr != nil {
		return nil, respondErr
	}
	return responseBody, nil
}

// UnmarshalAndRespondOnFail unmarshals a JSON request body into a typed
// request struct.
//
// This is a generic function that handles the common pattern of unmarshaling
// and validating incoming JSON requests. If unmarshaling fails, it sends the
// provided error response to the client with a 400 Bad Request status.
//
// Type Parameters:
//   - Req: The request type to unmarshal into
//   - Res: The response type for error cases
//
// Parameters:
//   - requestBody: The raw JSON request body to unmarshal
//   - w: The response writer for error handling
//   - errorResponseForBadRequest: A response object to send if unmarshaling
//     fails
//
// Returns:
//   - *Req: A pointer to the unmarshaled request struct, or nil if
//     unmarshaling failed
//   - *sdkErrors.SDKError: ErrDataUnmarshalFailure if unmarshaling fails, or
//     nil on success
//
// The function handles all error logging and response writing for the error
// case. Callers should check if the returned pointer is nil before proceeding.
func UnmarshalAndRespondOnFail[Req any, Res any](
	requestBody []byte,
	w http.ResponseWriter,
	errorResponseForBadRequest Res,
) (*Req, *sdkErrors.SDKError) {
	var request Req

	if unmarshalErr := json.Unmarshal(requestBody, &request); unmarshalErr != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(unmarshalErr)

		responseBodyForBadRequest, err := MarshalBodyAndRespondOnMarshalFail(
			errorResponseForBadRequest, w,
		)
		if noResponseSentYet := err == nil; noResponseSentYet {
			respondErr := Respond(http.StatusBadRequest, responseBodyForBadRequest, w)
			if respondErr != nil {
				failErr = failErr.Wrap(respondErr)
			}
		}

		// If marshal succeeded, we already responded with a 400 Bad Request with
		// the errorResponseForBadRequest.
		// Otherwise, if marshal failed (err != nil; very unlikely), we already
		// responded with a 400 Bad Request in MarshalBodyAndRespondOnMarshalFail.
		// Either way, we don't need to respond again. Just return the error.
		return nil, failErr
	}

	// We were able to unmarshal the request successfully.
	// We didn't send any failure response to the client so far.
	// Return a pointer to the request to be handled by the calling site.
	return &request, nil
}

// ExtractPeerSPIFFEIDAndRespondOnFail extracts and validates the peer
// SPIFFE ID from an HTTP request. If the SPIFFE ID cannot be extracted or is
// nil, it writes an unauthorized response using the provided error response
// object and returns an error.
//
// This function is generic and can be used with any response type that needs
// to be returned in case of authentication failure.
//
// Parameters:
//   - r *http.Request: The HTTP request containing peer SPIFFE ID
//   - w http.ResponseWriter: Response writer for error responses
//   - errorResponse T: The error response object to marshal and send if
//     validation fails
//
// Returns:
//   - *spiffeid.ID: The extracted SPIFFE ID if successful
//   - *sdkErrors.SDKError: ErrAccessUnauthorized if extraction fails or ID is
//     invalid, nil otherwise
//
// Example usage:
//
//	peerID, err := net.ExtractPeerSPIFFEIDAndRespondOnFail(
//	    r, w,
//	    reqres.ShardGetResponse{Err: data.ErrUnauthorized},
//	)
//	if err != nil {
//	    return err
//	}
func ExtractPeerSPIFFEIDAndRespondOnFail[T any](
	r *http.Request,
	w http.ResponseWriter,
	errorResponse T,
) (*spiffeid.ID, *sdkErrors.SDKError) {
	peerSPIFFEID, err := spiffe.IDFromRequest(r)
	if err != nil {
		failErr := sdkErrors.ErrSPIFFEFailedToExtractX509SVID.Wrap(err)

		responseBody, err := MarshalBodyAndRespondOnMarshalFail(
			errorResponse, w,
		)
		if notRespondedYet := err == nil; notRespondedYet {
			respondErr := Respond(http.StatusUnauthorized, responseBody, w)
			if respondErr != nil {
				failErr = failErr.Wrap(respondErr)
			}
		}

		notAuthorizedErr := sdkErrors.ErrAccessUnauthorized.Wrap(failErr)
		return nil, notAuthorizedErr
	}

	err = validation.ValidateSPIFFEID(peerSPIFFEID.String())
	if err != nil {
		failErr := sdkErrors.ErrSPIFFEInvalidSPIFFEID.Wrap(err)

		responseBody, err := MarshalBodyAndRespondOnMarshalFail(
			errorResponse, w,
		)
		if notRespondedYet := err == nil; notRespondedYet {
			respondErr := Respond(http.StatusUnauthorized, responseBody, w)
			if respondErr != nil {
				failErr = failErr.Wrap(respondErr)
			}
		}

		notAuthorizedErr := sdkErrors.ErrAccessUnauthorized.Wrap(failErr)
		return nil, notAuthorizedErr
	}

	return peerSPIFFEID, nil
}

// ReadRequestBodyAndRespondOnFail reads the entire request body from an HTTP
// request.
//
// On error, this function writes a 400 Bad Request status to the response
// writer and returns the error for propagation to the caller. If writing the
// error response fails, it returns a 500 Internal Server Error.
//
// Parameters:
//   - w: http.ResponseWriter - The response writer for error handling
//   - r: *http.Request - The incoming HTTP request
//
// Returns:
//   - []byte: The request body as a byte slice, or nil if reading failed
//   - *sdkErrors.SDKError: sdkErrors.ErrDataReadFailure if reading fails,
//     nil on success
func ReadRequestBodyAndRespondOnFail(
	w http.ResponseWriter, r *http.Request,
) ([]byte, *sdkErrors.SDKError) {
	body, err := RequestBody(r)
	if err != nil {
		failErr := sdkErrors.ErrDataReadFailure.Wrap(err)
		failErr.Msg = "problem reading request body"

		// do not send the wrapped error to the client as it may contain
		// error details that an attacker can use and exploit.
		failJSON, err := json.Marshal(sdkErrors.ErrDataReadFailure)
		if err != nil {
			// Cannot even parse a generic struct, this is an internal error.
			w.WriteHeader(http.StatusInternalServerError)
			_, writeErr := w.Write(failJSON)
			if writeErr != nil {
				// Cannot even write the error response, this is a critical error.
				failErr = failErr.Wrap(writeErr)
				failErr.Msg = "problem writing response"
			}

			return nil, failErr
		}

		w.WriteHeader(http.StatusBadRequest)
		_, writeErr := w.Write(failJSON)
		if writeErr != nil {
			failErr = failErr.Wrap(writeErr)
			failErr.Msg = "problem writing response"
			// Cannot even write the error response, this is a critical error.
			// We can only return the error at this point.
			return nil, failErr
		}

		return nil, failErr
	}

	return body, nil
}

// SPIFFEIDPredicate is a function type that validates a SPIFFE ID string.
// It returns true if the SPIFFE ID passes validation, false otherwise.
type SPIFFEIDPredicate func(string) bool

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
	predicateFn SPIFFEIDPredicate, failureResponse any,
	w http.ResponseWriter, r *http.Request,
) *sdkErrors.SDKError {
	peerSPIFFEID, err := ExtractPeerSPIFFEIDAndRespondOnFail(
		r, w, failureResponse,
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

// RespondFallbackWithStatus writes a fallback JSON response with the given HTTP
// status code and error code. It sets appropriate headers to prevent caching.
//
// This function is used when the primary response handling fails or when a
// generic error response needs to be sent.
//
// Parameters:
//   - w: The HTTP response writer
//   - status: The HTTP status code to return
//   - code: The error code to include in the response body
//
// Returns:
//   - *sdkErrors.SDKError: An error if marshaling or writing fails,
//     nil on success
func RespondFallbackWithStatus(
	w http.ResponseWriter, status int, code sdkErrors.ErrorCode,
) *sdkErrors.SDKError {
	body, err := MarshalBodyAndRespondOnMarshalFail(
		reqres.FallbackResponse{Err: code}, w,
	)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	// Add cache invalidation headers
	w.Header().Set(
		"Cache-Control",
		"no-store, no-cache, must-revalidate, private",
	)
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	w.WriteHeader(status)

	if _, err := w.Write(body); err != nil {
		failErr := sdkErrors.ErrAPIInternal.Wrap(err)
		return failErr
	}

	return nil
}
