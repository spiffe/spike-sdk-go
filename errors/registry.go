//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

import "sync"

type ErrorCode string

// errorRegistry maps ErrorCodes to their corresponding SDKError instances.
// This map is automatically populated when errors are defined using the
// register() function, ensuring FromCode() always has up-to-date mappings.
// Access is protected by errorRegistryMu for thread safety.
var (
	errorRegistry   = make(map[ErrorCode]*SDKError)
	errorRegistryMu sync.RWMutex
)

// register creates a new SDKError, adds it to the global registry, and
// returns it. This ensures that all defined errors are automatically available
// in FromCode(). This function is thread-safe.
//
// Parameters:
//   - code: The error code string
//   - msg: The human-readable error message
//   - wrapped: Optional wrapped error (typically nil for predefined errors)
//
// Returns:
//   - *SDKError: The newly created and registered error
func register(code string, msg string, wrapped error) *SDKError {
	err := New(ErrorCode(code), msg, wrapped)
	errorRegistryMu.Lock()
	errorRegistry[err.Code] = err
	errorRegistryMu.Unlock()
	return err
}
