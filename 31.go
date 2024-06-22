package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 RBTC = 100,000,000 rsat
func RootstockTestnet() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           RootstockTestnet
	var   key string = createkey(chainid.RootstockTestnet)

	exponents[key] = value
}
