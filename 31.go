package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 RBTC = 100,000,000 rsat
const RootstockTestnet = 8

func init() {
	const value =                      RootstockTestnet
	var key string = createkey(chainid.RootstockTestnet)

	exponents[key] = value
}
