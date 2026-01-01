//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package fs

import "sync"

// restrictedPaths contains system directories that should not be used
// for data storage for security and operational reasons.
var restrictedPaths = []string{
	"/", "/etc", "/sys", "/proc", "/dev", "/bin", "/sbin",
	"/usr", "/lib", "/lib64", "/boot", "/root",
}

const spikeHiddenFolderName = ".spike"
const spikeDataFolderName = "data"
const spikeRecoveryFolderName = "recover"

// Cached directory paths and sync.Once for one-time initialization.
var (
	nexusDataPath     string
	nexusDataOnce     sync.Once
	pilotRecoveryPath string
	pilotRecoveryOnce sync.Once
)
