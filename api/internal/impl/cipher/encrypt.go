//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"context"
	"io"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// EncryptStream encrypts data from a reader using streaming mode using the
// default Cipher instance.
// It sends the reader content as the request body and returns the encrypted
// ciphertext bytes. The data is treated as binary (application/octet-stream)
// as encryption operates on raw bytes.
//
// This is a convenience function that uses the default Cipher instance.
// For testing or custom configuration, create a Cipher instance directly.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - r: io.Reader containing the data to encrypt
//
// Returns:
//   - ([]byte, nil) containing the encrypted ciphertext if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - Errors from streamPost(): if the streaming request fails
//   - ErrNetReadingResponseBody: if reading the response fails
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	reader := bytes.NewReader([]byte("sensitive data"))
//	ciphertext, err := EncryptStream(ctx, source, reader)
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
func EncryptStream(
	ctx context.Context, source *workloadapi.X509Source, r io.Reader,
) ([]byte, *sdkErrors.SDKError) {
	return NewCipher().EncryptStream(ctx, source, r)
}

// EncryptStream encrypts data from a reader using streaming mode.
// It sends the reader content as the request body and returns the encrypted
// ciphertext bytes. The data is treated as binary (application/octet-stream)
// as encryption operates on raw bytes.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - r: io.Reader containing the data to encrypt
//
// Returns:
//   - ([]byte, nil) containing the encrypted ciphertext if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - Errors from streamPost(): if the streaming request fails
//   - ErrNetReadingResponseBody: if reading the response fails
//
// Example:
//
//	cipher := NewCipher()
//	reader := bytes.NewReader([]byte("sensitive data"))
//	ciphertext, err := cipher.EncryptStream(ctx, source, reader)
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
func (c *Cipher) EncryptStream(
	ctx context.Context, source *workloadapi.X509Source, r io.Reader,
) ([]byte, *sdkErrors.SDKError) {
	return c.streamOperation(ctx, source, r, url.CipherEncrypt(), "EncryptStream")
}

// Encrypt encrypts data with structured parameters using the default Cipher
// instance, and returns the ciphertext together with the version and nonce
// required to decrypt it.
//
// It sends plaintext and algorithm to SPIKE Nexus and returns everything
// decryption needs.
//
// This is a convenience function that uses the default Cipher instance.
// For testing or custom configuration, create a Cipher instance directly.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - plaintext: The data to encrypt
//   - algorithm: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - (*data.EncryptedData, nil) containing the ciphertext and the parameters
//     needed to decrypt it if successful
//   - (nil, *sdkErrors.SDKError) if an error occurs:
//   - ErrSPIFFENilX509Source: if source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from httpPost(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	data := []byte("secret message")
//	encrypted, err := Encrypt(ctx, source, data, "AES-GCM")
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
func Encrypt(
	ctx context.Context, source *workloadapi.X509Source, plaintext []byte, algorithm string,
) (*data.EncryptedData, *sdkErrors.SDKError) {
	return NewCipher().Encrypt(ctx, source, plaintext, algorithm)
}

// Encrypt encrypts data with structured parameters and returns the ciphertext
// together with the version and nonce required to decrypt it.
//
// It sends plaintext and algorithm to SPIKE Nexus. The version and the nonce
// are chosen by SPIKE Nexus, so a caller cannot supply them, and the response
// is the only place they are returned.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - plaintext: The data to encrypt
//   - algorithm: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - (*data.EncryptedData, nil) containing the ciphertext and the parameters
//     needed to decrypt it if successful
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
//	cipher := NewCipher()
//	data := []byte("secret message")
//	encrypted, err := cipher.Encrypt(ctx, source, data, "AES-GCM")
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
//	// encrypted.Version, encrypted.Nonce and encrypted.Ciphertext are all
//	// needed to decrypt the payload later.
func (c *Cipher) Encrypt(
	ctx context.Context, source *workloadapi.X509Source, plaintext []byte, algorithm string,
) (*data.EncryptedData, *sdkErrors.SDKError) {
	payload := reqres.CipherEncryptRequest{
		Plaintext: plaintext,
		Algorithm: algorithm,
	}

	var res reqres.CipherEncryptResponse
	if err := c.jsonOperation(
		ctx, source, payload, url.CipherEncrypt(), &res,
	); err != nil {
		return nil, err
	}

	return &data.EncryptedData{
		Version:    res.Version,
		Nonce:      res.Nonce,
		Ciphertext: res.Ciphertext,
	}, nil
}
