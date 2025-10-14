//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

import "errors"

var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrParseFailure = errors.New("failed to parse request body")
var ErrReadFailure = errors.New("failed to read request body")
var ErrMarshalFailure = errors.New("failed to marshal response body")
var ErrAlreadyInitialized = errors.New("already initialized")
var ErrMissingRootKey = errors.New("missing root key")
var ErrInvalidInput = errors.New("invalid input")
var ErrInvalidPermission = errors.New("invalid permission")
var ErrPeerConnection = errors.New("problem connecting to peer")
var ErrReadingResponseBody = errors.New("problem reading response body")
var ErrNilX509Source = errors.New("nil X509Source")
