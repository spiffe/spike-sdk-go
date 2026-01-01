//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package validation

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
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
		return sdkErrors.ErrDataInvalidInput.Clone()
	}

	// Validate format
	if match, _ := regexp.MatchString(validNamePattern, name); !match {
		return sdkErrors.ErrDataInvalidInput.Clone()
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
		return sdkErrors.ErrDataInvalidInput.Clone()
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
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	return nil
}

// ValidatePathPattern validates the given path pattern string for correctness.
//
// This function is used for validating path patterns that may contain regex
// metacharacters for matching multiple paths. The path pattern must be between
// 1 and 500 characters and may include regex anchors (^, $) and other regex
// special characters (?, +, *, |, [], {}, \, etc.) along with alphanumeric
// characters, underscores, hyphens, forward slashes, and periods.
//
// Use ValidatePath instead if you need to validate literal paths without
// regex anchors.
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
		return sdkErrors.ErrDataInvalidInput.Clone()
	}

	// Validate format
	// Allow regex special characters along with alphanumeric and basic symbols
	if match, _ := regexp.MatchString(validPathPattern, pathPattern); !match {
		return sdkErrors.ErrDataInvalidInput.Clone()
	}

	return nil
}

// ValidatePath checks if the given path is valid based on predefined rules.
//
// This function validates paths that should not contain regex anchor
// metacharacters (^ or $). Unlike ValidatePathPattern, this function is for
// validating literal paths. The path must be between 1 and 500 characters.
//
// Note: While this function excludes regex anchors (^, $), it still allows
// other special characters that may appear in actual paths such as ?, +, *,
// |, [], {}, \, /, etc. Use ValidatePathPattern if you need to validate
// patterns that include regex anchors.
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
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	if match, _ := regexp.MatchString(validPath, path); !match {
		return sdkErrors.ErrDataInvalidInput.Clone()
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
		return sdkErrors.ErrDataInvalidInput.Clone()
	}
	return nil
}

// validPermissions contains the set of valid policy permissions supported by
// the SPIKE system. These are sourced from the SDK to prevent typos.
//
// Valid permissions are:
//   - read: Read access to resources
//   - write: Write access to resources
//   - list: List access to resources
//   - execute: Execute access to resources
//   - super: Superuser access (grants all permissions)
var validPermissions = []data.PolicyPermission{
	data.PermissionRead,
	data.PermissionWrite,
	data.PermissionList,
	data.PermissionExecute,
	data.PermissionSuper,
}

// ValidPermission checks if the given permission string is valid.
//
// Parameters:
//   - perm: The permission string to validate.
//
// Returns:
//   - true if the permission is found in ValidPermissions, false otherwise.
func ValidPermission(perm string) bool {
	for _, p := range validPermissions {
		if string(p) == perm {
			return true
		}
	}
	return false
}

// ValidPermissionsList returns a comma-separated string of valid permissions,
// suitable for display in error messages.
//
// Returns:
//   - string: A comma-separated list of valid permissions.
func ValidPermissionsList() string {
	perms := make([]string, len(validPermissions))
	for i, p := range validPermissions {
		perms[i] = string(p)
	}
	return strings.Join(perms, ", ")
}

// ValidatePermissions validates policy permissions from a comma-separated
// string and returns a slice of PolicyPermission values. It returns an error
// if any permission is invalid or if the string contains no valid permissions.
//
// Valid permissions are:
//   - read: Read access to resources
//   - write: Write access to resources
//   - list: List access to resources
//   - execute: Execute access to resources
//   - super: Superuser access (grants all permissions)
//
// Parameters:
//   - permsStr: Comma-separated string of permissions
//     (e.g., "read,write,execute")
//
// Returns:
//   - []data.PolicyPermission: Validated policy permissions
//   - *sdkErrors.SDKError: An error if any permission is invalid
func ValidatePermissions(permsStr string) (
	[]data.PolicyPermission, *sdkErrors.SDKError,
) {
	var permissions []string
	for _, p := range strings.Split(permsStr, ",") {
		perm := strings.TrimSpace(p)
		if perm != "" {
			permissions = append(permissions, perm)
		}
	}

	perms := make([]data.PolicyPermission, 0, len(permissions))
	for _, perm := range permissions {
		if !ValidPermission(perm) {
			failErr := *sdkErrors.ErrAccessInvalidPermission.Clone()
			failErr.Msg = fmt.Sprintf(
				"invalid permission: '%s'. valid permissions: '%s'",
				perm, ValidPermissionsList(),
			)
			return nil, &failErr
		}
		perms = append(perms, data.PolicyPermission(perm))
	}

	if len(perms) == 0 {
		failErr := *sdkErrors.ErrAccessInvalidPermission.Clone()
		failErr.Msg = "no valid permissions specified" +
			". valid permissions are: " + ValidPermissionsList()
		return nil, &failErr
	}

	return perms, nil
}

// NonNilContextOrDie checks if the provided context is nil and terminates the
// program if so.
//
// This function is used to ensure that all operations requiring a context
// receive a valid one. A nil context indicates a programming error that
// should never occur in production, so the function terminates the program
// immediately via log.FatalErr.
//
// Parameters:
//   - ctx: The context to validate
//   - fName: The calling function name for logging purposes
func NonNilContextOrDie(ctx context.Context, fName string) {
	if ctx == nil {
		failErr := *sdkErrors.ErrNilContext.Clone()
		log.FatalErr(fName, failErr)
	}
}
