package config

import (
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logicfool/GoSniper/contracts"

	"github.com/logicfool/GoSniper/sniper"
	"github.com/logicfool/GoSniper/utils"
)

type (
	Address string

	BaseConfig struct ***REMOVED***
		Chains      []Network   `json:"networks"`
		Exchanges   []Exchanges `json:"exchanges"`
		PrivateKeys []string    `json:"private_keys"`
	***REMOVED***

	Network struct ***REMOVED***
		RPC  string `json:"node_url"`
		Id   int    `json:"chain_id"`
		Name string `json:"name"`
	***REMOVED***
	Contracts struct ***REMOVED***
	***REMOVED***

	Exchanges struct ***REMOVED***
		ChainId  int    `json:"network"`
		Name     string `json:"name"`
		Router   string `json:"router"`
		Factory  string `json:"factory"`
		HoneyPot string `json:"honeypot"`
	***REMOVED***

	// Exchanges is for baseconfig to read the data and then its converted to EXChange when setting up the main config...
	Exchange struct ***REMOVED***
		Router  *contracts.Unirouter
		Factory *contracts.Unifactory
		Honey   *contracts.HoneyPot
	***REMOVED***

	PrivateKey struct ***REMOVED***
		Key string
	***REMOVED***

	MainConfig struct ***REMOVED***
		Client      *ethclient.Client
		Exchanges   []Exchange
		WalletAuths []*bind.TransactOpts
	***REMOVED***
)

func Read_config_from_file(f string) (*BaseConfig, error) ***REMOVED***
	file, err := os.Open(f)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := BaseConfig***REMOVED******REMOVED***
	err = decoder.Decode(&config)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return &config, nil
***REMOVED***

func BaseConfToMainConf(config *BaseConfig, chainid int) (*MainConfig, error) ***REMOVED***
	main_config := MainConfig***REMOVED******REMOVED***
	for _, network := range config.Chains ***REMOVED***
		if network.Id == chainid ***REMOVED***
			main_config.Client = sniper.GetClient(network.RPC)
			// get all exchanges in config.exchanges and create exchange structs
			for _, exchange := range config.Exchanges ***REMOVED***
				if exchange.ChainId == network.Id ***REMOVED***
					exchange_struct := Exchange***REMOVED******REMOVED***
					exchange_struct.Router = contracts.GetRouter(utils.ToChecksumAddress(exchange.Router), main_config.Client)
					exchange_struct.Factory = contracts.GetFactory(utils.ToChecksumAddress(exchange.Factory), main_config.Client)
					// exchange_struct.Honey = contracts.GetHoneyPot(utils.ToChecksumAddress(exchange.HoneyPot), main_config.Client)
					main_config.Exchanges = append(main_config.Exchanges, exchange_struct)
				***REMOVED***
			***REMOVED***
			// get all private keys in config.private_keys and create wallet auths
			for _, private_key := range config.PrivateKeys ***REMOVED***
				wallet_auth := sniper.GetAccountAuth(main_config.Client, private_key, chainid)
				main_config.WalletAuths = append(main_config.WalletAuths, wallet_auth)
			***REMOVED***

			break
		***REMOVED***
	***REMOVED***
	return &main_config, nil
***REMOVED***

func QuickMaintainConfig(chainId int) *MainConfig ***REMOVED***
	config, err := Read_config_from_file("config.json")
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	main_config, err := BaseConfToMainConf(config, chainId)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return main_config
***REMOVED***
