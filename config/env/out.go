//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strings"
)

// BannerEnabledVal returns whether to show the initial banner on app start
// based on the SPIKE_BANNER_ENABLED environment variable.
//
// The function reads the SPIKE_BANNER_ENABLED environment variable and returns:
//   - true if the variable is not set (default behavior)
//   - true if the variable is set to "true" (case-insensitive)
//   - false for any other value
//
// The environment variable value is trimmed of whitespace and converted to
// lowercase before comparison.
func BannerEnabledVal() bool {
	s := os.Getenv(BannerEnabled)
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return true
	}
	return s == "true"
}

// ShowMemoryWarningVal returns whether to display a warning when the system
// cannot lock memory based on the SPIKE_PILOT_SHOW_MEMORY_WARNING environment
// variable.
//
// The function reads the SPIKE_PILOT_SHOW_MEMORY_WARNING environment variable
// and returns:
//   - false if the variable is not set (default behavior)
//   - true if the variable is set to "true" (case-insensitive)
//   - false for any other value
//
// The environment variable value is trimmed of whitespace and converted to
// lowercase before comparison.
//
// This warning is typically shown when memory locking fails, which could
// impact security-sensitive operations that require pages to remain in RAM.
func ShowMemoryWarningVal() bool {
	s := os.Getenv(PilotShowMemoryWarning)
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return false
	}
	return s == "true"
}

// StackTracesOnLogFatalVal returns whether to print stack traces when
// `log.FatalLn` is called, based on the SPIKE_STACK_TRACES_ON_LOG_FATAL
// environment variable.
//
// The function reads the SPIKE_STACK_TRACES_ON_LOG_FATAL environment variable
// and returns:
//   - false if the variable is not set (default behavior - clean exit)
//   - true if the variable is set to "true" (case-insensitive - panic with stack trace)
//   - false for any other value
//
// The environment variable value is trimmed of whitespace and converted to
// lowercase before comparison.
//
// By default, log.FatalLn performs a clean exit to avoid leaking sensitive
// information in production stack traces. Enable this for development and
// testing purposes.
func StackTracesOnLogFatalVal() bool {
	s := os.Getenv(StackTracesOnLogFatal)
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return false
	}
	return s == "true"
}
