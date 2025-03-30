//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/api/internal/env"
)

// Init returns the URL for initializing SPIKE Nexus.
func Init() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlInit),
	)
	return u
}

// InitState returns the URL for checking the initialization state of
// SPIKE Nexus.
func InitState() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlInit),
	)
	params := url.Values{}
	params.Add(KeyApiAction, string(ActionCheck))
	return u + "?" + params.Encode()
}
