//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "os"

// BootstrapConfigMapNameVal returns the name of the ConfigMap used to store
// SPIKE Bootstrap state information.
//
// It retrieves the ConfigMap name from the SPIKE_BOOTSTRAP_CONFIGMAP_NAME
// environment variable. If the environment variable is not set, it returns
// the default value "spike-bootstrap-state".
//
// Returns:
//   - A string containing the ConfigMap name for storing bootstrap state
func BootstrapConfigMapNameVal() string {
	cn := os.Getenv(BootstrapConfigMapName)
	if cn == "" {
		return "spike-bootstrap-state"
	}
	return cn
}
