package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/services"
	"github.com/ardata-tech/hauska-go/types"
)

// revenueClient implements the RevenueClient interface
type revenueClient struct {
	service *services.RevenueService
}

// NewRevenueClient creates a new revenue client
func NewRevenueClient(service *services.RevenueService) RevenueClient {
	return &revenueClient{
		service: service,
	}
}

// Revenue distribution
func (c *revenueClient) DistributeRevenue(ctx context.Context, from string, amount *big.Int, assetOwner, integrationPartner, orgContract string) (*types.TransactionResult, error) {
	return c.service.DistributeRevenue(ctx, from, amount, assetOwner, integrationPartner, orgContract)
}

func (c *revenueClient) GetRevenueStats(ctx context.Context, orgContract string) (*types.RevenueStats, error) {
	return c.service.GetRevenueStats(ctx, orgContract)
}

func (c *revenueClient) GetEarnings(ctx context.Context, account string) (*big.Int, error) {
	return c.service.GetEarnings(ctx, account)
}

// Fee management
func (c *revenueClient) SetCustomFees(ctx context.Context, orgContract string, hauskaFeePct, integratorFeePct uint32) (*types.TransactionResult, error) {
	return c.service.SetCustomFees(ctx, orgContract, hauskaFeePct, integratorFeePct)
}

func (c *revenueClient) RemoveCustomFees(ctx context.Context, orgContract string) (*types.TransactionResult, error) {
	return c.service.RemoveCustomFees(ctx, orgContract)
}

func (c *revenueClient) GetCustomFees(ctx context.Context, orgContract string) (*types.PlatformFees, error) {
	return c.service.GetCustomFees(ctx, orgContract)
}
