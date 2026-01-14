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

// ExtractPeerSPIFFEIDFromRequestAndRespondOnFail extracts and validates
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
func ExtractPeerSPIFFEIDFromRequestAndRespondOnFail[T any](
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
