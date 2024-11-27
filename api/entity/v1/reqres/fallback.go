//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
)

// FallbackResponse is a generic response for any error.
type FallbackResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}
