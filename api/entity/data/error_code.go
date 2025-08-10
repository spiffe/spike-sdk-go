//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

type ErrorCode string

const ErrBadInput = ErrorCode("bad_request")
const ErrServerFault = ErrorCode("server_fault")
const ErrUnauthorized = ErrorCode("unauthorized")
const ErrInternal = ErrorCode("internal_error")
const ErrLowEntropy = ErrorCode("low_entropy")
const ErrNotReady = ErrorCode("not_ready")
const ErrNotAlive = ErrorCode("not_alive")
const ErrAlreadyInitialized = ErrorCode("already_initialized")
const ErrNotFound = ErrorCode("not_found")
const ErrSuccess = ErrorCode("success")
