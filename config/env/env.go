//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

// Sort alphabetically.

const BannerEnabled = "SPIKE_BANNER_ENABLED"
const BootstrapConfigMapName = "SPIKE_BOOTSTRAP_CONFIGMAP_NAME"
const BootstrapForce = "SPIKE_BOOTSTRAP_FORCE"
const KeeperTLSPort = "SPIKE_KEEPER_TLS_PORT"
const NexusAPIURL = "SPIKE_NEXUS_API_URL"
const NexusBackendStore = "SPIKE_NEXUS_BACKEND_STORE"
const NexusDBBusyTimeoutMS = "SPIKE_NEXUS_DB_BUSY_TIMEOUT_MS"
const NexusDBConnMaxLifetime = "SPIKE_NEXUS_DB_CONN_MAX_LIFETIME"
const NexusDBInitializationTimeout = "SPIKE_NEXUS_DB_INITIALIZATION_TIMEOUT"
const NexusDBJournalMode = "SPIKE_NEXUS_DB_JOURNAL_MODE"
const NexusDBMaxIdleConns = "SPIKE_NEXUS_DB_MAX_IDLE_CONNS"
const NexusDBMaxOpenConns = "SPIKE_NEXUS_DB_MAX_OPEN_CONNS"
const NexusDBOperationTimeout = "SPIKE_NEXUS_DB_OPERATION_TIMEOUT"
const NexusDBSkipSchemaCreation = "SPIKE_NEXUS_DB_SKIP_SCHEMA_CREATION"
const NexusDataDir = "SPIKE_NEXUS_DATA_DIR"
const NexusKeeperPeers = "SPIKE_NEXUS_KEEPER_PEERS"
const NexusKeeperUpdateInterval = "SPIKE_NEXUS_KEEPER_UPDATE_INTERVAL"
const NexusMaxEntryVersions = "SPIKE_NEXUS_MAX_SECRET_VERSIONS"
const NexusPBKDF2IterationCount = "SPIKE_NEXUS_PBKDF2_ITERATION_COUNT"
const NexusRecoveryMaxInterval = "SPIKE_NEXUS_RECOVERY_MAX_INTERVAL"
const NexusShamirShares = "SPIKE_NEXUS_SHAMIR_SHARES"
const NexusShamirThreshold = "SPIKE_NEXUS_SHAMIR_THRESHOLD"
const NexusTLSPort = "SPIKE_NEXUS_TLS_PORT"
const PilotRecoveryDir = "SPIKE_PILOT_RECOVERY_DIR"
const PilotShowMemoryWarning = "SPIKE_PILOT_SHOW_MEMORY_WARNING"
const SPIFFEEndpointSocket = "SPIFFE_ENDPOINT_SOCKET"
const SystemLogLevel = "SPIKE_SYSTEM_LOG_LEVEL"
const TrustRoot = "SPIKE_TRUST_ROOT"
const TrustRootBootstrap = "SPIKE_TRUST_ROOT_BOOTSTRAP"
const TrustRootKeeper = "SPIKE_TRUST_ROOT_KEEPER"
const TrustRootLiteWorkload = "SPIKE_TRUST_ROOT_LITE_WORKLOAD"
const TrustRootNexus = "SPIKE_TRUST_ROOT_NEXUS"
const TrustRootPilot = "SPIKE_TRUST_ROOT_PILOT"
