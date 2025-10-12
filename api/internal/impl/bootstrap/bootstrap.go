package bootstrap

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudflare/circl/secretsharing"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/config/env"
	"github.com/spiffe/spike-sdk-go/crypto"
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
//   - source: X509Source for establishing mTLS connection to SPIKE Keeper
//   - keeperShare: The secret share to contribute to the Keeper
//   - keeperID: The unique identifier of the target Keeper
//
// Returns:
//   - nil if successful
//   - error if marshaling, validation, network request, or server-side
//     processing fails
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	err = Contribute(source, keeperShare, "keeper-1")
//	if err != nil {
//	    log.Printf("Failed to contribute share: %v", err)
//	}
func Contribute(
	source *workloadapi.X509Source,
	keeperShare secretsharing.Share,
	keeperID string,
) error {
	const fName = "Contribute"

	if source == nil {
		return errors.New("nil X509Source")
	}

	contribution, err := keeperShare.Value.MarshalBinary()
	if err != nil {
		log.FatalLn(fName, "message", "Failed to marshal share",
			"err", err, "keeper_id", keeperID)
	}

	if len(contribution) != crypto.AES256KeySize {
		log.FatalLn(fName,
			"message", "invalid contribution length",
			"len", len(contribution), "keeper_id", keeperID)
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

		md, err := json.Marshal(scr)
		if err != nil {
			log.FatalLn(fName,
				"message", "Failed to marshal request",
				"err", err, "keeper_id", keeperID)
		}

		log.Log().Info(fName, "payload", fmt.Sprintf("%x", sha256.Sum256(md)))

		u := url.KeeperBootstrapContributeEndpoint(keeperAPIRoot)

		_, err = net.Post(client, u, md)
		if err != nil {
			log.Log().Info(fName, "message",
				"Failed to post",
				"err", err, "keeper_id", keeperID)
			return err
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
//   - source: X509Source for establishing mTLS connection to SPIKE Nexus
//   - randomText: The original random text that was encrypted
//   - nonce: The nonce used during encryption
//   - ciphertext: The encrypted random text
//
// Returns:
//   - nil if verification succeeds (hash matches)
//   - error if marshaling, network request, parsing, or hash verification fails
//
// Example:
//
//	source, err := workloadapi.NewX509Source(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer source.Close()
//
//	err = Verify(source, randomText, nonce, ciphertext)
//	if err != nil {
//	    log.Printf("Bootstrap verification failed: %v", err)
//	}
func Verify(
	source *workloadapi.X509Source,
	randomText string,
	nonce, ciphertext []byte,
) error {
	const fName = "Verify"

	if source == nil {
		return errors.New("nil X509Source")
	}

	client := net.CreateMTLSClientForNexus(source)

	request := reqres.BootstrapVerifyRequest{
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}

	md, err := json.Marshal(request)
	if err != nil {
		log.FatalLn(fName,
			"message", "Failed to marshal verification request",
			"err", err)
	}

	log.Log().Info(fName, "payload", fmt.Sprintf("%x", sha256.Sum256(md)))

	// Send the verification request to SPIKE Nexus
	nexusAPIRoot := env.NexusAPIRootVal()
	verifyURL := url.NexusVerifyEndpoint(nexusAPIRoot)

	log.Log().Info(fName, "message",
		"Sending verification request to SPIKE Nexus", "url", verifyURL)

	responseBody, err := net.Post(client, verifyURL, md)
	if err != nil {
		log.Log().Error(fName, "message",
			"Failed to post verification request",
			"err", err)
		return err
	}

	// Parse the response
	var verifyResponse struct {
		Hash string `json:"hash"`
		Err  string `json:"err"`
	}
	if err := json.Unmarshal(responseBody, &verifyResponse); err != nil {
		log.Log().Error(fName, "message",
			"Failed to parse verification response", "err", err.Error())
		return err
	}

	// Compute the expected hash
	expectedHash := sha256.Sum256([]byte(randomText))
	expectedHashHex := hex.EncodeToString(expectedHash[:])

	// Verify the hash matches
	if verifyResponse.Hash != expectedHashHex {
		log.FatalLn(fName, "message",
			"Verification failed: hash mismatch",
			"expected", expectedHashHex,
			"received", verifyResponse.Hash)
		return errors.New("verification failed")
	}

	return nil
}
