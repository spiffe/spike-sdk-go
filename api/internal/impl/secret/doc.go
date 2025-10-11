//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package secret provides internal implementation for secret management
// operations. It includes functions for creating, retrieving, updating,
// deleting, undeleting, and listing secrets, as well as accessing secret
// metadata and version information. All operations require mutual TLS
// authentication using SPIFFE X.509 certificates.
package secret
