package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 RBTC = 100,000,000 rsat
const RootstockMainnet = 8

func init() {
	const value =                      RootstockMainnet
	var key string = createkey(chainid.RootstockMainnet)

	exponents[key] = value
}
