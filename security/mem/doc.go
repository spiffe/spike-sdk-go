//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package mem provides utilities for secure memory operations to protect
// sensitive data such as cryptographic keys and secrets.
//
// The package includes functionality for:
//   - Securely erasing memory by overwriting with zeros or multiple patterns
//   - Preventing memory from being swapped to disk (Unix-like systems)
//   - Checking if memory regions contain only zeros
//   - Handling both fixed-size arrays and byte slices
//
// Memory Clearing:
//
// ClearRawBytes performs a single-pass zero overwrite, which is sufficient
// for most use cases according to NIST SP 800-88 Rev. 1 guidelines:
//
//	key := &[32]byte{...}
//	defer mem.ClearRawBytes(key)
//
// ClearRawBytesParanoid performs multiple passes with different patterns
// for extreme security requirements:
//
//	sensitiveData := &[64]byte{...}
//	defer mem.ClearRawBytesParanoid(sensitiveData)
//
// For byte slices, use ClearBytes to clear the underlying array data:
//
//	token := []byte("secret-token")
//	defer mem.ClearBytes(token)
//
// Memory Locking:
//
// Lock prevents process memory from being swapped to disk, reducing the risk
// of secrets being written to persistent storage:
//
//	if mem.Lock() {
//	    // Memory is locked, proceed with sensitive operations
//	}
//
// Important Notes:
//
// ClearRawBytes only clears the direct memory of the provided value. For
// structs containing pointers, slices, or maps, you must clear the referenced
// data separately before clearing the struct itself.
package mem
