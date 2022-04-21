package sniper

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/logicfool/GoSniper/structs"
	"github.com/logicfool/GoSniper/utils"
)

// A back Compatibility with Routers in Case The Contract is not deployed on a chain

func (client *MainClient) LegacyRouterETHBuy(tokenAddress interface{}, amount interface{}, wallets_to_use interface{}) []*types.Transaction {
	var weth common.Address
	var exchange structs.Exchange
	var path []common.Address
	var buy_amount *big.Int
	var wallet_to_use []*bind.TransactOpts
	var tx_hashes []*types.Transaction
	// if tokenAddress is a string, convert it to a common.Address
	exchange = client.MainExchange
	weth, _ = exchange.Router.WETH(&bind.CallOpts{})

	// fmt.Println(exchange)
	switch tokenAddress.(type) {
	case string:
		path = []common.Address{weth, utils.ToChecksumAddress(tokenAddress.(string))}

	case []common.Address:
		path = tokenAddress.([]common.Address)
	}
	switch amount.(type) {
	case *big.Int:
		buy_amount = amount.(*big.Int)
	case *big.Float:
		buy_amount = utils.FloatToBigInt(amount.(*big.Float))
	case int:
		buy_amount = big.NewInt(int64(amount.(int)))
	}
	// var pair =
	// fmt.Println(path)
	amountsout := client.GetAmountsOut(path, buy_amount, exchange)
	minamountout := amountsout[len(amountsout)-1]
	deadline := big.NewInt(time.Now().Add(time.Minute * 5).Unix())

	if wallets_to_use == nil {
		wallet_to_use = client.UpdateWalletAuthValues(buy_amount)
	} else {
		wallet_to_use = wallets_to_use.([]*bind.TransactOpts)
	}

	for _, wallet := range wallet_to_use {
		// fmt.Println(wallet)
		tx, err := exchange.Router.SwapExactETHForTokensSupportingFeeOnTransferTokens(wallet, minamountout, path, wallet.From, deadline)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Tx Hash %v: %v\n", wallet.From.Hex(), tx.Hash().Hex())
		tx_hashes = append(tx_hashes, tx)
	}
	//for each element in tx hashes get the receipt if successful or not
	for _, tx := range tx_hashes {
		receipt := client.GetTxReceipt(tx.Hash().Hex())
		if receipt.Status == 1 {
			fmt.Printf("Tx %v: Successful\n", tx.Hash().Hex())
		} else {
			fmt.Printf("Tx %v: Failed\n", tx.Hash().Hex())
		}
	}
	return tx_hashes
	// fmt.Printf("%T", contract)
	// make sure amount is already in wei and not in ether
}
