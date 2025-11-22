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

// Restore constructs the full API endpoint URL for operator restore requests.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the operator restore path to create a complete endpoint URL for submitting
// recovery shards during the restoration process.
//
// Returns:
//   - string: The complete endpoint URL for operator restore requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := Restore()
func Restore() string {
	const fName = "Restore"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusOperatorRestore))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus operator restore path"
		log.FatalErr(fName, *failErr)
	}
	return u
}

// Recover constructs the full API endpoint URL for operator recover requests.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the operator recover path to create a complete endpoint URL for initiating
// the recovery process and retrieving recovery shards.
//
// Returns:
//   - string: The complete endpoint URL for operator recover requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := Recover()
func Recover() string {
	const fName = "Recover"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusOperatorRecover))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus operator recover path"
		log.FatalErr(fName, *failErr)
	}
	return u
}
