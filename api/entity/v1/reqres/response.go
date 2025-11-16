package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

// SharPutResponse variants
var (
	ShardPutBadInput     = ShardPutResponse{Err: data.ErrBadInput}
	ShardPutUnauthorized = ShardPutResponse{Err: data.ErrUnauthorized}
	ShardPutSuccess      = ShardPutResponse{Err: data.ErrSuccess}
	ShardPutInternal     = ShardPutResponse{Err: data.ErrInternal}
)

// ShardGetResponse variants
var (
	ShardGetBadInput     = ShardGetResponse{Err: data.ErrBadInput}
	ShardGetUnauthorized = ShardGetResponse{Err: data.ErrUnauthorized}
	ShardGetSuccess      = ShardGetResponse{Err: data.ErrSuccess}
	ShardGetInternal     = ShardGetResponse{Err: data.ErrInternal}
)

// SecretReadResponse variants
var (
	SecretReadBadInput     = SecretReadResponse{Err: data.ErrBadInput}
	SecretReadUnauthorized = SecretReadResponse{Err: data.ErrUnauthorized}
	SecretReadSuccess      = SecretReadResponse{Err: data.ErrSuccess}
	SecretReadInternal     = SecretReadResponse{Err: data.ErrInternal}
	SecretReadNotFound     = SecretReadResponse{Err: data.ErrNotFound}
)

// SecretDeleteResponse variants
var (
	SecretDeleteBadInput     = SecretDeleteResponse{Err: data.ErrBadInput}
	SecretDeleteUnauthorized = SecretDeleteResponse{Err: data.ErrUnauthorized}
	SecretDeleteSuccess      = SecretDeleteResponse{Err: data.ErrSuccess}
	SecretDeleteInternal     = SecretDeleteResponse{Err: data.ErrInternal}
)

// SecretUndeleteResponse variants
var (
	SecretUndeleteBadInput     = SecretUndeleteResponse{Err: data.ErrBadInput}
	SecretUndeleteUnauthorized = SecretUndeleteResponse{Err: data.ErrUnauthorized}
	SecretUndeleteSuccess      = SecretUndeleteResponse{Err: data.ErrSuccess}
	SecretUndeleteInternal     = SecretUndeleteResponse{Err: data.ErrInternal}
)

// SecretListResponse variants
var (
	SecretListBadInput     = SecretListResponse{Err: data.ErrBadInput}
	SecretListUnauthorized = SecretListResponse{Err: data.ErrUnauthorized}
	SecretListInternal     = SecretListResponse{Err: data.ErrInternal}
	SecretListNotFound     = SecretListResponse{Err: data.ErrNotFound}
)

// SecretPutResponse variants
var (
	SecretPutBadInput     = SecretPutResponse{Err: data.ErrBadInput}
	SecretPutUnauthorized = SecretPutResponse{Err: data.ErrUnauthorized}
	SecretPutSuccess      = SecretPutResponse{Err: data.ErrSuccess}
	SecretPutInternal     = SecretPutResponse{Err: data.ErrInternal}
)

// SecretMetadataResponse variants
var (
	SecretMetadataBadInput     = SecretMetadataResponse{Err: data.ErrBadInput}
	SecretMetadataUnauthorized = SecretMetadataResponse{Err: data.ErrUnauthorized}
	SecretMetadataInternal     = SecretMetadataResponse{Err: data.ErrInternal}
	SecretMetadataNotFound     = SecretMetadataResponse{Err: data.ErrNotFound}
)

// PolicyCreateResponse variants
var (
	PolicyCreateBadInput     = PolicyCreateResponse{Err: data.ErrBadInput}
	PolicyCreateUnauthorized = PolicyCreateResponse{Err: data.ErrUnauthorized}
	PolicyCreateSuccess      = PolicyCreateResponse{Err: data.ErrSuccess}
	PolicyCreateInternal     = PolicyCreateResponse{Err: data.ErrInternal}
)

// PolicyReadResponse variants
var (
	PolicyReadBadInput     = PolicyReadResponse{Err: data.ErrBadInput}
	PolicyReadUnauthorized = PolicyReadResponse{Err: data.ErrUnauthorized}
	PolicyReadSuccess      = PolicyReadResponse{Err: data.ErrSuccess}
	PolicyReadInternal     = PolicyReadResponse{Err: data.ErrInternal}
	PolicyReadNotFound     = PolicyReadResponse{Err: data.ErrNotFound}
)

// PolicyDeleteResponse variants
var (
	PolicyDeleteBadInput     = PolicyDeleteResponse{Err: data.ErrBadInput}
	PolicyDeleteUnauthorized = PolicyDeleteResponse{Err: data.ErrUnauthorized}
	PolicyDeleteSuccess      = PolicyDeleteResponse{Err: data.ErrSuccess}
	PolicyDeleteInternal     = PolicyDeleteResponse{Err: data.ErrInternal}
)

// PolicyListResponse variants
var (
	PolicyListBadInput     = PolicyListResponse{Err: data.ErrBadInput}
	PolicyListUnauthorized = PolicyListResponse{Err: data.ErrUnauthorized}
	PolicyListInternal     = PolicyListResponse{Err: data.ErrInternal}
	PolicyListNotFound     = PolicyListResponse{Err: data.ErrNotFound}
)

// CipherDecryptResponse variants
var (
	CipherDecryptBadInput     = CipherDecryptResponse{Err: data.ErrBadInput}
	CipherDecryptUnauthorized = CipherDecryptResponse{Err: data.ErrUnauthorized}
	CipherDecryptInternal     = CipherDecryptResponse{Err: data.ErrInternal}
)

// CipherEncryptResponse variants
var (
	CipherEncryptBadInput     = CipherEncryptResponse{Err: data.ErrBadInput}
	CipherEncryptUnauthorized = CipherEncryptResponse{Err: data.ErrUnauthorized}
	CipherEncryptInternal     = CipherEncryptResponse{Err: data.ErrInternal}
)

// RecoverResponse variants
var (
	RecoverBadInput     = RecoverResponse{Err: data.ErrBadInput}
	RecoverUnauthorized = RecoverResponse{Err: data.ErrUnauthorized}
	RecoverInternal     = RecoverResponse{Err: data.ErrInternal}
)

// RestoreResponse variants
var (
	RestoreBadInput     = RestoreResponse{Err: data.ErrBadInput}
	RestoreUnauthorized = RestoreResponse{Err: data.ErrUnauthorized}
	RestoreInternal     = RestoreResponse{Err: data.ErrInternal}
)
