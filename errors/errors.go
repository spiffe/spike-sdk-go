//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

//
// General error codes
//

var ErrBadRequest = New(errCodeBadRequest, "bad request", nil)
var ErrEmptyPayload = New(errCodeEmptyPayload, "empty payload", nil)
var ErrFound = New(errCodeFound, "found", nil)
var ErrInternal = New(errCodeInternal, "internal error", nil)
var ErrNilContext = New(errCodeNilContext, "nil context", nil)
var ErrServerFault = New(errCodeServerFault, "server fault", nil)
var ErrSuccess = New(errCodeSuccess, "success", nil)
var ErrNotFound = New(errCodeNotFound, "not found", nil)
var ErrPostFailed = New(errCodePostFailed, "post failed", nil)
var ErrGeneralFailure = New(errCodeGeneralFailure, "general failure", nil)

//
// Cluster operations
//

var ErrK8sReconciliationFailed = New(errCodeClusterReconciliationFailed, "reconciliation failed", nil)

//
// Entity operations
//

var ErrEntityExists = New(errCodeEntityExists, "entity already exists", nil)
var ErrEntityInvalid = New(errCodeEntityInvalid, "entity is invalid", nil)
var ErrEntityNotFound = New(errCodeEntityNotFound, "entity not found", nil)

//
// State operations
//

var ErrAlreadyInitialized = New(errCodeAlreadyInitialized, "already initialized", nil)
var ErrInitializationFailed = New(errCodeInitializationFailed, "initialization failed", nil)
var ErrNotAlive = New(errCodeNotAlive, "not alive", nil)
var ErrNotReady = New(errCodeNotReady, "not ready", nil)

//
// Policy/RBAC/ABAC
//

var ErrUnauthorized = New(errCodeUnauthorized, "unauthorized", nil)

//
// CRUD operations
//

var ErrDeletionFailed = New(errCodeDeletionFailed, "deletion failed", nil)
var ErrDeletionSuccess = New(errCodeDeletionSuccess, "deletion success", nil)
var ErrUndeleteFailed = New(errCodeUndeleteFailed, "undeletion failed", nil)
var ErrUndeleteSuccess = New(errCodeUndeleteSuccess, "undeletion success", nil)
var ErrCreationFailed = New(errCodeCreationFailed, "creation failed", nil)

//
// Root key management
//

var ErrRootKeyEmpty = New(errCodeRootKeyEmpty, "root key empty", nil)
var ErrRootKeyMissing = New(errCodeRootKeyMissing, "root key missing", nil)
var ErrRootKeyNotEmpty = New(errCodeRootKeyNotEmpty, "root key not empty", nil)
var ErrRootKeySetSuccess = New(errCodeRootKeySetSuccess, "root key set success", nil)
var ErrRootKeySkipCreationForInMemoryMode = New(errCodeRootKeySkipCreationForInMemoryMode, "root key skip creation for in memory mode", nil)
var ErrRootKeyUpdateSkippedKeyEmpty = New(errCodeRootKeyUpdateSkippedKeyEmpty, "root key update skipped key empty", nil)

//
// Shamir-related
//

var ErrShamirDuplicateIndex = New(errCodeShamirDuplicateIndex, "shamir duplicate index", nil)
var ErrShamirEmptyShard = New(errCodeShamirEmptyShard, "shamir empty shard", nil)
var ErrShamirInvalidIndex = New(errCodeShamirInvalidIndex, "shamir invalid index", nil)
var ErrShamirNilShard = New(errCodeShamirNilShard, "shamir nil shard", nil)
var ErrShamirNotEnoughShards = New(errCodeShamirNotEnoughShards, "shamir not enough shards", nil)
var ErrShamirReconstructionFailed = New(errCodeShamirReconstructionFailed, "shamir reconstruction failed", nil)

//
// Crypto operations
//

var ErrLowEntropy = New(errCodeLowEntropy, "low entropy", nil)
var ErrCryptoCipherNotAvailable = New(errCodeCryptoCipherNotAvailable, "cipher not available", nil)
var ErrCryptoCipherVerificationSuccess = New(errCodeCryptoCipherVerificationSuccess, "cipher verification success", nil)
var ErrCryptoDecryptionFailed = New(errCodeCryptoDecryptionFailed, "decryption failed", nil)
var ErrCryptoEncryptionFailed = New(errCodeCryptoEncryptionFailed, "encryption failed", nil)
var ErrCryptoFailedToReadNonce = New(errCodeCryptoFailedToReadNonce, "failed to read nonce", nil)
var ErrCryptoFailedToReadVersion = New(errCodeCryptoFailedToReadVersion, "failed to read version", nil)
var ErrCryptoNonceGenerationFailed = New(errCodeCryptoNonceGenerationFailed, "nonce generation failed", nil)
var ErrCryptoFailedToCreateCipher = New(errCodeCryptoFailedToCreateCipher, "failed to create cipher", nil)
var ErrCryptoFailedToCreateGCM = New(errCodeCryptoFailedToCreateGCM, "failed to create GCM", nil)
var ErrCryptoInvalidEncryptionKeyLength = New(errCodeCryptoInvalidEncryptionKeyLength, "invalid encryption key length", nil)

//
// Store configuration
//

var ErrStoreInvalidConfiguration = New(errCodeStoreInvalidConfiguration, "invalid store configuration", nil)
var ErrStoreInvalidEncryptionKey = New(errCodeStoreInvalidEncryptionKey, "invalid store encryption key", nil)

//
// Store operations
//

var ErrStoreLoadFailed = New(errCodeStoreLoadFailed, "failed to load data", nil)
var ErrStoreQueryFailed = New(errCodeStoreQueryFailed, "failed to query data", nil)
var ErrStoreSaveFailed = New(errCodeStoreSaveFailed, "failed to save data", nil)
var ErrStoreQueryFailure = New(errCodeStoreQueryFailure, "failed to query for the requested data", nil)
var ErrStoreResultSetFailedToLoad = New(errCodeStoreResultSetFailedToLoad, "result set failed to load", nil)
var ErrStoreVersionNotFound = New(errCodeStoreVersionNotFound, "version not found", nil)
var ErrStoreItemSoftDeleted = New(errCodeStoreItemSoftDeleted, "item marked as deleted", nil)
var ErrStoreInvalidVersion = New(errCodeStoreInvalidVersion, "invalid version", nil)

//
// Filesystem operations
//

var ErrDirectoryCreationFailed = New(errCodeDirectoryCreationFailed, "directory creation failed", nil)
var ErrFSFailedToCheckDirectory = New(errCodeFSFailedToCheckDirectory, "failed to check directory", nil)
var ErrFSFailedToCreateDirectory = New(errCodeFSFailedToCreateDirectory, "failed to create directory", nil)
var ErrFSFailedToResolvePath = New(errCodeFSFailedToResolvePath, "failed to resolve filesystem path", nil)
var ErrFSFileIsNotADirectory = New(errCodeFSFileIsNotADirectory, "file is not a directory", nil)
var ErrFSInvalidDirectory = New(errCodeFSInvalidDirectory, "invalid directory", nil)
var ErrFSParentDirectoryDoesNotExist = New(errCodeFSParentDirectoryDoesNotExist, "parent directory does not exist", nil)
var ErrFSPathCannotBeEmpty = New(errCodeFSPathCannotBeEmpty, "filesystem path cannot be empty", nil)
var ErrFSPathRestricted = New(errCodeFSPathRestricted, "filesystem path is restricted for security reasons", nil)

//
// File I/O
//

var ErrFileCloseFailed = New(errCodeFileCloseFailed, "file close failed", nil)
var ErrStreamCloseFailed = New(errCodeStreamCloseFailed, "stream close failed", nil)
var ErrFileOpenFailed = New(errCodeFileOpenFailed, "file open failed", nil)

//
// Input/Output
//

var ErrInvalidInput = New(errCodeInvalidInput, "invalid input", nil)
var ErrInvalidPermission = New(errCodeInvalidPermission, "invalid permission", nil)
var ErrMarshalFailure = New(errCodeMarshalFailure, "failed to marshal response body", nil)
var ErrParseFailure = New(errCodeParseFailure, "failed to parse request body", nil)
var ErrReadFailure = New(errCodeReadFailure, "failed to read request body", nil)
var ErrUnmarshalFailure = New(errCodeUnmarshalFailure, "failed to unmarshal request body", nil)

//
// Network/Peer
//

var ErrPeerConnection = New(errCodePeerConnection, "problem connecting to peer", nil)
var ErrReadingResponseBody = New(errCodeReadingResponseBody, "problem reading response body", nil)
var ErrReadingRequestBody = New(errCodeReadingRequestBody, "problem reading requestbody", nil)

//
// Transaction operations
//

var ErrTransactionBeginFailed = New(errCodeTransactionBeginFailed, "failed to begin transaction", nil)
var ErrTransactionCommitFailed = New(errCodeTransactionCommitFailed, "failed to commit transaction", nil)
var ErrTransactionFailed = New(errCodeTransactionFailed, "transaction failed", nil)
var ErrTransactionRollbackFailed = New(errCodeTransactionRollbackFailed, "failed to rollback transaction", nil)

//
// Recovery operations
//

var ErrRecoveryRetryFailed = New(errCodeRecoveryRetryFailed, "recovery retry failed", nil)

//
// X509/SPIFFE
//

var ErrNilX509Source = New(errCodeNilX509Source, "nil X509Source", nil)

// MaybeError converts an error to its string representation if the error is
// not nil. If the error is nil, it returns an empty string.
//
// Parameters:
//   - err: the error to convert to a string
//
// Returns:
//   - string: the error message if err is non-nil, empty string otherwise
func MaybeError(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
