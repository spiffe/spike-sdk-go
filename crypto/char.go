//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"fmt"
)

// secureRandomStringFromCharClass generates a cryptographically secure random
// string of the specified length using characters from the given character
// class.
//
// Parameters:
//   - charClass: character class specification supporting:
//   - Predefined classes: \w (word chars), \d (digits), \x (symbols)
//   - Custom ranges: "A-Z", "a-z", "0-9", or combinations like "A-Za-z0-9"
//   - Individual characters: any literal characters
//   - length: number of characters in the resulting string
//
// Returns:
//   - string: the generated random string
//   - error: non-nil if the character class is empty, invalid, results in
//     an empty character set, or if the cryptographic random number generator
//     fails
func secureRandomStringFromCharClass(
	charClass string, length int,
) (string, error) {
	chars, err := expandCharacterClass(charClass)
	if err != nil {
		return "", err
	}

	if len(chars) == 0 {
		return "", fmt.Errorf("empty character set")
	}

	result := make([]byte, length)
	randomBytes := make([]byte, length)
	if _, err := reader(randomBytes); err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		result[i] = chars[randomBytes[i]%byte(len(chars))]
	}

	return string(result), nil
}

// expandCharacterClass expands character class expressions into a string
// containing all valid characters from the class. It handles both predefined
// character classes and custom character ranges.
//
// Parameters:
//   - charClass: character class expression supporting:
//   - Predefined classes:
//   - \w: word characters (a-z, A-Z, 0-9, and underscore)
//   - \d: digits (0-9)
//   - \x: symbols (printable ASCII excluding letters and digits)
//   - Custom ranges:
//   - Single characters: included as-is
//   - Range notation: "A-Z" expands to all uppercase letters
//   - Combined ranges: "A-Za-z0-9" expands to alphanumeric characters
//
// Returns:
//   - string: expanded character set containing all characters from the class
//   - error: non-nil if the character class is empty, contains invalid range
//     specifications (e.g., "Z-A"), or results in an empty character set
func expandCharacterClass(charClass string) (string, error) {
	// Check for empty character class first
	if len(charClass) == 0 {
		return "", fmt.Errorf("empty character class")
	}

	charSet := make(map[byte]bool) // Use map to avoid duplicates

	// Handle predefined character classes
	switch charClass {
	case "\\w":
		// Word characters: letters, digits, underscore
		for c := 'a'; c <= 'z'; c++ {
			charSet[byte(c)] = true
		}
		for c := 'A'; c <= 'Z'; c++ {
			charSet[byte(c)] = true
		}
		for c := '0'; c <= '9'; c++ {
			charSet[byte(c)] = true
		}
		charSet['_'] = true
	case "\\d":
		// Digits
		for c := '0'; c <= '9'; c++ {
			charSet[byte(c)] = true
		}
	case "\\x":
		// Symbols (printable ASCII excluding letters and digits)
		for c := 32; c <= 126; c++ {
			ch := byte(c)
			if !((ch >= 'a' && ch <= 'z') ||
				(ch >= 'A' && ch <= 'Z') ||
				(ch >= '0' && ch <= '9')) {
				charSet[ch] = true
			}
		}
	default:
		// Handle character ranges and individual characters like A-Za-z0-9
		i := 0
		for i < len(charClass) {
			if i+2 < len(charClass) && charClass[i+1] == '-' {
				// Range specification
				start := charClass[i]
				end := charClass[i+2]

				// Only allow forward ranges (`start <= end`)
				if start > end {
					return "",
						fmt.Errorf("invalid range specified: %c-%c", start, end)
				}

				// Add all characters in range
				for c := start; c <= end; c++ {
					charSet[c] = true
				}
				i += 3
			} else {
				// Single character
				charSet[charClass[i]] = true
				i++
			}
		}
	}

	// Convert map to slice
	chars := make([]byte, 0, len(charSet))
	for char := range charSet {
		chars = append(chars, char)
	}

	// Final check for the empty result (this catches edge cases)
	if len(chars) == 0 {
		return "", fmt.Errorf("character class resulted in empty character set")
	}

	return string(chars), nil
}
