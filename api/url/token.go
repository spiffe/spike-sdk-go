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

// Init constructs the full API endpoint URL for initializing SPIKE Nexus.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the init path to create a complete endpoint URL for initializing SPIKE Nexus
// with the root encryption key.
//
// Returns:
//   - string: The complete endpoint URL for initialization requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := Init()
func Init() string {
	const fName = "Init"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusInit))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus init path"
		log.FatalErr(fName, *failErr)
	}
	return u
}

// InitState constructs the full API endpoint URL for checking the
// initialization state of SPIKE Nexus.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the init path and adds query parameters to specify the check action for
// verifying whether SPIKE Nexus has been initialized.
//
// Returns:
//   - string: The complete endpoint URL for initialization state check requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := InitState()
func InitState() string {
	const fName = "InitState"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusInit))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus init path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionCheck))
	return u + "?" + params.Encode()
}
