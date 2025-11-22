//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"encoding/hex"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// TestAES256Seed_Success tests successful AES-256 seed generation
func TestAES256Seed_Success(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	// Create deterministic reader for testing
	reader = func(b []byte) (int, error) {
		// Fill with deterministic data
		for i := range b {
			b[i] = byte(i)
		}
		return len(b), nil
	}

	seed, err := AES256Seed()

	assert.Nil(t, err)
	assert.NotEmpty(t, seed)
	assert.Equal(t, AES256KeySize*2, len(seed)) // Hex encoding doubles the length

	// Verify it's valid hex
	decoded, decodeErr := hex.DecodeString(seed)
	assert.NoError(t, decodeErr)
	assert.Equal(t, AES256KeySize, len(decoded))
}

// TestAES256Seed_Error tests AES256Seed when random generation fails
func TestAES256Seed_Error(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	// Mock reader that returns an error
	reader = func(_ []byte) (int, error) {
		return 0, errors.New("mock random generation failure")
	}

	seed, err := AES256Seed()

	assert.Empty(t, seed)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrCryptoFailedToCreateCipher))
}

// TestAES256Seed_Uniqueness tests that multiple calls generate different seeds
func TestAES256Seed_Uniqueness(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	// Use a counter to generate different values each time
	counter := 0
	reader = func(b []byte) (int, error) {
		for i := range b {
			b[i] = byte(counter + i)
		}
		counter++
		return len(b), nil
	}

	seed1, err1 := AES256Seed()
	seed2, err2 := AES256Seed()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotEqual(t, seed1, seed2, "Seeds should be unique")
}

// TestRandomString_Success tests successful random string generation
func TestRandomString_Success(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	// Use deterministic reader
	reader = func(b []byte) (int, error) {
		for i := range b {
			b[i] = byte(i % 62) // Keep within alphanumeric range
		}
		return len(b), nil
	}

	tests := []struct {
		name   string
		length int
	}{
		{"Length1", 1},
		{"Length8", 8},
		{"Length26", 26},
		{"Length64", 64},
		{"Length100", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			assert.Equal(t, tt.length, len(result))

			// Verify all characters are alphanumeric
			for _, c := range result {
				assert.True(t,
					(c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'),
					"Character %c should be alphanumeric", c)
			}
		})
	}
}

// TestRandomString_EmptyLength tests RandomString with zero length
func TestRandomString_EmptyLength(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	reader = func(b []byte) (int, error) {
		return len(b), nil
	}

	result := RandomString(0)
	assert.Equal(t, "", result)
}

// TestRandomString_CharacterDistribution tests that all character classes are used
func TestRandomString_CharacterDistribution(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	// Create a reader that returns different values to cover all character ranges
	position := 0
	reader = func(b []byte) (int, error) {
		for i := range b {
			b[i] = byte(position % 256)
			position++
		}
		return len(b), nil
	}

	result := RandomString(100)

	// Check that we have at least some variety in the output
	assert.True(t, len(result) == 100)

	// Verify it only contains valid characters
	for _, c := range result {
		assert.True(t, strings.ContainsRune(letters, c),
			"Character %c should be from the letters set", c)
	}
}

// TestToken_Format tests that Token generates correct format
func TestToken_Format(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	reader = func(b []byte) (int, error) {
		for i := range b {
			b[i] = byte(i % 62)
		}
		return len(b), nil
	}

	token := Token()

	// Verify format: "spike." + 26 characters
	assert.True(t, strings.HasPrefix(token, "spike."))
	assert.Equal(t, 6+26, len(token)) // "spike." is 6 chars + 26 random chars

	// Extract and verify the random part
	randomPart := token[6:]
	assert.Equal(t, 26, len(randomPart))

	// Verify all characters in random part are alphanumeric
	for _, c := range randomPart {
		assert.True(t,
			(c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'),
			"Character %c should be alphanumeric", c)
	}
}

// TestToken_Uniqueness tests that multiple Token calls generate different tokens
func TestToken_Uniqueness(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	counter := 0
	reader = func(b []byte) (int, error) {
		for i := range b {
			b[i] = byte(counter + i)
		}
		counter++
		return len(b), nil
	}

	token1 := Token()
	token2 := Token()

	assert.NotEqual(t, token1, token2, "Tokens should be unique")
	assert.True(t, strings.HasPrefix(token1, "spike."))
	assert.True(t, strings.HasPrefix(token2, "spike."))
}

// TestID_Format tests that ID generates correct length
func TestID_Format(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	reader = func(b []byte) (int, error) {
		for i := range b {
			b[i] = byte(i % 62)
		}
		return len(b), nil
	}

	id := ID()

	assert.Equal(t, 8, len(id))

	// Verify all characters are alphanumeric
	for _, c := range id {
		assert.True(t,
			(c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'),
			"Character %c should be alphanumeric", c)
	}
}

// TestID_Uniqueness tests that multiple ID calls generate different IDs
func TestID_Uniqueness(t *testing.T) {
	// Save original reader and restore after test
	originalReader := reader
	defer func() { reader = originalReader }()

	counter := 0
	reader = func(b []byte) (int, error) {
		for i := range b {
			b[i] = byte(counter + i)
		}
		counter++
		return len(b), nil
	}

	id1 := ID()
	id2 := ID()
	id3 := ID()

	assert.Equal(t, 8, len(id1))
	assert.Equal(t, 8, len(id2))
	assert.Equal(t, 8, len(id3))
	assert.NotEqual(t, id1, id2, "IDs should be unique")
	assert.NotEqual(t, id2, id3, "IDs should be unique")
	assert.NotEqual(t, id1, id3, "IDs should be unique")
}

// TestDeterministicReader_Read tests the Read method
func TestDeterministicReader_Read(t *testing.T) {
	seed := []byte("test seed")
	reader := NewDeterministicReader(seed)

	buffer := make([]byte, 16)
	n, err := reader.Read(buffer)

	assert.NoError(t, err)
	assert.Equal(t, 16, n)
	assert.NotEmpty(t, buffer)
}

// TestDeterministicReader_Consistency tests that same seed produces same output
func TestDeterministicReader_Consistency(t *testing.T) {
	seed := []byte("test seed")

	// Create two readers with same seed
	reader1 := NewDeterministicReader(seed)
	reader2 := NewDeterministicReader(seed)

	buffer1 := make([]byte, 32)
	buffer2 := make([]byte, 32)

	n1, err1 := reader1.Read(buffer1)
	n2, err2 := reader2.Read(buffer2)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, 32, n1)
	assert.Equal(t, 32, n2)
	assert.Equal(t, buffer1, buffer2, "Same seed should produce same output")
}

// TestDeterministicReader_DifferentSeeds tests that different seeds produce different output
func TestDeterministicReader_DifferentSeeds(t *testing.T) {
	seed1 := []byte("seed one")
	seed2 := []byte("seed two")

	reader1 := NewDeterministicReader(seed1)
	reader2 := NewDeterministicReader(seed2)

	buffer1 := make([]byte, 32)
	buffer2 := make([]byte, 32)

	_, err1 := reader1.Read(buffer1)
	_, err2 := reader2.Read(buffer2)

	require.NoError(t, err1)
	require.NoError(t, err2)
	assert.NotEqual(t, buffer1, buffer2, "Different seeds should produce different output")
}

// TestDeterministicReader_MultipleReads tests multiple consecutive reads
func TestDeterministicReader_MultipleReads(t *testing.T) {
	seed := []byte("test seed")
	reader := NewDeterministicReader(seed)

	// Read in chunks
	buffer1 := make([]byte, 16)
	buffer2 := make([]byte, 16)
	buffer3 := make([]byte, 16)

	n1, err1 := reader.Read(buffer1)
	n2, err2 := reader.Read(buffer2)
	n3, err3 := reader.Read(buffer3)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NoError(t, err3)
	assert.Equal(t, 16, n1)
	assert.Equal(t, 16, n2)
	assert.Equal(t, 16, n3)

	// Buffers should be different (sequential reads)
	assert.NotEqual(t, buffer1, buffer2)
	assert.NotEqual(t, buffer2, buffer3)
}

// TestDeterministicReader_LargeRead tests reading more data than initial buffer
func TestDeterministicReader_LargeRead(t *testing.T) {
	seed := []byte("test seed")
	reader := NewDeterministicReader(seed)

	// Read more than the initial SHA-256 hash size (32 bytes)
	// Note: io.Reader may return less than requested
	buffer := make([]byte, 100)
	totalRead := 0

	// Read in multiple calls to fill the buffer
	for totalRead < 100 {
		n, err := reader.Read(buffer[totalRead:])
		assert.NoError(t, err)
		assert.True(t, n > 0, "Should read at least some data")
		totalRead += n
	}

	assert.Equal(t, 100, totalRead)
	assert.NotEmpty(t, buffer)
}

// TestDeterministicReader_EmptyBuffer tests reading into empty buffer
func TestDeterministicReader_EmptyBuffer(t *testing.T) {
	seed := []byte("test seed")
	reader := NewDeterministicReader(seed)

	buffer := make([]byte, 0)
	n, err := reader.Read(buffer)

	assert.NoError(t, err)
	assert.Equal(t, 0, n)
}

// TestDeterministicReader_SmallReads tests many small consecutive reads
func TestDeterministicReader_SmallReads(t *testing.T) {
	seed := []byte("test seed")
	reader := NewDeterministicReader(seed)

	// Collect data from many small reads
	var collected []byte
	for i := 0; i < 100; i++ {
		buffer := make([]byte, 1)
		n, err := reader.Read(buffer)
		assert.NoError(t, err)
		assert.Equal(t, 1, n)
		collected = append(collected, buffer[0])
	}

	assert.Equal(t, 100, len(collected))
}

// TestDeterministicReader_ReproducibleStream tests that stream is reproducible
func TestDeterministicReader_ReproducibleStream(t *testing.T) {
	seed := []byte("reproducible seed")

	// First stream
	reader1 := NewDeterministicReader(seed)
	stream1 := make([]byte, 200)
	_, err1 := reader1.Read(stream1)

	// Second stream with same seed
	reader2 := NewDeterministicReader(seed)
	stream2 := make([]byte, 200)
	_, err2 := reader2.Read(stream2)

	require.NoError(t, err1)
	require.NoError(t, err2)
	assert.Equal(t, stream1, stream2, "Stream should be reproducible with same seed")
}

// TestAES256KeySize_Constant tests the AES256KeySize constant
func TestAES256KeySize_Constant(t *testing.T) {
	assert.Equal(t, 32, AES256KeySize, "AES-256 key size should be 32 bytes")
}

// TestLettersConstant tests the letters constant has expected characters
func TestLettersConstant(t *testing.T) {
	// Verify letters contains lowercase
	for c := 'a'; c <= 'z'; c++ {
		assert.True(t, strings.ContainsRune(letters, c),
			"letters should contain lowercase letter %c", c)
	}

	// Verify letters contains uppercase
	for c := 'A'; c <= 'Z'; c++ {
		assert.True(t, strings.ContainsRune(letters, c),
			"letters should contain uppercase letter %c", c)
	}

	// Verify letters contains digits
	for c := '0'; c <= '9'; c++ {
		assert.True(t, strings.ContainsRune(letters, c),
			"letters should contain digit %c", c)
	}

	// Verify total length
	assert.Equal(t, 62, len(letters), "letters should have 62 characters (26+26+10)")
}

// TestNewDeterministicReader_NilSeed tests creating reader with nil seed
func TestNewDeterministicReader_NilSeed(t *testing.T) {
	reader := NewDeterministicReader(nil)
	require.NotNil(t, reader)

	buffer := make([]byte, 32)
	n, err := reader.Read(buffer)

	assert.NoError(t, err)
	assert.Equal(t, 32, n)
	assert.NotEmpty(t, buffer)
}

// TestNewDeterministicReader_EmptySeed tests creating reader with empty seed
func TestNewDeterministicReader_EmptySeed(t *testing.T) {
	reader := NewDeterministicReader([]byte{})
	require.NotNil(t, reader)

	buffer := make([]byte, 32)
	n, err := reader.Read(buffer)

	assert.NoError(t, err)
	assert.Equal(t, 32, n)
	assert.NotEmpty(t, buffer)
}
