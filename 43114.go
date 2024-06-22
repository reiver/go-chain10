package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 AVAX = 1,000,000,000 nanoAVAX = 10⁹ nanoAVAX
const AvalancheCChain = 9

func init() {
	const value =                      AvalancheCChain
	var key string = createkey(chainid.AvalancheCChain)

	exponents[key] = value
}
