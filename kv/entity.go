//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

import "time"

// Version represents a single version of versioned data along with its
// metadata. Each version maintains its own set of key-value pairs and tracking
// information.
type Version struct {
	// Data contains the actual key-value pairs stored in this version
	Data map[string]string

	// CreatedTime is when this version was created
	CreatedTime time.Time

	// Version is the numeric identifier for this version
	Version int

	// DeletedTime indicates when this version was marked as deleted
	// A nil value means the version is active/not deleted
	DeletedTime *time.Time
}

// Metadata tracks control information for versioned data stored at a path.
// It maintains version boundaries and timestamps for the overall data
// collection.
type Metadata struct {
	// CurrentVersion is the newest/latest version number
	CurrentVersion int

	// OldestVersion is the oldest available version number
	OldestVersion int

	// CreatedTime is when the data at this path was first created
	CreatedTime time.Time

	// UpdatedTime is when the data was last modified
	UpdatedTime time.Time

	// MaxVersions is the maximum number of versions to retain
	// When exceeded, older versions are automatically pruned
	MaxVersions int
}

// Value represents a versioned collection of key-value pairs stored at a
// specific path. It maintains both the version history and metadata about the
// collection as a whole.
type Value struct {
	// Versions maps version numbers to their corresponding Version objects
	Versions map[int]Version

	// Metadata contains control information about this versioned data
	Metadata Metadata
}
