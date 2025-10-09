//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package system provides utilities for application lifecycle management
// and graceful shutdown handling.
//
// The package includes functionality for:
//   - Graceful shutdown on OS signals (SIGINT, SIGTERM)
//   - Condition-based watching with polling and exit actions
//   - Application keep-alive for long-running services
//
// Signal Handling:
//
// KeepAlive blocks until the application receives a termination signal,
// enabling graceful shutdown:
//
//	func main() {
//	    setupApp()
//	    defer cleanup()
//	    system.KeepAlive()
//	}
//
// Condition Watching:
//
// Watch continuously polls a condition and executes an action when the
// condition becomes true, useful for initialization monitoring:
//
//	config := system.WatchConfig{
//	    WaitTimeBeforeExit: 5 * time.Second,
//	    PollInterval: 1 * time.Second,
//	    InitializationPredicate: func() bool {
//	        return isServiceReady()
//	    },
//	    ExitAction: func() {
//	        log.Println("Service ready")
//	    },
//	}
//	go system.Watch(config)
//
// This package is typically used in main functions to control application
// lifecycle and ensure proper cleanup during shutdown.
package system
