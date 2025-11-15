//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

// BootstrapVerifyRequest for verifying SPIKE Nexus initialization.
type BootstrapVerifyRequest struct {
	// Nonce used for encryption
	Nonce []byte `json:"nonce"`
	// Encrypted ciphertext to verify
	Ciphertext []byte `json:"ciphertext"`
}

// BootstrapVerifyResponse contains the hash of the decrypted plaintext.
type BootstrapVerifyResponse struct {
	// Hash of the decrypted plaintext
	Hash string `json:"hash"`
	// Error code if operation failed
	Err data.ErrorCode `json:"err,omitempty"`
}

func (b BootstrapVerifyResponse) Success() BootstrapVerifyResponse {
	b.Err = data.ErrSuccess
	return b
}
