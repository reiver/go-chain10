package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 MATIC = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const FantomOpera =  18

func init() {
	const value =                      FantomOpera
	var key string = createkey(chainid.FantomOpera)

	exponents[key] = value
}
