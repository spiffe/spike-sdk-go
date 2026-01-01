//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package fs provides filesystem utilities for managing SPIKE data directories.
//
// This package handles the creation and resolution of directories used by SPIKE
// components for storing encrypted backups, recovery shards, and other
// persistent data. It ensures directories are created with secure permissions
// (0700) and validates paths to prevent usage of restricted system directories.
//
// # Directory Resolution
//
// The package provides two main directory accessors:
//
//   - [NexusDataFolder]: Returns the path where Nexus stores encrypted backups
//   - [PilotRecoveryFolder]: Returns the path where recovery shards are stored
//
// Each function follows a resolution order:
//
//  1. Environment variable (SPIKE_NEXUS_DATA_DIR or SPIKE_PILOT_RECOVERY_DIR)
//  2. User's home directory (~/.spike/data or ~/.spike/recover)
//  3. Temporary directory with user isolation (/tmp/.spike-$USER/...)
//
// # Thread Safety
//
// Directory paths are resolved once on first access using [sync.Once] and
// cached for subsequent calls, making all accessor functions safe for
// concurrent use.
//
// # Security
//
// The package validates custom directory paths to prevent usage of restricted
// system locations such as /, /etc, /sys, /proc, /dev, /bin, /sbin, /usr,
// /lib, /lib64, /boot, and /root. All directories are created with mode 0700
// to restrict access to the owner only.
package fs
