package chain10

import (
	"strconv"
)

func createkey(chainid uint64) string {
	return strconv.FormatUint(chainid, 10)
}
