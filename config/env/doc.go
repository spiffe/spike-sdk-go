//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package env provides environment variable configuration for SPIKE components.
// It defines constants for all SPIKE environment variables and provides utility
// functions to read and parse these variables with appropriate defaults.
//
// The package covers configuration for:
//   - SPIKE Nexus (API URLs, backend storage, database settings, TLS ports)
//   - SPIKE Keeper (peers, update intervals, TLS ports)
//   - SPIKE Pilot (recovery directories, memory warnings)
//   - Trust roots (bootstrap, keeper, nexus, pilot, lite workload)
//   - Shamir secret sharing (shares, threshold)
//   - Recovery and validation settings
//   - Logging and system-level configuration
//
// All configuration values can be customized via environment variables with
// sensible defaults provided when variables are not set.
package env
