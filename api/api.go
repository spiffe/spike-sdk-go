//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

type Api struct {
	source *workloadapi.X509Source
}

func New() *Api {
	defaultEndpointSocket := spiffe.EndpointSocket()

	source, _, err := spiffe.Source(context.Background(), defaultEndpointSocket)
	if err != nil {
		return nil
	}

	return &Api{source: source}
}

func NewWithSource(source *workloadapi.X509Source) *Api {
	return &Api{source: source}
}

func (a *Api) Close() {
	spiffe.CloseSource(a.source)
}

func (a *Api) CreatePolicy(
	name string, spiffeIdPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) error {
	return CreatePolicy(a.source,
		name, spiffeIdPattern, pathPattern, permissions)
}

func (a *Api) DeletePolicy(name string) error {
	return DeletePolicy(a.source, name)
}

func (a *Api) GetPolicy(name string) (*data.Policy, error) {
	return GetPolicy(a.source, name)
}

func (a *Api) ListPolicies() (*[]data.Policy, error) {
	return ListPolicies(a.source)
}

func (a *Api) DeleteSecretVersions(path string, versions []int) error {
	return DeleteSecret(a.source, path, versions)
}

func (a *Api) DeleteSecret(path string) error {
	return DeleteSecret(a.source, path, []int{})
}

func (a *Api) GetSecretVersioned(path string, version int) (*data.Secret, error) {
	return GetSecret(a.source, path, version)
}

func (a *Api) GetSecret(path string) (*data.Secret, error) {
	return GetSecret(a.source, path, 0)
}

func (a *Api) ListSecretKeys() (*[]string, error) {
	return ListSecretKeys(a.source)
}

func (a *Api) GetSecretMetadata(
	path string, version int,
) (*data.SecretMetadata, error) {
	return GetSecretMetadata(a.source, path, version)
}

func (a *Api) PutSecret(path string, data map[string]string) error {
	return PutSecret(a.source, path, data)
}

func (a *Api) UndeleteSecret(path string, versions []int) error {
	return UndeleteSecret(a.source, path, versions)
}

func (a *Api) Init() error {
	return Init(a.source)
}
