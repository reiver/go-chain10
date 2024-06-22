package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 MATIC = 1,000,000,000,000,000,000 wei
func PolygonMainnet() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           PolygonMainnet
	var   key string = createkey(chainid.PolygonMainnet)

	exponents[key] = value
}
