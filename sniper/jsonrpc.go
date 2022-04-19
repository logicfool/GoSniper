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

func (client *MainClient) SubscribePendingTx() ***REMOVED***
	utils.ColorPrint("Scanning Mempool txs activated!", "yellow")
	logs := make(chan *common.Hash)
	context := context.Background()
	sub, err := client.Rpc.EthSubscribe(context, logs, "newPendingTransactions")
	if err != nil ***REMOVED***
		fmt.Println("Error During Sub1: ", err)
	***REMOVED***
	for ***REMOVED***
		select ***REMOVED***
		case err := <-sub.Err():
			fmt.Println("Error During Sub2: ", err)
		case vLog := <-logs:
			// fmt.Printf("%T: %v \n", vLog, vLog) // pointer to event log
			// client.WaitGroup.Add(1)
			// fmt.Println(*vLog)
			go client.ProcessMempool(*vLog)
		***REMOVED***
	***REMOVED***

	//client.Client.PendingTransactionCount()Call(&txs, "eth_pendingTransactions"
***REMOVED***

func (client *MainClient) ToCallArg(msg ethereum.CallMsg) interface***REMOVED******REMOVED*** ***REMOVED***
	arg := map[string]interface***REMOVED******REMOVED******REMOVED***
		"from": msg.From,
		"to":   msg.To,
	***REMOVED***
	if len(msg.Data) > 0 ***REMOVED***
		arg["data"] = hexutil.Bytes(msg.Data)
	***REMOVED***
	if msg.Value != nil ***REMOVED***
		arg["value"] = (*hexutil.Big)(msg.Value)
	***REMOVED***
	if msg.Gas != 0 ***REMOVED***
		arg["gas"] = hexutil.Uint64(msg.Gas)
	***REMOVED***
	if msg.GasPrice != nil ***REMOVED***
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	***REMOVED***
	return arg
***REMOVED***

func (client *MainClient) EthCall(msg interface***REMOVED******REMOVED***, args ...interface***REMOVED******REMOVED***) ([]byte, error) ***REMOVED***
	// var res interface***REMOVED******REMOVED***
	// var args map[string]interface***REMOVED******REMOVED***
	// args = make(map[string]interface***REMOVED******REMOVED***)
	// args["to"] = "0x0000000000000000000000000000000000000000"
	// cli := client.Rpc.Call(&res, "eth_call")
	// _ = cli
	// https://geth.ethereum.org/docs/rpc/ns-eth
	ctx := context.Background()
	var hex hexutil.Bytes
	// client.Rpc.CallContext()
	err := client.Rpc.CallContext(ctx, &hex, "eth_call", msg, args)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return hex, nil

***REMOVED***

func (client *MainClient) GetETHPriceFromFeed() (*big.Int, error) ***REMOVED***
	var hex1 hexutil.Bytes
	// fmt.Println("DataFeed: ", client.Network.DataFeed, client.Network.DataFeed == " ", client.Network.DataFeed == "")
	if client.Network.DataFeed == "" ***REMOVED***
		// fmt.Println("DataFeed: ", client.Network.DataFeed, client.Network.DataFeed == " ", client.Network.DataFeed == "")
		return nil, errors.New("No Data Feed")
	***REMOVED***
	encodedfunc := "0xfeaf968c"
	arg := map[string]interface***REMOVED******REMOVED******REMOVED***
		"from": client.WalletAuths[0].From,
		"to":   common.HexToAddress(client.Network.DataFeed),
		"data": encodedfunc,
	***REMOVED***
	err := client.Rpc.Call(&hex1, "eth_call", arg, "latest")
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	decodebyte := hex.EncodeToString(hex1[32:64])
	nn, success := new(big.Int).SetString(decodebyte, 16)
	if success == false ***REMOVED***
		return nil, errors.New("Error in converting hex to big int")
	***REMOVED***
	nn = nn.Mul(nn, big.NewInt(10000000000))
	return nn, nil
***REMOVED***
