package reqres

import (
	sdkErrors "github.com/spiffe/spike-sdk-go/errors"
)

// SharPutResponse variants
var (
	ShardPutBadInput     = ShardPutResponse{Err: sdkErrors.ErrCodeBadInput}
	ShardPutUnauthorized = ShardPutResponse{Err: sdkErrors.ErrCodeUnauthorized}
	ShardPutSuccess      = ShardPutResponse{Err: sdkErrors.ErrCodeSuccess}
	ShardPutInternal     = ShardPutResponse{Err: sdkErrors.ErrCodeInternal}
)

// ShardGetResponse variants
var (
	ShardGetBadInput     = ShardGetResponse{Err: sdkErrors.ErrCodeBadInput}
	ShardGetUnauthorized = ShardGetResponse{Err: sdkErrors.ErrCodeUnauthorized}
	ShardGetSuccess      = ShardGetResponse{Err: sdkErrors.ErrCodeSuccess}
	ShardGetInternal     = ShardGetResponse{Err: sdkErrors.ErrCodeInternal}
)

// SecretReadResponse variants
var (
	SecretReadBadInput     = SecretReadResponse{Err: sdkErrors.ErrCodeBadInput}
	SecretReadUnauthorized = SecretReadResponse{Err: sdkErrors.ErrCodeUnauthorized}
	SecretReadSuccess      = SecretReadResponse{Err: sdkErrors.ErrCodeSuccess}
	SecretReadInternal     = SecretReadResponse{Err: sdkErrors.ErrCodeInternal}
	SecretReadNotFound     = SecretReadResponse{Err: sdkErrors.ErrCodeNotFound}
)

// SecretDeleteResponse variants
var (
	SecretDeleteBadInput     = SecretDeleteResponse{Err: sdkErrors.ErrCodeBadInput}
	SecretDeleteUnauthorized = SecretDeleteResponse{Err: sdkErrors.ErrCodeUnauthorized}
	SecretDeleteSuccess      = SecretDeleteResponse{Err: sdkErrors.ErrCodeSuccess}
	SecretDeleteInternal     = SecretDeleteResponse{Err: sdkErrors.ErrCodeInternal}
)

// SecretUndeleteResponse variants
var (
	SecretUndeleteBadInput     = SecretUndeleteResponse{Err: sdkErrors.ErrCodeBadInput}
	SecretUndeleteUnauthorized = SecretUndeleteResponse{Err: sdkErrors.ErrCodeUnauthorized}
	SecretUndeleteSuccess      = SecretUndeleteResponse{Err: sdkErrors.ErrCodeSuccess}
	SecretUndeleteInternal     = SecretUndeleteResponse{Err: sdkErrors.ErrCodeInternal}
)

// SecretListResponse variants
var (
	SecretListBadInput     = SecretListResponse{Err: sdkErrors.ErrCodeBadInput}
	SecretListUnauthorized = SecretListResponse{Err: sdkErrors.ErrCodeUnauthorized}
	SecretListInternal     = SecretListResponse{Err: sdkErrors.ErrCodeInternal}
	SecretListNotFound     = SecretListResponse{Err: sdkErrors.ErrCodeNotFound}
)

// SecretPutResponse variants
var (
	SecretPutBadInput     = SecretPutResponse{Err: sdkErrors.ErrCodeBadInput}
	SecretPutUnauthorized = SecretPutResponse{Err: sdkErrors.ErrCodeUnauthorized}
	SecretPutSuccess      = SecretPutResponse{Err: sdkErrors.ErrCodeSuccess}
	SecretPutInternal     = SecretPutResponse{Err: sdkErrors.ErrCodeInternal}
)

// SecretMetadataResponse variants
var (
	SecretMetadataBadInput     = SecretMetadataResponse{Err: sdkErrors.ErrCodeBadInput}
	SecretMetadataUnauthorized = SecretMetadataResponse{Err: sdkErrors.ErrCodeUnauthorized}
	SecretMetadataInternal     = SecretMetadataResponse{Err: sdkErrors.ErrCodeInternal}
	SecretMetadataNotFound     = SecretMetadataResponse{Err: sdkErrors.ErrCodeNotFound}
)

// PolicyCreateResponse variants
var (
	PolicyCreateBadInput     = PolicyCreateResponse{Err: sdkErrors.ErrCodeBadInput}
	PolicyCreateUnauthorized = PolicyCreateResponse{Err: sdkErrors.ErrCodeUnauthorized}
	PolicyCreateSuccess      = PolicyCreateResponse{Err: sdkErrors.ErrCodeSuccess}
	PolicyCreateInternal     = PolicyCreateResponse{Err: sdkErrors.ErrCodeInternal}
)

// PolicyReadResponse variants
var (
	PolicyReadBadInput     = PolicyReadResponse{Err: sdkErrors.ErrCodeBadInput}
	PolicyReadUnauthorized = PolicyReadResponse{Err: sdkErrors.ErrCodeUnauthorized}
	PolicyReadSuccess      = PolicyReadResponse{Err: sdkErrors.ErrCodeSuccess}
	PolicyReadInternal     = PolicyReadResponse{Err: sdkErrors.ErrCodeInternal}
	PolicyReadNotFound     = PolicyReadResponse{Err: sdkErrors.ErrCodeNotFound}
)

// PolicyDeleteResponse variants
var (
	PolicyDeleteBadInput     = PolicyDeleteResponse{Err: sdkErrors.ErrCodeBadInput}
	PolicyDeleteUnauthorized = PolicyDeleteResponse{Err: sdkErrors.ErrCodeUnauthorized}
	PolicyDeleteSuccess      = PolicyDeleteResponse{Err: sdkErrors.ErrCodeSuccess}
	PolicyDeleteInternal     = PolicyDeleteResponse{Err: sdkErrors.ErrCodeInternal}
)

// PolicyListResponse variants
var (
	PolicyListBadInput     = PolicyListResponse{Err: sdkErrors.ErrCodeBadInput}
	PolicyListUnauthorized = PolicyListResponse{Err: sdkErrors.ErrCodeUnauthorized}
	PolicyListInternal     = PolicyListResponse{Err: sdkErrors.ErrCodeInternal}
	PolicyListNotFound     = PolicyListResponse{Err: sdkErrors.ErrCodeNotFound}
)

// CipherDecryptResponse variants
var (
	CipherDecryptBadInput     = CipherDecryptResponse{Err: sdkErrors.ErrCodeBadInput}
	CipherDecryptUnauthorized = CipherDecryptResponse{Err: sdkErrors.ErrCodeUnauthorized}
	CipherDecryptInternal     = CipherDecryptResponse{Err: sdkErrors.ErrCodeInternal}
)

// CipherEncryptResponse variants
var (
	CipherEncryptBadInput     = CipherEncryptResponse{Err: sdkErrors.ErrCodeBadInput}
	CipherEncryptUnauthorized = CipherEncryptResponse{Err: sdkErrors.ErrCodeUnauthorized}
	CipherEncryptInternal     = CipherEncryptResponse{Err: sdkErrors.ErrCodeInternal}
)

// RecoverResponse variants
var (
	RecoverBadInput     = RecoverResponse{Err: sdkErrors.ErrCodeBadInput}
	RecoverUnauthorized = RecoverResponse{Err: sdkErrors.ErrCodeUnauthorized}
	RecoverInternal     = RecoverResponse{Err: sdkErrors.ErrCodeInternal}
)

// RestoreResponse variants
var (
	RestoreBadInput     = RestoreResponse{Err: sdkErrors.ErrCodeBadInput}
	RestoreUnauthorized = RestoreResponse{Err: sdkErrors.ErrCodeUnauthorized}
	RestoreInternal     = RestoreResponse{Err: sdkErrors.ErrCodeInternal}
)

// BoostrapResponse variants
var (
	BootstrapBadInput     = BootstrapVerifyResponse{Err: sdkErrors.ErrCodeBadInput}
	BootstrapUnauthorized = BootstrapVerifyResponse{Err: sdkErrors.ErrCodeUnauthorized}
	BootstrapInternal     = BootstrapVerifyResponse{Err: sdkErrors.ErrCodeInternal}
)
