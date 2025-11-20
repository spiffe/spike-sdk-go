package reqres

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// SharPutResponse variants
var (
	ShardPutBadInput     = ShardPutResponse{Err: sdkErrors.ErrBadRequest.Code}
	ShardPutUnauthorized = ShardPutResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	ShardPutSuccess      = ShardPutResponse{Err: sdkErrors.ErrSuccess.Code}
	ShardPutInternal     = ShardPutResponse{Err: sdkErrors.ErrInternal.Code}
)

// ShardGetResponse variants
var (
	ShardGetBadInput     = ShardGetResponse{Err: sdkErrors.ErrBadRequest.Code}
	ShardGetUnauthorized = ShardGetResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	ShardGetSuccess      = ShardGetResponse{Err: sdkErrors.ErrSuccess.Code}
	ShardGetInternal     = ShardGetResponse{Err: sdkErrors.ErrInternal.Code}
)

// SecretReadResponse variants
var (
	SecretReadBadInput     = SecretReadResponse{Err: sdkErrors.ErrBadRequest.Code}
	SecretReadUnauthorized = SecretReadResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	SecretReadSuccess      = SecretReadResponse{Err: sdkErrors.ErrSuccess.Code}
	SecretReadInternal     = SecretReadResponse{Err: sdkErrors.ErrInternal.Code}
	SecretReadNotFound     = SecretReadResponse{Err: sdkErrors.ErrNotFound.Code}
)

// SecretDeleteResponse variants
var (
	SecretDeleteBadInput     = SecretDeleteResponse{Err: sdkErrors.ErrBadRequest.Code}
	SecretDeleteUnauthorized = SecretDeleteResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	SecretDeleteSuccess      = SecretDeleteResponse{Err: sdkErrors.ErrSuccess.Code}
	SecretDeleteInternal     = SecretDeleteResponse{Err: sdkErrors.ErrInternal.Code}
)

// SecretUndeleteResponse variants
var (
	SecretUndeleteBadInput     = SecretUndeleteResponse{Err: sdkErrors.ErrBadRequest.Code}
	SecretUndeleteUnauthorized = SecretUndeleteResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	SecretUndeleteSuccess      = SecretUndeleteResponse{Err: sdkErrors.ErrSuccess.Code}
	SecretUndeleteInternal     = SecretUndeleteResponse{Err: sdkErrors.ErrInternal.Code}
)

// SecretListResponse variants
var (
	SecretListBadInput     = SecretListResponse{Err: sdkErrors.ErrBadRequest.Code}
	SecretListUnauthorized = SecretListResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	SecretListInternal     = SecretListResponse{Err: sdkErrors.ErrInternal.Code}
	SecretListNotFound     = SecretListResponse{Err: sdkErrors.ErrNotFound.Code}
)

// SecretPutResponse variants
var (
	SecretPutBadInput     = SecretPutResponse{Err: sdkErrors.ErrBadRequest.Code}
	SecretPutUnauthorized = SecretPutResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	SecretPutSuccess      = SecretPutResponse{Err: sdkErrors.ErrSuccess.Code}
	SecretPutInternal     = SecretPutResponse{Err: sdkErrors.ErrInternal.Code}
)

// SecretMetadataResponse variants
var (
	SecretMetadataBadInput     = SecretMetadataResponse{Err: sdkErrors.ErrBadRequest.Code}
	SecretMetadataUnauthorized = SecretMetadataResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	SecretMetadataInternal     = SecretMetadataResponse{Err: sdkErrors.ErrInternal.Code}
	SecretMetadataNotFound     = SecretMetadataResponse{Err: sdkErrors.ErrNotFound.Code}
)

// PolicyCreateResponse variants
var (
	PolicyCreateBadInput     = PolicyCreateResponse{Err: sdkErrors.ErrBadRequest.Code}
	PolicyCreateUnauthorized = PolicyCreateResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	PolicyCreateSuccess      = PolicyCreateResponse{Err: sdkErrors.ErrSuccess.Code}
	PolicyCreateInternal     = PolicyCreateResponse{Err: sdkErrors.ErrInternal.Code}
)

// PolicyReadResponse variants
var (
	PolicyReadBadInput     = PolicyReadResponse{Err: sdkErrors.ErrBadRequest.Code}
	PolicyReadUnauthorized = PolicyReadResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	PolicyReadSuccess      = PolicyReadResponse{Err: sdkErrors.ErrSuccess.Code}
	PolicyReadInternal     = PolicyReadResponse{Err: sdkErrors.ErrInternal.Code}
	PolicyReadNotFound     = PolicyReadResponse{Err: sdkErrors.ErrNotFound.Code}
)

// PolicyDeleteResponse variants
var (
	PolicyDeleteBadInput     = PolicyDeleteResponse{Err: sdkErrors.ErrBadRequest.Code}
	PolicyDeleteUnauthorized = PolicyDeleteResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	PolicyDeleteSuccess      = PolicyDeleteResponse{Err: sdkErrors.ErrSuccess.Code}
	PolicyDeleteInternal     = PolicyDeleteResponse{Err: sdkErrors.ErrInternal.Code}
)

// PolicyListResponse variants
var (
	PolicyListBadInput     = PolicyListResponse{Err: sdkErrors.ErrBadRequest.Code}
	PolicyListUnauthorized = PolicyListResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	PolicyListInternal     = PolicyListResponse{Err: sdkErrors.ErrInternal.Code}
	PolicyListNotFound     = PolicyListResponse{Err: sdkErrors.ErrNotFound.Code}
)

// CipherDecryptResponse variants
var (
	CipherDecryptBadInput     = CipherDecryptResponse{Err: sdkErrors.ErrBadRequest.Code}
	CipherDecryptUnauthorized = CipherDecryptResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	CipherDecryptInternal     = CipherDecryptResponse{Err: sdkErrors.ErrInternal.Code}
)

// CipherEncryptResponse variants
var (
	CipherEncryptBadInput     = CipherEncryptResponse{Err: sdkErrors.ErrBadRequest.Code}
	CipherEncryptUnauthorized = CipherEncryptResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	CipherEncryptInternal     = CipherEncryptResponse{Err: sdkErrors.ErrInternal.Code}
)

// RecoverResponse variants
var (
	RecoverBadInput     = RecoverResponse{Err: sdkErrors.ErrBadRequest.Code}
	RecoverUnauthorized = RecoverResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	RecoverInternal     = RecoverResponse{Err: sdkErrors.ErrInternal.Code}
)

// RestoreResponse variants
var (
	RestoreBadInput     = RestoreResponse{Err: sdkErrors.ErrBadRequest.Code}
	RestoreUnauthorized = RestoreResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	RestoreInternal     = RestoreResponse{Err: sdkErrors.ErrInternal.Code}
)

// BoostrapResponse variants
var (
	BootstrapBadInput     = BootstrapVerifyResponse{Err: sdkErrors.ErrBadRequest.Code}
	BootstrapUnauthorized = BootstrapVerifyResponse{Err: sdkErrors.ErrAccessUnauthorized.Code}
	BootstrapInternal     = BootstrapVerifyResponse{Err: sdkErrors.ErrInternal.Code}
)
