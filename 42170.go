package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const ArbitrumNova = 18

func init() {
	const value =                      ArbitrumNova
	var key string = createkey(chainid.ArbitrumNova)

	exponents[key] = value
}
