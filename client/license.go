package client

import (
	"context"
	"math/big"

	"github.com/ardata-tech/hauska-go/services"
	"github.com/ardata-tech/hauska-go/types"
)

// licenseClient implements the LicenseClient interface
type licenseClient struct {
	service *services.LicenseService
}

// NewLicenseClient creates a new license client
func NewLicenseClient(service *services.LicenseService) LicenseClient {
	return &licenseClient{
		service: service,
	}
}

// License operations
func (c *licenseClient) GetLicense(ctx context.Context, licenseID *big.Int) (*types.LicenseData, error) {
	return c.service.GetLicense(ctx, licenseID)
}

func (c *licenseClient) GetLicensesByUser(ctx context.Context, user string) ([]*big.Int, error) {
	return c.service.GetLicensesByUser(ctx, user)
}

func (c *licenseClient) IsLicenseValid(ctx context.Context, licenseID *big.Int) (bool, error) {
	return c.service.IsLicenseValid(ctx, licenseID)
}

func (c *licenseClient) RelicenseAsset(ctx context.Context, orgContract string, assetID *big.Int, licensee string, existingLicenseID *big.Int) (*types.TransactionResult, error) {
	// TODO: Add this method to LicenseService
	return nil, types.ErrNotImplemented
}

func (c *licenseClient) RenewLicense(ctx context.Context, licenseID *big.Int, additionalDuration *big.Int) (*types.TransactionResult, error) {
	// TODO: Add this method to LicenseService
	return nil, types.ErrNotImplemented
}

func (c *licenseClient) RevokeLicense(ctx context.Context, licenseID *big.Int) (*types.TransactionResult, error) {
	// TODO: Add this method to LicenseService
	return nil, types.ErrNotImplemented
}

// Batch operations
func (c *licenseClient) LicenseMultipleAssets(ctx context.Context, requests []*types.LicenseAssetRequest) ([]*types.TransactionResult, error) {
	// TODO: Add this method to LicenseService for batch operations
	return nil, types.ErrNotImplemented
}

// Query operations
func (c *licenseClient) GetAssetLicensees(ctx context.Context, orgContract string, assetID *big.Int) ([]string, error) {
	// TODO: Add this method to LicenseService
	return nil, types.ErrNotImplemented
}

func (c *licenseClient) GetLicenseHistory(ctx context.Context, assetID *big.Int) ([]*types.LicenseData, error) {
	// TODO: Add this method to LicenseService
	return nil, types.ErrNotImplemented
}
