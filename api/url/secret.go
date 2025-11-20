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

// SecretGet constructs the full API endpoint URL for retrieving secrets.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the secrets path and adds query parameters to specify the get action.
//
// Returns:
//   - string: The complete endpoint URL for secret retrieval requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := SecretGet()
func SecretGet() string {
	const fName = "SecretGet"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusSecrets))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus secrets path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionGet))
	return u + "?" + params.Encode()
}

// SecretPut constructs the full API endpoint URL for creating or updating
// secrets.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the secrets path to create a complete endpoint URL for storing secrets.
//
// Returns:
//   - string: The complete endpoint URL for secret creation/update requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := SecretPut()
func SecretPut() string {
	const fName = "SecretPut"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusSecrets))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus secrets path"
		log.FatalErr(fName, *failErr)
	}
	return u
}

// SecretDelete constructs the full API endpoint URL for deleting secrets.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the secrets path and adds query parameters to specify the delete action.
//
// Returns:
//   - string: The complete endpoint URL for secret deletion requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := SecretDelete()
func SecretDelete() string {
	const fName = "SecretDelete"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusSecrets))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus secrets path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionDelete))
	return u + "?" + params.Encode()
}

// SecretUndelete constructs the full API endpoint URL for restoring deleted
// secrets.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the secrets path and adds query parameters to specify the undelete action.
//
// Returns:
//   - string: The complete endpoint URL for secret restoration requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := SecretUndelete()
func SecretUndelete() string {
	const fName = "SecretUndelete"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusSecrets))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus secrets path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionUndelete))
	return u + "?" + params.Encode()
}

// SecretList constructs the full API endpoint URL for listing secrets.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the secrets path and adds query parameters to specify the list action.
//
// Returns:
//   - string: The complete endpoint URL for secret list requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := SecretList()
func SecretList() string {
	const fName = "SecretList"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusSecrets))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus secrets path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionList))
	return u + "?" + params.Encode()
}

// SecretMetadataGet constructs the full API endpoint URL for retrieving secret
// metadata.
//
// It joins the SPIKE Nexus API root URL (from environment configuration) with
// the secrets metadata path and adds query parameters to specify the get action.
//
// Returns:
//   - string: The complete endpoint URL for secret metadata retrieval requests
//
// Note: The function will fatally crash (via log.FatalErr) if URL path joining
// fails.
//
// Example:
//
//	endpoint := SecretMetadataGet()
func SecretMetadataGet() string {
	const fName = "SecretMetadataGet"

	u, err := url.JoinPath(env.NexusAPIRootVal(), string(NexusSecretsMetadata))
	if err != nil {
		failErr := sdkErrors.ErrNetURLJoinPathFailed.Wrap(err)
		failErr.Msg = "failed to join SPIKE Nexus secrets metadata path"
		log.FatalErr(fName, *failErr)
	}
	params := url.Values{}
	params.Add(KeyAPIAction, string(ActionGet))
	return u + "?" + params.Encode()
}
