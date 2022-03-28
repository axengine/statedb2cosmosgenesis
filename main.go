package main

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"strconv"
)

func main() {
	globalInitCosmosConfig("evmos")

	if len(os.Args) != 3 {
		panic("usage ./cmd evmos_genesis.json mondo_genesis.json")
	}

	evmos, err := loadEvmosGenesis(os.Args[1])
	if err != nil {
		panic(err)
	}

	mondo, err := loadMondoGensis(os.Args[2])
	if err != nil {
		panic(err)
	}

	for k, v := range mondo.Accounts {
		address := evmAddressToCosmosAddress(k.Hex())
		if !evmos.AppState.Auth.accountsMap[address] {
			if needAddToAuth(v) {
				evmos.AppState.Auth.Accounts = append(evmos.AppState.Auth.Accounts, Account{
					Type: "/ethermint.types.v1.EthAccount",
					BaseAccount: BaseAccount{
						AccountNumber: "0",
						Address:       address,
						PubKey:        nil,
						Sequence:      strconv.FormatInt(int64(v.Nonce), 10),
					},
					CodeHash:    common.HexToHash(v.CodeHash).Hex(),
					Name:        "",
					Permissions: nil,
				})
				evmos.AppState.Auth.accountsMap[address] = true
			}
		}

		if !evmos.AppState.Bank.balancesMap[address] {
			if needAddToBank(v) {
				evmos.AppState.Bank.Balances = append(evmos.AppState.Bank.Balances, Balance{
					Address: address,
					Coins: []Coin{{
						Amount: v.Balance,
						Denom:  "aevmos",
					}},
				})
				evmos.AppState.Bank.balancesMap[address] = true
			}
		}

		if needAddToEvm(v) {
			ethAct := ETHAccount{
				Address: k.Hex(),
				Code:    v.Code,
				Storage: nil,
			}
			for k, v := range v.Storage {
				ethAct.Storage = append(ethAct.Storage, Storage{
					Key:   k.Hex(),
					Value: common.HexToHash(v).Hex(),
				})
			}

			evmos.AppState.Evm.Accounts = append(evmos.AppState.Evm.Accounts, ethAct)
		}
	}

	genesisMap, err := loadEvmosGenesisToMap(os.Args[1])
	if err != nil {
		panic(err)
	}

	appState := genesisMap["app_state"].(map[string]interface{})
	appState["auth"] = evmos.AppState.Auth
	appState["bank"] = evmos.AppState.Bank
	appState["evm"] = evmos.AppState.Evm

	genesisMap["app_state"] = appState

	bz, _ := json.MarshalIndent(genesisMap, "", "  ")
	fmt.Println(string(bz))
}

func needAddToAuth(act DumpAccount) bool {
	return act.Nonce > 0
}

func needAddToBank(act DumpAccount) bool {
	return act.Balance != "0"
}

func needAddToEvm(act DumpAccount) bool {
	return len(act.Code) > 0
}

func loadEvmosGenesis(fname string) (*Genesis, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var genesisData Genesis

	if err := json.NewDecoder(f).Decode(&genesisData); err != nil {
		return nil, err
	}
	genesisData.AppState.Auth.accountsMap = make(map[string]bool)
	genesisData.AppState.Bank.balancesMap = make(map[string]bool)
	for _, v := range genesisData.AppState.Auth.Accounts {
		genesisData.AppState.Auth.accountsMap[v.BaseAccount.Address] = true
		genesisData.AppState.Bank.balancesMap[v.BaseAccount.Address] = true
	}

	return &genesisData, nil
}

func loadEvmosGenesisToMap(fname string) (map[string]interface{}, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var genesisData = make(map[string]interface{})

	if err := json.NewDecoder(f).Decode(&genesisData); err != nil {
		return nil, err
	}

	return genesisData, nil
}

func loadMondoGensis(fname string) (*Dump, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var dump Dump
	if err = json.NewDecoder(f).Decode(&dump); err != nil {
		return nil, err
	}
	return &dump, nil
}

func globalInitCosmosConfig(addressPrefix string) {
	sdk.GetConfig().SetBech32PrefixForAccount(addressPrefix, "")
	sdk.GetConfig().Seal()
}

func evmAddressToCosmosAddress(address string) string {
	return sdk.AccAddress(common.HexToAddress(address).Bytes()).String()
}
