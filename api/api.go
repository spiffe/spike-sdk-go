//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	impl "github.com/spiffe/spike-sdk-go/api/internal/impl/api"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

// Api is the SPIKE API.
type Api struct {
	source *workloadapi.X509Source
}

// New creates a new SPIKE API.
func New() *Api {
	defaultEndpointSocket := spiffe.EndpointSocket()

	source, _, err := spiffe.Source(context.Background(), defaultEndpointSocket)
	if err != nil {
		return nil
	}

	return &Api{source: source}
}

// NewWithSource creates a new SPIKE API with a given X509Source.
func NewWithSource(source *workloadapi.X509Source) *Api {
	return &Api{source: source}
}

// Close closes the SPIKE API.
func (a *Api) Close() {
	spiffe.CloseSource(a.source)
}

// CreatePolicy creates a new policy.
func (a *Api) CreatePolicy(
	name string, spiffeIdPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) error {
	return impl.CreatePolicy(a.source,
		name, spiffeIdPattern, pathPattern, permissions)
}

// DeletePolicy deletes a policy.
func (a *Api) DeletePolicy(name string) error {
	return impl.DeletePolicy(a.source, name)
}

// GetPolicy gets a policy.
func (a *Api) GetPolicy(name string) (*data.Policy, error) {
	return impl.GetPolicy(a.source, name)
}

// ListPolicies lists all policies.
func (a *Api) ListPolicies() (*[]data.Policy, error) {
	return impl.ListPolicies(a.source)
}

// DeleteSecretVersions deletes secret versions.
func (a *Api) DeleteSecretVersions(path string, versions []int) error {
	return impl.DeleteSecret(a.source, path, versions)
}

// DeleteSecret deletes a secret.
func (a *Api) DeleteSecret(path string) error {
	return impl.DeleteSecret(a.source, path, []int{})
}

// GetSecretVersioned gets a secret versioned.
func (a *Api) GetSecretVersioned(
	path string, version int,
) (*data.Secret, error) {
	return impl.GetSecret(a.source, path, version)
}

// GetSecret gets a secret.
func (a *Api) GetSecret(path string) (*data.Secret, error) {
	return impl.GetSecret(a.source, path, 0)
}

// ListSecretKeys lists all secret key paths
func (a *Api) ListSecretKeys() (*[]string, error) {
	return impl.ListSecretKeys(a.source)
}

// GetSecretMetadata gets secret metadata.
func (a *Api) GetSecretMetadata(
	path string, version int,
) (*data.SecretMetadata, error) {
	return impl.GetSecretMetadata(a.source, path, version)
}

// PutSecret puts a secret.
func (a *Api) PutSecret(path string, data map[string]string) error {
	return impl.PutSecret(a.source, path, data)
}

// UndeleteSecret undeletes secret versions.
func (a *Api) UndeleteSecret(path string, versions []int) error {
	return impl.UndeleteSecret(a.source, path, versions)
}

// Init initializes SPIKE Nexus.
func (a *Api) Init() error {
	return impl.Init(a.source)
}
