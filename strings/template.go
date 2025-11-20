//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package strings

import (
	"regexp"
	"strconv"
	stdstrings "strings"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
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
// # Security: Fatal Exit on CSPRNG Failure
//
// This function uses crypto/rand.Read() for generating random characters. If the
// cryptographic random number generator fails, this function will terminate the
// program with log.FatalErr() rather than returning an error.
//
// This design decision is intentional and critical for security:
//
//  1. CSPRNG failures indicate fundamental system compromise or misconfiguration
//  2. This function is used for generating security-sensitive strings (passwords,
//     tokens, API keys, secrets) where weak randomness would be catastrophic
//  3. Silently falling back to weaker randomness or returning an error that could
//     be ignored would create a false sense of security
//  4. A CSPRNG failure is an exceptional, unrecoverable system-level error
//
// DO NOT modify this behavior to return errors for CSPRNG failures, as it would
// compromise the security guarantees of all code using this function.
//
// # Parameters
//
// template: A string containing literal text and generator expressions.
//
//	Generator expressions are replaced with random characters.
//
// # Returns
//
// Returns:
//   - string: The generated string with all generator expressions replaced,
//     empty on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrStringInvalidLength: if length specification is not a valid number
//   - ErrStringNegativeLength: if length is negative
//   - ErrStringEmptyCharacterClass: if character class is empty
//   - ErrStringInvalidRange: if character range is invalid
//   - ErrStringEmptyCharacterSet: if character set is empty
//
// Note: CSPRNG failures (crypto/rand.Read) cause immediate program termination
// via log.FatalErr() and do not return an error.
func StringFromTemplate(template string) (string, *sdkErrors.SDKError) {
	// Regular expression to match generator expressions like [a-z]{5} or [\w]{3}
	// Modified to capture any content in braces, not just digits
	// Changed + to * to allow empty character classes like []
	re := regexp.MustCompile(`\[([^]]*)]\{([^}]+)}`)

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

		// Parse length - this will now catch non-numeric values
		length, parseErr := strconv.Atoi(lengthStr)
		if parseErr != nil {
			failErr := sdkErrors.ErrStringInvalidLength.Wrap(parseErr)
			failErr.Msg = "invalid length specification in template"
			return "", failErr
		}

		// Validate that length is non-negative
		if length < 0 {
			failErr := sdkErrors.ErrStringNegativeLength
			failErr.Msg = "length cannot be negative in template"
			return "", failErr
		}

		// Generate random string based on character class
		randomStr, err := secureRandomStringFromCharClass(charClass, length)
		if err != nil {
			return "", err
		}

		// Replace the first occurrence of the pattern
		result = stdstrings.Replace(result, fullMatch, randomStr, 1)
	}

	return result, nil
}
