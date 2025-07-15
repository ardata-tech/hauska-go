package utils

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// AddressUtils provides utility functions for Ethereum addresses
type AddressUtils struct{}

// IsValidAddress checks if a string is a valid Ethereum address
func (a *AddressUtils) IsValidAddress(address string) bool {
	// Basic regex check for Ethereum address format
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

// ToChecksumAddress converts an address to checksum format
func (a *AddressUtils) ToChecksumAddress(address string) string {
	// TODO: Implement EIP-55 checksum
	// For now, return lowercase
	return strings.ToLower(address)
}

// IsZeroAddress checks if an address is the zero address
func (a *AddressUtils) IsZeroAddress(address string) bool {
	return strings.ToLower(address) == "0x0000000000000000000000000000000000000000" ||
		address == "0x0"
}

// AmountUtils provides utility functions for amount conversions
type AmountUtils struct{}

// USDCToWei converts USDC amount (6 decimals) to wei representation
func (a *AmountUtils) USDCToWei(amount string) (*big.Int, error) {
	// Parse as decimal to handle floating point
	d, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, fmt.Errorf("invalid amount format: %v", err)
	}

	// USDC has 6 decimals
	multiplier := decimal.NewFromInt(1000000) // 10^6
	result := d.Mul(multiplier)

	if !result.IsInteger() {
		return nil, fmt.Errorf("amount has too many decimal places")
	}

	return result.BigInt(), nil
}

// WeiToUSDC converts wei representation to USDC amount string
func (a *AmountUtils) WeiToUSDC(wei *big.Int) string {
	// Convert to decimal with 6 decimal places
	d := decimal.NewFromBigInt(wei, -6)
	return d.String()
}

// EthToWei converts ETH amount to wei
func (a *AmountUtils) EthToWei(amount string) (*big.Int, error) {
	d, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, fmt.Errorf("invalid amount format: %v", err)
	}

	// ETH has 18 decimals
	multiplier := decimal.NewFromBigInt(big.NewInt(10).Exp(big.NewInt(10), big.NewInt(18), nil), 0)
	result := d.Mul(multiplier)

	if !result.IsInteger() {
		return nil, fmt.Errorf("amount has too many decimal places")
	}

	return result.BigInt(), nil
}

// WeiToEth converts wei to ETH amount string
func (a *AmountUtils) WeiToEth(wei *big.Int) string {
	d := decimal.NewFromBigInt(wei, -18)
	return d.String()
}

// HashUtils provides utility functions for hashing
type HashUtils struct{}

// GenerateAssetHash generates a hash for asset content
func (h *HashUtils) GenerateAssetHash(content []byte) [32]byte {
	hash := sha256.Sum256(content)
	return hash
}

// GenerateAssetHashFromString generates a hash from string content
func (h *HashUtils) GenerateAssetHashFromString(content string) [32]byte {
	return h.GenerateAssetHash([]byte(content))
}

// TimeUtils provides utility functions for time handling
type TimeUtils struct{}

// ToUnixTimestamp converts time to Unix timestamp
func (t *TimeUtils) ToUnixTimestamp(time time.Time) *big.Int {
	return big.NewInt(time.Unix())
}

// FromUnixTimestamp converts Unix timestamp to time
func (t *TimeUtils) FromUnixTimestamp(timestamp *big.Int) time.Time {
	return time.Unix(timestamp.Int64(), 0)
}

// DurationToSeconds converts duration to seconds as big.Int
func (t *TimeUtils) DurationToSeconds(duration time.Duration) *big.Int {
	return big.NewInt(int64(duration.Seconds()))
}

// ValidationUtils provides validation functions
type ValidationUtils struct{}

// ValidateIPFSHash validates an IPFS hash format
func (v *ValidationUtils) ValidateIPFSHash(hash string) error {
	if len(hash) == 0 {
		return fmt.Errorf("IPFS hash cannot be empty")
	}

	// Basic validation - IPFS hashes typically start with "Qm" or "ba"
	if !strings.HasPrefix(hash, "Qm") && !strings.HasPrefix(hash, "ba") {
		return fmt.Errorf("invalid IPFS hash format")
	}

	// Basic length check (CIDv0 is typically 46 characters)
	if len(hash) < 40 {
		return fmt.Errorf("IPFS hash too short")
	}

	return nil
}

// ValidateAssetPrice validates asset price is reasonable
func (v *ValidationUtils) ValidateAssetPrice(price *big.Int) error {
	if price == nil {
		return fmt.Errorf("price cannot be nil")
	}

	if price.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	// Maximum price check (e.g., 1 million USDC)
	maxPrice := big.NewInt(1000000000000) // 1M USDC with 6 decimals
	if price.Cmp(maxPrice) > 0 {
		return fmt.Errorf("price exceeds maximum allowed")
	}

	return nil
}

// ValidateResellerFee validates reseller fee is reasonable
func (v *ValidationUtils) ValidateResellerFee(fee, assetPrice *big.Int) error {
	if fee == nil {
		return fmt.Errorf("reseller fee cannot be nil")
	}

	if fee.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("reseller fee cannot be negative")
	}

	// Reseller fee should not exceed asset price
	if assetPrice != nil && fee.Cmp(assetPrice) > 0 {
		return fmt.Errorf("reseller fee cannot exceed asset price")
	}

	return nil
}

// ConversionUtils provides conversion utilities
type ConversionUtils struct{}

// StringArrayToBigIntArray converts string array to big.Int array
func (c *ConversionUtils) StringArrayToBigIntArray(strs []string) ([]*big.Int, error) {
	result := make([]*big.Int, len(strs))

	for i, str := range strs {
		num, ok := new(big.Int).SetString(str, 10)
		if !ok {
			return nil, fmt.Errorf("invalid number format: %s", str)
		}
		result[i] = num
	}

	return result, nil
}

// BigIntArrayToStringArray converts big.Int array to string array
func (c *ConversionUtils) BigIntArrayToStringArray(nums []*big.Int) []string {
	result := make([]string, len(nums))

	for i, num := range nums {
		result[i] = num.String()
	}

	return result
}

// Global utility instances
var (
	Address    = &AddressUtils{}
	Amount     = &AmountUtils{}
	Hash       = &HashUtils{}
	Time       = &TimeUtils{}
	Validation = &ValidationUtils{}
	Conversion = &ConversionUtils{}
)
