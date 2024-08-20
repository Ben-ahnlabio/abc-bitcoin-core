package service

import interfaces "github.com/ahnlabio/bitcoin-core/electrum-api/interfaces"

type BitcoinApiService struct {
	Elemctrum interfaces.Electrum
}

func NewBitcoinApiService(electrum interfaces.Electrum) *BitcoinApiService {
	return &BitcoinApiService{
		Elemctrum: electrum,
	}
}

func (s BitcoinApiService) GetBalance(address string) (*GetBalanceResult, error) {
	_, err := validateAddress(address)
	if err != nil {
		return nil, InvalidAddressError(err)
	}

	result, err := s.Elemctrum.GetBalance(address)
	if err != nil {
		return nil, ElectrumError(err)
	}

	return &GetBalanceResult{
		Address:     address,
		Confirmed:   int(result.Confirmed),
		Unconfirmed: int(result.Unconfirmed),
	}, nil
}

func (s BitcoinApiService) GetTransaction(txId string) (*GetTransactionResult, error) {
	result, err := s.Elemctrum.GetTransaction(txId)
	if err != nil {
		return nil, ElectrumError(err)
	}

	return &GetTransactionResult{
		BlockHash:     result.Blockhash,
		TxHash:        result.Hash,
		Confirmations: int(result.Confirmations),
	}, nil
}

func (s BitcoinApiService) GetUTXO(address string) (*GetUTXOResult, error) {
	result, err := s.Elemctrum.GetListUnspent(address)
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

func (s BitcoinApiService) GetHistory(address string) (*GetHistoryResult, error) {
	result, err := s.Elemctrum.GetHistory(address)
	if err != nil {
		return nil, ElectrumError(err)
	}

	histories := make([]*History, 0)
	for _, history := range result {
		histories = append(histories, &History{
			Height: history.Height,
			TxHash: history.Hash,
		})
	}

	return &GetHistoryResult{
		Address:   address,
		Histories: histories,
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

type GetHistoryResult struct {
	Address   string     `json:"address"`
	Histories []*History `json:"histories"`
}

type History struct {
	Height int32  `json:"height"`
	TxHash string `json:"tx_hash"`
	Fee    uint32 `json:"fee,omitempty"`
}
