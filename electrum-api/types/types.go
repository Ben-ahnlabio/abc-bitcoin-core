package types

type IElectrum interface {
	GetBalance(address string) (*ElectrumBalance, error)
	GetTransaction(txId string) (*ElectrumTransaction, error)
	GetListUnspent(address string) ([]*ElectrumUtxo, error)
	GetHistory(address string) ([]*ElectrumHistory, error)
}

type ElectrumBalance struct {
	Confirmed   float64 `json:"confirmed"`
	Unconfirmed float64 `json:"unconfirmed"`
}

type ElectrumTransaction struct {
	Blockhash     string `json:"blockhash"`
	Hash          string `json:"hash"`
	Confirmations int32  `json:"confirmations"`
}

type ElectrumUtxo struct {
	Height   uint32 `json:"height"`
	Position uint32 `json:"tx_pos"`
	Hash     string `json:"tx_hash"`
	Value    uint64 `json:"value"`
}

type ElectrumHistory struct {
	Hash   string `json:"tx_hash"`
	Height int32  `json:"height"`
	Fee    uint32 `json:"fee,omitempty"`
}

type IBtcService interface {
	GetBalance(address string) (*GetBalanceResult, error)
	GetTransaction(txId string) (*GetTransactionResult, error)
	GetUTXO(address string) (*GetUTXOResult, error)
	GetHistory(address string) (*GetHistoryResult, error)
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
