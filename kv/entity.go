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

	// Version is the numeric identifier for this version. Version numbers
	// start at 1 and increment with each update.
	Version int

	// DeletedTime indicates when this version was marked as deleted
	// A nil value means the version is active/not deleted
	DeletedTime *time.Time
}

// Metadata tracks control information for versioned data stored at a path.
// It maintains version boundaries and timestamps for the overall data
// collection.
type Metadata struct {
	// CurrentVersion is the newest/latest non-deleted version number.
	// Version numbers start at 1. A value of 0 indicates that all versions
	// have been deleted (no valid version exists).
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

// HasValidVersions returns true if the Value has at least one non-deleted
// version. It iterates through all versions to check their DeletedTime.
//
// Returns:
//   - true if any version has DeletedTime == nil (active version exists)
//   - false if all versions are deleted or no versions exist
//
// Note: This method performs a full scan of all versions. For stores where
// CurrentVersion is maintained correctly (like SPIKE Nexus), checking
// Metadata.CurrentVersion != 0 is more efficient.
func (v *Value) HasValidVersions() bool {
	for _, version := range v.Versions {
		if version.DeletedTime == nil {
			return true
		}
	}
	return false
}

// Empty returns true if the Value has no valid (non-deleted) versions.
// This is the inverse of HasValidVersions() and is useful for identifying
// secrets that can be purged from storage.
//
// Returns:
//   - true if all versions are deleted or no versions exist
//   - false if at least one active version exists
//
// Example:
//
//	if secret.IsEmpty() {
//	    // Safe to remove from storage
//	    kv.Destroy(path)
//	}
func (v *Value) Empty() bool {
	return !v.HasValidVersions()
}
