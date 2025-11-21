//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package errors

// FromCode maps an ErrorCode to its corresponding SDKError using the
// automatically populated error registry. This is used to convert error codes
// received from API responses back to proper SDKError instances.
//
// The registry is automatically populated when errors are defined using the
// register() function, ensuring new errors are immediately available without
// manual updates to this function.
//
// If the error code is not recognized, it returns ErrGeneralFailure.
//
// Parameters:
//   - code: the error code to map
//
// Returns:
//   - *SDKError: the corresponding SDK error instance
func FromCode(code ErrorCode) *SDKError {
	// Defensive coding: While concurrent reads to a map are safe, unless a
	// write happens concurrently; if we enable dynamic error registration
	// later down the line, without a mutex the behavior of this code will be
	// undeterministic.
	errorRegistryMu.RLock()
	err, ok := errorRegistry[code]
	errorRegistryMu.RUnlock()

	if ok {
		return err
	}
	return ErrGeneralFailure
}

// MaybeError converts an error to its string representation if the error is
// not nil. If the error is nil, it returns an empty string.
//
// Parameters:
//   - err: the error to convert to a string
//
// Returns:
//   - string: the error message if err is non-nil, empty string otherwise
func MaybeError(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
