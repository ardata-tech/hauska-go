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

// HauskaRevenueDistributorMetaData contains all meta data concerning the HauskaRevenueDistributor contract.
var HauskaRevenueDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"hauskaFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"integratorFee\",\"type\":\"uint32\"}],\"name\":\"CustomFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EarningsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hauskaAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"integratorAmount\",\"type\":\"uint256\"}],\"name\":\"RevenueDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORIZED_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTOR_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"addAuthorizedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creatorEarnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customHauskaFees\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customIntegratorFees\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"integrationPartner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"distributeRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getEarnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"getRevenueStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creatorTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hauskaTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"integratorTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasCustomFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hauskaEarnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"integratorEarnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"orgCreatorRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"orgHauskaRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"orgIntegratorRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"removeAuthorizedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"removeCustomFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"hauskaFeePct\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"integratorFeePct\",\"type\":\"uint32\"}],\"name\":\"setCustomFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalRevenueDistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HauskaRevenueDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use HauskaRevenueDistributorMetaData.ABI instead.
var HauskaRevenueDistributorABI = HauskaRevenueDistributorMetaData.ABI

// HauskaRevenueDistributor is an auto generated Go binding around an Ethereum contract.
type HauskaRevenueDistributor struct {
	HauskaRevenueDistributorCaller     // Read-only binding to the contract
	HauskaRevenueDistributorTransactor // Write-only binding to the contract
	HauskaRevenueDistributorFilterer   // Log filterer for contract events
}

// HauskaRevenueDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type HauskaRevenueDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaRevenueDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HauskaRevenueDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaRevenueDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HauskaRevenueDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaRevenueDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HauskaRevenueDistributorSession struct {
	Contract     *HauskaRevenueDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HauskaRevenueDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HauskaRevenueDistributorCallerSession struct {
	Contract *HauskaRevenueDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// HauskaRevenueDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HauskaRevenueDistributorTransactorSession struct {
	Contract     *HauskaRevenueDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// HauskaRevenueDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type HauskaRevenueDistributorRaw struct {
	Contract *HauskaRevenueDistributor // Generic contract binding to access the raw methods on
}

// HauskaRevenueDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HauskaRevenueDistributorCallerRaw struct {
	Contract *HauskaRevenueDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// HauskaRevenueDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HauskaRevenueDistributorTransactorRaw struct {
	Contract *HauskaRevenueDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHauskaRevenueDistributor creates a new instance of HauskaRevenueDistributor, bound to a specific deployed contract.
func NewHauskaRevenueDistributor(address common.Address, backend bind.ContractBackend) (*HauskaRevenueDistributor, error) {
	contract, err := bindHauskaRevenueDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributor{HauskaRevenueDistributorCaller: HauskaRevenueDistributorCaller{contract: contract}, HauskaRevenueDistributorTransactor: HauskaRevenueDistributorTransactor{contract: contract}, HauskaRevenueDistributorFilterer: HauskaRevenueDistributorFilterer{contract: contract}}, nil
}

// NewHauskaRevenueDistributorCaller creates a new read-only instance of HauskaRevenueDistributor, bound to a specific deployed contract.
func NewHauskaRevenueDistributorCaller(address common.Address, caller bind.ContractCaller) (*HauskaRevenueDistributorCaller, error) {
	contract, err := bindHauskaRevenueDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorCaller{contract: contract}, nil
}

// NewHauskaRevenueDistributorTransactor creates a new write-only instance of HauskaRevenueDistributor, bound to a specific deployed contract.
func NewHauskaRevenueDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*HauskaRevenueDistributorTransactor, error) {
	contract, err := bindHauskaRevenueDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorTransactor{contract: contract}, nil
}

// NewHauskaRevenueDistributorFilterer creates a new log filterer instance of HauskaRevenueDistributor, bound to a specific deployed contract.
func NewHauskaRevenueDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*HauskaRevenueDistributorFilterer, error) {
	contract, err := bindHauskaRevenueDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorFilterer{contract: contract}, nil
}

// bindHauskaRevenueDistributor binds a generic wrapper to an already deployed contract.
func bindHauskaRevenueDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HauskaRevenueDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaRevenueDistributor *HauskaRevenueDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaRevenueDistributor.Contract.HauskaRevenueDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaRevenueDistributor *HauskaRevenueDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.HauskaRevenueDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaRevenueDistributor *HauskaRevenueDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.HauskaRevenueDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaRevenueDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.contract.Transact(opts, method, params...)
}

// AUTHORIZEDCONTRACTROLE is a free data retrieval call binding the contract method 0x3182cd8c.
//
// Solidity: function AUTHORIZED_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) AUTHORIZEDCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "AUTHORIZED_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHORIZEDCONTRACTROLE is a free data retrieval call binding the contract method 0x3182cd8c.
//
// Solidity: function AUTHORIZED_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) AUTHORIZEDCONTRACTROLE() ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.AUTHORIZEDCONTRACTROLE(&_HauskaRevenueDistributor.CallOpts)
}

// AUTHORIZEDCONTRACTROLE is a free data retrieval call binding the contract method 0x3182cd8c.
//
// Solidity: function AUTHORIZED_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) AUTHORIZEDCONTRACTROLE() ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.AUTHORIZEDCONTRACTROLE(&_HauskaRevenueDistributor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.DEFAULTADMINROLE(&_HauskaRevenueDistributor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.DEFAULTADMINROLE(&_HauskaRevenueDistributor.CallOpts)
}

// DISTRIBUTORADMINROLE is a free data retrieval call binding the contract method 0x0e6fcd3a.
//
// Solidity: function DISTRIBUTOR_ADMIN_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) DISTRIBUTORADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "DISTRIBUTOR_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISTRIBUTORADMINROLE is a free data retrieval call binding the contract method 0x0e6fcd3a.
//
// Solidity: function DISTRIBUTOR_ADMIN_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) DISTRIBUTORADMINROLE() ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.DISTRIBUTORADMINROLE(&_HauskaRevenueDistributor.CallOpts)
}

// DISTRIBUTORADMINROLE is a free data retrieval call binding the contract method 0x0e6fcd3a.
//
// Solidity: function DISTRIBUTOR_ADMIN_ROLE() view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) DISTRIBUTORADMINROLE() ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.DISTRIBUTORADMINROLE(&_HauskaRevenueDistributor.CallOpts)
}

// CreatorEarnings is a free data retrieval call binding the contract method 0xfa3c30a8.
//
// Solidity: function creatorEarnings(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) CreatorEarnings(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "creatorEarnings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorEarnings is a free data retrieval call binding the contract method 0xfa3c30a8.
//
// Solidity: function creatorEarnings(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) CreatorEarnings(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.CreatorEarnings(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// CreatorEarnings is a free data retrieval call binding the contract method 0xfa3c30a8.
//
// Solidity: function creatorEarnings(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) CreatorEarnings(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.CreatorEarnings(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// CustomHauskaFees is a free data retrieval call binding the contract method 0xbafe3b23.
//
// Solidity: function customHauskaFees(address ) view returns(uint32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) CustomHauskaFees(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "customHauskaFees", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CustomHauskaFees is a free data retrieval call binding the contract method 0xbafe3b23.
//
// Solidity: function customHauskaFees(address ) view returns(uint32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) CustomHauskaFees(arg0 common.Address) (uint32, error) {
	return _HauskaRevenueDistributor.Contract.CustomHauskaFees(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// CustomHauskaFees is a free data retrieval call binding the contract method 0xbafe3b23.
//
// Solidity: function customHauskaFees(address ) view returns(uint32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) CustomHauskaFees(arg0 common.Address) (uint32, error) {
	return _HauskaRevenueDistributor.Contract.CustomHauskaFees(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// CustomIntegratorFees is a free data retrieval call binding the contract method 0xae3a262e.
//
// Solidity: function customIntegratorFees(address ) view returns(uint32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) CustomIntegratorFees(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "customIntegratorFees", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CustomIntegratorFees is a free data retrieval call binding the contract method 0xae3a262e.
//
// Solidity: function customIntegratorFees(address ) view returns(uint32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) CustomIntegratorFees(arg0 common.Address) (uint32, error) {
	return _HauskaRevenueDistributor.Contract.CustomIntegratorFees(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// CustomIntegratorFees is a free data retrieval call binding the contract method 0xae3a262e.
//
// Solidity: function customIntegratorFees(address ) view returns(uint32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) CustomIntegratorFees(arg0 common.Address) (uint32, error) {
	return _HauskaRevenueDistributor.Contract.CustomIntegratorFees(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) Factory() (common.Address, error) {
	return _HauskaRevenueDistributor.Contract.Factory(&_HauskaRevenueDistributor.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) Factory() (common.Address, error) {
	return _HauskaRevenueDistributor.Contract.Factory(&_HauskaRevenueDistributor.CallOpts)
}

// GetEarnings is a free data retrieval call binding the contract method 0x131b9c04.
//
// Solidity: function getEarnings(address account) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) GetEarnings(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "getEarnings", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEarnings is a free data retrieval call binding the contract method 0x131b9c04.
//
// Solidity: function getEarnings(address account) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) GetEarnings(account common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.GetEarnings(&_HauskaRevenueDistributor.CallOpts, account)
}

// GetEarnings is a free data retrieval call binding the contract method 0x131b9c04.
//
// Solidity: function getEarnings(address account) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) GetEarnings(account common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.GetEarnings(&_HauskaRevenueDistributor.CallOpts, account)
}

// GetRevenueStats is a free data retrieval call binding the contract method 0x08f56c08.
//
// Solidity: function getRevenueStats(address orgContract) view returns(uint256 total, uint256 creatorTotal, uint256 hauskaTotal, uint256 integratorTotal)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) GetRevenueStats(opts *bind.CallOpts, orgContract common.Address) (struct {
	Total           *big.Int
	CreatorTotal    *big.Int
	HauskaTotal     *big.Int
	IntegratorTotal *big.Int
}, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "getRevenueStats", orgContract)

	outstruct := new(struct {
		Total           *big.Int
		CreatorTotal    *big.Int
		HauskaTotal     *big.Int
		IntegratorTotal *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatorTotal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HauskaTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IntegratorTotal = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRevenueStats is a free data retrieval call binding the contract method 0x08f56c08.
//
// Solidity: function getRevenueStats(address orgContract) view returns(uint256 total, uint256 creatorTotal, uint256 hauskaTotal, uint256 integratorTotal)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) GetRevenueStats(orgContract common.Address) (struct {
	Total           *big.Int
	CreatorTotal    *big.Int
	HauskaTotal     *big.Int
	IntegratorTotal *big.Int
}, error) {
	return _HauskaRevenueDistributor.Contract.GetRevenueStats(&_HauskaRevenueDistributor.CallOpts, orgContract)
}

// GetRevenueStats is a free data retrieval call binding the contract method 0x08f56c08.
//
// Solidity: function getRevenueStats(address orgContract) view returns(uint256 total, uint256 creatorTotal, uint256 hauskaTotal, uint256 integratorTotal)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) GetRevenueStats(orgContract common.Address) (struct {
	Total           *big.Int
	CreatorTotal    *big.Int
	HauskaTotal     *big.Int
	IntegratorTotal *big.Int
}, error) {
	return _HauskaRevenueDistributor.Contract.GetRevenueStats(&_HauskaRevenueDistributor.CallOpts, orgContract)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.GetRoleAdmin(&_HauskaRevenueDistributor.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaRevenueDistributor.Contract.GetRoleAdmin(&_HauskaRevenueDistributor.CallOpts, role)
}

// HasCustomFees is a free data retrieval call binding the contract method 0xfbb49a50.
//
// Solidity: function hasCustomFees(address ) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) HasCustomFees(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "hasCustomFees", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCustomFees is a free data retrieval call binding the contract method 0xfbb49a50.
//
// Solidity: function hasCustomFees(address ) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) HasCustomFees(arg0 common.Address) (bool, error) {
	return _HauskaRevenueDistributor.Contract.HasCustomFees(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// HasCustomFees is a free data retrieval call binding the contract method 0xfbb49a50.
//
// Solidity: function hasCustomFees(address ) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) HasCustomFees(arg0 common.Address) (bool, error) {
	return _HauskaRevenueDistributor.Contract.HasCustomFees(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaRevenueDistributor.Contract.HasRole(&_HauskaRevenueDistributor.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaRevenueDistributor.Contract.HasRole(&_HauskaRevenueDistributor.CallOpts, role, account)
}

// HauskaEarnings is a free data retrieval call binding the contract method 0xa1f53608.
//
// Solidity: function hauskaEarnings() view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) HauskaEarnings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "hauskaEarnings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HauskaEarnings is a free data retrieval call binding the contract method 0xa1f53608.
//
// Solidity: function hauskaEarnings() view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) HauskaEarnings() (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.HauskaEarnings(&_HauskaRevenueDistributor.CallOpts)
}

// HauskaEarnings is a free data retrieval call binding the contract method 0xa1f53608.
//
// Solidity: function hauskaEarnings() view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) HauskaEarnings() (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.HauskaEarnings(&_HauskaRevenueDistributor.CallOpts)
}

// IntegratorEarnings is a free data retrieval call binding the contract method 0x76057aef.
//
// Solidity: function integratorEarnings(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) IntegratorEarnings(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "integratorEarnings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegratorEarnings is a free data retrieval call binding the contract method 0x76057aef.
//
// Solidity: function integratorEarnings(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) IntegratorEarnings(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.IntegratorEarnings(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// IntegratorEarnings is a free data retrieval call binding the contract method 0x76057aef.
//
// Solidity: function integratorEarnings(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) IntegratorEarnings(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.IntegratorEarnings(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// OrgCreatorRevenue is a free data retrieval call binding the contract method 0x9ebf5eb4.
//
// Solidity: function orgCreatorRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) OrgCreatorRevenue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "orgCreatorRevenue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrgCreatorRevenue is a free data retrieval call binding the contract method 0x9ebf5eb4.
//
// Solidity: function orgCreatorRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) OrgCreatorRevenue(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.OrgCreatorRevenue(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// OrgCreatorRevenue is a free data retrieval call binding the contract method 0x9ebf5eb4.
//
// Solidity: function orgCreatorRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) OrgCreatorRevenue(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.OrgCreatorRevenue(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// OrgHauskaRevenue is a free data retrieval call binding the contract method 0xdf199cfb.
//
// Solidity: function orgHauskaRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) OrgHauskaRevenue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "orgHauskaRevenue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrgHauskaRevenue is a free data retrieval call binding the contract method 0xdf199cfb.
//
// Solidity: function orgHauskaRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) OrgHauskaRevenue(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.OrgHauskaRevenue(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// OrgHauskaRevenue is a free data retrieval call binding the contract method 0xdf199cfb.
//
// Solidity: function orgHauskaRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) OrgHauskaRevenue(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.OrgHauskaRevenue(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// OrgIntegratorRevenue is a free data retrieval call binding the contract method 0xfa18679f.
//
// Solidity: function orgIntegratorRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) OrgIntegratorRevenue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "orgIntegratorRevenue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrgIntegratorRevenue is a free data retrieval call binding the contract method 0xfa18679f.
//
// Solidity: function orgIntegratorRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) OrgIntegratorRevenue(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.OrgIntegratorRevenue(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// OrgIntegratorRevenue is a free data retrieval call binding the contract method 0xfa18679f.
//
// Solidity: function orgIntegratorRevenue(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) OrgIntegratorRevenue(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.OrgIntegratorRevenue(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaRevenueDistributor.Contract.SupportsInterface(&_HauskaRevenueDistributor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaRevenueDistributor.Contract.SupportsInterface(&_HauskaRevenueDistributor.CallOpts, interfaceId)
}

// TotalRevenueDistributed is a free data retrieval call binding the contract method 0xbc58bbc5.
//
// Solidity: function totalRevenueDistributed(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) TotalRevenueDistributed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "totalRevenueDistributed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRevenueDistributed is a free data retrieval call binding the contract method 0xbc58bbc5.
//
// Solidity: function totalRevenueDistributed(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) TotalRevenueDistributed(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.TotalRevenueDistributed(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// TotalRevenueDistributed is a free data retrieval call binding the contract method 0xbc58bbc5.
//
// Solidity: function totalRevenueDistributed(address ) view returns(uint256)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) TotalRevenueDistributed(arg0 common.Address) (*big.Int, error) {
	return _HauskaRevenueDistributor.Contract.TotalRevenueDistributed(&_HauskaRevenueDistributor.CallOpts, arg0)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaRevenueDistributor.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) UsdcToken() (common.Address, error) {
	return _HauskaRevenueDistributor.Contract.UsdcToken(&_HauskaRevenueDistributor.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorCallerSession) UsdcToken() (common.Address, error) {
	return _HauskaRevenueDistributor.Contract.UsdcToken(&_HauskaRevenueDistributor.CallOpts)
}

// AddAuthorizedContract is a paid mutator transaction binding the contract method 0x98eaa4a7.
//
// Solidity: function addAuthorizedContract(address contractAddress) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) AddAuthorizedContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "addAuthorizedContract", contractAddress)
}

// AddAuthorizedContract is a paid mutator transaction binding the contract method 0x98eaa4a7.
//
// Solidity: function addAuthorizedContract(address contractAddress) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) AddAuthorizedContract(contractAddress common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.AddAuthorizedContract(&_HauskaRevenueDistributor.TransactOpts, contractAddress)
}

// AddAuthorizedContract is a paid mutator transaction binding the contract method 0x98eaa4a7.
//
// Solidity: function addAuthorizedContract(address contractAddress) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) AddAuthorizedContract(contractAddress common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.AddAuthorizedContract(&_HauskaRevenueDistributor.TransactOpts, contractAddress)
}

// DistributeRevenue is a paid mutator transaction binding the contract method 0x998af887.
//
// Solidity: function distributeRevenue(address from, uint256 amount, address assetOwner, address integrationPartner, address orgContract) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) DistributeRevenue(opts *bind.TransactOpts, from common.Address, amount *big.Int, assetOwner common.Address, integrationPartner common.Address, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "distributeRevenue", from, amount, assetOwner, integrationPartner, orgContract)
}

// DistributeRevenue is a paid mutator transaction binding the contract method 0x998af887.
//
// Solidity: function distributeRevenue(address from, uint256 amount, address assetOwner, address integrationPartner, address orgContract) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) DistributeRevenue(from common.Address, amount *big.Int, assetOwner common.Address, integrationPartner common.Address, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.DistributeRevenue(&_HauskaRevenueDistributor.TransactOpts, from, amount, assetOwner, integrationPartner, orgContract)
}

// DistributeRevenue is a paid mutator transaction binding the contract method 0x998af887.
//
// Solidity: function distributeRevenue(address from, uint256 amount, address assetOwner, address integrationPartner, address orgContract) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) DistributeRevenue(from common.Address, amount *big.Int, assetOwner common.Address, integrationPartner common.Address, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.DistributeRevenue(&_HauskaRevenueDistributor.TransactOpts, from, amount, assetOwner, integrationPartner, orgContract)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.GrantRole(&_HauskaRevenueDistributor.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.GrantRole(&_HauskaRevenueDistributor.TransactOpts, role, account)
}

// RemoveAuthorizedContract is a paid mutator transaction binding the contract method 0xe6b165ed.
//
// Solidity: function removeAuthorizedContract(address contractAddress) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) RemoveAuthorizedContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "removeAuthorizedContract", contractAddress)
}

// RemoveAuthorizedContract is a paid mutator transaction binding the contract method 0xe6b165ed.
//
// Solidity: function removeAuthorizedContract(address contractAddress) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) RemoveAuthorizedContract(contractAddress common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RemoveAuthorizedContract(&_HauskaRevenueDistributor.TransactOpts, contractAddress)
}

// RemoveAuthorizedContract is a paid mutator transaction binding the contract method 0xe6b165ed.
//
// Solidity: function removeAuthorizedContract(address contractAddress) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) RemoveAuthorizedContract(contractAddress common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RemoveAuthorizedContract(&_HauskaRevenueDistributor.TransactOpts, contractAddress)
}

// RemoveCustomFees is a paid mutator transaction binding the contract method 0x407aed1e.
//
// Solidity: function removeCustomFees(address orgContract) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) RemoveCustomFees(opts *bind.TransactOpts, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "removeCustomFees", orgContract)
}

// RemoveCustomFees is a paid mutator transaction binding the contract method 0x407aed1e.
//
// Solidity: function removeCustomFees(address orgContract) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) RemoveCustomFees(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RemoveCustomFees(&_HauskaRevenueDistributor.TransactOpts, orgContract)
}

// RemoveCustomFees is a paid mutator transaction binding the contract method 0x407aed1e.
//
// Solidity: function removeCustomFees(address orgContract) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) RemoveCustomFees(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RemoveCustomFees(&_HauskaRevenueDistributor.TransactOpts, orgContract)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RenounceRole(&_HauskaRevenueDistributor.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RenounceRole(&_HauskaRevenueDistributor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RevokeRole(&_HauskaRevenueDistributor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.RevokeRole(&_HauskaRevenueDistributor.TransactOpts, role, account)
}

// SetCustomFees is a paid mutator transaction binding the contract method 0x26c9b12d.
//
// Solidity: function setCustomFees(address orgContract, uint32 hauskaFeePct, uint32 integratorFeePct) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactor) SetCustomFees(opts *bind.TransactOpts, orgContract common.Address, hauskaFeePct uint32, integratorFeePct uint32) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.contract.Transact(opts, "setCustomFees", orgContract, hauskaFeePct, integratorFeePct)
}

// SetCustomFees is a paid mutator transaction binding the contract method 0x26c9b12d.
//
// Solidity: function setCustomFees(address orgContract, uint32 hauskaFeePct, uint32 integratorFeePct) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorSession) SetCustomFees(orgContract common.Address, hauskaFeePct uint32, integratorFeePct uint32) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.SetCustomFees(&_HauskaRevenueDistributor.TransactOpts, orgContract, hauskaFeePct, integratorFeePct)
}

// SetCustomFees is a paid mutator transaction binding the contract method 0x26c9b12d.
//
// Solidity: function setCustomFees(address orgContract, uint32 hauskaFeePct, uint32 integratorFeePct) returns()
func (_HauskaRevenueDistributor *HauskaRevenueDistributorTransactorSession) SetCustomFees(orgContract common.Address, hauskaFeePct uint32, integratorFeePct uint32) (*types.Transaction, error) {
	return _HauskaRevenueDistributor.Contract.SetCustomFees(&_HauskaRevenueDistributor.TransactOpts, orgContract, hauskaFeePct, integratorFeePct)
}

// HauskaRevenueDistributorCustomFeesSetIterator is returned from FilterCustomFeesSet and is used to iterate over the raw logs and unpacked data for CustomFeesSet events raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorCustomFeesSetIterator struct {
	Event *HauskaRevenueDistributorCustomFeesSet // Event containing the contract specifics and raw log

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
func (it *HauskaRevenueDistributorCustomFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaRevenueDistributorCustomFeesSet)
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
		it.Event = new(HauskaRevenueDistributorCustomFeesSet)
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
func (it *HauskaRevenueDistributorCustomFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaRevenueDistributorCustomFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaRevenueDistributorCustomFeesSet represents a CustomFeesSet event raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorCustomFeesSet struct {
	OrgContract   common.Address
	HauskaFee     uint32
	IntegratorFee uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCustomFeesSet is a free log retrieval operation binding the contract event 0xee0f051af6f1aee452b05393d97a084ad7ba2e242842eccb13f9c284c60a2c5f.
//
// Solidity: event CustomFeesSet(address indexed orgContract, uint32 hauskaFee, uint32 integratorFee)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) FilterCustomFeesSet(opts *bind.FilterOpts, orgContract []common.Address) (*HauskaRevenueDistributorCustomFeesSetIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}

	logs, sub, err := _HauskaRevenueDistributor.contract.FilterLogs(opts, "CustomFeesSet", orgContractRule)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorCustomFeesSetIterator{contract: _HauskaRevenueDistributor.contract, event: "CustomFeesSet", logs: logs, sub: sub}, nil
}

// WatchCustomFeesSet is a free log subscription operation binding the contract event 0xee0f051af6f1aee452b05393d97a084ad7ba2e242842eccb13f9c284c60a2c5f.
//
// Solidity: event CustomFeesSet(address indexed orgContract, uint32 hauskaFee, uint32 integratorFee)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) WatchCustomFeesSet(opts *bind.WatchOpts, sink chan<- *HauskaRevenueDistributorCustomFeesSet, orgContract []common.Address) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}

	logs, sub, err := _HauskaRevenueDistributor.contract.WatchLogs(opts, "CustomFeesSet", orgContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaRevenueDistributorCustomFeesSet)
				if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "CustomFeesSet", log); err != nil {
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

// ParseCustomFeesSet is a log parse operation binding the contract event 0xee0f051af6f1aee452b05393d97a084ad7ba2e242842eccb13f9c284c60a2c5f.
//
// Solidity: event CustomFeesSet(address indexed orgContract, uint32 hauskaFee, uint32 integratorFee)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) ParseCustomFeesSet(log types.Log) (*HauskaRevenueDistributorCustomFeesSet, error) {
	event := new(HauskaRevenueDistributorCustomFeesSet)
	if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "CustomFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaRevenueDistributorEarningsWithdrawnIterator is returned from FilterEarningsWithdrawn and is used to iterate over the raw logs and unpacked data for EarningsWithdrawn events raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorEarningsWithdrawnIterator struct {
	Event *HauskaRevenueDistributorEarningsWithdrawn // Event containing the contract specifics and raw log

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
func (it *HauskaRevenueDistributorEarningsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaRevenueDistributorEarningsWithdrawn)
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
		it.Event = new(HauskaRevenueDistributorEarningsWithdrawn)
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
func (it *HauskaRevenueDistributorEarningsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaRevenueDistributorEarningsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaRevenueDistributorEarningsWithdrawn represents a EarningsWithdrawn event raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorEarningsWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEarningsWithdrawn is a free log retrieval operation binding the contract event 0x48dc35af7b45e2a81fffad55f6e2fafacdb1d3d0d50d24ebdc16324f5ba757f1.
//
// Solidity: event EarningsWithdrawn(address indexed recipient, uint256 amount)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) FilterEarningsWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*HauskaRevenueDistributorEarningsWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _HauskaRevenueDistributor.contract.FilterLogs(opts, "EarningsWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorEarningsWithdrawnIterator{contract: _HauskaRevenueDistributor.contract, event: "EarningsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEarningsWithdrawn is a free log subscription operation binding the contract event 0x48dc35af7b45e2a81fffad55f6e2fafacdb1d3d0d50d24ebdc16324f5ba757f1.
//
// Solidity: event EarningsWithdrawn(address indexed recipient, uint256 amount)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) WatchEarningsWithdrawn(opts *bind.WatchOpts, sink chan<- *HauskaRevenueDistributorEarningsWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _HauskaRevenueDistributor.contract.WatchLogs(opts, "EarningsWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaRevenueDistributorEarningsWithdrawn)
				if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "EarningsWithdrawn", log); err != nil {
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

// ParseEarningsWithdrawn is a log parse operation binding the contract event 0x48dc35af7b45e2a81fffad55f6e2fafacdb1d3d0d50d24ebdc16324f5ba757f1.
//
// Solidity: event EarningsWithdrawn(address indexed recipient, uint256 amount)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) ParseEarningsWithdrawn(log types.Log) (*HauskaRevenueDistributorEarningsWithdrawn, error) {
	event := new(HauskaRevenueDistributorEarningsWithdrawn)
	if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "EarningsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaRevenueDistributorRevenueDistributedIterator is returned from FilterRevenueDistributed and is used to iterate over the raw logs and unpacked data for RevenueDistributed events raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRevenueDistributedIterator struct {
	Event *HauskaRevenueDistributorRevenueDistributed // Event containing the contract specifics and raw log

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
func (it *HauskaRevenueDistributorRevenueDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaRevenueDistributorRevenueDistributed)
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
		it.Event = new(HauskaRevenueDistributorRevenueDistributed)
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
func (it *HauskaRevenueDistributorRevenueDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaRevenueDistributorRevenueDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaRevenueDistributorRevenueDistributed represents a RevenueDistributed event raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRevenueDistributed struct {
	From             common.Address
	Amount           *big.Int
	AssetOwner       common.Address
	OwnerAmount      *big.Int
	HauskaAmount     *big.Int
	IntegratorAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRevenueDistributed is a free log retrieval operation binding the contract event 0xd364c3f5b6023a844ed5a2f07b23845760b4b41abc1accc6f782d9a0300b8f5b.
//
// Solidity: event RevenueDistributed(address indexed from, uint256 amount, address indexed assetOwner, uint256 ownerAmount, uint256 hauskaAmount, uint256 integratorAmount)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) FilterRevenueDistributed(opts *bind.FilterOpts, from []common.Address, assetOwner []common.Address) (*HauskaRevenueDistributorRevenueDistributedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	var assetOwnerRule []interface{}
	for _, assetOwnerItem := range assetOwner {
		assetOwnerRule = append(assetOwnerRule, assetOwnerItem)
	}

	logs, sub, err := _HauskaRevenueDistributor.contract.FilterLogs(opts, "RevenueDistributed", fromRule, assetOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorRevenueDistributedIterator{contract: _HauskaRevenueDistributor.contract, event: "RevenueDistributed", logs: logs, sub: sub}, nil
}

// WatchRevenueDistributed is a free log subscription operation binding the contract event 0xd364c3f5b6023a844ed5a2f07b23845760b4b41abc1accc6f782d9a0300b8f5b.
//
// Solidity: event RevenueDistributed(address indexed from, uint256 amount, address indexed assetOwner, uint256 ownerAmount, uint256 hauskaAmount, uint256 integratorAmount)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) WatchRevenueDistributed(opts *bind.WatchOpts, sink chan<- *HauskaRevenueDistributorRevenueDistributed, from []common.Address, assetOwner []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	var assetOwnerRule []interface{}
	for _, assetOwnerItem := range assetOwner {
		assetOwnerRule = append(assetOwnerRule, assetOwnerItem)
	}

	logs, sub, err := _HauskaRevenueDistributor.contract.WatchLogs(opts, "RevenueDistributed", fromRule, assetOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaRevenueDistributorRevenueDistributed)
				if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RevenueDistributed", log); err != nil {
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

// ParseRevenueDistributed is a log parse operation binding the contract event 0xd364c3f5b6023a844ed5a2f07b23845760b4b41abc1accc6f782d9a0300b8f5b.
//
// Solidity: event RevenueDistributed(address indexed from, uint256 amount, address indexed assetOwner, uint256 ownerAmount, uint256 hauskaAmount, uint256 integratorAmount)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) ParseRevenueDistributed(log types.Log) (*HauskaRevenueDistributorRevenueDistributed, error) {
	event := new(HauskaRevenueDistributorRevenueDistributed)
	if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RevenueDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaRevenueDistributorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRoleAdminChangedIterator struct {
	Event *HauskaRevenueDistributorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HauskaRevenueDistributorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaRevenueDistributorRoleAdminChanged)
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
		it.Event = new(HauskaRevenueDistributorRoleAdminChanged)
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
func (it *HauskaRevenueDistributorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaRevenueDistributorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaRevenueDistributorRoleAdminChanged represents a RoleAdminChanged event raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HauskaRevenueDistributorRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HauskaRevenueDistributor.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorRoleAdminChangedIterator{contract: _HauskaRevenueDistributor.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HauskaRevenueDistributorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HauskaRevenueDistributor.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaRevenueDistributorRoleAdminChanged)
				if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) ParseRoleAdminChanged(log types.Log) (*HauskaRevenueDistributorRoleAdminChanged, error) {
	event := new(HauskaRevenueDistributorRoleAdminChanged)
	if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaRevenueDistributorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRoleGrantedIterator struct {
	Event *HauskaRevenueDistributorRoleGranted // Event containing the contract specifics and raw log

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
func (it *HauskaRevenueDistributorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaRevenueDistributorRoleGranted)
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
		it.Event = new(HauskaRevenueDistributorRoleGranted)
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
func (it *HauskaRevenueDistributorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaRevenueDistributorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaRevenueDistributorRoleGranted represents a RoleGranted event raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaRevenueDistributorRoleGrantedIterator, error) {

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

	logs, sub, err := _HauskaRevenueDistributor.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorRoleGrantedIterator{contract: _HauskaRevenueDistributor.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HauskaRevenueDistributorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaRevenueDistributor.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaRevenueDistributorRoleGranted)
				if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) ParseRoleGranted(log types.Log) (*HauskaRevenueDistributorRoleGranted, error) {
	event := new(HauskaRevenueDistributorRoleGranted)
	if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaRevenueDistributorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRoleRevokedIterator struct {
	Event *HauskaRevenueDistributorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HauskaRevenueDistributorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaRevenueDistributorRoleRevoked)
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
		it.Event = new(HauskaRevenueDistributorRoleRevoked)
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
func (it *HauskaRevenueDistributorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaRevenueDistributorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaRevenueDistributorRoleRevoked represents a RoleRevoked event raised by the HauskaRevenueDistributor contract.
type HauskaRevenueDistributorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaRevenueDistributorRoleRevokedIterator, error) {

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

	logs, sub, err := _HauskaRevenueDistributor.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaRevenueDistributorRoleRevokedIterator{contract: _HauskaRevenueDistributor.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HauskaRevenueDistributorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaRevenueDistributor.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaRevenueDistributorRoleRevoked)
				if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HauskaRevenueDistributor *HauskaRevenueDistributorFilterer) ParseRoleRevoked(log types.Log) (*HauskaRevenueDistributorRoleRevoked, error) {
	event := new(HauskaRevenueDistributorRoleRevoked)
	if err := _HauskaRevenueDistributor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
