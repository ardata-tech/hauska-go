package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/services"
	"github.com/ardata-tech/hauska-go/types"
)

// groupClient implements the GroupClient interface
type groupClient struct {
	service *services.GroupService
}

// NewGroupClient creates a new group client
func NewGroupClient(service *services.GroupService) GroupClient {
	return &groupClient{
		service: service,
	}
}

// Group operations
func (c *groupClient) CreateAssetGroup(ctx context.Context, orgContract, groupName string, assetIDs []*big.Int, groupPrice *big.Int, creator string) (*types.TransactionResult, error) {
	return c.service.CreateGroup(ctx, orgContract, groupName, assetIDs, groupPrice, creator)
}

func (c *groupClient) GetAssetGroup(ctx context.Context, orgContract string, groupID *big.Int) (*types.AssetGroup, error) {
	return c.service.GetGroup(ctx, orgContract, groupID)
}

func (c *groupClient) GetGroupsByCreator(ctx context.Context, orgContract, creator string) ([]*big.Int, error) {
	return c.service.GetGroupsByCreator(ctx, orgContract, creator)
}

func (c *groupClient) AddAssetToGroup(ctx context.Context, orgContract string, groupID, assetID *big.Int, caller string) (*types.TransactionResult, error) {
	return c.service.AddAssetToGroup(ctx, orgContract, groupID, assetID, caller)
}

func (c *groupClient) RemoveAssetFromGroup(ctx context.Context, orgContract string, groupID, assetID *big.Int, caller string) (*types.TransactionResult, error) {
	return c.service.RemoveAssetFromGroup(ctx, orgContract, groupID, assetID, caller)
}

func (c *groupClient) UpdateGroupPrice(ctx context.Context, orgContract string, groupID *big.Int, newPrice *big.Int, caller string) (*types.TransactionResult, error) {
	return c.service.UpdateGroupPrice(ctx, orgContract, groupID, newPrice, caller)
}

func (c *groupClient) RemoveAssetGroup(ctx context.Context, orgContract string, groupID *big.Int, caller string) (*types.TransactionResult, error) {
	return c.service.RemoveGroup(ctx, orgContract, groupID, caller)
}

func (c *groupClient) GetGroupCount(ctx context.Context, orgContract string) (*big.Int, error) {
	return c.service.GetGroupCount(ctx, orgContract)
}

// Alias methods and helpers
func (c *groupClient) GetGroup(ctx context.Context, orgContract string, groupID *big.Int) (*types.AssetGroup, error) {
	return c.service.GetGroup(ctx, orgContract, groupID)
}

func (c *groupClient) IsGroupActive(ctx context.Context, orgContract string, groupID *big.Int) (bool, error) {
	return c.service.IsGroupActive(ctx, orgContract, groupID)
}

func (c *groupClient) GetGroupAssets(ctx context.Context, orgContract string, groupID *big.Int) ([]*big.Int, error) {
	return c.service.GetGroupAssets(ctx, orgContract, groupID)
}

func (c *groupClient) CalculateGroupLicenseFee(ctx context.Context, orgContract string, groupID *big.Int) (*big.Int, error) {
	return c.service.CalculateGroupLicenseFee(ctx, orgContract, groupID)
}
