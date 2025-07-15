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

// GroupService implements group management operations
type GroupService struct {
	client       *ethclient.Client
	groupManager *contracts.HauskaGroupManager
	auth         *bind.TransactOpts
}

// NewGroupService creates a new group service instance
func NewGroupService(client *ethclient.Client, groupManagerAddr common.Address, auth *bind.TransactOpts) (*GroupService, error) {
	groupManager, err := contracts.NewHauskaGroupManager(groupManagerAddr, client)
	if err != nil {
		return nil, err
	}

	return &GroupService{
		client:       client,
		groupManager: groupManager,
		auth:         auth,
	}, nil
}

// CreateGroup creates a new asset group
func (s *GroupService) CreateGroup(ctx context.Context, orgContract, groupName string, assetIDs []*big.Int, groupPrice *big.Int, creator string) (*types.TransactionResult, error) {
	tx, err := s.groupManager.CreateGroup(s.auth, groupName, assetIDs, groupPrice, common.HexToAddress(creator))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetGroup returns group details by organization and group ID
func (s *GroupService) GetGroup(ctx context.Context, orgContract string, groupID *big.Int) (*types.AssetGroup, error) {
	group, err := s.groupManager.GetGroup(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), groupID)
	if err != nil {
		return nil, err
	}

	return &types.AssetGroup{
		GroupID:    group.GroupId,
		Name:       group.Name,
		Members:    group.Members,
		GroupPrice: group.GroupPrice,
		Owner:      group.Owner,
		IsActive:   true, // Groups returned from GetGroup are active
	}, nil
}

// GetGroupsByCreator returns all groups created by a specific creator
func (s *GroupService) GetGroupsByCreator(ctx context.Context, orgContract, creator string) ([]*big.Int, error) {
	return s.groupManager.GetGroupsByCreator(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), common.HexToAddress(creator))
}

// AddAssetToGroup adds an asset to an existing group
func (s *GroupService) AddAssetToGroup(ctx context.Context, orgContract string, groupID, assetID *big.Int, caller string) (*types.TransactionResult, error) {
	tx, err := s.groupManager.AddAssetToGroup(s.auth, groupID, assetID, common.HexToAddress(caller))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// RemoveAssetFromGroup removes an asset from a group
func (s *GroupService) RemoveAssetFromGroup(ctx context.Context, orgContract string, groupID, assetID *big.Int, caller string) (*types.TransactionResult, error) {
	tx, err := s.groupManager.RemoveAssetFromGroup(s.auth, groupID, assetID, common.HexToAddress(caller))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// UpdateGroupPrice updates the price of a group
func (s *GroupService) UpdateGroupPrice(ctx context.Context, orgContract string, groupID *big.Int, newPrice *big.Int, caller string) (*types.TransactionResult, error) {
	tx, err := s.groupManager.UpdateGroupPrice(s.auth, groupID, newPrice, common.HexToAddress(caller))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// RemoveGroup removes a group (makes it inactive)
func (s *GroupService) RemoveGroup(ctx context.Context, orgContract string, groupID *big.Int, caller string) (*types.TransactionResult, error) {
	tx, err := s.groupManager.RemoveGroup(s.auth, groupID, common.HexToAddress(caller))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetGroupCount returns the total number of groups for an organization
func (s *GroupService) GetGroupCount(ctx context.Context, orgContract string) (*big.Int, error) {
	return s.groupManager.GetGroupCount(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract))
}

// IsGroupActive checks if a group is active
func (s *GroupService) IsGroupActive(ctx context.Context, orgContract string, groupID *big.Int) (bool, error) {
	// Get group details and check if it exists
	group, err := s.GetGroup(ctx, orgContract, groupID)
	if err != nil {
		return false, err
	}
	return group.IsActive, nil
}

// GetGroupAssets returns all asset IDs in a group
func (s *GroupService) GetGroupAssets(ctx context.Context, orgContract string, groupID *big.Int) ([]*big.Int, error) {
	group, err := s.GetGroup(ctx, orgContract, groupID)
	if err != nil {
		return nil, err
	}
	return group.Members, nil
}

// CalculateGroupLicenseFee calculates the total fee for licensing a group
func (s *GroupService) CalculateGroupLicenseFee(ctx context.Context, orgContract string, groupID *big.Int) (*big.Int, error) {
	group, err := s.GetGroup(ctx, orgContract, groupID)
	if err != nil {
		return nil, err
	}
	return group.GroupPrice, nil
}

// waitForTransaction waits for a transaction to be confirmed and returns the result
func (s *GroupService) waitForTransaction(ctx context.Context, txHash common.Hash) (*types.TransactionResult, error) {
	receipt, err := s.client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	return &types.TransactionResult{
		TxHash:      txHash,
		BlockNumber: receipt.BlockNumber,
		GasUsed:     receipt.GasUsed,
		Status:      uint64(receipt.Status),
	}, nil
}
