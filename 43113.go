package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 AVAX = 1,000,000,000 nanoAVAX
const AvalancheFujiTestnet = 9

func init() {
	const value =                      AvalancheFujiTestnet
	var key string = createkey(chainid.AvalancheFujiTestnet)

	exponents[key] = value
}
