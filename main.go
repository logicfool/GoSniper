package main

import (
	"fmt"

	"github.com/logicfool/GoSniper/sniper"
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

	MainConf := sniper.QuickMaintainConfig(56)
	fmt.Println(MainConf.GetClient())
	fmt.Println(MainConf.WeiToEther(1000000000000000000, 18))
	// weth, err := MainConf.Exchanges[0].Router.WETH(&bind.CallOpts***REMOVED******REMOVED***)
	// if err != nil ***REMOVED***
	// ***REMOVED***
	// path := []common.Address***REMOVED***weth, utils.ToChecksumAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")***REMOVED***
	// deadline := time.Now().Add(time.Minute * 5).Unix()
	// buy_conf, err := sniper.BuyToken(MainConf.Exchanges[0].Router, path, MainConf.WalletAuths[0], big.NewInt(1))
***REMOVED***
