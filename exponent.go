package chain10

import (
	"math/big"
)

var exponents map[string]func()*big.Int = map[string]func()*big.Int{}

func Exponent(chainid uint64) *big.Int {

	var key string = createkey(chainid)

	fn, found := exponents[key]
	if !found {
		return nil
	}

	return fn()
}
