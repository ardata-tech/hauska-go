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

// LicenseService implements license management operations
type LicenseService struct {
	client         *ethclient.Client
	licenseManager *contracts.HauskaLicenseManager
	auth           *bind.TransactOpts
}

// NewLicenseService creates a new license service instance
func NewLicenseService(client *ethclient.Client, licenseManagerAddr common.Address, auth *bind.TransactOpts) (*LicenseService, error) {
	licenseManager, err := contracts.NewHauskaLicenseManager(licenseManagerAddr, client)
	if err != nil {
		return nil, err
	}

	return &LicenseService{
		client:         client,
		licenseManager: licenseManager,
		auth:           auth,
	}, nil
}

// GetLicense returns license details by ID
func (s *LicenseService) GetLicense(ctx context.Context, licenseID *big.Int) (*types.LicenseData, error) {
	// TODO: Implement contract call
	// 1. Call licenseManager.GetLicenseDetails()
	// 2. Convert to LicenseData struct
	// 3. Return license details

	return nil, types.ErrNotImplemented
}

// GetLicensesByUser returns all licenses owned by a user
func (s *LicenseService) GetLicensesByUser(ctx context.Context, user string) ([]*big.Int, error) {
	// TODO: Implement contract call
	// 1. Call licenseManager.GetUserLicenses()
	// 2. Return license IDs

	return nil, types.ErrNotImplemented
}

// IsLicenseValid checks if a license is valid (active and not expired)
func (s *LicenseService) IsLicenseValid(ctx context.Context, licenseID *big.Int) (bool, error) {
	// TODO: Implement contract call
	// 1. Call licenseManager.IsLicenseValid()
	// 2. Return boolean result

	return false, types.ErrNotImplemented
}

// RelicenseAsset allows a license holder to create a sub-license
func (s *LicenseService) RelicenseAsset(ctx context.Context, orgContract string, assetID *big.Int, licensee string, existingLicenseID *big.Int) (*types.TransactionResult, error) {
	// TODO: Implement contract call
	// 1. Validate caller has resell permissions
	// 2. Check USDC allowance
	// 3. Call licenseManager.RelicenseAsset()
	// 4. Wait for confirmation
	// 5. Return transaction result

	return nil, types.ErrNotImplemented
}

// RenewLicense extends the duration of an existing license
func (s *LicenseService) RenewLicense(ctx context.Context, licenseID *big.Int, additionalDuration *big.Int) (*types.TransactionResult, error) {
	// TODO: Implement contract call
	// 1. Validate caller is license owner
	// 2. Check USDC allowance
	// 3. Call licenseManager.RenewLicense()
	// 4. Wait for confirmation
	// 5. Return transaction result

	return nil, types.ErrNotImplemented
}

// RevokeLicense revokes a license (admin only)
func (s *LicenseService) RevokeLicense(ctx context.Context, licenseID *big.Int) (*types.TransactionResult, error) {
	// TODO: Implement contract call
	// 1. Validate caller has admin role
	// 2. Call licenseManager.RevokeLicense()
	// 3. Wait for confirmation
	// 4. Return transaction result

	return nil, types.ErrNotImplemented
}

// LicenseMultipleAssets purchases licenses for multiple assets in batch
func (s *LicenseService) LicenseMultipleAssets(ctx context.Context, requests []*types.LicenseAssetRequest) ([]*types.TransactionResult, error) {
	// TODO: Implement batch operation
	// 1. Validate all requests
	// 2. Check total USDC allowance
	// 3. Execute licenses in sequence or parallel
	// 4. Return transaction results

	return nil, types.ErrNotImplemented
}

// GetAssetLicensees returns all addresses that have licensed an asset
func (s *LicenseService) GetAssetLicensees(ctx context.Context, orgContract string, assetID *big.Int) ([]string, error) {
	// TODO: Implement using events
	// 1. Query LicenseGranted events for asset
	// 2. Filter for active licenses
	// 3. Return licensee addresses

	return nil, types.ErrNotImplemented
}

// GetLicenseHistory returns the license history for an asset
func (s *LicenseService) GetLicenseHistory(ctx context.Context, assetID *big.Int) ([]*types.LicenseData, error) {
	// TODO: Implement using events
	// 1. Query all license-related events for asset
	// 2. Build history of licenses
	// 3. Return sorted by timestamp

	return nil, types.ErrNotImplemented
}
