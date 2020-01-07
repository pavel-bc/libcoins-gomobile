package main

//go:generate gopherjs build --minify -o libcoins.js
import (
    "github.com/pavel-bc/libcoins/libcoins"
    "github.com/gopherjs/gopherjs/js"
)

func main() {
    js.Module.Get("exports").Set("BitcoinSign", libcoins.BitcoinSign)
}
