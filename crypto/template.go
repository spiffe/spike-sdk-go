//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// StringFromTemplate creates a string based on a template with embedded
// generator expressions.
//
// # Template Syntax
//
// Generator expressions follow the pattern: [character_class]{length}
// Where:
//   - character_class defines which characters can be generated
//   - length specifies how many characters to generate (must be a positive
//     integer)
//
// # Supported Character Classes
//
// ## Predefined Classes
//
//   - \w  : Word characters (a-z, A-Z, 0-9, _)
//   - \d  : Digits (0-9)
//   - \x  : Symbols (printable ASCII excluding letters and digits:
//     !"#$%&'()*+,-./:;<=>?@[\]^`{|}~ and space)
//
// ## Character Ranges
//
//   - a-z : Lowercase letters
//   - A-Z : Uppercase letters
//   - 0-9 : Digits
//   - a-Z : All letters (equivalent to a-zA-Z)
//
// ## Multiple Ranges and Characters
//
// You can combine multiple ranges and individual characters within a
// single class:
//   - [a-zA-Z0-9]  : Letters and digits
//   - [A-Za-z0-6]  : Letters and digits 0-6
//   - [0-9a-fA-F]  : Hexadecimal characters
//   - [A-Ca-c1-3]  : A,B,C,a,b,c,1,2,3
//
// Individual characters can be mixed with ranges:
//   - [a-z_.-]     : Lowercase letters plus underscore, period, and hyphen
//   - [A-Z0-9!@#]  : Uppercase letters, digits, and specific symbols
//
// # Template Examples
//
//	StringFromTemplate("user[0-9]{4}")                   // "user1234"
//	StringFromTemplate("pass[a-zA-Z0-9]{12}")            // "passA3kL9mX2nQ8z"
//	StringFromTemplate("prefix[\w]{8}suffix")            // "prefixaB3_kM9Zsuffix"
//	StringFromTemplate("id[0-9a-f]{8}-[0-9a-f]{4}")      // "a1b2c3d4-ef56"
//	StringFromTemplate("admin[a-z]{3}[A-Z]{2}[0-9]{3}")  // "adminxyzAB123"
//
// # Error Conditions
//
// The function returns an error for:
//   - Invalid ranges where start > end: [z-a] or [9-0]
//   - Empty character classes: []
//   - Invalid length specifications: non-numeric values
//   - Malformed expressions: missing brackets or braces
//
// # Implementation Notes
//
// Character ranges are inclusive on both ends. When multiple ranges overlap
// (e.g., [a-zA-Z] contains both a-z and A-Z), duplicate characters are
// automatically deduplicated.
//
// Ranges must follow ASCII ordering. Cross-case ranges like [a-Z] work because
// they span the ASCII range from 'a' (97) to 'Z' (90), but this includes
// punctuation characters between uppercase and lowercase letters.
//
// # Limitations and Assumptions
//
// This implementation assumes reasonable usage patterns:
//   - Character classes should be logically organized
//   - Ranges should follow natural ordering (a-z, not z-a)
//   - Individual characters mixed with ranges are supported but should be
//     used judiciously
//   - Unicode characters beyond ASCII are not explicitly supported
//   - Escape sequences beyond \w, \d, \x are not supported
//   - Character class negation (^) is not supported
//   - POSIX character classes ([:alpha:], [:digit:]) are not supported
//
// The function prioritizes common use cases for password generation, API keys,
// tokens, and identifiers while maintaining simplicity and predictability.
//
// # Parameters
//
// template: A string containing literal text and generator expressions.
//
//	Generator expressions are replaced with random characters.
//
// # Returns
//
// Returns the generated string with all generator expressions replaced by
// random characters matching their specifications, or an error if any
// generator expression is invalid.
func StringFromTemplate(template string) (string, error) {
	// Regular expression to match generator expressions like [a-z]{5} or [\w]{3}
	re := regexp.MustCompile(`\[([^]]+)]\{(\d+)}`)

	result := template

	// Find all matches and replace them
	for {
		match := re.FindStringSubmatch(result)
		if match == nil {
			break
		}

		fullMatch := match[0]
		charClass := match[1]
		lengthStr := match[2]

		// Parse length
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return "", fmt.Errorf("invalid length: %s", lengthStr)
		}

		// StringFromTemplate random string based on character class
		randomStr, err := secureRandomStringFromCharClass(charClass, length)
		if err != nil {
			return "", err
		}

		// Replace the first occurrence of the pattern
		result = strings.Replace(result, fullMatch, randomStr, 1)
	}

	return result, nil
}
