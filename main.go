package hauska

import (
	"context"

	"github.com/ardata-tech/hauska-go/client"
	"github.com/ardata-tech/hauska-go/services"
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

	// Initialize services
	factoryService := services.NewFactoryService()
	licenseService := services.NewLicenseService()
	assetService := services.NewAssetService()
	groupService := services.NewGroupService()
	revenueService := services.NewRevenueService()
	nftService := services.NewNFTService()
	usdcService := services.NewUSDCService()

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
func (sdk *SDK) NewOrganizationClient(orgAddress string) client.OrganizationClient {
	orgService := services.NewOrganizationService(orgAddress)
	sdk.Services.Organization = orgService
	return client.NewOrganizationClient(orgService)
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
