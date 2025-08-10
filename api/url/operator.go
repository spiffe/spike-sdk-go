//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/api/internal/env"
)

func Restore() string {
	u, _ := url.JoinPath(
		env.NexusAPIRoot(),
		string(SpikeNexusURLOperatorRestore),
	)
	return u
}

func Recover() string {
	u, _ := url.JoinPath(
		env.NexusAPIRoot(),
		string(SpikeNexusURLOperatorRecover),
	)
	return u
}
