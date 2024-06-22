package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
const Zora = 18

func init() {
	const value =                      Zora
	var key string = createkey(chainid.Zora)

	exponents[key] = value
}
