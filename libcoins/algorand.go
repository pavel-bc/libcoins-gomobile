package libcoins

import (
    "github.com/algorand/go-algorand-sdk/types"
)

func AlgoIsValidAddress(addr string) bool {
	_, err := types.DecodeAddress(addr)
	if err != nil {
        return false
    }

    return true
}
