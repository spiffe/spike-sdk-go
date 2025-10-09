//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package spiffeid provides utilities for constructing and validating
// SPIFFE IDs for SPIKE system components.
//
// The package includes functionality for:
//   - Constructing standardized SPIFFE IDs for SPIKE components (Nexus, Keeper,
//     Pilot, Bootstrap, LiteWorkload)
//   - Validating SPIFFE IDs against expected patterns for each component type
//   - Supporting multiple trust domains via environment configuration
//   - Handling both exact matches and extended IDs with instance metadata
//   - Peer authorization for inter-component communication
//
// SPIFFE ID Construction:
//
// Each SPIKE component has a standardized SPIFFE ID pattern. Functions are
// provided to construct these IDs:
//
//	nexusID := spiffeid.Nexus("example.org")
//	// Returns: "spiffe://example.org/spike/nexus"
//
//	pilotID := spiffeid.Pilot("example.org")
//	// Returns: "spiffe://example.org/spike/pilot/role/superuser"
//
// Identity Validation:
//
// Validation functions check if a given SPIFFE ID matches the expected pattern
// for a component type, supporting both exact and extended matches:
//
//	if spiffeid.IsNexus("spiffe://example.org/spike/nexus") {
//	    // Handle Nexus-specific logic
//	}
//
//	// Extended IDs with metadata are also recognized
//	if spiffeid.IsNexus("spiffe://example.org/spike/nexus/instance-0") {
//	    // Also recognized as Nexus
//	}
//
// Supported Components:
//   - Nexus: The central secrets management service
//   - Keeper: Distributed key storage component
//   - Pilot: Administrative/superuser role
//   - Bootstrap: Initial setup component
//   - LiteWorkload: Workloads with limited encryption-only access
//
// Multi-Trust Domain Support:
//
// All validation functions support multiple trust domains configured via
// environment variables, allowing SPIKE deployments across different trust
// boundaries.
package spiffeid
