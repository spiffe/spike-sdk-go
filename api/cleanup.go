//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import "github.com/spiffe/spike-sdk-go/spiffe"

// Close releases any resources held by the API instance.
// It ensures proper cleanup of the underlying source.
func (a *API) Close() {
	if a.source == nil {
		return
	}
	spiffe.CloseSource(a.source)
}
