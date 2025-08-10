//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

//go:build !windows

// Package mem provides utilities for secure mem operations.
package mem

import "syscall"

// Lock attempts to lock the process memory to prevent swapping.
// Returns true if successful, false if not supported or failed.
func Lock() bool {
	// Attempt to lock all current and future memory
	if err := syscall.Mlockall(syscall.MCL_CURRENT | syscall.MCL_FUTURE); err != nil {
		return false
	}

	return true
}
