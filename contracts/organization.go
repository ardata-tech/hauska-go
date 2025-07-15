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

// HauskaOrgContractMetaData contains all meta data concerning the HauskaOrgContract contract.
var HauskaOrgContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_principal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_integrationPartner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AssetTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toOrg\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"AssetTransferredCrossOrg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"moduleName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"ModuleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRINCIPAL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"addCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetRetrieval\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"assetCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumIHauskaStructs.FxPool\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"grossPrice\",\"type\":\"uint256\"}],\"name\":\"calculateCreatorAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"creatorAmount\",\"type\":\"uint256\"}],\"name\":\"calculateGrossPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"assetHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEncrypted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBeLicensed\",\"type\":\"bool\"},{\"internalType\":\"enumIHauskaStructs.FxPool\",\"name\":\"fxPool\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"timeStamp\",\"type\":\"string\"},{\"internalType\":\"enumIHauskaStructs.CountryCode[]\",\"name\":\"geoRestrictions\",\"type\":\"uint8[]\"}],\"name\":\"createAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"groupAssets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"groupPrice\",\"type\":\"uint256\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getAssetGroupsByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getAssetsByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"getGroupLicensingApprovalAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getLicensingApprovalAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDCToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"groupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"integrationPartner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCreator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOrganizationMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"isValidAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"enumIHauskaStructs.LicensePermissions[]\",\"name\":\"permissions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"resellerFee\",\"type\":\"uint256\"}],\"name\":\"licenseAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"enumIHauskaStructs.LicensePermissions[]\",\"name\":\"permissions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"resellerFee\",\"type\":\"uint256\"}],\"name\":\"licenseGroup\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"licenseManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"principal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"removeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"removeGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenueDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetRegistry\",\"type\":\"address\"}],\"name\":\"setAssetRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_groupManager\",\"type\":\"address\"}],\"name\":\"setGroupManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_licenseManager\",\"type\":\"address\"}],\"name\":\"setLicenseManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_revenueDistributor\",\"type\":\"address\"}],\"name\":\"setRevenueDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toOrg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAssetCrossOrg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"verifyAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HauskaOrgContractABI is the input ABI used to generate the binding from.
// Deprecated: Use HauskaOrgContractMetaData.ABI instead.
var HauskaOrgContractABI = HauskaOrgContractMetaData.ABI

// HauskaOrgContract is an auto generated Go binding around an Ethereum contract.
type HauskaOrgContract struct {
	HauskaOrgContractCaller     // Read-only binding to the contract
	HauskaOrgContractTransactor // Write-only binding to the contract
	HauskaOrgContractFilterer   // Log filterer for contract events
}

// HauskaOrgContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type HauskaOrgContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaOrgContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HauskaOrgContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaOrgContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HauskaOrgContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HauskaOrgContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HauskaOrgContractSession struct {
	Contract     *HauskaOrgContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HauskaOrgContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HauskaOrgContractCallerSession struct {
	Contract *HauskaOrgContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HauskaOrgContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HauskaOrgContractTransactorSession struct {
	Contract     *HauskaOrgContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HauskaOrgContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type HauskaOrgContractRaw struct {
	Contract *HauskaOrgContract // Generic contract binding to access the raw methods on
}

// HauskaOrgContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HauskaOrgContractCallerRaw struct {
	Contract *HauskaOrgContractCaller // Generic read-only contract binding to access the raw methods on
}

// HauskaOrgContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HauskaOrgContractTransactorRaw struct {
	Contract *HauskaOrgContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHauskaOrgContract creates a new instance of HauskaOrgContract, bound to a specific deployed contract.
func NewHauskaOrgContract(address common.Address, backend bind.ContractBackend) (*HauskaOrgContract, error) {
	contract, err := bindHauskaOrgContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContract{HauskaOrgContractCaller: HauskaOrgContractCaller{contract: contract}, HauskaOrgContractTransactor: HauskaOrgContractTransactor{contract: contract}, HauskaOrgContractFilterer: HauskaOrgContractFilterer{contract: contract}}, nil
}

// NewHauskaOrgContractCaller creates a new read-only instance of HauskaOrgContract, bound to a specific deployed contract.
func NewHauskaOrgContractCaller(address common.Address, caller bind.ContractCaller) (*HauskaOrgContractCaller, error) {
	contract, err := bindHauskaOrgContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractCaller{contract: contract}, nil
}

// NewHauskaOrgContractTransactor creates a new write-only instance of HauskaOrgContract, bound to a specific deployed contract.
func NewHauskaOrgContractTransactor(address common.Address, transactor bind.ContractTransactor) (*HauskaOrgContractTransactor, error) {
	contract, err := bindHauskaOrgContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractTransactor{contract: contract}, nil
}

// NewHauskaOrgContractFilterer creates a new log filterer instance of HauskaOrgContract, bound to a specific deployed contract.
func NewHauskaOrgContractFilterer(address common.Address, filterer bind.ContractFilterer) (*HauskaOrgContractFilterer, error) {
	contract, err := bindHauskaOrgContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractFilterer{contract: contract}, nil
}

// bindHauskaOrgContract binds a generic wrapper to an already deployed contract.
func bindHauskaOrgContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HauskaOrgContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaOrgContract *HauskaOrgContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaOrgContract.Contract.HauskaOrgContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaOrgContract *HauskaOrgContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.HauskaOrgContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaOrgContract *HauskaOrgContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.HauskaOrgContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HauskaOrgContract *HauskaOrgContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HauskaOrgContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HauskaOrgContract *HauskaOrgContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HauskaOrgContract *HauskaOrgContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.contract.Transact(opts, method, params...)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCaller) CREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractSession) CREATORROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.CREATORROLE(&_HauskaOrgContract.CallOpts)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) CREATORROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.CREATORROLE(&_HauskaOrgContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.DEFAULTADMINROLE(&_HauskaOrgContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.DEFAULTADMINROLE(&_HauskaOrgContract.CallOpts)
}

// PRINCIPALROLE is a free data retrieval call binding the contract method 0x25c3385f.
//
// Solidity: function PRINCIPAL_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCaller) PRINCIPALROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "PRINCIPAL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRINCIPALROLE is a free data retrieval call binding the contract method 0x25c3385f.
//
// Solidity: function PRINCIPAL_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractSession) PRINCIPALROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.PRINCIPALROLE(&_HauskaOrgContract.CallOpts)
}

// PRINCIPALROLE is a free data retrieval call binding the contract method 0x25c3385f.
//
// Solidity: function PRINCIPAL_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) PRINCIPALROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.PRINCIPALROLE(&_HauskaOrgContract.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCaller) VERIFIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "VERIFIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractSession) VERIFIERROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.VERIFIERROLE(&_HauskaOrgContract.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) VERIFIERROLE() ([32]byte, error) {
	return _HauskaOrgContract.Contract.VERIFIERROLE(&_HauskaOrgContract.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) AssetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "assetRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) AssetRegistry() (common.Address, error) {
	return _HauskaOrgContract.Contract.AssetRegistry(&_HauskaOrgContract.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) AssetRegistry() (common.Address, error) {
	return _HauskaOrgContract.Contract.AssetRegistry(&_HauskaOrgContract.CallOpts)
}

// AssetRetrieval is a free data retrieval call binding the contract method 0x70b1d2e4.
//
// Solidity: function assetRetrieval(uint256 assetId) view returns(string assetCID, string metadataCID)
func (_HauskaOrgContract *HauskaOrgContractCaller) AssetRetrieval(opts *bind.CallOpts, assetId *big.Int) (struct {
	AssetCID    string
	MetadataCID string
}, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "assetRetrieval", assetId)

	outstruct := new(struct {
		AssetCID    string
		MetadataCID string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetCID = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.MetadataCID = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// AssetRetrieval is a free data retrieval call binding the contract method 0x70b1d2e4.
//
// Solidity: function assetRetrieval(uint256 assetId) view returns(string assetCID, string metadataCID)
func (_HauskaOrgContract *HauskaOrgContractSession) AssetRetrieval(assetId *big.Int) (struct {
	AssetCID    string
	MetadataCID string
}, error) {
	return _HauskaOrgContract.Contract.AssetRetrieval(&_HauskaOrgContract.CallOpts, assetId)
}

// AssetRetrieval is a free data retrieval call binding the contract method 0x70b1d2e4.
//
// Solidity: function assetRetrieval(uint256 assetId) view returns(string assetCID, string metadataCID)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) AssetRetrieval(assetId *big.Int) (struct {
	AssetCID    string
	MetadataCID string
}, error) {
	return _HauskaOrgContract.Contract.AssetRetrieval(&_HauskaOrgContract.CallOpts, assetId)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 assetId) view returns(uint256, address, address, address, string, string, bytes32, uint256, bool, uint256, uint256, uint256, bool, bool, uint8, string)
func (_HauskaOrgContract *HauskaOrgContractCaller) Assets(opts *bind.CallOpts, assetId *big.Int) (*big.Int, common.Address, common.Address, common.Address, string, string, [32]byte, *big.Int, bool, *big.Int, *big.Int, *big.Int, bool, bool, uint8, string, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "assets", assetId)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(common.Address), *new(common.Address), *new(string), *new(string), *new([32]byte), *new(*big.Int), *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), *new(bool), *new(uint8), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)
	out6 := *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(bool)).(*bool)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	out11 := *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	out12 := *abi.ConvertType(out[12], new(bool)).(*bool)
	out13 := *abi.ConvertType(out[13], new(bool)).(*bool)
	out14 := *abi.ConvertType(out[14], new(uint8)).(*uint8)
	out15 := *abi.ConvertType(out[15], new(string)).(*string)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, out12, out13, out14, out15, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 assetId) view returns(uint256, address, address, address, string, string, bytes32, uint256, bool, uint256, uint256, uint256, bool, bool, uint8, string)
func (_HauskaOrgContract *HauskaOrgContractSession) Assets(assetId *big.Int) (*big.Int, common.Address, common.Address, common.Address, string, string, [32]byte, *big.Int, bool, *big.Int, *big.Int, *big.Int, bool, bool, uint8, string, error) {
	return _HauskaOrgContract.Contract.Assets(&_HauskaOrgContract.CallOpts, assetId)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 assetId) view returns(uint256, address, address, address, string, string, bytes32, uint256, bool, uint256, uint256, uint256, bool, bool, uint8, string)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) Assets(assetId *big.Int) (*big.Int, common.Address, common.Address, common.Address, string, string, [32]byte, *big.Int, bool, *big.Int, *big.Int, *big.Int, bool, bool, uint8, string, error) {
	return _HauskaOrgContract.Contract.Assets(&_HauskaOrgContract.CallOpts, assetId)
}

// CalculateCreatorAmount is a free data retrieval call binding the contract method 0xc4f3691c.
//
// Solidity: function calculateCreatorAmount(uint256 grossPrice) view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCaller) CalculateCreatorAmount(opts *bind.CallOpts, grossPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "calculateCreatorAmount", grossPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCreatorAmount is a free data retrieval call binding the contract method 0xc4f3691c.
//
// Solidity: function calculateCreatorAmount(uint256 grossPrice) view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) CalculateCreatorAmount(grossPrice *big.Int) (*big.Int, error) {
	return _HauskaOrgContract.Contract.CalculateCreatorAmount(&_HauskaOrgContract.CallOpts, grossPrice)
}

// CalculateCreatorAmount is a free data retrieval call binding the contract method 0xc4f3691c.
//
// Solidity: function calculateCreatorAmount(uint256 grossPrice) view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) CalculateCreatorAmount(grossPrice *big.Int) (*big.Int, error) {
	return _HauskaOrgContract.Contract.CalculateCreatorAmount(&_HauskaOrgContract.CallOpts, grossPrice)
}

// CalculateGrossPrice is a free data retrieval call binding the contract method 0x0a4f18be.
//
// Solidity: function calculateGrossPrice(uint256 creatorAmount) view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCaller) CalculateGrossPrice(opts *bind.CallOpts, creatorAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "calculateGrossPrice", creatorAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateGrossPrice is a free data retrieval call binding the contract method 0x0a4f18be.
//
// Solidity: function calculateGrossPrice(uint256 creatorAmount) view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) CalculateGrossPrice(creatorAmount *big.Int) (*big.Int, error) {
	return _HauskaOrgContract.Contract.CalculateGrossPrice(&_HauskaOrgContract.CallOpts, creatorAmount)
}

// CalculateGrossPrice is a free data retrieval call binding the contract method 0x0a4f18be.
//
// Solidity: function calculateGrossPrice(uint256 creatorAmount) view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) CalculateGrossPrice(creatorAmount *big.Int) (*big.Int, error) {
	return _HauskaOrgContract.Contract.CalculateGrossPrice(&_HauskaOrgContract.CallOpts, creatorAmount)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) Creators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "creators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _HauskaOrgContract.Contract.Creators(&_HauskaOrgContract.CallOpts, arg0)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _HauskaOrgContract.Contract.Creators(&_HauskaOrgContract.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) Factory() (common.Address, error) {
	return _HauskaOrgContract.Contract.Factory(&_HauskaOrgContract.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) Factory() (common.Address, error) {
	return _HauskaOrgContract.Contract.Factory(&_HauskaOrgContract.CallOpts)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCaller) GetAssetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getAssetCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) GetAssetCount() (*big.Int, error) {
	return _HauskaOrgContract.Contract.GetAssetCount(&_HauskaOrgContract.CallOpts)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetAssetCount() (*big.Int, error) {
	return _HauskaOrgContract.Contract.GetAssetCount(&_HauskaOrgContract.CallOpts)
}

// GetAssetGroupsByOwner is a free data retrieval call binding the contract method 0x8ee1e995.
//
// Solidity: function getAssetGroupsByOwner(address owner) view returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractCaller) GetAssetGroupsByOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getAssetGroupsByOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetGroupsByOwner is a free data retrieval call binding the contract method 0x8ee1e995.
//
// Solidity: function getAssetGroupsByOwner(address owner) view returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractSession) GetAssetGroupsByOwner(owner common.Address) ([]*big.Int, error) {
	return _HauskaOrgContract.Contract.GetAssetGroupsByOwner(&_HauskaOrgContract.CallOpts, owner)
}

// GetAssetGroupsByOwner is a free data retrieval call binding the contract method 0x8ee1e995.
//
// Solidity: function getAssetGroupsByOwner(address owner) view returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetAssetGroupsByOwner(owner common.Address) ([]*big.Int, error) {
	return _HauskaOrgContract.Contract.GetAssetGroupsByOwner(&_HauskaOrgContract.CallOpts, owner)
}

// GetAssetsByOwner is a free data retrieval call binding the contract method 0x276f0934.
//
// Solidity: function getAssetsByOwner(address owner) view returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractCaller) GetAssetsByOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getAssetsByOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetsByOwner is a free data retrieval call binding the contract method 0x276f0934.
//
// Solidity: function getAssetsByOwner(address owner) view returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractSession) GetAssetsByOwner(owner common.Address) ([]*big.Int, error) {
	return _HauskaOrgContract.Contract.GetAssetsByOwner(&_HauskaOrgContract.CallOpts, owner)
}

// GetAssetsByOwner is a free data retrieval call binding the contract method 0x276f0934.
//
// Solidity: function getAssetsByOwner(address owner) view returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetAssetsByOwner(owner common.Address) ([]*big.Int, error) {
	return _HauskaOrgContract.Contract.GetAssetsByOwner(&_HauskaOrgContract.CallOpts, owner)
}

// GetCreatorCount is a free data retrieval call binding the contract method 0xdfe6712e.
//
// Solidity: function getCreatorCount() view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCaller) GetCreatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getCreatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreatorCount is a free data retrieval call binding the contract method 0xdfe6712e.
//
// Solidity: function getCreatorCount() view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) GetCreatorCount() (*big.Int, error) {
	return _HauskaOrgContract.Contract.GetCreatorCount(&_HauskaOrgContract.CallOpts)
}

// GetCreatorCount is a free data retrieval call binding the contract method 0xdfe6712e.
//
// Solidity: function getCreatorCount() view returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetCreatorCount() (*big.Int, error) {
	return _HauskaOrgContract.Contract.GetCreatorCount(&_HauskaOrgContract.CallOpts)
}

// GetGroupLicensingApprovalAmount is a free data retrieval call binding the contract method 0x3e645893.
//
// Solidity: function getGroupLicensingApprovalAmount(uint256 groupId) view returns(address spender, uint256 amount)
func (_HauskaOrgContract *HauskaOrgContractCaller) GetGroupLicensingApprovalAmount(opts *bind.CallOpts, groupId *big.Int) (struct {
	Spender common.Address
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getGroupLicensingApprovalAmount", groupId)

	outstruct := new(struct {
		Spender common.Address
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Spender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGroupLicensingApprovalAmount is a free data retrieval call binding the contract method 0x3e645893.
//
// Solidity: function getGroupLicensingApprovalAmount(uint256 groupId) view returns(address spender, uint256 amount)
func (_HauskaOrgContract *HauskaOrgContractSession) GetGroupLicensingApprovalAmount(groupId *big.Int) (struct {
	Spender common.Address
	Amount  *big.Int
}, error) {
	return _HauskaOrgContract.Contract.GetGroupLicensingApprovalAmount(&_HauskaOrgContract.CallOpts, groupId)
}

// GetGroupLicensingApprovalAmount is a free data retrieval call binding the contract method 0x3e645893.
//
// Solidity: function getGroupLicensingApprovalAmount(uint256 groupId) view returns(address spender, uint256 amount)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetGroupLicensingApprovalAmount(groupId *big.Int) (struct {
	Spender common.Address
	Amount  *big.Int
}, error) {
	return _HauskaOrgContract.Contract.GetGroupLicensingApprovalAmount(&_HauskaOrgContract.CallOpts, groupId)
}

// GetLicensingApprovalAmount is a free data retrieval call binding the contract method 0xfeda91e2.
//
// Solidity: function getLicensingApprovalAmount(uint256 assetId) view returns(address spender, uint256 amount)
func (_HauskaOrgContract *HauskaOrgContractCaller) GetLicensingApprovalAmount(opts *bind.CallOpts, assetId *big.Int) (struct {
	Spender common.Address
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getLicensingApprovalAmount", assetId)

	outstruct := new(struct {
		Spender common.Address
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Spender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLicensingApprovalAmount is a free data retrieval call binding the contract method 0xfeda91e2.
//
// Solidity: function getLicensingApprovalAmount(uint256 assetId) view returns(address spender, uint256 amount)
func (_HauskaOrgContract *HauskaOrgContractSession) GetLicensingApprovalAmount(assetId *big.Int) (struct {
	Spender common.Address
	Amount  *big.Int
}, error) {
	return _HauskaOrgContract.Contract.GetLicensingApprovalAmount(&_HauskaOrgContract.CallOpts, assetId)
}

// GetLicensingApprovalAmount is a free data retrieval call binding the contract method 0xfeda91e2.
//
// Solidity: function getLicensingApprovalAmount(uint256 assetId) view returns(address spender, uint256 amount)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetLicensingApprovalAmount(assetId *big.Int) (struct {
	Spender common.Address
	Amount  *big.Int
}, error) {
	return _HauskaOrgContract.Contract.GetLicensingApprovalAmount(&_HauskaOrgContract.CallOpts, assetId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaOrgContract.Contract.GetRoleAdmin(&_HauskaOrgContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HauskaOrgContract.Contract.GetRoleAdmin(&_HauskaOrgContract.CallOpts, role)
}

// GetUSDCToken is a free data retrieval call binding the contract method 0xf1752409.
//
// Solidity: function getUSDCToken() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) GetUSDCToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "getUSDCToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDCToken is a free data retrieval call binding the contract method 0xf1752409.
//
// Solidity: function getUSDCToken() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) GetUSDCToken() (common.Address, error) {
	return _HauskaOrgContract.Contract.GetUSDCToken(&_HauskaOrgContract.CallOpts)
}

// GetUSDCToken is a free data retrieval call binding the contract method 0xf1752409.
//
// Solidity: function getUSDCToken() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GetUSDCToken() (common.Address, error) {
	return _HauskaOrgContract.Contract.GetUSDCToken(&_HauskaOrgContract.CallOpts)
}

// GroupManager is a free data retrieval call binding the contract method 0xb26bc483.
//
// Solidity: function groupManager() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) GroupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "groupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GroupManager is a free data retrieval call binding the contract method 0xb26bc483.
//
// Solidity: function groupManager() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) GroupManager() (common.Address, error) {
	return _HauskaOrgContract.Contract.GroupManager(&_HauskaOrgContract.CallOpts)
}

// GroupManager is a free data retrieval call binding the contract method 0xb26bc483.
//
// Solidity: function groupManager() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) GroupManager() (common.Address, error) {
	return _HauskaOrgContract.Contract.GroupManager(&_HauskaOrgContract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaOrgContract.Contract.HasRole(&_HauskaOrgContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HauskaOrgContract.Contract.HasRole(&_HauskaOrgContract.CallOpts, role, account)
}

// IntegrationPartner is a free data retrieval call binding the contract method 0xd8eda83e.
//
// Solidity: function integrationPartner() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) IntegrationPartner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "integrationPartner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IntegrationPartner is a free data retrieval call binding the contract method 0xd8eda83e.
//
// Solidity: function integrationPartner() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) IntegrationPartner() (common.Address, error) {
	return _HauskaOrgContract.Contract.IntegrationPartner(&_HauskaOrgContract.CallOpts)
}

// IntegrationPartner is a free data retrieval call binding the contract method 0xd8eda83e.
//
// Solidity: function integrationPartner() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) IntegrationPartner() (common.Address, error) {
	return _HauskaOrgContract.Contract.IntegrationPartner(&_HauskaOrgContract.CallOpts)
}

// IsCreator is a free data retrieval call binding the contract method 0xefd46065.
//
// Solidity: function isCreator(address ) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCaller) IsCreator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "isCreator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCreator is a free data retrieval call binding the contract method 0xefd46065.
//
// Solidity: function isCreator(address ) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractSession) IsCreator(arg0 common.Address) (bool, error) {
	return _HauskaOrgContract.Contract.IsCreator(&_HauskaOrgContract.CallOpts, arg0)
}

// IsCreator is a free data retrieval call binding the contract method 0xefd46065.
//
// Solidity: function isCreator(address ) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) IsCreator(arg0 common.Address) (bool, error) {
	return _HauskaOrgContract.Contract.IsCreator(&_HauskaOrgContract.CallOpts, arg0)
}

// IsOrganizationMember is a free data retrieval call binding the contract method 0xa1bf7134.
//
// Solidity: function isOrganizationMember(address account) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCaller) IsOrganizationMember(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "isOrganizationMember", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrganizationMember is a free data retrieval call binding the contract method 0xa1bf7134.
//
// Solidity: function isOrganizationMember(address account) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractSession) IsOrganizationMember(account common.Address) (bool, error) {
	return _HauskaOrgContract.Contract.IsOrganizationMember(&_HauskaOrgContract.CallOpts, account)
}

// IsOrganizationMember is a free data retrieval call binding the contract method 0xa1bf7134.
//
// Solidity: function isOrganizationMember(address account) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) IsOrganizationMember(account common.Address) (bool, error) {
	return _HauskaOrgContract.Contract.IsOrganizationMember(&_HauskaOrgContract.CallOpts, account)
}

// IsValidAsset is a free data retrieval call binding the contract method 0xdb0aeda2.
//
// Solidity: function isValidAsset(uint256 assetId) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCaller) IsValidAsset(opts *bind.CallOpts, assetId *big.Int) (bool, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "isValidAsset", assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAsset is a free data retrieval call binding the contract method 0xdb0aeda2.
//
// Solidity: function isValidAsset(uint256 assetId) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractSession) IsValidAsset(assetId *big.Int) (bool, error) {
	return _HauskaOrgContract.Contract.IsValidAsset(&_HauskaOrgContract.CallOpts, assetId)
}

// IsValidAsset is a free data retrieval call binding the contract method 0xdb0aeda2.
//
// Solidity: function isValidAsset(uint256 assetId) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) IsValidAsset(assetId *big.Int) (bool, error) {
	return _HauskaOrgContract.Contract.IsValidAsset(&_HauskaOrgContract.CallOpts, assetId)
}

// LicenseManager is a free data retrieval call binding the contract method 0x0ad51144.
//
// Solidity: function licenseManager() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) LicenseManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "licenseManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseManager is a free data retrieval call binding the contract method 0x0ad51144.
//
// Solidity: function licenseManager() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) LicenseManager() (common.Address, error) {
	return _HauskaOrgContract.Contract.LicenseManager(&_HauskaOrgContract.CallOpts)
}

// LicenseManager is a free data retrieval call binding the contract method 0x0ad51144.
//
// Solidity: function licenseManager() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) LicenseManager() (common.Address, error) {
	return _HauskaOrgContract.Contract.LicenseManager(&_HauskaOrgContract.CallOpts)
}

// Principal is a free data retrieval call binding the contract method 0xba5d3078.
//
// Solidity: function principal() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) Principal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "principal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Principal is a free data retrieval call binding the contract method 0xba5d3078.
//
// Solidity: function principal() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) Principal() (common.Address, error) {
	return _HauskaOrgContract.Contract.Principal(&_HauskaOrgContract.CallOpts)
}

// Principal is a free data retrieval call binding the contract method 0xba5d3078.
//
// Solidity: function principal() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) Principal() (common.Address, error) {
	return _HauskaOrgContract.Contract.Principal(&_HauskaOrgContract.CallOpts)
}

// RevenueDistributor is a free data retrieval call binding the contract method 0xeae0a488.
//
// Solidity: function revenueDistributor() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCaller) RevenueDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "revenueDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenueDistributor is a free data retrieval call binding the contract method 0xeae0a488.
//
// Solidity: function revenueDistributor() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractSession) RevenueDistributor() (common.Address, error) {
	return _HauskaOrgContract.Contract.RevenueDistributor(&_HauskaOrgContract.CallOpts)
}

// RevenueDistributor is a free data retrieval call binding the contract method 0xeae0a488.
//
// Solidity: function revenueDistributor() view returns(address)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) RevenueDistributor() (common.Address, error) {
	return _HauskaOrgContract.Contract.RevenueDistributor(&_HauskaOrgContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HauskaOrgContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaOrgContract.Contract.SupportsInterface(&_HauskaOrgContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HauskaOrgContract *HauskaOrgContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HauskaOrgContract.Contract.SupportsInterface(&_HauskaOrgContract.CallOpts, interfaceId)
}

// AddCreator is a paid mutator transaction binding the contract method 0x3b9dce05.
//
// Solidity: function addCreator(address creator) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) AddCreator(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "addCreator", creator)
}

// AddCreator is a paid mutator transaction binding the contract method 0x3b9dce05.
//
// Solidity: function addCreator(address creator) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) AddCreator(creator common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.AddCreator(&_HauskaOrgContract.TransactOpts, creator)
}

// AddCreator is a paid mutator transaction binding the contract method 0x3b9dce05.
//
// Solidity: function addCreator(address creator) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) AddCreator(creator common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.AddCreator(&_HauskaOrgContract.TransactOpts, creator)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x0b649745.
//
// Solidity: function createAsset(string assetCID, string metadataCID, bytes32 assetHash, uint256 price, bool isEncrypted, bool canBeLicensed, uint8 fxPool, string timeStamp, uint8[] geoRestrictions) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactor) CreateAsset(opts *bind.TransactOpts, assetCID string, metadataCID string, assetHash [32]byte, price *big.Int, isEncrypted bool, canBeLicensed bool, fxPool uint8, timeStamp string, geoRestrictions []uint8) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "createAsset", assetCID, metadataCID, assetHash, price, isEncrypted, canBeLicensed, fxPool, timeStamp, geoRestrictions)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x0b649745.
//
// Solidity: function createAsset(string assetCID, string metadataCID, bytes32 assetHash, uint256 price, bool isEncrypted, bool canBeLicensed, uint8 fxPool, string timeStamp, uint8[] geoRestrictions) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) CreateAsset(assetCID string, metadataCID string, assetHash [32]byte, price *big.Int, isEncrypted bool, canBeLicensed bool, fxPool uint8, timeStamp string, geoRestrictions []uint8) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.CreateAsset(&_HauskaOrgContract.TransactOpts, assetCID, metadataCID, assetHash, price, isEncrypted, canBeLicensed, fxPool, timeStamp, geoRestrictions)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x0b649745.
//
// Solidity: function createAsset(string assetCID, string metadataCID, bytes32 assetHash, uint256 price, bool isEncrypted, bool canBeLicensed, uint8 fxPool, string timeStamp, uint8[] geoRestrictions) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) CreateAsset(assetCID string, metadataCID string, assetHash [32]byte, price *big.Int, isEncrypted bool, canBeLicensed bool, fxPool uint8, timeStamp string, geoRestrictions []uint8) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.CreateAsset(&_HauskaOrgContract.TransactOpts, assetCID, metadataCID, assetHash, price, isEncrypted, canBeLicensed, fxPool, timeStamp, geoRestrictions)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x9db614b9.
//
// Solidity: function createGroup(string groupName, uint256[] groupAssets, uint256 groupPrice) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactor) CreateGroup(opts *bind.TransactOpts, groupName string, groupAssets []*big.Int, groupPrice *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "createGroup", groupName, groupAssets, groupPrice)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x9db614b9.
//
// Solidity: function createGroup(string groupName, uint256[] groupAssets, uint256 groupPrice) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) CreateGroup(groupName string, groupAssets []*big.Int, groupPrice *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.CreateGroup(&_HauskaOrgContract.TransactOpts, groupName, groupAssets, groupPrice)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x9db614b9.
//
// Solidity: function createGroup(string groupName, uint256[] groupAssets, uint256 groupPrice) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) CreateGroup(groupName string, groupAssets []*big.Int, groupPrice *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.CreateGroup(&_HauskaOrgContract.TransactOpts, groupName, groupAssets, groupPrice)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.GrantRole(&_HauskaOrgContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.GrantRole(&_HauskaOrgContract.TransactOpts, role, account)
}

// LicenseAsset is a paid mutator transaction binding the contract method 0x0a0d673b.
//
// Solidity: function licenseAsset(uint256 assetId, uint8[] permissions, uint256 resellerFee) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactor) LicenseAsset(opts *bind.TransactOpts, assetId *big.Int, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "licenseAsset", assetId, permissions, resellerFee)
}

// LicenseAsset is a paid mutator transaction binding the contract method 0x0a0d673b.
//
// Solidity: function licenseAsset(uint256 assetId, uint8[] permissions, uint256 resellerFee) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) LicenseAsset(assetId *big.Int, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.LicenseAsset(&_HauskaOrgContract.TransactOpts, assetId, permissions, resellerFee)
}

// LicenseAsset is a paid mutator transaction binding the contract method 0x0a0d673b.
//
// Solidity: function licenseAsset(uint256 assetId, uint8[] permissions, uint256 resellerFee) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) LicenseAsset(assetId *big.Int, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.LicenseAsset(&_HauskaOrgContract.TransactOpts, assetId, permissions, resellerFee)
}

// LicenseGroup is a paid mutator transaction binding the contract method 0x939ba772.
//
// Solidity: function licenseGroup(uint256 groupId, uint8[] permissions, uint256 resellerFee) returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractTransactor) LicenseGroup(opts *bind.TransactOpts, groupId *big.Int, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "licenseGroup", groupId, permissions, resellerFee)
}

// LicenseGroup is a paid mutator transaction binding the contract method 0x939ba772.
//
// Solidity: function licenseGroup(uint256 groupId, uint8[] permissions, uint256 resellerFee) returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractSession) LicenseGroup(groupId *big.Int, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.LicenseGroup(&_HauskaOrgContract.TransactOpts, groupId, permissions, resellerFee)
}

// LicenseGroup is a paid mutator transaction binding the contract method 0x939ba772.
//
// Solidity: function licenseGroup(uint256 groupId, uint8[] permissions, uint256 resellerFee) returns(uint256[])
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) LicenseGroup(groupId *big.Int, permissions []uint8, resellerFee *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.LicenseGroup(&_HauskaOrgContract.TransactOpts, groupId, permissions, resellerFee)
}

// RemoveCreator is a paid mutator transaction binding the contract method 0xbd6018bb.
//
// Solidity: function removeCreator(address creator) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) RemoveCreator(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "removeCreator", creator)
}

// RemoveCreator is a paid mutator transaction binding the contract method 0xbd6018bb.
//
// Solidity: function removeCreator(address creator) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) RemoveCreator(creator common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RemoveCreator(&_HauskaOrgContract.TransactOpts, creator)
}

// RemoveCreator is a paid mutator transaction binding the contract method 0xbd6018bb.
//
// Solidity: function removeCreator(address creator) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) RemoveCreator(creator common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RemoveCreator(&_HauskaOrgContract.TransactOpts, creator)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x5fe1410d.
//
// Solidity: function removeGroup(uint256 groupId) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) RemoveGroup(opts *bind.TransactOpts, groupId *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "removeGroup", groupId)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x5fe1410d.
//
// Solidity: function removeGroup(uint256 groupId) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) RemoveGroup(groupId *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RemoveGroup(&_HauskaOrgContract.TransactOpts, groupId)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x5fe1410d.
//
// Solidity: function removeGroup(uint256 groupId) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) RemoveGroup(groupId *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RemoveGroup(&_HauskaOrgContract.TransactOpts, groupId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RenounceRole(&_HauskaOrgContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RenounceRole(&_HauskaOrgContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RevokeRole(&_HauskaOrgContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.RevokeRole(&_HauskaOrgContract.TransactOpts, role, account)
}

// SetAssetRegistry is a paid mutator transaction binding the contract method 0xe84d033c.
//
// Solidity: function setAssetRegistry(address _assetRegistry) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) SetAssetRegistry(opts *bind.TransactOpts, _assetRegistry common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "setAssetRegistry", _assetRegistry)
}

// SetAssetRegistry is a paid mutator transaction binding the contract method 0xe84d033c.
//
// Solidity: function setAssetRegistry(address _assetRegistry) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) SetAssetRegistry(_assetRegistry common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetAssetRegistry(&_HauskaOrgContract.TransactOpts, _assetRegistry)
}

// SetAssetRegistry is a paid mutator transaction binding the contract method 0xe84d033c.
//
// Solidity: function setAssetRegistry(address _assetRegistry) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) SetAssetRegistry(_assetRegistry common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetAssetRegistry(&_HauskaOrgContract.TransactOpts, _assetRegistry)
}

// SetGroupManager is a paid mutator transaction binding the contract method 0xb328b69d.
//
// Solidity: function setGroupManager(address _groupManager) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) SetGroupManager(opts *bind.TransactOpts, _groupManager common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "setGroupManager", _groupManager)
}

// SetGroupManager is a paid mutator transaction binding the contract method 0xb328b69d.
//
// Solidity: function setGroupManager(address _groupManager) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) SetGroupManager(_groupManager common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetGroupManager(&_HauskaOrgContract.TransactOpts, _groupManager)
}

// SetGroupManager is a paid mutator transaction binding the contract method 0xb328b69d.
//
// Solidity: function setGroupManager(address _groupManager) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) SetGroupManager(_groupManager common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetGroupManager(&_HauskaOrgContract.TransactOpts, _groupManager)
}

// SetLicenseManager is a paid mutator transaction binding the contract method 0x66559fbb.
//
// Solidity: function setLicenseManager(address _licenseManager) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) SetLicenseManager(opts *bind.TransactOpts, _licenseManager common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "setLicenseManager", _licenseManager)
}

// SetLicenseManager is a paid mutator transaction binding the contract method 0x66559fbb.
//
// Solidity: function setLicenseManager(address _licenseManager) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) SetLicenseManager(_licenseManager common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetLicenseManager(&_HauskaOrgContract.TransactOpts, _licenseManager)
}

// SetLicenseManager is a paid mutator transaction binding the contract method 0x66559fbb.
//
// Solidity: function setLicenseManager(address _licenseManager) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) SetLicenseManager(_licenseManager common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetLicenseManager(&_HauskaOrgContract.TransactOpts, _licenseManager)
}

// SetRevenueDistributor is a paid mutator transaction binding the contract method 0x179f9e65.
//
// Solidity: function setRevenueDistributor(address _revenueDistributor) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) SetRevenueDistributor(opts *bind.TransactOpts, _revenueDistributor common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "setRevenueDistributor", _revenueDistributor)
}

// SetRevenueDistributor is a paid mutator transaction binding the contract method 0x179f9e65.
//
// Solidity: function setRevenueDistributor(address _revenueDistributor) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) SetRevenueDistributor(_revenueDistributor common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetRevenueDistributor(&_HauskaOrgContract.TransactOpts, _revenueDistributor)
}

// SetRevenueDistributor is a paid mutator transaction binding the contract method 0x179f9e65.
//
// Solidity: function setRevenueDistributor(address _revenueDistributor) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) SetRevenueDistributor(_revenueDistributor common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.SetRevenueDistributor(&_HauskaOrgContract.TransactOpts, _revenueDistributor)
}

// TransferAsset is a paid mutator transaction binding the contract method 0xfa62ee8d.
//
// Solidity: function transferAsset(uint256 assetId, address newOwner) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) TransferAsset(opts *bind.TransactOpts, assetId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "transferAsset", assetId, newOwner)
}

// TransferAsset is a paid mutator transaction binding the contract method 0xfa62ee8d.
//
// Solidity: function transferAsset(uint256 assetId, address newOwner) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) TransferAsset(assetId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.TransferAsset(&_HauskaOrgContract.TransactOpts, assetId, newOwner)
}

// TransferAsset is a paid mutator transaction binding the contract method 0xfa62ee8d.
//
// Solidity: function transferAsset(uint256 assetId, address newOwner) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) TransferAsset(assetId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.TransferAsset(&_HauskaOrgContract.TransactOpts, assetId, newOwner)
}

// TransferAssetCrossOrg is a paid mutator transaction binding the contract method 0x244c8f2b.
//
// Solidity: function transferAssetCrossOrg(uint256 assetId, address toOrg, address newOwner) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactor) TransferAssetCrossOrg(opts *bind.TransactOpts, assetId *big.Int, toOrg common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "transferAssetCrossOrg", assetId, toOrg, newOwner)
}

// TransferAssetCrossOrg is a paid mutator transaction binding the contract method 0x244c8f2b.
//
// Solidity: function transferAssetCrossOrg(uint256 assetId, address toOrg, address newOwner) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractSession) TransferAssetCrossOrg(assetId *big.Int, toOrg common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.TransferAssetCrossOrg(&_HauskaOrgContract.TransactOpts, assetId, toOrg, newOwner)
}

// TransferAssetCrossOrg is a paid mutator transaction binding the contract method 0x244c8f2b.
//
// Solidity: function transferAssetCrossOrg(uint256 assetId, address toOrg, address newOwner) returns(uint256)
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) TransferAssetCrossOrg(assetId *big.Int, toOrg common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.TransferAssetCrossOrg(&_HauskaOrgContract.TransactOpts, assetId, toOrg, newOwner)
}

// VerifyAsset is a paid mutator transaction binding the contract method 0x59fd285c.
//
// Solidity: function verifyAsset(uint256 assetId) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactor) VerifyAsset(opts *bind.TransactOpts, assetId *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.contract.Transact(opts, "verifyAsset", assetId)
}

// VerifyAsset is a paid mutator transaction binding the contract method 0x59fd285c.
//
// Solidity: function verifyAsset(uint256 assetId) returns()
func (_HauskaOrgContract *HauskaOrgContractSession) VerifyAsset(assetId *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.VerifyAsset(&_HauskaOrgContract.TransactOpts, assetId)
}

// VerifyAsset is a paid mutator transaction binding the contract method 0x59fd285c.
//
// Solidity: function verifyAsset(uint256 assetId) returns()
func (_HauskaOrgContract *HauskaOrgContractTransactorSession) VerifyAsset(assetId *big.Int) (*types.Transaction, error) {
	return _HauskaOrgContract.Contract.VerifyAsset(&_HauskaOrgContract.TransactOpts, assetId)
}

// HauskaOrgContractAssetTransferredIterator is returned from FilterAssetTransferred and is used to iterate over the raw logs and unpacked data for AssetTransferred events raised by the HauskaOrgContract contract.
type HauskaOrgContractAssetTransferredIterator struct {
	Event *HauskaOrgContractAssetTransferred // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractAssetTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractAssetTransferred)
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
		it.Event = new(HauskaOrgContractAssetTransferred)
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
func (it *HauskaOrgContractAssetTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractAssetTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractAssetTransferred represents a AssetTransferred event raised by the HauskaOrgContract contract.
type HauskaOrgContractAssetTransferred struct {
	AssetId *big.Int
	From    common.Address
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetTransferred is a free log retrieval operation binding the contract event 0xa993eb3a10693085bc7afc1de0202310fbd5992b9e51edd263b198f62f20cdae.
//
// Solidity: event AssetTransferred(uint256 indexed assetId, address indexed from, address indexed to)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterAssetTransferred(opts *bind.FilterOpts, assetId []*big.Int, from []common.Address, to []common.Address) (*HauskaOrgContractAssetTransferredIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "AssetTransferred", assetIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractAssetTransferredIterator{contract: _HauskaOrgContract.contract, event: "AssetTransferred", logs: logs, sub: sub}, nil
}

// WatchAssetTransferred is a free log subscription operation binding the contract event 0xa993eb3a10693085bc7afc1de0202310fbd5992b9e51edd263b198f62f20cdae.
//
// Solidity: event AssetTransferred(uint256 indexed assetId, address indexed from, address indexed to)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchAssetTransferred(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractAssetTransferred, assetId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "AssetTransferred", assetIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractAssetTransferred)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "AssetTransferred", log); err != nil {
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

// ParseAssetTransferred is a log parse operation binding the contract event 0xa993eb3a10693085bc7afc1de0202310fbd5992b9e51edd263b198f62f20cdae.
//
// Solidity: event AssetTransferred(uint256 indexed assetId, address indexed from, address indexed to)
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseAssetTransferred(log types.Log) (*HauskaOrgContractAssetTransferred, error) {
	event := new(HauskaOrgContractAssetTransferred)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "AssetTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaOrgContractAssetTransferredCrossOrgIterator is returned from FilterAssetTransferredCrossOrg and is used to iterate over the raw logs and unpacked data for AssetTransferredCrossOrg events raised by the HauskaOrgContract contract.
type HauskaOrgContractAssetTransferredCrossOrgIterator struct {
	Event *HauskaOrgContractAssetTransferredCrossOrg // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractAssetTransferredCrossOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractAssetTransferredCrossOrg)
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
		it.Event = new(HauskaOrgContractAssetTransferredCrossOrg)
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
func (it *HauskaOrgContractAssetTransferredCrossOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractAssetTransferredCrossOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractAssetTransferredCrossOrg represents a AssetTransferredCrossOrg event raised by the HauskaOrgContract contract.
type HauskaOrgContractAssetTransferredCrossOrg struct {
	AssetId    *big.Int
	ToOrg      common.Address
	NewAssetId *big.Int
	NewOwner   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetTransferredCrossOrg is a free log retrieval operation binding the contract event 0xb01de4563989286aa5112427f3d54a73c73109f0f77ec2a6fd9f798d313fe71c.
//
// Solidity: event AssetTransferredCrossOrg(uint256 indexed assetId, address indexed toOrg, uint256 indexed newAssetId, address newOwner)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterAssetTransferredCrossOrg(opts *bind.FilterOpts, assetId []*big.Int, toOrg []common.Address, newAssetId []*big.Int) (*HauskaOrgContractAssetTransferredCrossOrgIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var toOrgRule []interface{}
	for _, toOrgItem := range toOrg {
		toOrgRule = append(toOrgRule, toOrgItem)
	}
	var newAssetIdRule []interface{}
	for _, newAssetIdItem := range newAssetId {
		newAssetIdRule = append(newAssetIdRule, newAssetIdItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "AssetTransferredCrossOrg", assetIdRule, toOrgRule, newAssetIdRule)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractAssetTransferredCrossOrgIterator{contract: _HauskaOrgContract.contract, event: "AssetTransferredCrossOrg", logs: logs, sub: sub}, nil
}

// WatchAssetTransferredCrossOrg is a free log subscription operation binding the contract event 0xb01de4563989286aa5112427f3d54a73c73109f0f77ec2a6fd9f798d313fe71c.
//
// Solidity: event AssetTransferredCrossOrg(uint256 indexed assetId, address indexed toOrg, uint256 indexed newAssetId, address newOwner)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchAssetTransferredCrossOrg(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractAssetTransferredCrossOrg, assetId []*big.Int, toOrg []common.Address, newAssetId []*big.Int) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var toOrgRule []interface{}
	for _, toOrgItem := range toOrg {
		toOrgRule = append(toOrgRule, toOrgItem)
	}
	var newAssetIdRule []interface{}
	for _, newAssetIdItem := range newAssetId {
		newAssetIdRule = append(newAssetIdRule, newAssetIdItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "AssetTransferredCrossOrg", assetIdRule, toOrgRule, newAssetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractAssetTransferredCrossOrg)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "AssetTransferredCrossOrg", log); err != nil {
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

// ParseAssetTransferredCrossOrg is a log parse operation binding the contract event 0xb01de4563989286aa5112427f3d54a73c73109f0f77ec2a6fd9f798d313fe71c.
//
// Solidity: event AssetTransferredCrossOrg(uint256 indexed assetId, address indexed toOrg, uint256 indexed newAssetId, address newOwner)
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseAssetTransferredCrossOrg(log types.Log) (*HauskaOrgContractAssetTransferredCrossOrg, error) {
	event := new(HauskaOrgContractAssetTransferredCrossOrg)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "AssetTransferredCrossOrg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaOrgContractCreatorAddedIterator is returned from FilterCreatorAdded and is used to iterate over the raw logs and unpacked data for CreatorAdded events raised by the HauskaOrgContract contract.
type HauskaOrgContractCreatorAddedIterator struct {
	Event *HauskaOrgContractCreatorAdded // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractCreatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractCreatorAdded)
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
		it.Event = new(HauskaOrgContractCreatorAdded)
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
func (it *HauskaOrgContractCreatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractCreatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractCreatorAdded represents a CreatorAdded event raised by the HauskaOrgContract contract.
type HauskaOrgContractCreatorAdded struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatorAdded is a free log retrieval operation binding the contract event 0x021a687fbe334e9e815cee8b399c0bc42e82356eb7f63a09ddb558a25d3dcdbd.
//
// Solidity: event CreatorAdded(address indexed creator)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterCreatorAdded(opts *bind.FilterOpts, creator []common.Address) (*HauskaOrgContractCreatorAddedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "CreatorAdded", creatorRule)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractCreatorAddedIterator{contract: _HauskaOrgContract.contract, event: "CreatorAdded", logs: logs, sub: sub}, nil
}

// WatchCreatorAdded is a free log subscription operation binding the contract event 0x021a687fbe334e9e815cee8b399c0bc42e82356eb7f63a09ddb558a25d3dcdbd.
//
// Solidity: event CreatorAdded(address indexed creator)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchCreatorAdded(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractCreatorAdded, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "CreatorAdded", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractCreatorAdded)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "CreatorAdded", log); err != nil {
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

// ParseCreatorAdded is a log parse operation binding the contract event 0x021a687fbe334e9e815cee8b399c0bc42e82356eb7f63a09ddb558a25d3dcdbd.
//
// Solidity: event CreatorAdded(address indexed creator)
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseCreatorAdded(log types.Log) (*HauskaOrgContractCreatorAdded, error) {
	event := new(HauskaOrgContractCreatorAdded)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "CreatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaOrgContractCreatorRemovedIterator is returned from FilterCreatorRemoved and is used to iterate over the raw logs and unpacked data for CreatorRemoved events raised by the HauskaOrgContract contract.
type HauskaOrgContractCreatorRemovedIterator struct {
	Event *HauskaOrgContractCreatorRemoved // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractCreatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractCreatorRemoved)
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
		it.Event = new(HauskaOrgContractCreatorRemoved)
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
func (it *HauskaOrgContractCreatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractCreatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractCreatorRemoved represents a CreatorRemoved event raised by the HauskaOrgContract contract.
type HauskaOrgContractCreatorRemoved struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatorRemoved is a free log retrieval operation binding the contract event 0x94a409f50d78dc8628b7499ba5af0848a134b9a935a59c0a45313b67f66920f8.
//
// Solidity: event CreatorRemoved(address indexed creator)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterCreatorRemoved(opts *bind.FilterOpts, creator []common.Address) (*HauskaOrgContractCreatorRemovedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "CreatorRemoved", creatorRule)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractCreatorRemovedIterator{contract: _HauskaOrgContract.contract, event: "CreatorRemoved", logs: logs, sub: sub}, nil
}

// WatchCreatorRemoved is a free log subscription operation binding the contract event 0x94a409f50d78dc8628b7499ba5af0848a134b9a935a59c0a45313b67f66920f8.
//
// Solidity: event CreatorRemoved(address indexed creator)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchCreatorRemoved(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractCreatorRemoved, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "CreatorRemoved", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractCreatorRemoved)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "CreatorRemoved", log); err != nil {
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

// ParseCreatorRemoved is a log parse operation binding the contract event 0x94a409f50d78dc8628b7499ba5af0848a134b9a935a59c0a45313b67f66920f8.
//
// Solidity: event CreatorRemoved(address indexed creator)
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseCreatorRemoved(log types.Log) (*HauskaOrgContractCreatorRemoved, error) {
	event := new(HauskaOrgContractCreatorRemoved)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "CreatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaOrgContractModuleSetIterator is returned from FilterModuleSet and is used to iterate over the raw logs and unpacked data for ModuleSet events raised by the HauskaOrgContract contract.
type HauskaOrgContractModuleSetIterator struct {
	Event *HauskaOrgContractModuleSet // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractModuleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractModuleSet)
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
		it.Event = new(HauskaOrgContractModuleSet)
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
func (it *HauskaOrgContractModuleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractModuleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractModuleSet represents a ModuleSet event raised by the HauskaOrgContract contract.
type HauskaOrgContractModuleSet struct {
	ModuleName    string
	ModuleAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterModuleSet is a free log retrieval operation binding the contract event 0xebcfeb887c9fe4eaee3b330008d80b3d2c92a4fd53a9ba188170c974e31278e8.
//
// Solidity: event ModuleSet(string moduleName, address moduleAddress)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterModuleSet(opts *bind.FilterOpts) (*HauskaOrgContractModuleSetIterator, error) {

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "ModuleSet")
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractModuleSetIterator{contract: _HauskaOrgContract.contract, event: "ModuleSet", logs: logs, sub: sub}, nil
}

// WatchModuleSet is a free log subscription operation binding the contract event 0xebcfeb887c9fe4eaee3b330008d80b3d2c92a4fd53a9ba188170c974e31278e8.
//
// Solidity: event ModuleSet(string moduleName, address moduleAddress)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchModuleSet(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractModuleSet) (event.Subscription, error) {

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "ModuleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractModuleSet)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "ModuleSet", log); err != nil {
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

// ParseModuleSet is a log parse operation binding the contract event 0xebcfeb887c9fe4eaee3b330008d80b3d2c92a4fd53a9ba188170c974e31278e8.
//
// Solidity: event ModuleSet(string moduleName, address moduleAddress)
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseModuleSet(log types.Log) (*HauskaOrgContractModuleSet, error) {
	event := new(HauskaOrgContractModuleSet)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "ModuleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaOrgContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HauskaOrgContract contract.
type HauskaOrgContractRoleAdminChangedIterator struct {
	Event *HauskaOrgContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractRoleAdminChanged)
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
		it.Event = new(HauskaOrgContractRoleAdminChanged)
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
func (it *HauskaOrgContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractRoleAdminChanged represents a RoleAdminChanged event raised by the HauskaOrgContract contract.
type HauskaOrgContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HauskaOrgContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractRoleAdminChangedIterator{contract: _HauskaOrgContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractRoleAdminChanged)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseRoleAdminChanged(log types.Log) (*HauskaOrgContractRoleAdminChanged, error) {
	event := new(HauskaOrgContractRoleAdminChanged)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaOrgContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HauskaOrgContract contract.
type HauskaOrgContractRoleGrantedIterator struct {
	Event *HauskaOrgContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractRoleGranted)
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
		it.Event = new(HauskaOrgContractRoleGranted)
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
func (it *HauskaOrgContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractRoleGranted represents a RoleGranted event raised by the HauskaOrgContract contract.
type HauskaOrgContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaOrgContractRoleGrantedIterator, error) {

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

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractRoleGrantedIterator{contract: _HauskaOrgContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractRoleGranted)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseRoleGranted(log types.Log) (*HauskaOrgContractRoleGranted, error) {
	event := new(HauskaOrgContractRoleGranted)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HauskaOrgContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HauskaOrgContract contract.
type HauskaOrgContractRoleRevokedIterator struct {
	Event *HauskaOrgContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HauskaOrgContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HauskaOrgContractRoleRevoked)
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
		it.Event = new(HauskaOrgContractRoleRevoked)
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
func (it *HauskaOrgContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HauskaOrgContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HauskaOrgContractRoleRevoked represents a RoleRevoked event raised by the HauskaOrgContract contract.
type HauskaOrgContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaOrgContract *HauskaOrgContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HauskaOrgContractRoleRevokedIterator, error) {

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

	logs, sub, err := _HauskaOrgContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HauskaOrgContractRoleRevokedIterator{contract: _HauskaOrgContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HauskaOrgContract *HauskaOrgContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HauskaOrgContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HauskaOrgContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HauskaOrgContractRoleRevoked)
				if err := _HauskaOrgContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HauskaOrgContract *HauskaOrgContractFilterer) ParseRoleRevoked(log types.Log) (*HauskaOrgContractRoleRevoked, error) {
	event := new(HauskaOrgContractRoleRevoked)
	if err := _HauskaOrgContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
