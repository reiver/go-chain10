package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
const BlastSepoliaTestnet = 18

func init() {
	const value =                      BlastSepoliaTestnet
	var key string = createkey(chainid.BlastSepoliaTestnet)

	exponents[key] = value
}
