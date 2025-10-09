//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package api provides the primary client interface for interacting with
// SPIKE services.
//
// The API type serves as the main entry point for all SPIKE operations,
// supporting:
//   - Secret management (create, read, update, delete, list, and version
//     control)
//   - Policy management (create, read, update, delete, and list access
//     control policies)
//   - Cryptographic operations (encrypt and decrypt via streaming or
//     JSON modes)
//   - Operator functions (recover and restore using Shamir secret sharing)
//
// All operations use mutual TLS authentication with SPIFFE X.509 certificates
// and communicate exclusively with SPIKE Nexus servers by default.
//
// Example usage:
//
//	// Create a new API client
//	api := api.New()
//	if api == nil {
//	    log.Fatal("Failed to initialize SPIKE API")
//	}
//	defer api.Close()
//
//	// Store a secret
//	err := api.PutSecret("app/db/password", map[string]string{
//	    "username": "admin",
//	    "password": "secret123",
//	})
//
//	// Retrieve a secret
//	secret, err := api.GetSecret("app/db/password")
//
//	// Create an access policy
//	err = api.CreatePolicy(
//	    "db-access",
//	    "spiffe://example.org/app/*",
//	    "app/db/*",
//	    []data.PolicyPermission{data.PermissionRead},
//	)
package api
