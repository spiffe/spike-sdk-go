//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package validation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestValidateName_Valid tests ValidateName with valid names
func TestValidateName_Valid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Simple", "simple-name"},
		{"WithUnderscore", "name_with_underscore"},
		{"WithSpace", "name with space"},
		{"Alphanumeric", "Name123"},
		{"Mixed", "My-Policy_Name 123"},
		{"SingleChar", "a"},
		{"MaxLength", strings.Repeat("a", 250)},
		{"AllNumbers", "12345"},
		{"AllDashes", "----"},
		{"AllUnderscores", "____"},
		{"AllSpaces", "    "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input)
			assert.Nil(t, err, "Expected valid name: %s", tt.input)
		})
	}
}

// TestValidateName_Invalid tests ValidateName with invalid names
func TestValidateName_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"TooLong", strings.Repeat("a", 251)},
		{"WithSlash", "name/with/slash"},
		{"WithDot", "name.with.dot"},
		{"WithSpecialChars", "name@example"},
		{"WithParentheses", "name(test)"},
		{"WithBrackets", "name[test]"},
		{"WithBraces", "name{test}"},
		{"WithAsterisk", "name*"},
		{"WithQuestion", "name?"},
		{"WithPlus", "name+"},
		{"WithEquals", "name=value"},
		{"WithPipe", "name|other"},
		{"WithBackslash", "name\\test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input)
			assert.NotNil(t, err, "Expected invalid name: %s", tt.input)
			assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
		})
	}
}

// TestValidateSPIFFEIDPattern_Valid tests ValidateSPIFFEIDPattern with valid patterns
func TestValidateSPIFFEIDPattern_Valid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Simple", "spiffe://example.org/service"},
		{"WithAnchor", "^spiffe://example.org/service$"},
		{"WithWildcard", "spiffe://example.org/.*"},
		{"WithPlus", "spiffe://example.org/service.+"},
		{"WithQuestion", "spiffe://example.org/service.?"},
		{"WithCharClass", "spiffe://example.org/[a-z]+"},
		{"WithUnderscore", "spiffe://example.org/service_name"},
		{"WithDot", "spiffe://trust.example.org/service"},
		{"WithDash", "spiffe://example-org.com/service"},
		{"MultiplePath", "spiffe://example.org/path/to/service"},
		{"WithNumbers", "spiffe://example.org/service123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSPIFFEIDPattern(tt.input)
			assert.Nil(t, err, "Expected valid SPIFFE ID pattern: %s", tt.input)
		})
	}
}

// TestValidateSPIFFEIDPattern_Invalid tests ValidateSPIFFEIDPattern with invalid patterns
func TestValidateSPIFFEIDPattern_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"NoScheme", "example.org/service"},
		{"WrongScheme", "http://example.org/service"},
		{"NoTrustDomain", "spiffe:///service"},
		{"TrailingSlash", "spiffe://example.org/"},
		{"Spaces", "spiffe://example.org/service name"},
		{"InvalidChars", "spiffe://example.org/service@test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSPIFFEIDPattern(tt.input)
			assert.NotNil(t, err, "Expected invalid SPIFFE ID pattern: %s", tt.input)
			assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
		})
	}
}

// TestValidateSPIFFEID_Valid tests ValidateSPIFFEID with valid IDs
func TestValidateSPIFFEID_Valid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Simple", "spiffe://example.org/service"},
		{"WithUnderscore", "spiffe://example.org/service_name"},
		{"WithDot", "spiffe://trust.example.org/service.api"},
		{"WithDash", "spiffe://example-org.com/service-name"},
		{"MultiplePath", "spiffe://example.org/path/to/service"},
		{"WithNumbers", "spiffe://example.org/service123"},
		{"Subdomain", "spiffe://sub.example.org/workload"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSPIFFEID(tt.input)
			assert.Nil(t, err, "Expected valid SPIFFE ID: %s", tt.input)
		})
	}
}

// TestValidateSPIFFEID_Invalid tests ValidateSPIFFEID with invalid IDs
func TestValidateSPIFFEID_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"NoScheme", "example.org/service"},
		{"WrongScheme", "http://example.org/service"},
		{"NoTrustDomain", "spiffe:///service"},
		{"WithAnchor", "^spiffe://example.org/service$"},
		{"WithWildcard", "spiffe://example.org/.*"},
		{"WithRegex", "spiffe://example.org/service.+"},
		{"WithCharClass", "spiffe://example.org/[a-z]"},
		{"Spaces", "spiffe://example.org/service name"},
		{"SpecialChars", "spiffe://example.org/service@test"},
		{"Parentheses", "spiffe://example.org/(service)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSPIFFEID(tt.input)
			assert.NotNil(t, err, "Expected invalid SPIFFE ID: %s", tt.input)
			assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
		})
	}
}

// TestValidatePathPattern_Valid tests ValidatePathPattern with valid patterns
func TestValidatePathPattern_Valid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"SimplePath", "app/secrets"},
		{"WithWildcard", "app/.*/secrets"},
		{"WithPlus", "app/secret.+"},
		{"WithQuestion", "app/secret.?"},
		{"WithCharClass", "app/[a-z]+"},
		{"WithParens", "app/(prod|dev)/secrets"},
		{"WithUnderscore", "app/secret_name"},
		{"WithDot", "app/secret.json"},
		{"WithDash", "app/secret-name"},
		{"WithAnchor", "^app/secrets$"},
		{"WithPipe", "app|config"},
		{"WithBraces", "app/{id}"},
		{"WithBrackets", "app[0-9]"},
		{"MaxLength", strings.Repeat("a", 500)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePathPattern(tt.input)
			assert.Nil(t, err, "Expected valid path pattern: %s", tt.input)
		})
	}
}

// TestValidatePathPattern_Invalid tests ValidatePathPattern with invalid patterns
func TestValidatePathPattern_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"TooLong", strings.Repeat("a", 501)},
		{"WithSpace", "app/secret name"},
		{"WithAt", "app/@secret"},
		{"WithHash", "app/#secret"},
		{"WithPercent", "app/%secret"},
		{"WithAmpersand", "app/&secret"},
		{"WithEquals", "app/secret=value"},
		{"WithColon", "app/secret:value"},
		{"WithSemicolon", "app/secret;value"},
		{"WithComma", "app/secret,value"},
		{"WithLessThan", "app/<secret"},
		{"WithGreaterThan", "app/>secret"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePathPattern(tt.input)
			assert.NotNil(t, err, "Expected invalid path pattern: %s", tt.input)
			assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
		})
	}
}

// TestValidatePath_Valid tests ValidatePath with valid paths
func TestValidatePath_Valid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"SimplePath", "app/secrets"},
		{"WithUnderscore", "app/secret_name"},
		{"WithDot", "app/secret.json"},
		{"WithDash", "app/secret-name"},
		{"Nested", "app/prod/database/credentials"},
		{"WithNumbers", "app/secret123"},
		{"SingleLevel", "secrets"},
		{"WithParentheses", "app/(prod)/secrets"},
		{"WithBrackets", "app/[index]/secrets"},
		{"WithBraces", "app/{id}/secrets"},
		{"WithPipe", "app|config"},
		{"WithBackslash", "app\\path"},
		{"WithPlus", "app+secrets"},
		{"WithAsterisk", "app*secrets"},
		{"WithQuestion", "app?secrets"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePath(tt.input)
			assert.Nil(t, err, "Expected valid path: %s", tt.input)
		})
	}
}

// TestValidatePath_Invalid tests ValidatePath with invalid paths
func TestValidatePath_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"TooLong", strings.Repeat("a", 501)},
		{"WithSpace", "app/secret name"},
		{"WithAt", "app/@secret"},
		{"WithHash", "app/#secret"},
		{"WithPercent", "app/%secret"},
		{"WithAmpersand", "app/&secret"},
		{"WithEquals", "app/secret=value"},
		{"WithColon", "app/secret:value"},
		{"WithSemicolon", "app/secret;value"},
		{"WithComma", "app/secret,value"},
		{"WithLessThan", "app/<secret"},
		{"WithGreaterThan", "app/>secret"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePath(tt.input)
			assert.NotNil(t, err, "Expected invalid path: %s", tt.input)
			assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
		})
	}
}

// TestValidatePolicyID_Valid tests ValidatePolicyID with valid UUIDs
func TestValidatePolicyID_Valid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"UUID_v4", "550e8400-e29b-41d4-a716-446655440000"},
		{"UUID_v1", "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
		{"UUID_v5", "886313e1-3b8a-5372-9b90-0c9aee199e5d"},
		{"AllZeros", "00000000-0000-0000-0000-000000000000"},
		{"AllOnes", "11111111-1111-1111-1111-111111111111"},
		{"MixedCase", "A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11"},
		{"Lowercase", "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePolicyID(tt.input)
			assert.Nil(t, err, "Expected valid policy ID: %s", tt.input)
		})
	}
}

// TestValidatePolicyID_Invalid tests ValidatePolicyID with invalid UUIDs
func TestValidatePolicyID_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"TooShort", "550e8400-e29b-41d4-a716"},
		{"TooLong", "550e8400-e29b-41d4-a716-446655440000-extra"},
		{"WrongFormat", "550e8400-e29b-41d4-a716-44665544000"},
		{"InvalidChars", "550e8400-e29b-41d4-a716-44665544000g"},
		{"Spaces", "550e8400 e29b 41d4 a716 446655440000"},
		{"MissingSegment", "550e8400--41d4-a716-446655440000"},
		{"NotUUID", "not-a-uuid"},
		{"JustHyphens", "----"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePolicyID(tt.input)
			if err == nil {
				t.Fatalf("Expected invalid policy ID: %s, but got nil error", tt.input)
			}
			assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
		})
	}
}

// TestValidatePermissions_Valid tests ValidatePermissions with valid permissions
func TestValidatePermissions_Valid(t *testing.T) {
	tests := []struct {
		name  string
		input []data.PolicyPermission
	}{
		{"SingleRead", []data.PolicyPermission{data.PermissionRead}},
		{"SingleWrite", []data.PolicyPermission{data.PermissionWrite}},
		{"SingleList", []data.PolicyPermission{data.PermissionList}},
		{"SingleSuper", []data.PolicyPermission{data.PermissionSuper}},
		{"ReadWrite", []data.PolicyPermission{data.PermissionRead, data.PermissionWrite}},
		{"AllPermissions", []data.PolicyPermission{
			data.PermissionRead,
			data.PermissionWrite,
			data.PermissionList,
			data.PermissionSuper,
		}},
		{"Empty", []data.PolicyPermission{}},
		{"Duplicates", []data.PolicyPermission{
			data.PermissionRead,
			data.PermissionRead,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePermissions(tt.input)
			assert.Nil(t, err, "Expected valid permissions")
		})
	}
}

// TestValidatePermissions_Invalid tests ValidatePermissions with invalid permissions
func TestValidatePermissions_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input []data.PolicyPermission
	}{
		{"Invalid", []data.PolicyPermission{"invalid"}},
		{"MixedWithInvalid", []data.PolicyPermission{
			data.PermissionRead,
			"invalid",
		}},
		{"UnknownPermission", []data.PolicyPermission{"execute"}},
		{"Typo", []data.PolicyPermission{"reed"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePermissions(tt.input)
			assert.NotNil(t, err, "Expected invalid permissions")
			assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
		})
	}
}

// TestValidateName_BoundaryConditions tests boundary conditions for name validation
func TestValidateName_BoundaryConditions(t *testing.T) {
	// Test boundary at 250 characters
	exactly250 := strings.Repeat("a", 250)
	err := ValidateName(exactly250)
	assert.Nil(t, err)

	// Test boundary at 251 characters
	exactly251 := strings.Repeat("a", 251)
	err = ValidateName(exactly251)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
}

// TestValidatePathPattern_BoundaryConditions tests boundary conditions for path pattern validation
func TestValidatePathPattern_BoundaryConditions(t *testing.T) {
	// Test boundary at 500 characters
	exactly500 := strings.Repeat("a", 500)
	err := ValidatePathPattern(exactly500)
	assert.Nil(t, err)

	// Test boundary at 501 characters
	exactly501 := strings.Repeat("a", 501)
	err = ValidatePathPattern(exactly501)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
}

// TestValidatePath_BoundaryConditions tests boundary conditions for path validation
func TestValidatePath_BoundaryConditions(t *testing.T) {
	// Test boundary at 500 characters
	exactly500 := strings.Repeat("a", 500)
	err := ValidatePath(exactly500)
	assert.Nil(t, err)

	// Test boundary at 501 characters
	exactly501 := strings.Repeat("a", 501)
	err = ValidatePath(exactly501)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
}

// TestValidationConstants tests that constants are defined correctly
func TestValidationConstants(t *testing.T) {
	assert.Equal(t, 250, maxNameLength)
	assert.Equal(t, 500, maxPathPatternLength)
}
