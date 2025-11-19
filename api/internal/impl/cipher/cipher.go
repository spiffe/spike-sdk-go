//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"io"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"

	"github.com/spiffe/spike-sdk-go/net"
)

// Cipher encapsulates cipher operations with configurable HTTP client
// dependencies. This struct-based approach enables clean dependency injection
// for testing without relying on the global mutable state.
//
// The zero value is not usable; instances should be created using NewCipher().
type Cipher struct {
	// createMTLSHTTPClientFromSource creates an mTLS HTTP client
	// from an X509Source
	createMTLSHTTPClientFromSource func(*workloadapi.X509Source) *http.Client

	// httpPost performs a POST request and returns the response body
	httpPost func(*http.Client, string, []byte) ([]byte, *sdkErrors.SDKError)

	// streamPost performs a streaming POST request with a custom content type
	streamPost func(
		*http.Client, string, io.Reader, net.ContentType,
	) (io.ReadCloser, *sdkErrors.SDKError)
}

// NewCipher creates a new Cipher instance with default production dependencies.
// The returned Cipher is ready to use for encryption and decryption operations.
//
// For testing, create a Cipher with custom dependencies by directly
// constructing the struct with test doubles.
//
// Example:
//
//	cipher := NewCipher()
//	plaintext, err := cipher.EncryptJSON(source, data, "AES-GCM")
func NewCipher() *Cipher {
	return &Cipher{
		createMTLSHTTPClientFromSource: net.CreateMTLSClientForNexus,
		httpPost:                       net.Post,
		streamPost:                     net.StreamPostWithContentType,
	}
}
