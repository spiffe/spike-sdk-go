package net

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
	"net/http"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// DecryptDataStreaming performs decryption for streaming mode requests.
//
// Parameters:
//   - nonce: The nonce bytes
//   - ciphertext: The encrypted data
//   - c: The cipher to use for decryption
//   - w: The HTTP response writer for error responses
//
// Returns:
//   - plaintext: The decrypted data if successful
//   - *sdkErrors.SDKError: An error if decryption fails
func DecryptDataStreaming(
	nonce, ciphertext []byte, c cipher.AEAD, w http.ResponseWriter,
) ([]byte, *sdkErrors.SDKError) {
	plaintext, err := c.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		http.Error(w, "decryption failed", http.StatusBadRequest)
		return nil, sdkErrors.ErrCryptoDecryptionFailed.Wrap(err)
	}

	return plaintext, nil
}

// DecryptDataJSON performs decryption for JSON mode requests.
//
// Parameters:
//   - nonce: The nonce bytes
//   - ciphertext: The encrypted data
//   - c: The cipher to use for decryption
//   - w: The HTTP response writer for error responses
//
// Returns:
//   - plaintext: The decrypted data if successful
//   - *sdkErrors.SDKError: An error if decryption fails
func DecryptDataJSON(
	nonce, ciphertext []byte, c cipher.AEAD, w http.ResponseWriter,
) ([]byte, *sdkErrors.SDKError) {
	plaintext, err := c.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		failErr := Fail(
			reqres.CipherDecryptResponse{}.Internal(), w,
			http.StatusInternalServerError,
		)
		if failErr != nil {
			return nil, sdkErrors.ErrCryptoDecryptionFailed.Wrap(err).Wrap(failErr)
		}
		return nil, sdkErrors.ErrCryptoDecryptionFailed.Wrap(err)
	}

	return plaintext, nil
}

// GenerateNonceOrFailStreaming generates a cryptographically secure random
// nonce for streaming mode requests.
//
// Parameters:
//   - c: The cipher to determine nonce size
//   - w: The HTTP response writer for error responses
//
// Returns:
//   - nonce: The generated nonce bytes if successful
//   - *sdkErrors.SDKError: An error if nonce generation fails
func GenerateNonceOrFailStreaming(
	c cipher.AEAD, w http.ResponseWriter,
) ([]byte, *sdkErrors.SDKError) {
	nonce := make([]byte, c.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		http.Error(
			w, string(sdkErrors.ErrCryptoNonceGenerationFailed.Code),
			http.StatusInternalServerError,
		)
		return nil, sdkErrors.ErrCryptoNonceGenerationFailed.Wrap(err)
	}

	return nonce, nil
}

// GenerateNonceOrFailJSON generates a cryptographically secure random nonce
// for JSON mode requests.
//
// Parameters:
//   - c: The cipher to determine nonce size
//   - w: The HTTP response writer for error responses
//   - errorResponse: The error response to send on failure
//
// Returns:
//   - nonce: The generated nonce bytes if successful
//   - *sdkErrors.SDKError: An error if nonce generation fails
func GenerateNonceOrFailJSON[T any](
	c cipher.AEAD, w http.ResponseWriter, errorResponse T,
) ([]byte, *sdkErrors.SDKError) {
	nonce := make([]byte, c.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		failErr := Fail(errorResponse, w, http.StatusInternalServerError)
		if failErr != nil {
			return nil, sdkErrors.ErrCryptoNonceGenerationFailed.Wrap(
				err).Wrap(failErr)
		}
		return nil, sdkErrors.ErrCryptoNonceGenerationFailed.Wrap(err)
	}

	return nonce, nil
}

// EncryptDataStreaming generates a nonce, performs encryption, and returns
// the nonce and ciphertext for streaming mode requests.
//
// Parameters:
//   - plaintext: The data to encrypt
//   - c: The cipher to use for encryption
//   - w: The HTTP response writer for error responses
//
// Returns:
//   - nonce: The generated nonce bytes
//   - ciphertext: The encrypted data
//   - *sdkErrors.SDKError: An error if nonce generation fails
func EncryptDataStreaming(
	plaintext []byte, c cipher.AEAD, w http.ResponseWriter,
) ([]byte, []byte, *sdkErrors.SDKError) {
	nonce, err := GenerateNonceOrFailStreaming(c, w)
	if err != nil {
		return nil, nil, err
	}

	ciphertext := c.Seal(nil, nonce, plaintext, nil)
	return nonce, ciphertext, nil
}

// EncryptDataJSON generates a nonce, performs encryption, and returns the
// nonce and ciphertext for JSON mode requests.
//
// Parameters:
//   - plaintext: The data to encrypt
//   - c: The cipher to use for encryption
//   - w: The HTTP response writer for error responses
//
// Returns:
//   - nonce: The generated nonce bytes
//   - ciphertext: The encrypted data
//   - *sdkErrors.SDKError: An error if nonce generation fails
func EncryptDataJSON(
	plaintext []byte, c cipher.AEAD, w http.ResponseWriter,
) ([]byte, []byte, *sdkErrors.SDKError) {
	nonce, err := GenerateNonceOrFailJSON(
		c, w, reqres.CipherEncryptResponse{}.Internal(),
	)
	if err != nil {
		return nil, nil, err
	}

	ciphertext := c.Seal(nil, nonce, plaintext, nil)
	return nonce, ciphertext, nil
}

// ReadStreamingDecryptRequestData reads the binary data from a streaming mode
// decryption request (version, nonce, ciphertext).
//
// This function does NOT perform authentication - the caller must have already
// called the guard function.
//
// The streaming format is: version byte + nonce and ciphertext
//
// Parameters:
//   - w: The HTTP response writer for error responses
//   - r: The HTTP request containing the binary data
//   - c: The cipher to determine nonce size
//
// Returns:
//   - version: The protocol version byte
//   - nonce: The nonce bytes
//   - ciphertext: The encrypted data
//   - *sdkErrors.SDKError: An error if reading fails
func ReadStreamingDecryptRequestData(
	w http.ResponseWriter, r *http.Request, c cipher.AEAD,
) (byte, []byte, []byte, *sdkErrors.SDKError) {
	const fName = "readStreamingDecryptRequestData"

	// Read the version byte
	ver := make([]byte, 1)
	n, err := io.ReadFull(r.Body, ver)
	if err != nil || n != 1 {
		failErr := sdkErrors.ErrCryptoFailedToReadVersion.Clone()
		log.WarnErr(fName, *failErr)
		http.Error(
			w, string(failErr.Code), http.StatusBadRequest,
		)
		return 0, nil, nil, failErr
	}

	version := ver[0]

	// Validate version matches the expected value
	if version != spikeCipherVersion {
		failErr := sdkErrors.ErrCryptoUnsupportedCipherVersion.Clone()
		log.WarnErr(fName, *failErr)
		http.Error(
			w, string(failErr.Code), http.StatusBadRequest,
		)
		return 0, nil, nil, failErr
	}

	// Read the nonce
	bytesToRead := c.NonceSize()
	nonce := make([]byte, bytesToRead)
	n, err = io.ReadFull(r.Body, nonce)
	if err != nil || n != bytesToRead {
		failErr := sdkErrors.ErrCryptoFailedToReadNonce.Clone()
		log.WarnErr(fName, *failErr)
		http.Error(
			w, string(failErr.Code), http.StatusBadRequest,
		)
		return 0, nil, nil, failErr
	}

	// Read the remaining body as ciphertext
	ciphertext, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		failErr := sdkErrors.ErrDataReadFailure.Wrap(readErr)
		failErr.Msg = "failed to read ciphertext"
		log.WarnErr(fName, *failErr)
		http.Error(
			w, string(failErr.Code), http.StatusBadRequest,
		)
		return 0, nil, nil, failErr
	}

	return version, nonce, ciphertext, nil
}

// ReadJSONDecryptRequestWithoutGuard reads and parses a JSON mode decryption
// request without performing guard validation.
//
// Parameters:
//   - w: The HTTP response writer for error responses
//   - r: The HTTP request containing the JSON data
//
// Returns:
//   - *reqres.CipherDecryptRequest: The parsed request
//   - *sdkErrors.SDKError: An error if reading or parsing fails
func ReadJSONDecryptRequestWithoutGuard(
	w http.ResponseWriter, r *http.Request,
) (*reqres.CipherDecryptRequest, *sdkErrors.SDKError) {
	requestBody, err := ReadRequestBodyAndRespondOnFail(w, r)
	if err != nil {
		return nil, err
	}

	request, unmarshalErr := UnmarshalAndRespondOnFail[
		reqres.CipherDecryptRequest, reqres.CipherDecryptResponse](
		requestBody, w,
		reqres.CipherDecryptResponse{}.BadRequest(),
	)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return request, nil
}

// ReadStreamingEncryptRequestWithoutGuard reads a streaming mode encryption
// request without performing guard validation. The raw binary body is wrapped
// in a CipherEncryptRequest to provide a unified interface with the JSON
// reader.
//
// Parameters:
//   - w: The HTTP response writer for error responses
//   - r: The HTTP request containing the binary data
//
// Returns:
//   - *reqres.CipherEncryptRequest: The request with plaintext populated
//   - *sdkErrors.SDKError: An error if reading fails
func ReadStreamingEncryptRequestWithoutGuard(
	w http.ResponseWriter, r *http.Request,
) (*reqres.CipherEncryptRequest, *sdkErrors.SDKError) {
	plaintext, err := ReadRequestBodyAndRespondOnFail(w, r)
	if err != nil {
		return nil, err
	}

	return &reqres.CipherEncryptRequest{Plaintext: plaintext}, nil
}

// ReadJSONEncryptRequestWithoutGuard reads and parses a JSON mode encryption
// request without performing guard validation.
//
// Parameters:
//   - w: The HTTP response writer for error responses
//   - r: The HTTP request containing the JSON data
//
// Returns:
//   - *reqres.CipherEncryptRequest: The parsed request
//   - *sdkErrors.SDKError: An error if reading or parsing fails
func ReadJSONEncryptRequestWithoutGuard(
	w http.ResponseWriter, r *http.Request,
) (*reqres.CipherEncryptRequest, *sdkErrors.SDKError) {
	requestBody, err := ReadRequestBodyAndRespondOnFail(w, r)
	if err != nil {
		return nil, err
	}

	request, unmarshalErr := UnmarshalAndRespondOnFail[
		reqres.CipherEncryptRequest, reqres.CipherEncryptResponse](
		requestBody, w,
		reqres.CipherEncryptResponse{}.BadRequest(),
	)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return request, nil
}
