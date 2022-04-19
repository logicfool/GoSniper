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

// ChiGasTokenMetaData contains all meta data concerning the ChiGasToken contract.
var ChiGasTokenMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"free\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"freeFrom\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"freeFromUpTo\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"freeUpTo\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***]",
	Sigs: map[string]string***REMOVED***
		"d8ccd0f3": "free(uint256)",
		"5f2e2b45": "freeFrom(address,uint256)",
		"079d229f": "freeFromUpTo(address,uint256)",
		"6366b936": "freeUpTo(uint256)",
	***REMOVED***,
***REMOVED***

// ChiGasTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ChiGasTokenMetaData.ABI instead.
var ChiGasTokenABI = ChiGasTokenMetaData.ABI

// Deprecated: Use ChiGasTokenMetaData.Sigs instead.
// ChiGasTokenFuncSigs maps the 4-byte function signature to its string representation.
var ChiGasTokenFuncSigs = ChiGasTokenMetaData.Sigs

// ChiGasToken is an auto generated Go binding around an Ethereum contract.
type ChiGasToken struct ***REMOVED***
	ChiGasTokenCaller     // Read-only binding to the contract
	ChiGasTokenTransactor // Write-only binding to the contract
	ChiGasTokenFilterer   // Log filterer for contract events
***REMOVED***

// ChiGasTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChiGasTokenCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ChiGasTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChiGasTokenTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ChiGasTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChiGasTokenFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ChiGasTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChiGasTokenSession struct ***REMOVED***
	Contract     *ChiGasToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// ChiGasTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChiGasTokenCallerSession struct ***REMOVED***
	Contract *ChiGasTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
***REMOVED***

// ChiGasTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChiGasTokenTransactorSession struct ***REMOVED***
	Contract     *ChiGasTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
***REMOVED***

// ChiGasTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChiGasTokenRaw struct ***REMOVED***
	Contract *ChiGasToken // Generic contract binding to access the raw methods on
***REMOVED***

// ChiGasTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChiGasTokenCallerRaw struct ***REMOVED***
	Contract *ChiGasTokenCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// ChiGasTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChiGasTokenTransactorRaw struct ***REMOVED***
	Contract *ChiGasTokenTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewChiGasToken creates a new instance of ChiGasToken, bound to a specific deployed contract.
func NewChiGasToken(address common.Address, backend bind.ContractBackend) (*ChiGasToken, error) ***REMOVED***
	contract, err := bindChiGasToken(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ChiGasToken***REMOVED***ChiGasTokenCaller: ChiGasTokenCaller***REMOVED***contract: contract***REMOVED***, ChiGasTokenTransactor: ChiGasTokenTransactor***REMOVED***contract: contract***REMOVED***, ChiGasTokenFilterer: ChiGasTokenFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewChiGasTokenCaller creates a new read-only instance of ChiGasToken, bound to a specific deployed contract.
func NewChiGasTokenCaller(address common.Address, caller bind.ContractCaller) (*ChiGasTokenCaller, error) ***REMOVED***
	contract, err := bindChiGasToken(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ChiGasTokenCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewChiGasTokenTransactor creates a new write-only instance of ChiGasToken, bound to a specific deployed contract.
func NewChiGasTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ChiGasTokenTransactor, error) ***REMOVED***
	contract, err := bindChiGasToken(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ChiGasTokenTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewChiGasTokenFilterer creates a new log filterer instance of ChiGasToken, bound to a specific deployed contract.
func NewChiGasTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ChiGasTokenFilterer, error) ***REMOVED***
	contract, err := bindChiGasToken(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ChiGasTokenFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindChiGasToken binds a generic wrapper to an already deployed contract.
func bindChiGasToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(ChiGasTokenABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChiGasToken *ChiGasTokenRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _ChiGasToken.Contract.ChiGasTokenCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChiGasToken *ChiGasTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.ChiGasTokenTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_ChiGasToken *ChiGasTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.ChiGasTokenTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChiGasToken *ChiGasTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _ChiGasToken.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChiGasToken *ChiGasTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_ChiGasToken *ChiGasTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactor) Free(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.contract.Transact(opts, "free", value)
***REMOVED***

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenSession) Free(value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.Free(&_ChiGasToken.TransactOpts, value)
***REMOVED***

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactorSession) Free(value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.Free(&_ChiGasToken.TransactOpts, value)
***REMOVED***

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactor) FreeFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.contract.Transact(opts, "freeFrom", from, value)
***REMOVED***

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenSession) FreeFrom(from common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.FreeFrom(&_ChiGasToken.TransactOpts, from, value)
***REMOVED***

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactorSession) FreeFrom(from common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.FreeFrom(&_ChiGasToken.TransactOpts, from, value)
***REMOVED***

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactor) FreeFromUpTo(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.contract.Transact(opts, "freeFromUpTo", from, value)
***REMOVED***

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenSession) FreeFromUpTo(from common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.FreeFromUpTo(&_ChiGasToken.TransactOpts, from, value)
***REMOVED***

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactorSession) FreeFromUpTo(from common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.FreeFromUpTo(&_ChiGasToken.TransactOpts, from, value)
***REMOVED***

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactor) FreeUpTo(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.contract.Transact(opts, "freeUpTo", value)
***REMOVED***

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenSession) FreeUpTo(value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.FreeUpTo(&_ChiGasToken.TransactOpts, value)
***REMOVED***

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_ChiGasToken *ChiGasTokenTransactorSession) FreeUpTo(value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _ChiGasToken.Contract.FreeUpTo(&_ChiGasToken.TransactOpts, value)
***REMOVED***

// GoSniperMetaData contains all meta data concerning the GoSniper contract.
var GoSniperMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"_chi\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"***REMOVED***,***REMOVED***\"stateMutability\":\"payable\",\"type\":\"fallback\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"name\":\"Allowed\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"pricefeed\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address[][]\",\"name\":\"all_pairs_to_get\",\"type\":\"address[][]\"***REMOVED***],\"name\":\"GetPrices\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"ETHValue\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"tok_addy\",\"type\":\"address\"***REMOVED***],\"name\":\"RescueBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"bool\",\"name\":\"_res\",\"type\":\"bool\"***REMOVED***],\"name\":\"SetAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"bool\",\"name\":\"_res\",\"type\":\"bool\"***REMOVED***],\"name\":\"SetAllowedBulk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address[]\",\"name\":\"majoraddresses\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"tokenamount\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_amountoutMIN\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"beforebal\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"extrasettings\",\"type\":\"uint256[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_split\",\"type\":\"uint256\"***REMOVED***],\"name\":\"TradeDirectlyByPair\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address[]\",\"name\":\"majoraddresses\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_amountoutMIN\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"beforebal\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"extrasettings\",\"type\":\"uint256[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_split\",\"type\":\"uint256\"***REMOVED***],\"name\":\"TradeDirectlyByPairWithETH\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address[]\",\"name\":\"majoraddresses\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_amountoutMIN\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"beforebal\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"extrasettings\",\"type\":\"uint256[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_split\",\"type\":\"uint256\"***REMOVED***],\"name\":\"TradeDirectlyByPairWithETHGasRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address[]\",\"name\":\"majoraddresses\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"tokenamount\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_amountoutMIN\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"beforebal\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"extrasettings\",\"type\":\"uint256[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"_split\",\"type\":\"uint256\"***REMOVED***],\"name\":\"TradeDirectlyByPairWithGasRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"tokenamount\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"***REMOVED***],\"name\":\"UnsafeSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"WETH\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"chi\",\"outputs\":[***REMOVED***\"internalType\":\"contractChiGasToken\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"owner\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_chi\",\"type\":\"address\"***REMOVED***],\"name\":\"setCHI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"***REMOVED***],\"name\":\"setWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"***REMOVED***],\"name\":\"sortTokens\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"***REMOVED***],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"stateMutability\":\"payable\",\"type\":\"receive\"***REMOVED***]",
	Sigs: map[string]string***REMOVED***
		"77a7dbc6": "Allowed(address)",
		"1741d40e": "GetPrices(address,address,address[][])",
		"791b45d9": "RescueBalance(address)",
		"91d426a6": "SetAllowed(address,bool)",
		"063774f7": "SetAllowedBulk(address[],bool)",
		"3541c60f": "TradeDirectlyByPair(address[],address[],uint256,address,uint256,uint256,uint256[],uint256)",
		"ea952c6a": "TradeDirectlyByPairWithETH(address[],address[],address,uint256,uint256,uint256[],uint256)",
		"17960dce": "TradeDirectlyByPairWithETHGasRefund(address[],address[],address,uint256,uint256,uint256[],uint256)",
		"3e005d76": "TradeDirectlyByPairWithGasRefund(address[],address[],uint256,address,uint256,uint256,uint256[],uint256)",
		"568a0432": "UnsafeSwap(address,address[],uint256,uint256,uint256)",
		"ad5c4648": "WETH()",
		"c92aecc4": "chi()",
		"8da5cb5b": "owner()",
		"a73448c9": "setCHI(address)",
		"5b769f3c": "setWETH(address)",
		"544caa56": "sortTokens(address,address)",
		"f2fde38b": "transferOwnership(address)",
	***REMOVED***,
	Bin: "0x60806040523480156200001157600080fd5b506040516200306138038062003061833981016040819052620000349162000071565b600080546001600160a01b03199081163317909155600280546001600160a01b0394851690831617905560018054929093169116179055620000c8565b6000806040838503121562000084578182fd5b82516200009181620000af565b6020840151909250620000a481620000af565b809150509250929050565b6001600160a01b0381168114620000c557600080fd5b50565b612f8980620000d86000396000f3fe6080604052600436106101025760003560e01c806377a7dbc611610095578063a73448c911610064578063a73448c91461028b578063ad5c4648146102ab578063c92aecc4146102c0578063ea952c6a146102d5578063f2fde38b146102e857610109565b806377a7dbc614610216578063791b45d9146102365780638da5cb5b1461024957806391d426a61461026b57610109565b80633e005d76116100d15780633e005d7614610195578063544caa56146101a8578063568a0432146101d65780635b769f3c146101f657610109565b8063063774f71461010b5780631741d40e1461012b57806317960dce146101625780633541c60f1461017557610109565b3661010957005b005b34801561011757600080fd5b506101096101263660046128b3565b610308565b34801561013757600080fd5b5061014b6101463660046125f5565b610397565b604051610159929190612dc9565b60405180910390f35b610109610170366004612743565b6108fc565b6101886101833660046127f6565b610915565b6040516101599190612b27565b6101096101a33660046127f6565b611324565b3480156101b457600080fd5b506101c86101c33660046125c1565b6113fd565b604051610159929190612ae9565b3480156101e257600080fd5b506101096101f13660046126a4565b611487565b34801561020257600080fd5b5061010961021136600461259f565b611525565b34801561022257600080fd5b5061018861023136600461259f565b611571565b61010961024436600461259f565b611586565b34801561025557600080fd5b5061025e61167d565b6040516101599190612abc565b34801561027757600080fd5b5061010961028636600461270c565b61168c565b34801561029757600080fd5b506101096102a636600461259f565b6116e1565b3480156102b757600080fd5b5061025e61172d565b3480156102cc57600080fd5b5061025e61173c565b6101886102e3366004612743565b61174b565b3480156102f457600080fd5b5061010961030336600461259f565b611fe2565b6000546001600160a01b0316331461033b5760405162461bcd60e51b815260040161033290612c02565b60405180910390fd5b60005b825181101561039257816003600085848151811061035857fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905560010161033e565b505050565b6000606060405162461bcd60e51b815260040161033290612d55565b60608152602001906001900390816103b357905050905060005b85518110156108f05760608682815181106103e457fe5b60200260200101519050805160011415610499576040805160028082526060820183529091602083019080368337505060025482519296506001600160a01b03169186915060009061043257fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508060008151811061045f57fe5b60200260200101518460018151811061047457fe5b60200260200101906001600160a01b031690816001600160a01b031681525050610607565b60025481516001600160a01b039091169082906000906104b557fe5b60200260200101516001600160a01b03161415610533576040805160028082526060820183529091602083019080368337505060025482519296506001600160a01b03169186915060009061050657fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508060018151811061045f57fe5b805160010167ffffffffffffffff8111801561054e57600080fd5b50604051908082528060200260200182016040528015610578578160200160208202803683370190505b5060025481519195506001600160a01b031690859060009061059657fe5b60200260200101906001600160a01b031690816001600160a01b03168152505060005b8151811015610605578181815181106105ce57fe5b60200260200101518582600101815181106105e557fe5b6001600160a01b03909216602092830291909101909101526001016105b9565b505b60405163d06ca61f60e01b81526001600160a01b038a169063d06ca61f9061063d90670de0b6b3a7640000908890600401612b32565b60006040518083038186803b15801561065557600080fd5b505afa92505050801561068a57506040513d6000823e601f3d908101601f1916820160405261068791908101906128f8565b60015b610724576040805160028082526060820183529091602083019080368337019050508383815181106106b857fe5b602002602001018190525060008383815181106106d157fe5b60200260200101516000815181106106e557fe5b60200260200101818152505060008383815181106106ff57fe5b602002602001015160018151811061071357fe5b6020026020010181815250506108e7565b5060405163d06ca61f60e01b81526060906001600160a01b038b169063d06ca61f9061075e90670de0b6b3a7640000908990600401612b32565b60006040518083038186803b15801561077657600080fd5b505afa15801561078a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107b291908101906128f8565b6040805160028082526060820183529293509190602083019080368337019050508484815181106107df57fe5b6020026020010181905250806001825103815181106107fa57fe5b602002602001015184848151811061080e57fe5b602002602001015160008151811061082257fe5b6020026020010181815250508460018651038151811061083e57fe5b60200260200101516001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561087e57600080fd5b505afa158015610892573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b69190612a10565b60ff168484815181106108c557fe5b60200260200101516001815181106108d957fe5b602002602001018181525050505b506001016103cd565b50915050935093915050565b61090b8787878787878761174b565b5050505050505050565b6002548751600091829182916001600160a01b0316908b90600190811061093857fe5b60200260200101516001600160a01b031614156109625750506001600160a01b03861631306109f8565b8960018151811061096f57fe5b60200260200101516001600160a01b03166370a08231896040518263ffffffff1660e01b81526004016109a29190612abc565b60206040518083038186803b1580156109ba57600080fd5b505afa1580156109ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f291906129f8565b91508790505b85821115610a185760405162461bcd60e51b815260040161033290612c5c565b600085600081518110610a2757fe5b60200260200101511180610a4f5750600085600181518110610a4557fe5b6020026020010151115b80610a6e5750600085600281518110610a6457fe5b6020026020010151115b156110c15789600181518110610a8057fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610ab39190612abc565b60206040518083038186803b158015610acb57600080fd5b505afa158015610adf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0391906129f8565b915060008b600081518110610b1457fe5b60200260200101516001600160a01b031663d06ca61f60648c81610b3457fe5b046002028d6040518363ffffffff1660e01b8152600401610b56929190612b32565b60006040518083038186803b158015610b6e57600080fd5b505afa158015610b82573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610baa91908101906128f8565b600181518110610bb657fe5b602002602001015190508a600081518110610bcd57fe5b60200260200101516001600160a01b031663a9059cbb8d600281518110610bf057fe5b602002602001015160648d81610c0257fe5b046002026040518363ffffffff1660e01b8152600401610c23929190612ad0565b602060405180830381600087803b158015610c3d57600080fd5b505af1158015610c51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c759190612988565b610c7b57fe5b610cb18b8d600281518110610c8c57fe5b60200260200101518e600081518110610ca157fe5b602002602001015130600061202e565b60648a046002028a0399506000838c600181518110610ccc57fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610cff9190612abc565b60206040518083038186803b158015610d1757600080fd5b505afa158015610d2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4f91906129f8565b039050600087600181518110610d6157fe5b60200260200101511115610db9576000828281610d7a57fe5b04606402606403905087600181518110610d9057fe5b6020026020010151811115610db75760405162461bcd60e51b815260040161033290612d78565b505b60008c600081518110610dc857fe5b602002602001015190508c600181518110610ddf57fe5b60200260200101518d600081518110610df457fe5b60200260200101906001600160a01b031690816001600160a01b031681525050808d600181518110610e2257fe5b60200260200101906001600160a01b031690816001600160a01b031681525050508c600081518110610e5057fe5b60200260200101516001600160a01b031663d06ca61f828e6040518363ffffffff1660e01b8152600401610e85929190612b32565b60006040518083038186803b158015610e9d57600080fd5b505afa158015610eb1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ed991908101906128f8565b600181518110610ee557fe5b602002602001015191508b600181518110610efc57fe5b60200260200101516001600160a01b031663a9059cbb8e600281518110610f1f57fe5b6020026020010151836040518363ffffffff1660e01b8152600401610f45929190612ad0565b602060405180830381600087803b158015610f5f57600080fd5b505af1158015610f73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f979190612988565b610f9d57fe5b610fc38c8e600281518110610fae57fe5b60200260200101518f600081518110610ca157fe5b808c600181518110610fd157fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016110049190612abc565b60206040518083038186803b15801561101c57600080fd5b505afa158015611030573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105491906129f8565b03905060008760028151811061106657fe5b602002602001015111156110be57600082828161107f57fe5b0460640260640390508760028151811061109557fe5b60200260200101518111156110bc5760405162461bcd60e51b815260040161033290612b82565b505b50505b8861115957896000815181106110d357fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016111069190612abc565b60206040518083038186803b15801561111e57600080fd5b505afa158015611132573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115691906129f8565b98505b6000848a8161116457fe5b04905060005b858110156111ec576111a68c60008151811061118257fe5b60200260200101518b8f60028151811061119857fe5b60200260200101518561230d565b6111e48c8e6002815181106111b757fe5b60200260200101518f6000815181106111cc57fe5b6020026020010151868a8e816111de57fe5b0461202e565b60010161116a565b506002548b516001600160a01b03909116908c90600190811061120b57fe5b60200260200101516001600160a01b03161415611312576002546040516370a0823160e01b81526000916001600160a01b0316906370a0823190611253903090600401612abc565b60206040518083038186803b15801561126b57600080fd5b505afa15801561127f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a391906129f8565b600254604051632e1a7d4d60e01b81529192506001600160a01b031690632e1a7d4d906112d4908490600401612dc0565b600060405180830381600087803b1580156112ee57600080fd5b505af1158015611302573d6000803e3d6000fd5b505050506113108a826123fe565b505b5060019b9a5050505050505050505050565b6001546001600160a01b03161561090b5760005a905061134a8989898989898989610915565b50615208601036020160005a828401039050600061a3db61374a830160015460405163079d229f60e01b81529290910492506001600160a01b03169063079d229f9061139c9033908590600401612ad0565b602060405180830381600087803b1580156113b657600080fd5b505af11580156113ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ee91906129f8565b50505050505050505050505050565b600080826001600160a01b0316846001600160a01b031614156114325760405162461bcd60e51b815260040161033290612bcb565b826001600160a01b0316846001600160a01b031610611452578284611455565b83835b90925090506001600160a01b0382166114805760405162461bcd60e51b815260040161033290612ca1565b9250929050565b60008590506114ac8560008151811061149c57fe5b602002602001015133888761230d565b6040805160008152602081019182905263022c0d9f60e01b9091526001600160a01b0382169063022c0d9f906114eb9086908690339060248101612e57565b600060405180830381600087803b15801561150557600080fd5b505af1158015611519573d6000803e3d6000fd5b50505050505050505050565b6000546001600160a01b0316331461154f5760405162461bcd60e51b815260040161033290612c02565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60036020526000908152604090205460ff1681565b6000546001600160a01b031633146115b05760405162461bcd60e51b815260040161033290612c02565b6001600160a01b0381166115dc5760005447906115d6906001600160a01b0316826123fe565b5061167a565b6040516370a0823160e01b81526000906001600160a01b038316906370a082319061160b903090600401612abc565b60206040518083038186803b15801561162357600080fd5b505afa158015611637573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061165b91906129f8565b60005490915061167890839030906001600160a01b03168461230d565b505b50565b6000546001600160a01b031681565b6000546001600160a01b031633146116b65760405162461bcd60e51b815260040161033290612c02565b6001600160a01b03919091166000908152600360205260409020805460ff1916911515919091179055565b6000546001600160a01b0316331461170b5760405162461bcd60e51b815260040161033290612c02565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6002546001600160a01b031681565b6001546001600160a01b031681565b6000808760018151811061175b57fe5b60200260200101516001600160a01b03166370a08231886040518263ffffffff1660e01b815260040161178e9190612abc565b60206040518083038186803b1580156117a657600080fd5b505afa1580156117ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117de91906129f8565b9050848111156118005760405162461bcd60e51b815260040161033290612c5c565b5060025460408051630d0e30db60e41b8152905134926001600160a01b03169163d0e30db091849160048082019260009290919082900301818588803b15801561184957600080fd5b505af115801561185d573d6000803e3d6000fd5b505050505060008460008151811061187157fe5b60200260200101511180611899575060008460018151811061188f57fe5b6020026020010151115b806118b857506000846002815181106118ae57fe5b6020026020010151115b15611edd576000886001815181106118cc57fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016118ff9190612abc565b60206040518083038186803b15801561191757600080fd5b505afa15801561192b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061194f91906129f8565b905060008a60008151811061196057fe5b60200260200101516001600160a01b031663d06ca61f6064858161198057fe5b046002028c6040518363ffffffff1660e01b81526004016119a2929190612b32565b60006040518083038186803b1580156119ba57600080fd5b505afa1580156119ce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119f691908101906128f8565b600181518110611a0257fe5b6020908102919091010151600280548d519293506001600160a01b03169163a9059cbb918e918110611a3057fe5b602002602001015160648681611a4257fe5b046002026040518363ffffffff1660e01b8152600401611a63929190612ad0565b602060405180830381600087803b158015611a7d57600080fd5b505af1158015611a91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ab59190612988565b611abb57fe5b6064830460020234039250611aec8a8c600281518110611ad757fe5b60200260200101518d600081518110610ca157fe5b6000828b600181518110611afc57fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611b2f9190612abc565b60206040518083038186803b158015611b4757600080fd5b505afa158015611b5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b7f91906129f8565b039050600087600181518110611b9157fe5b60200260200101511115611be9576000828281611baa57fe5b04606402606403905087600181518110611bc057fe5b6020026020010151811115611be75760405162461bcd60e51b815260040161033290612d78565b505b60008b600081518110611bf857fe5b602002602001015190508b600181518110611c0f57fe5b60200260200101518c600081518110611c2457fe5b60200260200101906001600160a01b031690816001600160a01b031681525050808c600181518110611c5257fe5b60200260200101906001600160a01b031690816001600160a01b031681525050508b600081518110611c8057fe5b60200260200101516001600160a01b031663d06ca61f828d6040518363ffffffff1660e01b8152600401611cb5929190612b32565b60006040518083038186803b158015611ccd57600080fd5b505afa158015611ce1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d0991908101906128f8565b600181518110611d1557fe5b602002602001015191508a600181518110611d2c57fe5b60200260200101516001600160a01b031663a9059cbb8d600281518110611d4f57fe5b6020026020010151836040518363ffffffff1660e01b8152600401611d75929190612ad0565b602060405180830381600087803b158015611d8f57600080fd5b505af1158015611da3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dc79190612988565b611dcd57fe5b611dde8b8d600281518110610c8c57fe5b808b600181518110611dec57fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611e1f9190612abc565b60206040518083038186803b158015611e3757600080fd5b505afa158015611e4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e6f91906129f8565b039050600087600281518110611e8157fe5b60200260200101511115611ed9576000828281611e9a57fe5b04606402606403905087600281518110611eb057fe5b6020026020010151811115611ed75760405162461bcd60e51b815260040161033290612b82565b505b5050505b60005b83811015611fd257600280548b516001600160a01b039091169163a9059cbb918d91908110611f0b57fe5b6020026020010151868581611f1c57fe5b046040518363ffffffff1660e01b8152600401611f3a929190612ad0565b602060405180830381600087803b158015611f5457600080fd5b505af1158015611f68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f8c9190612988565b611f9257fe5b611fca898b600281518110611fa357fe5b60200260200101518c600081518110611fb857fe5b60200260200101518b888c816111de57fe5b600101611ee0565b5060019998505050505050505050565b6000546001600160a01b0316331461200c5760405162461bcd60e51b815260040161033290612c02565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000808660008151811061203e57fe5b60200260200101518760018151811061205357fe5b602002602001015191509150600061206b83836113fd565b5090506000879050600080600080846001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b1580156120b257600080fd5b505afa1580156120c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120ea91906129a4565b506001600160701b031691506001600160701b03169150600080876001600160a01b03168a6001600160a01b031614612124578284612127565b83835b915091506121b1828b6001600160a01b03166370a082318a6040518263ffffffff1660e01b815260040161215b9190612abc565b60206040518083038186803b15801561217357600080fd5b505afa158015612187573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ab91906129f8565b9061248b565b604051630153543560e21b81529096506001600160a01b038e169063054d50d4906121e490899086908690600401612e8e565b60206040518083038186803b1580156121fc57600080fd5b505afa158015612210573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061223491906129f8565b94508a1561225c578a85101561225c5760405162461bcd60e51b815260040161033290612c25565b50505050600080856001600160a01b0316886001600160a01b03161461228457826000612288565b6000835b6040805160008152602081019182905263022c0d9f60e01b90915291935091506001600160a01b0386169063022c0d9f906122cc90859085908f9060248101612e57565b600060405180830381600087803b1580156122e657600080fd5b505af11580156122fa573d6000803e3d6000fd5b5050505050505050505050505050505050565b60006060856001600160a01b03166323b872dd86868660405160240161233593929190612b03565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161236e9190612aa0565b6000604051808303816000865af19150503d80600081146123ab576040519150601f19603f3d011682016040523d82523d6000602084013e6123b0565b606091505b50915091508180156123da5750805115806123da5750808060200190518101906123da9190612988565b6123f65760405162461bcd60e51b815260040161033290612d11565b505050505050565b604080516000808252602082019092526001600160a01b0384169083906040516124289190612aa0565b60006040518083038185875af1925050503d8060008114612465576040519150601f19603f3d011682016040523d82523d6000602084013e61246a565b606091505b50509050806103925760405162461bcd60e51b815260040161033290612cce565b808203828111156124ae5760405162461bcd60e51b815260040161033290612b53565b92915050565b80356001600160a01b03811681146124ae57600080fd5b600082601f8301126124db578081fd5b81356124ee6124e982612ecb565b612ea4565b81815291506020808301908481018184028601820187101561250f57600080fd5b60005b848110156125365761252488836124b4565b84529282019290820190600101612512565b505050505092915050565b600082601f830112612551578081fd5b813561255f6124e982612ecb565b81815291506020808301908481018184028601820187101561258057600080fd5b60005b8481101561253657813584529282019290820190600101612583565b6000602082840312156125b0578081fd5b6125ba83836124b4565b9392505050565b600080604083850312156125d3578081fd5b6125dd84846124b4565b91506125ec84602085016124b4565b90509250929050565b600080600060608486031215612609578081fd5b833561261481612f1b565b925060208481013561262581612f1b565b9250604085013567ffffffffffffffff811115612640578283fd5b8501601f81018713612650578283fd5b803561265e6124e982612ecb565b81815283810190838501865b84811015612693576126818c8884358901016124cb565b8452928601929086019060010161266a565b505080955050505050509250925092565b600080600080600060a086880312156126bb578081fd5b6126c587876124b4565b9450602086013567ffffffffffffffff8111156126e0578182fd5b6126ec888289016124cb565b959895975050505060408401359360608101359360809091013592509050565b6000806040838503121561271e578182fd5b61272884846124b4565b9150602083013561273881612f30565b809150509250929050565b600080600080600080600060e0888a03121561275d578182fd5b873567ffffffffffffffff80821115612774578384fd5b6127808b838c016124cb565b985060208a0135915080821115612795578384fd5b6127a18b838c016124cb565b97506127b08b60408c016124b4565b965060608a0135955060808a0135945060a08a01359150808211156127d3578384fd5b506127e08a828b01612541565b92505060c0880135905092959891949750929550565b600080600080600080600080610100898b031215612812578182fd5b883567ffffffffffffffff80821115612829578384fd5b6128358c838d016124cb565b995060208b013591508082111561284a578384fd5b6128568c838d016124cb565b985060408b0135975061286c8c60608d016124b4565b965060808b0135955060a08b0135945060c08b013591508082111561288f578384fd5b5061289c8b828c01612541565b92505060e089013590509295985092959890939650565b600080604083850312156128c5578182fd5b823567ffffffffffffffff8111156128db578283fd5b6128e7858286016124cb565b925050602083013561273881612f30565b6000602080838503121561290a578182fd5b825167ffffffffffffffff811115612920578283fd5b8301601f81018513612930578283fd5b805161293e6124e982612ecb565b818152838101908385018584028501860189101561295a578687fd5b8694505b8385101561297c57805183526001949094019391850191850161295e565b50979650505050505050565b600060208284031215612999578081fd5b81516125ba81612f30565b6000806000606084860312156129b8578081fd5b83516129c381612f3e565b60208501519093506129d481612f3e565b604085015190925063ffffffff811681146129ed578182fd5b809150509250925092565b600060208284031215612a09578081fd5b5051919050565b600060208284031215612a21578081fd5b815160ff811681146125ba578182fd5b6000815180845260208085019450808401835b83811015612a695781516001600160a01b031687529582019590820190600101612a44565b509495945050505050565b60008151808452612a8c816020860160208601612eeb565b601f01601f19169290920160200192915050565b60008251612ab2818460208701612eeb565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b901515815260200190565b600083825260406020830152612b4b6040830184612a31565b949350505050565b60208082526015908201527464732d6d6174682d7375622d756e646572666c6f7760581b604082015260600190565b60208082526029908201527f53656c6c2054617820546f6f2048696768207468616e2074686520646566696e60408201526865642076616c75652160b81b606082015260800190565b6020808252601b908201527f204572726f723a204944454e544943414c5f4144445245535345530000000000604082015260600190565b6020808252600990820152682737ba1027bbb732b960b91b604082015260600190565b6020808252601e908201527f4572726f723a20536c69707061676520414d4f554e545f4f55545f4d494e0000604082015260600190565b60208082526025908201527f4572726f723a20544f4b454e5f534e4950455f53554343455346554c4c5f414c604082015264524541445960d81b606082015260800190565b6020808252601390820152724572726f723a205a45524f5f4144445245535360681b604082015260600190565b60208082526023908201527f5472616e7366657248656c7065723a204554485f5452414e534645525f46414960408201526213115160ea1b606082015260800190565b60208082526024908201527f5472616e7366657248656c7065723a205452414e534645525f46524f4d5f46416040820152631253115160e21b606082015260800190565b6020808252600990820152680d0cae4ca40c4d2e6d60bb1b604082015260600190565b60208082526028908201527f4275792054617820546f6f2048696768207468616e2074686520646566696e65604082015267642076616c75652160c01b606082015260800190565b90815260200190565b600060408201848352602060408185015281855180845260608601915060608382028701019350828701855b82811015612e4957878603605f19018452815180518088529086019086880190895b81811015612e3357835183529288019291880191600101612e17565b5090975050509284019290840190600101612df5565b509398975050505050505050565b600085825284602083015260018060a01b038416604083015260806060830152612e846080830184612a74565b9695505050505050565b9283526020830191909152604082015260600190565b60405181810167ffffffffffffffff81118282101715612ec357600080fd5b604052919050565b600067ffffffffffffffff821115612ee1578081fd5b5060209081020190565b60005b83811015612f06578181015183820152602001612eee565b83811115612f15576000848401525b50505050565b6001600160a01b038116811461167a57600080fd5b801515811461167a57600080fd5b6001600160701b038116811461167a57600080fdfea2646970667358221220f35b607ec4a0135bac4bb6c25344031522937cb805b15e03b6b14d2855060da364736f6c634300060c0033",
***REMOVED***

// GoSniperABI is the input ABI used to generate the binding from.
// Deprecated: Use GoSniperMetaData.ABI instead.
var GoSniperABI = GoSniperMetaData.ABI

// Deprecated: Use GoSniperMetaData.Sigs instead.
// GoSniperFuncSigs maps the 4-byte function signature to its string representation.
var GoSniperFuncSigs = GoSniperMetaData.Sigs

// GoSniperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GoSniperMetaData.Bin instead.
var GoSniperBin = GoSniperMetaData.Bin

// DeployGoSniper deploys a new Ethereum contract, binding an instance of GoSniper to it.
func DeployGoSniper(auth *bind.TransactOpts, backend bind.ContractBackend, _weth common.Address, _chi common.Address) (common.Address, *types.Transaction, *GoSniper, error) ***REMOVED***
	parsed, err := GoSniperMetaData.GetAbi()
	if err != nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, err
	***REMOVED***
	if parsed == nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, errors.New("GetABI returned nil")
	***REMOVED***

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoSniperBin), backend, _weth, _chi)
	if err != nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, err
	***REMOVED***
	return address, tx, &GoSniper***REMOVED***GoSniperCaller: GoSniperCaller***REMOVED***contract: contract***REMOVED***, GoSniperTransactor: GoSniperTransactor***REMOVED***contract: contract***REMOVED***, GoSniperFilterer: GoSniperFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// GoSniper is an auto generated Go binding around an Ethereum contract.
type GoSniper struct ***REMOVED***
	GoSniperCaller     // Read-only binding to the contract
	GoSniperTransactor // Write-only binding to the contract
	GoSniperFilterer   // Log filterer for contract events
***REMOVED***

// GoSniperCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoSniperCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// GoSniperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoSniperTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// GoSniperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoSniperFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// GoSniperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoSniperSession struct ***REMOVED***
	Contract     *GoSniper         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// GoSniperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoSniperCallerSession struct ***REMOVED***
	Contract *GoSniperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
***REMOVED***

// GoSniperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoSniperTransactorSession struct ***REMOVED***
	Contract     *GoSniperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
***REMOVED***

// GoSniperRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoSniperRaw struct ***REMOVED***
	Contract *GoSniper // Generic contract binding to access the raw methods on
***REMOVED***

// GoSniperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoSniperCallerRaw struct ***REMOVED***
	Contract *GoSniperCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// GoSniperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoSniperTransactorRaw struct ***REMOVED***
	Contract *GoSniperTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewGoSniper creates a new instance of GoSniper, bound to a specific deployed contract.
func NewGoSniper(address common.Address, backend bind.ContractBackend) (*GoSniper, error) ***REMOVED***
	contract, err := bindGoSniper(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &GoSniper***REMOVED***GoSniperCaller: GoSniperCaller***REMOVED***contract: contract***REMOVED***, GoSniperTransactor: GoSniperTransactor***REMOVED***contract: contract***REMOVED***, GoSniperFilterer: GoSniperFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewGoSniperCaller creates a new read-only instance of GoSniper, bound to a specific deployed contract.
func NewGoSniperCaller(address common.Address, caller bind.ContractCaller) (*GoSniperCaller, error) ***REMOVED***
	contract, err := bindGoSniper(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &GoSniperCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewGoSniperTransactor creates a new write-only instance of GoSniper, bound to a specific deployed contract.
func NewGoSniperTransactor(address common.Address, transactor bind.ContractTransactor) (*GoSniperTransactor, error) ***REMOVED***
	contract, err := bindGoSniper(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &GoSniperTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewGoSniperFilterer creates a new log filterer instance of GoSniper, bound to a specific deployed contract.
func NewGoSniperFilterer(address common.Address, filterer bind.ContractFilterer) (*GoSniperFilterer, error) ***REMOVED***
	contract, err := bindGoSniper(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &GoSniperFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindGoSniper binds a generic wrapper to an already deployed contract.
func bindGoSniper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(GoSniperABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoSniper *GoSniperRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _GoSniper.Contract.GoSniperCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoSniper *GoSniperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.GoSniperTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_GoSniper *GoSniperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.GoSniperTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoSniper *GoSniperCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _GoSniper.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoSniper *GoSniperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_GoSniper *GoSniperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Allowed is a free data retrieval call binding the contract method 0x77a7dbc6.
//
// Solidity: function Allowed(address ) view returns(bool)
func (_GoSniper *GoSniperCaller) Allowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _GoSniper.contract.Call(opts, &out, "Allowed", arg0)

	if err != nil ***REMOVED***
		return *new(bool), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

***REMOVED***

// Allowed is a free data retrieval call binding the contract method 0x77a7dbc6.
//
// Solidity: function Allowed(address ) view returns(bool)
func (_GoSniper *GoSniperSession) Allowed(arg0 common.Address) (bool, error) ***REMOVED***
	return _GoSniper.Contract.Allowed(&_GoSniper.CallOpts, arg0)
***REMOVED***

// Allowed is a free data retrieval call binding the contract method 0x77a7dbc6.
//
// Solidity: function Allowed(address ) view returns(bool)
func (_GoSniper *GoSniperCallerSession) Allowed(arg0 common.Address) (bool, error) ***REMOVED***
	return _GoSniper.Contract.Allowed(&_GoSniper.CallOpts, arg0)
***REMOVED***

// GetPrices is a free data retrieval call binding the contract method 0x1741d40e.
//
// Solidity: function GetPrices(address _router, address pricefeed, address[][] all_pairs_to_get) view returns(uint256 ETHValue, uint256[][])
func (_GoSniper *GoSniperCaller) GetPrices(opts *bind.CallOpts, _router common.Address, pricefeed common.Address, all_pairs_to_get [][]common.Address) (*big.Int, [][]*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _GoSniper.contract.Call(opts, &out, "GetPrices", _router, pricefeed, all_pairs_to_get)

	if err != nil ***REMOVED***
		return *new(*big.Int), *new([][]*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return out0, out1, err

***REMOVED***

// GetPrices is a free data retrieval call binding the contract method 0x1741d40e.
//
// Solidity: function GetPrices(address _router, address pricefeed, address[][] all_pairs_to_get) view returns(uint256 ETHValue, uint256[][])
func (_GoSniper *GoSniperSession) GetPrices(_router common.Address, pricefeed common.Address, all_pairs_to_get [][]common.Address) (*big.Int, [][]*big.Int, error) ***REMOVED***
	return _GoSniper.Contract.GetPrices(&_GoSniper.CallOpts, _router, pricefeed, all_pairs_to_get)
***REMOVED***

// GetPrices is a free data retrieval call binding the contract method 0x1741d40e.
//
// Solidity: function GetPrices(address _router, address pricefeed, address[][] all_pairs_to_get) view returns(uint256 ETHValue, uint256[][])
func (_GoSniper *GoSniperCallerSession) GetPrices(_router common.Address, pricefeed common.Address, all_pairs_to_get [][]common.Address) (*big.Int, [][]*big.Int, error) ***REMOVED***
	return _GoSniper.Contract.GetPrices(&_GoSniper.CallOpts, _router, pricefeed, all_pairs_to_get)
***REMOVED***

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_GoSniper *GoSniperCaller) WETH(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _GoSniper.contract.Call(opts, &out, "WETH")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_GoSniper *GoSniperSession) WETH() (common.Address, error) ***REMOVED***
	return _GoSniper.Contract.WETH(&_GoSniper.CallOpts)
***REMOVED***

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_GoSniper *GoSniperCallerSession) WETH() (common.Address, error) ***REMOVED***
	return _GoSniper.Contract.WETH(&_GoSniper.CallOpts)
***REMOVED***

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(address)
func (_GoSniper *GoSniperCaller) Chi(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _GoSniper.contract.Call(opts, &out, "chi")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(address)
func (_GoSniper *GoSniperSession) Chi() (common.Address, error) ***REMOVED***
	return _GoSniper.Contract.Chi(&_GoSniper.CallOpts)
***REMOVED***

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(address)
func (_GoSniper *GoSniperCallerSession) Chi() (common.Address, error) ***REMOVED***
	return _GoSniper.Contract.Chi(&_GoSniper.CallOpts)
***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GoSniper *GoSniperCaller) Owner(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _GoSniper.contract.Call(opts, &out, "owner")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GoSniper *GoSniperSession) Owner() (common.Address, error) ***REMOVED***
	return _GoSniper.Contract.Owner(&_GoSniper.CallOpts)
***REMOVED***

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GoSniper *GoSniperCallerSession) Owner() (common.Address, error) ***REMOVED***
	return _GoSniper.Contract.Owner(&_GoSniper.CallOpts)
***REMOVED***

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_GoSniper *GoSniperCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct ***REMOVED***
	Token0 common.Address
	Token1 common.Address
***REMOVED***, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _GoSniper.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

	outstruct := new(struct ***REMOVED***
		Token0 common.Address
		Token1 common.Address
	***REMOVED***)
	if err != nil ***REMOVED***
		return *outstruct, err
	***REMOVED***

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

***REMOVED***

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_GoSniper *GoSniperSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct ***REMOVED***
	Token0 common.Address
	Token1 common.Address
***REMOVED***, error) ***REMOVED***
	return _GoSniper.Contract.SortTokens(&_GoSniper.CallOpts, tokenA, tokenB)
***REMOVED***

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_GoSniper *GoSniperCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct ***REMOVED***
	Token0 common.Address
	Token1 common.Address
***REMOVED***, error) ***REMOVED***
	return _GoSniper.Contract.SortTokens(&_GoSniper.CallOpts, tokenA, tokenB)
***REMOVED***

// RescueBalance is a paid mutator transaction binding the contract method 0x791b45d9.
//
// Solidity: function RescueBalance(address tok_addy) payable returns()
func (_GoSniper *GoSniperTransactor) RescueBalance(opts *bind.TransactOpts, tok_addy common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "RescueBalance", tok_addy)
***REMOVED***

// RescueBalance is a paid mutator transaction binding the contract method 0x791b45d9.
//
// Solidity: function RescueBalance(address tok_addy) payable returns()
func (_GoSniper *GoSniperSession) RescueBalance(tok_addy common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.RescueBalance(&_GoSniper.TransactOpts, tok_addy)
***REMOVED***

// RescueBalance is a paid mutator transaction binding the contract method 0x791b45d9.
//
// Solidity: function RescueBalance(address tok_addy) payable returns()
func (_GoSniper *GoSniperTransactorSession) RescueBalance(tok_addy common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.RescueBalance(&_GoSniper.TransactOpts, tok_addy)
***REMOVED***

// SetAllowed is a paid mutator transaction binding the contract method 0x91d426a6.
//
// Solidity: function SetAllowed(address _addr, bool _res) returns()
func (_GoSniper *GoSniperTransactor) SetAllowed(opts *bind.TransactOpts, _addr common.Address, _res bool) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "SetAllowed", _addr, _res)
***REMOVED***

// SetAllowed is a paid mutator transaction binding the contract method 0x91d426a6.
//
// Solidity: function SetAllowed(address _addr, bool _res) returns()
func (_GoSniper *GoSniperSession) SetAllowed(_addr common.Address, _res bool) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetAllowed(&_GoSniper.TransactOpts, _addr, _res)
***REMOVED***

// SetAllowed is a paid mutator transaction binding the contract method 0x91d426a6.
//
// Solidity: function SetAllowed(address _addr, bool _res) returns()
func (_GoSniper *GoSniperTransactorSession) SetAllowed(_addr common.Address, _res bool) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetAllowed(&_GoSniper.TransactOpts, _addr, _res)
***REMOVED***

// SetAllowedBulk is a paid mutator transaction binding the contract method 0x063774f7.
//
// Solidity: function SetAllowedBulk(address[] _addrs, bool _res) returns()
func (_GoSniper *GoSniperTransactor) SetAllowedBulk(opts *bind.TransactOpts, _addrs []common.Address, _res bool) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "SetAllowedBulk", _addrs, _res)
***REMOVED***

// SetAllowedBulk is a paid mutator transaction binding the contract method 0x063774f7.
//
// Solidity: function SetAllowedBulk(address[] _addrs, bool _res) returns()
func (_GoSniper *GoSniperSession) SetAllowedBulk(_addrs []common.Address, _res bool) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetAllowedBulk(&_GoSniper.TransactOpts, _addrs, _res)
***REMOVED***

// SetAllowedBulk is a paid mutator transaction binding the contract method 0x063774f7.
//
// Solidity: function SetAllowedBulk(address[] _addrs, bool _res) returns()
func (_GoSniper *GoSniperTransactorSession) SetAllowedBulk(_addrs []common.Address, _res bool) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetAllowedBulk(&_GoSniper.TransactOpts, _addrs, _res)
***REMOVED***

// TradeDirectlyByPair is a paid mutator transaction binding the contract method 0x3541c60f.
//
// Solidity: function TradeDirectlyByPair(address[] majoraddresses, address[] path, uint256 tokenamount, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns(bool)
func (_GoSniper *GoSniperTransactor) TradeDirectlyByPair(opts *bind.TransactOpts, majoraddresses []common.Address, path []common.Address, tokenamount *big.Int, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "TradeDirectlyByPair", majoraddresses, path, tokenamount, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPair is a paid mutator transaction binding the contract method 0x3541c60f.
//
// Solidity: function TradeDirectlyByPair(address[] majoraddresses, address[] path, uint256 tokenamount, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns(bool)
func (_GoSniper *GoSniperSession) TradeDirectlyByPair(majoraddresses []common.Address, path []common.Address, tokenamount *big.Int, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPair(&_GoSniper.TransactOpts, majoraddresses, path, tokenamount, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPair is a paid mutator transaction binding the contract method 0x3541c60f.
//
// Solidity: function TradeDirectlyByPair(address[] majoraddresses, address[] path, uint256 tokenamount, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns(bool)
func (_GoSniper *GoSniperTransactorSession) TradeDirectlyByPair(majoraddresses []common.Address, path []common.Address, tokenamount *big.Int, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPair(&_GoSniper.TransactOpts, majoraddresses, path, tokenamount, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithETH is a paid mutator transaction binding the contract method 0xea952c6a.
//
// Solidity: function TradeDirectlyByPairWithETH(address[] majoraddresses, address[] path, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns(bool)
func (_GoSniper *GoSniperTransactor) TradeDirectlyByPairWithETH(opts *bind.TransactOpts, majoraddresses []common.Address, path []common.Address, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "TradeDirectlyByPairWithETH", majoraddresses, path, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithETH is a paid mutator transaction binding the contract method 0xea952c6a.
//
// Solidity: function TradeDirectlyByPairWithETH(address[] majoraddresses, address[] path, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns(bool)
func (_GoSniper *GoSniperSession) TradeDirectlyByPairWithETH(majoraddresses []common.Address, path []common.Address, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPairWithETH(&_GoSniper.TransactOpts, majoraddresses, path, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithETH is a paid mutator transaction binding the contract method 0xea952c6a.
//
// Solidity: function TradeDirectlyByPairWithETH(address[] majoraddresses, address[] path, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns(bool)
func (_GoSniper *GoSniperTransactorSession) TradeDirectlyByPairWithETH(majoraddresses []common.Address, path []common.Address, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPairWithETH(&_GoSniper.TransactOpts, majoraddresses, path, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithETHGasRefund is a paid mutator transaction binding the contract method 0x17960dce.
//
// Solidity: function TradeDirectlyByPairWithETHGasRefund(address[] majoraddresses, address[] path, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns()
func (_GoSniper *GoSniperTransactor) TradeDirectlyByPairWithETHGasRefund(opts *bind.TransactOpts, majoraddresses []common.Address, path []common.Address, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "TradeDirectlyByPairWithETHGasRefund", majoraddresses, path, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithETHGasRefund is a paid mutator transaction binding the contract method 0x17960dce.
//
// Solidity: function TradeDirectlyByPairWithETHGasRefund(address[] majoraddresses, address[] path, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns()
func (_GoSniper *GoSniperSession) TradeDirectlyByPairWithETHGasRefund(majoraddresses []common.Address, path []common.Address, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPairWithETHGasRefund(&_GoSniper.TransactOpts, majoraddresses, path, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithETHGasRefund is a paid mutator transaction binding the contract method 0x17960dce.
//
// Solidity: function TradeDirectlyByPairWithETHGasRefund(address[] majoraddresses, address[] path, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns()
func (_GoSniper *GoSniperTransactorSession) TradeDirectlyByPairWithETHGasRefund(majoraddresses []common.Address, path []common.Address, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPairWithETHGasRefund(&_GoSniper.TransactOpts, majoraddresses, path, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithGasRefund is a paid mutator transaction binding the contract method 0x3e005d76.
//
// Solidity: function TradeDirectlyByPairWithGasRefund(address[] majoraddresses, address[] path, uint256 tokenamount, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns()
func (_GoSniper *GoSniperTransactor) TradeDirectlyByPairWithGasRefund(opts *bind.TransactOpts, majoraddresses []common.Address, path []common.Address, tokenamount *big.Int, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "TradeDirectlyByPairWithGasRefund", majoraddresses, path, tokenamount, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithGasRefund is a paid mutator transaction binding the contract method 0x3e005d76.
//
// Solidity: function TradeDirectlyByPairWithGasRefund(address[] majoraddresses, address[] path, uint256 tokenamount, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns()
func (_GoSniper *GoSniperSession) TradeDirectlyByPairWithGasRefund(majoraddresses []common.Address, path []common.Address, tokenamount *big.Int, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPairWithGasRefund(&_GoSniper.TransactOpts, majoraddresses, path, tokenamount, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// TradeDirectlyByPairWithGasRefund is a paid mutator transaction binding the contract method 0x3e005d76.
//
// Solidity: function TradeDirectlyByPairWithGasRefund(address[] majoraddresses, address[] path, uint256 tokenamount, address _to, uint256 _amountoutMIN, uint256 beforebal, uint256[] extrasettings, uint256 _split) payable returns()
func (_GoSniper *GoSniperTransactorSession) TradeDirectlyByPairWithGasRefund(majoraddresses []common.Address, path []common.Address, tokenamount *big.Int, _to common.Address, _amountoutMIN *big.Int, beforebal *big.Int, extrasettings []*big.Int, _split *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TradeDirectlyByPairWithGasRefund(&_GoSniper.TransactOpts, majoraddresses, path, tokenamount, _to, _amountoutMIN, beforebal, extrasettings, _split)
***REMOVED***

// UnsafeSwap is a paid mutator transaction binding the contract method 0x568a0432.
//
// Solidity: function UnsafeSwap(address _pair, address[] path, uint256 tokenamount, uint256 amount0, uint256 amount1) returns()
func (_GoSniper *GoSniperTransactor) UnsafeSwap(opts *bind.TransactOpts, _pair common.Address, path []common.Address, tokenamount *big.Int, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "UnsafeSwap", _pair, path, tokenamount, amount0, amount1)
***REMOVED***

// UnsafeSwap is a paid mutator transaction binding the contract method 0x568a0432.
//
// Solidity: function UnsafeSwap(address _pair, address[] path, uint256 tokenamount, uint256 amount0, uint256 amount1) returns()
func (_GoSniper *GoSniperSession) UnsafeSwap(_pair common.Address, path []common.Address, tokenamount *big.Int, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.UnsafeSwap(&_GoSniper.TransactOpts, _pair, path, tokenamount, amount0, amount1)
***REMOVED***

// UnsafeSwap is a paid mutator transaction binding the contract method 0x568a0432.
//
// Solidity: function UnsafeSwap(address _pair, address[] path, uint256 tokenamount, uint256 amount0, uint256 amount1) returns()
func (_GoSniper *GoSniperTransactorSession) UnsafeSwap(_pair common.Address, path []common.Address, tokenamount *big.Int, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.UnsafeSwap(&_GoSniper.TransactOpts, _pair, path, tokenamount, amount0, amount1)
***REMOVED***

// SetCHI is a paid mutator transaction binding the contract method 0xa73448c9.
//
// Solidity: function setCHI(address _chi) returns()
func (_GoSniper *GoSniperTransactor) SetCHI(opts *bind.TransactOpts, _chi common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "setCHI", _chi)
***REMOVED***

// SetCHI is a paid mutator transaction binding the contract method 0xa73448c9.
//
// Solidity: function setCHI(address _chi) returns()
func (_GoSniper *GoSniperSession) SetCHI(_chi common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetCHI(&_GoSniper.TransactOpts, _chi)
***REMOVED***

// SetCHI is a paid mutator transaction binding the contract method 0xa73448c9.
//
// Solidity: function setCHI(address _chi) returns()
func (_GoSniper *GoSniperTransactorSession) SetCHI(_chi common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetCHI(&_GoSniper.TransactOpts, _chi)
***REMOVED***

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_GoSniper *GoSniperTransactor) SetWETH(opts *bind.TransactOpts, _weth common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "setWETH", _weth)
***REMOVED***

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_GoSniper *GoSniperSession) SetWETH(_weth common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetWETH(&_GoSniper.TransactOpts, _weth)
***REMOVED***

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_GoSniper *GoSniperTransactorSession) SetWETH(_weth common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.SetWETH(&_GoSniper.TransactOpts, _weth)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GoSniper *GoSniperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.Transact(opts, "transferOwnership", newOwner)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GoSniper *GoSniperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TransferOwnership(&_GoSniper.TransactOpts, newOwner)
***REMOVED***

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GoSniper *GoSniperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.TransferOwnership(&_GoSniper.TransactOpts, newOwner)
***REMOVED***

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GoSniper *GoSniperTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.RawTransact(opts, calldata)
***REMOVED***

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GoSniper *GoSniperSession) Fallback(calldata []byte) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.Fallback(&_GoSniper.TransactOpts, calldata)
***REMOVED***

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GoSniper *GoSniperTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.Fallback(&_GoSniper.TransactOpts, calldata)
***REMOVED***

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GoSniper *GoSniperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _GoSniper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
***REMOVED***

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GoSniper *GoSniperSession) Receive() (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.Receive(&_GoSniper.TransactOpts)
***REMOVED***

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GoSniper *GoSniperTransactorSession) Receive() (*types.Transaction, error) ***REMOVED***
	return _GoSniper.Contract.Receive(&_GoSniper.TransactOpts)
***REMOVED***

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Approval\",\"type\":\"event\"***REMOVED***,***REMOVED***\"anonymous\":false,\"inputs\":[***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"Transfer\",\"type\":\"event\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***],\"name\":\"allowance\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"approve\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"***REMOVED***],\"name\":\"balanceOf\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"decimals\",\"outputs\":[***REMOVED***\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[***REMOVED***\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"***REMOVED***,***REMOVED***\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"name\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"symbol\",\"outputs\":[***REMOVED***\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transfer\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"***REMOVED***],\"name\":\"transferFrom\",\"outputs\":[***REMOVED***\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***]",
	Sigs: map[string]string***REMOVED***
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"d0e30db0": "deposit()",
		"feaf968c": "latestRoundData()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	***REMOVED***,
***REMOVED***

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct ***REMOVED***
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
***REMOVED***

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct ***REMOVED***
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct ***REMOVED***
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
***REMOVED***

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct ***REMOVED***
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct ***REMOVED***
	Contract *IERC20 // Generic contract binding to access the raw methods on
***REMOVED***

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct ***REMOVED***
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct ***REMOVED***
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) ***REMOVED***
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IERC20***REMOVED***IERC20Caller: IERC20Caller***REMOVED***contract: contract***REMOVED***, IERC20Transactor: IERC20Transactor***REMOVED***contract: contract***REMOVED***, IERC20Filterer: IERC20Filterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) ***REMOVED***
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IERC20Caller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) ***REMOVED***
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IERC20Transactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) ***REMOVED***
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IERC20Filterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
***REMOVED***

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) ***REMOVED***
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) ***REMOVED***
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
***REMOVED***

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) ***REMOVED***
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil ***REMOVED***
		return *new(uint8), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) ***REMOVED***
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
***REMOVED***

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) ***REMOVED***
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
***REMOVED***

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IERC20 *IERC20Caller) LatestRoundData(opts *bind.CallOpts) (struct ***REMOVED***
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
***REMOVED***, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IERC20.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct ***REMOVED***
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	***REMOVED***)
	if err != nil ***REMOVED***
		return *outstruct, err
	***REMOVED***

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

***REMOVED***

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IERC20 *IERC20Session) LatestRoundData() (struct ***REMOVED***
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
***REMOVED***, error) ***REMOVED***
	return _IERC20.Contract.LatestRoundData(&_IERC20.CallOpts)
***REMOVED***

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IERC20 *IERC20CallerSession) LatestRoundData() (struct ***REMOVED***
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
***REMOVED***, error) ***REMOVED***
	return _IERC20.Contract.LatestRoundData(&_IERC20.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) ***REMOVED***
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
***REMOVED***

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) ***REMOVED***
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil ***REMOVED***
		return *new(string), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) ***REMOVED***
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
***REMOVED***

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) ***REMOVED***
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) ***REMOVED***
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
***REMOVED***

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) ***REMOVED***
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.contract.Transact(opts, "approve", spender, value)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
***REMOVED***

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IERC20 *IERC20Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _IERC20.contract.Transact(opts, "deposit")
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IERC20 *IERC20Session) Deposit() (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Deposit(&_IERC20.TransactOpts)
***REMOVED***

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IERC20 *IERC20TransactorSession) Deposit() (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Deposit(&_IERC20.TransactOpts)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.contract.Transact(opts, "transfer", to, value)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
***REMOVED***

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
***REMOVED***

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IERC20 *IERC20Transactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.contract.Transact(opts, "withdraw", arg0)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IERC20 *IERC20Session) Withdraw(arg0 *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Withdraw(&_IERC20.TransactOpts, arg0)
***REMOVED***

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IERC20 *IERC20TransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IERC20.Contract.Withdraw(&_IERC20.TransactOpts, arg0)
***REMOVED***

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct ***REMOVED***
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct ***REMOVED***
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IERC20ApprovalIterator***REMOVED***contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) ***REMOVED***

	var ownerRule []interface***REMOVED******REMOVED***
	for _, ownerItem := range owner ***REMOVED***
		ownerRule = append(ownerRule, ownerItem)
	***REMOVED***
	var spenderRule []interface***REMOVED******REMOVED***
	for _, spenderItem := range spender ***REMOVED***
		spenderRule = append(spenderRule, spenderItem)
	***REMOVED***

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) ***REMOVED***
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct ***REMOVED***
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool ***REMOVED***
	// If the iterator failed, stop iterating
	if it.fail != nil ***REMOVED***
		return false
	***REMOVED***
	// If the iterator completed, deliver directly whatever's available
	if it.done ***REMOVED***
		select ***REMOVED***
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error ***REMOVED***
	return it.fail
***REMOVED***

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error ***REMOVED***
	it.sub.Unsubscribe()
	return nil
***REMOVED***

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct ***REMOVED***
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
***REMOVED***

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IERC20TransferIterator***REMOVED***contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub***REMOVED***, nil
***REMOVED***

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) ***REMOVED***

	var fromRule []interface***REMOVED******REMOVED***
	for _, fromItem := range from ***REMOVED***
		fromRule = append(fromRule, fromItem)
	***REMOVED***
	var toRule []interface***REMOVED******REMOVED***
	for _, toItem := range to ***REMOVED***
		toRule = append(toRule, toItem)
	***REMOVED***

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return event.NewSubscription(func(quit <-chan struct***REMOVED******REMOVED***) error ***REMOVED***
		defer sub.Unsubscribe()
		for ***REMOVED***
			select ***REMOVED***
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) ***REMOVED***
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	event.Raw = log
	return event, nil
***REMOVED***

// IFactoryMetaData contains all meta data concerning the IFactory contract.
var IFactoryMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"name\":\"allPairs\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"***REMOVED***],\"name\":\"getPair\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***]",
	Sigs: map[string]string***REMOVED***
		"1e3dd18b": "allPairs(uint256)",
		"574f2ba3": "allPairsLength()",
		"e6a43905": "getPair(address,address)",
	***REMOVED***,
***REMOVED***

// IFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IFactoryMetaData.ABI instead.
var IFactoryABI = IFactoryMetaData.ABI

// Deprecated: Use IFactoryMetaData.Sigs instead.
// IFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IFactoryFuncSigs = IFactoryMetaData.Sigs

// IFactory is an auto generated Go binding around an Ethereum contract.
type IFactory struct ***REMOVED***
	IFactoryCaller     // Read-only binding to the contract
	IFactoryTransactor // Write-only binding to the contract
	IFactoryFilterer   // Log filterer for contract events
***REMOVED***

// IFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFactoryCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFactoryTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFactoryFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFactorySession struct ***REMOVED***
	Contract     *IFactory         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// IFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFactoryCallerSession struct ***REMOVED***
	Contract *IFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
***REMOVED***

// IFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFactoryTransactorSession struct ***REMOVED***
	Contract     *IFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
***REMOVED***

// IFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFactoryRaw struct ***REMOVED***
	Contract *IFactory // Generic contract binding to access the raw methods on
***REMOVED***

// IFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFactoryCallerRaw struct ***REMOVED***
	Contract *IFactoryCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// IFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFactoryTransactorRaw struct ***REMOVED***
	Contract *IFactoryTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewIFactory creates a new instance of IFactory, bound to a specific deployed contract.
func NewIFactory(address common.Address, backend bind.ContractBackend) (*IFactory, error) ***REMOVED***
	contract, err := bindIFactory(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IFactory***REMOVED***IFactoryCaller: IFactoryCaller***REMOVED***contract: contract***REMOVED***, IFactoryTransactor: IFactoryTransactor***REMOVED***contract: contract***REMOVED***, IFactoryFilterer: IFactoryFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewIFactoryCaller creates a new read-only instance of IFactory, bound to a specific deployed contract.
func NewIFactoryCaller(address common.Address, caller bind.ContractCaller) (*IFactoryCaller, error) ***REMOVED***
	contract, err := bindIFactory(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IFactoryCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewIFactoryTransactor creates a new write-only instance of IFactory, bound to a specific deployed contract.
func NewIFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFactoryTransactor, error) ***REMOVED***
	contract, err := bindIFactory(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IFactoryTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewIFactoryFilterer creates a new log filterer instance of IFactory, bound to a specific deployed contract.
func NewIFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFactoryFilterer, error) ***REMOVED***
	contract, err := bindIFactory(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IFactoryFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindIFactory binds a generic wrapper to an already deployed contract.
func bindIFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(IFactoryABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactory *IFactoryRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _IFactory.Contract.IFactoryCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactory *IFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _IFactory.Contract.IFactoryTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_IFactory *IFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _IFactory.Contract.IFactoryTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactory *IFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _IFactory.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactory *IFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _IFactory.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_IFactory *IFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _IFactory.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IFactory *IFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IFactory *IFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) ***REMOVED***
	return _IFactory.Contract.AllPairs(&_IFactory.CallOpts, arg0)
***REMOVED***

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IFactory *IFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) ***REMOVED***
	return _IFactory.Contract.AllPairs(&_IFactory.CallOpts, arg0)
***REMOVED***

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IFactory *IFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IFactory *IFactorySession) AllPairsLength() (*big.Int, error) ***REMOVED***
	return _IFactory.Contract.AllPairsLength(&_IFactory.CallOpts)
***REMOVED***

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IFactory *IFactoryCallerSession) AllPairsLength() (*big.Int, error) ***REMOVED***
	return _IFactory.Contract.AllPairsLength(&_IFactory.CallOpts)
***REMOVED***

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IFactory *IFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IFactory *IFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) ***REMOVED***
	return _IFactory.Contract.GetPair(&_IFactory.CallOpts, tokenA, tokenB)
***REMOVED***

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IFactory *IFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) ***REMOVED***
	return _IFactory.Contract.GetPair(&_IFactory.CallOpts, tokenA, tokenB)
***REMOVED***

// ILiquidityPairMetaData contains all meta data concerning the ILiquidityPair contract.
var ILiquidityPairMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[],\"name\":\"factory\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[***REMOVED***\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"***REMOVED***,***REMOVED***\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"***REMOVED***,***REMOVED***\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"***REMOVED***],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"token0\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***]",
	Sigs: map[string]string***REMOVED***
		"c45a0155": "factory()",
		"0902f1ac": "getReserves()",
		"022c0d9f": "swap(uint256,uint256,address,bytes)",
		"fff6cae9": "sync()",
		"0dfe1681": "token0()",
	***REMOVED***,
***REMOVED***

// ILiquidityPairABI is the input ABI used to generate the binding from.
// Deprecated: Use ILiquidityPairMetaData.ABI instead.
var ILiquidityPairABI = ILiquidityPairMetaData.ABI

// Deprecated: Use ILiquidityPairMetaData.Sigs instead.
// ILiquidityPairFuncSigs maps the 4-byte function signature to its string representation.
var ILiquidityPairFuncSigs = ILiquidityPairMetaData.Sigs

// ILiquidityPair is an auto generated Go binding around an Ethereum contract.
type ILiquidityPair struct ***REMOVED***
	ILiquidityPairCaller     // Read-only binding to the contract
	ILiquidityPairTransactor // Write-only binding to the contract
	ILiquidityPairFilterer   // Log filterer for contract events
***REMOVED***

// ILiquidityPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILiquidityPairCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ILiquidityPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILiquidityPairTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ILiquidityPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILiquidityPairFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// ILiquidityPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILiquidityPairSession struct ***REMOVED***
	Contract     *ILiquidityPair   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// ILiquidityPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILiquidityPairCallerSession struct ***REMOVED***
	Contract *ILiquidityPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
***REMOVED***

// ILiquidityPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILiquidityPairTransactorSession struct ***REMOVED***
	Contract     *ILiquidityPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
***REMOVED***

// ILiquidityPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILiquidityPairRaw struct ***REMOVED***
	Contract *ILiquidityPair // Generic contract binding to access the raw methods on
***REMOVED***

// ILiquidityPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILiquidityPairCallerRaw struct ***REMOVED***
	Contract *ILiquidityPairCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// ILiquidityPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILiquidityPairTransactorRaw struct ***REMOVED***
	Contract *ILiquidityPairTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewILiquidityPair creates a new instance of ILiquidityPair, bound to a specific deployed contract.
func NewILiquidityPair(address common.Address, backend bind.ContractBackend) (*ILiquidityPair, error) ***REMOVED***
	contract, err := bindILiquidityPair(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ILiquidityPair***REMOVED***ILiquidityPairCaller: ILiquidityPairCaller***REMOVED***contract: contract***REMOVED***, ILiquidityPairTransactor: ILiquidityPairTransactor***REMOVED***contract: contract***REMOVED***, ILiquidityPairFilterer: ILiquidityPairFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewILiquidityPairCaller creates a new read-only instance of ILiquidityPair, bound to a specific deployed contract.
func NewILiquidityPairCaller(address common.Address, caller bind.ContractCaller) (*ILiquidityPairCaller, error) ***REMOVED***
	contract, err := bindILiquidityPair(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ILiquidityPairCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewILiquidityPairTransactor creates a new write-only instance of ILiquidityPair, bound to a specific deployed contract.
func NewILiquidityPairTransactor(address common.Address, transactor bind.ContractTransactor) (*ILiquidityPairTransactor, error) ***REMOVED***
	contract, err := bindILiquidityPair(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ILiquidityPairTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewILiquidityPairFilterer creates a new log filterer instance of ILiquidityPair, bound to a specific deployed contract.
func NewILiquidityPairFilterer(address common.Address, filterer bind.ContractFilterer) (*ILiquidityPairFilterer, error) ***REMOVED***
	contract, err := bindILiquidityPair(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &ILiquidityPairFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindILiquidityPair binds a generic wrapper to an already deployed contract.
func bindILiquidityPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(ILiquidityPairABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILiquidityPair *ILiquidityPairRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _ILiquidityPair.Contract.ILiquidityPairCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILiquidityPair *ILiquidityPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.ILiquidityPairTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_ILiquidityPair *ILiquidityPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.ILiquidityPairTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILiquidityPair *ILiquidityPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _ILiquidityPair.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILiquidityPair *ILiquidityPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_ILiquidityPair *ILiquidityPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ILiquidityPair *ILiquidityPairCaller) Factory(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ILiquidityPair.contract.Call(opts, &out, "factory")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ILiquidityPair *ILiquidityPairSession) Factory() (common.Address, error) ***REMOVED***
	return _ILiquidityPair.Contract.Factory(&_ILiquidityPair.CallOpts)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ILiquidityPair *ILiquidityPairCallerSession) Factory() (common.Address, error) ***REMOVED***
	return _ILiquidityPair.Contract.Factory(&_ILiquidityPair.CallOpts)
***REMOVED***

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ILiquidityPair *ILiquidityPairCaller) GetReserves(opts *bind.CallOpts) (struct ***REMOVED***
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
***REMOVED***, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ILiquidityPair.contract.Call(opts, &out, "getReserves")

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
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ILiquidityPair *ILiquidityPairSession) GetReserves() (struct ***REMOVED***
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
***REMOVED***, error) ***REMOVED***
	return _ILiquidityPair.Contract.GetReserves(&_ILiquidityPair.CallOpts)
***REMOVED***

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ILiquidityPair *ILiquidityPairCallerSession) GetReserves() (struct ***REMOVED***
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
***REMOVED***, error) ***REMOVED***
	return _ILiquidityPair.Contract.GetReserves(&_ILiquidityPair.CallOpts)
***REMOVED***

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() pure returns(address)
func (_ILiquidityPair *ILiquidityPairCaller) Token0(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _ILiquidityPair.contract.Call(opts, &out, "token0")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() pure returns(address)
func (_ILiquidityPair *ILiquidityPairSession) Token0() (common.Address, error) ***REMOVED***
	return _ILiquidityPair.Contract.Token0(&_ILiquidityPair.CallOpts)
***REMOVED***

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() pure returns(address)
func (_ILiquidityPair *ILiquidityPairCallerSession) Token0() (common.Address, error) ***REMOVED***
	return _ILiquidityPair.Contract.Token0(&_ILiquidityPair.CallOpts)
***REMOVED***

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ILiquidityPair *ILiquidityPairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
***REMOVED***

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ILiquidityPair *ILiquidityPairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.Swap(&_ILiquidityPair.TransactOpts, amount0Out, amount1Out, to, data)
***REMOVED***

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ILiquidityPair *ILiquidityPairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.Swap(&_ILiquidityPair.TransactOpts, amount0Out, amount1Out, to, data)
***REMOVED***

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ILiquidityPair *ILiquidityPairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.contract.Transact(opts, "sync")
***REMOVED***

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ILiquidityPair *ILiquidityPairSession) Sync() (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.Sync(&_ILiquidityPair.TransactOpts)
***REMOVED***

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ILiquidityPair *ILiquidityPairTransactorSession) Sync() (*types.Transaction, error) ***REMOVED***
	return _ILiquidityPair.Contract.Sync(&_ILiquidityPair.TransactOpts)
***REMOVED***

// IRouterMetaData contains all meta data concerning the IRouter contract.
var IRouterMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[],\"name\":\"factory\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"***REMOVED***],\"name\":\"getAmountOut\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***],\"name\":\"getAmountsOut\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactETHForTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactTokensForETH\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***]",
	Sigs: map[string]string***REMOVED***
		"c45a0155": "factory()",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
		"b6f9de95": "swapExactETHForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)",
		"2679fb5f": "swapExactTokensForETH(uint256,address[],uint256)",
		"791ac947": "swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
	***REMOVED***,
***REMOVED***

// IRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IRouterMetaData.ABI instead.
var IRouterABI = IRouterMetaData.ABI

// Deprecated: Use IRouterMetaData.Sigs instead.
// IRouterFuncSigs maps the 4-byte function signature to its string representation.
var IRouterFuncSigs = IRouterMetaData.Sigs

// IRouter is an auto generated Go binding around an Ethereum contract.
type IRouter struct ***REMOVED***
	IRouterCaller     // Read-only binding to the contract
	IRouterTransactor // Write-only binding to the contract
	IRouterFilterer   // Log filterer for contract events
***REMOVED***

// IRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRouterCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRouterTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRouterFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// IRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRouterSession struct ***REMOVED***
	Contract     *IRouter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// IRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRouterCallerSession struct ***REMOVED***
	Contract *IRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
***REMOVED***

// IRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRouterTransactorSession struct ***REMOVED***
	Contract     *IRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
***REMOVED***

// IRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRouterRaw struct ***REMOVED***
	Contract *IRouter // Generic contract binding to access the raw methods on
***REMOVED***

// IRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRouterCallerRaw struct ***REMOVED***
	Contract *IRouterCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// IRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRouterTransactorRaw struct ***REMOVED***
	Contract *IRouterTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewIRouter creates a new instance of IRouter, bound to a specific deployed contract.
func NewIRouter(address common.Address, backend bind.ContractBackend) (*IRouter, error) ***REMOVED***
	contract, err := bindIRouter(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IRouter***REMOVED***IRouterCaller: IRouterCaller***REMOVED***contract: contract***REMOVED***, IRouterTransactor: IRouterTransactor***REMOVED***contract: contract***REMOVED***, IRouterFilterer: IRouterFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewIRouterCaller creates a new read-only instance of IRouter, bound to a specific deployed contract.
func NewIRouterCaller(address common.Address, caller bind.ContractCaller) (*IRouterCaller, error) ***REMOVED***
	contract, err := bindIRouter(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IRouterCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewIRouterTransactor creates a new write-only instance of IRouter, bound to a specific deployed contract.
func NewIRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IRouterTransactor, error) ***REMOVED***
	contract, err := bindIRouter(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IRouterTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewIRouterFilterer creates a new log filterer instance of IRouter, bound to a specific deployed contract.
func NewIRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IRouterFilterer, error) ***REMOVED***
	contract, err := bindIRouter(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &IRouterFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindIRouter binds a generic wrapper to an already deployed contract.
func bindIRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(IRouterABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouter *IRouterRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _IRouter.Contract.IRouterCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouter *IRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.IRouterTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_IRouter *IRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.IRouterTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouter *IRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _IRouter.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouter *IRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_IRouter *IRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IRouter *IRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IRouter.contract.Call(opts, &out, "factory")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IRouter *IRouterSession) Factory() (common.Address, error) ***REMOVED***
	return _IRouter.Contract.Factory(&_IRouter.CallOpts)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IRouter *IRouterCallerSession) Factory() (common.Address, error) ***REMOVED***
	return _IRouter.Contract.Factory(&_IRouter.CallOpts)
***REMOVED***

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IRouter *IRouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IRouter.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IRouter *IRouterSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	return _IRouter.Contract.GetAmountOut(&_IRouter.CallOpts, amountIn, reserveIn, reserveOut)
***REMOVED***

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IRouter *IRouterCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	return _IRouter.Contract.GetAmountOut(&_IRouter.CallOpts, amountIn, reserveIn, reserveOut)
***REMOVED***

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IRouter *IRouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _IRouter.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil ***REMOVED***
		return *new([]*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

***REMOVED***

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IRouter *IRouterSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	return _IRouter.Contract.GetAmountsOut(&_IRouter.CallOpts, amountIn, path)
***REMOVED***

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IRouter *IRouterCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	return _IRouter.Contract.GetAmountsOut(&_IRouter.CallOpts, amountIn, path)
***REMOVED***

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactETHForTokens(&_IRouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactETHForTokens(&_IRouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IRouter *IRouterTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IRouter *IRouterSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IRouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IRouter *IRouterTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IRouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x2679fb5f.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, address[] path, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.contract.Transact(opts, "swapExactTokensForETH", amountIn, path, deadline)
***REMOVED***

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x2679fb5f.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, address[] path, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterSession) SwapExactTokensForETH(amountIn *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactTokensForETH(&_IRouter.TransactOpts, amountIn, path, deadline)
***REMOVED***

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x2679fb5f.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, address[] path, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactTokensForETH(&_IRouter.TransactOpts, amountIn, path, deadline)
***REMOVED***

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IRouter *IRouterTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IRouter *IRouterSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IRouter *IRouterTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IRouter *IRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _IRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData***REMOVED***
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209483e6c4d0fccd848ed665e72db3a7012b3b1088daa58c47d06481b2be71c2c364736f6c634300060c0033",
***REMOVED***

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) ***REMOVED***
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, err
	***REMOVED***
	if parsed == nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, errors.New("GetABI returned nil")
	***REMOVED***

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, err
	***REMOVED***
	return address, tx, &SafeMath***REMOVED***SafeMathCaller: SafeMathCaller***REMOVED***contract: contract***REMOVED***, SafeMathTransactor: SafeMathTransactor***REMOVED***contract: contract***REMOVED***, SafeMathFilterer: SafeMathFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct ***REMOVED***
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
***REMOVED***

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct ***REMOVED***
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct ***REMOVED***
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
***REMOVED***

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct ***REMOVED***
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
***REMOVED***

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct ***REMOVED***
	Contract *SafeMath // Generic contract binding to access the raw methods on
***REMOVED***

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct ***REMOVED***
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct ***REMOVED***
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) ***REMOVED***
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &SafeMath***REMOVED***SafeMathCaller: SafeMathCaller***REMOVED***contract: contract***REMOVED***, SafeMathTransactor: SafeMathTransactor***REMOVED***contract: contract***REMOVED***, SafeMathFilterer: SafeMathFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) ***REMOVED***
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &SafeMathCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) ***REMOVED***
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &SafeMathTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) ***REMOVED***
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &SafeMathFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _SafeMath.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData***REMOVED***
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122042dac9a155eb229b19a071a2e89c08036a89c16a8e7a5423cd038288d861430a64736f6c634300060c0033",
***REMOVED***

// TransferHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperMetaData.ABI instead.
var TransferHelperABI = TransferHelperMetaData.ABI

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferHelperMetaData.Bin instead.
var TransferHelperBin = TransferHelperMetaData.Bin

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) ***REMOVED***
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, err
	***REMOVED***
	if parsed == nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, errors.New("GetABI returned nil")
	***REMOVED***

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil ***REMOVED***
		return common.Address***REMOVED******REMOVED***, nil, nil, err
	***REMOVED***
	return address, tx, &TransferHelper***REMOVED***TransferHelperCaller: TransferHelperCaller***REMOVED***contract: contract***REMOVED***, TransferHelperTransactor: TransferHelperTransactor***REMOVED***contract: contract***REMOVED***, TransferHelperFilterer: TransferHelperFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct ***REMOVED***
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
***REMOVED***

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct ***REMOVED***
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct ***REMOVED***
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
***REMOVED***

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct ***REMOVED***
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
***REMOVED***

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct ***REMOVED***
	Contract *TransferHelper // Generic contract binding to access the raw methods on
***REMOVED***

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct ***REMOVED***
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct ***REMOVED***
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) ***REMOVED***
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &TransferHelper***REMOVED***TransferHelperCaller: TransferHelperCaller***REMOVED***contract: contract***REMOVED***, TransferHelperTransactor: TransferHelperTransactor***REMOVED***contract: contract***REMOVED***, TransferHelperFilterer: TransferHelperFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) ***REMOVED***
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &TransferHelperCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) ***REMOVED***
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &TransferHelperTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) ***REMOVED***
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &TransferHelperFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _TransferHelper.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
***REMOVED***
