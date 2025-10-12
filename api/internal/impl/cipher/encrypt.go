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
	apiErr "github.com/spiffe/spike-sdk-go/api/errors"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/net"
)

// indirections for testability within this package
var (
	createMTLSClient = net.CreateMTLSClientForNexus
	// streamPost                = net.StreamPost
	streamPostWithContentType = net.StreamPostWithContentType
	httpPost                  = net.Post
)

// EncryptStream encrypts data from a reader using streaming mode.
// It sends the reader content as the request body with the specified content
// type and returns the encrypted ciphertext bytes.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - r: io.Reader containing the data to encrypt
//   - contentType: Content type for the request (defaults to
//     "application/octet-stream" if empty)
//
// Returns:
//   - ([]byte, nil) containing the encrypted ciphertext if successful
//   - (nil, error) if an error occurs during encryption
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
//	ciphertext, err := EncryptStream(source, reader, "text/plain")
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
func EncryptStream(
	source *workloadapi.X509Source, r io.Reader, contentType string,
) ([]byte, error) {
	if source == nil {
		return []byte{}, errors.New("nil X509Source")
	}

	const fName = "encryptStream"

	client := createMTLSClient(source)

	if contentType == "" {
		contentType = "application/octet-stream"
	}
	rc, err := streamPostWithContentType(
		client, url.CipherEncrypt(), r, contentType,
	)
	if err != nil {
		return []byte{}, err
	}
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			log.Log().Info(fName,
				"message", "Error closing response body",
				"err", err.Error())
		}
	}(rc)
	b, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// EncryptJSON encrypts data using JSON mode with structured parameters.
// It sends plaintext and algorithm as JSON and returns encrypted ciphertext
// bytes.
//
// Parameters:
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - plaintext: The data to encrypt
//   - algorithm: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - ([]byte, nil) containing the encrypted ciphertext if successful
//   - ([]byte{}, nil) if the data is not found
//   - (nil, error) if an error occurs during encryption
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
//	ciphertext, err := EncryptJSON(source, data, "AES-GCM")
//	if err != nil {
//	    log.Printf("Encryption failed: %v", err)
//	}
func EncryptJSON(
	source *workloadapi.X509Source, plaintext []byte, algorithm string,
) ([]byte, error) {
	if source == nil {
		return []byte{}, errors.New("nil X509Source")
	}

	client := createMTLSClient(source)

	payload := reqres.CipherEncryptRequest{
		Plaintext: plaintext,
		Algorithm: algorithm,
	}
	mr, err := json.Marshal(payload)
	if err != nil {
		return []byte{},
			errors.Join(errors.New("cipher.EncryptJSON: marshal request"), err)
	}
	body, err := httpPost(client, url.CipherEncrypt(), mr)
	if err != nil {
		if errors.Is(err, apiErr.ErrNotFound) {
			return []byte{}, nil
		}
		return []byte{}, err
	}
	var res reqres.CipherEncryptResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return []byte{},
			errors.Join(errors.New("cipher.EncryptJSON: unmarshal response"), err)
	}
	if res.Err != "" {
		return []byte{}, errors.New(string(res.Err))
	}
	return res.Ciphertext, nil
}
