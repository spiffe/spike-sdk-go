//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

import "errors"

// FOR FUTURE IMPLEMENTATION (see ADR-0028)

// 1. All SDK and non-cli SPIKE errors shall be replaced with sentinel errors
//    that stem from the `SDKEror` construct down below.
// 2. All comparisons shall be done with the `Is()` function.
// 3. Any missing context shall be a part of the `Msg` of the error.
// 4. etc.
// 5. import sdk errors as `sdkErrors` consistently to make search/find easier.

type SDKError struct {
	// Need to implement Wrap and Unwrap methods

	Code    ErrorCode
	Msg     string
	Wrapped error
}

func (e *SDKError) Error() string {
	// Maybe also include the wrapped errors recursively?
	return e.Msg
}

func (e *SDKError) Is(target error) bool {
	// Usage: if errors.Is(err, &sdk.Error{Code: sdk.ErrNotReady}) {

	var t *SDKError
	if errors.As(target, &t) {
		return e.Code == t.Code
	}
	return false
}
