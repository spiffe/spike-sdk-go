//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/errors"
)

// FallbackResponse is a generic response for any error.
type FallbackResponse struct {
	Err errors.ErrorCode `json:"err,omitempty"`
}
