package structs

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/logicfool/GoSniper/contracts"
	"github.com/logicfool/GoSniper/utils"
	// "github.com/logicfool/GoSniper/utils"
)

// Majorly Taken from https://github.com/umbracle/ethgo/blob/master/structs.go

// Config Structs

type (
	// Address string

	BaseConfig struct ***REMOVED***
		Chains      []Network      `json:"networks"`
		Exchanges   []Exchanges    `json:"exchanges"`
		PrivateKeys []string       `json:"private_keys"`
		Sniper      SniperSettings `json:"sniper"`
	***REMOVED***

	Network struct ***REMOVED***
		RPC               string `json:"node_url"`
		Id                int    `json:"chain_id"`
		Name              string `json:"name"`
		BuyContract       string `json:"contract"`
		DataFeed          string `json:"datafeed,omitempty"`
		Explorer          string `json:"explorer"`
		GasLimit          int    `json:"gaslimit,omitempty"`
		GasPrice          int    `json:"gasprice,omitempty"`
		EIP1599Compatible bool   `json:"eip1599_compatible,omitempty"`
	***REMOVED***
	Contracts struct ***REMOVED***
	***REMOVED***

	Exchanges struct ***REMOVED***
		ChainId   int    `json:"network"`
		Name      string `json:"name"`
		Router    string `json:"router"`
		Factory   string `json:"factory"`
		HoneyPot  string `json:"honeypot"`
		Init_code string `json:"init_code"`
	***REMOVED***

	PrivateKey struct ***REMOVED***
		Key string
	***REMOVED***

	SniperSettings struct ***REMOVED***
		Slippage                    int            `json:"slippage,omitempty"`
		MultiWallet                 bool           `json:"multiwallet,omitempty"`
		MinLiq                      float64        `json:"min_liquidity,omitempty"`
		BuySplit                    int            `json:"buy_split,omitempty"`
		SellSplit                   int            `json:"sell_split,omitempty"`
		HoneyPotCheck               bool           `json:"honeypot_check,omitempty"`
		LiquidityAddFunctions       []string       `json:"liquidity_add_functions,omitempty"`
		LiquidityRemoveFunctions    []string       `json:"liquidity_remove_functions,omitempty"`
		ChiGasEnabled               bool           `json:"chigas_enabled,omitempty"`
		AdvancedLiquiditySniping    bool           `json:"advanced_liquidity_sniping,omitempty"`
		MutliMainSnipe              bool           `json:"mutli_main_snipe,omitempty"`
		Amount                      float64        `json:",omitempty"`
		Token                       common.Address `json:",omitempty"`
		Pair                        common.Address `json:",omitempty"`
		Watch                       bool           `json:",omitempty"`
		MaxBuyTax                   float64        `json:",omitempty"`
		MaxSellTax                  float64        `json:",omitempty"`
		MempoolLiquidityCheck       bool           `json:",omitempty"`
		MempoolRemoveLiquidityCheck bool           `json:",omitempty"`
		DisableTradingCheck         bool           `json:",omitempty"`
		Network                     string         `json:",omitempty"`
		Exchange                    string         `json:",omitempty"`
		ToBuy                       bool           `json:",omitempty"`
		ToSell                      bool           `json:",omitempty"`
		BlocksToSkip                int            `json:",omitempty"`
		TakeProfitPercentage        float64        `json:",omitempty"`
		StopLossPercentage          float64        `json:",omitempty"`
		BuyPrice                    float64        `json:",omitempty"`
		SellPrice                   float64        `json:",omitempty"`
		DisableReserveChecks        bool           `json:",omitempty"`
		Freefire                    bool           `json:",omitempty"`
		GasPrice                    int            `json:",omitempty"`
		GasLimit                    int            `json:",omitempty"`
		FreeFireLimit               int            `json:",omitempty"`
		NumberOfMultipleWallets     int            `json:",omitempty"`
	***REMOVED***

	// Exchanges is for baseconfig to read the data and then its converted to EXChange when setting up the main config...
	Exchange struct ***REMOVED***
		Name           string
		Router         *contracts.Unirouter
		Factory        *contracts.Unifactory
		Sniper         *contracts.GoSniper
		RouterAddress  common.Address
		FactoryAddress common.Address
		SniperAddress  common.Address
		Init_code      string
		WETH           *contracts.WETH
		WETHAddress    common.Address
	***REMOVED***

	CurrentSnipeSession struct ***REMOVED***
		Token                  common.Address
		Pair                   common.Address
		TokenContract          *contracts.Token
		PairContract           *contracts.Token
		AllTokens              []*contracts.Token
		AllPairs               map[common.Address]map[common.Address]common.Address
		AllBalanceOfUserTokens map[int]map[common.Address]map[common.Address]*big.Int
		AmountToTrade          *big.Int
		AmountBought           *big.Int
		BnbSpent               *big.Int //make price checking..
		Reserve0               *big.Int
		Reserve1               *big.Int
		BuySuccess             bool
		SellSuccess            bool
		SuccessAddress         common.Address
		Buy                    bool
		Sell                   bool
		Split                  int
		LiquidityAdded         bool
		LiquidityRemoved       bool
		Stop                   bool
		GasPrice               *big.Int
		GasLimit               uint64
		PriceBought            *big.Float
		RouterABI              abi.ABI //to not create one every loop!
		Wallets_to_use         []*bind.TransactOpts
		MajorAddresses         []common.Address
		ExtraSettings          []*big.Int
		Path                   []common.Address
		// OneSuccessful          bool
	***REMOVED***

	LiquidityAddFunctions    map[string]string
	LiquidityRemoveFunctions map[string]string
)

var (
	// ZeroAddress is an address of all zeros
	ZeroAddress = Address***REMOVED******REMOVED***

	// ZeroHash is a hash of all zeros
	ZeroHash = Hash***REMOVED******REMOVED***
)

// Address is an Ethereum address
type Address [20]byte

// BytesToAddress converts bytes to an address object
func BytesToAddress(b []byte) Address ***REMOVED***
	var a Address

	size := len(b)
	min := min(size, 20)

	copy(a[20-min:], b[len(b)-min:])
	return a
***REMOVED***

// Address implements the ethgo.Key interface Address method.
func (a Address) Address() Address ***REMOVED***
	return a
***REMOVED***

// Sign implements the ethgo.Key interface Sign method.
func (a Address) Sign(hash []byte) ([]byte, error) ***REMOVED***
	panic("an address cannot sign messages")
***REMOVED***

// MarshalText implements the marshal interface
func (a Address) MarshalText() ([]byte, error) ***REMOVED***
	return []byte(a.String()), nil
***REMOVED***

// Bytes returns the bytes of the Address
func (a Address) Bytes() []byte ***REMOVED***
	return a[:]
***REMOVED***

func (a Address) String() string ***REMOVED***
	return a.checksumEncode()
***REMOVED***

func (a Address) checksumEncode() string ***REMOVED***
	address := strings.ToLower(hex.EncodeToString(a[:]))
	hash := hex.EncodeToString(utils.Keccak256([]byte(address)))

	ret := "0x"
	for i := 0; i < len(address); i++ ***REMOVED***
		character := string(address[i])

		num, _ := strconv.ParseInt(string(hash[i]), 16, 64)
		if num > 7 ***REMOVED***
			ret += strings.ToUpper(character)
		***REMOVED*** else ***REMOVED***
			ret += character
		***REMOVED***
	***REMOVED***

	return ret
***REMOVED***

// Hash is an Ethereum hash
type Hash [32]byte

// HexToHash converts an hex string value to a hash object

// BytesToHash converts bytes to a hash object
func BytesToHash(b []byte) Hash ***REMOVED***
	var h Hash

	size := len(b)
	min := min(size, 32)

	copy(h[32-min:], b[len(b)-min:])
	return h
***REMOVED***

// MarshalText implements the marshal interface
func (h Hash) MarshalText() ([]byte, error) ***REMOVED***
	return []byte(h.String()), nil
***REMOVED***

// Bytes returns the bytes of the Hash
func (h Hash) Bytes() []byte ***REMOVED***
	return h[:]
***REMOVED***

func (h Hash) String() string ***REMOVED***
	return "0x" + hex.EncodeToString(h[:])
***REMOVED***

func (h Hash) Location() string ***REMOVED***
	return h.String()
***REMOVED***

type Block struct ***REMOVED***
	Number             uint64
	Hash               Hash
	ParentHash         Hash
	Sha3Uncles         Hash
	TransactionsRoot   Hash
	StateRoot          Hash
	ReceiptsRoot       Hash
	Miner              Address
	Difficulty         *big.Int
	ExtraData          []byte
	GasLimit           uint64
	GasUsed            uint64
	Timestamp          uint64
	Transactions       []*Transaction
	TransactionsHashes []Hash
	Uncles             []Hash
***REMOVED***

func (b *Block) Copy() *Block ***REMOVED***
	bb := new(Block)
	*bb = *b
	if b.Difficulty != nil ***REMOVED***
		bb.Difficulty = new(big.Int).Set(b.Difficulty)
	***REMOVED***
	bb.ExtraData = append(bb.ExtraData[:0], b.ExtraData...)
	bb.Transactions = make([]*Transaction, len(b.Transactions))
	for indx, txn := range b.Transactions ***REMOVED***
		bb.Transactions[indx] = txn.Copy()
	***REMOVED***
	return bb
***REMOVED***

type TransactionType int

const (
	TransactionLegacy TransactionType = 0
	// eip-2930
	TransactionAccessList TransactionType = 1
	// eip-1559
	TransactionDynamicFee TransactionType = 2
)

type Transaction struct ***REMOVED***
	Type TransactionType

	// legacy values
	Hash     Hash
	From     Address
	To       *Address
	Input    []byte
	GasPrice uint64
	Gas      uint64
	Value    *big.Int
	Nonce    uint64
	V        []byte
	R        []byte
	S        []byte

	// jsonrpc values
	BlockHash   Hash
	BlockNumber uint64
	TxnIndex    uint64

	// eip-2930 values
	ChainID    *big.Int
	AccessList AccessList

	// eip-1559 values
	MaxPriorityFeePerGas *big.Int
	MaxFeePerGas         *big.Int
***REMOVED***

func (t *Transaction) Copy() *Transaction ***REMOVED***
	tt := new(Transaction)
	if t.To != nil ***REMOVED***
		to := Address(*t.To)
		tt.To = &to
	***REMOVED***
	tt.Input = append(tt.Input[:0], t.Input...)
	if t.Value != nil ***REMOVED***
		tt.Value = new(big.Int).Set(t.Value)
	***REMOVED***
	tt.V = append(tt.V[:0], t.V...)
	tt.R = append(tt.R[:0], t.R...)
	tt.S = append(tt.S[:0], t.S...)
	if t.ChainID != nil ***REMOVED***
		tt.ChainID = new(big.Int).Set(t.ChainID)
	***REMOVED***
	if t.MaxPriorityFeePerGas != nil ***REMOVED***
		tt.MaxPriorityFeePerGas = new(big.Int).Set(t.MaxPriorityFeePerGas)
	***REMOVED***
	if t.MaxFeePerGas != nil ***REMOVED***
		tt.MaxFeePerGas = new(big.Int).Set(t.MaxFeePerGas)
	***REMOVED***
	return tt
***REMOVED***

type AccessEntry struct ***REMOVED***
	Address Address
	Storage []Hash
***REMOVED***

type AccessList []AccessEntry

type CallMsg struct ***REMOVED***
	From     Address
	To       *Address
	Data     []byte
	GasPrice uint64
	Gas      *big.Int
	Value    *big.Int
***REMOVED***

type LogFilter struct ***REMOVED***
	Address   []Address
	Topics    [][]*Hash
	BlockHash *Hash
	From      *BlockNumber
	To        *BlockNumber
***REMOVED***

func (l *LogFilter) SetFromUint64(num uint64) ***REMOVED***
	b := BlockNumber(num)
	l.From = &b
***REMOVED***

func (l *LogFilter) SetToUint64(num uint64) ***REMOVED***
	b := BlockNumber(num)
	l.To = &b
***REMOVED***

func (l *LogFilter) SetTo(b BlockNumber) ***REMOVED***
	l.To = &b
***REMOVED***

type Receipt struct ***REMOVED***
	TransactionHash   Hash
	TransactionIndex  uint64
	ContractAddress   Address
	BlockHash         Hash
	From              Address
	BlockNumber       uint64
	GasUsed           uint64
	CumulativeGasUsed uint64
	LogsBloom         []byte
	Logs              []*Log
	Status            uint64
***REMOVED***

func (r *Receipt) Copy() *Receipt ***REMOVED***
	rr := new(Receipt)
	*rr = *r
	rr.LogsBloom = append(rr.LogsBloom[:0], r.LogsBloom...)
	rr.Logs = make([]*Log, len(r.Logs))
	for indx, log := range r.Logs ***REMOVED***
		rr.Logs[indx] = log.Copy()
	***REMOVED***
	return rr
***REMOVED***

type Log struct ***REMOVED***
	Removed          bool
	LogIndex         uint64
	TransactionIndex uint64
	TransactionHash  Hash
	BlockHash        Hash
	BlockNumber      uint64
	Address          Address
	Topics           []Hash
	Data             []byte
***REMOVED***

func (l *Log) Copy() *Log ***REMOVED***
	ll := new(Log)
	*ll = *l
	ll.Data = append(ll.Data[:0], l.Data...)
	return ll
***REMOVED***

type BlockNumber int

const (
	Latest   BlockNumber = -1
	Earliest BlockNumber = -2
	Pending  BlockNumber = -3
)

func (b BlockNumber) Location() string ***REMOVED***
	return b.String()
***REMOVED***

func (b BlockNumber) String() string ***REMOVED***
	switch b ***REMOVED***
	case Latest:
		return "latest"
	case Earliest:
		return "earliest"
	case Pending:
		return "pending"
	***REMOVED***
	if b < 0 ***REMOVED***
		panic("internal. blocknumber is negative")
	***REMOVED***
	return fmt.Sprintf("0x%x", uint64(b))
***REMOVED***

func EncodeBlock(block ...BlockNumber) BlockNumber ***REMOVED***
	if len(block) != 1 ***REMOVED***
		return Latest
	***REMOVED***
	return block[0]
***REMOVED***

type BlockNumberOrHash interface ***REMOVED***
	Location() string
***REMOVED***

func min(i, j int) int ***REMOVED***
	if i < j ***REMOVED***
		return i
	***REMOVED***
	return j
***REMOVED***

type Key interface ***REMOVED***
	Address() Address
	Sign(hash []byte) ([]byte, error)
***REMOVED***

func completeHex(str string, num int) []byte ***REMOVED***
	num = num * 2
	str = strings.TrimPrefix(str, "0x")

	size := len(str)
	if size < num ***REMOVED***
		for i := size; i < num; i++ ***REMOVED***
			str = "0" + str
		***REMOVED***
	***REMOVED*** else ***REMOVED***
		diff := size - num
		str = str[diff:]
	***REMOVED***
	return []byte("0x" + str)
***REMOVED***
