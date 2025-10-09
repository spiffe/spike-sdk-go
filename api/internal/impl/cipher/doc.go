//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package cipher provides internal implementation for cryptographic operations
// including encryption and decryption. It supports two modes of operation:
// streaming mode for handling large data efficiently, and JSON mode for
// structured request/response communication. All operations require mutual TLS
// authentication using SPIFFE X.509 certificates.
package cipher
