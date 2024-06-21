package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
func ArbitrumNova() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           ArbitrumNova
	var   key string = createkey(chainid.ArbitrumNova)

	exponents[key] = value
}
