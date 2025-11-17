//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

import (
	"errors"
	"fmt"
)

var ErrAlreadyInitialized = errors.New("already initialized")
var ErrCipherDecryptionFailed = errors.New("decryption failed")
var ErrCipherFailedToReadNonce = errors.New("failed to read nonce")
var ErrCipherFailedToReadVersion = errors.New("failed to read version")
var ErrCipherNonceGenerationFailed = errors.New("nonce generation failed")
var ErrCipherNotAvailable = errors.New("cipher not available")
var ErrCipherVerificationSuccess = errors.New("cipher verification success")
var ErrCreationFailed = errors.New("creation failed")
var ErrDataLoadFailed = errors.New("failed to load data")
var ErrDataQueryFailed = errors.New("failed to query data")
var ErrDeletionFailed = errors.New("deletion failed")
var ErrDirectoryCreationFailed = errors.New("directory creation failed")
var ErrFileCloseFailed = errors.New("file close failed")
var ErrFileOpenFailed = errors.New("file open failed")
var ErrFound = errors.New("found")
var ErrInvalidInput = errors.New("invalid input")
var ErrInvalidPermission = errors.New("invalid permission")
var ErrMarshalFailure = errors.New("failed to marshal response body")
var ErrMissingRootKey = errors.New("missing root key")
var ErrNilX509Source = errors.New("nil X509Source")
var ErrNotFound = errors.New("not found")
var ErrParseFailure = errors.New("failed to parse request body")
var ErrPeerConnection = errors.New("problem connecting to peer")
var ErrQueryFailure = errors.New("failed to query for the requested data")
var ErrReadFailure = errors.New("failed to read request body")
var ErrReadingResponseBody = errors.New("problem reading response body")
var ErrRecoveryRetryFailed = errors.New("recovery retry failed")
var ErrTransactionBeginFailed = errors.New("failed to begin transaction")
var ErrTransactionCommitFailed = errors.New("failed to commit transaction")
var ErrTransactionFailed = errors.New("transaction failed")
var ErrTransactionRollbackFailed = errors.New("failed to rollback transaction")
var ErrUnauthorized = errors.New("unauthorized")
var ErrUndeleteFailed = errors.New("undeletion failed")
var ErrUndeleteSuccess = errors.New("undeletion success")
var ErrUnmarshalFailure = errors.New("failed to unmarshal request body")

// ErrFailedFor returns an error message indicating that an action failed
// for a specific entity.
// i.e.: "[encoding] of [path] failed for [spiffeid]: [spiffe://spike.ist/]"
func ErrFailedFor(action, whatFailed, forWhat, identifier string) error {
	return fmt.Errorf("%s of %s failed for %s: '%s'",
		action, whatFailed, forWhat, identifier,
	)
}

// ErrInvalidFor returns an error message indicating that an entity is invalid
// for a specific purpose.
// i.e.: "[encoding] is invalid for [purpose]: [spiffe://spike.ist/]"
func ErrInvalidFor(whatsInvalid, forWhat, identifier string) error {
	return fmt.Errorf("%s is invalid for %s: '%s'",
		whatsInvalid, forWhat, identifier,
	)
}
