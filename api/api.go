//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\ Copyright 2024-present SPIKE contributors.
// \\\\\ SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"io"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	acl2 "github.com/spiffe/spike-sdk-go/api/internal/impl/acl"
	cipher2 "github.com/spiffe/spike-sdk-go/api/internal/impl/cipher"
	operator2 "github.com/spiffe/spike-sdk-go/api/internal/impl/operator"
	secret2 "github.com/spiffe/spike-sdk-go/api/internal/impl/secret"

	"github.com/spiffe/spike-sdk-go/api/entity/data"
	"github.com/spiffe/spike-sdk-go/predicate"
	"github.com/spiffe/spike-sdk-go/spiffe"
)

// indirection for testability: allows stubbing cipher calls in unit tests
var (
	cipherEncryptFunc = cipher2.Encrypt
	cipherDecryptFunc = cipher2.Decrypt
)

// API is the SPIKE API.
type API struct {
	source    *workloadapi.X509Source
	predicate predicate.Predicate
}

// New creates and returns a new instance of API configured with a SPIFFE source.
// It automatically discovers and connects to the SPIFFE Workload API endpoint
// using the default socket path and creates an X.509 source for authentication.
// The API client is configured to communicate exclusively with SPIKE Nexus servers.
//
// Returns:
//   - *API: A configured API instance ready for use, or nil if initialization
//     fails
//
// The function will return nil if:
//   - The SPIFFE Workload API is not available
//   - The default endpoint socket cannot be accessed
//   - The X.509 source creation fails
//
// Example usage:
//
//	// Create API client that connects to SPIKE Nexus
//	impl := New()
//	if impl == nil {
//	    log.Fatal("Failed to initialize SPIKE API")
//	}
//	defer impl.Close()
func New() *API {
	defaultEndpointSocket := spiffe.EndpointSocket()

	source, _, err := spiffe.Source(
		context.Background(), defaultEndpointSocket,
	)
	if err != nil {
		return nil
	}

	// API Client can only talk to SPIKE Nexus as a peer.
	return &API{source: source, predicate: predicate.AllowNexus}
}

// NewWithSource initializes a new API instance with a pre-configured
// X509Source. This constructor is useful when you already have an X.509 source
// or need custom source configuration. The API instance will be configured to
// only communicate with SPIKE Nexus servers.
//
// Parameters:
//   - source: A pre-configured X509Source that provides the client's identity
//     certificates and trusted roots for server validation
//
// Returns:
//   - *API: A configured API instance using the provided source
//
// Note: The API client created with this function is restricted to communicate
// only with SPIKE Nexus instances (using predicate.AllowNexus). If you need
// to connect to different servers, use New() with a custom predicate instead.
//
// Example usage:
//
//		// Use with custom-configured source
//		source, err := workloadapi.NewX509Source(ctx,
//	 	workloadapi.WithClientOptions(...))
//		if err != nil {
//		    log.Fatal("Failed to create X509Source")
//		}
//		impl := NewWithSource(source)
//		defer impl.Close()
func NewWithSource(source *workloadapi.X509Source) *API {
	return &API{
		source: source,
		// API Client can only talk to SPIKE Nexus as a peer.
		predicate: predicate.AllowNexus,
	}
}

// Close releases any resources held by the API instance.
// It ensures proper cleanup of the underlying source.
func (a *API) Close() {
	spiffe.CloseSource(a.source)
}

// CreatePolicy creates a new policy in the system. It establishes a mutual
// TLS connection using the X.509 source and sends a policy creation request
// to the server.
//
// The function takes the following parameters:
//   - name string: The name of the policy to be created
//   - spiffeIdPattern string: The SPIFFE ID pattern that this policy will apply
//     to
//   - pathPattern string: The path pattern that this policy will match against
//   - permissions []data.PolicyPermission: A slice of PolicyPermission defining
//     the access rights for this policy
//
// The function returns an error if any of the following operations fail:
//   - Marshaling the policy creation request
//   - Creating the mTLS client
//   - Making the HTTP POST request
//   - Unmarshaling the response
//   - Server-side policy creation (indicated in the response)
//
// Example usage:
//
//	permissions := []data.PolicyPermission{
//	    {
//	        Action: "read",
//	        Resource: "documents/*",
//	    },
//	}
//
//	err = impl.CreatePolicy(
//	    "doc-reader",
//	    "spiffe://example.org/service/*",
//	    "/impl/documents/*",
//	    permissions,
//	)
//	if err != nil {
//	    log.Printf("Failed to create policy: %v", err)
//	    return
//	}
func (a *API) CreatePolicy(
	name string, SPIFFEIDPattern string, pathPattern string,
	permissions []data.PolicyPermission,
) error {
	return acl2.CreatePolicy(a.source,
		name, SPIFFEIDPattern, pathPattern, permissions, a.predicate)
}

// DeletePolicy removes an existing policy from the system using its name.
//
// The function takes the following parameters:
//   - name string: The name of the policy to be deleted
//
// The function returns an error if any of the following operations fail:
//   - Marshaling the policy deletion request
//   - Creating the mTLS client
//   - Making the HTTP POST request
//   - Unmarshaling the response
//   - Server-side policy deletion (indicated in the response)
//
// Example usage:
//
//	err = impl.DeletePolicy("doc-reader")
//	if err != nil {
//	    log.Printf("Failed to delete policy: %v", err)
//	    return
//	}
func (a *API) DeletePolicy(name string) error {
	return acl2.DeletePolicy(a.source, name, a.predicate)
}

// GetPolicy retrieves a policy from the system using its name.
//
// The function takes the following parameters:
//   - name string: The name of the policy to retrieve
//
// The function returns:
//   - (*data.Policy, nil) if the policy is found
//   - (nil, nil) if the policy is not found
//   - (nil, error) if an error occurs during the operation
//
// Errors can occur during:
//   - Marshaling the policy retrieval request
//   - Creating the mTLS client
//   - Making the HTTP POST request (except for not found cases)
//   - Unmarshaling the response
//   - Server-side policy retrieval (indicated in the response)
//
// Example usage:
//
//	policy, err := impl.GetPolicy("doc-reader")
//	if err != nil {
//	    log.Printf("Error retrieving policy: %v", err)
//	    return
//	}
//	if policy == nil {
//	    log.Printf("Policy not found")
//	    return
//	}
//
//	log.Printf("Found policy: %+v", policy)
func (a *API) GetPolicy(name string) (*data.Policy, error) {
	return acl2.GetPolicy(a.source, name, a.predicate)
}

// ListPolicies retrieves policies from the system, optionally filtering by
// SPIFFE ID and path patterns.
//
// The function takes the following parameters:
//   - spiffeIdPattern string: The SPIFFE ID pattern to filter policies.
//     An empty string matches all SPIFFE IDs.
//   - pathPattern string: The path pattern to filter policies.
//     An empty string matches all paths.
//
// The function returns:
//   - (*[]data.Policy, nil) containing all matching policies if successful
//   - (nil, nil) if no policies are found
//   - (nil, error) if an error occurs during the operation
//
// Note: The returned slice pointer should be dereferenced before use:
//
//	policies := *result
//
// Errors can occur during:
//   - Marshaling the policy list request
//   - Creating the mTLS client
//   - Making the HTTP POST request (except for not found cases)
//   - Unmarshaling the response
//   - Server-side policy listing (indicated in the response)
//
// Example usage:
//
//	// List all policies
//	result, err := impl.ListPolicies("", "")
//	if err != nil {
//	    log.Printf("Error listing policies: %v", err)
//	    return
//	}
//	if result == nil {
//	    log.Printf("No policies found")
//	    return
//	}
//
//	policies := *result
//	for _, policy := range policies {
//	    log.Printf("Found policy: %+v", policy)
//	}
func (a *API) ListPolicies(
	SPIFFEIDPattern, pathPattern string,
) (*[]data.Policy, error) {
	return acl2.ListPolicies(a.source, SPIFFEIDPattern, pathPattern, a.predicate)
}

// DeleteSecretVersions deletes specified versions of a secret at the given
// path
//
// It constructs a delete request and sends it to the secrets API endpoint.
//
// Parameters:
//   - path string: Path to the secret to delete
//   - versions []int: Array of version numbers to delete
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or wrapped
//     error on request/parsing failure
//
// Example:
//
//	err := impl.DeleteSecretVersions("secret/path", []int{1, 2})
func (a *API) DeleteSecretVersions(path string, versions []int) error {
	return secret2.Delete(a.source, path, versions, a.predicate)
}

// DeleteSecret deletes the entire secret at the given path
//
// Parameters:
//   - path string: Path to the secret to delete
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or wrapped
//     error on request/parsing failure
//
// Example:
//
//	err := impl.DeleteSecret("secret/path")
func (a *API) DeleteSecret(path string) error {
	return secret2.Delete(a.source, path, []int{}, a.predicate)
}

// GetSecretVersion retrieves a specific version of a secret at the given
// path.
//
// Parameters:
//   - path string: Path to the secret to retrieve
//   - version int: Version number of the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	secret, err := impl.GetSecretVersion("secret/path", 1)
func (a *API) GetSecretVersion(
	path string, version int,
) (*data.Secret, error) {
	return secret2.Get(a.source, path, version, a.predicate)
}

// GetSecret retrieves the latest version of the secret at the given path.
//
// Parameters:
//   - path string: Path to the secret to retrieve
//
// Returns:
//   - *data.Secret: Secret data if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request/parsing failure
//
// Example:
//
//	secret, err := impl.GetSecret("secret/path")
func (a *API) GetSecret(path string) (*data.Secret, error) {
	return secret2.Get(a.source, path, 0, a.predicate)
}

// ListSecretKeys retrieves all secret keys.
//
// Returns:
//   - *[]string: Pointer to an array of secret keys if found, nil if none found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	keys, err := impl.ListSecretKeys()
func (a *API) ListSecretKeys() (*[]string, error) {
	return secret2.ListKeys(a.source, a.predicate)
}

// GetSecretMetadata retrieves metadata for a specific version of a secret at
// the given path.
//
// Parameters:
//   - path string: Path to the secret to retrieve metadata for
//   - version int: Version number of the secret to retrieve metadata for
//
// Returns:
//   - *data.SecretMetadata: Secret metadata if found, nil if secret not found
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	metadata, err := impl.GetSecretMetadata("secret/path", 1)
func (a *API) GetSecretMetadata(
	path string, version int,
) (*data.SecretMetadata, error) {
	return secret2.GetMetadata(a.source, path, version, a.predicate)
}

// PutSecret creates or updates a secret at the specified path with the given
// values.
//
// Parameters:
//   - path string: Path where the secret should be stored
//   - data map[string]string: Map of key-value pairs representing the secret
//     data
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	err := impl.PutSecret("secret/path", map[string]string{"key": "value"})
func (a *API) PutSecret(path string, data map[string]string) error {
	return secret2.Put(a.source, path, data, a.predicate)
}

// UndeleteSecret restores previously deleted versions of a secret at the
// specified path.
//
// Parameters:
//   - path string: Path to the secret to restore
//   - versions []int: Array of version numbers to restore. Empty array
//     attempts no restoration
//
// Returns:
//   - error: nil on success, unauthorized error if not logged in, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	err := impl.UndeleteSecret("secret/path", []int{1, 2})
func (a *API) UndeleteSecret(path string, versions []int) error {
	return secret2.Undelete(a.source, path, versions, a.predicate)
}

// Recover returns recovery partitions for SPIKE Nexus to be used in a
// break-the-glass recovery operation if the SPIKE Nexus auto-recovery mechanism
// isn't successful.
//
// The returned shards are sensitive and should be securely stored out-of-band
// in encrypted form.
//
// Returns:
//   - *[][32]byte: Pointer to an array of recovery shards as 32-byte arrays
//   - error: nil on success, unauthorized error if not authorized, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	shards, err := impl.Recover()
func (a *API) Recover() (map[int]*[32]byte, error) {
	return operator2.Recover(a.source)
}

// Restore SPIKE Nexus backing using recovery shards when SPIKE Keepers cannot
// provide adequate shards and SPIKE Nexus cannot recall its root key either.
//
// This is a break-the-glass superuser-only operation that a well-architected
// SPIKE deployment should not need.
//
// Parameters:
//   - shard *[32]byte: Pointer to a 32-byte array containing the shard to seed
//
// Returns:
//   - *data.RestorationStatus: Status of the restoration process if successful
//   - error: nil on success, unauthorized error if not authorized, or
//     wrapped error on request-parsing failure
//
// Example:
//
//	status, err := impl.Restore(shardPtr)
func (a *API) Restore(
	index int, shard *[32]byte,
) (*data.RestorationStatus, error) {
	return operator2.Restore(a.source, index, shard)
}

// CipherEncryptStream encrypts data from a reader using streaming mode.
// It sends the reader content as the request body with the specified content type.
//
// Parameters:
//   - reader io.Reader: The data source to encrypt
//   - contentType string: The MIME type of the data (e.g., "application/json")
//
// Returns:
//   - []byte: The encrypted data if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//	reader := strings.NewReader("sensitive data")
//	encrypted, err := impl.CipherEncryptStream(reader, "text/plain")
func (a *API) CipherEncryptStream(
	reader io.Reader, contentType string,
) ([]byte, error) {
	return cipherEncryptFunc(
		a.source, cipher2.ModeStream, reader,
		contentType, nil, "", a.predicate,
	)
}

// CipherEncryptJSON encrypts data using JSON mode with structured parameters.
// It sends plaintext and algorithm as JSON and returns encrypted bytes.
//
// Parameters:
//   - plaintext []byte: The data to encrypt
//   - algorithm string: The encryption algorithm to use (e.g., "AES-GCM")
//
// Returns:
//   - []byte: The encrypted data if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//	data := []byte("secret message")
//	encrypted, err := impl.CipherEncryptJSON(data, "AES-GCM")
func (a *API) CipherEncryptJSON(
	plaintext []byte, algorithm string,
) ([]byte, error) {
	return cipherEncryptFunc(
		a.source, cipher2.ModeJSON, nil, "",
		plaintext, algorithm, a.predicate,
	)
}

// CipherDecryptStream decrypts data from a reader using streaming mode.
// It sends the reader content as the request body with the specified
// content type.
//
// Parameters:
//   - reader io.Reader: The encrypted data source to decrypt
//   - contentType string: The MIME type of the data
//     (e.g., "application/octet-stream")
//
// Returns:
//   - []byte: The decrypted plaintext if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//		reader := bytes.NewReader(encryptedData)
//		plaintext, err := impl.CipherDecryptStream(
//	 	reader, "application/octet-stream")
func (a *API) CipherDecryptStream(
	reader io.Reader, contentType string,
) ([]byte, error) {
	return cipherDecryptFunc(
		a.source, cipher2.ModeStream, reader,
		contentType, 0, nil, nil, "", a.predicate,
	)
}

// CipherDecryptJSON decrypts data using JSON mode with structured parameters.
// It sends version, nonce, ciphertext, and algorithm as JSON and returns
// plaintext.
//
// Parameters:
//   - version byte: The cipher version used during encryption
//   - nonce []byte: The nonce bytes used during encryption
//   - ciphertext []byte: The encrypted data to decrypt
//   - algorithm string: The encryption algorithm used (e.g., "AES-GCM")
//
// Returns:
//   - []byte: The decrypted plaintext if successful
//   - error: nil on success, or an error if the operation fails
//
// Example:
//
//	plaintext, err := impl.CipherDecryptJSON(1, nonce, ciphertext, "AES-GCM")
func (a *API) CipherDecryptJSON(
	version byte, nonce, ciphertext []byte, algorithm string,
) ([]byte, error) {
	return cipherDecryptFunc(
		a.source, cipher2.ModeJSON, nil, "",
		version, nonce, ciphertext, algorithm, a.predicate,
	)
}
