//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/log"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	apiErr "github.com/spiffe/spike-sdk-go/api/errors"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/net"
	"github.com/spiffe/spike-sdk-go/predicate"
)

// indirections for testability within this package
var (
	createMTLSClient          = net.CreateMTLSClientWithPredicate
	streamPost                = net.StreamPost
	streamPostWithContentType = net.StreamPostWithContentType
	httpPost                  = net.Post
)

// Encrypt encrypts either via streaming or JSON based on mode.
// Stream mode: send r as body with contentType. Returns ciphertext bytes.
// JSON mode: send plaintext + algorithm; returns ciphertext bytes.
func Encrypt(
	source *workloadapi.X509Source, mode Mode, r io.Reader,
	contentType string, plaintext []byte, algorithm string,
	allow predicate.Predicate,
) ([]byte, error) {
	const fName = "encrypt"

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
			client, url.CipherEncrypt(), r, contentType,
		)
		if err != nil {
			return nil, err
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

	case ModeJSON:
		payload := reqres.CipherEncryptRequest{
			Plaintext: plaintext,
			Algorithm: algorithm,
		}
		mr, err := json.Marshal(payload)
		if err != nil {
			return nil,
				errors.Join(errors.New("cipher.Encrypt: marshal request"), err)
		}
		body, err := httpPost(client, url.CipherEncrypt(), mr)
		if err != nil {
			if errors.Is(err, apiErr.ErrNotFound) {
				return nil, nil
			}
			return nil, err
		}
		var res reqres.CipherEncryptResponse
		if err := json.Unmarshal(body, &res); err != nil {
			return nil,
				errors.Join(errors.New("cipher.Encrypt: unmarshal response"), err)
		}
		if res.Err != "" {
			return nil, errors.New(string(res.Err))
		}
		return res.Ciphertext, nil
	}

	return nil, errors.New("cipher.Encrypt: unsupported mode")
}
