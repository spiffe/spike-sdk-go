//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

type ErrorCode string

const ErrAlreadyInitialized = ErrorCode("already_initialized")
const ErrBadInput = ErrorCode("bad_request")
const ErrCreationFailed = ErrorCode("creation_failed")
const ErrDeletionFailed = ErrorCode("deletion_failed")
const ErrDeletionSuccess = ErrorCode("deletion_success")
const ErrEmptyPayload = ErrorCode("empty_payload")
const ErrInternal = ErrorCode("internal_error")
const ErrLowEntropy = ErrorCode("low_entropy")
const ErrNotAlive = ErrorCode("not_alive")
const ErrNotFound = ErrorCode("not_found")
const ErrNotReady = ErrorCode("not_ready")
const ErrQueryFailure = ErrorCode("query_failed")
const ErrServerFault = ErrorCode("server_fault")
const ErrSuccess = ErrorCode("success")
const ErrUnauthorized = ErrorCode("unauthorized")
const ErrShamirNotEnoughShards = ErrorCode("shamir_not_enough_shards")
const ErrShamirDuplicateIndex = ErrorCode("shamir_duplicate_index")
const ErrShamirNilShard = ErrorCode("shamir_nil_shard")
const ErrShamirEmptyShard = ErrorCode("shamir_empty_shard")
const ErrShamirInvalidIndex = ErrorCode("shamir_invalid_index")
