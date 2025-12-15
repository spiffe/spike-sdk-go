//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

//go:build !windows

package mem

import (
	"syscall"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Lock attempts to lock all current and future process memory to prevent
// swapping to disk. This is a security measure to protect sensitive data
// (such as encryption keys, secrets, and credentials) from being written
// to swap space where it could potentially be recovered.
//
// The function uses syscall.Mlockall with MCL_CURRENT | MCL_FUTURE flags
// to lock both existing memory pages and any pages allocated in the future.
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or ErrSystemMemLockFailed if the
//     memory lock operation fails (e.g., insufficient privileges, resource
//     limits exceeded)
//
// Note: This function is only available on non-Windows platforms. On Linux,
// the process typically needs CAP_IPC_LOCK capability or sufficient RLIMIT_MEMLOCK.
//
// Example:
//
//	if err := mem.Lock(); err != nil {
//	    log.Printf("Warning: could not lock memory: %v", err)
//	    // Decide whether to continue without memory locking
//	}
func Lock() *sdkErrors.SDKError {
	// Attempt to lock all current and future memory
	if err := syscall.Mlockall(
		syscall.MCL_CURRENT | syscall.MCL_FUTURE); err != nil {
		return sdkErrors.ErrSystemMemLockFailed.Clone()
	}
	return nil
}
