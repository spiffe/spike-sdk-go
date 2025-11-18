//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

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
//   - (nil, error) if an error occurs during decryption
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
	source *workloadapi.X509Source, r io.Reader, contentType string,
) ([]byte, error) {
	if source == nil {
		return []byte{}, sdkErrors.ErrNilX509Source
	}

	const fName = "decryptStream"

	client := createMTLSClient(source)

	if contentType == "" {
		contentType = "application/octet-stream"
	}
	rc, err := streamPostWithContentType(
		client, url.CipherDecrypt(), r, contentType,
	)
	if err != nil {
		return []byte{}, err
	}
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			log.Log().Error(fName, "err", err.Error())
		}
	}(rc)
	b, err := io.ReadAll(rc)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
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
//   - ([]byte{}, nil) if the data is not found
//   - (nil, error) if an error occurs during decryption
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
) ([]byte, error) {
	if source == nil {
		return []byte{}, sdkErrors.ErrNilX509Source
	}

	client := createMTLSClient(source)

	payload := reqres.CipherDecryptRequest{
		Version:    version,
		Nonce:      nonce,
		Ciphertext: ciphertext,
		Algorithm:  algorithm,
	}
	mr, err := json.Marshal(payload)
	if err != nil {
		return []byte{},
			errors.Join(errors.New("cipher.DecryptJSON: marshal request"), err)
	}
	body, err := httpPost(client, url.CipherDecrypt(), mr)
	if err != nil {
		if errors.Is(err, sdkErrors.ErrNotFound) {
			return []byte{}, nil
		}
		return []byte{}, err
	}
	var res reqres.CipherDecryptResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return []byte{},
			errors.Join(errors.New("cipher.DecryptJSON: unmarshal response"), err)
	}
	if res.Err != "" {
		return []byte{}, errors.New(string(res.Err))
	}
	return res.Plaintext, nil
}
