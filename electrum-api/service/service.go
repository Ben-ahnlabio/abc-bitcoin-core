package service

import types "github.com/ahnlabio/bitcoin-core/electrum-api/types"

type BtcService struct {
	Elemctrum types.IElectrum
}

func NewBitcoinApiService(electrum types.IElectrum) *BtcService {
	return &BtcService{
		Elemctrum: electrum,
	}
}

func (s BtcService) GetBalance(address string) (*types.GetBalanceResult, error) {
	_, err := validateAddress(address)
	if err != nil {
		return nil, InvalidAddressError(err)
	}

	result, err := s.Elemctrum.GetBalance(address)
	if err != nil {
		return nil, ElectrumError(err)
	}

	return &types.GetBalanceResult{
		Address:     address,
		Confirmed:   int(result.Confirmed),
		Unconfirmed: int(result.Unconfirmed),
	}, nil
}

func (s BtcService) GetTransaction(txId string) (*types.GetTransactionResult, error) {
	result, err := s.Elemctrum.GetTransaction(txId)
	if err != nil {
		return nil, ElectrumError(err)
	}

	return &types.GetTransactionResult{
		BlockHash:     result.Blockhash,
		TxHash:        result.Hash,
		Confirmations: int(result.Confirmations),
	}, nil
}

func (s BtcService) GetUTXO(address string) (*types.GetUTXOResult, error) {
	result, err := s.Elemctrum.GetListUnspent(address)
	if err != nil {
		return nil, ElectrumError(err)
	}

	utxos := make([]*types.UTXO, 0)
	for _, utxo := range result {
		utxos = append(utxos, &types.UTXO{
			Height:   utxo.Height,
			Position: utxo.Position,
			Hash:     utxo.Hash,
			Value:    utxo.Value,
		})
	}

	return &types.GetUTXOResult{
		Address: address,
		UTXOs:   utxos,
	}, nil

}

func (s BtcService) GetHistory(address string) (*types.GetHistoryResult, error) {
	result, err := s.Elemctrum.GetHistory(address)
	if err != nil {
		return nil, ElectrumError(err)
	}

	histories := make([]*types.History, 0)
	for _, history := range result {
		histories = append(histories, &types.History{
			Height: history.Height,
			TxHash: history.Hash,
		})
	}

	return &types.GetHistoryResult{
		Address:   address,
		Histories: histories,
	}, nil
}
