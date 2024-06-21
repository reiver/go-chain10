package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
func ArbitrumOne() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           ArbitrumOne
	var   key string = createkey(chainid.ArbitrumOne)

	exponents[key] = value
}
