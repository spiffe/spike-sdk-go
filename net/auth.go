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
	peerSPIFFEID, idErr := ExtractPeerSPIFFEIDAndRespondOnFail(
		w, r, unauthorizedRes,
	)
	if idErr != nil {
		return idErr
	}

	if authErr := AuthorizedAndRespondOnFailWithPredicate(
		func(peerSPIFFEID string) bool {
			return accessCheck(
				peerSPIFFEID, policyCheck,
			)
		},
		peerSPIFFEID.String(),
		unauthorizedRes, w,
	); authErr != nil {
		return authErr
	}
	return nil
}

// AuthorizeAndRespondOnFailNoPolicy performs access-based authorization for an
// HTTP request without policy validation and automatically responds with an
// unauthorized error if the check fails.
//
// This is a simplified version of AuthorizeAndRespondOnFail that uses
// predicate.AllowAllPolicies, bypassing policy-level checks while still
// performing SPIFFE ID extraction and access validation. Use this when you need
// to verify the caller's identity but don't require fine-grained policy control.
//
// Type Parameters:
//   - U: The type of the unauthorized response body to send on failure
//
// Parameters:
//   - unauthorizedRes: The response body to return if authorization fails
//   - accessCheck: Function that validates access for a SPIFFE ID (policy
//     parameter will receive AllowAllPolicies)
//   - w: The HTTP response writer
//   - r: The HTTP request containing the peer's TLS credentials
//
// Returns:
//   - *sdkErrors.SDKError: nil if authorized, otherwise an error describing the
//     authorization failure
func AuthorizeAndRespondOnFailNoPolicy[U any](
	unauthorizedRes U,
	accessCheck predicate.WithPolicyAccessChecker,
	w http.ResponseWriter, r *http.Request,
) *sdkErrors.SDKError {
	peerSPIFFEID, idErr := ExtractPeerSPIFFEIDAndRespondOnFail(
		w, r, unauthorizedRes,
	)
	if idErr != nil {
		return idErr
	}

	if authErr := AuthorizedAndRespondOnFailWithPredicate(
		func(peerSPIFFEID string) bool {
			return accessCheck(peerSPIFFEID, predicate.AllowAllPolicies)
		},
		peerSPIFFEID.String(),
		unauthorizedRes, w,
	); authErr != nil {
		return authErr
	}
	return nil
}

// AuthorizedAndRespondOnFailWithPredicate extracts the peer SPIFFE ID from an
// HTTP request and validates it using the provided predicate function. If the
// SPIFFE ID extraction fails or the predicate returns false, it sends an
// HTTP 401 Unauthorized response with the provided failure response body.
//
// This function combines SPIFFE ID extraction with custom authorization logic,
// making it useful for route handlers that need to verify the caller's identity
// against specific criteria (e.g., checking if the caller is a known service,
// validating trust domain membership, or matching against an allowlist).
//
// Parameters:
//   - predicateFn: A function that takes a SPIFFE ID string and returns true
//     if the caller is authorized, false otherwise
//   - failureResponse: The response object to send if authorization fails
//   - w: The HTTP response writer for error responses
//   - r: The incoming HTTP request containing the peer's SPIFFE ID
//
// Returns:
//   - *sdkErrors.SDKError: nil if the SPIFFE ID was successfully extracted and
//     the predicate returned true; otherwise returns ErrAccessUnauthorized
//     (potentially wrapping additional errors from response writing)
//
// Example usage:
//
//	isAuthorizedService := func(spiffeID string) bool {
//	    return strings.HasPrefix(spiffeID, "spiffe://example.org/service/")
//	}
//
//	if err := net.RespondUnauthorizedOnPredicateFail(
//	    isAuthorizedService,
//	    reqres.SecretGetResponse{Err: data.ErrUnauthorized},
//	    w, r,
//	); err != nil {
//	    return err
//	}
func AuthorizedAndRespondOnFailWithPredicate(
	predicateFn predicate.Predicate, peerSPIFFEID string,
	failureResponse any,
	w http.ResponseWriter,
) *sdkErrors.SDKError {
	if !predicateFn(peerSPIFFEID) {
		failErr := Fail(
			failureResponse, w,
			http.StatusUnauthorized,
		)
		if failErr != nil {
			return sdkErrors.ErrAccessUnauthorized.Wrap(failErr)
		}
		return sdkErrors.ErrAccessUnauthorized.Clone()
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
