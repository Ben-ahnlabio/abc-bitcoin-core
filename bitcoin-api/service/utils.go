package service

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func validateAddress(address string) (btcjson.ValidateAddressChainResult, error) {
	result := btcjson.ValidateAddressChainResult{}
	net, err := GetNetworkFromAddress(address)
	if err != nil {
		return result, err
	}

	addr, err := btcutil.DecodeAddress(address, net)
	if err != nil {
		// Return the default value (false) for IsValid.
		return result, err
	}

	switch addr := addr.(type) {
	case *btcutil.AddressPubKeyHash:
		result.IsScript = btcjson.Bool(false)
		result.IsWitness = btcjson.Bool(false)

	case *btcutil.AddressScriptHash:
		result.IsScript = btcjson.Bool(true)
		result.IsWitness = btcjson.Bool(false)

	case *btcutil.AddressPubKey:
		result.IsScript = btcjson.Bool(false)
		result.IsWitness = btcjson.Bool(false)

	case *btcutil.AddressWitnessPubKeyHash:
		result.IsScript = btcjson.Bool(false)
		result.IsWitness = btcjson.Bool(true)
		result.WitnessVersion = btcjson.Int32(int32(addr.WitnessVersion()))
		result.WitnessProgram = btcjson.String(hex.EncodeToString(addr.WitnessProgram()))

	case *btcutil.AddressWitnessScriptHash:
		result.IsScript = btcjson.Bool(true)
		result.IsWitness = btcjson.Bool(true)
		result.WitnessVersion = btcjson.Int32(int32(addr.WitnessVersion()))
		result.WitnessProgram = btcjson.String(hex.EncodeToString(addr.WitnessProgram()))

	default:
		// Handle the case when a new Address is supported by btcutil, but none
		// of the cases were matched in the switch block. The current behaviour
		// is to do nothing, and only populate the Address and IsValid fields.
	}

	result.Address = addr.EncodeAddress()
	result.IsValid = true
	return result, nil

}

func GetNetworkFromAddress(address string) (*chaincfg.Params, error) {
	// address '1', 'bc1q', 'bc1p', '3' 으로 시작하면 메인넷
	// address 'm', 'tb1', 'n', '2' 으로 시작하면 테스트넷
	// 참고 : http://cryptostudy.xyz/crypto/article/241-%EB%B9%84%ED%8A%B8%EC%BD%94%EC%9D%B8-%EC%A3%BC%EC%86%8C#google_vignette

	if address[0] == '1' || address[0:3] == "bc1" || address[0] == '3' {
		return &chaincfg.MainNetParams, nil
	}
	if address[0] == 'm' || address[0:3] == "tb1" || address[0] == 'n' || address[0] == '2' {
		return &chaincfg.TestNet3Params, nil
	}
	// address is invalid bitcoin address
	return nil, fmt.Errorf("unknown network. to address is invalid bitcoin address. address: %s", address)
}
