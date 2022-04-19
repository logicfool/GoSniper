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

// UnirouterMetaData contains all meta data concerning the Unirouter contract.
var UnirouterMetaData = &bind.MetaData***REMOVED***
	ABI: "[***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"WETH\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"addLiquidity\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"addLiquidityETH\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[],\"name\":\"factory\",\"outputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"***REMOVED***],\"name\":\"getAmountIn\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"***REMOVED***],\"name\":\"getAmountOut\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***],\"name\":\"getAmountsIn\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***],\"name\":\"getAmountsOut\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"view\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"***REMOVED***],\"name\":\"quote\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"pure\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"removeLiquidity\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"removeLiquidityETH\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"***REMOVED***,***REMOVED***\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"***REMOVED***],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"***REMOVED***,***REMOVED***\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"***REMOVED***],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"***REMOVED***,***REMOVED***\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"***REMOVED***,***REMOVED***\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"***REMOVED***],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapETHForExactTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactETHForTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactTokensForETH\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactTokensForTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapTokensForExactETH\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"swapTokensForExactTokens\",\"outputs\":[***REMOVED***\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"***REMOVED***],\"stateMutability\":\"nonpayable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"stateMutability\":\"payable\",\"type\":\"receive\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountBNBMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"addLiquidityBNB\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountBNB\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***,***REMOVED***\"inputs\":[***REMOVED***\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"***REMOVED***],\"name\":\"addLiquidityAVAX\",\"outputs\":[***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"***REMOVED***,***REMOVED***\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"***REMOVED***],\"stateMutability\":\"payable\",\"type\":\"function\"***REMOVED***]",
***REMOVED***

// UnirouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UnirouterMetaData.ABI instead.
var UnirouterABI = UnirouterMetaData.ABI

// Unirouter is an auto generated Go binding around an Ethereum contract.
type Unirouter struct ***REMOVED***
	UnirouterCaller     // Read-only binding to the contract
	UnirouterTransactor // Write-only binding to the contract
	UnirouterFilterer   // Log filterer for contract events
***REMOVED***

// UnirouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnirouterCaller struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// UnirouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnirouterTransactor struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// UnirouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnirouterFilterer struct ***REMOVED***
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
***REMOVED***

// UnirouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnirouterSession struct ***REMOVED***
	Contract     *Unirouter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
***REMOVED***

// UnirouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnirouterCallerSession struct ***REMOVED***
	Contract *UnirouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
***REMOVED***

// UnirouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnirouterTransactorSession struct ***REMOVED***
	Contract     *UnirouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
***REMOVED***

// UnirouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnirouterRaw struct ***REMOVED***
	Contract *Unirouter // Generic contract binding to access the raw methods on
***REMOVED***

// UnirouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnirouterCallerRaw struct ***REMOVED***
	Contract *UnirouterCaller // Generic read-only contract binding to access the raw methods on
***REMOVED***

// UnirouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnirouterTransactorRaw struct ***REMOVED***
	Contract *UnirouterTransactor // Generic write-only contract binding to access the raw methods on
***REMOVED***

// NewUnirouter creates a new instance of Unirouter, bound to a specific deployed contract.
func NewUnirouter(address common.Address, backend bind.ContractBackend) (*Unirouter, error) ***REMOVED***
	contract, err := bindUnirouter(address, backend, backend, backend)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &Unirouter***REMOVED***UnirouterCaller: UnirouterCaller***REMOVED***contract: contract***REMOVED***, UnirouterTransactor: UnirouterTransactor***REMOVED***contract: contract***REMOVED***, UnirouterFilterer: UnirouterFilterer***REMOVED***contract: contract***REMOVED******REMOVED***, nil
***REMOVED***

// NewUnirouterCaller creates a new read-only instance of Unirouter, bound to a specific deployed contract.
func NewUnirouterCaller(address common.Address, caller bind.ContractCaller) (*UnirouterCaller, error) ***REMOVED***
	contract, err := bindUnirouter(address, caller, nil, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &UnirouterCaller***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewUnirouterTransactor creates a new write-only instance of Unirouter, bound to a specific deployed contract.
func NewUnirouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UnirouterTransactor, error) ***REMOVED***
	contract, err := bindUnirouter(address, nil, transactor, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &UnirouterTransactor***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// NewUnirouterFilterer creates a new log filterer instance of Unirouter, bound to a specific deployed contract.
func NewUnirouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UnirouterFilterer, error) ***REMOVED***
	contract, err := bindUnirouter(address, nil, nil, filterer)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return &UnirouterFilterer***REMOVED***contract: contract***REMOVED***, nil
***REMOVED***

// bindUnirouter binds a generic wrapper to an already deployed contract.
func bindUnirouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) ***REMOVED***
	parsed, err := abi.JSON(strings.NewReader(UnirouterABI))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unirouter *UnirouterRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Unirouter.Contract.UnirouterCaller.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unirouter *UnirouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.UnirouterTransactor.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Unirouter *UnirouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.UnirouterTransactor.contract.Transact(opts, method, params...)
***REMOVED***

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unirouter *UnirouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface***REMOVED******REMOVED***, method string, params ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return _Unirouter.Contract.contract.Call(opts, result, method, params...)
***REMOVED***

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unirouter *UnirouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.contract.Transfer(opts)
***REMOVED***

// Transact invokes the (paid) contract method with params as input values.
func (_Unirouter *UnirouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface***REMOVED******REMOVED***) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.contract.Transact(opts, method, params...)
***REMOVED***

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Unirouter *UnirouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unirouter.contract.Call(opts, &out, "WETH")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Unirouter *UnirouterSession) WETH() (common.Address, error) ***REMOVED***
	return _Unirouter.Contract.WETH(&_Unirouter.CallOpts)
***REMOVED***

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Unirouter *UnirouterCallerSession) WETH() (common.Address, error) ***REMOVED***
	return _Unirouter.Contract.WETH(&_Unirouter.CallOpts)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Unirouter *UnirouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unirouter.contract.Call(opts, &out, "factory")

	if err != nil ***REMOVED***
		return *new(common.Address), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Unirouter *UnirouterSession) Factory() (common.Address, error) ***REMOVED***
	return _Unirouter.Contract.Factory(&_Unirouter.CallOpts)
***REMOVED***

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Unirouter *UnirouterCallerSession) Factory() (common.Address, error) ***REMOVED***
	return _Unirouter.Contract.Factory(&_Unirouter.CallOpts)
***REMOVED***

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Unirouter *UnirouterCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unirouter.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Unirouter *UnirouterSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountIn(&_Unirouter.CallOpts, amountOut, reserveIn, reserveOut)
***REMOVED***

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Unirouter *UnirouterCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountIn(&_Unirouter.CallOpts, amountOut, reserveIn, reserveOut)
***REMOVED***

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Unirouter *UnirouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unirouter.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Unirouter *UnirouterSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountOut(&_Unirouter.CallOpts, amountIn, reserveIn, reserveOut)
***REMOVED***

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Unirouter *UnirouterCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountOut(&_Unirouter.CallOpts, amountIn, reserveIn, reserveOut)
***REMOVED***

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Unirouter *UnirouterCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unirouter.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil ***REMOVED***
		return *new([]*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

***REMOVED***

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountsIn(&_Unirouter.CallOpts, amountOut, path)
***REMOVED***

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Unirouter *UnirouterCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountsIn(&_Unirouter.CallOpts, amountOut, path)
***REMOVED***

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Unirouter *UnirouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unirouter.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil ***REMOVED***
		return *new([]*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

***REMOVED***

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountsOut(&_Unirouter.CallOpts, amountIn, path)
***REMOVED***

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Unirouter *UnirouterCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.GetAmountsOut(&_Unirouter.CallOpts, amountIn, path)
***REMOVED***

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Unirouter *UnirouterCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) ***REMOVED***
	var out []interface***REMOVED******REMOVED***
	err := _Unirouter.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil ***REMOVED***
		return *new(*big.Int), err
	***REMOVED***

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

***REMOVED***

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Unirouter *UnirouterSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.Quote(&_Unirouter.CallOpts, amountA, reserveA, reserveB)
***REMOVED***

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Unirouter *UnirouterCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) ***REMOVED***
	return _Unirouter.Contract.Quote(&_Unirouter.CallOpts, amountA, reserveA, reserveB)
***REMOVED***

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Unirouter *UnirouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
***REMOVED***

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Unirouter *UnirouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidity(&_Unirouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
***REMOVED***

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Unirouter *UnirouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidity(&_Unirouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
***REMOVED***

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_Unirouter *UnirouterTransactor) AddLiquidityAVAX(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "addLiquidityAVAX", token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
***REMOVED***

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_Unirouter *UnirouterSession) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidityAVAX(&_Unirouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
***REMOVED***

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_Unirouter *UnirouterTransactorSession) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidityAVAX(&_Unirouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
***REMOVED***

// AddLiquidityBNB is a paid mutator transaction binding the contract method 0xeaaed442.
//
// Solidity: function addLiquidityBNB(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountBNB, uint256 liquidity)
func (_Unirouter *UnirouterTransactor) AddLiquidityBNB(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "addLiquidityBNB", token, amountTokenDesired, amountTokenMin, amountBNBMin, to, deadline)
***REMOVED***

// AddLiquidityBNB is a paid mutator transaction binding the contract method 0xeaaed442.
//
// Solidity: function addLiquidityBNB(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountBNB, uint256 liquidity)
func (_Unirouter *UnirouterSession) AddLiquidityBNB(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidityBNB(&_Unirouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountBNBMin, to, deadline)
***REMOVED***

// AddLiquidityBNB is a paid mutator transaction binding the contract method 0xeaaed442.
//
// Solidity: function addLiquidityBNB(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountBNB, uint256 liquidity)
func (_Unirouter *UnirouterTransactorSession) AddLiquidityBNB(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidityBNB(&_Unirouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountBNBMin, to, deadline)
***REMOVED***

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Unirouter *UnirouterTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Unirouter *UnirouterSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidityETH(&_Unirouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Unirouter *UnirouterTransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.AddLiquidityETH(&_Unirouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Unirouter *UnirouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
***REMOVED***

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Unirouter *UnirouterSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidity(&_Unirouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
***REMOVED***

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Unirouter *UnirouterTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidity(&_Unirouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
***REMOVED***

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Unirouter *UnirouterTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Unirouter *UnirouterSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETH(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Unirouter *UnirouterTransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETH(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_Unirouter *UnirouterTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_Unirouter *UnirouterSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_Unirouter *UnirouterTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
***REMOVED***

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_Unirouter *UnirouterTransactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_Unirouter *UnirouterSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETHWithPermit(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_Unirouter *UnirouterTransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETHWithPermit(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_Unirouter *UnirouterTransactor) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_Unirouter *UnirouterSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_Unirouter *UnirouterTransactorSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Unirouter *UnirouterTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Unirouter *UnirouterSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityWithPermit(&_Unirouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Unirouter *UnirouterTransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.RemoveLiquidityWithPermit(&_Unirouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
***REMOVED***

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
***REMOVED***

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapETHForExactTokens(&_Unirouter.TransactOpts, amountOut, path, to, deadline)
***REMOVED***

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapETHForExactTokens(&_Unirouter.TransactOpts, amountOut, path, to, deadline)
***REMOVED***

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactETHForTokens(&_Unirouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactETHForTokens(&_Unirouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Unirouter *UnirouterTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Unirouter *UnirouterSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Unirouter *UnirouterTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForETH(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForETH(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Unirouter *UnirouterTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Unirouter *UnirouterSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Unirouter *UnirouterTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForTokens(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForTokens(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Unirouter *UnirouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Unirouter *UnirouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Unirouter *UnirouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Unirouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
***REMOVED***

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
***REMOVED***

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapTokensForExactETH(&_Unirouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
***REMOVED***

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapTokensForExactETH(&_Unirouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
***REMOVED***

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
***REMOVED***

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapTokensForExactTokens(&_Unirouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
***REMOVED***

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Unirouter *UnirouterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.SwapTokensForExactTokens(&_Unirouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
***REMOVED***

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Unirouter *UnirouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) ***REMOVED***
	return _Unirouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
***REMOVED***

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Unirouter *UnirouterSession) Receive() (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.Receive(&_Unirouter.TransactOpts)
***REMOVED***

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Unirouter *UnirouterTransactorSession) Receive() (*types.Transaction, error) ***REMOVED***
	return _Unirouter.Contract.Receive(&_Unirouter.TransactOpts)
***REMOVED***
