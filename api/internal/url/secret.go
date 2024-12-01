//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/api/internal/env"
)

// UrlSecretGet returns the URL for getting a secret.
func SecretGet() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(spikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusGet))
	return u + "?" + params.Encode()
}

// UrlSecretPut returns the URL for putting a secret.
func SecretPut() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(spikeNexusUrlSecrets),
	)
	return u
}

// UrlSecretDelete returns the URL for deleting a secret.
func SecretDelete() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(spikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusDelete))
	return u + "?" + params.Encode()
}

// UrlSecretUndelete returns the URL for undeleting a secret.
func SecretUndelete() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(spikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusUndelete))
	return u + "?" + params.Encode()
}

// UrlSecretList returns the URL for listing secrets.
func SecretList() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(spikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusList))
	return u + "?" + params.Encode()
}

// UrlSecretMetadataGet returns the URL for getting a secret metadata.
func SecretMetadataGet() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(spikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(keyApiAction, string(actionNexusGetMetadata))
	return u + "?" + params.Encode()
}
