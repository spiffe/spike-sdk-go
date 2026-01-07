//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strconv"
)

// CryptoMaxCiphertextSizeVal returns the maximum allowed ciphertext size in
// bytes.
// It reads the value from the SPIKE_NEXUS_CRYPTO_MAX_CIPHERTEXT_SIZE
// environment variable. If the variable is not set or contains an invalid
// value, it defaults to 65,536 bytes (64 KB).
func CryptoMaxCiphertextSizeVal() int {
	p := os.Getenv(NexusCryptoMaxCiphertextSize)
	if p != "" {
		mv, err := strconv.Atoi(p)
		if err == nil && mv > 0 {
			return mv
		}
	}
	return 65536
}

// CryptoMaxPlaintextSizeVal returns the maximum allowed plaintext size in
// bytes.
// It is calculated as CryptoMaxCiphertextSizeVal minus 16 bytes, which accounts
// for the authentication tag overhead used in authenticated encryption schemes
// such as AES-GCM.
func CryptoMaxPlaintextSizeVal() int {
	return CryptoMaxCiphertextSizeVal() - 16
}
