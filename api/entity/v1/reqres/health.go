package reqres

import (
	"time"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
)

type HealthReadRequest struct {
	Version int `json:"version,omitempty"` // Optional specific version
}

type HealthReadResponse struct {
	StatusResponse StatusResponse `json:"status_response"`
	Err            data.ErrorCode `json:"err,omitempty"`
}

type NexusStatusRequest struct{}

type StatusResponse struct {
	Health    string    `json:"health"`
	Timestamp time.Time `json:"timestamp"`
	// Keepers       KeeperStatus  `json:"keepers"` --- IGNORE ---
	RootKey       data.RootKeyState      `json:"root_key"`
	BackingStore  data.BackingStoreState `json:"backing_store"`
	FIPSMode      bool                   `json:"fips_mode"`
	SecretsCount  *int                   `json:"secrets_count,omitempty"`
	UptimeSeconds int64                  `json:"uptime_seconds"`
}
