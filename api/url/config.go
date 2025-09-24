//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

type APIAction string

const KeyAPIAction = "action"

const ActionCheck APIAction = "check"
const ActionGet APIAction = "get"
const ActionDelete APIAction = "delete"
const ActionUndelete APIAction = "undelete"
const ActionList APIAction = "list"
const ActionDefault APIAction = ""
const ActionRead APIAction = "read"

type APIURL string

const NexusSecrets APIURL = "/v1/store/secrets"
const NexusSecretsMetadata APIURL = "/v1/store/secrets/metadata"
const NexusInit APIURL = "/v1/auth/initialization"

const NexusPolicy APIURL = "/v1/acl/policy"

const NexusOperatorRecover APIURL = "/v1/operator/recover"
const NexusOperatorRestore APIURL = "/v1/operator/restore"

const NexusCipherEncrypt APIURL = "/v1/cipher/encrypt"
const NexusCipherDecrypt APIURL = "/v1/cipher/decrypt"

const KeeperKeep APIURL = "/v1/store/keep"
const KeeperContribute APIURL = "/v1/store/contribute"
const KeeperShard APIURL = "/v1/store/shard"

const NexusHealth APIURL = "/v1/health/status"
