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
const errCodeGeneralFailure = ErrorCode("general_failure")
const errCodeInternal = ErrorCode("internal_error")
const errCodeNilContext = ErrorCode("nil_context")
const errCodeNotFound = ErrorCode("not_found")
const errCodePostFailed = ErrorCode("post_failed")
const errCodeServerFault = ErrorCode("server_fault")
const errCodeSuccess = ErrorCode("success")

var ErrBadRequest = register(New(errCodeBadRequest, "bad request", nil))
var ErrEmptyPayload = register(New(errCodeEmptyPayload, "empty payload", nil))
var ErrFound = register(New(errCodeFound, "found", nil))
var ErrGeneralFailure = register(New(errCodeGeneralFailure, "general failure", nil))
var ErrInternal = register(New(errCodeInternal, "internal error", nil))
var ErrNilContext = register(New(errCodeNilContext, "nil context", nil))
var ErrNotFound = register(New(errCodeNotFound, "not found", nil))
var ErrPostFailed = register(New(errCodePostFailed, "post failed", nil))
var ErrServerFault = register(New(errCodeServerFault, "server fault", nil))
var ErrSuccess = register(New(errCodeSuccess, "success", nil))

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

const errCodeStateAlreadyInitialized = ErrorCode("state_already_initialized")
const errCodeStateInitializationFailed = ErrorCode("state_initialization_failed")
const errCodeStateNotAlive = ErrorCode("state_not_alive")
const errCodeStateNotReady = ErrorCode("state_not_ready")

var ErrStateAlreadyInitialized = register(New(errCodeStateAlreadyInitialized, "already initialized", nil))
var ErrStateInitializationFailed = register(New(errCodeStateInitializationFailed, "initialization failed", nil))
var ErrStateNotAlive = register(New(errCodeStateNotAlive, "not alive", nil))
var ErrStateNotReady = register(New(errCodeStateNotReady, "not ready", nil))

//
// Policy/RBAC/ABAC
//

const errCodeAccessInvalidPermission = ErrorCode("access_invalid_permission")
const errCodeAccessUnauthorized = ErrorCode("access_unauthorized")

var ErrAccessInvalidPermission = register(New(errCodeAccessInvalidPermission, "invalid permission", nil))
var ErrAccessUnauthorized = register(New(errCodeAccessUnauthorized, "unauthorized", nil))

//
// CRUD operations
//

const errCodeObjectCreationFailed = ErrorCode("object_creation_failed")
const errCodeObjectDeletionFailed = ErrorCode("object_deletion_failed")
const errCodeObjectDeletionSuccess = ErrorCode("object_deletion_success")
const errCodeObjectUndeletionFailed = ErrorCode("object_undeletion_failed")
const errCodeObjectUndeletionSuccess = ErrorCode("object_undeletion_success")

var ErrObjectCreationFailed = register(New(errCodeObjectCreationFailed, "creation failed", nil))
var ErrObjectDeletionFailed = register(New(errCodeObjectDeletionFailed, "deletion failed", nil))
var ErrObjectDeletionSuccess = register(New(errCodeObjectDeletionSuccess, "deletion success", nil))
var ErrObjectUndeletionFailed = register(New(errCodeObjectUndeletionFailed, "undeletion failed", nil))
var ErrObjectUndeletionSuccess = register(New(errCodeObjectUndeletionSuccess, "undeletion success", nil))

//
// Root key management
//

const errCodeRootKeyEmpty = ErrorCode("root_key_empty")
const errCodeRootKeyMissing = ErrorCode("root_key_missing")
const errCodeRootKeyNotEmpty = ErrorCode("root_key_not_empty")
const errCodeRootKeySetSuccess = ErrorCode("root_key_set_success")
const errCodeRootKeySkipCreationForInMemoryMode = ErrorCode("root_key_skip_creation_for_in_memory_mode")
const errCodeRootKeyUpdateSkippedKeyEmpty = ErrorCode("root_key_update_skipped_key_empty")

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

const errCodeCryptoCipherNotAvailable = ErrorCode("crypto_cipher_not_available")
const errCodeCryptoCipherVerificationFailed = ErrorCode("crypto_cipher_verification_failed")
const errCodeCryptoCipherVerificationSuccess = ErrorCode("crypto_cipher_verification_success")
const errCodeCryptoDecryptionFailed = ErrorCode("crypto_decryption_failed")
const errCodeCryptoEncryptionFailed = ErrorCode("crypto_encryption_failed")
const errCodeCryptoFailedToCreateCipher = ErrorCode("crypto_failed_to_create_cipher")
const errCodeCryptoFailedToCreateGCM = ErrorCode("crypto_failed_to_create_gcm")
const errCodeCryptoFailedToReadNonce = ErrorCode("crypto_failed_to_read_nonce")
const errCodeCryptoFailedToReadVersion = ErrorCode("crypto_failed_to_read_version")
const errCodeCryptoInvalidEncryptionKeyLength = ErrorCode("crypto_invalid_encryption_key_length")
const errCodeCryptoLowEntropy = ErrorCode("low_entropy")
const errCodeCryptoNonceGenerationFailed = ErrorCode("crypto_nonce_generation_failed")
const errCodeCryptoRandomGenerationFailed = ErrorCode("crypto_random_generation_failed")

var ErrCryptoCipherNotAvailable = register(New(errCodeCryptoCipherNotAvailable, "cipher not available", nil))
var ErrCryptoCipherVerificationFailed = register(New(errCodeCryptoCipherVerificationFailed, "cipher verification failed", nil))
var ErrCryptoCipherVerificationSuccess = register(New(errCodeCryptoCipherVerificationSuccess, "cipher verification success", nil))
var ErrCryptoDecryptionFailed = register(New(errCodeCryptoDecryptionFailed, "decryption failed", nil))
var ErrCryptoEncryptionFailed = register(New(errCodeCryptoEncryptionFailed, "encryption failed", nil))
var ErrCryptoFailedToCreateCipher = register(New(errCodeCryptoFailedToCreateCipher, "failed to create cipher", nil))
var ErrCryptoFailedToCreateGCM = register(New(errCodeCryptoFailedToCreateGCM, "failed to create GCM", nil))
var ErrCryptoFailedToReadNonce = register(New(errCodeCryptoFailedToReadNonce, "failed to read nonce", nil))
var ErrCryptoFailedToReadVersion = register(New(errCodeCryptoFailedToReadVersion, "failed to read version", nil))
var ErrCryptoInvalidEncryptionKeyLength = register(New(errCodeCryptoInvalidEncryptionKeyLength, "invalid encryption key length", nil))
var ErrCryptoLowEntropy = register(New(errCodeCryptoLowEntropy, "low entropy", nil))
var ErrCryptoNonceGenerationFailed = register(New(errCodeCryptoNonceGenerationFailed, "nonce generation failed", nil))
var ErrCryptoRandomGenerationFailed = register(New(errCodeCryptoRandomGenerationFailed, "random generation failed", nil))

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

const errCodeStoreInvalidVersion = ErrorCode("store_invalid_version")
const errCodeStoreItemSoftDeleted = ErrorCode("store_item_soft_deleted")
const errCodeStoreLoadFailed = ErrorCode("store_load_failed")
const errCodeStoreQueryFailed = ErrorCode("store_query_failed")
const errCodeStoreQueryFailure = ErrorCode("store_query_failure")
const errCodeStoreResultSetFailedToLoad = ErrorCode("store_result_set_failed_to_load")
const errCodeStoreSaveFailed = ErrorCode("store_save_failed")
const errCodeStoreVersionNotFound = ErrorCode("store_version_not_found")

var ErrStoreInvalidVersion = register(New(errCodeStoreInvalidVersion, "invalid version", nil))
var ErrStoreItemSoftDeleted = register(New(errCodeStoreItemSoftDeleted, "item marked as deleted", nil))
var ErrStoreLoadFailed = register(New(errCodeStoreLoadFailed, "failed to load data", nil))
var ErrStoreQueryFailed = register(New(errCodeStoreQueryFailed, "failed to query data", nil))
var ErrStoreQueryFailure = register(New(errCodeStoreQueryFailure, "failed to query for the requested data", nil))
var ErrStoreResultSetFailedToLoad = register(New(errCodeStoreResultSetFailedToLoad, "result set failed to load", nil))
var ErrStoreSaveFailed = register(New(errCodeStoreSaveFailed, "failed to save data", nil))
var ErrStoreVersionNotFound = register(New(errCodeStoreVersionNotFound, "version not found", nil))

//
// Filesystem operations
//

const errCodeFSDirectoryCreationFailed = ErrorCode("fs_directory_creation_failed")
const errCodeFSFailedToCheckDirectory = ErrorCode("fs_failed_to_check_directory")
const errCodeFSFailedToCreateDirectory = ErrorCode("fs_failed_to_create_directory")
const errCodeFSFailedToResolvePath = ErrorCode("fs_failed_to_resolve_path")
const errCodeFSFileCloseFailed = ErrorCode("fs_file_close_failed")
const errCodeFSFileIsNotADirectory = ErrorCode("fs_file_is_not_a_directory")
const errCodeFSFileOpenFailed = ErrorCode("fs_file_open_failed")
const errCodeFSInvalidDirectory = ErrorCode("fs_invalid_directory")
const errCodeFSParentDirectoryDoesNotExist = ErrorCode("fs_parent_directory_does_not_exist")
const errCodeFSPathCannotBeEmpty = ErrorCode("fs_path_cannot_be_empty")
const errCodeFSPathRestricted = ErrorCode("fs_path_restricted")
const errCodeFSStreamCloseFailed = ErrorCode("fs_stream_close_failed")
const errCodeFSStreamOpenFailed = ErrorCode("fs_stream_open_failed")

var ErrDirectoryCreationFailed = register(New(errCodeFSDirectoryCreationFailed, "directory creation failed", nil))
var ErrFSFailedToCheckDirectory = register(New(errCodeFSFailedToCheckDirectory, "failed to check directory", nil))
var ErrFSFailedToCreateDirectory = register(New(errCodeFSFailedToCreateDirectory, "failed to create directory", nil))
var ErrFSFailedToResolvePath = register(New(errCodeFSFailedToResolvePath, "failed to resolve filesystem path", nil))
var ErrFSFileCloseFailed = register(New(errCodeFSFileCloseFailed, "file close failed", nil))
var ErrFSFileIsNotADirectory = register(New(errCodeFSFileIsNotADirectory, "file is not a directory", nil))
var ErrFSFileOpenFailed = register(New(errCodeFSFileOpenFailed, "file open failed", nil))
var ErrFSInvalidDirectory = register(New(errCodeFSInvalidDirectory, "invalid directory", nil))
var ErrFSParentDirectoryDoesNotExist = register(New(errCodeFSParentDirectoryDoesNotExist, "parent directory does not exist", nil))
var ErrFSPathCannotBeEmpty = register(New(errCodeFSPathCannotBeEmpty, "filesystem path cannot be empty", nil))
var ErrFSPathRestricted = register(New(errCodeFSPathRestricted, "filesystem path is restricted for security reasons", nil))
var ErrFSStreamCloseFailed = register(New(errCodeFSStreamCloseFailed, "stream close failed", nil))
var ErrFSStreamOpenFailed = register(New(errCodeFSStreamOpenFailed, "stream open failed", nil))

//
// Data Processing
//

const errCodeDataInvalidInput = ErrorCode("data_invalid_input")
const errCodeDataMarshalFailure = ErrorCode("data_marshal_failure")
const errCodeDataParseFailure = ErrorCode("data_parse_failure")
const errCodeDataReadFailure = ErrorCode("data_read_failure")
const errCodeDataUnmarshalFailure = ErrorCode("data_unmarshal_failure")

var ErrDataInvalidInput = register(New(errCodeDataInvalidInput, "invalid input", nil))
var ErrDataMarshalFailure = register(New(errCodeDataMarshalFailure, "failed to marshal response body", nil))
var ErrDataParseFailure = register(New(errCodeDataParseFailure, "failed to parse request body", nil))
var ErrDataReadFailure = register(New(errCodeDataReadFailure, "failed to read request body", nil))
var ErrDataUnmarshalFailure = register(New(errCodeDataUnmarshalFailure, "failed to unmarshal request body", nil))

//
// String/Template operations
//

const errCodeStringEmptyCharacterClass = ErrorCode("string_empty_character_class")
const errCodeStringInvalidRange = ErrorCode("string_invalid_range")
const errCodeStringEmptyCharacterSet = ErrorCode("string_empty_character_set")
const errCodeStringInvalidLength = ErrorCode("string_invalid_length")
const errCodeStringNegativeLength = ErrorCode("string_negative_length")

var ErrStringEmptyCharacterClass = register(New(errCodeStringEmptyCharacterClass, "empty character class", nil))
var ErrStringInvalidRange = register(New(errCodeStringInvalidRange, "invalid character range", nil))
var ErrStringEmptyCharacterSet = register(New(errCodeStringEmptyCharacterSet, "character class resulted in empty set", nil))
var ErrStringInvalidLength = register(New(errCodeStringInvalidLength, "invalid length specification", nil))
var ErrStringNegativeLength = register(New(errCodeStringNegativeLength, "length cannot be negative", nil))

//
// Network/Peer
//

const errCodeNetPeerConnection = ErrorCode("net_peer_connection")
const errCodeNetReadingRequestBody = ErrorCode("net_reading_request_body")
const errCodeNetReadingResponseBody = ErrorCode("net_reading_response_body")
const errCodeNetURLJoinPathFailed = ErrorCode("net_url_join_path_failed")

var ErrNetPeerConnection = register(New(errCodeNetPeerConnection, "problem connecting to peer", nil))
var ErrNetReadingRequestBody = register(New(errCodeNetReadingRequestBody, "problem reading request body", nil))
var ErrNetReadingResponseBody = register(New(errCodeNetReadingResponseBody, "problem reading response body", nil))
var ErrNetURLJoinPathFailed = register(New(errCodeNetURLJoinPathFailed, "failed to join URL path", nil))

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
// Retry operations
//

const errCodeRetryMaxElapsedTimeReached = ErrorCode("retry_max_elapsed_time_reached")
const errCodeRetryContextCanceled = ErrorCode("retry_context_canceled")
const errCodeRetryOperationFailed = ErrorCode("retry_operation_failed")

var ErrRetryMaxElapsedTimeReached = register(New(errCodeRetryMaxElapsedTimeReached, "maximum elapsed time for retries reached", nil))
var ErrRetryContextCanceled = register(New(errCodeRetryContextCanceled, "retry canceled due to context cancellation", nil))
var ErrRetryOperationFailed = register(New(errCodeRetryOperationFailed, "retry operation failed", nil))

//
// X509/SPIFFE
//

const errCodeSPIFFEEmptyTrustDomain = ErrorCode("spiffe_empty_trust_domain")
const errCodeSPIFFEFailedToCloseX509Source = ErrorCode("spiffe_failed_to_close_source")
const errCodeSPIFFEFailedToCreateX509Source = ErrorCode("spiffe_failed_to_create_x509_source")
const errCodeSPIFFEFailedToExtractX509SVID = ErrorCode("spiffe_failed_to_extract_x509_svid")
const errCodeSPIFFEMultipleTrustDomains = ErrorCode("spiffe_multiple_trust_domains")
const errCodeSPIFFENilX509Source = ErrorCode("spiffe_nil_x509_source")
const errCodeSPIFFENoPeerCertificates = ErrorCode("spiffe_no_peer_certificates")
const errCodeSPIFFEUnableToFetchX509Source = ErrorCode("spiffe_unable_to_fetch_x509_source")

var ErrSPIFFEEmptyTrustDomain = register(New(errCodeSPIFFEEmptyTrustDomain, "empty trust domain", nil))
var ErrSPIFFEFailedToCreateX509Source = register(New(errCodeSPIFFEFailedToCreateX509Source, "failed to create X509Source", nil))
var ErrSPIFFEFailedToExtractX509SVID = register(New(errCodeSPIFFEFailedToExtractX509SVID, "failed to extract X509 SVID", nil))
var ErrSPIFFEMultipleTrustDomains = register(New(errCodeSPIFFEMultipleTrustDomains, "provide a single trust domain", nil))
var ErrSPIFFENilX509Source = register(New(errCodeSPIFFENilX509Source, "nil X509Source", nil))
var ErrSPIFFENoPeerCertificates = register(New(errCodeSPIFFENoPeerCertificates, "no peer certificates", nil))
var ErrSPIFFEUnableToFetchX509Source = register(New(errCodeSPIFFEUnableToFetchX509Source, "unable to fetch X509Source", nil))
var ErrSPIFFEFailedToCloseX509Source = register(New(errCodeSPIFFEFailedToCloseX509Source, "failed to close X509Source", nil))

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
