package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 RBTC = 100,000,000 rsat
func RootstockMainnet() *big.Int {
	return new(big.Int).SetUint64(8)
}

func init() {
	var value func()*big.Int =           RootstockMainnet
	var   key string = createkey(chainid.RootstockMainnet)

	exponents[key] = value
}
