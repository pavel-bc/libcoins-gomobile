package main

import "C"
import "fmt"
import "github.com/pavel-bc/libcoins/libcoins"

//export BitcoinSign
func BitcoinSign(cKey, cMsg *C.char) *C.char {
	key := C.GoString(cKey)
	msg := C.GoString(cMsg)
	res := libcoins.BitcoinSign(key, msg).Unwrap()
	return C.CString(res)
}

//export Print
func Print(data string) {
	fmt.Println(data)
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
