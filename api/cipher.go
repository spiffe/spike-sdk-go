//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"io"

	"github.com/spiffe/spike-sdk-go/api/internal/impl/cipher"
)

// CipherEncryptStream encrypts data from a reader using streaming mode.
// It sends the reader content as the request body with the specified content type.
//
// Parameters:
//   - reader io.Reader: The data source to encrypt
//   - contentType string: The MIME type of the data (e.g., "application/json")
//
// Returns:
//   - []byte: The encrypted data if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//	reader := strings.NewReader("sensitive data")
//	encrypted, err := api.CipherEncryptStream(reader, "text/plain")
func (a *API) CipherEncryptStream(
	reader io.Reader, contentType string,
) ([]byte, error) {
	return cipher.EncryptStream(a.source, reader, contentType)
}

// CipherEncryptJSON encrypts data using JSON mode with structured parameters.
// It sends plaintext and algorithm as JSON and returns encrypted bytes.
//
// Parameters:
//   - plaintext []byte: The data to encrypt
//   - algorithm string: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - []byte: The encrypted data if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//	data := []byte("secret message")
//	encrypted, err := api.CipherEncryptJSON(data, "AES-GCM")
func (a *API) CipherEncryptJSON(
	plaintext []byte, algorithm string,
) ([]byte, error) {
	return cipher.EncryptJSON(a.source, plaintext, algorithm)
}

// CipherDecryptStream decrypts data from a reader using streaming mode.
// It sends the reader content as the request body with the specified
// content type.
//
// Parameters:
//   - reader io.Reader: The encrypted data source to decrypt
//   - contentType string: The MIME type of the data
//     (e.g., "application/octet-stream")
//
// Returns:
//   - []byte: The decrypted plaintext if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//		reader := bytes.NewReader(encryptedData)
//		plaintext, err := api.CipherDecryptStream(
//	 	reader, "application/octet-stream")
func (a *API) CipherDecryptStream(
	reader io.Reader, contentType string,
) ([]byte, error) {
	return cipher.DecryptStream(a.source, reader, contentType)
}

// CipherDecryptJSON decrypts data using JSON mode with structured parameters.
// It sends version, nonce, ciphertext, and algorithm as JSON and returns
// plaintext.
//
// Parameters:
//   - version byte: The cipher version used during encryption
//   - nonce []byte: The nonce bytes used during encryption
//   - ciphertext []byte: The encrypted data to decrypt
//   - algorithm string: The encryption algorithm used (e.g., "AES-GCM")
//
// Returns:
//   - []byte: The decrypted plaintext if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//	plaintext, err := api.CipherDecryptJSON(1, nonce, ciphertext, "AES-GCM")
func (a *API) CipherDecryptJSON(
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, error) {
	return cipher.DecryptJSON(a.source, version, nonce, ciphertext, algorithm)
}
