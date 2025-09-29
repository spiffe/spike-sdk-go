//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/predicate"
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
//   - err: any error that occurred during reading or closing the body
//
// Example:
//
//	body, err := RequestBody(req)
//	if err != nil {
//	    log.Printf("Failed to read request body: %v", err)
//	    return
//	}
//	// Process body data...
func RequestBody(r *http.Request) (bod []byte, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err = errors.Join(err, b.Close())
	}(r.Body)

	return body, err
}

// CreateMTLSServer creates an HTTP server configured for mutual TLS (mTLS)
// authentication using SPIFFE X.509 certificates. It sets up the server with a
// custom authorizer that validates client SPIFFE IDs against a provided
// predicate function.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. It must be initialized and valid.
//   - tlsPort: The network address and port for the server to listen on
//     (e.g., ":8443").
//   - predicate: A function that takes a SPIFFE ID string and returns true if
//     the client should be allowed access, false otherwise.
//
// Returns:
//   - *http.Server: A configured HTTP server ready to be started with TLS
//     enabled.
//   - error: An error if the server configuration fails.
//
// The server uses the provided X509Source for both its own identity and for
// validating client certificates. Client connections are only accepted if their
// SPIFFE ID passes the provided predicate function.
func CreateMTLSServer(source *workloadapi.X509Source,
	tlsPort string,
	predicate func(string) bool) (*http.Server, error) {
	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if predicate(id.String()) {
			return nil
		}

		return fmt.Errorf(
			"authorizer: TLS Config: untrusted spiffe id: '%s'", id.String(),
		)
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:              tlsPort,
		TLSConfig:         tlsConfig,
		ReadHeaderTimeout: 10 * time.Second,
		// ^ Timeout for reading request headers,
		// it helps prevent slowloris attacks
	}
	return server, nil
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
//   - error: An error if the client creation fails
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
//	client, err := CreateMTLSClientWithPredicate(source,
//	 func(serverID string) bool {
//	    return strings.Contains(serverID, "/ns/backend/")
//	})
func CreateMTLSClientWithPredicate(
	source *workloadapi.X509Source,
	predicate predicate.Predicate,
) (*http.Client, error) {
	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if predicate(id.String()) {
			return nil
		}

		return fmt.Errorf(
			"TLS Config: untrusted spiffe id: '%s'", id.String(),
		)
	})

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     tlsConfig,
			IdleConnTimeout:     30 * time.Second,
			MaxIdleConns:        100,
			MaxConnsPerHost:     10,
			MaxIdleConnsPerHost: 10,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		},
		Timeout: 60 * time.Second,
	}

	return client, nil
}

// CreateMTLSClient creates an HTTP client configured for mutual TLS
// authentication using SPIFFE workload identities.
// It uses the provided X.509 source for client certificates and validates peer
// certificates against a predicate function.
//
// Parameters:
//   - source: An X509Source that provides the client's identity certificates
//     and trusted roots
//
// Returns:
//   - *http.Client: A configured HTTP client that will use mTLS for all
//     connections
//   - error: An error if the client creation fails
//
// The returned client will:
//   - Present client certificates from the provided X509Source
//   - Validate peer certificates using the same X509Source
//   - Only accept peer certificates with SPIFFE IDs that pass the predicate
//     function
func CreateMTLSClient(source *workloadapi.X509Source) (*http.Client, error) {
	return CreateMTLSClientWithPredicate(source, predicate.AllowAll)
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
//   - error: Returns nil if the server starts successfully, otherwise returns
//     an error explaining the failure. Specific error cases include:
//   - If the source is nil
//   - If server creation fails
//   - If the server fails to start or encounters an error while running
//
// The function uses empty strings for the certificate and key file parameters
// in ListenAndServeTLS as the certificates are provided by the X509Source. The
// server's mTLS configuration is determined by the CreateMTLSServer function.
func ServeWithPredicate(source *workloadapi.X509Source,
	initializeRoutes func(),
	predicate func(string) bool,
	tlsPort string) error {
	if source == nil {
		return errors.New("serve: got nil source while trying to serve")
	}

	initializeRoutes()

	server, err := CreateMTLSServer(source, tlsPort, predicate)
	if err != nil {
		return err
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		return errors.Join(
			err,
			errors.New("serve: failed to listen and serve"),
		)
	}

	return nil
}

// Serve initializes and starts an HTTPS server using mTLS
// authentication with SPIFFE X.509 certificates. It sets up the server routes
// using the provided initialization function and listens for incoming
// connections on the specified port.
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
//   - error: Returns nil if the server starts successfully, otherwise returns
//     an error explaining the failure. Specific error cases include:
//   - If `source` is nil
//   - If server creation fails
//   - If the server fails to start or encounters an error while running
//
// The function uses empty strings for the certificate and key file parameters
// in ListenAndServeTLS as the certificates are provided by the X509Source. The
// server's mTLS configuration is determined by the CreateMTLSServer function.
func Serve(
	source *workloadapi.X509Source,
	initializeRoutes func(),
	tlsPort string) error {
	return ServeWithPredicate(
		source, initializeRoutes,
		func(string) bool { return true }, tlsPort)
}
