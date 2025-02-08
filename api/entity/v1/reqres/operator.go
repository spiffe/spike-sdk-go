//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

type RestoreRequest struct{}
type RestoreResponse struct {
	Shards []string       `json:"shards"`
	Err    data.ErrorCode `json:"err,omitempty"`
}

type RecoverRequest struct {
	Shard string `json:"shard"`
}
type RecoverResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}
