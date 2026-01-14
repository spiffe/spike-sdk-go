//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

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
