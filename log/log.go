//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package log provides a lightweight thread-safe logging facility
// using structured logging (slog) with JSON output format. It offers a
// singleton logger instance with configurable log levels through environment
// variables and convenience methods for fatal error logging.
package log

import (
	"log/slog"
	"os"
	"strings"
	"sync"

	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

var logger *slog.Logger
var loggerMutex sync.Mutex

// Log returns a thread-safe singleton instance of slog.Logger configured for
// JSON output. If the logger hasn't been initialized, it creates a new instance
// with the log level specified by the environment. Further calls return the
// same logger instance.
//
// By convention, when using the returned logger, the first argument (msg)
// should be the function name (fName) from which the logging call is made.
func Log() *slog.Logger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	if logger != nil {
		return logger
	}

	opts := &slog.HandlerOptions{
		Level: Level(),
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger = slog.New(handler)
	return logger
}

// FatalLn logs a message at Fatal level with a line feed.
// The fName parameter indicates the function name from which the call is made.
// The args parameter contains the values to be logged, which will be formatted
// and joined.
//
// By default, this function exits cleanly with status code 1 to avoid leaking
// sensitive information through stack traces in production. To enable stack
// traces for development and testing, set SPIKE_STACK_TRACES_ON_LOG_FATAL=true.
func FatalLn(fName string, args ...any) {
	Log().Error(fName, args...)
	fatalExit(fName, args)
}

// FatalErr logs an SDK error at Fatal level and exits the program.
// The fName parameter indicates the function name from which the call is made.
// The err parameter is an SDKError that will be logged with its message, code,
// and error text as structured fields.
//
// By default, this function exits cleanly with status code 1 to avoid leaking
// sensitive information through stack traces in production. To enable stack
// traces for development and testing, set SPIKE_STACK_TRACES_ON_LOG_FATAL=true.
func FatalErr(fName string, err sdkErrors.SDKError) {
	FatalLn(
		fName,
		"message", err.Msg,
		"code", err.Code,
		"err", err.Error(),
	)
}

// Cannot get from env.go because of circular dependency.
const systemLogLevelEnvVar = "SPIKE_SYSTEM_LOG_LEVEL"

// Level returns the logging level for the SPIKE components.
//
// It reads from the SPIKE_SYSTEM_LOG_LEVEL environment variable and
// converts it to the corresponding slog.Level value.
// Valid values (case-insensitive) are:
//   - "DEBUG": returns slog.LevelDebug
//   - "INFO": returns slog.LevelInfo
//   - "WARN": returns slog.LevelWarn
//   - "ERROR": returns slog.LevelError
//
// If the environment variable is not set or contains an invalid value,
// it returns the default level slog.LevelWarn.
func Level() slog.Level {
	level := os.Getenv(systemLogLevelEnvVar)
	level = strings.ToUpper(level)

	switch level {
	case "DEBUG":
		return slog.LevelDebug // -4
	case "INFO":
		return slog.LevelInfo // 0
	case "WARN":
		return slog.LevelWarn // 4
	case "ERROR":
		return slog.LevelError // 8
	default:
		return slog.LevelWarn // 4
	}
}
