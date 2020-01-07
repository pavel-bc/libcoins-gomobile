package libcoins

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type BitcoinSignResult struct {
	Result   string
	Error    error
	HasError bool
}

func (res *BitcoinSignResult) Unwrap() string {
	if res.HasError {
		return res.Error.Error()
	}

	return res.Result
}

func BitcoinSign(key, message string) *BitcoinSignResult {
	pkBytes, err := hex.DecodeString(key)
	if err != nil {
		return &BitcoinSignResult{HasError: true, Error: err}
	}

	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)
	hash := chainhash.DoubleHashB([]byte(message))
	signature, err := privKey.Sign(hash)
	if err != nil {
		return &BitcoinSignResult{HasError: true, Error: err}
	}

	return &BitcoinSignResult{Result: hex.EncodeToString(signature.Serialize())}
}
