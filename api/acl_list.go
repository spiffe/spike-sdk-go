//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"encoding/json"
	"errors"
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

func ListPolicies(source *workloadapi.X509Source) (*[]data.Policy, error) {
	r := reqres.PolicyListRequest{}
	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New("listPolicies: I am having problem generating the payload"),
			err,
		)
	}

	var truer = func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)
	if err != nil {
		return nil, err
	}

	body, err := net.Post(client, url.PolicyList(), mr)
	if err != nil {
		if errors.Is(err, net.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var res reqres.PolicyListResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Join(
			errors.New("getPolicy: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return nil, errors.New(string(res.Err))
	}

	return &res.Policies, nil
}
