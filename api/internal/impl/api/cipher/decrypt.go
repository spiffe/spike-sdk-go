//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	reqres "github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	apiErr "github.com/spiffe/spike-sdk-go/api/errors"
	"github.com/spiffe/spike-sdk-go/api/url"
)

// Decrypt decrypts via streaming or JSON based on mode.
// Stream mode: send r as body with contentType. Returns plaintext bytes.
// JSON mode: send version, nonce, ciphertext, algorithm; returns plaintext bytes.
func Decrypt(
	source *workloadapi.X509Source,
	mode Mode,
	r io.Reader,
	contentType string,
	version byte,
	nonce, ciphertext []byte,
	algorithm string,
) ([]byte, error) {
	client, err := createMTLSClient(source)
	if err != nil {
		return nil, err
	}

	switch mode {
	case ModeStream:
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		rc, err := streamPostWithContentType(client, url.CipherDecrypt(), r, contentType)
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		b, err := io.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		return b, nil

	case ModeJSON:
		payload := reqres.CipherDecryptRequest{Version: version, Nonce: nonce, Ciphertext: ciphertext, Algorithm: algorithm}
		mr, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.Join(errors.New("cipher.Decrypt: marshal request"), err)
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
			return nil, errors.Join(errors.New("cipher.Decrypt: unmarshal response"), err)
		}
		if res.Err != "" {
			return nil, errors.New(string(res.Err))
		}
		return res.Plaintext, nil
	}

	return nil, errors.New("cipher.Decrypt: unsupported mode")
}
