//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// PolicyPutRequest for policy creation.
type PolicyPutRequest struct {
	Name            string                  `json:"name"`
	SPIFFEIDPattern string                  `json:"spiffeidPattern"`
	PathPattern     string                  `json:"pathPattern"`
	Permissions     []data.PolicyPermission `json:"permissions"`
}

// PolicyPutResponse for policy creation.
type PolicyPutResponse struct {
	ID  string              `json:"id,omitempty"`
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyPutResponse) Success() PolicyPutResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return p
}
func (p PolicyPutResponse) NotFound() PolicyPutResponse {
	p.Err = sdkErrors.ErrNotFound.Code
	return p
}
func (p PolicyPutResponse) BadRequest() PolicyPutResponse {
	p.Err = sdkErrors.ErrBadRequest.Code
	return p
}
func (p PolicyPutResponse) Unauthorized() PolicyPutResponse {
	p.Err = sdkErrors.ErrAccessUnauthorized.Code
	return p
}
func (p PolicyPutResponse) Internal() PolicyPutResponse {
	p.Err = sdkErrors.ErrInternal.Code
	return p
}

// PolicyReadRequest to read a policy.
type PolicyReadRequest struct {
	ID string `json:"id"`
}

// PolicyReadResponse to read a policy.
type PolicyReadResponse struct {
	data.Policy
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (s PolicyReadResponse) Success() PolicyReadResponse {
	s.Err = sdkErrors.ErrSuccess.Code
	return s
}
func (s PolicyReadResponse) NotFound() PolicyReadResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s PolicyReadResponse) BadRequest() PolicyReadResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s PolicyReadResponse) Unauthorized() PolicyReadResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s PolicyReadResponse) Internal() PolicyReadResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}

// PolicyDeleteRequest to delete a policy.
type PolicyDeleteRequest struct {
	ID string `json:"id"`
}

// PolicyDeleteResponse to delete a policy.
type PolicyDeleteResponse struct {
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyDeleteResponse) Success() PolicyDeleteResponse {
	p.Err = sdkErrors.ErrSuccess.Code
	return p
}
func (s PolicyDeleteResponse) NotFound() PolicyDeleteResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s PolicyDeleteResponse) BadRequest() PolicyDeleteResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s PolicyDeleteResponse) Unauthorized() PolicyDeleteResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s PolicyDeleteResponse) Internal() PolicyDeleteResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}

// PolicyListRequest to list policies.
type PolicyListRequest struct {
	SPIFFEIDPattern string `json:"spiffeidPattern"`
	PathPattern     string `json:"pathPattern"`
}

// PolicyListResponse to list policies.
type PolicyListResponse struct {
	Policies []data.Policy       `json:"policies"`
	Err      sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyListResponse) Success() PolicyListResponse {
	p.Err = sdkErrors.ErrSuccess.Code
	return p
}
func (s PolicyListResponse) NotFound() PolicyListResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s PolicyListResponse) BadRequest() PolicyListResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s PolicyListResponse) Unauthorized() PolicyListResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s PolicyListResponse) Internal() PolicyListResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}

// PolicyAccessCheckRequest to validate policy access.
type PolicyAccessCheckRequest struct {
	SPIFFEID string `json:"spiffeid"`
	Path     string `json:"path"`
	Action   string `json:"action"`
}

// PolicyAccessCheckResponse to validate policy access.
type PolicyAccessCheckResponse struct {
	Allowed          bool                `json:"allowed"`
	MatchingPolicies []string            `json:"matchingPolicies"`
	Err              sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyAccessCheckResponse) Success() PolicyAccessCheckResponse {
	p.Err = sdkErrors.ErrSuccess.Code
	return p
}
func (s PolicyAccessCheckResponse) NotFound() PolicyAccessCheckResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s PolicyAccessCheckResponse) BadRequest() PolicyAccessCheckResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s PolicyAccessCheckResponse) Unauthorized() PolicyAccessCheckResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s PolicyAccessCheckResponse) Internal() PolicyAccessCheckResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}
