//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package retry provides a flexible and type-safe retry mechanism with
// exponential backoff. It allows for customizable retry strategies and
// notifications while maintaining context awareness and cancellation support.
package retry

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Default configuration values for the exponential backoff strategy
const (
	// Initial wait time between retries
	defaultInitialInterval = 500 * time.Millisecond
	// Maximum wait time between retries
	defaultMaxInterval = 3 * time.Second
	// Maximum total time for all retry attempts
	defaultMaxElapsedTime = 30 * time.Second
	// Factor by which the wait time increases
	defaultMultiplier = 2.0
)

// Retrier defines the interface for retry operations with backoff support.
// Implementations of this interface provide different retry strategies.
type Retrier interface {
	// RetryWithBackoff executes an operation with backoff strategy.
	// It will repeatedly execute the operation until it succeeds or
	// the context is cancelled. The backoff strategy determines the
	// delay between retry attempts.
	//
	// Parameters:
	//   - ctx: Context for cancellation and timeout control
	//   - op: The operation to retry, returns error if attempt failed
	//
	// Returns:
	//   - error: Last error from the operation or nil if successful
	RetryWithBackoff(ctx context.Context, op func() error) error
}

// TypedRetrier provides type-safe retry operations for functions that return
// both a value and an error. It wraps a base Retrier to provide typed results.
type TypedRetrier[T any] struct {
	retrier Retrier
}

// NewTypedRetrier creates a new TypedRetrier with the given base Retrier.
// This allows for type-safe retry operations while reusing existing retry logic.
//
// Example:
//
//	retrier := NewTypedRetrier[string](NewExponentialRetrier())
//	result, err := retrier.RetryWithBackoff(ctx, func() (string, error) {
//	    return callExternalService()
//	})
func NewTypedRetrier[T any](r Retrier) *TypedRetrier[T] {
	return &TypedRetrier[T]{retrier: r}
}

// RetryWithBackoff executes a typed operation with backoff strategy.
// It preserves the return value while maintaining retry functionality.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - op: The operation to retry, returns both a value and an error
//
// Returns:
//   - T: The result value from the successful operation
//   - error: Last error from the operation or nil if successful
func (r *TypedRetrier[T]) RetryWithBackoff(
	ctx context.Context,
	op func() (T, error),
) (T, error) {
	var result T
	err := r.retrier.RetryWithBackoff(ctx, func() error {
		var err error
		result, err = op()
		return err
	})
	return result, err
}

// NotifyFn is a callback function type for retry notifications.
// It provides information about each retry attempt including the error,
// current interval duration, and total elapsed time.
type NotifyFn func(err error, duration, totalDuration time.Duration)

// RetrierOption is a function type for configuring ExponentialRetrier.
// It follows the functional options pattern for flexible configuration.
type RetrierOption func(*ExponentialRetrier)

// ExponentialRetrier implements Retrier using exponential backoff strategy.
// It provides configurable retry intervals and maximum attempt durations..
type ExponentialRetrier struct {
	newBackOff func() backoff.BackOff
	notify     NotifyFn
}

// BackOffOption is a function type for configuring ExponentialBackOff.
// It allows fine-tuning of the backoff strategy parameters.
type BackOffOption func(*backoff.ExponentialBackOff)

// NewExponentialRetrier creates a new ExponentialRetrier with configurable
// settings.
//
// Example:
//
//	retrier := NewExponentialRetrier(
//	    WithBackOffOptions(
//	        WithInitialInterval(100 * time.Millisecond),
//	        WithMaxInterval(5 * time.Second),
//	    ),
//	    WithNotify(func(err error, d, total time.Duration) {
//	        log.Printf("Retry attempt failed: %v", err)
//	    }),
//	)
func NewExponentialRetrier(opts ...RetrierOption) *ExponentialRetrier {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = defaultInitialInterval
	b.MaxInterval = defaultMaxInterval
	b.MaxElapsedTime = defaultMaxElapsedTime
	b.Multiplier = defaultMultiplier

	r := &ExponentialRetrier{
		newBackOff: func() backoff.BackOff {
			return b
		},
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// RetryWithBackoff implements the Retrier interface
func (r *ExponentialRetrier) RetryWithBackoff(
	ctx context.Context,
	operation func() error,
) error {
	b := r.newBackOff()
	totalDuration := time.Duration(0)
	return backoff.RetryNotify(
		operation,
		backoff.WithContext(b, ctx),
		func(err error, duration time.Duration) {
			totalDuration += duration
			if r.notify != nil {
				r.notify(err, duration, totalDuration)
			}
		},
	)
}

// WithBackOffOptions configures the backoff settings using the provided
// options. Multiple options can be combined to customize the retry behavior.
//
// Example:
//
//	retrier := NewExponentialRetrier(
//	    WithBackOffOptions(
//	        WithInitialInterval(1 * time.Second),
//	        WithMaxElapsedTime(1 * time.Minute),
//	    ),
//	)
func WithBackOffOptions(opts ...BackOffOption) RetrierOption {
	return func(r *ExponentialRetrier) {
		b := r.newBackOff().(*backoff.ExponentialBackOff)
		for _, opt := range opts {
			opt(b)
		}
	}
}

// WithInitialInterval sets the initial interval between retries.
// This is the starting point for the exponential backoff calculation.
func WithInitialInterval(d time.Duration) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.InitialInterval = d
	}
}

// WithMaxInterval sets the maximum interval between retries.
// The interval will never exceed this value, regardless of the multiplier.s
func WithMaxInterval(d time.Duration) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.MaxInterval = d
	}
}

// WithMaxElapsedTime sets the maximum total time for retries.
// The retry operation will stop after this duration, even if not successful.
func WithMaxElapsedTime(d time.Duration) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.MaxElapsedTime = d
	}
}

// WithMultiplier sets the multiplier for increasing intervals.
// Each retry interval is multiplied by this value, up to MaxInterval.
func WithMultiplier(m float64) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.Multiplier = m
	}
}

// WithNotify is an option to set the notification callback.
// The callback is called after each failed attempt with retry statistics.
//
// Example:
//
//	retrier := NewExponentialRetrier(
//	    WithNotify(func(err error, d, total time.Duration) {
//	        log.Printf("Attempt failed after %v, total time %v: %v",
//	            d, total, err)
//	    }),
//	)
func WithNotify(fn NotifyFn) RetrierOption {
	return func(r *ExponentialRetrier) {
		r.notify = fn
	}
}

// Handler represents a function that returns a value and an error.
// It's used with the Do helper function for simple retry operations.
type Handler[T any] func() (T, error)

// Do provides a simplified way to retry a typed operation with default
// settings.
// It creates a TypedRetrier with default exponential backoff configuration.
//
// Example:
//
//	result, err := Do(ctx, func() (string, error) {
//	    return fetchData()
//	})
func Do[T any](
	ctx context.Context,
	handler Handler[T],
	options ...RetrierOption,
) (T, error) {
	return NewTypedRetrier[T](
		NewExponentialRetrier(options...),
	).RetryWithBackoff(ctx, handler)
}
