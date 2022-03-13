package config

import (
	"encoding/json"
	"os"
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

	PrivateKey struct ***REMOVED***
		Key string
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
