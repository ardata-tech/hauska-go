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

// HauskaAssetNFTMetaData contains all meta data concerning the HauskaAssetNFT contract.
var HauskaAssetNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AssetNFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AssetNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_REGISTRY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetHasNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getTokenIdForAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetRegistry\",\"type\":\"address\"}],\"name\":\"grantAssetRegistryRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMinterRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintAssetNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToAssetId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToOrgContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HauskaAssetNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use HauskaAssetNFTMetaData.ABI instead.
var HauskaAssetNFTABI = HauskaAssetNFTMetaData.ABI

// HauskaAssetNFT is an auto generated Go binding around an Ethereum contract.
type HauskaAssetNFT struct {
	HauskaAssetNFTCaller     // Read-only binding to the contract
	HauskaAssetNFTTransactor // Write-only binding to the contract
	HauskaAssetNFTFilterer   // Log filterer for contract events
}

// HauskaAssetNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type HauskaAssetNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaAssetNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HauskaAssetNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaAssetNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HauskaAssetNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaAssetNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HauskaAssetNFTSession struct {
	Contract     *HauskaAssetNFT   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HauskaAssetNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HauskaAssetNFTCallerSession struct {
	Contract *HauskaAssetNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HauskaAssetNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HauskaAssetNFTTransactorSession struct {
	Contract     *HauskaAssetNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HauskaAssetNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type HauskaAssetNFTRaw struct {
	Contract *HauskaAssetNFT // Generic contract binding to access the raw methods on
}

// HauskaAssetNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HauskaAssetNFTCallerRaw struct {
	Contract *HauskaAssetNFTCaller // Generic read-only contract binding to access the raw methods on
}

// HauskaAssetNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HauskaAssetNFTTransactorRaw struct {
	Contract *HauskaAssetNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHauskaAssetNFT creates a new instance of HauskaAssetNFT, bound to a specific deployed contract.
func NewHauskaAssetNFT(address common.Address, backend bind.ContractBackend) (*HauskaAssetNFT, error) {
	contract, err := bindHauskaAssetNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFT{HauskaAssetNFTCaller: HauskaAssetNFTCaller{contract: contract}, HauskaAssetNFTTransactor: HauskaAssetNFTTransactor{contract: contract}, HauskaAssetNFTFilterer: HauskaAssetNFTFilterer{contract: contract}}, nil
}

// NewHauskaAssetNFTCaller creates a new read-only instance of HauskaAssetNFT, bound to a specific deployed contract.
func NewHauskaAssetNFTCaller(address common.Address, caller bind.ContractCaller) (*HauskaAssetNFTCaller, error) {
	contract, err := bindHauskaAssetNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTCaller{contract: contract}, nil
}

// NewHauskaAssetNFTTransactor creates a new write-only instance of HauskaAssetNFT, bound to a specific deployed contract.
func NewHauskaAssetNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*HauskaAssetNFTTransactor, error) {
	contract, err := bindHauskaAssetNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTTransactor{contract: contract}, nil
}

// NewHauskaAssetNFTFilterer creates a new log filterer instance of HauskaAssetNFT, bound to a specific deployed contract.
func NewHauskaAssetNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*HauskaAssetNFTFilterer, error) {
	contract, err := bindHauskaAssetNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTFilterer{contract: contract}, nil
}

// bindHauskaAssetNFT binds a generic wrapper to an already deployed contract.
func bindHauskaAssetNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HauskaAssetNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaAssetNFT *HauskaAssetNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaAssetNFT.Contract.HauskaAssetNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaAssetNFT *HauskaAssetNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.HauskaAssetNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaAssetNFT *HauskaAssetNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.HauskaAssetNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaAssetNFT *HauskaAssetNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaAssetNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaAssetNFT *HauskaAssetNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaAssetNFT *HauskaAssetNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.contract.Transact(opts, method, params...)
}

// ASSETREGISTRYROLE is a free data retrieval call binding the contract method 0x0680f211.
//
// Solidity: function ASSET_REGISTRY_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) ASSETREGISTRYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "ASSET_REGISTRY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETREGISTRYROLE is a free data retrieval call binding the contract method 0x0680f211.
//
// Solidity: function ASSET_REGISTRY_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTSession) ASSETREGISTRYROLE() ([32]byte, error) {
	return _HauskaAssetNFT.Contract.ASSETREGISTRYROLE(&_HauskaAssetNFT.CallOpts)
}

// ASSETREGISTRYROLE is a free data retrieval call binding the contract method 0x0680f211.
//
// Solidity: function ASSET_REGISTRY_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) ASSETREGISTRYROLE() ([32]byte, error) {
	return _HauskaAssetNFT.Contract.ASSETREGISTRYROLE(&_HauskaAssetNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaAssetNFT.Contract.DEFAULTADMINROLE(&_HauskaAssetNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaAssetNFT.Contract.DEFAULTADMINROLE(&_HauskaAssetNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTSession) MINTERROLE() ([32]byte, error) {
	return _HauskaAssetNFT.Contract.MINTERROLE(&_HauskaAssetNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) MINTERROLE() ([32]byte, error) {
	return _HauskaAssetNFT.Contract.MINTERROLE(&_HauskaAssetNFT.CallOpts)
}

// AssetHasNFT is a free data retrieval call binding the contract method 0xafc9d472.
//
// Solidity: function assetHasNFT(address orgContract, uint256 assetId) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) AssetHasNFT(opts *bind.CallOpts, orgContract common.Address, assetId *big.Int) (bool, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "assetHasNFT", orgContract, assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetHasNFT is a free data retrieval call binding the contract method 0xafc9d472.
//
// Solidity: function assetHasNFT(address orgContract, uint256 assetId) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTSession) AssetHasNFT(orgContract common.Address, assetId *big.Int) (bool, error) {
	return _HauskaAssetNFT.Contract.AssetHasNFT(&_HauskaAssetNFT.CallOpts, orgContract, assetId)
}

// AssetHasNFT is a free data retrieval call binding the contract method 0xafc9d472.
//
// Solidity: function assetHasNFT(address orgContract, uint256 assetId) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) AssetHasNFT(orgContract common.Address, assetId *big.Int) (bool, error) {
	return _HauskaAssetNFT.Contract.AssetHasNFT(&_HauskaAssetNFT.CallOpts, orgContract, assetId)
}

// AssetToTokenId is a free data retrieval call binding the contract method 0xcdf2db12.
//
// Solidity: function assetToTokenId(address , uint256 ) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) AssetToTokenId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "assetToTokenId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetToTokenId is a free data retrieval call binding the contract method 0xcdf2db12.
//
// Solidity: function assetToTokenId(address , uint256 ) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTSession) AssetToTokenId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.AssetToTokenId(&_HauskaAssetNFT.CallOpts, arg0, arg1)
}

// AssetToTokenId is a free data retrieval call binding the contract method 0xcdf2db12.
//
// Solidity: function assetToTokenId(address , uint256 ) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) AssetToTokenId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.AssetToTokenId(&_HauskaAssetNFT.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.BalanceOf(&_HauskaAssetNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.BalanceOf(&_HauskaAssetNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _HauskaAssetNFT.Contract.GetApproved(&_HauskaAssetNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _HauskaAssetNFT.Contract.GetApproved(&_HauskaAssetNFT.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaAssetNFT.Contract.GetRoleAdmin(&_HauskaAssetNFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaAssetNFT.Contract.GetRoleAdmin(&_HauskaAssetNFT.CallOpts, role)
}

// GetTokenIdForAsset is a free data retrieval call binding the contract method 0xc82c41e5.
//
// Solidity: function getTokenIdForAsset(address orgContract, uint256 assetId) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) GetTokenIdForAsset(opts *bind.CallOpts, orgContract common.Address, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "getTokenIdForAsset", orgContract, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenIdForAsset is a free data retrieval call binding the contract method 0xc82c41e5.
//
// Solidity: function getTokenIdForAsset(address orgContract, uint256 assetId) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTSession) GetTokenIdForAsset(orgContract common.Address, assetId *big.Int) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.GetTokenIdForAsset(&_HauskaAssetNFT.CallOpts, orgContract, assetId)
}

// GetTokenIdForAsset is a free data retrieval call binding the contract method 0xc82c41e5.
//
// Solidity: function getTokenIdForAsset(address orgContract, uint256 assetId) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) GetTokenIdForAsset(orgContract common.Address, assetId *big.Int) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.GetTokenIdForAsset(&_HauskaAssetNFT.CallOpts, orgContract, assetId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaAssetNFT.Contract.HasRole(&_HauskaAssetNFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaAssetNFT.Contract.HasRole(&_HauskaAssetNFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _HauskaAssetNFT.Contract.IsApprovedForAll(&_HauskaAssetNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _HauskaAssetNFT.Contract.IsApprovedForAll(&_HauskaAssetNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTSession) Name() (string, error) {
	return _HauskaAssetNFT.Contract.Name(&_HauskaAssetNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) Name() (string, error) {
	return _HauskaAssetNFT.Contract.Name(&_HauskaAssetNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _HauskaAssetNFT.Contract.OwnerOf(&_HauskaAssetNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _HauskaAssetNFT.Contract.OwnerOf(&_HauskaAssetNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaAssetNFT.Contract.SupportsInterface(&_HauskaAssetNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaAssetNFT.Contract.SupportsInterface(&_HauskaAssetNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTSession) Symbol() (string, error) {
	return _HauskaAssetNFT.Contract.Symbol(&_HauskaAssetNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) Symbol() (string, error) {
	return _HauskaAssetNFT.Contract.Symbol(&_HauskaAssetNFT.CallOpts)
}

// TokenToAssetId is a free data retrieval call binding the contract method 0x51cf0ce8.
//
// Solidity: function tokenToAssetId(uint256 ) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) TokenToAssetId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "tokenToAssetId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToAssetId is a free data retrieval call binding the contract method 0x51cf0ce8.
//
// Solidity: function tokenToAssetId(uint256 ) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTSession) TokenToAssetId(arg0 *big.Int) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.TokenToAssetId(&_HauskaAssetNFT.CallOpts, arg0)
}

// TokenToAssetId is a free data retrieval call binding the contract method 0x51cf0ce8.
//
// Solidity: function tokenToAssetId(uint256 ) view returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) TokenToAssetId(arg0 *big.Int) (*big.Int, error) {
	return _HauskaAssetNFT.Contract.TokenToAssetId(&_HauskaAssetNFT.CallOpts, arg0)
}

// TokenToOrgContract is a free data retrieval call binding the contract method 0x7ca0d35a.
//
// Solidity: function tokenToOrgContract(uint256 ) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) TokenToOrgContract(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "tokenToOrgContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenToOrgContract is a free data retrieval call binding the contract method 0x7ca0d35a.
//
// Solidity: function tokenToOrgContract(uint256 ) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTSession) TokenToOrgContract(arg0 *big.Int) (common.Address, error) {
	return _HauskaAssetNFT.Contract.TokenToOrgContract(&_HauskaAssetNFT.CallOpts, arg0)
}

// TokenToOrgContract is a free data retrieval call binding the contract method 0x7ca0d35a.
//
// Solidity: function tokenToOrgContract(uint256 ) view returns(address)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) TokenToOrgContract(arg0 *big.Int) (common.Address, error) {
	return _HauskaAssetNFT.Contract.TokenToOrgContract(&_HauskaAssetNFT.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _HauskaAssetNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _HauskaAssetNFT.Contract.TokenURI(&_HauskaAssetNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_HauskaAssetNFT *HauskaAssetNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _HauskaAssetNFT.Contract.TokenURI(&_HauskaAssetNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.Approve(&_HauskaAssetNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.Approve(&_HauskaAssetNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.Burn(&_HauskaAssetNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.Burn(&_HauskaAssetNFT.TransactOpts, tokenId)
}

// GrantAssetRegistryRole is a paid mutator transaction binding the contract method 0x5cb2330c.
//
// Solidity: function grantAssetRegistryRole(address assetRegistry) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) GrantAssetRegistryRole(opts *bind.TransactOpts, assetRegistry common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "grantAssetRegistryRole", assetRegistry)
}

// GrantAssetRegistryRole is a paid mutator transaction binding the contract method 0x5cb2330c.
//
// Solidity: function grantAssetRegistryRole(address assetRegistry) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) GrantAssetRegistryRole(assetRegistry common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.GrantAssetRegistryRole(&_HauskaAssetNFT.TransactOpts, assetRegistry)
}

// GrantAssetRegistryRole is a paid mutator transaction binding the contract method 0x5cb2330c.
//
// Solidity: function grantAssetRegistryRole(address assetRegistry) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) GrantAssetRegistryRole(assetRegistry common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.GrantAssetRegistryRole(&_HauskaAssetNFT.TransactOpts, assetRegistry)
}

// GrantMinterRole is a paid mutator transaction binding the contract method 0x3dd1eb61.
//
// Solidity: function grantMinterRole(address minter) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) GrantMinterRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "grantMinterRole", minter)
}

// GrantMinterRole is a paid mutator transaction binding the contract method 0x3dd1eb61.
//
// Solidity: function grantMinterRole(address minter) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) GrantMinterRole(minter common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.GrantMinterRole(&_HauskaAssetNFT.TransactOpts, minter)
}

// GrantMinterRole is a paid mutator transaction binding the contract method 0x3dd1eb61.
//
// Solidity: function grantMinterRole(address minter) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) GrantMinterRole(minter common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.GrantMinterRole(&_HauskaAssetNFT.TransactOpts, minter)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.GrantRole(&_HauskaAssetNFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.GrantRole(&_HauskaAssetNFT.TransactOpts, role, account)
}

// MintAssetNFT is a paid mutator transaction binding the contract method 0x702b076d.
//
// Solidity: function mintAssetNFT(address to, address orgContract, uint256 assetId, string uri) returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) MintAssetNFT(opts *bind.TransactOpts, to common.Address, orgContract common.Address, assetId *big.Int, uri string) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "mintAssetNFT", to, orgContract, assetId, uri)
}

// MintAssetNFT is a paid mutator transaction binding the contract method 0x702b076d.
//
// Solidity: function mintAssetNFT(address to, address orgContract, uint256 assetId, string uri) returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTSession) MintAssetNFT(to common.Address, orgContract common.Address, assetId *big.Int, uri string) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.MintAssetNFT(&_HauskaAssetNFT.TransactOpts, to, orgContract, assetId, uri)
}

// MintAssetNFT is a paid mutator transaction binding the contract method 0x702b076d.
//
// Solidity: function mintAssetNFT(address to, address orgContract, uint256 assetId, string uri) returns(uint256)
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) MintAssetNFT(to common.Address, orgContract common.Address, assetId *big.Int, uri string) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.MintAssetNFT(&_HauskaAssetNFT.TransactOpts, to, orgContract, assetId, uri)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.RenounceRole(&_HauskaAssetNFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.RenounceRole(&_HauskaAssetNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.RevokeRole(&_HauskaAssetNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.RevokeRole(&_HauskaAssetNFT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.SafeTransferFrom(&_HauskaAssetNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.SafeTransferFrom(&_HauskaAssetNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.SafeTransferFrom0(&_HauskaAssetNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.SafeTransferFrom0(&_HauskaAssetNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.SetApprovalForAll(&_HauskaAssetNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.SetApprovalForAll(&_HauskaAssetNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.TransferFrom(&_HauskaAssetNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.TransferFrom(&_HauskaAssetNFT.TransactOpts, from, to, tokenId)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string uri) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _HauskaAssetNFT.contract.Transact(opts, "updateTokenURI", tokenId, uri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string uri) returns()
func (_HauskaAssetNFT *HauskaAssetNFTSession) UpdateTokenURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.UpdateTokenURI(&_HauskaAssetNFT.TransactOpts, tokenId, uri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string uri) returns()
func (_HauskaAssetNFT *HauskaAssetNFTTransactorSession) UpdateTokenURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _HauskaAssetNFT.Contract.UpdateTokenURI(&_HauskaAssetNFT.TransactOpts, tokenId, uri)
}

// HauskaAssetNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTApprovalIterator struct {
	Event *HauskaAssetNFTApproval // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTApproval)
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
		it.Event = new(HauskaAssetNFTApproval)
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
func (it *HauskaAssetNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTApproval represents a Approval event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*HauskaAssetNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTApprovalIterator{contract: _HauskaAssetNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTApproval)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseApproval(log types.Log) (*HauskaAssetNFTApproval, error) {
	event := new(HauskaAssetNFTApproval)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTApprovalForAllIterator struct {
	Event *HauskaAssetNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTApprovalForAll)
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
		it.Event = new(HauskaAssetNFTApprovalForAll)
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
func (it *HauskaAssetNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTApprovalForAll represents a ApprovalForAll event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*HauskaAssetNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTApprovalForAllIterator{contract: _HauskaAssetNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTApprovalForAll)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseApprovalForAll(log types.Log) (*HauskaAssetNFTApprovalForAll, error) {
	event := new(HauskaAssetNFTApprovalForAll)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTAssetNFTMintedIterator is returned from FilterAssetNFTMinted and is used to iterate over the raw logs and unpacked data for AssetNFTMinted events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTAssetNFTMintedIterator struct {
	Event *HauskaAssetNFTAssetNFTMinted // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTAssetNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTAssetNFTMinted)
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
		it.Event = new(HauskaAssetNFTAssetNFTMinted)
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
func (it *HauskaAssetNFTAssetNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTAssetNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTAssetNFTMinted represents a AssetNFTMinted event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTAssetNFTMinted struct {
	OrgContract common.Address
	AssetId     *big.Int
	TokenId     *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetNFTMinted is a free log retrieval operation binding the contract event 0xd70c8babb3e4632b2d61a510607c94740a3876bab8ead3fd152b4daf08926188.
//
// Solidity: event AssetNFTMinted(address indexed orgContract, uint256 indexed assetId, uint256 indexed tokenId, address owner)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterAssetNFTMinted(opts *bind.FilterOpts, orgContract []common.Address, assetId []*big.Int, tokenId []*big.Int) (*HauskaAssetNFTAssetNFTMintedIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "AssetNFTMinted", orgContractRule, assetIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTAssetNFTMintedIterator{contract: _HauskaAssetNFT.contract, event: "AssetNFTMinted", logs: logs, sub: sub}, nil
}

// WatchAssetNFTMinted is a free log subscription operation binding the contract event 0xd70c8babb3e4632b2d61a510607c94740a3876bab8ead3fd152b4daf08926188.
//
// Solidity: event AssetNFTMinted(address indexed orgContract, uint256 indexed assetId, uint256 indexed tokenId, address owner)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchAssetNFTMinted(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTAssetNFTMinted, orgContract []common.Address, assetId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "AssetNFTMinted", orgContractRule, assetIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTAssetNFTMinted)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "AssetNFTMinted", log); err != nil {
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

// ParseAssetNFTMinted is a log parse operation binding the contract event 0xd70c8babb3e4632b2d61a510607c94740a3876bab8ead3fd152b4daf08926188.
//
// Solidity: event AssetNFTMinted(address indexed orgContract, uint256 indexed assetId, uint256 indexed tokenId, address owner)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseAssetNFTMinted(log types.Log) (*HauskaAssetNFTAssetNFTMinted, error) {
	event := new(HauskaAssetNFTAssetNFTMinted)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "AssetNFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTAssetNFTTransferredIterator is returned from FilterAssetNFTTransferred and is used to iterate over the raw logs and unpacked data for AssetNFTTransferred events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTAssetNFTTransferredIterator struct {
	Event *HauskaAssetNFTAssetNFTTransferred // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTAssetNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTAssetNFTTransferred)
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
		it.Event = new(HauskaAssetNFTAssetNFTTransferred)
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
func (it *HauskaAssetNFTAssetNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTAssetNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTAssetNFTTransferred represents a AssetNFTTransferred event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTAssetNFTTransferred struct {
	TokenId *big.Int
	From    common.Address
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetNFTTransferred is a free log retrieval operation binding the contract event 0xd70c55a7794811b5345eb8740e3bdc08fa0fc55e141360e6302148109e5c30d3.
//
// Solidity: event AssetNFTTransferred(uint256 indexed tokenId, address indexed from, address indexed to)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterAssetNFTTransferred(opts *bind.FilterOpts, tokenId []*big.Int, from []common.Address, to []common.Address) (*HauskaAssetNFTAssetNFTTransferredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "AssetNFTTransferred", tokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTAssetNFTTransferredIterator{contract: _HauskaAssetNFT.contract, event: "AssetNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchAssetNFTTransferred is a free log subscription operation binding the contract event 0xd70c55a7794811b5345eb8740e3bdc08fa0fc55e141360e6302148109e5c30d3.
//
// Solidity: event AssetNFTTransferred(uint256 indexed tokenId, address indexed from, address indexed to)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchAssetNFTTransferred(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTAssetNFTTransferred, tokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "AssetNFTTransferred", tokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTAssetNFTTransferred)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "AssetNFTTransferred", log); err != nil {
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

// ParseAssetNFTTransferred is a log parse operation binding the contract event 0xd70c55a7794811b5345eb8740e3bdc08fa0fc55e141360e6302148109e5c30d3.
//
// Solidity: event AssetNFTTransferred(uint256 indexed tokenId, address indexed from, address indexed to)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseAssetNFTTransferred(log types.Log) (*HauskaAssetNFTAssetNFTTransferred, error) {
	event := new(HauskaAssetNFTAssetNFTTransferred)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "AssetNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTBatchMetadataUpdateIterator struct {
	Event *HauskaAssetNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTBatchMetadataUpdate)
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
		it.Event = new(HauskaAssetNFTBatchMetadataUpdate)
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
func (it *HauskaAssetNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*HauskaAssetNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTBatchMetadataUpdateIterator{contract: _HauskaAssetNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTBatchMetadataUpdate)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*HauskaAssetNFTBatchMetadataUpdate, error) {
	event := new(HauskaAssetNFTBatchMetadataUpdate)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTMetadataUpdateIterator struct {
	Event *HauskaAssetNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTMetadataUpdate)
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
		it.Event = new(HauskaAssetNFTMetadataUpdate)
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
func (it *HauskaAssetNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTMetadataUpdate represents a MetadataUpdate event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*HauskaAssetNFTMetadataUpdateIterator, error) {

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTMetadataUpdateIterator{contract: _HauskaAssetNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTMetadataUpdate)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseMetadataUpdate(log types.Log) (*HauskaAssetNFTMetadataUpdate, error) {
	event := new(HauskaAssetNFTMetadataUpdate)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTRoleAdminChangedIterator struct {
	Event *HauskaAssetNFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTRoleAdminChanged)
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
		it.Event = new(HauskaAssetNFTRoleAdminChanged)
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
func (it *HauskaAssetNFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTRoleAdminChanged represents a RoleAdminChanged event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HauskaAssetNFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTRoleAdminChangedIterator{contract: _HauskaAssetNFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTRoleAdminChanged)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseRoleAdminChanged(log types.Log) (*HauskaAssetNFTRoleAdminChanged, error) {
	event := new(HauskaAssetNFTRoleAdminChanged)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTRoleGrantedIterator struct {
	Event *HauskaAssetNFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTRoleGranted)
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
		it.Event = new(HauskaAssetNFTRoleGranted)
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
func (it *HauskaAssetNFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTRoleGranted represents a RoleGranted event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaAssetNFTRoleGrantedIterator, error) {

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

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTRoleGrantedIterator{contract: _HauskaAssetNFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTRoleGranted)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseRoleGranted(log types.Log) (*HauskaAssetNFTRoleGranted, error) {
	event := new(HauskaAssetNFTRoleGranted)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTRoleRevokedIterator struct {
	Event *HauskaAssetNFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTRoleRevoked)
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
		it.Event = new(HauskaAssetNFTRoleRevoked)
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
func (it *HauskaAssetNFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTRoleRevoked represents a RoleRevoked event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaAssetNFTRoleRevokedIterator, error) {

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

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTRoleRevokedIterator{contract: _HauskaAssetNFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTRoleRevoked)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseRoleRevoked(log types.Log) (*HauskaAssetNFTRoleRevoked, error) {
	event := new(HauskaAssetNFTRoleRevoked)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaAssetNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HauskaAssetNFT contract.
type HauskaAssetNFTTransferIterator struct {
	Event *HauskaAssetNFTTransfer // Event containing the contract specifics and raw log

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
func (it *HauskaAssetNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaAssetNFTTransfer)
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
		it.Event = new(HauskaAssetNFTTransfer)
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
func (it *HauskaAssetNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaAssetNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaAssetNFTTransfer represents a Transfer event raised by the HauskaAssetNFT contract.
type HauskaAssetNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*HauskaAssetNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaAssetNFTTransferIterator{contract: _HauskaAssetNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HauskaAssetNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HauskaAssetNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaAssetNFTTransfer)
				if err := _HauskaAssetNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_HauskaAssetNFT *HauskaAssetNFTFilterer) ParseTransfer(log types.Log) (*HauskaAssetNFTTransfer, error) {
	event := new(HauskaAssetNFTTransfer)
	if err := _HauskaAssetNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
