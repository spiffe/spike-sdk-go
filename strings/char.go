//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package strings

import (
	"crypto/rand"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

// secureRandomStringFromCharClass generates a cryptographically secure random
// string of the specified length using characters from the given character
// class.
//
// # Security: Fatal Exit on CSPRNG Failure
//
// This function uses crypto/rand.Read() as its source of randomness. If the
// cryptographic random number generator fails, this function will terminate
// the program with log.FatalErr() rather than returning an error.
//
// This design decision is intentional and critical for security:
//
//  1. CSPRNG failures indicate fundamental system compromise or misconfiguration
//  2. This function generates security-sensitive strings (passwords, tokens,
//     API keys, secrets) where weak randomness would be catastrophic
//  3. Silently falling back to weaker randomness or continuing execution would
//     create a false sense of security
//  4. A CSPRNG failure is an exceptional, unrecoverable system-level error
//     (kernel entropy depletion, hardware failure, or system compromise)
//  5. Consistent with other security-critical operations in the SDK (Shamir
//     secret sharing, SVID acquisition) that also fatal exit on failure
//
// DO NOT remove this fatal exit behavior. Allowing the function to return
// an error that could be ignored would compromise the security guarantees
// of all code using this function.
//
// Parameters:
//   - charClass: character class specification supporting:
//   - Predefined classes: \w (word chars), \d (digits), \x (symbols)
//   - Custom ranges: "A-Z", "a-z", "0-9", or combinations like "A-Za-z0-9"
//   - Individual characters: any literal characters
//   - length: number of characters in the resulting string
//
// Returns:
//   - string: the generated random string, empty on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrStringEmptyCharacterClass: if character class is empty
//   - ErrStringInvalidRange: if character range is invalid
//   - ErrStringEmptyCharacterSet: if the character set is empty
//
// Note: CSPRNG failures (crypto/rand.Read) cause immediate program termination
// via log.FatalErr() for security reasons (cannot generate secure random data).
// This is intentional and critical - DO NOT remove this fatal exit behavior.
func secureRandomStringFromCharClass(
	charClass string, length int,
) (string, *sdkErrors.SDKError) {
	const fName = "secureRandomStringFromCharClass"

	chars, err := expandCharacterClass(charClass)
	if err != nil {
		return "", err
	}

	if len(chars) == 0 {
		failErr := sdkErrors.ErrStringEmptyCharacterSet.Clone()
		failErr.Msg = "character class resulted in empty character set"
		return "", failErr
	}

	result := make([]byte, length)
	randomBytes := make([]byte, length)
	if _, randErr := rand.Read(randomBytes); randErr != nil {
		failErr := sdkErrors.ErrCryptoRandomGenerationFailed.Wrap(randErr)
		failErr.Msg = "cryptographic random number generator failed"
		log.FatalErr(fName, *failErr)
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
//   - string: expanded character set containing all characters from the class,
//     empty on error
//   - *sdkErrors.SDKError: nil on success, or one of the following errors:
//   - ErrStringEmptyCharacterClass: if character class is empty
//   - ErrStringInvalidRange: if character range is invalid (e.g., "Z-A")
//   - ErrStringEmptyCharacterSet: if expansion results in an empty set
func expandCharacterClass(charClass string) (string, *sdkErrors.SDKError) {
	// Check for empty character class first
	if len(charClass) == 0 {
		failErr := sdkErrors.ErrStringEmptyCharacterClass.Clone()
		failErr.Msg = "character class cannot be empty"
		return "", failErr
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
					failErr := sdkErrors.ErrStringInvalidRange.Clone()
					failErr.Msg = "invalid character range: start > end"
					return "", failErr
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
		failErr := sdkErrors.ErrStringEmptyCharacterSet.Clone()
		failErr.Msg = "character class resulted in empty character set"
		return "", failErr
	}

	return string(chars), nil
}
