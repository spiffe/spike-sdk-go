package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

// ShardContributionRequest represents a request to submit a shard contribution.
// KeeperId specifies the identifier of the keeper responsible for the shard.
// Shard represents the shard data being contributed to the system.
// Version optionally specifies the version of the shard being submitted.
type ShardContributionRequest struct {
	KeeperId string `json:"id"`
	Shard    string `json:"shard"`
}

// ShardContributionResponse represents the response structure for a shard
// contribution.
type ShardContributionResponse struct {
	Err data.ErrorCode `json:"err,omitempty"`
}

// ShardRequest represents a request to handle data partitioning or sharding.
type ShardRequest struct {
}

// ShardResponse represents the result of an operation on a specific data shard.
// The struct includes the shard identifier and an associated error code.
type ShardResponse struct {
	Shard string `json:"shard"`
	Err   data.ErrorCode
}
