package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/services"
	"github.com/ardata-tech/hauska-go/types"
)

// assetClient implements the AssetClient interface
type assetClient struct {
	service *services.AssetService
}

// NewAssetClient creates a new asset client
func NewAssetClient(service *services.AssetService) AssetClient {
	return &assetClient{
		service: service,
	}
}

// Asset registry operations
func (c *assetClient) RegisterAsset(ctx context.Context, orgContract string, asset *types.VerifiedDigitalAsset, creator string) (*types.TransactionResult, error) {
	return c.service.RegisterAsset(ctx, orgContract, asset, creator)
}

func (c *assetClient) GetRegisteredAsset(ctx context.Context, orgContract string, assetID *big.Int) (*types.VerifiedDigitalAsset, error) {
	return c.service.GetAsset(ctx, orgContract, assetID)
}

func (c *assetClient) IsAssetVerified(ctx context.Context, orgContract string, assetID *big.Int) (bool, error) {
	return c.service.IsAssetVerified(ctx, orgContract, assetID)
}

func (c *assetClient) GetAssetsByCreator(ctx context.Context, orgContract, creator string) ([]*big.Int, error) {
	return c.service.GetAssetsByCreator(ctx, orgContract, creator)
}

func (c *assetClient) TransferAssetOwnership(ctx context.Context, orgContract string, assetID *big.Int, newOwner, caller string) (*types.TransactionResult, error) {
	return c.service.TransferAssetOwnership(ctx, orgContract, assetID, newOwner, caller)
}

// Cross-organization operations
func (c *assetClient) TransferAssetCrossOrg(ctx context.Context, fromOrg, toOrg string, assetID *big.Int, newOwner string) (*types.TransactionResult, error) {
	return c.service.TransferAssetCrossOrg(ctx, fromOrg, toOrg, assetID, newOwner)
}

// Asset NFT operations
func (c *assetClient) GetAssetNFTTokenID(ctx context.Context, orgContract string, assetID *big.Int) (*big.Int, error) {
	return c.service.GetAssetNFTTokenID(ctx, orgContract, assetID)
}

func (c *assetClient) GetAssetNFTOwner(ctx context.Context, tokenID *big.Int) (string, error) {
	// TODO: This should delegate to NFT service
	return "", types.ErrNotImplemented
}

func (c *assetClient) TransferAssetNFT(ctx context.Context, from, to string, tokenID *big.Int) (*types.TransactionResult, error) {
	return c.service.TransferAssetNFT(ctx, from, to, tokenID)
}
