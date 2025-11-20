//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"crypto/sha256"
)

// DeterministicReader implements io.Reader to generate deterministic
// pseudo-random data based on a seed. It uses SHA-256 hashing to create a
// repeatable stream of bytes.
type DeterministicReader struct {
	data []byte
	pos  int
}

// Read implements io.Reader interface. It returns deterministic data by reading
// from the internal buffer and generating new data using SHA-256 when needed.
//
// If the current position reaches the end of the data buffer, it generates
// a new block by hashing the current data. This ensures a continuous,
// deterministic stream of data.
//
// This implementation properly satisfies the io.Reader interface contract.
// The error return is always nil since deterministic hashing operations cannot
// fail, but is required for io.Reader interface compliance.
//
// Parameters:
//   - p []byte: Buffer to read data into
//
// Returns:
//   - n int: Number of bytes read
//   - err error: Always nil (deterministic reads never fail)
func (r *DeterministicReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		// Generate more deterministic data if needed
		hash := sha256.Sum256(r.data)
		r.data = hash[:]
		r.pos = 0
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

// NewDeterministicReader creates a new DeterministicReader initialized with
// the SHA-256 hash of the provided seed data.
//
// Parameters:
//   - seed []byte: Initial seed data to generate the deterministic stream
//
// Returns:
//   - *DeterministicReader: New reader instance initialized with the seed
func NewDeterministicReader(seed []byte) *DeterministicReader {
	hash := sha256.Sum256(seed)
	return &DeterministicReader{
		data: hash[:],
		pos:  0,
	}
}
