//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"encoding/json"
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

func CreatePolicy(source *workloadapi.X509Source,
	name string, spiffeIdPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) error {
	r := reqres.PolicyCreateRequest{
		Name:            name,
		SpiffeIdPattern: spiffeIdPattern,
		PathPattern:     pathPattern,
		Permissions:     permissions,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New(
				"createPolicy: I am having problem generating the payload",
			),
			err,
		)
	}

	var truer = func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)
	if err != nil {
		return err
	}

	body, err := net.Post(client, url.PolicyCreate(), mr)
	if err != nil {
		return err
	}

	res := reqres.PolicyCreateResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Join(
			errors.New("createPolicy: Problem parsing response body"),
			err,
		)
	}
	if res.Err != "" {
		return errors.New(string(res.Err))
	}

	return nil
}
