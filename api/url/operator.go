//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/config/env"
)

// Restore returns the URL for operator's restore endpoint.
func Restore() string {
	u, _ := url.JoinPath(
		env.NexusAPIRootVal(),
		string(NexusOperatorRestore),
	)
	return u
}

// Recover returns the URL for operator's recover endpoint.
func Recover() string {
	u, _ := url.JoinPath(
		env.NexusAPIRootVal(),
		string(NexusOperatorRecover),
	)
	return u
}
