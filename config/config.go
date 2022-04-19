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

func Read_config_from_file(f string) (*structs.BaseConfig, error) ***REMOVED***
	// if config file does not exist create one
	if _, err := os.Stat(f); os.IsNotExist(err) ***REMOVED***
		utils.ColorPrint("[+] Config file not found, creating one Edit and restart the bot!", "green")
		config, err := Create_a_new_Config()
		os.Exit(0)
		return config, err
	***REMOVED***
	file, err := os.Open(f)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := structs.BaseConfig***REMOVED******REMOVED***
	err = decoder.Decode(&config)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return &config, nil
***REMOVED***

func Create_a_new_Config() (*structs.BaseConfig, error) ***REMOVED***
	config := structs.BaseConfig***REMOVED******REMOVED***
	config.Chains = []structs.Network***REMOVED******REMOVED***
	config.Exchanges = []structs.Exchanges***REMOVED******REMOVED***
	config.PrivateKeys = []string***REMOVED******REMOVED***
	config.Sniper = structs.SniperSettings***REMOVED******REMOVED***
	config.Sniper.Slippage = 5
	config.Sniper.MultiWallet = false
	config.Sniper.MinLiq = 0
	config.Sniper.BuySplit = 1
	config.Sniper.SellSplit = 1
	config.Sniper.HoneyPotCheck = false
	// ***REMOVED***"f305d719","e8e33700"***REMOVED***
	config.Sniper.LiquidityAddFunctions = []string***REMOVED******REMOVED***
	config.Sniper.LiquidityAddFunctions = append(config.Sniper.LiquidityAddFunctions, "f305d719")
	config.Sniper.LiquidityAddFunctions = append(config.Sniper.LiquidityAddFunctions, "e8e33700")
	// "baa2abde","02751cec","2195995c","ded9382a","af2979eb","7d72d6e0"
	config.Sniper.LiquidityRemoveFunctions = []string***REMOVED******REMOVED***
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "baa2abde")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "02751cec")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "2195995c")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "ded9382a")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "af2979eb")
	config.Sniper.LiquidityRemoveFunctions = append(config.Sniper.LiquidityRemoveFunctions, "7d72d6e0")
	config.Sniper.ChiGasEnabled = false
	config.Sniper.AdvancedLiquiditySniping = false
	network := structs.Network***REMOVED******REMOVED***
	network.RPC = "http://localhost:8545"
	network.Id = 0
	network.Name = "Local"
	config.Chains = append(config.Chains, network)

	exchange := structs.Exchanges***REMOVED******REMOVED***
	exchange.ChainId = 1
	exchange.Name = "Dummy"
	exchange.Router = "0xdummya6f5b5e8d8c5b9a8f8c8c8d8d8c5b9a8f8c"
	exchange.Factory = "0xdummya6f5b5e8d8c5b9a8f8c8c8d8d8c5b9a8f8c"
	exchange.Init_code = "get init_code from pairfor function of router contract search its source code and get it from there"
	config.Exchanges = append(config.Exchanges, exchange)

	private_key := structs.PrivateKey***REMOVED******REMOVED***
	private_key.Key = "private_key_here"
	config.PrivateKeys = append(config.PrivateKeys, private_key.Key)

	file, err := os.Create("config.json")
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	file.Close()
	return &config, nil
***REMOVED***

func Update_and_save_new_config(conff structs.BaseConfig) ***REMOVED***
	file, err := os.Open("config.json")
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := structs.BaseConfig***REMOVED******REMOVED***
	err = decoder.Decode(&config)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	config.Chains = conff.Chains
	config.Exchanges = conff.Exchanges
	config.PrivateKeys = conff.PrivateKeys
	config.Sniper = conff.Sniper
	file, err = os.Create("config.json")
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	file.Close()

***REMOVED***

func ParseArgsAndConfig() *structs.BaseConfig ***REMOVED***
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
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
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
	gasprice := flag.Int("gasprice", 0, "Gas Price")
	gaslimit := flag.Int("gaslimit", 0, "Gas Limit")
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
	if *help == true ***REMOVED***
		Showhelp()
		os.Exit(0)
	***REMOVED***
	if *stoploss > 0 ***REMOVED***
		stp := *stoploss
		stp = (stp - float64(100))
		stoploss = &stp
	***REMOVED***
	snipersettings := structs.SniperSettings***REMOVED******REMOVED***
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
	if *onlybuy == false && *onlysell == false ***REMOVED***
		snipersettings.ToBuy = true
		snipersettings.ToSell = true
	***REMOVED***
	if *onlybuy == true ***REMOVED***
		snipersettings.ToBuy = true
		snipersettings.ToSell = false
	***REMOVED***
	if *onlysell == true ***REMOVED***
		snipersettings.ToBuy = false
		snipersettings.ToSell = true
	***REMOVED***
	if *onlybuy == true && *onlysell == true ***REMOVED***
		utils.ColorPrint("[ERROR] You can't use Onlybuy and Only Sell at the same time! if you wana both buy and sell dont pass any of it", "red")
		os.Exit(0)
	***REMOVED***
	configf.Sniper = snipersettings
	return configf
***REMOVED***

// func
