//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffe

import (
	"context"
	"net/http"
	"os"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"

	"github.com/spiffe/spike-sdk-go/config/env"
)

// EndpointSocket returns the UNIX domain socket address for the SPIFFE
// Workload API endpoint.
//
// The function first checks for the SPIFFE_ENDPOINT_SOCKET environment
// variable. If set, it returns that value. Otherwise, it returns a default
// development
//
//	socket path:
//
// "unix:///tmp/spire-agent/public/api.sock"
//
// For production deployments, especially in Kubernetes environments, it's
// recommended to set SPIFFE_ENDPOINT_SOCKET to a more restricted socket path,
// such as: "unix:///run/spire/agent/sockets/spire.sock"
//
// Default socket paths by environment:
//   - Development (Linux): unix:///tmp/spire-agent/public/api.sock
//   - Kubernetes: unix:///run/spire/agent/sockets/spire.sock
//
// Returns:
//   - string: The UNIX domain socket address for the SPIFFE Workload API
//     endpoint
//
// Environment Variables:
//   - SPIFFE_ENDPOINT_SOCKET: Override the default socket path
func EndpointSocket() string {
	p := os.Getenv(env.SPIFFEEndpointSocket)
	if p != "" {
		return p
	}

	return "unix:///tmp/spire-agent/public/api.sock"
}

// Source creates a new SPIFFE X.509 source and returns the associated SVID ID.
// It establishes a connection to the Workload API at the specified socket path
// and retrieves the X.509 SVID for the workload.
//
// The returned X509Source should be closed when no longer needed.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - socketPath: The Workload API endpoint location
//     (e.g., "unix:///path/to/socket")
//
// Returns:
//   - *workloadapi.X509Source: An X509Source that can be used to fetch and
//     monitor X.509 SVIDs
//   - string: The string representation of the current SVID ID
//   - *sdkErrors.SDKError: ErrSPIFFEFailedToCreateX509Source if source creation
//     fails, or ErrSPIFFEUnableToFetchX509Source if initial SVID fetch fails
func Source(ctx context.Context, socketPath string) (
	*workloadapi.X509Source, string, *sdkErrors.SDKError,
) {
	source, err := workloadapi.NewX509Source(ctx,
		workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)))

	if err != nil {
		return nil, "", sdkErrors.ErrSPIFFEFailedToCreateX509Source.Wrap(err)
	}

	sv, err := source.GetX509SVID()
	if err != nil {
		return nil, "", sdkErrors.ErrSPIFFEUnableToFetchX509Source.Wrap(err)
	}

	return source, sv.ID.String(), nil
}

// IDFromRequest extracts the SPIFFE ID from the TLS peer certificate of
// an HTTP request. It checks if the incoming request has a valid TLS connection
// and at least one peer certificate. The first certificate in the chain is used
// to extract the SPIFFE ID.
//
// Note: This function assumes that the request is already over a secured TLS
// connection and will fail if the TLS connection state is not available or
// the peer certificates are missing.
//
// Parameters:
//   - r: The HTTP request from which the SPIFFE ID is to be extracted
//
// Returns:
//   - *spiffeid.ID: The SPIFFE ID extracted from the first peer certificate,
//     or nil if extraction fails
//   - *sdkErrors.SDKError: ErrSPIFFENoPeerCertificates if peer certificates are
//     absent, or ErrSPIFFEFailedToExtractX509SVID if extraction fails
func IDFromRequest(r *http.Request) (*spiffeid.ID, *sdkErrors.SDKError) {
	tlsConnectionState := r.TLS
	if len(tlsConnectionState.PeerCertificates) == 0 {
		return nil, sdkErrors.ErrSPIFFENoPeerCertificates
	}

	id, err := x509svid.IDFromCert(tlsConnectionState.PeerCertificates[0])
	if err != nil {
		return nil, sdkErrors.ErrSPIFFEFailedToExtractX509SVID.Wrap(err)
	}

	return &id, nil
}

// CloseSource safely closes an X509Source.
//
// This function should be called when the X509Source is no longer needed,
// typically during application shutdown or cleanup. It handles nil sources
// gracefully.
//
// Parameters:
//   - source: The X509Source to close, may be nil
//
// Returns:
//   - *sdkErrors.SDKError: nil if successful or source is nil,
//     ErrSPIFFEFailedToCloseX509Source if closure fails
func CloseSource(source *workloadapi.X509Source) *sdkErrors.SDKError {
	if source == nil {
		return nil
	}

	if err := source.Close(); err != nil {
		return sdkErrors.ErrSPIFFEFailedToCloseX509Source.Wrap(err)
	}

	return nil
}
