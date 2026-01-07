//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffeid

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestKeeper_ValidTrustRoot tests Keeper SPIFFE ID construction
func TestKeeper_ValidTrustRoot(t *testing.T) {
	tests := []struct {
		name      string
		trustRoot string
		expected  string
	}{
		{
			name:      "SimpleDomain",
			trustRoot: "example.org",
			expected:  "spiffe://example.org/spike/keeper",
		},
		{
			name:      "SubDomain",
			trustRoot: "trust.example.org",
			expected:  "spiffe://trust.example.org/spike/keeper",
		},
		{
			name:      "LocalDomain",
			trustRoot: "local.domain",
			expected:  "spiffe://local.domain/spike/keeper",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Keeper(tt.trustRoot)
			assert.Equal(t, tt.expected, result)
			assert.True(t, strings.HasPrefix(result, "spiffe://"))
			assert.Contains(t, result, "/spike/keeper")
		})
	}
}

// TestNexus_ValidTrustRoot tests Nexus SPIFFE ID construction
func TestNexus_ValidTrustRoot(t *testing.T) {
	tests := []struct {
		name      string
		trustRoot string
		expected  string
	}{
		{
			name:      "SimpleDomain",
			trustRoot: "example.org",
			expected:  "spiffe://example.org/spike/nexus",
		},
		{
			name:      "SubDomain",
			trustRoot: "trust.example.org",
			expected:  "spiffe://trust.example.org/spike/nexus",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Nexus(tt.trustRoot)
			assert.Equal(t, tt.expected, result)
			assert.True(t, strings.HasPrefix(result, "spiffe://"))
			assert.Contains(t, result, "/spike/nexus")
		})
	}
}

// TestPilot_ValidTrustRoot tests Pilot SPIFFE ID construction
func TestPilot_ValidTrustRoot(t *testing.T) {
	tests := []struct {
		name      string
		trustRoot string
		expected  string
	}{
		{
			name:      "SimpleDomain",
			trustRoot: "example.org",
			expected:  "spiffe://example.org/spike/pilot/role/superuser",
		},
		{
			name:      "SubDomain",
			trustRoot: "trust.example.org",
			expected:  "spiffe://trust.example.org/spike/pilot/role/superuser",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Pilot(tt.trustRoot)
			assert.Equal(t, tt.expected, result)
			assert.True(t, strings.HasPrefix(result, "spiffe://"))
			assert.Contains(t, result, "/spike/pilot/role/superuser")
		})
	}
}

// TestBootstrap_ValidTrustRoot tests Bootstrap SPIFFE ID construction
func TestBootstrap_ValidTrustRoot(t *testing.T) {
	tests := []struct {
		name      string
		trustRoot string
		expected  string
	}{
		{
			name:      "SimpleDomain",
			trustRoot: "example.org",
			expected:  "spiffe://example.org/spike/bootstrap",
		},
		{
			name:      "SubDomain",
			trustRoot: "trust.example.org",
			expected:  "spiffe://trust.example.org/spike/bootstrap",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Bootstrap(tt.trustRoot)
			assert.Equal(t, tt.expected, result)
			assert.True(t, strings.HasPrefix(result, "spiffe://"))
			assert.Contains(t, result, "/spike/bootstrap")
		})
	}
}

// TestLiteWorkload_ValidTrustRoot tests LiteWorkload SPIFFE ID construction
func TestLiteWorkload_ValidTrustRoot(t *testing.T) {
	tests := []struct {
		name      string
		trustRoot string
		expected  string
	}{
		{
			name:      "SimpleDomain",
			trustRoot: "example.org",
			expected:  "spiffe://example.org/spike/workload/role/lite",
		},
		{
			name:      "SubDomain",
			trustRoot: "trust.example.org",
			expected:  "spiffe://trust.example.org/spike/workload/role/lite",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LiteWorkload(tt.trustRoot)
			assert.Equal(t, tt.expected, result)
			assert.True(t, strings.HasPrefix(result, "spiffe://"))
			assert.Contains(t, result, "/spike/workload/role/lite")
		})
	}
}

// TestPilotRecover_ValidTrustRoot tests PilotRecover SPIFFE ID construction
func TestPilotRecover_ValidTrustRoot(t *testing.T) {
	tests := []struct {
		name      string
		trustRoot string
		expected  string
	}{
		{
			name:      "SimpleDomain",
			trustRoot: "example.org",
			expected:  "spiffe://example.org/spike/pilot/role/recover",
		},
		{
			name:      "SubDomain",
			trustRoot: "trust.example.org",
			expected:  "spiffe://trust.example.org/spike/pilot/role/recover",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PilotRecover(tt.trustRoot)
			assert.Equal(t, tt.expected, result)
			assert.True(t, strings.HasPrefix(result, "spiffe://"))
			assert.Contains(t, result, "/spike/pilot/role/recover")
		})
	}
}

// TestPilotRestore_ValidTrustRoot tests PilotRestore SPIFFE ID construction
func TestPilotRestore_ValidTrustRoot(t *testing.T) {
	tests := []struct {
		name      string
		trustRoot string
		expected  string
	}{
		{
			name:      "SimpleDomain",
			trustRoot: "example.org",
			expected:  "spiffe://example.org/spike/pilot/role/restore",
		},
		{
			name:      "SubDomain",
			trustRoot: "trust.example.org",
			expected:  "spiffe://trust.example.org/spike/pilot/role/restore",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PilotRestore(tt.trustRoot)
			assert.Equal(t, tt.expected, result)
			assert.True(t, strings.HasPrefix(result, "spiffe://"))
			assert.Contains(t, result, "/spike/pilot/role/restore")
		})
	}
}

// TestIsPilotOperator_ExactMatch tests IsPilotOperator with an exact match
func TestIsPilotOperator_ExactMatch(t *testing.T) {
	// Set up environment variable
	os.Setenv("SPIKE_TRUST_ROOT_PILOT", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_PILOT")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "ExactMatch",
			spiffeID: "spiffe://example.org/spike/pilot/role/superuser",
			expected: true,
		},
		{
			name:     "ExtendedMatch",
			spiffeID: "spiffe://example.org/spike/pilot/role/superuser/instance-0",
			expected: true,
		},
		{
			name:     "NoMatch",
			spiffeID: "spiffe://other.org/spike/pilot/role/superuser",
			expected: false,
		},
		{
			name:     "PartialMatch",
			spiffeID: "spiffe://example.org/spike/keeper",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPilotOperator(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsLiteWorkload_ExactMatch tests IsLiteWorkload with an exact match
func TestIsLiteWorkload_ExactMatch(t *testing.T) {
	// Set up environment variable
	os.Setenv("SPIKE_TRUST_ROOT_LITE_WORKLOAD", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_LITE_WORKLOAD")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "ExactMatch",
			spiffeID: "spiffe://example.org/spike/workload/role/lite",
			expected: true,
		},
		{
			name:     "ExtendedMatch",
			spiffeID: "spiffe://example.org/spike/workload/role/lite/instance-0",
			expected: true,
		},
		{
			name:     "NoMatch",
			spiffeID: "spiffe://other.org/spike/workload/role/lite",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLiteWorkload(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsPilotRecover_ExactMatch tests IsPilotRecover with an exact match
func TestIsPilotRecover_ExactMatch(t *testing.T) {
	// Set up environment variable
	os.Setenv("SPIKE_TRUST_ROOT_PILOT", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_PILOT")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "ExactMatch",
			spiffeID: "spiffe://example.org/spike/pilot/role/recover",
			expected: true,
		},
		{
			name:     "ExtendedMatch",
			spiffeID: "spiffe://example.org/spike/pilot/role/recover/instance-0",
			expected: true,
		},
		{
			name:     "NoMatch",
			spiffeID: "spiffe://other.org/spike/pilot/role/recover",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPilotRecover(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsPilotRestore_ExactMatch tests IsPilotRestore with an exact match
func TestIsPilotRestore_ExactMatch(t *testing.T) {
	// Set up environment variable
	os.Setenv("SPIKE_TRUST_ROOT_PILOT", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_PILOT")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "ExactMatch",
			spiffeID: "spiffe://example.org/spike/pilot/role/restore",
			expected: true,
		},
		{
			name:     "ExtendedMatch",
			spiffeID: "spiffe://example.org/spike/pilot/role/restore/instance-0",
			expected: true,
		},
		{
			name:     "NoMatch",
			spiffeID: "spiffe://other.org/spike/pilot/role/restore",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPilotRestore(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsBootstrap_ExactMatch tests IsBootstrap with an exact match
func TestIsBootstrap_ExactMatch(t *testing.T) {
	// Set up environment variable
	os.Setenv("SPIKE_TRUST_ROOT_BOOTSTRAP", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_BOOTSTRAP")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "ExactMatch",
			spiffeID: "spiffe://example.org/spike/bootstrap",
			expected: true,
		},
		{
			name:     "ExtendedMatch",
			spiffeID: "spiffe://example.org/spike/bootstrap/instance-0",
			expected: true,
		},
		{
			name:     "NoMatch",
			spiffeID: "spiffe://other.org/spike/bootstrap",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBootstrap(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsKeeper_ExactMatch tests IsKeeper with an exact match
func TestIsKeeper_ExactMatch(t *testing.T) {
	// Set up environment variable
	os.Setenv("SPIKE_TRUST_ROOT_KEEPER", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_KEEPER")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "ExactMatch",
			spiffeID: "spiffe://example.org/spike/keeper",
			expected: true,
		},
		{
			name:     "ExtendedMatch",
			spiffeID: "spiffe://example.org/spike/keeper/instance-0",
			expected: true,
		},
		{
			name:     "NoMatch",
			spiffeID: "spiffe://other.org/spike/keeper",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsKeeper(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsNexus_ExactMatch tests IsNexus with an exact match
func TestIsNexus_ExactMatch(t *testing.T) {
	// Set up environment variable
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "ExactMatch",
			spiffeID: "spiffe://example.org/spike/nexus",
			expected: true,
		},
		{
			name:     "ExtendedMatch",
			spiffeID: "spiffe://example.org/spike/nexus/instance-0",
			expected: true,
		},
		{
			name:     "NoMatch",
			spiffeID: "spiffe://other.org/spike/nexus",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNexus(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestPeerCanTalkToAnyone tests the debug function
func TestPeerCanTalkToAnyone(t *testing.T) {
	tests := []struct {
		name  string
		peer1 string
		peer2 string
	}{
		{"BothEmpty", "", ""},
		{"OneEmpty", "spiffe://example.org/spike/nexus", ""},
		{"BothValid", "spiffe://example.org/spike/nexus", "spiffe://example.org/spike/keeper"},
		{"Invalid", "not-a-spiffe-id", "also-not-a-spiffe-id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PeerCanTalkToAnyone(tt.peer1, tt.peer2)
			assert.True(t, result, "PeerCanTalkToAnyone should always return true")
		})
	}
}

// TestPeerCanTalkToKeeper tests keeper authorization
func TestPeerCanTalkToKeeper(t *testing.T) {
	// Set up environment variables
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org")
	os.Setenv("SPIKE_TRUST_ROOT_BOOTSTRAP", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_BOOTSTRAP")

	tests := []struct {
		name     string
		peerID   string
		expected bool
	}{
		{
			name:     "NexusCanTalk",
			peerID:   "spiffe://example.org/spike/nexus",
			expected: true,
		},
		{
			name:     "BootstrapCanTalk",
			peerID:   "spiffe://example.org/spike/bootstrap",
			expected: true,
		},
		{
			name:     "NexusExtendedCanTalk",
			peerID:   "spiffe://example.org/spike/nexus/instance-0",
			expected: true,
		},
		{
			name:     "BootstrapExtendedCanTalk",
			peerID:   "spiffe://example.org/spike/bootstrap/instance-0",
			expected: true,
		},
		{
			name:     "PilotCannotTalk",
			peerID:   "spiffe://example.org/spike/pilot/role/superuser",
			expected: false,
		},
		{
			name:     "KeeperCannotTalk",
			peerID:   "spiffe://example.org/spike/keeper",
			expected: false,
		},
		{
			name:     "InvalidCannotTalk",
			peerID:   "spiffe://other.org/spike/nexus",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PeerCanTalkToKeeper(tt.peerID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestMultipleTrustDomains tests validation with multiple trust domains
func TestMultipleTrustDomains(t *testing.T) {
	// Set up an environment variable with multiple trust domains
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org, other.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{
			name:     "FirstDomain",
			spiffeID: "spiffe://example.org/spike/nexus",
			expected: true,
		},
		{
			name:     "SecondDomain",
			spiffeID: "spiffe://other.org/spike/nexus",
			expected: true,
		},
		{
			name:     "ThirdDomainNotConfigured",
			spiffeID: "spiffe://third.org/spike/nexus",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNexus(tt.spiffeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestSPIFFEIDFormat tests the format of generated SPIFFE IDs
func TestSPIFFEIDFormat(t *testing.T) {
	trustRoot := "example.org"

	ids := map[string]string{
		"Keeper":       Keeper(trustRoot),
		"Nexus":        Nexus(trustRoot),
		"Pilot":        Pilot(trustRoot),
		"Bootstrap":    Bootstrap(trustRoot),
		"LiteWorkload": LiteWorkload(trustRoot),
		"PilotRecover": PilotRecover(trustRoot),
		"PilotRestore": PilotRestore(trustRoot),
	}

	for name, id := range ids {
		t.Run(name, func(t *testing.T) {
			// All IDs should start with spiffe://
			assert.True(t, strings.HasPrefix(id, "spiffe://"),
				"%s ID should start with spiffe://", name)

			// All IDs should contain the trust root
			assert.Contains(t, id, trustRoot,
				"%s ID should contain trust root", name)

			// All IDs should contain /spike/
			assert.Contains(t, id, "/spike/",
				"%s ID should contain /spike/", name)

			// ID should not contain double slashes except in spiffe://
			cleaned := strings.Replace(id, "spiffe://", "", 1)
			assert.False(t, strings.Contains(cleaned, "//"),
				"%s ID should not contain double slashes", name)
		})
	}
}
