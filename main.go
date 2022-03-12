package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logicfool/GoSniper/config"
	"github.com/logicfool/GoSniper/consts"
	"github.com/logicfool/GoSniper/sniper"
	"github.com/logicfool/GoSniper/utils"
)

func main() ***REMOVED***
	conn, err := ethclient.Dial("https://speedy-nodes-nyc.moralis.io/78fd8cc9150d7e2a80c27dc4/eth/mainnet")
	if err != nil ***REMOVED***
		fmt.Printf("%v\n", err)
		return
	***REMOVED***
	conn.Close()
	// ctx := context.Background()
	// tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("0x0701ece3270b544197ed49800f27e6fd16291b508dd880e09fcdf809b23b6c0a"))
	// if !pending ***REMOVED***
	// 	fmt.Println(tx)
	// ***REMOVED***
	ctx := context.Background()
	grades := []int***REMOVED***1, 12, 221***REMOVED***
	fmt.Println(utils.WeiToEtherFloat(big.NewInt(1000000000000000000)))
	fmt.Println(utils.EtherToWei(1))
	fmt.Println(consts.ChainIdMap["eth"], consts.ZeroAddress)
	fmt.Println(consts.IsOwner, *&grades)
	helo := utils.ToChecksumAddress("0xabd3e1274a93a09bea79d6e8be9259eaaf57e7dd")
	fmt.Println(helo)
	conf, err := config.Read_config_from_file("config.json")
	fmt.Println(*&conf.Chains[0].RPC)
	var client ethclient.Client = sniper.GetClient(*&conf.Chains[0])
	fmt.Println(client.ChainID(ctx))
	// var chainID, error = conn.ChainID(ctx)
	// if error != nil ***REMOVED***
	// 	log.Fatalf("Something Wrong! %v", err)
	// ***REMOVED***
	// fmt.Printf("%v\n", chainID)
***REMOVED***
