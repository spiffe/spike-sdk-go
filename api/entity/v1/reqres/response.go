package reqres

// SharPutResponse variants
var (
	ShardPutBadRequest   = ShardPutResponse{}.BadRequest()
	ShardPutUnauthorized = ShardPutResponse{}.Unauthorized()
	ShardPutSuccess      = ShardPutResponse{}.Success()
	ShardPutInternal     = ShardPutResponse{}.Internal()
)

// ShardGetResponse variants
var (
	ShardGetNotFound     = ShardGetResponse{}.NotFound()
	ShardGetBadRequest   = ShardGetResponse{}.BadRequest()
	ShardGetUnauthorized = ShardGetResponse{}.Unauthorized()
	ShardGetSuccess      = ShardGetResponse{}.Success()
	ShardGetInternal     = ShardGetResponse{}.Internal()
)

// SecretGetResponse variants
var (
	SecretReadNotFound     = SecretGetResponse{}.NotFound()
	SecretReadBadRequest   = SecretGetResponse{}.BadRequest()
	SecretReadUnauthorized = SecretGetResponse{}.Unauthorized()
	SecretReadSuccess      = SecretGetResponse{}.Success()
	SecretReadInternal     = SecretGetResponse{}.Internal()
)

// SecretDeleteResponse variants
var (
	SecretDeleteNotFound     = SecretDeleteResponse{}.NotFound()
	SecretDeleteBadRequest   = SecretDeleteResponse{}.BadRequest()
	SecretDeleteUnauthorized = SecretDeleteResponse{}.Unauthorized()
	SecretDeleteSuccess      = SecretDeleteResponse{}.Success()
	SecretDeleteInternal     = SecretDeleteResponse{}.Internal()
)

// SecretUndeleteResponse variants
var (
	SecretUndeleteNotFound     = SecretUndeleteResponse{}.NotFound()
	SecretUndeleteBadRequest   = SecretUndeleteResponse{}.BadRequest()
	SecretUndeleteUnauthorized = SecretUndeleteResponse{}.Unauthorized()
	SecretUndeleteSuccess      = SecretUndeleteResponse{}.Success()
	SecretUndeleteInternal     = SecretUndeleteResponse{}.Internal()
)

// SecretListResponse variants
var (
	SecretListBadRequest   = SecretListResponse{}.BadRequest()
	SecretListUnauthorized = SecretListResponse{}.Unauthorized()
	SecretListSuccess      = SecretListResponse{}.Success()
	SecretListInternal     = SecretListResponse{}.Internal()
)

// SecretPutResponse variants
var (
	SecretPutBadRequest   = SecretPutResponse{}.BadRequest()
	SecretPutUnauthorized = SecretPutResponse{}.Unauthorized()
	SecretPutSuccess      = SecretPutResponse{}.Success()
	SecretPutInternal     = SecretPutResponse{}.Internal()
)

// SecretMetadataResponse variants
var (
	SecretMetadataNotFound     = SecretMetadataResponse{}.NotFound()
	SecretMetadataBadRequest   = SecretMetadataResponse{}.BadRequest()
	SecretMetadataUnauthorized = SecretMetadataResponse{}.Unauthorized()
	SecretMetadataSuccess      = SecretMetadataResponse{}.Success()
	SecretMetadataInternal     = SecretMetadataResponse{}.Internal()
)

// PolicyPutResponse variants
var (
	PolicyPutBadRequest   = PolicyPutResponse{}.BadRequest()
	PolicyPutUnauthorized = PolicyPutResponse{}.Unauthorized()
	PolicyPutSuccess      = PolicyPutResponse{}.Success()
	PolicyPutInternal     = PolicyPutResponse{}.Internal()
)

// PolicyReadResponse variants
var (
	PolicyReadNotFound     = PolicyReadResponse{}.NotFound()
	PolicyReadBadRequest   = PolicyReadResponse{}.BadRequest()
	PolicyReadUnauthorized = PolicyReadResponse{}.Unauthorized()
	PolicyReadSuccess      = PolicyReadResponse{}.Success()
	PolicyReadInternal     = PolicyReadResponse{}.Internal()
)

// PolicyDeleteResponse variants
var (
	PolicyDeleteNotFound     = PolicyDeleteResponse{}.NotFound()
	PolicyDeleteBadRequest   = PolicyDeleteResponse{}.BadRequest()
	PolicyDeleteUnauthorized = PolicyDeleteResponse{}.Unauthorized()
	PolicyDeleteSuccess      = PolicyDeleteResponse{}.Success()
	PolicyDeleteInternal     = PolicyDeleteResponse{}.Internal()
)

// PolicyListResponse variants
var (
	PolicyListBadRequest   = PolicyListResponse{}.BadRequest()
	PolicyListUnauthorized = PolicyListResponse{}.Unauthorized()
	PolicyListSuccess      = PolicyListResponse{}.Success()
	PolicyListInternal     = PolicyListResponse{}.Internal()
)

// PolicyAccessCheckResponse variants
var (
	PolicyAccessCheckNotFound     = PolicyAccessCheckResponse{}.NotFound()
	PolicyAccessCheckBadRequest   = PolicyAccessCheckResponse{}.BadRequest()
	PolicyAccessCheckUnauthorized = PolicyAccessCheckResponse{}.Unauthorized()
	PolicyAccessCheckSuccess      = PolicyAccessCheckResponse{}.Success()
	PolicyAccessCheckInternal     = PolicyAccessCheckResponse{}.Internal()
)

// CipherDecryptResponse variants
var (
	CipherDecryptBadRequest   = CipherDecryptResponse{}.BadRequest()
	CipherDecryptUnauthorized = CipherDecryptResponse{}.Unauthorized()
	CipherDecryptSuccess      = CipherDecryptResponse{}.Success()
	CipherDecryptInternal     = CipherDecryptResponse{}.Internal()
)

// CipherEncryptResponse variants
var (
	CipherEncryptBadRequest   = CipherEncryptResponse{}.BadRequest()
	CipherEncryptUnauthorized = CipherEncryptResponse{}.Unauthorized()
	CipherEncryptSuccess      = CipherEncryptResponse{}.Success()
	CipherEncryptInternal     = CipherEncryptResponse{}.Internal()
)

// RecoverResponse variants
var (
	RecoverBadRequest   = RecoverResponse{}.BadRequest()
	RecoverUnauthorized = RecoverResponse{}.Unauthorized()
	RecoverSuccess      = RecoverResponse{}.Success()
	RecoverInternal     = RecoverResponse{}.Internal()
)

// RestoreResponse variants
var (
	RestoreBadRequest   = RestoreResponse{}.BadRequest()
	RestoreUnauthorized = RestoreResponse{}.Unauthorized()
	RestoreSuccess      = RestoreResponse{}.Success()
	RestoreInternal     = RestoreResponse{}.Internal()
)

// BoostrapResponse variants
var (
	BootstrapBadRequest   = BootstrapVerifyResponse{}.BadRequest()
	BootstrapUnauthorized = BootstrapVerifyResponse{}.Unauthorized()
	BootstrapSuccess      = BootstrapVerifyResponse{}.Success()
	BootstrapInternal     = BootstrapVerifyResponse{}.Internal()
)
