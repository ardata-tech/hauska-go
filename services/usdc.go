package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ardata-tech/hauska-go/contracts"
	"github.com/ardata-tech/hauska-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// USDCService implements USDC token operations
type USDCService struct {
	client       *ethclient.Client
	usdcContract *contracts.MockUSDC
	auth         *bind.TransactOpts
}

// NewUSDCService creates a new USDC service instance
func NewUSDCService(client *ethclient.Client, usdcAddr common.Address, auth *bind.TransactOpts) (*USDCService, error) {
	usdcContract, err := contracts.NewMockUSDC(usdcAddr, client)
	if err != nil {
		return nil, err
	}

	return &USDCService{
		client:       client,
		usdcContract: usdcContract,
		auth:         auth,
	}, nil
}

// GetBalance returns USDC balance for an account
func (s *USDCService) GetBalance(ctx context.Context, account string) (*big.Int, error) {
	return s.usdcContract.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(account))
}

// Transfer transfers USDC to another account
func (s *USDCService) Transfer(ctx context.Context, to string, amount *big.Int) (*types.TransactionResult, error) {
	tx, err := s.usdcContract.Transfer(s.auth, common.HexToAddress(to), amount)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// Approve approves spender to transfer USDC on behalf of caller
func (s *USDCService) Approve(ctx context.Context, spender string, amount *big.Int) (*types.TransactionResult, error) {
	tx, err := s.usdcContract.Approve(s.auth, common.HexToAddress(spender), amount)
	if err != nil {
		return nil, err
	}

	return s.waitForTransaction(ctx, tx.Hash())
}

// GetAllowance returns the allowance granted to spender by owner
func (s *USDCService) GetAllowance(ctx context.Context, owner, spender string) (*big.Int, error) {
	return s.usdcContract.Allowance(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner), common.HexToAddress(spender))
}

// GetDecimals returns the number of decimals for USDC
func (s *USDCService) GetDecimals(ctx context.Context) (uint8, error) {
	return s.usdcContract.Decimals(&bind.CallOpts{Context: ctx})
}

// GetSymbol returns the token symbol
func (s *USDCService) GetSymbol(ctx context.Context) (string, error) {
	return s.usdcContract.Symbol(&bind.CallOpts{Context: ctx})
}

// GetName returns the token name
func (s *USDCService) GetName(ctx context.Context) (string, error) {
	return s.usdcContract.Name(&bind.CallOpts{Context: ctx})
}

// GetTotalSupply returns the total supply of USDC
func (s *USDCService) GetTotalSupply(ctx context.Context) (*big.Int, error) {
	return s.usdcContract.TotalSupply(&bind.CallOpts{Context: ctx})
}

// FormatAmount formats a USDC amount from wei to human-readable format
func (s *USDCService) FormatAmount(amount *big.Int) string {
	// USDC has 6 decimals
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

	// Get integer part
	integerPart := new(big.Int).Div(amount, divisor)

	// Get fractional part
	remainder := new(big.Int).Mod(amount, divisor)

	if remainder.Cmp(big.NewInt(0)) == 0 {
		return integerPart.String()
	}

	// Format with decimal places, removing trailing zeros
	fractionalStr := remainder.String()
	fractionalStr = strings.TrimRight(fractionalStr, "0")

	if len(fractionalStr) == 0 {
		return integerPart.String()
	}

	// Pad with leading zeros if necessary
	for len(fractionalStr) < 6 {
		fractionalStr = "0" + fractionalStr
	}

	return integerPart.String() + "." + fractionalStr
}

// ParseAmount parses a human-readable USDC amount to wei
func (s *USDCService) ParseAmount(amount string) (*big.Int, error) {
	parts := strings.Split(amount, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid amount format")
	}

	integerPart := parts[0]
	fractionalPart := ""
	if len(parts) == 2 {
		fractionalPart = parts[1]
	}

	// Pad or truncate fractional part to 6 digits
	for len(fractionalPart) < 6 {
		fractionalPart += "0"
	}
	if len(fractionalPart) > 6 {
		fractionalPart = fractionalPart[:6]
	}

	// Convert to big.Int
	fullAmount := integerPart + fractionalPart
	result, ok := new(big.Int).SetString(fullAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount format")
	}

	return result, nil
}

// Helper functions

func (s *USDCService) waitForTransaction(ctx context.Context, txHash common.Hash) (*types.TransactionResult, error) {
	// TODO: Implement transaction waiting and parsing
	// This is a placeholder implementation
	return &types.TransactionResult{
		TxHash: txHash,
		Status: 1,
	}, nil
}
