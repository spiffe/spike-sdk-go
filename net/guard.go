//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// GuardFunc is a function type for request guard/validation functions.
// Guard functions validate requests and return an error if validation fails.
// They typically check authentication, authorization, and input validation.
//
// Type Parameters:
//   - Req: The request type to validate
//
// Parameters:
//   - request: The request to validate
//   - w: http.ResponseWriter for writing error responses
//   - r: *http.Request for accessing request context
//
// Returns:
//   - *sdkErrors.SDKError: nil if validation passes, error otherwise
type GuardFunc[Req any] func(
	Req, http.ResponseWriter, *http.Request,
) *sdkErrors.SDKError

// ReadParseAndGuard reads the HTTP request body, parses it, and executes
// a guard function in a single operation. This function combines
// readAndParseRequest with guard execution to further reduce boilerplate.
//
// This function performs the following steps:
//  1. Reads the request body from the HTTP request
//  2. Unmarshals the body into the request type
//  3. Executes the guard function for validation
//  4. Returns the parsed request and any errors
//
// Type Parameters:
//   - Req: The request type to unmarshal into
//   - Res: The response type for error cases
//
// Parameters:
//   - w: The response writer for error handling
//   - r: The incoming HTTP request
//   - errorResponse: A response object to send if parsing fails
//   - guard: The guard function to execute for validation
//
// Returns:
//   - *Req: A pointer to the parsed request struct, or nil if any step failed
//   - *sdkErrors.SDKError: ErrDataReadFailure, ErrDataParseFailure, or error
//     from the guard function
//
// Example usage:
//
//	request, err := net.ReadParseAndGuard[
//	    reqres.ShardPutRequest,
//	    reqres.ShardPutResponse](
//	    w, r,
//	    reqres.ShardPutResponse{Err: data.ErrBadInput},
//	    guardShardPutRequest,
//	)
//	if err != nil {
//	    return err
//	}
func ReadParseAndGuard[Req any, Res any](
	w http.ResponseWriter, r *http.Request, errorResponse Res,
	guard GuardFunc[Req],
) (*Req, *sdkErrors.SDKError) {
	request, err := readAndParseRequest[Req, Res](w, r, errorResponse)
	if err != nil {
		return nil, err
	}

	if err = guard(*request, w, r); err != nil {
		return nil, err
	}

	return request, nil
}
