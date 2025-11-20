//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"io"

	"github.com/spiffe/spike-sdk-go/api/internal/impl/cipher"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// CipherEncryptStream encrypts data from a reader using streaming mode.
//
// It sends the reader content as the request body with the specified content
// type to SPIKE Nexus for encryption.
//
// Parameters:
//   - reader: The data source to encrypt
//   - contentType: The MIME type of the data (e.g., "application/json")
//
// Returns:
//   - []byte: The encrypted ciphertext if successful, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - Errors from streamPost(): if the streaming request fails
//   - ErrNetReadingResponseBody: if reading the response fails
//
// Example:
//
//	reader := strings.NewReader("sensitive data")
//	encrypted, err := api.CipherEncryptStream(reader, "text/plain")
func (a *API) CipherEncryptStream(
	reader io.Reader, contentType string,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.EncryptStream(a.source, reader, contentType)
}

// CipherEncryptJSON encrypts data using JSON mode with structured parameters.
//
// It sends plaintext and algorithm as JSON to SPIKE Nexus and returns the
// encrypted ciphertext bytes.
//
// Parameters:
//   - plaintext: The data to encrypt
//   - algorithm: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - []byte: The encrypted ciphertext if successful, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from httpPost(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	data := []byte("secret message")
//	encrypted, err := api.CipherEncryptJSON(data, "AES-GCM")
func (a *API) CipherEncryptJSON(
	plaintext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.EncryptJSON(a.source, plaintext, algorithm)
}

// CipherDecryptStream decrypts data from a reader using streaming mode.
//
// It sends the reader content as the request body with the specified content
// type to SPIKE Nexus for decryption.
//
// Parameters:
//   - reader: The encrypted data source to decrypt
//   - contentType: The MIME type of the data (e.g., "application/octet-stream")
//
// Returns:
//   - []byte: The decrypted plaintext if successful, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - Errors from streamPost(): if the streaming request fails
//   - ErrNetReadingResponseBody: if reading the response fails
//
// Example:
//
//	reader := bytes.NewReader(encryptedData)
//	plaintext, err := api.CipherDecryptStream(reader, "application/octet-stream")
func (a *API) CipherDecryptStream(
	reader io.Reader, contentType string,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.DecryptStream(a.source, reader, contentType)
}

// CipherDecryptJSON decrypts data using JSON mode with structured parameters.
//
// It sends version, nonce, ciphertext, and algorithm as JSON to SPIKE Nexus
// and returns the decrypted plaintext.
//
// Parameters:
//   - version: The cipher version used during encryption
//   - nonce: The nonce bytes used during encryption
//   - ciphertext: The encrypted data to decrypt
//   - algorithm: The encryption algorithm used (e.g., "AES-GCM")
//
// Returns:
//   - []byte: The decrypted plaintext if successful, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from httpPost(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	plaintext, err := api.CipherDecryptJSON(1, nonce, ciphertext, "AES-GCM")
func (a *API) CipherDecryptJSON(
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.DecryptJSON(a.source, version, nonce, ciphertext, algorithm)
}
