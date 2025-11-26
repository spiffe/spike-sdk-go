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

func (r BootstrapVerifyResponse) Success() BootstrapVerifyResponse {
	r.Err = ""
	return r
}
func (r BootstrapVerifyResponse) NotFound() BootstrapVerifyResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrAPIResponseCodeInvalid)
	return r
}
func (r BootstrapVerifyResponse) BadRequest() BootstrapVerifyResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r BootstrapVerifyResponse) Unauthorized() BootstrapVerifyResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r BootstrapVerifyResponse) Internal() BootstrapVerifyResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r BootstrapVerifyResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}
