//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package validation provides input validation utilities for SPIKE API
// operations.
//
// The package includes validation functions for:
//   - Names: alphanumeric strings with length and format constraints
//   - SPIFFE IDs: both exact IDs and regex patterns for identity matching
//   - Paths: resource paths with support for wildcards and special characters
//   - Path patterns: regex patterns for path matching in policies
//   - Policy IDs: UUID format validation
//   - Permissions: verification against allowed permission types
//
// All validation functions return errors.ErrInvalidInput when validation fails.
//
// Name Validation:
//
// Names must be 1-250 characters and contain only alphanumeric characters,
// hyphens, underscores, and spaces:
//
//	if err := validation.ValidateName("my-policy-name"); err != nil {
//	    log.Fatal("Invalid name")
//	}
//
// SPIFFE ID Validation:
//
// Validates both raw SPIFFE IDs and regex patterns:
//
//	// Raw SPIFFE ID
//	err := validation.ValidateSPIFFEID("spiffe://example.org/service/api")
//
//	// SPIFFE ID pattern with wildcards
//	err = validation.ValidateSPIFFEIDPattern("^spiffe://example.org/.*$")
//
// Path Validation:
//
// Validates resource paths and path patterns:
//
//	// Simple path
//	err := validation.ValidatePath("app/secrets/database")
//
//	// Path pattern with wildcards
//	err = validation.ValidatePathPattern("app/.*/database")
//
// Permission Validation:
//
// Verifies that permissions are in the allowed set:
//
//	permissions := []data.PolicyPermission{
//	    data.PermissionRead,
//	    data.PermissionWrite,
//	}
//	if err := validation.ValidatePermissions(permissions); err != nil {
//	    log.Fatal("Invalid permissions")
//	}
package validation
