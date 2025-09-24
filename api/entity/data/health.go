package data

import "time"

type StatusResponse struct {
	Health    string    `json:"health"`
	Timestamp time.Time `json:"timestamp"`
	// Keepers       KeeperStatus  `json:"keepers"` --- IGNORE ---
	RootKey       RootKeyStatus `json:"root_key"`
	BackingStore  BackingStore  `json:"backing_store"`
	FIPSMode      bool          `json:"fips_mode"`
	SecretsCount  *int          `json:"secrets_count,omitempty"`
	UptimeSeconds int64         `json:"uptime_seconds"`
}

// KeeperStatus represents the status of the keeper cluster used for
// Shamir secret sharing. It tracks both the configured keeper instances
// and the threshold required for root key reconstruction.
type KeeperStatus struct {
	Status      string `json:"status"`
	ActiveCount int    `json:"actsve_count"`
}

// RootKeyStatus indicates whether the root encryption key is available
// for cryptographic operations and where it's sourced from.
type RootKeyStatus struct {
	Status string `json:"status"`
}

// BackingStore represents the connection status and performance metrics
// of the persistent storage backend used for secret storage.
type BackingStore struct {
	Status         string `json:"status"`
	Type           string `json:"type"`
	ResponseTimeMs *int   `json:"response_time_ms,omitempty"`
}

type Result struct {
	RootKey      RootKeyStatus
	BackingStore BackingStore
	FipsMode     bool
	SecretsCount *int
	Health       string
}
