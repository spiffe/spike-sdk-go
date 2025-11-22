//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

import (
	"regexp"
	"time"
)

type PolicyPermission string

// PermissionRead gives permission to read secrets.
// This DOES NOT include listing secrets.
const PermissionRead PolicyPermission = "read"

// PermissionWrite gives permission to write (including
// create, update and delete) secrets.
const PermissionWrite PolicyPermission = "write"

// PermissionList gives permission to list available secrets or resources.
const PermissionList PolicyPermission = "list"

// TODO: verify proper usage of PermissionExecute in SPIKE code.
// UPDATE CHANGELOG if behavior changes.

// PermissionExecute grants the ability to execute specified resources.
// One such resource is encryption and decryption operations that
// don't necessarily persist anything but execute an internal command.
const PermissionExecute PolicyPermission = "execute"

// PermissionSuper gives superuser permissions.
// The user is the alpha and the omega.
const PermissionSuper PolicyPermission = "super"

// Policy represents a security policy applied within SPIKE.
// It includes details such as ID, name, patterns, permissions, and metadata.
type Policy struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	SPIFFEIDPattern string             `json:"spiffiedPattern"`
	PathPattern     string             `json:"pathPattern"`
	Permissions     []PolicyPermission `json:"permissions"`
	CreatedAt       time.Time          `json:"createdAt"`
	CreatedBy       string             `json:"createdBy"`

	// Unexported fields won't be serialized to JSON
	IDRegex   *regexp.Regexp `json:"-"`
	PathRegex *regexp.Regexp `json:"-"`
}

// PolicySpec defines the specification of a policy configuration.
// Name specifies the name of the policy.
// SpiffeIDPattern specifies the SPIFFE ID regex pattern for the policy.
// PathPattern defines the path regex pattern associated with the policy.
// Permissions lists the permissions granted by the policy.
type PolicySpec struct {
	Name            string             `yaml:"name"`
	SpiffeIDPattern string             `yaml:"spiffeidPattern"`
	PathPattern     string             `yaml:"pathPattern"`
	Permissions     []PolicyPermission `json:"permissions"`
}
