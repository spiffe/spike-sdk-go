//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

//go:build !windows

package mem

import (
	"syscall"

	"github.com/spiffe/spike-sdk-go/log"
)

// Lock attempts to lock the process memory to prevent swapping.
// Returns true if successful, false if not supported or failed.
func Lock() bool {
	const fName = "Lock"
	// Attempt to lock all current and future memory
	if err := syscall.Mlockall(
		syscall.MCL_CURRENT | syscall.MCL_FUTURE); err != nil {
		log.Log().Warn(fName, "msg", "Failed to lock memory", "err", err.Error())
		return false
	}

	return true
}
