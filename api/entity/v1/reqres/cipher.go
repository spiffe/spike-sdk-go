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

func (r CipherEncryptResponse) Success() CipherEncryptResponse {
	r.Err = ""
	return r
}
func (r CipherEncryptResponse) NotFound() CipherEncryptResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrAPIResponseCodeInvalid)
	return r
}
func (r CipherEncryptResponse) BadRequest() CipherEncryptResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r CipherEncryptResponse) Unauthorized() CipherEncryptResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r CipherEncryptResponse) Internal() CipherEncryptResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r CipherEncryptResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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

func (r CipherDecryptResponse) Success() CipherDecryptResponse {
	r.Err = ""
	return r
}
func (r CipherDecryptResponse) NotFound() CipherDecryptResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrAPIResponseCodeInvalid)
	return r
}
func (r CipherDecryptResponse) BadRequest() CipherDecryptResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r CipherDecryptResponse) Unauthorized() CipherDecryptResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r CipherDecryptResponse) Internal() CipherDecryptResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r CipherDecryptResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}
