package spiffe

import (
	"context"
	"errors"
	"os"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
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
	p := os.Getenv("SPIFFE_ENDPOINT_SOCKET")
	if p != "" {
		return p
	}

	return "unix:///tmp/spire-agent/public/api.sock"
}

// Source creates a new SPIFFE X.509 source and returns the associated SVID ID.
// It establishes a connection to the Workload API at the specified socket path
// and retrieves the X.509 SVID for the workload.
//
// The function takes a context for cancellation and timeout control, and a
// socket path string specifying the Workload API endpoint location.
//
// It returns:
//   - An X509Source that can be used to fetch and monitor X.509 SVIDs
//   - The string representation of the current SVID ID
//   - An error if the source creation or initial SVID fetch fails
//
// The returned X509Source should be closed when no longer needed.
func Source(ctx context.Context, socketPath string) (
	*workloadapi.X509Source, string, error,
) {
	source, err := workloadapi.NewX509Source(ctx,
		workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)))

	if err != nil {
		return nil, "", errors.Join(
			errors.New("failed to create X509Source"),
			err,
		)
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		return nil, "", errors.Join(
			errors.New("unable to get X509SVID"),
			err,
		)
	}

	return source, svid.ID.String(), nil
}
