//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestExpandCharacterClass_PredefinedClasses tests predefined character classes
func TestExpandCharacterClass_PredefinedClasses(t *testing.T) {
	tests := []struct {
		name          string
		charClass     string
		expectedChars map[byte]bool
	}{
		{
			name:      "WordClass",
			charClass: "\\w",
			expectedChars: func() map[byte]bool {
				m := make(map[byte]bool)
				for c := 'a'; c <= 'z'; c++ {
					m[byte(c)] = true
				}
				for c := 'A'; c <= 'Z'; c++ {
					m[byte(c)] = true
				}
				for c := '0'; c <= '9'; c++ {
					m[byte(c)] = true
				}
				m['_'] = true
				return m
			}(),
		},
		{
			name:      "DigitClass",
			charClass: "\\d",
			expectedChars: func() map[byte]bool {
				m := make(map[byte]bool)
				for c := '0'; c <= '9'; c++ {
					m[byte(c)] = true
				}
				return m
			}(),
		},
		{
			name:      "SymbolClass",
			charClass: "\\x",
			expectedChars: func() map[byte]bool {
				m := make(map[byte]bool)
				for c := 32; c <= 126; c++ {
					ch := byte(c)
					if !((ch >= 'a' && ch <= 'z') ||
						(ch >= 'A' && ch <= 'Z') ||
						(ch >= '0' && ch <= '9')) {
						m[ch] = true
					}
				}
				return m
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandCharacterClass(tt.charClass)
			assert.Nil(t, err)
			assert.NotEmpty(t, result)

			// Verify all expected characters are present
			for expectedChar := range tt.expectedChars {
				assert.Contains(t, result, string(expectedChar),
					"Expected character %c (%d) to be in result", expectedChar, expectedChar)
			}

			// Verify no unexpected characters
			for _, c := range result {
				assert.True(t, tt.expectedChars[byte(c)],
					"Unexpected character %c (%d) in result", c, c)
			}
		})
	}
}

// TestExpandCharacterClass_CustomRanges tests custom character ranges
func TestExpandCharacterClass_CustomRanges(t *testing.T) {
	tests := []struct {
		name           string
		charClass      string
		expectedLength int
		expectedStart  byte
		expectedEnd    byte
	}{
		{"UppercaseRange", "A-Z", 26, 'A', 'Z'},
		{"LowercaseRange", "a-z", 26, 'a', 'z'},
		{"DigitRange", "0-9", 10, '0', '9'},
		{"SmallRange", "A-C", 3, 'A', 'C'},
		{"SingleCharRange", "A-A", 1, 'A', 'A'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandCharacterClass(tt.charClass)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedLength, len(result))

			// Verify range
			for c := tt.expectedStart; c <= tt.expectedEnd; c++ {
				assert.Contains(t, result, string(c))
			}
		})
	}
}

// TestExpandCharacterClass_CombinedRanges tests combined character ranges
func TestExpandCharacterClass_CombinedRanges(t *testing.T) {
	tests := []struct {
		name           string
		charClass      string
		expectedLength int
	}{
		{"AlphanumericLower", "a-z0-9", 36},
		{"AlphanumericUpper", "A-Z0-9", 36},
		{"AlphanumericBoth", "A-Za-z0-9", 62},
		{"MultipleRanges", "A-C0-2", 6}, // ABC012
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandCharacterClass(tt.charClass)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedLength, len(result))
		})
	}
}

// TestExpandCharacterClass_SingleCharacters tests individual characters
func TestExpandCharacterClass_SingleCharacters(t *testing.T) {
	tests := []struct {
		name      string
		charClass string
		expected  string
	}{
		{"SingleChar", "A", "A"},
		{"MultipleChars", "ABC", "ABC"},
		{"MixedCharsAndRange", "XA-CY", "XABCY"},
		{"SpecialChars", "!@#", "!@#"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandCharacterClass(tt.charClass)
			assert.Nil(t, err)
			assert.Equal(t, len(tt.expected), len(result))

			// Check all expected characters are present
			for _, c := range tt.expected {
				assert.Contains(t, result, string(c))
			}
		})
	}
}

// TestExpandCharacterClass_EmptyCharClass tests empty character class
func TestExpandCharacterClass_EmptyCharClass(t *testing.T) {
	result, err := expandCharacterClass("")
	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrStringEmptyCharacterClass))
}

// TestExpandCharacterClass_InvalidRange tests invalid character ranges
func TestExpandCharacterClass_InvalidRange(t *testing.T) {
	tests := []struct {
		name      string
		charClass string
	}{
		{"BackwardRange", "Z-A"},
		{"BackwardDigitRange", "9-0"},
		{"BackwardLowerRange", "z-a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandCharacterClass(tt.charClass)
			assert.Empty(t, result)
			assert.NotNil(t, err)
			assert.True(t, err.Is(sdkErrors.ErrStringInvalidRange))
		})
	}
}

// TestExpandCharacterClass_OverlappingRanges tests overlapping ranges
func TestExpandCharacterClass_OverlappingRanges(t *testing.T) {
	// Overlapping ranges should not produce duplicates
	result, err := expandCharacterClass("A-CA-C")
	assert.Nil(t, err)
	// Should have only 3 unique characters: A, B, C
	assert.Equal(t, 3, len(result))
	assert.Contains(t, result, "A")
	assert.Contains(t, result, "B")
	assert.Contains(t, result, "C")
}

// TestExpandCharacterClass_EdgeCases tests edge cases
func TestExpandCharacterClass_EdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		charClass string
		wantErr   bool
	}{
		{"DashAtEnd", "ABC-", false},     // Dash at end is treated as literal
		{"DashAtStart", "-ABC", false},   // Dash at start is treated as literal
		{"OnlyDash", "-", false},         // Single dash is treated as literal
		{"MultipleDashes", "---", false}, // Multiple dashes treated as literals
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandCharacterClass(tt.charClass)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Empty(t, result)
			} else {
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
			}
		})
	}
}

// TestSecureRandomStringFromCharClass_ValidClasses tests valid character classes
func TestSecureRandomStringFromCharClass_ValidClasses(t *testing.T) {
	tests := []struct {
		name      string
		charClass string
		length    int
	}{
		{"WordClass10", "\\w", 10},
		{"DigitClass10", "\\d", 10},
		{"SymbolClass10", "\\x", 10},
		{"UppercaseRange", "A-Z", 20},
		{"LowercaseRange", "a-z", 20},
		{"Alphanumeric", "A-Za-z0-9", 30},
		{"SingleChar", "A", 5},
		{"SpecialChars", "!@#$%", 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := secureRandomStringFromCharClass(tt.charClass, tt.length)
			assert.Nil(t, err)
			assert.Equal(t, tt.length, len(result))

			// Verify all characters are from the expanded char class
			expandedChars, expandErr := expandCharacterClass(tt.charClass)
			require.Nil(t, expandErr)
			for _, c := range result {
				assert.Contains(t, expandedChars, string(c),
					"Character %c should be from character class %s", c, tt.charClass)
			}
		})
	}
}

// TestSecureRandomStringFromCharClass_DifferentLengths tests different string lengths
func TestSecureRandomStringFromCharClass_DifferentLengths(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"Length0", 0},
		{"Length1", 1},
		{"Length10", 10},
		{"Length100", 100},
		{"Length1000", 1000},
	}

	charClass := "A-Za-z0-9"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := secureRandomStringFromCharClass(charClass, tt.length)
			assert.Nil(t, err)
			assert.Equal(t, tt.length, len(result))
		})
	}
}

// TestSecureRandomStringFromCharClass_EmptyCharClass tests empty character class
func TestSecureRandomStringFromCharClass_EmptyCharClass(t *testing.T) {
	result, err := secureRandomStringFromCharClass("", 10)
	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrStringEmptyCharacterClass))
}

// TestSecureRandomStringFromCharClass_InvalidRange tests invalid character range
func TestSecureRandomStringFromCharClass_InvalidRange(t *testing.T) {
	result, err := secureRandomStringFromCharClass("Z-A", 10)
	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrStringInvalidRange))
}

// TestSecureRandomStringFromCharClass_Randomness tests that multiple calls produce different results
func TestSecureRandomStringFromCharClass_Randomness(t *testing.T) {
	charClass := "A-Za-z0-9"
	length := 20

	// Generate multiple strings
	results := make(map[string]bool)
	iterations := 10

	for i := 0; i < iterations; i++ {
		result, err := secureRandomStringFromCharClass(charClass, length)
		require.Nil(t, err)
		results[result] = true
	}

	// All results should be unique (extremely unlikely to have duplicates with 62^20 possibilities)
	assert.Equal(t, iterations, len(results), "Expected all generated strings to be unique")
}

// TestSecureRandomStringFromCharClass_CharacterDistribution tests character distribution
func TestSecureRandomStringFromCharClass_CharacterDistribution(t *testing.T) {
	charClass := "ABC"
	length := 300

	result, err := secureRandomStringFromCharClass(charClass, length)
	require.Nil(t, err)

	// Count occurrences of each character
	counts := make(map[rune]int)
	for _, c := range result {
		counts[c]++
	}

	// With 3 characters and 300 iterations, we expect roughly 100 of each
	// Allow for statistical variation (at least 50 of each)
	assert.GreaterOrEqual(t, counts['A'], 50, "Character A should appear at least 50 times")
	assert.GreaterOrEqual(t, counts['B'], 50, "Character B should appear at least 50 times")
	assert.GreaterOrEqual(t, counts['C'], 50, "Character C should appear at least 50 times")
}

// TestExpandCharacterClass_Consistency tests that same input produces same output
func TestExpandCharacterClass_Consistency(t *testing.T) {
	charClass := "A-Za-z0-9"

	result1, err1 := expandCharacterClass(charClass)
	result2, err2 := expandCharacterClass(charClass)

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, len(result1), len(result2))

	// Convert to maps for comparison (order doesn't matter)
	map1 := make(map[byte]bool)
	for _, c := range result1 {
		map1[byte(c)] = true
	}
	map2 := make(map[byte]bool)
	for _, c := range result2 {
		map2[byte(c)] = true
	}

	assert.Equal(t, map1, map2, "Same input should produce same character set")
}

// TestExpandCharacterClass_WordClassSize tests that \w has correct size
func TestExpandCharacterClass_WordClassSize(t *testing.T) {
	result, err := expandCharacterClass("\\w")
	assert.Nil(t, err)
	// \w = a-z (26) + A-Z (26) + 0-9 (10) + _ (1) = 63 characters
	assert.Equal(t, 63, len(result))
}

// TestExpandCharacterClass_DigitClassSize tests that \d has correct size
func TestExpandCharacterClass_DigitClassSize(t *testing.T) {
	result, err := expandCharacterClass("\\d")
	assert.Nil(t, err)
	// \d = 0-9 = 10 characters
	assert.Equal(t, 10, len(result))
}

// TestExpandCharacterClass_SymbolClassSize tests that \x has correct size
func TestExpandCharacterClass_SymbolClassSize(t *testing.T) {
	result, err := expandCharacterClass("\\x")
	assert.Nil(t, err)
	// \x = printable ASCII (95) - letters (52) - digits (10) = 33 characters
	assert.Equal(t, 33, len(result))
}

// TestSecureRandomStringFromCharClass_AllPredefinedClasses tests all predefined classes work
func TestSecureRandomStringFromCharClass_AllPredefinedClasses(t *testing.T) {
	tests := []string{"\\w", "\\d", "\\x"}

	for _, charClass := range tests {
		t.Run(charClass, func(t *testing.T) {
			result, err := secureRandomStringFromCharClass(charClass, 20)
			assert.Nil(t, err)
			assert.Equal(t, 20, len(result))
		})
	}
}
