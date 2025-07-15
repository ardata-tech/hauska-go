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

// AssetService implements asset registry operations
type AssetService struct {
	client        *ethclient.Client
	assetRegistry *contracts.HauskaAssetRegistry
	assetNFT      *contracts.HauskaAssetNFT
	auth          *bind.TransactOpts
}

// NewAssetService creates a new asset service instance
func NewAssetService(client *ethclient.Client, assetRegistryAddr, assetNFTAddr common.Address, auth *bind.TransactOpts) (*AssetService, error) {
	assetRegistry, err := contracts.NewHauskaAssetRegistry(assetRegistryAddr, client)
	if err != nil {
		return nil, err
	}

	assetNFT, err := contracts.NewHauskaAssetNFT(assetNFTAddr, client)
	if err != nil {
		return nil, err
	}

	return &AssetService{
		client:        client,
		assetRegistry: assetRegistry,
		assetNFT:      assetNFT,
		auth:          auth,
	}, nil
}

// RegisterAsset registers a new asset in the registry
func (s *AssetService) RegisterAsset(ctx context.Context, orgContract string, asset *types.VerifiedDigitalAsset, creator string) (*types.TransactionResult, error) {
	// Convert to contract struct
	contractAsset := contracts.IHauskaStructsVerifiedDigitalAsset{
		AssetId:                asset.AssetID,
		Creator:                asset.Creator,
		Owner:                  asset.Owner,
		Partner:                asset.Partner,
		IpfsHash:               asset.IPFSHash,
		MetadataHash:           asset.MetadataHash,
		AssetHash:              asset.AssetHash,
		Version:                asset.Version,
		IsVerified:             asset.IsVerified,
		CreationTime:           asset.CreationTime,
		LastTransferTime:       asset.LastTransferTime,
		Price:                  asset.Price,
		Encrypted:              asset.Encrypted,
		CanBeLicensed:          asset.CanBeLicensed,
		FxPool:                 uint8(asset.FxPool),
		EventTimestamp:         asset.EventTimestamp,
		GeographicRestrictions: convertCountryCodes(asset.GeographicRestrictions),
	}

	tx, err := s.assetRegistry.RegisterAsset(s.auth, contractAsset, common.HexToAddress(creator))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAsset returns asset details by organization and asset ID
func (s *AssetService) GetAsset(ctx context.Context, orgContract string, assetID *big.Int) (*types.VerifiedDigitalAsset, error) {
	contractAsset, err := s.assetRegistry.GetAsset(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), assetID)
	if err != nil {
		return nil, err
	}

	// Convert from contract struct to types
	asset := &types.VerifiedDigitalAsset{
		AssetID:                contractAsset.AssetId,
		Creator:                contractAsset.Creator,
		Owner:                  contractAsset.Owner,
		Partner:                contractAsset.Partner,
		IPFSHash:               contractAsset.IpfsHash,
		MetadataHash:           contractAsset.MetadataHash,
		AssetHash:              contractAsset.AssetHash,
		Version:                contractAsset.Version,
		IsVerified:             contractAsset.IsVerified,
		CreationTime:           contractAsset.CreationTime,
		LastTransferTime:       contractAsset.LastTransferTime,
		Price:                  contractAsset.Price,
		Encrypted:              contractAsset.Encrypted,
		CanBeLicensed:          contractAsset.CanBeLicensed,
		FxPool:                 types.FxPool(contractAsset.FxPool),
		EventTimestamp:         contractAsset.EventTimestamp,
		GeographicRestrictions: convertCountryCodesFromContract(contractAsset.GeographicRestrictions),
	}

	return asset, nil
}

// VerifyAsset verifies an asset (marks it as verified)
func (s *AssetService) VerifyAsset(ctx context.Context, orgContract string, assetID *big.Int, verifier string) (*types.TransactionResult, error) {
	tx, err := s.assetRegistry.VerifyAsset(s.auth, common.HexToAddress(orgContract), assetID, common.HexToAddress(verifier))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// IsAssetVerified checks if an asset is verified
func (s *AssetService) IsAssetVerified(ctx context.Context, orgContract string, assetID *big.Int) (bool, error) {
	return s.assetRegistry.IsAssetVerified(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), assetID)
}

// GetAssetsByCreator returns all assets created by a specific creator
func (s *AssetService) GetAssetsByCreator(ctx context.Context, orgContract, creator string) ([]*big.Int, error) {
	return s.assetRegistry.GetAssetsByCreator(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), common.HexToAddress(creator))
}

// GetAssetsByOwner returns all assets owned by a specific owner
func (s *AssetService) GetAssetsByOwner(ctx context.Context, orgContract, owner string) ([]*big.Int, error) {
	// Get the organization contract instance
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return nil, err
	}
	return orgContractInstance.GetAssetsByOwner(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner))
}

// TransferAssetOwnership transfers ownership of an asset
func (s *AssetService) TransferAssetOwnership(ctx context.Context, orgContract string, assetID *big.Int, newOwner, caller string) (*types.TransactionResult, error) {
	tx, err := s.assetRegistry.TransferAssetOwnership(s.auth, common.HexToAddress(orgContract), assetID, common.HexToAddress(newOwner), common.HexToAddress(caller))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// TransferAssetCrossOrg transfers an asset between organizations
func (s *AssetService) TransferAssetCrossOrg(ctx context.Context, fromOrg, toOrg string, assetID *big.Int, newOwner string) (*types.TransactionResult, error) {
	tx, err := s.assetRegistry.TransferAssetCrossOrg(s.auth, common.HexToAddress(fromOrg), common.HexToAddress(toOrg), assetID, common.HexToAddress(newOwner))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAssetCount returns the total number of assets for an organization
func (s *AssetService) GetAssetCount(ctx context.Context, orgContract string) (*big.Int, error) {
	// TODO: Implement contract call
	// 1. Call assetRegistry.GetAssetCount()
	// 2. Return count

	return nil, types.ErrNotImplemented
}

// GetAssetNFTAddress returns the address of the asset NFT contract
func (s *AssetService) GetAssetNFTAddress(ctx context.Context) (string, error) {
	// TODO: Implement contract call
	// 1. Call assetRegistry.AssetNFT()
	// 2. Return address as string

	return "", types.ErrNotImplemented
}

// GetAssetNFTTokenID returns the NFT token ID for an asset
func (s *AssetService) GetAssetNFTTokenID(ctx context.Context, orgContract string, assetID *big.Int) (*big.Int, error) {
	return s.assetNFT.GetTokenIdForAsset(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), assetID)
}

// TransferAssetNFT transfers an asset NFT
func (s *AssetService) TransferAssetNFT(ctx context.Context, from, to string, tokenID *big.Int) (*types.TransactionResult, error) {
	tx, err := s.assetNFT.SafeTransferFrom(s.auth, common.HexToAddress(from), common.HexToAddress(to), tokenID)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAssetNFTOwner returns the owner of an asset NFT
func (s *AssetService) GetAssetNFTOwner(ctx context.Context, tokenID *big.Int) (string, error) {
	owner, err := s.assetNFT.OwnerOf(&bind.CallOpts{Context: ctx}, tokenID)
	if err != nil {
		return "", err
	}
	return owner.Hex(), nil
}

// Helper functions

func convertCountryCodes(codes []types.CountryCode) []uint8 {
	result := make([]uint8, len(codes))
	for i, code := range codes {
		result[i] = uint8(code)
	}
	return result
}

func convertCountryCodesFromContract(codes []uint8) []types.CountryCode {
	result := make([]types.CountryCode, len(codes))
	for i, code := range codes {
		result[i] = types.CountryCode(code)
	}
	return result
}

func (s *AssetService) waitForTransaction(ctx context.Context, txHash common.Hash) (*types.TransactionResult, error) {
	// TODO: Implement transaction waiting and parsing
	// This is a placeholder implementation
	return &types.TransactionResult{
		TxHash: txHash,
		Status: 1,
	}, nil
}
