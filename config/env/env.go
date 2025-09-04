//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

const KeeperTLSPort = "SPIKE_KEEPER_TLS_PORT"
const NexusKeeperPeers = "SPIKE_NEXUS_KEEPER_PEERS"
const NexusTLSPort = "SPIKE_NEXUS_TLS_PORT"
const NexusMaxSecretVersions = "SPIKE_NEXUS_MAX_SECRET_VERSIONS"
const NexusBackendStore = "SPIKE_NEXUS_BACKEND_STORE"
const NexusDBOperationTimeout = "SPIKE_NEXUS_DB_OPERATION_TIMEOUT"
const NexusDBJournalMode = "SPIKE_NEXUS_DB_JOURNAL_MODE"
const NexusDBBusyTimeoutMS = "SPIKE_NEXUS_DB_BUSY_TIMEOUT_MS"
const NexusDBMaxOpenConns = "SPIKE_NEXUS_DB_MAX_OPEN_CONNS"
const NexusDBMaxIdleConns = "SPIKE_NEXUS_DB_MAX_IDLE_CONNS"
const NexusDBConnMaxLifetime = "SPIKE_NEXUS_DB_CONN_MAX_LIFETIME"
const NexusDBInitializationTimeout = "SPIKE_NEXUS_DB_INITIALIZATION_TIMEOUT"
const NexusDBSkipSchemaCreation = "SPIKE_NEXUS_DB_SKIP_SCHEMA_CREATION"
const NexusPBKDF2IterationCount = "SPIKE_NEXUS_PBKDF2_ITERATION_COUNT"
const NexusRecoveryMaxInterval = "SPIKE_NEXUS_RECOVERY_MAX_INTERVAL"
const NexusShamirShares = "SPIKE_NEXUS_SHAMIR_SHARES"
const NexusShamirThreshold = "SPIKE_NEXUS_SHAMIR_THRESHOLD"
const NexusKeeperUpdateInterval = "SPIKE_NEXUS_KEEPER_UPDATE_INTERVAL"
const PilotShowMemoryWarning = "SPIKE_PILOT_SHOW_MEMORY_WARNING"
const SystemLogLevel = "SPIKE_SYSTEM_LOG_LEVEL"
const NexusAPIURL = "SPIKE_NEXUS_API_URL"
const TrustRoot = "SPIKE_TRUST_ROOT"
const TrustRootKeeper = "SPIKE_TRUST_ROOT_KEEPER"
const TrustRootPilot = "SPIKE_TRUST_ROOT_PILOT"
const TrustRootNexus = "SPIKE_TRUST_ROOT_NEXUS"
const TrustRootLiteWorkload = "SPIKE_TRUST_ROOT_LITE_WORKLOAD"
const BannerEnabled = "SPIKE_BANNER_ENABLED"
const SPIFFEEndpointSocket = "SPIFFE_ENDPOINT_SOCKET"
