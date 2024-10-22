package config

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/logicfool/GoSniper/structs"
	"github.com/logicfool/GoSniper/utils"
)

func Read_config_from_file(f string) (*structs.BaseConfig, error) {
	// if config file does not exist create one
	if _, err := os.Stat(f); os.IsNotExist(err) {
		utils.ColorPrint("[+] Config file not found, creating one Edit and restart the bot!", "green")
		config, err := Create_a_new_Config()
		os.Exit(0)
		return config, err
	}
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := structs.BaseConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return &config, nil
}

func Create_a_new_Config() (*structs.BaseConfig, error) {
	config := structs.BaseConfig{}
	config.Chains = []structs.Network{}
	config.Exchanges = []structs.Exchanges{}
	config.PrivateKeys = []string{}
	config.Sniper = structs.SniperSettings{}
	config.Sniper.Slippage = 5
	config.Sniper.MultiWallet = false
	config.Sniper.MinLiq = 0
	config.Sniper.BuySplit = 1
	config.Sniper.SellSplit = 1
	config.Sniper.HoneyPotCheck = false
	// {"f305d719","e8e33700"}
	config.Sniper.LiquidityAddFunctions = []string{}
	config.Sniper.LiquidityAddFunctions = append(config.Sniper.LiquidityAddFunctions, "f305d719")
	config.Sniper.LiquidityAddFunctions = append(config.Sniper.LiquidityAddFunctions, "e8e33700")
	// "baa2abde","02751cec","2195995c","ded9382a","af2979eb","7d72d6e0"
	config.Sniper.LiquidityRemoveFunctions = []string{}
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "baa2abde")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "02751cec")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "2195995c")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "ded9382a")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "af2979eb")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "7d72d6e0")
	config.Sniper.ChiGasEnabled = false
	config.Sniper.AdvancedLiquiditySniping = false
	network := structs.Network{}
	network.RPC = "http://localhost:8545"
	network.Id = 0
	network.Name = "Local"
	config.Chains = append(config.Chains, network)

	exchange := structs.Exchanges{}
	exchange.ChainId = 1
	exchange.Name = "Dummy"
	exchange.Router = "0xdummya6f5b5e8d8c5b9a8f8c8c8d8d8c5b9a8f8c"
	exchange.Factory = "0xdummya6f5b5e8d8c5b9a8f8c8c8d8d8c5b9a8f8c"
	exchange.Init_code = "get init_code from pairfor function of router contract search its source code and get it from there"
	config.Exchanges = append(config.Exchanges, exchange)

	private_key := structs.PrivateKey{}
	private_key.Key = "private_key_here"
	config.PrivateKeys = append(config.PrivateKeys, private_key.Key)

	file, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}
	file.Close()
	return &config, nil
}

func Update_and_save_new_config(conff structs.BaseConfig) {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := structs.BaseConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	config.Chains = conff.Chains
	config.Exchanges = conff.Exchanges
	config.PrivateKeys = conff.PrivateKeys
	config.Sniper = structs.SniperSettings{}
	config.Sniper.Slippage = conff.Sniper.Slippage
	config.Sniper.MultiWallet = conff.Sniper.MultiWallet
	config.Sniper.MinLiq = conff.Sniper.MinLiq
	config.Sniper.BuySplit = conff.Sniper.BuySplit
	config.Sniper.SellSplit = conff.Sniper.SellSplit
	config.Sniper.HoneyPotCheck = conff.Sniper.HoneyPotCheck
	config.Sniper.LiquidityAddFunctions = conff.Sniper.LiquidityAddFunctions
	config.Sniper.LiquidityRemoveFunctions = conff.Sniper.LiquidityRemoveFunctions
	config.Sniper.ChiGasEnabled = conff.Sniper.ChiGasEnabled
	config.Sniper.GasLimit = conff.Sniper.GasLimit
	// config.Sniper.
	// config.Sniper = conff.Sniper
	file, err = os.Create("config.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}
	file.Close()

}

func ParseArgsAndConfig() *structs.BaseConfig {
	/* Flasg are gonna be:
	--token : string "Token to buy"
	--pair : string "To buy token in a different pair"
	--watch : bool "Run sniper in watch mode"
	--maxbt : float "Maximum Buy Tax"
	--maxst : float "Maximum Sell Tax"
	--minliq : float "Minimum Liquidity"
	--slippage : float "Slippage"
	--multiwallet : bool "Multi wallet"
	--buysplit : bool "Buy split"
	--sellsplit : bool "Sell split"
	--hpc : bool "Honeypot Check"
	--mlc : bool "Mempool Liquidity Check"
	--mrlc : bool "Mempool Remove Liquidity Check"
	--dtc : bool "Disable Trading Check"
	*/
	configf, err := Read_config_from_file("config.json")
	if err != nil {
		panic(err)
	}
	// token := flag.String("token", "0x0", "Token to buy")
	help := flag.Bool("help", false, "Show Help")
	pair := flag.String("pair", "0x0", "To buy token in a different pair")
	watch := flag.Bool("watch", false, "Run sniper in watch mode")
	maxbt := flag.Float64("maxbt", 0.0, "Maximum Buy Tax")
	maxst := flag.Float64("maxst", 0.0, "Maximum Sell Tax")
	skipblocks := flag.Int("skipblocks", 0, "Number of blocks to Skip After Liquidity is added")
	minliq := flag.Float64("minliq", float64(configf.Sniper.MinLiq), "Minimum Liquidity")
	slippage := flag.Int("slippage", configf.Sniper.Slippage, "Slippage")
	multiwallet := flag.Bool("multiwallet", configf.Sniper.MultiWallet, "Multi wallet")
	multiwalletmain := flag.Bool("multiwalletmain", configf.Sniper.MutliMainSnipe, "Multi wallet main")
	buysplit := flag.Int("buysplit", configf.Sniper.BuySplit, "Buy split")
	sellsplit := flag.Int("sellsplit", configf.Sniper.SellSplit, "Sell split")
	hpc := flag.Bool("hpc", configf.Sniper.HoneyPotCheck, "Honeypot Check")
	mlc := flag.Bool("mlc", false, "Mempool Liquidity Check")
	mrlc := flag.Bool("mrlc", false, "Mempool Remove Liquidity Check")
	dtc := flag.Bool("dtc", false, "Disable Trading Check")
	network := flag.String("network", "bsc", "Blockchain Network to use")
	exchange := flag.String("exchange", "", "Exchange to Use from the network!")
	onlybuy := flag.Bool("onlybuy", false, "Only buy")
	onlysell := flag.Bool("onlysell", false, "Only sell")
	chigasenabled := flag.Bool("chigas", configf.Sniper.ChiGasEnabled, "To Use ChiGas In tx's")
	takeprofit := flag.Float64("takeprofit", 0, "Take Profit")
	stoploss := flag.Float64("stoploss", 0, "Stop Loss")
	buyprice := flag.Float64("buyprice", 0.0, "Specific Buy Price Of token")
	sellprice := flag.Float64("sellprice", 0.0, "Specific Sell Price Of token")
	gasprice := flag.Int("gasprice", configf.Sniper.GasPrice, "Gas Price")
	gaslimit := flag.Int("gaslimit", configf.Sniper.GasLimit, "Gas Limit")
	freefire := flag.Bool("freefire", false, "Fire Snipes till one of the tx succeeds")
	freefirelimit := flag.Int("fflimit", 0, "Fire Snipes till one of the tx succeeds")
	dreservechecks := flag.Bool("dreservechecks", false, "Disable Reserve Checks") //for tests!!
	NumberOfMultipleWallets := flag.Int("nwm", 10, "Number Of Multiple Wallets")
	// deploy_contract := flag.Bool("deploysniper", false, "To deploy the sniper contract")
	flag.Usage = Showhelp
	flag.Parse()
	token := flag.Arg(0)
	amount := flag.Arg(1)
	// Some checks
	if *help == true {
		Showhelp()
		os.Exit(0)
	}
	if *stoploss > 0 {
		stp := *stoploss
		stp = (stp - float64(100))
		stoploss = &stp
	}
	snipersettings := structs.SniperSettings{}
	// Save all the parsed parameters in SniperSettings
	snipersettings.Token = common.HexToAddress(token)
	snipersettings.Amount, _ = strconv.ParseFloat(amount, 64)
	snipersettings.Pair = common.HexToAddress(*pair)
	snipersettings.Watch = *watch
	snipersettings.MaxBuyTax = *maxbt
	snipersettings.MaxSellTax = *maxst
	snipersettings.BlocksToSkip = *skipblocks
	snipersettings.MinLiq = *minliq
	snipersettings.Slippage = *slippage
	snipersettings.MultiWallet = *multiwallet
	snipersettings.MutliMainSnipe = *multiwalletmain
	snipersettings.BuySplit = *buysplit
	snipersettings.SellSplit = *sellsplit
	snipersettings.HoneyPotCheck = *hpc
	snipersettings.MempoolLiquidityCheck = *mlc
	snipersettings.MempoolRemoveLiquidityCheck = *mrlc
	snipersettings.AdvancedLiquiditySniping = configf.Sniper.AdvancedLiquiditySniping
	snipersettings.DisableTradingCheck = *dtc
	snipersettings.Network = *network
	snipersettings.Exchange = *exchange
	snipersettings.ChiGasEnabled = *chigasenabled
	snipersettings.LiquidityAddFunctions = configf.Sniper.LiquidityAddFunctions
	snipersettings.LiquidityRemoveFunctions = configf.Sniper.LiquidityRemoveFunctions
	snipersettings.TakeProfitPercentage = *takeprofit
	snipersettings.StopLossPercentage = *stoploss
	snipersettings.BuyPrice = *buyprice
	snipersettings.SellPrice = *sellprice
	snipersettings.DisableReserveChecks = *dreservechecks
	snipersettings.Freefire = *freefire
	snipersettings.FreeFireLimit = *freefirelimit
	snipersettings.GasPrice = *gasprice
	snipersettings.GasLimit = *gaslimit
	snipersettings.NumberOfMultipleWallets = *NumberOfMultipleWallets
	if *onlybuy == false && *onlysell == false {
		snipersettings.ToBuy = true
		snipersettings.ToSell = true
	}
	if *onlybuy == true {
		snipersettings.ToBuy = true
		snipersettings.ToSell = false
	}
	if *onlysell == true {
		snipersettings.ToBuy = false
		snipersettings.ToSell = true
	}
	if *onlybuy == true && *onlysell == true {
		utils.ColorPrint("[ERROR] You can't use Onlybuy and Only Sell at the same time! if you wana both buy and sell dont pass any of it", "red")
		os.Exit(0)
	}
	configf.Sniper = snipersettings
	return configf
}

// func
