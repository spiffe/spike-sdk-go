package api

import (
	"context"

	"github.com/cloudflare/circl/secretsharing"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/bootstrap"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Contribute sends a secret share contribution to a SPIKE Keeper during the
// bootstrap process.
//
// It establishes a mutual TLS connection to the specified Keeper and transmits
// the keeper's share of the secret. The function marshals the share value,
// validates its length, and sends it securely to the Keeper. After sending, the
// contribution is zeroed out in memory for security.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - keeperShare: The secret share to contribute to the Keeper
//   - keeperID: The unique identifier of the target Keeper
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - Errors from net.Post(): if the HTTP request fails
//
// Note: The function will fatally crash (via log.FatalErr) if:
//   - Marshal failures (ErrDataMarshalFailure)
//   - Share length validation fails (ErrCryptoInvalidEncryptionKeyLength)
func (a *API) Contribute(
	ctx context.Context, keeperShare secretsharing.Share, keeperID string,
) *sdkErrors.SDKError {
	return bootstrap.Contribute(ctx, a.source, keeperShare, keeperID)
}

// Verify performs bootstrap verification with SPIKE Nexus by sending encrypted
// random text and validating that Nexus can decrypt it correctly.
//
// This ensures that the bootstrap process completed successfully and Nexus has
// the correct master key. The function sends the nonce and ciphertext to Nexus,
// receives back a hash, and compares it against the expected hash of the
// original random text. A match confirms successful bootstrap.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - randomText: The original random text that was encrypted
//   - nonce: The nonce used during encryption
//   - cipherText: The encrypted random text
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - Errors from net.Post(): if the HTTP request fails
//
// Note: The function will fatally crash (via log.FatalErr) if:
//   - Marshal failures (ErrDataMarshalFailure)
//   - Response parsing failures (ErrDataUnmarshalFailure)
//   - Hash verification fails (ErrCryptoCipherVerificationFailed)
func (a *API) Verify(
	ctx context.Context, randomText string, nonce, cipherText []byte,
) *sdkErrors.SDKError {
	return bootstrap.Verify(ctx, a.source, randomText, nonce, cipherText)
}
