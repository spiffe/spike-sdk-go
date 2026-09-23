//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package fs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spiffe/spike-sdk-go/config/env"
)

// TestTempDirUser tests that USER is reduced to a single path element, so the
// /tmp fallback directories cannot leave /tmp.
func TestTempDirUser(t *testing.T) {
	tests := []struct {
		name     string
		user     string
		expected string
	}{
		{"PlainUser", "alice", "alice"},
		{"Empty", "", "spike"},
		{"Dot", ".", "spike"},
		{"DotDot", "..", "spike"},
		{"SeparatorOnly", "/", "spike"},
		{"Traversal", "/../../etc", "etc"},
		{"TrailingDotDot", "alice/..", "spike"},
		{"NestedPath", "a/b", "b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("USER", tt.user)
			user := tempDirUser()
			assert.Equal(t, tt.expected, user)

			tempDir := filepath.Join("/tmp", ".spike-"+user)
			assert.Equal(t, "/tmp", filepath.Dir(tempDir),
				"fallback directory must stay directly under /tmp")
		})
	}
}

// TestTryCustomNexusDataDir_CreatesValidatedPath tests that a relative
// SPIKE_NEXUS_DATA_DIR is created at the absolute path that was validated.
func TestTryCustomNexusDataDir_CreatesValidatedPath(t *testing.T) {
	base := t.TempDir()
	t.Chdir(base)
	t.Setenv(env.NexusDataDir, "custom")

	path := tryCustomNexusDataDir("test")

	expected := filepath.Join(base, "custom", spikeDataFolderName)
	assert.Equal(t, expected, path)
	info, statErr := os.Stat(expected)
	require.NoError(t, statErr)
	assert.True(t, info.IsDir())
}

// TestTryCustomPilotRecoveryDir_CreatesValidatedPath tests that a relative
// SPIKE_PILOT_RECOVERY_DIR is created at the absolute path that was validated.
func TestTryCustomPilotRecoveryDir_CreatesValidatedPath(t *testing.T) {
	base := t.TempDir()
	t.Chdir(base)
	t.Setenv(env.PilotRecoveryDir, "custom")

	path := tryCustomPilotRecoveryDir("test")

	expected := filepath.Join(base, "custom", spikeRecoveryFolderName)
	assert.Equal(t, expected, path)
	info, statErr := os.Stat(expected)
	require.NoError(t, statErr)
	assert.True(t, info.IsDir())
}

// TestTryCustomNexusDataDir_RejectsRestrictedPath tests that a restricted
// directory is refused and nothing is created.
func TestTryCustomNexusDataDir_RejectsRestrictedPath(t *testing.T) {
	t.Setenv(env.NexusDataDir, "/")

	assert.Empty(t, tryCustomNexusDataDir("test"))
}
