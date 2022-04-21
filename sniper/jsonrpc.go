package sniper

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/logicfool/GoSniper/utils"
)

var Quit chan (bool) = make(chan bool)

// Read pending transaction from mempool

func (client *MainClient) SubscribePendingTx() {
	utils.ColorPrint("Scanning Mempool txs activated!", "yellow")
	logs := make(chan *common.Hash)
	context := context.Background()
	sub, err := client.Rpc.EthSubscribe(context, logs, "newPendingTransactions")
	if err != nil {
		fmt.Println("Error During Sub1: ", err)
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Error During Sub2: ", err)
		case vLog := <-logs:
			// fmt.Printf("%T: %v \n", vLog, vLog) // pointer to event log
			// client.WaitGroup.Add(1)
			// fmt.Println(*vLog)
			go client.ProcessMempool(*vLog)
		}
	}

	//client.Client.PendingTransactionCount()Call(&txs, "eth_pendingTransactions"
}

func (client *MainClient) ToCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func (client *MainClient) EthCall(msg interface{}, args ...interface{}) ([]byte, error) {
	// var res interface{}
	// var args map[string]interface{}
	// args = make(map[string]interface{})
	// args["to"] = "0x0000000000000000000000000000000000000000"
	// cli := client.Rpc.Call(&res, "eth_call")
	// _ = cli
	// https://geth.ethereum.org/docs/rpc/ns-eth
	ctx := context.Background()
	var hex hexutil.Bytes
	// client.Rpc.CallContext()
	err := client.Rpc.CallContext(ctx, &hex, "eth_call", msg, args)
	if err != nil {
		return nil, err
	}
	return hex, nil

}

func (client *MainClient) GetETHPriceFromFeed() (*big.Int, error) {
	var hex1 hexutil.Bytes
	// fmt.Println("DataFeed: ", client.Network.DataFeed, client.Network.DataFeed == " ", client.Network.DataFeed == "")
	if client.Network.DataFeed == "" {
		// fmt.Println("DataFeed: ", client.Network.DataFeed, client.Network.DataFeed == " ", client.Network.DataFeed == "")
		return nil, errors.New("No Data Feed")
	}
	encodedfunc := "0xfeaf968c"
	arg := map[string]interface{}{
		"from": client.WalletAuths[0].From,
		"to":   common.HexToAddress(client.Network.DataFeed),
		"data": encodedfunc,
	}
	err := client.Rpc.Call(&hex1, "eth_call", arg, "latest")
	if err != nil {
		return nil, err
	}
	decodebyte := hex.EncodeToString(hex1[32:64])
	nn, success := new(big.Int).SetString(decodebyte, 16)
	if success == false {
		return nil, errors.New("Error in converting hex to big int")
	}
	nn = nn.Mul(nn, big.NewInt(10000000000))
	return nn, nil
}
