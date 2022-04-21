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
var UnifactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnifactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UnifactoryMetaData.ABI instead.
var UnifactoryABI = UnifactoryMetaData.ABI

// Unifactory is an auto generated Go binding around an Ethereum contract.
type Unifactory struct {
	UnifactoryCaller     // Read-only binding to the contract
	UnifactoryTransactor // Write-only binding to the contract
	UnifactoryFilterer   // Log filterer for contract events
}

// UnifactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnifactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnifactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnifactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnifactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnifactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnifactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnifactorySession struct {
	Contract     *Unifactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnifactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnifactoryCallerSession struct {
	Contract *UnifactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UnifactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnifactoryTransactorSession struct {
	Contract     *UnifactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnifactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnifactoryRaw struct {
	Contract *Unifactory // Generic contract binding to access the raw methods on
}

// UnifactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnifactoryCallerRaw struct {
	Contract *UnifactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UnifactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnifactoryTransactorRaw struct {
	Contract *UnifactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnifactory creates a new instance of Unifactory, bound to a specific deployed contract.
func NewUnifactory(address common.Address, backend bind.ContractBackend) (*Unifactory, error) {
	contract, err := bindUnifactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unifactory{UnifactoryCaller: UnifactoryCaller{contract: contract}, UnifactoryTransactor: UnifactoryTransactor{contract: contract}, UnifactoryFilterer: UnifactoryFilterer{contract: contract}}, nil
}

// NewUnifactoryCaller creates a new read-only instance of Unifactory, bound to a specific deployed contract.
func NewUnifactoryCaller(address common.Address, caller bind.ContractCaller) (*UnifactoryCaller, error) {
	contract, err := bindUnifactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnifactoryCaller{contract: contract}, nil
}

// NewUnifactoryTransactor creates a new write-only instance of Unifactory, bound to a specific deployed contract.
func NewUnifactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UnifactoryTransactor, error) {
	contract, err := bindUnifactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnifactoryTransactor{contract: contract}, nil
}

// NewUnifactoryFilterer creates a new log filterer instance of Unifactory, bound to a specific deployed contract.
func NewUnifactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UnifactoryFilterer, error) {
	contract, err := bindUnifactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnifactoryFilterer{contract: contract}, nil
}

// bindUnifactory binds a generic wrapper to an already deployed contract.
func bindUnifactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnifactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unifactory *UnifactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unifactory.Contract.UnifactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unifactory *UnifactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unifactory.Contract.UnifactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unifactory *UnifactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unifactory.Contract.UnifactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unifactory *UnifactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unifactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unifactory *UnifactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unifactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unifactory *UnifactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unifactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Unifactory *UnifactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Unifactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Unifactory *UnifactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Unifactory.Contract.AllPairs(&_Unifactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Unifactory *UnifactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Unifactory.Contract.AllPairs(&_Unifactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Unifactory *UnifactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unifactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Unifactory *UnifactorySession) AllPairsLength() (*big.Int, error) {
	return _Unifactory.Contract.AllPairsLength(&_Unifactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Unifactory *UnifactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _Unifactory.Contract.AllPairsLength(&_Unifactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Unifactory *UnifactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unifactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Unifactory *UnifactorySession) FeeTo() (common.Address, error) {
	return _Unifactory.Contract.FeeTo(&_Unifactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Unifactory *UnifactoryCallerSession) FeeTo() (common.Address, error) {
	return _Unifactory.Contract.FeeTo(&_Unifactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Unifactory *UnifactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unifactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Unifactory *UnifactorySession) FeeToSetter() (common.Address, error) {
	return _Unifactory.Contract.FeeToSetter(&_Unifactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Unifactory *UnifactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _Unifactory.Contract.FeeToSetter(&_Unifactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Unifactory *UnifactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Unifactory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Unifactory *UnifactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Unifactory.Contract.GetPair(&_Unifactory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Unifactory *UnifactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Unifactory.Contract.GetPair(&_Unifactory.CallOpts, arg0, arg1)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Unifactory *UnifactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Unifactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Unifactory *UnifactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Unifactory.Contract.CreatePair(&_Unifactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Unifactory *UnifactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Unifactory.Contract.CreatePair(&_Unifactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Unifactory *UnifactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Unifactory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Unifactory *UnifactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Unifactory.Contract.SetFeeTo(&_Unifactory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Unifactory *UnifactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Unifactory.Contract.SetFeeTo(&_Unifactory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Unifactory *UnifactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _Unifactory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Unifactory *UnifactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Unifactory.Contract.SetFeeToSetter(&_Unifactory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Unifactory *UnifactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Unifactory.Contract.SetFeeToSetter(&_Unifactory.TransactOpts, _feeToSetter)
}

// UnifactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Unifactory contract.
type UnifactoryPairCreatedIterator struct {
	Event *UnifactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *UnifactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnifactoryPairCreated)
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
		it.Event = new(UnifactoryPairCreated)
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
func (it *UnifactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnifactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnifactoryPairCreated represents a PairCreated event raised by the Unifactory contract.
type UnifactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Unifactory *UnifactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*UnifactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Unifactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &UnifactoryPairCreatedIterator{contract: _Unifactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Unifactory *UnifactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *UnifactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Unifactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnifactoryPairCreated)
				if err := _Unifactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Unifactory *UnifactoryFilterer) ParsePairCreated(log types.Log) (*UnifactoryPairCreated, error) {
	event := new(UnifactoryPairCreated)
	if err := _Unifactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
