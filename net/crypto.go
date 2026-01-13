package net

import (
	"net/http"

	"github.com/spiffe/spike-sdk-go/config/env"
	"github.com/spiffe/spike-sdk-go/crypto"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

const spikeCipherVersion = 0x01

// RespondCryptoErrOnVersionMismatch validates that the protocol version
// is supported.
//
// Parameters:
//   - version: The protocol version byte to validate
//   - w: The HTTP response writer for error responses
//   - errorResponse: The error response to send on failure
//
// Returns:
//   - nil if the version is valid
//   - *sdkErrors.SDKError if the version is unsupported
func RespondCryptoErrOnVersionMismatch[T any](
	version byte, w http.ResponseWriter, errorResponse T,
) *sdkErrors.SDKError {
	if version != spikeCipherVersion {
		failErr := Fail(errorResponse, w, http.StatusBadRequest)
		if failErr != nil {
			return sdkErrors.ErrCryptoUnsupportedCipherVersion.Wrap(failErr)
		}
		return sdkErrors.ErrCryptoUnsupportedCipherVersion.Clone()
	}
	return nil
}

// RespondCryptoErrOnInvalidNonceSize validates that the nonce is exactly the
// expected size.
//
// Parameters:
//   - nonce: The nonce bytes to validate
//   - w: The HTTP response writer for error responses
//   - errorResponse: The error response to send on failure
//
// Returns:
//   - nil if the nonce size is valid
//   - *sdkErrors.SDKError if the nonce size is invalid
func RespondCryptoErrOnInvalidNonceSize[T any](
	nonce []byte, w http.ResponseWriter, errorResponse T,
) *sdkErrors.SDKError {
	if len(nonce) != crypto.GCMNonceSize {
		failErr := Fail(errorResponse, w, http.StatusBadRequest)
		if failErr != nil {
			return sdkErrors.ErrDataInvalidInput.Wrap(failErr)
		}
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	return nil
}

// RespondCryptoErrOnLargeCipherText validates that the ciphertext does not
// exceed the maximum allowed size.
//
// Parameters:
//   - ciphertext: The ciphertext bytes to validate
//   - w: The HTTP response writer for error responses
//   - errorResponse: The error response to send on failure
//
// Returns:
//   - nil if the ciphertext size is valid
//   - *sdkErrors.SDKError if the ciphertext is too large
func RespondCryptoErrOnLargeCipherText[T any](
	ciphertext []byte, w http.ResponseWriter, errorResponse T,
) *sdkErrors.SDKError {
	if len(ciphertext) > env.CryptoMaxCiphertextSizeVal() {
		failErr := Fail(errorResponse, w, http.StatusBadRequest)
		if failErr != nil {
			return sdkErrors.ErrDataInvalidInput.Wrap(failErr)
		}
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	return nil
}

// RespondCryptoErrOnLargePlaintext validates that the plaintext does not
// exceed the maximum allowed size.
//
// Parameters:
//   - plaintext: The plaintext bytes to validate
//   - w: The HTTP response writer for error responses
//   - errorResponse: The error response to send on failure
//
// Returns:
//   - nil if the plaintext size is valid
//   - *sdkErrors.SDKError if the plaintext is too large
func RespondCryptoErrOnLargePlaintext[T any](
	plaintext []byte, w http.ResponseWriter, errorResponse T,
) *sdkErrors.SDKError {
	if len(plaintext) > env.CryptoMaxPlaintextSizeVal() {
		failErr := Fail(errorResponse, w, http.StatusBadRequest)
		if failErr != nil {
			return sdkErrors.ErrDataInvalidInput.Wrap(failErr)
		}
		return sdkErrors.ErrDataInvalidInput
	}
	return nil
}
