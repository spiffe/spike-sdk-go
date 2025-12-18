//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package retry

import (
	"context"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
	"github.com/stretchr/testify/assert"
)

func TestExponentialRetrier_Success(t *testing.T) {
	retrier := NewExponentialRetrier()

	// Operation that succeeds immediately
	err := retrier.RetryWithBackoff(context.Background(), func() *sdkErrors.SDKError {
		return nil
	})

	assert.Nil(t, err)
}

func TestExponentialRetrier_EventualSuccess(t *testing.T) {
	attempts := 0
	maxAttempts := 3

	retrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithInitialInterval(1*time.Millisecond),
			WithMaxInterval(5*time.Millisecond),
		),
	)

	err := retrier.RetryWithBackoff(context.Background(), func() *sdkErrors.SDKError {
		attempts++
		if attempts < maxAttempts {
			return sdkErrors.ErrRetryOperationFailed
		}
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, maxAttempts, attempts)
}

func TestExponentialRetrier_Failure(t *testing.T) {
	retrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithMaxElapsedTime(10*time.Millisecond),
			WithInitialInterval(1*time.Millisecond),
		),
	)

	expectedErr := sdkErrors.ErrRetryOperationFailed
	err := retrier.RetryWithBackoff(context.Background(), func() *sdkErrors.SDKError {
		return expectedErr
	})

	assert.Error(t, err)
	assert.True(t, err.Is(expectedErr))
}

func TestExponentialRetrier_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	retrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithInitialInterval(100 * time.Millisecond),
		),
	)

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	err := retrier.RetryWithBackoff(ctx, func() *sdkErrors.SDKError {
		return sdkErrors.ErrRetryOperationFailed
	})

	assert.Error(t, err)
	assert.True(t, err.Is(sdkErrors.ErrRetryContextCanceled))
}

func TestExponentialRetrier_ContextDeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	retrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithInitialInterval(100 * time.Millisecond),
		),
	)

	err := retrier.RetryWithBackoff(ctx, func() *sdkErrors.SDKError {
		return sdkErrors.ErrRetryOperationFailed
	})

	assert.Error(t, err)
	// Context deadline exceeded should be wrapped as ErrRetryMaxElapsedTimeReached
	assert.True(t, err.Is(sdkErrors.ErrRetryMaxElapsedTimeReached))
}

func TestExponentialRetrier_Notification(t *testing.T) {
	var notifications []time.Duration
	totalDurations := make([]time.Duration, 0)

	retrier := NewExponentialRetrier(
		WithNotify(func(_ *sdkErrors.SDKError, duration, totalDuration time.Duration) {
			notifications = append(notifications, duration)
			totalDurations = append(totalDurations, totalDuration)
		}),
		WithBackOffOptions(
			// NOTE: These intervals are intentionally set to reasonable values
			// (10ms+) to ensure test reliability across different systems and
			// CI environments. DO NOT reduce these values without good reason:
			//
			// 1. System timer precision: Many systems have timer resolution
			//    of 1-15ms, making sub-millisecond intervals unreliable
			// 2. Scheduling jitter: OS thread scheduling can cause actual
			//    sleep durations to vary significantly from intended values
			// 3. CI environment variance: Different CI systems (GitHub Actions,
			//    local Docker, etc.) have different performance characteristics
			// 4. Test determinism: Smaller intervals make the test more likely
			//    to fail due to timing race conditions
			//
			// If you need faster tests, consider mocking time instead of
			// reducing these intervals.
			WithInitialInterval(100*time.Millisecond),
			WithMaxInterval(500*time.Millisecond),
			WithMaxElapsedTime(2000*time.Millisecond),
			// Disable randomization (jitter) to ensure deterministic intervals
			// for testing. By default, backoff uses a RandomizationFactor of 0.5,
			// which randomizes intervals by ±50%. This can cause the second retry
			// to have a shorter duration than the first (e.g., first retry: 150ms,
			// second retry: 100ms), breaking the monotonically increasing assumption
			// in our assertions below (lines 134-135).
			WithRandomizationFactor(0),
		),
	)

	attempts := 0
	_ = retrier.RetryWithBackoff(context.Background(), func() *sdkErrors.SDKError {
		attempts++
		if attempts < 3 {
			return sdkErrors.ErrRetryOperationFailed
		}
		return nil
	})

	assert.Equal(t, 2, len(notifications))
	assert.Equal(t, 2, len(totalDurations))

	// Verify that durations are increasing
	assert.Less(t, notifications[0], notifications[1])
	assert.Less(t, totalDurations[0], totalDurations[1])
}

func TestTypedRetrier_Success(t *testing.T) {
	baseRetrier := NewExponentialRetrier()
	typedRetrier := NewTypedRetrier[string](baseRetrier)

	expected := "success"
	result, err := typedRetrier.RetryWithBackoff(context.Background(), func() (string, *sdkErrors.SDKError) {
		return expected, nil
	})

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestTypedRetrier_EventualSuccess(t *testing.T) {
	baseRetrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithInitialInterval(1*time.Millisecond),
			WithMaxInterval(5*time.Millisecond),
		),
	)
	typedRetrier := NewTypedRetrier[string](baseRetrier)

	attempts := 0
	expected := "final-value"

	result, err := typedRetrier.RetryWithBackoff(context.Background(), func() (string, *sdkErrors.SDKError) {
		attempts++
		if attempts < 3 {
			return "", sdkErrors.ErrRetryOperationFailed
		}
		return expected, nil
	})

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 3, attempts)
}

func TestTypedRetrier_Failure(t *testing.T) {
	baseRetrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithMaxElapsedTime(10*time.Millisecond),
			WithInitialInterval(1*time.Millisecond),
		),
	)
	typedRetrier := NewTypedRetrier[int](baseRetrier)

	expectedErr := sdkErrors.ErrRetryOperationFailed
	result, err := typedRetrier.RetryWithBackoff(context.Background(), func() (int, *sdkErrors.SDKError) {
		return 0, expectedErr
	})

	assert.Error(t, err)
	assert.True(t, err.Is(expectedErr))
	assert.Equal(t, 0, result)
}

func TestTypedRetrier_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	baseRetrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithInitialInterval(100 * time.Millisecond),
		),
	)
	typedRetrier := NewTypedRetrier[string](baseRetrier)

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	result, err := typedRetrier.RetryWithBackoff(ctx, func() (string, *sdkErrors.SDKError) {
		return "", sdkErrors.ErrRetryOperationFailed
	})

	assert.Error(t, err)
	assert.True(t, err.Is(sdkErrors.ErrRetryContextCanceled))
	assert.Equal(t, "", result)
}

func TestDo_Success(t *testing.T) {
	expected := "success-value"
	result, err := Do(context.Background(), func() (string, *sdkErrors.SDKError) {
		return expected, nil
	})

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestDo_Failure(t *testing.T) {
	expectedErr := sdkErrors.ErrRetryOperationFailed
	result, err := Do(
		context.Background(),
		func() (int, *sdkErrors.SDKError) {
			return 0, expectedErr
		},
		WithBackOffOptions(
			WithMaxElapsedTime(10*time.Millisecond),
			WithInitialInterval(1*time.Millisecond),
		),
	)

	assert.Error(t, err)
	assert.True(t, err.Is(expectedErr))
	assert.Equal(t, 0, result)
}

func TestDo_WithOptions(t *testing.T) {
	attempts := 0
	expected := "eventual-success"

	result, err := Do(
		context.Background(),
		func() (string, *sdkErrors.SDKError) {
			attempts++
			if attempts < 3 {
				return "", sdkErrors.ErrRetryOperationFailed
			}
			return expected, nil
		},
		WithBackOffOptions(
			WithInitialInterval(1*time.Millisecond),
			WithMaxInterval(5*time.Millisecond),
		),
	)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 3, attempts)
}

func TestForever_EventualSuccess(t *testing.T) {
	attempts := 0
	expected := "eventual-success"

	result, err := Forever(
		context.Background(),
		func() (string, *sdkErrors.SDKError) {
			attempts++
			if attempts < 5 {
				return "", sdkErrors.ErrRetryOperationFailed
			}
			return expected, nil
		},
		WithBackOffOptions(
			WithInitialInterval(1*time.Millisecond),
			WithMaxInterval(5*time.Millisecond),
		),
	)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 5, attempts)
}

func TestForever_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	attempts := 0

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	result, err := Forever(
		ctx,
		func() (string, *sdkErrors.SDKError) {
			attempts++
			return "", sdkErrors.ErrRetryOperationFailed
		},
		WithBackOffOptions(
			WithInitialInterval(1*time.Millisecond),
		),
	)

	assert.Error(t, err)
	assert.True(t, err.Is(sdkErrors.ErrRetryContextCanceled))
	assert.Equal(t, "", result)
	// Should have retried multiple times before cancellation
	assert.Greater(t, attempts, 1)
}

func TestForever_UserOverride(t *testing.T) {
	attempts := 0

	// User overrides Forever's MaxElapsedTime=0 with a specific value
	// This should stop retrying after 10ms
	result, err := Forever(
		context.Background(),
		func() (int, *sdkErrors.SDKError) {
			attempts++
			return 0, sdkErrors.ErrRetryOperationFailed
		},
		WithBackOffOptions(
			WithMaxElapsedTime(10*time.Millisecond),
			WithInitialInterval(1*time.Millisecond),
		),
	)

	assert.Error(t, err)
	assert.True(t, err.Is(sdkErrors.ErrRetryOperationFailed))
	assert.Equal(t, 0, result)
	// Should have stopped after MaxElapsedTime, not retry forever
	assert.Greater(t, attempts, 1)
	assert.Less(t, attempts, 100) // Shouldn't retry too many times
}

func TestForever_VerifiesMaxElapsedTimeZero(t *testing.T) {
	// This test verifies that Forever actually sets MaxElapsedTime to 0
	// by checking the backoff configuration
	retrier := NewExponentialRetrier(
		WithBackOffOptions(WithMaxElapsedTime(forever)),
	)
	b := retrier.newBackOff().(*backoff.ExponentialBackOff)

	assert.Equal(t, time.Duration(0), b.MaxElapsedTime)
}

func TestWithMaxAttempts_EventualSuccess(t *testing.T) {
	attempts := 0

	result, err := WithMaxAttempts(
		context.Background(),
		5,
		func() (bool, *sdkErrors.SDKError) {
			attempts++
			if attempts < 3 {
				return false, sdkErrors.ErrRetryOperationFailed
			}
			return true, nil
		},
	)

	assert.Nil(t, err)
	assert.True(t, result)
	assert.Equal(t, 3, attempts)
}

func TestWithMaxAttempts_MaxAttemptsReached(t *testing.T) {
	attempts := 0

	result, err := WithMaxAttempts(
		context.Background(),
		3,
		func() (bool, *sdkErrors.SDKError) {
			attempts++
			return false, sdkErrors.ErrRetryOperationFailed
		},
	)

	assert.False(t, result)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrRetryOperationFailed))
	assert.Equal(t, 3, attempts)
}

func TestWithMaxAttempts_InvalidMaxAttempts(t *testing.T) {
	result, err := WithMaxAttempts(
		context.Background(),
		0,
		func() (bool, *sdkErrors.SDKError) {
			return true, nil
		},
	)

	assert.False(t, result)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrDataInvalidInput))
}

func TestWithMaxAttempts_FalseWithoutErrorRetries(t *testing.T) {
	attempts := 0

	result, err := WithMaxAttempts(
		context.Background(),
		2,
		func() (bool, *sdkErrors.SDKError) {
			attempts++
			return false, nil
		},
	)

	assert.False(t, result)
	assert.NotNil(t, err)
	assert.True(t, err.Is(sdkErrors.ErrRetryOperationFailed))
	assert.Equal(t, 2, attempts)
}

func TestBackOffOptions(t *testing.T) {
	initialInterval := 100 * time.Millisecond
	maxInterval := 1 * time.Second
	maxElapsedTime := 5 * time.Second
	multiplier := 2.5

	retrier := NewExponentialRetrier(
		WithBackOffOptions(
			WithInitialInterval(initialInterval),
			WithMaxInterval(maxInterval),
			WithMaxElapsedTime(maxElapsedTime),
			WithMultiplier(multiplier),
		),
	)

	// Access the backoff configuration
	b := retrier.newBackOff().(*backoff.ExponentialBackOff)

	assert.Equal(t, initialInterval, b.InitialInterval)
	assert.Equal(t, maxInterval, b.MaxInterval)
	assert.Equal(t, maxElapsedTime, b.MaxElapsedTime)
	assert.Equal(t, multiplier, b.Multiplier)
}

func TestDefaultSettings(t *testing.T) {
	retrier := NewExponentialRetrier()
	b := retrier.newBackOff().(*backoff.ExponentialBackOff)

	assert.Equal(t, defaultInitialInterval, b.InitialInterval)
	assert.Equal(t, defaultMaxInterval, b.MaxInterval)
	assert.Equal(t, defaultMaxElapsedTime, b.MaxElapsedTime)
	assert.Equal(t, defaultMultiplier, b.Multiplier)
}
