package utils

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type (
	Address string
)

func WeiToEtherFloatByDecimals(decimals int, wei *big.Int) *big.Float ***REMOVED***
	if decimals == 0 ***REMOVED***
		decimals = 18
	***REMOVED***
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(math.Pow10(decimals)))
***REMOVED***

func WeiToEtherFloat(wei *big.Int) *big.Float ***REMOVED***
	return WeiToEtherFloatByDecimals(18, wei)
***REMOVED***

func EtherToWei(ether float64) *big.Int ***REMOVED***
	// convert float to wei
	return new(big.Int).Mul(big.NewInt(int64(ether*math.Pow10(18))), big.NewInt(1000000000000000000))

	// return new(big.Int).Mul(big.NewInt(int64(ether)), big.NewInt(1000000000000000000))
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
