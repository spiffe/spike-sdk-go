//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"context"
	"io"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// DecryptStream decrypts data from a reader using streaming mode using the
// default Cipher instance.
// It sends the reader content as the request body and returns the decrypted
// plaintext bytes. The data is treated as binary (application/octet-stream)
// as decryption operates on raw encrypted bytes.
//
// This is a convenience function that uses the default Cipher instance.
// For testing or custom configuration, create a Cipher instance directly.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - r: io.Reader containing the encrypted data
//
// Returns:
//   - ([]byte, nil) containing the decrypted plaintext if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - Errors from streamPost(): if the streaming request fails
//   - ErrNetReadingResponseBody: if reading the response fails
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	reader := bytes.NewReader(encryptedData)
//	plaintext, err := DecryptStream(ctx, source, reader)
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func DecryptStream(
	ctx context.Context, source *workloadapi.X509Source, r io.Reader,
) ([]byte, *sdkErrors.SDKError) {
	return NewCipher().DecryptStream(ctx, source, r)
}

// DecryptStream decrypts data from a reader using streaming mode.
// It sends the reader content as the request body and returns the decrypted
// plaintext bytes. The data is treated as binary (application/octet-stream)
// as decryption operates on raw encrypted bytes.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - r: io.Reader containing the encrypted data
//
// Returns:
//   - ([]byte, nil) containing the decrypted plaintext if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - Errors from streamPost(): if the streaming request fails
//   - ErrNetReadingResponseBody: if reading the response fails
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	cipher := NewCipher()
//	reader := bytes.NewReader(encryptedData)
//	plaintext, err := cipher.DecryptStream(ctx, source, reader)
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func (c *Cipher) DecryptStream(
	ctx context.Context, source *workloadapi.X509Source, r io.Reader,
) ([]byte, *sdkErrors.SDKError) {
	return c.streamOperation(ctx, source, r, url.CipherDecrypt(), "DecryptStream")
}

// Decrypt decrypts data with structured parameters using
// the default Cipher instance.
// It sends version, nonce, ciphertext, and algorithm and returns
// decrypted plaintext bytes.
//
// This is a convenience function that uses the default Cipher instance.
// For testing or custom configuration, create a Cipher instance directly.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - version: The cipher version used during encryption
//   - nonce: The nonce bytes used during encryption
//   - ciphertext: The encrypted data to decrypt
//   - algorithm: The encryption algorithm used (e.g., "AES-GCM")
//
// Returns:
//   - ([]byte, nil) containing the decrypted plaintext if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from httpPost(): if the HTTP request fails (e.g., ErrAPINotFound,
//     ErrAccessUnauthorized, ErrAPIBadRequest, ErrStateNotReady,
//     ErrNetPeerConnection)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	plaintext, err := Decrypt(ctx, source, 1, nonce, ciphertext, "AES-GCM")
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func Decrypt(
	ctx context.Context,
	source *workloadapi.X509Source,
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	return NewCipher().Decrypt(ctx, source, version, nonce, ciphertext, algorithm)
}

// Decrypt decrypts data with structured parameters.
// It sends version, nonce, ciphertext, and algorithm and returns
// decrypted plaintext bytes.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - version: The cipher version used during encryption
//   - nonce: The nonce bytes used during encryption
//   - ciphertext: The encrypted data to decrypt
//   - algorithm: The encryption algorithm used (e.g., "AES-GCM")
//
// Returns:
//   - ([]byte, nil) containing the decrypted plaintext if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from httpPost(): if the HTTP request fails (e.g., ErrAPINotFound,
//     ErrAccessUnauthorized, ErrAPIBadRequest, ErrStateNotReady,
//     ErrNetPeerConnection)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	cipher := NewCipher()
//	plaintext, err := cipher.Decrypt(ctx, source, 1, nonce, ciphertext, "AES-GCM")
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func (c *Cipher) Decrypt(
	ctx context.Context,
	source *workloadapi.X509Source,
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	payload := reqres.CipherDecryptRequest{
		Version:    version,
		Nonce:      nonce,
		Ciphertext: ciphertext,
		Algorithm:  algorithm,
	}

	var res reqres.CipherDecryptResponse
	if err := c.jsonOperation(
		ctx, source, payload, url.CipherDecrypt(), &res,
	); err != nil {
		return nil, err
	}

	return res.Plaintext, nil
}
