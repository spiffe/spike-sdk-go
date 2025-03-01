//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import "errors"

var (
	ErrVersionNotFound = errors.New("version not found")
	ErrItemNotFound    = errors.New("item not found")
	ErrItemSoftDeleted = errors.New("item marked as deleted")
	ErrInvalidVersion  = errors.New("invalid version")
)
