//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

// Package retry provides a flexible and type-safe retry mechanism with
// exponential backoff. It allows for customizable retry strategies and
// notifications while maintaining context awareness and cancellation support.
package retry

import (
	"context"
	"errors"
	"time"

	"github.com/cenkalti/backoff/v4"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// Default configuration values for the exponential backoff strategy
const (
	// Initial wait time between retries
	defaultInitialInterval = 500 * time.Millisecond
	// Maximum wait time between retries
	defaultMaxInterval = 60 * time.Second
	// Maximum total time for all retry attempts
	defaultMaxElapsedTime = 1200 * time.Second
	// A zero max elapsed time means try forever.
	forever = 0
	// Factor by which the wait time increases
	defaultMultiplier = 2.0
)

// Retrier defines the interface for retry operations with backoff support.
// Implementations of this interface provide different retry strategies.
type Retrier interface {
	// RetryWithBackoff executes an operation with a backoff strategy.
	// It will repeatedly execute the operation until it succeeds or
	// the context is canceled. The backoff strategy determines the
	// delay between retry attempts.
	//
	// Parameters:
	//   - ctx: Context for cancellation and timeout control
	//   - op: The operation to retry, returns error if the attempt failed
	//
	// Returns:
	//   - *sdkErrors.SDKError: nil if successful, or one of the following:
	//   - ErrRetryMaxElapsedTimeReached: if maximum elapsed time is reached
	//   - ErrRetryContextCanceled: if context is canceled
	//   - The last error from the operation
	RetryWithBackoff(
		ctx context.Context, op func() *sdkErrors.SDKError,
	) *sdkErrors.SDKError
}

// TypedRetrier provides type-safe retry operations for functions that return
// both a value and an error. It wraps a base Retrier to provide typed results.
type TypedRetrier[T any] struct {
	retrier Retrier
}

// NewTypedRetrier creates a new TypedRetrier with the given base Retrier.
// This allows for type-safe retry operations while reusing existing retry
// logic.
//
// Parameters:
//   - r: The base Retrier implementation to wrap
//
// Returns:
//   - *TypedRetrier[T]: A new TypedRetrier instance for the specified type
//
// Example:
//
//		retrier := NewTypedRetrier[string](NewExponentialRetrier())
//		result, err := retrier.RetryWithBackoff(ctx, func() (
//	 	string, *sdkErrors.SDKError) {
//		    return callExternalService()
//		})
func NewTypedRetrier[T any](r Retrier) *TypedRetrier[T] {
	return &TypedRetrier[T]{retrier: r}
}

// RetryWithBackoff executes a typed operation with a backoff strategy.
// It preserves the return value while maintaining retry functionality.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - op: The operation to retry, returns both a value and an error
//
// Returns:
//   - T: The result value from the successful operation
//   - *sdkErrors.SDKError: nil if successful, or one of the following errors:
//   - ErrRetryMaxElapsedTimeReached: if maximum elapsed time is reached
//   - ErrRetryContextCanceled: if context is canceled
//   - The wrapped error from the operation if it fails
func (r *TypedRetrier[T]) RetryWithBackoff(
	ctx context.Context,
	op func() (T, *sdkErrors.SDKError),
) (T, *sdkErrors.SDKError) {
	var result T
	err := r.retrier.RetryWithBackoff(ctx, func() *sdkErrors.SDKError {
		var opErr *sdkErrors.SDKError
		result, opErr = op()
		return opErr
	})
	return result, err
}

// NotifyFn is a callback function type for retry notifications.
// It provides information about each retry attempt, including the error,
// current interval duration, and total elapsed time.
type NotifyFn func(
	err *sdkErrors.SDKError, duration, totalDuration time.Duration,
)

// RetrierOption is a function type for configuring ExponentialRetrier.
// It follows the functional options pattern for flexible configuration.
type RetrierOption func(*ExponentialRetrier)

// ExponentialRetrier implements Retrier using exponential backoff strategy.
// It provides configurable retry intervals and maximum attempt durations.
type ExponentialRetrier struct {
	newBackOff func() backoff.BackOff
	notify     NotifyFn
}

// BackOffOption is a function type for configuring ExponentialBackOff.
// It allows fine-tuning of the backoff strategy parameters.
type BackOffOption func(*backoff.ExponentialBackOff)

// NewExponentialRetrier creates a new ExponentialRetrier with configurable
// settings. Default values provide sensible backoff behavior for most use
// cases.
//
// Default settings:
//   - InitialInterval: 500ms
//   - MaxInterval: 60s
//   - MaxElapsedTime: 1200s (20 minutes)
//   - Multiplier: 2.0
//
// Parameters:
//   - opts: Optional configuration functions to customize retry behavior
//
// Returns:
//   - *ExponentialRetrier: A configured retrier instance ready for use
//
// Example:
//
//	retrier := NewExponentialRetrier(
//	    WithBackOffOptions(
//	        WithInitialInterval(100 * time.Millisecond),
//	        WithMaxInterval(5 * time.Second),
//	    ),
//	    WithNotify(func(err *sdkErrors.SDKError, d, total time.Duration) {
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

// RetryWithBackoff implements the Retrier interface using exponential backoff.
// It executes the operation repeatedly until success or context cancellation.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - operation: The function to retry that returns an error
//
// Returns:
//   - *sdkErrors.SDKError: nil if the operation eventually succeeds, or one of:
//   - ErrRetryMaxElapsedTimeReached: if maximum elapsed time is reached
//   - ErrRetryContextCanceled: if context is canceled
//   - The last error from the operation
func (r *ExponentialRetrier) RetryWithBackoff(
	ctx context.Context,
	operation func() *sdkErrors.SDKError,
) *sdkErrors.SDKError {
	b := r.newBackOff()
	totalDuration := time.Duration(0)

	wrappedOp := func() error {
		sdkErr := operation()
		if sdkErr == nil {
			return nil
		}

		if sdkErr.Is(sdkErrors.ErrRetryMaximumAttemptsReached) {
			return backoff.Permanent(sdkErr)
		}

		return sdkErr
	}

	err := backoff.RetryNotify(
		wrappedOp,
		backoff.WithContext(b, ctx),
		func(err error, duration time.Duration) {
			totalDuration += duration

			if r.notify == nil {
				return
			}

			var sdkErr *sdkErrors.SDKError
			if errors.As(err, &sdkErr) {
				r.notify(sdkErr, duration, totalDuration)
			} else {
				wrapped := sdkErrors.ErrRetryOperationFailed.Wrap(err)
				r.notify(wrapped, duration, totalDuration)
			}
		},
	)

	if err == nil {
		return nil
	}

	var sdkErr *sdkErrors.SDKError
	if errors.As(err, &sdkErr) {
		return sdkErr
	}

	if errors.Is(err, context.Canceled) {
		failErr := sdkErrors.ErrRetryContextCanceled.Wrap(err)
		failErr.Msg = "retry operation canceled"
		return failErr
	}

	if errors.Is(err, context.DeadlineExceeded) {
		failErr := sdkErrors.ErrRetryMaxElapsedTimeReached.Wrap(err)
		failErr.Msg = "maximum retry elapsed time exceeded"
		return failErr
	}

	failErr := sdkErrors.ErrRetryOperationFailed.Wrap(err)
	failErr.Msg = "retry operation failed"
	return failErr
}

// WithBackOffOptions configures the backoff settings using the provided
// options. Multiple options can be combined to customize the retry behavior.
//
// Parameters:
//   - opts: One or more BackOffOption functions to configure the backoff
//     strategy
//
// Returns:
//   - RetrierOption: A configuration function for ExponentialRetrier
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
//
// Parameters:
//   - d: The initial wait duration before the first retry
//
// Returns:
//   - BackOffOption: A configuration function for ExponentialBackOff
func WithInitialInterval(d time.Duration) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.InitialInterval = d
	}
}

// WithMaxInterval sets the maximum interval between retries.
// The interval will never exceed this value, regardless of the multiplier.
//
// Parameters:
//   - d: The maximum wait duration between retry attempts
//
// Returns:
//   - BackOffOption: A configuration function for ExponentialBackOff
func WithMaxInterval(d time.Duration) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.MaxInterval = d
	}
}

// WithMaxElapsedTime sets the maximum total time for retries.
// The retry operation will stop after this duration, even if not successful.
// Set to 0 to retry indefinitely (until context is canceled).
//
// Parameters:
//   - d: The maximum total duration for all retry attempts
//
// Returns:
//   - BackOffOption: A configuration function for ExponentialBackOff
func WithMaxElapsedTime(d time.Duration) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.MaxElapsedTime = d
	}
}

// WithMultiplier sets the multiplier for increasing intervals.
// Each retry interval is multiplied by this value, up to MaxInterval.
//
// Parameters:
//   - m: The multiplier factor (e.g., 2.0 doubles the interval each time)
//
// Returns:
//   - BackOffOption: A configuration function for ExponentialBackOff
func WithMultiplier(m float64) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.Multiplier = m
	}
}

// WithRandomizationFactor sets the randomization factor for backoff intervals.
// The actual interval will be randomized between
// [interval * (1 - factor), interval * (1 + factor)].
//
// A factor of 0 disables randomization (deterministic intervals).
// A factor of 0.5 (the default) means intervals can vary by ±50%.
// This randomization helps prevent thundering herd issues in distributed
// systems.
//
// Parameters:
//   - factor: The randomization factor (0.0 to 1.0)
//
// Returns:
//   - BackOffOption: A configuration function for ExponentialBackOff
func WithRandomizationFactor(factor float64) BackOffOption {
	return func(b *backoff.ExponentialBackOff) {
		b.RandomizationFactor = factor
	}
}

// WithNotify is an option to set the notification callback.
// The callback is called after each failed attempt, allowing you to log
// or monitor retry behavior.
//
// Parameters:
//   - fn: Callback function invoked after each failed retry attempt
//
// Returns:
//   - RetrierOption: A configuration function for ExponentialRetrier
//
// Example:
//
//	retrier := NewExponentialRetrier(
//	    WithNotify(func(err *sdkErrors.SDKError, d, total time.Duration) {
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
type Handler[T any] func() (T, *sdkErrors.SDKError)

// Do provides a simplified way to retry a typed operation with configurable
// settings. It creates a TypedRetrier with exponential backoff and applies
// any provided options.
//
// This is a convenience function for common retry scenarios where you don't
// need to create and manage a retrier instance explicitly.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - handler: The function to retry that returns a value and error
//   - options: Optional configuration for the retry behavior
//
// Returns:
//   - T: The result value from the successful operation
//   - *sdkErrors.SDKError: nil if successful, or one of the following errors:
//   - ErrRetryMaxElapsedTimeReached: if maximum elapsed time is reached
//   - ErrRetryContextCanceled: if context is canceled
//   - The wrapped error from the handler if it fails
//
// Example:
//
//	result, err := Do(ctx, func() (string, *sdkErrors.SDKError) {
//	    return fetchData()
//	}, WithNotify(logRetryAttempts))
func Do[T any](
	ctx context.Context, handler Handler[T], options ...RetrierOption,
) (T, *sdkErrors.SDKError) {
	return NewTypedRetrier[T](
		NewExponentialRetrier(options...),
	).RetryWithBackoff(ctx, handler)
}

// Forever retries an operation indefinitely with exponential backoff until it
// succeeds or the context is canceled. It sets MaxElapsedTime to 0, which means
// the retry loop will continue forever (or until the context is canceled).
//
// This is a convenience function that sets up exponential backoff with sensible
// defaults for infinite retry scenarios.
//
// Default settings:
//   - InitialInterval: 500ms
//   - MaxInterval: 60s
//   - MaxElapsedTime: 0 (retry forever)
//   - Multiplier: 2.0
//
// Parameters:
//   - ctx: Context for cancellation control (the only way to stop retrying)
//   - handler: The function to retry that returns a value and error
//   - options: Optional configuration for retry behavior
//
// Note: User-provided options are applied AFTER the default settings and will
// override them. If you pass WithBackOffOptions(WithMaxElapsedTime(...)), it
// will override the "forever" behavior. This allows power users to customize
// the retry behavior while keeping the convenience of preset defaults.
//
// Returns:
//   - T: The result value from the successful operation
//   - *sdkErrors.SDKError: nil if successful, or one of the following errors:
//   - ErrRetryContextCanceled: if context is canceled
//   - The wrapped error from the handler if all retries fail
//
// Example:
//
//		// Retry forever with custom notification
//		result, err := Forever(ctx, func() (string, *sdkErrors.SDKError) {
//		    return fetchData()
//		}, WithNotify(func(err *sdkErrors.SDKError, d, total time.Duration) {
//		    log.Printf("Retry failed: %v (attempt duration: %v, total: %v)",
//		      err, d, total)
//		}))
//
//		// Override behavior (will now stop after 1 minute
//	 //	instead of retrying forever)
//		result, err := Forever(ctx, func() (string, *sdkErrors.SDKError) {
//		    return fetchData()
//		}, WithBackOffOptions(WithMaxElapsedTime(1 * time.Minute)))
func Forever[T any](
	ctx context.Context, handler Handler[T], options ...RetrierOption,
) (T, *sdkErrors.SDKError) {
	ro := WithBackOffOptions(WithMaxElapsedTime(forever))
	ros := []RetrierOption{ro}
	ros = append(ros, options...)

	return NewTypedRetrier[T](
		NewExponentialRetrier(ros...),
	).RetryWithBackoff(ctx, handler)
}

// WithMaxAttempts retries an operation up to a maximum number of attempts with
// exponential backoff. It stops retrying when the operation succeeds, the
// maximum number of attempts is reached, or the context is canceled.
//
// This is a convenience function for retry scenarios where you need to limit
// the number of attempts rather than the total elapsed time.
//
// Default settings:
//   - InitialInterval: 500ms
//   - MaxInterval: 60s
//   - MaxElapsedTime: 0 (unlimited, controlled by maxAttempts)
//   - Multiplier: 2.0
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - maxAttempts: Maximum number of retry attempts (must be > 0)
//   - handler: The function to retry that returns (success bool, error)
//   - success: true if the operation succeeded, false to retry
//   - error: nil on success, or the error that occurred
//   - options: Optional configuration for the retry behavior
//
// Returns:
//   - bool: true if the operation eventually succeeded; false otherwise
//   - *sdkErrors.SDKError: nil if successful, or one of the following errors:
//   - ErrRetryContextCanceled: if context is canceled
//   - ErrRetryOperationFailed: if max attempts reached without success
//   - The last error returned by the handler or a retry framework error
//
// Example:
//
//	success, err := WithMaxAttempts(ctx, 5, func() (bool, *sdkErrors.SDKError) {
//	    result, err := callService()
//	    if err != nil {
//	        return false, err
//	    }
//	    return true, nil
//	})
//
//	// With custom backoff options:
//	success, err := WithMaxAttempts(ctx, 5, handler,
//	    WithBackOffOptions(
//	        WithInitialInterval(2 * time.Second),
//	        WithMaxInterval(30 * time.Second),
//	    ),
//	)
func WithMaxAttempts(
	ctx context.Context,
	maxAttempts int,
	handler func() (bool, *sdkErrors.SDKError),
	options ...RetrierOption,
) (bool, *sdkErrors.SDKError) {
	if maxAttempts <= 0 {
		failErr := sdkErrors.ErrDataInvalidInput.Clone()
		failErr.Msg = "maxAttempts must be greater than 0"
		return false, failErr
	}

	attempts := 0

	// Prepend MaxElapsedTime(forever) so user options can override if needed
	opts := []RetrierOption{WithBackOffOptions(WithMaxElapsedTime(forever))}
	opts = append(opts, options...)

	return NewTypedRetrier[bool](
		NewExponentialRetrier(opts...),
	).RetryWithBackoff(ctx, func() (bool, *sdkErrors.SDKError) {
		attempts++

		success, err := handler()
		if success {
			return true, nil
		}

		if attempts >= maxAttempts {
			failErr := sdkErrors.ErrRetryMaximumAttemptsReached.Clone()
			failErr.Msg = "maximum retry attempts reached"
			return false, failErr
		}

		if err != nil {
			return false, err
		}

		retryErr := sdkErrors.ErrRetryOperationFailed.Clone()
		retryErr.Msg = "operation failed, will retry"
		return false, retryErr
	})
}
