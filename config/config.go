package config

import (
	"encoding/json"
	"os"
)

type (
	Address string

	Config struct ***REMOVED***
		Chains []Network `json:"networks"`
	***REMOVED***

	Network struct ***REMOVED***
		RPC  string `json:"node_url"`
		Id   int    `json:"chain_id"`
		Name string `json:"name"`
	***REMOVED***
	Contracts struct ***REMOVED***
	***REMOVED***

	Exchanges struct ***REMOVED***
		ChainId int
	***REMOVED***
)

func Read_config_from_file(f string) (*Config, error) ***REMOVED***
	file, err := os.Open(f)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config***REMOVED******REMOVED***
	err = decoder.Decode(&config)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return &config, nil
***REMOVED***
