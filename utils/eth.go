package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/sha3"
)

type (
	Address string
)

type etherconversion struct ***REMOVED***
	wei  int
	gwei int
***REMOVED***

var Ethermeasures = etherconversion***REMOVED***18, 9***REMOVED***

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
	coin.SetInt(big.NewInt(1))

	bigval.Mul(&bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
***REMOVED***

func BigIntToBigFloat(val *big.Int) *big.Float ***REMOVED***
	return new(big.Float).SetInt(val)
***REMOVED***

// func FindNetworkByname(name *string, configf *structs.BaseConfig) structs.Network ***REMOVED***
// 	for _, network := range configf.Chains ***REMOVED***
// 		if network.Name == *name ***REMOVED***
// 			return network
// 		***REMOVED***
// 	***REMOVED***
// 	// return first network
// 	return configf.Chains[0]
// ***REMOVED***

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
	case uint64:
		return new(big.Float).SetUint64(wei.(uint64))
	case uint:
		return new(big.Float).SetUint64(uint64(wei.(uint)))
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
		val := new(big.Float).SetUint64(uint64(wei.(int)))
		return WeiToEtherByDecimals(val, decimals)
	case uint64:
		val := new(big.Float).SetUint64(wei.(uint64))
		return WeiToEtherByDecimals(val, decimals)
	case float64:
		val := new(big.Float).SetFloat64(wei.(float64))
		return WeiToEtherByDecimals(val, decimals)
	case float32:
		val := new(big.Float).SetFloat64(float64(wei.(float32)))
		return WeiToEtherByDecimals(val, decimals)
	***REMOVED***
	panic("Invalid Wei Type!")
***REMOVED***

func convert(val uint64, decimals int64) *big.Int ***REMOVED***
	v := big.NewInt(int64(val))
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil)
	return v.Mul(v, exp)
***REMOVED***

func EtherToWei(ether interface***REMOVED******REMOVED***) *big.Int ***REMOVED***
	switch ether.(type) ***REMOVED***
	case float64:
		// return big.NewInt(int64(ether.(float64) * 1000000000000000000))
		v := big.NewFloat(ether.(float64))
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil).Uint64()
		return FloatToBigInt(v.Mul(v, big.NewFloat(float64(exp))))

		// return convert(uint64(ether.(float64)), 18)
	case float32:
		// return big.NewInt(int64(ether.(float32) * 1000000000000000000))
		v := big.NewFloat(float64(ether.(float32)))
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil).Uint64()
		return FloatToBigInt(v.Mul(v, big.NewFloat(float64(exp))))
		// return convert(uint64(ether.(float32)), 18)
	case int:
		// return big.NewInt(int64(ether.(int) * 1000000000000000000))
		return convert(uint64(ether.(int)), 18)
	case uint64:
		// return big.NewInt(int64(ether.(uint64) * 1000000000000000000))
		return convert(ether.(uint64), 18)
	case *big.Int:
		return ether.(*big.Int).Mul(ether.(*big.Int), big.NewInt(1000000000000000000))
	case *big.Float:
		return FloatToBigInt(ether.(*big.Float).Mul(ether.(*big.Float), big.NewFloat(1000000000000000000)))
	***REMOVED***
	panic(fmt.Sprintf("Invalid Ether Type! %T", ether))
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
		// return big.NewInt(int64(amount.(float64) * math.Pow10(int(decimals))))
		// return convert(uint64(amount.(float64)), decimals)
		v := big.NewFloat(amount.(float64))
		exp := float64(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil).Uint64())
		return FloatToBigInt(v.Mul(v, big.NewFloat(exp)))
	case float32:
		v := big.NewFloat(float64(amount.(float32)))
		exp := float64(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil).Uint64())
		return FloatToBigInt(v.Mul(v, big.NewFloat(exp)))
	case int:
		// return big.NewInt(int64(float64(amount.(int)) * math.Pow10(decimals)))
		return convert(uint64(amount.(int)), int64(decimals))
	case *big.Int:
		return amount.(*big.Int).Mul(amount.(*big.Int), big.NewInt(int64(math.Pow10(int(decimals)))))
		// case *big.Float:
		// 	return FloatToBigInt(amount.(*big.Float).Mul(amount.(*big.Float), big.NewFloat(math.Pow10(decimals))))
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

// Keccak256 calculates the Keccak256
func Keccak256(v ...[]byte) []byte ***REMOVED***
	h := sha3.NewLegacyKeccak256()
	for _, i := range v ***REMOVED***
		h.Write(i)
	***REMOVED***
	return h.Sum(nil)
***REMOVED***

func EncodeFunctionName(name interface***REMOVED******REMOVED***) string ***REMOVED***
	var namebyte []byte
	switch name.(type) ***REMOVED***
	case string:
		// cast to []byte
		namebyte = []byte(name.(string))
	case []byte:
		namebyte = name.([]byte)
	***REMOVED***

	return hex.EncodeToString(Keccak256(namebyte))[0:8]
***REMOVED***

func HexToBigInt(hex string) *big.Int ***REMOVED***
	n, _ := new(big.Int).SetString(hex, 16)
	return n
***REMOVED***

func StringToBigInt(ss string) *big.Int ***REMOVED***
	n, _ := new(big.Int).SetString(ss, 10)
	return n
***REMOVED***

// function sortTokens(address tokenA, address tokenB) public pure returns (address token0, address token1) ***REMOVED***
// 	require(tokenA != tokenB, ' Error: IDENTICAL_ADDRESSES');
// 	(token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
// 	require(token0 != address(0), 'Error: ZERO_ADDRESS');
// ***REMOVED***

func SortTokens(tokenA, tokenB common.Address) (common.Address, common.Address, error) ***REMOVED***
	if tokenA == tokenB ***REMOVED***
		// panic("Error: IDENTICAL_ADDRESSES")
		return common.Address***REMOVED******REMOVED***, common.Address***REMOVED******REMOVED***, errors.New("Error: IDENTICAL_ADDRESSES")
	***REMOVED***
	// convert common.Address to int
	tokenAInt := new(big.Int).SetBytes(tokenA.Bytes())
	tokenBInt := new(big.Int).SetBytes(tokenB.Bytes())
	// compare
	if tokenAInt.Cmp(tokenBInt) < 0 ***REMOVED***
		return tokenA, tokenB, nil
	***REMOVED***
	return tokenB, tokenA, nil
***REMOVED***

func GetPair(factoryAddress common.Address, token0 common.Address, token1 common.Address, initCodePairHash string) (common.Address, error) ***REMOVED***
	token0, token1, err := SortTokens(token0, token1)
	if err != nil ***REMOVED***
		// panic(err)
		return common.Address***REMOVED******REMOVED***, err

	***REMOVED***
	if initCodePairHash[:2] == "0x" ***REMOVED***
		initCodePairHash = initCodePairHash[2:]
	***REMOVED***

	tokenPackedHash := Keccak256(common.Hex2Bytes(token0.String()[2:] + token1.String()[2:]))
	allPacked := common.Hex2Bytes("ff" + factoryAddress.String()[2:] + hexutil.Encode(tokenPackedHash)[2:] + initCodePairHash)
	allPackedHash := Keccak256(allPacked)
	allPackedHashBigInt := new(big.Int).SetBytes(allPackedHash)
	pairAddress := common.BigToAddress(allPackedHashBigInt)
	return pairAddress, nil
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
