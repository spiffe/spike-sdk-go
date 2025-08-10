//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

type RestorationStatus struct {
	ShardsCollected int  `json:"collected"`
	ShardsRemaining int  `json:"remaining"`
	Restored        bool `json:"restored"`
}
