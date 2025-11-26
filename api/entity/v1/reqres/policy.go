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

func (r PolicyPutResponse) Success() PolicyPutResponse {
	r.Err = ""
	return r
}
func (r PolicyPutResponse) NotFound() PolicyPutResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r PolicyPutResponse) BadRequest() PolicyPutResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r PolicyPutResponse) Unauthorized() PolicyPutResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r PolicyPutResponse) Internal() PolicyPutResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r PolicyPutResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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

func (r PolicyReadResponse) Success() PolicyReadResponse {
	r.Err = ""
	return r
}
func (r PolicyReadResponse) NotFound() PolicyReadResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r PolicyReadResponse) BadRequest() PolicyReadResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r PolicyReadResponse) Unauthorized() PolicyReadResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r PolicyReadResponse) Internal() PolicyReadResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r PolicyReadResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}

// PolicyDeleteRequest to delete a policy.
type PolicyDeleteRequest struct {
	ID string `json:"id"`
}

// PolicyDeleteResponse to delete a policy.
type PolicyDeleteResponse struct {
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (r PolicyDeleteResponse) Success() PolicyDeleteResponse {
	r.Err = ""
	return r
}
func (r PolicyDeleteResponse) NotFound() PolicyDeleteResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r PolicyDeleteResponse) BadRequest() PolicyDeleteResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r PolicyDeleteResponse) Unauthorized() PolicyDeleteResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r PolicyDeleteResponse) Internal() PolicyDeleteResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r PolicyDeleteResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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

func (r PolicyListResponse) Success() PolicyListResponse {
	r.Err = ""
	return r
}
func (r PolicyListResponse) NotFound() PolicyListResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrAPIResponseCodeInvalid)
	return r
}
func (r PolicyListResponse) BadRequest() PolicyListResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r PolicyListResponse) Unauthorized() PolicyListResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r PolicyListResponse) Internal() PolicyListResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r PolicyListResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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

func (r PolicyAccessCheckResponse) Success() PolicyAccessCheckResponse {
	r.Err = ""
	return r
}
func (r PolicyAccessCheckResponse) NotFound() PolicyAccessCheckResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r PolicyAccessCheckResponse) BadRequest() PolicyAccessCheckResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r PolicyAccessCheckResponse) Unauthorized() PolicyAccessCheckResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r PolicyAccessCheckResponse) Internal() PolicyAccessCheckResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r PolicyAccessCheckResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}
