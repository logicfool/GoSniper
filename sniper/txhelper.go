package sniper

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/logicfool/GoSniper/contracts"
)

func (client *MainClient) GetSignedRawTx(Wallet *bind.TransactOpts, function_name string, params ...interface{}) *types.Transaction {
	var data []byte
	var err error
	abi, _ := contracts.GoSniperMetaData.GetAbi()
	data, err = abi.Pack(function_name, params...)
	if err != nil {
		panic(err)
	}
	toaddr := client.MainExchange.SniperAddress
	nonce := Noncehandler[Wallet.From]
	baseTx := &types.LegacyTx{
		To:       &toaddr,
		Nonce:    uint64(nonce),
		GasPrice: client.Session.GasPrice,
		Gas:      client.Session.GasLimit,
		Value:    Wallet.Value,
		Data:     data,
	}
	txs := types.NewTx(baseTx)
	signedTx, err := Wallet.Signer(Wallet.From, txs)
	if err != nil {
		panic(err)
	}
	return signedTx
	// err = client.Client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	panic(err)
	// }
}

func (client *MainClient) SendRawTx(signedTx *types.Transaction) (common.Hash, error) {
	err := client.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return signedTx.Hash(), err
	}
	// signedTx
	return signedTx.Hash(), nil

}

// ----------------------------------------------------------------------------------------------------------------
func (client *MainClient) FormAndSaverawTxs() {

}
