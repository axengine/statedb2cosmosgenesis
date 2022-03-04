package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Genesis struct {
	AppState AppState `json:"app_state"`
}

type AppState struct {
	Auth Auth `json:"auth"`
	Bank Bank `json:"bank"`
	Evm  Evm  `json:"evm"`
}

type Auth struct {
	Accounts    []Account `json:"accounts"`
	accountsMap map[string]bool
	Params      interface{} `json:"params"`
}

type Account struct {
	Type        string      `json:"@type"`
	BaseAccount BaseAccount `json:"base_account"`
	CodeHash    string      `json:"code_hash,omitempty"`
	Name        string      `json:"name,omitempty"`
	Permissions []string    `json:"permissions,omitempty"`
}

type BaseAccount struct {
	AccountNumber string      `json:"account_number"`
	Address       string      `json:"address"`
	PubKey        interface{} `json:"pub_key"`
	Sequence      string      `json:"sequence"`
}

type Bank struct {
	Balances       []Balance `json:"balances"`
	balancesMap    map[string]bool
	Denom_metadata interface{} `json:"denom_metadata"`
	Params         interface{} `json:"params"`
	Supply         interface{} `json:"supply"`
}

type Balance struct {
	Address string `json:"address"`
	Coins   []Coin `json:"coins"`
}

type Coin struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}

type Evm struct {
	Accounts []ETHAccount `json:"accounts"`
	Params   interface{}  `json:"params"`
}

type ETHAccount struct {
	Address string    `json:"address"`
	Code    string    `json:"code"`
	Storage []Storage `json:"storage"`
}

type Storage struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// DumpAccount for go-ethereum v1.10.3
type DumpAccount struct {
	Balance   string                 `json:"balance"`
	Nonce     uint64                 `json:"nonce"`
	Root      string                 `json:"root"`
	CodeHash  string                 `json:"codeHash"`
	Code      string                 `json:"code,omitempty"`
	Storage   map[common.Hash]string `json:"storage,omitempty"`
	Address   *common.Address        `json:"address,omitempty"` // Address only present in iterative (line-by-line) mode
	SecureKey hexutil.Bytes          `json:"key,omitempty"`     // If we don't have address, we can output the key

}

type Dump struct {
	Root     string                         `json:"root"`
	Accounts map[common.Address]DumpAccount `json:"accounts"`
}
