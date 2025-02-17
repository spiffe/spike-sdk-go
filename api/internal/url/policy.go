//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/api/internal/env"
)

// PolicyCreate returns the URL for creating a policy.
func PolicyCreate() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlPolicy),
	)
	return u
}

// PolicyList returns the URL for listing policies.
func PolicyList() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlPolicy),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusList))
	return u + "?" + params.Encode()
}

// PolicyDelete returns the URL for deleting a policy.
func PolicyDelete() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlPolicy),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusDelete))
	return u + "?" + params.Encode()
}

// PolicyGet returns the URL for getting a policy.
func PolicyGet() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlPolicy),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusGet))
	return u + "?" + params.Encode()
}
