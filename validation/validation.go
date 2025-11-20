//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package validation

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

const validNamePattern = `^[a-zA-Z0-9-_ ]+$`
const maxNameLength = 250
const validSPIFFEIDPattern = `^\^?spiffe://[\\a-zA-Z0-9.\-*()+?\[\]]+(/[\\/a-zA-Z0-9._\-*()+?\[\]]+)*\$?$`
const validRawSPIFFEIDPattern = `^spiffe://[a-zA-Z0-9.-]+(/[a-zA-Z0-9._-]+)*$`
const maxPathPatternLength = 500
const validPathPattern = `^[a-zA-Z0-9._\-/^$()?+*|[\]{}\\]+$`
const validPath = `^[a-zA-Z0-9._\-/()?+*|[\]{}\\]+$`

// ValidateName checks if the provided name meets length and format constraints.
//
// The name must be between 1 and 250 characters and contain only alphanumeric
// characters, hyphens, underscores, and spaces.
//
// Parameters:
//   - name: The name string to validate
//
// Returns:
//   - *sdkErrors.SDKError: nil if valid, or one of the following errors:
//   - ErrDataInvalidInput: if name is empty, exceeds 250 characters, or
//     contains invalid characters
func ValidateName(name string) *sdkErrors.SDKError {
	// Validate length
	if len(name) == 0 || len(name) > maxNameLength {
		return sdkErrors.ErrDataInvalidInput
	}

	// Validate format
	if match, _ := regexp.MatchString(validNamePattern, name); !match {
		return sdkErrors.ErrDataInvalidInput
	}

	return nil
}

// ValidateSPIFFEIDPattern validates whether the given SPIFFE ID pattern string
// conforms to the expected format.
//
// The pattern may include regex special characters for matching multiple
// SPIFFE IDs.
// It must start with "spiffe://" and follow the SPIFFE ID specification with
// optional regex metacharacters.
//
// Parameters:
//   - SPIFFEIDPattern: The SPIFFE ID pattern string to validate
//
// Returns:
//   - *sdkErrors.SDKError: nil if valid, or one of the following errors:
//   - ErrDataInvalidInput: if the pattern does not conform to the expected
//     format
func ValidateSPIFFEIDPattern(SPIFFEIDPattern string) *sdkErrors.SDKError {
	// Validate SPIFFEIDPattern
	if match, _ := regexp.MatchString(
		validSPIFFEIDPattern, SPIFFEIDPattern); !match {
		return sdkErrors.ErrDataInvalidInput
	}

	return nil
}

// ValidateSPIFFEID validates if the given SPIFFE ID matches the expected
// format.
//
// Unlike ValidateSPIFFEIDPattern, this function validates raw SPIFFE IDs
// without regex metacharacters. The ID must strictly conform to the SPIFFE
// specification:
// "spiffe://<trust-domain>/<path>".
//
// Parameters:
//   - SPIFFEID: The SPIFFE ID string to validate
//
// Returns:
//   - *sdkErrors.SDKError: nil if valid, or one of the following errors:
//   - ErrDataInvalidInput: if the SPIFFE ID does not conform to the expected
//     format
func ValidateSPIFFEID(SPIFFEID string) *sdkErrors.SDKError {
	if match, _ := regexp.MatchString(
		validRawSPIFFEIDPattern, SPIFFEID); !match {
		return sdkErrors.ErrDataInvalidInput
	}
	return nil
}

// ValidatePathPattern validates the given path pattern string for correctness.
//
// The path pattern must be between 1 and 500 characters and may include regex
// special characters for matching multiple paths. Allowed characters include
// alphanumeric, underscores, hyphens, forward slashes, periods, and regex
// metacharacters.
//
// Parameters:
//   - pathPattern: The path pattern string to validate
//
// Returns:
//   - *sdkErrors.SDKError: nil if valid, or one of the following errors:
//   - ErrDataInvalidInput: if the pattern is empty, exceeds 500 characters, or
//     contains invalid characters
func ValidatePathPattern(pathPattern string) *sdkErrors.SDKError {
	// Validate length
	if len(pathPattern) == 0 || len(pathPattern) > maxPathPatternLength {
		return sdkErrors.ErrDataInvalidInput
	}

	// Validate format
	// Allow regex special characters along with alphanumeric and basic symbols
	if match, _ := regexp.MatchString(validPathPattern, pathPattern); !match {
		return sdkErrors.ErrDataInvalidInput
	}

	return nil
}

// ValidatePath checks if the given path is valid based on predefined rules.
//
// Unlike ValidatePathPattern, this function validates literal paths without
// regex metacharacters (except those used in actual paths). The path must be
// between 1 and 500 characters and contain only allowed characters.
//
// Parameters:
//   - path: The path string to validate
//
// Returns:
//   - *sdkErrors.SDKError: nil if valid, or one of the following errors:
//   - ErrDataInvalidInput: if the path is empty, exceeds 500 characters, or
//     contains invalid characters
func ValidatePath(path string) *sdkErrors.SDKError {
	if len(path) == 0 || len(path) > maxPathPatternLength {
		return sdkErrors.ErrDataInvalidInput
	}
	if match, _ := regexp.MatchString(validPath, path); !match {
		return sdkErrors.ErrDataInvalidInput
	}
	return nil
}

// ValidatePolicyID verifies if the given policy ID is a valid UUID format.
//
// The policy ID must conform to the UUID specification (RFC 4122).
// This function uses the google/uuid package for validation.
//
// Parameters:
//   - policyID: The policy ID string to validate
//
// Returns:
//   - *sdkErrors.SDKError: nil if valid, or one of the following errors:
//   - ErrDataInvalidInput: if the policy ID is not a valid UUID
func ValidatePolicyID(policyID string) *sdkErrors.SDKError {
	err := uuid.Validate(policyID)
	if err != nil {
		return sdkErrors.ErrDataInvalidInput
	}
	return nil
}

// ValidatePermissions checks if all provided permissions are valid.
//
// Permissions are compared against a predefined list of allowed permissions:
//   - PermissionList: list secrets
//   - PermissionRead: read secret values
//   - PermissionWrite: create/update secrets
//   - PermissionSuper: administrative access
//
// Parameters:
//   - permissions: Slice of policy permissions to validate
//
// Returns:
//   - *sdkErrors.SDKError: nil if all permissions are valid, or one of the
//     following errors:
//   - ErrDataInvalidInput: if any permission is not in the allowed list
func ValidatePermissions(
	permissions []data.PolicyPermission,
) *sdkErrors.SDKError {
	allowedPermissions := []data.PolicyPermission{
		data.PermissionList,
		data.PermissionRead,
		data.PermissionWrite,
		data.PermissionSuper,
	}

	for _, permission := range permissions {
		isAllowed := false
		for _, allowedPermission := range allowedPermissions {
			if permission == allowedPermission {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			return sdkErrors.ErrDataInvalidInput
		}
	}

	return nil
}
