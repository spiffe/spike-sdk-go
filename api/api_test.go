//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/api/cipher"
	"github.com/spiffe/spike-sdk-go/predicate"
)

// stub source (nil pointer acceptable for our stubs)
var fakeSource *workloadapi.X509Source

func TestAPI_CipherStreamMethods(t *testing.T) {
	a := NewWithSource(fakeSource, predicate.AllowAll)

	// backup and restore
	origEnc, origDec := cipherEncryptFunc, cipherDecryptFunc
	t.Cleanup(func() { cipherEncryptFunc, cipherDecryptFunc = origEnc, origDec })

	cipherEncryptFunc = func(_ *workloadapi.X509Source,
		mode cipher.Mode, r io.Reader, contentType string,
		_ []byte, _ string, _ predicate.Predicate,
	) ([]byte, error) {
		if mode != cipher.ModeStream || contentType != "application/octet-stream" {
			return nil, errors.New("bad mode or content-type")
		}
		b, _ := io.ReadAll(r)
		if string(b) != "plain" {
			return nil, errors.New("unexpected input")
		}
		return []byte("cipher"), nil
	}
	cipherDecryptFunc = func(_ *workloadapi.X509Source, mode cipher.Mode,
		r io.Reader, contentType string, _ byte, _, _ []byte, _ string,
		_ predicate.Predicate) ([]byte, error) {
		if mode != cipher.ModeStream || contentType != "application/octet-stream" {
			return nil, errors.New("bad mode or content-type")
		}
		b, _ := io.ReadAll(r)
		if string(b) != "cipher" {
			return nil, errors.New("unexpected input")
		}
		return []byte("plain"), nil
	}

	out, err := a.CipherEncryptStream(
		bytes.NewReader([]byte("plain")), "application/octet-stream")
	if err != nil {
		t.Fatalf("CipherEncryptStream error: %v", err)
	}
	if string(out) != "cipher" {
		t.Fatalf("unexpected encrypt out: %s", string(out))
	}

	out2, err := a.CipherDecryptStream(
		bytes.NewReader([]byte("cipher")), "application/octet-stream")
	if err != nil {
		t.Fatalf("CipherDecryptStream error: %v", err)
	}
	if string(out2) != "plain" {
		t.Fatalf("unexpected decrypt out: %s", string(out2))
	}

	// error path
	cipherEncryptFunc = func(_ *workloadapi.X509Source,
		_ cipher.Mode, _ io.Reader, _ string, _ []byte, _ string,
		_ predicate.Predicate) ([]byte, error) {
		return nil, errors.New("boom")
	}
	if _, err := a.CipherEncryptStream(
		bytes.NewReader(nil), "application/octet-stream"); err == nil {
		t.Fatalf("expected error from CipherEncryptStream")
	}
}

func TestAPI_CipherJSONMethods(t *testing.T) {
	a := NewWithSource(fakeSource, predicate.AllowAll)

	origEnc, origDec := cipherEncryptFunc, cipherDecryptFunc
	t.Cleanup(func() { cipherEncryptFunc, cipherDecryptFunc = origEnc, origDec })

	cipherEncryptFunc = func(
		_ *workloadapi.X509Source, mode cipher.Mode, _ io.Reader,
		_ string, plaintext []byte, algorithm string, _ predicate.Predicate,
	) ([]byte, error) {
		if mode != cipher.ModeJSON {
			return nil, errors.New("bad mode")
		}
		if string(plaintext) != "p" || algorithm != "alg" {
			return nil, errors.New("bad input")
		}
		return []byte{2}, nil
	}
	cipherDecryptFunc = func(
		_ *workloadapi.X509Source, mode cipher.Mode, _ io.Reader,
		_ string, version byte, _, _ []byte,
		algorithm string, _ predicate.Predicate,
	) ([]byte, error) {
		if mode != cipher.ModeJSON {
			return nil, errors.New("bad mode")
		}
		if version != 1 || algorithm != "alg" {
			return nil, errors.New("bad input")
		}
		return []byte("p"), nil
	}

	out, err := a.CipherEncryptJSON([]byte("p"), "alg")
	if err != nil {
		t.Fatalf("CipherEncryptJSON error: %v", err)
	}
	if len(out) != 1 {
		t.Fatalf("unexpected encrypt json response length: %d", len(out))
	}

	outp, err := a.CipherDecryptJSON(1, []byte{1}, []byte{2}, "alg")
	if err != nil {
		t.Fatalf("CipherDecryptJSON error: %v", err)
	}
	if string(outp) != "p" {
		t.Fatalf("unexpected decrypt json: %s", string(outp))
	}
}
