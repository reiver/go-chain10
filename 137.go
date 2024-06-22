package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 MATIC = 1,000,000,000,000,000,000 wei
const PolygonMainnet = 18

func init() {
	const value =                      PolygonMainnet
	var key string = createkey(chainid.PolygonMainnet)

	exponents[key] = value
}
