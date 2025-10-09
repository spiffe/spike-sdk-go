//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package net provides network utilities for secure mTLS communication using
// SPIFFE X.509 certificates.
//
// The package includes functionality for:
//   - Creating mTLS HTTP clients and servers with SPIFFE authentication
//   - Predicate-based authorization for fine-grained access control
//   - HTTP POST operations with both JSON and streaming support
//   - Request and response body handling with proper resource cleanup
//   - Common HTTP error handling and status code mapping
//
// All network operations use mutual TLS authentication, where both client and
// server verify each other's SPIFFE identities. Predicates can be used to
// restrict which peer SPIFFE IDs are allowed to connect.
//
// Example usage:
//
//	// Create an mTLS client that only connects to SPIKE Nexus servers
//	source, _ := workloadapi.NewX509Source(ctx)
//	client, err := net.CreateMTLSClientWithPredicate(source, predicate.AllowNexus)
//
//	// Make a secure POST request
//	payload := []byte(`{"key": "value"}`)
//	response, err := net.Post(client, "https://nexus:8443/api/v1/secrets", payload)
//
//	// Create an mTLS server with custom predicate
//	server, err := net.CreateMTLSServerWithPredicate(source, ":8443",
//	    func(id string) bool {
//	        return strings.HasPrefix(id, "spiffe://example.org/")
//	    })
package net
