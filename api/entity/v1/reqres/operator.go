//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// RestoreRequest for disaster recovery.
type RestoreRequest struct {
	ID    int       `json:"id"`
	Shard *[32]byte `json:"shard"`
}

// RestoreResponse for disaster recovery.
type RestoreResponse struct {
	data.RestorationStatus
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (r RestoreResponse) Success() RestoreResponse {
	r.Err = sdkErrors.ErrSuccess.Code
	return r
}
func (s RestoreResponse) NotFound() RestoreResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s RestoreResponse) BadRequest() RestoreResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s RestoreResponse) Unauthorized() RestoreResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s RestoreResponse) Internal() RestoreResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}

// RecoverRequest for disaster recovery.
type RecoverRequest struct {
}

// RecoverResponse for disaster recovery.
type RecoverResponse struct {
	Shards map[int]*[32]byte   `json:"shards"`
	Err    sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (r RecoverResponse) Success() RecoverResponse {
	r.Err = sdkErrors.ErrSuccess.Code
	return r
}
func (s RecoverResponse) NotFound() RecoverResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s RecoverResponse) BadRequest() RecoverResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s RecoverResponse) Unauthorized() RecoverResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s RecoverResponse) Internal() RecoverResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}
