package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/types"
)

// FactoryClient provides high-level access to factory contract operations
type FactoryClient interface {
	// Organization management
	CreateOrganization(ctx context.Context, principal, integrationPartner string) (*types.TransactionResult, error)
	GetAllOrganizations(ctx context.Context) ([]string, error)
	GetOrganization(ctx context.Context, principal string) (*types.Organization, error)
	RemoveOrganization(ctx context.Context, orgAddress string) (*types.TransactionResult, error)

	// Platform management
	GetPlatformFees(ctx context.Context) (*types.PlatformFees, error)
	UpdatePlatformFees(ctx context.Context, hauskaFeePct, integratorFeePct uint32) (*types.TransactionResult, error)

	// Module information
	GetModules(ctx context.Context) (*types.ModuleAddresses, error)
	IsValidOrgContract(ctx context.Context, address string) (bool, error)

	// Statistics
	GetNumberOfOrganizations(ctx context.Context) (*big.Int, error)
	GetNumberOfCreators(ctx context.Context, orgAddress string) (*big.Int, error)
}

// OrganizationClient provides high-level access to organization contract operations
type OrganizationClient interface {
	// Creator management
	AddCreator(ctx context.Context, creatorAddress string) (*types.TransactionResult, error)
	RemoveCreator(ctx context.Context, creatorAddress string) (*types.TransactionResult, error)
	GetCreators(ctx context.Context) ([]string, error)
	IsCreator(ctx context.Context, address string) (bool, error)

	// Asset management
	CreateAsset(ctx context.Context, req *types.CreateAssetRequest) (*types.TransactionResult, error)
	GetAsset(ctx context.Context, assetID *big.Int) (*types.VerifiedDigitalAsset, error)
	GetAssetsByOwner(ctx context.Context, owner string) ([]*big.Int, error)
	GetAssetsByCreator(ctx context.Context, creator string) ([]*big.Int, error)
	VerifyAsset(ctx context.Context, assetID *big.Int) (*types.TransactionResult, error)
	UpdateAsset(ctx context.Context, assetID *big.Int, updates *types.CreateAssetRequest) (*types.TransactionResult, error)
	TransferAsset(ctx context.Context, assetID *big.Int, newOwner string) (*types.TransactionResult, error)

	// License management
	LicenseAsset(ctx context.Context, req *types.LicenseAssetRequest) (*types.TransactionResult, error)
	LicenseGroup(ctx context.Context, groupID *big.Int, licensee string, permissions []types.LicensePermissions, resellerFee *big.Int) (*types.TransactionResult, error)
	GetUserLicenses(ctx context.Context, user string) ([]*big.Int, error)
	IsAssetLicensedBy(ctx context.Context, assetID *big.Int, user string) (bool, error)

	// Group management
	CreateGroup(ctx context.Context, req *types.CreateGroupRequest) (*types.TransactionResult, error)
	GetGroup(ctx context.Context, groupID *big.Int) (*types.AssetGroup, error)
	GetGroupsByOwner(ctx context.Context, owner string) ([]*big.Int, error)
	AddAssetToGroup(ctx context.Context, groupID, assetID *big.Int) (*types.TransactionResult, error)
	RemoveAssetFromGroup(ctx context.Context, groupID, assetID *big.Int) (*types.TransactionResult, error)
	UpdateGroupPrice(ctx context.Context, groupID *big.Int, newPrice *big.Int) (*types.TransactionResult, error)
	RemoveGroup(ctx context.Context, groupID *big.Int) (*types.TransactionResult, error)

	// Information
	GetInfo(ctx context.Context) (*types.Organization, error)
	GetAssetCount(ctx context.Context) (*big.Int, error)
	GetCreatorCount(ctx context.Context) (*big.Int, error)
	GetGroupCount(ctx context.Context) (*big.Int, error)
}

// LicenseClient provides high-level access to license management operations
type LicenseClient interface {
	// License operations
	GetLicense(ctx context.Context, licenseID *big.Int) (*types.LicenseData, error)
	GetLicensesByUser(ctx context.Context, user string) ([]*big.Int, error)
	IsLicenseValid(ctx context.Context, licenseID *big.Int) (bool, error)
	RelicenseAsset(ctx context.Context, orgContract string, assetID *big.Int, licensee string, existingLicenseID *big.Int) (*types.TransactionResult, error)
	RenewLicense(ctx context.Context, licenseID *big.Int, additionalDuration *big.Int) (*types.TransactionResult, error)
	RevokeLicense(ctx context.Context, licenseID *big.Int) (*types.TransactionResult, error)

	// Batch operations
	LicenseMultipleAssets(ctx context.Context, requests []*types.LicenseAssetRequest) ([]*types.TransactionResult, error)

	// Query operations
	GetAssetLicensees(ctx context.Context, orgContract string, assetID *big.Int) ([]string, error)
	GetLicenseHistory(ctx context.Context, assetID *big.Int) ([]*types.LicenseData, error)
}

// AssetClient provides high-level access to asset management operations
type AssetClient interface {
	// Asset registry operations
	RegisterAsset(ctx context.Context, orgContract string, asset *types.VerifiedDigitalAsset, creator string) (*types.TransactionResult, error)
	GetRegisteredAsset(ctx context.Context, orgContract string, assetID *big.Int) (*types.VerifiedDigitalAsset, error)
	IsAssetVerified(ctx context.Context, orgContract string, assetID *big.Int) (bool, error)
	GetAssetsByCreator(ctx context.Context, orgContract, creator string) ([]*big.Int, error)
	TransferAssetOwnership(ctx context.Context, orgContract string, assetID *big.Int, newOwner, caller string) (*types.TransactionResult, error)

	// Cross-organization operations
	TransferAssetCrossOrg(ctx context.Context, fromOrg, toOrg string, assetID *big.Int, newOwner string) (*types.TransactionResult, error)

	// Asset NFT operations
	GetAssetNFTTokenID(ctx context.Context, orgContract string, assetID *big.Int) (*big.Int, error)
	GetAssetNFTOwner(ctx context.Context, tokenID *big.Int) (string, error)
	TransferAssetNFT(ctx context.Context, from, to string, tokenID *big.Int) (*types.TransactionResult, error)
}

// GroupClient provides high-level access to group management operations
type GroupClient interface {
	// Group operations
	CreateAssetGroup(ctx context.Context, orgContract, groupName string, assetIDs []*big.Int, groupPrice *big.Int, creator string) (*types.TransactionResult, error)
	GetAssetGroup(ctx context.Context, orgContract string, groupID *big.Int) (*types.AssetGroup, error)
	GetGroupsByCreator(ctx context.Context, orgContract, creator string) ([]*big.Int, error)
	AddAssetToGroup(ctx context.Context, orgContract string, groupID, assetID *big.Int, caller string) (*types.TransactionResult, error)
	RemoveAssetFromGroup(ctx context.Context, orgContract string, groupID, assetID *big.Int, caller string) (*types.TransactionResult, error)
	UpdateGroupPrice(ctx context.Context, orgContract string, groupID *big.Int, newPrice *big.Int, caller string) (*types.TransactionResult, error)
	RemoveAssetGroup(ctx context.Context, orgContract string, groupID *big.Int, caller string) (*types.TransactionResult, error)
	GetGroupCount(ctx context.Context, orgContract string) (*big.Int, error)
}

// RevenueClient provides high-level access to revenue distribution operations
type RevenueClient interface {
	// Revenue distribution
	DistributeRevenue(ctx context.Context, from string, amount *big.Int, assetOwner, integrationPartner, orgContract string) (*types.TransactionResult, error)
	GetRevenueStats(ctx context.Context, orgContract string) (*types.RevenueStats, error)
	GetEarnings(ctx context.Context, account string) (*big.Int, error)

	// Fee management
	SetCustomFees(ctx context.Context, orgContract string, hauskaFeePct, integratorFeePct uint32) (*types.TransactionResult, error)
	RemoveCustomFees(ctx context.Context, orgContract string) (*types.TransactionResult, error)
	GetCustomFees(ctx context.Context, orgContract string) (*types.PlatformFees, error)
}

// USDCClient provides high-level access to USDC token operations
type USDCClient interface {
	// Balance operations
	GetBalance(ctx context.Context, account string) (*big.Int, error)
	Transfer(ctx context.Context, to string, amount *big.Int) (*types.TransactionResult, error)

	// Allowance operations
	Approve(ctx context.Context, spender string, amount *big.Int) (*types.TransactionResult, error)
	GetAllowance(ctx context.Context, owner, spender string) (*big.Int, error)

	// Token information
	GetDecimals(ctx context.Context) (uint8, error)
	GetSymbol(ctx context.Context) (string, error)
	GetName(ctx context.Context) (string, error)
	GetTotalSupply(ctx context.Context) (*big.Int, error)

	// Utility functions
	FormatAmount(amount *big.Int) string
	ParseAmount(amount string) (*big.Int, error)
}
