//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

// EncryptedData is the result of encrypting a payload: the ciphertext
// together with the version and nonce that SPIKE Nexus used to produce it.
//
// Decryption needs all three, and SPIKE Nexus validates each of them, so a
// caller that intends to decrypt the result later needs this rather than the
// ciphertext alone. The algorithm is not part of the result: the caller chooses
// it and passes it to both operations.
type EncryptedData struct {
	// Version determines the decryption method, and exists for future
	// compatibility.
	Version byte `json:"version"`
	// Nonce is the per-operation nonce generated during encryption.
	Nonce []byte `json:"nonce"`
	// Ciphertext is the encrypted payload.
	Ciphertext []byte `json:"ciphertext"`
}
