//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"encoding/json"
	"io"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/net"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// DecryptStream decrypts data from a reader using streaming mode using the
// default Cipher instance.
// It sends the reader content as the request body with the specified content
// type and returns the decrypted plaintext bytes.
//
// This is a convenience function that uses the default Cipher instance.
// For testing or custom configuration, create a Cipher instance directly.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - r: io.Reader containing the encrypted data
//   - contentType: Content type for the request (defaults to
//     "application/octet-stream" if empty)
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
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	reader := bytes.NewReader(encryptedData)
//	plaintext, err := DecryptStream(source, reader, "application/octet-stream")
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func DecryptStream(
	source *workloadapi.X509Source, r io.Reader, contentType net.ContentType,
) ([]byte, *sdkErrors.SDKError) {
	return NewCipher().DecryptStream(source, r, contentType)
}

// DecryptStream decrypts data from a reader using streaming mode.
// It sends the reader content as the request body with the specified content
// type and returns the decrypted plaintext bytes.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - r: io.Reader containing the encrypted data
//   - contentType: Content type for the request (defaults to
//     "application/octet-stream" if empty)
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
//	cipher := NewCipher()
//	reader := bytes.NewReader(encryptedData)
//	plaintext, err := cipher.DecryptStream(source, reader, "application/octet-stream")
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func (c *Cipher) DecryptStream(
	source *workloadapi.X509Source, r io.Reader, contentType net.ContentType,
) ([]byte, *sdkErrors.SDKError) {
	const fName = "DecryptStream"

	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source
	}

	client := c.createMTLSHTTPClientFromSource(source)

	if contentType == "" {
		contentType = net.ContentTypeOctetStream
	}

	rc, err := c.streamPost(client, url.CipherDecrypt(), r, contentType)
	if err != nil {
		return nil, err
	}

	defer func(rc io.ReadCloser) {
		if rc == nil {
			return
		}
		err := rc.Close()
		if err != nil {
			failErr := sdkErrors.ErrFSStreamCloseFailed.Wrap(err)
			failErr.Msg = "failed to close response body"
			log.WarnErr(fName, *failErr)
		}
	}(rc)

	b, err := io.ReadAll(rc)
	if err != nil {
		failErr := sdkErrors.ErrNetReadingResponseBody.Wrap(err)
		failErr.Msg = "failed to read response body"
		return nil, failErr
	}

	return b, nil
}

// DecryptJSON decrypts data using JSON mode with structured parameters using
// the default Cipher instance.
// It sends version, nonce, ciphertext, and algorithm as JSON and returns
// decrypted plaintext bytes.
//
// This is a convenience function that uses the default Cipher instance.
// For testing or custom configuration, create a Cipher instance directly.
//
// Parameters:
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
//   - Errors from httpPost(): if the HTTP request fails (e.g., ErrNotFound,
//     ErrAccessUnauthorized, ErrBadRequest, ErrStateNotReady, ErrNetPeerConnection)
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
//	plaintext, err := DecryptJSON(source, 1, nonce, ciphertext, "AES-GCM")
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func DecryptJSON(
	source *workloadapi.X509Source,
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	return NewCipher().DecryptJSON(source, version, nonce, ciphertext, algorithm)
}

// DecryptJSON decrypts data using JSON mode with structured parameters.
// It sends version, nonce, ciphertext, and algorithm as JSON and returns
// decrypted plaintext bytes.
//
// Parameters:
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
//   - Errors from httpPost(): if the HTTP request fails (e.g., ErrNotFound,
//     ErrAccessUnauthorized, ErrBadRequest, ErrStateNotReady, ErrNetPeerConnection)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	cipher := NewCipher()
//	plaintext, err := cipher.DecryptJSON(source, 1, nonce, ciphertext, "AES-GCM")
//	if err != nil {
//	    log.Printf("Decryption failed: %v", err)
//	}
func (c *Cipher) DecryptJSON(
	source *workloadapi.X509Source,
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source
	}

	client := c.createMTLSHTTPClientFromSource(source)

	payload := reqres.CipherDecryptRequest{
		Version:    version,
		Nonce:      nonce,
		Ciphertext: ciphertext,
		Algorithm:  algorithm,
	}

	mr, err := json.Marshal(payload)
	if err != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(err)
		failErr.Msg = "problem generating the payload"
		return nil, failErr
	}

	body, err := c.httpPost(client, url.CipherDecrypt(), mr)
	if err != nil {
		return nil, err
	}

	var res reqres.CipherDecryptResponse
	if err := json.Unmarshal(body, &res); err != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(err)
		failErr.Msg = "problem parsing response body"
		return nil, failErr
	}
	if res.Err != "" {
		return nil, sdkErrors.FromCode(res.Err)
	}
	return res.Plaintext, nil
}
