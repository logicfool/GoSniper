package sniper

import (
	"os"

	"github.com/logicfool/GoSniper/utils"
)

// var client ethclient.Client = GetClient(*&conf.Chains[0])

// conf, err := config.Read_config_from_file("config.json")

func (client *MainClient) Start() ***REMOVED***
	// client.WaitGroup.Add(1)
	if client.Sniper.MempoolLiquidityCheck ***REMOVED***
		go client.SubscribePendingTx()
	***REMOVED***
	// utils.ColorPrint(fmt.Sprintf("Multi Main Wallet: %v", client.Sniper.MutliMainSnipe), "yellow")
	// os.Exit(0)
	// Buy Scope
	// for client.Sniper.ToBuy ***REMOVED***
	// 	// _ = Eligible_wallets
	// 	// client.BuyTokenWithPair()
	// ***REMOVED***
	client.PreTxCheckandGetWallets()

	// Sell Scope
	// for client.Sniper.ToSell ***REMOVED***
	// 	client.PreTxCheckandGetWallets()
	// 	// _ = Eligible_wallets
	// 	client.BuyTokenWithPair()
	// ***REMOVED***
	client.WaitGroup.Wait()
	// client.Stop()
***REMOVED***

func (client *MainClient) Stop() ***REMOVED***
	client.Sniper.ToBuy = false
	client.Sniper.ToSell = false
	utils.ColorPrint("Stopping Sniper", "yellow")
	os.Exit(0)
	// Quit <- true
	// send true to quit without blocking the main thread
	// client.WaitGroup.Done()

	// os.Exit(0)
***REMOVED***
