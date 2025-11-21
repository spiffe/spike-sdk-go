//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

//
// General error codes
//

var ErrGeneralFailure = register("gen_general_failure", "general failure", nil)
var ErrNilContext = register("gen_nil_context", "nil context", nil)

//
// Cluster operations
//

var ErrK8sReconciliationFailed = register("k8s_reconciliation_failed", "reconciliation failed", nil)

//
// API/HTTP operations
//

var ErrAPIBadRequest = register("api_bad_request", "bad request", nil)
var ErrAPIEmptyPayload = register("api_empty_payload", "empty payload", nil)
var ErrAPIFound = register("api_found", "found", nil)
var ErrAPIInternal = register("api_internal_error", "internal error", nil)
var ErrAPINotFound = register("api_not_found", "not found", nil)
var ErrAPIPostFailed = register("api_post_failed", "post failed", nil)
var ErrAPIResponseCodeInvalid = register("api_response_code_invalid", "invalid API response code", nil)
var ErrAPIServerFault = register("api_server_fault", "server fault", nil)
var ErrAPISuccess = register("api_success", "success", nil)

//
// Entity operations
//

var ErrEntityDeleted = register("entity_deleted", "entity marked as deleted", nil)
var ErrEntityExists = register("entity_exists", "entity already exists", nil)
var ErrEntityInvalid = register("entity_invalid", "entity is invalid", nil)
var ErrEntityLoadFailed = register("entity_load_failed", "failed to load entity", nil)
var ErrEntityNotFound = register("entity_not_found", "entity not found", nil)
var ErrEntityQueryFailed = register("entity_query_failed", "failed to query entities", nil)
var ErrEntitySaveFailed = register("entity_save_failed", "failed to save entity", nil)
var ErrEntityVersionInvalid = register("entity_version_invalid", "invalid version", nil)
var ErrEntityVersionNotFound = register("entity_version_not_found", "version not found", nil)

//
// State operations
//

var ErrStateAlreadyInitialized = register("state_already_initialized", "already initialized", nil)
var ErrStateInitializationFailed = register("state_initialization_failed", "initialization failed", nil)
var ErrStateNotAlive = register("state_not_alive", "not alive", nil)
var ErrStateNotReady = register("state_not_ready", "not ready", nil)

//
// Policy/RBAC/ABAC
//

var ErrAccessInvalidPermission = register("access_invalid_permission", "invalid permission", nil)
var ErrAccessUnauthorized = register("access_unauthorized", "unauthorized", nil)

//
// CRUD operations
//

var ErrObjectCreationFailed = register("object_creation_failed", "creation failed", nil)
var ErrObjectDeletionFailed = register("object_deletion_failed", "deletion failed", nil)
var ErrObjectDeletionSuccess = register("object_deletion_success", "deletion success", nil)
var ErrObjectUndeletionFailed = register("object_undeletion_failed", "undeletion failed", nil)
var ErrObjectUndeletionSuccess = register("object_undeletion_success", "undeletion success", nil)

//
// Root key management
//

var ErrRootKeyEmpty = register("root_key_empty", "root key empty", nil)
var ErrRootKeyMissing = register("root_key_missing", "root key missing", nil)
var ErrRootKeyNotEmpty = register("root_key_not_empty", "root key not empty", nil)
var ErrRootKeySetSuccess = register("root_key_set_success", "root key set success", nil)
var ErrRootKeySkipCreationForInMemoryMode = register("root_key_skip_creation_for_in_memory_mode", "root key skip creation for in memory mode", nil)
var ErrRootKeyUpdateSkippedKeyEmpty = register("root_key_update_skipped_key_empty", "root key update skipped key empty", nil)

//
// Shamir-related
//

var ErrShamirDuplicateIndex = register("shamir_duplicate_index", "shamir duplicate index", nil)
var ErrShamirEmptyShard = register("shamir_empty_shard", "shamir empty shard", nil)
var ErrShamirInvalidIndex = register("shamir_invalid_index", "shamir invalid index", nil)
var ErrShamirNilShard = register("shamir_nil_shard", "shamir nil shard", nil)
var ErrShamirNotEnoughShards = register("shamir_not_enough_shards", "shamir not enough shards", nil)
var ErrShamirReconstructionFailed = register("shamir_reconstruction_failed", "shamir reconstruction failed", nil)

//
// Crypto operations
//

var ErrCryptoCipherNotAvailable = register("crypto_cipher_not_available", "cipher not available", nil)
var ErrCryptoCipherVerificationFailed = register("crypto_cipher_verification_failed", "cipher verification failed", nil)
var ErrCryptoCipherVerificationSuccess = register("crypto_cipher_verification_success", "cipher verification success", nil)
var ErrCryptoDecryptionFailed = register("crypto_decryption_failed", "decryption failed", nil)
var ErrCryptoEncryptionFailed = register("crypto_encryption_failed", "encryption failed", nil)
var ErrCryptoFailedToCreateCipher = register("crypto_failed_to_create_cipher", "failed to create cipher", nil)
var ErrCryptoFailedToCreateGCM = register("crypto_failed_to_create_gcm", "failed to create GCM", nil)
var ErrCryptoFailedToReadNonce = register("crypto_failed_to_read_nonce", "failed to read nonce", nil)
var ErrCryptoFailedToReadVersion = register("crypto_failed_to_read_version", "failed to read version", nil)
var ErrCryptoInvalidEncryptionKeyLength = register("crypto_invalid_encryption_key_length", "invalid encryption key length", nil)
var ErrCryptoLowEntropy = register("crypto_low_entropy", "low entropy", nil)
var ErrCryptoNonceGenerationFailed = register("crypto_nonce_generation_failed", "nonce generation failed", nil)
var ErrCryptoRandomGenerationFailed = register("crypto_random_generation_failed", "random generation failed", nil)

//
// Backing store infrastructure (internal store operations)
//

var ErrStoreInvalidConfiguration = register("store_invalid_configuration", "invalid store configuration", nil)
var ErrStoreInvalidEncryptionKey = register("store_invalid_encryption_key", "invalid store encryption key", nil)
var ErrStoreResultSetFailedToLoad = register("store_result_set_failed_to_load", "result set failed to load", nil)

//
// Filesystem operations
//

var ErrFSDirectoryCreationFailed = register("fs_directory_creation_failed", "directory creation failed", nil)
var ErrFSFailedToCheckDirectory = register("fs_failed_to_check_directory", "failed to check directory", nil)
var ErrFSFailedToCreateDirectory = register("fs_failed_to_create_directory", "failed to create directory", nil)
var ErrFSFailedToResolvePath = register("fs_failed_to_resolve_path", "failed to resolve filesystem path", nil)
var ErrFSFileCloseFailed = register("fs_file_close_failed", "file close failed", nil)
var ErrFSFileIsNotADirectory = register("fs_file_is_not_a_directory", "file is not a directory", nil)
var ErrFSFileOpenFailed = register("fs_file_open_failed", "file open failed", nil)
var ErrFSInvalidDirectory = register("fs_invalid_directory", "invalid directory", nil)
var ErrFSParentDirectoryDoesNotExist = register("fs_parent_directory_does_not_exist", "parent directory does not exist", nil)
var ErrFSPathCannotBeEmpty = register("fs_path_cannot_be_empty", "filesystem path cannot be empty", nil)
var ErrFSPathRestricted = register("fs_path_restricted", "filesystem path is restricted for security reasons", nil)
var ErrFSStreamCloseFailed = register("fs_stream_close_failed", "stream close failed", nil)
var ErrFSStreamOpenFailed = register("fs_stream_open_failed", "stream open failed", nil)

//
// Data Processing
//

var ErrDataInvalidInput = register("data_invalid_input", "invalid input", nil)
var ErrDataMarshalFailure = register("data_marshal_failure", "failed to marshal response body", nil)
var ErrDataParseFailure = register("data_parse_failure", "failed to parse request body", nil)
var ErrDataReadFailure = register("data_read_failure", "failed to read request body", nil)
var ErrDataUnmarshalFailure = register("data_unmarshal_failure", "failed to unmarshal request body", nil)

//
// String/Template operations
//

var ErrStringEmptyCharacterClass = register("string_empty_character_class", "empty character class", nil)
var ErrStringInvalidRange = register("string_invalid_range", "invalid character range", nil)
var ErrStringEmptyCharacterSet = register("string_empty_character_set", "character class resulted in empty set", nil)
var ErrStringInvalidLength = register("string_invalid_length", "invalid length specification", nil)
var ErrStringNegativeLength = register("string_negative_length", "length cannot be negative", nil)

//
// Network/Peer
//

var ErrNetPeerConnection = register("net_peer_connection", "problem connecting to peer", nil)
var ErrNetReadingRequestBody = register("net_reading_request_body", "problem reading request body", nil)
var ErrNetReadingResponseBody = register("net_reading_response_body", "problem reading response body", nil)
var ErrNetURLJoinPathFailed = register("net_url_join_path_failed", "failed to join URL path", nil)

//
// Transaction operations
//

var ErrTransactionBeginFailed = register("transaction_begin_failed", "failed to begin transaction", nil)
var ErrTransactionCommitFailed = register("transaction_commit_failed", "failed to commit transaction", nil)
var ErrTransactionFailed = register("transaction_failed", "transaction failed", nil)
var ErrTransactionRollbackFailed = register("transaction_rollback_failed", "failed to rollback transaction", nil)

//
// Recovery operations
//

var ErrRecoveryRetryFailed = register("recovery_retry_failed", "recovery retry failed", nil)
var ErrRecoveryRetryLimitReached = register("recovery_retry_limit_reached", "recovery retry limit reached", nil)
var ErrRecoveryFailed = register("recovery_failed", "recovery failed", nil)

//
// Retry operations
//

var ErrRetryMaxElapsedTimeReached = register("retry_max_elapsed_time_reached", "maximum elapsed time for retries reached", nil)
var ErrRetryContextCanceled = register("retry_context_canceled", "retry canceled due to context cancellation", nil)
var ErrRetryOperationFailed = register("retry_operation_failed", "retry operation failed", nil)

//
// X509/SPIFFE
//

var ErrSPIFFEEmptyTrustDomain = register("spiffe_empty_trust_domain", "empty trust domain", nil)
var ErrSPIFFEFailedToCreateX509Source = register("spiffe_failed_to_create_x509_source", "failed to create X509Source", nil)
var ErrSPIFFEFailedToExtractX509SVID = register("spiffe_failed_to_extract_x509_svid", "failed to extract X509 SVID", nil)
var ErrSPIFFEMultipleTrustDomains = register("spiffe_multiple_trust_domains", "provide a single trust domain", nil)
var ErrSPIFFENilX509Source = register("spiffe_nil_x509_source", "nil X509Source", nil)
var ErrSPIFFENoPeerCertificates = register("spiffe_no_peer_certificates", "no peer certificates", nil)
var ErrSPIFFEUnableToFetchX509Source = register("spiffe_unable_to_fetch_x509_source", "unable to fetch X509Source", nil)
var ErrSPIFFEFailedToCloseX509Source = register("spiffe_failed_to_close_source", "failed to close X509Source", nil)
