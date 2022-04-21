package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func GetRouter(address common.Address, client *ethclient.Client) *Unirouter {
// 	contract, err := NewUnirouter(address, client)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return contract
// }

// func GetFactory(address common.Address, client *ethclient.Client) *Unifactory {
// 	contract, err := NewUnifactory(address, client)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return contract
// }

// func GetErc20(address common.Address, client *ethclient.Client) *Erc20 {
// 	contract, err := NewErc20(address, client)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return contract
// }

// func GetPair(address common.Address, client *ethclient.Client) *Pair {
// 	contract, err := NewPair(address, client)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return contract
// }

type (
	Token struct {
		Address  common.Address
		Name     string
		Symbol   string
		Decimals *big.Int
		Contract ERC20
	}
)

func GetContract(address common.Address, client *ethclient.Client, contract_type string) interface{} {
	var contract interface{}
	if contract_type == "router" {
		contract, _ = NewUnirouter(address, client)
	} else if contract_type == "factory" {
		contract, _ = NewUnifactory(address, client)
	} else if contract_type == "erc20" || contract_type == "token" {
		TOKEN, _ := NewERC20(address, client)
		Name, err := TOKEN.Name(&bind.CallOpts{})
		if err != nil {
			Name = "Not Found!"
		}
		Symbol, err := TOKEN.Symbol(&bind.CallOpts{})
		if err != nil {
			Symbol = "Not Found!"
		}
		Decimals, err := TOKEN.Decimals(&bind.CallOpts{})
		if err != nil {
			Decimals = uint8(0)
		}
		Token1 := Token{
			Address:  address,
			Name:     Name,
			Symbol:   Symbol,
			Decimals: new(big.Int).SetUint64(uint64(Decimals)),
			Contract: *TOKEN,
		}
		contract = &Token1
	} else if contract_type == "pair" {
		contract, _ = NewPair(address, client)
	} else if contract_type == "sniper" {
		contract, _ = NewGoSniper(address, client)
	} else if contract_type == "weth" {
		contract, _ = NewWETH(address, client)
	} else {
		panic("Unknown contract type")
	}
	return contract
}
