//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
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
	httpPost func(
		context.Context, *http.Client, string, []byte,
	) ([]byte, *sdkErrors.SDKError)

	// streamPost performs a streaming POST request with binary data
	// (always uses application/octet-stream content type)
	streamPost func(
		context.Context, *http.Client, string, io.Reader,
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
//	plaintext, err := cipher.Encrypt(source, data, "AES-GCM")
func NewCipher() *Cipher {
	return &Cipher{
		createMTLSHTTPClientFromSource: net.CreateMTLSClientForNexus,
		httpPost:                       net.Post,
		streamPost:                     net.StreamPost,
	}
}

// streamOperation performs a streaming encryption or decryption operation.
// This is a common helper that removes duplication between EncryptStream
// and DecryptStream.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection
//   - r: io.Reader containing the data to process
//   - urlPath: The API endpoint URL
//   - fName: Function name for logging purposes
//
// Returns:
//   - []byte: The processed data if successful
//   - *sdkErrors.SDKError: Error if the operation fails
func (c *Cipher) streamOperation(
	ctx context.Context,
	source *workloadapi.X509Source,
	r io.Reader,
	urlPath string,
	fName string,
) ([]byte, *sdkErrors.SDKError) {
	if source == nil {
		return nil, sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	client := c.createMTLSHTTPClientFromSource(source)
	rc, err := c.streamPost(ctx, client, urlPath, r)
	if err != nil {
		return nil, err
	}

	defer func(rc io.ReadCloser) {
		if rc == nil {
			return
		}
		closeErr := rc.Close()
		if closeErr != nil {
			failErr := sdkErrors.ErrFSStreamCloseFailed.Wrap(closeErr)
			failErr.Msg = "failed to close response body"
			log.WarnErr(fName, *failErr)
		}
	}(rc)

	b, readErr := io.ReadAll(rc)
	if readErr != nil {
		failErr := sdkErrors.ErrNetReadingResponseBody.Wrap(readErr)
		failErr.Msg = "failed to read response body"
		return nil, failErr
	}
	return b, nil
}

// jsonOperation performs a JSON-based operation with generic request/response
// handling. This helper removes duplication between Encrypt and Decrypt
// operations.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection
//   - request: The request payload (will be marshaled to JSON)
//   - urlPath: The API endpoint URL
//   - response: Pointer to response struct that implements ResponseWithError
//
// Returns:
//   - *sdkErrors.SDKError: Error if the operation fails, nil on success
func (c *Cipher) jsonOperation(
	ctx context.Context,
	source *workloadapi.X509Source, request any, urlPath string, response any,
) *sdkErrors.SDKError {
	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	client := c.createMTLSHTTPClientFromSource(source)

	mr, marshalErr := json.Marshal(request)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "problem generating the payload"
		return failErr
	}

	body, err := c.httpPost(ctx, client, urlPath, mr)
	if err != nil {
		return err
	}

	if unmarshalErr := json.Unmarshal(body, response); unmarshalErr != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(unmarshalErr)
		failErr.Msg = "problem parsing response body"
		return failErr
	}

	// Type assertion to check error code
	// Doing this with generics would be tricky in Go's current type system.
	if respWithErr, ok := response.(net.ResponseWithError); ok {
		if errCode := respWithErr.ErrorCode(); errCode != "" {
			return sdkErrors.FromCode(errCode)
		}
	}

	return nil
}
