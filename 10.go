package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
const  OPMainnet = 18

func init() {
	const value =                      OPMainnet
	var key string = createkey(chainid.OPMainnet)

	exponents[key] = value
}
