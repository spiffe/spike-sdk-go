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

const spikeNexusUrlSecrets ApiUrl = "/v1/store/secrets"
const spikeNexusUrlSecretsMetadata ApiUrl = "/v1/store/secrets/metadata"
const spikeNexusUrlInit ApiUrl = "/v1/auth/initialization"

const spikeNexusUrlPolicy ApiUrl = "/v1/acl/policy"

const spikeNexusUrlRecover ApiUrl = "/v1/operator/recover"
const spikeNexusUrlRestore ApiUrl = "/v1/operator/restore"

const spikeKeeperUrlKeep ApiUrl = "/v1/store/keep"
