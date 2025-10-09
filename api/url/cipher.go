//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

import (
	"net/url"

	"github.com/spiffe/spike-sdk-go/config/env"
)

// CipherEncrypt returns the URL for encrypting a text.
func CipherEncrypt() string {
	u, _ := url.JoinPath(
		env.NexusAPIRootVal(),
		string(NexusCipherEncrypt),
	)
	return u
}

// CipherDecrypt returns the URL for decrypting a text.
func CipherDecrypt() string {
	u, _ := url.JoinPath(
		env.NexusAPIRootVal(),
		string(NexusCipherDecrypt),
	)
	return u
}
