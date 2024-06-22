package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei
const Holesky = 18

func init() {
	const value =                      Holesky
	var key string = createkey(chainid.Holesky)

	exponents[key] = value
}
