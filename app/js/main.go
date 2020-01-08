package main

//go:generate gopherjs build --minify -o libcoins.js
import (
    "github.com/pavel-bc/libcoins/libcoins"
    "github.com/gopherjs/gopherjs/js"
)

func main() {
    js.Global.Set("BitcoinSign", libcoins.BitcoinSign)
    js.Global.Set("EthereumSign", libcoins.EthereumSign)
    js.Global.Set("AlgoIsValidAddress", libcoins.AlgoIsValidAddress)
}
