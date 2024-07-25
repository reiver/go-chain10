package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const ShapeSepolia = 18

func init() {
	const value =                      ShapeSepolia
	var key string = createkey(chainid.ShapeSepolia)

	exponents[key] = value
}
