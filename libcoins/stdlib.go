package libcoins

// #include <stdlib.h>
import "C"

type cCharPointer struct {
	Ptr *C.char
}

func Hello() string {
	return "hi"
}
