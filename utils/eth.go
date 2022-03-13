package utils

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

type (
	Address string
)

func FloatToBigInt(val interface***REMOVED******REMOVED***) *big.Int ***REMOVED***
	// Convert a float to a big.Int
	var bigval big.Float
	switch v := val.(type) ***REMOVED***
	case float64:
		bigval.SetFloat64(v)
	case *big.Float:
		bigval = *v
	***REMOVED***
	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(&bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
***REMOVED***

func ParseBigFloat(value string) (*big.Float, error) ***REMOVED***
	// Parse a string into a big.Float
	// Thanks to StackoverFlow
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
***REMOVED***

func WeiToEther(wei interface***REMOVED******REMOVED***) *big.Float ***REMOVED***
	// Useful when converting from Wei to Ether .. Can also be used in tokens with decimal 18
	switch wei.(type) ***REMOVED***
	case *big.Int:
		return new(big.Float).Quo(new(big.Float).SetInt(wei.(*big.Int)), big.NewFloat(params.Ether))
	case *big.Float:
		return new(big.Float).Quo(wei.(*big.Float), big.NewFloat(params.Ether))
	case float64:
		return new(big.Float).Quo(new(big.Float).SetInt(big.NewInt(0).Mul(big.NewInt(0).SetUint64(uint64(wei.(float64))), big.NewInt(1000000000000000000))), big.NewFloat(params.Ether))
	case int:
		return WeiToEther(big.NewInt(int64(wei.(int))))
	***REMOVED***
	panic("Type Not Supported!")
***REMOVED***

func WeiToEtherByDecimals(wei interface***REMOVED******REMOVED***, decimals int) *big.Float ***REMOVED***
	// Useful when converting from Wei to Ether in tokens as they have decimals != 18
	if decimals == 0 ***REMOVED***
		decimals = 18
	***REMOVED***
	switch wei.(type) ***REMOVED***
	case *big.Int:
		fl := new(big.Float)
		fl.SetInt(wei.(*big.Int))
		value := new(big.Float).Quo(fl, big.NewFloat(math.Pow10(decimals)))
		return value
	case *big.Float:
		return new(big.Float).Quo(wei.(*big.Float), big.NewFloat(math.Pow10(decimals)))
	case int:
		return WeiToEtherByDecimals(big.NewInt(int64(wei.(int))), decimals)
	***REMOVED***
	panic("Invalid Wei Type!")
***REMOVED***

func EtherToWei(ether interface***REMOVED******REMOVED***) *big.Int ***REMOVED***
	switch ether.(type) ***REMOVED***
	case float64:
		return big.NewInt(int64(ether.(float64) * 1000000000000000000))
	case float32:
		return big.NewInt(int64(ether.(float32) * 1000000000000000000))
	case int:
		return big.NewInt(int64(ether.(int) * 1000000000000000000))
	case *big.Int:
		return ether.(*big.Int).Mul(ether.(*big.Int), big.NewInt(1000000000000000000))
	case *big.Float:
		return FloatToBigInt(ether.(*big.Float).Mul(ether.(*big.Float), big.NewFloat(1000000000000000000)))
	***REMOVED***
	panic("Invalid Ether Type!")
***REMOVED***

func EtherToWeiByDecimals(amount interface***REMOVED******REMOVED***, decimals int) *big.Int ***REMOVED***
	if decimals == 0 ***REMOVED***
		decimals = 18
	***REMOVED***
	if decimals < 0 ***REMOVED***
		panic("decimals must be greater than 0")
	***REMOVED***
	switch amount.(type) ***REMOVED***
	case float64:
		return big.NewInt(int64(amount.(float64) * math.Pow10(decimals)))
	case int:
		return big.NewInt(int64(float64(amount.(int)) * math.Pow10(decimals)))
	case *big.Int:
		return amount.(*big.Int).Mul(amount.(*big.Int), big.NewInt(int64(math.Pow10(decimals))))
	case *big.Float:
		return FloatToBigInt(amount.(*big.Float).Mul(amount.(*big.Float), big.NewFloat(math.Pow10(decimals))))
	***REMOVED***
	panic("Invalid Ether Type!")
***REMOVED***

func (a Address) Addr() common.Address ***REMOVED***
	if len(a) == 0 ***REMOVED***
		panic("empty address")
	***REMOVED***
	return common.HexToAddress(string(a))
***REMOVED***

func ToChecksumAddress(a string) common.Address ***REMOVED***
	if len(a) == 0 ***REMOVED***
		panic("empty address")
	***REMOVED***
	b := Address(a)
	return b.Addr()
***REMOVED***

func updateNonce() ***REMOVED******REMOVED***

func sendtx() ***REMOVED******REMOVED***

// func (a string) Addr() common.Address ***REMOVED***
// 	if len(a) == 0 ***REMOVED***
// 		panic("empty address")
// 	***REMOVED***
// 	return common.HexToAddress(string(a))
// ***REMOVED***

// func (a string) Hex() string ***REMOVED***
// 	if len(a) == 0 ***REMOVED***
// 		panic("empty address")
// 	***REMOVED***
// 	return string(a)
// ***REMOVED***
