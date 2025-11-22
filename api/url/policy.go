//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/config/env"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// PolicyCreate constructs the full API endpoint URL for creating policies.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the policy path to create a complete endpoint URL for creating new policies.
//
// Returns:
//   - string: The complete endpoint URL for policy creation requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := PolicyCreate()
func PolicyCreate() string {
	const fName = "PolicyCreate"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusPolicy))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus policy path"
		log.FatalErr(fName, *failErr)
	}
	return u
}

// PolicyList constructs the full API endpoint URL for listing policies.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the policy path and adds query parameters to specify the list action.
//
// Returns:
//   - string: The complete endpoint URL for policy list requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := PolicyList()
func PolicyList() string {
	const fName = "PolicyList"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusPolicy))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus policy path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionList))
	return u + "?" + params.Encode()
}

// PolicyDelete constructs the full API endpoint URL for deleting policies.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the policy path and adds query parameters to specify the delete action.
//
// Returns:
//   - string: The complete endpoint URL for policy deletion requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := PolicyDelete()
func PolicyDelete() string {
	const fName = "PolicyDelete"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusPolicy))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus policy path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionDelete))
	return u + "?" + params.Encode()
}

// PolicyGet constructs the full API endpoint URL for retrieving a policy.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the policy path and adds query parameters to specify the get action.
//
// Returns:
//   - string: The complete endpoint URL for policy retrieval requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := PolicyGet()
func PolicyGet() string {
	const fName = "PolicyGet"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusPolicy))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus policy path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionGet))
	return u + "?" + params.Encode()
}
