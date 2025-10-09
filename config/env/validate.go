//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "net/url"

// validURL validates that a URL is properly formatted and uses HTTPS
func validURL(urlStr string) bool {
	pu, err := url.Parse(urlStr)
	if err != nil {
		return false
	}
	return pu.Scheme == "https" && pu.Host != ""
}
