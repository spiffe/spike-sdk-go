//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/internal/impl/cipher"
	"github.com/spiffe/spike-sdk-go/predicate"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

// indirection for testability: allows stubbing cipher calls in unit tests
var (
	cipherEncrypt = cipher.Encrypt
	cipherDecrypt = cipher.Decrypt
)

// API is the SPIKE API.
type API struct {
	source    *workloadapi.X509Source
	predicate predicate.Predicate
}

// New creates and returns a new instance of API configured with a SPIFFE source.
// It automatically discovers and connects to the SPIFFE Workload API endpoint
// using the default socket path and creates an X.509 source for authentication.
// The API client is configured to communicate exclusively with SPIKE Nexus.
//
// Returns:
//   - *API: A configured API instance ready for use, or nil if initialization
//     fails
//
// The function will return nil if:
//   - The SPIFFE Workload API is not available
//   - The default endpoint socket cannot be accessed
//   - The X.509 source creation fails
//
// Example usage:
//
//	// Create API client that connects to SPIKE Nexus
//	api := New()
//	if api == nil {
//	    log.Fatal("Failed to initialize SPIKE API")
//	}
//	defer api.Close()
func New() *API {
	defaultEndpointSocket := spiffe.EndpointSocket()

	source, _, err := spiffe.Source(
		context.Background(), defaultEndpointSocket,
	)
	if err != nil {
		return nil
	}

	// API Client can only talk to SPIKE Nexus as a peer.
	return &API{source: source, predicate: predicate.AllowNexus}
}

// NewWithSource initializes a new API instance with a pre-configured
// X509Source. This constructor is useful when you already have an X.509 source
// or need custom source configuration. The API instance will be configured to
// only communicate with SPIKE Nexus servers.
//
// Parameters:
//   - source: A pre-configured X509Source that provides the client's identity
//     certificates and trusted roots for server validation
//
// Returns:
//   - *API: A configured API instance using the provided source
//
// Note: The API client created with this function is restricted to communicate
// only with SPIKE Nexus instances (using predicate.AllowNexus). If you need
// to connect to different servers, use New() with a custom predicate instead.
//
// Example usage:
//
//		// Use with custom-configured source
//		source, err := workloadapi.NewX509Source(ctx,
//	 	workloadapi.WithClientOptions(...))
//		if err != nil {
//		    log.Fatal("Failed to create X509Source")
//		}
//		api := NewWithSource(source)
//		defer api.Close()
func NewWithSource(source *workloadapi.X509Source) *API {
	return &API{
		source: source,
		// API Client can only talk to SPIKE Nexus as a peer.
		predicate: predicate.AllowNexus,
	}
}
