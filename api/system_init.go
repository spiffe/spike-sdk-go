//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/api/entity/v1/reqres"
	"os"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/internal/config"
	"github.com/spiffe/spike-sdk-go/api/internal/url"
	"github.com/spiffe/spike-sdk-go/net"
)

// Init sends an initialization request to SPIKE Nexus.
func Init(source *workloadapi.X509Source) error {
	r := reqres.InitRequest{}
	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New("initialization: I am having problem generating the payload"),
			err,
		)
	}

	// TODO: maybe restrict init and ACL operations at this level without doing any policy checks.
	var truer = func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)

	body, err := net.Post(client, url.Init(), mr)

	var res reqres.InitResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Join(
			errors.New("initialization: Problem parsing response body"),
			err,
		)
	}

	if res.RecoveryToken == "" {
		fmt.Println("Failed to get recovery token")
		return errors.New("failed to get recovery token")
	}

	err = os.WriteFile(
		config.SpikePilotRootKeyRecoveryFile(), []byte(res.RecoveryToken), 0600,
	)
	if err != nil {
		fmt.Println("Failed to save token to file:")
		fmt.Println(err.Error())
		return errors.New("failed to save token to file")
	}

	return nil
}

// CheckInitState sends a checkInitState request to SPIKE Nexus.
func CheckInitState(source *workloadapi.X509Source) (data.InitState, error) {
	r := reqres.CheckInitStateRequest{}
	mr, err := json.Marshal(r)
	if err != nil {
		return data.NotInitialized, errors.Join(
			errors.New(
				"checkInitState: I am having problem generating the payload",
			),
			err,
		)
	}

	var truer = func(string) bool { return true }
	client, err := net.CreateMtlsClient(source, truer)
	body, err := net.Post(client, url.InitState(), mr)

	if err != nil {
		return data.NotInitialized, errors.Join(
			errors.New(
				"checkInitState: I am having problem sending the request",
			), err)
	}

	var res reqres.CheckInitStateResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return data.NotInitialized, errors.Join(
			errors.New("checkInitState: Problem parsing response body"),
			err,
		)
	}

	state := res.State

	return state, nil
}
