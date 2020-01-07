package libcoins

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func EthereumSign(key, message string) string {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return ""
	}

	hash := sha256.Sum256([]byte(message))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return ""
	}

	return fmt.Sprintf("0x%x, 0x%x", r, s)
}
