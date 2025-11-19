//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

import "sync"

type ErrorCode string

// errorRegistry maps ErrorCodes to their corresponding SDKError instances.
// This map is automatically populated when errors are defined using the
// register() function, ensuring FromCode() always has up-to-date mappings.
// Access is protected by errorRegistryMu for thread safety.
var (
	errorRegistry   = make(map[ErrorCode]*SDKError)
	errorRegistryMu sync.RWMutex
)

// register adds an error to the global registry and returns it.
// This ensures that all defined errors are automatically available in
// FromCode(). This function is thread-safe.
//
// Parameters:
//   - err: The SDKError to register
//
// Returns:
//   - *SDKError: The same error that was passed in
func register(err *SDKError) *SDKError {
	errorRegistryMu.Lock()
	errorRegistry[err.Code] = err
	errorRegistryMu.Unlock()
	return err
}

//
// General error codes
//

const errCodeBadRequest = ErrorCode("bad_request")
const errCodeEmptyPayload = ErrorCode("empty_payload")
const errCodeFound = ErrorCode("found")
const errCodeInternal = ErrorCode("internal_error")
const errCodeNilContext = ErrorCode("nil_context")
const errCodeServerFault = ErrorCode("server_fault")
const errCodeSuccess = ErrorCode("success")
const errCodeNotFound = ErrorCode("not_found")
const errCodePostFailed = ErrorCode("post_failed")

var ErrBadRequest = register(New(errCodeBadRequest, "bad request", nil))
var ErrEmptyPayload = register(New(errCodeEmptyPayload, "empty payload", nil))
var ErrFound = register(New(errCodeFound, "found", nil))
var ErrInternal = register(New(errCodeInternal, "internal error", nil))
var ErrNilContext = register(New(errCodeNilContext, "nil context", nil))
var ErrServerFault = register(New(errCodeServerFault, "server fault", nil))
var ErrSuccess = register(New(errCodeSuccess, "success", nil))
var ErrNotFound = register(New(errCodeNotFound, "not found", nil))
var ErrPostFailed = register(New(errCodePostFailed, "post failed", nil))
var ErrGeneralFailure = register(New(errCodeGeneralFailure, "general failure", nil))

//
// Cluster operations
//

const errCodeK8sReconciliationFailed = ErrorCode("cluster_reconciliation_failed")

var ErrK8sReconciliationFailed = register(New(errCodeK8sReconciliationFailed, "reconciliation failed", nil))

//
// Entity operations
//

const errCodeEntityExists = ErrorCode("entity_exists")
const errCodeEntityInvalid = ErrorCode("entity_invalid")
const errCodeEntityNotFound = ErrorCode("entity_not_found")

var ErrEntityExists = register(New(errCodeEntityExists, "entity already exists", nil))
var ErrEntityInvalid = register(New(errCodeEntityInvalid, "entity is invalid", nil))
var ErrEntityNotFound = register(New(errCodeEntityNotFound, "entity not found", nil))

//
// State operations
//

const errCodeAlreadyInitialized = ErrorCode("already_initialized")
const errCodeInitializationFailed = ErrorCode("initialization_failed")
const errCodeNotAlive = ErrorCode("not_alive")
const errCodeNotReady = ErrorCode("not_ready")

var ErrAlreadyInitialized = register(New(errCodeAlreadyInitialized, "already initialized", nil))
var ErrInitializationFailed = register(New(errCodeInitializationFailed, "initialization failed", nil))
var ErrNotAlive = register(New(errCodeNotAlive, "not alive", nil))
var ErrNotReady = register(New(errCodeNotReady, "not ready", nil))

//
// Policy/RBAC/ABAC
//

const errCodeUnauthorized = ErrorCode("unauthorized")

var ErrUnauthorized = register(New(errCodeUnauthorized, "unauthorized", nil))

//
// CRUD operations
//

const errCodeDeletionFailed = ErrorCode("deletion_failed")
const errCodeDeletionSuccess = ErrorCode("deletion_success")
const errCodeUndeleteFailed = ErrorCode("undeletion_failed")
const errCodeUndeleteSuccess = ErrorCode("undeletion_success")
const errCodeCreationFailed = ErrorCode("creation_failed")

var ErrDeletionFailed = register(New(errCodeDeletionFailed, "deletion failed", nil))
var ErrDeletionSuccess = register(New(errCodeDeletionSuccess, "deletion success", nil))
var ErrUndeleteFailed = register(New(errCodeUndeleteFailed, "undeletion failed", nil))
var ErrUndeleteSuccess = register(New(errCodeUndeleteSuccess, "undeletion success", nil))
var ErrCreationFailed = register(New(errCodeCreationFailed, "creation failed", nil))

//
// Root key management
//

const errCodeRootKeyEmpty = ErrorCode("root_key_empty")
const errCodeRootKeyMissing = ErrorCode("root_key_missing")
const errCodeRootKeyNotEmpty = ErrorCode("root_key_not_empty")
const errCodeRootKeySetSuccess = ErrorCode("root_key_set_success")
const errCodeRootKeySkipCreationForInMemoryMode = ErrorCode("root_key_skip_creation_for_in_memory_mode")
const errCodeRootKeyUpdateSkippedKeyEmpty = ErrorCode("root_key_update_skipped_key_empty")
const errCodeGeneralFailure = ErrorCode("general_failure")

var ErrRootKeyEmpty = register(New(errCodeRootKeyEmpty, "root key empty", nil))
var ErrRootKeyMissing = register(New(errCodeRootKeyMissing, "root key missing", nil))
var ErrRootKeyNotEmpty = register(New(errCodeRootKeyNotEmpty, "root key not empty", nil))
var ErrRootKeySetSuccess = register(New(errCodeRootKeySetSuccess, "root key set success", nil))
var ErrRootKeySkipCreationForInMemoryMode = register(New(errCodeRootKeySkipCreationForInMemoryMode, "root key skip creation for in memory mode", nil))
var ErrRootKeyUpdateSkippedKeyEmpty = register(New(errCodeRootKeyUpdateSkippedKeyEmpty, "root key update skipped key empty", nil))

//
// Shamir-related
//

const errCodeShamirDuplicateIndex = ErrorCode("shamir_duplicate_index")
const errCodeShamirEmptyShard = ErrorCode("shamir_empty_shard")
const errCodeShamirInvalidIndex = ErrorCode("shamir_invalid_index")
const errCodeShamirNilShard = ErrorCode("shamir_nil_shard")
const errCodeShamirNotEnoughShards = ErrorCode("shamir_not_enough_shards")
const errCodeShamirReconstructionFailed = ErrorCode("shamir_reconstruction_failed")

var ErrShamirDuplicateIndex = register(New(errCodeShamirDuplicateIndex, "shamir duplicate index", nil))
var ErrShamirEmptyShard = register(New(errCodeShamirEmptyShard, "shamir empty shard", nil))
var ErrShamirInvalidIndex = register(New(errCodeShamirInvalidIndex, "shamir invalid index", nil))
var ErrShamirNilShard = register(New(errCodeShamirNilShard, "shamir nil shard", nil))
var ErrShamirNotEnoughShards = register(New(errCodeShamirNotEnoughShards, "shamir not enough shards", nil))
var ErrShamirReconstructionFailed = register(New(errCodeShamirReconstructionFailed, "shamir reconstruction failed", nil))

//
// Crypto operations
//

const errCodeCryptoLowEntropy = ErrorCode("low_entropy")
const errCodeCryptoCipherNotAvailable = ErrorCode("crypto_cipher_not_available")
const errCodeCryptoCipherVerificationSuccess = ErrorCode("crypto_cipher_verification_success")
const errCodeCryptoCipherVerificationFailed = ErrorCode("crypto_cipher_verification_failed")
const errCodeCryptoDecryptionFailed = ErrorCode("crypto_decryption_failed")
const errCodeCryptoEncryptionFailed = ErrorCode("crypto_encryption_failed")
const errCodeCryptoFailedToReadNonce = ErrorCode("crypto_failed_to_read_nonce")
const errCodeCryptoFailedToReadVersion = ErrorCode("crypto_failed_to_read_version")
const errCodeCryptoNonceGenerationFailed = ErrorCode("crypto_nonce_generation_failed")
const errCodeCryptoFailedToCreateCipher = ErrorCode("crypto_failed_to_create_cipher")
const errCodeCryptoFailedToCreateGCM = ErrorCode("crypto_failed_to_create_gcm")
const errCodeCryptoInvalidEncryptionKeyLength = ErrorCode("crypto_invalid_encryption_key_length")

var ErrCryptoLowEntropy = register(New(errCodeCryptoLowEntropy, "low entropy", nil))
var ErrCryptoCipherNotAvailable = register(New(errCodeCryptoCipherNotAvailable, "cipher not available", nil))
var ErrCryptoCipherVerificationSuccess = register(New(errCodeCryptoCipherVerificationSuccess, "cipher verification success", nil))
var ErrCryptoCipherVerificationFailed = register(New(errCodeCryptoCipherVerificationFailed, "cipher verification failed", nil))
var ErrCryptoDecryptionFailed = register(New(errCodeCryptoDecryptionFailed, "decryption failed", nil))
var ErrCryptoEncryptionFailed = register(New(errCodeCryptoEncryptionFailed, "encryption failed", nil))
var ErrCryptoFailedToReadNonce = register(New(errCodeCryptoFailedToReadNonce, "failed to read nonce", nil))
var ErrCryptoFailedToReadVersion = register(New(errCodeCryptoFailedToReadVersion, "failed to read version", nil))
var ErrCryptoNonceGenerationFailed = register(New(errCodeCryptoNonceGenerationFailed, "nonce generation failed", nil))
var ErrCryptoFailedToCreateCipher = register(New(errCodeCryptoFailedToCreateCipher, "failed to create cipher", nil))
var ErrCryptoFailedToCreateGCM = register(New(errCodeCryptoFailedToCreateGCM, "failed to create GCM", nil))
var ErrCryptoInvalidEncryptionKeyLength = register(New(errCodeCryptoInvalidEncryptionKeyLength, "invalid encryption key length", nil))

//
// Store configuration
//

const errCodeStoreInvalidConfiguration = ErrorCode("store_invalid_configuration")
const errCodeStoreInvalidEncryptionKey = ErrorCode("store_invalid_encryption_key")

var ErrStoreInvalidConfiguration = register(New(errCodeStoreInvalidConfiguration, "invalid store configuration", nil))
var ErrStoreInvalidEncryptionKey = register(New(errCodeStoreInvalidEncryptionKey, "invalid store encryption key", nil))

//
// Store operations
//

const errCodeStoreLoadFailed = ErrorCode("store_load_failed")
const errCodeStoreQueryFailed = ErrorCode("store_query_failed")
const errCodeStoreSaveFailed = ErrorCode("store_save_failed")
const errCodeStoreQueryFailure = ErrorCode("store_query_failure")
const errCodeStoreResultSetFailedToLoad = ErrorCode("store_result_set_failed_to_load")
const errCodeStoreVersionNotFound = ErrorCode("store_version_not_found")
const errCodeStoreItemSoftDeleted = ErrorCode("store_item_soft_deleted")
const errCodeStoreInvalidVersion = ErrorCode("store_invalid_version")

var ErrStoreLoadFailed = register(New(errCodeStoreLoadFailed, "failed to load data", nil))
var ErrStoreQueryFailed = register(New(errCodeStoreQueryFailed, "failed to query data", nil))
var ErrStoreSaveFailed = register(New(errCodeStoreSaveFailed, "failed to save data", nil))
var ErrStoreQueryFailure = register(New(errCodeStoreQueryFailure, "failed to query for the requested data", nil))
var ErrStoreResultSetFailedToLoad = register(New(errCodeStoreResultSetFailedToLoad, "result set failed to load", nil))
var ErrStoreVersionNotFound = register(New(errCodeStoreVersionNotFound, "version not found", nil))
var ErrStoreItemSoftDeleted = register(New(errCodeStoreItemSoftDeleted, "item marked as deleted", nil))
var ErrStoreInvalidVersion = register(New(errCodeStoreInvalidVersion, "invalid version", nil))

//
// Filesystem operations
//

const errCodeDirectoryCreationFailed = ErrorCode("directory_creation_failed")
const errCodeFSFailedToCheckDirectory = ErrorCode("fs_failed_to_check_directory")
const errCodeFSFailedToCreateDirectory = ErrorCode("fs_failed_to_create_directory")
const errCodeFSFailedToResolvePath = ErrorCode("fs_failed_to_resolve_path")
const errCodeFSFileIsNotADirectory = ErrorCode("fs_file_is_not_a_directory")
const errCodeFSInvalidDirectory = ErrorCode("fs_invalid_directory")
const errCodeFSParentDirectoryDoesNotExist = ErrorCode("fs_parent_directory_does_not_exist")
const errCodeFSPathCannotBeEmpty = ErrorCode("fs_path_cannot_be_empty")
const errCodeFSPathRestricted = ErrorCode("fs_path_restricted")

var ErrDirectoryCreationFailed = register(New(errCodeDirectoryCreationFailed, "directory creation failed", nil))
var ErrFSFailedToCheckDirectory = register(New(errCodeFSFailedToCheckDirectory, "failed to check directory", nil))
var ErrFSFailedToCreateDirectory = register(New(errCodeFSFailedToCreateDirectory, "failed to create directory", nil))
var ErrFSFailedToResolvePath = register(New(errCodeFSFailedToResolvePath, "failed to resolve filesystem path", nil))
var ErrFSFileIsNotADirectory = register(New(errCodeFSFileIsNotADirectory, "file is not a directory", nil))
var ErrFSInvalidDirectory = register(New(errCodeFSInvalidDirectory, "invalid directory", nil))
var ErrFSParentDirectoryDoesNotExist = register(New(errCodeFSParentDirectoryDoesNotExist, "parent directory does not exist", nil))
var ErrFSPathCannotBeEmpty = register(New(errCodeFSPathCannotBeEmpty, "filesystem path cannot be empty", nil))
var ErrFSPathRestricted = register(New(errCodeFSPathRestricted, "filesystem path is restricted for security reasons", nil))

//
// File I/O
//

const errCodeFileCloseFailed = ErrorCode("file_close_failed")
const errCodeFileOpenFailed = ErrorCode("file_open_failed")
const errCodeStreamCloseFailed = ErrorCode("stream_close_failed")
const errCodeStreamOpenFailed = ErrorCode("stream_open_failed")

var ErrFileCloseFailed = register(New(errCodeFileCloseFailed, "file close failed", nil))
var ErrStreamCloseFailed = register(New(errCodeStreamCloseFailed, "stream close failed", nil))
var ErrStreamOpenFailed = register(New(errCodeStreamOpenFailed, "stream open failed", nil))
var ErrFileOpenFailed = register(New(errCodeFileOpenFailed, "file open failed", nil))

//
// Input/Output
//

const errCodeInvalidInput = ErrorCode("invalid_input")
const errCodeInvalidPermission = ErrorCode("invalid_permission")
const errCodeMarshalFailure = ErrorCode("marshal_failure")
const errCodeParseFailure = ErrorCode("parse_failure")
const errCodeReadFailure = ErrorCode("read_failure")
const errCodeUnmarshalFailure = ErrorCode("unmarshal_failure")

var ErrInvalidInput = register(New(errCodeInvalidInput, "invalid input", nil))
var ErrInvalidPermission = register(New(errCodeInvalidPermission, "invalid permission", nil))
var ErrMarshalFailure = register(New(errCodeMarshalFailure, "failed to marshal response body", nil))
var ErrParseFailure = register(New(errCodeParseFailure, "failed to parse request body", nil))
var ErrReadFailure = register(New(errCodeReadFailure, "failed to read request body", nil))
var ErrUnmarshalFailure = register(New(errCodeUnmarshalFailure, "failed to unmarshal request body", nil))

//
// Network/Peer
//

const errCodePeerConnection = ErrorCode("peer_connection")
const errCodeReadingResponseBody = ErrorCode("reading_response_body")
const errCodeReadingRequestBody = ErrorCode("reading_request_body")

var ErrPeerConnection = register(New(errCodePeerConnection, "problem connecting to peer", nil))
var ErrReadingResponseBody = register(New(errCodeReadingResponseBody, "problem reading response body", nil))
var ErrReadingRequestBody = register(New(errCodeReadingRequestBody, "problem reading requestbody", nil))

//
// Transaction operations
//

const errCodeTransactionBeginFailed = ErrorCode("transaction_begin_failed")
const errCodeTransactionCommitFailed = ErrorCode("transaction_commit_failed")
const errCodeTransactionFailed = ErrorCode("transaction_failed")
const errCodeTransactionRollbackFailed = ErrorCode("transaction_rollback_failed")

var ErrTransactionBeginFailed = register(New(errCodeTransactionBeginFailed, "failed to begin transaction", nil))
var ErrTransactionCommitFailed = register(New(errCodeTransactionCommitFailed, "failed to commit transaction", nil))
var ErrTransactionFailed = register(New(errCodeTransactionFailed, "transaction failed", nil))
var ErrTransactionRollbackFailed = register(New(errCodeTransactionRollbackFailed, "failed to rollback transaction", nil))

//
// Recovery operations
//

const errCodeRecoveryRetryFailed = ErrorCode("recovery_retry_failed")
const errCodeRecoveryRetryLimitReached = ErrorCode("recovery_retry_limit_reached")
const errCodeRecoveryFailed = ErrorCode("recovery_failed")

var ErrRecoveryRetryFailed = register(New(errCodeRecoveryRetryFailed, "recovery retry failed", nil))
var ErrRecoveryRetryLimitReached = register(New(errCodeRecoveryRetryLimitReached, "recovery retry limit reached", nil))
var ErrRecoveryFailed = register(New(errCodeRecoveryFailed, "recovery failed", nil))

//
// X509/SPIFFE
//

const errCodeSPIFFENilX509Source = ErrorCode("spiffe_nil_x509_source")
const errCodeSPIFFEEmptyTrustDomain = ErrorCode("spiffe_empty_trust_domain")
const errCodeSPIFFEMultipleTrustDomains = ErrorCode("spiffe_multiple_trust_domains")
const errCodeSPIFFEFailedToCreateX509Source = ErrorCode("spiffe_failed_to_create_x509_source")
const errCodeSPIFFEUnableToFetchX509Source = ErrorCode("spiffe_unable_to_fetch_x509_source")
const errCodeSPIFFENoPeerCertificates = ErrorCode("spiffe_no_peer_certificates")
const errCodeSPIFFEFailedToExtractX509SVID = ErrorCode("spiffe_failed_to_extract_x509_svid")
const errCodeSPIFFEFailedToCloseSource = ErrorCode("spiffe_failed_to_close_source")

var ErrSPIFFENilX509Source = register(New(errCodeSPIFFENilX509Source, "nil X509Source", nil))
var ErrSPIFFEEmptyTrustDomain = register(New(errCodeSPIFFEEmptyTrustDomain, "empty trust domain", nil))
var ErrSPIFFEMultipleTrustDomains = register(New(errCodeSPIFFEMultipleTrustDomains, "provide a single trust domain", nil))
var ErrSPIFFEFailedToCreateX509Source = register(New(errCodeSPIFFEFailedToCreateX509Source, "failed to create X509Source", nil))
var ErrSPIFFEUnableToFetchX509Source = register(New(errCodeSPIFFEUnableToFetchX509Source, "unable to fetch X509Source", nil))
var ErrSPIFFENoPeerCertificates = register(New(errCodeSPIFFENoPeerCertificates, "no peer certificates", nil))
var ErrSPIFFEFailedToExtractX509SVID = register(New(errCodeSPIFFEFailedToExtractX509SVID, "failed to extract X509 SVID", nil))

// FromCode maps an ErrorCode to its corresponding SDKError using the
// automatically populated error registry. This is used to convert error codes
// received from API responses back to proper SDKError instances.
//
// The registry is automatically populated when errors are defined using the
// register() function, ensuring new errors are immediately available without
// manual updates to this function.
//
// If the error code is not recognized, it returns ErrGeneralFailure.
//
// Parameters:
//   - code: the error code to map
//
// Returns:
//   - *SDKError: the corresponding SDK error instance
func FromCode(code ErrorCode) *SDKError {
	// Defensive coding: While concurrent reads to a map are safe, unless a
	// write happens concurrently; if we enable dynamic error registration
	// later down the line, without a mutex the behavior of this code will be
	// undeterministic.
	errorRegistryMu.RLock()
	err, ok := errorRegistry[code]
	errorRegistryMu.RUnlock()

	if ok {
		return err
	}
	return ErrGeneralFailure
}

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
