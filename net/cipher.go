package net

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
	"net/http"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
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
