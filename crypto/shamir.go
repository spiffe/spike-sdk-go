//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"github.com/cloudflare/circl/group"
	shamir "github.com/cloudflare/circl/secretsharing"
	"github.com/spiffe/spike-sdk-go/security/mem"

	"github.com/spiffe/spike-sdk-go/config/env"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// VerifyShamirReconstruction verifies that a set of secret shares can
// correctly reconstruct the original secret. It performs this verification by
// attempting to recover the secret using the minimum required number of shares
// and comparing the result with the original secret.
//
// This function is intended for validating newly generated shares, not for
// restore operations. During a restore, the original secret is unknown, and
// successful reconstruction via secretsharing.Recover() is itself proof that
// the shards are mathematically valid.
//
// Parameters:
//   - secret group.Scalar: The original secret to verify against.
//   - shares []shamir.Share: The generated secret shares to verify.
//
// The function will:
//   - Calculate the threshold (t) from the environment configuration.
//   - Attempt to reconstruct the secret using exactly t+1 shares.
//   - Compare the reconstructed secret with the original.
//   - Zero out the reconstructed secret regardless of success or failure.
//
// If the verification fails, the function will:
//   - Log a fatal error and exit if recovery fails.
//   - Log a fatal error and exit if the recovered secret does not match the
//     original.
//
// Security:
//   - The reconstructed secret is always zeroed out to prevent memory leaks.
//   - In case of fatal errors, the reconstructed secret is explicitly zeroed
//     before logging since deferred functions will not run after log.FatalErr.
func VerifyShamirReconstruction(secret group.Scalar, shares []shamir.Share) {
	const fName = "VerifyShamirReconstruction"

	thresholdVal := env.ShamirThresholdVal()
	if thresholdVal < 1 {
		failErr := *sdkErrors.ErrShamirReconstructionFailed.Clone()
		failErr.Msg = "shamir threshold must be at least 1"
		log.FatalErr(fName, failErr)
	}
	// #nosec G115 -- thresholdVal is validated >= 1 above, so thresholdVal-1 >= 0
	t := uint(thresholdVal - 1) // Need t+1 shares to reconstruct

	reconstructed, err := shamir.Recover(t, shares[:thresholdVal])
	// Security: Ensure that the secret is zeroed out if the check fails.
	defer func() {
		if reconstructed == nil {
			return
		}
		reconstructed.SetUint64(0)
	}()

	if err != nil {
		// deferred will not run in a fatal crash.
		reconstructed.SetUint64(0)

		failErr := sdkErrors.ErrShamirReconstructionFailed.Wrap(err)
		failErr.Msg = "failed to recover root key"
		log.FatalErr(fName, *failErr)
	}
	if !secret.IsEqual(reconstructed) {
		// deferred will not run in a fatal crash.
		reconstructed.SetUint64(0)

		failErr := *sdkErrors.ErrShamirReconstructionFailed.Clone()
		failErr.Msg = "recovered secret does not match original"
		log.FatalErr(fName, failErr)
	}
}

// ComputeShares generates a set of Shamir secret shares from the root key.
// The function uses a deterministic random reader seeded with the root key,
// which ensures that the same shares are always generated for a given root key.
// This deterministic behavior is crucial for the system's reliability, allowing
// shares to be recomputed as needed while maintaining consistency.
//
// Parameters:
//   - rk *[32]byte: The root key used to generate the secret shares. Since this
//     is a pointer type and the root key is typically a global variable in
//     SPIKE Nexus, the caller must acquire a mutex lock before calling this
//     function and release it afterward to ensure thread safety.
//
// Returns:
//   - group.Scalar: The root secret as a P256 scalar (caller must zero after
//     use)
//   - []shamir.Share: The computed shares with monotonically increasing IDs
//     starting from 1 (caller must zero after use)
//
// The function will log a fatal error and exit if:
//   - The root key is nil or zeroed
//   - The root key fails to unmarshal into a scalar
//   - The generated shares fail reconstruction verification
func ComputeShares(rk *[32]byte) (group.Scalar, []shamir.Share) {
	const fName = "ComputeShares"

	if rk == nil || mem.Zeroed32(rk) {
		failErr := sdkErrors.ErrRootKeyEmpty.Clone()
		log.FatalErr(fName, *failErr)
	}

	g := group.P256
	t := uint(env.ShamirThresholdVal() - 1) // Need t+1 shares to reconstruct
	n := uint(env.ShamirSharesVal())        // Total number of shares

	rootSecret := g.NewScalar()
	if err := rootSecret.UnmarshalBinary(rk[:]); err != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(err)
		log.FatalErr(fName, *failErr)
	}

	// Using the root key as the seed is secure because Shamir Secret Sharing
	// security does not depend on the random seed; it depends on the shards
	// being kept secret. Using a deterministic reader ensures identical shares
	// are generated for the same root key, which simplifies synchronization
	// after Nexus restarts.
	reader := NewDeterministicReader(rk[:])
	ss := shamir.New(reader, t, rootSecret)
	shares := ss.Share(n)

	// Verify the generated shares can reconstruct the original secret.
	// This crashes via log.FatalErr if reconstruction fails.
	VerifyShamirReconstruction(rootSecret, shares)

	return rootSecret, shares
}
