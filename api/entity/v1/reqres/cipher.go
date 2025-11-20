//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// CipherEncryptRequest for encrypting data
type CipherEncryptRequest struct {
	// Plaintext data to encrypt
	Plaintext []byte `json:"plaintext"`
	// Optional: specify encryption algorithm/version
	Algorithm string `json:"algorithm,omitempty"`
}

// CipherEncryptResponse contains encrypted data
type CipherEncryptResponse struct {
	// Version byte for future compatibility
	Version byte `json:"version"`
	// Nonce used for encryption
	Nonce []byte `json:"nonce"`
	// Encrypted ciphertext
	Ciphertext []byte `json:"ciphertext"`
	// Error code if operation failed
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (c CipherEncryptResponse) Success() CipherEncryptResponse {
	c.Err = sdkErrors.ErrSuccess.Code
	return c
}
func (s CipherEncryptResponse) NotFound() CipherEncryptResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s CipherEncryptResponse) BadRequest() CipherEncryptResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s CipherEncryptResponse) Unauthorized() CipherEncryptResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s CipherEncryptResponse) Internal() CipherEncryptResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}

// CipherDecryptRequest for decrypting data
type CipherDecryptRequest struct {
	// Version byte to determine decryption method
	Version byte `json:"version"`
	// Nonce used during encryption
	Nonce []byte `json:"nonce"`
	// Encrypted ciphertext to decrypt
	Ciphertext []byte `json:"ciphertext"`
	// Optional: specify decryption algorithm/version
	Algorithm string `json:"algorithm,omitempty"`
}

// CipherDecryptResponse contains decrypted data
type CipherDecryptResponse struct {
	// Decrypted plaintext data
	Plaintext []byte `json:"plaintext"`
	// Error code if operation failed
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (c CipherDecryptResponse) Success() CipherDecryptResponse {
	c.Err = sdkErrors.ErrSuccess.Code
	return c
}
func (s CipherDecryptResponse) NotFound() CipherDecryptResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s CipherDecryptResponse) BadRequest() CipherDecryptResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s CipherDecryptResponse) Unauthorized() CipherDecryptResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s CipherDecryptResponse) Internal() CipherDecryptResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}
