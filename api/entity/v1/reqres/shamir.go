//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

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
	Err data.ErrorCode `json:"err,omitempty"`
}

func (s ShardPutResponse) Success() ShardPutResponse {
	s.Err = data.ErrSuccess
	return s
}

// ShardGetRequest represents a request to get a Shamir shard.
type ShardGetRequest struct {
}

// ShardGetResponse represents the response that returns a Shamir shard.
// The struct includes the shard identifier and an associated error code.
type ShardGetResponse struct {
	Shard *[32]byte `json:"shard"`
	Err   data.ErrorCode
}

func (s ShardGetResponse) Success() ShardGetResponse {
	s.Err = data.ErrSuccess
	return s
}
