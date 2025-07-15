package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ardata-tech/hauska-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// GoChain network configuration
	rpcURL := "http://localhost:8545"
	chainID := int64(31337)
	targetAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	// Placeholder contract addresses (these would need to be deployed)
	factoryAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512" // Example address
	usdcAddress := "0x5FbDB2315678afecb367f032d93F642f64180aa3"    // Example address

	fmt.Printf("Hauska SDK GoChain Balance Example\n")
	fmt.Printf("==================================\n")
	fmt.Printf("RPC URL: %s\n", rpcURL)
	fmt.Printf("Chain ID: %d\n", chainID)
	fmt.Printf("Currency: GO\n")
	fmt.Printf("Target Address: %s\n", targetAddress)
	fmt.Printf("Factory Address: %s\n", factoryAddress)
	fmt.Printf("USDC Address: %s\n", usdcAddress)
	fmt.Printf("----------------------------------------\n")

	// Connect to the GoChain network
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the GoChain network: %v", err)
	}
	defer client.Close()

	// Create Hauska SDK configuration
	config := &hauska.Config{
		Client:         client,
		FactoryAddress: common.HexToAddress(factoryAddress),
		USDCAddress:    common.HexToAddress(usdcAddress),
		ChainID:        big.NewInt(chainID),
		GasLimit:       3000000,
		GasPrice:       big.NewInt(20000000000), // 20 gwei
	}

	fmt.Printf("✓ Connected to GoChain network\n")

	// For this example, we'll fetch the balance directly using the client
	// since we don't need the full Hauska SDK functionality for balance queries
	address := common.HexToAddress(targetAddress)

	// Fetch the balance
	fmt.Printf("\nFetching balance for address: %s\n", address.Hex())

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	// Convert balance from wei to GO
	balanceInGO := weiToGO(balance)

	fmt.Printf("Raw balance: %s wei\n", balance.String())
	fmt.Printf("Formatted balance: %s GO\n", balanceInGO)

	// Check if this is a funded account (typical for development)
	if balance.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("✓ Account has funds available\n")
	} else {
		fmt.Printf("⚠ Account has no funds\n")
	}

	// Get network information
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Warning: Failed to get network ID: %v", err)
	} else {
		fmt.Printf("Network ID: %d\n", networkID.Int64())
	}

	// Get latest block to verify network connectivity
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("Warning: Failed to get latest block: %v", err)
	} else {
		fmt.Printf("Latest block: %d\n", blockNumber)
	}

	// Example: Show how you would initialize the SDK (if contracts were deployed)
	fmt.Printf("\nHauska SDK Configuration:\n")
	fmt.Printf("Chain ID: %s\n", config.ChainID.String())
	fmt.Printf("Gas Limit: %d\n", config.GasLimit)
	fmt.Printf("Gas Price: %s wei\n", config.GasPrice.String())

	// Note: To actually use the Hauska SDK, you would need to:
	// 1. Deploy the Hauska contracts to your GoChain network
	// 2. Update the contract addresses in the config
	// 3. Set up authentication (private key or TransactOpts)
	//
	// Example:
	// sdk, err := hauska.NewSDK(config)
	// if err != nil {
	//     log.Fatalf("Failed to create SDK: %v", err)
	// }
}

// weiToGO converts wei to GO (equivalent to ETH conversion)
func weiToGO(wei *big.Int) string {
	// 1 GO = 10^18 wei (same as ETH)
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Create a float representation for better readability
	weiFloat := new(big.Float).SetInt(wei)
	divisorFloat := new(big.Float).SetInt(divisor)
	result := new(big.Float).Quo(weiFloat, divisorFloat)

	return result.Text('f', 6) // 6 decimal places
}
