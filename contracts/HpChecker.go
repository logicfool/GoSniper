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
var HoneyPotMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***],\"name\":\"checkHPAndTaxes\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"_val\",\"type\":\"bool\"***REMOVED***],\"name\":\"setLocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"***REMOVED***]",
***REMOVED***

// HoneyPotABI is the input ABI used to generate the binding from.
// Deprecated: Use HoneyPotMetaData.ABI instead.
var HoneyPotABI = HoneyPotMetaData.ABI

// HoneyPot is an auto generated Go binding around an Ethereum contract.
type HoneyPot struct ***REMOVED***
	HoneyPotCaller     // Read-only binding to the contract
	HoneyPotTransactor // Write-only binding to the contract
	HoneyPotFilterer   // Log filterer for contract events
***REMOVED***

// HoneyPotCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoneyPotCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// HoneyPotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoneyPotTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// HoneyPotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoneyPotFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// HoneyPotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoneyPotSession struct ***REMOVED***
	Contract     *HoneyPot         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// HoneyPotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoneyPotCallerSession struct ***REMOVED***
	Contract *HoneyPotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
***REMOVED***

// HoneyPotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoneyPotTransactorSession struct ***REMOVED***
	Contract     *HoneyPotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
***REMOVED***

// HoneyPotRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoneyPotRaw struct ***REMOVED***
	Contract *HoneyPot // Generic contract binding to access the raw methods on
***REMOVED***

// HoneyPotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoneyPotCallerRaw struct ***REMOVED***
	Contract *HoneyPotCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// HoneyPotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoneyPotTransactorRaw struct ***REMOVED***
	Contract *HoneyPotTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewHoneyPot creates a new instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPot(address common.Address, backend bind.ContractBackend) (*HoneyPot, error) ***REMOVED***
	contract, err := bindHoneyPot(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &HoneyPot***REMOVED***HoneyPotCaller: HoneyPotCaller***REMOVED***contract: contract***REMOVED***, HoneyPotTransactor: HoneyPotTransactor***REMOVED***contract: contract***REMOVED***, HoneyPotFilterer: HoneyPotFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewHoneyPotCaller creates a new read-only instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPotCaller(address common.Address, caller bind.ContractCaller) (*HoneyPotCaller, error) ***REMOVED***
	contract, err := bindHoneyPot(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &HoneyPotCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewHoneyPotTransactor creates a new write-only instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPotTransactor(address common.Address, transactor bind.ContractTransactor) (*HoneyPotTransactor, error) ***REMOVED***
	contract, err := bindHoneyPot(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &HoneyPotTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewHoneyPotFilterer creates a new log filterer instance of HoneyPot, bound to a specific deployed contract.
func NewHoneyPotFilterer(address common.Address, filterer bind.ContractFilterer) (*HoneyPotFilterer, error) ***REMOVED***
	contract, err := bindHoneyPot(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &HoneyPotFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindHoneyPot binds a generic wrapper to an already deployed contract.
func bindHoneyPot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(HoneyPotABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoneyPot *HoneyPotRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _HoneyPot.Contract.HoneyPotCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoneyPot *HoneyPotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.HoneyPotTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_HoneyPot *HoneyPotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.HoneyPotTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoneyPot *HoneyPotCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _HoneyPot.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoneyPot *HoneyPotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_HoneyPot *HoneyPotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// CheckHPAndTaxes is a paid mutator transaction binding the contract method 0x7f48c2de.
//
// Solidity: function checkHPAndTaxes(address _router, address[] path) payable returns(uint256, uint256, uint256, uint256)
func (_HoneyPot *HoneyPotTransactor) CheckHPAndTaxes(opts *bind.TransactOpts, _router common.Address, path []common.Address) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.contract.Transact(opts, "checkHPAndTaxes", _router, path)
***REMOVED***

// CheckHPAndTaxes is a paid mutator transaction binding the contract method 0x7f48c2de.
//
// Solidity: function checkHPAndTaxes(address _router, address[] path) payable returns(uint256, uint256, uint256, uint256)
func (_HoneyPot *HoneyPotSession) CheckHPAndTaxes(_router common.Address, path []common.Address) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.CheckHPAndTaxes(&_HoneyPot.TransactOpts, _router, path)
***REMOVED***

// CheckHPAndTaxes is a paid mutator transaction binding the contract method 0x7f48c2de.
//
// Solidity: function checkHPAndTaxes(address _router, address[] path) payable returns(uint256, uint256, uint256, uint256)
func (_HoneyPot *HoneyPotTransactorSession) CheckHPAndTaxes(_router common.Address, path []common.Address) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.CheckHPAndTaxes(&_HoneyPot.TransactOpts, _router, path)
***REMOVED***

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(bool _val) returns()
func (_HoneyPot *HoneyPotTransactor) SetLocked(opts *bind.TransactOpts, _val bool) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.contract.Transact(opts, "setLocked", _val)
***REMOVED***

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(bool _val) returns()
func (_HoneyPot *HoneyPotSession) SetLocked(_val bool) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.SetLocked(&_HoneyPot.TransactOpts, _val)
***REMOVED***

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(bool _val) returns()
func (_HoneyPot *HoneyPotTransactorSession) SetLocked(_val bool) (*types.Transaction, error) ***REMOVED***
	return _HoneyPot.Contract.SetLocked(&_HoneyPot.TransactOpts, _val)
***REMOVED***
