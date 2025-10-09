//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package spiffe provides utilities for working with SPIFFE (Secure Production
// Identity Framework for Everyone) and the SPIFFE Workload API.
//
// The package includes functionality for:
//   - Discovering and connecting to SPIFFE Workload API endpoints
//   - Creating and managing X.509 SVID sources for workload identity
//   - Extracting SPIFFE IDs from HTTP requests over mTLS connections
//   - Proper resource cleanup and connection management
//
// SPIFFE provides a standard way to identify and authenticate workloads in
// distributed systems. This package simplifies the integration with SPIFFE
// infrastructure by handling the connection setup, certificate management,
// and identity extraction.
//
// Example usage:
//
//	// Get the SPIFFE Workload API socket
//	socket := spiffe.EndpointSocket()
//
//	// Create an X.509 source
//	source, svidID, err := spiffe.Source(context.Background(), socket)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer spiffe.CloseSource(source)
//
//	// Extract SPIFFE ID from an mTLS HTTP request
//	spiffeID, err := spiffe.IDFromRequest(req)
//	if err != nil {
//	    log.Printf("Failed to extract SPIFFE ID: %v", err)
//	}
package spiffe
