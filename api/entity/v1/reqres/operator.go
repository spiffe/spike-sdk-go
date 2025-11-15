//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

// RestoreRequest for disaster recovery.
type RestoreRequest struct {
	ID    int       `json:"id"`
	Shard *[32]byte `json:"shard"`
}

// RestoreResponse for disaster recovery.
type RestoreResponse struct {
	data.RestorationStatus
	Err data.ErrorCode `json:"err,omitempty"`
}

func (r RestoreResponse) Success() RestoreResponse {
	r.Err = data.ErrSuccess
	return r
}

// RecoverRequest for disaster recovery.
type RecoverRequest struct {
}

// RecoverResponse for disaster recovery.
type RecoverResponse struct {
	Shards map[int]*[32]byte `json:"shards"`
	Err    data.ErrorCode    `json:"err,omitempty"`
}

func (r RecoverResponse) Success() RecoverResponse {
	r.Err = data.ErrSuccess
	return r
}
