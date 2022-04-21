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

	BaseConfig struct {
		Chains      []Network      `json:"networks"`
		Exchanges   []Exchanges    `json:"exchanges"`
		PrivateKeys []string       `json:"private_keys"`
		Sniper      SniperSettings `json:"sniper"`
	}

	Network struct {
		RPC               string `json:"node_url"`
		Id                int    `json:"chain_id"`
		Name              string `json:"name"`
		BuyContract       string `json:"contract"`
		DataFeed          string `json:"datafeed,omitempty"`
		Explorer          string `json:"explorer"`
		GasLimit          int    `json:"gaslimit,omitempty"`
		GasPrice          int    `json:"gasprice,omitempty"`
		EIP1599Compatible bool   `json:"eip1599_compatible,omitempty"`
	}
	Contracts struct {
	}

	Exchanges struct {
		ChainId   int    `json:"network"`
		Name      string `json:"name"`
		Router    string `json:"router"`
		Factory   string `json:"factory"`
		HoneyPot  string `json:"honeypot"`
		Init_code string `json:"init_code"`
	}

	PrivateKey struct {
		Key string
	}

	SniperSettings struct {
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
	}

	// Exchanges is for baseconfig to read the data and then its converted to EXChange when setting up the main config...
	Exchange struct {
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
	}

	CurrentSnipeSession struct {
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
	}

	LiquidityAddFunctions    map[string]string
	LiquidityRemoveFunctions map[string]string
)

var (
	// ZeroAddress is an address of all zeros
	ZeroAddress = Address{}

	// ZeroHash is a hash of all zeros
	ZeroHash = Hash{}
)

// Address is an Ethereum address
type Address [20]byte

// BytesToAddress converts bytes to an address object
func BytesToAddress(b []byte) Address {
	var a Address

	size := len(b)
	min := min(size, 20)

	copy(a[20-min:], b[len(b)-min:])
	return a
}

// Address implements the ethgo.Key interface Address method.
func (a Address) Address() Address {
	return a
}

// Sign implements the ethgo.Key interface Sign method.
func (a Address) Sign(hash []byte) ([]byte, error) {
	panic("an address cannot sign messages")
}

// MarshalText implements the marshal interface
func (a Address) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

// Bytes returns the bytes of the Address
func (a Address) Bytes() []byte {
	return a[:]
}

func (a Address) String() string {
	return a.checksumEncode()
}

func (a Address) checksumEncode() string {
	address := strings.ToLower(hex.EncodeToString(a[:]))
	hash := hex.EncodeToString(utils.Keccak256([]byte(address)))

	ret := "0x"
	for i := 0; i < len(address); i++ {
		character := string(address[i])

		num, _ := strconv.ParseInt(string(hash[i]), 16, 64)
		if num > 7 {
			ret += strings.ToUpper(character)
		} else {
			ret += character
		}
	}

	return ret
}

// Hash is an Ethereum hash
type Hash [32]byte

// HexToHash converts an hex string value to a hash object

// BytesToHash converts bytes to a hash object
func BytesToHash(b []byte) Hash {
	var h Hash

	size := len(b)
	min := min(size, 32)

	copy(h[32-min:], b[len(b)-min:])
	return h
}

// MarshalText implements the marshal interface
func (h Hash) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

// Bytes returns the bytes of the Hash
func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

func (h Hash) Location() string {
	return h.String()
}

type Block struct {
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
}

func (b *Block) Copy() *Block {
	bb := new(Block)
	*bb = *b
	if b.Difficulty != nil {
		bb.Difficulty = new(big.Int).Set(b.Difficulty)
	}
	bb.ExtraData = append(bb.ExtraData[:0], b.ExtraData...)
	bb.Transactions = make([]*Transaction, len(b.Transactions))
	for indx, txn := range b.Transactions {
		bb.Transactions[indx] = txn.Copy()
	}
	return bb
}

type TransactionType int

const (
	TransactionLegacy TransactionType = 0
	// eip-2930
	TransactionAccessList TransactionType = 1
	// eip-1559
	TransactionDynamicFee TransactionType = 2
)

type Transaction struct {
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
}

func (t *Transaction) Copy() *Transaction {
	tt := new(Transaction)
	if t.To != nil {
		to := Address(*t.To)
		tt.To = &to
	}
	tt.Input = append(tt.Input[:0], t.Input...)
	if t.Value != nil {
		tt.Value = new(big.Int).Set(t.Value)
	}
	tt.V = append(tt.V[:0], t.V...)
	tt.R = append(tt.R[:0], t.R...)
	tt.S = append(tt.S[:0], t.S...)
	if t.ChainID != nil {
		tt.ChainID = new(big.Int).Set(t.ChainID)
	}
	if t.MaxPriorityFeePerGas != nil {
		tt.MaxPriorityFeePerGas = new(big.Int).Set(t.MaxPriorityFeePerGas)
	}
	if t.MaxFeePerGas != nil {
		tt.MaxFeePerGas = new(big.Int).Set(t.MaxFeePerGas)
	}
	return tt
}

type AccessEntry struct {
	Address Address
	Storage []Hash
}

type AccessList []AccessEntry

type CallMsg struct {
	From     Address
	To       *Address
	Data     []byte
	GasPrice uint64
	Gas      *big.Int
	Value    *big.Int
}

type LogFilter struct {
	Address   []Address
	Topics    [][]*Hash
	BlockHash *Hash
	From      *BlockNumber
	To        *BlockNumber
}

func (l *LogFilter) SetFromUint64(num uint64) {
	b := BlockNumber(num)
	l.From = &b
}

func (l *LogFilter) SetToUint64(num uint64) {
	b := BlockNumber(num)
	l.To = &b
}

func (l *LogFilter) SetTo(b BlockNumber) {
	l.To = &b
}

type Receipt struct {
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
}

func (r *Receipt) Copy() *Receipt {
	rr := new(Receipt)
	*rr = *r
	rr.LogsBloom = append(rr.LogsBloom[:0], r.LogsBloom...)
	rr.Logs = make([]*Log, len(r.Logs))
	for indx, log := range r.Logs {
		rr.Logs[indx] = log.Copy()
	}
	return rr
}

type Log struct {
	Removed          bool
	LogIndex         uint64
	TransactionIndex uint64
	TransactionHash  Hash
	BlockHash        Hash
	BlockNumber      uint64
	Address          Address
	Topics           []Hash
	Data             []byte
}

func (l *Log) Copy() *Log {
	ll := new(Log)
	*ll = *l
	ll.Data = append(ll.Data[:0], l.Data...)
	return ll
}

type BlockNumber int

const (
	Latest   BlockNumber = -1
	Earliest BlockNumber = -2
	Pending  BlockNumber = -3
)

func (b BlockNumber) Location() string {
	return b.String()
}

func (b BlockNumber) String() string {
	switch b {
	case Latest:
		return "latest"
	case Earliest:
		return "earliest"
	case Pending:
		return "pending"
	}
	if b < 0 {
		panic("internal. blocknumber is negative")
	}
	return fmt.Sprintf("0x%x", uint64(b))
}

func EncodeBlock(block ...BlockNumber) BlockNumber {
	if len(block) != 1 {
		return Latest
	}
	return block[0]
}

type BlockNumberOrHash interface {
	Location() string
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

type Key interface {
	Address() Address
	Sign(hash []byte) ([]byte, error)
}

func completeHex(str string, num int) []byte {
	num = num * 2
	str = strings.TrimPrefix(str, "0x")

	size := len(str)
	if size < num {
		for i := size; i < num; i++ {
			str = "0" + str
		}
	} else {
		diff := size - num
		str = str[diff:]
	}
	return []byte("0x" + str)
}
