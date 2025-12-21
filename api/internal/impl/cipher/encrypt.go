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
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
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
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
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

// Encrypt encrypts data with structured parameters using
// the default Cipher instance.
// It sends plaintext and algorithm and returns encrypted ciphertext
// bytes.
//
// This is a convenience function that uses the default Cipher instance.
// For testing or custom configuration, create a Cipher instance directly.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - plaintext: The data to encrypt
//   - algorithm: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - ([]byte, nil) containing the encrypted ciphertext if successful
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
//	data := []byte("secret message")
//	ciphertext, err := Encrypt(ctx, source, data, "AES-GCM")
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
func Encrypt(
	ctx context.Context,
	source *workloadapi.X509Source, plaintext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	return NewCipher().Encrypt(ctx, source, plaintext, algorithm)
}

// Encrypt encrypts data with structured parameters.
// It sends plaintext and algorithm and returns encrypted ciphertext
// bytes.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - plaintext: The data to encrypt
//   - algorithm: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - ([]byte, nil) containing the encrypted ciphertext if successful
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
//	data := []byte("secret message")
//	ciphertext, err := cipher.Encrypt(ctx, source, data, "AES-GCM")
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
func (c *Cipher) Encrypt(
	ctx context.Context,
	source *workloadapi.X509Source, plaintext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
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

	return res.Ciphertext, nil
}
