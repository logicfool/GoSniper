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

// UnifactoryMetaData contains all meta data concerning the Unifactory contract.
var UnifactoryMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"name\":\"PairCreated\",\"type\":\"event\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"name\":\"allPairs\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"***REMOVED***],\"name\":\"createPair\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":true,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"name\":\"getPair\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"***REMOVED***],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"constant\":false,\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"***REMOVED***],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***]",
***REMOVED***

// UnifactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UnifactoryMetaData.ABI instead.
var UnifactoryABI = UnifactoryMetaData.ABI

// Unifactory is an auto generated Go binding around an Ethereum contract.
type Unifactory struct ***REMOVED***
	UnifactoryCaller     // Read-only binding to the contract
	UnifactoryTransactor // Write-only binding to the contract
	UnifactoryFilterer   // Log filterer for contract events
***REMOVED***

// UnifactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnifactoryCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// UnifactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnifactoryTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// UnifactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnifactoryFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// UnifactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnifactorySession struct ***REMOVED***
	Contract     *Unifactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// UnifactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnifactoryCallerSession struct ***REMOVED***
	Contract *UnifactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
***REMOVED***

// UnifactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnifactoryTransactorSession struct ***REMOVED***
	Contract     *UnifactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
***REMOVED***

// UnifactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnifactoryRaw struct ***REMOVED***
	Contract *Unifactory // Generic contract binding to access the raw methods on
***REMOVED***

// UnifactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnifactoryCallerRaw struct ***REMOVED***
	Contract *UnifactoryCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// UnifactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnifactoryTransactorRaw struct ***REMOVED***
	Contract *UnifactoryTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewUnifactory creates a new instance of Unifactory, bound to a specific deployed contract.
func NewUnifactory(address common.Address, backend bind.ContractBackend) (*Unifactory, error) ***REMOVED***
	contract, err := bindUnifactory(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Unifactory***REMOVED***UnifactoryCaller: UnifactoryCaller***REMOVED***contract: contract***REMOVED***, UnifactoryTransactor: UnifactoryTransactor***REMOVED***contract: contract***REMOVED***, UnifactoryFilterer: UnifactoryFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewUnifactoryCaller creates a new read-only instance of Unifactory, bound to a specific deployed contract.
func NewUnifactoryCaller(address common.Address, caller bind.ContractCaller) (*UnifactoryCaller, error) ***REMOVED***
	contract, err := bindUnifactory(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &UnifactoryCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewUnifactoryTransactor creates a new write-only instance of Unifactory, bound to a specific deployed contract.
func NewUnifactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UnifactoryTransactor, error) ***REMOVED***
	contract, err := bindUnifactory(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &UnifactoryTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewUnifactoryFilterer creates a new log filterer instance of Unifactory, bound to a specific deployed contract.
func NewUnifactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UnifactoryFilterer, error) ***REMOVED***
	contract, err := bindUnifactory(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &UnifactoryFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindUnifactory binds a generic wrapper to an already deployed contract.
func bindUnifactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(UnifactoryABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unifactory *UnifactoryRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Unifactory.Contract.UnifactoryCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unifactory *UnifactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.UnifactoryTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Unifactory *UnifactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.UnifactoryTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unifactory *UnifactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Unifactory.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unifactory *UnifactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Unifactory *UnifactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Unifactory *UnifactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unifactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Unifactory *UnifactorySession) AllPairs(arg0 *big.Int) (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.AllPairs(&_Unifactory.CallOpts, arg0)
***REMOVED***

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Unifactory *UnifactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.AllPairs(&_Unifactory.CallOpts, arg0)
***REMOVED***

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Unifactory *UnifactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unifactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Unifactory *UnifactorySession) AllPairsLength() (*big.Int, error) ***REMOVED***
	return _Unifactory.Contract.AllPairsLength(&_Unifactory.CallOpts)
***REMOVED***

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Unifactory *UnifactoryCallerSession) AllPairsLength() (*big.Int, error) ***REMOVED***
	return _Unifactory.Contract.AllPairsLength(&_Unifactory.CallOpts)
***REMOVED***

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Unifactory *UnifactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unifactory.contract.Call(opts, &out, "feeTo")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Unifactory *UnifactorySession) FeeTo() (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.FeeTo(&_Unifactory.CallOpts)
***REMOVED***

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Unifactory *UnifactoryCallerSession) FeeTo() (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.FeeTo(&_Unifactory.CallOpts)
***REMOVED***

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Unifactory *UnifactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unifactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Unifactory *UnifactorySession) FeeToSetter() (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.FeeToSetter(&_Unifactory.CallOpts)
***REMOVED***

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Unifactory *UnifactoryCallerSession) FeeToSetter() (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.FeeToSetter(&_Unifactory.CallOpts)
***REMOVED***

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Unifactory *UnifactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unifactory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Unifactory *UnifactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.GetPair(&_Unifactory.CallOpts, arg0, arg1)
***REMOVED***

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Unifactory *UnifactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) ***REMOVED***
	return _Unifactory.Contract.GetPair(&_Unifactory.CallOpts, arg0, arg1)
***REMOVED***

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Unifactory *UnifactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.contract.Transact(opts, "createPair", tokenA, tokenB)
***REMOVED***

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Unifactory *UnifactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.CreatePair(&_Unifactory.TransactOpts, tokenA, tokenB)
***REMOVED***

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Unifactory *UnifactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.CreatePair(&_Unifactory.TransactOpts, tokenA, tokenB)
***REMOVED***

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Unifactory *UnifactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.contract.Transact(opts, "setFeeTo", _feeTo)
***REMOVED***

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Unifactory *UnifactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.SetFeeTo(&_Unifactory.TransactOpts, _feeTo)
***REMOVED***

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Unifactory *UnifactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.SetFeeTo(&_Unifactory.TransactOpts, _feeTo)
***REMOVED***

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Unifactory *UnifactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
***REMOVED***

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Unifactory *UnifactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.SetFeeToSetter(&_Unifactory.TransactOpts, _feeToSetter)
***REMOVED***

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Unifactory *UnifactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) ***REMOVED***
	return _Unifactory.Contract.SetFeeToSetter(&_Unifactory.TransactOpts, _feeToSetter)
***REMOVED***

// UnifactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Unifactory contract.
type UnifactoryPairCreatedIterator struct ***REMOVED***
	Event *UnifactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *UnifactoryPairCreatedIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(UnifactoryPairCreated)
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
		it.Event = new(UnifactoryPairCreated)
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
func (it *UnifactoryPairCreatedIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnifactoryPairCreatedIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// UnifactoryPairCreated represents a PairCreated event raised by the Unifactory contract.
type UnifactoryPairCreated struct ***REMOVED***
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Unifactory *UnifactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*UnifactoryPairCreatedIterator, error) ***REMOVED***

	var token0Rule []interface***REMOVED******REMOVED***
	for _, token0Item := range token0 ***REMOVED***
		token0Rule = append(token0Rule, token0Item)
	***REMOVED***
	var token1Rule []interface***REMOVED******REMOVED***
	for _, token1Item := range token1 ***REMOVED***
		token1Rule = append(token1Rule, token1Item)
	***REMOVED***

	logs, sub, err := _Unifactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &UnifactoryPairCreatedIterator***REMOVED***contract: _Unifactory.contract, event: "PairCreated", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Unifactory *UnifactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *UnifactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) ***REMOVED***

	var token0Rule []interface***REMOVED******REMOVED***
	for _, token0Item := range token0 ***REMOVED***
		token0Rule = append(token0Rule, token0Item)
	***REMOVED***
	var token1Rule []interface***REMOVED******REMOVED***
	for _, token1Item := range token1 ***REMOVED***
		token1Rule = append(token1Rule, token1Item)
	***REMOVED***

	logs, sub, err := _Unifactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnifactoryPairCreated)
				if err := _Unifactory.contract.UnpackLog(event, "PairCreated", log); err != nil ***REMOVED***
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Unifactory *UnifactoryFilterer) ParsePairCreated(log types.Log) (*UnifactoryPairCreated, error) ***REMOVED***
	event := new(UnifactoryPairCreated)
	if err := _Unifactory.contract.UnpackLog(event, "PairCreated", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***
