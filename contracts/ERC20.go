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

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"***REMOVED***,***REMOVED***\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"addresspayable\",\"name\":\"feeReceiver\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"constructor\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Approval\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"***REMOVED***],\"name\":\"OwnershipTransferred\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Transfer\",\"type\":\"event\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***],\"name\":\"allowance\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"***REMOVED***],\"name\":\"approve\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"***REMOVED***],\"name\":\"balanceOf\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"decimals\",\"outputs\":[***REMOVED***\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"***REMOVED***],\"name\":\"decreaseAllowance\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"generator\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"***REMOVED***],\"name\":\"increaseAllowance\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"name\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"owner\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"symbol\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transfer\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transferFrom\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"***REMOVED***],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"version\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***]",
***REMOVED***

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct ***REMOVED***
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
***REMOVED***

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct ***REMOVED***
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct ***REMOVED***
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
***REMOVED***

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct ***REMOVED***
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct ***REMOVED***
	Contract *Erc20 // Generic contract binding to access the raw methods on
***REMOVED***

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct ***REMOVED***
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct ***REMOVED***
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) ***REMOVED***
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Erc20***REMOVED***Erc20Caller: Erc20Caller***REMOVED***contract: contract***REMOVED***, Erc20Transactor: Erc20Transactor***REMOVED***contract: contract***REMOVED***, Erc20Filterer: Erc20Filterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) ***REMOVED***
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Erc20Caller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) ***REMOVED***
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Erc20Transactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) ***REMOVED***
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Erc20Filterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, owner, spender)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, owner, spender)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(account common.Address) (*big.Int, error) ***REMOVED***
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, account)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20 *Erc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) ***REMOVED***
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, account)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil ***REMOVED***
		return *new(uint8), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Session) Decimals() (uint8, error) ***REMOVED***
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20CallerSession) Decimals() (uint8, error) ***REMOVED***
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
***REMOVED***

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() pure returns(string)
func (_Erc20 *Erc20Caller) Generator(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "generator")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() pure returns(string)
func (_Erc20 *Erc20Session) Generator() (string, error) ***REMOVED***
	return _Erc20.Contract.Generator(&_Erc20.CallOpts)
***REMOVED***

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() pure returns(string)
func (_Erc20 *Erc20CallerSession) Generator() (string, error) ***REMOVED***
	return _Erc20.Contract.Generator(&_Erc20.CallOpts)
***REMOVED***

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Erc20 *Erc20Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "getOwner")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Erc20 *Erc20Session) GetOwner() (common.Address, error) ***REMOVED***
	return _Erc20.Contract.GetOwner(&_Erc20.CallOpts)
***REMOVED***

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Erc20 *Erc20CallerSession) GetOwner() (common.Address, error) ***REMOVED***
	return _Erc20.Contract.GetOwner(&_Erc20.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) ***REMOVED***
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) ***REMOVED***
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "owner")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Session) Owner() (common.Address, error) ***REMOVED***
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20CallerSession) Owner() (common.Address, error) ***REMOVED***
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) ***REMOVED***
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) ***REMOVED***
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) ***REMOVED***
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) ***REMOVED***
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
***REMOVED***

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Erc20 *Erc20Caller) Version(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Erc20.contract.Call(opts, &out, "version")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Erc20 *Erc20Session) Version() (string, error) ***REMOVED***
	return _Erc20.Contract.Version(&_Erc20.CallOpts)
***REMOVED***

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Erc20 *Erc20CallerSession) Version() (string, error) ***REMOVED***
	return _Erc20.Contract.Version(&_Erc20.CallOpts)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.contract.Transact(opts, "approve", spender, amount)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
***REMOVED***

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20 *Erc20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
***REMOVED***

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20 *Erc20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.DecreaseAllowance(&_Erc20.TransactOpts, spender, subtractedValue)
***REMOVED***

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20 *Erc20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.DecreaseAllowance(&_Erc20.TransactOpts, spender, subtractedValue)
***REMOVED***

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20 *Erc20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
***REMOVED***

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20 *Erc20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.IncreaseAllowance(&_Erc20.TransactOpts, spender, addedValue)
***REMOVED***

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20 *Erc20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.IncreaseAllowance(&_Erc20.TransactOpts, spender, addedValue)
***REMOVED***

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc20 *Erc20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Erc20.contract.Transact(opts, "renounceOwnership")
***REMOVED***

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc20 *Erc20Session) RenounceOwnership() (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.RenounceOwnership(&_Erc20.TransactOpts)
***REMOVED***

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc20 *Erc20TransactorSession) RenounceOwnership() (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.RenounceOwnership(&_Erc20.TransactOpts)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.contract.Transact(opts, "transfer", recipient, amount)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _Erc20.contract.Transact(opts, "transferOwnership", newOwner)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
***REMOVED***

// Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.
type Erc20ApprovalIterator struct ***REMOVED***
	Event *Erc20Approval // Event containing the contract specifics and raw log

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
func (it *Erc20ApprovalIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(Erc20Approval)
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
		it.Event = new(Erc20Approval)
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
func (it *Erc20ApprovalIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ApprovalIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct ***REMOVED***
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Erc20ApprovalIterator***REMOVED***contract: _Erc20.contract, event: "Approval", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Approval)
				if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
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
func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) ***REMOVED***
	event := new(Erc20Approval)
	if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// Erc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc20 contract.
type Erc20OwnershipTransferredIterator struct ***REMOVED***
	Event *Erc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc20OwnershipTransferredIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(Erc20OwnershipTransferred)
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
		it.Event = new(Erc20OwnershipTransferred)
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
func (it *Erc20OwnershipTransferredIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20OwnershipTransferredIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// Erc20OwnershipTransferred represents a OwnershipTransferred event raised by the Erc20 contract.
type Erc20OwnershipTransferred struct ***REMOVED***
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc20OwnershipTransferredIterator, error) ***REMOVED***

	var previousOwnerRule []interface***REMOVED******REMOVED***
	for _, previousOwnerItem := range previousOwner ***REMOVED***
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	***REMOVED***
	var newOwnerRule []interface***REMOVED******REMOVED***
	for _, newOwnerItem := range newOwner ***REMOVED***
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	***REMOVED***

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Erc20OwnershipTransferredIterator***REMOVED***contract: _Erc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) ***REMOVED***

	var previousOwnerRule []interface***REMOVED******REMOVED***
	for _, previousOwnerItem := range previousOwner ***REMOVED***
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	***REMOVED***
	var newOwnerRule []interface***REMOVED******REMOVED***
	for _, newOwnerItem := range newOwner ***REMOVED***
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	***REMOVED***

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20OwnershipTransferred)
				if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil ***REMOVED***
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
func (_Erc20 *Erc20Filterer) ParseOwnershipTransferred(log types.Log) (*Erc20OwnershipTransferred, error) ***REMOVED***
	event := new(Erc20OwnershipTransferred)
	if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct ***REMOVED***
	Event *Erc20Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20TransferIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
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
		it.Event = new(Erc20Transfer)
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
func (it *Erc20TransferIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct ***REMOVED***
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Erc20TransferIterator***REMOVED***contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
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
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) ***REMOVED***
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***
