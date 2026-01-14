//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

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
