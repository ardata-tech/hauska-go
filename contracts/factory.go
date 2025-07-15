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

// HauskaContractFactoryMetaData contains all meta data concerning the HauskaContractFactory contract.
var HauskaContractFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdcToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"principal\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"integrationPartner\",\"type\":\"address\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"licenseManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"groupManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revenueDistributor\",\"type\":\"address\"}],\"name\":\"ModulesDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"hauskaFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"integratorFee\",\"type\":\"uint32\"}],\"name\":\"PlatformFeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"principalEntity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"integrationPartner\",\"type\":\"address\"}],\"name\":\"createContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssetNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"principal\",\"type\":\"address\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_licenseManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_groupManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_revenueDistributor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"getNumberOfCreators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfOrganizations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFees\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"integratorFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"hauskaFee\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"groupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hauskaFeePct\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"integratorFeePct\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isValidOrgContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"licenseManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"principalToContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registerAssetHash\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractToRemove\",\"type\":\"address\"}],\"name\":\"removeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenueDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_hauskaFeePct\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_integratorFeePct\",\"type\":\"uint32\"}],\"name\":\"updatePlatformFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HauskaContractFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use HauskaContractFactoryMetaData.ABI instead.
var HauskaContractFactoryABI = HauskaContractFactoryMetaData.ABI

// HauskaContractFactory is an auto generated Go binding around an Ethereum contract.
type HauskaContractFactory struct {
	HauskaContractFactoryCaller     // Read-only binding to the contract
	HauskaContractFactoryTransactor // Write-only binding to the contract
	HauskaContractFactoryFilterer   // Log filterer for contract events
}

// HauskaContractFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HauskaContractFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaContractFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HauskaContractFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaContractFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HauskaContractFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaContractFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HauskaContractFactorySession struct {
	Contract     *HauskaContractFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HauskaContractFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HauskaContractFactoryCallerSession struct {
	Contract *HauskaContractFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// HauskaContractFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HauskaContractFactoryTransactorSession struct {
	Contract     *HauskaContractFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// HauskaContractFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HauskaContractFactoryRaw struct {
	Contract *HauskaContractFactory // Generic contract binding to access the raw methods on
}

// HauskaContractFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HauskaContractFactoryCallerRaw struct {
	Contract *HauskaContractFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// HauskaContractFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HauskaContractFactoryTransactorRaw struct {
	Contract *HauskaContractFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHauskaContractFactory creates a new instance of HauskaContractFactory, bound to a specific deployed contract.
func NewHauskaContractFactory(address common.Address, backend bind.ContractBackend) (*HauskaContractFactory, error) {
	contract, err := bindHauskaContractFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactory{HauskaContractFactoryCaller: HauskaContractFactoryCaller{contract: contract}, HauskaContractFactoryTransactor: HauskaContractFactoryTransactor{contract: contract}, HauskaContractFactoryFilterer: HauskaContractFactoryFilterer{contract: contract}}, nil
}

// NewHauskaContractFactoryCaller creates a new read-only instance of HauskaContractFactory, bound to a specific deployed contract.
func NewHauskaContractFactoryCaller(address common.Address, caller bind.ContractCaller) (*HauskaContractFactoryCaller, error) {
	contract, err := bindHauskaContractFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryCaller{contract: contract}, nil
}

// NewHauskaContractFactoryTransactor creates a new write-only instance of HauskaContractFactory, bound to a specific deployed contract.
func NewHauskaContractFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*HauskaContractFactoryTransactor, error) {
	contract, err := bindHauskaContractFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryTransactor{contract: contract}, nil
}

// NewHauskaContractFactoryFilterer creates a new log filterer instance of HauskaContractFactory, bound to a specific deployed contract.
func NewHauskaContractFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*HauskaContractFactoryFilterer, error) {
	contract, err := bindHauskaContractFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryFilterer{contract: contract}, nil
}

// bindHauskaContractFactory binds a generic wrapper to an already deployed contract.
func bindHauskaContractFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HauskaContractFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaContractFactory *HauskaContractFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaContractFactory.Contract.HauskaContractFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaContractFactory *HauskaContractFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.HauskaContractFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaContractFactory *HauskaContractFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.HauskaContractFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaContractFactory *HauskaContractFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaContractFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaContractFactory *HauskaContractFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaContractFactory *HauskaContractFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactoryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactorySession) ADMINROLE() ([32]byte, error) {
	return _HauskaContractFactory.Contract.ADMINROLE(&_HauskaContractFactory.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) ADMINROLE() ([32]byte, error) {
	return _HauskaContractFactory.Contract.ADMINROLE(&_HauskaContractFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaContractFactory.Contract.DEFAULTADMINROLE(&_HauskaContractFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaContractFactory.Contract.DEFAULTADMINROLE(&_HauskaContractFactory.CallOpts)
}

// AllContracts is a free data retrieval call binding the contract method 0xe54e26b9.
//
// Solidity: function allContracts(uint256 ) view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) AllContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "allContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllContracts is a free data retrieval call binding the contract method 0xe54e26b9.
//
// Solidity: function allContracts(uint256 ) view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) AllContracts(arg0 *big.Int) (common.Address, error) {
	return _HauskaContractFactory.Contract.AllContracts(&_HauskaContractFactory.CallOpts, arg0)
}

// AllContracts is a free data retrieval call binding the contract method 0xe54e26b9.
//
// Solidity: function allContracts(uint256 ) view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) AllContracts(arg0 *big.Int) (common.Address, error) {
	return _HauskaContractFactory.Contract.AllContracts(&_HauskaContractFactory.CallOpts, arg0)
}

// AssetNFT is a free data retrieval call binding the contract method 0xac6a258d.
//
// Solidity: function assetNFT() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) AssetNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "assetNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetNFT is a free data retrieval call binding the contract method 0xac6a258d.
//
// Solidity: function assetNFT() view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) AssetNFT() (common.Address, error) {
	return _HauskaContractFactory.Contract.AssetNFT(&_HauskaContractFactory.CallOpts)
}

// AssetNFT is a free data retrieval call binding the contract method 0xac6a258d.
//
// Solidity: function assetNFT() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) AssetNFT() (common.Address, error) {
	return _HauskaContractFactory.Contract.AssetNFT(&_HauskaContractFactory.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) AssetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "assetRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) AssetRegistry() (common.Address, error) {
	return _HauskaContractFactory.Contract.AssetRegistry(&_HauskaContractFactory.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) AssetRegistry() (common.Address, error) {
	return _HauskaContractFactory.Contract.AssetRegistry(&_HauskaContractFactory.CallOpts)
}

// ContractIndex is a free data retrieval call binding the contract method 0x4cfb222d.
//
// Solidity: function contractIndex(address ) view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactoryCaller) ContractIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "contractIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractIndex is a free data retrieval call binding the contract method 0x4cfb222d.
//
// Solidity: function contractIndex(address ) view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactorySession) ContractIndex(arg0 common.Address) (*big.Int, error) {
	return _HauskaContractFactory.Contract.ContractIndex(&_HauskaContractFactory.CallOpts, arg0)
}

// ContractIndex is a free data retrieval call binding the contract method 0x4cfb222d.
//
// Solidity: function contractIndex(address ) view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) ContractIndex(arg0 common.Address) (*big.Int, error) {
	return _HauskaContractFactory.Contract.ContractIndex(&_HauskaContractFactory.CallOpts, arg0)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns(address[])
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetAllContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getAllContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns(address[])
func (_HauskaContractFactory *HauskaContractFactorySession) GetAllContracts() ([]common.Address, error) {
	return _HauskaContractFactory.Contract.GetAllContracts(&_HauskaContractFactory.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns(address[])
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetAllContracts() ([]common.Address, error) {
	return _HauskaContractFactory.Contract.GetAllContracts(&_HauskaContractFactory.CallOpts)
}

// GetAssetNFT is a free data retrieval call binding the contract method 0xbd1f3c29.
//
// Solidity: function getAssetNFT() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetAssetNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getAssetNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAssetNFT is a free data retrieval call binding the contract method 0xbd1f3c29.
//
// Solidity: function getAssetNFT() view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) GetAssetNFT() (common.Address, error) {
	return _HauskaContractFactory.Contract.GetAssetNFT(&_HauskaContractFactory.CallOpts)
}

// GetAssetNFT is a free data retrieval call binding the contract method 0xbd1f3c29.
//
// Solidity: function getAssetNFT() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetAssetNFT() (common.Address, error) {
	return _HauskaContractFactory.Contract.GetAssetNFT(&_HauskaContractFactory.CallOpts)
}

// GetContract is a free data retrieval call binding the contract method 0x0ad1c2fa.
//
// Solidity: function getContract(address principal) view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetContract(opts *bind.CallOpts, principal common.Address) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getContract", principal)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x0ad1c2fa.
//
// Solidity: function getContract(address principal) view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) GetContract(principal common.Address) (common.Address, error) {
	return _HauskaContractFactory.Contract.GetContract(&_HauskaContractFactory.CallOpts, principal)
}

// GetContract is a free data retrieval call binding the contract method 0x0ad1c2fa.
//
// Solidity: function getContract(address principal) view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetContract(principal common.Address) (common.Address, error) {
	return _HauskaContractFactory.Contract.GetContract(&_HauskaContractFactory.CallOpts, principal)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address _licenseManager, address _assetRegistry, address _groupManager, address _revenueDistributor)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetModules(opts *bind.CallOpts) (struct {
	LicenseManager     common.Address
	AssetRegistry      common.Address
	GroupManager       common.Address
	RevenueDistributor common.Address
}, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getModules")

	outstruct := new(struct {
		LicenseManager     common.Address
		AssetRegistry      common.Address
		GroupManager       common.Address
		RevenueDistributor common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LicenseManager = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AssetRegistry = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.GroupManager = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.RevenueDistributor = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address _licenseManager, address _assetRegistry, address _groupManager, address _revenueDistributor)
func (_HauskaContractFactory *HauskaContractFactorySession) GetModules() (struct {
	LicenseManager     common.Address
	AssetRegistry      common.Address
	GroupManager       common.Address
	RevenueDistributor common.Address
}, error) {
	return _HauskaContractFactory.Contract.GetModules(&_HauskaContractFactory.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address _licenseManager, address _assetRegistry, address _groupManager, address _revenueDistributor)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetModules() (struct {
	LicenseManager     common.Address
	AssetRegistry      common.Address
	GroupManager       common.Address
	RevenueDistributor common.Address
}, error) {
	return _HauskaContractFactory.Contract.GetModules(&_HauskaContractFactory.CallOpts)
}

// GetNumberOfCreators is a free data retrieval call binding the contract method 0x301a86c1.
//
// Solidity: function getNumberOfCreators(address orgContract) view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetNumberOfCreators(opts *bind.CallOpts, orgContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getNumberOfCreators", orgContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfCreators is a free data retrieval call binding the contract method 0x301a86c1.
//
// Solidity: function getNumberOfCreators(address orgContract) view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactorySession) GetNumberOfCreators(orgContract common.Address) (*big.Int, error) {
	return _HauskaContractFactory.Contract.GetNumberOfCreators(&_HauskaContractFactory.CallOpts, orgContract)
}

// GetNumberOfCreators is a free data retrieval call binding the contract method 0x301a86c1.
//
// Solidity: function getNumberOfCreators(address orgContract) view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetNumberOfCreators(orgContract common.Address) (*big.Int, error) {
	return _HauskaContractFactory.Contract.GetNumberOfCreators(&_HauskaContractFactory.CallOpts, orgContract)
}

// GetNumberOfOrganizations is a free data retrieval call binding the contract method 0xbd3a694e.
//
// Solidity: function getNumberOfOrganizations() view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetNumberOfOrganizations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getNumberOfOrganizations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfOrganizations is a free data retrieval call binding the contract method 0xbd3a694e.
//
// Solidity: function getNumberOfOrganizations() view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactorySession) GetNumberOfOrganizations() (*big.Int, error) {
	return _HauskaContractFactory.Contract.GetNumberOfOrganizations(&_HauskaContractFactory.CallOpts)
}

// GetNumberOfOrganizations is a free data retrieval call binding the contract method 0xbd3a694e.
//
// Solidity: function getNumberOfOrganizations() view returns(uint256)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetNumberOfOrganizations() (*big.Int, error) {
	return _HauskaContractFactory.Contract.GetNumberOfOrganizations(&_HauskaContractFactory.CallOpts)
}

// GetPlatformFees is a free data retrieval call binding the contract method 0x55237200.
//
// Solidity: function getPlatformFees() view returns(uint32 integratorFee, uint32 hauskaFee)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetPlatformFees(opts *bind.CallOpts) (struct {
	IntegratorFee uint32
	HauskaFee     uint32
}, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getPlatformFees")

	outstruct := new(struct {
		IntegratorFee uint32
		HauskaFee     uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IntegratorFee = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.HauskaFee = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetPlatformFees is a free data retrieval call binding the contract method 0x55237200.
//
// Solidity: function getPlatformFees() view returns(uint32 integratorFee, uint32 hauskaFee)
func (_HauskaContractFactory *HauskaContractFactorySession) GetPlatformFees() (struct {
	IntegratorFee uint32
	HauskaFee     uint32
}, error) {
	return _HauskaContractFactory.Contract.GetPlatformFees(&_HauskaContractFactory.CallOpts)
}

// GetPlatformFees is a free data retrieval call binding the contract method 0x55237200.
//
// Solidity: function getPlatformFees() view returns(uint32 integratorFee, uint32 hauskaFee)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetPlatformFees() (struct {
	IntegratorFee uint32
	HauskaFee     uint32
}, error) {
	return _HauskaContractFactory.Contract.GetPlatformFees(&_HauskaContractFactory.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaContractFactory.Contract.GetRoleAdmin(&_HauskaContractFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaContractFactory.Contract.GetRoleAdmin(&_HauskaContractFactory.CallOpts, role)
}

// GroupManager is a free data retrieval call binding the contract method 0xb26bc483.
//
// Solidity: function groupManager() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) GroupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "groupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GroupManager is a free data retrieval call binding the contract method 0xb26bc483.
//
// Solidity: function groupManager() view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) GroupManager() (common.Address, error) {
	return _HauskaContractFactory.Contract.GroupManager(&_HauskaContractFactory.CallOpts)
}

// GroupManager is a free data retrieval call binding the contract method 0xb26bc483.
//
// Solidity: function groupManager() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) GroupManager() (common.Address, error) {
	return _HauskaContractFactory.Contract.GroupManager(&_HauskaContractFactory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaContractFactory.Contract.HasRole(&_HauskaContractFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaContractFactory.Contract.HasRole(&_HauskaContractFactory.CallOpts, role, account)
}

// HauskaFeePct is a free data retrieval call binding the contract method 0x605c346a.
//
// Solidity: function hauskaFeePct() view returns(uint32)
func (_HauskaContractFactory *HauskaContractFactoryCaller) HauskaFeePct(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "hauskaFeePct")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// HauskaFeePct is a free data retrieval call binding the contract method 0x605c346a.
//
// Solidity: function hauskaFeePct() view returns(uint32)
func (_HauskaContractFactory *HauskaContractFactorySession) HauskaFeePct() (uint32, error) {
	return _HauskaContractFactory.Contract.HauskaFeePct(&_HauskaContractFactory.CallOpts)
}

// HauskaFeePct is a free data retrieval call binding the contract method 0x605c346a.
//
// Solidity: function hauskaFeePct() view returns(uint32)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) HauskaFeePct() (uint32, error) {
	return _HauskaContractFactory.Contract.HauskaFeePct(&_HauskaContractFactory.CallOpts)
}

// IntegratorFeePct is a free data retrieval call binding the contract method 0x4520ca75.
//
// Solidity: function integratorFeePct() view returns(uint32)
func (_HauskaContractFactory *HauskaContractFactoryCaller) IntegratorFeePct(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "integratorFeePct")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// IntegratorFeePct is a free data retrieval call binding the contract method 0x4520ca75.
//
// Solidity: function integratorFeePct() view returns(uint32)
func (_HauskaContractFactory *HauskaContractFactorySession) IntegratorFeePct() (uint32, error) {
	return _HauskaContractFactory.Contract.IntegratorFeePct(&_HauskaContractFactory.CallOpts)
}

// IntegratorFeePct is a free data retrieval call binding the contract method 0x4520ca75.
//
// Solidity: function integratorFeePct() view returns(uint32)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) IntegratorFeePct() (uint32, error) {
	return _HauskaContractFactory.Contract.IntegratorFeePct(&_HauskaContractFactory.CallOpts)
}

// IsValidOrgContract is a free data retrieval call binding the contract method 0xfbed1ddb.
//
// Solidity: function isValidOrgContract(address contractAddress) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCaller) IsValidOrgContract(opts *bind.CallOpts, contractAddress common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "isValidOrgContract", contractAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOrgContract is a free data retrieval call binding the contract method 0xfbed1ddb.
//
// Solidity: function isValidOrgContract(address contractAddress) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactorySession) IsValidOrgContract(contractAddress common.Address) (bool, error) {
	return _HauskaContractFactory.Contract.IsValidOrgContract(&_HauskaContractFactory.CallOpts, contractAddress)
}

// IsValidOrgContract is a free data retrieval call binding the contract method 0xfbed1ddb.
//
// Solidity: function isValidOrgContract(address contractAddress) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) IsValidOrgContract(contractAddress common.Address) (bool, error) {
	return _HauskaContractFactory.Contract.IsValidOrgContract(&_HauskaContractFactory.CallOpts, contractAddress)
}

// LicenseManager is a free data retrieval call binding the contract method 0x0ad51144.
//
// Solidity: function licenseManager() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) LicenseManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "licenseManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseManager is a free data retrieval call binding the contract method 0x0ad51144.
//
// Solidity: function licenseManager() view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) LicenseManager() (common.Address, error) {
	return _HauskaContractFactory.Contract.LicenseManager(&_HauskaContractFactory.CallOpts)
}

// LicenseManager is a free data retrieval call binding the contract method 0x0ad51144.
//
// Solidity: function licenseManager() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) LicenseManager() (common.Address, error) {
	return _HauskaContractFactory.Contract.LicenseManager(&_HauskaContractFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HauskaContractFactory *HauskaContractFactorySession) Paused() (bool, error) {
	return _HauskaContractFactory.Contract.Paused(&_HauskaContractFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) Paused() (bool, error) {
	return _HauskaContractFactory.Contract.Paused(&_HauskaContractFactory.CallOpts)
}

// PrincipalToContract is a free data retrieval call binding the contract method 0xa830164b.
//
// Solidity: function principalToContract(address ) view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) PrincipalToContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "principalToContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrincipalToContract is a free data retrieval call binding the contract method 0xa830164b.
//
// Solidity: function principalToContract(address ) view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) PrincipalToContract(arg0 common.Address) (common.Address, error) {
	return _HauskaContractFactory.Contract.PrincipalToContract(&_HauskaContractFactory.CallOpts, arg0)
}

// PrincipalToContract is a free data retrieval call binding the contract method 0xa830164b.
//
// Solidity: function principalToContract(address ) view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) PrincipalToContract(arg0 common.Address) (common.Address, error) {
	return _HauskaContractFactory.Contract.PrincipalToContract(&_HauskaContractFactory.CallOpts, arg0)
}

// RegisterAssetHash is a free data retrieval call binding the contract method 0x62f2c1a3.
//
// Solidity: function registerAssetHash(bytes32 , string ) pure returns()
func (_HauskaContractFactory *HauskaContractFactoryCaller) RegisterAssetHash(opts *bind.CallOpts, arg0 [32]byte, arg1 string) error {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "registerAssetHash", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RegisterAssetHash is a free data retrieval call binding the contract method 0x62f2c1a3.
//
// Solidity: function registerAssetHash(bytes32 , string ) pure returns()
func (_HauskaContractFactory *HauskaContractFactorySession) RegisterAssetHash(arg0 [32]byte, arg1 string) error {
	return _HauskaContractFactory.Contract.RegisterAssetHash(&_HauskaContractFactory.CallOpts, arg0, arg1)
}

// RegisterAssetHash is a free data retrieval call binding the contract method 0x62f2c1a3.
//
// Solidity: function registerAssetHash(bytes32 , string ) pure returns()
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) RegisterAssetHash(arg0 [32]byte, arg1 string) error {
	return _HauskaContractFactory.Contract.RegisterAssetHash(&_HauskaContractFactory.CallOpts, arg0, arg1)
}

// RevenueDistributor is a free data retrieval call binding the contract method 0xeae0a488.
//
// Solidity: function revenueDistributor() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) RevenueDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "revenueDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenueDistributor is a free data retrieval call binding the contract method 0xeae0a488.
//
// Solidity: function revenueDistributor() view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) RevenueDistributor() (common.Address, error) {
	return _HauskaContractFactory.Contract.RevenueDistributor(&_HauskaContractFactory.CallOpts)
}

// RevenueDistributor is a free data retrieval call binding the contract method 0xeae0a488.
//
// Solidity: function revenueDistributor() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) RevenueDistributor() (common.Address, error) {
	return _HauskaContractFactory.Contract.RevenueDistributor(&_HauskaContractFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaContractFactory.Contract.SupportsInterface(&_HauskaContractFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaContractFactory.Contract.SupportsInterface(&_HauskaContractFactory.CallOpts, interfaceId)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) UsdcToken() (common.Address, error) {
	return _HauskaContractFactory.Contract.UsdcToken(&_HauskaContractFactory.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) UsdcToken() (common.Address, error) {
	return _HauskaContractFactory.Contract.UsdcToken(&_HauskaContractFactory.CallOpts)
}

// ValidContracts is a free data retrieval call binding the contract method 0x487f6630.
//
// Solidity: function validContracts(address ) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCaller) ValidContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaContractFactory.contract.Call(opts, &out, "validContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidContracts is a free data retrieval call binding the contract method 0x487f6630.
//
// Solidity: function validContracts(address ) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactorySession) ValidContracts(arg0 common.Address) (bool, error) {
	return _HauskaContractFactory.Contract.ValidContracts(&_HauskaContractFactory.CallOpts, arg0)
}

// ValidContracts is a free data retrieval call binding the contract method 0x487f6630.
//
// Solidity: function validContracts(address ) view returns(bool)
func (_HauskaContractFactory *HauskaContractFactoryCallerSession) ValidContracts(arg0 common.Address) (bool, error) {
	return _HauskaContractFactory.Contract.ValidContracts(&_HauskaContractFactory.CallOpts, arg0)
}

// CreateContract is a paid mutator transaction binding the contract method 0x3c46c43c.
//
// Solidity: function createContract(address principalEntity, address integrationPartner) returns(address)
func (_HauskaContractFactory *HauskaContractFactoryTransactor) CreateContract(opts *bind.TransactOpts, principalEntity common.Address, integrationPartner common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "createContract", principalEntity, integrationPartner)
}

// CreateContract is a paid mutator transaction binding the contract method 0x3c46c43c.
//
// Solidity: function createContract(address principalEntity, address integrationPartner) returns(address)
func (_HauskaContractFactory *HauskaContractFactorySession) CreateContract(principalEntity common.Address, integrationPartner common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.CreateContract(&_HauskaContractFactory.TransactOpts, principalEntity, integrationPartner)
}

// CreateContract is a paid mutator transaction binding the contract method 0x3c46c43c.
//
// Solidity: function createContract(address principalEntity, address integrationPartner) returns(address)
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) CreateContract(principalEntity common.Address, integrationPartner common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.CreateContract(&_HauskaContractFactory.TransactOpts, principalEntity, integrationPartner)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.GrantRole(&_HauskaContractFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.GrantRole(&_HauskaContractFactory.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HauskaContractFactory *HauskaContractFactorySession) Pause() (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.Pause(&_HauskaContractFactory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) Pause() (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.Pause(&_HauskaContractFactory.TransactOpts)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xc375c2ef.
//
// Solidity: function removeContract(address contractToRemove) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactor) RemoveContract(opts *bind.TransactOpts, contractToRemove common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "removeContract", contractToRemove)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xc375c2ef.
//
// Solidity: function removeContract(address contractToRemove) returns()
func (_HauskaContractFactory *HauskaContractFactorySession) RemoveContract(contractToRemove common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.RemoveContract(&_HauskaContractFactory.TransactOpts, contractToRemove)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xc375c2ef.
//
// Solidity: function removeContract(address contractToRemove) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) RemoveContract(contractToRemove common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.RemoveContract(&_HauskaContractFactory.TransactOpts, contractToRemove)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactorySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.RenounceRole(&_HauskaContractFactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.RenounceRole(&_HauskaContractFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.RevokeRole(&_HauskaContractFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.RevokeRole(&_HauskaContractFactory.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HauskaContractFactory *HauskaContractFactorySession) Unpause() (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.Unpause(&_HauskaContractFactory.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.Unpause(&_HauskaContractFactory.TransactOpts)
}

// UpdatePlatformFees is a paid mutator transaction binding the contract method 0x49f01972.
//
// Solidity: function updatePlatformFees(uint32 _hauskaFeePct, uint32 _integratorFeePct) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactor) UpdatePlatformFees(opts *bind.TransactOpts, _hauskaFeePct uint32, _integratorFeePct uint32) (*types.Transaction, error) {
	return _HauskaContractFactory.contract.Transact(opts, "updatePlatformFees", _hauskaFeePct, _integratorFeePct)
}

// UpdatePlatformFees is a paid mutator transaction binding the contract method 0x49f01972.
//
// Solidity: function updatePlatformFees(uint32 _hauskaFeePct, uint32 _integratorFeePct) returns()
func (_HauskaContractFactory *HauskaContractFactorySession) UpdatePlatformFees(_hauskaFeePct uint32, _integratorFeePct uint32) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.UpdatePlatformFees(&_HauskaContractFactory.TransactOpts, _hauskaFeePct, _integratorFeePct)
}

// UpdatePlatformFees is a paid mutator transaction binding the contract method 0x49f01972.
//
// Solidity: function updatePlatformFees(uint32 _hauskaFeePct, uint32 _integratorFeePct) returns()
func (_HauskaContractFactory *HauskaContractFactoryTransactorSession) UpdatePlatformFees(_hauskaFeePct uint32, _integratorFeePct uint32) (*types.Transaction, error) {
	return _HauskaContractFactory.Contract.UpdatePlatformFees(&_HauskaContractFactory.TransactOpts, _hauskaFeePct, _integratorFeePct)
}

// HauskaContractFactoryContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the HauskaContractFactory contract.
type HauskaContractFactoryContractCreatedIterator struct {
	Event *HauskaContractFactoryContractCreated // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryContractCreated)
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
		it.Event = new(HauskaContractFactoryContractCreated)
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
func (it *HauskaContractFactoryContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryContractCreated represents a ContractCreated event raised by the HauskaContractFactory contract.
type HauskaContractFactoryContractCreated struct {
	Principal          common.Address
	ContractAddress    common.Address
	IntegrationPartner common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0x1202c61d7d89c76ba1493b085733ede04e071a2a76bb0fbae1345f128fe8b29d.
//
// Solidity: event ContractCreated(address indexed principal, address indexed contractAddress, address indexed integrationPartner)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterContractCreated(opts *bind.FilterOpts, principal []common.Address, contractAddress []common.Address, integrationPartner []common.Address) (*HauskaContractFactoryContractCreatedIterator, error) {

	var principalRule []interface{}
	for _, principalItem := range principal {
		principalRule = append(principalRule, principalItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var integrationPartnerRule []interface{}
	for _, integrationPartnerItem := range integrationPartner {
		integrationPartnerRule = append(integrationPartnerRule, integrationPartnerItem)
	}

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "ContractCreated", principalRule, contractAddressRule, integrationPartnerRule)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryContractCreatedIterator{contract: _HauskaContractFactory.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0x1202c61d7d89c76ba1493b085733ede04e071a2a76bb0fbae1345f128fe8b29d.
//
// Solidity: event ContractCreated(address indexed principal, address indexed contractAddress, address indexed integrationPartner)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryContractCreated, principal []common.Address, contractAddress []common.Address, integrationPartner []common.Address) (event.Subscription, error) {

	var principalRule []interface{}
	for _, principalItem := range principal {
		principalRule = append(principalRule, principalItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var integrationPartnerRule []interface{}
	for _, integrationPartnerItem := range integrationPartner {
		integrationPartnerRule = append(integrationPartnerRule, integrationPartnerItem)
	}

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "ContractCreated", principalRule, contractAddressRule, integrationPartnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryContractCreated)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ParseContractCreated is a log parse operation binding the contract event 0x1202c61d7d89c76ba1493b085733ede04e071a2a76bb0fbae1345f128fe8b29d.
//
// Solidity: event ContractCreated(address indexed principal, address indexed contractAddress, address indexed integrationPartner)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParseContractCreated(log types.Log) (*HauskaContractFactoryContractCreated, error) {
	event := new(HauskaContractFactoryContractCreated)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryContractRemovedIterator is returned from FilterContractRemoved and is used to iterate over the raw logs and unpacked data for ContractRemoved events raised by the HauskaContractFactory contract.
type HauskaContractFactoryContractRemovedIterator struct {
	Event *HauskaContractFactoryContractRemoved // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryContractRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryContractRemoved)
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
		it.Event = new(HauskaContractFactoryContractRemoved)
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
func (it *HauskaContractFactoryContractRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryContractRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryContractRemoved represents a ContractRemoved event raised by the HauskaContractFactory contract.
type HauskaContractFactoryContractRemoved struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractRemoved is a free log retrieval operation binding the contract event 0x8d30d41865a0b811b9545d879520d2dde9f4cc49e4241f486ad9752bc904b565.
//
// Solidity: event ContractRemoved(address indexed contractAddress)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterContractRemoved(opts *bind.FilterOpts, contractAddress []common.Address) (*HauskaContractFactoryContractRemovedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "ContractRemoved", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryContractRemovedIterator{contract: _HauskaContractFactory.contract, event: "ContractRemoved", logs: logs, sub: sub}, nil
}

// WatchContractRemoved is a free log subscription operation binding the contract event 0x8d30d41865a0b811b9545d879520d2dde9f4cc49e4241f486ad9752bc904b565.
//
// Solidity: event ContractRemoved(address indexed contractAddress)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchContractRemoved(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryContractRemoved, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "ContractRemoved", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryContractRemoved)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "ContractRemoved", log); err != nil {
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

// ParseContractRemoved is a log parse operation binding the contract event 0x8d30d41865a0b811b9545d879520d2dde9f4cc49e4241f486ad9752bc904b565.
//
// Solidity: event ContractRemoved(address indexed contractAddress)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParseContractRemoved(log types.Log) (*HauskaContractFactoryContractRemoved, error) {
	event := new(HauskaContractFactoryContractRemoved)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "ContractRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryModulesDeployedIterator is returned from FilterModulesDeployed and is used to iterate over the raw logs and unpacked data for ModulesDeployed events raised by the HauskaContractFactory contract.
type HauskaContractFactoryModulesDeployedIterator struct {
	Event *HauskaContractFactoryModulesDeployed // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryModulesDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryModulesDeployed)
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
		it.Event = new(HauskaContractFactoryModulesDeployed)
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
func (it *HauskaContractFactoryModulesDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryModulesDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryModulesDeployed represents a ModulesDeployed event raised by the HauskaContractFactory contract.
type HauskaContractFactoryModulesDeployed struct {
	LicenseManager     common.Address
	AssetRegistry      common.Address
	GroupManager       common.Address
	RevenueDistributor common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterModulesDeployed is a free log retrieval operation binding the contract event 0x987bcafbf3cc3a68a0af70bff9583cb5d7cfa501f4f2b2573a4759511397bd33.
//
// Solidity: event ModulesDeployed(address licenseManager, address assetRegistry, address groupManager, address revenueDistributor)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterModulesDeployed(opts *bind.FilterOpts) (*HauskaContractFactoryModulesDeployedIterator, error) {

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "ModulesDeployed")
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryModulesDeployedIterator{contract: _HauskaContractFactory.contract, event: "ModulesDeployed", logs: logs, sub: sub}, nil
}

// WatchModulesDeployed is a free log subscription operation binding the contract event 0x987bcafbf3cc3a68a0af70bff9583cb5d7cfa501f4f2b2573a4759511397bd33.
//
// Solidity: event ModulesDeployed(address licenseManager, address assetRegistry, address groupManager, address revenueDistributor)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchModulesDeployed(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryModulesDeployed) (event.Subscription, error) {

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "ModulesDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryModulesDeployed)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "ModulesDeployed", log); err != nil {
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

// ParseModulesDeployed is a log parse operation binding the contract event 0x987bcafbf3cc3a68a0af70bff9583cb5d7cfa501f4f2b2573a4759511397bd33.
//
// Solidity: event ModulesDeployed(address licenseManager, address assetRegistry, address groupManager, address revenueDistributor)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParseModulesDeployed(log types.Log) (*HauskaContractFactoryModulesDeployed, error) {
	event := new(HauskaContractFactoryModulesDeployed)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "ModulesDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HauskaContractFactory contract.
type HauskaContractFactoryPausedIterator struct {
	Event *HauskaContractFactoryPaused // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryPaused)
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
		it.Event = new(HauskaContractFactoryPaused)
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
func (it *HauskaContractFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryPaused represents a Paused event raised by the HauskaContractFactory contract.
type HauskaContractFactoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterPaused(opts *bind.FilterOpts) (*HauskaContractFactoryPausedIterator, error) {

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryPausedIterator{contract: _HauskaContractFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryPaused) (event.Subscription, error) {

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryPaused)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParsePaused(log types.Log) (*HauskaContractFactoryPaused, error) {
	event := new(HauskaContractFactoryPaused)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryPlatformFeesUpdatedIterator is returned from FilterPlatformFeesUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeesUpdated events raised by the HauskaContractFactory contract.
type HauskaContractFactoryPlatformFeesUpdatedIterator struct {
	Event *HauskaContractFactoryPlatformFeesUpdated // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryPlatformFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryPlatformFeesUpdated)
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
		it.Event = new(HauskaContractFactoryPlatformFeesUpdated)
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
func (it *HauskaContractFactoryPlatformFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryPlatformFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryPlatformFeesUpdated represents a PlatformFeesUpdated event raised by the HauskaContractFactory contract.
type HauskaContractFactoryPlatformFeesUpdated struct {
	HauskaFee     uint32
	IntegratorFee uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeesUpdated is a free log retrieval operation binding the contract event 0x6940c1f13c19b7b46ebd19df2c2a905b51e5f9d6a09f3800cc7a55649d29dde0.
//
// Solidity: event PlatformFeesUpdated(uint32 hauskaFee, uint32 integratorFee)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterPlatformFeesUpdated(opts *bind.FilterOpts) (*HauskaContractFactoryPlatformFeesUpdatedIterator, error) {

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "PlatformFeesUpdated")
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryPlatformFeesUpdatedIterator{contract: _HauskaContractFactory.contract, event: "PlatformFeesUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeesUpdated is a free log subscription operation binding the contract event 0x6940c1f13c19b7b46ebd19df2c2a905b51e5f9d6a09f3800cc7a55649d29dde0.
//
// Solidity: event PlatformFeesUpdated(uint32 hauskaFee, uint32 integratorFee)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchPlatformFeesUpdated(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryPlatformFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "PlatformFeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryPlatformFeesUpdated)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "PlatformFeesUpdated", log); err != nil {
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

// ParsePlatformFeesUpdated is a log parse operation binding the contract event 0x6940c1f13c19b7b46ebd19df2c2a905b51e5f9d6a09f3800cc7a55649d29dde0.
//
// Solidity: event PlatformFeesUpdated(uint32 hauskaFee, uint32 integratorFee)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParsePlatformFeesUpdated(log types.Log) (*HauskaContractFactoryPlatformFeesUpdated, error) {
	event := new(HauskaContractFactoryPlatformFeesUpdated)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "PlatformFeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HauskaContractFactory contract.
type HauskaContractFactoryRoleAdminChangedIterator struct {
	Event *HauskaContractFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryRoleAdminChanged)
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
		it.Event = new(HauskaContractFactoryRoleAdminChanged)
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
func (it *HauskaContractFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the HauskaContractFactory contract.
type HauskaContractFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HauskaContractFactoryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryRoleAdminChangedIterator{contract: _HauskaContractFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryRoleAdminChanged)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*HauskaContractFactoryRoleAdminChanged, error) {
	event := new(HauskaContractFactoryRoleAdminChanged)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HauskaContractFactory contract.
type HauskaContractFactoryRoleGrantedIterator struct {
	Event *HauskaContractFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryRoleGranted)
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
		it.Event = new(HauskaContractFactoryRoleGranted)
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
func (it *HauskaContractFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryRoleGranted represents a RoleGranted event raised by the HauskaContractFactory contract.
type HauskaContractFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaContractFactoryRoleGrantedIterator, error) {

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

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryRoleGrantedIterator{contract: _HauskaContractFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryRoleGranted)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParseRoleGranted(log types.Log) (*HauskaContractFactoryRoleGranted, error) {
	event := new(HauskaContractFactoryRoleGranted)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HauskaContractFactory contract.
type HauskaContractFactoryRoleRevokedIterator struct {
	Event *HauskaContractFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryRoleRevoked)
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
		it.Event = new(HauskaContractFactoryRoleRevoked)
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
func (it *HauskaContractFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryRoleRevoked represents a RoleRevoked event raised by the HauskaContractFactory contract.
type HauskaContractFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaContractFactoryRoleRevokedIterator, error) {

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

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryRoleRevokedIterator{contract: _HauskaContractFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryRoleRevoked)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParseRoleRevoked(log types.Log) (*HauskaContractFactoryRoleRevoked, error) {
	event := new(HauskaContractFactoryRoleRevoked)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaContractFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HauskaContractFactory contract.
type HauskaContractFactoryUnpausedIterator struct {
	Event *HauskaContractFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *HauskaContractFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaContractFactoryUnpaused)
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
		it.Event = new(HauskaContractFactoryUnpaused)
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
func (it *HauskaContractFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaContractFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaContractFactoryUnpaused represents a Unpaused event raised by the HauskaContractFactory contract.
type HauskaContractFactoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HauskaContractFactoryUnpausedIterator, error) {

	logs, sub, err := _HauskaContractFactory.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HauskaContractFactoryUnpausedIterator{contract: _HauskaContractFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HauskaContractFactoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _HauskaContractFactory.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaContractFactoryUnpaused)
				if err := _HauskaContractFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HauskaContractFactory *HauskaContractFactoryFilterer) ParseUnpaused(log types.Log) (*HauskaContractFactoryUnpaused, error) {
	event := new(HauskaContractFactoryUnpaused)
	if err := _HauskaContractFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
