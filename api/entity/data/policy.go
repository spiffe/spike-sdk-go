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

// PermissionExecute grants the ability to execute specified resources.
// One such resource is encryption and decryption operations that
// don't necessarily persist anything but execute an internal command.
const PermissionExecute PolicyPermission = "execute"

// PermissionSuper gives superuser permissions.
// The user is the alpha and the omega.
const PermissionSuper PolicyPermission = "super"

type Policy struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	SPIFFEIDPattern string             `json:"spiffeidPattern"`
	PathPattern     string             `json:"pathPattern"`
	Permissions     []PolicyPermission `json:"permissions"`
	CreatedAt       time.Time          `json:"createdAt"`
	CreatedBy       string             `json:"createdBy"`

	// Unexported fields won't be serialized to JSON
	IDRegex   *regexp.Regexp `json:"-"`
	PathRegex *regexp.Regexp `json:"-"`
}
