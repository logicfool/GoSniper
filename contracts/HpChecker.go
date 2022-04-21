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
)

// HoneyPotMetaData contains all meta data concerning the HoneyPot contract.
var HoneyPotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"checkHPAndTaxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_val\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// HoneyPotABI is the input ABI used to generate the binding from.
// Deprecated: Use HoneyPotMetaData.ABI instead.
var HoneyPotABI = HoneyPotMetaData.ABI

// HoneyPot is an auto generated Go binding around an Ethereum contract.
type HoneyPot struct {
	HoneyPotCaller     // Read-only binding to the contract
	HoneyPotTransactor // Write-only binding to the contract
	HoneyPotFilterer   // Log filterer for contract events
}

// HoneyPotCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoneyPotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneyPotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoneyPotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneyPotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoneyPotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneyPotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoneyPotSession struct {
	Contract     *HoneyPot         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoneyPotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoneyPotCallerSession struct {
	Contract *HoneyPotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HoneyPotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoneyPotTransactorSession struct {
	Contract     *HoneyPotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HoneyPotRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoneyPotRaw struct {
	Contract *HoneyPot // Generic contract binding to access the raw methods on
}

// HoneyPotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoneyPotCallerRaw struct {
	Contract *HoneyPotCaller // Generic read-only contract binding to access the raw methods on
}

// HoneyPotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoneyPotTransactorRaw struct {
	Contract *HoneyPotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoneyPot creates a new instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPot(address common.Address, backend bind.ContractBackend) (*HoneyPot, error) {
	contract, err := bindHoneyPot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoneyPot{HoneyPotCaller: HoneyPotCaller{contract: contract}, HoneyPotTransactor: HoneyPotTransactor{contract: contract}, HoneyPotFilterer: HoneyPotFilterer{contract: contract}}, nil
}

// NewHoneyPotCaller creates a new read-only instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPotCaller(address common.Address, caller bind.ContractCaller) (*HoneyPotCaller, error) {
	contract, err := bindHoneyPot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoneyPotCaller{contract: contract}, nil
}

// NewHoneyPotTransactor creates a new write-only instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPotTransactor(address common.Address, transactor bind.ContractTransactor) (*HoneyPotTransactor, error) {
	contract, err := bindHoneyPot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoneyPotTransactor{contract: contract}, nil
}

// NewHoneyPotFilterer creates a new log filterer instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPotFilterer(address common.Address, filterer bind.ContractFilterer) (*HoneyPotFilterer, error) {
	contract, err := bindHoneyPot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoneyPotFilterer{contract: contract}, nil
}

// bindHoneyPot binds a generic wrapper to an already deployed contract.
func bindHoneyPot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoneyPotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoneyPot *HoneyPotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoneyPot.Contract.HoneyPotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoneyPot *HoneyPotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoneyPot.Contract.HoneyPotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoneyPot *HoneyPotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoneyPot.Contract.HoneyPotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoneyPot *HoneyPotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoneyPot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoneyPot *HoneyPotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoneyPot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoneyPot *HoneyPotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoneyPot.Contract.contract.Transact(opts, method, params...)
}

// CheckHPAndTaxes is a paid mutator transaction binding the contract method 0x7f48c2de.
//
// Solidity: function checkHPAndTaxes(address _router, address[] path) payable returns(uint256, uint256, uint256, uint256)
func (_HoneyPot *HoneyPotTransactor) CheckHPAndTaxes(opts *bind.TransactOpts, _router common.Address, path []common.Address) (*types.Transaction, error) {
	return _HoneyPot.contract.Transact(opts, "checkHPAndTaxes", _router, path)
}

// CheckHPAndTaxes is a paid mutator transaction binding the contract method 0x7f48c2de.
//
// Solidity: function checkHPAndTaxes(address _router, address[] path) payable returns(uint256, uint256, uint256, uint256)
func (_HoneyPot *HoneyPotSession) CheckHPAndTaxes(_router common.Address, path []common.Address) (*types.Transaction, error) {
	return _HoneyPot.Contract.CheckHPAndTaxes(&_HoneyPot.TransactOpts, _router, path)
}

// CheckHPAndTaxes is a paid mutator transaction binding the contract method 0x7f48c2de.
//
// Solidity: function checkHPAndTaxes(address _router, address[] path) payable returns(uint256, uint256, uint256, uint256)
func (_HoneyPot *HoneyPotTransactorSession) CheckHPAndTaxes(_router common.Address, path []common.Address) (*types.Transaction, error) {
	return _HoneyPot.Contract.CheckHPAndTaxes(&_HoneyPot.TransactOpts, _router, path)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(bool _val) returns()
func (_HoneyPot *HoneyPotTransactor) SetLocked(opts *bind.TransactOpts, _val bool) (*types.Transaction, error) {
	return _HoneyPot.contract.Transact(opts, "setLocked", _val)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(bool _val) returns()
func (_HoneyPot *HoneyPotSession) SetLocked(_val bool) (*types.Transaction, error) {
	return _HoneyPot.Contract.SetLocked(&_HoneyPot.TransactOpts, _val)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(bool _val) returns()
func (_HoneyPot *HoneyPotTransactorSession) SetLocked(_val bool) (*types.Transaction, error) {
	return _HoneyPot.Contract.SetLocked(&_HoneyPot.TransactOpts, _val)
}
