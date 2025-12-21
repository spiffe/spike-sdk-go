//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"io"

	"github.com/spiffe/spike-sdk-go/api/internal/impl/cipher"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// CipherEncryptStream encrypts data from a reader using streaming mode.
//
// It sends the reader content as the request body to SPIKE Nexus for encryption.
// The data is treated as binary (application/octet-stream) regardless of its
// original format, as encryption operates on raw bytes.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - reader: The data source to encrypt
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	reader := strings.NewReader("sensitive data")
//	encrypted, err := api.CipherEncryptStream(ctx, reader)
func (a *API) CipherEncryptStream(
	ctx context.Context, reader io.Reader,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.EncryptStream(ctx, a.source, reader)
}

// CipherEncrypt encrypts data with structured parameters.
//
// It sends plaintext and algorithm to SPIKE Nexus and returns the
// encrypted ciphertext bytes.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	data := []byte("secret message")
//	encrypted, err := api.CipherEncrypt(ctx, data, "AES-GCM")
func (a *API) CipherEncrypt(
	ctx context.Context, plaintext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.Encrypt(ctx, a.source, plaintext, algorithm)
}

// CipherDecryptStream decrypts data from a reader using streaming mode.
//
// It sends the reader content as the request body to SPIKE Nexus for decryption.
// The data is treated as binary (application/octet-stream) as decryption
// operates on raw encrypted bytes.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - reader: The encrypted data source to decrypt
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	reader := bytes.NewReader(encryptedData)
//	plaintext, err := api.CipherDecryptStream(ctx, reader)
func (a *API) CipherDecryptStream(
	ctx context.Context, reader io.Reader,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.DecryptStream(ctx, a.source, reader)
}

// CipherDecrypt decrypts data with structured parameters.
//
// It sends version, nonce, ciphertext, and algorithm to SPIKE Nexus
// and returns the decrypted plaintext.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	plaintext, err := api.CipherDecrypt(ctx, 1, nonce, ciphertext, "AES-GCM")
func (a *API) CipherDecrypt(
	ctx context.Context,
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	return cipher.Decrypt(ctx, a.source, version, nonce, ciphertext, algorithm)
}
