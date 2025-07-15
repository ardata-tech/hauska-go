package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/services"
	"github.com/ardata-tech/hauska-go/types"
)

// usdcClient implements the USDCClient interface
type usdcClient struct {
	service *services.USDCService
}

// NewUSDCClient creates a new USDC client
func NewUSDCClient(service *services.USDCService) USDCClient {
	return &usdcClient{
		service: service,
	}
}

// Balance operations
func (c *usdcClient) GetBalance(ctx context.Context, account string) (*big.Int, error) {
	return c.service.GetBalance(ctx, account)
}

func (c *usdcClient) Transfer(ctx context.Context, to string, amount *big.Int) (*types.TransactionResult, error) {
	return c.service.Transfer(ctx, to, amount)
}

// Allowance operations
func (c *usdcClient) Approve(ctx context.Context, spender string, amount *big.Int) (*types.TransactionResult, error) {
	return c.service.Approve(ctx, spender, amount)
}

func (c *usdcClient) GetAllowance(ctx context.Context, owner, spender string) (*big.Int, error) {
	return c.service.GetAllowance(ctx, owner, spender)
}

// Token information
func (c *usdcClient) GetDecimals(ctx context.Context) (uint8, error) {
	return c.service.GetDecimals(ctx)
}

func (c *usdcClient) GetSymbol(ctx context.Context) (string, error) {
	return c.service.GetSymbol(ctx)
}

func (c *usdcClient) GetName(ctx context.Context) (string, error) {
	return c.service.GetName(ctx)
}

func (c *usdcClient) GetTotalSupply(ctx context.Context) (*big.Int, error) {
	return c.service.GetTotalSupply(ctx)
}

// Utility functions
func (c *usdcClient) FormatAmount(amount *big.Int) string {
	return c.service.FormatAmount(amount)
}

func (c *usdcClient) ParseAmount(amount string) (*big.Int, error) {
	return c.service.ParseAmount(amount)
}
