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

// CipherEncrypt constructs the full API endpoint URL for encryption requests.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the cipher encrypt path to create a complete endpoint URL for encrypting
// data.
//
// Returns:
//   - string: The complete endpoint URL for encryption requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := CipherEncrypt()
func CipherEncrypt() string {
	const fName = "CipherEncrypt"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusCipherEncrypt))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus cipher encrypt path"
		log.FatalErr(fName, *failErr)
	}
	return u
}

// CipherDecrypt constructs the full API endpoint URL for decryption requests.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the cipher decrypt path to create a complete endpoint URL for decrypting
// data.
//
// Returns:
//   - string: The complete endpoint URL for decryption requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := CipherDecrypt()
func CipherDecrypt() string {
	const fName = "CipherDecrypt"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusCipherDecrypt))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus cipher decrypt path"
		log.FatalErr(fName, *failErr)
	}
	return u
}
