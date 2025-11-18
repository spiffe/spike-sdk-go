//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

const letters = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a cryptographically secure random string of the
// specified length using alphanumeric characters (a-z, A-Z, 0-9).
//
// Parameters:
//   - n: length of the random string to generate
//
// Returns:
//   - string: the generated random alphanumeric string
//   - error: non-nil if the cryptographic random number generator fails
func RandomString(n int) (string, error) {
	bytes := make([]byte, n)

	if _, err := reader(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}

// Token generates a cryptographically secure random token with the "spike."
// prefix. The token consists of the prefix followed by 26 random alphanumeric
// characters, resulting in a format like "spike.AbCd1234EfGh5678IjKl9012Mn".
//
// Returns:
//   - string: the generated token in the format "spike.<26-char-random-string>"
//   - error: non-nil if the cryptographic random number generator fails
func Token() (string, error) {
	id, err := RandomString(26)
	if err != nil {
		return "", err
	}
	return "spike." + id, nil
}

// ID generates a cryptographically secure random identifier consisting of
// 8 alphanumeric characters. Suitable for use as short, unique identifiers.
//
// Returns:
//   - string: the generated 8-character random alphanumeric string
//   - error: non-nil if the cryptographic random number generator fails
func ID() (string, error) {
	return RandomString(8)
}
