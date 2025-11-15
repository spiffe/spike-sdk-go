//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

type ErrorCode string

const ErrAlreadyInitialized = ErrorCode("already_initialized")
const ErrBadInput = ErrorCode("bad_request")
const ErrCreationFailed = ErrorCode("creation_failed")
const ErrEmptyPayload = ErrorCode("empty_payload")
const ErrInternal = ErrorCode("internal_error")
const ErrLowEntropy = ErrorCode("low_entropy")
const ErrNotAlive = ErrorCode("not_alive")
const ErrNotFound = ErrorCode("not_found")
const ErrNotReady = ErrorCode("not_ready")
const ErrServerFault = ErrorCode("server_fault")
const ErrSuccess = ErrorCode("success")
const ErrUnauthorized = ErrorCode("unauthorized")
