package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ardata-tech/hauska-go"
	"github.com/ardata-tech/hauska-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Printf("Hauska SDK Asset Service Example\n")
	fmt.Printf("=================================\n")

	// Network configuration
	rpcURL := "http://localhost:8545"
	chainID := int64(31337)

	// Example contract addresses (replace with actual deployed addresses)
	factoryAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	usdcAddress := "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	organizationAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"

	// Example private key (replace with your own - DO NOT use in production)
	privateKeyHex := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	fmt.Printf("RPC URL: %s\n", rpcURL)
	fmt.Printf("Chain ID: %d\n", chainID)
	fmt.Printf("Organization Address: %s\n", organizationAddress)
	fmt.Printf("----------------------------------------\n")

	// Setup client and authentication
	client, auth, fromAddress := setupClientAndAuth(rpcURL, privateKeyHex, chainID)
	defer client.Close()

	fmt.Printf("From Address: %s\n", fromAddress.Hex())

	// Create Hauska SDK configuration
	config := &hauska.Config{
		Client:         client,
		FactoryAddress: common.HexToAddress(factoryAddress),
		USDCAddress:    common.HexToAddress(usdcAddress),
		Auth:           auth,
		ChainID:        big.NewInt(chainID),
		GasLimit:       3000000,
	}

	// Initialize SDK
	sdk, err := hauska.NewSDK(config)
	if err != nil {
		log.Fatalf("Failed to create SDK: %v", err)
	}

	fmt.Printf("✓ SDK initialized successfully\n\n")

	// Asset Service Examples
	ctx := context.Background()

	// Demonstrate the SDK is working
	demonstrateAssetService(sdk, ctx, organizationAddress, fromAddress)
}

func demonstrateAssetService(sdk *hauska.SDK, ctx context.Context, organizationAddress string, fromAddress common.Address) {
	fmt.Printf("Asset Service Operations:\n")
	fmt.Printf("========================\n")

	// Reference types package to avoid "unused import" error
	_ = types.FxPoolVDAS

	// Get SDK config for client and auth access
	config := sdk.GetConfig()

	// Verify what's deployed at the organization address
	fmt.Printf("📝 Verifying contract at organization address...\n")
	code, err := config.Client.CodeAt(ctx, common.HexToAddress(organizationAddress), nil)
	if err != nil {
		fmt.Printf("   ❌ Error getting contract code: %v\n", err)
	} else if len(code) == 0 {
		fmt.Printf("   ❌ No contract deployed at address %s\n", organizationAddress)
		fmt.Printf("   📝 Make sure you have deployed the organization contract first\n")
		return
	} else {
		fmt.Printf("   ✅ Contract found at address (code length: %d bytes)\n", len(code))
	}

	// 0. First check and add creator role if needed
	fmt.Printf("0. Checking and setting up creator role:\n")

	// Create organization client to manage roles
	orgClient, err := sdk.NewOrganizationClient(organizationAddress)
	if err != nil {
		fmt.Printf("   ❌ Error creating organization client: %v\n", err)
		return
	}

	// Check if user has creator role
	isCreator, err := orgClient.IsCreator(ctx, fromAddress.Hex())
	if err != nil {
		fmt.Printf("   ❌ Error checking creator role: %v\n", err)
	} else if !isCreator {
		fmt.Printf("   📝 Adding creator role to address: %s\n", fromAddress.Hex())

		// Update nonce before transaction
		nonce, err := config.Client.PendingNonceAt(ctx, config.Auth.From)
		if err != nil {
			fmt.Printf("   ❌ Error getting nonce: %v\n", err)
		} else {
			config.Auth.Nonce = big.NewInt(int64(nonce))

			result, err := orgClient.AddCreator(ctx, fromAddress.Hex())
			if err != nil {
				fmt.Printf("   ❌ Error adding creator role: %v\n", err)
				fmt.Printf("   📝 Note: You may need admin privileges to add creators\n")
			} else {
				fmt.Printf("   ✅ Creator role added! TX: %s\n", result.TxHash.Hex())
			}
		}
	} else {
		fmt.Printf("   ✅ Address already has creator role\n")
	}
	fmt.Println()

	// 1. Example: Asset Registration
	fmt.Printf("1. Asset registration example:\n")
	fmt.Printf("   Creating asset metadata...\n")

	// Get current timestamp for creation time
	currentTime := big.NewInt(time.Now().Unix())

	// Create asset metadata
	asset := &types.VerifiedDigitalAsset{
		AssetID:                big.NewInt(1),
		Creator:                fromAddress,
		Owner:                  fromAddress,
		Partner:                common.Address{},
		IPFSHash:               "QmYourAssetIPFSHash",
		MetadataHash:           "QmYourMetadataIPFSHash",
		AssetHash:              [32]byte{1, 2, 3}, // Your actual hash
		Version:                big.NewInt(1),
		IsVerified:             false,
		CreationTime:           currentTime,
		LastTransferTime:       currentTime,
		Price:                  big.NewInt(1000000), // 1 USDC
		Encrypted:              false,
		CanBeLicensed:          true,
		FxPool:                 types.FxPoolVDAS,
		EventTimestamp:         "2024-01-01T00:00:00Z",
		GeographicRestrictions: []types.CountryCode{},
	}

	fmt.Printf("   Registering asset...\n")

	// Update nonce before transaction
	nonce, err := config.Client.PendingNonceAt(ctx, config.Auth.From)
	if err != nil {
		fmt.Printf("   ❌ Error getting nonce: %v\n", err)
	} else {
		config.Auth.Nonce = big.NewInt(int64(nonce))

		result, err := sdk.Asset.RegisterAsset(ctx, organizationAddress, asset, fromAddress.Hex())
		if err != nil {
			fmt.Printf("   ❌ Error registering asset: %v\n", err)
			fmt.Printf("   📝 Tip: Make sure you have creator role in the organization\n")
		} else {
			fmt.Printf("   ✅ Asset registered! TX: %s\n", result.TxHash.Hex())
		}
	}
	fmt.Println()

	// 2. Example: Get asset details (will fail if no assets exist)
	fmt.Printf("2. Asset retrieval example:\n")
	assetID := big.NewInt(1)
	fmt.Printf("   Organization: %s\n", organizationAddress)
	fmt.Printf("   Asset ID: %s\n", assetID.String())

	retrievedAsset, err := sdk.Asset.GetAsset(ctx, organizationAddress, assetID)
	if err != nil {
		fmt.Printf("   ❌ Error getting asset (expected if no assets exist): %v\n", err)
	} else {
		fmt.Printf("   ✅ Asset found:\n")
		fmt.Printf("      Asset ID: %s\n", retrievedAsset.AssetID.String())
		fmt.Printf("      Creator: %s\n", retrievedAsset.Creator.Hex())
		fmt.Printf("      Owner: %s\n", retrievedAsset.Owner.Hex())
		fmt.Printf("      Price: %s wei\n", retrievedAsset.Price.String())
		fmt.Printf("      IPFS Hash: %s\n", retrievedAsset.IPFSHash)
		fmt.Printf("      Verified: %t\n", retrievedAsset.IsVerified)
		fmt.Printf("      Can be licensed: %t\n", retrievedAsset.CanBeLicensed)
	}
	fmt.Println()

	// 3. Example: Get assets by owner (using alternative approach)
	fmt.Printf("3. Getting assets by owner:\n")
	fmt.Printf("   Owner: %s\n", fromAddress.Hex())

	// Use asset enumeration to find owned assets
	fmt.Printf("   📝 Checking asset ownership through asset enumeration...\n")
	var ownerAssets []*big.Int

	// Check first 10 assets to find owned ones
	for i := int64(1); i <= 10; i++ {
		testAssetID := big.NewInt(i)
		asset, err := sdk.Asset.GetAsset(ctx, organizationAddress, testAssetID)
		if err == nil && asset.Owner == fromAddress {
			ownerAssets = append(ownerAssets, testAssetID)
		}
	}

	fmt.Printf("   ✅ Assets owned: %v\n", ownerAssets)
	if len(ownerAssets) == 0 {
		fmt.Printf("   📝 No assets found for this owner\n")
	}
	fmt.Println()

	// 4. Example: Get assets by creator
	fmt.Printf("4. Getting assets by creator:\n")
	fmt.Printf("   Creator: %s\n", fromAddress.Hex())

	creatorAssets, err := sdk.Asset.GetAssetsByCreator(ctx, organizationAddress, fromAddress.Hex())
	if err != nil {
		fmt.Printf("   ❌ Error getting creator assets: %v\n", err)
	} else {
		fmt.Printf("   ✅ Assets created: %v\n", creatorAssets)
		if len(creatorAssets) == 0 {
			fmt.Printf("   📝 No assets found for this creator\n")
		}
	}
	fmt.Println()

	// 5. Asset modification examples
	fmt.Printf("5. Asset modification examples:\n")

	// 5a. Check if asset is verified
	fmt.Printf("   📝 Checking asset verification status:\n")
	isVerified, err := sdk.Asset.IsAssetVerified(ctx, organizationAddress, assetID)
	if err != nil {
		fmt.Printf("   ❌ Error checking verification status: %v\n", err)
	} else {
		fmt.Printf("   ✅ Asset verification status: %t\n", isVerified)
	}
	fmt.Println()

	// 5b. Transfer asset ownership (with proper nonce handling)
	fmt.Printf("   📝 Transferring asset ownership:\n")
	newOwner := "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"

	// Update nonce before transaction
	nonce, err = config.Client.PendingNonceAt(ctx, config.Auth.From)
	if err != nil {
		fmt.Printf("   ❌ Error getting nonce: %v\n", err)
	} else {
		config.Auth.Nonce = big.NewInt(int64(nonce))

		transferResult, err := sdk.Asset.TransferAssetOwnership(ctx, organizationAddress, assetID, newOwner, fromAddress.Hex())
		if err != nil {
			fmt.Printf("   ❌ Error transferring asset: %v\n", err)
			fmt.Printf("   📝 Note: Asset may not exist or you may not be the owner\n")
		} else {
			fmt.Printf("   ✅ Asset ownership transferred! TX: %s\n", transferResult.TxHash.Hex())
			fmt.Printf("   New owner: %s\n", newOwner)
		}
	}
	fmt.Println()

	// 5c. Cross-organization transfer (example)
	fmt.Printf("   📝 Cross-organization transfer example (commented out):\n")
	fmt.Printf("   // targetOrg := \"0x...\"\n")
	fmt.Printf("   // result, err := sdk.Asset.TransferAssetCrossOrg(ctx, organizationAddress, targetOrg, assetID, newOwner)\n")
	fmt.Printf("   // This transfers an asset between different organizations\n")
	fmt.Println()

	// 5d. NFT operations (with error handling)
	fmt.Printf("   📝 Asset NFT operations:\n")
	tokenID, err := sdk.Asset.GetAssetNFTTokenID(ctx, organizationAddress, assetID)
	if err != nil {
		fmt.Printf("   ❌ Error getting NFT token ID: %v\n", err)
		fmt.Printf("   📝 Note: NFT contract may not be deployed or configured\n")
	} else {
		fmt.Printf("   ✅ Asset NFT Token ID: %s\n", tokenID.String())

		// Get NFT owner
		nftOwner, err := sdk.Asset.GetAssetNFTOwner(ctx, tokenID)
		if err != nil {
			fmt.Printf("   ❌ Error getting NFT owner: %v\n", err)
		} else {
			fmt.Printf("   ✅ NFT Owner: %s\n", nftOwner)
		}
	}
	fmt.Println()

	// Note about other operations
	fmt.Printf("   📝 Note: Asset verification, updates, and some other operations\n")
	fmt.Printf("   are available through the Organization client, not Asset client.\n")
	fmt.Printf("   These operations require organization-level permissions.\n")

	fmt.Printf("🎉 Asset Service Example Complete!\n")
}

func setupClientAndAuth(rpcURL, privateKeyHex string, chainID int64) (*ethclient.Client, *bind.TransactOpts, common.Address) {
	// Connect to the network
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the network: %v", err)
	}

	// Setup authentication
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:]) // Remove 0x prefix
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("Failed to create auth: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return client, auth, fromAddress
}
