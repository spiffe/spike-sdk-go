//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	data2 "github.com/spiffe/spike-sdk-go/api/entity/data"
)

type PolicyCreateRequest struct {
	Name            string                   `json:"name"`
	SpiffeIdPattern string                   `json:"spiffe_id_pattern"`
	PathPattern     string                   `json:"path_pattern"`
	Permissions     []data2.PolicyPermission `json:"permissions"`
}

type PolicyCreateResponse struct {
	Id  string          `json:"id,omitempty"`
	Err data2.ErrorCode `json:"err,omitempty"`
}

type PolicyReadRequest struct {
	Id string `json:"id"`
}

type PolicyReadResponse struct {
	data2.Policy
	Err data2.ErrorCode `json:"err,omitempty"`
}

type PolicyDeleteRequest struct {
	Id string `json:"id"`
}

type PolicyDeleteResponse struct {
	Err data2.ErrorCode `json:"err,omitempty"`
}

type PolicyListRequest struct{}

type PolicyListResponse struct {
	Policies []data2.Policy  `json:"policies"`
	Err      data2.ErrorCode `json:"err,omitempty"`
}

type PolicyAccessCheckRequest struct {
	SpiffeId string `json:"spiffe_id"`
	Path     string `json:"path"`
	Action   string `json:"action"`
}

type PolicyAccessCheckResponse struct {
	Allowed          bool            `json:"allowed"`
	MatchingPolicies []string        `json:"matching_policies"`
	Err              data2.ErrorCode `json:"err,omitempty"`
}
