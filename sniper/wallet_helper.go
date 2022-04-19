package sniper

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (client *MainClient) GenerateWallets(num int) ***REMOVED***
	var all_wallets []*bind.TransactOpts
	for i := 0; i < num; i++ ***REMOVED***
		privateKey, err := crypto.GenerateKey()
		if err != nil ***REMOVED***
			log.Fatal(err)
		***REMOVED***

		private_key_bytes := crypto.FromECDSA(privateKey)
		private_key_string := hexutil.Encode(private_key_bytes)[2:]
		fmt.Println("private key: ", private_key_string)
		wallet_bind := GetAccountAuth(client.Client, private_key_string, int(client.Network.Id), client.Network.GasPrice, client.Network.GasLimit)
		all_wallets = append(all_wallets, wallet_bind)
		client.Config.PrivateKeys = append(client.Config.PrivateKeys, private_key_string)
	***REMOVED***
***REMOVED***
