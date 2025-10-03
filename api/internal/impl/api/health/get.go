package health

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/url"
	"github.com/spiffe/spike-sdk-go/log"
	"github.com/spiffe/spike-sdk-go/net"
)

// GetSystemStatus fetches the current system status from the operator endpoint using POST
func GetSystemStatus(ctx context.Context, source *workloadapi.X509Source) (*reqres.StatusResponse, error) {
	client, err := net.CreateMTLSClient(source)
	if err != nil {
		log.Log().Error("[ERROR] Creating mTLS client:", err)
		return nil, err
	}

	payload, _ := json.Marshal(reqres.HealthReadRequest{})

	body, err := net.Post(client, url.GetHealth(), payload)
	if err != nil {
		log.Log().Error("POST request failed", "error", err)
		return nil, err
	}

	var res reqres.HealthReadResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Log().Error("[ERROR] Unmarshalling response failed:", err)
		return nil, errors.Join(errors.New("GetSystemStatus: cannot unmarshal response"), err)
	}
	if res.Err != "" {
		return nil, errors.New(string(res.Err))
	}

	return &res.StatusResponse, nil
}
