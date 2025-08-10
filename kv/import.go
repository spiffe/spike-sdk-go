//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package kv

// ImportSecrets hydrates the key-value store with secrets loaded from
// persistent storage or a similar medium. It takes a map of path to secret
// values and adds them to the in-memory store. This is typically used during
// initialization or recovery after a system crash.
//
// If a secret already exists in the store, it will be overwritten with the
// imported value. The method preserves all version history and metadata from
// the imported secrets.
//
// Example usage:
//
//	secrets, err := persistentStore.LoadAllSecrets(context.Background())
//	if err != nil {
//	    log.Fatalf("Failed to load secrets: %v", err)
//	}
//	kvStore.ImportSecrets(secrets)
func (kv *KV) ImportSecrets(secrets map[string]*Value) {
	for path, secret := range secrets {
		// Create a deep copy of the secret to avoid sharing memory
		newSecret := &Value{
			Versions: make(map[int]Version, len(secret.Versions)),
			Metadata: Metadata{
				CreatedTime:    secret.Metadata.CreatedTime,
				UpdatedTime:    secret.Metadata.UpdatedTime,
				MaxVersions:    kv.maxSecretVersions, // Use the KV store's setting
				CurrentVersion: secret.Metadata.CurrentVersion,
				OldestVersion:  secret.Metadata.OldestVersion,
			},
		}

		// Copy all versions
		for versionNum, version := range secret.Versions {
			// Deep copy the data map
			dataCopy := make(map[string]string, len(version.Data))
			for k, v := range version.Data {
				dataCopy[k] = v
			}

			// Create the version copy
			versionCopy := Version{
				Data:        dataCopy,
				CreatedTime: version.CreatedTime,
				Version:     versionNum,
			}

			// Copy deleted time if set
			if version.DeletedTime != nil {
				deletedTime := *version.DeletedTime
				versionCopy.DeletedTime = &deletedTime
			}

			newSecret.Versions[versionNum] = versionCopy
		}

		// Store the copied secret
		kv.data[path] = newSecret
	}
}
