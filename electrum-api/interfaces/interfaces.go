package interfaces

type Electrum interface {
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
