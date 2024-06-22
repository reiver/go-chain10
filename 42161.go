package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const ArbitrumOne = 18

func init() {
	const value =                      ArbitrumOne
	var key string = createkey(chainid.ArbitrumOne)

	exponents[key] = value
}
