package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

type ShardContributionRequest struct {
	KeeperId string `json:"id"`
	Shard    string `json:"shard"`
	Version  int    `json:"version,omitempty"` // Optional specific version
}

type ShardContributionResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}
