//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

type ErrorCode string

const ErrAlreadyInitialized = ErrorCode("already_initialized")
const ErrBadInput = ErrorCode("bad_request")
const ErrCreationFailed = ErrorCode("creation_failed")
const ErrCryptoCipherNotAvailable = ErrorCode("crypto_cipher_not_available")
const ErrCryptoCipherVerificationSuccess = ErrorCode("crypto_cipher_verification_success")
const ErrCryptoDecryptionFailed = ErrorCode("crypto_decryption_failed")
const ErrCryptoEncryptionFailed = ErrorCode("crypto_encryption_failed")
const ErrCryptoFailedToReadNonce = ErrorCode("crypto_failed_to_read_nonce")
const ErrCryptoFailedToReadVersion = ErrorCode("crypto_failed_to_read_version")
const ErrCryptoNonceGenerationFailed = ErrorCode("crypto_nonce_generation_failed")
const ErrDeletionFailed = ErrorCode("deletion_failed")
const ErrDeletionSuccess = ErrorCode("deletion_success")
const ErrEmptyPayload = ErrorCode("empty_payload")
const ErrFound = ErrorCode("found")
const ErrInternal = ErrorCode("internal_error")
const ErrLowEntropy = ErrorCode("low_entropy")
const ErrNotAlive = ErrorCode("not_alive")
const ErrNotFound = ErrorCode("not_found")
const ErrNotReady = ErrorCode("not_ready")
const ErrQueryFailure = ErrorCode("query_failed")
const ErrServerFault = ErrorCode("server_fault")
const ErrShamirDuplicateIndex = ErrorCode("shamir_duplicate_index")
const ErrShamirEmptyShard = ErrorCode("shamir_empty_shard")
const ErrShamirInvalidIndex = ErrorCode("shamir_invalid_index")
const ErrShamirNilShard = ErrorCode("shamir_nil_shard")
const ErrShamirNotEnoughShards = ErrorCode("shamir_not_enough_shards")
const ErrSuccess = ErrorCode("success")
const ErrUnauthorized = ErrorCode("unauthorized")
const ErrUndeleteFailed = ErrorCode("undeletion_failed")
const ErrUndeleteSuccess = ErrorCode("undeletion_success")
const ErrNilContext = ErrorCode("nil_context")
const ErrRootKeySkipCreationForInMemoryMode = ErrorCode("root_key_skip_creation_for_in_memory_mode")
const ErrRootKeyEmpty = ErrorCode("root_key_empty")
const ErrResultSetFailedToLoad = ErrorCode("result_set_failed_to_load")
