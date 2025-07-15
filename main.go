package hauska

import (
	"context"

	"github.com/ardata-tech/hauska-go/client"
	"github.com/ardata-tech/hauska-go/services"
	"github.com/ethereum/go-ethereum/common"
)

// SDK is the main entry point for the Hauska Go SDK
type SDK struct {
	config *Config

	// High-level clients
	Factory      client.FactoryClient
	Organization client.OrganizationClient
	License      client.LicenseClient
	Asset        client.AssetClient
	Group        client.GroupClient
	Revenue      client.RevenueClient
	USDC         client.USDCClient

	// Low-level services (for advanced usage)
	Services *Services
}

// Services holds all low-level service implementations
type Services struct {
	Factory      *services.FactoryService
	Organization *services.OrganizationService
	License      *services.LicenseService
	Asset        *services.AssetService
	Group        *services.GroupService
	Revenue      *services.RevenueService
	NFT          *services.NFTService
	USDC         *services.USDCService
}

// NewSDK creates a new instance of the Hauska SDK
func NewSDK(config *Config) (*SDK, error) {
	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Initialize services with proper parameters
	factoryService, err := services.NewFactoryService(config.Client, config.FactoryAddress, config.Auth)
	if err != nil {
		return nil, err
	}

	usdcService, err := services.NewUSDCService(config.Client, config.USDCAddress, config.Auth)
	if err != nil {
		return nil, err
	}

	// Get module addresses from factory if not provided
	moduleAddresses := config.ModuleAddresses
	if moduleAddresses == nil {
		modules, err := factoryService.GetModules(context.Background())
		if err != nil {
			return nil, err
		}
		moduleAddresses = &ModuleAddresses{
			LicenseManager:     modules.LicenseManager,
			AssetRegistry:      modules.AssetRegistry,
			GroupManager:       modules.GroupManager,
			RevenueDistributor: modules.RevenueDistributor,
			AssetNFT:           modules.AssetNFT,
		}
	}

	licenseService, err := services.NewLicenseService(config.Client, moduleAddresses.LicenseManager, config.Auth)
	if err != nil {
		return nil, err
	}

	assetService, err := services.NewAssetService(config.Client, moduleAddresses.AssetRegistry, moduleAddresses.AssetNFT, config.Auth)
	if err != nil {
		return nil, err
	}

	nftService, err := services.NewNFTService(config.Client, moduleAddresses.AssetNFT, config.Auth)
	if err != nil {
		return nil, err
	}

	// Initialize placeholder services (these need proper implementation later)
	groupService := services.NewGroupService()
	revenueService := services.NewRevenueService()

	// Note: OrganizationService requires an organization address,
	// so it will be created when needed via GetOrganization() method

	// Create high-level clients
	factoryClient := client.NewFactoryClient(factoryService)
	licenseClient := client.NewLicenseClient(licenseService)
	assetClient := client.NewAssetClient(assetService)
	groupClient := client.NewGroupClient(groupService)
	revenueClient := client.NewRevenueClient(revenueService)
	usdcClient := client.NewUSDCClient(usdcService)

	// Create SDK instance
	sdk := &SDK{
		config: config,

		// High-level clients
		Factory: factoryClient,
		License: licenseClient,
		Asset:   assetClient,
		Group:   groupClient,
		Revenue: revenueClient,
		USDC:    usdcClient,

		// Low-level services
		Services: &Services{
			Factory: factoryService,
			License: licenseService,
			Asset:   assetService,
			Group:   groupService,
			Revenue: revenueService,
			NFT:     nftService,
			USDC:    usdcService,
		},
	}

	return sdk, nil
}

// NewOrganizationClient creates a new organization client for a specific organization
func (sdk *SDK) NewOrganizationClient(orgAddress string) (client.OrganizationClient, error) {
	orgService, err := services.NewOrganizationService(sdk.config.Client, common.HexToAddress(orgAddress), sdk.config.Auth)
	if err != nil {
		return nil, err
	}
	sdk.Services.Organization = orgService
	return client.NewOrganizationClient(orgService), nil
}

// GetConfig returns the SDK configuration
func (sdk *SDK) GetConfig() *Config {
	return sdk.config
}

// Close cleans up SDK resources
func (sdk *SDK) Close() error {
	// TODO: Close connections, cleanup resources
	return nil
}

// Health checks the health of all SDK components
func (sdk *SDK) Health(ctx context.Context) error {
	// TODO: Implement health checks
	// 1. Check Ethereum client connection
	// 2. Verify contract addresses are valid
	// 3. Check if contracts are deployed
	// 4. Validate module addresses

	return ErrNotImplemented
}
