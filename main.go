package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/logicfool/GoSniper/config"
	"github.com/logicfool/GoSniper/utils"
)

func main() ***REMOVED***
	// conn.Close()

	// ctx := context.Background()
	// grades := []int***REMOVED***1, 12, 221***REMOVED***
	// fmt.Println(utils.WeiToEtherFloat(big.NewInt(1000000000000000000)))
	// fmt.Println(utils.EtherToWei(1))
	// fmt.Println(consts.ChainIdMap["eth"], consts.ZeroAddress)
	// fmt.Println(consts.IsOwner, *&grades)
	// helo := utils.ToChecksumAddress("0xabd3e1274a93a09bea79d6e8be9259eaaf57e7dd")
	// fmt.Println(helo)
	// conf, err := config.Read_config_from_file("config.json")
	// fmt.Println(*&conf.Exchanges[0])
	// var client *ethclient.Client = sniper.GetClient(*&conf.Chains[1])
	// // fmt.Println(client.ChainID(ctx))
	// fmt.Printf("%T", err)
	// // var chainID, error = conn.ChainID(ctx)
	// // if error != nil ***REMOVED***
	// // 	log.Fatalf("Something Wrong! %v", err)
	// // ***REMOVED***
	// // fmt.Printf("%v\n", chainID)
	// contract, err := contracts.GetRouter(utils.ToChecksumAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), client)
	// output, err := contract.WETH(&bind.CallOpts***REMOVED******REMOVED***)
	// fmt.Println("WETH ADDRESS FROM UNI ROUTER", output, err)

	MainConf := config.QuickMaintainConfig(56)
	weth, err := MainConf.Exchanges[0].Router.WETH(&bind.CallOpts***REMOVED******REMOVED***)
	if err != nil ***REMOVED***
	***REMOVED***
	path := []common.Address***REMOVED***weth, utils.ToChecksumAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")***REMOVED***
	// interact and buy a token from pancakeswap contract
	// deadline is 5 mins from now
	// path:= [weth,utils.ToChecksumAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")]
	deadline := time.Now().Add(time.Minute * 5).Unix()
	tx, err := MainConf.Exchanges[0].Router.SwapExactETHForTokensSupportingFeeOnTransferTokens(MainConf.WalletAuths[0], big.NewInt(0), path, MainConf.WalletAuths[0].From, big.NewInt(deadline))
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	fmt.Println(tx.Hash().Hex())

***REMOVED***
