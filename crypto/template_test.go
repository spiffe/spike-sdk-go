//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"regexp"
	"strings"
	"testing"
)

func TestStringFromTemplate(t *testing.T) {
	tests := []struct {
		name        string
		template    string
		expectError bool
		validate    func(t *testing.T, result string) // Custom validation function
	}{
		{
			name:     "simple word characters",
			template: `football[\w]{8}bartender`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "football") {
					t.Errorf("result should start with 'football', got: %s", result)
				}
				if !strings.HasSuffix(result, "bartender") {
					t.Errorf("result should end with 'bartender', got: %s", result)
				}
				// Extract the middle part and check it's 8 word characters
				middle := result[8 : len(result)-9] // football(8chars)bartender
				if len(middle) != 8 {
					t.Errorf("middle part should be 8 characters, got %d: %s", len(middle), middle)
				}
				matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]{8}$`, middle)
				if !matched {
					t.Errorf("middle part should contain only word characters, got: %s", middle)
				}
			},
		},
		{
			name:     "lowercase letters and digits",
			template: `admin[a-z0-9]{3}`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "admin") {
					t.Errorf("result should start with 'admin', got: %s", result)
				}
				suffix := result[5:] // Everything after "admin"
				if len(suffix) != 3 {
					t.Errorf("suffix should be 3 characters, got %d: %s", len(suffix), suffix)
				}
				matched, _ := regexp.MatchString(`^[a-z0-9]{3}$`, suffix)
				if !matched {
					t.Errorf("suffix should contain only lowercase letters and digits, got: %s", suffix)
				}
			},
		},
		{
			name:     "multiple patterns",
			template: `admin[a-z0-9]{3}something[\w]{3}`,
			validate: func(t *testing.T, result string) {
				// Should be: admin + 3 chars + something + 3 chars
				expectedLen := 5 + 3 + 9 + 3 // admin(5) + random(3) + something(9) + random(3)
				if len(result) != expectedLen {
					t.Errorf("result should be %d characters, got %d: %s", expectedLen, len(result), result)
				}
				if !strings.HasPrefix(result, "admin") {
					t.Errorf("result should start with 'admin', got: %s", result)
				}
				if !strings.Contains(result, "something") {
					t.Errorf("result should contain 'something', got: %s", result)
				}
			},
		},
		{
			name:     "letters and digits mixed",
			template: `pass[a-zA-Z0-9]{12}`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "pass") {
					t.Errorf("result should start with 'pass', got: %s", result)
				}
				suffix := result[4:]
				if len(suffix) != 12 {
					t.Errorf("suffix should be 12 characters, got %d: %s", len(suffix), suffix)
				}
				matched, _ := regexp.MatchString(`^[a-zA-Z0-9]{12}$`, suffix)
				if !matched {
					t.Errorf("suffix should contain only letters and digits, got: %s", suffix)
				}
			},
		},
		{
			name:        "cross-case range a-Z",
			template:    `fail[a-Z]{8}`,
			expectError: true,
		},
		{
			name:     "mixed case letters",
			template: `pass[A-Za-z]{8}`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "pass") {
					t.Errorf("result should start with 'pass', got: %s", result)
				}
				suffix := result[4:]
				if len(suffix) != 8 {
					t.Errorf("suffix should be 8 characters, got %d: %s", len(suffix), suffix)
				}
				matched, _ := regexp.MatchString(`^[A-Za-z]{8}$`, suffix)
				if !matched {
					t.Errorf("suffix should contain only letters, got: %s", suffix)
				}
			},
		},
		{
			name:     "digits only",
			template: `football[\d]{8}bartender`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "football") {
					t.Errorf("result should start with 'football', got: %s", result)
				}
				if !strings.HasSuffix(result, "bartender") {
					t.Errorf("result should end with 'bartender', got: %s", result)
				}
				middle := result[8 : len(result)-9]
				if len(middle) != 8 {
					t.Errorf("middle part should be 8 characters, got %d: %s", len(middle), middle)
				}
				matched, _ := regexp.MatchString(`^[0-9]{8}$`, middle)
				if !matched {
					t.Errorf("middle part should contain only digits, got: %s", middle)
				}
			},
		},
		{
			name:     "symbols only",
			template: `football[\x]{4}`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "football") {
					t.Errorf("result should start with 'football', got: %s", result)
				}
				suffix := result[8:]
				if len(suffix) != 4 {
					t.Errorf("suffix should be 4 characters, got %d: %s", len(suffix), suffix)
				}
				// \x should be symbols (printable ASCII excluding letters and digits)
				for _, char := range suffix {
					if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
						t.Errorf("suffix should contain only symbols, but found alphanumeric: %c in %s", char, suffix)
					}
				}
			},
		},
		{
			name:     "multiple different patterns",
			template: `user[A-Z]{2}[0-9]{4}`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "user") {
					t.Errorf("result should start with 'user', got: %s", result)
				}
				suffix := result[4:]
				if len(suffix) != 6 { // 2 uppercase + 4 digits
					t.Errorf("suffix should be 6 characters, got %d: %s", len(suffix), suffix)
				}
				// The first 2 should be uppercase letters
				upperPart := suffix[:2]
				matched, _ := regexp.MatchString(`^[A-Z]{2}$`, upperPart)
				if !matched {
					t.Errorf("first part should be uppercase letters, got: %s", upperPart)
				}
				// The last 4 should be digits
				digitPart := suffix[2:]
				matched, _ = regexp.MatchString(`^[0-9]{4}$`, digitPart)
				if !matched {
					t.Errorf("second part should be digits, got: %s", digitPart)
				}
			},
		},
		{
			name:     "no patterns",
			template: `simple`,
			validate: func(t *testing.T, result string) {
				if result != "simple" {
					t.Errorf("result should be 'simple', got: %s", result)
				}
			},
		},
		{
			name:     "pattern in middle",
			template: `prefix[\w]{10}suffix`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "prefix") {
					t.Errorf("result should start with 'prefix', got: %s", result)
				}
				if !strings.HasSuffix(result, "suffix") {
					t.Errorf("result should end with 'suffix', got: %s", result)
				}
				middle := result[6 : len(result)-6] // prefix(6chars)suffix(6chars)
				if len(middle) != 10 {
					t.Errorf("middle part should be 10 characters, got %d: %s", len(middle), middle)
				}
			},
		},
		{
			name:     "hex characters",
			template: `test[0-9a-fA-F]{8}`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "test") {
					t.Errorf("result should start with 'test', got: %s", result)
				}
				suffix := result[4:]
				if len(suffix) != 8 {
					t.Errorf("suffix should be 8 characters, got %d: %s", len(suffix), suffix)
				}
				matched, _ := regexp.MatchString(`^[0-9a-fA-F]{8}$`, suffix)
				if !matched {
					t.Errorf("suffix should contain only hex characters, got: %s", suffix)
				}
			},
		},
		{
			name:     "multiple small ranges",
			template: `prefix[A-Ca-c1-3]{5}suffix`,
			validate: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "prefix") {
					t.Errorf("result should start with 'prefix', got: %s", result)
				}
				if !strings.HasSuffix(result, "suffix") {
					t.Errorf("result should end with 'suffix', got: %s", result)
				}
				middle := result[6 : len(result)-6]
				if len(middle) != 5 {
					t.Errorf("middle part should be 5 characters, got %d: %s", len(middle), middle)
				}
				// Should only contain A, B, C, a, b, c, 1, 2, 3.
				matched, _ := regexp.MatchString(`^[ABCabc123]{5}$`, middle)
				if !matched {
					t.Errorf("middle part should contain only A,B,C,a,b,c,1,2,3, got: %s", middle)
				}
			},
		},
		// Error cases
		{
			name:        "invalid range z-a",
			template:    `pass[z-a]{8}`,
			expectError: true,
		},
		{
			name:        "empty character class",
			template:    `[]{5}`,
			expectError: true,
		},
		{
			name:     "zero length",
			template: `[a-z]{0}`,
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("result should be empty string for zero length, got: %s", result)
				}
			},
		},
		{
			name:        "invalid length - non-numeric",
			template:    `[a-z]{abc}`,
			expectError: true,
		},
		{
			name:     "malformed - missing closing bracket",
			template: `[a-z{5}`,
			validate: func(t *testing.T, result string) {
				// This should be treated as literal text since it doesn't match the pattern
				if result != "[a-z{5}" {
					t.Errorf("malformed pattern should be treated as literal, got: %s", result)
				}
			},
		},
		{
			name:     "malformed - missing opening brace",
			template: `[a-z]5}`,
			validate: func(t *testing.T, result string) {
				// This should be treated as literal text since it doesn't match the pattern
				if result != "[a-z]5}" {
					t.Errorf("malformed pattern should be treated as literal, got: %s", result)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringFromTemplate(tt.template)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none, result: %s", result)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.validate != nil {
				tt.validate(t, result)
			}
		})
	}
}

// TestStringFromTemplateConsistency tests that the function produces different results
// on multiple calls (since it should be generating random strings)
func TestStringFromTemplateConsistency(t *testing.T) {
	template := `test[a-zA-Z0-9]{10}`

	results := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		result, err := StringFromTemplate(template)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if !strings.HasPrefix(result, "test") {
			t.Errorf("result should start with 'test', got: %s", result)
		}

		if len(result) != 14 { // "test" + 10 random chars
			t.Errorf("result should be 14 characters, got %d: %s", len(result), result)
		}

		results[result] = true
	}

	// We should have generated many different strings
	// With 10 random alphanumeric characters, the chance of collision is very low
	if len(results) < iterations/2 {
		t.Errorf("expected more unique results, got %d unique out of %d iterations", len(results), iterations)
	}
}

// TestStringFromTemplateLength tests various length specifications
func TestStringFromTemplateLength(t *testing.T) {
	testCases := []struct {
		template    string
		expectedLen int
	}{
		{`[a-z]{1}`, 1},
		{`[a-z]{5}`, 5},
		{`[a-z]{20}`, 20},
		{`prefix[a-z]{10}suffix`, 22},    // 6 + 10 + 6
		{`[a-z]{5}[A-Z]{3}[0-9]{2}`, 10}, // 5 + 3 + 2
	}

	for _, tc := range testCases {
		t.Run(tc.template, func(t *testing.T) {
			result, err := StringFromTemplate(tc.template)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(result) != tc.expectedLen {
				t.Errorf("expected length %d, got %d: %s", tc.expectedLen, len(result), result)
			}
		})
	}
}

// BenchmarkStringFromTemplate benchmarks the function performance
func BenchmarkStringFromTemplate(b *testing.B) {
	template := `user[a-zA-Z0-9]{16}[A-Z]{4}[0-9]{8}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := StringFromTemplate(template)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// TestStringFromTemplateEdgeCases tests edge cases and boundary conditions
func TestStringFromTemplateEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		template    string
		expectError bool
		validate    func(t *testing.T, result string)
	}{
		{
			name:     "empty template",
			template: "",
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("empty template should return empty string, got: %s", result)
				}
			},
		},
		{
			name:     "single character class",
			template: `[a]{1}`,
			validate: func(t *testing.T, result string) {
				if result != "a" {
					t.Errorf("single character class should return that character, got: %s", result)
				}
			},
		},
		{
			name:     "large length",
			template: `[a]{100}`,
			validate: func(t *testing.T, result string) {
				if len(result) != 100 {
					t.Errorf("result should be 100 characters, got %d", len(result))
				}
				for _, char := range result {
					if char != 'a' {
						t.Errorf("all characters should be 'a', found: %c", char)
					}
				}
			},
		},
		{
			name:     "consecutive patterns",
			template: `[a]{2}[b]{2}[c]{2}`,
			validate: func(t *testing.T, result string) {
				expected := "aabbcc"
				if result != expected {
					t.Errorf("expected %s, got %s", expected, result)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringFromTemplate(tt.template)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none, result: %s", result)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.validate != nil {
				tt.validate(t, result)
			}
		})
	}
}
