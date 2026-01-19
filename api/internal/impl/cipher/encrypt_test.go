//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

type rtFunc func(*http.Request) (*http.Response, error)

func (f rtFunc) RoundTrip(_ *http.Request) (*http.Response, error) {
	return f(nil)
}

func fakeClient(rt http.RoundTripper) *http.Client {
	return &http.Client{Transport: rt}
}

func TestEncryptOctetStream(t *testing.T) {
	// Create a Cipher with test doubles injected
	cipher := &Cipher{
		createMTLSHTTPClientFromSource: func(_ *workloadapi.X509Source) *http.Client {
			return fakeClient(rtFunc(func(_ *http.Request) (*http.Response, error) {
				return nil, nil
			}))
		},
		streamPost: func(_ context.Context, _ *http.Client, path string, body io.Reader) (io.ReadCloser, *sdkErrors.SDKError) {
			if path == "" {
				t.Fatalf("empty path")
			}
			b, _ := io.ReadAll(body)
			if string(b) != "plain" {
				t.Fatalf("unexpected body: %q", string(b))
			}
			return io.NopCloser(bytes.NewReader([]byte("cipher"))), nil
		},
		httpPost: func(_ context.Context, _ *http.Client, _ string, _ []byte) ([]byte, *sdkErrors.SDKError) {
			return nil, nil
		},
	}

	out, err := cipher.EncryptStream(
		context.Background(), &workloadapi.X509Source{}, bytes.NewReader([]byte("plain")),
	)
	if err != nil {
		t.Fatalf("EncryptStream error: %v", err)
	}
	if string(out) != "cipher" {
		t.Fatalf("unexpected out: %s", string(out))
	}
}
