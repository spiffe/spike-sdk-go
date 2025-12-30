//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/config/env"
	"github.com/spiffe/spike-sdk-go/predicate"
)

// CreateMTLSClientWithPredicate creates an HTTP client configured for
// mutual TLS authentication using SPIFFE workload identities.
//
// Parameters:
//   - source: An X509Source that provides:
//   - The client's own identity certificate (presented to servers)
//   - Trusted roots for validating server certificates
//   - predicate: A function that validates SERVER (peer) SPIFFE IDs.
//     Returns true if the SERVER's ID should be trusted.
//     NOTE: This predicate checks the SERVER's identity, NOT the client's.
//
// Returns:
//   - *http.Client: A configured HTTP client that will use mTLS for all
//     connections
//
// The returned client will:
//   - Present its own client certificate from the X509Source to servers
//   - Validate server certificates using the same X509Source's trust bundle
//   - Only accept connections to servers whose SPIFFE IDs pass the predicate
//
// Example:
//
//	// This predicate allows the client to connect only to servers with
//	// SPIFFE IDs in the "backend" service namespace
//	client := CreateMTLSClientWithPredicate(source,
//	 func(serverID string) bool {
//	    return strings.Contains(serverID, "/ns/backend/")
//	})
func CreateMTLSClientWithPredicate(
	source *workloadapi.X509Source,
	predicate predicate.Predicate,
) *http.Client {
	authorizer := AuthorizerWithPredicate(predicate)
	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     tlsConfig,
			IdleConnTimeout:     env.HTTPClientIdleConnTimeoutVal(),
			MaxIdleConns:        env.HTTPClientMaxIdleConnsVal(),
			MaxConnsPerHost:     env.HTTPClientMaxConnsPerHostVal(),
			MaxIdleConnsPerHost: env.HTTPClientMaxIdleConnsPerHostVal(),
			DialContext: (&net.Dialer{
				Timeout:   env.HTTPClientDialerTimeoutVal(),
				KeepAlive: env.HTTPClientDialerKeepAliveVal(),
			}).DialContext,
			TLSHandshakeTimeout:   env.HTTPClientTLSHandshakeTimeoutVal(),
			ResponseHeaderTimeout: env.HTTPClientResponseHeaderTimeoutVal(),
			ExpectContinueTimeout: env.HTTPClientExpectContinueTimeoutVal(),
		},
		Timeout: env.HTTPClientTimeoutVal(),
	}

	return client
}

// CreateMTLSClient creates an HTTP client configured for mutual TLS
// authentication using SPIFFE workload identities.
//
// WARNING: This function accepts ALL server SPIFFE IDs without validation.
// For production use, consider using CreateMTLSClientWithPredicate to restrict
// which servers this client will connect to for better security.
//
// Parameters:
//   - source: An X509Source that provides the client's identity certificates
//     and trusted roots
//
// Returns:
//   - *http.Client: A configured HTTP client that will use mTLS for all
//     connections
//
// The returned client will:
//   - Present client certificates from the provided X509Source
//   - Validate server certificates using the same X509Source
//   - Accept connections to ANY server with a valid SPIFFE certificate
func CreateMTLSClient(source *workloadapi.X509Source) *http.Client {
	return CreateMTLSClientWithPredicate(source, predicate.AllowAll)
}

// CreateMTLSClientForNexus creates an HTTP client configured for mutual TLS
// authentication with SPIKE Nexus using the provided X509Source. The client
// is configured with a predicate that validates peer IDs against the trusted
// Nexus root. Only peers that pass the spiffeid.IsNexus validation will be
// accepted for connections.
//
// Parameters:
//   - source: An X509Source that provides the client's identity certificates
//     and trusted roots
//
// Returns:
//   - *http.Client: A configured HTTP client for connecting to SPIKE Nexus
func CreateMTLSClientForNexus(source *workloadapi.X509Source) *http.Client {
	return CreateMTLSClientWithPredicate(source, predicate.AllowNexus)
}

// CreateMTLSClientForKeeper creates an HTTP client configured for mutual
// TLS authentication using the provided X509Source. The client is configured
// with a predicate that validates peer IDs against the trusted keeper root.
// Only peers that pass the spiffeid.IsKeeper validation will be accepted for
// connections.
//
// Parameters:
//   - source: An X509Source that provides the client's identity certificates
//     and trusted roots
//
// Returns:
//   - *http.Client: A configured HTTP client for connecting to SPIKE Keeper
func CreateMTLSClientForKeeper(source *workloadapi.X509Source) *http.Client {
	return CreateMTLSClientWithPredicate(source, predicate.AllowKeeper)
}
