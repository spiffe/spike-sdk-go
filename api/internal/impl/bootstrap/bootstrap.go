//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package bootstrap

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/circl/secretsharing"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/config/env"
	"github.com/spiffe/spike-sdk-go/crypto"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/net"
	"github.com/spiffe/spike-sdk-go/security/mem"
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
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Keeper
//   - keeperShare: The secret share to contribute to the Keeper
//   - keeperID: The unique identifier of the target Keeper
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - Errors from net.Post(): if the HTTP request fails (e.g., ErrAPINotFound,
//     ErrAccessUnauthorized, ErrAPIBadRequest, ErrStateNotReady,
//     ErrNetPeerConnection)
//
// Note: The function will fatally crash (via log.FatalErr) for unrecoverable
// errors such as marshal failures (ErrDataMarshalFailure) or invalid
// contribution length (ErrCryptoInvalidEncryptionKeyLength).
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	err = Contribute(ctx, source, keeperShare, "keeper-1")
//	if err != nil {
//	    log.Printf("Failed to contribute share: %v", err)
//	}
func Contribute(
	ctx context.Context,
	source *workloadapi.X509Source,
	keeperShare secretsharing.Share,
	keeperID string,
) *sdkErrors.SDKError {
	const fName = "Contribute"

	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	contribution, err := keeperShare.Value.MarshalBinary()
	if err != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(err)
		failErr.Msg = "failed to marshal share"
		log.FatalErr(fName, *failErr)
	}

	if len(contribution) != crypto.AES256KeySize {
		failErr := sdkErrors.ErrCryptoInvalidEncryptionKeyLength.Clone()
		failErr.Msg = fmt.Sprintf(
			"invalid contribution length: expected %d, got %d",
			crypto.AES256KeySize, len(contribution),
		)
		log.FatalErr(fName, *failErr)
	}

	scr := reqres.ShardPutRequest{}
	shard := new([crypto.AES256KeySize]byte)
	copy(shard[:], contribution)

	// Security: Zero out contribution as soon as we don't need it.
	mem.ClearBytes(contribution)

	scr.Shard = shard

	client := net.CreateMTLSClientForKeeper(source)

	for kid, keeperAPIRoot := range env.KeepersVal() {
		if kid != keeperID {
			// These are not the keepers we are looking for...
			continue
		}

		md, marshalErr := json.Marshal(scr)
		if marshalErr != nil {
			failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
			failErr.Msg = "failed to marshal request"
			log.FatalErr(fName, *failErr)
		}

		u := url.KeeperBootstrapContributeEndpoint(keeperAPIRoot)

		_, sdkErr := net.Post(ctx, client, u, md)
		if sdkErr != nil {
			return sdkErr
		}
	}

	return nil
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
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - randomText: The original random text that was encrypted
//   - nonce: The nonce used during encryption
//   - ciphertext: The encrypted random text
//
// Returns:
//   - *sdkErrors.SDKError: nil on success (hash matches), or one of the following errors:
//   - ErrSPIFFENilX509Source: if source is nil
//   - Errors from net.Post(): if the HTTP request fails (e.g., ErrAPINotFound,
//     ErrAccessUnauthorized, ErrAPIBadRequest, ErrStateNotReady, ErrNetPeerConnection)
//
// Note: The function will fatally crash (via log.FatalErr) for unrecoverable
// errors such as marshal failures (ErrDataMarshalFailure), response parsing
// failures (ErrDataUnmarshalFailure), or hash verification failures
// (ErrCryptoCipherVerificationFailed). These indicate potential security
// issues and the application should not continue.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	err = Verify(ctx, source, randomText, nonce, ciphertext)
//	if err != nil {
//	    log.Printf("Bootstrap verification failed: %v", err)
//	}
func Verify(
	ctx context.Context,
	source *workloadapi.X509Source,
	randomText string,
	nonce, ciphertext []byte,
) *sdkErrors.SDKError {
	const fName = "Verify"

	if source == nil {
		return sdkErrors.ErrSPIFFENilX509Source.Clone()
	}

	client := net.CreateMTLSClientForNexus(source)

	request := reqres.BootstrapVerifyRequest{
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}

	md, marshalErr := json.Marshal(request)
	if marshalErr != nil {
		failErr := sdkErrors.ErrDataMarshalFailure.Wrap(marshalErr)
		failErr.Msg = "failed to marshal verification request"
		log.FatalErr(fName, *failErr)
	}

	// Send the verification request to SPIKE Nexus
	nexusAPIRoot := env.NexusAPIRootVal()
	verifyURL := url.NexusBootstrapVerifyEndpoint(nexusAPIRoot)

	log.Info(
		fName,
		"message", "sending verification request to SPIKE Nexus",
		"url", verifyURL,
	)

	responseBody, err := net.Post(ctx, client, verifyURL, md)
	if err != nil {
		return err
	}

	// Parse the response
	var verifyResponse struct {
		Hash string `json:"hash"`
		Err  string `json:"err"`
	}
	if unmarshalErr := json.Unmarshal(
		responseBody, &verifyResponse,
	); unmarshalErr != nil {
		failErr := sdkErrors.ErrDataUnmarshalFailure.Wrap(unmarshalErr)
		failErr.Msg = "failed to parse verification response"
		// If SPIKE Keeper is sending gibberish, it may be a malicious actor.
		// Fatally crash here to prevent a possible compromise.
		log.FatalErr(fName, *failErr)
	}

	// Compute the expected hash
	expectedHash := sha256.Sum256([]byte(randomText))
	expectedHashHex := hex.EncodeToString(expectedHash[:])

	// Verify the hash matches
	if verifyResponse.Hash != expectedHashHex {
		failErr := sdkErrors.ErrCryptoCipherVerificationFailed.Clone()
		failErr.Msg = "verification failed: hash mismatch"
		log.FatalErr(fName, *failErr)
	}

	return nil
}
