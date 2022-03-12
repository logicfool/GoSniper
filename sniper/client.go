package sniper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logicfool/GoSniper/utils"
	// "github.com/logicfool/GoSniper/config"
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
	auth.Value = utils.EtherToWei(0.1) // in wei
	auth.GasLimit = uint64(10000000)   // in units
	auth.GasPrice = big.NewInt(150)
	return auth
***REMOVED***
