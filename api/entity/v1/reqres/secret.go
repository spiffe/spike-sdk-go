//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// SecretMetadataRequest for get secrets metadata
type SecretMetadataRequest struct {
	Path    string `json:"path"`
	Version int    `json:"version,omitempty"` // Optional specific version
}

// SecretMetadataResponse for secrets versions and metadata
type SecretMetadataResponse struct {
	data.SecretMetadata
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (m SecretMetadataResponse) Success() SecretMetadataResponse {
	m.Err = sdkErrors.ErrCodeSuccess
	return m
}

// SecretPutRequest for creating/updating secrets
type SecretPutRequest struct {
	Path   string              `json:"path"`
	Values map[string]string   `json:"values"`
	Err    sdkErrors.ErrorCode `json:"err,omitempty"`
}

// SecretPutResponse is after a successful secret write operation.
type SecretPutResponse struct {
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (s SecretPutResponse) Success() SecretPutResponse {
	s.Err = sdkErrors.ErrCodeSuccess
	return s
}

// SecretReadRequest is for getting secrets
type SecretReadRequest struct {
	Path    string `json:"path"`
	Version int    `json:"version,omitempty"` // Optional specific version
}

// SecretReadResponse is for getting secrets
type SecretReadResponse struct {
	data.Secret
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (s SecretReadResponse) Success() SecretReadResponse {
	s.Err = sdkErrors.ErrCodeSuccess
	return s
}

// SecretDeleteRequest for soft-deleting secret versions
type SecretDeleteRequest struct {
	Path     string `json:"path"`
	Versions []int  `json:"versions"` // Empty means the latest version
}

// SecretDeleteResponse after soft-delete
type SecretDeleteResponse struct {
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (s SecretDeleteResponse) Success() SecretDeleteResponse {
	s.Err = sdkErrors.ErrCodeSuccess
	return s
}

// SecretUndeleteRequest for recovering soft-deleted versions
type SecretUndeleteRequest struct {
	Path     string `json:"path"`
	Versions []int  `json:"versions"`
}

// SecretUndeleteResponse after recovery
type SecretUndeleteResponse struct {
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (s SecretUndeleteResponse) Success() SecretUndeleteResponse {
	s.Err = sdkErrors.ErrCodeSuccess
	return s
}

// SecretListRequest for listing secrets
type SecretListRequest struct {
}

// SecretListResponse for listing secrets
type SecretListResponse struct {
	Keys []string            `json:"keys"`
	Err  sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (s SecretListResponse) Success() SecretListResponse {
	s.Err = sdkErrors.ErrCodeSuccess
	return s
}
