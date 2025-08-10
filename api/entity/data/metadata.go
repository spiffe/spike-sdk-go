//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

import "time"

// SecretVersionInfo for secrets version
type SecretVersionInfo struct {
	CreatedTime time.Time  `json:"createdTime"`
	Version     int        `json:"version"`
	DeletedTime *time.Time `json:"deletedTime"`
}

// SecretMetaDataContent for secrets raw metadata
type SecretMetaDataContent struct {
	CurrentVersion int       `json:"currentVersion"`
	OldestVersion  int       `json:"oldestVersion"`
	CreatedTime    time.Time `json:"createdTime"`
	UpdatedTime    time.Time `json:"updatedTime"`
	MaxVersions    int       `json:"maxVersions"`
}

// SecretMetadata for secrets metadata
type SecretMetadata struct {
	Versions map[int]SecretVersionInfo `json:"versions,omitempty"`
	Metadata SecretMetaDataContent     `json:"metadata,omitempty"`
}
