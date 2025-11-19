//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func TestDecrypt_OctetStream(t *testing.T) {
	// Create a Cipher with test doubles injected
	cipher := &Cipher{
		createMTLSHTTPClientFromSource: func(_ *workloadapi.X509Source) *http.Client {
			return fakeClient(rtFunc(func(_ *http.Request) (*http.Response, error) {
				return nil, nil
			}))
		},
		streamPost: func(_ *http.Client, _ string, body io.Reader, ct string) (io.ReadCloser, error) {
			if ct != "application/octet-stream" {
				t.Fatalf("unexpected ct: %s", ct)
			}
			b, _ := io.ReadAll(body)
			if string(b) != "cipher" {
				t.Fatalf("unexpected body: %q", string(b))
			}
			return io.NopCloser(bytes.NewReader([]byte("plain"))), nil
		},
		httpPost: func(_ *http.Client, _ string, _ []byte) ([]byte, error) {
			return nil, nil
		},
	}

	out, err := cipher.DecryptStream(
		&workloadapi.X509Source{}, bytes.NewReader([]byte("cipher")),
		"application/octet-stream",
	)
	if err != nil {
		t.Fatalf("DecryptStream error: %v", err)
	}
	if string(out) != "plain" {
		t.Fatalf("unexpected out: %s", string(out))
	}
}
