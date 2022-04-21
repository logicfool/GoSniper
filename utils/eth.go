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

type etherconversion struct {
	wei  int
	gwei int
}

var Ethermeasures = etherconversion{18, 9}

func FloatToBigInt(val interface{}) *big.Int {
	// Convert a float to a big.Int
	var bigval big.Float
	switch v := val.(type) {
	case float64:
		bigval.SetFloat64(v)
	case *big.Float:
		bigval = *v
	}
	coin := new(big.Float)
	coin.SetInt(big.NewInt(1))

	bigval.Mul(&bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func BigIntToBigFloat(val *big.Int) *big.Float {
	return new(big.Float).SetInt(val)
}

// func FindNetworkByname(name *string, configf *structs.BaseConfig) structs.Network {
// 	for _, network := range configf.Chains {
// 		if network.Name == *name {
// 			return network
// 		}
// 	}
// 	// return first network
// 	return configf.Chains[0]
// }

func ParseBigFloat(value string) (*big.Float, error) {
	// Parse a string into a big.Float
	// Thanks to StackoverFlow
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

func WeiToEther(wei interface{}) *big.Float {
	// Useful when converting from Wei to Ether .. Can also be used in tokens with decimal 18
	switch wei.(type) {
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
	}
	panic("Type Not Supported!")
}

func WeiToEtherByDecimals(wei interface{}, decimals int) *big.Float {
	// Useful when converting from Wei to Ether in tokens as they have decimals != 18
	if decimals == 0 {
		decimals = 18
	}
	switch wei.(type) {
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
	}
	panic("Invalid Wei Type!")
}

func convert(val uint64, decimals int64) *big.Int {
	v := big.NewInt(int64(val))
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil)
	return v.Mul(v, exp)
}

func EtherToWei(ether interface{}) *big.Int {
	switch ether.(type) {
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
	}
	panic(fmt.Sprintf("Invalid Ether Type! %T", ether))
}

func EtherToWeiByDecimals(amount interface{}, decimals int) *big.Int {
	if decimals == 0 {
		decimals = 18
	}
	if decimals < 0 {
		panic("decimals must be greater than 0")
	}
	switch amount.(type) {
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
	}
	panic("Invalid Ether Type!")
}

func (a Address) Addr() common.Address {
	if len(a) == 0 {
		panic("empty address")
	}
	return common.HexToAddress(string(a))
}

func ToChecksumAddress(a string) common.Address {
	if len(a) == 0 {
		panic("empty address")
	}
	b := Address(a)
	return b.Addr()
}

// Keccak256 calculates the Keccak256
func Keccak256(v ...[]byte) []byte {
	h := sha3.NewLegacyKeccak256()
	for _, i := range v {
		h.Write(i)
	}
	return h.Sum(nil)
}

func EncodeFunctionName(name interface{}) string {
	var namebyte []byte
	switch name.(type) {
	case string:
		// cast to []byte
		namebyte = []byte(name.(string))
	case []byte:
		namebyte = name.([]byte)
	}

	return hex.EncodeToString(Keccak256(namebyte))[0:8]
}

func HexToBigInt(hex string) *big.Int {
	n, _ := new(big.Int).SetString(hex, 16)
	return n
}

func StringToBigInt(ss string) *big.Int {
	n, _ := new(big.Int).SetString(ss, 10)
	return n
}

// function sortTokens(address tokenA, address tokenB) public pure returns (address token0, address token1) {
// 	require(tokenA != tokenB, ' Error: IDENTICAL_ADDRESSES');
// 	(token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
// 	require(token0 != address(0), 'Error: ZERO_ADDRESS');
// }

func SortTokens(tokenA, tokenB common.Address) (common.Address, common.Address, error) {
	if tokenA == tokenB {
		// panic("Error: IDENTICAL_ADDRESSES")
		return common.Address{}, common.Address{}, errors.New("Error: IDENTICAL_ADDRESSES")
	}
	// convert common.Address to int
	tokenAInt := new(big.Int).SetBytes(tokenA.Bytes())
	tokenBInt := new(big.Int).SetBytes(tokenB.Bytes())
	// compare
	if tokenAInt.Cmp(tokenBInt) < 0 {
		return tokenA, tokenB, nil
	}
	return tokenB, tokenA, nil
}

func GetPair(factoryAddress common.Address, token0 common.Address, token1 common.Address, initCodePairHash string) (common.Address, error) {
	token0, token1, err := SortTokens(token0, token1)
	if err != nil {
		// panic(err)
		return common.Address{}, err

	}
	if initCodePairHash[:2] == "0x" {
		initCodePairHash = initCodePairHash[2:]
	}

	tokenPackedHash := Keccak256(common.Hex2Bytes(token0.String()[2:] + token1.String()[2:]))
	allPacked := common.Hex2Bytes("ff" + factoryAddress.String()[2:] + hexutil.Encode(tokenPackedHash)[2:] + initCodePairHash)
	allPackedHash := Keccak256(allPacked)
	allPackedHashBigInt := new(big.Int).SetBytes(allPackedHash)
	pairAddress := common.BigToAddress(allPackedHashBigInt)
	return pairAddress, nil
}

func updateNonce() {}

func sendtx() {}

// func (a string) Addr() common.Address {
// 	if len(a) == 0 {
// 		panic("empty address")
// 	}
// 	return common.HexToAddress(string(a))
// }

// func (a string) Hex() string {
// 	if len(a) == 0 {
// 		panic("empty address")
// 	}
// 	return string(a)
// }
