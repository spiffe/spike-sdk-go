//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/secret"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// DeleteSecretVersions deletes specified versions of a secret at the given
// path.
//
// Parameters:
//   - path: Path to the secret to delete
//   - versions: Array of version numbers to delete
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	err := api.DeleteSecretVersions("secret/path", []int{1, 2})
//	if err != nil {
//	    log.Printf("Failed to delete secret versions: %v", err)
//	}
func (a *API) DeleteSecretVersions(
	path string, versions []int,
) *sdkErrors.SDKError {
	return secret.Delete(a.source, path, versions)
}

// DeleteSecret deletes the entire secret at the given path.
//
// Parameters:
//   - path: Path to the secret to delete
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	err := api.DeleteSecret("secret/path")
//	if err != nil {
//	    log.Printf("Failed to delete secret: %v", err)
//	}
func (a *API) DeleteSecret(path string) *sdkErrors.SDKError {
	return secret.Delete(a.source, path, []int{})
}

// GetSecretVersion retrieves a specific version of a secret at the given
// path.
//
// Parameters:
//   - path: Path to the secret to retrieve
//   - version: Version number of the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil if not found or on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrAPINotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (nil, nil) if the secret is not found (ErrAPINotFound)
//
// Example:
//
//	secret, err := api.GetSecretVersion("secret/path", 1)
//	if err != nil {
//	    log.Printf("Error retrieving secret: %v", err)
//	    return
//	}
//	if secret == nil {
//	    log.Printf("Secret not found")
//	    return
//	}
func (a *API) GetSecretVersion(
	path string, version int,
) (*data.Secret, *sdkErrors.SDKError) {
	return secret.Get(a.source, path, version)
}

// GetSecret retrieves the latest version of the secret at the given path.
//
// Parameters:
//   - path: Path to the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil if not found or on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrAPINotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (nil, nil) if the secret is not found (ErrAPINotFound)
//
// Example:
//
//	secret, err := api.GetSecret("secret/path")
//	if err != nil {
//	    log.Printf("Error retrieving secret: %v", err)
//	    return
//	}
//	if secret == nil {
//	    log.Printf("Secret not found")
//	    return
//	}
func (a *API) GetSecret(path string) (*data.Secret, *sdkErrors.SDKError) {
	return secret.Get(a.source, path, 0)
}

// ListSecretKeys retrieves all secret keys.
//
// Returns:
//   - *[]string: Array of secret keys if found, empty array if none found,
//     nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrAPINotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (&[]string{}, nil) if no secrets are found (ErrAPINotFound)
//
// Example:
//
//	keys, err := api.ListSecretKeys()
//	if err != nil {
//	    log.Printf("Error listing keys: %v", err)
//	    return
//	}
//	for _, key := range *keys {
//	    log.Printf("Found key: %s", key)
//	}
func (a *API) ListSecretKeys() (*[]string, *sdkErrors.SDKError) {
	return secret.ListKeys(a.source)
}

// GetSecretMetadata retrieves metadata for a specific version of a secret at
// the given path.
//
// Parameters:
//   - path: Path to the secret to retrieve metadata for
//   - version: Version number of the secret to retrieve metadata for
//
// Returns:
//   - *data.SecretMetadata: Secret metadata if found, nil if not found or on
//     error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails (except ErrAPINotFound)
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Note: Returns (nil, nil) if the secret metadata is not found (ErrAPINotFound)
//
// Example:
//
//	metadata, err := api.GetSecretMetadata("secret/path", 1)
//	if err != nil {
//	    log.Printf("Error retrieving metadata: %v", err)
//	    return
//	}
//	if metadata == nil {
//	    log.Printf("Metadata not found")
//	    return
//	}
func (a *API) GetSecretMetadata(
	path string, version int,
) (*data.SecretMetadata, *sdkErrors.SDKError) {
	return secret.GetMetadata(a.source, path, version)
}

// PutSecret creates or updates a secret at the specified path with the given
// values.
//
// Parameters:
//   - path: Path where the secret should be stored
//   - data: Map of key-value pairs representing the secret data
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	err := api.PutSecret("secret/path", map[string]string{"key": "value"})
//	if err != nil {
//	    log.Printf("Failed to put secret: %v", err)
//	}
func (a *API) PutSecret(
	path string, data map[string]string,
) *sdkErrors.SDKError {
	return secret.Put(a.source, path, data)
}

// UndeleteSecret restores previously deleted versions of a secret at the
// specified path.
//
// Parameters:
//   - path: Path to the secret to restore
//   - versions: Array of version numbers to restore (empty array attempts no
//     restoration)
//
// Returns:
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	err := api.UndeleteSecret("secret/path", []int{1, 2})
//	if err != nil {
//	    log.Printf("Failed to undelete secret: %v", err)
//	}
func (a *API) UndeleteSecret(path string, versions []int) *sdkErrors.SDKError {
	return secret.Undelete(a.source, path, versions)
}
