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

func TestDecrypt_OctetStream(t *testing.T) {
	// Create a Cipher with test doubles injected
	cipher := &Cipher{
		createMTLSHTTPClientFromSource: func(_ *workloadapi.X509Source) *http.Client {
			return fakeClient(rtFunc(func(_ *http.Request) (*http.Response, error) {
				return nil, nil
			}))
		},
		streamPost: func(_ context.Context, _ *http.Client, _ string, body io.Reader) (io.ReadCloser, *sdkErrors.SDKError) {
			b, _ := io.ReadAll(body)
			if string(b) != "cipher" {
				t.Fatalf("unexpected body: %q", string(b))
			}
			return io.NopCloser(bytes.NewReader([]byte("plain"))), nil
		},
		httpPost: func(_ context.Context, _ *http.Client, _ string, _ []byte) ([]byte, *sdkErrors.SDKError) {
			return nil, nil
		},
	}

	out, err := cipher.DecryptStream(
		context.Background(), &workloadapi.X509Source{}, bytes.NewReader([]byte("cipher")),
	)
	if err != nil {
		t.Fatalf("DecryptStream error: %v", err)
	}
	if string(out) != "plain" {
		t.Fatalf("unexpected out: %s", string(out))
	}
}
