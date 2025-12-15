//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

//go:build !windows

package mem

import (
	"syscall"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Lock attempts to lock the process memory to prevent swapping.
// Returns true if successful, false if not supported or failed.
func Lock() *sdkErrors.SDKError {
	const fName = "Lock"
	// Attempt to lock all current and future memory
	if err := syscall.Mlockall(
		syscall.MCL_CURRENT | syscall.MCL_FUTURE); err != nil {
		return sdkErrors.ErrSystemMemLockFailed.Clone()
	}

	return nil
}
