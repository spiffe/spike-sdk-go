//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package mem provides utilities for secure mem operations.
package mem

import (
	"crypto/rand"
	"runtime"
	"unsafe"
)

// ClearRawBytes securely erases all bytes in the provided value by overwriting
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
func ClearRawBytes[T any](s *T) {
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

// ClearRawBytesParanoid provides a more thorough memory wiping method for
// highly-sensitive data.
//
// It performs multiple passes using different patterns (zeros, ones,
// random data, and alternating bits) to minimize potential data remanence
// concerns from sophisticated physical memory attacks.
//
// This method is designed for extremely security-sensitive applications where:
//  1. An attacker might have physical access to RAM
//  2. Cold boot attacks or specialized memory forensics equipment might be
//     employed
//  3. The data being protected is critically sensitive (e.g., high-value
//     encryption keys)
//
// For most applications, the standard Clear() method is sufficient as:
//   - Modern RAM technologies (DDR4/DDR5) make data remanence attacks
//     increasingly difficult
//   - Successful attacks typically require specialized equipment and immediate
//     (sub-second) physical access.
//   - The time window for such attacks is extremely short after power loss
//   - The detectable signal from previous memory states diminishes rapidly with
//     a single overwrite
//
// This method is provided for users with extreme security requirements or in
// regulated environments where multiple-pass overwrite policies are mandated.
func ClearRawBytesParanoid[T any](s *T) {
	if s == nil {
		return
	}

	p := unsafe.Pointer(s)
	size := unsafe.Sizeof(*s)
	b := (*[1 << 30]byte)(p)[:size:size]

	// Pattern overwrite cycles:
	// 1. All zeros
	// 2. All ones (0xFF)
	// 3. Random data
	// 4. Alternating 0x55/0xAA (01010101/10101010)
	// 5. Final zero out

	// Zero out all bytes (first pass)
	for i := range b {
		b[i] = 0
	}
	runtime.KeepAlive(s)

	// Fill with ones (second pass)
	for i := range b {
		b[i] = 0xFF
	}
	runtime.KeepAlive(s)

	// Fill with random data (third pass)
	_, err := rand.Read(b)
	if err != nil {
		panic("")
	}
	runtime.KeepAlive(s)

	// Alternating bit pattern (fourth pass)
	for i := range b {
		if i%2 == 0 {
			b[i] = 0x55 // 01010101
		} else {
			b[i] = 0xAA // 10101010
		}
	}
	runtime.KeepAlive(s)

	// Final zero out (fifth pass)
	for i := range b {
		b[i] = 0
	}
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
// This is especially important for slices because executing `mem.Clear` on
// a slice it will only zero out the slice header structure itself, NOT the
// underlying array data that the slice points to.
//
// When we pass a byte slice s to the function Clear[T any](s *T),
// we are passing a pointer to the slice header, not a pointer to the
// underlying array. The slice header contains three fields:
//   - A pointer to the underlying array
//   - The length of the slice
//   - The capacity of the slice
//
// mem.Clear(s) will zero out this slice header structure, but not the
// actual array data the slice points to
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
