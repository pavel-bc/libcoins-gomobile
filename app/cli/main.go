package main

import (
	"fmt"

	"github.com/pavel-bc/libcoins/libcoins"
)

func main() {
	message := "Hello, world!"

	btcecKey := "22a47fa09a223f2aa079edf85a7c2d4f8720ee63e502ee2869afab7de234b80c"
	btcResult := libcoins.BitcoinSign(btcecKey, message)
	fmt.Printf("Bitcoin signature: %s\n", btcResult.Unwrap())

	ecdsaKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	ecdsaSig := libcoins.EthereumSign(ecdsaKey, message)
	fmt.Printf("Ethereum signature: %s\n", ecdsaSig)
}
