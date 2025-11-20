//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

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
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (b BootstrapVerifyResponse) Success() BootstrapVerifyResponse {
	b.Err = sdkErrors.ErrSuccess.Code
	return b
}
func (s BootstrapVerifyResponse) NotFound() BootstrapVerifyResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s BootstrapVerifyResponse) BadRequest() BootstrapVerifyResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s BootstrapVerifyResponse) Unauthorized() BootstrapVerifyResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s BootstrapVerifyResponse) Internal() BootstrapVerifyResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}
