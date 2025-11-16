//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

import "errors"

var ErrAlreadyInitialized = errors.New("already initialized")
var ErrCreationFailed = errors.New("creation failed")
var ErrDeletionFailed = errors.New("deletion failed")
var ErrFound = errors.New("found")
var ErrInvalidInput = errors.New("invalid input")
var ErrInvalidPermission = errors.New("invalid permission")
var ErrMarshalFailure = errors.New("failed to marshal response body")
var ErrMissingRootKey = errors.New("missing root key")
var ErrNilX509Source = errors.New("nil X509Source")
var ErrNotFound = errors.New("not found")
var ErrParseFailure = errors.New("failed to parse request body")
var ErrPeerConnection = errors.New("problem connecting to peer")
var ErrQueryFailure = errors.New("failed to query for the requested data")
var ErrReadFailure = errors.New("failed to read request body")
var ErrReadingResponseBody = errors.New("problem reading response body")
var ErrUnauthorized = errors.New("unauthorized")
