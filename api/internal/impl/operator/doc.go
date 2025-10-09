//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package operator provides internal implementation for operator-level
// functions including recovery and restoration operations using Shamir
// secret sharing.
// These operations enable disaster recovery scenarios by distributing and
// reconstructing cryptographic secrets across multiple shards. All operations
// require SPIKE Pilot authentication and use mutual TLS connections to
// SPIKE Nexus.
package operator
