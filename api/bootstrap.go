package api

import (
	"github.com/cloudflare/circl/secretsharing"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/bootstrap"
)

// Contribute sends a secret share contribution to a SPIKE Keeper during the
// bootstrap process. It establishes a mutual TLS connection to the specified
// Keeper and transmits the keeper's share of the secret.
//
// The function marshals the share value, validates its length, and sends it
// securely to the Keeper. After sending, the contribution is zeroed out in
// memory for security.
//
// Parameters:
//   - keeperShare: The secret share to contribute to the Keeper
//   - keeperID: The unique identifier of the target Keeper
//
// Returns:
//   - nil if successful
//   - error if marshaling, validation, network request, or server-side
//     processing fails
func (a *API) Contribute(
	keeperShare secretsharing.Share, keeperID string,
) error {
	return bootstrap.Contribute(a.source, keeperShare, keeperID)
}

// Verify performs bootstrap verification with SPIKE Nexus by sending encrypted
// random text and validating that Nexus can decrypt it correctly. This ensures
// that the bootstrap process completed successfully and Nexus has the correct
// master key.
//
// The function sends the nonce and ciphertext to Nexus, receives back a hash,
// and compares it against the expected hash of the original random text. A
// match confirms successful bootstrap.
//
// Parameters:
//   - randomText: The original random text that was encrypted
//   - nonce: The nonce used during encryption
//   - cipherText: The encrypted random text
//
// Returns:
//   - nil if verification succeeds (hash matches)
//   - error if marshaling, network request, parsing, or hash verification fails
func (a *API) Verify(
	randomText string, nonce, cipherText []byte,
) error {
	return bootstrap.Verify(a.source, randomText, nonce, cipherText)
}
