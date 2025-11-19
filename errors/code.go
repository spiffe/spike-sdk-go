//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

type ErrorCode string

// General error codes.
const errCodeBadRequest = ErrorCode("bad_request")
const errCodeEmptyPayload = ErrorCode("empty_payload")
const errCodeFound = ErrorCode("found")
const errCodeInternal = ErrorCode("internal_error")
const errCodeNilContext = ErrorCode("nil_context")
const errCodeServerFault = ErrorCode("server_fault")
const errCodeSuccess = ErrorCode("success")
const errCodeNotFound = ErrorCode("not_found")
const errCodePostFailed = ErrorCode("post_failed")

// Entity operations
const errCodeEntityExists = ErrorCode("entity_exists")
const errCodeEntityInvalid = ErrorCode("entity_invalid")
const errCodeEntityNotFound = ErrorCode("entity_not_found")

// State operations
const errCodeAlreadyInitialized = ErrorCode("already_initialized")
const errCodeInitializationFailed = ErrorCode("initialization_failed")
const errCodeNotAlive = ErrorCode("not_alive")
const errCodeNotReady = ErrorCode("not_ready")

// Policy/RBAC/ABAC
const errCodeUnauthorized = ErrorCode("unauthorized")

// CRUD operations
const errCodeDeletionFailed = ErrorCode("deletion_failed")
const errCodeDeletionSuccess = ErrorCode("deletion_success")
const errCodeUndeleteFailed = ErrorCode("undeletion_failed")
const errCodeUndeleteSuccess = ErrorCode("undeletion_success")
const errCodeCreationFailed = ErrorCode("creation_failed")

// Root key management
const errCodeRootKeyEmpty = ErrorCode("root_key_empty")
const errCodeRootKeyMissing = ErrorCode("root_key_missing")
const errCodeRootKeyNotEmpty = ErrorCode("root_key_not_empty")
const errCodeRootKeySetSuccess = ErrorCode("root_key_set_success")
const errCodeRootKeySkipCreationForInMemoryMode = ErrorCode("root_key_skip_creation_for_in_memory_mode")
const errCodeRootKeyUpdateSkippedKeyEmpty = ErrorCode("root_key_update_skipped_key_empty")

// Shamir-related
const errCodeShamirDuplicateIndex = ErrorCode("shamir_duplicate_index")
const errCodeShamirEmptyShard = ErrorCode("shamir_empty_shard")
const errCodeShamirInvalidIndex = ErrorCode("shamir_invalid_index")
const errCodeShamirNilShard = ErrorCode("shamir_nil_shard")
const errCodeShamirNotEnoughShards = ErrorCode("shamir_not_enough_shards")
const errCodeShamirReconstructionFailed = ErrorCode("shamir_reconstruction_failed")

// Crypto operations
const errCodeLowEntropy = ErrorCode("low_entropy")
const errCodeCryptoCipherNotAvailable = ErrorCode("crypto_cipher_not_available")
const errCodeCryptoCipherVerificationSuccess = ErrorCode("crypto_cipher_verification_success")
const errCodeCryptoDecryptionFailed = ErrorCode("crypto_decryption_failed")
const errCodeCryptoEncryptionFailed = ErrorCode("crypto_encryption_failed")
const errCodeCryptoFailedToReadNonce = ErrorCode("crypto_failed_to_read_nonce")
const errCodeCryptoFailedToReadVersion = ErrorCode("crypto_failed_to_read_version")
const errCodeCryptoNonceGenerationFailed = ErrorCode("crypto_nonce_generation_failed")
const errCodeCryptoFailedToCreateCipher = ErrorCode("crypto_failed_to_create_cipher")
const errCodeCryptoFailedToCreateGCM = ErrorCode("crypto_failed_to_create_gcm")
const errCodeCryptoInvalidEncryptionKeyLength = ErrorCode("crypto_invalid_encryption_key_length")

// Store configuration
const errCodeStoreInvalidConfiguration = ErrorCode("store_invalid_configuration")
const errCodeStoreInvalidEncryptionKey = ErrorCode("store_invalid_encryption_key")

// Store operations
const errCodeStoreLoadFailed = ErrorCode("store_load_failed")
const errCodeStoreQueryFailed = ErrorCode("store_query_failed")
const errCodeStoreSaveFailed = ErrorCode("store_save_failed")
const errCodeStoreQueryFailure = ErrorCode("store_query_failure")
const errCodeStoreResultSetFailedToLoad = ErrorCode("store_result_set_failed_to_load")
const errCodeStoreVersionNotFound = ErrorCode("store_version_not_found")
const errCodeStoreItemSoftDeleted = ErrorCode("store_item_soft_deleted")
const errCodeStoreInvalidVersion = ErrorCode("store_invalid_version")

// Filesystem operations
const errCodeDirectoryCreationFailed = ErrorCode("directory_creation_failed")
const errCodeFSFailedToCheckDirectory = ErrorCode("fs_failed_to_check_directory")
const errCodeFSFailedToCreateDirectory = ErrorCode("fs_failed_to_create_directory")
const errCodeFSFailedToResolvePath = ErrorCode("fs_failed_to_resolve_path")
const errCodeFSFileIsNotADirectory = ErrorCode("fs_file_is_not_a_directory")
const errCodeFSInvalidDirectory = ErrorCode("fs_invalid_directory")
const errCodeFSParentDirectoryDoesNotExist = ErrorCode("fs_parent_directory_does_not_exist")
const errCodeFSPathCannotBeEmpty = ErrorCode("fs_path_cannot_be_empty")
const errCodeFSPathRestricted = ErrorCode("fs_path_restricted")

// File I/O
const errCodeFileCloseFailed = ErrorCode("file_close_failed")
const errCodeFileOpenFailed = ErrorCode("file_open_failed")

// Input/Output
const errCodeInvalidInput = ErrorCode("invalid_input")
const errCodeInvalidPermission = ErrorCode("invalid_permission")
const errCodeMarshalFailure = ErrorCode("marshal_failure")
const errCodeParseFailure = ErrorCode("parse_failure")
const errCodeReadFailure = ErrorCode("read_failure")
const errCodeUnmarshalFailure = ErrorCode("unmarshal_failure")

// Network/Peer
const errCodePeerConnection = ErrorCode("peer_connection")
const errCodeReadingResponseBody = ErrorCode("reading_response_body")

// Transaction operations
const errCodeTransactionBeginFailed = ErrorCode("transaction_begin_failed")
const errCodeTransactionCommitFailed = ErrorCode("transaction_commit_failed")
const errCodeTransactionFailed = ErrorCode("transaction_failed")
const errCodeTransactionRollbackFailed = ErrorCode("transaction_rollback_failed")

// Recovery operations
const errCodeRecoveryRetryFailed = ErrorCode("recovery_retry_failed")

// X509/SPIFFE
const errCodeNilX509Source = ErrorCode("nil_x509_source")
