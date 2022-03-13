// This Module is for setting up the MainClient Struct and adding functions to it for ease of usability.
package sniper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logicfool/GoSniper/config"
	"github.com/logicfool/GoSniper/contracts"
	"github.com/logicfool/GoSniper/utils"
	// "github.com/logicfool/GoSniper/config"
)

type (
	// Exchanges is for baseconfig to read the data and then its converted to EXChange when setting up the main config...
	Exchange struct ***REMOVED***
		Name    string
		Router  *contracts.Unirouter
		Factory *contracts.Unifactory
		Honey   *contracts.HoneyPot
	***REMOVED***

	MainClient struct ***REMOVED***
		Client      *ethclient.Client
		Exchanges   []Exchange
		WalletAuths []*bind.TransactOpts
	***REMOVED***
)

func GetClient(RPC string) *ethclient.Client ***REMOVED***
	client, err := ethclient.Dial(RPC)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return client
***REMOVED***

func GetAccountAuth(client *ethclient.Client, accountAddress string, chainId int) *bind.TransactOpts ***REMOVED***
	var chainid *big.Int
	// #TODO: Get private key from config
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok ***REMOVED***
		panic("invalid key")
	***REMOVED***

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	fmt.Println("nounce=", nonce)

	if chainId == 0 ***REMOVED***
		chainiid, err := client.ChainID(context.Background())
		if err != nil ***REMOVED***
			panic(err)
		***REMOVED***
		chainid = chainiid
	***REMOVED*** else ***REMOVED***
		chainid = big.NewInt(int64(chainId))
	***REMOVED***
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(100000000000000000) //utils.EtherToWei(0.1) // in wei
	// auth.GasLimit = uint64(10000000)
	// auth.GasPrice = big.NewInt(1000000000) // in wei
	return auth
***REMOVED***

// Taking in the normal config and creating a MainConf with it

func BaseConfToMainConf(config *config.BaseConfig, chainid int) (*MainClient, error) ***REMOVED***
	main_config := MainClient***REMOVED******REMOVED***
	for _, network := range config.Chains ***REMOVED***
		if network.Id == chainid ***REMOVED***
			main_config.Client = GetClient(network.RPC)
			// get all exchanges in config.exchanges and create exchange structs
			for _, exchange := range config.Exchanges ***REMOVED***
				if exchange.ChainId == network.Id ***REMOVED***
					exchange_struct := Exchange***REMOVED******REMOVED***
					exchange_struct.Router = contracts.GetRouter(utils.ToChecksumAddress(exchange.Router), main_config.Client)
					exchange_struct.Factory = contracts.GetFactory(utils.ToChecksumAddress(exchange.Factory), main_config.Client)
					exchange_struct.Name = exchange.Name
					// exchange_struct.Honey = contracts.GetHoneyPot(utils.ToChecksumAddress(exchange.HoneyPot), main_config.Client)
					main_config.Exchanges = append(main_config.Exchanges, exchange_struct)
				***REMOVED***
			***REMOVED***
			// get all private keys in config.private_keys and create wallet auths
			for _, private_key := range config.PrivateKeys ***REMOVED***
				wallet_auth := GetAccountAuth(main_config.Client, private_key, chainid)
				main_config.WalletAuths = append(main_config.WalletAuths, wallet_auth)
			***REMOVED***

			break
		***REMOVED***
	***REMOVED***
	return &main_config, nil
***REMOVED***

func QuickMaintainConfig(chainId int) *MainClient ***REMOVED***
	config, err := config.Read_config_from_file("config.json")
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	main_config, err := BaseConfToMainConf(config, chainId)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return main_config
***REMOVED***

// Dispatching new methods in MainClient
func (config *MainClient) GetClient() *ethclient.Client ***REMOVED***
	return config.Client
***REMOVED***

func (config *MainClient) GetWalletAuth(accountAddress string) *bind.TransactOpts ***REMOVED***
	for _, wallet_auth := range config.WalletAuths ***REMOVED***
		if wallet_auth.From.Hex() == accountAddress ***REMOVED***
			return wallet_auth
		***REMOVED***
	***REMOVED***
	return nil
***REMOVED***

func (config *MainClient) UpdateAllWalletNonces() ***REMOVED***
	for _, wallet_auth := range config.WalletAuths ***REMOVED***
		nonce, err := config.Client.PendingNonceAt(context.Background(), wallet_auth.From)
		if err != nil ***REMOVED***
			panic(err)
		***REMOVED***

		wallet_auth.Nonce = big.NewInt(int64(nonce))
	***REMOVED***
***REMOVED***

func (config *MainClient) GetExchange(exchangeName string) *Exchange ***REMOVED***
	for _, exchange := range config.Exchanges ***REMOVED***
		if exchange.Name == exchangeName ***REMOVED***
			return &exchange
		***REMOVED***
	***REMOVED***
	return nil
***REMOVED***

func (config *MainClient) UpdateWalletAuthValues(value interface***REMOVED******REMOVED***) ***REMOVED***
	for _, wallet_auth := range config.WalletAuths ***REMOVED***
		switch value.(type) ***REMOVED***
		case *big.Int:
			wallet_auth.Value = value.(*big.Int)
		case *int:
			wallet_auth.Value = big.NewInt(int64(*value.(*int)))
		***REMOVED***
	***REMOVED***
***REMOVED***

func (config *MainClient) WeiToEther(args ...interface***REMOVED******REMOVED***) *big.Float ***REMOVED***
	// Supported Args: Wei,Decimals
	lenarg := len(args)
	if lenarg > 2 ***REMOVED***
		panic("Too many arguments for WeiToEther")
	***REMOVED*** else if lenarg < 1 ***REMOVED***
		panic("Too few arguments for WeiToEther")
	***REMOVED***
	if lenarg == 1 ***REMOVED***
		return utils.WeiToEther(args[0])
	***REMOVED***
	return utils.WeiToEtherByDecimals(args[0], args[1].(int))
***REMOVED***

func (config *MainClient) EtherToWei(args ...interface***REMOVED******REMOVED***) *big.Int ***REMOVED***
	// Supported Args: Ether,Decimals
	lenarg := len(args)
	if lenarg > 2 ***REMOVED***
		panic("Too many arguments for EtherToWei")
	***REMOVED*** else if lenarg < 1 ***REMOVED***
		panic("Too few arguments for EtherToWei")
	***REMOVED***
	if lenarg == 1 ***REMOVED***
		return utils.EtherToWei(args[0])
	***REMOVED***
	return utils.EtherToWeiByDecimals(args[0], args[1].(int))
***REMOVED***

func (config *MainClient) BuyTokenWithETH(tokenAddress interface***REMOVED******REMOVED***, amount interface***REMOVED******REMOVED***, exchangeName interface***REMOVED******REMOVED***) ***REMOVED***
	var haspath bool = false
	switch tokenAddress.(type) ***REMOVED***
	case string:
		tokenAddress = utils.ToChecksumAddress(tokenAddress.(string))
	case []common.Address:
		tokenAddress = tokenAddress.([]common.Address)
		haspath = true
	***REMOVED***
	// make sure amount is already in wei and not in ether
	fmt.Println(haspath)
***REMOVED***
