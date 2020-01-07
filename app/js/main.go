package main

//go:generate gopherjs build --minify -o libcoins.js
import (
    "github.com/pavel-bc/libcoins/libcoins"
    "github.com/gopherjs/gopherjs/js"
)

func main() {
    js.Global.Set("BitcoinSign", BitcoinSign)
}

func BitcoinSign(key, message string) string {
    return libcoins.BitcoinSign(key, message).Unwrap()
}
