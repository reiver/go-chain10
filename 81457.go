package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const Blast = 18

func init() {
	const value =                      Blast
	var key string = createkey(chainid.Blast)

	exponents[key] = value
}
