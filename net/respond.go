//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
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
