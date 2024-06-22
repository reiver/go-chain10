package chain10

import (
	"math/big"

	"github.com/reiver/go-chainid"
)

// 1 ETC = 1,000,000,000,000,000,000 wei
func EthereumClassic() *big.Int {
	return new(big.Int).SetUint64(18)
}

func init() {
	var value func()*big.Int =           EthereumClassic
	var   key string = createkey(chainid.EthereumClassic)

	exponents[key] = value
}
