package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func GetRouter(address common.Address, client *ethclient.Client) *Unirouter ***REMOVED***
// 	contract, err := NewUnirouter(address, client)
// 	if err != nil ***REMOVED***
// 		panic(err)
// 	***REMOVED***
// 	return contract
// ***REMOVED***

// func GetFactory(address common.Address, client *ethclient.Client) *Unifactory ***REMOVED***
// 	contract, err := NewUnifactory(address, client)
// 	if err != nil ***REMOVED***
// 		panic(err)
// 	***REMOVED***
// 	return contract
// ***REMOVED***

// func GetErc20(address common.Address, client *ethclient.Client) *Erc20 ***REMOVED***
// 	contract, err := NewErc20(address, client)
// 	if err != nil ***REMOVED***
// 		panic(err)
// 	***REMOVED***
// 	return contract
// ***REMOVED***

// func GetPair(address common.Address, client *ethclient.Client) *Pair ***REMOVED***
// 	contract, err := NewPair(address, client)
// 	if err != nil ***REMOVED***
// 		panic(err)
// 	***REMOVED***
// 	return contract
// ***REMOVED***

type (
	Token struct ***REMOVED***
		Address  common.Address
		Name     string
		Symbol   string
		Decimals *big.Int
		Contract ERC20
	***REMOVED***
)

func GetContract(address common.Address, client *ethclient.Client, contract_type string) interface***REMOVED******REMOVED*** ***REMOVED***
	var contract interface***REMOVED******REMOVED***
	if contract_type == "router" ***REMOVED***
		contract, _ = NewUnirouter(address, client)
	***REMOVED*** else if contract_type == "factory" ***REMOVED***
		contract, _ = NewUnifactory(address, client)
	***REMOVED*** else if contract_type == "erc20" || contract_type == "token" ***REMOVED***
		TOKEN, _ := NewERC20(address, client)
		Name, err := TOKEN.Name(&bind.CallOpts***REMOVED******REMOVED***)
		if err != nil ***REMOVED***
			Name = "Not Found!"
		***REMOVED***
		Symbol, err := TOKEN.Symbol(&bind.CallOpts***REMOVED******REMOVED***)
		if err != nil ***REMOVED***
			Symbol = "Not Found!"
		***REMOVED***
		Decimals, err := TOKEN.Decimals(&bind.CallOpts***REMOVED******REMOVED***)
		if err != nil ***REMOVED***
			Decimals = uint8(0)
		***REMOVED***
		Token1 := Token***REMOVED***
			Address:  address,
			Name:     Name,
			Symbol:   Symbol,
			Decimals: new(big.Int).SetUint64(uint64(Decimals)),
			Contract: *TOKEN,
		***REMOVED***
		contract = &Token1
	***REMOVED*** else if contract_type == "pair" ***REMOVED***
		contract, _ = NewPair(address, client)
	***REMOVED*** else if contract_type == "sniper" ***REMOVED***
		contract, _ = NewGoSniper(address, client)
	***REMOVED*** else if contract_type == "weth" ***REMOVED***
		contract, _ = NewWETH(address, client)
	***REMOVED*** else ***REMOVED***
		panic("Unknown contract type")
	***REMOVED***
	return contract
***REMOVED***
