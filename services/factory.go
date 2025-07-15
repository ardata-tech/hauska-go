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

// FactoryService implements factory contract operations
type FactoryService struct {
	client  *ethclient.Client
	factory *contracts.HauskaContractFactory
	auth    *bind.TransactOpts
}

// NewFactoryService creates a new factory service instance
func NewFactoryService(client *ethclient.Client, factoryAddr common.Address, auth *bind.TransactOpts) (*FactoryService, error) {
	factory, err := contracts.NewHauskaContractFactory(factoryAddr, client)
	if err != nil {
		return nil, err
	}

	return &FactoryService{
		client:  client,
		factory: factory,
		auth:    auth,
	}, nil
}

// CreateOrganization creates a new organization contract
func (s *FactoryService) CreateOrganization(ctx context.Context, principal, integrationPartner string) (*types.TransactionResult, error) {
	tx, err := s.factory.CreateContract(s.auth, common.HexToAddress(principal), common.HexToAddress(integrationPartner))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAllOrganizations returns all organization contract addresses
func (s *FactoryService) GetAllOrganizations(ctx context.Context) ([]string, error) {
	contracts, err := s.factory.GetAllContracts(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	addresses := make([]string, len(contracts))
	for i, addr := range contracts {
		addresses[i] = addr.Hex()
	}

	return addresses, nil
}

// GetOrganization returns organization information by principal address
func (s *FactoryService) GetOrganization(ctx context.Context, principal string) (*types.Organization, error) {
	orgAddr, err := s.factory.GetContract(&bind.CallOpts{Context: ctx}, common.HexToAddress(principal))
	if err != nil {
		return nil, err
	}

	// Create organization contract instance to fetch details
	orgContract, err := contracts.NewHauskaOrgContract(orgAddr, s.client)
	if err != nil {
		return nil, err
	}

	// Fetch organization details
	integrationPartner, err := orgContract.IntegrationPartner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	factory, err := orgContract.Factory(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	licenseManager, err := orgContract.LicenseManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	assetRegistry, err := orgContract.AssetRegistry(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	groupManager, err := orgContract.GroupManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	revenueDistributor, err := orgContract.RevenueDistributor(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	assetCount, err := orgContract.GetAssetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	creatorCount, err := orgContract.GetCreatorCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	// Fetch creators by iterating through creator count
	creators := make([]common.Address, creatorCount.Int64())
	for i := int64(0); i < creatorCount.Int64(); i++ {
		creator, err := orgContract.Creators(&bind.CallOpts{Context: ctx}, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		creators[i] = creator
	}

	// For now, set groupCount to 0 as GetGroupCount is not available
	groupCount := big.NewInt(0)

	return &types.Organization{
		Address:            orgAddr,
		Principal:          common.HexToAddress(principal),
		IntegrationPartner: integrationPartner,
		Factory:            factory,
		LicenseManager:     licenseManager,
		AssetRegistry:      assetRegistry,
		GroupManager:       groupManager,
		RevenueDistributor: revenueDistributor,
		Creators:           creators,
		AssetCount:         assetCount,
		CreatorCount:       creatorCount,
		GroupCount:         groupCount,
	}, nil
}

// RemoveOrganization removes an organization contract
func (s *FactoryService) RemoveOrganization(ctx context.Context, orgAddress string) (*types.TransactionResult, error) {
	tx, err := s.factory.RemoveContract(s.auth, common.HexToAddress(orgAddress))
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetPlatformFees returns current platform fee percentages
func (s *FactoryService) GetPlatformFees(ctx context.Context) (*types.PlatformFees, error) {
	fees, err := s.factory.GetPlatformFees(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return &types.PlatformFees{
		HauskaFeePct:     fees.HauskaFee,
		IntegratorFeePct: fees.IntegratorFee,
	}, nil
}

// UpdatePlatformFees updates platform fee percentages
func (s *FactoryService) UpdatePlatformFees(ctx context.Context, hauskaFeePct, integratorFeePct uint32) (*types.TransactionResult, error) {
	tx, err := s.factory.UpdatePlatformFees(s.auth, hauskaFeePct, integratorFeePct)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetModules returns addresses of all deployed modules
func (s *FactoryService) GetModules(ctx context.Context) (*types.ModuleAddresses, error) {
	modules, err := s.factory.GetModules(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return &types.ModuleAddresses{
		LicenseManager:     modules.LicenseManager,
		AssetRegistry:      modules.AssetRegistry,
		GroupManager:       modules.GroupManager,
		RevenueDistributor: modules.RevenueDistributor,
	}, nil
}

// IsValidOrgContract checks if an address is a valid organization contract
func (s *FactoryService) IsValidOrgContract(ctx context.Context, address string) (bool, error) {
	return s.factory.ValidContracts(&bind.CallOpts{Context: ctx}, common.HexToAddress(address))
}

// GetNumberOfOrganizations returns total number of organizations
func (s *FactoryService) GetNumberOfOrganizations(ctx context.Context) (*big.Int, error) {
	return s.factory.GetNumberOfOrganizations(&bind.CallOpts{Context: ctx})
}

// GetNumberOfCreators returns total number of creators in an organization
func (s *FactoryService) GetNumberOfCreators(ctx context.Context, orgAddress string) (*big.Int, error) {
	orgContract, err := contracts.NewHauskaOrgContract(common.HexToAddress(orgAddress), s.client)
	if err != nil {
		return nil, err
	}

	return orgContract.GetCreatorCount(&bind.CallOpts{Context: ctx})
}

// Helper functions

func (s *FactoryService) waitForTransaction(ctx context.Context, txHash common.Hash) (*types.TransactionResult, error) {
	// TODO: Implement transaction waiting and parsing
	// This is a placeholder implementation
	return &types.TransactionResult{
		TxHash: txHash,
		Status: 1,
	}, nil
}
