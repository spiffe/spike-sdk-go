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

const SpikeNexusURLSecrets APIURL = "/v1/store/secrets"
const SpikeNexusURLSecretsMetadata APIURL = "/v1/store/secrets/metadata"
const SpikeNexusURLInit APIURL = "/v1/auth/initialization"

const SpikeNexusURLPolicy APIURL = "/v1/acl/policy"

const SpikeNexusURLOperatorRecover APIURL = "/v1/operator/recover"
const SpikeNexusURLOperatorRestore APIURL = "/v1/operator/restore"

const SpikeKeeperURLKeep APIURL = "/v1/store/keep"
const SpikeKeeperURLContribute APIURL = "/v1/store/contribute"
const SpikeKeeperURLShard APIURL = "/v1/store/shard"
