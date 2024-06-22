package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETC = 1,000,000,000,000,000,000 wei
const EthereumClassic = 18

func init() {
	const value =                      EthereumClassic
	var key string = createkey(chainid.EthereumClassic)

	exponents[key] = value
}
