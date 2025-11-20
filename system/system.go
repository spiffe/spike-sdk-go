//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package system

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// KeepAlive blocks the current goroutine until it receives either a
// SIGINT (Ctrl+C) or SIGTERM signal, enabling graceful shutdown of the
// application.
//
// The function creates a buffered channel to handle OS signals and uses
// signal.Notify to register for SIGINT and SIGTERM signals. It then blocks
// until a signal is received.
//
// An optional callback can be provided to handle the received signal. If no
// callback is provided, no action is taken when a signal is received (the
// function simply returns). This allows callers to handle logging, cleanup,
// or other actions as needed.
//
// This is typically used in the main function to prevent the program from
// exiting immediately and to ensure proper cleanup when the program is
// terminated.
//
// Parameters:
//   - onSignal: Optional callback invoked when a signal is received, with the
//     signal as parameter. If not provided, the function returns silently.
//
// Example usage:
//
//	func main() {
//	    // Initialize your application
//	    setupApp()
//
//	    // Keep the application running until shutdown signal
//	    KeepAlive(func(sig os.Signal) {
//	        log.Printf("Received %v signal, shutting down gracefully...\n", sig)
//	    })
//
//	    // Perform cleanup
//	    cleanup()
//	}
//
// Example without callback:
//
//	func main() {
//	    setupApp()
//	    KeepAlive()  // Simply blocks until signal, no logging
//	    cleanup()
//	}
func KeepAlive(onSignal ...func(os.Signal)) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	if len(onSignal) > 0 && onSignal[0] != nil {
		onSignal[0](sig)
	}
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

	// OnTick is an optional callback invoked on each polling interval.
	// If nil, no action is taken on tick.
	OnTick func()

	// OnInitialized is an optional callback invoked when the initialization
	// predicate returns true, before waiting and executing the exit action.
	// If nil, no action is taken on initialization.
	OnInitialized func()
}

// Watch continuously polls a condition at regular intervals and executes an
// exit action once the condition is met. It will poll using the
// InitializationPredicate function at intervals specified by PollInterval.
// When the predicate returns true, it invokes the OnInitialized callback (if
// provided), waits for WaitTimeBeforeExit duration, and then executes
// ExitAction.
//
// The OnTick callback (if provided) is invoked on each polling interval before
// checking the initialization predicate. The OnInitialized callback (if
// provided)
// is invoked when the predicate first returns true.
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
//		OnTick: func() {
//			log.Println("Checking service status...")
//		},
//		OnInitialized: func() {
//			log.Println("Service initialized successfully")
//		},
//		ExitAction: func() {
//			log.Println("Shutting down watcher")
//			os.Exit(0)
//		},
//	}
//	go Watch(config)
func Watch(config WatchConfig) {
	interval := config.PollInterval
	ticker := time.NewTicker(interval)

	for range ticker.C {
		if config.OnTick != nil {
			config.OnTick()
		}

		if config.InitializationPredicate() {
			if config.OnInitialized != nil {
				config.OnInitialized()
			}
			time.Sleep(config.WaitTimeBeforeExit)
			config.ExitAction()
		}
	}
}
