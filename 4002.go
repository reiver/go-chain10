package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 MATIC = 1,000,000,000,000,000,000 wei
const FantomTestnet = 18

func init() {
	const value =                      FantomTestnet
	var key string = createkey(chainid.FantomTestnet)

	exponents[key] = value
}
