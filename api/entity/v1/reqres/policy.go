//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/errors"
)

// PolicyCreateRequest for policy creation.
type PolicyCreateRequest struct {
	Name            string                  `json:"name"`
	SPIFFEIDPattern string                  `json:"spiffeidPattern"`
	PathPattern     string                  `json:"pathPattern"`
	Permissions     []data.PolicyPermission `json:"permissions"`
}

// PolicyCreateResponse for policy creation.
type PolicyCreateResponse struct {
	ID  string           `json:"id,omitempty"`
	Err errors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyCreateResponse) Success() PolicyCreateResponse {
	p.Err = errors.ErrCodeSuccess
	return p
}

// PolicyReadRequest to read a policy.
type PolicyReadRequest struct {
	ID string `json:"id"`
}

// PolicyReadResponse to read a policy.
type PolicyReadResponse struct {
	data.Policy
	Err errors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyReadResponse) Success() PolicyReadResponse {
	p.Err = errors.ErrCodeSuccess
	return p
}

// PolicyDeleteRequest to delete a policy.
type PolicyDeleteRequest struct {
	ID string `json:"id"`
}

// PolicyDeleteResponse to delete a policy.
type PolicyDeleteResponse struct {
	Err errors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyDeleteResponse) Success() PolicyDeleteResponse {
	p.Err = errors.ErrCodeSuccess
	return p
}

// PolicyListRequest to list policies.
type PolicyListRequest struct {
	SPIFFEIDPattern string `json:"spiffeidPattern"`
	PathPattern     string `json:"pathPattern"`
}

// PolicyListResponse to list policies.
type PolicyListResponse struct {
	Policies []data.Policy    `json:"policies"`
	Err      errors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyListResponse) Success() PolicyListResponse {
	p.Err = errors.ErrCodeSuccess
	return p
}

// PolicyAccessCheckRequest to validate policy access.
type PolicyAccessCheckRequest struct {
	SPIFFEID string `json:"spiffeid"`
	Path     string `json:"path"`
	Action   string `json:"action"`
}

// PolicyAccessCheckResponse to validate policy access.
type PolicyAccessCheckResponse struct {
	Allowed          bool             `json:"allowed"`
	MatchingPolicies []string         `json:"matchingPolicies"`
	Err              errors.ErrorCode `json:"err,omitempty"`
}

func (p PolicyAccessCheckResponse) Success() PolicyAccessCheckResponse {
	p.Err = errors.ErrCodeSuccess
	return p
}
