//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package crypto provides cryptographic utilities for SPIKE.
//
// It includes functionality for:
//   - Generating cryptographically secure random strings and identifiers
//   - Creating AES-256 encryption keys
//   - Character class-based random string generation with support for
//     predefined classes (\w, \d, \x) and custom ranges (e.g., A-Za-z0-9)
//   - Deterministic readers for testing and reproducible random data
//   - Template-based string generation
//
// All random generation uses cryptographically secure random number generators
// to ensure suitability for security-sensitive operations.
package crypto
