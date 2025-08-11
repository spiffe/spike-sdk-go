//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package mock

import (
	"context"
)

// Retrier implements Retrier for testing
type Retrier struct {
	RetryFunc func(context.Context, func() error) error
}

// RetryWithBackoff implements the Retrier interface
func (m *Retrier) RetryWithBackoff(
	ctx context.Context,
	operation func() error,
) error {
	if m.RetryFunc != nil {
		return m.RetryFunc(ctx, operation)
	}
	return nil
}
