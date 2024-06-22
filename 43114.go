package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 AVAX = 1,000,000,000 nanoAVAX
func AvalancheCChain() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           AvalancheCChain
	var   key string = createkey(chainid.AvalancheCChain)

	exponents[key] = value
}
