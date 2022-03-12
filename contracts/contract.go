package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetRouter(address common.Address, client *ethclient.Client) *Unirouter ***REMOVED***
	contract, err := NewUnirouter(address, client)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return contract
***REMOVED***

func GetFactory(address common.Address, client *ethclient.Client) *Unifactory ***REMOVED***
	contract, err := NewUnifactory(address, client)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return contract
***REMOVED***

func GetErc20(address common.Address, client *ethclient.Client) *Erc20 ***REMOVED***
	contract, err := NewErc20(address, client)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return contract
***REMOVED***

func GetPair(address common.Address, client *ethclient.Client) *Pair ***REMOVED***
	contract, err := NewPair(address, client)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return contract
***REMOVED***
