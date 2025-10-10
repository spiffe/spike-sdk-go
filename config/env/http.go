//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import (
	"os"
	"strconv"
	"time"
)

// HTTPClientDialerKeepAliveVal returns the keep-alive duration for the HTTP
// client's dialer. It can be configured using the
// SPIKE_HTTP_CLIENT_DIALER_KEEP_ALIVE environment variable.
// The value should be a valid Go duration string (e.g., "30s", "1m").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 30 seconds.
func HTTPClientDialerKeepAliveVal() time.Duration {
	p := os.Getenv(HTTPClientDialerKeepAlive)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 30 * time.Second
}

// HTTPClientDialerTimeoutVal returns the timeout duration for the HTTP
// client's dialer. It can be configured using the
// SPIKE_HTTP_CLIENT_DIALER_TIMEOUT environment variable.
// The value should be a valid Go duration string (e.g., "30s", "1m").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 30 seconds.
func HTTPClientDialerTimeoutVal() time.Duration {
	p := os.Getenv(HTTPClientDialerTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 30 * time.Second
}

// HTTPClientExpectContinueTimeoutVal returns the timeout for Expect: 100-continue
// responses from the server. It can be configured using the
// SPIKE_HTTP_CLIENT_EXPECT_CONTINUE_TIMEOUT environment variable.
// The value should be a valid Go duration string (e.g., "5s", "10s").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 5 seconds.
func HTTPClientExpectContinueTimeoutVal() time.Duration {
	p := os.Getenv(HTTPClientExpectContinueTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 5 * time.Second
}

// HTTPClientIdleConnTimeoutVal returns the maximum duration an idle connection
// will remain idle before closing. It can be configured using the
// SPIKE_HTTP_CLIENT_IDLE_CONN_TIMEOUT environment variable.
// The value should be a valid Go duration string (e.g., "30s", "1m").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 30 seconds.
func HTTPClientIdleConnTimeoutVal() time.Duration {
	p := os.Getenv(HTTPClientIdleConnTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 30 * time.Second
}

// HTTPClientMaxConnsPerHostVal returns the maximum number of connections
// per host. It can be configured using the SPIKE_HTTP_CLIENT_MAX_CONNS_PER_HOST
// environment variable. The value must be a positive integer.
//
// If the environment variable is not set or contains an invalid value,
// it defaults to 10 connections.
func HTTPClientMaxConnsPerHostVal() int {
	p := os.Getenv(HTTPClientMaxConnsPerHost)
	if p != "" {
		moc, err := strconv.Atoi(p)
		if err == nil && moc > 0 {
			return moc
		}
	}

	return 10
}

// HTTPClientMaxIdleConnsVal returns the maximum number of idle connections
// across all hosts. It can be configured using the
// SPIKE_HTTP_CLIENT_MAX_IDLE_CONNS environment variable.
// The value must be a positive integer.
//
// If the environment variable is not set or contains an invalid value,
// it defaults to 100 connections.
func HTTPClientMaxIdleConnsVal() int {
	p := os.Getenv(HTTPClientMaxIdleConns)
	if p != "" {
		mic, err := strconv.Atoi(p)
		if err == nil && mic > 0 {
			return mic
		}
	}

	return 100
}

// HTTPClientMaxIdleConnsPerHostVal returns the maximum number of idle
// connections per host. It can be configured using the
// SPIKE_HTTP_CLIENT_MAX_IDLE_CONNS_PER_HOST environment variable.
// The value must be a positive integer.
//
// If the environment variable is not set or contains an invalid value,
// it defaults to 10 connections.
func HTTPClientMaxIdleConnsPerHostVal() int {
	p := os.Getenv(HTTPClientMaxIdleConnsPerHost)
	if p != "" {
		mic, err := strconv.Atoi(p)
		if err == nil && mic > 0 {
			return mic
		}
	}

	return 10
}

// HTTPClientResponseHeaderTimeoutVal returns the timeout for waiting for a
// server's response headers. It can be configured using the
// SPIKE_HTTP_CLIENT_RESPONSE_HEADER_TIMEOUT environment variable.
// The value should be a valid Go duration string (e.g., "10s", "30s").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 10 seconds.
func HTTPClientResponseHeaderTimeoutVal() time.Duration {
	p := os.Getenv(HTTPClientResponseHeaderTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 10 * time.Second
}

// HTTPClientTimeoutVal returns the overall timeout for HTTP client requests.
// It can be configured using the SPIKE_HTTP_CLIENT_TIMEOUT environment variable.
// The value should be a valid Go duration string (e.g., "60s", "2m").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 60 seconds.
func HTTPClientTimeoutVal() time.Duration {
	p := os.Getenv(HTTPClientTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 60 * time.Second
}

// HTTPClientTLSHandshakeTimeoutVal returns the timeout for the TLS handshake.
// It can be configured using the SPIKE_HTTP_CLIENT_TLS_HANDSHAKE_TIMEOUT
// environment variable. The value should be a valid Go duration string
// (e.g., "10s", "30s").
//
// If the environment variable is not set or contains an invalid duration,
// it defaults to 10 seconds.
func HTTPClientTLSHandshakeTimeoutVal() time.Duration {
	p := os.Getenv(HTTPClientTLSHandshakeTimeout)
	if p != "" {
		d, err := time.ParseDuration(p)
		if err == nil {
			return d
		}
	}

	return 10 * time.Second
}
