package url

import (
	"net/url"
)

// ShardFromKeeperAPIRoot constructs the full URL for the keeper shard endpoint
// by joining the SPIKE Keeper API root with the shard path.
//
// This function is used during recovery operations to build the endpoint URL
// for retrieving Shamir secret shards from SPIKE Keeper instances.
//
// Parameters:
//   - keeperAPIRoot: The base URL of the keeper API
//     (e.g., "https://keeper.example.com:8443")
//
// Returns:
//   - string: The complete URL to the shard endpoint, or an error message
//     string if URL construction fails
//
// Example:
//
//	url := ShardFromKeeperAPIRoot("https://keeper.example.com:8443")
//	// Returns: "https://keeper.example.com:8443/v1/shard"
func ShardFromKeeperAPIRoot(keeperAPIRoot string) string {
	u, err := url.JoinPath(keeperAPIRoot, string(KeeperShard))
	if err != nil {
		return "parseError: Bad Keeper API Root"
	}
	return u
}
