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
	"github.com/spiffe/spike-sdk-go/predicate"
)

// Decrypt decrypts data using SPIKE's cipher service, supporting both streaming
// and JSON modes. It requires a SPIFFE X.509 source for establishing a mutual
// TLS connection to make the decryption request.
//
// The function supports two modes:
//   - Stream mode: Sends the reader content as the request body with the
//     specified content type. Returns decrypted plaintext bytes.
//   - JSON mode: Sends structured data (version, nonce, ciphertext, algorithm)
//     as JSON. Returns decrypted plaintext bytes.
//
// The function takes the following parameters:
//   - source: A pointer to a workloadapi.X509Source for establishing mTLS
//     connection
//   - mode: The decryption mode (ModeStream or ModeJSON)
//   - r: An io.Reader containing data to decrypt (used in stream mode)
//   - contentType: The content type for stream mode (defaults to
//     "application/octet-stream" if empty)
//   - version: The cipher version (used in JSON mode)
//   - nonce: The nonce bytes (used in JSON mode)
//   - ciphertext: The encrypted data bytes (used in JSON mode)
//   - algorithm: The encryption algorithm name (used in JSON mode)
//   - allow: A predicate.Predicate that determines which server certificates
//     to trust during the mTLS connection
//
// The function returns:
//   - ([]byte, nil) containing the decrypted plaintext if successful
//   - (nil, nil) if the data is not found (JSON mode only)
//   - (nil, error) if an error occurs during the operation
//
// Errors can occur during:
//   - Creating the mTLS client
//   - Marshaling the decryption request (JSON mode)
//   - Making the HTTP POST request
//   - Reading the response (stream mode)
//   - Unmarshaling the response (JSON mode)
//   - Server-side decryption (indicated in the response)
//
// Example usage (JSON mode):
//
//	source, err := workloadapi.NewX509Source(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	plaintext, err := Decrypt(
//	    source,
//	    ModeJSON,
//	    nil, // not used in JSON mode
//	    "",  // not used in JSON mode
//	    1,   // version
//	    nonce,
//	    ciphertext,
//	    "AES-GCM",
//	    predicate.AllowAll,
//	)
//	if err != nil {
//	    log.Printf("Failed to decrypt: %v", err)
//	    return
//	}
//
// Example usage (Stream mode):
//
//	reader := bytes.NewReader(encryptedData)
//	plaintext, err := Decrypt(
//	    source,
//	    ModeStream,
//	    reader,
//	    "application/octet-stream",
//	    0,   // not used in stream mode
//	    nil, // not used in stream mode
//	    nil, // not used in stream mode
//	    "",  // not used in stream mode
//	    predicate.AllowAll,
//	)
func Decrypt(
	source *workloadapi.X509Source,
	mode Mode,
	r io.Reader,
	contentType string, version byte,
	nonce, ciphertext []byte, algorithm string,
	allow predicate.Predicate,
) ([]byte, error) {
	const fName = "Decrypt"

	client, err := createMTLSClient(source, allow)
	if err != nil {
		return nil, err
	}

	switch mode {
	case ModeStream:
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		rc, err := streamPostWithContentType(
			client, url.CipherDecrypt(), r, contentType,
		)
		if err != nil {
			return nil, err
		}
		defer func(rc io.ReadCloser) {
			err := rc.Close()
			if err != nil {
				log.Log().Error(fName, "err", err.Error())
			}
		}(rc)
		b, err := io.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		return b, nil

	case ModeJSON:
		payload := reqres.CipherDecryptRequest{
			Version:    version,
			Nonce:      nonce,
			Ciphertext: ciphertext,
			Algorithm:  algorithm,
		}
		mr, err := json.Marshal(payload)
		if err != nil {
			return nil,
				errors.Join(errors.New("cipher.Decrypt: marshal request"), err)
		}
		body, err := httpPost(client, url.CipherDecrypt(), mr)
		if err != nil {
			if errors.Is(err, apiErr.ErrNotFound) {
				return nil, nil
			}
			return nil, err
		}
		var res reqres.CipherDecryptResponse
		if err := json.Unmarshal(body, &res); err != nil {
			return nil,
				errors.Join(errors.New("cipher.Decrypt: unmarshal response"), err)
		}
		if res.Err != "" {
			return nil, errors.New(string(res.Err))
		}
		return res.Plaintext, nil
	}

	return nil, errors.New("cipher.Decrypt: unsupported mode")
}
