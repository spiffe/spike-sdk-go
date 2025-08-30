//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func TestDecrypt_OctetStream(t *testing.T) {
	origCreate := createMTLSClient
	origPost := streamPost
	origPostCT := streamPostWithContentType
	t.Cleanup(func() { createMTLSClient = origCreate; streamPost = origPost; streamPostWithContentType = origPostCT })

	createMTLSClient = func(_ *workloadapi.X509Source) (*http.Client, error) {
		return &http.Client{}, nil
	}
	streamPostWithContentType = func(_ *http.Client, _ string, body io.Reader, ct string) (io.ReadCloser, error) {
		if ct != "application/octet-stream" {
			t.Fatalf("unexpected ct: %s", ct)
		}
		b, _ := io.ReadAll(body)
		if string(b) != "cipher" {
			t.Fatalf("unexpected body: %q", string(b))
		}
		return io.NopCloser(bytes.NewReader([]byte("plain"))), nil
	}

	out, err := Decrypt(nil, ModeStream, bytes.NewReader([]byte("cipher")), "application/octet-stream", 0, nil, nil, "")
	if err != nil {
		t.Fatalf("Decrypt error: %v", err)
	}
	if string(out) != "plain" {
		t.Fatalf("unexpected out: %s", string(out))
	}
}
