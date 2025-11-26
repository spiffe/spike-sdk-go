//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// ShardPutRequest represents a request to submit a shard contribution.
// KeeperId specifies the identifier of the keeper responsible for the shard.
// Shard represents the shard data being contributed to the system.
// Version optionally specifies the version of the shard being submitted.
type ShardPutRequest struct {
	Shard *[32]byte `json:"shard"`
}

// ShardPutResponse represents the response structure for a shard
// contribution.
type ShardPutResponse struct {
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (r ShardPutResponse) Success() ShardPutResponse {
	r.Err = ""
	return r
}
func (r ShardPutResponse) NotFound() ShardPutResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrAPIResponseCodeInvalid)
	return r
}
func (r ShardPutResponse) BadRequest() ShardPutResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r ShardPutResponse) Unauthorized() ShardPutResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r ShardPutResponse) Internal() ShardPutResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r ShardPutResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}

// ShardGetRequest represents a request to get a Shamir shard.
type ShardGetRequest struct {
}

// ShardGetResponse represents the response that returns a Shamir shard.
// The struct includes the shard identifier and an associated error code.
type ShardGetResponse struct {
	Shard *[32]byte `json:"shard"`
	Err   sdkErrors.ErrorCode
}

func (r ShardGetResponse) Success() ShardGetResponse {
	r.Err = ""
	return r
}
func (r ShardGetResponse) NotFound() ShardGetResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r ShardGetResponse) BadRequest() ShardGetResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r ShardGetResponse) Unauthorized() ShardGetResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r ShardGetResponse) Internal() ShardGetResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r ShardGetResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}
