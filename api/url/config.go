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

const NexusURLSecrets APIURL = "/v1/store/secrets"
const NexusURLSecretsMetadata APIURL = "/v1/store/secrets/metadata"
const NexusURLInit APIURL = "/v1/auth/initialization"

const NexusURLPolicy APIURL = "/v1/acl/policy"

const NexusURLOperatorRecover APIURL = "/v1/operator/recover"
const NexusURLOperatorRestore APIURL = "/v1/operator/restore"

const KeeperURLKeep APIURL = "/v1/store/keep"
const KeeperURLContribute APIURL = "/v1/store/contribute"
const KeeperURLShard APIURL = "/v1/store/shard"
