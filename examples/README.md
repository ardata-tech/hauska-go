# Hauska Go SDK Examples

This directory contains examples demonstrating how to use the Hauska Go SDK with different blockchain networks.

## GoChain Network Configuration

All examples are configured to work with a GoChain network using the following settings:

- **Network Name**: GoChain
- **RPC URL**: http://localhost:8545
- **Chain ID**: 31337
- **Currency**: GO
- **Test Address**: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

## Examples

### 1. Basic Balance Check (`basic_balance/`)

A simple example that demonstrates how to:
- Connect to a GoChain network
- Fetch the ETH (GO) balance of an address
- Convert wei to GO currency
- Display network information

**To run:**
```bash
cd examples/basic_balance
go run main.go
```

### 2. Hauska SDK Integration (`hauska_balance/`)

A more advanced example that shows how to:
- Set up the Hauska SDK configuration for GoChain
- Connect to the network using the SDK structure
- Fetch balance information
- Prepare for full Hauska contract integration

**To run:**
```bash
cd examples/hauska_balance
go run main.go
```

## Prerequisites

1. **Go Installation**: Make sure you have Go 1.23+ installed
2. **GoChain Node**: Ensure you have a GoChain node running on `http://localhost:8545`
3. **Dependencies**: Run `go mod tidy` in the project root to install dependencies

## Testing with Local Development

If you're running a local GoChain node or using a development network like Hardhat/Ganache configured for GoChain:

1. Make sure your node is running on the correct port (8545)
2. Ensure the chain ID matches (31337)
3. The test address should have some GO tokens for the balance check to show meaningful results

## Next Steps

To use the full Hauska SDK functionality:

1. Deploy the Hauska smart contracts to your GoChain network
2. Update the contract addresses in the SDK configuration
3. Set up proper authentication (private key or transaction options)
4. Use the SDK's high-level clients for asset management, licensing, etc.

## Common Issues

- **Connection refused**: Ensure your GoChain node is running on localhost:8545
- **Wrong network**: Verify the chain ID matches your network configuration
- **Zero balance**: The test address might not have any GO tokens

## Network Information

The examples use the following test address: `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266`

This is a common test address used in development environments. In a real GoChain network, you would use actual addresses with GO tokens.
