package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

// SharPutResponse variants
var (
	ShardPutBadInput     = ShardPutResponse{Err: data.ErrBadInput}
	ShardPutUnauthorized = ShardPutResponse{Err: data.ErrUnauthorized}
	ShardPutSuccess      = ShardPutResponse{Err: data.ErrSuccess}
)

// ShardGetResponse variants
var (
	ShardGetBadInput     = ShardGetResponse{Err: data.ErrBadInput}
	ShardGetUnauthorized = ShardGetResponse{Err: data.ErrUnauthorized}
	ShardGetSuccess      = ShardGetResponse{Err: data.ErrSuccess}
)

// SecretReadResponse variants
var (
	SecretReadBadInput     = SecretReadResponse{Err: data.ErrBadInput}
	SecretReadUnauthorized = SecretReadResponse{Err: data.ErrUnauthorized}
	SecretReadSuccess      = SecretReadResponse{Err: data.ErrSuccess}
)

// SecretDeleteResponse variants
var (
	SecretDeleteBadInput     = SecretDeleteResponse{Err: data.ErrBadInput}
	SecretDeleteUnauthorized = SecretDeleteResponse{Err: data.ErrUnauthorized}
	SecretDeleteSuccess      = SecretDeleteResponse{Err: data.ErrSuccess}
)

// SecretUndeleteResponse variants
var (
	SecretUndeleteBadInput     = SecretUndeleteResponse{Err: data.ErrBadInput}
	SecretUndeleteUnauthorized = SecretUndeleteResponse{Err: data.ErrUnauthorized}
	SecretUndeleteSuccess      = SecretUndeleteResponse{Err: data.ErrSuccess}
)

// SecretListResponse variants
var (
	SecretListBadInput     = SecretListResponse{Err: data.ErrBadInput}
	SecretListUnauthorized = SecretListResponse{Err: data.ErrUnauthorized}
)

// SecretPutResponse variants
var (
	SecretPutBadInput     = SecretPutResponse{Err: data.ErrBadInput}
	SecretPutUnauthorized = SecretPutResponse{Err: data.ErrUnauthorized}
	SecretPutSuccess      = SecretPutResponse{Err: data.ErrSuccess}
)

// SecretMetadataResponse variants
var (
	SecretMetadataBadInput     = SecretMetadataResponse{Err: data.ErrBadInput}
	SecretMetadataUnauthorized = SecretMetadataResponse{Err: data.ErrUnauthorized}
)

// PolicyCreateResponse variants
var (
	PolicyCreateBadInput     = PolicyCreateResponse{Err: data.ErrBadInput}
	PolicyCreateUnauthorized = PolicyCreateResponse{Err: data.ErrUnauthorized}
	PolicyCreateSuccess      = PolicyCreateResponse{Err: data.ErrSuccess}
)

// PolicyReadResponse variants
var (
	PolicyReadBadInput     = PolicyReadResponse{Err: data.ErrBadInput}
	PolicyReadUnauthorized = PolicyReadResponse{Err: data.ErrUnauthorized}
	PolicyReadSuccess      = PolicyReadResponse{Err: data.ErrSuccess}
)

// PolicyDeleteResponse variants
var (
	PolicyDeleteBadInput     = PolicyDeleteResponse{Err: data.ErrBadInput}
	PolicyDeleteUnauthorized = PolicyDeleteResponse{Err: data.ErrUnauthorized}
	PolicyDeleteSuccess      = PolicyDeleteResponse{Err: data.ErrSuccess}
)

// PolicyListResponse variants
var (
	PolicyListBadInput     = PolicyListResponse{Err: data.ErrBadInput}
	PolicyListUnauthorized = PolicyListResponse{Err: data.ErrUnauthorized}
)

// CipherDecryptResponse variants
var (
	CipherDecryptBadInput     = CipherDecryptResponse{Err: data.ErrBadInput}
	CipherDecryptUnauthorized = CipherDecryptResponse{Err: data.ErrUnauthorized}
)

// CipherEncryptResponse variants
var (
	CipherEncryptBadInput     = CipherEncryptResponse{Err: data.ErrBadInput}
	CipherEncryptUnauthorized = CipherEncryptResponse{Err: data.ErrUnauthorized}
)

// RecoverResponse variants
var (
	RecoverBadInput     = RecoverResponse{Err: data.ErrBadInput}
	RecoverUnauthorized = RecoverResponse{Err: data.ErrUnauthorized}
)

// RestoreResponse variants
var (
	RestoreBadInput     = RestoreResponse{Err: data.ErrBadInput}
	RestoreUnauthorized = RestoreResponse{Err: data.ErrUnauthorized}
)
