//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
)

// SecretMetadataRequest for get secrets metadata
type SecretMetadataRequest struct {
	Path    string `json:"path"`
	Version int    `json:"version,omitempty"` // Optional specific version
}

// SecretMetadataResponse for secrets versions and metadata
type SecretMetadataResponse struct {
	data.SecretMetadata
	Err data.ErrorCode `json:"err,omitempty"`
}

// SecretPutRequest for creating/updating secrets
type SecretPutRequest struct {
	Path   string            `json:"path"`
	Values map[string]string `json:"values"`
	Err    data.ErrorCode    `json:"err,omitempty"`
}

// SecretPutResponse is after a successful secret write operation.
type SecretPutResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}

// SecretReadRequest is for getting secrets
type SecretReadRequest struct {
	Path    string `json:"path"`
	Version int    `json:"version,omitempty"` // Optional specific version
}

// SecretReadResponse is for getting secrets
type SecretReadResponse struct {
	data.Secret
	Err data.ErrorCode `json:"err,omitempty"`
}

// SecretDeleteRequest for soft-deleting secret versions
type SecretDeleteRequest struct {
	Path     string `json:"path"`
	Versions []int  `json:"versions"` // Empty means the latest version
}

// SecretDeleteResponse after soft-delete
type SecretDeleteResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}

// SecretUndeleteRequest for recovering soft-deleted versions
type SecretUndeleteRequest struct {
	Path     string `json:"path"`
	Versions []int  `json:"versions"`
}

// SecretUndeleteResponse after recovery
type SecretUndeleteResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}

// SecretListRequest for listing secrets
type SecretListRequest struct {
}

// SecretListResponse for listing secrets
type SecretListResponse struct {
	Keys []string       `json:"keys"`
	Err  data.ErrorCode `json:"err,omitempty"`
}

// SecretEncryptRequest for encrypting data
type SecretEncryptRequest struct {
	// Plaintext data to encrypt
	Plaintext []byte `json:"plaintext"`
	// Optional: specify encryption algorithm/version
	Algorithm string `json:"algorithm,omitempty"`
}

// SecretEncryptResponse contains encrypted data
type SecretEncryptResponse struct {
	// Version byte for future compatibility
	Version byte `json:"version"`
	// Nonce used for encryption
	Nonce []byte `json:"nonce"`
	// Encrypted ciphertext
	Ciphertext []byte `json:"ciphertext"`
	// Error code if operation failed
	Err data.ErrorCode `json:"err,omitempty"`
}

// SecretDecryptRequest for decrypting data
type SecretDecryptRequest struct {
	// Version byte to determine decryption method
	Version byte `json:"version"`
	// Nonce used during encryption
	Nonce []byte `json:"nonce"`
	// Encrypted ciphertext to decrypt
	Ciphertext []byte `json:"ciphertext"`
	// Optional: specify decryption algorithm/version
	Algorithm string `json:"algorithm,omitempty"`
}

// SecretDecryptResponse contains decrypted data
type SecretDecryptResponse struct {
	// Decrypted plaintext data
	Plaintext []byte `json:"plaintext"`
	// Error code if operation failed
	Err data.ErrorCode `json:"err,omitempty"`
}
