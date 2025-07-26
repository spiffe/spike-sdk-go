//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package system

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// KeepAlive blocks the current goroutine until it receives either a
// SIGINT (Ctrl+C) or SIGTERM signal, enabling graceful shutdown of the
// application. Upon receiving a termination signal, it logs the signal type
// and begins the shutdown process.
//
// The function creates a buffered channel to handle OS signals and uses
// signal.Notify to register for SIGINT and SIGTERM signals. It then blocks
// until a signal is received.
//
// This is typically used in the main function to prevent the program from
// exiting immediately and to ensure proper cleanup when the program is
// terminated.
//
// Example usage:
//
//	func main() {
//	    // Initialize your application
//	    setupApp()
//
//	    // Keep the application running until shutdown signal
//	    KeepAlive()
//
//	    // Perform cleanup
//	    cleanup()
//	}
func KeepAlive() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	log.Printf("\nReceived %v signal, shutting down gracefully...\n", sig)
}

// WatchConfig defines the configuration for the Watch function.
type WatchConfig struct {
	// WaitTimeBeforeExit specifies how long to wait after initialization
	// before executing the exit action.
	WaitTimeBeforeExit time.Duration

	// PollInterval defines how frequently to check the initialization predicate.
	PollInterval time.Duration

	// InitializationPredicate is a function that returns true when the watched
	// condition is met and initialization is complete.
	InitializationPredicate func() bool

	// ExitAction is the function to execute after the initialization predicate
	// returns true and the wait time has elapsed.
	ExitAction func()
}

// Watch continuously polls a condition at regular intervals and executes an
// exit action once the condition is met. It will poll using the
// InitializationPredicate function at intervals specified by PollInterval.
// When the predicate returns true, it waits for WaitTimeBeforeExit duration
// and then executes ExitAction.
//
// This function runs indefinitely until the exit action is called, so it
// should typically be run in a goroutine if the exit action doesn't terminate
// the program.
//
// Example:
//
//	config := WatchConfig{
//		WaitTimeBeforeExit: 5 * time.Second,
//		PollInterval: 1 * time.Second,
//		InitializationPredicate: func() bool {
//			return isServiceReady()
//		},
//		ExitAction: func() {
//			log.Println("Service initialized, shutting down watcher")
//			os.Exit(0)
//		},
//	}
//	go Watch(config)
func Watch(config WatchConfig) {
	interval := config.PollInterval
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			log.Println("tick")
			if config.InitializationPredicate() {
				log.Println("initialized...")
				time.Sleep(config.WaitTimeBeforeExit)
				config.ExitAction()
			}
		}
	}
}
