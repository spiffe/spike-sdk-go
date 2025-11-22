//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package log

import (
	"fmt"
	"os"
	"strings"
)

// Cannot import "env" here because of circular dependency.
const stackTracesOnLogFatalEnvVar = "SPIKE_STACK_TRACES_ON_LOG_FATAL"

// stackTracesOnLogFatalVal checks if stack traces should be enabled for fatal
// log calls by reading the SPIKE_STACK_TRACES_ON_LOG_FATAL environment
// variable.
//
// Returns:
//   - bool: true if the environment variable is set to "true"
//     (case-insensitive),
//     false otherwise or if the variable is empty/unset
func stackTracesOnLogFatalVal() bool {
	s := os.Getenv(stackTracesOnLogFatalEnvVar)
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return false
	}
	return s == "true"
}

// fatalExit terminates the program with exit code 1, or panics with a stack
// trace if SPIKE_STACK_TRACES_ON_LOG_FATAL is enabled. This provides a way
// to get detailed stack traces for debugging during development while using
// clean exits in production.
//
// Parameters:
//   - fName: the name of the calling function for stack trace identification
//   - args: variadic arguments to include in the panic message if stack traces
//     are enabled
func fatalExit(fName string, args []any) {
	if stackTracesOnLogFatalVal() {
		ss := make([]string, len(args))
		for i, arg := range args {
			ss[i] = fmt.Sprint(arg)
		}
		panic(fName + " " + strings.Join(ss, ","))
	}
	os.Exit(1)
}
