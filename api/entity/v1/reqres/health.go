package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

type HealthReadRequest struct {
	Version int `json:"version,omitempty"` // Optional specific version
}

type HealthReadResponse struct {
	StatusResponse data.StatusResponse `json:"status_response"`
	Err            data.ErrorCode      `json:"err,omitempty"`
}
