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

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"***REMOVED***,***REMOVED***\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"addresspayable\",\"name\":\"feeReceiver\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"constructor\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Approval\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"***REMOVED***],\"name\":\"OwnershipTransferred\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Transfer\",\"type\":\"event\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***],\"name\":\"allowance\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"***REMOVED***],\"name\":\"approve\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"***REMOVED***],\"name\":\"balanceOf\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"decimals\",\"outputs\":[***REMOVED***\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"***REMOVED***],\"name\":\"decreaseAllowance\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"generator\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"***REMOVED***],\"name\":\"increaseAllowance\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"name\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"owner\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"symbol\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transfer\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transferFrom\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"***REMOVED***],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"version\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"name\":\"wad\",\"type\":\"uint256\"***REMOVED***],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***]",
***REMOVED***

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct ***REMOVED***
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
***REMOVED***

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct ***REMOVED***
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct ***REMOVED***
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
***REMOVED***

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct ***REMOVED***
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct ***REMOVED***
	Contract *ERC20 // Generic contract binding to access the raw methods on
***REMOVED***

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct ***REMOVED***
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct ***REMOVED***
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) ***REMOVED***
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ERC20***REMOVED***ERC20Caller: ERC20Caller***REMOVED***contract: contract***REMOVED***, ERC20Transactor: ERC20Transactor***REMOVED***contract: contract***REMOVED***, ERC20Filterer: ERC20Filterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) ***REMOVED***
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ERC20Caller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) ***REMOVED***
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ERC20Transactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) ***REMOVED***
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ERC20Filterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) ***REMOVED***
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) ***REMOVED***
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil ***REMOVED***
		return *new(uint8), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) ***REMOVED***
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) ***REMOVED***
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
***REMOVED***

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() pure returns(string)
func (_ERC20 *ERC20Caller) Generator(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "generator")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() pure returns(string)
func (_ERC20 *ERC20Session) Generator() (string, error) ***REMOVED***
	return _ERC20.Contract.Generator(&_ERC20.CallOpts)
***REMOVED***

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() pure returns(string)
func (_ERC20 *ERC20CallerSession) Generator() (string, error) ***REMOVED***
	return _ERC20.Contract.Generator(&_ERC20.CallOpts)
***REMOVED***

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ERC20 *ERC20Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "getOwner")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ERC20 *ERC20Session) GetOwner() (common.Address, error) ***REMOVED***
	return _ERC20.Contract.GetOwner(&_ERC20.CallOpts)
***REMOVED***

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ERC20 *ERC20CallerSession) GetOwner() (common.Address, error) ***REMOVED***
	return _ERC20.Contract.GetOwner(&_ERC20.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) ***REMOVED***
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) ***REMOVED***
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "owner")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Session) Owner() (common.Address, error) ***REMOVED***
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20CallerSession) Owner() (common.Address, error) ***REMOVED***
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) ***REMOVED***
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) ***REMOVED***
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) ***REMOVED***
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) ***REMOVED***
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
***REMOVED***

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ERC20 *ERC20Caller) Version(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ERC20.contract.Call(opts, &out, "version")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ERC20 *ERC20Session) Version() (string, error) ***REMOVED***
	return _ERC20.Contract.Version(&_ERC20.CallOpts)
***REMOVED***

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ERC20 *ERC20CallerSession) Version() (string, error) ***REMOVED***
	return _ERC20.Contract.Version(&_ERC20.CallOpts)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
***REMOVED***

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
***REMOVED***

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
***REMOVED***

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ERC20 *ERC20Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "deposit")
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ERC20 *ERC20Session) Deposit() (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Deposit(&_ERC20.TransactOpts)
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ERC20 *ERC20TransactorSession) Deposit() (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Deposit(&_ERC20.TransactOpts)
***REMOVED***

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
***REMOVED***

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
***REMOVED***

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
***REMOVED***

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20 *ERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "renounceOwnership")
***REMOVED***

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20 *ERC20Session) RenounceOwnership() (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.RenounceOwnership(&_ERC20.TransactOpts)
***REMOVED***

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20 *ERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.RenounceOwnership(&_ERC20.TransactOpts)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "transferOwnership", newOwner)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.TransferOwnership(&_ERC20.TransactOpts, newOwner)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.TransferOwnership(&_ERC20.TransactOpts, newOwner)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_ERC20 *ERC20Transactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.contract.Transact(opts, "withdraw", wad)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_ERC20 *ERC20Session) Withdraw(wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Withdraw(&_ERC20.TransactOpts, wad)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_ERC20 *ERC20TransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ERC20.Contract.Withdraw(&_ERC20.TransactOpts, wad)
***REMOVED***

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct ***REMOVED***
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct ***REMOVED***
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ERC20ApprovalIterator***REMOVED***contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) ***REMOVED***
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// ERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20 contract.
type ERC20OwnershipTransferredIterator struct ***REMOVED***
	Event *ERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20OwnershipTransferredIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(ERC20OwnershipTransferred)
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
		it.Event = new(ERC20OwnershipTransferred)
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
func (it *ERC20OwnershipTransferredIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20OwnershipTransferredIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// ERC20OwnershipTransferred represents a OwnershipTransferred event raised by the ERC20 contract.
type ERC20OwnershipTransferred struct ***REMOVED***
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20 *ERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20OwnershipTransferredIterator, error) ***REMOVED***

	var previousOwnerRule []interface***REMOVED******REMOVED***
	for _, previousOwnerItem := range previousOwner ***REMOVED***
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	***REMOVED***
	var newOwnerRule []interface***REMOVED******REMOVED***
	for _, newOwnerItem := range newOwner ***REMOVED***
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	***REMOVED***

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ERC20OwnershipTransferredIterator***REMOVED***contract: _ERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20 *ERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) ***REMOVED***

	var previousOwnerRule []interface***REMOVED******REMOVED***
	for _, previousOwnerItem := range previousOwner ***REMOVED***
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	***REMOVED***
	var newOwnerRule []interface***REMOVED******REMOVED***
	for _, newOwnerItem := range newOwner ***REMOVED***
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	***REMOVED***

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20OwnershipTransferred)
				if err := _ERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil ***REMOVED***
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20 *ERC20Filterer) ParseOwnershipTransferred(log types.Log) (*ERC20OwnershipTransferred, error) ***REMOVED***
	event := new(ERC20OwnershipTransferred)
	if err := _ERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct ***REMOVED***
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct ***REMOVED***
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ERC20TransferIterator***REMOVED***contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) ***REMOVED***
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***
