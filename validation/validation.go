package validation

import (
	"github.com/google/uuid"
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"regexp"

	"github.com/spiffe/spike-sdk-go/api/errors"
)

const validNamePattern = `^[a-zA-Z0-9-_ ]+$`
const maxNameLength = 250
const validSpiffeIdPattern = `^\^spiffe://[a-zA-Z0-9.\-*]+(/[a-zA-Z0-9._\-*]+)*$`
const validRawSpiffeIdPattern = `^spiffe://[a-zA-Z0-9.-]+(/[a-zA-Z0-9._-]+)*$`
const maxPathPatternLength = 500
const validPathPattern = `^[a-zA-Z0-9._\-/^$()?+*|[\]{}\\]+$`
const validPath = `^[a-zA-Z0-9._\-/()?+*|[\]{}\\]+$`

// ValidateName checks if the provided name meets length and format constraints.
// It returns an error if the name is invalid, otherwise it returns nil.
func ValidateName(name string) error {
	// Validate length
	if len(name) == 0 || len(name) > maxNameLength {
		return errors.ErrInvalidInput
	}

	// Validate format
	if match, _ := regexp.MatchString(validNamePattern, name); !match {
		return errors.ErrInvalidInput
	}

	return nil
}

// ValidateSpiffeIdPattern validates whether the given SPIFFE ID pattern string
// conforms to the expected format and returns an error if it does not.
func ValidateSpiffeIdPattern(spiffeIdPattern string) error {
	// Validate SpiffeIdPattern
	if match, _ := regexp.MatchString(
		validSpiffeIdPattern, spiffeIdPattern); !match {
		return errors.ErrInvalidInput
	}

	return nil
}

// ValidateSpiffeId validates if the given SPIFFE ID matches the expected format.
// Returns an error if the SPIFFE ID is invalid.
func ValidateSpiffeId(spiffeId string) error {
	if match, _ := regexp.MatchString(
		validRawSpiffeIdPattern, spiffeId); !match {
		return errors.ErrInvalidInput
	}
	return nil
}

// ValidatePathPattern validates the given path pattern string for correctness.
// Returns an error if the pattern is empty, too long, or has invalid characters.
func ValidatePathPattern(pathPattern string) error {
	// Validate length
	if len(pathPattern) == 0 || len(pathPattern) > maxPathPatternLength {
		return errors.ErrInvalidInput
	}

	// Validate format
	// Allow regex special characters along with alphanumeric and basic symbols
	if match, _ := regexp.MatchString(validPathPattern, pathPattern); !match {
		return errors.ErrInvalidInput
	}

	return nil
}

// ValidatePath checks if the given path is valid based on predefined rules.
// It returns an error if the path is empty, too long, or contains invalid
// characters.
func ValidatePath(path string) error {
	if len(path) == 0 || len(path) > maxPathPatternLength {
		return errors.ErrInvalidInput
	}
	if match, _ := regexp.MatchString(validPath, path); !match {
		return errors.ErrInvalidInput
	}
	return nil
}

// ValidatePolicyId verifies if the given policyId is a valid UUID format.
// Returns errors.ErrInvalidInput if the validation fails.
func ValidatePolicyId(policyId string) error {
	err := uuid.Validate(policyId)
	if err != nil {
		return errors.ErrInvalidInput
	}
	return nil
}

// ValidatePermissions checks if all provided permissions are valid.
// Permissions are compared against a predefined list of allowed permissions.
// Returns ErrInvalidInput if any permission is invalid, nil otherwise.
func ValidatePermissions(permissions []data.PolicyPermission) error {
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
			return errors.ErrInvalidInput
		}
	}

	return nil
}
