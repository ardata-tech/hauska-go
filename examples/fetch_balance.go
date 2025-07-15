package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// GoChain network configuration
	rpcURL := "http://localhost:8545"
	chainID := int64(31337)
	targetAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	fmt.Printf("Connecting to GoChain network...\n")
	fmt.Printf("RPC URL: %s\n", rpcURL)
	fmt.Printf("Chain ID: %d\n", chainID)
	fmt.Printf("Currency: GO\n")
	fmt.Printf("Target Address: %s\n", targetAddress)
	fmt.Printf("----------------------------------------\n")

	// Connect to the GoChain network
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the GoChain network: %v", err)
	}
	defer client.Close()

	// Verify we're connected to the correct network
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	if networkID.Int64() != chainID {
		log.Printf("Warning: Connected to network ID %d, expected %d", networkID.Int64(), chainID)
	} else {
		fmt.Printf("✓ Successfully connected to network ID: %d\n", networkID.Int64())
	}

	// Convert string address to common.Address
	address := common.HexToAddress(targetAddress)

	// Fetch the balance
	fmt.Printf("\nFetching balance for address: %s\n", address.Hex())

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	// Convert balance from wei to GO (same as ETH, 18 decimals)
	balanceInGO := weiToGO(balance)

	fmt.Printf("Balance: %s wei\n", balance.String())
	fmt.Printf("Balance: %s GO\n", balanceInGO)

	// Get additional account information
	nonce, err := client.NonceAt(context.Background(), address, nil)
	if err != nil {
		log.Printf("Warning: Failed to get nonce: %v", err)
	} else {
		fmt.Printf("Nonce: %d\n", nonce)
	}

	// Get latest block number to show network is active
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("Warning: Failed to get latest block: %v", err)
	} else {
		fmt.Printf("Latest block: %d\n", blockNumber)
	}

	fmt.Printf("\n✓ Successfully fetched balance from GoChain network!\n")
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
