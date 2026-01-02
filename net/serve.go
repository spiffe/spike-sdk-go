//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/config/env"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/predicate"
)

// CreateMTLSServerWithPredicate creates an HTTP server configured for mutual
// TLS (mTLS) authentication using SPIFFE X.509 certificates. It sets up the
// server with a custom authorizer that validates client SPIFFE IDs against a
// provided predicate function.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. Must not be nil.
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
//
// Note: Terminates the program via log.FatalErr if `source` is nil, as this
// indicates a critical configuration error that should be caught during
// development.
func CreateMTLSServerWithPredicate(source *workloadapi.X509Source,
	tlsPort string,
	predicate func(string) bool) *http.Server {
	const fName = "CreateMTLSServerWithPredicate"

	if source == nil {
		failErr := sdkErrors.ErrSPIFFENilX509Source.Clone()
		failErr.Msg = "source cannot be nil"
		log.FatalErr(fName, *failErr)
	}

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
//     validates client certificates. Must not be nil.
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
//
// Note: Terminates the program via log.FatalErr if `source` is nil, as this
// indicates a critical configuration error that should be caught during
// development.
func CreateMTLSServer(source *workloadapi.X509Source,
	tlsPort string) *http.Server {
	return CreateMTLSServerWithPredicate(source, tlsPort, predicate.AllowAll)
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
//   - ErrFSStreamOpenFailed: if the server fails to start or encounters an
//     error while running
//
// The function uses empty strings for the certificate and key file parameters
// in ListenAndServeTLS as the certificates are provided by the X509Source. The
// server's mTLS configuration is determined by the CreateMTLSServer function.
func ServeWithPredicate(source *workloadapi.X509Source,
	initializeRoutes func(),
	predicate func(string) bool,
	tlsPort string) *sdkErrors.SDKError {
	if source == nil {
		failErr := sdkErrors.ErrSPIFFENilX509Source.Clone()
		failErr.Msg = "got nil source while trying to serve"
		return failErr
	}

	initializeRoutes()

	server := CreateMTLSServerWithPredicate(source, tlsPort, predicate)
	defer func(server *http.Server) {
		err := server.Close()
		if err != nil {
			failErr := sdkErrors.ErrFSStreamCloseFailed.Clone()
			failErr.Msg = "failed to close server"
			failErr = failErr.Wrap(err)
			log.WarnErr("ServeWithPredicate", *failErr)
		}
	}(server)

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
//   - ErrFSStreamOpenFailed: if the server fails to start or encounters an
//     error while running
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

// ServeWithRoute is a convenience wrapper around ServeWithPredicate that
// initializes and starts an HTTPS server using mTLS authentication with
// SPIFFE X.509 certificates. Unlike ServeWithPredicate, this function
// terminates the program on failure instead of returning an error.
//
// Parameters:
//   - appName: The application name used for error logging context.
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. Must not be nil.
//   - initializeRoutes: A function that sets up the HTTP route handlers for
//     the server. This function is called before the server starts.
//   - spiffeIDPredicate: A function that takes a client SPIFFE ID string and
//     returns true if the client should be allowed access, false otherwise.
//   - tlsPort: The network address and port for the server to listen on
//     (e.g., ":8443").
//
// The function terminates the program via log.FatalErr if:
//   - source is nil
//   - the server fails to start or encounters an error while running
//
// Use this function when server startup failures should be fatal. For more
// granular error handling, use ServeWithPredicate instead.
func ServeWithRoute(
	appName string,
	source *workloadapi.X509Source,
	initializeRoutes func(),
	spiffeIDPredicate func(string) bool,
	tlsPort string,
) {
	if source == nil {
		log.FatalErr(appName, *sdkErrors.ErrSPIFFENilX509Source)
	}

	if err := ServeWithPredicate(
		source,
		initializeRoutes,
		spiffeIDPredicate,
		tlsPort,
	); err != nil {
		log.FatalErr(appName, *err)
	}
}
