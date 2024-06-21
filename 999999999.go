package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
func ZoraSepoliaTestnet() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           ZoraSepoliaTestnet
	var   key string = createkey(chainid.ZoraSepoliaTestnet)

	exponents[key] = value
}

