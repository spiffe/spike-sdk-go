//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

type ApiAction string

const KeyApiAction = "action"

const ActionCheck ApiAction = "check"
const ActionGet ApiAction = "get"
const ActionDelete ApiAction = "delete"
const ActionUndelete ApiAction = "undelete"
const ActionList ApiAction = "list"
const ActionDefault ApiAction = ""
const ActionRead ApiAction = "read"

type ApiUrl string

const SpikeNexusUrlSecrets ApiUrl = "/v1/store/secrets"
const SpikeNexusUrlSecretsMetadata ApiUrl = "/v1/store/secrets/metadata"
const SpikeNexusUrlInit ApiUrl = "/v1/auth/initialization"

const SpikeNexusUrlPolicy ApiUrl = "/v1/acl/policy"

const SpikeNexusUrlOperatorRecover ApiUrl = "/v1/operator/recover"
const SpikeNexusUrlOperatorRestore ApiUrl = "/v1/operator/restore"

const SpikeKeeperUrlKeep ApiUrl = "/v1/store/keep"
const SpikeKeeperUrlContribute ApiUrl = "/v1/store/contribute"
const SpikeKeeperUrlShard ApiUrl = "/v1/store/shard"
