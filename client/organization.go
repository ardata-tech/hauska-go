package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/services"
	"github.com/ardata-tech/hauska-go/types"
)

// organizationClient implements the OrganizationClient interface
type organizationClient struct {
	service *services.OrganizationService
}

// NewOrganizationClient creates a new organization client
func NewOrganizationClient(service *services.OrganizationService) OrganizationClient {
	return &organizationClient{
		service: service,
	}
}

// Creator management
func (c *organizationClient) AddCreator(ctx context.Context, creatorAddress string) (*types.TransactionResult, error) {
	return c.service.AddCreator(ctx, creatorAddress)
}

func (c *organizationClient) RemoveCreator(ctx context.Context, creatorAddress string) (*types.TransactionResult, error) {
	return c.service.RemoveCreator(ctx, creatorAddress)
}

func (c *organizationClient) GetCreators(ctx context.Context) ([]string, error) {
	return c.service.GetCreators(ctx)
}

func (c *organizationClient) IsCreator(ctx context.Context, address string) (bool, error) {
	return c.service.IsCreator(ctx, address)
}

// Asset management
func (c *organizationClient) CreateAsset(ctx context.Context, req *types.CreateAssetRequest) (*types.TransactionResult, error) {
	return c.service.CreateAsset(ctx, req)
}

func (c *organizationClient) GetAsset(ctx context.Context, assetID *big.Int) (*types.VerifiedDigitalAsset, error) {
	return c.service.GetAsset(ctx, assetID)
}

func (c *organizationClient) GetAssetsByOwner(ctx context.Context, owner string) ([]*big.Int, error) {
	return c.service.GetAssetsByOwner(ctx, owner)
}

func (c *organizationClient) GetAssetsByCreator(ctx context.Context, creator string) ([]*big.Int, error) {
	return c.service.GetAssetsByCreator(ctx, creator)
}

func (c *organizationClient) VerifyAsset(ctx context.Context, assetID *big.Int) (*types.TransactionResult, error) {
	return c.service.VerifyAsset(ctx, assetID)
}

func (c *organizationClient) UpdateAsset(ctx context.Context, assetID *big.Int, updates *types.CreateAssetRequest) (*types.TransactionResult, error) {
	return c.service.UpdateAsset(ctx, assetID, updates)
}

func (c *organizationClient) TransferAsset(ctx context.Context, assetID *big.Int, newOwner string) (*types.TransactionResult, error) {
	return c.service.TransferAsset(ctx, assetID, newOwner)
}

// License management
func (c *organizationClient) LicenseAsset(ctx context.Context, req *types.LicenseAssetRequest) (*types.TransactionResult, error) {
	return c.service.LicenseAsset(ctx, req)
}

func (c *organizationClient) LicenseGroup(ctx context.Context, groupID *big.Int, licensee string, permissions []types.LicensePermissions, resellerFee *big.Int) (*types.TransactionResult, error) {
	return c.service.LicenseGroup(ctx, groupID, licensee, permissions, resellerFee)
}

func (c *organizationClient) GetUserLicenses(ctx context.Context, user string) ([]*big.Int, error) {
	return c.service.GetUserLicenses(ctx, user)
}

func (c *organizationClient) IsAssetLicensedBy(ctx context.Context, assetID *big.Int, user string) (bool, error) {
	return c.service.IsAssetLicensedBy(ctx, assetID, user)
}

// Group management
func (c *organizationClient) CreateGroup(ctx context.Context, req *types.CreateGroupRequest) (*types.TransactionResult, error) {
	return c.service.CreateGroup(ctx, req)
}

func (c *organizationClient) GetGroup(ctx context.Context, groupID *big.Int) (*types.AssetGroup, error) {
	return c.service.GetGroup(ctx, groupID)
}

func (c *organizationClient) GetGroupsByOwner(ctx context.Context, owner string) ([]*big.Int, error) {
	return c.service.GetGroupsByOwner(ctx, owner)
}

func (c *organizationClient) AddAssetToGroup(ctx context.Context, groupID, assetID *big.Int) (*types.TransactionResult, error) {
	return c.service.AddAssetToGroup(ctx, groupID, assetID)
}

func (c *organizationClient) RemoveAssetFromGroup(ctx context.Context, groupID, assetID *big.Int) (*types.TransactionResult, error) {
	return c.service.RemoveAssetFromGroup(ctx, groupID, assetID)
}

func (c *organizationClient) UpdateGroupPrice(ctx context.Context, groupID *big.Int, newPrice *big.Int) (*types.TransactionResult, error) {
	return c.service.UpdateGroupPrice(ctx, groupID, newPrice)
}

func (c *organizationClient) RemoveGroup(ctx context.Context, groupID *big.Int) (*types.TransactionResult, error) {
	return c.service.RemoveGroup(ctx, groupID)
}

// Information
func (c *organizationClient) GetInfo(ctx context.Context) (*types.Organization, error) {
	return c.service.GetInfo(ctx)
}

func (c *organizationClient) GetAssetCount(ctx context.Context) (*big.Int, error) {
	return c.service.GetAssetCount(ctx)
}

func (c *organizationClient) GetCreatorCount(ctx context.Context) (*big.Int, error) {
	return c.service.GetCreatorCount(ctx)
}

func (c *organizationClient) GetGroupCount(ctx context.Context) (*big.Int, error) {
	return c.service.GetGroupCount(ctx)
}
