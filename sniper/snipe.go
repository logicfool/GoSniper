package sniper

import (
	"os"

	"github.com/logicfool/GoSniper/utils"
)

// var client ethclient.Client = GetClient(*&conf.Chains[0])

// conf, err := config.Read_config_from_file("config.json")

func (client *MainClient) Start() {
	// client.WaitGroup.Add(1)
	if client.Sniper.MempoolLiquidityCheck {
		go client.SubscribePendingTx()
	}
	// utils.ColorPrint(fmt.Sprintf("Multi Main Wallet: %v", client.Sniper.MutliMainSnipe), "yellow")
	// os.Exit(0)
	// Buy Scope
	// for client.Sniper.ToBuy {
	// 	// _ = Eligible_wallets
	// 	// client.BuyTokenWithPair()
	// }
	client.PreTxCheckandGetWallets()

	// Sell Scope
	// for client.Sniper.ToSell {
	// 	client.PreTxCheckandGetWallets()
	// 	// _ = Eligible_wallets
	// 	client.BuyTokenWithPair()
	// }
	client.WaitGroup.Wait()
	// client.Stop()
}

func (client *MainClient) Stop() {
	client.Sniper.ToBuy = false
	client.Sniper.ToSell = false
	utils.ColorPrint("Stopping Sniper", "yellow")
	os.Exit(0)
	// Quit <- true
	// send true to quit without blocking the main thread
	// client.WaitGroup.Done()

	// os.Exit(0)
}
