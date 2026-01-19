//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/spiffe/spike-sdk-go/log"
)

// CreateCipher generates a new AES-256-GCM cipher for authenticated encryption.
//
// This function creates a cryptographically secure random 256-bit key and
// initializes an AES block cipher in Galois/Counter Mode (GCM). GCM provides
// both confidentiality and authenticity, making it suitable for secure
// encryption operations.
//
// The function terminates the program via log.FatalLn if key generation,
// cipher creation, or GCM initialization fails. This fail-fast behavior
// ensures that cryptographic operations never proceed with invalid state.
//
// Returns:
//   - cipher.AEAD: An authenticated encryption cipher ready for use with
//     Seal and Open operations
func CreateCipher() cipher.AEAD {
	key := make([]byte, AES256KeySize) // AES-256 key
	if _, randErr := rand.Read(key); randErr != nil {
		log.FatalLn("createCipher", "message",
			"Failed to generate test key", "err", randErr)
	}

	block, cipherErr := aes.NewCipher(key)
	if cipherErr != nil {
		log.FatalLn("createCipher", "message",
			"Failed to create cipher", "err", cipherErr)
	}

	gcm, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		log.FatalLn("createCipher", "message",
			"Failed to create GCM", "err", gcmErr)
	}

	return gcm
}
