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

// NFTService implements NFT-related operations for both Asset and License NFTs
type NFTService struct {
	client   *ethclient.Client
	assetNFT *contracts.HauskaAssetNFT
	auth     *bind.TransactOpts
}

// NewNFTService creates a new NFT service instance
func NewNFTService(client *ethclient.Client, assetNFTAddr common.Address, auth *bind.TransactOpts) (*NFTService, error) {
	assetNFT, err := contracts.NewHauskaAssetNFT(assetNFTAddr, client)
	if err != nil {
		return nil, err
	}

	return &NFTService{
		client:   client,
		assetNFT: assetNFT,
		auth:     auth,
	}, nil
}

// Asset NFT Methods

// GetAssetNFTOwner returns the owner of an asset NFT
func (s *NFTService) GetAssetNFTOwner(ctx context.Context, tokenID *big.Int) (string, error) {
	owner, err := s.assetNFT.OwnerOf(&bind.CallOpts{Context: ctx}, tokenID)
	if err != nil {
		return "", err
	}
	return owner.Hex(), nil
}

// GetAssetNFTTokenURI returns the token URI for an asset NFT
func (s *NFTService) GetAssetNFTTokenURI(ctx context.Context, tokenID *big.Int) (string, error) {
	return s.assetNFT.TokenURI(&bind.CallOpts{Context: ctx}, tokenID)
}

// GetAssetNFTTokenID returns the token ID for a specific asset
func (s *NFTService) GetAssetNFTTokenID(ctx context.Context, orgContract string, assetID *big.Int) (*big.Int, error) {
	return s.assetNFT.GetTokenIdForAsset(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), assetID)
}

// TransferAssetNFT transfers an asset NFT to a new owner
func (s *NFTService) TransferAssetNFT(ctx context.Context, from, to string, tokenID *big.Int) (*types.TransactionResult, error) {
	tx, err := s.assetNFT.SafeTransferFrom(s.auth, common.HexToAddress(from), common.HexToAddress(to), tokenID)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// ApproveAssetNFT approves another address to transfer a specific asset NFT
func (s *NFTService) ApproveAssetNFT(ctx context.Context, to string, tokenID *big.Int) (*types.TransactionResult, error) {
	tx, err := s.assetNFT.Approve(s.auth, common.HexToAddress(to), tokenID)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAssetNFTBalance returns the NFT balance for an address
func (s *NFTService) GetAssetNFTBalance(ctx context.Context, owner string) (*big.Int, error) {
	return s.assetNFT.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner))
}

// GetAssetNFTApproved returns the approved address for a token ID
func (s *NFTService) GetAssetNFTApproved(ctx context.Context, tokenID *big.Int) (string, error) {
	approved, err := s.assetNFT.GetApproved(&bind.CallOpts{Context: ctx}, tokenID)
	if err != nil {
		return "", err
	}
	return approved.Hex(), nil
}

// IsAssetNFTApprovedForAll returns whether the operator is approved for all tokens of the owner
func (s *NFTService) IsAssetNFTApprovedForAll(ctx context.Context, owner, operator string) (bool, error) {
	return s.assetNFT.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner), common.HexToAddress(operator))
}

// SetAssetNFTApprovalForAll sets or unsets the approval of a given operator
func (s *NFTService) SetAssetNFTApprovalForAll(ctx context.Context, operator string, approved bool) (*types.TransactionResult, error) {
	tx, err := s.assetNFT.SetApprovalForAll(s.auth, common.HexToAddress(operator), approved)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// MintAssetNFT mints a new asset NFT
func (s *NFTService) MintAssetNFT(ctx context.Context, to, orgContract string, assetID *big.Int, uri string) (*types.TransactionResult, error) {
	tx, err := s.assetNFT.MintAssetNFT(s.auth, common.HexToAddress(to), common.HexToAddress(orgContract), assetID, uri)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// BurnAssetNFT burns an asset NFT
func (s *NFTService) BurnAssetNFT(ctx context.Context, tokenID *big.Int) (*types.TransactionResult, error) {
	tx, err := s.assetNFT.Burn(s.auth, tokenID)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// UpdateAssetNFTTokenURI updates the token URI for an asset NFT
func (s *NFTService) UpdateAssetNFTTokenURI(ctx context.Context, tokenID *big.Int, uri string) (*types.TransactionResult, error) {
	tx, err := s.assetNFT.UpdateTokenURI(s.auth, tokenID, uri)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// CheckAssetHasNFT checks if an asset has an associated NFT
func (s *NFTService) CheckAssetHasNFT(ctx context.Context, orgContract string, assetID *big.Int) (bool, error) {
	return s.assetNFT.AssetHasNFT(&bind.CallOpts{Context: ctx}, common.HexToAddress(orgContract), assetID)
}

// GetAssetNFTName returns the name of the NFT contract
func (s *NFTService) GetAssetNFTName(ctx context.Context) (string, error) {
	return s.assetNFT.Name(&bind.CallOpts{Context: ctx})
}

// GetAssetNFTSymbol returns the symbol of the NFT contract
func (s *NFTService) GetAssetNFTSymbol(ctx context.Context) (string, error) {
	return s.assetNFT.Symbol(&bind.CallOpts{Context: ctx})
}

// Helper functions

func (s *NFTService) waitForTransaction(ctx context.Context, txHash common.Hash) (*types.TransactionResult, error) {
	// TODO: Implement transaction waiting and parsing
	// This is a placeholder implementation
	return &types.TransactionResult{
		TxHash: txHash,
		Status: 1,
	}, nil
}

// RevokeLicenseNFT revokes a license NFT (admin only)
func (s *NFTService) RevokeLicenseNFT(ctx context.Context, tokenID *big.Int) (*types.TransactionResult, error) {
	// TODO: Implement contract call
	// 1. Call licenseNFT.RevokeLicense()
	// 2. Wait for transaction confirmation
	// 3. Return transaction result

	return nil, types.ErrNotImplemented
}

// Common NFT Methods

// GetNFTBalance returns the NFT balance for a user for a specific NFT contract
func (s *NFTService) GetNFTBalance(ctx context.Context, nftType string, owner string) (*big.Int, error) {
	// TODO: Implement contract call
	// 1. Determine which NFT contract to use based on nftType
	// 2. Call BalanceOf()
	// 3. Return balance

	return nil, types.ErrNotImplemented
}

// SetApprovalForAll sets approval for all NFTs of a specific type
func (s *NFTService) SetApprovalForAll(ctx context.Context, nftType string, operator string, approved bool) (*types.TransactionResult, error) {
	// TODO: Implement contract call
	// 1. Determine which NFT contract to use based on nftType
	// 2. Call SetApprovalForAll()
	// 3. Wait for transaction confirmation
	// 4. Return transaction result

	return nil, types.ErrNotImplemented
}

// IsApprovedForAll checks if an operator is approved for all NFTs of an owner
func (s *NFTService) IsApprovedForAll(ctx context.Context, nftType string, owner, operator string) (bool, error) {
	// TODO: Implement contract call
	// 1. Determine which NFT contract to use based on nftType
	// 2. Call IsApprovedForAll()
	// 3. Return boolean result

	return false, types.ErrNotImplemented
}

// GetApproved returns the approved address for a specific NFT token
func (s *NFTService) GetApproved(ctx context.Context, nftType string, tokenID *big.Int) (string, error) {
	// TODO: Implement contract call
	// 1. Determine which NFT contract to use based on nftType
	// 2. Call GetApproved()
	// 3. Return approved address as string

	return "", types.ErrNotImplemented
}
