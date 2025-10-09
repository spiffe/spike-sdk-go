//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strconv"
	"time"
)

// DatabaseJournalModeVal returns the SQLite journal mode to use.
// It can be configured using the SPIKE_NEXUS_DB_JOURNAL_MODE environment
// variable.
//
// If the environment variable is not set, it defaults to "WAL"
// (Write-Ahead Logging).
func DatabaseJournalModeVal() string {
	s := os.Getenv(NexusDBJournalMode)
	if s != "" {
		return s
	}
	return "WAL"
}

// DatabaseBusyTimeoutMsVal returns the SQLite busy timeout in milliseconds.
// It can be configured using the SPIKE_NEXUS_DB_BUSY_TIMEOUT_MS environment
// variable. The value must be a positive integer.
//
// If the environment variable is not set or contains an invalid value,
// it defaults to 5000 milliseconds (5 seconds).
func DatabaseBusyTimeoutMsVal() int {
	p := os.Getenv(NexusDBBusyTimeoutMS)
	if p != "" {
		bt, err := strconv.Atoi(p)
		if err == nil && bt > 0 {
			return bt
		}
	}

	return 5000
}

// DatabaseMaxOpenConnsVal returns the maximum number of open database
// connections. It can be configured using the SPIKE_NEXUS_DB_MAX_OPEN_CONNS
// environment variable. The value must be a positive integer.
//
// If the environment variable is not set or contains an invalid value,
// it defaults to 10 connections.
func DatabaseMaxOpenConnsVal() int {
	p := os.Getenv(NexusDBMaxOpenConns)
	if p != "" {
		moc, err := strconv.Atoi(p)
		if err == nil && moc > 0 {
			return moc
		}
	}

	return 10
}

// DatabaseMaxIdleConnsVal returns the maximum number of idle database
// connections. It can be configured using the SPIKE_NEXUS_DB_MAX_IDLE_CONNS
// environment variable. The value must be a positive integer.
//
// If the environment variable is not set or contains an invalid value,
// it defaults to 5 connections.
func DatabaseMaxIdleConnsVal() int {
	p := os.Getenv(NexusDBMaxIdleConns)
	if p != "" {
		mic, err := strconv.Atoi(p)
		if err == nil && mic > 0 {
			return mic
		}
	}

	return 5
}

// DatabaseConnMaxLifetimeSecVal returns the maximum lifetime duration for a
// database connection. It can be configured using the
// SPIKE_NEXUS_DB_CONN_MAX_LIFETIME environment variable.
// The value should be a valid Go duration string (e.g., "1h", "30m").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 1 hour.
func DatabaseConnMaxLifetimeSecVal() time.Duration {
	p := os.Getenv(NexusDBConnMaxLifetime)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return time.Hour
}

// DatabaseOperationTimeoutVal returns the duration to use for database
// operations. It can be configured using the SPIKE_NEXUS_DB_OPERATION_TIMEOUT
// environment variable. The value should be a valid Go duration string
// (e.g., "10s", "1m").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 15 seconds.
func DatabaseOperationTimeoutVal() time.Duration {
	p := os.Getenv(NexusDBOperationTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 15 * time.Second
}

// DatabaseInitializationTimeoutVal returns the duration to wait for database
// initialization.
//
// The timeout is read from the environment variable
// `SPIKE_NEXUS_DB_INITIALIZATION_TIMEOUT`. If this variable is set and its
// value can be parsed as a duration (e.g., "1m30s"), it is used.
// Otherwise, the function defaults to a timeout of 30 seconds.
func DatabaseInitializationTimeoutVal() time.Duration {
	p := os.Getenv(NexusDBInitializationTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}
	return 30 * time.Second
}

// DatabaseSkipSchemaCreationVal determines if schema creation should be
// skipped. It checks the "SPIKE_NEXUS_DB_SKIP_SCHEMA_CREATION" env variable
// to decide.
// If the env variable is set and its value is "true", it returns true.
// Otherwise, it returns false.
func DatabaseSkipSchemaCreationVal() bool {
	p := os.Getenv(NexusDBSkipSchemaCreation)
	if p != "" {
		s, err := strconv.ParseBool(p)
		if err == nil {
			return s
		}
	}
	return false
}
