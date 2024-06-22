package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 AVAX = 1,000,000,000 nanoAVAX
const AvalancheCChain = 18

func init() {
	const value =                      AvalancheCChain
	var key string = createkey(chainid.AvalancheCChain)

	exponents[key] = value
}
