//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
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
	m.Err = sdkErrors.ErrSuccess.Code
	return m
}
func (s SecretMetadataResponse) NotFound() SecretMetadataResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s SecretMetadataResponse) BadRequest() SecretMetadataResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s SecretMetadataResponse) Unauthorized() SecretMetadataResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s SecretMetadataResponse) Internal() SecretMetadataResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
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
	s.Err = sdkErrors.ErrSuccess.Code
	return s
}
func (s SecretPutResponse) NotFound() SecretPutResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s SecretPutResponse) BadRequest() SecretPutResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s SecretPutResponse) Unauthorized() SecretPutResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s SecretPutResponse) Internal() SecretPutResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}

// SecretGetRequest is for getting secrets
type SecretGetRequest struct {
	Path    string `json:"path"`
	Version int    `json:"version,omitempty"` // Optional specific version
}

// SecretGetResponse is for getting secrets
type SecretGetResponse struct {
	data.Secret
	Err sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (s SecretGetResponse) Success() SecretGetResponse {
	s.Err = sdkErrors.ErrSuccess.Code
	return s
}
func (s SecretGetResponse) NotFound() SecretGetResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s SecretGetResponse) BadRequest() SecretGetResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s SecretGetResponse) Unauthorized() SecretGetResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s SecretGetResponse) Internal() SecretGetResponse {
	s.Err = sdkErrors.ErrInternal.Code
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

func (s SecretDeleteResponse) NotFound() SecretDeleteResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s SecretDeleteResponse) BadRequest() SecretDeleteResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s SecretDeleteResponse) Unauthorized() SecretDeleteResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s SecretDeleteResponse) Internal() SecretDeleteResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}

func (s SecretDeleteResponse) Success() SecretDeleteResponse {
	s.Err = sdkErrors.ErrSuccess.Code
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
	s.Err = sdkErrors.ErrSuccess.Code
	return s
}
func (s SecretUndeleteResponse) NotFound() SecretUndeleteResponse {
	s.Err = sdkErrors.ErrNotFound.Code
	return s
}
func (s SecretUndeleteResponse) BadRequest() SecretUndeleteResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s SecretUndeleteResponse) Unauthorized() SecretUndeleteResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s SecretUndeleteResponse) Internal() SecretUndeleteResponse {
	s.Err = sdkErrors.ErrInternal.Code
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
	s.Err = sdkErrors.ErrSuccess.Code
	return s
}
func (s SecretListResponse) NotFound() SecretListResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrEntityResponseCodeInvalid)
	return s
}
func (s SecretListResponse) BadRequest() SecretListResponse {
	s.Err = sdkErrors.ErrBadRequest.Code
	return s
}
func (s SecretListResponse) Unauthorized() SecretListResponse {
	s.Err = sdkErrors.ErrAccessUnauthorized.Code
	return s
}
func (s SecretListResponse) Internal() SecretListResponse {
	s.Err = sdkErrors.ErrInternal.Code
	return s
}
