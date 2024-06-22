package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 MATIC = 1,000,000,000,000,000,000 wei
func FantomTestnet() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           FantomTestnet
	var   key string = createkey(chainid.FantomTestnet)

	exponents[key] = value
}
