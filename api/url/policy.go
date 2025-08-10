//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
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
		env.NexusAPIRoot(),
		string(NexusURLPolicy),
	)
	return u
}

// PolicyList returns the URL for listing policies.
func PolicyList() string {
	u, _ := url.JoinPath(
		env.NexusAPIRoot(),
		string(NexusURLPolicy),
	)
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionList))
	return u + "?" + params.Encode()
}

// PolicyDelete returns the URL for deleting a policy.
func PolicyDelete() string {
	u, _ := url.JoinPath(
		env.NexusAPIRoot(),
		string(NexusURLPolicy),
	)
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionDelete))
	return u + "?" + params.Encode()
}

// PolicyGet returns the URL for getting a policy.
func PolicyGet() string {
	u, _ := url.JoinPath(
		env.NexusAPIRoot(),
		string(NexusURLPolicy),
	)
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionGet))
	return u + "?" + params.Encode()
}
