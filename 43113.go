package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 AVAX = 1,000,000,000 nanoAVAX
func AvalancheFujiTestnet() *big.Int {
	return new(big.Int).SetUint64(9)
}

func init() {
	var value func()*big.Int =           AvalancheFujiTestnet
	var   key string = createkey(chainid.AvalancheFujiTestnet)

	exponents[key] = value
}
