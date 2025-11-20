//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"

	"github.com/spiffe/spike-sdk-go/config/env"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/predicate"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

// RequestBody reads and returns the entire request body as a byte slice.
// It reads all data from r.Body and ensures the body is properly closed
// after reading, even if an error occurs during the read operation.
//
// The function uses errors.Join to combine any read error with potential
// close errors, ensuring that close failures are not silently ignored.
//
// Parameters:
//   - r: HTTP request containing the body to read
//
// Returns:
//   - bod: byte slice containing the full request body data
//   - err: *sdkErrors.SDKError with ErrNetReadingRequestBody or
//     ErrFSStreamCloseFailed if an error occurred during reading or
//     closing the body
//
// Example:
//
//	body, err := RequestBody(req)
//	if err != nil {
//	    log.Printf("Failed to read request body: %v", err)
//	    return
//	}
//	// Process body data...
func RequestBody(r *http.Request) (bod []byte, err *sdkErrors.SDKError) {
	const fName = "RequestBody"

	body, e := io.ReadAll(r.Body)
	if e != nil {
		failErr := sdkErrors.ErrNetReadingRequestBody.Wrap(e)
		return nil, failErr
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		failErr := sdkErrors.ErrFSStreamCloseFailed
		log.WarnErr(fName, *failErr)
	}(r.Body)

	return body, err
}

// AuthorizerWithPredicate creates a TLS authorizer that validates SPIFFE IDs
// using the provided predicate function.
//
// The authorizer checks each connecting peer's SPIFFE ID against the predicate.
// If the predicate returns true, the connection is authorized. If false, the
// connection is rejected with ErrAccessUnauthorized.
//
// Parameters:
//   - predicate: Function that takes a SPIFFE ID string and returns true to
//     allow the connection, false to reject it
//
// Returns:
//   - tlsconfig.Authorizer: A TLS authorizer that can be used with mTLS configs
//
// Example:
//
//	// Allow only production namespace
//	authorizer := AuthorizerWithPredicate(func(id string) bool {
//	    return strings.Contains(id, "/ns/production/")
//	})
func AuthorizerWithPredicate(predicate func(string) bool) tlsconfig.Authorizer {
	return tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if predicate(id.String()) {
			return nil
		}

		failErr := sdkErrors.ErrAccessUnauthorized
		failErr.Msg = fmt.Sprintf("unauthorized spiffe id: '%s'", id.String())

		return failErr
	})
}

// CreateMTLSServerWithPredicate creates an HTTP server configured for mutual
// TLS (mTLS) authentication using SPIFFE X.509 certificates. It sets up the
// server with a custom authorizer that validates client SPIFFE IDs against a
// provided predicate function.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. It must be initialized and valid.
//   - tlsPort: The network address and port for the server to listen on
//     (e.g., ":8443").
//   - predicate: A function that takes a client SPIFFE ID string and returns
//     true if the client should be allowed access, false otherwise.
//
// Returns:
//   - *http.Server: A configured HTTP server ready to be started with TLS
//     enabled.
//
// The server uses the provided X509Source for both its own identity and for
// validating client certificates. Client connections are only accepted if their
// SPIFFE ID passes the provided predicate function.
func CreateMTLSServerWithPredicate(source *workloadapi.X509Source,
	tlsPort string,
	predicate func(string) bool) *http.Server {
	authorizer := AuthorizerWithPredicate(predicate)
	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:              tlsPort,
		TLSConfig:         tlsConfig,
		ReadHeaderTimeout: env.HTTPServerReadHeaderTimeoutVal(),
		// ^ Timeout for reading request headers,
		// it helps prevent slowloris attacks
	}
	return server
}

// CreateMTLSServer creates an HTTP server configured for mutual TLS (mTLS)
// authentication using SPIFFE X.509 certificates.
//
// WARNING: This function accepts ALL client SPIFFE IDs without validation.
// For production use, consider using CreateMTLSServerWithPredicate to restrict
// which clients can connect to this server for better security.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. It must be initialized and valid.
//   - tlsPort: The network address and port for the server to listen on
//     (e.g., ":8443").
//
// Returns:
//   - *http.Server: A configured HTTP server ready to be started with TLS
//     enabled.
//
// The server uses the provided X509Source for both its own identity and for
// validating client certificates. Client connections are accepted from ANY
// client with a valid SPIFFE certificate.
func CreateMTLSServer(source *workloadapi.X509Source,
	tlsPort string) *http.Server {
	return CreateMTLSServerWithPredicate(source, tlsPort, predicate.AllowAll)
}

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

// Source creates and returns a new SPIFFE X509Source for workload API
// communication. It establishes a connection to the SPIFFE workload API using
// the default endpoint socket with a configurable timeout to prevent indefinite
// blocking on socket issues.
//
// The timeout can be configured using the SPIKE_SPIFFE_SOURCE_TIMEOUT
// environment variable (default: 30s).
//
// The function will terminate the program with exit code 1 if the source
// creation fails or times out.
//
// Returns:
//   - *workloadapi.X509Source: A new X509Source for SPIFFE workload API
//     communication
func Source() *workloadapi.X509Source {
	const fName = "Source"

	ctx, cancel := context.WithTimeout(
		context.Background(),
		env.SPIFFESourceTimeoutVal(),
	)
	defer cancel()

	source, _, err := spiffe.Source(ctx, spiffe.EndpointSocket())
	if err != nil {
		failErr := sdkErrors.ErrObjectCreationFailed.Wrap(err)
		log.FatalErr(fName, *failErr)
	}
	return source
}

// ServeWithPredicate initializes and starts an HTTPS server using mTLS
// authentication with SPIFFE X.509 certificates. It sets up the server routes
// using the provided initialization function and listens for incoming
// connections on the specified port.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. Must not be nil.
//   - initializeRoutes: A function that sets up the HTTP route handlers for the
//     server. This function is called before the server starts.
//   - predicate: a predicate function to pass to CreateMTLSServer.
//   - tlsPort: The network address and port for the server to listen to on
//     (e.g., ":8443").
//
// Returns:
//   - *sdkErrors.SDKError: Returns nil if the server starts successfully,
//     otherwise returns one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrFSStreamOpenFailed: if the server fails to start or encounters an error
//     while running
//
// The function uses empty strings for the certificate and key file parameters
// in ListenAndServeTLS as the certificates are provided by the X509Source. The
// server's mTLS configuration is determined by the CreateMTLSServer function.
func ServeWithPredicate(source *workloadapi.X509Source,
	initializeRoutes func(),
	predicate func(string) bool,
	tlsPort string) *sdkErrors.SDKError {
	if source == nil {
		failErr := sdkErrors.ErrSPIFFENilX509Source
		failErr.Msg = "got nil source while trying to serve"
		return failErr
	}

	initializeRoutes()

	server := CreateMTLSServerWithPredicate(source, tlsPort, predicate)

	if err := server.ListenAndServeTLS("", ""); err != nil {
		failErr := sdkErrors.ErrFSStreamOpenFailed.Wrap(err)
		failErr.Msg = "failed to listen and serve"
		return failErr
	}

	return nil
}

// Serve initializes and starts an HTTPS server using mTLS
// authentication with SPIFFE X.509 certificates. It sets up the server routes
// using the provided initialization function and listens for incoming
// connections on the specified port.
//
// WARNING: This function accepts ALL client SPIFFE IDs without validation.
// For production use, consider using ServeWithPredicate to restrict
// which clients can connect to this server for better security.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. Must not be nil.
//   - initializeRoutes: A function that sets up the HTTP route handlers for the
//     server. This function is called before the server starts.
//   - tlsPort: The network address and port for the server to listen on
//     (e.g., ":8443").
//
// Returns:
//   - *sdkErrors.SDKError: Returns nil if the server starts successfully,
//     otherwise returns one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrFSStreamOpenFailed: if the server fails to start or encounters an error
//     while running
//
// The function uses empty strings for the certificate and key file parameters
// in ListenAndServeTLS as the certificates are provided by the X509Source. The
// server's mTLS configuration is determined by the CreateMTLSServer function.
func Serve(
	source *workloadapi.X509Source,
	initializeRoutes func(),
	tlsPort string) *sdkErrors.SDKError {
	return ServeWithPredicate(
		source, initializeRoutes,
		predicate.AllowAll, tlsPort,
	)
}
