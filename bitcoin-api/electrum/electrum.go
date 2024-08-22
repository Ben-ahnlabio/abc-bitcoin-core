package electrum

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/ahnlabio/bitcoin-core/bitcoin-api/types"
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

func (e Electrum) GetBalance(address string) (*types.ElectrumBalance, error) {

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
	defer client.Shutdown()

	ctx := context.TODO()
	getBalanceResult, err := client.GetBalance(ctx, scriptHash)
	if err != nil {
		log.Printf("electrumx client GetBalance error: %s", err)
		return nil, err
	}

	log.Printf("getBalanceResult: %v scriptHash: %s", getBalanceResult, scriptHash)
	return &types.ElectrumBalance{
		Confirmed:   getBalanceResult.Confirmed,
		Unconfirmed: getBalanceResult.Unconfirmed,
	}, nil
}

func (e Electrum) GetHistory(address string) ([]*types.ElectrumHistory, error) {
	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	client, err := e.sslClient()
	if err != nil {
		return nil, err
	}
	defer client.Shutdown()

	ctx := context.TODO()
	history, err := client.GetHistory(ctx, scriptHash)
	if err != nil {
		log.Printf("electrumx client GetHistory error: %s", err)
		return nil, err
	}

	log.Printf("history: %v scriptHash: %s", history, scriptHash)

	var histories []*types.ElectrumHistory
	for _, h := range history {
		histories = append(histories, &types.ElectrumHistory{
			Hash:   h.Hash,
			Height: h.Height,
			Fee:    h.Fee,
		})
	}
	return histories, nil
}

func (e Electrum) GetListUnspent(address string) ([]*types.ElectrumUtxo, error) {
	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	client, err := e.sslClient()
	if err != nil {
		return nil, err
	}
	defer client.Shutdown()

	ctx := context.TODO()
	listUnspent, err := client.ListUnspent(ctx, scriptHash)
	if err != nil {
		log.Printf("electrumx client GetListUnspent error: %s", err)
		return nil, err
	}

	log.Printf("[Electrum] listUnspent: %v scriptHash: %s", listUnspent, scriptHash)

	var utxos []*types.ElectrumUtxo
	for _, u := range listUnspent {
		utxos = append(utxos, &types.ElectrumUtxo{
			Height:   u.Height,
			Position: u.Position,
			Hash:     u.Hash,
			Value:    u.Value,
		})
	}

	return utxos, nil
}

func (e Electrum) GetTransaction(txHash string) (*types.ElectrumTransaction, error) {
	client, err := e.sslClient()
	if err != nil {
		return nil, err
	}
	defer client.Shutdown()

	ctx := context.TODO()

	getTransactionResult, err := client.GetTransaction(ctx, txHash)
	if err != nil {
		log.Printf("electrumx client GetTransaction error: %s", err)
		return nil, err
	}

	log.Printf("getTransactionResult: %v txHash: %s", getTransactionResult, txHash)
	return &types.ElectrumTransaction{
		Blockhash:     getTransactionResult.Blockhash,
		Hash:          getTransactionResult.Hash,
		Confirmations: getTransactionResult.Confirmations,
	}, nil

}

func (e Electrum) GetServerVersion() (string, error) {
	client, err := e.sslClient()
	if err != nil {
		return "", err
	}
	defer client.Shutdown()

	ctx := context.TODO()

	err = client.Ping(ctx)
	if err != nil {
		log.Printf("electrumx client Ping error: %s", err)
	}

	version, protocolVer, err := client.ServerVersion(ctx)
	if err != nil {
		log.Printf("electrumx client ServerVersion error: %s", err)
		return "", err
	}

	log.Printf("version: %s, protocol version: %s", version, protocolVer)
	return version, nil
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
