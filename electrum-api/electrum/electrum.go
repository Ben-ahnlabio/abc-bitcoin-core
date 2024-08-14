package electrum

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/checksum0/go-electrum/electrum"
)

type Electrum struct {
	Host string
	Port string
}

func NewElectrum(host, port string) *Electrum {
	return &Electrum{
		Host: host,
		Port: port,
	}
}

func (e Electrum) GetBalance(address string) (*electrum.GetBalanceResult, error) {
	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	fmt.Println("GetBalance scriptHash: ", scriptHash)
	client, err := e.SSLClient()
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
	return &getBalanceResult, nil
}

func (e Electrum) GetHistory(address string) ([]*electrum.GetMempoolResult, error) {
	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	client, err := e.SSLClient()
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
	return history, nil
}

func (e Electrum) GetListUnspent(address string) ([]*electrum.ListUnspentResult, error) {
	scriptHash, err := electrum.AddressToElectrumScriptHash(address)
	if err != nil {
		log.Printf("electrum.AddressToElectrumScriptHash error: %s", err)
		return nil, err
	}

	client, err := e.SSLClient()
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
	return listUnspent, nil
}

func (e Electrum) GetTransaction(txHash string) (*electrum.GetTransactionResult, error) {
	client, err := e.SSLClient()
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
	return getTransactionResult, nil
}

func (e Electrum) SSLClient() (*electrum.Client, error) {
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

func (e Electrum) TCPClient() (*electrum.Client, error) {
	address := fmt.Sprintf("%s:%s", e.Host, e.Port)
	ctx := context.TODO()
	client, err := electrum.NewClientTCP(ctx, address)
	if err != nil {
		log.Printf("electrum.NewClientTCP error: %s", err)
		return nil, err
	}
	return client, nil
}
