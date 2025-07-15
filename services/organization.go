package services

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/contracts"
	"github.com/ardata-tech/hauska-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// OrganizationService implements organization contract operations
type OrganizationService struct {
	client      *ethclient.Client
	orgContract *contracts.HauskaOrgContract
	auth        *bind.TransactOpts
	orgAddress  common.Address
}

// NewOrganizationService creates a new organization service instance
func NewOrganizationService(client *ethclient.Client, orgAddress common.Address, auth *bind.TransactOpts) (*OrganizationService, error) {
	orgContract, err := contracts.NewHauskaOrgContract(orgAddress, client)
	if err != nil {
		return nil, err
	}

	return &OrganizationService{
		client:      client,
		orgContract: orgContract,
		auth:        auth,
		orgAddress:  orgAddress,
	}, nil
}

// AddCreator adds a new creator to the organization
func (s *OrganizationService) AddCreator(ctx context.Context, creatorAddress string) (*types.TransactionResult, error) {
	tx, err := s.orgContract.AddCreator(s.auth, common.HexToAddress(creatorAddress))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// RemoveCreator removes a creator from the organization
func (s *OrganizationService) RemoveCreator(ctx context.Context, creatorAddress string) (*types.TransactionResult, error) {
	tx, err := s.orgContract.RemoveCreator(s.auth, common.HexToAddress(creatorAddress))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetCreators returns all creators in the organization
func (s *OrganizationService) GetCreators(ctx context.Context) ([]string, error) {
	creatorRole, err := s.orgContract.CREATORROLE(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	// TODO: Use event filtering to get all creators with CREATOR_ROLE
	// For now, return empty slice as the contract doesn't have a direct getCreators method
	_ = creatorRole
	return []string{}, nil
}

// IsCreator checks if an address has creator role
func (s *OrganizationService) IsCreator(ctx context.Context, address string) (bool, error) {
	creatorRole, err := s.orgContract.CREATORROLE(&bind.CallOpts{Context: ctx})
	if err != nil {
		return false, err
	}

	hasRole, err := s.orgContract.HasRole(&bind.CallOpts{Context: ctx}, creatorRole, common.HexToAddress(address))
	if err != nil {
		return false, err
	}

	return hasRole, nil
}

// Asset management methods

// CreateAsset creates a new asset in the organization
func (s *OrganizationService) CreateAsset(ctx context.Context, req *types.CreateAssetRequest) (*types.TransactionResult, error) {
	geoRestrictions := s.convertCountryCodes(req.GeographicRestrictions)

	tx, err := s.orgContract.CreateAsset(
		s.auth,
		req.AssetCID,
		req.MetadataCID,
		req.AssetHash,
		req.Price,
		req.IsEncrypted,
		req.CanBeLicensed,
		uint8(req.FxPool),
		req.Timestamp,
		geoRestrictions,
	)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAsset returns asset details by ID
func (s *OrganizationService) GetAsset(ctx context.Context, assetID *big.Int) (*types.VerifiedDigitalAsset, error) {
	id, creator, owner, partner, ipfsHash, metadataHash, assetHash, version, isVerified,
		creationTime, lastTransferTime, price, encrypted, canBeLicensed, fxPool, eventTimestamp, err := s.orgContract.Assets(&bind.CallOpts{Context: ctx}, assetID)
	if err != nil {
		return nil, err
	}

	return s.convertToVerifiedAsset(
		id, version, creationTime, lastTransferTime, price,
		creator, owner, partner,
		ipfsHash, metadataHash, assetHash,
		isVerified, encrypted, canBeLicensed,
		fxPool, eventTimestamp,
	), nil
}

// GetAssetsByOwner returns all assets owned by a specific address
func (s *OrganizationService) GetAssetsByOwner(ctx context.Context, owner string) ([]*big.Int, error) {
	assetIDs, err := s.orgContract.GetAssetsByOwner(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner))
	if err != nil {
		return nil, err
	}

	return assetIDs, nil
}

// GetAssetsByCreator returns all assets created by a specific creator
func (s *OrganizationService) GetAssetsByCreator(ctx context.Context, creator string) ([]*big.Int, error) {
	// The organization contract doesn't have this method directly
	// Would need to use AssetRegistry or event filtering
	// TODO: Implement via AssetRegistry contract call or event parsing
	return []*big.Int{}, nil
}

// VerifyAsset verifies an asset (marks it as verified)
func (s *OrganizationService) VerifyAsset(ctx context.Context, assetID *big.Int) (*types.TransactionResult, error) {
	tx, err := s.orgContract.VerifyAsset(s.auth, assetID)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// UpdateAsset updates asset metadata and properties
func (s *OrganizationService) UpdateAsset(ctx context.Context, assetID *big.Int, updates *types.CreateAssetRequest) (*types.TransactionResult, error) {
	// The organization contract doesn't have an update method directly
	// Asset updates would typically be handled via AssetRegistry
	// TODO: Implement via AssetRegistry contract
	return &types.TransactionResult{Status: 1}, nil
}

// TransferAsset transfers asset ownership to a new owner
func (s *OrganizationService) TransferAsset(ctx context.Context, assetID *big.Int, newOwner string) (*types.TransactionResult, error) {
	tx, err := s.orgContract.TransferAsset(s.auth, assetID, common.HexToAddress(newOwner))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// License management methods

// LicenseAsset licenses an asset to a user
func (s *OrganizationService) LicenseAsset(ctx context.Context, req *types.LicenseAssetRequest) (*types.TransactionResult, error) {
	permissions := make([]uint8, len(req.Permissions))
	for i, perm := range req.Permissions {
		permissions[i] = uint8(perm)
	}

	tx, err := s.orgContract.LicenseAsset(s.auth, req.AssetID, permissions, req.ResellerFee)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// LicenseGroup licenses an asset group to a user
func (s *OrganizationService) LicenseGroup(ctx context.Context, groupID *big.Int, licensee string, permissions []types.LicensePermissions, resellerFee *big.Int) (*types.TransactionResult, error) {
	permissionInts := make([]uint8, len(permissions))
	for i, perm := range permissions {
		permissionInts[i] = uint8(perm)
	}

	tx, err := s.orgContract.LicenseGroup(s.auth, groupID, permissionInts, resellerFee)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetUserLicenses returns all licenses owned by a user in this organization
func (s *OrganizationService) GetUserLicenses(ctx context.Context, user string) ([]*big.Int, error) {
	// The organization contract doesn't have this method directly
	// Would need to call LicenseManager contract
	// TODO: Implement via LicenseManager contract
	return []*big.Int{}, nil
}

// IsAssetLicensedBy checks if a user has a valid license for an asset
func (s *OrganizationService) IsAssetLicensedBy(ctx context.Context, assetID *big.Int, user string) (bool, error) {
	// The organization contract doesn't have this method directly
	// Would need to check via LicenseManager contract
	// TODO: Implement via LicenseManager contract call
	return false, nil
}

// Group management methods

// CreateGroup creates a new asset group
func (s *OrganizationService) CreateGroup(ctx context.Context, req *types.CreateGroupRequest) (*types.TransactionResult, error) {
	tx, err := s.orgContract.CreateGroup(s.auth, req.GroupName, req.AssetIDs, req.GroupPrice)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetGroup returns group details by ID
func (s *OrganizationService) GetGroup(ctx context.Context, groupID *big.Int) (*types.AssetGroup, error) {
	// The organization contract doesn't have a direct getGroup method
	// Would need to call GroupManager contract
	// TODO: Implement via GroupManager contract
	return &types.AssetGroup{
		GroupID:  groupID,
		IsActive: true,
	}, nil
}

// GetGroupsByOwner returns all groups owned by a specific address
func (s *OrganizationService) GetGroupsByOwner(ctx context.Context, owner string) ([]*big.Int, error) {
	groupIDs, err := s.orgContract.GetAssetGroupsByOwner(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner))
	if err != nil {
		return nil, err
	}

	return groupIDs, nil
}

// AddAssetToGroup adds an asset to an existing group
func (s *OrganizationService) AddAssetToGroup(ctx context.Context, groupID, assetID *big.Int) (*types.TransactionResult, error) {
	// The organization contract doesn't have this method directly
	// Would need to call GroupManager contract
	// TODO: Implement via GroupManager contract
	return &types.TransactionResult{Status: 1}, nil
}

// RemoveAssetFromGroup removes an asset from a group
func (s *OrganizationService) RemoveAssetFromGroup(ctx context.Context, groupID, assetID *big.Int) (*types.TransactionResult, error) {
	// The organization contract doesn't have this method directly
	// Would need to call GroupManager contract
	// TODO: Implement via GroupManager contract
	return &types.TransactionResult{Status: 1}, nil
}

// UpdateGroupPrice updates the price of a group
func (s *OrganizationService) UpdateGroupPrice(ctx context.Context, groupID *big.Int, newPrice *big.Int) (*types.TransactionResult, error) {
	// The organization contract doesn't have this method directly
	// Would need to call GroupManager contract
	// TODO: Implement via GroupManager contract
	return &types.TransactionResult{Status: 1}, nil
}

// RemoveGroup removes a group (makes it inactive)
func (s *OrganizationService) RemoveGroup(ctx context.Context, groupID *big.Int) (*types.TransactionResult, error) {
	tx, err := s.orgContract.RemoveGroup(s.auth, groupID)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// Information methods

// GetInfo returns organization information
func (s *OrganizationService) GetInfo(ctx context.Context) (*types.Organization, error) {
	assetRegistry, err := s.orgContract.AssetRegistry(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	licenseManager, err := s.orgContract.LicenseManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	groupManager, err := s.orgContract.GroupManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	revenueDistributor, err := s.orgContract.RevenueDistributor(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	principal, err := s.orgContract.Principal(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	integrationPartner, err := s.orgContract.IntegrationPartner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	factory, err := s.orgContract.Factory(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	assetCount, err := s.orgContract.GetAssetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	creatorCount, err := s.orgContract.GetCreatorCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return &types.Organization{
		Address:            s.orgAddress,
		Principal:          principal,
		IntegrationPartner: integrationPartner,
		Factory:            factory,
		LicenseManager:     licenseManager,
		AssetRegistry:      assetRegistry,
		GroupManager:       groupManager,
		RevenueDistributor: revenueDistributor,
		AssetCount:         assetCount,
		CreatorCount:       creatorCount,
	}, nil
}

// GetAssetCount returns the total number of assets in the organization
func (s *OrganizationService) GetAssetCount(ctx context.Context) (*big.Int, error) {
	count, err := s.orgContract.GetAssetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return count, nil
}

// GetCreatorCount returns the total number of creators in the organization
func (s *OrganizationService) GetCreatorCount(ctx context.Context) (*big.Int, error) {
	count, err := s.orgContract.GetCreatorCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return count, nil
}

// GetGroupCount returns the total number of groups in the organization
func (s *OrganizationService) GetGroupCount(ctx context.Context) (*big.Int, error) {
	// The organization contract doesn't have this method directly
	// Would need to call GroupManager contract
	// TODO: Implement via GroupManager contract
	return big.NewInt(0), nil
}

// waitForTransaction waits for a transaction to be mined and returns the result
func (s *OrganizationService) waitForTransaction(ctx context.Context, txHash common.Hash) (*types.TransactionResult, error) {
	// TODO: Implement transaction waiting and parsing
	// This is a placeholder implementation
	return &types.TransactionResult{
		TxHash: txHash,
		Status: 1,
	}, nil
}

// convertToVerifiedAsset converts contract asset data to VerifiedDigitalAsset struct
func (s *OrganizationService) convertToVerifiedAsset(
	assetID, version, creationTime, lastTransferTime, price *big.Int,
	creator, owner, partner common.Address,
	ipfsHash, metadataHash string,
	assetHash [32]byte,
	isVerified, encrypted, canBeLicensed bool,
	fxPool uint8,
	eventTimestamp string,
) *types.VerifiedDigitalAsset {
	return &types.VerifiedDigitalAsset{
		AssetID:          assetID,
		Creator:          creator,
		Owner:            owner,
		Partner:          partner,
		IPFSHash:         ipfsHash,
		MetadataHash:     metadataHash,
		AssetHash:        assetHash,
		Version:          version,
		IsVerified:       isVerified,
		CreationTime:     creationTime,
		LastTransferTime: lastTransferTime,
		Price:            price,
		Encrypted:        encrypted,
		CanBeLicensed:    canBeLicensed,
		FxPool:           types.FxPool(fxPool),
		EventTimestamp:   eventTimestamp,
		// GeographicRestrictions will need to be fetched separately if needed
	}
}

// convertCountryCodes converts Go CountryCode slice to uint8 slice for contract calls
func (s *OrganizationService) convertCountryCodes(codes []types.CountryCode) []uint8 {
	result := make([]uint8, len(codes))
	for i, code := range codes {
		result[i] = uint8(code)
	}
	return result
}
