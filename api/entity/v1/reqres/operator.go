//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

type RestoreRequest struct {
	Shard string `json:"shard"`
}
type RestoreResponse struct {
	data.RestorationStatus
	Err data.ErrorCode `json:"err,omitempty"`
}

type RecoverRequest struct {
}

type RecoverResponse struct {
	Shards []string       `json:"shards"`
	Err    data.ErrorCode `json:"err,omitempty"`
}
