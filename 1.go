package chain10

import (
	"github.com/reiver/go-chainid"
)

// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
const EthereumMainnet = 18

func init() {
	const value =                      EthereumMainnet
	var key string = createkey(chainid.EthereumMainnet)

	exponents[key] = value
}
