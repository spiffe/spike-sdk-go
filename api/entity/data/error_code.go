//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

type ErrorCode string

const ErrBadInput = ErrorCode("bad_request")

// ErrBadInput can be used for all of these.
//const ErrInvalidName = ErrorCode("invalid_name")
//const ErrInvalidSpiffeIdPattern = ErrorCode("invalid_spiffe_id_pattern")
//const ErrInvalidPermission = ErrorCode("invalid_permission")
//const ErrInvalidPathPattern = ErrorCode("invalid_path_pattern")

// Use ErrInternal instead.
// const ErrServerFault = ErrorCode("server_fault")

const ErrInternal = ErrorCode("internal_error")

const ErrUnauthorized = ErrorCode("unauthorized")
const ErrLowEntropy = ErrorCode("low_entropy")

const ErrNotFound = ErrorCode("not_found")
const ErrSuccess = ErrorCode("success")
