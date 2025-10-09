//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package acl provides internal implementation for access control list (ACL)
// policy management. It includes functions for creating, listing, retrieving,
// and deleting policies that control access to SPIKE resources. All operations
// require mutual TLS authentication using SPIFFE X.509 certificates and support
// predicate-based trust validation for server connections.
package acl
