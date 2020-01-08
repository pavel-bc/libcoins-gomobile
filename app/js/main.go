package main

//go:generate gopherjs build --minify -o libcoins.js
import (
    "github.com/pavel-bc/libcoins/libcoins"
    "github.com/gopherjs/gopherjs/js"
)

func main() {
    js.Global.Set("BitcoinSign", BitcoinSign)
    js.Global.Set("EthereumSign", EthereumSign)
    js.Global.Set("AlgoIsValidAddress", AlgoIsValidAddress)
}

func BitcoinSign(key, message string) string {
    return libcoins.BitcoinSign(key, message).Unwrap()
}

func EthereumSign(key, message string) string {
    return libcoins.EthereumSign(key, message)
}

func AlgoIsValidAddress(addr string) bool {
    return libcoins.AlgoIsValidAddress(addr)
}
