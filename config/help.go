package config

import "github.com/logicfool/GoSniper/utils"

func Showhelp() {
	// Usage Show
	utils.ColorPrint("Usage: ./Sniper [OPTIONS] TokenAddress Amount\n", "green", false)
	// Start Options
	utils.ColorPrint("Options:\n", "yellow", false)
	utils.ColorPrint("\t--help: Show this help Message\n", "blue", false)
	utils.ColorPrint("\t--pair {pairaddress}: Pair Address To buy token in a different pair\n", "blue", false)
	utils.ColorPrint("\t--watch: Run sniper in watch mode\n", "blue", false)
	utils.ColorPrint("\t--maxbt {buytax}: Maximum Buy Tax to check if higher bot will wait for it to be lowered!\n", "blue", false)
	utils.ColorPrint("\t--maxst {selltax}: Maximum Sell Tax to check if higher bot will wait for it to be lowered!\n", "blue", false)
	// utils.ColorPrint("\t--skipblocks {blocknumer}: Number of blocks to Skip After Liquidity is added\n", "blue", false)
	// utils.ColorPrint("\t--minliq {liqamount}: Minimum Liquidity in pool if lower bot wont buy!\n", "blue", false)
	utils.ColorPrint("\t--slippage {slippage}: Slippage to used when Trading\n", "blue", false)
	utils.ColorPrint("\t--multiwallet: Enable Multi wallet\n", "blue", false)
	utils.ColorPrint("\t--buysplit {split}: Buy in Splits\n", "blue", false)
	utils.ColorPrint("\t--sellsplit {split}: Sell in Splits\n", "blue", false)
	utils.ColorPrint("\t--hpc: Enable Honeypot Check\n", "blue", false)
	utils.ColorPrint("\t--mlc: Enable Mempool Liquidity Check\n", "blue", false)
	utils.ColorPrint("\t--mrlc: Enable Mempool Remove Liquidity Check\n", "blue", false)
	utils.ColorPrint("\t--dtc: Disable Trading Check\n", "blue", false)
	utils.ColorPrint("\t--network {network}: Blockchain Network to use\n", "blue", false)
	utils.ColorPrint("\t--exchange {exchange}: Exchange to Use from the network!\n", "blue", false)
	utils.ColorPrint("\t--onlybuy: Only buy\n", "blue", false)
	utils.ColorPrint("\t--onlysell: Only sell\n", "blue", false)
	utils.ColorPrint("\t--chigas: Enable ChiGas In tx's\n", "blue", false)
	utils.ColorPrint("\t--sellprice {price}: Sell Price To sell the token on\n", "blue", false)
	utils.ColorPrint("\t--buyprice {price}: Buy Price To buy the token on\n", "blue", false)
	utils.ColorPrint("\t--takeprofit {profit}: Take Profit %\n", "blue", false)
	utils.ColorPrint("\t--stoploss {loss}: Stop Loss %\n", "blue", false)
	// End Options

	// Notes
	utils.ColorPrint("Notes:\n", "yellow", false)
	utils.ColorPrint("\t - To add an exchange you need the router/factory address as well as its INIT_HASH_CODE you can get it from the router's source code once you are on the page find `PairFor` Function and you will see the init code in hex value there copy it and paste it in config!\n", "cyan", false)
	utils.ColorPrint("\t - You need to add datafeed address in the network when adding a new network else Price Features wont work!You can get the datafeed address from https://data.chain.link/\n", "cyan", false)

	utils.ColorPrint("\t - if you dont pass an exchange in options the First Exchange in the config of the selected network will be selected!\n", "cyan", false)
	utils.ColorPrint("\t - if you dont pass a network in options the First Network in the config will be selected!\n", "cyan", false)
	utils.ColorPrint("\t - If you Enable Multi Wallet the same amount you want will be bought in all the wallets present in the config.. If you disable multiwallet the first wallet from the config will be selected as primary wallet\n", "cyan", false)

}
