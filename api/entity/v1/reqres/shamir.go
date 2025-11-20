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

func (s ShardPutResponse) Success() ShardPutResponse {
	s.Err = sdkErrors.ErrSuccess.Code
	return s
}
func (s ShardPutResponse) NotFound() ShardPutResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s ShardPutResponse) BadRequest() ShardPutResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s ShardPutResponse) Unauthorized() ShardPutResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s ShardPutResponse) Internal() ShardPutResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
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

func (s ShardGetResponse) Success() ShardGetResponse {
	s.Err = sdkErrors.ErrSuccess.Code
	return s
}
func (s ShardGetResponse) NotFound() ShardGetResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s ShardGetResponse) BadRequest() ShardGetResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s ShardGetResponse) Unauthorized() ShardGetResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s ShardGetResponse) Internal() ShardGetResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}
