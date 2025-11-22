//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// KeeperBootstrapContributeEndpoint constructs the full API endpoint URL for
// SPIKE Keeper contribution requests.
//
// It joins the provided keeper API root URL with the KeeperContribute path
// segment to create a complete endpoint URL for submitting secret shares to
// keepers.
//
// Parameters:
//   - keeperAPIRoot: The base URL of the SPIKE Keeper API
//
// Returns:
//   - string: The complete endpoint URL for keeper contribution requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := KeeperBootstrapContributeEndpoint("https://keeper.example.com")
func KeeperBootstrapContributeEndpoint(keeperAPIRoot string) string {
	const fName = "KeeperBootstrapContributeEndpoint"

	u, err := url.JoinPath(
		keeperAPIRoot, string(KeeperContribute),
	)
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Keeper API path"
		log.FatalErr(fName, *failErr)
	}
	return u
}

// NexusBootstrapVerifyEndpoint constructs the full API endpoint URL for
// bootstrap verification requests.
//
// It joins the provided Nexus API root URL with the bootstrap verify path to
// create a complete endpoint URL for verifying that SPIKE Nexus has been
// properly initialized with the root key.
//
// Parameters:
//   - nexusAPIRoot: The base URL of the SPIKE Nexus API
//
// Returns:
//   - string: The complete endpoint URL for bootstrap verification requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := NexusBootstrapVerifyEndpoint("https://nexus.example.com")
func NexusBootstrapVerifyEndpoint(nexusAPIRoot string) string {
	const fName = "NexusBootstrapVerifyEndpoint"

	u, err := url.JoinPath(nexusAPIRoot, string(NexusBootstrapVerify))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus API path"
		log.FatalErr(fName, *failErr)
	}
	return u
}
