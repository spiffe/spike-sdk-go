//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/kv"
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

func (r SecretMetadataResponse) Success() SecretMetadataResponse {
	r.Err = ""
	return r
}
func (r SecretMetadataResponse) NotFound() SecretMetadataResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r SecretMetadataResponse) BadRequest() SecretMetadataResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r SecretMetadataResponse) Unauthorized() SecretMetadataResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r SecretMetadataResponse) Internal() SecretMetadataResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r SecretMetadataResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}

// ValueToSecretMetadataSuccessResponse converts a key-value store secret value
// into a secret metadata response.
//
// The function transforms the internal kv.Value representation into the API
// response format by:
//   - Converting all secret versions into a map of version info
//   - Extracting metadata including current/oldest versions and timestamps
//   - Preserving version-specific details like creation and deletion times
//
// This conversion is used when clients request secret metadata without
// retrieving the actual secret data, allowing them to inspect version history
// and lifecycle information.
//
// Parameters:
//   - secret: The key-value store secret value containing version history
//     and metadata
//
// Returns:
//   - reqres.SecretMetadataResponse: The formatted metadata response containing
//     version information and metadata suitable for API responses
func ValueToSecretMetadataSuccessResponse(
	secret *kv.Value,
) SecretMetadataResponse {
	versions := make(map[int]data.SecretVersionInfo)
	for _, version := range secret.Versions {
		versions[version.Version] = data.SecretVersionInfo{
			CreatedTime: version.CreatedTime,
			Version:     version.Version,
			DeletedTime: version.DeletedTime,
		}
	}

	return SecretMetadataResponse{
		SecretMetadata: data.SecretMetadata{
			Versions: versions,
			Metadata: data.SecretMetaDataContent{
				CurrentVersion: secret.Metadata.CurrentVersion,
				OldestVersion:  secret.Metadata.OldestVersion,
				CreatedTime:    secret.Metadata.CreatedTime,
				UpdatedTime:    secret.Metadata.UpdatedTime,
				MaxVersions:    secret.Metadata.MaxVersions,
			},
		},
	}.Success()
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

func (r SecretPutResponse) Success() SecretPutResponse {
	r.Err = ""
	return r
}
func (r SecretPutResponse) NotFound() SecretPutResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrAPIResponseCodeInvalid)
	return r
}
func (r SecretPutResponse) BadRequest() SecretPutResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r SecretPutResponse) Unauthorized() SecretPutResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r SecretPutResponse) Internal() SecretPutResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r SecretPutResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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

func (r SecretGetResponse) Success() SecretGetResponse {
	r.Err = ""
	return r
}
func (r SecretGetResponse) NotFound() SecretGetResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r SecretGetResponse) BadRequest() SecretGetResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r SecretGetResponse) Unauthorized() SecretGetResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r SecretGetResponse) Internal() SecretGetResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r SecretGetResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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

func (r SecretDeleteResponse) NotFound() SecretDeleteResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r SecretDeleteResponse) BadRequest() SecretDeleteResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r SecretDeleteResponse) Unauthorized() SecretDeleteResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r SecretDeleteResponse) Internal() SecretDeleteResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}

func (r SecretDeleteResponse) Success() SecretDeleteResponse {
	r.Err = ""
	return r
}
func (r SecretDeleteResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
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

func (r SecretUndeleteResponse) Success() SecretUndeleteResponse {
	r.Err = ""
	return r
}
func (r SecretUndeleteResponse) NotFound() SecretUndeleteResponse {
	r.Err = sdkErrors.ErrAPINotFound.Code
	return r
}
func (r SecretUndeleteResponse) BadRequest() SecretUndeleteResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r SecretUndeleteResponse) Unauthorized() SecretUndeleteResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r SecretUndeleteResponse) Internal() SecretUndeleteResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r SecretUndeleteResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}

// SecretListRequest for listing secrets
type SecretListRequest struct {
}

// SecretListResponse for listing secrets
type SecretListResponse struct {
	Keys []string            `json:"keys"`
	Err  sdkErrors.ErrorCode `json:"err,omitempty"`
}

func (r SecretListResponse) Success() SecretListResponse {
	r.Err = ""
	return r
}
func (r SecretListResponse) NotFound() SecretListResponse {
	log.FatalErr("NotFound", *sdkErrors.ErrAPIResponseCodeInvalid)
	return r
}
func (r SecretListResponse) BadRequest() SecretListResponse {
	r.Err = sdkErrors.ErrAPIBadRequest.Code
	return r
}
func (r SecretListResponse) Unauthorized() SecretListResponse {
	r.Err = sdkErrors.ErrAccessUnauthorized.Code
	return r
}
func (r SecretListResponse) Internal() SecretListResponse {
	r.Err = sdkErrors.ErrAPIInternal.Code
	return r
}
func (r SecretListResponse) ErrorCode() sdkErrors.ErrorCode {
	return r.Err
}
