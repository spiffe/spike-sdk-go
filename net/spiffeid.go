//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/spiffe"
	"github.com/spiffe/spike-sdk-go/validation"
)

// ExtractPeerSPIFFEIDAndRespondOnFail extracts and validates
// the peer SPIFFE ID from an HTTP request. If the SPIFFE ID cannot be extracted
// or is nil, it writes an unauthorized response using the provided error
// response object and returns an error.
//
// Type Parameters:
//   - T: The type of the error response object. This can be any type that is
//     JSON-serializable, typically a struct representing the API error response
//     format expected by clients.
//
// Parameters:
//   - w http.ResponseWriter: Response writer for error responses
//   - r *http.Request: The HTTP request containing peer SPIFFE ID
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
//	peerID, err := auth.ExtractPeerSPIFFEID(
//	    r, w,
//	    reqres.ShardGetResponse{Err: data.ErrUnauthorized},
//	)
//	if err != nil {
//	    return err
//	}
func ExtractPeerSPIFFEIDAndRespondOnFail[T any](
	w http.ResponseWriter, r *http.Request,
	errorResponse T,
) (*spiffeid.ID, *sdkErrors.SDKError) {
	peerSPIFFEID, err := spiffe.IDFromRequest(r)
	if err != nil {
		failErr := sdkErrors.ErrSPIFFEFailedToExtractX509SVID.Wrap(err)

		responseBody, err := MarshalBodyAndRespondOnMarshalFail(
			errorResponse, w,
		)

		e := RespondUnauthorizedAndWrapError(err, failErr, w, responseBody)

		return nil, e
	}

	err = validation.ValidateSPIFFEID(peerSPIFFEID.String())
	if err != nil {
		failErr := sdkErrors.ErrSPIFFEInvalidSPIFFEID.Wrap(err)

		responseBody, err := MarshalBodyAndRespondOnMarshalFail(
			errorResponse, w,
		)

		e := RespondUnauthorizedAndWrapError(err, failErr, w, responseBody)

		return nil, e
	}

	return peerSPIFFEID, nil
}

// RespondErrOnBadSPIFFEIDPattern validates the given SPIFFE ID pattern and
// writes an error response if the validation fails.
//
// This function checks if the SPIFFE ID pattern conforms to the expected format
// using validation.ValidateSPIFFEIDPattern. The pattern may include regex
// special characters for matching multiple SPIFFE IDs. If validation fails, it
// sends the provided error response to the client with a 400 Bad Request status
// code.
//
// Type Parameters:
//   - T: The response type to send to the client in case of validation failure
//
// Parameters:
//   - SPIFFEIDPattern: string - The SPIFFE ID pattern to validate
//   - badInputResp: T - The error response object to send if validation fails
//   - w: http.ResponseWriter - The response writer for error handling
//
// Returns:
//   - *sdkErrors.SDKError: ErrDataInvalidInput if the pattern is invalid,
//     nil if validation succeeds
func RespondErrOnBadSPIFFEIDPattern[T any](
	SPIFFEIDPattern string, badInputResp T, w http.ResponseWriter,
) *sdkErrors.SDKError {
	if err := validation.ValidateSPIFFEIDPattern(SPIFFEIDPattern); err != nil {
		failEr := Fail(badInputResp, w, http.StatusBadRequest)
		if failEr != nil {
			return sdkErrors.ErrDataInvalidInput.Wrap(failEr)
		}
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	return nil
}
