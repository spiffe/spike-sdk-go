//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package auth

// These paths are kv-store keys.
// They are NOT Unix file paths.
// They should NOT start with a trailing slash.

const PathSystemPolicyAccess = "spike/system/acl"
const PathSystemCipherExecute = "spike/system/cipher/exec"
const PathSystemSecretAccess = "spike/system/secret"
