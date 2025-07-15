// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IHauskaStructsVerifiedDigitalAsset is an auto generated low-level Go binding around an user-defined struct.
type IHauskaStructsVerifiedDigitalAsset struct {
	AssetId                *big.Int
	Creator                common.Address
	Owner                  common.Address
	Partner                common.Address
	IpfsHash               string
	MetadataHash           string
	AssetHash              [32]byte
	Version                *big.Int
	IsVerified             bool
	CreationTime           *big.Int
	LastTransferTime       *big.Int
	Price                  *big.Int
	Encrypted              bool
	CanBeLicensed          bool
	FxPool                 uint8
	EventTimestamp         string
	GeographicRestrictions []uint8
}

// HauskaAssetRegistryMetaData contains all meta data concerning the HauskaAssetRegistry contract.
var HauskaAssetRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetHash\",\"type\":\"bytes32\"}],\"name\":\"AssetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromOrg\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toOrg\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transferredBy\",\"type\":\"address\"}],\"name\":\"AssetTransferredCrossOrg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"}],\"name\":\"AssetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"AssetVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORIZED_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORG_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRY_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"addOrgContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetNFT\",\"outputs\":[{\"internalType\":\"contractHauskaAssetNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataHash\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"assetHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"creationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastTransferTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"encrypted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBeLicensed\",\"type\":\"bool\"},{\"internalType\":\"enumIHauskaStructs.FxPool\",\"name\":\"fxPool\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"eventTimestamp\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creatorAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataHash\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"assetHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"creationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastTransferTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"encrypted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBeLicensed\",\"type\":\"bool\"},{\"internalType\":\"enumIHauskaStructs.FxPool\",\"name\":\"fxPool\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"eventTimestamp\",\"type\":\"string\"},{\"internalType\":\"enumIHauskaStructs.CountryCode[]\",\"name\":\"geographicRestrictions\",\"type\":\"uint8[]\"}],\"internalType\":\"structIHauskaStructs.VerifiedDigitalAsset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"getAssetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getAssetsByCreator\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalAssetHashExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ipfsHashToAssetId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"isAssetVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataHash\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"assetHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"creationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastTransferTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"encrypted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBeLicensed\",\"type\":\"bool\"},{\"internalType\":\"enumIHauskaStructs.FxPool\",\"name\":\"fxPool\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"eventTimestamp\",\"type\":\"string\"},{\"internalType\":\"enumIHauskaStructs.CountryCode[]\",\"name\":\"geographicRestrictions\",\"type\":\"uint8[]\"}],\"internalType\":\"structIHauskaStructs.VerifiedDigitalAsset\",\"name\":\"asset\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"registerAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"removeOrgContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetNFT\",\"type\":\"address\"}],\"name\":\"setAssetNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromOrg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toOrg\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAssetCrossOrg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newAssetId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"transferAssetOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newIpfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newMetadataHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canBeLicensed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"updateAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"verifyAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HauskaAssetRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use HauskaAssetRegistryMetaData.ABI instead.
var HauskaAssetRegistryABI = HauskaAssetRegistryMetaData.ABI

// HauskaAssetRegistry is an auto generated Go binding around an Ethereum contract.
type HauskaAssetRegistry struct {
	HauskaAssetRegistryCaller     // Read-only binding to the contract
	HauskaAssetRegistryTransactor // Write-only binding to the contract
	HauskaAssetRegistryFilterer   // Log filterer for contract events
}

// HauskaAssetRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HauskaAssetRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaAssetRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HauskaAssetRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaAssetRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HauskaAssetRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaAssetRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HauskaAssetRegistrySession struct {
	Contract     *HauskaAssetRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HauskaAssetRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HauskaAssetRegistryCallerSession struct {
	Contract *HauskaAssetRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// HauskaAssetRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HauskaAssetRegistryTransactorSession struct {
	Contract     *HauskaAssetRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// HauskaAssetRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HauskaAssetRegistryRaw struct {
	Contract *HauskaAssetRegistry // Generic contract binding to access the raw methods on
}

// HauskaAssetRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HauskaAssetRegistryCallerRaw struct {
	Contract *HauskaAssetRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// HauskaAssetRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HauskaAssetRegistryTransactorRaw struct {
	Contract *HauskaAssetRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHauskaAssetRegistry creates a new instance of HauskaAssetRegistry, bound to a specific deployed contract.
func NewHauskaAssetRegistry(address common.Address, backend bind.ContractBackend) (*HauskaAssetRegistry, error) {
	contract, err := bindHauskaAssetRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistry{HauskaAssetRegistryCaller: HauskaAssetRegistryCaller{contract: contract}, HauskaAssetRegistryTransactor: HauskaAssetRegistryTransactor{contract: contract}, HauskaAssetRegistryFilterer: HauskaAssetRegistryFilterer{contract: contract}}, nil
}

// NewHauskaAssetRegistryCaller creates a new read-only instance of HauskaAssetRegistry, bound to a specific deployed contract.
func NewHauskaAssetRegistryCaller(address common.Address, caller bind.ContractCaller) (*HauskaAssetRegistryCaller, error) {
	contract, err := bindHauskaAssetRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryCaller{contract: contract}, nil
}

// NewHauskaAssetRegistryTransactor creates a new write-only instance of HauskaAssetRegistry, bound to a specific deployed contract.
func NewHauskaAssetRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*HauskaAssetRegistryTransactor, error) {
	contract, err := bindHauskaAssetRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryTransactor{contract: contract}, nil
}

// NewHauskaAssetRegistryFilterer creates a new log filterer instance of HauskaAssetRegistry, bound to a specific deployed contract.
func NewHauskaAssetRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*HauskaAssetRegistryFilterer, error) {
	contract, err := bindHauskaAssetRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryFilterer{contract: contract}, nil
}

// bindHauskaAssetRegistry binds a generic wrapper to an already deployed contract.
func bindHauskaAssetRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HauskaAssetRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaAssetRegistry *HauskaAssetRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaAssetRegistry.Contract.HauskaAssetRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaAssetRegistry *HauskaAssetRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.HauskaAssetRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaAssetRegistry *HauskaAssetRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.HauskaAssetRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaAssetRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.contract.Transact(opts, method, params...)
}

// AUTHORIZEDCONTRACTROLE is a free data retrieval call binding the contract method 0x3182cd8c.
//
// Solidity: function AUTHORIZED_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) AUTHORIZEDCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "AUTHORIZED_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHORIZEDCONTRACTROLE is a free data retrieval call binding the contract method 0x3182cd8c.
//
// Solidity: function AUTHORIZED_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) AUTHORIZEDCONTRACTROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.AUTHORIZEDCONTRACTROLE(&_HauskaAssetRegistry.CallOpts)
}

// AUTHORIZEDCONTRACTROLE is a free data retrieval call binding the contract method 0x3182cd8c.
//
// Solidity: function AUTHORIZED_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) AUTHORIZEDCONTRACTROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.AUTHORIZEDCONTRACTROLE(&_HauskaAssetRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.DEFAULTADMINROLE(&_HauskaAssetRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.DEFAULTADMINROLE(&_HauskaAssetRegistry.CallOpts)
}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) ORGCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "ORG_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) ORGCONTRACTROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.ORGCONTRACTROLE(&_HauskaAssetRegistry.CallOpts)
}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) ORGCONTRACTROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.ORGCONTRACTROLE(&_HauskaAssetRegistry.CallOpts)
}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) REGISTRYADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "REGISTRY_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) REGISTRYADMINROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.REGISTRYADMINROLE(&_HauskaAssetRegistry.CallOpts)
}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) REGISTRYADMINROLE() ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.REGISTRYADMINROLE(&_HauskaAssetRegistry.CallOpts)
}

// AssetCounter is a free data retrieval call binding the contract method 0xd571661d.
//
// Solidity: function assetCounter(address ) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) AssetCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "assetCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetCounter is a free data retrieval call binding the contract method 0xd571661d.
//
// Solidity: function assetCounter(address ) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) AssetCounter(arg0 common.Address) (*big.Int, error) {
	return _HauskaAssetRegistry.Contract.AssetCounter(&_HauskaAssetRegistry.CallOpts, arg0)
}

// AssetCounter is a free data retrieval call binding the contract method 0xd571661d.
//
// Solidity: function assetCounter(address ) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) AssetCounter(arg0 common.Address) (*big.Int, error) {
	return _HauskaAssetRegistry.Contract.AssetCounter(&_HauskaAssetRegistry.CallOpts, arg0)
}

// AssetNFT is a free data retrieval call binding the contract method 0xac6a258d.
//
// Solidity: function assetNFT() view returns(address)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) AssetNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "assetNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetNFT is a free data retrieval call binding the contract method 0xac6a258d.
//
// Solidity: function assetNFT() view returns(address)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) AssetNFT() (common.Address, error) {
	return _HauskaAssetRegistry.Contract.AssetNFT(&_HauskaAssetRegistry.CallOpts)
}

// AssetNFT is a free data retrieval call binding the contract method 0xac6a258d.
//
// Solidity: function assetNFT() view returns(address)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) AssetNFT() (common.Address, error) {
	return _HauskaAssetRegistry.Contract.AssetNFT(&_HauskaAssetRegistry.CallOpts)
}

// AssetVerifiers is a free data retrieval call binding the contract method 0xcb41118e.
//
// Solidity: function assetVerifiers(address , uint256 ) view returns(address)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) AssetVerifiers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "assetVerifiers", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetVerifiers is a free data retrieval call binding the contract method 0xcb41118e.
//
// Solidity: function assetVerifiers(address , uint256 ) view returns(address)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) AssetVerifiers(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _HauskaAssetRegistry.Contract.AssetVerifiers(&_HauskaAssetRegistry.CallOpts, arg0, arg1)
}

// AssetVerifiers is a free data retrieval call binding the contract method 0xcb41118e.
//
// Solidity: function assetVerifiers(address , uint256 ) view returns(address)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) AssetVerifiers(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _HauskaAssetRegistry.Contract.AssetVerifiers(&_HauskaAssetRegistry.CallOpts, arg0, arg1)
}

// Assets is a free data retrieval call binding the contract method 0x51eb27c0.
//
// Solidity: function assets(address , uint256 ) view returns(uint256 assetId, address creator, address owner, address partner, string ipfsHash, string metadataHash, bytes32 assetHash, uint256 version, bool isVerified, uint256 creationTime, uint256 lastTransferTime, uint256 price, bool encrypted, bool canBeLicensed, uint8 fxPool, string eventTimestamp)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) Assets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	AssetId          *big.Int
	Creator          common.Address
	Owner            common.Address
	Partner          common.Address
	IpfsHash         string
	MetadataHash     string
	AssetHash        [32]byte
	Version          *big.Int
	IsVerified       bool
	CreationTime     *big.Int
	LastTransferTime *big.Int
	Price            *big.Int
	Encrypted        bool
	CanBeLicensed    bool
	FxPool           uint8
	EventTimestamp   string
}, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "assets", arg0, arg1)

	outstruct := new(struct {
		AssetId          *big.Int
		Creator          common.Address
		Owner            common.Address
		Partner          common.Address
		IpfsHash         string
		MetadataHash     string
		AssetHash        [32]byte
		Version          *big.Int
		IsVerified       bool
		CreationTime     *big.Int
		LastTransferTime *big.Int
		Price            *big.Int
		Encrypted        bool
		CanBeLicensed    bool
		FxPool           uint8
		EventTimestamp   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Partner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IpfsHash = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.MetadataHash = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.AssetHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.Version = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IsVerified = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.CreationTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.LastTransferTime = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.Encrypted = *abi.ConvertType(out[12], new(bool)).(*bool)
	outstruct.CanBeLicensed = *abi.ConvertType(out[13], new(bool)).(*bool)
	outstruct.FxPool = *abi.ConvertType(out[14], new(uint8)).(*uint8)
	outstruct.EventTimestamp = *abi.ConvertType(out[15], new(string)).(*string)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0x51eb27c0.
//
// Solidity: function assets(address , uint256 ) view returns(uint256 assetId, address creator, address owner, address partner, string ipfsHash, string metadataHash, bytes32 assetHash, uint256 version, bool isVerified, uint256 creationTime, uint256 lastTransferTime, uint256 price, bool encrypted, bool canBeLicensed, uint8 fxPool, string eventTimestamp)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) Assets(arg0 common.Address, arg1 *big.Int) (struct {
	AssetId          *big.Int
	Creator          common.Address
	Owner            common.Address
	Partner          common.Address
	IpfsHash         string
	MetadataHash     string
	AssetHash        [32]byte
	Version          *big.Int
	IsVerified       bool
	CreationTime     *big.Int
	LastTransferTime *big.Int
	Price            *big.Int
	Encrypted        bool
	CanBeLicensed    bool
	FxPool           uint8
	EventTimestamp   string
}, error) {
	return _HauskaAssetRegistry.Contract.Assets(&_HauskaAssetRegistry.CallOpts, arg0, arg1)
}

// Assets is a free data retrieval call binding the contract method 0x51eb27c0.
//
// Solidity: function assets(address , uint256 ) view returns(uint256 assetId, address creator, address owner, address partner, string ipfsHash, string metadataHash, bytes32 assetHash, uint256 version, bool isVerified, uint256 creationTime, uint256 lastTransferTime, uint256 price, bool encrypted, bool canBeLicensed, uint8 fxPool, string eventTimestamp)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) Assets(arg0 common.Address, arg1 *big.Int) (struct {
	AssetId          *big.Int
	Creator          common.Address
	Owner            common.Address
	Partner          common.Address
	IpfsHash         string
	MetadataHash     string
	AssetHash        [32]byte
	Version          *big.Int
	IsVerified       bool
	CreationTime     *big.Int
	LastTransferTime *big.Int
	Price            *big.Int
	Encrypted        bool
	CanBeLicensed    bool
	FxPool           uint8
	EventTimestamp   string
}, error) {
	return _HauskaAssetRegistry.Contract.Assets(&_HauskaAssetRegistry.CallOpts, arg0, arg1)
}

// CreatorAssets is a free data retrieval call binding the contract method 0x16721ee1.
//
// Solidity: function creatorAssets(address , address , uint256 ) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) CreatorAssets(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "creatorAssets", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorAssets is a free data retrieval call binding the contract method 0x16721ee1.
//
// Solidity: function creatorAssets(address , address , uint256 ) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) CreatorAssets(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _HauskaAssetRegistry.Contract.CreatorAssets(&_HauskaAssetRegistry.CallOpts, arg0, arg1, arg2)
}

// CreatorAssets is a free data retrieval call binding the contract method 0x16721ee1.
//
// Solidity: function creatorAssets(address , address , uint256 ) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) CreatorAssets(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _HauskaAssetRegistry.Contract.CreatorAssets(&_HauskaAssetRegistry.CallOpts, arg0, arg1, arg2)
}

// GetAsset is a free data retrieval call binding the contract method 0xdfd01ff3.
//
// Solidity: function getAsset(address orgContract, uint256 assetId) view returns((uint256,address,address,address,string,string,bytes32,uint256,bool,uint256,uint256,uint256,bool,bool,uint8,string,uint8[]))
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) GetAsset(opts *bind.CallOpts, orgContract common.Address, assetId *big.Int) (IHauskaStructsVerifiedDigitalAsset, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "getAsset", orgContract, assetId)

	if err != nil {
		return *new(IHauskaStructsVerifiedDigitalAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(IHauskaStructsVerifiedDigitalAsset)).(*IHauskaStructsVerifiedDigitalAsset)

	return out0, err

}

// GetAsset is a free data retrieval call binding the contract method 0xdfd01ff3.
//
// Solidity: function getAsset(address orgContract, uint256 assetId) view returns((uint256,address,address,address,string,string,bytes32,uint256,bool,uint256,uint256,uint256,bool,bool,uint8,string,uint8[]))
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) GetAsset(orgContract common.Address, assetId *big.Int) (IHauskaStructsVerifiedDigitalAsset, error) {
	return _HauskaAssetRegistry.Contract.GetAsset(&_HauskaAssetRegistry.CallOpts, orgContract, assetId)
}

// GetAsset is a free data retrieval call binding the contract method 0xdfd01ff3.
//
// Solidity: function getAsset(address orgContract, uint256 assetId) view returns((uint256,address,address,address,string,string,bytes32,uint256,bool,uint256,uint256,uint256,bool,bool,uint8,string,uint8[]))
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) GetAsset(orgContract common.Address, assetId *big.Int) (IHauskaStructsVerifiedDigitalAsset, error) {
	return _HauskaAssetRegistry.Contract.GetAsset(&_HauskaAssetRegistry.CallOpts, orgContract, assetId)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xdd7e37d6.
//
// Solidity: function getAssetCount(address orgContract) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) GetAssetCount(opts *bind.CallOpts, orgContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "getAssetCount", orgContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetCount is a free data retrieval call binding the contract method 0xdd7e37d6.
//
// Solidity: function getAssetCount(address orgContract) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) GetAssetCount(orgContract common.Address) (*big.Int, error) {
	return _HauskaAssetRegistry.Contract.GetAssetCount(&_HauskaAssetRegistry.CallOpts, orgContract)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xdd7e37d6.
//
// Solidity: function getAssetCount(address orgContract) view returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) GetAssetCount(orgContract common.Address) (*big.Int, error) {
	return _HauskaAssetRegistry.Contract.GetAssetCount(&_HauskaAssetRegistry.CallOpts, orgContract)
}

// GetAssetsByCreator is a free data retrieval call binding the contract method 0x296a6a29.
//
// Solidity: function getAssetsByCreator(address orgContract, address creator) view returns(uint256[])
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) GetAssetsByCreator(opts *bind.CallOpts, orgContract common.Address, creator common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "getAssetsByCreator", orgContract, creator)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetsByCreator is a free data retrieval call binding the contract method 0x296a6a29.
//
// Solidity: function getAssetsByCreator(address orgContract, address creator) view returns(uint256[])
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) GetAssetsByCreator(orgContract common.Address, creator common.Address) ([]*big.Int, error) {
	return _HauskaAssetRegistry.Contract.GetAssetsByCreator(&_HauskaAssetRegistry.CallOpts, orgContract, creator)
}

// GetAssetsByCreator is a free data retrieval call binding the contract method 0x296a6a29.
//
// Solidity: function getAssetsByCreator(address orgContract, address creator) view returns(uint256[])
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) GetAssetsByCreator(orgContract common.Address, creator common.Address) ([]*big.Int, error) {
	return _HauskaAssetRegistry.Contract.GetAssetsByCreator(&_HauskaAssetRegistry.CallOpts, orgContract, creator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.GetRoleAdmin(&_HauskaAssetRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaAssetRegistry.Contract.GetRoleAdmin(&_HauskaAssetRegistry.CallOpts, role)
}

// GlobalAssetHashExists is a free data retrieval call binding the contract method 0x7ed3263e.
//
// Solidity: function globalAssetHashExists(bytes32 ) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) GlobalAssetHashExists(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "globalAssetHashExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GlobalAssetHashExists is a free data retrieval call binding the contract method 0x7ed3263e.
//
// Solidity: function globalAssetHashExists(bytes32 ) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) GlobalAssetHashExists(arg0 [32]byte) (bool, error) {
	return _HauskaAssetRegistry.Contract.GlobalAssetHashExists(&_HauskaAssetRegistry.CallOpts, arg0)
}

// GlobalAssetHashExists is a free data retrieval call binding the contract method 0x7ed3263e.
//
// Solidity: function globalAssetHashExists(bytes32 ) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) GlobalAssetHashExists(arg0 [32]byte) (bool, error) {
	return _HauskaAssetRegistry.Contract.GlobalAssetHashExists(&_HauskaAssetRegistry.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaAssetRegistry.Contract.HasRole(&_HauskaAssetRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaAssetRegistry.Contract.HasRole(&_HauskaAssetRegistry.CallOpts, role, account)
}

// IpfsHashToAssetId is a free data retrieval call binding the contract method 0x31c7720c.
//
// Solidity: function ipfsHashToAssetId(string ) view returns(address orgContract, uint256 assetId)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) IpfsHashToAssetId(opts *bind.CallOpts, arg0 string) (struct {
	OrgContract common.Address
	AssetId     *big.Int
}, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "ipfsHashToAssetId", arg0)

	outstruct := new(struct {
		OrgContract common.Address
		AssetId     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrgContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AssetId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IpfsHashToAssetId is a free data retrieval call binding the contract method 0x31c7720c.
//
// Solidity: function ipfsHashToAssetId(string ) view returns(address orgContract, uint256 assetId)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) IpfsHashToAssetId(arg0 string) (struct {
	OrgContract common.Address
	AssetId     *big.Int
}, error) {
	return _HauskaAssetRegistry.Contract.IpfsHashToAssetId(&_HauskaAssetRegistry.CallOpts, arg0)
}

// IpfsHashToAssetId is a free data retrieval call binding the contract method 0x31c7720c.
//
// Solidity: function ipfsHashToAssetId(string ) view returns(address orgContract, uint256 assetId)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) IpfsHashToAssetId(arg0 string) (struct {
	OrgContract common.Address
	AssetId     *big.Int
}, error) {
	return _HauskaAssetRegistry.Contract.IpfsHashToAssetId(&_HauskaAssetRegistry.CallOpts, arg0)
}

// IsAssetVerified is a free data retrieval call binding the contract method 0xd46e938f.
//
// Solidity: function isAssetVerified(address orgContract, uint256 assetId) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) IsAssetVerified(opts *bind.CallOpts, orgContract common.Address, assetId *big.Int) (bool, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "isAssetVerified", orgContract, assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetVerified is a free data retrieval call binding the contract method 0xd46e938f.
//
// Solidity: function isAssetVerified(address orgContract, uint256 assetId) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) IsAssetVerified(orgContract common.Address, assetId *big.Int) (bool, error) {
	return _HauskaAssetRegistry.Contract.IsAssetVerified(&_HauskaAssetRegistry.CallOpts, orgContract, assetId)
}

// IsAssetVerified is a free data retrieval call binding the contract method 0xd46e938f.
//
// Solidity: function isAssetVerified(address orgContract, uint256 assetId) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) IsAssetVerified(orgContract common.Address, assetId *big.Int) (bool, error) {
	return _HauskaAssetRegistry.Contract.IsAssetVerified(&_HauskaAssetRegistry.CallOpts, orgContract, assetId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HauskaAssetRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaAssetRegistry.Contract.SupportsInterface(&_HauskaAssetRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaAssetRegistry *HauskaAssetRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaAssetRegistry.Contract.SupportsInterface(&_HauskaAssetRegistry.CallOpts, interfaceId)
}

// AddOrgContract is a paid mutator transaction binding the contract method 0xa85e5b61.
//
// Solidity: function addOrgContract(address orgContract) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) AddOrgContract(opts *bind.TransactOpts, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "addOrgContract", orgContract)
}

// AddOrgContract is a paid mutator transaction binding the contract method 0xa85e5b61.
//
// Solidity: function addOrgContract(address orgContract) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) AddOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.AddOrgContract(&_HauskaAssetRegistry.TransactOpts, orgContract)
}

// AddOrgContract is a paid mutator transaction binding the contract method 0xa85e5b61.
//
// Solidity: function addOrgContract(address orgContract) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) AddOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.AddOrgContract(&_HauskaAssetRegistry.TransactOpts, orgContract)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.GrantRole(&_HauskaAssetRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.GrantRole(&_HauskaAssetRegistry.TransactOpts, role, account)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x5fcfc4b8.
//
// Solidity: function registerAsset((uint256,address,address,address,string,string,bytes32,uint256,bool,uint256,uint256,uint256,bool,bool,uint8,string,uint8[]) asset, address creator) returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) RegisterAsset(opts *bind.TransactOpts, asset IHauskaStructsVerifiedDigitalAsset, creator common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "registerAsset", asset, creator)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x5fcfc4b8.
//
// Solidity: function registerAsset((uint256,address,address,address,string,string,bytes32,uint256,bool,uint256,uint256,uint256,bool,bool,uint8,string,uint8[]) asset, address creator) returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) RegisterAsset(asset IHauskaStructsVerifiedDigitalAsset, creator common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RegisterAsset(&_HauskaAssetRegistry.TransactOpts, asset, creator)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x5fcfc4b8.
//
// Solidity: function registerAsset((uint256,address,address,address,string,string,bytes32,uint256,bool,uint256,uint256,uint256,bool,bool,uint8,string,uint8[]) asset, address creator) returns(uint256)
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) RegisterAsset(asset IHauskaStructsVerifiedDigitalAsset, creator common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RegisterAsset(&_HauskaAssetRegistry.TransactOpts, asset, creator)
}

// RemoveOrgContract is a paid mutator transaction binding the contract method 0xc009e23b.
//
// Solidity: function removeOrgContract(address orgContract) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) RemoveOrgContract(opts *bind.TransactOpts, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "removeOrgContract", orgContract)
}

// RemoveOrgContract is a paid mutator transaction binding the contract method 0xc009e23b.
//
// Solidity: function removeOrgContract(address orgContract) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) RemoveOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RemoveOrgContract(&_HauskaAssetRegistry.TransactOpts, orgContract)
}

// RemoveOrgContract is a paid mutator transaction binding the contract method 0xc009e23b.
//
// Solidity: function removeOrgContract(address orgContract) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) RemoveOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RemoveOrgContract(&_HauskaAssetRegistry.TransactOpts, orgContract)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RenounceRole(&_HauskaAssetRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RenounceRole(&_HauskaAssetRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RevokeRole(&_HauskaAssetRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.RevokeRole(&_HauskaAssetRegistry.TransactOpts, role, account)
}

// SetAssetNFT is a paid mutator transaction binding the contract method 0xaede4559.
//
// Solidity: function setAssetNFT(address _assetNFT) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) SetAssetNFT(opts *bind.TransactOpts, _assetNFT common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "setAssetNFT", _assetNFT)
}

// SetAssetNFT is a paid mutator transaction binding the contract method 0xaede4559.
//
// Solidity: function setAssetNFT(address _assetNFT) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) SetAssetNFT(_assetNFT common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.SetAssetNFT(&_HauskaAssetRegistry.TransactOpts, _assetNFT)
}

// SetAssetNFT is a paid mutator transaction binding the contract method 0xaede4559.
//
// Solidity: function setAssetNFT(address _assetNFT) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) SetAssetNFT(_assetNFT common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.SetAssetNFT(&_HauskaAssetRegistry.TransactOpts, _assetNFT)
}

// TransferAssetCrossOrg is a paid mutator transaction binding the contract method 0xe6a550a4.
//
// Solidity: function transferAssetCrossOrg(address fromOrg, address toOrg, uint256 assetId, address newOwner) returns(uint256 newAssetId)
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) TransferAssetCrossOrg(opts *bind.TransactOpts, fromOrg common.Address, toOrg common.Address, assetId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "transferAssetCrossOrg", fromOrg, toOrg, assetId, newOwner)
}

// TransferAssetCrossOrg is a paid mutator transaction binding the contract method 0xe6a550a4.
//
// Solidity: function transferAssetCrossOrg(address fromOrg, address toOrg, uint256 assetId, address newOwner) returns(uint256 newAssetId)
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) TransferAssetCrossOrg(fromOrg common.Address, toOrg common.Address, assetId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.TransferAssetCrossOrg(&_HauskaAssetRegistry.TransactOpts, fromOrg, toOrg, assetId, newOwner)
}

// TransferAssetCrossOrg is a paid mutator transaction binding the contract method 0xe6a550a4.
//
// Solidity: function transferAssetCrossOrg(address fromOrg, address toOrg, uint256 assetId, address newOwner) returns(uint256 newAssetId)
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) TransferAssetCrossOrg(fromOrg common.Address, toOrg common.Address, assetId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.TransferAssetCrossOrg(&_HauskaAssetRegistry.TransactOpts, fromOrg, toOrg, assetId, newOwner)
}

// TransferAssetOwnership is a paid mutator transaction binding the contract method 0x197540af.
//
// Solidity: function transferAssetOwnership(address orgContract, uint256 assetId, address newOwner, address caller) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) TransferAssetOwnership(opts *bind.TransactOpts, orgContract common.Address, assetId *big.Int, newOwner common.Address, caller common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "transferAssetOwnership", orgContract, assetId, newOwner, caller)
}

// TransferAssetOwnership is a paid mutator transaction binding the contract method 0x197540af.
//
// Solidity: function transferAssetOwnership(address orgContract, uint256 assetId, address newOwner, address caller) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) TransferAssetOwnership(orgContract common.Address, assetId *big.Int, newOwner common.Address, caller common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.TransferAssetOwnership(&_HauskaAssetRegistry.TransactOpts, orgContract, assetId, newOwner, caller)
}

// TransferAssetOwnership is a paid mutator transaction binding the contract method 0x197540af.
//
// Solidity: function transferAssetOwnership(address orgContract, uint256 assetId, address newOwner, address caller) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) TransferAssetOwnership(orgContract common.Address, assetId *big.Int, newOwner common.Address, caller common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.TransferAssetOwnership(&_HauskaAssetRegistry.TransactOpts, orgContract, assetId, newOwner, caller)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x60b5b99d.
//
// Solidity: function updateAsset(address orgContract, uint256 assetId, string newIpfsHash, string newMetadataHash, uint256 newPrice, bool canBeLicensed, address caller) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) UpdateAsset(opts *bind.TransactOpts, orgContract common.Address, assetId *big.Int, newIpfsHash string, newMetadataHash string, newPrice *big.Int, canBeLicensed bool, caller common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "updateAsset", orgContract, assetId, newIpfsHash, newMetadataHash, newPrice, canBeLicensed, caller)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x60b5b99d.
//
// Solidity: function updateAsset(address orgContract, uint256 assetId, string newIpfsHash, string newMetadataHash, uint256 newPrice, bool canBeLicensed, address caller) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) UpdateAsset(orgContract common.Address, assetId *big.Int, newIpfsHash string, newMetadataHash string, newPrice *big.Int, canBeLicensed bool, caller common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.UpdateAsset(&_HauskaAssetRegistry.TransactOpts, orgContract, assetId, newIpfsHash, newMetadataHash, newPrice, canBeLicensed, caller)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x60b5b99d.
//
// Solidity: function updateAsset(address orgContract, uint256 assetId, string newIpfsHash, string newMetadataHash, uint256 newPrice, bool canBeLicensed, address caller) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) UpdateAsset(orgContract common.Address, assetId *big.Int, newIpfsHash string, newMetadataHash string, newPrice *big.Int, canBeLicensed bool, caller common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.UpdateAsset(&_HauskaAssetRegistry.TransactOpts, orgContract, assetId, newIpfsHash, newMetadataHash, newPrice, canBeLicensed, caller)
}

// VerifyAsset is a paid mutator transaction binding the contract method 0xe883844c.
//
// Solidity: function verifyAsset(address orgContract, uint256 assetId, address verifier) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactor) VerifyAsset(opts *bind.TransactOpts, orgContract common.Address, assetId *big.Int, verifier common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.contract.Transact(opts, "verifyAsset", orgContract, assetId, verifier)
}

// VerifyAsset is a paid mutator transaction binding the contract method 0xe883844c.
//
// Solidity: function verifyAsset(address orgContract, uint256 assetId, address verifier) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistrySession) VerifyAsset(orgContract common.Address, assetId *big.Int, verifier common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.VerifyAsset(&_HauskaAssetRegistry.TransactOpts, orgContract, assetId, verifier)
}

// VerifyAsset is a paid mutator transaction binding the contract method 0xe883844c.
//
// Solidity: function verifyAsset(address orgContract, uint256 assetId, address verifier) returns()
func (_HauskaAssetRegistry *HauskaAssetRegistryTransactorSession) VerifyAsset(orgContract common.Address, assetId *big.Int, verifier common.Address) (*types.Transaction, error) {
	return _HauskaAssetRegistry.Contract.VerifyAsset(&_HauskaAssetRegistry.TransactOpts, orgContract, assetId, verifier)
}

// HauskaAssetRegistryAssetRegisteredIterator is returned from FilterAssetRegistered and is used to iterate over the raw logs and unpacked data for AssetRegistered events raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetRegisteredIterator struct {
	Event *HauskaAssetRegistryAssetRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HauskaAssetRegistryAssetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetRegistryAssetRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HauskaAssetRegistryAssetRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HauskaAssetRegistryAssetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetRegistryAssetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetRegistryAssetRegistered represents a AssetRegistered event raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetRegistered struct {
	OrgContract common.Address
	AssetId     *big.Int
	Creator     common.Address
	AssetHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetRegistered is a free log retrieval operation binding the contract event 0x2c3b865998b4abe6a1c4845b2daeef80c4178bf6fd5615c13f03d6e44c39d6f2.
//
// Solidity: event AssetRegistered(address indexed orgContract, uint256 indexed assetId, address indexed creator, bytes32 assetHash)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) FilterAssetRegistered(opts *bind.FilterOpts, orgContract []common.Address, assetId []*big.Int, creator []common.Address) (*HauskaAssetRegistryAssetRegisteredIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.FilterLogs(opts, "AssetRegistered", orgContractRule, assetIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryAssetRegisteredIterator{contract: _HauskaAssetRegistry.contract, event: "AssetRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetRegistered is a free log subscription operation binding the contract event 0x2c3b865998b4abe6a1c4845b2daeef80c4178bf6fd5615c13f03d6e44c39d6f2.
//
// Solidity: event AssetRegistered(address indexed orgContract, uint256 indexed assetId, address indexed creator, bytes32 assetHash)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) WatchAssetRegistered(opts *bind.WatchOpts, sink chan<- *HauskaAssetRegistryAssetRegistered, orgContract []common.Address, assetId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.WatchLogs(opts, "AssetRegistered", orgContractRule, assetIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetRegistryAssetRegistered)
				if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetRegistered is a log parse operation binding the contract event 0x2c3b865998b4abe6a1c4845b2daeef80c4178bf6fd5615c13f03d6e44c39d6f2.
//
// Solidity: event AssetRegistered(address indexed orgContract, uint256 indexed assetId, address indexed creator, bytes32 assetHash)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) ParseAssetRegistered(log types.Log) (*HauskaAssetRegistryAssetRegistered, error) {
	event := new(HauskaAssetRegistryAssetRegistered)
	if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetRegistryAssetTransferredCrossOrgIterator is returned from FilterAssetTransferredCrossOrg and is used to iterate over the raw logs and unpacked data for AssetTransferredCrossOrg events raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetTransferredCrossOrgIterator struct {
	Event *HauskaAssetRegistryAssetTransferredCrossOrg // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HauskaAssetRegistryAssetTransferredCrossOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetRegistryAssetTransferredCrossOrg)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HauskaAssetRegistryAssetTransferredCrossOrg)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HauskaAssetRegistryAssetTransferredCrossOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetRegistryAssetTransferredCrossOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetRegistryAssetTransferredCrossOrg represents a AssetTransferredCrossOrg event raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetTransferredCrossOrg struct {
	FromOrg       common.Address
	ToOrg         common.Address
	AssetId       *big.Int
	NewAssetId    *big.Int
	TransferredBy common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAssetTransferredCrossOrg is a free log retrieval operation binding the contract event 0x9dda9b6c1846dc46857869cc78a347178e0e284601729195e375836bb3f1931f.
//
// Solidity: event AssetTransferredCrossOrg(address indexed fromOrg, address indexed toOrg, uint256 indexed assetId, uint256 newAssetId, address transferredBy)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) FilterAssetTransferredCrossOrg(opts *bind.FilterOpts, fromOrg []common.Address, toOrg []common.Address, assetId []*big.Int) (*HauskaAssetRegistryAssetTransferredCrossOrgIterator, error) {

	var fromOrgRule []interface{}
	for _, fromOrgItem := range fromOrg {
		fromOrgRule = append(fromOrgRule, fromOrgItem)
	}
	var toOrgRule []interface{}
	for _, toOrgItem := range toOrg {
		toOrgRule = append(toOrgRule, toOrgItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.FilterLogs(opts, "AssetTransferredCrossOrg", fromOrgRule, toOrgRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryAssetTransferredCrossOrgIterator{contract: _HauskaAssetRegistry.contract, event: "AssetTransferredCrossOrg", logs: logs, sub: sub}, nil
}

// WatchAssetTransferredCrossOrg is a free log subscription operation binding the contract event 0x9dda9b6c1846dc46857869cc78a347178e0e284601729195e375836bb3f1931f.
//
// Solidity: event AssetTransferredCrossOrg(address indexed fromOrg, address indexed toOrg, uint256 indexed assetId, uint256 newAssetId, address transferredBy)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) WatchAssetTransferredCrossOrg(opts *bind.WatchOpts, sink chan<- *HauskaAssetRegistryAssetTransferredCrossOrg, fromOrg []common.Address, toOrg []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var fromOrgRule []interface{}
	for _, fromOrgItem := range fromOrg {
		fromOrgRule = append(fromOrgRule, fromOrgItem)
	}
	var toOrgRule []interface{}
	for _, toOrgItem := range toOrg {
		toOrgRule = append(toOrgRule, toOrgItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.WatchLogs(opts, "AssetTransferredCrossOrg", fromOrgRule, toOrgRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetRegistryAssetTransferredCrossOrg)
				if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetTransferredCrossOrg", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetTransferredCrossOrg is a log parse operation binding the contract event 0x9dda9b6c1846dc46857869cc78a347178e0e284601729195e375836bb3f1931f.
//
// Solidity: event AssetTransferredCrossOrg(address indexed fromOrg, address indexed toOrg, uint256 indexed assetId, uint256 newAssetId, address transferredBy)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) ParseAssetTransferredCrossOrg(log types.Log) (*HauskaAssetRegistryAssetTransferredCrossOrg, error) {
	event := new(HauskaAssetRegistryAssetTransferredCrossOrg)
	if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetTransferredCrossOrg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetRegistryAssetUpdatedIterator is returned from FilterAssetUpdated and is used to iterate over the raw logs and unpacked data for AssetUpdated events raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetUpdatedIterator struct {
	Event *HauskaAssetRegistryAssetUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HauskaAssetRegistryAssetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetRegistryAssetUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HauskaAssetRegistryAssetUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HauskaAssetRegistryAssetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetRegistryAssetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetRegistryAssetUpdated represents a AssetUpdated event raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetUpdated struct {
	OrgContract common.Address
	AssetId     *big.Int
	NewVersion  *big.Int
	UpdatedBy   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetUpdated is a free log retrieval operation binding the contract event 0xd7cb9aa4a3392025204212a5b56625bd06cbf380665987943d2db85a2d9235d6.
//
// Solidity: event AssetUpdated(address indexed orgContract, uint256 indexed assetId, uint256 newVersion, address updatedBy)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) FilterAssetUpdated(opts *bind.FilterOpts, orgContract []common.Address, assetId []*big.Int) (*HauskaAssetRegistryAssetUpdatedIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.FilterLogs(opts, "AssetUpdated", orgContractRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryAssetUpdatedIterator{contract: _HauskaAssetRegistry.contract, event: "AssetUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetUpdated is a free log subscription operation binding the contract event 0xd7cb9aa4a3392025204212a5b56625bd06cbf380665987943d2db85a2d9235d6.
//
// Solidity: event AssetUpdated(address indexed orgContract, uint256 indexed assetId, uint256 newVersion, address updatedBy)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) WatchAssetUpdated(opts *bind.WatchOpts, sink chan<- *HauskaAssetRegistryAssetUpdated, orgContract []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.WatchLogs(opts, "AssetUpdated", orgContractRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetRegistryAssetUpdated)
				if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetUpdated is a log parse operation binding the contract event 0xd7cb9aa4a3392025204212a5b56625bd06cbf380665987943d2db85a2d9235d6.
//
// Solidity: event AssetUpdated(address indexed orgContract, uint256 indexed assetId, uint256 newVersion, address updatedBy)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) ParseAssetUpdated(log types.Log) (*HauskaAssetRegistryAssetUpdated, error) {
	event := new(HauskaAssetRegistryAssetUpdated)
	if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetRegistryAssetVerifiedIterator is returned from FilterAssetVerified and is used to iterate over the raw logs and unpacked data for AssetVerified events raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetVerifiedIterator struct {
	Event *HauskaAssetRegistryAssetVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HauskaAssetRegistryAssetVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetRegistryAssetVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HauskaAssetRegistryAssetVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HauskaAssetRegistryAssetVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetRegistryAssetVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetRegistryAssetVerified represents a AssetVerified event raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryAssetVerified struct {
	OrgContract common.Address
	AssetId     *big.Int
	Verifier    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetVerified is a free log retrieval operation binding the contract event 0xa42b8d3300f0dcf5b507c7496082291ffb40df98da1a4565fdf6cc4c3eb1f47a.
//
// Solidity: event AssetVerified(address indexed orgContract, uint256 indexed assetId, address indexed verifier)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) FilterAssetVerified(opts *bind.FilterOpts, orgContract []common.Address, assetId []*big.Int, verifier []common.Address) (*HauskaAssetRegistryAssetVerifiedIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.FilterLogs(opts, "AssetVerified", orgContractRule, assetIdRule, verifierRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryAssetVerifiedIterator{contract: _HauskaAssetRegistry.contract, event: "AssetVerified", logs: logs, sub: sub}, nil
}

// WatchAssetVerified is a free log subscription operation binding the contract event 0xa42b8d3300f0dcf5b507c7496082291ffb40df98da1a4565fdf6cc4c3eb1f47a.
//
// Solidity: event AssetVerified(address indexed orgContract, uint256 indexed assetId, address indexed verifier)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) WatchAssetVerified(opts *bind.WatchOpts, sink chan<- *HauskaAssetRegistryAssetVerified, orgContract []common.Address, assetId []*big.Int, verifier []common.Address) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.WatchLogs(opts, "AssetVerified", orgContractRule, assetIdRule, verifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetRegistryAssetVerified)
				if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetVerified is a log parse operation binding the contract event 0xa42b8d3300f0dcf5b507c7496082291ffb40df98da1a4565fdf6cc4c3eb1f47a.
//
// Solidity: event AssetVerified(address indexed orgContract, uint256 indexed assetId, address indexed verifier)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) ParseAssetVerified(log types.Log) (*HauskaAssetRegistryAssetVerified, error) {
	event := new(HauskaAssetRegistryAssetVerified)
	if err := _HauskaAssetRegistry.contract.UnpackLog(event, "AssetVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryRoleAdminChangedIterator struct {
	Event *HauskaAssetRegistryRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HauskaAssetRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetRegistryRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HauskaAssetRegistryRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HauskaAssetRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HauskaAssetRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryRoleAdminChangedIterator{contract: _HauskaAssetRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HauskaAssetRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetRegistryRoleAdminChanged)
				if err := _HauskaAssetRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*HauskaAssetRegistryRoleAdminChanged, error) {
	event := new(HauskaAssetRegistryRoleAdminChanged)
	if err := _HauskaAssetRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryRoleGrantedIterator struct {
	Event *HauskaAssetRegistryRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HauskaAssetRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetRegistryRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HauskaAssetRegistryRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HauskaAssetRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetRegistryRoleGranted represents a RoleGranted event raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaAssetRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryRoleGrantedIterator{contract: _HauskaAssetRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HauskaAssetRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetRegistryRoleGranted)
				if err := _HauskaAssetRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) ParseRoleGranted(log types.Log) (*HauskaAssetRegistryRoleGranted, error) {
	event := new(HauskaAssetRegistryRoleGranted)
	if err := _HauskaAssetRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryRoleRevokedIterator struct {
	Event *HauskaAssetRegistryRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HauskaAssetRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetRegistryRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HauskaAssetRegistryRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HauskaAssetRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetRegistryRoleRevoked represents a RoleRevoked event raised by the HauskaAssetRegistry contract.
type HauskaAssetRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaAssetRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetRegistryRoleRevokedIterator{contract: _HauskaAssetRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HauskaAssetRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _HauskaAssetRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetRegistryRoleRevoked)
				if err := _HauskaAssetRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetRegistry *HauskaAssetRegistryFilterer) ParseRoleRevoked(log types.Log) (*HauskaAssetRegistryRoleRevoked, error) {
	event := new(HauskaAssetRegistryRoleRevoked)
	if err := _HauskaAssetRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
