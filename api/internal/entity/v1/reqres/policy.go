//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity"
	data "github.com/spiffe/spike-sdk-go/api/internal/entity/data"
)

type PolicyCreateRequest struct {
	Name            string                    `json:"name"`
	SpiffeIdPattern string                    `json:"spiffe_id_pattern"`
	PathPattern     string                    `json:"path_pattern"`
	Permissions     []entity.PolicyPermission `json:"permissions"`
}

type PolicyCreateResponse struct {
	Id  string         `json:"id,omitempty"`
	Err data.ErrorCode `json:"err,omitempty"`
}

type PolicyReadRequest struct {
	Id string `json:"id"`
}

type PolicyReadResponse struct {
	entity.Policy
	Err data.ErrorCode `json:"err,omitempty"`
}

type PolicyDeleteRequest struct {
	Id string `json:"id"`
}

type PolicyDeleteResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}

type PolicyListRequest struct{}

type PolicyListResponse struct {
	Policies []entity.Policy `json:"policies"`
	Err      data.ErrorCode  `json:"err,omitempty"`
}

type PolicyAccessCheckRequest struct {
	SpiffeId string `json:"spiffe_id"`
	Path     string `json:"path"`
	Action   string `json:"action"`
}

type PolicyAccessCheckResponse struct {
	Allowed          bool           `json:"allowed"`
	MatchingPolicies []string       `json:"matching_policies"`
	Err              data.ErrorCode `json:"err,omitempty"`
}
