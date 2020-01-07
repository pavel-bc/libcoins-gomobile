package main

import "C"
import "fmt"
import "github.com/pavel-bc/libcoins/libcoins"

//export BitcoinSign
func BitcoinSign(key, message string) string {
	return libcoins.BitcoinSign(key, message).Unwrap()
}

//export Print
func Print(data string) {
    fmt.Println(data)
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
