//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package strings provides utility functions for string operations.
package strings

// MaybeError converts an error to its string representation if the error is
// not nil. If the error is nil, it returns an empty string. This is useful
// for cases where you need a string representation of an error that may or
// may not exist.
func MaybeError(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
