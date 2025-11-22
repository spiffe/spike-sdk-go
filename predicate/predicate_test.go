//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package predicate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAllowAll_ReturnsTrue tests that AllowAll accepts any SPIFFE ID
func TestAllowAll_ReturnsTrue(t *testing.T) {
	tests := []struct {
		name     string
		spiffeID string
	}{
		{"ValidSPIFFEID", "spiffe://example.org/service"},
		{"Empty", ""},
		{"NexusSPIFFEID", "spiffe://example.org/spike/nexus"},
		{"KeeperSPIFFEID", "spiffe://example.org/spike/keeper"},
		{"RandomString", "not-a-spiffe-id"},
		{"WithSpecialChars", "spiffe://example.org/service@#$"},
		{"VeryLong", "spiffe://example.org/" + string(make([]byte, 1000))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowAll(tt.spiffeID)
			assert.True(t, result, "AllowAll should return true for: %s", tt.spiffeID)
		})
	}
}

// TestDenyAll_ReturnsFalse tests that DenyAll rejects all SPIFFE IDs
func TestDenyAll_ReturnsFalse(t *testing.T) {
	tests := []struct {
		name     string
		spiffeID string
	}{
		{"ValidSPIFFEID", "spiffe://example.org/service"},
		{"Empty", ""},
		{"NexusSPIFFEID", "spiffe://example.org/spike/nexus"},
		{"KeeperSPIFFEID", "spiffe://example.org/spike/keeper"},
		{"RandomString", "not-a-spiffe-id"},
		{"WithSpecialChars", "spiffe://example.org/service@#$"},
		{"VeryLong", "spiffe://example.org/" + string(make([]byte, 1000))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DenyAll(tt.spiffeID)
			assert.False(t, result, "DenyAll should return false for: %s", tt.spiffeID)
		})
	}
}

// TestAllowNexus_ValidNexusID tests that AllowNexus accepts valid Nexus SPIFFE IDs
func TestAllowNexus_ValidNexusID(t *testing.T) {
	// Set up environment variable for Nexus trust root
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")

	tests := []struct {
		name     string
		spiffeID string
	}{
		{"ExactMatch", "spiffe://example.org/spike/nexus"},
		{"WithInstance", "spiffe://example.org/spike/nexus/instance-1"},
		{"WithPath", "spiffe://example.org/spike/nexus/path/to/service"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowNexus(tt.spiffeID)
			assert.True(t, result, "AllowNexus should return true for: %s", tt.spiffeID)
		})
	}
}

// TestAllowNexus_InvalidNexusID tests that AllowNexus rejects non-Nexus SPIFFE IDs
func TestAllowNexus_InvalidNexusID(t *testing.T) {
	// Set up environment variable for Nexus trust root
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")

	tests := []struct {
		name     string
		spiffeID string
	}{
		{"Empty", ""},
		{"KeeperSPIFFEID", "spiffe://example.org/spike/keeper"},
		{"PilotSPIFFEID", "spiffe://example.org/spike/pilot"},
		{"RegularService", "spiffe://example.org/service"},
		{"WrongDomain", "spiffe://other.org/spike/nexus"},
		{"PartialMatch", "spiffe://example.org/spike/nexu"},
		{"NotSPIFFEID", "not-a-spiffe-id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowNexus(tt.spiffeID)
			assert.False(t, result, "AllowNexus should return false for: %s", tt.spiffeID)
		})
	}
}

// TestAllowKeeper_ValidKeeperID tests that AllowKeeper accepts valid Keeper SPIFFE IDs
func TestAllowKeeper_ValidKeeperID(t *testing.T) {
	// Set up environment variable for Keeper trust root
	os.Setenv("SPIKE_TRUST_ROOT_KEEPER", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_KEEPER")

	tests := []struct {
		name     string
		spiffeID string
	}{
		{"ExactMatch", "spiffe://example.org/spike/keeper"},
		{"WithInstance", "spiffe://example.org/spike/keeper/instance-1"},
		{"WithPath", "spiffe://example.org/spike/keeper/path/to/service"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowKeeper(tt.spiffeID)
			assert.True(t, result, "AllowKeeper should return true for: %s", tt.spiffeID)
		})
	}
}

// TestAllowKeeper_InvalidKeeperID tests that AllowKeeper rejects non-Keeper SPIFFE IDs
func TestAllowKeeper_InvalidKeeperID(t *testing.T) {
	// Set up environment variable for Keeper trust root
	os.Setenv("SPIKE_TRUST_ROOT_KEEPER", "example.org")
	defer os.Unsetenv("SPIKE_TRUST_ROOT_KEEPER")

	tests := []struct {
		name     string
		spiffeID string
	}{
		{"Empty", ""},
		{"NexusSPIFFEID", "spiffe://example.org/spike/nexus"},
		{"PilotSPIFFEID", "spiffe://example.org/spike/pilot"},
		{"RegularService", "spiffe://example.org/service"},
		{"WrongDomain", "spiffe://other.org/spike/keeper"},
		{"PartialMatch", "spiffe://example.org/spike/keepe"},
		{"NotSPIFFEID", "not-a-spiffe-id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowKeeper(tt.spiffeID)
			assert.False(t, result, "AllowKeeper should return false for: %s", tt.spiffeID)
		})
	}
}

// TestAllowKeeperPeer_ValidPeers tests that AllowKeeperPeer accepts Nexus and Bootstrap peers
func TestAllowKeeperPeer_ValidPeers(t *testing.T) {
	// Set up environment variables for trust roots
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org")
	os.Setenv("SPIKE_TRUST_ROOT_BOOTSTRAP", "example.org")
	defer func() {
		os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")
		os.Unsetenv("SPIKE_TRUST_ROOT_BOOTSTRAP")
	}()

	tests := []struct {
		name     string
		spiffeID string
	}{
		{"NexusExact", "spiffe://example.org/spike/nexus"},
		{"NexusWithInstance", "spiffe://example.org/spike/nexus/instance-1"},
		{"BootstrapExact", "spiffe://example.org/spike/bootstrap"},
		{"BootstrapWithInstance", "spiffe://example.org/spike/bootstrap/instance-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowKeeperPeer(tt.spiffeID)
			assert.True(t, result, "AllowKeeperPeer should return true for: %s", tt.spiffeID)
		})
	}
}

// TestAllowKeeperPeer_InvalidPeers tests that AllowKeeperPeer rejects unauthorized peers
func TestAllowKeeperPeer_InvalidPeers(t *testing.T) {
	// Set up environment variables for trust roots
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org")
	os.Setenv("SPIKE_TRUST_ROOT_BOOTSTRAP", "example.org")
	defer func() {
		os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")
		os.Unsetenv("SPIKE_TRUST_ROOT_BOOTSTRAP")
	}()

	tests := []struct {
		name     string
		spiffeID string
	}{
		{"Empty", ""},
		{"KeeperSPIFFEID", "spiffe://example.org/spike/keeper"},
		{"PilotSPIFFEID", "spiffe://example.org/spike/pilot"},
		{"RegularService", "spiffe://example.org/service"},
		{"WrongDomain", "spiffe://other.org/spike/nexus"},
		{"NotSPIFFEID", "not-a-spiffe-id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowKeeperPeer(tt.spiffeID)
			assert.False(t, result, "AllowKeeperPeer should return false for: %s", tt.spiffeID)
		})
	}
}

// TestPredicate_TypeFunction tests that Predicate type works as expected
func TestPredicate_TypeFunction(t *testing.T) {
	// Create a custom predicate
	customPredicate := Predicate(func(spiffeID string) bool {
		return spiffeID == "spiffe://example.org/custom"
	})

	// Test that it works correctly
	assert.True(t, customPredicate("spiffe://example.org/custom"))
	assert.False(t, customPredicate("spiffe://example.org/other"))
	assert.False(t, customPredicate(""))
}

// TestPredicate_Composition tests that predicates can be composed
func TestPredicate_Composition(t *testing.T) {
	// Create a composite predicate using OR logic
	allowNexusOrKeeper := func(spiffeID string) bool {
		return AllowNexus(spiffeID) || AllowKeeper(spiffeID)
	}

	// Set up environment variables
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org")
	os.Setenv("SPIKE_TRUST_ROOT_KEEPER", "example.org")
	defer func() {
		os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")
		os.Unsetenv("SPIKE_TRUST_ROOT_KEEPER")
	}()

	// Test composition
	assert.True(t, allowNexusOrKeeper("spiffe://example.org/spike/nexus"))
	assert.True(t, allowNexusOrKeeper("spiffe://example.org/spike/keeper"))
	assert.False(t, allowNexusOrKeeper("spiffe://example.org/spike/pilot"))
}

// TestAllowAll_IsPredicateType tests that AllowAll is of type Predicate
func TestAllowAll_IsPredicateType(_ *testing.T) {
	var _ Predicate = AllowAll
}

// TestDenyAll_IsPredicateType tests that DenyAll is of type Predicate
func TestDenyAll_IsPredicateType(_ *testing.T) {
	var _ Predicate = DenyAll
}

// TestAllowNexus_IsPredicateType tests that AllowNexus is of type Predicate
func TestAllowNexus_IsPredicateType(_ *testing.T) {
	var _ Predicate = AllowNexus
}

// TestAllowKeeper_IsPredicateType tests that AllowKeeper is of type Predicate
func TestAllowKeeper_IsPredicateType(_ *testing.T) {
	var _ Predicate = AllowKeeper
}

// TestAllowKeeperPeer_MultipleTrustDomains tests AllowKeeperPeer with multiple trust domains
func TestAllowKeeperPeer_MultipleTrustDomains(t *testing.T) {
	// Set up multiple trust domains
	os.Setenv("SPIKE_TRUST_ROOT_NEXUS", "example.org,dev.example.org")
	os.Setenv("SPIKE_TRUST_ROOT_BOOTSTRAP", "example.org,dev.example.org")
	defer func() {
		os.Unsetenv("SPIKE_TRUST_ROOT_NEXUS")
		os.Unsetenv("SPIKE_TRUST_ROOT_BOOTSTRAP")
	}()

	tests := []struct {
		name     string
		spiffeID string
		expected bool
	}{
		{"NexusMainDomain", "spiffe://example.org/spike/nexus", true},
		{"NexusDevDomain", "spiffe://dev.example.org/spike/nexus", true},
		{"BootstrapMainDomain", "spiffe://example.org/spike/bootstrap", true},
		{"BootstrapDevDomain", "spiffe://dev.example.org/spike/bootstrap", true},
		{"UntrustedDomain", "spiffe://untrusted.org/spike/nexus", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AllowKeeperPeer(tt.spiffeID)
			assert.Equal(t, tt.expected, result, "AllowKeeperPeer(%s) should return %v", tt.spiffeID, tt.expected)
		})
	}
}
