package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/services"
	"github.com/ardata-tech/hauska-go/types"
)

// factoryClient implements the FactoryClient interface
type factoryClient struct {
	service *services.FactoryService
}

// NewFactoryClient creates a new factory client
func NewFactoryClient(service *services.FactoryService) FactoryClient {
	return &factoryClient{
		service: service,
	}
}

// Organization management
func (c *factoryClient) CreateOrganization(ctx context.Context, principal, integrationPartner string) (*types.TransactionResult, error) {
	return c.service.CreateOrganization(ctx, principal, integrationPartner)
}

func (c *factoryClient) GetAllOrganizations(ctx context.Context) ([]string, error) {
	return c.service.GetAllOrganizations(ctx)
}

func (c *factoryClient) GetOrganization(ctx context.Context, principal string) (*types.Organization, error) {
	return c.service.GetOrganization(ctx, principal)
}

func (c *factoryClient) RemoveOrganization(ctx context.Context, orgAddress string) (*types.TransactionResult, error) {
	return c.service.RemoveOrganization(ctx, orgAddress)
}

// Platform management
func (c *factoryClient) GetPlatformFees(ctx context.Context) (*types.PlatformFees, error) {
	return c.service.GetPlatformFees(ctx)
}

func (c *factoryClient) UpdatePlatformFees(ctx context.Context, hauskaFeePct, integratorFeePct uint32) (*types.TransactionResult, error) {
	return c.service.UpdatePlatformFees(ctx, hauskaFeePct, integratorFeePct)
}

// Module information
func (c *factoryClient) GetModules(ctx context.Context) (*types.ModuleAddresses, error) {
	return c.service.GetModules(ctx)
}

func (c *factoryClient) IsValidOrgContract(ctx context.Context, address string) (bool, error) {
	return c.service.IsValidOrgContract(ctx, address)
}

// Statistics
func (c *factoryClient) GetNumberOfOrganizations(ctx context.Context) (*big.Int, error) {
	return c.service.GetNumberOfOrganizations(ctx)
}

func (c *factoryClient) GetNumberOfCreators(ctx context.Context, orgAddress string) (*big.Int, error) {
	return c.service.GetNumberOfCreators(ctx, orgAddress)
}
