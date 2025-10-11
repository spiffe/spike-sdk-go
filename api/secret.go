//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/internal/impl/secret"
)

// DeleteSecretVersions deletes specified versions of a secret at the given
// path
//
// It constructs a delete request and sends it to the secrets API endpoint.
//
// Parameters:
//   - path string: Path to the secret to delete
//   - versions []int: Array of version numbers to delete
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or wrapped
//     error on request/parsing failure
//
// Example:
//
//	err := api.DeleteSecretVersions("secret/path", []int{1, 2})
func (a *API) DeleteSecretVersions(path string, versions []int) error {
	return secret.Delete(a.source, path, versions)
}

// DeleteSecret deletes the entire secret at the given path
//
// Parameters:
//   - path string: Path to the secret to delete
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or wrapped
//     error on request/parsing failure
//
// Example:
//
//	err := api.DeleteSecret("secret/path")
func (a *API) DeleteSecret(path string) error {
	return secret.Delete(a.source, path, []int{})
}

// GetSecretVersion retrieves a specific version of a secret at the given
// path.
//
// Parameters:
//   - path string: Path to the secret to retrieve
//   - version int: Version number of the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	secret, err := api.GetSecretVersion("secret/path", 1)
func (a *API) GetSecretVersion(
	path string, version int,
) (*data.Secret, error) {
	return secret.Get(a.source, path, version)
}

// GetSecret retrieves the latest version of the secret at the given path.
//
// Parameters:
//   - path string: Path to the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	secret, err := api.GetSecret("secret/path")
func (a *API) GetSecret(path string) (*data.Secret, error) {
	return secret.Get(a.source, path, 0)
}

// ListSecretKeys retrieves all secret keys.
//
// Returns:
//   - *[]string: Pointer to an array of secret keys if found, nil if none found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	keys, err := api.ListSecretKeys()
func (a *API) ListSecretKeys() (*[]string, error) {
	return secret.ListKeys(a.source)
}

// GetSecretMetadata retrieves metadata for a specific version of a secret at
// the given path.
//
// Parameters:
//   - path string: Path to the secret to retrieve metadata for
//   - version int: Version number of the secret to retrieve metadata for
//
// Returns:
//   - *data.SecretMetadata: Secret metadata if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	metadata, err := api.GetSecretMetadata("secret/path", 1)
func (a *API) GetSecretMetadata(
	path string, version int,
) (*data.SecretMetadata, error) {
	return secret.GetMetadata(a.source, path, version)
}

// PutSecret creates or updates a secret at the specified path with the given
// values.
//
// Parameters:
//   - path string: Path where the secret should be stored
//   - data map[string]string: Map of key-value pairs representing the secret
//     data
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	err := api.PutSecret("secret/path", map[string]string{"key": "value"})
func (a *API) PutSecret(path string, data map[string]string) error {
	return secret.Put(a.source, path, data)
}

// UndeleteSecret restores previously deleted versions of a secret at the
// specified path.
//
// Parameters:
//   - path string: Path to the secret to restore
//   - versions []int: Array of version numbers to restore. Empty array
//     attempts no restoration
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	err := api.UndeleteSecret("secret/path", []int{1, 2})
func (a *API) UndeleteSecret(path string, versions []int) error {
	return secret.Undelete(a.source, path, versions)
}
