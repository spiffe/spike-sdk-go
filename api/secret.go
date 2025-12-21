//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/secret"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// DeleteSecretVersions deletes specified versions of a secret at the given
// path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	err := api.DeleteSecretVersions(ctx, "secret/path", []int{1, 2})
//	if err != nil {
//	    log.Printf("Failed to delete secret versions: %v", err)
//	}
func (a *API) DeleteSecretVersions(
	ctx context.Context, path string, versions []int,
) *sdkErrors.SDKError {
	return secret.Delete(ctx, a.source, path, versions)
}

// DeleteSecret deletes the entire secret at the given path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	err := api.DeleteSecret(ctx, "secret/path")
//	if err != nil {
//	    log.Printf("Failed to delete secret: %v", err)
//	}
func (a *API) DeleteSecret(ctx context.Context, path string) *sdkErrors.SDKError {
	return secret.Delete(ctx, a.source, path, []int{})
}

// GetSecretVersion retrieves a specific version of a secret at the given
// path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - path: Path to the secret to retrieve
//   - version: Version number of the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPINotFound: if the secret is not found
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	secret, err := api.GetSecretVersion(ctx, "secret/path", 1)
//	if err != nil {
//	    if err.Is(sdkErrors.ErrAPINotFound) {
//	        log.Printf("Secret not found")
//	        return
//	    }
//	    log.Printf("Error retrieving secret: %v", err)
//	    return
//	}
func (a *API) GetSecretVersion(
	ctx context.Context, path string, version int,
) (*data.Secret, *sdkErrors.SDKError) {
	return secret.Get(ctx, a.source, path, version)
}

// GetSecret retrieves the latest version of the secret at the given path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - path: Path to the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPINotFound: if the secret is not found
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	secret, err := api.GetSecret(ctx, "secret/path")
//	if err != nil {
//	    if err.Is(sdkErrors.ErrAPINotFound) {
//	        log.Printf("Secret not found")
//	        return
//	    }
//	    log.Printf("Error retrieving secret: %v", err)
//	    return
//	}
func (a *API) GetSecret(
	ctx context.Context, path string,
) (*data.Secret, *sdkErrors.SDKError) {
	return secret.Get(ctx, a.source, path, 0)
}

// ListSecretKeys retrieves all secret keys.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	keys, err := api.ListSecretKeys(ctx)
//	if err != nil {
//	    log.Printf("Error listing keys: %v", err)
//	    return
//	}
//	for _, key := range *keys {
//	    log.Printf("Found key: %s", key)
//	}
func (a *API) ListSecretKeys(ctx context.Context) (*[]string, *sdkErrors.SDKError) {
	return secret.ListKeys(ctx, a.source)
}

// GetSecretMetadata retrieves metadata for a specific version of a secret at
// the given path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
//   - path: Path to the secret to retrieve metadata for
//   - version: Version number of the secret to retrieve metadata for
//
// Returns:
//   - *data.SecretMetadata: Secret metadata if found, nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrSPIFFENilX509Source: if the X509 source is nil
//   - ErrDataMarshalFailure: if request serialization fails
//   - ErrAPINotFound: if the secret metadata is not found
//   - Errors from net.Post(): if the HTTP request fails
//   - ErrDataUnmarshalFailure: if response parsing fails
//   - Error from FromCode(): if the server returns an error
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	metadata, err := api.GetSecretMetadata(ctx, "secret/path", 1)
//	if err != nil {
//	    if err.Is(sdkErrors.ErrAPINotFound) {
//	        log.Printf("Metadata not found")
//	        return
//	    }
//	    log.Printf("Error retrieving metadata: %v", err)
//	    return
//	}
func (a *API) GetSecretMetadata(
	ctx context.Context, path string, version int,
) (*data.SecretMetadata, *sdkErrors.SDKError) {
	return secret.GetMetadata(ctx, a.source, path, version)
}

// PutSecret creates or updates a secret at the specified path with the given
// values.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	err := api.PutSecret(ctx, "secret/path", map[string]string{"key": "value"})
//	if err != nil {
//	    log.Printf("Failed to put secret: %v", err)
//	}
func (a *API) PutSecret(
	ctx context.Context, path string, data map[string]string,
) *sdkErrors.SDKError {
	return secret.Put(ctx, a.source, path, data)
}

// UndeleteSecret restores previously deleted versions of a secret at the
// specified path.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control. If nil,
//     context.Background() is used.
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
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	err := api.UndeleteSecret(ctx, "secret/path", []int{1, 2})
//	if err != nil {
//	    log.Printf("Failed to undelete secret: %v", err)
//	}
func (a *API) UndeleteSecret(
	ctx context.Context, path string, versions []int,
) *sdkErrors.SDKError {
	return secret.Undelete(ctx, a.source, path, versions)
}
