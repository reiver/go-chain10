package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const Sepolia = 18

func init() {
	const value =                      Sepolia
	var key string = createkey(chainid.Sepolia)

	exponents[key] = value
}
