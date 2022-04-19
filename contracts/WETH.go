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

// WETHMetaData contains all meta data concerning the WETH contract.
var WETHMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"string\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"name\":\"guy\",\"type\":\"address\"***REMOVED***,***REMOVED***\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"approve\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"name\":\"src\",\"type\":\"address\"***REMOVED***,***REMOVED***\"name\":\"dst\",\"type\":\"address\"***REMOVED***,***REMOVED***\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transferFrom\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"uint8\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[***REMOVED***\"name\":\"\",\"type\":\"address\"***REMOVED***],\"name\":\"balanceOf\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"string\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"name\":\"dst\",\"type\":\"address\"***REMOVED***,***REMOVED***\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transfer\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[***REMOVED***\"name\":\"\",\"type\":\"address\"***REMOVED***,***REMOVED***\"name\":\"\",\"type\":\"address\"***REMOVED***],\"name\":\"allowance\",\"outputs\":[***REMOVED***\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"name\":\"src\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Approval\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"name\":\"src\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Transfer\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Deposit\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"name\":\"src\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Withdrawal\",\"type\":\"event\"***REMOVED***]",
***REMOVED***

// WETHABI is the input ABI used to generate the binding from.
// Deprecated: Use WETHMetaData.ABI instead.
var WETHABI = WETHMetaData.ABI

// WETH is an auto generated Go binding around an Ethereum contract.
type WETH struct ***REMOVED***
	WETHCaller     // Read-only binding to the contract
	WETHTransactor // Write-only binding to the contract
	WETHFilterer   // Log filterer for contract events
***REMOVED***

// WETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// WETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// WETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// WETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHSession struct ***REMOVED***
	Contract     *WETH             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// WETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHCallerSession struct ***REMOVED***
	Contract *WETHCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
***REMOVED***

// WETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHTransactorSession struct ***REMOVED***
	Contract     *WETHTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// WETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHRaw struct ***REMOVED***
	Contract *WETH // Generic contract binding to access the raw methods on
***REMOVED***

// WETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHCallerRaw struct ***REMOVED***
	Contract *WETHCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// WETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHTransactorRaw struct ***REMOVED***
	Contract *WETHTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewWETH creates a new instance of WETH, bound to a specific deployed contract.
func NewWETH(address common.Address, backend bind.ContractBackend) (*WETH, error) ***REMOVED***
	contract, err := bindWETH(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETH***REMOVED***WETHCaller: WETHCaller***REMOVED***contract: contract***REMOVED***, WETHTransactor: WETHTransactor***REMOVED***contract: contract***REMOVED***, WETHFilterer: WETHFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewWETHCaller creates a new read-only instance of WETH, bound to a specific deployed contract.
func NewWETHCaller(address common.Address, caller bind.ContractCaller) (*WETHCaller, error) ***REMOVED***
	contract, err := bindWETH(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETHCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewWETHTransactor creates a new write-only instance of WETH, bound to a specific deployed contract.
func NewWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHTransactor, error) ***REMOVED***
	contract, err := bindWETH(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETHTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewWETHFilterer creates a new log filterer instance of WETH, bound to a specific deployed contract.
func NewWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHFilterer, error) ***REMOVED***
	contract, err := bindWETH(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETHFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindWETH binds a generic wrapper to an already deployed contract.
func bindWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(WETHABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH *WETHRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _WETH.Contract.WETHCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH *WETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.WETHTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_WETH *WETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.WETHTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH *WETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _WETH.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH *WETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_WETH *WETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WETH *WETHCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _WETH.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WETH *WETHSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) ***REMOVED***
	return _WETH.Contract.Allowance(&_WETH.CallOpts, arg0, arg1)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WETH *WETHCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) ***REMOVED***
	return _WETH.Contract.Allowance(&_WETH.CallOpts, arg0, arg1)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WETH *WETHCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _WETH.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WETH *WETHSession) BalanceOf(arg0 common.Address) (*big.Int, error) ***REMOVED***
	return _WETH.Contract.BalanceOf(&_WETH.CallOpts, arg0)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WETH *WETHCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) ***REMOVED***
	return _WETH.Contract.BalanceOf(&_WETH.CallOpts, arg0)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH *WETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _WETH.contract.Call(opts, &out, "decimals")

	if err != nil ***REMOVED***
		return *new(uint8), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH *WETHSession) Decimals() (uint8, error) ***REMOVED***
	return _WETH.Contract.Decimals(&_WETH.CallOpts)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH *WETHCallerSession) Decimals() (uint8, error) ***REMOVED***
	return _WETH.Contract.Decimals(&_WETH.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH *WETHCaller) Name(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _WETH.contract.Call(opts, &out, "name")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH *WETHSession) Name() (string, error) ***REMOVED***
	return _WETH.Contract.Name(&_WETH.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH *WETHCallerSession) Name() (string, error) ***REMOVED***
	return _WETH.Contract.Name(&_WETH.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH *WETHCaller) Symbol(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _WETH.contract.Call(opts, &out, "symbol")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH *WETHSession) Symbol() (string, error) ***REMOVED***
	return _WETH.Contract.Symbol(&_WETH.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH *WETHCallerSession) Symbol() (string, error) ***REMOVED***
	return _WETH.Contract.Symbol(&_WETH.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _WETH.contract.Call(opts, &out, "totalSupply")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHSession) TotalSupply() (*big.Int, error) ***REMOVED***
	return _WETH.Contract.TotalSupply(&_WETH.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHCallerSession) TotalSupply() (*big.Int, error) ***REMOVED***
	return _WETH.Contract.TotalSupply(&_WETH.CallOpts)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WETH *WETHTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.contract.Transact(opts, "approve", guy, wad)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WETH *WETHSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Approve(&_WETH.TransactOpts, guy, wad)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WETH *WETHTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Approve(&_WETH.TransactOpts, guy, wad)
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH *WETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _WETH.contract.Transact(opts, "deposit")
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH *WETHSession) Deposit() (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Deposit(&_WETH.TransactOpts)
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH *WETHTransactorSession) Deposit() (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Deposit(&_WETH.TransactOpts)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WETH *WETHTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.contract.Transact(opts, "transfer", dst, wad)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WETH *WETHSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Transfer(&_WETH.TransactOpts, dst, wad)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WETH *WETHTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Transfer(&_WETH.TransactOpts, dst, wad)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WETH *WETHTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.contract.Transact(opts, "transferFrom", src, dst, wad)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WETH *WETHSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.TransferFrom(&_WETH.TransactOpts, src, dst, wad)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WETH *WETHTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.TransferFrom(&_WETH.TransactOpts, src, dst, wad)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETH *WETHTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.contract.Transact(opts, "withdraw", wad)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETH *WETHSession) Withdraw(wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Withdraw(&_WETH.TransactOpts, wad)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETH *WETHTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Withdraw(&_WETH.TransactOpts, wad)
***REMOVED***

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WETH *WETHTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) ***REMOVED***
	return _WETH.contract.RawTransact(opts, calldata)
***REMOVED***

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WETH *WETHSession) Fallback(calldata []byte) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Fallback(&_WETH.TransactOpts, calldata)
***REMOVED***

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WETH *WETHTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) ***REMOVED***
	return _WETH.Contract.Fallback(&_WETH.TransactOpts, calldata)
***REMOVED***

// WETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WETH contract.
type WETHApprovalIterator struct ***REMOVED***
	Event *WETHApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
***REMOVED***

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WETHApprovalIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(WETHApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
				it.fail = err
				return false
			***REMOVED***
			it.Event.Raw = log
			return true

		default:
			return false
		***REMOVED***
	***REMOVED***
	// Iterator still in progress, wait for either a data or an error event
	select ***REMOVED***
	case log := <-it.logs:
		it.Event = new(WETHApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
			it.fail = err
			return false
		***REMOVED***
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	***REMOVED***
***REMOVED***

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WETHApprovalIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHApprovalIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// WETHApproval represents a Approval event raised by the WETH contract.
type WETHApproval struct ***REMOVED***
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WETH *WETHFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WETHApprovalIterator, error) ***REMOVED***

	var srcRule []interface***REMOVED******REMOVED***
	for _, srcItem := range src ***REMOVED***
		srcRule = append(srcRule, srcItem)
	***REMOVED***
	var guyRule []interface***REMOVED******REMOVED***
	for _, guyItem := range guy ***REMOVED***
		guyRule = append(guyRule, guyItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETHApprovalIterator***REMOVED***contract: _WETH.contract, event: "Approval", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WETH *WETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETHApproval, src []common.Address, guy []common.Address) (event.Subscription, error) ***REMOVED***

	var srcRule []interface***REMOVED******REMOVED***
	for _, srcItem := range src ***REMOVED***
		srcRule = append(srcRule, srcItem)
	***REMOVED***
	var guyRule []interface***REMOVED******REMOVED***
	for _, guyItem := range guy ***REMOVED***
		guyRule = append(guyRule, guyItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHApproval)
				if err := _WETH.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
					return err
				***REMOVED***
				event.Raw = log

				select ***REMOVED***
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				***REMOVED***
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			***REMOVED***
		***REMOVED***
	***REMOVED***), nil
***REMOVED***

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WETH *WETHFilterer) ParseApproval(log types.Log) (*WETHApproval, error) ***REMOVED***
	event := new(WETHApproval)
	if err := _WETH.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// WETHDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WETH contract.
type WETHDepositIterator struct ***REMOVED***
	Event *WETHDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
***REMOVED***

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WETHDepositIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(WETHDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
				it.fail = err
				return false
			***REMOVED***
			it.Event.Raw = log
			return true

		default:
			return false
		***REMOVED***
	***REMOVED***
	// Iterator still in progress, wait for either a data or an error event
	select ***REMOVED***
	case log := <-it.logs:
		it.Event = new(WETHDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
			it.fail = err
			return false
		***REMOVED***
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	***REMOVED***
***REMOVED***

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WETHDepositIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHDepositIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// WETHDeposit represents a Deposit event raised by the WETH contract.
type WETHDeposit struct ***REMOVED***
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WETH *WETHFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WETHDepositIterator, error) ***REMOVED***

	var dstRule []interface***REMOVED******REMOVED***
	for _, dstItem := range dst ***REMOVED***
		dstRule = append(dstRule, dstItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETHDepositIterator***REMOVED***contract: _WETH.contract, event: "Deposit", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WETH *WETHFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WETHDeposit, dst []common.Address) (event.Subscription, error) ***REMOVED***

	var dstRule []interface***REMOVED******REMOVED***
	for _, dstItem := range dst ***REMOVED***
		dstRule = append(dstRule, dstItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHDeposit)
				if err := _WETH.contract.UnpackLog(event, "Deposit", log); err != nil ***REMOVED***
					return err
				***REMOVED***
				event.Raw = log

				select ***REMOVED***
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				***REMOVED***
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			***REMOVED***
		***REMOVED***
	***REMOVED***), nil
***REMOVED***

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WETH *WETHFilterer) ParseDeposit(log types.Log) (*WETHDeposit, error) ***REMOVED***
	event := new(WETHDeposit)
	if err := _WETH.contract.UnpackLog(event, "Deposit", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// WETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WETH contract.
type WETHTransferIterator struct ***REMOVED***
	Event *WETHTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
***REMOVED***

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WETHTransferIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(WETHTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
				it.fail = err
				return false
			***REMOVED***
			it.Event.Raw = log
			return true

		default:
			return false
		***REMOVED***
	***REMOVED***
	// Iterator still in progress, wait for either a data or an error event
	select ***REMOVED***
	case log := <-it.logs:
		it.Event = new(WETHTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
			it.fail = err
			return false
		***REMOVED***
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	***REMOVED***
***REMOVED***

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WETHTransferIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHTransferIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// WETHTransfer represents a Transfer event raised by the WETH contract.
type WETHTransfer struct ***REMOVED***
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WETH *WETHFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WETHTransferIterator, error) ***REMOVED***

	var srcRule []interface***REMOVED******REMOVED***
	for _, srcItem := range src ***REMOVED***
		srcRule = append(srcRule, srcItem)
	***REMOVED***
	var dstRule []interface***REMOVED******REMOVED***
	for _, dstItem := range dst ***REMOVED***
		dstRule = append(dstRule, dstItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETHTransferIterator***REMOVED***contract: _WETH.contract, event: "Transfer", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WETH *WETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETHTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) ***REMOVED***

	var srcRule []interface***REMOVED******REMOVED***
	for _, srcItem := range src ***REMOVED***
		srcRule = append(srcRule, srcItem)
	***REMOVED***
	var dstRule []interface***REMOVED******REMOVED***
	for _, dstItem := range dst ***REMOVED***
		dstRule = append(dstRule, dstItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHTransfer)
				if err := _WETH.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
					return err
				***REMOVED***
				event.Raw = log

				select ***REMOVED***
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				***REMOVED***
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			***REMOVED***
		***REMOVED***
	***REMOVED***), nil
***REMOVED***

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WETH *WETHFilterer) ParseTransfer(log types.Log) (*WETHTransfer, error) ***REMOVED***
	event := new(WETHTransfer)
	if err := _WETH.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// WETHWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WETH contract.
type WETHWithdrawalIterator struct ***REMOVED***
	Event *WETHWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
***REMOVED***

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WETHWithdrawalIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(WETHWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
				it.fail = err
				return false
			***REMOVED***
			it.Event.Raw = log
			return true

		default:
			return false
		***REMOVED***
	***REMOVED***
	// Iterator still in progress, wait for either a data or an error event
	select ***REMOVED***
	case log := <-it.logs:
		it.Event = new(WETHWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil ***REMOVED***
			it.fail = err
			return false
		***REMOVED***
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	***REMOVED***
***REMOVED***

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WETHWithdrawalIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHWithdrawalIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// WETHWithdrawal represents a Withdrawal event raised by the WETH contract.
type WETHWithdrawal struct ***REMOVED***
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WETH *WETHFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WETHWithdrawalIterator, error) ***REMOVED***

	var srcRule []interface***REMOVED******REMOVED***
	for _, srcItem := range src ***REMOVED***
		srcRule = append(srcRule, srcItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &WETHWithdrawalIterator***REMOVED***contract: _WETH.contract, event: "Withdrawal", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WETH *WETHFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WETHWithdrawal, src []common.Address) (event.Subscription, error) ***REMOVED***

	var srcRule []interface***REMOVED******REMOVED***
	for _, srcItem := range src ***REMOVED***
		srcRule = append(srcRule, srcItem)
	***REMOVED***

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHWithdrawal)
				if err := _WETH.contract.UnpackLog(event, "Withdrawal", log); err != nil ***REMOVED***
					return err
				***REMOVED***
				event.Raw = log

				select ***REMOVED***
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				***REMOVED***
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			***REMOVED***
		***REMOVED***
	***REMOVED***), nil
***REMOVED***

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WETH *WETHFilterer) ParseWithdrawal(log types.Log) (*WETHWithdrawal, error) ***REMOVED***
	event := new(WETHWithdrawal)
	if err := _WETH.contract.UnpackLog(event, "Withdrawal", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***
