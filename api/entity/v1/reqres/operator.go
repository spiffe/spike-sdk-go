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
func (r RestoreResponse) NotFound() RestoreResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return r
}
func (r RestoreResponse) BadRequest() RestoreResponse {
	r.Err = sdkErrors.ErrBadRequest.Code
	return r
}
func (r RestoreResponse) Unauthorized() RestoreResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r RestoreResponse) Internal() RestoreResponse {
	r.Err = sdkErrors.ErrInternal.Code
	return r
}
func (r RestoreResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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
func (r RecoverResponse) NotFound() RecoverResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return r
}
func (r RecoverResponse) BadRequest() RecoverResponse {
	r.Err = sdkErrors.ErrBadRequest.Code
	return r
}
func (r RecoverResponse) Unauthorized() RecoverResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r RecoverResponse) Internal() RecoverResponse {
	r.Err = sdkErrors.ErrInternal.Code
	return r
}
func (r RecoverResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}
