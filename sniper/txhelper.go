package sniper

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/logicfool/GoSniper/contracts"
)

func (client *MainClient) GetSignedRawTx(Wallet *bind.TransactOpts, function_name string, params ...interface***REMOVED******REMOVED***) *types.Transaction ***REMOVED***
	var data []byte
	var err error
	abi, _ := contracts.GoSniperMetaData.GetAbi()
	data, err = abi.Pack(function_name, params...)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	toaddr := client.MainExchange.SniperAddress
	nonce := Noncehandler[Wallet.From]
	baseTx := &types.LegacyTx***REMOVED***
		To:       &toaddr,
		Nonce:    uint64(nonce),
		GasPrice: client.Session.GasPrice,
		Gas:      client.Session.GasLimit,
		Value:    Wallet.Value,
		Data:     data,
	***REMOVED***
	txs := types.NewTx(baseTx)
	signedTx, err := Wallet.Signer(Wallet.From, txs)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return signedTx
	// err = client.Client.SendTransaction(context.Background(), signedTx)
	// if err != nil ***REMOVED***
	// 	panic(err)
	// ***REMOVED***
***REMOVED***

func (client *MainClient) SendRawTx(signedTx *types.Transaction) (common.Hash, error) ***REMOVED***
	err := client.Client.SendTransaction(context.Background(), signedTx)
	if err != nil ***REMOVED***
		return signedTx.Hash(), err
	***REMOVED***
	// signedTx
	return signedTx.Hash(), nil

***REMOVED***

// ----------------------------------------------------------------------------------------------------------------
func (client *MainClient) FormAndSaverawTxs() ***REMOVED***

***REMOVED***
