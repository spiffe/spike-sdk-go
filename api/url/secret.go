//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/api/internal/env"
)

// SecretGet returns the URL for getting a secret.
func SecretGet() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(KeyApiAction, string(ActionGet))
	return u + "?" + params.Encode()
}

// SecretPut returns the URL for putting a secret.
func SecretPut() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlSecrets),
	)
	return u
}

// SecretDelete returns the URL for deleting a secret.
func SecretDelete() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(KeyApiAction, string(ActionDelete))
	return u + "?" + params.Encode()
}

// SecretUndelete returns the URL for undeleting a secret.
func SecretUndelete() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(KeyApiAction, string(ActionUndelete))
	return u + "?" + params.Encode()
}

// SecretList returns the URL for listing secrets.
func SecretList() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlSecrets),
	)
	params := url.Values{}
	params.Add(KeyApiAction, string(ActionList))
	return u + "?" + params.Encode()
}

// SecretMetadataGet returns the URL for getting a secret metadata.
func SecretMetadataGet() string {
	u, _ := url.JoinPath(
		env.NexusApiRoot(),
		string(SpikeNexusUrlSecretsMetadata),
	)
	params := url.Values{}
	params.Add(KeyApiAction, string(ActionGet))
	return u + "?" + params.Encode()
}
