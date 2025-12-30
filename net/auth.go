//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"fmt"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// AuthorizerWithPredicate creates a TLS authorizer that validates SPIFFE IDs
// using the provided predicate function.
//
// The authorizer checks each connecting peer's SPIFFE ID against the predicate.
// If the predicate returns true, the connection is authorized. If false, the
// connection is rejected with ErrAccessUnauthorized.
//
// Parameters:
//   - predicate: Function that takes a SPIFFE ID string and returns true to
//     allow the connection, false to reject it
//
// Returns:
//   - tlsconfig.Authorizer: A TLS authorizer that can be used with mTLS configs
//
// Example:
//
//	// Allow only production namespace
//	authorizer := AuthorizerWithPredicate(func(id string) bool {
//	    return strings.Contains(id, "/ns/production/")
//	})
func AuthorizerWithPredicate(predicate func(string) bool) tlsconfig.Authorizer {
	return tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if predicate(id.String()) {
			return nil
		}

		failErr := sdkErrors.ErrAccessUnauthorized.Clone()
		failErr.Msg = fmt.Sprintf("unauthorized spiffe id: '%s'", id.String())

		return failErr
	})
}
