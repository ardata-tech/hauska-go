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

// IHauskaStructsAssetGroup is an auto generated low-level Go binding around an user-defined struct.
type IHauskaStructsAssetGroup struct {
	GroupId    *big.Int
	Members    []*big.Int
	Name       string
	GroupPrice *big.Int
	Owner      common.Address
}

// HauskaGroupManagerMetaData contains all meta data concerning the HauskaGroupManager contract.
var HauskaGroupManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"AssetAddedToGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"AssetRemovedFromGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GroupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"GroupUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORG_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"addAssetToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"addOrgContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"assetIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"groupPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creatorGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"getGroup\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"members\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"groupPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIHauskaStructs.AssetGroup\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"getGroupCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getGroupsByCreator\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"orgGroupCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orgGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"groupPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"removeAssetFromGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"removeGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"name\":\"removeOrgContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"updateGroupPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HauskaGroupManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use HauskaGroupManagerMetaData.ABI instead.
var HauskaGroupManagerABI = HauskaGroupManagerMetaData.ABI

// HauskaGroupManager is an auto generated Go binding around an Ethereum contract.
type HauskaGroupManager struct {
	HauskaGroupManagerCaller     // Read-only binding to the contract
	HauskaGroupManagerTransactor // Write-only binding to the contract
	HauskaGroupManagerFilterer   // Log filterer for contract events
}

// HauskaGroupManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type HauskaGroupManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaGroupManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HauskaGroupManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaGroupManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HauskaGroupManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaGroupManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HauskaGroupManagerSession struct {
	Contract     *HauskaGroupManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HauskaGroupManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HauskaGroupManagerCallerSession struct {
	Contract *HauskaGroupManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HauskaGroupManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HauskaGroupManagerTransactorSession struct {
	Contract     *HauskaGroupManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HauskaGroupManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type HauskaGroupManagerRaw struct {
	Contract *HauskaGroupManager // Generic contract binding to access the raw methods on
}

// HauskaGroupManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HauskaGroupManagerCallerRaw struct {
	Contract *HauskaGroupManagerCaller // Generic read-only contract binding to access the raw methods on
}

// HauskaGroupManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HauskaGroupManagerTransactorRaw struct {
	Contract *HauskaGroupManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHauskaGroupManager creates a new instance of HauskaGroupManager, bound to a specific deployed contract.
func NewHauskaGroupManager(address common.Address, backend bind.ContractBackend) (*HauskaGroupManager, error) {
	contract, err := bindHauskaGroupManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManager{HauskaGroupManagerCaller: HauskaGroupManagerCaller{contract: contract}, HauskaGroupManagerTransactor: HauskaGroupManagerTransactor{contract: contract}, HauskaGroupManagerFilterer: HauskaGroupManagerFilterer{contract: contract}}, nil
}

// NewHauskaGroupManagerCaller creates a new read-only instance of HauskaGroupManager, bound to a specific deployed contract.
func NewHauskaGroupManagerCaller(address common.Address, caller bind.ContractCaller) (*HauskaGroupManagerCaller, error) {
	contract, err := bindHauskaGroupManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerCaller{contract: contract}, nil
}

// NewHauskaGroupManagerTransactor creates a new write-only instance of HauskaGroupManager, bound to a specific deployed contract.
func NewHauskaGroupManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*HauskaGroupManagerTransactor, error) {
	contract, err := bindHauskaGroupManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerTransactor{contract: contract}, nil
}

// NewHauskaGroupManagerFilterer creates a new log filterer instance of HauskaGroupManager, bound to a specific deployed contract.
func NewHauskaGroupManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*HauskaGroupManagerFilterer, error) {
	contract, err := bindHauskaGroupManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerFilterer{contract: contract}, nil
}

// bindHauskaGroupManager binds a generic wrapper to an already deployed contract.
func bindHauskaGroupManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HauskaGroupManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaGroupManager *HauskaGroupManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaGroupManager.Contract.HauskaGroupManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaGroupManager *HauskaGroupManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.HauskaGroupManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaGroupManager *HauskaGroupManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.HauskaGroupManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaGroupManager *HauskaGroupManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaGroupManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaGroupManager *HauskaGroupManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaGroupManager *HauskaGroupManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaGroupManager.Contract.DEFAULTADMINROLE(&_HauskaGroupManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaGroupManager.Contract.DEFAULTADMINROLE(&_HauskaGroupManager.CallOpts)
}

// GROUPADMINROLE is a free data retrieval call binding the contract method 0x0cec8e7a.
//
// Solidity: function GROUP_ADMIN_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCaller) GROUPADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "GROUP_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GROUPADMINROLE is a free data retrieval call binding the contract method 0x0cec8e7a.
//
// Solidity: function GROUP_ADMIN_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerSession) GROUPADMINROLE() ([32]byte, error) {
	return _HauskaGroupManager.Contract.GROUPADMINROLE(&_HauskaGroupManager.CallOpts)
}

// GROUPADMINROLE is a free data retrieval call binding the contract method 0x0cec8e7a.
//
// Solidity: function GROUP_ADMIN_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) GROUPADMINROLE() ([32]byte, error) {
	return _HauskaGroupManager.Contract.GROUPADMINROLE(&_HauskaGroupManager.CallOpts)
}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCaller) ORGCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "ORG_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerSession) ORGCONTRACTROLE() ([32]byte, error) {
	return _HauskaGroupManager.Contract.ORGCONTRACTROLE(&_HauskaGroupManager.CallOpts)
}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) ORGCONTRACTROLE() ([32]byte, error) {
	return _HauskaGroupManager.Contract.ORGCONTRACTROLE(&_HauskaGroupManager.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaGroupManager *HauskaGroupManagerCaller) AssetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "assetRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaGroupManager *HauskaGroupManagerSession) AssetRegistry() (common.Address, error) {
	return _HauskaGroupManager.Contract.AssetRegistry(&_HauskaGroupManager.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) AssetRegistry() (common.Address, error) {
	return _HauskaGroupManager.Contract.AssetRegistry(&_HauskaGroupManager.CallOpts)
}

// CreatorGroups is a free data retrieval call binding the contract method 0x842fd84c.
//
// Solidity: function creatorGroups(address , address , uint256 ) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerCaller) CreatorGroups(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "creatorGroups", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorGroups is a free data retrieval call binding the contract method 0x842fd84c.
//
// Solidity: function creatorGroups(address , address , uint256 ) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerSession) CreatorGroups(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _HauskaGroupManager.Contract.CreatorGroups(&_HauskaGroupManager.CallOpts, arg0, arg1, arg2)
}

// CreatorGroups is a free data retrieval call binding the contract method 0x842fd84c.
//
// Solidity: function creatorGroups(address , address , uint256 ) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) CreatorGroups(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _HauskaGroupManager.Contract.CreatorGroups(&_HauskaGroupManager.CallOpts, arg0, arg1, arg2)
}

// GetGroup is a free data retrieval call binding the contract method 0x6a0f9e76.
//
// Solidity: function getGroup(address orgContract, uint256 groupId) view returns((uint256,uint256[],string,uint256,address))
func (_HauskaGroupManager *HauskaGroupManagerCaller) GetGroup(opts *bind.CallOpts, orgContract common.Address, groupId *big.Int) (IHauskaStructsAssetGroup, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "getGroup", orgContract, groupId)

	if err != nil {
		return *new(IHauskaStructsAssetGroup), err
	}

	out0 := *abi.ConvertType(out[0], new(IHauskaStructsAssetGroup)).(*IHauskaStructsAssetGroup)

	return out0, err

}

// GetGroup is a free data retrieval call binding the contract method 0x6a0f9e76.
//
// Solidity: function getGroup(address orgContract, uint256 groupId) view returns((uint256,uint256[],string,uint256,address))
func (_HauskaGroupManager *HauskaGroupManagerSession) GetGroup(orgContract common.Address, groupId *big.Int) (IHauskaStructsAssetGroup, error) {
	return _HauskaGroupManager.Contract.GetGroup(&_HauskaGroupManager.CallOpts, orgContract, groupId)
}

// GetGroup is a free data retrieval call binding the contract method 0x6a0f9e76.
//
// Solidity: function getGroup(address orgContract, uint256 groupId) view returns((uint256,uint256[],string,uint256,address))
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) GetGroup(orgContract common.Address, groupId *big.Int) (IHauskaStructsAssetGroup, error) {
	return _HauskaGroupManager.Contract.GetGroup(&_HauskaGroupManager.CallOpts, orgContract, groupId)
}

// GetGroupCount is a free data retrieval call binding the contract method 0x96a2f896.
//
// Solidity: function getGroupCount(address orgContract) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerCaller) GetGroupCount(opts *bind.CallOpts, orgContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "getGroupCount", orgContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGroupCount is a free data retrieval call binding the contract method 0x96a2f896.
//
// Solidity: function getGroupCount(address orgContract) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerSession) GetGroupCount(orgContract common.Address) (*big.Int, error) {
	return _HauskaGroupManager.Contract.GetGroupCount(&_HauskaGroupManager.CallOpts, orgContract)
}

// GetGroupCount is a free data retrieval call binding the contract method 0x96a2f896.
//
// Solidity: function getGroupCount(address orgContract) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) GetGroupCount(orgContract common.Address) (*big.Int, error) {
	return _HauskaGroupManager.Contract.GetGroupCount(&_HauskaGroupManager.CallOpts, orgContract)
}

// GetGroupsByCreator is a free data retrieval call binding the contract method 0x530fbd59.
//
// Solidity: function getGroupsByCreator(address orgContract, address creator) view returns(uint256[])
func (_HauskaGroupManager *HauskaGroupManagerCaller) GetGroupsByCreator(opts *bind.CallOpts, orgContract common.Address, creator common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "getGroupsByCreator", orgContract, creator)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGroupsByCreator is a free data retrieval call binding the contract method 0x530fbd59.
//
// Solidity: function getGroupsByCreator(address orgContract, address creator) view returns(uint256[])
func (_HauskaGroupManager *HauskaGroupManagerSession) GetGroupsByCreator(orgContract common.Address, creator common.Address) ([]*big.Int, error) {
	return _HauskaGroupManager.Contract.GetGroupsByCreator(&_HauskaGroupManager.CallOpts, orgContract, creator)
}

// GetGroupsByCreator is a free data retrieval call binding the contract method 0x530fbd59.
//
// Solidity: function getGroupsByCreator(address orgContract, address creator) view returns(uint256[])
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) GetGroupsByCreator(orgContract common.Address, creator common.Address) ([]*big.Int, error) {
	return _HauskaGroupManager.Contract.GetGroupsByCreator(&_HauskaGroupManager.CallOpts, orgContract, creator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaGroupManager.Contract.GetRoleAdmin(&_HauskaGroupManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaGroupManager.Contract.GetRoleAdmin(&_HauskaGroupManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaGroupManager *HauskaGroupManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaGroupManager *HauskaGroupManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaGroupManager.Contract.HasRole(&_HauskaGroupManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaGroupManager.Contract.HasRole(&_HauskaGroupManager.CallOpts, role, account)
}

// OrgGroupCounts is a free data retrieval call binding the contract method 0x6843b7e8.
//
// Solidity: function orgGroupCounts(address ) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerCaller) OrgGroupCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "orgGroupCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrgGroupCounts is a free data retrieval call binding the contract method 0x6843b7e8.
//
// Solidity: function orgGroupCounts(address ) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerSession) OrgGroupCounts(arg0 common.Address) (*big.Int, error) {
	return _HauskaGroupManager.Contract.OrgGroupCounts(&_HauskaGroupManager.CallOpts, arg0)
}

// OrgGroupCounts is a free data retrieval call binding the contract method 0x6843b7e8.
//
// Solidity: function orgGroupCounts(address ) view returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) OrgGroupCounts(arg0 common.Address) (*big.Int, error) {
	return _HauskaGroupManager.Contract.OrgGroupCounts(&_HauskaGroupManager.CallOpts, arg0)
}

// OrgGroups is a free data retrieval call binding the contract method 0x04bcd1c7.
//
// Solidity: function orgGroups(address , uint256 ) view returns(uint256 groupId, string name, uint256 groupPrice, address owner)
func (_HauskaGroupManager *HauskaGroupManagerCaller) OrgGroups(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	GroupId    *big.Int
	Name       string
	GroupPrice *big.Int
	Owner      common.Address
}, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "orgGroups", arg0, arg1)

	outstruct := new(struct {
		GroupId    *big.Int
		Name       string
		GroupPrice *big.Int
		Owner      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GroupId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.GroupPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// OrgGroups is a free data retrieval call binding the contract method 0x04bcd1c7.
//
// Solidity: function orgGroups(address , uint256 ) view returns(uint256 groupId, string name, uint256 groupPrice, address owner)
func (_HauskaGroupManager *HauskaGroupManagerSession) OrgGroups(arg0 common.Address, arg1 *big.Int) (struct {
	GroupId    *big.Int
	Name       string
	GroupPrice *big.Int
	Owner      common.Address
}, error) {
	return _HauskaGroupManager.Contract.OrgGroups(&_HauskaGroupManager.CallOpts, arg0, arg1)
}

// OrgGroups is a free data retrieval call binding the contract method 0x04bcd1c7.
//
// Solidity: function orgGroups(address , uint256 ) view returns(uint256 groupId, string name, uint256 groupPrice, address owner)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) OrgGroups(arg0 common.Address, arg1 *big.Int) (struct {
	GroupId    *big.Int
	Name       string
	GroupPrice *big.Int
	Owner      common.Address
}, error) {
	return _HauskaGroupManager.Contract.OrgGroups(&_HauskaGroupManager.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaGroupManager *HauskaGroupManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HauskaGroupManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaGroupManager *HauskaGroupManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaGroupManager.Contract.SupportsInterface(&_HauskaGroupManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaGroupManager *HauskaGroupManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaGroupManager.Contract.SupportsInterface(&_HauskaGroupManager.CallOpts, interfaceId)
}

// AddAssetToGroup is a paid mutator transaction binding the contract method 0xe6e1f5c7.
//
// Solidity: function addAssetToGroup(uint256 groupId, uint256 assetId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) AddAssetToGroup(opts *bind.TransactOpts, groupId *big.Int, assetId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "addAssetToGroup", groupId, assetId, caller)
}

// AddAssetToGroup is a paid mutator transaction binding the contract method 0xe6e1f5c7.
//
// Solidity: function addAssetToGroup(uint256 groupId, uint256 assetId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) AddAssetToGroup(groupId *big.Int, assetId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.AddAssetToGroup(&_HauskaGroupManager.TransactOpts, groupId, assetId, caller)
}

// AddAssetToGroup is a paid mutator transaction binding the contract method 0xe6e1f5c7.
//
// Solidity: function addAssetToGroup(uint256 groupId, uint256 assetId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) AddAssetToGroup(groupId *big.Int, assetId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.AddAssetToGroup(&_HauskaGroupManager.TransactOpts, groupId, assetId, caller)
}

// AddOrgContract is a paid mutator transaction binding the contract method 0xa85e5b61.
//
// Solidity: function addOrgContract(address orgContract) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) AddOrgContract(opts *bind.TransactOpts, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "addOrgContract", orgContract)
}

// AddOrgContract is a paid mutator transaction binding the contract method 0xa85e5b61.
//
// Solidity: function addOrgContract(address orgContract) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) AddOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.AddOrgContract(&_HauskaGroupManager.TransactOpts, orgContract)
}

// AddOrgContract is a paid mutator transaction binding the contract method 0xa85e5b61.
//
// Solidity: function addOrgContract(address orgContract) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) AddOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.AddOrgContract(&_HauskaGroupManager.TransactOpts, orgContract)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xaf2c01b8.
//
// Solidity: function createGroup(string groupName, uint256[] assetIds, uint256 groupPrice, address creator) returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerTransactor) CreateGroup(opts *bind.TransactOpts, groupName string, assetIds []*big.Int, groupPrice *big.Int, creator common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "createGroup", groupName, assetIds, groupPrice, creator)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xaf2c01b8.
//
// Solidity: function createGroup(string groupName, uint256[] assetIds, uint256 groupPrice, address creator) returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerSession) CreateGroup(groupName string, assetIds []*big.Int, groupPrice *big.Int, creator common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.CreateGroup(&_HauskaGroupManager.TransactOpts, groupName, assetIds, groupPrice, creator)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xaf2c01b8.
//
// Solidity: function createGroup(string groupName, uint256[] assetIds, uint256 groupPrice, address creator) returns(uint256)
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) CreateGroup(groupName string, assetIds []*big.Int, groupPrice *big.Int, creator common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.CreateGroup(&_HauskaGroupManager.TransactOpts, groupName, assetIds, groupPrice, creator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.GrantRole(&_HauskaGroupManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.GrantRole(&_HauskaGroupManager.TransactOpts, role, account)
}

// RemoveAssetFromGroup is a paid mutator transaction binding the contract method 0xe4e75bc3.
//
// Solidity: function removeAssetFromGroup(uint256 groupId, uint256 assetId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) RemoveAssetFromGroup(opts *bind.TransactOpts, groupId *big.Int, assetId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "removeAssetFromGroup", groupId, assetId, caller)
}

// RemoveAssetFromGroup is a paid mutator transaction binding the contract method 0xe4e75bc3.
//
// Solidity: function removeAssetFromGroup(uint256 groupId, uint256 assetId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) RemoveAssetFromGroup(groupId *big.Int, assetId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RemoveAssetFromGroup(&_HauskaGroupManager.TransactOpts, groupId, assetId, caller)
}

// RemoveAssetFromGroup is a paid mutator transaction binding the contract method 0xe4e75bc3.
//
// Solidity: function removeAssetFromGroup(uint256 groupId, uint256 assetId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) RemoveAssetFromGroup(groupId *big.Int, assetId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RemoveAssetFromGroup(&_HauskaGroupManager.TransactOpts, groupId, assetId, caller)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0xc5795d5c.
//
// Solidity: function removeGroup(uint256 groupId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) RemoveGroup(opts *bind.TransactOpts, groupId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "removeGroup", groupId, caller)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0xc5795d5c.
//
// Solidity: function removeGroup(uint256 groupId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) RemoveGroup(groupId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RemoveGroup(&_HauskaGroupManager.TransactOpts, groupId, caller)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0xc5795d5c.
//
// Solidity: function removeGroup(uint256 groupId, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) RemoveGroup(groupId *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RemoveGroup(&_HauskaGroupManager.TransactOpts, groupId, caller)
}

// RemoveOrgContract is a paid mutator transaction binding the contract method 0xc009e23b.
//
// Solidity: function removeOrgContract(address orgContract) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) RemoveOrgContract(opts *bind.TransactOpts, orgContract common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "removeOrgContract", orgContract)
}

// RemoveOrgContract is a paid mutator transaction binding the contract method 0xc009e23b.
//
// Solidity: function removeOrgContract(address orgContract) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) RemoveOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RemoveOrgContract(&_HauskaGroupManager.TransactOpts, orgContract)
}

// RemoveOrgContract is a paid mutator transaction binding the contract method 0xc009e23b.
//
// Solidity: function removeOrgContract(address orgContract) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) RemoveOrgContract(orgContract common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RemoveOrgContract(&_HauskaGroupManager.TransactOpts, orgContract)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RenounceRole(&_HauskaGroupManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RenounceRole(&_HauskaGroupManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RevokeRole(&_HauskaGroupManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.RevokeRole(&_HauskaGroupManager.TransactOpts, role, account)
}

// UpdateGroupPrice is a paid mutator transaction binding the contract method 0x2b584506.
//
// Solidity: function updateGroupPrice(uint256 groupId, uint256 newPrice, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactor) UpdateGroupPrice(opts *bind.TransactOpts, groupId *big.Int, newPrice *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.contract.Transact(opts, "updateGroupPrice", groupId, newPrice, caller)
}

// UpdateGroupPrice is a paid mutator transaction binding the contract method 0x2b584506.
//
// Solidity: function updateGroupPrice(uint256 groupId, uint256 newPrice, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerSession) UpdateGroupPrice(groupId *big.Int, newPrice *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.UpdateGroupPrice(&_HauskaGroupManager.TransactOpts, groupId, newPrice, caller)
}

// UpdateGroupPrice is a paid mutator transaction binding the contract method 0x2b584506.
//
// Solidity: function updateGroupPrice(uint256 groupId, uint256 newPrice, address caller) returns()
func (_HauskaGroupManager *HauskaGroupManagerTransactorSession) UpdateGroupPrice(groupId *big.Int, newPrice *big.Int, caller common.Address) (*types.Transaction, error) {
	return _HauskaGroupManager.Contract.UpdateGroupPrice(&_HauskaGroupManager.TransactOpts, groupId, newPrice, caller)
}

// HauskaGroupManagerAssetAddedToGroupIterator is returned from FilterAssetAddedToGroup and is used to iterate over the raw logs and unpacked data for AssetAddedToGroup events raised by the HauskaGroupManager contract.
type HauskaGroupManagerAssetAddedToGroupIterator struct {
	Event *HauskaGroupManagerAssetAddedToGroup // Event containing the contract specifics and raw log

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
func (it *HauskaGroupManagerAssetAddedToGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaGroupManagerAssetAddedToGroup)
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
		it.Event = new(HauskaGroupManagerAssetAddedToGroup)
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
func (it *HauskaGroupManagerAssetAddedToGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaGroupManagerAssetAddedToGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaGroupManagerAssetAddedToGroup represents a AssetAddedToGroup event raised by the HauskaGroupManager contract.
type HauskaGroupManagerAssetAddedToGroup struct {
	OrgContract common.Address
	GroupId     *big.Int
	AssetId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetAddedToGroup is a free log retrieval operation binding the contract event 0xd82fb8c2248aa7b0938ef0df4027edbbeeb69ab2e9cb52123f760e056f849f0f.
//
// Solidity: event AssetAddedToGroup(address indexed orgContract, uint256 indexed groupId, uint256 assetId)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) FilterAssetAddedToGroup(opts *bind.FilterOpts, orgContract []common.Address, groupId []*big.Int) (*HauskaGroupManagerAssetAddedToGroupIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.FilterLogs(opts, "AssetAddedToGroup", orgContractRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerAssetAddedToGroupIterator{contract: _HauskaGroupManager.contract, event: "AssetAddedToGroup", logs: logs, sub: sub}, nil
}

// WatchAssetAddedToGroup is a free log subscription operation binding the contract event 0xd82fb8c2248aa7b0938ef0df4027edbbeeb69ab2e9cb52123f760e056f849f0f.
//
// Solidity: event AssetAddedToGroup(address indexed orgContract, uint256 indexed groupId, uint256 assetId)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) WatchAssetAddedToGroup(opts *bind.WatchOpts, sink chan<- *HauskaGroupManagerAssetAddedToGroup, orgContract []common.Address, groupId []*big.Int) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.WatchLogs(opts, "AssetAddedToGroup", orgContractRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaGroupManagerAssetAddedToGroup)
				if err := _HauskaGroupManager.contract.UnpackLog(event, "AssetAddedToGroup", log); err != nil {
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

// ParseAssetAddedToGroup is a log parse operation binding the contract event 0xd82fb8c2248aa7b0938ef0df4027edbbeeb69ab2e9cb52123f760e056f849f0f.
//
// Solidity: event AssetAddedToGroup(address indexed orgContract, uint256 indexed groupId, uint256 assetId)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) ParseAssetAddedToGroup(log types.Log) (*HauskaGroupManagerAssetAddedToGroup, error) {
	event := new(HauskaGroupManagerAssetAddedToGroup)
	if err := _HauskaGroupManager.contract.UnpackLog(event, "AssetAddedToGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaGroupManagerAssetRemovedFromGroupIterator is returned from FilterAssetRemovedFromGroup and is used to iterate over the raw logs and unpacked data for AssetRemovedFromGroup events raised by the HauskaGroupManager contract.
type HauskaGroupManagerAssetRemovedFromGroupIterator struct {
	Event *HauskaGroupManagerAssetRemovedFromGroup // Event containing the contract specifics and raw log

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
func (it *HauskaGroupManagerAssetRemovedFromGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaGroupManagerAssetRemovedFromGroup)
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
		it.Event = new(HauskaGroupManagerAssetRemovedFromGroup)
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
func (it *HauskaGroupManagerAssetRemovedFromGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaGroupManagerAssetRemovedFromGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaGroupManagerAssetRemovedFromGroup represents a AssetRemovedFromGroup event raised by the HauskaGroupManager contract.
type HauskaGroupManagerAssetRemovedFromGroup struct {
	OrgContract common.Address
	GroupId     *big.Int
	AssetId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetRemovedFromGroup is a free log retrieval operation binding the contract event 0xb9f865abbd60ec0111b3333ac3f34b2d97125a764e21f7109705d10794ab305f.
//
// Solidity: event AssetRemovedFromGroup(address indexed orgContract, uint256 indexed groupId, uint256 assetId)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) FilterAssetRemovedFromGroup(opts *bind.FilterOpts, orgContract []common.Address, groupId []*big.Int) (*HauskaGroupManagerAssetRemovedFromGroupIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.FilterLogs(opts, "AssetRemovedFromGroup", orgContractRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerAssetRemovedFromGroupIterator{contract: _HauskaGroupManager.contract, event: "AssetRemovedFromGroup", logs: logs, sub: sub}, nil
}

// WatchAssetRemovedFromGroup is a free log subscription operation binding the contract event 0xb9f865abbd60ec0111b3333ac3f34b2d97125a764e21f7109705d10794ab305f.
//
// Solidity: event AssetRemovedFromGroup(address indexed orgContract, uint256 indexed groupId, uint256 assetId)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) WatchAssetRemovedFromGroup(opts *bind.WatchOpts, sink chan<- *HauskaGroupManagerAssetRemovedFromGroup, orgContract []common.Address, groupId []*big.Int) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.WatchLogs(opts, "AssetRemovedFromGroup", orgContractRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaGroupManagerAssetRemovedFromGroup)
				if err := _HauskaGroupManager.contract.UnpackLog(event, "AssetRemovedFromGroup", log); err != nil {
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

// ParseAssetRemovedFromGroup is a log parse operation binding the contract event 0xb9f865abbd60ec0111b3333ac3f34b2d97125a764e21f7109705d10794ab305f.
//
// Solidity: event AssetRemovedFromGroup(address indexed orgContract, uint256 indexed groupId, uint256 assetId)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) ParseAssetRemovedFromGroup(log types.Log) (*HauskaGroupManagerAssetRemovedFromGroup, error) {
	event := new(HauskaGroupManagerAssetRemovedFromGroup)
	if err := _HauskaGroupManager.contract.UnpackLog(event, "AssetRemovedFromGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaGroupManagerGroupCreatedIterator is returned from FilterGroupCreated and is used to iterate over the raw logs and unpacked data for GroupCreated events raised by the HauskaGroupManager contract.
type HauskaGroupManagerGroupCreatedIterator struct {
	Event *HauskaGroupManagerGroupCreated // Event containing the contract specifics and raw log

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
func (it *HauskaGroupManagerGroupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaGroupManagerGroupCreated)
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
		it.Event = new(HauskaGroupManagerGroupCreated)
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
func (it *HauskaGroupManagerGroupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaGroupManagerGroupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaGroupManagerGroupCreated represents a GroupCreated event raised by the HauskaGroupManager contract.
type HauskaGroupManagerGroupCreated struct {
	OrgContract common.Address
	GroupId     *big.Int
	Name        string
	Creator     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGroupCreated is a free log retrieval operation binding the contract event 0x7c3d7e35d9229130c2dc5f5f23f1dc8f70a778bc95606f5a114c4e3ecae2a5d2.
//
// Solidity: event GroupCreated(address indexed orgContract, uint256 indexed groupId, string name, address indexed creator)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) FilterGroupCreated(opts *bind.FilterOpts, orgContract []common.Address, groupId []*big.Int, creator []common.Address) (*HauskaGroupManagerGroupCreatedIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.FilterLogs(opts, "GroupCreated", orgContractRule, groupIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerGroupCreatedIterator{contract: _HauskaGroupManager.contract, event: "GroupCreated", logs: logs, sub: sub}, nil
}

// WatchGroupCreated is a free log subscription operation binding the contract event 0x7c3d7e35d9229130c2dc5f5f23f1dc8f70a778bc95606f5a114c4e3ecae2a5d2.
//
// Solidity: event GroupCreated(address indexed orgContract, uint256 indexed groupId, string name, address indexed creator)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) WatchGroupCreated(opts *bind.WatchOpts, sink chan<- *HauskaGroupManagerGroupCreated, orgContract []common.Address, groupId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.WatchLogs(opts, "GroupCreated", orgContractRule, groupIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaGroupManagerGroupCreated)
				if err := _HauskaGroupManager.contract.UnpackLog(event, "GroupCreated", log); err != nil {
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

// ParseGroupCreated is a log parse operation binding the contract event 0x7c3d7e35d9229130c2dc5f5f23f1dc8f70a778bc95606f5a114c4e3ecae2a5d2.
//
// Solidity: event GroupCreated(address indexed orgContract, uint256 indexed groupId, string name, address indexed creator)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) ParseGroupCreated(log types.Log) (*HauskaGroupManagerGroupCreated, error) {
	event := new(HauskaGroupManagerGroupCreated)
	if err := _HauskaGroupManager.contract.UnpackLog(event, "GroupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaGroupManagerGroupUpdatedIterator is returned from FilterGroupUpdated and is used to iterate over the raw logs and unpacked data for GroupUpdated events raised by the HauskaGroupManager contract.
type HauskaGroupManagerGroupUpdatedIterator struct {
	Event *HauskaGroupManagerGroupUpdated // Event containing the contract specifics and raw log

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
func (it *HauskaGroupManagerGroupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaGroupManagerGroupUpdated)
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
		it.Event = new(HauskaGroupManagerGroupUpdated)
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
func (it *HauskaGroupManagerGroupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaGroupManagerGroupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaGroupManagerGroupUpdated represents a GroupUpdated event raised by the HauskaGroupManager contract.
type HauskaGroupManagerGroupUpdated struct {
	OrgContract common.Address
	GroupId     *big.Int
	NewPrice    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGroupUpdated is a free log retrieval operation binding the contract event 0x336a857c369bd2769bf49b8157ac34d27f780014189c9a67f72597fe4e213f9c.
//
// Solidity: event GroupUpdated(address indexed orgContract, uint256 indexed groupId, uint256 newPrice)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) FilterGroupUpdated(opts *bind.FilterOpts, orgContract []common.Address, groupId []*big.Int) (*HauskaGroupManagerGroupUpdatedIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.FilterLogs(opts, "GroupUpdated", orgContractRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerGroupUpdatedIterator{contract: _HauskaGroupManager.contract, event: "GroupUpdated", logs: logs, sub: sub}, nil
}

// WatchGroupUpdated is a free log subscription operation binding the contract event 0x336a857c369bd2769bf49b8157ac34d27f780014189c9a67f72597fe4e213f9c.
//
// Solidity: event GroupUpdated(address indexed orgContract, uint256 indexed groupId, uint256 newPrice)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) WatchGroupUpdated(opts *bind.WatchOpts, sink chan<- *HauskaGroupManagerGroupUpdated, orgContract []common.Address, groupId []*big.Int) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _HauskaGroupManager.contract.WatchLogs(opts, "GroupUpdated", orgContractRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaGroupManagerGroupUpdated)
				if err := _HauskaGroupManager.contract.UnpackLog(event, "GroupUpdated", log); err != nil {
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

// ParseGroupUpdated is a log parse operation binding the contract event 0x336a857c369bd2769bf49b8157ac34d27f780014189c9a67f72597fe4e213f9c.
//
// Solidity: event GroupUpdated(address indexed orgContract, uint256 indexed groupId, uint256 newPrice)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) ParseGroupUpdated(log types.Log) (*HauskaGroupManagerGroupUpdated, error) {
	event := new(HauskaGroupManagerGroupUpdated)
	if err := _HauskaGroupManager.contract.UnpackLog(event, "GroupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaGroupManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HauskaGroupManager contract.
type HauskaGroupManagerRoleAdminChangedIterator struct {
	Event *HauskaGroupManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HauskaGroupManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaGroupManagerRoleAdminChanged)
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
		it.Event = new(HauskaGroupManagerRoleAdminChanged)
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
func (it *HauskaGroupManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaGroupManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaGroupManagerRoleAdminChanged represents a RoleAdminChanged event raised by the HauskaGroupManager contract.
type HauskaGroupManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HauskaGroupManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HauskaGroupManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerRoleAdminChangedIterator{contract: _HauskaGroupManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HauskaGroupManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HauskaGroupManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaGroupManagerRoleAdminChanged)
				if err := _HauskaGroupManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HauskaGroupManager *HauskaGroupManagerFilterer) ParseRoleAdminChanged(log types.Log) (*HauskaGroupManagerRoleAdminChanged, error) {
	event := new(HauskaGroupManagerRoleAdminChanged)
	if err := _HauskaGroupManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaGroupManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HauskaGroupManager contract.
type HauskaGroupManagerRoleGrantedIterator struct {
	Event *HauskaGroupManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *HauskaGroupManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaGroupManagerRoleGranted)
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
		it.Event = new(HauskaGroupManagerRoleGranted)
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
func (it *HauskaGroupManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaGroupManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaGroupManagerRoleGranted represents a RoleGranted event raised by the HauskaGroupManager contract.
type HauskaGroupManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaGroupManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _HauskaGroupManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerRoleGrantedIterator{contract: _HauskaGroupManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HauskaGroupManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaGroupManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaGroupManagerRoleGranted)
				if err := _HauskaGroupManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HauskaGroupManager *HauskaGroupManagerFilterer) ParseRoleGranted(log types.Log) (*HauskaGroupManagerRoleGranted, error) {
	event := new(HauskaGroupManagerRoleGranted)
	if err := _HauskaGroupManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaGroupManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HauskaGroupManager contract.
type HauskaGroupManagerRoleRevokedIterator struct {
	Event *HauskaGroupManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HauskaGroupManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaGroupManagerRoleRevoked)
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
		it.Event = new(HauskaGroupManagerRoleRevoked)
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
func (it *HauskaGroupManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaGroupManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaGroupManagerRoleRevoked represents a RoleRevoked event raised by the HauskaGroupManager contract.
type HauskaGroupManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaGroupManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _HauskaGroupManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaGroupManagerRoleRevokedIterator{contract: _HauskaGroupManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaGroupManager *HauskaGroupManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HauskaGroupManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaGroupManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaGroupManagerRoleRevoked)
				if err := _HauskaGroupManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HauskaGroupManager *HauskaGroupManagerFilterer) ParseRoleRevoked(log types.Log) (*HauskaGroupManagerRoleRevoked, error) {
	event := new(HauskaGroupManagerRoleRevoked)
	if err := _HauskaGroupManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
