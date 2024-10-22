// This Module is for setting up the MainClient Struct and adding functions to it for ease of usability.
package sniper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/logicfool/GoSniper/config"
	parseconfig "github.com/logicfool/GoSniper/config"
	"github.com/logicfool/GoSniper/consts"
	"github.com/logicfool/GoSniper/contracts"
	"github.com/logicfool/GoSniper/structs"
	"github.com/logicfool/GoSniper/utils"
	// "github.com/logicfool/GoSniper/config"
)

type (
	MainClient struct {
		Client       *ethclient.Client
		Rpc          *rpc.Client
		MainExchange structs.Exchange
		Exchanges    []structs.Exchange
		WalletAuths  []*bind.TransactOpts
		Sniper       *structs.SniperSettings
		Session      structs.CurrentSnipeSession
		PrivateKeys  []string
		Network      *structs.Network
		WaitGroup    *sync.WaitGroup
		Config       structs.BaseConfig
	}
)

var (
	Round        int                      = 1
	mutex                                 = &sync.Mutex{}
	Noncehandler map[common.Address]int64 = make(map[common.Address]int64)
)

func GetClient(RPC string) (*ethclient.Client, *rpc.Client) {
	client, err := ethclient.Dial(RPC)
	rpc_cli, err := rpc.DialContext(context.Background(), RPC)
	if err != nil {
		utils.ColorPrint(fmt.Sprintf("[-] Error : %v, Mempool Features Will be off as it requires a WSS url", err), "red")
		rpc_cli = nil
	}
	return client, rpc_cli
}

func GetAccountAuth(client *ethclient.Client, accountAddress string, chainId int, gasprice int, gaslimit int) *bind.TransactOpts {
	var chainid *big.Int
	// #TODO: Get private key from config
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	// fmt.Println("nounce=", nonce)

	if chainId == 0 {
		chainiid, err := client.ChainID(context.Background())
		if err != nil {
			panic(err)
		}
		chainid = chainiid
	} else {
		chainid = big.NewInt(int64(chainId))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	Noncehandler[fromAddress] = int64(nonce)
	auth.Value = big.NewInt(100000000000000000) //utils.EtherToWei(0.1) // in wei
	if gasprice != 0 {
		auth.GasPrice = utils.EtherToWeiByDecimals(gasprice, 9)
	}
	if gaslimit != 0 {
		auth.GasLimit = uint64(gaslimit)
	}
	// auth.GasLimit = uint64(10000000)
	// auth.GasPrice = big.NewInt(1000000000) // in wei
	return auth
}

// Taking in the normal config and creating a MainConf with it
func BaseConfToMainConf(config *structs.BaseConfig, chainid int) (*MainClient, error) {
	main_config := MainClient{}
	var sniper_contract *contracts.GoSniper
	var weth common.Address
	var err error
	var WETH *contracts.WETH
	var rawurll string
	for _, network := range config.Chains {
		if network.Id == chainid {
			main_config.Client, main_config.Rpc = GetClient(network.RPC)
			rawurll = network.RPC
			// main_config.ChainId = *big.NewInt(int64(network.Id))
			main_config.Network = &network
			if network.BuyContract != "" {
				sniper_contract = contracts.GetContract(utils.ToChecksumAddress(network.BuyContract), main_config.Client, "sniper").(*contracts.GoSniper)
			} else {
				sniper_contract = nil
			}
			// get all exchanges in config.exchanges and create exchange structs
			for _, exchange := range config.Exchanges {
				if exchange.ChainId == network.Id {
					exchange_struct := structs.Exchange{}
					exchange_struct.Router = contracts.GetContract(utils.ToChecksumAddress(exchange.Router), main_config.Client, "router").(*contracts.Unirouter)
					exchange_struct.Factory = contracts.GetContract(utils.ToChecksumAddress(exchange.Factory), main_config.Client, "factory").(*contracts.Unifactory)
					exchange_struct.Sniper = sniper_contract
					exchange_struct.RouterAddress = utils.ToChecksumAddress(exchange.Router)
					exchange_struct.FactoryAddress = utils.ToChecksumAddress(exchange.Factory)
					if sniper_contract != nil {
						exchange_struct.SniperAddress = utils.ToChecksumAddress(network.BuyContract)
						exchange_struct.Init_code = exchange.Init_code
					} else {
						exchange_struct.SniperAddress = common.HexToAddress("0x0")
						exchange_struct.Init_code = ""
					}
					if (weth == common.Address{}) {
						if exchange_struct.SniperAddress != common.HexToAddress("0x0") {
							weth, err = exchange_struct.Sniper.WETH(&bind.CallOpts{})
							if err != nil {
								panic(err)
							}
						} else {
							weth, err = exchange_struct.Router.WETH(&bind.CallOpts{})
							if err != nil {
								panic(err)
							}
						}
					}
					if WETH == nil {
						WETH = contracts.GetContract(weth, main_config.Client, "weth").(*contracts.WETH)
					}
					exchange_struct.WETH = WETH
					exchange_struct.WETHAddress = weth
					exchange_struct.WETH = contracts.GetContract(utils.ToChecksumAddress(exchange.Factory), main_config.Client, "weth").(*contracts.WETH)
					exchange_struct.Name = exchange.Name
					main_config.Exchanges = append(main_config.Exchanges, exchange_struct)
					if exchange.Name == config.Sniper.Exchange {
						main_config.MainExchange = exchange_struct
						utils.ColorPrint(fmt.Sprintf("Exchange Selected: %v", main_config.MainExchange.Name), "yellow")
					}
				}
			}
			// if main_config.Exchanges is empty error out
			// get all private keys in config.private_keys and create wallet auths
			main_config.Config = *config
			break
		}
	}

	main_config.Sniper = &config.Sniper
	if main_config.Sniper.GasPrice > 0 {
		main_config.Network.GasPrice = main_config.Sniper.GasPrice
	}
	if main_config.Sniper.GasLimit > 0 {
		main_config.Network.GasLimit = main_config.Sniper.GasLimit
	}
	for _, private_key := range config.PrivateKeys {
		wallet_auth := GetAccountAuth(main_config.Client, private_key, int(main_config.Network.Id), main_config.Network.GasPrice, main_config.Network.GasLimit)
		main_config.WalletAuths = append(main_config.WalletAuths, wallet_auth)
	}
	main_config.Session = structs.CurrentSnipeSession{}
	main_config.Session.Token = main_config.Sniper.Token
	main_config.PrivateKeys = config.PrivateKeys
	if main_config.Sniper.ToBuy == true {
		main_config.Session.Buy = true
	}
	if main_config.Sniper.ToSell == true {
		main_config.Session.Sell = true
	}
	main_config.Session.AmountBought = big.NewInt(int64(0))
	main_config.Session.AllPairs = make(map[common.Address]map[common.Address]common.Address)
	main_config.Session.AllBalanceOfUserTokens = make(map[int]map[common.Address]map[common.Address]*big.Int)
	main_config.Session.RouterABI, _ = abi.JSON(strings.NewReader(contracts.UnirouterABI))
	main_config.Session.GasPrice = utils.EtherToWeiByDecimals(main_config.Network.GasPrice, 9)
	main_config.Session.GasLimit = uint64(main_config.Network.GasLimit)
	u, err := url.Parse(rawurll)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http", "https":
		main_config.Sniper.MempoolLiquidityCheck = false
		main_config.Sniper.MempoolRemoveLiquidityCheck = false
		main_config.Sniper.AdvancedLiquiditySniping = false
		utils.ColorPrint("[-] Error cannot use Liquidity Sniping Without a WSS connection!", "red")
	}
	main_config.WaitGroup = &sync.WaitGroup{}
	return &main_config, nil
}

func FindNetworkByname(name *string, configf *structs.BaseConfig) structs.Network {
	for _, network := range configf.Chains {
		if network.Name == *name {
			return network
		}
	}
	// return first network
	return configf.Chains[0]
}

func VerifyMainConf(config *MainClient) {

	if config.Sniper.Token.Hex() == common.HexToAddress("0x0").Hex() {
		utils.ColorPrint(fmt.Sprintf("[*] Token Not Set! Please pass the token then the amount!"), "red")
		os.Exit(1)
	}

	if config.Sniper.Amount == 0 {
		utils.ColorPrint(fmt.Sprintf("[*] Amount Not Passed! Exiting"), "red")
		os.Exit(1)

	}
	// fmt.Println(common.HexToAddress("0x0").Hex(), config.Sniper.Token.Hex())
	if len(config.WalletAuths) == 0 {
		utils.ColorPrint(fmt.Sprintf("[*] No Wallets in the config please add your private key in the config file!"), "red")
	}

	exchangeName := config.Sniper.Exchange
	if len(config.Exchanges) == 0 {
		utils.ColorPrint(fmt.Sprintf("[*] No exchanges found in config for the selected network Add some and restart!"), "red")
		// exit the program
		os.Exit(1)
	}
	if exchangeName == "" {
		config.MainExchange = config.Exchanges[0]
		utils.ColorPrint(fmt.Sprintf("[-] No Exchange passed defaulting to: %v", config.MainExchange.Name), "yellow")
	}
	// check if mainexchange is defined
	if config.MainExchange.Name == "" {
		config.MainExchange = config.Exchanges[0]
		utils.ColorPrint(fmt.Sprintf("[-] No Exchange with that name found! Defaultihng to: %v", config.MainExchange.Name), "yellow")
		// os.Exit(1)
	}
	// check if mainexchange.sniper is defined
	if config.MainExchange.Sniper == nil {
		utils.ColorPrint(fmt.Sprintf("[*] No Sniper Contract found for the current network: %v\nBot will try to use the router! Please Deploy a contract and then set it in the config to use contract sniping!", ""), "yellow")
		// os.Exit(1)
	}
	// Initating and saving the contracts fo token and pair in memory to be used later
	config.GetTokenContract(config.Sniper.Token)

	if config.Sniper.Pair.Hex() != config.MainExchange.WETHAddress.Hex() && config.Sniper.Pair.Hex() != common.HexToAddress("0x0").Hex() {
		config.GetTokenContract(config.Sniper.Pair)
	} else {
		config.Sniper.Pair = config.MainExchange.WETHAddress
	}
	// if sniper argument for number_of_multiple_wallets is not equal to the number provided in the config/parameter use wallet_helpers function to create a dn update new config!
	if config.Sniper.MultiWallet {
		if config.Sniper.NumberOfMultipleWallets != len(config.PrivateKeys) && config.Sniper.NumberOfMultipleWallets > len(config.PrivateKeys) {
			num_to_gen := config.Sniper.NumberOfMultipleWallets - len(config.PrivateKeys)
			config.GenerateWallets(num_to_gen)
		}
		parseconfig.Update_and_save_new_config(config.Config)
		// parse all the private keys and save them in the config then update the config again!
	}
}

func QuickMaintainConfig() *MainClient {
	utils.ColorPrint("[+] Starting Sniper", "green")
	utils.ColorPrint("[+] Loading config", "green")
	// snipersettings := config.ParseArgs()
	configf := config.ParseArgsAndConfig()
	// configf, err := config.Read_config_from_file()
	// merge config settings
	Network := FindNetworkByname(&configf.Sniper.Network, configf)
	main_config, err := BaseConfToMainConf(configf, Network.Id)
	VerifyMainConf(main_config)

	var wallets string
	for _, wallet := range main_config.WalletAuths {
		wallets += fmt.Sprintf("%v, ", wallet.From.Hex())
	}
	utils.ColorPrint(fmt.Sprintf("[+] Loaded Wallets: %v", wallets), "green")
	utils.ColorPrint(fmt.Sprintf("[+] Choosen Network: %v", Network.Name), "green")

	if err != nil {
		panic(err)
	}
	return main_config
}

/*
Important Utility Functions For Use declared Below ------------------------------------------------------------------------
*/

func (client *MainClient) GetClient() *ethclient.Client {
	return client.Client
}

func (client *MainClient) GetWalletAuth(accountAddress string) *bind.TransactOpts {
	for _, wallet_auth := range client.WalletAuths {
		if wallet_auth.From.Hex() == accountAddress {
			return wallet_auth
		}
	}
	return nil
}

func (client *MainClient) GetTokenContract(address interface{}) *contracts.Token {
	var tokAddress common.Address
	switch address.(type) {
	case common.Address:
		tokAddress = address.(common.Address)
	case string:
		tokAddress = common.HexToAddress(address.(string))

	}
	for _, token := range client.Session.AllTokens {
		if token.Address.Hex() == tokAddress.Hex() {
			return token
		}
	}
	contract := contracts.GetContract(tokAddress, client.Client, "erc20").(*contracts.Token)
	client.Session.AllTokens = append(client.Session.AllTokens, contract)
	return contract
}

// func (client *MainClient) UpdateWalletNonces(wallet common.Address) {
// 	if wallet.Hex() == common.HexToAddress("0x0").Hex() {
// 		for i, wallet_auth := range client.WalletAuths {
// 			nonce, err := client.Client.NonceAt(context.Background(), wallet_auth.From, nil)
// 			if err != nil {
// 				utils.ColorPrint(fmt.Sprintf("[-] Error getting nonce for wallet: %v", wallet.Hex()), "red")
// 			}

// 			wallet_auth.Nonce = big.NewInt(int64(nonce))
// 			client.WalletAuths[i] = wallet_auth
// 		}
// 		return
// 	}
// 	for i, wallet_auth := range client.WalletAuths {
// 		if wallet_auth.From.Hex() == wallet.Hex() {
// 			nonce := wallet_auth.Nonce.Add(wallet_auth.Nonce, big.NewInt(1))
// 			wallet_auth.Nonce = nonce
// 			client.WalletAuths[i] = wallet_auth
// 		}
// 	}
// }

func (client *MainClient) GetWalletbalance(wallet interface{}) *big.Int {
	var waddress common.Address
	switch wallet.(type) {
	case *bind.TransactOpts:
		waddress = wallet.(*bind.TransactOpts).From
	case common.Address:
		waddress = wallet.(common.Address)
	case string:
		waddress = common.HexToAddress(wallet.(string))
	}
	balance, err := client.Client.BalanceAt(context.Background(), waddress, nil)
	if err != nil {
		panic(err)
	}
	return balance
}

func (client *MainClient) GetExchange(exchangeName interface{}) *structs.Exchange {
	switch exchangeName.(type) {
	case string:
		for _, exchange := range client.Exchanges {
			if exchange.Name == exchangeName {
				return &exchange
			}
		}
		return nil
	case structs.Exchange:
		return exchangeName.(*structs.Exchange)
	}
	panic("Error Loading the Exchange!")
}

func (client *MainClient) UpdateWalletAuthValues(value interface{}) []*bind.TransactOpts {
	var wallet_auths []*bind.TransactOpts
	var gasPrice *big.Int
	var gasLimit uint64
	if client.Session.GasPrice == nil {
		gasPrice, _ = client.Client.SuggestGasPrice(context.Background())
	} else {
		gasPrice = client.Session.GasPrice
	}
	if client.Session.GasLimit != 0 {
		gasLimit = client.Session.GasLimit
	}

	// fmt.Println(gasPrice)
	// os.Exit(1)
	if client.Session.GasLimit == 0 {
		gasLimit = 0
	} else {
		gasLimit = client.Session.GasLimit
	}
	for _, wallet_auth := range client.WalletAuths {
		nonce, _ := client.Client.NonceAt(context.Background(), wallet_auth.From, nil)
		// wallet_auth.Nonce = big.NewInt(int64(nonce))
		Noncehandler[wallet_auth.From] = int64(nonce)
		switch value.(type) {
		case *big.Int:
			wallet_auth.Value = value.(*big.Int)
		case *int:
			wallet_auth.Value = big.NewInt(int64(*value.(*int)))
		case *int64:
			wallet_auth.Value = big.NewInt(int64(*value.(*int64)))
		}
		wallet_auth.GasPrice = gasPrice
		wallet_auth.GasLimit = gasLimit
		if client.Sniper.Pair.Hex() == common.HexToAddress("0x0").Hex() {
			wallet_balance := client.GetWalletbalance(wallet_auth.From.Hex())
			if wallet_balance.Cmp(wallet_auth.Value) < 0 {
				utils.ColorPrint(fmt.Sprintf("[-] Insufficient funds for %v", wallet_auth.From.Hex()), "red")
				continue
			}
		} else {

		}

		wallet_auths = append(wallet_auths, wallet_auth)
		if !client.Sniper.MultiWallet {
			break
		}
	}
	return wallet_auths
}

func (client *MainClient) WeiToEther(args ...interface{}) *big.Float {
	// Supported Args: Wei,Decimals
	lenarg := len(args)
	if lenarg > 2 {
		panic("Too many arguments for WeiToEther")
	} else if lenarg < 1 {
		panic("Too few arguments for WeiToEther")
	}
	if lenarg == 1 {
		return utils.WeiToEther(args[0])
	}
	return utils.WeiToEtherByDecimals(args[0], args[1].(int))
}

func (client *MainClient) EtherToWei(args ...interface{}) *big.Int {
	// Supported Args: Ether,Decimals
	lenarg := len(args)
	if lenarg > 2 {
		panic("Too many arguments for EtherToWei")
	} else if lenarg < 1 {
		panic("Too few arguments for EtherToWei")
	}
	if lenarg == 1 {
		return utils.EtherToWei(args[0])
	}
	return utils.EtherToWeiByDecimals(args[0], args[1].(int))
}

/*
Methods to Be used to Calculate and get data from chain and Contracts declared Below -------------------------------------------------------------------
*/

func (client *MainClient) SaveBalances() {
	for _, wallet_auth := range client.WalletAuths {
		ethbalance := client.GetWalletbalance(wallet_auth.From.Hex())
		mutex.Lock()
		if client.Session.AllBalanceOfUserTokens[Round] == nil {
			client.Session.AllBalanceOfUserTokens[Round] = make(map[common.Address]map[common.Address]*big.Int)
		}
		if client.Session.AllBalanceOfUserTokens[Round][wallet_auth.From] == nil {
			client.Session.AllBalanceOfUserTokens[Round][wallet_auth.From] = make(map[common.Address]*big.Int)
		}
		mutex.Unlock()
		mutex.Lock()
		client.Session.AllBalanceOfUserTokens[Round][wallet_auth.From][common.HexToAddress("0x0")] = ethbalance
		mutex.Unlock()
		balanceoftoken, err := client.GetTokenContract(client.Sniper.Token).Contract.BalanceOf(nil, wallet_auth.From)
		// fmt.Println(client.Sniper.Token, balanceoftoken)
		if err != nil {
			utils.ColorPrint(fmt.Sprintf("[-] Error in getting balance of Token %v for user %v with error: %v", client.Sniper.Token, wallet_auth.From, err), "red")
		}
		mutex.Lock()
		// fmt.Println(client.Sniper.Token, balanceoftoken)
		client.Session.AllBalanceOfUserTokens[Round][wallet_auth.From][client.Sniper.Token] = balanceoftoken
		mutex.Unlock()
		balanceofpair, err := client.GetTokenContract(client.Sniper.Pair).Contract.BalanceOf(nil, wallet_auth.From)
		if err != nil {
			utils.ColorPrint(fmt.Sprintf("[-] Error in getting balance of Token %v for user %v with error: %v", client.Sniper.Pair, wallet_auth.From, err), "red")
		}
		mutex.Lock()
		client.Session.AllBalanceOfUserTokens[Round][wallet_auth.From][client.Sniper.Pair] = balanceofpair
		mutex.Unlock()
	}
}

// Check and Confirm the allowance of a token to be swapped
func (client *MainClient) HandleAllowance(token *contracts.Token, Wallets_to_use []*bind.TransactOpts) {
	tokencon := token.Contract
	for _, wallet := range Wallets_to_use {
		allowance, err := tokencon.Allowance(&bind.CallOpts{}, wallet.From, client.MainExchange.SniperAddress)
		if err != nil {
			utils.ColorPrint(fmt.Sprintf("Error in getting allowance of %v ! You sure its a token? Please Restart to retry.", token.Name), "red")
			os.Exit(0)
		}
		// utils.ColorPrint(fmt.Sprintf("Allowance of %v for %v is Sufficient", token.Name, wallet.From.Hex()), "blue")
		// if allowance is less trhen consts.MAX_INT_ALLOWANCE then we need to approve
		var alamount *big.Int
		if client.Session.AmountToTrade == nil {
			alamount = utils.EtherToWei(1000)
		} else {
			alamount = client.Session.AmountToTrade
		}
		if allowance.Cmp(alamount) < 0 {
			wallet.Value = big.NewInt(0)
			utils.ColorPrint(fmt.Sprintf("Wallet %v Approving %v for trade on %v", wallet.From.Hex(), client.MainExchange.SniperAddress.Hex(), token.Name), "yellow")
			tx, err := tokencon.Approve(wallet, client.MainExchange.SniperAddress, consts.MAX_APPROVAL_INT)
			if err != nil {
				utils.ColorPrint(fmt.Sprintf("Error in Approving %v for trade on %v", wallet.From.Hex(), token.Name), "red")
				continue
			}
			rec := client.GetTxReceipt(tx.Hash().Hex())
			if rec.Status == types.ReceiptStatusSuccessful {
				utils.ColorPrint(fmt.Sprintf("Wallet %v Approved %v for trade on %v", wallet.From.Hex(), client.MainExchange.SniperAddress.Hex(), token.Name), "green")
			}
			if err != nil {
				panic(err)
			} else {
				wallet.Nonce = wallet.Nonce.Add(wallet.Nonce, big.NewInt(1))
			}
		}
	}
}

// A wrapper around Router's GetAmountsOut methods to get the approximate amount we might get when swapped
func (client *MainClient) GetAmountsOut(path []common.Address, amount *big.Int, args ...interface{}) []*big.Int {
	var res []*big.Int
	var err error
	var minamountout *big.Int
	exchange := client.MainExchange
	res, err = exchange.Router.GetAmountsOut(&bind.CallOpts{}, amount, path)
	if err != nil {
		panic(err)
	}
	if args != nil {
	}
	minamountout = new(big.Int).Set(res[1])
	slippage := client.Sniper.Slippage
	minamountout.Mul(minamountout, big.NewInt(int64(slippage)))
	minamountout.Div(minamountout, big.NewInt(int64(100)))
	minamountout.Sub(res[1], minamountout)
	ress := append([]*big.Int{}, res...)
	ress = append(ress, minamountout)
	return ress
}

// Function To Pull the Tx info from a  Tx Hash
func (client *MainClient) GetTxbyHash(hash interface{}) (*types.Transaction, bool) {
	var txhash common.Hash
	switch hash.(type) {
	case string:
		txhash = common.HexToHash(hash.(string))
	case common.Hash:
		txhash = hash.(common.Hash)
	}

	tx, pending, err := client.Client.TransactionByHash(context.Background(), txhash)
	for err != nil {
		// fmt.Printf("\rError in getting TX : %v Retrying", err)
	}
	if pending {
	}
	return tx, pending
}

// To get the Receipt of the Tx if it was successfull or not
func (client *MainClient) GetTxReceipt(hash string) *types.Receipt {
	receipt, err := client.Client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	fmt.Print("\033[s")
	for err != nil {
		receipt, err = client.Client.TransactionReceipt(context.Background(), common.HexToHash(hash))
		if err != nil {
			// mess := fmt.Sprintf("\r[-] Error in getting TX : %v Still Pending", err)
			// utils.ColorPrint(mess, "red", false)
			continue
		}
		// fmt.Print("\033[u")
	}
	return receipt
}

func (client *MainClient) GetTxReceiptArray(alltxs map[common.Address]string) (bool, common.Address) {
	onesuccesfful := false
	var byuser common.Address = common.Address{}
	for address, txhash := range alltxs {
		receipt := client.GetTxReceipt(txhash)
		if receipt.Status == types.ReceiptStatusSuccessful {

			utils.ColorPrint(fmt.Sprintf("[+] Transaction %v was Successful", txhash), "green")
			onesuccesfful = true
			byuser = address
		} else {
			utils.ColorPrint(fmt.Sprintf("[-] Transaction %v was Failed", txhash), "red")
		}
	}

	return onesuccesfful, byuser

}

func (client *MainClient) GetBuySellStatus(txh common.Hash, Address common.Address, typeoftx string) {
	receipt := client.GetTxReceipt(txh.Hex())
	if receipt.Status == types.ReceiptStatusSuccessful {
		utils.ColorPrint(fmt.Sprintf("[+] Transaction %v was Successful", txh.Hex()), "green")
		if typeoftx == "buy" {
			// client.Session.BuySuccess = true
			client.BuyAfterVerify(Address)
		} else {
			client.Session.SellSuccess = true
			client.SellAfterVerify()
		}
		client.Session.SuccessAddress = Address

	} else {
		utils.ColorPrint(fmt.Sprintf("[-] Transaction %v was Failed", txh.Hex()), "red")
	}
}

// Get the Pair of the contract using TokenA/TokenB and Init_Code if provided or fallback to normal method by calling Factory
func (client *MainClient) GetPair(tokenA common.Address, tokenB common.Address) common.Address {
	var pair_address common.Address
	var err error
	token1, token2, err := utils.SortTokens(tokenA, tokenB)

	// if pair available in client.Session.AllPairs
	if client.Session.AllPairs[token1] != nil {
		pair_address = client.Session.AllPairs[token1][token2]
		if pair_address != (common.Address{}) {
			return pair_address
		}
	} else {
		client.Session.AllPairs[token1] = make(map[common.Address]common.Address)
	}
	if client.MainExchange.Init_code != "" && client.MainExchange.Init_code != "0x" {
		// pair_address, err = client.MainExchange.Sniper.PairFor(&bind.CallOpts{}, client.MainExchange.FactoryAddress, token1, token2, client.MainExchange.Init_code)
		pair_address, err = utils.GetPair(client.MainExchange.FactoryAddress, token1, token2, client.MainExchange.Init_code)
		if err != nil {
			panic(err)
		}
		client.Session.AllPairs[token1][token2] = pair_address
		return pair_address
	}
	utils.ColorPrint("Init Code Not Provided for exchange in config Reverting to Default Method", "yellow")
	pair_address, err = client.MainExchange.Factory.GetPair(&bind.CallOpts{}, token1, token2)
	if err != nil {
		panic(err)
	}
	client.Session.AllPairs[token1][token2] = pair_address
	return pair_address
}

func (client *MainClient) GetCurrentPairPrice() *big.Float {
	if client.Network.DataFeed == "" {
		utils.ColorPrint("[-] No Data Feed Found for this network Price check not possible", "red")
		return nil
	}
	var path [][]common.Address
	if client.Session.Token.Hex() == client.Sniper.Token.Hex() {
		path = append(path, []common.Address{client.Sniper.Pair, client.Sniper.Token})
	} else {
		path = append(path, []common.Address{client.Sniper.Token, client.Sniper.Pair})
	}
	ethprice, allprices, err := client.MainExchange.Sniper.GetPrices(&bind.CallOpts{}, client.MainExchange.RouterAddress, common.HexToAddress(client.Network.DataFeed), path)
	if err != nil {
		utils.ColorPrint(fmt.Sprintf("Error in Fetching Token Prices Will retry! Error:%v", err), "red")
		return nil
	}
	// fix maths issue
	ethpricef := utils.BigIntToBigFloat(ethprice)
	// divide ethpricef by 10**18
	ethpricef = new(big.Float).Quo(ethpricef, big.NewFloat(math.Pow10(18)))
	// fmt.Println(ethpricef)
	pathprice := allprices[0]
	decimals := pathprice[1].Int64() // fmt.Println(decimals)
	if decimals == 0 {
		utils.ColorPrint("Error in Fetching Token Prices Will retry!", "red")
		return nil
	}
	tokenprice := utils.BigIntToBigFloat(pathprice[0])
	// divide token price by 10**decimals
	// tokenprice = tokenprice.Div(tokenprice, big.NewInt(10).Exp(big.NewInt(10), decimals, nil))
	// tokenprice = new(big.Float).Quo(tokenprice, big.NewFloat(math.Pow10(int(decimals))))
	tokenprice = utils.WeiToEtherByDecimals(tokenprice, int(decimals))
	return new(big.Float).Quo(ethpricef, tokenprice)
}

func (client *MainClient) CalculateBuyPrice(Wallet common.Address) {
	var amountinwallet *big.Int
	var amountspent *big.Int
	var err error
	ethprice, err := client.GetETHPriceFromFeed()
	if err != nil {
		utils.ColorPrint(fmt.Sprintf("Error in Fetching MainCoin Price From Feed you sure the address is correct in config?! Error:%v", err), "red")
		client.Session.PriceBought = big.NewFloat(0.0)
		return
	}
	amountoftokensbought := utils.WeiToEtherByDecimals(utils.BigIntToBigFloat(client.Session.AmountBought), int(client.GetTokenContract(client.Sniper.Token).Decimals.Int64()))
	if client.Sniper.Pair.Hex() == client.MainExchange.WETHAddress.Hex() {
		amountinwallet = client.GetWalletbalance(Wallet)
		balbefore := client.Session.AllBalanceOfUserTokens[Round][Wallet][common.HexToAddress("0x0")]
		amountspent = new(big.Int).Sub(balbefore, amountinwallet)
	} else {
		// convert it into eth!
		var res []*big.Int
		balbefore := client.Session.AllBalanceOfUserTokens[Round][Wallet][client.Sniper.Pair]
		amountinwallet, err = client.GetTokenContract(client.Sniper.Pair).Contract.BalanceOf(&bind.CallOpts{}, Wallet)
		if err != nil {
			utils.ColorPrint(fmt.Sprintf("[-] Error in fetching token balance %v And calculating Buy Price!", err), "red")
			return
		}
		amountspent = new(big.Int).Sub(balbefore, amountinwallet)
		path := []common.Address{client.Sniper.Pair, client.MainExchange.WETHAddress}
		res, err = client.MainExchange.Router.GetAmountsOut(&bind.CallOpts{}, amountspent, path)
		if err != nil {
			utils.ColorPrint(fmt.Sprintf("[-] Error in fetching token balance %v And calculating Buy Price!", err), "red")
			return
		}
		amountspent = res[1]
	}

	ethpricef := utils.BigIntToBigFloat(ethprice)
	ethpricef = ethpricef.Mul(ethpricef, utils.BigIntToBigFloat(amountspent))
	ethpricef = new(big.Float).Quo(ethpricef, big.NewFloat(math.Pow10(18)))
	ethpricef = new(big.Float).Quo(ethpricef, big.NewFloat(math.Pow10(18)))
	pricet := new(big.Float).Quo(ethpricef, amountoftokensbought)
	client.Session.PriceBought = pricet
}

// Function To get mempool data and process it
func (client *MainClient) ProcessMempool(txh common.Hash) {
	// defer client.WaitGroup.Done()
	tx, _ := client.GetTxbyHash(txh)
	data := tx.Data()
	if len(data) == 0 {
		return
	}
	if cap(data) < 4 {
		return
	}
	methodName, err := client.Session.RouterABI.MethodById(data[:4])
	if err != nil {
		return
	}
	// fmt.Println("[+] TX Hash:", txh.Hex())

	// encoded_func := hexutil.Encode(data[:4])
	if client.Sniper.MempoolLiquidityCheck {
		if strings.ToLower(methodName.Name[:12]) == "addliquidity" {
			// if utils.ArrayContainsString(client.Sniper.LiquidityAddFunctions, encoded_func) {
			var token common.Address = common.HexToAddress("0x0")
			var token1 common.Address = common.HexToAddress("0x0")
			var amount0 = new(big.Int).SetInt64(0)
			var amount1 = new(big.Int).SetInt64(0)
			if err != nil {
				utils.ColorPrint("[-] Error in getting ABI", "red")
			}
			decodedInputs, err := utils.DecodeTxParams(client.Session.RouterABI, data)
			if err != nil {
				return
			}
			if decodedInputs["token"] != nil {
				if decodedInputs["token"].(common.Address) == client.Session.Token {
					token = decodedInputs["token"].(common.Address)
					amount0 = decodedInputs["amountTokenDesired"].(*big.Int)
				} else {
					token1 = decodedInputs["token"].(common.Address)
					amount1 = decodedInputs["amountTokenDesired"].(*big.Int)
				}
			}
			if decodedInputs["tokenA"] != nil {
				// token = decodedInputs["tokenA"].(common.Address)
				if decodedInputs["tokenA"].(common.Address).Hex() == client.Session.Token.Hex() {
					token = decodedInputs["tokenA"].(common.Address)
					amount0 = decodedInputs["amountADesired"].(*big.Int)
				} else {
					token1 = decodedInputs["tokenA"].(common.Address)
					amount1 = decodedInputs["amountADesired"].(*big.Int)
				}
			}
			if decodedInputs["tokenB"] != nil {
				// token = decodedInputs["tokenB"].(common.Address)
				if decodedInputs["tokenB"].(common.Address).Hex() == client.Session.Token.Hex() {
					token = decodedInputs["tokenB"].(common.Address)
					amount0 = decodedInputs["amountBDesired"].(*big.Int)
				} else {
					token1 = decodedInputs["tokenB"].(common.Address)
					amount1 = decodedInputs["amountBDesired"].(*big.Int)
				}
			}

			if token1.Hex() == common.HexToAddress("0x0").Hex() {
				token1 = client.MainExchange.WETHAddress
				amount1 = tx.Value()
			}
			if token.Hex() != client.Session.Token.Hex() {
				token2 := token1
				token1 = token
				token = token2
			}
			// utils.ColorPrint(fmt.Sprintf("1. Found Liquidity Addition for %v Tx Hash: %vtx/%v ", token, client.Network.Explorer, txh), "yellow")
			if token.Hex() != client.Session.Token.Hex() && token1.Hex() != client.Session.Token.Hex() {
				return
			}

			if client.Sniper.AdvancedLiquiditySniping {
				sorttoken1, _, _ := utils.SortTokens(token, token1)
				if sorttoken1 == token {
					client.Session.Reserve0 = amount0
					client.Session.Reserve1 = amount1
				} else {
					client.Session.Reserve0 = amount1
					client.Session.Reserve1 = amount0
				}
			}
			// client.Session.GasPrice = big.NewInt(int64(float64(tx.GasPrice().Int64())))
			client.Session.GasPrice = tx.GasPrice()
			client.Session.GasLimit = uint64(float64(tx.Gas()) * 1.1)
			client.Session.LiquidityAdded = true
			utils.ColorPrint(fmt.Sprintf("Found Liquidity Addition for %v Tx Hash: %vtx/%v ", token, client.Network.Explorer, txh), "blue")
			go client.BuyTokenWithPair()
		}
	}

	if client.Sniper.MempoolRemoveLiquidityCheck {
		// if utils.ArrayContainsString(client.Sniper.LiquidityRemoveFunctions, encoded_func) {
		if strings.ToLower(methodName.Name[:15]) == "removeliquidity" {
			var token common.Address = common.HexToAddress("0x0")
			var token1 common.Address = common.HexToAddress("0x0")
			var token2 common.Address = common.HexToAddress("0x0")
			if err != nil {
				utils.ColorPrint("[-] Error in getting ABI", "red")
			}
			decodedInputs, err := utils.DecodeTxParams(client.Session.RouterABI, data)
			if err != nil {
				utils.ColorPrint(fmt.Sprintf("[-] Error in decoding inputs %v", err), "red")
				return
			}
			if decodedInputs["token"] != nil {
				token = decodedInputs["token"].(common.Address)
			}
			if decodedInputs["tokenA"] != nil {
				token1 = decodedInputs["tokenA"].(common.Address)
			}
			if decodedInputs["tokenB"] != nil {
				token2 = decodedInputs["tokenB"].(common.Address)
			}
			temp_addy := client.Session.Token
			if token1.Hex() == temp_addy.Hex() || token2.Hex() == temp_addy.Hex() || token.Hex() == temp_addy.Hex() {
				// Liquidity Remove Detected!
				client.Session.LiquidityRemoved = true
				// client.ProcessLiquidityRemoveSell()
			}
			// client.Session.LiquidityRemoved = true
		}
	}
	// Quit <- true
}

/*
	Trading Methods declared Below to verify and make trades with the sniper -------------------------------------------------
*/

func (client *MainClient) GetandSaveReserves() {
	pair_contract := contracts.GetContract(client.GetPair(client.Sniper.Token, client.Sniper.Pair), client.Client, "pair").(*contracts.Pair)
	for {
		if client.Session.Reserve0 == nil || client.Session.Reserve1 == nil {
			data, err := pair_contract.GetReserves(&bind.CallOpts{})
			// fmt.Println("/r[-] Reached Here Delete this line later!", client.GetPair(client.Sniper.Token, client.Sniper.Pair))
			if err != nil {
				// utils.ColorPrint("\r[-] No Contract Found at Pair Probs its not yet deployed or not created! Waiting...", "red")
				fmt.Print("\033[G\033[K")
				utils.ColorPrint(fmt.Sprintf("[-] Error: %v Pair: %v", err, client.GetPair(client.Sniper.Token, client.Sniper.Pair)), "red")
				fmt.Print("\033[A")
				continue
			}
			if data.Reserve1.Cmp(big.NewInt(0)) == 0 || data.Reserve0.Cmp(big.NewInt(0)) == 0 {
				fmt.Print("\033[G\033[K")
				utils.ColorPrint("[-] 0 Reserves Found in Liquidity Pair Contract ! Waiting for to be fullfilled...", "yellow")
				fmt.Print("\033[A")
				if client.Session.LiquidityAdded == false {
					continue
				}

			} else {
				// client.Session.GasPrice = utils.EtherToWeiByDecimals(50, 9)
				// client.Session.GasLimit = uint64(3000000)
				go client.BuyTokenWithPair()
				client.Session.LiquidityAdded = true
				client.Session.Reserve0 = data.Reserve0
				client.Session.Reserve1 = data.Reserve1

			}
		}
	}
}

// Function To Verify the rules and conditions before buying
func (client *MainClient) BuyPreVerify() {
	// pair_contract := contracts.GetContract(client.GetPair(client.Sniper.Token, client.Sniper.Pair), client.Client, "pair").(*contracts.Pair)
	var liqadd_check_done bool = false

	// _, _, _, _, _, _, _, _ = pair_contract, liqadd_check_done, ishp, client.Sniper.Pair, token1, eth0, eth1, err
	// fmt.Println(client.GetPair(client.Sniper.Token, client.Sniper.Pair))
	// data
	// fmt.Println(client.GetPair(client.Sniper.Token, client.Sniper.Pair), client.Sniper.Token, client.Sniper.Pair)
	for {
		// Buy Price Checks

		for {
			if client.Sniper.BuyPrice > 0 {
				currprice := client.GetCurrentPairPrice()
				if currprice.Cmp(big.NewFloat(client.Sniper.BuyPrice)) >= 0 {
					utils.ColorPrint("[+] Buy Price Reached! Buying!", "green")
				}
				mess := fmt.Sprintf("[+] To Buy Price: %v, Current Price: %v", client.Sniper.BuyPrice, currprice)
				fmt.Print("\033[G\033[K")
				utils.ColorPrint(mess, "blue")
				fmt.Print("\033[A")
				continue
			}
			break
		}

		// if advancebuy skip this checks
		if liqadd_check_done == false {
			if !client.Sniper.DisableReserveChecks {

			}
		}

		// if client.Sniper.MinLiq {
		// }

		if client.Session.LiquidityAdded == true {
			break
		}
	}

}

func (client *MainClient) SellPreVerify(wallets []*bind.TransactOpts) []*bind.TransactOpts {
	// pair_contract := contracts.GetContract(client.GetPair(client.Sniper.Token, client.Sniper.Pair), client.Client, "pair").(*contracts.Pair)
	token_contract := client.GetTokenContract(client.Sniper.Pair)
	var walls_to_return []*bind.TransactOpts = []*bind.TransactOpts{}
	for _, wallet := range wallets {
		bal, _ := token_contract.Contract.BalanceOf(&bind.CallOpts{}, wallet.From)
		if bal.Cmp(big.NewInt(0)) >= 1 {
			walls_to_return = append(walls_to_return, wallet)
		}
	}
	for {
		for {
			if client.Sniper.TakeProfitPercentage > 0 || client.Sniper.StopLossPercentage > 0 || client.Sniper.SellPrice > 0 {
				currprice := client.GetCurrentPairPrice()
				if client.Session.AmountBought == nil {
					client.Session.PriceBought = currprice
				}
				percentagedifference := big.NewFloat(0.0).Quo(currprice, client.Session.PriceBought)
				percentagedifference = big.NewFloat(0.0).Mul(percentagedifference, big.NewFloat(100.0))
				percentagedifference = big.NewFloat(0.0).Sub(percentagedifference, big.NewFloat(100.0))
				if client.Sniper.TakeProfitPercentage > 0 {
					if percentagedifference.Cmp(new(big.Float).SetFloat64(client.Sniper.TakeProfitPercentage)) == 1 {
						utils.ColorPrint("[-] Take Profit is reached! Selling Now!", "green")
						break
					}
				}
				if client.Sniper.StopLossPercentage > 0 {
					if percentagedifference.Cmp(new(big.Float).SetFloat64(client.Sniper.StopLossPercentage)) == -1 {
						utils.ColorPrint("[-] Stop Loss is reached! Selling Now!", "red")
						break
					}
				}
				if client.Sniper.SellPrice > 0 {
					if currprice.Cmp(new(big.Float).Set(big.NewFloat(client.Sniper.SellPrice))) == 1 {
						utils.ColorPrint("[-] Sell Price target is reached! Selling Now!", "green")
						break
					}
				}
				message := fmt.Sprintf("Profit: %v%%| Token Price: %v| Price Bought: %v", percentagedifference, currprice, client.Session.PriceBought)
				fmt.Print("\033[G\033[K")
				utils.ColorPrint(message, "blue")
				fmt.Print("\033[A")
				continue
			}
			break
		}

		// if Client.Sniper.MinLiq{}
		break

	}
	return walls_to_return
}

func (client *MainClient) BuyAfterVerify(bywallet common.Address) {
	if client.Session.BuySuccess != true {
		client.Session.BuySuccess = true
		token := client.GetTokenContract(client.Session.Token)
		balnow, _ := token.Contract.BalanceOf(&bind.CallOpts{}, bywallet)
		balbefore := client.Session.AllBalanceOfUserTokens[Round][bywallet][token.Address]
		totalAmountReceived := new(big.Int).Sub(balnow, balbefore)
		client.Session.AmountBought = totalAmountReceived
		fmt.Println("Total Amount received: ", client.Session.AmountBought)
		client.Sniper.ToBuy = false
		// client.Session.BuySuccess = true
		utils.ColorPrint("Buy(s) Successful", "yellow")
		client.CalculateBuyPrice(bywallet)
		Round++
	}
}

func (client *MainClient) SellAfterVerify() {
	if client.Session.SellSuccess != true {
		client.Session.Sell = false
		client.Sniper.ToSell = false
		client.Session.SellSuccess = true
		utils.ColorPrint("[+] Sell(s) Successful!! Exiting", "green")
	}
}

// Method to CrossVerify and wait for all conditions to be satisified to trade or perform a swap
func (client *MainClient) PreTxCheckandGetWallets() []*bind.TransactOpts {
	var amount *big.Int
	var wallets_to_use []*bind.TransactOpts
	client.GetTokenContract(client.Sniper.Token)
	// if (client.Sniper.Pair.Hex() != client.MainExchange.WETHAddress.Hex() && client.Sniper.Token.Hex() != client.MainExchange.WETHAddress.Hex()) && (client.Sniper.Pair.Hex() != common.HexToAddress("0x0").Hex() && client.Sniper.Token.Hex() != common.HexToAddress("0x0").Hex()) {
	// 	client.GetTokenContract(client.Sniper.Pair) // initalize once for faster use later
	// 	if client.Sniper.HoneyPotCheck {
	// 		// utils.ColorPrint("[-] Honeypot Checking with pair! To make sure honeypot checker works id have to allow the pair token on the contract Checking.... ", "yellow")
	// 		utils.ColorPrint("[-] Honeypot Checking/Tax Detection/Trading Disabled wont work with custom pair!", "red")
	// 		client.Sniper.HoneyPotCheck = false
	// 		client.Sniper.MaxBuyTax = 0.0
	// 		client.Sniper.MaxSellTax = 0.0
	// 	}
	// } else {
	// if client.Sniper.Pair is emppty set it as WETH
	if client.Sniper.Pair.Hex() == common.HexToAddress("0x0").Hex() {
		client.Sniper.Pair = client.MainExchange.WETHAddress
		// }
	}
	if client.Sniper.ChiGasEnabled {
		wallets_to_use = client.UpdateWalletAuthValues(big.NewInt(0))
		chigas := client.GetTokenContract(consts.CHIGASTOKEN)
		client.HandleAllowance(chigas, wallets_to_use)
	}
	go client.SaveBalances()
	pair_to_use := client.GetPair(client.Sniper.Token, client.Sniper.Pair)
	client.Session.Pair = pair_to_use

	var majoraddresses []common.Address
	majoraddresses = append(majoraddresses, client.MainExchange.RouterAddress)
	majoraddresses = append(majoraddresses, client.MainExchange.FactoryAddress)
	majoraddresses = append(majoraddresses, pair_to_use)
	client.Session.MajorAddresses = majoraddresses
	if client.Sniper.HoneyPotCheck || client.Sniper.DisableTradingCheck || client.Sniper.MaxBuyTax > 0 || client.Sniper.MaxSellTax > 0 {
		var ExtraSettings []*big.Int
		if client.Sniper.HoneyPotCheck || client.Sniper.DisableTradingCheck {
			ExtraSettings = append(ExtraSettings, big.NewInt(1))
		} else {
			ExtraSettings = append(ExtraSettings, big.NewInt(0))
		}
		if client.Sniper.MaxBuyTax > 0 {
			ExtraSettings = append(ExtraSettings, utils.FloatToBigInt(client.Sniper.MaxBuyTax))
		} else {
			ExtraSettings = append(ExtraSettings, big.NewInt(0))
		}
		if client.Sniper.MaxSellTax > 0 {
			ExtraSettings = append(ExtraSettings, utils.FloatToBigInt(client.Sniper.MaxSellTax))
		} else {
			ExtraSettings = append(ExtraSettings, big.NewInt(0))
		}
		client.Session.ExtraSettings = ExtraSettings
	} else {
		var ExtraSettings []*big.Int
		ExtraSettings = append(ExtraSettings, big.NewInt(0))
		ExtraSettings = append(ExtraSettings, big.NewInt(0))
		ExtraSettings = append(ExtraSettings, big.NewInt(0))
		client.Session.ExtraSettings = ExtraSettings
	}

	if client.Sniper.ToBuy == true {
		if client.Session.Buy {
			client.Session.Split = client.Sniper.BuySplit
			if client.Sniper.Amount != float64(0.0) {
				if client.Sniper.Pair.Hex() == client.MainExchange.WETHAddress.Hex() {
					amount = client.EtherToWei(client.Sniper.Amount)
					client.Session.AmountToTrade = amount
					client.Session.Wallets_to_use = client.UpdateWalletAuthValues(amount)
				} else {
					amount = utils.FloatToBigInt(client.Sniper.Amount)
					client.Session.PairContract = client.GetTokenContract(client.Sniper.Pair)
					client.Session.TokenContract = client.GetTokenContract(client.Sniper.Token)
					decimals := client.Session.TokenContract.Decimals
					amount.Mul(amount, big.NewInt(int64(math.Pow10(int(decimals.Int64())))))
					client.Session.AmountToTrade = amount
					client.Session.Wallets_to_use = client.UpdateWalletAuthValues(0)
					// wallets_to_use = client.UpdateWalletAuthValues(0)
					client.HandleAllowance(client.Session.PairContract, client.Session.Wallets_to_use)
				}
			} else {
				utils.ColorPrint("[-] Amount is not set", "red")
				os.Exit(1)
			}
			path := client.Session.Path
			path = append(path, client.Sniper.Pair)
			path = append(path, client.Sniper.Token)
			client.Session.Path = path
			// client.Session.Path
			if !client.Sniper.DisableReserveChecks {
				go client.GetandSaveReserves()
			}
			for {
				if client.Session.BuySuccess {
					break
				}
			}
		}
	}
	token1 := client.Sniper.Token
	client.Sniper.Token = client.Sniper.Pair
	client.Sniper.Pair = token1
	if client.Sniper.ToSell == true {
		wallets_to_use = client.UpdateWalletAuthValues(big.NewInt(int64(0)))
		client.SaveBalances()
		client.Session.Split = client.Sniper.SellSplit
		client.Session.PairContract = client.GetTokenContract(client.Sniper.Pair)
		client.Session.TokenContract = client.GetTokenContract(client.Sniper.Token)
		if client.Session.AmountBought.Cmp(big.NewInt(int64(0))) > 0 {
			client.Session.AmountToTrade = client.Session.AmountBought
		}
		if client.Session.Sell {
			utils.ColorPrint("[-] Sell Activated!", "red")
			// wallets_to_use = client.UpdateWalletAuthValues(big.NewInt(int64(0)))
			client.HandleAllowance(client.Session.PairContract, wallets_to_use)
			client.Session.Wallets_to_use = client.UpdateWalletAuthValues(big.NewInt(int64(0)))
		}
		client.Session.Wallets_to_use = client.SellPreVerify(wallets_to_use)
		client.BuyTokenWithPair()
		for {
			if client.Session.SellSuccess {
				break
			}
		}
		// return wallets_to_use
	}
	// os.Exit(0)
	return nil
}

// The Main Buy Function that Trades with the Swap Function Directly
func (client *MainClient) BuyTokenWithPair() {
	// token1 := client.Sniper.Token
	// token0 := client.Sniper.Pair
	Exchange := client.MainExchange
	// Sniper := Exchange.Sniper
	wallets_to_use := client.Session.Wallets_to_use
	// if Sniper == nil {
	// 	utils.ColorPrint("[*] No Sniper Contract Found! Sending the tx via router!", "red")
	// 	os.Exit(0)
	// }

	// var path []common.Address
	// path = append(path, token0)
	// path = append(path, token1)
	path := client.Session.Path
	var limit map[common.Address]int = make(map[common.Address]int)
	// tokenb := client.GetTokenContract(path[0])

	if client.Sniper.Pair == Exchange.WETHAddress {
		// if len(wallets_to_use) == 0 {
		// 	utils.ColorPrint("[*] No Wallets to use! Exiting", "red")
		// 	os.Exit(1)
		// }
		var amountout *big.Int = big.NewInt(0)
		if client.Sniper.Slippage > 0 {
			val := client.EtherToWei(client.Sniper.Amount)
			amountoutres := client.GetAmountsOut(path, val)
			amountout = amountoutres[len(amountoutres)-1]
		}
		go func() {
			for _, wallet := range wallets_to_use {
				go func(wallet *bind.TransactOpts) {
					tomul := 1
					var tx *types.Transaction
					var err error
					var balbefore *big.Int
					var toaddress common.Address
					mutex.Lock()
					limit[wallet.From] = 1000000000000000000
					mutex.Unlock()
					if client.Sniper.FreeFireLimit > 0 && client.Sniper.Freefire {
						limit[wallet.From] = client.Sniper.FreeFireLimit
					}
					// var rawtx *types.Transaction
					for {
						if client.Session.BuySuccess && client.Sniper.MutliMainSnipe {
							return
						}
						if limit[wallet.From] == 0 {
							return
						}
						toaddress = wallet.From
						balbefore = client.Session.AllBalanceOfUserTokens[Round][wallet.From][client.Sniper.Token]
						if client.Sniper.MutliMainSnipe {
							toaddress = wallets_to_use[0].From
							balbefore = client.Session.AllBalanceOfUserTokens[Round][toaddress][client.Sniper.Token]
						}
						// wallet.GasPrice = utils.EtherToWeiByDecimals(5, 9) //client.Session.GasPrice
						wallet.GasPrice = client.Session.GasPrice
						wallet.Nonce = big.NewInt(Noncehandler[wallet.From])
						// wallet.GasLimit = client.Session.GasLimit
						// }
						// utils.ColorPrint(fmt.Sprintf("%v, %v", toaddress, balbefore), "green")
						// fmt.Println("Nonce Of wallet ", wallet.From, ": ", Noncehandler[wallet.From])
						if client.Sniper.ChiGasEnabled {
							// Later Update with GASCHIFUNCTION AFTER NEW DEPLOYMENT
							// tx, err = Sniper.GoSniperTransactor.TradeDirectlyByPairWithETH(wallet, client.Session.MajorAddresses, path, toaddress, amountout, balbefore, client.Session.ExtraSettings, big.NewInt(int64(client.Session.Split)))
							tx = client.GetSignedRawTx(wallet, "TradeDirectlyByPairWithETH", client.Session.MajorAddresses, path, toaddress, amountout, balbefore, client.Session.ExtraSettings, big.NewInt(int64(client.Session.Split)))
							_, err = client.SendRawTx(tx)
						} else {
							tx = client.GetSignedRawTx(wallet, "TradeDirectlyByPairWithETH", client.Session.MajorAddresses, path, toaddress, amountout, balbefore, client.Session.ExtraSettings, big.NewInt(int64(client.Session.Split)))
							_, err = client.SendRawTx(tx)
						}
						if err != nil {
							// do proper formatting later on!
							if strings.Contains(err.Error(), "known") {
								mutex.Lock()
								Noncehandler[wallet.From]++
								mutex.Unlock()
								continue
							}
							if strings.Contains(err.Error(), "replacement") {
								mutex.Lock()
								Noncehandler[wallet.From]++
								mutex.Unlock()
								continue
							}
							if strings.Contains(err.Error(), "nonce") {
								mutex.Lock()
								Noncehandler[wallet.From]++
								mutex.Unlock()
								continue
							}
							if strings.Contains(err.Error(), " exceeds") {
								return
							}
							if strings.Contains(err.Error(), "insufficient") {
								return
							}
							utils.ColorPrint(fmt.Sprintf("Buy Error for %v Error: %v", wallet.From.Hex(), err.Error()), "red")
							continue
						}
						utils.ColorPrint(fmt.Sprintf("Buy Sent from %v Tx Hash: %v", wallet.From.Hex(), tx.Hash().Hex()), "green")
						go client.GetBuySellStatus(tx.Hash(), toaddress, "buy")
						// alltxhashes = append(alltxhashes, tx.Hash().Hex())
						// go client.UpdateWalletNonces(wallet.From)
						mutex.Lock()
						Noncehandler[wallet.From]++
						mutex.Unlock()
						// alltxhashes[wallet.From] = tx.Hash().Hex()

						go func() {
							if client.Session.GasPrice != nil {
								// client.Session.GasPrice = client.Session.GasPrice.Mul(client.Session.GasPrice, big.NewInt(1.1))
								client.Session.GasPrice = big.NewInt(client.Session.GasPrice.Int64() + int64(tomul))
								tomul = tomul + 100
							}
						}()
						// sleep for 3 seconds
						// time.Sleep(3 * time.Second)
						mutex.Lock()
						limit[wallet.From]--
						mutex.Unlock()
						if !client.Sniper.Freefire {
							return
						}
						if client.Session.BuySuccess && client.Sniper.MutliMainSnipe {
							return
						}
					}
				}(wallet)
				// fmt.Println(gasPrice)
				// os.Exit(1)
			}
		}()
	} else { //else if client.Session.LiquidityRemoved != true {
		var minamountout *big.Int = big.NewInt(0)
		if client.Sniper.Slippage > 0 {
			minamountoutres := client.GetAmountsOut(path, client.Session.AmountToTrade)
			minamountout = minamountoutres[len(minamountoutres)-1]
		}
		var tradableamount *big.Int = big.NewInt(0)
		// token_pair_contract := client.GetTokenContract(client.Sniper.Pair)
		func() {
			for {
				for _, wallet := range wallets_to_use {
					go func(wallet *bind.TransactOpts) {
						tomul := 1.0
						var tx *types.Transaction
						var err error
						var balbefore *big.Int
						var toaddress common.Address
						if client.Sniper.MutliMainSnipe {
							if client.Sniper.ToBuy {
								if client.Session.BuySuccess {
									return
								}
							} else {
								if client.Session.SellSuccess {
									return
								}
							}
						}
						toaddress = wallet.From
						if client.Sniper.Token == Exchange.WETHAddress {
							balbefore = client.Session.AllBalanceOfUserTokens[Round][wallet.From][common.HexToAddress("0x0")]
						} else {
							balbefore = client.Session.AllBalanceOfUserTokens[Round][wallet.From][client.Sniper.Token]
						}
						if client.Sniper.MutliMainSnipe {
							toaddress = wallets_to_use[0].From
							if client.Sniper.Token == Exchange.WETHAddress {
								balbefore = client.Session.AllBalanceOfUserTokens[Round][toaddress][common.HexToAddress("0x0")]
							} else {
								balbefore = client.Session.AllBalanceOfUserTokens[Round][toaddress][client.Sniper.Token]
							}
						}
						// fmt.Println(client.Session.AmountBought)
						if client.Sniper.ToBuy {
							tradableamount = client.Session.AmountToTrade
						} else if client.Session.AmountBought.Cmp(big.NewInt(0)) == 0 {
							amount := utils.FloatToBigInt(client.Sniper.Amount)

							// decimals := token_pair_contract.Decimals
							tradableamount, _ = client.Session.PairContract.Contract.BalanceOf(&bind.CallOpts{}, wallet.From)
							// fmt.Println("Balance Of use rof token: ", tradableamount)
							if amount.Cmp(big.NewInt(100)) != 0 {
								// get amount % of tradable amount
								tradableamount = tradableamount.Mul(tradableamount, amount)
								tradableamount = tradableamount.Div(tradableamount, big.NewInt(100))
								// tradableamount = tradableamount.Div(tradableamount, amount)
								// fmt.Println("trade Am: ", tradableamount, " Total Am: ", amount)
							}
						} else {
							// tradableamount, _ = token_pair_contract.Contract.BalanceOf(&bind.CallOpts{}, wallet.From)
							// if tradableamount.Cmp(client.Session.AmountToTrade) == 1 {
							tradableamount = client.Session.AmountToTrade
							// }
						}
						wallet.Value = big.NewInt(0)
						ball, _ := client.Session.PairContract.Contract.BalanceOf(&bind.CallOpts{}, wallet.From)
						wallet.GasPrice = client.Session.GasPrice
						wallet.Nonce = big.NewInt(Noncehandler[wallet.From])
						// wallet.GasPrice = utils.EtherToWeiByDecimals(20, 9) //client.Session.GasPrice
						// wallet.GasLimit = client.Session.GasLimit
						utils.ColorPrint(fmt.Sprintf("Balance of %v is %v of %v", wallet.From.Hex(), ball, client.Session.PairContract.Name), "green")
						if client.Sniper.ChiGasEnabled {
							tx, err = Exchange.Sniper.GoSniperTransactor.TradeDirectlyByPairWithGasRefund(wallet, client.Session.MajorAddresses, path, tradableamount, toaddress, minamountout, balbefore, client.Session.ExtraSettings, big.NewInt(int64(client.Session.Split)))
						} else {
							tx, err = Exchange.Sniper.GoSniperTransactor.TradeDirectlyByPair(wallet, client.Session.MajorAddresses, path, tradableamount, toaddress, minamountout, balbefore, client.Session.ExtraSettings, big.NewInt(int64(client.Session.Split)))
						}
						if err != nil {
							utils.ColorPrint(fmt.Sprintf("Tx Error for %v Error: %v", wallet.From.Hex(), err), "red")
							return
						}
						// go client.UpdateWalletNonces(wallet.From)
						Noncehandler[wallet.From] = Noncehandler[wallet.From] + 1
						go func() {
							if client.Session.GasPrice != nil {
								// client.Session.GasPrice = client.Session.GasPrice.Mul(client.Session.GasPrice, big.NewInt(1.1))
								client.Session.GasPrice = big.NewInt(int64(float64(client.Session.GasPrice.Int64()) * tomul))
								tomul = tomul + 0.1
							}
						}()

						if client.Sniper.ToBuy {
							go client.GetBuySellStatus(tx.Hash(), toaddress, "buy")
							utils.ColorPrint(fmt.Sprintf("Buy Sent from %v Tx Hash: %v", wallet.From.Hex(), tx.Hash().Hex()), "green")
						} else {
							go client.GetBuySellStatus(tx.Hash(), toaddress, "sell")
							utils.ColorPrint(fmt.Sprintf("Sell Sent from %v Tx Hash: %v", wallet.From.Hex(), tx.Hash().Hex()), "green")
						}
					}(wallet)
				}
				if !client.Sniper.Freefire {
					break
				}
				if client.Sniper.MutliMainSnipe {
					if client.Sniper.ToBuy {
						if client.Session.BuySuccess {
							return
						}
					} else {
						if client.Session.SellSuccess {
							return
						}
					}
				}
			}

		}()
	}
	// for {
	// 	if len(alltxhashes) > 1 {
	// 		break
	// 	}
	// 	// sleep for a second
	// 	time.Sleep(1 * time.Second)
	// }
	// if client.Session.Sell {

	// 	isonesuccessful, _ := client.GetTxReceiptArray(alltxhashes)
	// 	if isonesuccessful {
	// 		client.Session.Sell = false
	// 		client.Sniper.ToSell = false
	// 		client.Session.SellSuccess = true
	// 		utils.ColorPrint("[+] Sell(s) Successful!! Exiting", "green")
	// 	}

	// } else {
	// 	utils.ColorPrint("[*] Sell Failed! Retrying", "red")
	// }

}

// A function to use the advance liquidity guessing and sniping mechanism in order to get a lead But its riskier and unsafe..
func (client *MainClient) AdvanceUnsafeTrade() {

}
