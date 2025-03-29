//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package mem provides utilities for secure mem operations.
package mem

import (
	"runtime"
	"unsafe"
)

// Clear securely erases all bytes in the provided value by overwriting
// its mem with zeros. This ensures sensitive data like root keys and
// Shamir shards are properly cleaned from mem before garbage collection.
//
// According to NIST SP 800-88 Rev. 1 (Guidelines for Media Sanitization),
// a single overwrite pass with zeros is sufficient for modern storage
// devices, including RAM.
//
// Parameters:
//   - s: A pointer to any type of data that should be securely erased
//
// Usage:
//
//	type SensitiveData struct {
//	    Key    [32]byte
//	    Token  string
//	}
//
//	data := &SensitiveData{...}
//	defer mem.Clear(data)
//	// Use data...
func Clear[T any](s *T) {
	if s == nil {
		return
	}

	p := unsafe.Pointer(s)
	size := unsafe.Sizeof(*s)
	b := (*[1 << 30]byte)(p)[:size:size]

	// Zero out all bytes in mem
	for i := range b {
		b[i] = 0
	}

	// Make sure the data is actually wiped before gc has time to interfere
	runtime.KeepAlive(s)
}

// Zeroed32 checks if a 32-byte array contains only zero values.
// Returns true if all bytes are zero, false otherwise.
func Zeroed32(ar *[32]byte) bool {
	for _, v := range ar {
		if v != 0 {
			return false
		}
	}
	return true
}

// ClearBytes securely erases a byte slice by overwriting all bytes with zeros.
// This is a convenience wrapper around Clear for byte slices.
//
// Parameters:
//   - b: A byte slice that should be securely erased
//
// Usage:
//
//	key := []byte{...} // Sensitive cryptographic key
//	defer mem.ClearBytes(key)
//	// Use key...
func ClearBytes(b []byte) {
	if len(b) == 0 {
		return
	}

	for i := range b {
		b[i] = 0
	}

	// Make sure the data is actually wiped before gc has time to interfere
	runtime.KeepAlive(b)
}
