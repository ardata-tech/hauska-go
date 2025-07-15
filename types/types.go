package types

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Common errors
var (
	ErrNotImplemented = errors.New("not implemented")
)

// FxPool represents the type of asset pool
type FxPool uint8

const (
	FxPoolPII  FxPool = 0 // Personal Identifiable Information
	FxPoolDT   FxPool = 1 // Data Transfer
	FxPoolVDAS FxPool = 2 // Verified Digital Asset Storage
)

// LicensePermissions represents what the licensee can do with the asset
type LicensePermissions uint8

const (
	LicensePermissionsView   LicensePermissions = 0 // View only
	LicensePermissionsResell LicensePermissions = 1 // Can resell
	LicensePermissionsBoth   LicensePermissions = 2 // Both view and resell
)

// CountryCode represents supported country codes for geographic restrictions
type CountryCode uint8

const (
	CountryCodeUS CountryCode = 0 // United States
	CountryCodeRU CountryCode = 1 // Russia
	CountryCodeJP CountryCode = 2 // Japan
	CountryCodeCN CountryCode = 3 // China
	CountryCodeUK CountryCode = 4 // United Kingdom
	CountryCodeDE CountryCode = 5 // Germany
	CountryCodeFR CountryCode = 6 // France
	CountryCodeCA CountryCode = 7 // Canada
	CountryCodeAU CountryCode = 8 // Australia
	CountryCodeBR CountryCode = 9 // Brazil
)

// VerifiedDigitalAsset represents a digital asset in the Hauska ecosystem
type VerifiedDigitalAsset struct {
	AssetID                *big.Int       `json:"assetId"`
	Creator                common.Address `json:"creator"`
	Owner                  common.Address `json:"owner"`
	Partner                common.Address `json:"partner"`
	IPFSHash               string         `json:"ipfsHash"`
	MetadataHash           string         `json:"metadataHash"`
	AssetHash              [32]byte       `json:"assetHash"`
	Version                *big.Int       `json:"version"`
	IsVerified             bool           `json:"isVerified"`
	CreationTime           *big.Int       `json:"creationTime"`
	LastTransferTime       *big.Int       `json:"lastTransferTime"`
	Price                  *big.Int       `json:"price"` // In USDC units (6 decimals)
	Encrypted              bool           `json:"encrypted"`
	CanBeLicensed          bool           `json:"canBeLicensed"`
	FxPool                 FxPool         `json:"fxPool"`
	EventTimestamp         string         `json:"eventTimestamp"`
	GeographicRestrictions []CountryCode  `json:"geographicRestrictions"`
}

// Organization represents an organization contract and its details
type Organization struct {
	Address            common.Address   `json:"address"`
	Principal          common.Address   `json:"principal"`
	IntegrationPartner common.Address   `json:"integrationPartner"`
	Factory            common.Address   `json:"factory"`
	LicenseManager     common.Address   `json:"licenseManager"`
	AssetRegistry      common.Address   `json:"assetRegistry"`
	GroupManager       common.Address   `json:"groupManager"`
	RevenueDistributor common.Address   `json:"revenueDistributor"`
	Creators           []common.Address `json:"creators"`
	AssetCount         *big.Int         `json:"assetCount"`
	CreatorCount       *big.Int         `json:"creatorCount"`
	GroupCount         *big.Int         `json:"groupCount"`
}

// AssetGroup represents a group of assets that can be licensed together
type AssetGroup struct {
	GroupID    *big.Int       `json:"groupId"`
	Name       string         `json:"name"`
	Owner      common.Address `json:"owner"`
	Creator    common.Address `json:"creator"`
	Members    []*big.Int     `json:"members"`
	GroupPrice *big.Int       `json:"groupPrice"`
	CreatedAt  *big.Int       `json:"createdAt"`
	IsActive   bool           `json:"isActive"`
}

// LicenseData represents license information, compatible with both V1 and V2 managers
type LicenseData struct {
	LicenseID      *big.Int             `json:"licenseId"`
	AssetID        *big.Int             `json:"assetId"`
	Licensor       common.Address       `json:"licensor"`
	Licensee       common.Address       `json:"licensee"`
	Fee            *big.Int             `json:"fee"`
	ResellerFee    *big.Int             `json:"resellerFee"`
	Permissions    []LicensePermissions `json:"permissions"`
	OrgContract    common.Address       `json:"orgContract"`
	ExpirationTime *big.Int             `json:"expirationTime"`
	IsActive       bool                 `json:"isActive"`
	IssuedAt       *big.Int             `json:"issuedAt"`

	// NFT-specific fields (V2 only)
	TokenID         *big.Int       `json:"tokenId,omitempty"`
	OriginalCreator common.Address `json:"originalCreator,omitempty"`
	MetadataURI     string         `json:"metadataUri,omitempty"`
}

// PlatformFees represents the fee structure for the platform
type PlatformFees struct {
	HauskaFeePct     uint32 `json:"hauskaFeePct"`     // Basis points (100 = 1%)
	IntegratorFeePct uint32 `json:"integratorFeePct"` // Basis points (100 = 1%)
}

// ModuleAddresses holds all deployed module contract addresses
type ModuleAddresses struct {
	LicenseManager     common.Address `json:"licenseManager"`
	AssetRegistry      common.Address `json:"assetRegistry"`
	GroupManager       common.Address `json:"groupManager"`
	RevenueDistributor common.Address `json:"revenueDistributor"`
	AssetNFT           common.Address `json:"assetNft,omitempty"`
	LicenseNFT         common.Address `json:"licenseNft,omitempty"`
}

// TransactionResult represents the result of a blockchain transaction
type TransactionResult struct {
	TxHash      common.Hash `json:"txHash"`
	BlockNumber *big.Int    `json:"blockNumber"`
	GasUsed     uint64      `json:"gasUsed"`
	Status      uint64      `json:"status"`
	Events      []EventLog  `json:"events"`
	Timestamp   time.Time   `json:"timestamp"`
	Error       string      `json:"error,omitempty"`
}

// EventLog represents a parsed event from a transaction
type EventLog struct {
	Name    string                 `json:"name"`
	Address common.Address         `json:"address"`
	Data    map[string]interface{} `json:"data"`
}

// CreateAssetRequest represents a request to create a new asset
type CreateAssetRequest struct {
	AssetCID               string        `json:"assetCid"`
	MetadataCID            string        `json:"metadataCid"`
	AssetHash              [32]byte      `json:"assetHash"`
	Price                  *big.Int      `json:"price"`
	IsEncrypted            bool          `json:"isEncrypted"`
	CanBeLicensed          bool          `json:"canBeLicensed"`
	FxPool                 FxPool        `json:"fxPool"`
	Timestamp              string        `json:"timestamp"`
	GeographicRestrictions []CountryCode `json:"geographicRestrictions"`
}

// LicenseAssetRequest represents a request to license an asset
type LicenseAssetRequest struct {
	OrgContract common.Address       `json:"orgContract"`
	AssetID     *big.Int             `json:"assetId"`
	Licensee    common.Address       `json:"licensee"`
	Permissions []LicensePermissions `json:"permissions"`
	ResellerFee *big.Int             `json:"resellerFee"`
	Duration    *big.Int             `json:"duration,omitempty"` // For V2 with time-limited licenses
}

// CreateGroupRequest represents a request to create an asset group
type CreateGroupRequest struct {
	GroupName  string     `json:"groupName"`
	AssetIDs   []*big.Int `json:"assetIds"`
	GroupPrice *big.Int   `json:"groupPrice"`
}

// RevenueStats represents revenue statistics for an organization
type RevenueStats struct {
	TotalRevenue       *big.Int `json:"totalRevenue"`
	CreatorEarnings    *big.Int `json:"creatorEarnings"`
	PlatformEarnings   *big.Int `json:"platformEarnings"`
	IntegratorEarnings *big.Int `json:"integratorEarnings"`
	TotalLicenses      *big.Int `json:"totalLicenses"`
	ActiveLicenses     *big.Int `json:"activeLicenses"`
}

// NFTLicenseInfo represents NFT-specific license information
type NFTLicenseInfo struct {
	TokenID         *big.Int       `json:"tokenId"`
	Owner           common.Address `json:"owner"`
	AssetID         *big.Int       `json:"assetId"`
	OrgContract     common.Address `json:"orgContract"`
	OriginalCreator common.Address `json:"originalCreator"`
	IssuedAt        *big.Int       `json:"issuedAt"`
	ExpirationTime  *big.Int       `json:"expirationTime"`
	Permissions     uint8          `json:"permissions"`
	ResellerFee     *big.Int       `json:"resellerFee"`
	IsActive        bool           `json:"isActive"`
	MetadataURI     string         `json:"metadataUri"`
	CanResell       bool           `json:"canResell"`
}

// BatchOperationResult represents the result of a batch operation
type BatchOperationResult struct {
	SuccessCount int                 `json:"successCount"`
	FailureCount int                 `json:"failureCount"`
	Results      []TransactionResult `json:"results"`
	Errors       []error             `json:"errors"`
}

// RelicenseRequest represents a request to relicense an asset
type RelicenseRequest struct {
	OriginalLicenseID *big.Int             `json:"originalLicenseId"`
	NewLicensee       common.Address       `json:"newLicensee"`
	Markup            *big.Int             `json:"markup"`
	Permissions       []LicensePermissions `json:"permissions"`
}

// AssetQueryFilter represents filters for querying assets
type AssetQueryFilter struct {
	Creator      *common.Address `json:"creator,omitempty"`
	Owner        *common.Address `json:"owner,omitempty"`
	IsVerified   *bool           `json:"isVerified,omitempty"`
	CanLicense   *bool           `json:"canLicense,omitempty"`
	MinPrice     *big.Int        `json:"minPrice,omitempty"`
	MaxPrice     *big.Int        `json:"maxPrice,omitempty"`
	FxPool       *FxPool         `json:"fxPool,omitempty"`
	CreatedAfter *big.Int        `json:"createdAfter,omitempty"`
}

// LicenseQueryFilter represents filters for querying licenses
type LicenseQueryFilter struct {
	Licensee      *common.Address     `json:"licensee,omitempty"`
	Licensor      *common.Address     `json:"licensor,omitempty"`
	AssetID       *big.Int            `json:"assetId,omitempty"`
	IsActive      *bool               `json:"isActive,omitempty"`
	IsExpired     *bool               `json:"isExpired,omitempty"`
	HasPermission *LicensePermissions `json:"hasPermission,omitempty"`
	IssuedAfter   *big.Int            `json:"issuedAfter,omitempty"`
}

// ContractInfo represents information about a deployed contract
type ContractInfo struct {
	Address     common.Address `json:"address"`
	Name        string         `json:"name"`
	Version     string         `json:"version"`
	Deployed    time.Time      `json:"deployed"`
	TxHash      common.Hash    `json:"txHash"`
	BlockNumber *big.Int       `json:"blockNumber"`
}

// GasSettings represents gas configuration for transactions
type GasSettings struct {
	GasLimit             uint64   `json:"gasLimit"`
	GasPrice             *big.Int `json:"gasPrice,omitempty"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas,omitempty"`
}

// Helper functions for type conversions

// String returns string representation of FxPool
func (f FxPool) String() string {
	switch f {
	case FxPoolPII:
		return "PII"
	case FxPoolDT:
		return "DT"
	case FxPoolVDAS:
		return "VDAS"
	default:
		return "Unknown"
	}
}

// String returns string representation of LicensePermissions
func (l LicensePermissions) String() string {
	switch l {
	case LicensePermissionsView:
		return "View"
	case LicensePermissionsResell:
		return "Resell"
	case LicensePermissionsBoth:
		return "Both"
	default:
		return "Unknown"
	}
}

// String returns string representation of CountryCode
func (c CountryCode) String() string {
	switch c {
	case CountryCodeUS:
		return "US"
	case CountryCodeRU:
		return "RU"
	case CountryCodeJP:
		return "JP"
	case CountryCodeCN:
		return "CN"
	case CountryCodeUK:
		return "UK"
	case CountryCodeDE:
		return "DE"
	case CountryCodeFR:
		return "FR"
	case CountryCodeCA:
		return "CA"
	case CountryCodeAU:
		return "AU"
	case CountryCodeBR:
		return "BR"
	default:
		return "Unknown"
	}
}
