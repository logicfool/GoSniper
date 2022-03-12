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

// PairMetaData contains all meta data concerning the Pair contract.
var PairMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Approval\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***],\"name\":\"Burn\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Mint\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***],\"name\":\"Swap\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"***REMOVED***],\"name\":\"Sync\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Transfer\",\"type\":\"event\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[***REMOVED***\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[***REMOVED***\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"name\":\"allowance\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"approve\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"name\":\"balanceOf\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***],\"name\":\"burn\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[***REMOVED***\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[***REMOVED***\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"***REMOVED***,***REMOVED***\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"***REMOVED***,***REMOVED***\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"***REMOVED***],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"kLast\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***],\"name\":\"mint\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"name\":\"nonces\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"***REMOVED***],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"***REMOVED***],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transfer\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transferFrom\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***]",
***REMOVED***

// PairABI is the input ABI used to generate the binding from.
// Deprecated: Use PairMetaData.ABI instead.
var PairABI = PairMetaData.ABI

// Pair is an auto generated Go binding around an Ethereum contract.
type Pair struct ***REMOVED***
	PairCaller     // Read-only binding to the contract
	PairTransactor // Write-only binding to the contract
	PairFilterer   // Log filterer for contract events
***REMOVED***

// PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairSession struct ***REMOVED***
	Contract     *Pair             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairCallerSession struct ***REMOVED***
	Contract *PairCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
***REMOVED***

// PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairTransactorSession struct ***REMOVED***
	Contract     *PairTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairRaw struct ***REMOVED***
	Contract *Pair // Generic contract binding to access the raw methods on
***REMOVED***

// PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairCallerRaw struct ***REMOVED***
	Contract *PairCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairTransactorRaw struct ***REMOVED***
	Contract *PairTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewPair creates a new instance of Pair, bound to a specific deployed contract.
func NewPair(address common.Address, backend bind.ContractBackend) (*Pair, error) ***REMOVED***
	contract, err := bindPair(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Pair***REMOVED***PairCaller: PairCaller***REMOVED***contract: contract***REMOVED***, PairTransactor: PairTransactor***REMOVED***contract: contract***REMOVED***, PairFilterer: PairFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewPairCaller creates a new read-only instance of Pair, bound to a specific deployed contract.
func NewPairCaller(address common.Address, caller bind.ContractCaller) (*PairCaller, error) ***REMOVED***
	contract, err := bindPair(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewPairTransactor creates a new write-only instance of Pair, bound to a specific deployed contract.
func NewPairTransactor(address common.Address, transactor bind.ContractTransactor) (*PairTransactor, error) ***REMOVED***
	contract, err := bindPair(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewPairFilterer creates a new log filterer instance of Pair, bound to a specific deployed contract.
func NewPairFilterer(address common.Address, filterer bind.ContractFilterer) (*PairFilterer, error) ***REMOVED***
	contract, err := bindPair(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindPair binds a generic wrapper to an already deployed contract.
func bindPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(PairABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Pair.Contract.PairCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.PairTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.PairTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Pair.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pair *PairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil ***REMOVED***
		return *new([32]byte), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

***REMOVED***

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pair *PairSession) DOMAINSEPARATOR() ([32]byte, error) ***REMOVED***
	return _Pair.Contract.DOMAINSEPARATOR(&_Pair.CallOpts)
***REMOVED***

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pair *PairCallerSession) DOMAINSEPARATOR() ([32]byte, error) ***REMOVED***
	return _Pair.Contract.DOMAINSEPARATOR(&_Pair.CallOpts)
***REMOVED***

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pair *PairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pair *PairSession) MINIMUMLIQUIDITY() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.MINIMUMLIQUIDITY(&_Pair.CallOpts)
***REMOVED***

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pair *PairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.MINIMUMLIQUIDITY(&_Pair.CallOpts)
***REMOVED***

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pair *PairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil ***REMOVED***
		return *new([32]byte), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

***REMOVED***

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pair *PairSession) PERMITTYPEHASH() ([32]byte, error) ***REMOVED***
	return _Pair.Contract.PERMITTYPEHASH(&_Pair.CallOpts)
***REMOVED***

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pair *PairCallerSession) PERMITTYPEHASH() ([32]byte, error) ***REMOVED***
	return _Pair.Contract.PERMITTYPEHASH(&_Pair.CallOpts)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pair *PairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pair *PairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Allowance(&_Pair.CallOpts, arg0, arg1)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pair *PairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Allowance(&_Pair.CallOpts, arg0, arg1)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pair *PairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pair *PairSession) BalanceOf(arg0 common.Address) (*big.Int, error) ***REMOVED***
	return _Pair.Contract.BalanceOf(&_Pair.CallOpts, arg0)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pair *PairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) ***REMOVED***
	return _Pair.Contract.BalanceOf(&_Pair.CallOpts, arg0)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pair *PairCaller) Decimals(opts *bind.CallOpts) (uint8, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "decimals")

	if err != nil ***REMOVED***
		return *new(uint8), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pair *PairSession) Decimals() (uint8, error) ***REMOVED***
	return _Pair.Contract.Decimals(&_Pair.CallOpts)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pair *PairCallerSession) Decimals() (uint8, error) ***REMOVED***
	return _Pair.Contract.Decimals(&_Pair.CallOpts)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pair *PairCaller) Factory(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "factory")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pair *PairSession) Factory() (common.Address, error) ***REMOVED***
	return _Pair.Contract.Factory(&_Pair.CallOpts)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pair *PairCallerSession) Factory() (common.Address, error) ***REMOVED***
	return _Pair.Contract.Factory(&_Pair.CallOpts)
***REMOVED***

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pair *PairCaller) GetReserves(opts *bind.CallOpts) (struct ***REMOVED***
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
***REMOVED***, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct ***REMOVED***
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	***REMOVED***)
	if err != nil ***REMOVED***
		return *outstruct, err
	***REMOVED***

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

***REMOVED***

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pair *PairSession) GetReserves() (struct ***REMOVED***
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
***REMOVED***, error) ***REMOVED***
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
***REMOVED***

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pair *PairCallerSession) GetReserves() (struct ***REMOVED***
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
***REMOVED***, error) ***REMOVED***
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
***REMOVED***

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pair *PairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "kLast")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pair *PairSession) KLast() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.KLast(&_Pair.CallOpts)
***REMOVED***

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pair *PairCallerSession) KLast() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.KLast(&_Pair.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pair *PairCaller) Name(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "name")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pair *PairSession) Name() (string, error) ***REMOVED***
	return _Pair.Contract.Name(&_Pair.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pair *PairCallerSession) Name() (string, error) ***REMOVED***
	return _Pair.Contract.Name(&_Pair.CallOpts)
***REMOVED***

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pair *PairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pair *PairSession) Nonces(arg0 common.Address) (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Nonces(&_Pair.CallOpts, arg0)
***REMOVED***

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pair *PairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Nonces(&_Pair.CallOpts, arg0)
***REMOVED***

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pair *PairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pair *PairSession) Price0CumulativeLast() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Price0CumulativeLast(&_Pair.CallOpts)
***REMOVED***

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pair *PairCallerSession) Price0CumulativeLast() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Price0CumulativeLast(&_Pair.CallOpts)
***REMOVED***

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pair *PairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pair *PairSession) Price1CumulativeLast() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Price1CumulativeLast(&_Pair.CallOpts)
***REMOVED***

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pair *PairCallerSession) Price1CumulativeLast() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.Price1CumulativeLast(&_Pair.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pair *PairCaller) Symbol(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "symbol")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pair *PairSession) Symbol() (string, error) ***REMOVED***
	return _Pair.Contract.Symbol(&_Pair.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pair *PairCallerSession) Symbol() (string, error) ***REMOVED***
	return _Pair.Contract.Symbol(&_Pair.CallOpts)
***REMOVED***

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCaller) Token0(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "token0")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairSession) Token0() (common.Address, error) ***REMOVED***
	return _Pair.Contract.Token0(&_Pair.CallOpts)
***REMOVED***

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCallerSession) Token0() (common.Address, error) ***REMOVED***
	return _Pair.Contract.Token0(&_Pair.CallOpts)
***REMOVED***

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCaller) Token1(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "token1")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairSession) Token1() (common.Address, error) ***REMOVED***
	return _Pair.Contract.Token1(&_Pair.CallOpts)
***REMOVED***

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCallerSession) Token1() (common.Address, error) ***REMOVED***
	return _Pair.Contract.Token1(&_Pair.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Pair.contract.Call(opts, &out, "totalSupply")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairSession) TotalSupply() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.TotalSupply(&_Pair.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairCallerSession) TotalSupply() (*big.Int, error) ***REMOVED***
	return _Pair.Contract.TotalSupply(&_Pair.CallOpts)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pair *PairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "approve", spender, value)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pair *PairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Approve(&_Pair.TransactOpts, spender, value)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pair *PairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Approve(&_Pair.TransactOpts, spender, value)
***REMOVED***

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "burn", to)
***REMOVED***

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairSession) Burn(to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
***REMOVED***

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactorSession) Burn(to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
***REMOVED***

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pair *PairTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "initialize", _token0, _token1)
***REMOVED***

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pair *PairSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Initialize(&_Pair.TransactOpts, _token0, _token1)
***REMOVED***

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pair *PairTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Initialize(&_Pair.TransactOpts, _token0, _token1)
***REMOVED***

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "mint", to)
***REMOVED***

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairSession) Mint(to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
***REMOVED***

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactorSession) Mint(to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
***REMOVED***

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
***REMOVED***

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
***REMOVED***

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
***REMOVED***

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "skim", to)
***REMOVED***

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairSession) Skim(to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Skim(&_Pair.TransactOpts, to)
***REMOVED***

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairTransactorSession) Skim(to common.Address) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Skim(&_Pair.TransactOpts, to)
***REMOVED***

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
***REMOVED***

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
***REMOVED***

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
***REMOVED***

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "sync")
***REMOVED***

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairSession) Sync() (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Sync(&_Pair.TransactOpts)
***REMOVED***

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairTransactorSession) Sync() (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Sync(&_Pair.TransactOpts)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pair *PairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "transfer", to, value)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pair *PairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Transfer(&_Pair.TransactOpts, to, value)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pair *PairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.Transfer(&_Pair.TransactOpts, to, value)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pair *PairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.contract.Transact(opts, "transferFrom", from, to, value)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pair *PairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, from, to, value)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pair *PairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, from, to, value)
***REMOVED***

// PairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pair contract.
type PairApprovalIterator struct ***REMOVED***
	Event *PairApproval // Event containing the contract specifics and raw log

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
func (it *PairApprovalIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(PairApproval)
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
		it.Event = new(PairApproval)
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
func (it *PairApprovalIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairApprovalIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// PairApproval represents a Approval event raised by the Pair contract.
type PairApproval struct ***REMOVED***
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pair *PairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PairApprovalIterator, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairApprovalIterator***REMOVED***contract: _Pair.contract, event: "Approval", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pair *PairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairApproval)
				if err := _Pair.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
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
func (_Pair *PairFilterer) ParseApproval(log types.Log) (*PairApproval, error) ***REMOVED***
	event := new(PairApproval)
	if err := _Pair.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// PairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pair contract.
type PairBurnIterator struct ***REMOVED***
	Event *PairBurn // Event containing the contract specifics and raw log

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
func (it *PairBurnIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(PairBurn)
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
		it.Event = new(PairBurn)
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
func (it *PairBurnIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairBurnIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// PairBurn represents a Burn event raised by the Pair contract.
type PairBurn struct ***REMOVED***
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairBurnIterator, error) ***REMOVED***

	var senderRule []interface***REMOVED******REMOVED***
	for _, senderItem := range sender ***REMOVED***
		senderRule = append(senderRule, senderItem)
	***REMOVED***

	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairBurnIterator***REMOVED***contract: _Pair.contract, event: "Burn", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) ***REMOVED***

	var senderRule []interface***REMOVED******REMOVED***
	for _, senderItem := range sender ***REMOVED***
		senderRule = append(senderRule, senderItem)
	***REMOVED***

	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairBurn)
				if err := _Pair.contract.UnpackLog(event, "Burn", log); err != nil ***REMOVED***
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) ParseBurn(log types.Log) (*PairBurn, error) ***REMOVED***
	event := new(PairBurn)
	if err := _Pair.contract.UnpackLog(event, "Burn", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// PairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pair contract.
type PairMintIterator struct ***REMOVED***
	Event *PairMint // Event containing the contract specifics and raw log

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
func (it *PairMintIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(PairMint)
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
		it.Event = new(PairMint)
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
func (it *PairMintIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairMintIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// PairMint represents a Mint event raised by the Pair contract.
type PairMint struct ***REMOVED***
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*PairMintIterator, error) ***REMOVED***

	var senderRule []interface***REMOVED******REMOVED***
	for _, senderItem := range sender ***REMOVED***
		senderRule = append(senderRule, senderItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairMintIterator***REMOVED***contract: _Pair.contract, event: "Mint", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PairMint, sender []common.Address) (event.Subscription, error) ***REMOVED***

	var senderRule []interface***REMOVED******REMOVED***
	for _, senderItem := range sender ***REMOVED***
		senderRule = append(senderRule, senderItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairMint)
				if err := _Pair.contract.UnpackLog(event, "Mint", log); err != nil ***REMOVED***
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) ParseMint(log types.Log) (*PairMint, error) ***REMOVED***
	event := new(PairMint)
	if err := _Pair.contract.UnpackLog(event, "Mint", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// PairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pair contract.
type PairSwapIterator struct ***REMOVED***
	Event *PairSwap // Event containing the contract specifics and raw log

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
func (it *PairSwapIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(PairSwap)
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
		it.Event = new(PairSwap)
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
func (it *PairSwapIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairSwapIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// PairSwap represents a Swap event raised by the Pair contract.
type PairSwap struct ***REMOVED***
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairSwapIterator, error) ***REMOVED***

	var senderRule []interface***REMOVED******REMOVED***
	for _, senderItem := range sender ***REMOVED***
		senderRule = append(senderRule, senderItem)
	***REMOVED***

	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairSwapIterator***REMOVED***contract: _Pair.contract, event: "Swap", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) ***REMOVED***

	var senderRule []interface***REMOVED******REMOVED***
	for _, senderItem := range sender ***REMOVED***
		senderRule = append(senderRule, senderItem)
	***REMOVED***

	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairSwap)
				if err := _Pair.contract.UnpackLog(event, "Swap", log); err != nil ***REMOVED***
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) ParseSwap(log types.Log) (*PairSwap, error) ***REMOVED***
	event := new(PairSwap)
	if err := _Pair.contract.UnpackLog(event, "Swap", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// PairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Pair contract.
type PairSyncIterator struct ***REMOVED***
	Event *PairSync // Event containing the contract specifics and raw log

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
func (it *PairSyncIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(PairSync)
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
		it.Event = new(PairSync)
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
func (it *PairSyncIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairSyncIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// PairSync represents a Sync event raised by the Pair contract.
type PairSync struct ***REMOVED***
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pair *PairFilterer) FilterSync(opts *bind.FilterOpts) (*PairSyncIterator, error) ***REMOVED***

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Sync")
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairSyncIterator***REMOVED***contract: _Pair.contract, event: "Sync", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pair *PairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *PairSync) (event.Subscription, error) ***REMOVED***

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Sync")
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairSync)
				if err := _Pair.contract.UnpackLog(event, "Sync", log); err != nil ***REMOVED***
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pair *PairFilterer) ParseSync(log types.Log) (*PairSync, error) ***REMOVED***
	event := new(PairSync)
	if err := _Pair.contract.UnpackLog(event, "Sync", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// PairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pair contract.
type PairTransferIterator struct ***REMOVED***
	Event *PairTransfer // Event containing the contract specifics and raw log

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
func (it *PairTransferIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(PairTransfer)
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
		it.Event = new(PairTransfer)
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
func (it *PairTransferIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairTransferIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// PairTransfer represents a Transfer event raised by the Pair contract.
type PairTransfer struct ***REMOVED***
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pair *PairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PairTransferIterator, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &PairTransferIterator***REMOVED***contract: _Pair.contract, event: "Transfer", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pair *PairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairTransfer)
				if err := _Pair.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
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
func (_Pair *PairFilterer) ParseTransfer(log types.Log) (*PairTransfer, error) ***REMOVED***
	event := new(PairTransfer)
	if err := _Pair.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***
