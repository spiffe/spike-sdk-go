package validation

import (
	"regexp"

	"github.com/spiffe/spike-sdk-go/api/errors"
)

const validNamePattern = `^[a-zA-Z0-9-_ ]+$`
const maxNameLength = 250
const validSpiffeIdPattern = `^\^spiffe://[a-zA-Z0-9.-]+(/[a-zA-Z0-9._-]+)*$`
const validRawSpiffeIdPattern = `^spiffe://[a-zA-Z0-9.-]+(/[a-zA-Z0-9._-]+)*$`
const maxPathPatternLength = 500

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
	const validPathPattern = `^[a-zA-Z0-9._\-/^$()?+*|[\]{}\\]+$`
	if match, _ := regexp.MatchString(validPathPattern, pathPattern); !match {
		return errors.ErrInvalidInput
	}

	return nil
}
