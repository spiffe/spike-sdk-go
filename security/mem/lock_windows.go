//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

//go:build windows

// Package mem provides utilities for secure mem operations.
package mem

// Lock attempts to lock the process memory to prevent swapping.
// Returns true if successful, false if not supported or failed.
func Lock() bool {
	// `mlock` is only available on Unix-like systems
	return false
}
