package electrum

import (
	"testing"
)

func TestGetBalance(t *testing.T) {
	electrum := NewElectrum("localhost", "50002")
	address := "bc1qanfh6n9csne5swjer6wmd2djugcy5y6eqtws67"
	result, _ := electrum.GetBalance(address)
	expectedConfirmed := 2000.0
	expectedUnconfirmed := 0.0
	if result.Confirmed != expectedConfirmed {
		t.Fatalf("Expected %f, got %f", expectedConfirmed, result.Confirmed)
	}
	if result.Unconfirmed != expectedUnconfirmed {
		t.Fatalf("Expected %f, got %f", expectedUnconfirmed, result.Unconfirmed)
	}
}

func TestListunspent(t *testing.T) {
	electrum := NewElectrum("localhost", "50002")
	address := "bc1qanfh6n9csne5swjer6wmd2djugcy5y6eqtws67"
	result, _ := electrum.GetListUnspent(address)

	expectedLength := 1
	if len(result) != expectedLength {
		t.Fatalf("Expected %d, got %d", expectedLength, len(result))
	}

	for _, utxo := range result {
		t.Logf("UTXO: %v", utxo)
	}
}

func TestGetHistory(t *testing.T) {
	electrum := NewElectrum("localhost", "50002")
	address := "bc1qanfh6n9csne5swjer6wmd2djugcy5y6eqtws67"
	result, _ := electrum.GetHistory(address)

	expectedLength := 3
	if len(result) != expectedLength {
		t.Fatalf("Expected %d, got %d", expectedLength, len(result))
	}

	for _, history := range result {
		t.Logf("History: %v", history)
	}
}

func TestGetTransaction(t *testing.T) {
	electrum := NewElectrum("localhost", "50002")
	hashId := "894af5cb799532ca63f46decabe418a6df70a4942e678e572c309c578d5eaab7"

	result, _ := electrum.GetTransaction(hashId)
	t.Logf("Transaction: %v", result)

}
