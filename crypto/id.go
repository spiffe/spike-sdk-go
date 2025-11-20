//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/spiffe/spike-sdk-go/log"
)

const letters = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a cryptographically secure random string of the
// specified length using alphanumeric characters (a-z, A-Z, 0-9).
//
// Security Note: This function will fatally crash the process
// (via log.FatalErr) if the system's cryptographic random number generator
// fails. This is an intentional security decision, not a bug. Here's why:
//
//  1. CSPRNG failure indicates a critical system-level security compromise
//  2. Continuing with potentially weak/predictable random values would create
//     security vulnerabilities (weak tokens, predictable IDs, compromised
//     secrets)
//  3. There is no safe fallback - using non-cryptographic randomness or
//     deterministic values would be catastrophically insecure
//  4. This is consistent with other security-critical failures in the codebase
//     (SVID acquisition, Shamir operations, Pilot restrictions)
//  5. Failing loudly prevents silent security degradation and forces operators
//     to address the underlying system issue
//
// The crypto/rand documentation states that Read failures are extremely rare
// and indicate serious OS-level problems. When this happens, the entire
// system's security is compromised, not just this function.
//
// Parameters:
//   - n: length of the random string to generate
//
// Returns:
//   - string: the generated random alphanumeric string
func RandomString(n int) string {
	const fName = "RandomString"

	bytes := make([]byte, n)

	if _, err := reader(bytes); err != nil {
		failErr := sdkErrors.ErrCryptoRandomGenerationFailed.Wrap(err)
		failErr.Msg = "cryptographic random number generator failed"
		log.FatalErr(fName, *failErr)
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes)
}

// Token generates a cryptographically secure random token with the "spike."
// prefix. The token consists of the prefix followed by 26 random alphanumeric
// characters, resulting in a format like "spike.AbCd1234EfGh5678IjKl9012Mn".
//
// Security Note: This function will fatally crash the process if the
// cryptographic random number generator fails. See RandomString() documentation
// for the security rationale behind this behavior.
//
// Returns:
//   - string: the generated token in the format "spike.<26-char-random-string>"
func Token() string {
	id := RandomString(26)
	return "spike." + id
}

// ID generates a cryptographically secure random identifier consisting of
// 8 alphanumeric characters. Suitable for use as short, unique identifiers.
//
// Security Note: This function will fatally crash the process if the
// cryptographic random number generator fails. See RandomString() documentation
// for the security rationale behind this behavior.
//
// Returns:
//   - string: the generated 8-character random alphanumeric string
func ID() string {
	return RandomString(8)
}
