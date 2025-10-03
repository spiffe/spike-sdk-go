//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/api/internal/env"
)

// HealthStatus returns the full URL for the health status endpoint
func GetHealth() string {
	u, _ := url.JoinPath(
		env.NexusAPIRoot(),
		string(NexusHealth),
	)
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionGet))
	return u + "?" + params.Encode()

}
