//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package url

type SpikeNexusApiAction string

const keyApiAction = "action"

const actionNexusCheck SpikeNexusApiAction = "check"
const actionNexusGet SpikeNexusApiAction = "get"
const actionNexusDelete SpikeNexusApiAction = "delete"
const actionNexusUndelete SpikeNexusApiAction = "undelete"
const actionNexusList SpikeNexusApiAction = "list"
const actionNexusDefault SpikeNexusApiAction = ""

type SpikeKeeperApiAction string

const actionKeeperRead SpikeKeeperApiAction = "read"
const actionKeeperDefault SpikeKeeperApiAction = ""

type ApiUrl string

const SpikeNexusUrlSecrets ApiUrl = "/v1/store/secrets"
const SpikeNexusUrlSecretsMetadata ApiUrl = "/v1/store/secrets/metadata"
const SpikeNexusUrlInit ApiUrl = "/v1/auth/initialization"

const SpikeNexusUrlPolicy ApiUrl = "/v1/acl/policy"

const SpikeNexusUrlRecover ApiUrl = "/v1/operator/recover"
const SpikeNexusUrlRestore ApiUrl = "/v1/operator/restore"

const SpikeKeeperUrlKeep ApiUrl = "/v1/store/keep"

const SpikeNexusUrlOperatorRestore = "/v1/operator/restore"
const SpikeNexusUrlOperatorRecover = "/v1/operator/recover"

const SpikeKeeperUrlContribute ApiUrl = "/v1/store/contribute"
const SpikeKeeperUrlShard ApiUrl = "/v1/store/shard"
