# Hauska Go SDK

A Go SDK for interacting with Hauska smart contracts, providing idiomatic Go interfaces for blockchain-based digital asset management and licensing.

---

## ✅ What's Complete

* Full package structure and modular organization
* All core types, interfaces, and service definitions
* High-level client API design
* Error definitions and usage examples
* **Contract bindings generated using `abigen`**
* Automated build script with ABI extraction

## 🔄 What's Pending

* Full implementation of service methods using contract bindings
* Network integration and testing
* Documentation and unit tests

---

## 📁 Project Structure

```bash
sdk/
├── go.mod
├── sdk.go                 # Main SDK entrypoint
├── config.go              # Config definitions
├── errors.go              # Error constants
├── types/                 # Core domain types
├── services/              # Contract logic layer
├── client/                # High-level interface layer
├── contracts/             # Contract bindings (generated)
├── utils/                 # Utility helpers
└── examples/              # Sample usage patterns
```

**Contract Bindings**: The `contracts/` directory contains auto-generated Go bindings for all Hauska smart contracts, created using `abigen` and the automated build script.

---

## 🏗️ Architecture Overview

The Hauska Go SDK follows a **layered architecture** that provides clean separation of concerns and intuitive APIs for blockchain interaction:

```text
┌─────────────────────────────────────┐
│           SDK (main.go)             │  ← Main entry point
├─────────────────────────────────────┤
│         Clients (client/)           │  ← High-level API layer
├─────────────────────────────────────┤
│        Services (services/)         │  ← Business logic layer
├─────────────────────────────────────┤
│       Contracts (contracts/)        │  ← Auto-generated bindings
├─────────────────────────────────────┤
│        Ethereum Blockchain          │  ← Smart contracts
└─────────────────────────────────────┘
```

### Component Responsibilities

* **`types/`**: Core data structures, domain models, and shared interfaces
* **`contracts/`**: Auto-generated contract bindings (via `abigen`) for direct blockchain interaction
* **`services/`**: Business logic layer handling transaction management, validation, and error handling
* **`client/`**: High-level, user-friendly API abstractions over service operations
* **`utils/`**: Utility functions for address validation, amount conversions, and common operations
* **`config.go`**: Centralized configuration with Ethereum client setup and contract addresses

### Communication Flow

```go
// User calls SDK
sdk.Factory.CreateOrganization(ctx, principal, partner)
       ↓
// Client delegates to service
factoryClient.service.CreateOrganization(ctx, principal, partner)
       ↓
// Service calls contract binding
factoryContract.CreateContract(opts, principal, partner)
       ↓
// Contract binding makes blockchain call
ethereum.Call() → Smart Contract on blockchain
```

This design provides **separation of concerns**, **testability**, and **extensibility** while maintaining clean, intuitive APIs for developers.

---

## 🔧 Configuration

```go
config := &hauska.Config{
    Client:         ethClient,                     // Ethereum client
    FactoryAddress: "0xFactoryContractAddress",    // Deployed Factory contract
    USDCAddress:    "0xUSDCContractAddress",       // Deployed USDC ERC20
    PrivateKey:     "...",                         // User's private key
    ChainID:        1,                             // Ethereum mainnet
    GasLimit:       300000,
    GasPrice:       big.NewInt(20000000000),
}
```

---

## 🛠️ Installation

```bash
go get github.com/ardata-tech/hauska-go
```

---

## 🚀 Quick Start

```go
package main

import (
    "context"
    "log"
    "github.com/ardata-tech/hauska-go"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, _ := ethclient.Dial("http://localhost:8545")
    
    config := &hauska.Config{
        Client:         client,
        FactoryAddress: "0x...",
        USDCAddress:    "0x...",
        PrivateKey:     "...",
    }

    sdk, err := hauska.New(config)
    if err != nil {
        log.Fatal(err)
    }

    // Example call
    orgs, err := sdk.Clients.Factory.GetAllOrganizations(context.Background())
    if err != nil {
        log.Printf("Expected (scaffold): %v", err)
    }
}
```

---

## ❗ Error Handling

Defined in `errors.go`, the SDK includes:

* `ErrInvalidClient`, `ErrMissingAuth` – misconfigured SDK
* `ErrContractNotFound`, `ErrTransactionFailed` – blockchain issues
* `ErrAssetNotFound`, `ErrLicenseNotFound` – logical not-found errors
