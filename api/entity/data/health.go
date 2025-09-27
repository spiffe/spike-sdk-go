package data

// KeeperStatus represents the status of the keeper cluster used for
// Shamir secret sharing. It tracks both the configured keeper instances
// and the threshold required for root key reconstruction.
type KeeperState struct {
	Status      string `json:"status"`
	ActiveCount int    `json:"actsve_count"`
}

// RootKeyStatus indicates whether the root encryption key is available
// for cryptographic operations and where it's sourced from.
type RootKeyState struct {
	Status string `json:"status"`
}

// BackingStore represents the connection status and performance metrics
// of the persistent storage backend used for secret storage.
type BackingStoreState struct {
	Status         string `json:"status"`
	Type           string `json:"type"`
	ResponseTimeMs *int   `json:"response_time_ms,omitempty"`
}

type HealthResult struct {
	RootKey      RootKeyState
	BackingStore BackingStoreState
	FipsMode     bool
	SecretsCount *int
	Health       string
}
