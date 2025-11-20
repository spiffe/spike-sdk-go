//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import sdkErrors "github.com/spiffe/spike-sdk-go/errors"

// Get retrieves a versioned key-value data map from the store at the specified
// path.
//
// The function supports versioned data retrieval with the following behavior:
//   - If version is 0, returns the current version of the data
//   - If version is specified, returns that specific version if it exists
//   - Returns nil if the path doesn't exist
//   - Returns nil if the specified version doesn't exist
//   - Returns nil if the version has been deleted (DeletedTime is set)
//
// Parameters:
//   - path: The path to retrieve data from
//   - version: The specific version to retrieve (0 for current version)
//
// Returns:
//   - map[string]string: The key-value data at the specified path and version,
//     nil on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrEntityNotFound: if the path doesn't exist
//   - ErrStoreItemSoftDeleted: if the version doesn't exist or has been deleted
//
// Example:
//
//	// Get current version
//	data, err := kv.Get("secret/myapp", 0)
//	if err != nil {
//	    log.Printf("Failed to get secret: %v", err)
//	    return
//	}
//
//	// Get specific version
//	historicalData, err := kv.Get("secret/myapp", 2)
//	if err != nil {
//	    log.Printf("Failed to get version 2: %v", err)
//	    return
//	}
func (kv *KV) Get(path string, version int) (map[string]string, *sdkErrors.SDKError) {
	secret, exists := kv.data[path]
	if !exists {
		return nil, sdkErrors.ErrEntityNotFound
	}

	// If the version not specified, use the current version:
	if version == 0 {
		version = secret.Metadata.CurrentVersion
	}

	v, exists := secret.Versions[version]
	if !exists || v.DeletedTime != nil {
		return nil, sdkErrors.ErrStoreItemSoftDeleted
	}

	return v.Data, nil
}

// GetRawSecret retrieves a raw secret from the store at the specified path.
// This function is similar to Get, but it returns the raw Value object instead
// of the key-value data map, providing access to all versions and metadata.
//
// Parameters:
//   - path: The path to retrieve the secret from
//
// Returns:
//   - *Value: The complete secret object with all versions and metadata, nil
//     on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrEntityNotFound: if the path doesn't exist
//
// Example:
//
//	secret, err := kv.GetRawSecret("secret/myapp")
//	if err != nil {
//	    log.Printf("Failed to get raw secret: %v", err)
//	    return
//	}
//	log.Printf("Current version: %d", secret.Metadata.CurrentVersion)
func (kv *KV) GetRawSecret(path string) (*Value, *sdkErrors.SDKError) {
	secret, exists := kv.data[path]
	if !exists {
		return nil, sdkErrors.ErrEntityNotFound
	}

	return secret, nil
}
