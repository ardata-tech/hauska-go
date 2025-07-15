package hauska

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Config holds the configuration for the Hauska SDK
type Config struct {
	// Ethereum client connection
	Client *ethclient.Client

	// Contract addresses
	FactoryAddress common.Address
	USDCAddress    common.Address

	// Authentication
	PrivateKey string
	Auth       *bind.TransactOpts

	// Gas settings
	GasLimit             uint64
	GasPrice             *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int

	// Network settings
	ChainID *big.Int

	// Optional: Custom module addresses (if using FactoryLite)
	ModuleAddresses *ModuleAddresses
}

// ModuleAddresses holds addresses for all Hauska modules
type ModuleAddresses struct {
	LicenseManager     common.Address
	AssetRegistry      common.Address
	GroupManager       common.Address
	RevenueDistributor common.Address
	AssetNFT           common.Address
}

// DefaultConfig returns a default configuration for local development
func DefaultConfig() *Config {
	return &Config{
		GasLimit:             3000000,
		GasPrice:             big.NewInt(20000000000), // 20 gwei
		MaxFeePerGas:         big.NewInt(50000000000), // 50 gwei
		MaxPriorityFeePerGas: big.NewInt(2000000000),  // 2 gwei
		ChainID:              big.NewInt(31337),       // Hardhat local network
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Client == nil {
		return ErrInvalidClient
	}

	if c.FactoryAddress == (common.Address{}) {
		return ErrInvalidFactoryAddress
	}

	if c.USDCAddress == (common.Address{}) {
		return ErrInvalidUSDCAddress
	}

	if c.Auth == nil && c.PrivateKey == "" {
		return ErrMissingAuth
	}

	return nil
}

// TransactionOptions returns transaction options based on config
func (c *Config) TransactionOptions() *bind.TransactOpts {
	if c.Auth != nil {
		return c.Auth
	}

	// Create auth from private key if provided
	// This would be implemented when we add actual contract integration
	return &bind.TransactOpts{
		GasLimit: c.GasLimit,
		GasPrice: c.GasPrice,
	}
}

// CallOptions returns call options for read operations
func (c *Config) CallOptions() *bind.CallOpts {
	return &bind.CallOpts{
		// Add any default call options
	}
}

// WaitForTransaction waits for a transaction to be mined
func (c *Config) WaitForTransaction(tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(nil, c.Client, tx)
}
