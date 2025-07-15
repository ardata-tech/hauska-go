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

// HauskaLicenseManagerMetaData contains all meta data concerning the HauskaLicenseManager contract.
var HauskaLicenseManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originalLicenseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newLicenseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reseller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLicensee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"name\":\"AssetRelicensed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"licensee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"licenseIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"name\":\"GroupLicensed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"licensee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LicenseGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LicenseRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"LicenseRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIHauskaStructs.CountryCode\",\"name\":\"location\",\"type\":\"uint8\"}],\"name\":\"UserLocationSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORG_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetLicensedBy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"getLicenseDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"licensor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"licensee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserLicenses\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"internalType\":\"enumIHauskaStructs.LicensePermissions\",\"name\":\"permission\",\"type\":\"uint8\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasUserLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAssetLicensedBy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"isLicenseValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"licensee\",\"type\":\"address\"},{\"internalType\":\"enumIHauskaStructs.LicensePermissions[]\",\"name\":\"permissions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"resellerFee\",\"type\":\"uint256\"}],\"name\":\"licenseAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"licensee\",\"type\":\"address\"},{\"internalType\":\"enumIHauskaStructs.LicensePermissions[]\",\"name\":\"permissions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"resellerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"licenseAssetWithDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"licensee\",\"type\":\"address\"},{\"internalType\":\"enumIHauskaStructs.LicensePermissions[]\",\"name\":\"permissions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"resellerFee\",\"type\":\"uint256\"}],\"name\":\"licenseGroup\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"licensor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"licensee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resellerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permissions\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"orgContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originalLicenseId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newLicensee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"markup\",\"type\":\"uint256\"}],\"name\":\"relicenseAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalDuration\",\"type\":\"uint256\"}],\"name\":\"renewLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"revokeLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumIHauskaStructs.CountryCode\",\"name\":\"location\",\"type\":\"uint8\"}],\"name\":\"setUserLocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLocations\",\"outputs\":[{\"internalType\":\"enumIHauskaStructs.CountryCode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HauskaLicenseManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use HauskaLicenseManagerMetaData.ABI instead.
var HauskaLicenseManagerABI = HauskaLicenseManagerMetaData.ABI

// HauskaLicenseManager is an auto generated Go binding around an Ethereum contract.
type HauskaLicenseManager struct {
	HauskaLicenseManagerCaller     // Read-only binding to the contract
	HauskaLicenseManagerTransactor // Write-only binding to the contract
	HauskaLicenseManagerFilterer   // Log filterer for contract events
}

// HauskaLicenseManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type HauskaLicenseManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaLicenseManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HauskaLicenseManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaLicenseManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HauskaLicenseManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaLicenseManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HauskaLicenseManagerSession struct {
	Contract     *HauskaLicenseManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HauskaLicenseManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HauskaLicenseManagerCallerSession struct {
	Contract *HauskaLicenseManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// HauskaLicenseManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HauskaLicenseManagerTransactorSession struct {
	Contract     *HauskaLicenseManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// HauskaLicenseManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type HauskaLicenseManagerRaw struct {
	Contract *HauskaLicenseManager // Generic contract binding to access the raw methods on
}

// HauskaLicenseManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HauskaLicenseManagerCallerRaw struct {
	Contract *HauskaLicenseManagerCaller // Generic read-only contract binding to access the raw methods on
}

// HauskaLicenseManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HauskaLicenseManagerTransactorRaw struct {
	Contract *HauskaLicenseManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHauskaLicenseManager creates a new instance of HauskaLicenseManager, bound to a specific deployed contract.
func NewHauskaLicenseManager(address common.Address, backend bind.ContractBackend) (*HauskaLicenseManager, error) {
	contract, err := bindHauskaLicenseManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManager{HauskaLicenseManagerCaller: HauskaLicenseManagerCaller{contract: contract}, HauskaLicenseManagerTransactor: HauskaLicenseManagerTransactor{contract: contract}, HauskaLicenseManagerFilterer: HauskaLicenseManagerFilterer{contract: contract}}, nil
}

// NewHauskaLicenseManagerCaller creates a new read-only instance of HauskaLicenseManager, bound to a specific deployed contract.
func NewHauskaLicenseManagerCaller(address common.Address, caller bind.ContractCaller) (*HauskaLicenseManagerCaller, error) {
	contract, err := bindHauskaLicenseManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerCaller{contract: contract}, nil
}

// NewHauskaLicenseManagerTransactor creates a new write-only instance of HauskaLicenseManager, bound to a specific deployed contract.
func NewHauskaLicenseManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*HauskaLicenseManagerTransactor, error) {
	contract, err := bindHauskaLicenseManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerTransactor{contract: contract}, nil
}

// NewHauskaLicenseManagerFilterer creates a new log filterer instance of HauskaLicenseManager, bound to a specific deployed contract.
func NewHauskaLicenseManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*HauskaLicenseManagerFilterer, error) {
	contract, err := bindHauskaLicenseManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerFilterer{contract: contract}, nil
}

// bindHauskaLicenseManager binds a generic wrapper to an already deployed contract.
func bindHauskaLicenseManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HauskaLicenseManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaLicenseManager *HauskaLicenseManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaLicenseManager.Contract.HauskaLicenseManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaLicenseManager *HauskaLicenseManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.HauskaLicenseManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaLicenseManager *HauskaLicenseManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.HauskaLicenseManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaLicenseManager *HauskaLicenseManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaLicenseManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) ADMINROLE() ([32]byte, error) {
	return _HauskaLicenseManager.Contract.ADMINROLE(&_HauskaLicenseManager.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) ADMINROLE() ([32]byte, error) {
	return _HauskaLicenseManager.Contract.ADMINROLE(&_HauskaLicenseManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaLicenseManager.Contract.DEFAULTADMINROLE(&_HauskaLicenseManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaLicenseManager.Contract.DEFAULTADMINROLE(&_HauskaLicenseManager.CallOpts)
}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) ORGCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "ORG_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) ORGCONTRACTROLE() ([32]byte, error) {
	return _HauskaLicenseManager.Contract.ORGCONTRACTROLE(&_HauskaLicenseManager.CallOpts)
}

// ORGCONTRACTROLE is a free data retrieval call binding the contract method 0x96f5a291.
//
// Solidity: function ORG_CONTRACT_ROLE() view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) ORGCONTRACTROLE() ([32]byte, error) {
	return _HauskaLicenseManager.Contract.ORGCONTRACTROLE(&_HauskaLicenseManager.CallOpts)
}

// AssetLicensedBy is a free data retrieval call binding the contract method 0xe28384cf.
//
// Solidity: function assetLicensedBy(address , uint256 , address ) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) AssetLicensedBy(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "assetLicensedBy", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetLicensedBy is a free data retrieval call binding the contract method 0xe28384cf.
//
// Solidity: function assetLicensedBy(address , uint256 , address ) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) AssetLicensedBy(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.AssetLicensedBy(&_HauskaLicenseManager.CallOpts, arg0, arg1, arg2)
}

// AssetLicensedBy is a free data retrieval call binding the contract method 0xe28384cf.
//
// Solidity: function assetLicensedBy(address , uint256 , address ) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) AssetLicensedBy(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.AssetLicensedBy(&_HauskaLicenseManager.CallOpts, arg0, arg1, arg2)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) Factory() (common.Address, error) {
	return _HauskaLicenseManager.Contract.Factory(&_HauskaLicenseManager.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) Factory() (common.Address, error) {
	return _HauskaLicenseManager.Contract.Factory(&_HauskaLicenseManager.CallOpts)
}

// GetLicenseDetails is a free data retrieval call binding the contract method 0x967d56fd.
//
// Solidity: function getLicenseDetails(uint256 licenseId) view returns(uint256 assetId, address licensor, address licensee, uint256 fee, address orgContract)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) GetLicenseDetails(opts *bind.CallOpts, licenseId *big.Int) (struct {
	AssetId     *big.Int
	Licensor    common.Address
	Licensee    common.Address
	Fee         *big.Int
	OrgContract common.Address
}, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "getLicenseDetails", licenseId)

	outstruct := new(struct {
		AssetId     *big.Int
		Licensor    common.Address
		Licensee    common.Address
		Fee         *big.Int
		OrgContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Licensor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Licensee = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OrgContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetLicenseDetails is a free data retrieval call binding the contract method 0x967d56fd.
//
// Solidity: function getLicenseDetails(uint256 licenseId) view returns(uint256 assetId, address licensor, address licensee, uint256 fee, address orgContract)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) GetLicenseDetails(licenseId *big.Int) (struct {
	AssetId     *big.Int
	Licensor    common.Address
	Licensee    common.Address
	Fee         *big.Int
	OrgContract common.Address
}, error) {
	return _HauskaLicenseManager.Contract.GetLicenseDetails(&_HauskaLicenseManager.CallOpts, licenseId)
}

// GetLicenseDetails is a free data retrieval call binding the contract method 0x967d56fd.
//
// Solidity: function getLicenseDetails(uint256 licenseId) view returns(uint256 assetId, address licensor, address licensee, uint256 fee, address orgContract)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) GetLicenseDetails(licenseId *big.Int) (struct {
	AssetId     *big.Int
	Licensor    common.Address
	Licensee    common.Address
	Fee         *big.Int
	OrgContract common.Address
}, error) {
	return _HauskaLicenseManager.Contract.GetLicenseDetails(&_HauskaLicenseManager.CallOpts, licenseId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaLicenseManager.Contract.GetRoleAdmin(&_HauskaLicenseManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaLicenseManager.Contract.GetRoleAdmin(&_HauskaLicenseManager.CallOpts, role)
}

// GetUserLicenses is a free data retrieval call binding the contract method 0xa9377926.
//
// Solidity: function getUserLicenses(address orgContract, address user) view returns(uint256[])
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) GetUserLicenses(opts *bind.CallOpts, orgContract common.Address, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "getUserLicenses", orgContract, user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserLicenses is a free data retrieval call binding the contract method 0xa9377926.
//
// Solidity: function getUserLicenses(address orgContract, address user) view returns(uint256[])
func (_HauskaLicenseManager *HauskaLicenseManagerSession) GetUserLicenses(orgContract common.Address, user common.Address) ([]*big.Int, error) {
	return _HauskaLicenseManager.Contract.GetUserLicenses(&_HauskaLicenseManager.CallOpts, orgContract, user)
}

// GetUserLicenses is a free data retrieval call binding the contract method 0xa9377926.
//
// Solidity: function getUserLicenses(address orgContract, address user) view returns(uint256[])
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) GetUserLicenses(orgContract common.Address, user common.Address) ([]*big.Int, error) {
	return _HauskaLicenseManager.Contract.GetUserLicenses(&_HauskaLicenseManager.CallOpts, orgContract, user)
}

// HasPermission is a free data retrieval call binding the contract method 0xceb0d747.
//
// Solidity: function hasPermission(uint256 licenseId, uint8 permission) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) HasPermission(opts *bind.CallOpts, licenseId *big.Int, permission uint8) (bool, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "hasPermission", licenseId, permission)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0xceb0d747.
//
// Solidity: function hasPermission(uint256 licenseId, uint8 permission) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) HasPermission(licenseId *big.Int, permission uint8) (bool, error) {
	return _HauskaLicenseManager.Contract.HasPermission(&_HauskaLicenseManager.CallOpts, licenseId, permission)
}

// HasPermission is a free data retrieval call binding the contract method 0xceb0d747.
//
// Solidity: function hasPermission(uint256 licenseId, uint8 permission) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) HasPermission(licenseId *big.Int, permission uint8) (bool, error) {
	return _HauskaLicenseManager.Contract.HasPermission(&_HauskaLicenseManager.CallOpts, licenseId, permission)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.HasRole(&_HauskaLicenseManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.HasRole(&_HauskaLicenseManager.CallOpts, role, account)
}

// HasUserLocation is a free data retrieval call binding the contract method 0x1e62c6d5.
//
// Solidity: function hasUserLocation(address ) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) HasUserLocation(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "hasUserLocation", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserLocation is a free data retrieval call binding the contract method 0x1e62c6d5.
//
// Solidity: function hasUserLocation(address ) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) HasUserLocation(arg0 common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.HasUserLocation(&_HauskaLicenseManager.CallOpts, arg0)
}

// HasUserLocation is a free data retrieval call binding the contract method 0x1e62c6d5.
//
// Solidity: function hasUserLocation(address ) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) HasUserLocation(arg0 common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.HasUserLocation(&_HauskaLicenseManager.CallOpts, arg0)
}

// IsAssetLicensedBy is a free data retrieval call binding the contract method 0x610588d3.
//
// Solidity: function isAssetLicensedBy(address orgContract, uint256 assetId, address user) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) IsAssetLicensedBy(opts *bind.CallOpts, orgContract common.Address, assetId *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "isAssetLicensedBy", orgContract, assetId, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetLicensedBy is a free data retrieval call binding the contract method 0x610588d3.
//
// Solidity: function isAssetLicensedBy(address orgContract, uint256 assetId, address user) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) IsAssetLicensedBy(orgContract common.Address, assetId *big.Int, user common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.IsAssetLicensedBy(&_HauskaLicenseManager.CallOpts, orgContract, assetId, user)
}

// IsAssetLicensedBy is a free data retrieval call binding the contract method 0x610588d3.
//
// Solidity: function isAssetLicensedBy(address orgContract, uint256 assetId, address user) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) IsAssetLicensedBy(orgContract common.Address, assetId *big.Int, user common.Address) (bool, error) {
	return _HauskaLicenseManager.Contract.IsAssetLicensedBy(&_HauskaLicenseManager.CallOpts, orgContract, assetId, user)
}

// IsLicenseValid is a free data retrieval call binding the contract method 0x9f7fc6d3.
//
// Solidity: function isLicenseValid(uint256 licenseId) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) IsLicenseValid(opts *bind.CallOpts, licenseId *big.Int) (bool, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "isLicenseValid", licenseId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLicenseValid is a free data retrieval call binding the contract method 0x9f7fc6d3.
//
// Solidity: function isLicenseValid(uint256 licenseId) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) IsLicenseValid(licenseId *big.Int) (bool, error) {
	return _HauskaLicenseManager.Contract.IsLicenseValid(&_HauskaLicenseManager.CallOpts, licenseId)
}

// IsLicenseValid is a free data retrieval call binding the contract method 0x9f7fc6d3.
//
// Solidity: function isLicenseValid(uint256 licenseId) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) IsLicenseValid(licenseId *big.Int) (bool, error) {
	return _HauskaLicenseManager.Contract.IsLicenseValid(&_HauskaLicenseManager.CallOpts, licenseId)
}

// Licenses is a free data retrieval call binding the contract method 0x33790845.
//
// Solidity: function licenses(uint256 ) view returns(uint256 assetId, address licensor, address licensee, uint256 fee, uint256 resellerFee, uint8 permissions, address orgContract, uint256 expirationTime, bool isActive)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) Licenses(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AssetId        *big.Int
	Licensor       common.Address
	Licensee       common.Address
	Fee            *big.Int
	ResellerFee    *big.Int
	Permissions    uint8
	OrgContract    common.Address
	ExpirationTime *big.Int
	IsActive       bool
}, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "licenses", arg0)

	outstruct := new(struct {
		AssetId        *big.Int
		Licensor       common.Address
		Licensee       common.Address
		Fee            *big.Int
		ResellerFee    *big.Int
		Permissions    uint8
		OrgContract    common.Address
		ExpirationTime *big.Int
		IsActive       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Licensor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Licensee = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ResellerFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Permissions = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.OrgContract = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ExpirationTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Licenses is a free data retrieval call binding the contract method 0x33790845.
//
// Solidity: function licenses(uint256 ) view returns(uint256 assetId, address licensor, address licensee, uint256 fee, uint256 resellerFee, uint8 permissions, address orgContract, uint256 expirationTime, bool isActive)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) Licenses(arg0 *big.Int) (struct {
	AssetId        *big.Int
	Licensor       common.Address
	Licensee       common.Address
	Fee            *big.Int
	ResellerFee    *big.Int
	Permissions    uint8
	OrgContract    common.Address
	ExpirationTime *big.Int
	IsActive       bool
}, error) {
	return _HauskaLicenseManager.Contract.Licenses(&_HauskaLicenseManager.CallOpts, arg0)
}

// Licenses is a free data retrieval call binding the contract method 0x33790845.
//
// Solidity: function licenses(uint256 ) view returns(uint256 assetId, address licensor, address licensee, uint256 fee, uint256 resellerFee, uint8 permissions, address orgContract, uint256 expirationTime, bool isActive)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) Licenses(arg0 *big.Int) (struct {
	AssetId        *big.Int
	Licensor       common.Address
	Licensee       common.Address
	Fee            *big.Int
	ResellerFee    *big.Int
	Permissions    uint8
	OrgContract    common.Address
	ExpirationTime *big.Int
	IsActive       bool
}, error) {
	return _HauskaLicenseManager.Contract.Licenses(&_HauskaLicenseManager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaLicenseManager.Contract.SupportsInterface(&_HauskaLicenseManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaLicenseManager.Contract.SupportsInterface(&_HauskaLicenseManager.CallOpts, interfaceId)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) UsdcToken() (common.Address, error) {
	return _HauskaLicenseManager.Contract.UsdcToken(&_HauskaLicenseManager.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) UsdcToken() (common.Address, error) {
	return _HauskaLicenseManager.Contract.UsdcToken(&_HauskaLicenseManager.CallOpts)
}

// UserLicenses is a free data retrieval call binding the contract method 0x4a88a387.
//
// Solidity: function userLicenses(address , address , uint256 ) view returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) UserLicenses(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "userLicenses", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLicenses is a free data retrieval call binding the contract method 0x4a88a387.
//
// Solidity: function userLicenses(address , address , uint256 ) view returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) UserLicenses(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _HauskaLicenseManager.Contract.UserLicenses(&_HauskaLicenseManager.CallOpts, arg0, arg1, arg2)
}

// UserLicenses is a free data retrieval call binding the contract method 0x4a88a387.
//
// Solidity: function userLicenses(address , address , uint256 ) view returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) UserLicenses(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _HauskaLicenseManager.Contract.UserLicenses(&_HauskaLicenseManager.CallOpts, arg0, arg1, arg2)
}

// UserLocations is a free data retrieval call binding the contract method 0x10f41ddf.
//
// Solidity: function userLocations(address ) view returns(uint8)
func (_HauskaLicenseManager *HauskaLicenseManagerCaller) UserLocations(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _HauskaLicenseManager.contract.Call(opts, &out, "userLocations", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UserLocations is a free data retrieval call binding the contract method 0x10f41ddf.
//
// Solidity: function userLocations(address ) view returns(uint8)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) UserLocations(arg0 common.Address) (uint8, error) {
	return _HauskaLicenseManager.Contract.UserLocations(&_HauskaLicenseManager.CallOpts, arg0)
}

// UserLocations is a free data retrieval call binding the contract method 0x10f41ddf.
//
// Solidity: function userLocations(address ) view returns(uint8)
func (_HauskaLicenseManager *HauskaLicenseManagerCallerSession) UserLocations(arg0 common.Address) (uint8, error) {
	return _HauskaLicenseManager.Contract.UserLocations(&_HauskaLicenseManager.CallOpts, arg0)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address token, uint256 amount) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) EmergencyWithdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "emergencyWithdraw", token, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address token, uint256 amount) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerSession) EmergencyWithdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.EmergencyWithdraw(&_HauskaLicenseManager.TransactOpts, token, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address token, uint256 amount) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) EmergencyWithdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.EmergencyWithdraw(&_HauskaLicenseManager.TransactOpts, token, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.GrantRole(&_HauskaLicenseManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.GrantRole(&_HauskaLicenseManager.TransactOpts, role, account)
}

// LicenseAsset is a paid mutator transaction binding the contract method 0xbf514751.
//
// Solidity: function licenseAsset(address orgContract, uint256 assetId, address licensee, uint8[] permissions, uint256 resellerFee) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) LicenseAsset(opts *bind.TransactOpts, orgContract common.Address, assetId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "licenseAsset", orgContract, assetId, licensee, permissions, resellerFee)
}

// LicenseAsset is a paid mutator transaction binding the contract method 0xbf514751.
//
// Solidity: function licenseAsset(address orgContract, uint256 assetId, address licensee, uint8[] permissions, uint256 resellerFee) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) LicenseAsset(orgContract common.Address, assetId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.LicenseAsset(&_HauskaLicenseManager.TransactOpts, orgContract, assetId, licensee, permissions, resellerFee)
}

// LicenseAsset is a paid mutator transaction binding the contract method 0xbf514751.
//
// Solidity: function licenseAsset(address orgContract, uint256 assetId, address licensee, uint8[] permissions, uint256 resellerFee) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) LicenseAsset(orgContract common.Address, assetId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.LicenseAsset(&_HauskaLicenseManager.TransactOpts, orgContract, assetId, licensee, permissions, resellerFee)
}

// LicenseAssetWithDuration is a paid mutator transaction binding the contract method 0x78a98369.
//
// Solidity: function licenseAssetWithDuration(address orgContract, uint256 assetId, address licensee, uint8[] permissions, uint256 resellerFee, uint256 duration) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) LicenseAssetWithDuration(opts *bind.TransactOpts, orgContract common.Address, assetId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "licenseAssetWithDuration", orgContract, assetId, licensee, permissions, resellerFee, duration)
}

// LicenseAssetWithDuration is a paid mutator transaction binding the contract method 0x78a98369.
//
// Solidity: function licenseAssetWithDuration(address orgContract, uint256 assetId, address licensee, uint8[] permissions, uint256 resellerFee, uint256 duration) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) LicenseAssetWithDuration(orgContract common.Address, assetId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.LicenseAssetWithDuration(&_HauskaLicenseManager.TransactOpts, orgContract, assetId, licensee, permissions, resellerFee, duration)
}

// LicenseAssetWithDuration is a paid mutator transaction binding the contract method 0x78a98369.
//
// Solidity: function licenseAssetWithDuration(address orgContract, uint256 assetId, address licensee, uint8[] permissions, uint256 resellerFee, uint256 duration) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) LicenseAssetWithDuration(orgContract common.Address, assetId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.LicenseAssetWithDuration(&_HauskaLicenseManager.TransactOpts, orgContract, assetId, licensee, permissions, resellerFee, duration)
}

// LicenseGroup is a paid mutator transaction binding the contract method 0x191e49b0.
//
// Solidity: function licenseGroup(address orgContract, uint256 groupId, address licensee, uint8[] permissions, uint256 resellerFee) returns(uint256[])
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) LicenseGroup(opts *bind.TransactOpts, orgContract common.Address, groupId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "licenseGroup", orgContract, groupId, licensee, permissions, resellerFee)
}

// LicenseGroup is a paid mutator transaction binding the contract method 0x191e49b0.
//
// Solidity: function licenseGroup(address orgContract, uint256 groupId, address licensee, uint8[] permissions, uint256 resellerFee) returns(uint256[])
func (_HauskaLicenseManager *HauskaLicenseManagerSession) LicenseGroup(orgContract common.Address, groupId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.LicenseGroup(&_HauskaLicenseManager.TransactOpts, orgContract, groupId, licensee, permissions, resellerFee)
}

// LicenseGroup is a paid mutator transaction binding the contract method 0x191e49b0.
//
// Solidity: function licenseGroup(address orgContract, uint256 groupId, address licensee, uint8[] permissions, uint256 resellerFee) returns(uint256[])
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) LicenseGroup(orgContract common.Address, groupId *big.Int, licensee common.Address, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.LicenseGroup(&_HauskaLicenseManager.TransactOpts, orgContract, groupId, licensee, permissions, resellerFee)
}

// RelicenseAsset is a paid mutator transaction binding the contract method 0x3b8a87c5.
//
// Solidity: function relicenseAsset(uint256 originalLicenseId, address newLicensee, uint256 markup) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) RelicenseAsset(opts *bind.TransactOpts, originalLicenseId *big.Int, newLicensee common.Address, markup *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "relicenseAsset", originalLicenseId, newLicensee, markup)
}

// RelicenseAsset is a paid mutator transaction binding the contract method 0x3b8a87c5.
//
// Solidity: function relicenseAsset(uint256 originalLicenseId, address newLicensee, uint256 markup) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerSession) RelicenseAsset(originalLicenseId *big.Int, newLicensee common.Address, markup *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RelicenseAsset(&_HauskaLicenseManager.TransactOpts, originalLicenseId, newLicensee, markup)
}

// RelicenseAsset is a paid mutator transaction binding the contract method 0x3b8a87c5.
//
// Solidity: function relicenseAsset(uint256 originalLicenseId, address newLicensee, uint256 markup) returns(uint256)
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) RelicenseAsset(originalLicenseId *big.Int, newLicensee common.Address, markup *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RelicenseAsset(&_HauskaLicenseManager.TransactOpts, originalLicenseId, newLicensee, markup)
}

// RenewLicense is a paid mutator transaction binding the contract method 0x525acbc4.
//
// Solidity: function renewLicense(uint256 licenseId, uint256 additionalDuration) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) RenewLicense(opts *bind.TransactOpts, licenseId *big.Int, additionalDuration *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "renewLicense", licenseId, additionalDuration)
}

// RenewLicense is a paid mutator transaction binding the contract method 0x525acbc4.
//
// Solidity: function renewLicense(uint256 licenseId, uint256 additionalDuration) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerSession) RenewLicense(licenseId *big.Int, additionalDuration *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RenewLicense(&_HauskaLicenseManager.TransactOpts, licenseId, additionalDuration)
}

// RenewLicense is a paid mutator transaction binding the contract method 0x525acbc4.
//
// Solidity: function renewLicense(uint256 licenseId, uint256 additionalDuration) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) RenewLicense(licenseId *big.Int, additionalDuration *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RenewLicense(&_HauskaLicenseManager.TransactOpts, licenseId, additionalDuration)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RenounceRole(&_HauskaLicenseManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RenounceRole(&_HauskaLicenseManager.TransactOpts, role, account)
}

// RevokeLicense is a paid mutator transaction binding the contract method 0x439d4b5c.
//
// Solidity: function revokeLicense(uint256 licenseId) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) RevokeLicense(opts *bind.TransactOpts, licenseId *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "revokeLicense", licenseId)
}

// RevokeLicense is a paid mutator transaction binding the contract method 0x439d4b5c.
//
// Solidity: function revokeLicense(uint256 licenseId) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerSession) RevokeLicense(licenseId *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RevokeLicense(&_HauskaLicenseManager.TransactOpts, licenseId)
}

// RevokeLicense is a paid mutator transaction binding the contract method 0x439d4b5c.
//
// Solidity: function revokeLicense(uint256 licenseId) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) RevokeLicense(licenseId *big.Int) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RevokeLicense(&_HauskaLicenseManager.TransactOpts, licenseId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RevokeRole(&_HauskaLicenseManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.RevokeRole(&_HauskaLicenseManager.TransactOpts, role, account)
}

// SetUserLocation is a paid mutator transaction binding the contract method 0x6a4c2d9d.
//
// Solidity: function setUserLocation(address user, uint8 location) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactor) SetUserLocation(opts *bind.TransactOpts, user common.Address, location uint8) (*types.Transaction, error) {
	return _HauskaLicenseManager.contract.Transact(opts, "setUserLocation", user, location)
}

// SetUserLocation is a paid mutator transaction binding the contract method 0x6a4c2d9d.
//
// Solidity: function setUserLocation(address user, uint8 location) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerSession) SetUserLocation(user common.Address, location uint8) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.SetUserLocation(&_HauskaLicenseManager.TransactOpts, user, location)
}

// SetUserLocation is a paid mutator transaction binding the contract method 0x6a4c2d9d.
//
// Solidity: function setUserLocation(address user, uint8 location) returns()
func (_HauskaLicenseManager *HauskaLicenseManagerTransactorSession) SetUserLocation(user common.Address, location uint8) (*types.Transaction, error) {
	return _HauskaLicenseManager.Contract.SetUserLocation(&_HauskaLicenseManager.TransactOpts, user, location)
}

// HauskaLicenseManagerAssetRelicensedIterator is returned from FilterAssetRelicensed and is used to iterate over the raw logs and unpacked data for AssetRelicensed events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerAssetRelicensedIterator struct {
	Event *HauskaLicenseManagerAssetRelicensed // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerAssetRelicensedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerAssetRelicensed)
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
		it.Event = new(HauskaLicenseManagerAssetRelicensed)
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
func (it *HauskaLicenseManagerAssetRelicensedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerAssetRelicensedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerAssetRelicensed represents a AssetRelicensed event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerAssetRelicensed struct {
	OriginalLicenseId *big.Int
	NewLicenseId      *big.Int
	Reseller          common.Address
	NewLicensee       common.Address
	TotalFee          *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAssetRelicensed is a free log retrieval operation binding the contract event 0x16138af9a3affa2b2b8cac832780010625f04919725c0a23af8a8f86848c0695.
//
// Solidity: event AssetRelicensed(uint256 indexed originalLicenseId, uint256 indexed newLicenseId, address indexed reseller, address newLicensee, uint256 totalFee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterAssetRelicensed(opts *bind.FilterOpts, originalLicenseId []*big.Int, newLicenseId []*big.Int, reseller []common.Address) (*HauskaLicenseManagerAssetRelicensedIterator, error) {

	var originalLicenseIdRule []interface{}
	for _, originalLicenseIdItem := range originalLicenseId {
		originalLicenseIdRule = append(originalLicenseIdRule, originalLicenseIdItem)
	}
	var newLicenseIdRule []interface{}
	for _, newLicenseIdItem := range newLicenseId {
		newLicenseIdRule = append(newLicenseIdRule, newLicenseIdItem)
	}
	var resellerRule []interface{}
	for _, resellerItem := range reseller {
		resellerRule = append(resellerRule, resellerItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "AssetRelicensed", originalLicenseIdRule, newLicenseIdRule, resellerRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerAssetRelicensedIterator{contract: _HauskaLicenseManager.contract, event: "AssetRelicensed", logs: logs, sub: sub}, nil
}

// WatchAssetRelicensed is a free log subscription operation binding the contract event 0x16138af9a3affa2b2b8cac832780010625f04919725c0a23af8a8f86848c0695.
//
// Solidity: event AssetRelicensed(uint256 indexed originalLicenseId, uint256 indexed newLicenseId, address indexed reseller, address newLicensee, uint256 totalFee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchAssetRelicensed(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerAssetRelicensed, originalLicenseId []*big.Int, newLicenseId []*big.Int, reseller []common.Address) (event.Subscription, error) {

	var originalLicenseIdRule []interface{}
	for _, originalLicenseIdItem := range originalLicenseId {
		originalLicenseIdRule = append(originalLicenseIdRule, originalLicenseIdItem)
	}
	var newLicenseIdRule []interface{}
	for _, newLicenseIdItem := range newLicenseId {
		newLicenseIdRule = append(newLicenseIdRule, newLicenseIdItem)
	}
	var resellerRule []interface{}
	for _, resellerItem := range reseller {
		resellerRule = append(resellerRule, resellerItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "AssetRelicensed", originalLicenseIdRule, newLicenseIdRule, resellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerAssetRelicensed)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "AssetRelicensed", log); err != nil {
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

// ParseAssetRelicensed is a log parse operation binding the contract event 0x16138af9a3affa2b2b8cac832780010625f04919725c0a23af8a8f86848c0695.
//
// Solidity: event AssetRelicensed(uint256 indexed originalLicenseId, uint256 indexed newLicenseId, address indexed reseller, address newLicensee, uint256 totalFee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseAssetRelicensed(log types.Log) (*HauskaLicenseManagerAssetRelicensed, error) {
	event := new(HauskaLicenseManagerAssetRelicensed)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "AssetRelicensed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerGroupLicensedIterator is returned from FilterGroupLicensed and is used to iterate over the raw logs and unpacked data for GroupLicensed events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerGroupLicensedIterator struct {
	Event *HauskaLicenseManagerGroupLicensed // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerGroupLicensedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerGroupLicensed)
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
		it.Event = new(HauskaLicenseManagerGroupLicensed)
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
func (it *HauskaLicenseManagerGroupLicensedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerGroupLicensedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerGroupLicensed represents a GroupLicensed event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerGroupLicensed struct {
	OrgContract common.Address
	GroupId     *big.Int
	Licensee    common.Address
	LicenseIds  []*big.Int
	TotalFee    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGroupLicensed is a free log retrieval operation binding the contract event 0x637232b5da4cf471a2e0753a5291c9d4e4b725fdc9f2895d531abd8bd14f49b0.
//
// Solidity: event GroupLicensed(address indexed orgContract, uint256 indexed groupId, address indexed licensee, uint256[] licenseIds, uint256 totalFee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterGroupLicensed(opts *bind.FilterOpts, orgContract []common.Address, groupId []*big.Int, licensee []common.Address) (*HauskaLicenseManagerGroupLicensedIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var licenseeRule []interface{}
	for _, licenseeItem := range licensee {
		licenseeRule = append(licenseeRule, licenseeItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "GroupLicensed", orgContractRule, groupIdRule, licenseeRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerGroupLicensedIterator{contract: _HauskaLicenseManager.contract, event: "GroupLicensed", logs: logs, sub: sub}, nil
}

// WatchGroupLicensed is a free log subscription operation binding the contract event 0x637232b5da4cf471a2e0753a5291c9d4e4b725fdc9f2895d531abd8bd14f49b0.
//
// Solidity: event GroupLicensed(address indexed orgContract, uint256 indexed groupId, address indexed licensee, uint256[] licenseIds, uint256 totalFee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchGroupLicensed(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerGroupLicensed, orgContract []common.Address, groupId []*big.Int, licensee []common.Address) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var licenseeRule []interface{}
	for _, licenseeItem := range licensee {
		licenseeRule = append(licenseeRule, licenseeItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "GroupLicensed", orgContractRule, groupIdRule, licenseeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerGroupLicensed)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "GroupLicensed", log); err != nil {
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

// ParseGroupLicensed is a log parse operation binding the contract event 0x637232b5da4cf471a2e0753a5291c9d4e4b725fdc9f2895d531abd8bd14f49b0.
//
// Solidity: event GroupLicensed(address indexed orgContract, uint256 indexed groupId, address indexed licensee, uint256[] licenseIds, uint256 totalFee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseGroupLicensed(log types.Log) (*HauskaLicenseManagerGroupLicensed, error) {
	event := new(HauskaLicenseManagerGroupLicensed)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "GroupLicensed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerLicenseGrantedIterator is returned from FilterLicenseGranted and is used to iterate over the raw logs and unpacked data for LicenseGranted events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerLicenseGrantedIterator struct {
	Event *HauskaLicenseManagerLicenseGranted // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerLicenseGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerLicenseGranted)
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
		it.Event = new(HauskaLicenseManagerLicenseGranted)
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
func (it *HauskaLicenseManagerLicenseGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerLicenseGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerLicenseGranted represents a LicenseGranted event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerLicenseGranted struct {
	OrgContract common.Address
	LicenseId   *big.Int
	AssetId     *big.Int
	Licensee    common.Address
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLicenseGranted is a free log retrieval operation binding the contract event 0x73251989fba896c3df2b1d231e63769612e9592dad967dd281f353026bc5fe07.
//
// Solidity: event LicenseGranted(address indexed orgContract, uint256 indexed licenseId, uint256 indexed assetId, address licensee, uint256 fee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterLicenseGranted(opts *bind.FilterOpts, orgContract []common.Address, licenseId []*big.Int, assetId []*big.Int) (*HauskaLicenseManagerLicenseGrantedIterator, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "LicenseGranted", orgContractRule, licenseIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerLicenseGrantedIterator{contract: _HauskaLicenseManager.contract, event: "LicenseGranted", logs: logs, sub: sub}, nil
}

// WatchLicenseGranted is a free log subscription operation binding the contract event 0x73251989fba896c3df2b1d231e63769612e9592dad967dd281f353026bc5fe07.
//
// Solidity: event LicenseGranted(address indexed orgContract, uint256 indexed licenseId, uint256 indexed assetId, address licensee, uint256 fee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchLicenseGranted(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerLicenseGranted, orgContract []common.Address, licenseId []*big.Int, assetId []*big.Int) (event.Subscription, error) {

	var orgContractRule []interface{}
	for _, orgContractItem := range orgContract {
		orgContractRule = append(orgContractRule, orgContractItem)
	}
	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "LicenseGranted", orgContractRule, licenseIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerLicenseGranted)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "LicenseGranted", log); err != nil {
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

// ParseLicenseGranted is a log parse operation binding the contract event 0x73251989fba896c3df2b1d231e63769612e9592dad967dd281f353026bc5fe07.
//
// Solidity: event LicenseGranted(address indexed orgContract, uint256 indexed licenseId, uint256 indexed assetId, address licensee, uint256 fee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseLicenseGranted(log types.Log) (*HauskaLicenseManagerLicenseGranted, error) {
	event := new(HauskaLicenseManagerLicenseGranted)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "LicenseGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerLicenseRenewedIterator is returned from FilterLicenseRenewed and is used to iterate over the raw logs and unpacked data for LicenseRenewed events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerLicenseRenewedIterator struct {
	Event *HauskaLicenseManagerLicenseRenewed // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerLicenseRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerLicenseRenewed)
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
		it.Event = new(HauskaLicenseManagerLicenseRenewed)
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
func (it *HauskaLicenseManagerLicenseRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerLicenseRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerLicenseRenewed represents a LicenseRenewed event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerLicenseRenewed struct {
	LicenseId     *big.Int
	NewExpiration *big.Int
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLicenseRenewed is a free log retrieval operation binding the contract event 0xf5fc558f4b174c3024847df1ed85d463898e7e614dddc57d6d460ab376d4966e.
//
// Solidity: event LicenseRenewed(uint256 indexed licenseId, uint256 newExpiration, uint256 fee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterLicenseRenewed(opts *bind.FilterOpts, licenseId []*big.Int) (*HauskaLicenseManagerLicenseRenewedIterator, error) {

	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "LicenseRenewed", licenseIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerLicenseRenewedIterator{contract: _HauskaLicenseManager.contract, event: "LicenseRenewed", logs: logs, sub: sub}, nil
}

// WatchLicenseRenewed is a free log subscription operation binding the contract event 0xf5fc558f4b174c3024847df1ed85d463898e7e614dddc57d6d460ab376d4966e.
//
// Solidity: event LicenseRenewed(uint256 indexed licenseId, uint256 newExpiration, uint256 fee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchLicenseRenewed(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerLicenseRenewed, licenseId []*big.Int) (event.Subscription, error) {

	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "LicenseRenewed", licenseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerLicenseRenewed)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "LicenseRenewed", log); err != nil {
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

// ParseLicenseRenewed is a log parse operation binding the contract event 0xf5fc558f4b174c3024847df1ed85d463898e7e614dddc57d6d460ab376d4966e.
//
// Solidity: event LicenseRenewed(uint256 indexed licenseId, uint256 newExpiration, uint256 fee)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseLicenseRenewed(log types.Log) (*HauskaLicenseManagerLicenseRenewed, error) {
	event := new(HauskaLicenseManagerLicenseRenewed)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "LicenseRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerLicenseRevokedIterator is returned from FilterLicenseRevoked and is used to iterate over the raw logs and unpacked data for LicenseRevoked events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerLicenseRevokedIterator struct {
	Event *HauskaLicenseManagerLicenseRevoked // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerLicenseRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerLicenseRevoked)
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
		it.Event = new(HauskaLicenseManagerLicenseRevoked)
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
func (it *HauskaLicenseManagerLicenseRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerLicenseRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerLicenseRevoked represents a LicenseRevoked event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerLicenseRevoked struct {
	LicenseId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLicenseRevoked is a free log retrieval operation binding the contract event 0xf8ecc300d4504a2b49692052f138b7e697df679cf0580522720ce405bd01be55.
//
// Solidity: event LicenseRevoked(uint256 indexed licenseId)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterLicenseRevoked(opts *bind.FilterOpts, licenseId []*big.Int) (*HauskaLicenseManagerLicenseRevokedIterator, error) {

	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "LicenseRevoked", licenseIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerLicenseRevokedIterator{contract: _HauskaLicenseManager.contract, event: "LicenseRevoked", logs: logs, sub: sub}, nil
}

// WatchLicenseRevoked is a free log subscription operation binding the contract event 0xf8ecc300d4504a2b49692052f138b7e697df679cf0580522720ce405bd01be55.
//
// Solidity: event LicenseRevoked(uint256 indexed licenseId)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchLicenseRevoked(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerLicenseRevoked, licenseId []*big.Int) (event.Subscription, error) {

	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "LicenseRevoked", licenseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerLicenseRevoked)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "LicenseRevoked", log); err != nil {
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

// ParseLicenseRevoked is a log parse operation binding the contract event 0xf8ecc300d4504a2b49692052f138b7e697df679cf0580522720ce405bd01be55.
//
// Solidity: event LicenseRevoked(uint256 indexed licenseId)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseLicenseRevoked(log types.Log) (*HauskaLicenseManagerLicenseRevoked, error) {
	event := new(HauskaLicenseManagerLicenseRevoked)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "LicenseRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerRoleAdminChangedIterator struct {
	Event *HauskaLicenseManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerRoleAdminChanged)
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
		it.Event = new(HauskaLicenseManagerRoleAdminChanged)
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
func (it *HauskaLicenseManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerRoleAdminChanged represents a RoleAdminChanged event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HauskaLicenseManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerRoleAdminChangedIterator{contract: _HauskaLicenseManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerRoleAdminChanged)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseRoleAdminChanged(log types.Log) (*HauskaLicenseManagerRoleAdminChanged, error) {
	event := new(HauskaLicenseManagerRoleAdminChanged)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerRoleGrantedIterator struct {
	Event *HauskaLicenseManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerRoleGranted)
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
		it.Event = new(HauskaLicenseManagerRoleGranted)
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
func (it *HauskaLicenseManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerRoleGranted represents a RoleGranted event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaLicenseManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerRoleGrantedIterator{contract: _HauskaLicenseManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerRoleGranted)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseRoleGranted(log types.Log) (*HauskaLicenseManagerRoleGranted, error) {
	event := new(HauskaLicenseManagerRoleGranted)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerRoleRevokedIterator struct {
	Event *HauskaLicenseManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerRoleRevoked)
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
		it.Event = new(HauskaLicenseManagerRoleRevoked)
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
func (it *HauskaLicenseManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerRoleRevoked represents a RoleRevoked event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaLicenseManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerRoleRevokedIterator{contract: _HauskaLicenseManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerRoleRevoked)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseRoleRevoked(log types.Log) (*HauskaLicenseManagerRoleRevoked, error) {
	event := new(HauskaLicenseManagerRoleRevoked)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaLicenseManagerUserLocationSetIterator is returned from FilterUserLocationSet and is used to iterate over the raw logs and unpacked data for UserLocationSet events raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerUserLocationSetIterator struct {
	Event *HauskaLicenseManagerUserLocationSet // Event containing the contract specifics and raw log

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
func (it *HauskaLicenseManagerUserLocationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaLicenseManagerUserLocationSet)
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
		it.Event = new(HauskaLicenseManagerUserLocationSet)
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
func (it *HauskaLicenseManagerUserLocationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaLicenseManagerUserLocationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaLicenseManagerUserLocationSet represents a UserLocationSet event raised by the HauskaLicenseManager contract.
type HauskaLicenseManagerUserLocationSet struct {
	User     common.Address
	Location uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUserLocationSet is a free log retrieval operation binding the contract event 0x70aef35ba4f60283593fd43d2f9ffb93b830263ece43892a58ecababe3cee3d7.
//
// Solidity: event UserLocationSet(address indexed user, uint8 location)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) FilterUserLocationSet(opts *bind.FilterOpts, user []common.Address) (*HauskaLicenseManagerUserLocationSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.FilterLogs(opts, "UserLocationSet", userRule)
	if err != nil {
		return nil, err
	}
	return &HauskaLicenseManagerUserLocationSetIterator{contract: _HauskaLicenseManager.contract, event: "UserLocationSet", logs: logs, sub: sub}, nil
}

// WatchUserLocationSet is a free log subscription operation binding the contract event 0x70aef35ba4f60283593fd43d2f9ffb93b830263ece43892a58ecababe3cee3d7.
//
// Solidity: event UserLocationSet(address indexed user, uint8 location)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) WatchUserLocationSet(opts *bind.WatchOpts, sink chan<- *HauskaLicenseManagerUserLocationSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HauskaLicenseManager.contract.WatchLogs(opts, "UserLocationSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaLicenseManagerUserLocationSet)
				if err := _HauskaLicenseManager.contract.UnpackLog(event, "UserLocationSet", log); err != nil {
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

// ParseUserLocationSet is a log parse operation binding the contract event 0x70aef35ba4f60283593fd43d2f9ffb93b830263ece43892a58ecababe3cee3d7.
//
// Solidity: event UserLocationSet(address indexed user, uint8 location)
func (_HauskaLicenseManager *HauskaLicenseManagerFilterer) ParseUserLocationSet(log types.Log) (*HauskaLicenseManagerUserLocationSet, error) {
	event := new(HauskaLicenseManagerUserLocationSet)
	if err := _HauskaLicenseManager.contract.UnpackLog(event, "UserLocationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
