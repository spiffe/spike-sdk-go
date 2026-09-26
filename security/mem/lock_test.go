//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

//go:build !windows

package mem

import (
	"testing"

	"github.com/spiffe/spike-sdk-go/config/env"
)

// TestLock_Disabled verifies that an explicit opt-out short-circuits the
// mlockall syscall. This is safe to run anywhere because the syscall is never
// reached.
func TestLock_Disabled(t *testing.T) {
	t.Setenv(env.MemLockDisabled, "true")

	if err := Lock(); err != nil {
		t.Fatalf("Lock() with memory locking disabled = %v, want nil", err)
	}
}
