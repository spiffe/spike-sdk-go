//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

// ErrorResponder defines an interface for response types that can generate
// standard error responses. All SDK response types in the reqres package
// implement this interface through their NotFound() and Internal() methods.
type ErrorResponder[T any] interface {
	NotFound() T
	Internal() T
}

// InternalErrorResponder defines an interface for response types that can
// generate internal error responses. This is a subset of ErrorResponder for
// cases where only internal errors are possible (no "not found" scenario).
type InternalErrorResponder[T any] interface {
	Internal() T
}
