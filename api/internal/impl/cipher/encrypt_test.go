//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
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

// testNonce is the 12 bytes AES-GCM uses, which is the size SPIKE Nexus
// accepts, so the fixtures match what the server would produce.
var testNonce = []byte("0123456789ab")

// jsonCipher returns a Cipher whose JSON API calls are served by f, so the
// JSON path can be exercised without a server.
func jsonCipher(
	f func(context.Context, *http.Client, string, []byte) ([]byte, *sdkErrors.SDKError),
) *Cipher {
	return &Cipher{
		createMTLSHTTPClientFromSource: func(_ *workloadapi.X509Source) *http.Client {
			return fakeClient(rtFunc(func(_ *http.Request) (*http.Response, error) {
				return nil, nil
			}))
		},
		httpPost: f,
	}
}

func TestEncryptDataReturnsVersionNonceAndCiphertext(t *testing.T) {
	body, marshalErr := json.Marshal(reqres.CipherEncryptResponse{
		Version:    1,
		Nonce:      testNonce,
		Ciphertext: []byte("a ciphertext"),
	})
	if marshalErr != nil {
		t.Fatalf("marshal response: %v", marshalErr)
	}

	cipher := jsonCipher(func(
		_ context.Context, _ *http.Client, path string, _ []byte,
	) ([]byte, *sdkErrors.SDKError) {
		if path == "" {
			t.Fatalf("empty path")
		}
		return body, nil
	})

	encrypted, err := cipher.EncryptData(
		context.Background(), &workloadapi.X509Source{}, []byte("plain"), "AES-GCM",
	)
	if err != nil {
		t.Fatalf("EncryptData error: %v", err)
	}
	if encrypted == nil {
		t.Fatal("EncryptData returned no data")
	}
	if encrypted.Version != 1 {
		t.Errorf("unexpected version: %d", encrypted.Version)
	}
	if !bytes.Equal(encrypted.Nonce, testNonce) {
		t.Errorf("unexpected nonce: %q", string(encrypted.Nonce))
	}
	if string(encrypted.Ciphertext) != "a ciphertext" {
		t.Errorf("unexpected ciphertext: %q", string(encrypted.Ciphertext))
	}
}

func TestEncryptStillReturnsOnlyCiphertext(t *testing.T) {
	body, marshalErr := json.Marshal(reqres.CipherEncryptResponse{
		Version:    1,
		Nonce:      testNonce,
		Ciphertext: []byte("a ciphertext"),
	})
	if marshalErr != nil {
		t.Fatalf("marshal response: %v", marshalErr)
	}

	cipher := jsonCipher(func(
		_ context.Context, _ *http.Client, _ string, _ []byte,
	) ([]byte, *sdkErrors.SDKError) {
		return body, nil
	})

	out, err := cipher.Encrypt(
		context.Background(), &workloadapi.X509Source{}, []byte("plain"), "AES-GCM",
	)
	if err != nil {
		t.Fatalf("Encrypt error: %v", err)
	}
	if string(out) != "a ciphertext" {
		t.Fatalf("unexpected out: %q", string(out))
	}
}

// The parameters EncryptData returns are the ones Decrypt needs, which is what
// makes a JSON-mode encrypted value possible to decrypt again.
func TestEncryptDataRoundTripsThroughDecrypt(t *testing.T) {
	plaintext := []byte("secret message")
	nonce := testNonce
	ciphertext := []byte("a ciphertext")

	encryptBody, marshalErr := json.Marshal(reqres.CipherEncryptResponse{
		Version:    1,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	})
	if marshalErr != nil {
		t.Fatalf("marshal encrypt response: %v", marshalErr)
	}
	decryptBody, marshalErr := json.Marshal(reqres.CipherDecryptResponse{
		Plaintext: plaintext,
	})
	if marshalErr != nil {
		t.Fatalf("marshal decrypt response: %v", marshalErr)
	}

	var decryptRequest reqres.CipherDecryptRequest
	cipher := jsonCipher(func(
		_ context.Context, _ *http.Client, path string, payload []byte,
	) ([]byte, *sdkErrors.SDKError) {
		switch path {
		case url.CipherEncrypt():
			return encryptBody, nil
		case url.CipherDecrypt():
			if unmarshalErr := json.Unmarshal(payload, &decryptRequest); unmarshalErr != nil {
				t.Fatalf("unmarshal decrypt request: %v", unmarshalErr)
			}
			return decryptBody, nil
		}

		t.Fatalf("unexpected path: %q", path)
		return nil, nil
	})

	ctx := context.Background()
	source := &workloadapi.X509Source{}

	encrypted, err := cipher.EncryptData(ctx, source, []byte("plain"), "AES-GCM")
	if err != nil {
		t.Fatalf("EncryptData error: %v", err)
	}

	decrypted, err := cipher.Decrypt(
		ctx, source, encrypted.Version, encrypted.Nonce, encrypted.Ciphertext, "AES-GCM",
	)
	if err != nil {
		t.Fatalf("Decrypt error: %v", err)
	}

	// Checked against the response the server sent as well as against each
	// other, so dropping either field anywhere on the way is caught.
	if decryptRequest.Version != 1 || encrypted.Version != 1 {
		t.Errorf("unexpected version: decrypt got %d, encrypt returned %d",
			decryptRequest.Version, encrypted.Version)
	}
	if !bytes.Equal(decryptRequest.Nonce, nonce) || !bytes.Equal(encrypted.Nonce, nonce) {
		t.Errorf("unexpected nonce: decrypt got %q, encrypt returned %q",
			string(decryptRequest.Nonce), string(encrypted.Nonce))
	}
	if !bytes.Equal(decryptRequest.Ciphertext, encrypted.Ciphertext) {
		t.Errorf("decrypt got ciphertext %q, encrypt returned %q",
			string(decryptRequest.Ciphertext), string(encrypted.Ciphertext))
	}
	if !bytes.Equal(decrypted, plaintext) {
		t.Errorf("unexpected plaintext: %q", string(decrypted))
	}
}
