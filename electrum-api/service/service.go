package service

import (
	"github.com/ahnlabio/bitcoin-core/electrum-api/config"
	"github.com/ahnlabio/bitcoin-core/electrum-api/electrum"
)

func GetBalance(address string) (*GetBalanceResult, error) {
	_, err := validateAddress(address)
	if err != nil {
		return nil, InvalidAddressError(err)
	}

	appConfig := config.GetConfig()
	electrum := electrum.NewElectrum(appConfig.ElectrumHost, appConfig.ElectrumPort)
	result, err := electrum.GetBalance(address)
	if err != nil {
		return nil, ElectrumError(err)
	}

	return &GetBalanceResult{
		Address:     address,
		Confirmed:   int(result.Confirmed),
		Unconfirmed: int(result.Unconfirmed),
	}, nil
}

func GetTransaction(txId string) (*GetTransactionResult, error) {
	appConfig := config.GetConfig()
	electrum := electrum.NewElectrum(appConfig.ElectrumHost, appConfig.ElectrumPort)
	result, err := electrum.GetTransaction(txId)
	if err != nil {
		return nil, ElectrumError(err)
	}

	return &GetTransactionResult{
		BlockHash:     result.Blockhash,
		TxHash:        result.Hash,
		Confirmations: int(result.Confirmations),
	}, nil
}

func GetUTXO(address string) (*GetUTXOResult, error) {
	appConfig := config.GetConfig()
	electrum := electrum.NewElectrum(appConfig.ElectrumHost, appConfig.ElectrumPort)
	result, err := electrum.GetListUnspent(address)
	if err != nil {
		return nil, ElectrumError(err)
	}

	utxos := make([]*UTXO, 0)
	for _, utxo := range result {
		utxos = append(utxos, &UTXO{
			Height:   utxo.Height,
			Position: utxo.Position,
			Hash:     utxo.Hash,
			Value:    utxo.Value,
		})
	}

	return &GetUTXOResult{
		Address: address,
		UTXOs:   utxos,
	}, nil

}

type GetUTXOResult struct {
	Address string  `json:"address"`
	UTXOs   []*UTXO `json:"utxos"`
}

type UTXO struct {
	Height   uint32 `json:"height"`
	Position uint32 `json:"tx_pos"`
	Hash     string `json:"tx_hash"`
	Value    uint64 `json:"value"`
}

type GetTransactionResult struct {
	BlockHash     string `json:"block_hash"`
	TxHash        string `json:"tx_hash"`
	Confirmations int    `json:"confirmations"`
}

type GetBalanceResult struct {
	Address     string `json:"address"`
	Confirmed   int    `json:"confirmd"`
	Unconfirmed int    `json:"unconfirmd"`
}
