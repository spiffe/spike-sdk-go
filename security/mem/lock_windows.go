//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

//go:build windows

package mem

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Lock attempts to lock all current and future process memory to prevent
// swapping to disk. This is a security measure to protect sensitive data
// (such as encryption keys, secrets, and credentials) from being written
// to swap space where it could potentially be recovered.
//
// On Windows, this operation is not supported and always returns an error.
// Memory locking via mlock/mlockall is only available on Unix-like systems.
//
// Returns:
//   - *sdkErrors.SDKError: Always returns ErrSystemMemLockFailed on Windows
//     since memory locking is not supported
//
// Example:
//
//	if err := mem.Lock(); err != nil {
//	    log.Printf("Warning: could not lock memory: %v", err)
//	    // Decide whether to continue without memory locking
//	}
func Lock() *sdkErrors.SDKError {
	// mlock/mlockall is only available on Unix-like systems
	return sdkErrors.ErrSystemMemLockFailed.Clone()
}
