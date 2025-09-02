//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package cipher

// Mode selects how encrypt/decrypt requests are made to Nexus.
type Mode string

const (
	ModeStream Mode = "stream"
	ModeJSON   Mode = "json"
)
