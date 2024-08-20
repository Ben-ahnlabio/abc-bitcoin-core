package electrum

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/ahnlabio/bitcoin-core/electrum-api/interfaces"
	"github.com/checksum0/go-electrum/electrum"
)

type Electrum struct {
	Host string
	Port string
}

func NewElectrum(host, port string) Electrum {
	return Electrum{
		Host: host,
		Port: port,
	}
}

func (e Electrum) GetBalance(address string) (*interfaces.ElectrumBalance, error) {

	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	fmt.Println("GetBalance scriptHash: ", scriptHash)
	client, err := e.sslClient()
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	getBalanceResult, err := client.GetBalance(ctx, scriptHash)
	if err != nil {
		log.Printf("electrumx client GetBalance error: %s", err)
		return nil, err
	}

	log.Printf("getBalanceResult: %v scriptHash: %s", getBalanceResult, scriptHash)
	return &interfaces.ElectrumBalance{
		Confirmed:   getBalanceResult.Confirmed,
		Unconfirmed: getBalanceResult.Unconfirmed,
	}, nil
}

func (e Electrum) GetHistory(address string) ([]*interfaces.ElectrumHistory, error) {
	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	client, err := e.sslClient()
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	history, err := client.GetHistory(ctx, scriptHash)
	if err != nil {
		log.Printf("electrumx client GetHistory error: %s", err)
		return nil, err
	}

	log.Printf("history: %v scriptHash: %s", history, scriptHash)

	var histories []*interfaces.ElectrumHistory
	for _, h := range history {
		histories = append(histories, &interfaces.ElectrumHistory{
			Hash:   h.Hash,
			Height: h.Height,
			Fee:    h.Fee,
		})
	}
	return histories, nil
}

func (e Electrum) GetListUnspent(address string) ([]*interfaces.ElectrumUtxo, error) {
	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	client, err := e.sslClient()
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	listUnspent, err := client.ListUnspent(ctx, scriptHash)
	if err != nil {
		log.Printf("electrumx client GetListUnspent error: %s", err)
		return nil, err
	}

	log.Printf("[Electrum] listUnspent: %v scriptHash: %s", listUnspent, scriptHash)

	var utxos []*interfaces.ElectrumUtxo
	for _, u := range listUnspent {
		utxos = append(utxos, &interfaces.ElectrumUtxo{
			Height:   u.Height,
			Position: u.Position,
			Hash:     u.Hash,
			Value:    u.Value,
		})
	}

	return utxos, nil
}

func (e Electrum) GetTransaction(txHash string) (*interfaces.ElectrumTransaction, error) {
	client, err := e.sslClient()
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()

	getTransactionResult, err := client.GetTransaction(ctx, txHash)
	if err != nil {
		log.Printf("electrumx client GetTransaction error: %s", err)
		return nil, err
	}

	log.Printf("getTransactionResult: %v txHash: %s", getTransactionResult, txHash)
	return &interfaces.ElectrumTransaction{
		Blockhash:     getTransactionResult.Blockhash,
		Hash:          getTransactionResult.Hash,
		Confirmations: getTransactionResult.Confirmations,
	}, nil

}

func (e Electrum) sslClient() (*electrum.Client, error) {
	address := fmt.Sprintf("%s:%s", e.Host, e.Port)
	ctx := context.TODO()
	config := tls.Config{InsecureSkipVerify: true}
	client, err := electrum.NewClientSSL(ctx, address, &config)
	if err != nil {
		log.Printf("electrum.NewClientSSL error: %s", err)
		return nil, err
	}
	return client, nil
}

// func (e Electrum) tcpClient() (*electrum.Client, error) {
// 	address := fmt.Sprintf("%s:%s", e.Host, e.Port)
// 	ctx := context.TODO()
// 	client, err := electrum.NewClientTCP(ctx, address)
// 	if err != nil {
// 		log.Printf("electrum.NewClientTCP error: %s", err)
// 		return nil, err
// 	}
// 	return client, nil
// }
