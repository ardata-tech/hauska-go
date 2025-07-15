package hauska

import "errors"

// SDK error definitions
var (
	// Configuration errors
	ErrInvalidClient         = errors.New("invalid ethereum client")
	ErrInvalidFactoryAddress = errors.New("invalid factory contract address")
	ErrInvalidUSDCAddress    = errors.New("invalid USDC token address")
	ErrMissingAuth           = errors.New("missing authentication (private key or auth)")
	ErrInvalidPrivateKey     = errors.New("invalid private key format")

	// Contract errors
	ErrContractNotFound      = errors.New("contract not found")
	ErrInvalidContractCall   = errors.New("invalid contract call")
	ErrTransactionFailed     = errors.New("transaction failed")
	ErrInsufficientBalance   = errors.New("insufficient balance")
	ErrInsufficientAllowance = errors.New("insufficient allowance")

	// Asset errors
	ErrAssetNotFound      = errors.New("asset not found")
	ErrAssetNotVerified   = errors.New("asset not verified")
	ErrAssetNotLicensable = errors.New("asset cannot be licensed")
	ErrAssetAlreadyExists = errors.New("asset already exists")
	ErrInvalidAssetHash   = errors.New("invalid asset hash")

	// License errors
	ErrLicenseNotFound    = errors.New("license not found")
	ErrLicenseExpired     = errors.New("license expired")
	ErrLicenseNotActive   = errors.New("license not active")
	ErrAlreadyLicensed    = errors.New("already licensed")
	ErrInvalidLicenseType = errors.New("invalid license type")

	// Group errors
	ErrGroupNotFound       = errors.New("group not found")
	ErrGroupEmpty          = errors.New("group cannot be empty")
	ErrAssetNotInGroup     = errors.New("asset not in group")
	ErrAssetAlreadyInGroup = errors.New("asset already in group")

	// Organization errors
	ErrOrgNotFound   = errors.New("organization not found")
	ErrNotAuthorized = errors.New("not authorized")
	ErrInvalidRole   = errors.New("invalid role")

	// General errors
	ErrInvalidAddress    = errors.New("invalid address")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInvalidParameters = errors.New("invalid parameters")
	ErrTimeout           = errors.New("operation timeout")
	ErrNotImplemented    = errors.New("not implemented")
)
