//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "testing"

func TestMemLockDisabledVal(t *testing.T) {
	tests := []struct {
		name string
		set  bool
		val  string
		want bool
	}{
		{name: "unset defaults to locking enabled", set: false, want: false},
		{name: "true disables locking", set: true, val: "true", want: true},
		{name: "TRUE is case-insensitive", set: true, val: "TRUE", want: true},
		{name: "whitespace is trimmed", set: true, val: "  true  ", want: true},
		{name: "false keeps locking enabled", set: true, val: "false", want: false},
		{name: "unexpected value keeps locking enabled", set: true, val: "1", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(MemLockDisabled, tt.val)
			}
			if got := MemLockDisabledVal(); got != tt.want {
				t.Fatalf("MemLockDisabledVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
