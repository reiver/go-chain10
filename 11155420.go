package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const OPSepoliaTestnet = 18

func init() {
	const value =                      OPSepoliaTestnet
	var key string = createkey(chainid.OPSepoliaTestnet)

	exponents[key] = value
}
