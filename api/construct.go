//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/config/env"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

// API is the SPIKE API.
type API struct {
	source *workloadapi.X509Source
}

// New creates and returns a new instance of API configured with a SPIFFE
// source.
// It automatically discovers and connects to the SPIFFE Workload API endpoint
// using the default socket path and creates an X.509 source for authentication
// with a configurable timeout to prevent indefinite blocking on socket issues.
//
// The timeout can be configured using the SPIKE_SPIFFE_SOURCE_TIMEOUT
// environment variable (default: 30s).
//
// The API client is configured to communicate exclusively with SPIKE Nexus.
//
// Returns:
//   - *API: A configured API instance ready for use, or nil if initialization
//     fails
//
// The function will return nil if:
//   - The SPIFFE Workload API is not available
//   - The default endpoint socket cannot be accessed
//   - The X.509 source creation fails or times out
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

	ctx, cancel := context.WithTimeout(
		context.Background(),
		env.SPIFFESourceTimeoutVal(),
	)
	defer cancel()

	source, _, err := spiffe.Source(ctx, defaultEndpointSocket)
	if err != nil {
		return nil
	}

	return &API{source: source}
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
	}
}
