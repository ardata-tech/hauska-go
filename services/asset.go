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

// RegisterAsset registers a new asset in the registry via organization contract
func (s *AssetService) RegisterAsset(ctx context.Context, orgContract string, asset *types.VerifiedDigitalAsset, creator string) (*types.TransactionResult, error) {
	// Get the organization contract instance to register through it
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return nil, err
	}

	// Update nonce before transaction
	nonce, err := s.client.PendingNonceAt(ctx, s.auth.From)
	if err != nil {
		return nil, err
	}
	s.auth.Nonce = big.NewInt(int64(nonce))

	// Register through organization contract using CreateAsset method
	tx, err := orgContractInstance.CreateAsset(s.auth,
		asset.IPFSHash,
		asset.MetadataHash,
		asset.AssetHash,
		asset.Price,
		asset.Encrypted,
		asset.CanBeLicensed,
		uint8(asset.FxPool),
		asset.EventTimestamp,
		convertCountryCodes(asset.GeographicRestrictions),
	)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAsset returns asset details by organization and asset ID
func (s *AssetService) GetAsset(ctx context.Context, orgContract string, assetID *big.Int) (*types.VerifiedDigitalAsset, error) {
	// Get the organization contract instance
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return nil, err
	}

	// Get asset details from organization contract
	assetId, creator, owner, partner, ipfsHash, metadataHash, assetHash, version, isVerified, creationTime, lastTransferTime, price, encrypted, canBeLicensed, fxPool, eventTimestamp, err := orgContractInstance.Assets(&bind.CallOpts{Context: ctx}, assetID)
	if err != nil {
		return nil, err
	}

	// Convert from contract data to types
	asset := &types.VerifiedDigitalAsset{
		AssetID:                assetId,
		Creator:                creator,
		Owner:                  owner,
		Partner:                partner,
		IPFSHash:               ipfsHash,
		MetadataHash:           metadataHash,
		AssetHash:              assetHash,
		Version:                version,
		IsVerified:             isVerified,
		CreationTime:           creationTime,
		LastTransferTime:       lastTransferTime,
		Price:                  price,
		Encrypted:              encrypted,
		CanBeLicensed:          canBeLicensed,
		FxPool:                 types.FxPool(fxPool),
		EventTimestamp:         eventTimestamp,
		GeographicRestrictions: []types.CountryCode{}, // TODO: Get from contract if available
	}

	return asset, nil
}

// VerifyAsset verifies an asset (marks it as verified)
func (s *AssetService) VerifyAsset(ctx context.Context, orgContract string, assetID *big.Int, verifier string) (*types.TransactionResult, error) {
	// Get the organization contract instance
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return nil, err
	}

	// Update nonce before transaction
	nonce, err := s.client.PendingNonceAt(ctx, s.auth.From)
	if err != nil {
		return nil, err
	}
	s.auth.Nonce = big.NewInt(int64(nonce))

	tx, err := orgContractInstance.VerifyAsset(s.auth, assetID)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// IsAssetVerified checks if an asset is verified
func (s *AssetService) IsAssetVerified(ctx context.Context, orgContract string, assetID *big.Int) (bool, error) {
	// Get the organization contract instance
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return false, err
	}

	// Get asset details to check verification status
	_, _, _, _, _, _, _, _, isVerified, _, _, _, _, _, _, _, err := orgContractInstance.Assets(&bind.CallOpts{Context: ctx}, assetID)
	if err != nil {
		return false, err
	}

	return isVerified, nil
}

// GetAssetsByCreator returns all assets created by a specific creator
func (s *AssetService) GetAssetsByCreator(ctx context.Context, orgContract, creator string) ([]*big.Int, error) {
	// Get the organization contract instance
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return nil, err
	}

	// Get total asset count
	assetCount, err := orgContractInstance.GetAssetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	var creatorAssets []*big.Int
	creatorAddr := common.HexToAddress(creator)

	// Iterate through all assets to find ones created by the specified creator
	for i := int64(1); i <= assetCount.Int64(); i++ {
		assetID := big.NewInt(i)
		// Get asset details
		_, assetCreator, _, _, _, _, _, _, _, _, _, _, _, _, _, _, err := orgContractInstance.Assets(&bind.CallOpts{Context: ctx}, assetID)
		if err != nil {
			continue // Skip assets that don't exist or can't be read
		}

		if assetCreator == creatorAddr {
			creatorAssets = append(creatorAssets, assetID)
		}
	}

	return creatorAssets, nil
}

// GetAssetsByOwner returns all assets owned by a specific owner
func (s *AssetService) GetAssetsByOwner(ctx context.Context, orgContract, owner string) ([]*big.Int, error) {
	// Get the organization contract instance
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return nil, err
	}

	// Try to get asset count (this might not exist in all contracts)
	assetCount, err := orgContractInstance.GetAssetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		// If GetAssetCount doesn't exist, check a reasonable number of assets
		assetCount = big.NewInt(100)
	}

	var ownerAssets []*big.Int
	ownerAddr := common.HexToAddress(owner)

	// Iterate through all assets to find ones owned by the specified owner
	for i := int64(1); i <= assetCount.Int64() && i <= 100; i++ { // Limit to 100 for safety
		assetID := big.NewInt(i)
		// Get asset details
		_, _, assetOwner, _, _, _, _, _, _, _, _, _, _, _, _, _, err := orgContractInstance.Assets(&bind.CallOpts{Context: ctx}, assetID)
		if err != nil {
			continue // Skip assets that don't exist or can't be read
		}

		if assetOwner == ownerAddr {
			ownerAssets = append(ownerAssets, assetID)
		}
	}

	return ownerAssets, nil
}

// TransferAssetOwnership transfers ownership of an asset
func (s *AssetService) TransferAssetOwnership(ctx context.Context, orgContract string, assetID *big.Int, newOwner, caller string) (*types.TransactionResult, error) {
	// Get the organization contract instance
	orgContractInstance, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgContract), s.client)
	if err != nil {
		return nil, err
	}

	// Update nonce before transaction
	nonce, err := s.client.PendingNonceAt(ctx, s.auth.From)
	if err != nil {
		return nil, err
	}
	s.auth.Nonce = big.NewInt(int64(nonce))

	// Transfer through organization contract
	tx, err := orgContractInstance.TransferAsset(s.auth, assetID, common.HexToAddress(newOwner))
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

func (s *AssetService) waitForTransaction(ctx context.Context, txHash common.Hash) (*types.TransactionResult, error) {
	// TODO: Implement transaction waiting and parsing
	// This is a placeholder implementation
	return &types.TransactionResult{
		TxHash: txHash,
		Status: 1,
	}, nil
}
