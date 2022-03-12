package sniper

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logicfool/GoSniper/config"
)

type (
	CurrentState struct ***REMOVED***
	***REMOVED***
)

func GetClient(network config.Network) ethclient.Client ***REMOVED***
	client, err := ethclient.Dial(network.RPC)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return *client
***REMOVED***
