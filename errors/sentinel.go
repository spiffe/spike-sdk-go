//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

// -----------------------------------------------------------------------------
// SENTINEL ERRORS - IMPORTANT USAGE NOTES
// -----------------------------------------------------------------------------
//
// All errors defined in this file are SHARED GLOBAL POINTERS (*SDKError).
// They must NEVER be mutated directly.
//
// RULES FOR INTERNAL DEVELOPERS:
//
//  1. NEVER assign a sentinel to a variable and then modify it:
//
//     // WRONG - corrupts the shared sentinel:
//     failErr := ErrEntityNotFound
//     failErr.Msg = "custom"
//
//  2. ALWAYS use Clone() before modifying Msg:
//
//     // CORRECT:
//     failErr := ErrEntityNotFound.Clone()
//     failErr.Msg = "custom"
//
//  3. ALWAYS use Clone() when returning sentinels (defensive programming):
//
//     // CORRECT:
//     return ErrEntityNotFound.Clone()
//
//  4. Wrap() is safe because it creates a new instance:
//
//     // SAFE:
//     failErr := ErrEntityNotFound.Wrap(err)
//     failErr.Msg = "custom"
//
// These rules prevent consumers from accidentally mutating shared state.
// See doc.go for full documentation.
// -----------------------------------------------------------------------------

//
// General error codes
//

// ErrGeneralFailure indicates a general unspecified failure.
var ErrGeneralFailure = register("gen_general_failure", "general failure", nil)

// ErrNilContext indicates a nil context was provided.
var ErrNilContext = register("gen_nil_context", "nil context", nil)

//
// Cluster operations
//

// ErrK8sReconciliationFailed indicates Kubernetes reconciliation failed.
var ErrK8sReconciliationFailed = register("k8s_reconciliation_failed", "reconciliation failed", nil)

// ErrK8sClientFailed indicates Kubernetes client creation failed.
var ErrK8sClientFailed = register("k8s_client_failed", "failed to create Kubernetes client", nil)

// ErrK8sResourceLookupFailed indicates Kubernetes resource lookup failed.
var ErrK8sResourceLookupFailed = register("k8s_lookup_failed", "failed to lookup Kubernetes resource", nil)

//
// API/HTTP operations
//

// ErrAPIBadRequest indicates the API request was malformed or invalid.
var ErrAPIBadRequest = register("api_bad_request", "bad request", nil)

// ErrAPIEmptyPayload indicates the API request payload was empty.
var ErrAPIEmptyPayload = register("api_empty_payload", "empty payload", nil)

// ErrAPIFound indicates the requested resource was found.
var ErrAPIFound = register("api_found", "found", nil)

// ErrAPIInternal indicates an internal server error occurred.
var ErrAPIInternal = register("api_internal_error", "internal error", nil)

// ErrAPINotFound indicates the requested resource was not found.
var ErrAPINotFound = register("api_not_found", "not found", nil)

// ErrAPIPostFailed indicates the HTTP POST request failed.
var ErrAPIPostFailed = register("api_post_failed", "post failed", nil)

// ErrAPIResponseCodeInvalid indicates an invalid API response code was received.
var ErrAPIResponseCodeInvalid = register("api_response_code_invalid", "invalid API response code", nil)

// ErrAPIServerFault indicates a server-side fault occurred.
var ErrAPIServerFault = register("api_server_fault", "server fault", nil)

//
// Entity operations
//

// ErrEntityCreationFailed indicates object creation failed.
var ErrEntityCreationFailed = register("object_creation_failed", "creation failed", nil)

// ErrEntityDeleted indicates the entity is marked as deleted.
var ErrEntityDeleted = register("entity_deleted", "entity marked as deleted", nil)

// ErrEntityDeletionFailed indicates object deletion failed.
var ErrEntityDeletionFailed = register("object_deletion_failed", "deletion failed", nil)

// ErrEntityExists indicates the entity already exists.
var ErrEntityExists = register("entity_exists", "entity already exists", nil)

// ErrEntityInvalid indicates the entity data is invalid.
var ErrEntityInvalid = register("entity_invalid", "entity is invalid", nil)

// ErrEntityLoadFailed indicates the entity failed to load from storage.
var ErrEntityLoadFailed = register("entity_load_failed", "failed to load entity", nil)

// ErrEntityNotFound indicates the requested entity was not found.
var ErrEntityNotFound = register("entity_not_found", "entity not found", nil)

// ErrEntityQueryFailed indicates the entity query operation failed.
var ErrEntityQueryFailed = register("entity_query_failed", "failed to query entities", nil)

// ErrEntitySaveFailed indicates the entity failed to get saved to storage.
var ErrEntitySaveFailed = register("entity_save_failed", "failed to save entity", nil)

// ErrEntityVersionInvalid indicates an invalid entity version was specified.
var ErrEntityVersionInvalid = register("entity_version_invalid", "invalid version", nil)

// ErrEntityVersionNotFound indicates the requested entity version was not found.
var ErrEntityVersionNotFound = register("entity_version_not_found", "version not found", nil)

// ErrEntityUndeletionFailed indicates object undeletion (restore) failed.
var ErrEntityUndeletionFailed = register("object_undeletion_failed", "undeletion failed", nil)

//
// State operations
//

// ErrStateAlreadyInitialized indicates the system state is already initialized.
var ErrStateAlreadyInitialized = register("state_already_initialized", "already initialized", nil)

// ErrStateInitializationFailed indicates system state initialization failed.
var ErrStateInitializationFailed = register("state_initialization_failed", "initialization failed", nil)

// ErrStateNotAlive indicates the system state is not alive.
var ErrStateNotAlive = register("state_not_alive", "not alive", nil)

// ErrStateNotReady indicates the system state is not ready.
var ErrStateNotReady = register("state_not_ready", "not ready", nil)

// ErrStateIntegrityCheck indicates the system state integrity check failed.
var ErrStateIntegrityCheck = register("state_integrity_check", "state integrity check failed", nil)

//
// Policy/RBAC/ABAC
//

// ErrAccessInvalidPermission indicates an invalid permission was specified.
var ErrAccessInvalidPermission = register("access_invalid_permission", "invalid permission", nil)

// ErrAccessUnauthorized indicates the requester lacks authorization for the operation.
var ErrAccessUnauthorized = register("access_unauthorized", "unauthorized", nil)

//
// Root key management
//

// ErrRootKeyEmpty indicates the root key is empty.
var ErrRootKeyEmpty = register("root_key_empty", "root key empty", nil)

// ErrRootKeyMissing indicates the root key is missing.
var ErrRootKeyMissing = register("root_key_missing", "root key missing", nil)

// ErrRootKeyNotEmpty indicates the root key is not empty when expected to be.
var ErrRootKeyNotEmpty = register("root_key_not_empty", "root key not empty", nil)

// ErrRootKeySkipCreationForInMemoryMode indicates root key creation was skipped for in-memory mode.
var ErrRootKeySkipCreationForInMemoryMode = register("root_key_skip_creation_for_in_memory_mode", "root key skip creation for in memory mode", nil)

// ErrRootKeyUpdateSkippedKeyEmpty indicates the root key update was skipped because the key is empty.
var ErrRootKeyUpdateSkippedKeyEmpty = register("root_key_update_skipped_key_empty", "root key update skipped key empty", nil)

//
// Shamir-related
//

// ErrShamirDuplicateIndex indicates a duplicate shard index was provided.
var ErrShamirDuplicateIndex = register("shamir_duplicate_index", "shamir duplicate index", nil)

// ErrShamirEmptyShard indicates a Shamir shard is empty.
var ErrShamirEmptyShard = register("shamir_empty_shard", "shamir empty shard", nil)

// ErrShamirInvalidIndex indicates an invalid Shamir shard index.
var ErrShamirInvalidIndex = register("shamir_invalid_index", "shamir invalid index", nil)

// ErrShamirNilShard indicates a nil Shamir shard was provided.
var ErrShamirNilShard = register("shamir_nil_shard", "shamir nil shard", nil)

// ErrShamirNotEnoughShards indicates insufficient Shamir shards for reconstruction.
var ErrShamirNotEnoughShards = register("shamir_not_enough_shards", "shamir not enough shards", nil)

// ErrShamirReconstructionFailed indicates Shamir secret reconstruction failed.
var ErrShamirReconstructionFailed = register("shamir_reconstruction_failed", "shamir reconstruction failed", nil)

//
// Crypto operations
//

// ErrCryptoCipherNotAvailable indicates the requested cipher is not available.
var ErrCryptoCipherNotAvailable = register("crypto_cipher_not_available", "cipher not available", nil)

// ErrCryptoCipherVerificationFailed indicates cipher verification failed.
var ErrCryptoCipherVerificationFailed = register("crypto_cipher_verification_failed", "cipher verification failed", nil)

// ErrCryptoDecryptionFailed indicates data decryption failed.
var ErrCryptoDecryptionFailed = register("crypto_decryption_failed", "decryption failed", nil)

// ErrCryptoEncryptionFailed indicates data encryption failed.
var ErrCryptoEncryptionFailed = register("crypto_encryption_failed", "encryption failed", nil)

// ErrCryptoFailedToCreateCipher indicates cipher creation failed.
var ErrCryptoFailedToCreateCipher = register("crypto_failed_to_create_cipher", "failed to create cipher", nil)

// ErrCryptoFailedToCreateGCM indicates GCM mode creation failed.
var ErrCryptoFailedToCreateGCM = register("crypto_failed_to_create_gcm", "failed to create GCM", nil)

// ErrCryptoFailedToReadNonce indicates nonce reading failed.
var ErrCryptoFailedToReadNonce = register("crypto_failed_to_read_nonce", "failed to read nonce", nil)

// ErrCryptoFailedToReadVersion indicates version reading failed.
var ErrCryptoFailedToReadVersion = register("crypto_failed_to_read_version", "failed to read version", nil)

// ErrCryptoInvalidEncryptionKeyLength indicates the encryption key length is invalid.
var ErrCryptoInvalidEncryptionKeyLength = register("crypto_invalid_encryption_key_length", "invalid encryption key length", nil)

// ErrCryptoLowEntropy indicates insufficient entropy for cryptographic operations.
var ErrCryptoLowEntropy = register("crypto_low_entropy", "low entropy", nil)

// ErrCryptoNonceGenerationFailed indicates nonce generation failed.
var ErrCryptoNonceGenerationFailed = register("crypto_nonce_generation_failed", "nonce generation failed", nil)

// ErrCryptoNonceSizeMismatch indicates the nonce size does not match the cipher block size.
var ErrCryptoNonceSizeMismatch = register("crypto_nonce_size_mismatch", "nonce size mismatch", nil)

// ErrCryptoRandomGenerationFailed indicates random data generation failed.
var ErrCryptoRandomGenerationFailed = register("crypto_random_generation_failed", "random generation failed", nil)

// ErrCryptoUnsupportedCipherVersion indicates the requested cipher version is not supported.
var ErrCryptoUnsupportedCipherVersion = register("crypto_unsupported_version", "unsupported crypto version", nil)

//
// Backing store infrastructure (internal store operations)
//

// ErrStoreCloseFailed indicates the backing store close operation failed.
var ErrStoreCloseFailed = register("store_close_failed", "backing store close failed", nil)

// ErrStoreInvalidConfiguration indicates the backing store configuration is invalid.
var ErrStoreInvalidConfiguration = register("store_invalid_configuration", "invalid store configuration", nil)

// ErrStoreInvalidEncryptionKey indicates the store encryption key is invalid.
var ErrStoreInvalidEncryptionKey = register("store_invalid_encryption_key", "invalid store encryption key", nil)

//
// Filesystem operations
//

// ErrFSDirectoryCreationFailed indicates filesystem directory creation failed.
var ErrFSDirectoryCreationFailed = register("fs_directory_creation_failed", "directory creation failed", nil)

// ErrFSDirectoryDoesNotExist indicates the directory does not exist.
var ErrFSDirectoryDoesNotExist = register("fs_directory_does_not_exist", "directory does not exist", nil)

// ErrFSFailedToCheckDirectory indicates failed to check directory status.
var ErrFSFailedToCheckDirectory = register("fs_failed_to_check_directory", "failed to check directory", nil)

// ErrFSFailedToCreateDirectory indicates failed to create the directory.
var ErrFSFailedToCreateDirectory = register("fs_failed_to_create_directory", "failed to create directory", nil)

// ErrFSFailedToResolvePath indicates failed to resolve the filesystem path.
var ErrFSFailedToResolvePath = register("fs_failed_to_resolve_path", "failed to resolve filesystem path", nil)

// ErrFSFileCloseFailed indicates file close operation failed.
var ErrFSFileCloseFailed = register("fs_file_close_failed", "file close failed", nil)

// ErrFSFileIsNotADirectory indicates the path is not a directory.
var ErrFSFileIsNotADirectory = register("fs_file_is_not_a_directory", "file is not a directory", nil)

// ErrFSFileOpenFailed indicates file open operation failed.
var ErrFSFileOpenFailed = register("fs_file_open_failed", "file open failed", nil)

// ErrFSInvalidDirectory indicates an invalid directory path.
var ErrFSInvalidDirectory = register("fs_invalid_directory", "invalid directory", nil)

// ErrFSParentDirectoryDoesNotExist indicates the parent directory does not exist.
var ErrFSParentDirectoryDoesNotExist = register("fs_parent_directory_does_not_exist", "parent directory does not exist", nil)

// ErrFSPathCannotBeEmpty indicates the filesystem path cannot be empty.
var ErrFSPathCannotBeEmpty = register("fs_path_cannot_be_empty", "filesystem path cannot be empty", nil)

// ErrFSPathRestricted indicates the filesystem path is restricted for security reasons.
var ErrFSPathRestricted = register("fs_path_restricted", "filesystem path is restricted for security reasons", nil)

// ErrFSStreamCloseFailed indicates stream close operation failed.
var ErrFSStreamCloseFailed = register("fs_stream_close_failed", "stream close failed", nil)

// ErrFSStreamReadFailed indicates stream read operation failed.
var ErrFSStreamReadFailed = register("fs_stream_read_failed", "stream read failed", nil)

// ErrFSStreamWriteFailed indicates a stream write operation failed.
var ErrFSStreamWriteFailed = register("stream_write_failed", "stream write failed", nil)

// ErrFSStreamOpenFailed indicates stream open operation failed.
var ErrFSStreamOpenFailed = register("fs_stream_open_failed", "stream open failed", nil)

//
// Data Processing
//

// ErrDataInvalidInput indicates invalid input data was provided.
var ErrDataInvalidInput = register("data_invalid_input", "invalid input", nil)

// ErrDataMarshalFailure indicates data marshaling failed.
var ErrDataMarshalFailure = register("data_marshal_failure", "failed to marshal response body", nil)

// ErrDataParseFailure indicates data parsing failed.
var ErrDataParseFailure = register("data_parse_failure", "failed to parse request body", nil)

// ErrDataReadFailure indicates data reading failed.
var ErrDataReadFailure = register("data_read_failure", "failed to read request body", nil)

// ErrDataUnmarshalFailure indicates data unmarshaling failed.
var ErrDataUnmarshalFailure = register("data_unmarshal_failure", "failed to unmarshal request body", nil)

//
// String/Template operations
//

// ErrStringEmptyCharacterClass indicates an empty character class was specified.
var ErrStringEmptyCharacterClass = register("string_empty_character_class", "empty character class", nil)

// ErrStringEmptyCharacterSet indicates the character class resulted in an empty set.
var ErrStringEmptyCharacterSet = register("string_empty_character_set", "character class resulted in empty set", nil)

// ErrStringInvalidLength indicates an invalid length specification.
var ErrStringInvalidLength = register("string_invalid_length", "invalid length specification", nil)

// ErrStringInvalidRange indicates an invalid character range specification.
var ErrStringInvalidRange = register("string_invalid_range", "invalid character range", nil)

// ErrStringNegativeLength indicates the length cannot be negative.
var ErrStringNegativeLength = register("string_negative_length", "length cannot be negative", nil)

//
// Network/Peer
//

// ErrNetPeerConnection indicates a problem connecting to a network peer.
var ErrNetPeerConnection = register("net_peer_connection", "problem connecting to peer", nil)

// ErrNetReadingRequestBody indicates a problem reading the request body.
var ErrNetReadingRequestBody = register("net_reading_request_body", "problem reading request body", nil)

// ErrNetReadingResponseBody indicates a problem reading the response body.
var ErrNetReadingResponseBody = register("net_reading_response_body", "problem reading response body", nil)

// ErrNetURLJoinPathFailed indicates URL path joining failed.
var ErrNetURLJoinPathFailed = register("net_url_join_path_failed", "failed to join URL path", nil)

//
// Transaction operations
//

// ErrTransactionBeginFailed indicates `transaction begin` failed.
var ErrTransactionBeginFailed = register("transaction_begin_failed", "failed to begin transaction", nil)

// ErrTransactionCommitFailed indicates `transaction commit` failed.
var ErrTransactionCommitFailed = register("transaction_commit_failed", "failed to commit transaction", nil)

// ErrTransactionFailed indicates the transaction failed.
var ErrTransactionFailed = register("transaction_failed", "transaction failed", nil)

// ErrTransactionRollbackFailed indicates transaction rollback failed.
var ErrTransactionRollbackFailed = register("transaction_rollback_failed", "failed to rollback transaction", nil)

//
// Recovery operations
//

// ErrRecoveryFailed indicates the recovery operation failed.
var ErrRecoveryFailed = register("recovery_failed", "recovery failed", nil)

// ErrRecoveryRetryFailed indicates recovery retry operation failed.
var ErrRecoveryRetryFailed = register("recovery_retry_failed", "recovery retry failed", nil)

// ErrRecoveryRetryLimitReached indicates the recovery retry limit was reached.
var ErrRecoveryRetryLimitReached = register("recovery_retry_limit_reached", "recovery retry limit reached", nil)

//
// Retry operations
//

// ErrRetryContextCanceled indicates the retry was canceled due to context cancellation.
var ErrRetryContextCanceled = register("retry_context_canceled", "retry canceled due to context cancellation", nil)

// ErrRetryMaxElapsedTimeReached indicates the maximum elapsed time for retries was reached.
var ErrRetryMaxElapsedTimeReached = register("retry_max_elapsed_time_reached", "maximum elapsed time for retries reached", nil)

// ErrRetryOperationFailed indicates the retry operation failed.
var ErrRetryOperationFailed = register("retry_operation_failed", "retry operation failed", nil)

//
// X509/SPIFFE
//

// ErrSPIFFEEmptyTrustDomain indicates an empty SPIFFE trust domain was provided.
var ErrSPIFFEEmptyTrustDomain = register("spiffe_empty_trust_domain", "empty trust domain", nil)

// ErrSPIFFEFailedToCloseX509Source indicates X509Source close operation failed.
var ErrSPIFFEFailedToCloseX509Source = register("spiffe_failed_to_close_source", "failed to close X509Source", nil)

// ErrSPIFFEFailedToExtractX509SVID indicates X509 SVID extraction failed.
var ErrSPIFFEFailedToExtractX509SVID = register("spiffe_failed_to_extract_x509_svid", "failed to extract X509 SVID", nil)

// ErrSPIFFEInvalidSPIFFEID indicates an invalid SPIFFE ID was provided.
var ErrSPIFFEInvalidSPIFFEID = register("spiffe_invalid_spiffe_id", "invalid SPIFFE ID", nil)

// ErrSPIFFEInvalidTrustDomain indicates an invalid trust domain was provided.
var ErrSPIFFEInvalidTrustDomain = register("spiffe_invalid_trust_domain", "invalid trust domain", nil)

// ErrSPIFFEMultipleTrustDomains indicates multiple trust domains were provided when only one is allowed.
var ErrSPIFFEMultipleTrustDomains = register("spiffe_multiple_trust_domains", "provide a single trust domain", nil)

// ErrSPIFFENilX509Source indicates a nil X509Source was provided.
var ErrSPIFFENilX509Source = register("spiffe_nil_x509_source", "nil X509Source", nil)

// ErrSPIFFENoPeerCertificates indicates no peer certificates were found.
var ErrSPIFFENoPeerCertificates = register("spiffe_no_peer_certificates", "no peer certificates", nil)

// ErrSPIFFEUnableToFetchX509Source indicates unable to fetch X509Source.
var ErrSPIFFEUnableToFetchX509Source = register("spiffe_unable_to_fetch_x509_source", "unable to fetch X509Source", nil)

//
// Errors related to the underlying system
//

var ErrSystemMemLockFailed = register("system_mem_lock_failed", "failed to lock memory", nil)
