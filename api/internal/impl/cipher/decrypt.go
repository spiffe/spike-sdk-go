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
)

// DecryptStream decrypts data from a reader using streaming mode.
// It sends the reader content as the request body with the specified
// content type.
// Returns the decrypted plaintext bytes.
func DecryptStream(
	source *workloadapi.X509Source, r io.Reader, contentType string,
) ([]byte, error) {
	if source == nil {
		return []byte{}, errors.New("nil X509Source")
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
func DecryptJSON(
	source *workloadapi.X509Source,
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, error) {
	if source == nil {
		return []byte{}, errors.New("nil X509Source")
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
		if errors.Is(err, apiErr.ErrNotFound) {
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
