package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
func Zora() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           Zora
	var   key string = createkey(chainid.Zora)

	exponents[key] = value
}
