//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"io"
	"net/http"
)

// body reads and returns the entire response body from an HTTP response.
// The response body is read completely and returned as a byte slice.
//
// Parameters:
//   - r: the HTTP response to read from
//
// Returns:
//   - []byte: the complete response body
//   - error: non-nil if reading the response body fails
func body(r *http.Response) (bod []byte, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
