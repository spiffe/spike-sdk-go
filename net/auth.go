//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"fmt"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/spike-sdk-go/predicate"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// AuthorizeAndRespondOnFail performs policy-based authorization for an HTTP
// request and automatically responds with an unauthorized error if the check
// fails.
//
// This generic function combines SPIFFE ID extraction and policy-based access
// control into a single authorization step. It first extracts the peer's
// SPIFFE ID from the TLS connection, then validates access using the provided
// access and policy checkers.
//
// Type Parameters:
//   - U: The type of the unauthorized response body to send on failure
//
// Parameters:
//   - unauthorizedRes: The response body to return if authorization fails
//   - accessCheck: Function that validates access for a SPIFFE ID using a
//     policy
//   - policyCheck: Function that checks if a SPIFFE ID has required permissions
//   - w: The HTTP response writer
//   - r: The HTTP request containing the peer's TLS credentials
//
// Returns:
//   - *sdkErrors.SDKError: nil if authorized, otherwise an error describing the
//     authorization failure
func AuthorizeAndRespondOnFail[U any](
	unauthorizedRes U,
	accessCheck predicate.WithPolicyAccessChecker,
	policyCheck predicate.PolicyAccessChecker,
	w http.ResponseWriter, r *http.Request,
) *sdkErrors.SDKError {
	if _, idErr := ExtractPeerSPIFFEIDAndRespondOnFail(
		w, r, unauthorizedRes,
	); idErr != nil {
		return idErr
	}

	if authErr := RespondUnauthorizedOnPredicateFail(
		func(peerSPIFFEID string) bool {
			return accessCheck(
				peerSPIFFEID, policyCheck,
			)
		},
		unauthorizedRes, w, r,
	); authErr != nil {
		return authErr
	}
	return nil
}

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
