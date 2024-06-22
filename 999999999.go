package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
const ZoraSepoliaTestnet = 18

func init() {
	const value =                      ZoraSepoliaTestnet
	var key string = createkey(chainid.ZoraSepoliaTestnet)

	exponents[key] = value
}

